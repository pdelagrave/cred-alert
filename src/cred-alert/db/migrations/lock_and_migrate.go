package migrations

import (
	"database/sql"
	"hash/crc32"
	"strings"
	"time"

	"github.com/BurntSushi/migration"
	"github.com/jinzhu/gorm"
	"github.com/pivotal-golang/lager"
)

func LockDBAndMigrate(logger lager.Logger, driver, dbURI string) (*gorm.DB, error) {
	logger = logger.Session("lock-db-and-migrate")
	logger.Info("starting")

	lockDB, err := dbOpen(logger, driver, dbURI)
	if err != nil {
		logger.Fatal("failed", err)
		return nil, err
	}
	defer lockDB.Close()

	lockName := crc32.ChecksumIEEE([]byte(driver + dbURI))

	for {
		logger.Info("acquiring-lock")
		_, err = lockDB.Exec(`SELECT GET_LOCK(?,10);`, lockName)
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}

		logger.Info("migrating")
		_, err = migration.OpenWith(driver, dbURI, Migrations, migration.DefaultGetVersion, setVersion)
		if err != nil {
			logger.Fatal("failed", err)
		}

		logger.Info("releasing-lock")
		_, err = lockDB.Exec(`SELECT RELEASE_LOCK(?)`, lockName)
		if err != nil {
			logger.Error("failed", err)
		}

		break
	}

	logger.Info("done")
	return gorm.Open(driver, dbURI)
}

func dbOpen(logger lager.Logger, driver, dbURI string) (*sql.DB, error) {
	var err error
	var lockDB *sql.DB

	logger = logger.Session("db-open")
	logger.Info("starting")

	for {
		lockDB, err = sql.Open(driver, dbURI)
		if err != nil {
			if strings.Contains(err.Error(), " dial ") {
				logger.Error("retrying", err)
				time.Sleep(5 * time.Second)
				continue
			}
			logger.Fatal("failed", err)
			return nil, err
		}

		break
	}

	logger.Info("done")
	return lockDB, err
}

func setVersion(tx migration.LimitedTx, version int) error {
	_, err := tx.Exec("UPDATE migration_version SET version = ?", version)
	return err
}