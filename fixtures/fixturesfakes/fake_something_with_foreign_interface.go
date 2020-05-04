// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/v6/fixtures"
)

type FakeSomethingWithForeignInterface struct {
	StuffStub        func(int) string
	stuffMutex       sync.RWMutex
	stuffArgsForCall []struct {
		arg1 int
	}
	stuffReturns struct {
		result1 string
	}
	stuffReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSomethingWithForeignInterface) Stuff(arg1 int) string {
	fake.stuffMutex.Lock()
	defer fake.stuffMutex.Unlock()
	ret, specificReturn := fake.stuffReturnsOnCall[len(fake.stuffArgsForCall)]
	fake.stuffArgsForCall = append(fake.stuffArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("Stuff", []interface{}{arg1})
	if fake.StuffStub != nil {
		return fake.StuffStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stuffReturns
	return fakeReturns.result1
}

func (fake *FakeSomethingWithForeignInterface) StuffCallCount() int {
	fake.stuffMutex.RLock()
	defer fake.stuffMutex.RUnlock()
	return len(fake.stuffArgsForCall)
}

func (fake *FakeSomethingWithForeignInterface) StuffCalls(stub func(int) string) {
	fake.stuffMutex.Lock()
	defer fake.stuffMutex.Unlock()
	fake.StuffStub = stub
}

func (fake *FakeSomethingWithForeignInterface) StuffArgsForCall(i int) int {
	fake.stuffMutex.RLock()
	defer fake.stuffMutex.RUnlock()
	argsForCall := fake.stuffArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSomethingWithForeignInterface) StuffReturns(result1 string) {
	fake.stuffMutex.Lock()
	defer fake.stuffMutex.Unlock()
	fake.StuffStub = nil
	fake.stuffReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSomethingWithForeignInterface) StuffReturnsOnCall(i int, result1 string) {
	fake.stuffMutex.Lock()
	defer fake.stuffMutex.Unlock()
	fake.StuffStub = nil
	if fake.stuffReturnsOnCall == nil {
		fake.stuffReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.stuffReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSomethingWithForeignInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.stuffMutex.RLock()
	defer fake.stuffMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSomethingWithForeignInterface) recordInvocation(key string, args []interface{}) {
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

var _ fixtures.SomethingWithForeignInterface = new(FakeSomethingWithForeignInterface)
