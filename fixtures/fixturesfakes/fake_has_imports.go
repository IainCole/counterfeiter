// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/maxbrunsfeld/counterfeiter/v6/fixtures"
)

type FakeHasImports struct {
	DoThingsStub        func(io.Writer, *os.File) *http.Client
	doThingsMutex       sync.RWMutex
	doThingsArgsForCall []struct {
		arg1 io.Writer
		arg2 *os.File
	}
	doThingsReturns struct {
		result1 *http.Client
	}
	doThingsReturnsOnCall map[int]struct {
		result1 *http.Client
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHasImports) DoThings(arg1 io.Writer, arg2 *os.File) *http.Client {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	ret, specificReturn := fake.doThingsReturnsOnCall[len(fake.doThingsArgsForCall)]
	fake.doThingsArgsForCall = append(fake.doThingsArgsForCall, struct {
		arg1 io.Writer
		arg2 *os.File
	}{arg1, arg2})
	fake.recordInvocation("DoThings", []interface{}{arg1, arg2})
	if fake.DoThingsStub != nil {
		return fake.DoThingsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.doThingsReturns
	return fakeReturns.result1
}

func (fake *FakeHasImports) DoThingsCallCount() int {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	return len(fake.doThingsArgsForCall)
}

func (fake *FakeHasImports) DoThingsCalls(stub func(io.Writer, *os.File) *http.Client) {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	fake.DoThingsStub = stub
}

func (fake *FakeHasImports) DoThingsArgsForCall(i int) (io.Writer, *os.File) {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	argsForCall := fake.doThingsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHasImports) DoThingsReturns(result1 *http.Client) {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	fake.DoThingsStub = nil
	fake.doThingsReturns = struct {
		result1 *http.Client
	}{result1}
}

func (fake *FakeHasImports) DoThingsReturnsOnCall(i int, result1 *http.Client) {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	fake.DoThingsStub = nil
	if fake.doThingsReturnsOnCall == nil {
		fake.doThingsReturnsOnCall = make(map[int]struct {
			result1 *http.Client
		})
	}
	fake.doThingsReturnsOnCall[i] = struct {
		result1 *http.Client
	}{result1}
}

func (fake *FakeHasImports) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHasImports) recordInvocation(key string, args []interface{}) {
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

var _ fixtures.HasImports = new(FakeHasImports)
