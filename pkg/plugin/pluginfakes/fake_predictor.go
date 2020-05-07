// Code generated by counterfeiter. DO NOT EDIT.
package pluginfakes

import (
	"sync"

	"github.com/kubernetes-analysis/kubernetes-analysis/pkg/plugin"
)

type FakePredictor struct {
	PredictStub        func(string, string) (float64, error)
	predictMutex       sync.RWMutex
	predictArgsForCall []struct {
		arg1 string
		arg2 string
	}
	predictReturns struct {
		result1 float64
		result2 error
	}
	predictReturnsOnCall map[int]struct {
		result1 float64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePredictor) Predict(arg1 string, arg2 string) (float64, error) {
	fake.predictMutex.Lock()
	ret, specificReturn := fake.predictReturnsOnCall[len(fake.predictArgsForCall)]
	fake.predictArgsForCall = append(fake.predictArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Predict", []interface{}{arg1, arg2})
	fake.predictMutex.Unlock()
	if fake.PredictStub != nil {
		return fake.PredictStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.predictReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePredictor) PredictCallCount() int {
	fake.predictMutex.RLock()
	defer fake.predictMutex.RUnlock()
	return len(fake.predictArgsForCall)
}

func (fake *FakePredictor) PredictCalls(stub func(string, string) (float64, error)) {
	fake.predictMutex.Lock()
	defer fake.predictMutex.Unlock()
	fake.PredictStub = stub
}

func (fake *FakePredictor) PredictArgsForCall(i int) (string, string) {
	fake.predictMutex.RLock()
	defer fake.predictMutex.RUnlock()
	argsForCall := fake.predictArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePredictor) PredictReturns(result1 float64, result2 error) {
	fake.predictMutex.Lock()
	defer fake.predictMutex.Unlock()
	fake.PredictStub = nil
	fake.predictReturns = struct {
		result1 float64
		result2 error
	}{result1, result2}
}

func (fake *FakePredictor) PredictReturnsOnCall(i int, result1 float64, result2 error) {
	fake.predictMutex.Lock()
	defer fake.predictMutex.Unlock()
	fake.PredictStub = nil
	if fake.predictReturnsOnCall == nil {
		fake.predictReturnsOnCall = make(map[int]struct {
			result1 float64
			result2 error
		})
	}
	fake.predictReturnsOnCall[i] = struct {
		result1 float64
		result2 error
	}{result1, result2}
}

func (fake *FakePredictor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.predictMutex.RLock()
	defer fake.predictMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePredictor) recordInvocation(key string, args []interface{}) {
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

var _ plugin.Predictor = new(FakePredictor)
