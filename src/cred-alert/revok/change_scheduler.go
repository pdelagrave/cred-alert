package revok

import (
	"fmt"
	"hash/fnv"
	"os"

	"code.cloudfoundry.org/lager"

	"cred-alert/db"
)

//go:generate counterfeiter . Scheduler

type Scheduler interface {
	ScheduleWork(string, func())
}

type ChangeScheduler struct {
	logger    lager.Logger
	repoRepo  db.RepositoryRepository
	scheduler Scheduler
	fetcher   ChangeFetcher
}

func NewChangeScheduler(logger lager.Logger, repoRepo db.RepositoryRepository, scheduler Scheduler, fetcher ChangeFetcher) *ChangeScheduler {
	return &ChangeScheduler{
		logger:    logger,
		repoRepo:  repoRepo,
		scheduler: scheduler,
		fetcher:   fetcher,
	}
}

func (s *ChangeScheduler) Run(signals <-chan os.Signal, ready chan<- struct{}) error {
	close(ready)

	if err := s.ScheduleActiveRepos(s.logger); err != nil {
		return err
	}

	<-signals

	return nil
}

func (s *ChangeScheduler) ScheduleActiveRepos(logger lager.Logger) error {
	logger = logger.Session("schedule-repos")

	repos, err := s.repoRepo.Active()
	if err != nil {
		logger.Error("failed-to-fetch-active-repos", err)
		return err
	}

	for _, repo := range repos {
		repo := repo
		schedule := scheduleForRepo(repo)

		s.scheduler.ScheduleWork(schedule, func() {
			_ = s.fetcher.Fetch(logger, repo)
		})
	}

	logger.Info("finished-scheduling", lager.Data{
		"count": len(repos),
	})

	return nil
}

const buckets = 1440

func scheduleForRepo(repo db.Repository) string {
	h := fnv.New64a()
	h.Write([]byte(repo.Owner))
	h.Write([]byte("/"))
	h.Write([]byte(repo.Name))

	sum := h.Sum64()

	bucket := sum % buckets

	hour := bucket / 60
	minute := bucket % 60

	return fmt.Sprintf("0 %d %d * * *", minute, hour)
}
