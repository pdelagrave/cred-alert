// Code generated by counterfeiter. DO NOT EDIT.
package inflatorfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/cred-alert/inflator"
)

type FakeInflator struct {
	InflateStub        func(lager.Logger, string, string, string) error
	inflateMutex       sync.RWMutex
	inflateArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 string
	}
	inflateReturns struct {
		result1 error
	}
	inflateReturnsOnCall map[int]struct {
		result1 error
	}
	LogPathStub        func() string
	logPathMutex       sync.RWMutex
	logPathArgsForCall []struct{}
	logPathReturns     struct {
		result1 string
	}
	logPathReturnsOnCall map[int]struct {
		result1 string
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns     struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInflator) Inflate(arg1 lager.Logger, arg2 string, arg3 string, arg4 string) error {
	fake.inflateMutex.Lock()
	ret, specificReturn := fake.inflateReturnsOnCall[len(fake.inflateArgsForCall)]
	fake.inflateArgsForCall = append(fake.inflateArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Inflate", []interface{}{arg1, arg2, arg3, arg4})
	fake.inflateMutex.Unlock()
	if fake.InflateStub != nil {
		return fake.InflateStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.inflateReturns.result1
}

func (fake *FakeInflator) InflateCallCount() int {
	fake.inflateMutex.RLock()
	defer fake.inflateMutex.RUnlock()
	return len(fake.inflateArgsForCall)
}

func (fake *FakeInflator) InflateArgsForCall(i int) (lager.Logger, string, string, string) {
	fake.inflateMutex.RLock()
	defer fake.inflateMutex.RUnlock()
	return fake.inflateArgsForCall[i].arg1, fake.inflateArgsForCall[i].arg2, fake.inflateArgsForCall[i].arg3, fake.inflateArgsForCall[i].arg4
}

func (fake *FakeInflator) InflateReturns(result1 error) {
	fake.InflateStub = nil
	fake.inflateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInflator) InflateReturnsOnCall(i int, result1 error) {
	fake.InflateStub = nil
	if fake.inflateReturnsOnCall == nil {
		fake.inflateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.inflateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInflator) LogPath() string {
	fake.logPathMutex.Lock()
	ret, specificReturn := fake.logPathReturnsOnCall[len(fake.logPathArgsForCall)]
	fake.logPathArgsForCall = append(fake.logPathArgsForCall, struct{}{})
	fake.recordInvocation("LogPath", []interface{}{})
	fake.logPathMutex.Unlock()
	if fake.LogPathStub != nil {
		return fake.LogPathStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.logPathReturns.result1
}

func (fake *FakeInflator) LogPathCallCount() int {
	fake.logPathMutex.RLock()
	defer fake.logPathMutex.RUnlock()
	return len(fake.logPathArgsForCall)
}

func (fake *FakeInflator) LogPathReturns(result1 string) {
	fake.LogPathStub = nil
	fake.logPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInflator) LogPathReturnsOnCall(i int, result1 string) {
	fake.LogPathStub = nil
	if fake.logPathReturnsOnCall == nil {
		fake.logPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.logPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInflator) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.closeReturns.result1
}

func (fake *FakeInflator) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeInflator) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInflator) CloseReturnsOnCall(i int, result1 error) {
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInflator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.inflateMutex.RLock()
	defer fake.inflateMutex.RUnlock()
	fake.logPathMutex.RLock()
	defer fake.logPathMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInflator) recordInvocation(key string, args []interface{}) {
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

var _ inflator.Inflator = new(FakeInflator)
