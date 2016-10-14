// This file was generated by counterfeiter
package dbfakes

import (
	"cred-alert/db"
	"sync"

	"code.cloudfoundry.org/lager"
)

type FakeScanRepository struct {
	StartStub        func(logger lager.Logger, scanType, startSHA, stopSHA string, repository *db.Repository, fetch *db.Fetch) db.ActiveScan
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		logger     lager.Logger
		scanType   string
		startSHA   string
		stopSHA    string
		repository *db.Repository
		fetch      *db.Fetch
	}
	startReturns struct {
		result1 db.ActiveScan
	}
	ScansNotYetRunWithVersionStub        func(lager.Logger, int) ([]db.PriorScan, error)
	scansNotYetRunWithVersionMutex       sync.RWMutex
	scansNotYetRunWithVersionArgsForCall []struct {
		arg1 lager.Logger
		arg2 int
	}
	scansNotYetRunWithVersionReturns struct {
		result1 []db.PriorScan
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeScanRepository) Start(logger lager.Logger, scanType string, startSHA string, stopSHA string, repository *db.Repository, fetch *db.Fetch) db.ActiveScan {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		logger     lager.Logger
		scanType   string
		startSHA   string
		stopSHA    string
		repository *db.Repository
		fetch      *db.Fetch
	}{logger, scanType, startSHA, stopSHA, repository, fetch})
	fake.recordInvocation("Start", []interface{}{logger, scanType, startSHA, stopSHA, repository, fetch})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub(logger, scanType, startSHA, stopSHA, repository, fetch)
	} else {
		return fake.startReturns.result1
	}
}

func (fake *FakeScanRepository) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeScanRepository) StartArgsForCall(i int) (lager.Logger, string, string, string, *db.Repository, *db.Fetch) {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return fake.startArgsForCall[i].logger, fake.startArgsForCall[i].scanType, fake.startArgsForCall[i].startSHA, fake.startArgsForCall[i].stopSHA, fake.startArgsForCall[i].repository, fake.startArgsForCall[i].fetch
}

func (fake *FakeScanRepository) StartReturns(result1 db.ActiveScan) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 db.ActiveScan
	}{result1}
}

func (fake *FakeScanRepository) ScansNotYetRunWithVersion(arg1 lager.Logger, arg2 int) ([]db.PriorScan, error) {
	fake.scansNotYetRunWithVersionMutex.Lock()
	fake.scansNotYetRunWithVersionArgsForCall = append(fake.scansNotYetRunWithVersionArgsForCall, struct {
		arg1 lager.Logger
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("ScansNotYetRunWithVersion", []interface{}{arg1, arg2})
	fake.scansNotYetRunWithVersionMutex.Unlock()
	if fake.ScansNotYetRunWithVersionStub != nil {
		return fake.ScansNotYetRunWithVersionStub(arg1, arg2)
	} else {
		return fake.scansNotYetRunWithVersionReturns.result1, fake.scansNotYetRunWithVersionReturns.result2
	}
}

func (fake *FakeScanRepository) ScansNotYetRunWithVersionCallCount() int {
	fake.scansNotYetRunWithVersionMutex.RLock()
	defer fake.scansNotYetRunWithVersionMutex.RUnlock()
	return len(fake.scansNotYetRunWithVersionArgsForCall)
}

func (fake *FakeScanRepository) ScansNotYetRunWithVersionArgsForCall(i int) (lager.Logger, int) {
	fake.scansNotYetRunWithVersionMutex.RLock()
	defer fake.scansNotYetRunWithVersionMutex.RUnlock()
	return fake.scansNotYetRunWithVersionArgsForCall[i].arg1, fake.scansNotYetRunWithVersionArgsForCall[i].arg2
}

func (fake *FakeScanRepository) ScansNotYetRunWithVersionReturns(result1 []db.PriorScan, result2 error) {
	fake.ScansNotYetRunWithVersionStub = nil
	fake.scansNotYetRunWithVersionReturns = struct {
		result1 []db.PriorScan
		result2 error
	}{result1, result2}
}

func (fake *FakeScanRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.scansNotYetRunWithVersionMutex.RLock()
	defer fake.scansNotYetRunWithVersionMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeScanRepository) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ db.ScanRepository = new(FakeScanRepository)
