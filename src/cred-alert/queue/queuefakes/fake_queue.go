// This file was generated by counterfeiter
package queuefakes

import (
	"cred-alert/queue"
	"sync"
)

type FakeQueue struct {
	EnqueueStub        func(queue.Task) error
	enqueueMutex       sync.RWMutex
	enqueueArgsForCall []struct {
		arg1 queue.Task
	}
	enqueueReturns struct {
		result1 error
	}
	DequeueStub        func() (queue.AckTask, error)
	dequeueMutex       sync.RWMutex
	dequeueArgsForCall []struct{}
	dequeueReturns     struct {
		result1 queue.AckTask
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeQueue) Enqueue(arg1 queue.Task) error {
	fake.enqueueMutex.Lock()
	fake.enqueueArgsForCall = append(fake.enqueueArgsForCall, struct {
		arg1 queue.Task
	}{arg1})
	fake.recordInvocation("Enqueue", []interface{}{arg1})
	fake.enqueueMutex.Unlock()
	if fake.EnqueueStub != nil {
		return fake.EnqueueStub(arg1)
	} else {
		return fake.enqueueReturns.result1
	}
}

func (fake *FakeQueue) EnqueueCallCount() int {
	fake.enqueueMutex.RLock()
	defer fake.enqueueMutex.RUnlock()
	return len(fake.enqueueArgsForCall)
}

func (fake *FakeQueue) EnqueueArgsForCall(i int) queue.Task {
	fake.enqueueMutex.RLock()
	defer fake.enqueueMutex.RUnlock()
	return fake.enqueueArgsForCall[i].arg1
}

func (fake *FakeQueue) EnqueueReturns(result1 error) {
	fake.EnqueueStub = nil
	fake.enqueueReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeQueue) Dequeue() (queue.AckTask, error) {
	fake.dequeueMutex.Lock()
	fake.dequeueArgsForCall = append(fake.dequeueArgsForCall, struct{}{})
	fake.recordInvocation("Dequeue", []interface{}{})
	fake.dequeueMutex.Unlock()
	if fake.DequeueStub != nil {
		return fake.DequeueStub()
	} else {
		return fake.dequeueReturns.result1, fake.dequeueReturns.result2
	}
}

func (fake *FakeQueue) DequeueCallCount() int {
	fake.dequeueMutex.RLock()
	defer fake.dequeueMutex.RUnlock()
	return len(fake.dequeueArgsForCall)
}

func (fake *FakeQueue) DequeueReturns(result1 queue.AckTask, result2 error) {
	fake.DequeueStub = nil
	fake.dequeueReturns = struct {
		result1 queue.AckTask
		result2 error
	}{result1, result2}
}

func (fake *FakeQueue) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.enqueueMutex.RLock()
	defer fake.enqueueMutex.RUnlock()
	fake.dequeueMutex.RLock()
	defer fake.dequeueMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeQueue) recordInvocation(key string, args []interface{}) {
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

var _ queue.Queue = new(FakeQueue)
