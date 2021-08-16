// Code generated by counterfeiter. DO NOT EDIT.
package assertfakes

import (
	"sync"

	"github.com/stretchr/testify/assert"
)

type FakeTestingT struct {
	ErrorfStub        func(string, ...interface{})
	errorfMutex       sync.RWMutex
	errorfArgsForCall []struct {
		arg1 string
		arg2 []interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTestingT) Errorf(arg1 string, arg2 ...interface{}) {
	fake.errorfMutex.Lock()
	fake.errorfArgsForCall = append(fake.errorfArgsForCall, struct {
		arg1 string
		arg2 []interface{}
	}{arg1, arg2})
	stub := fake.ErrorfStub
	fake.recordInvocation("Errorf", []interface{}{arg1, arg2})
	fake.errorfMutex.Unlock()
	if stub != nil {
		fake.ErrorfStub(arg1, arg2...)
	}
}

func (fake *FakeTestingT) ErrorfCallCount() int {
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	return len(fake.errorfArgsForCall)
}

func (fake *FakeTestingT) ErrorfCalls(stub func(string, ...interface{})) {
	fake.errorfMutex.Lock()
	defer fake.errorfMutex.Unlock()
	fake.ErrorfStub = stub
}

func (fake *FakeTestingT) ErrorfArgsForCall(i int) (string, []interface{}) {
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	argsForCall := fake.errorfArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTestingT) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.errorfMutex.RLock()
	defer fake.errorfMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTestingT) recordInvocation(key string, args []interface{}) {
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

var _ assert.TestingT = new(FakeTestingT)