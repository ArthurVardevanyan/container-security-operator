// Code generated by counterfeiter. DO NOT EDIT.
package secscanfakes

import (
	"sync"

	"github.com/quay/container-security-operator/image"
	"github.com/quay/container-security-operator/secscan"
)

type FakeInterface struct {
	GetLayerDataFromTemplateStub        func(string, *image.Image, bool, bool) (*secscan.Layer, error)
	getLayerDataFromTemplateMutex       sync.RWMutex
	getLayerDataFromTemplateArgsForCall []struct {
		arg1 string
		arg2 *image.Image
		arg3 bool
		arg4 bool
	}
	getLayerDataFromTemplateReturns struct {
		result1 *secscan.Layer
		result2 error
	}
	getLayerDataFromTemplateReturnsOnCall map[int]struct {
		result1 *secscan.Layer
		result2 error
	}
	WellknownStub        func(string, string) (secscan.WellknownInterface, error)
	wellknownMutex       sync.RWMutex
	wellknownArgsForCall []struct {
		arg1 string
		arg2 string
	}
	wellknownReturns struct {
		result1 secscan.WellknownInterface
		result2 error
	}
	wellknownReturnsOnCall map[int]struct {
		result1 secscan.WellknownInterface
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInterface) GetLayerDataFromTemplate(arg1 string, arg2 *image.Image, arg3 bool, arg4 bool) (*secscan.Layer, error) {
	fake.getLayerDataFromTemplateMutex.Lock()
	ret, specificReturn := fake.getLayerDataFromTemplateReturnsOnCall[len(fake.getLayerDataFromTemplateArgsForCall)]
	fake.getLayerDataFromTemplateArgsForCall = append(fake.getLayerDataFromTemplateArgsForCall, struct {
		arg1 string
		arg2 *image.Image
		arg3 bool
		arg4 bool
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetLayerDataFromTemplate", []interface{}{arg1, arg2, arg3, arg4})
	fake.getLayerDataFromTemplateMutex.Unlock()
	if fake.GetLayerDataFromTemplateStub != nil {
		return fake.GetLayerDataFromTemplateStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getLayerDataFromTemplateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInterface) GetLayerDataFromTemplateCallCount() int {
	fake.getLayerDataFromTemplateMutex.RLock()
	defer fake.getLayerDataFromTemplateMutex.RUnlock()
	return len(fake.getLayerDataFromTemplateArgsForCall)
}

func (fake *FakeInterface) GetLayerDataFromTemplateCalls(stub func(string, *image.Image, bool, bool) (*secscan.Layer, error)) {
	fake.getLayerDataFromTemplateMutex.Lock()
	defer fake.getLayerDataFromTemplateMutex.Unlock()
	fake.GetLayerDataFromTemplateStub = stub
}

func (fake *FakeInterface) GetLayerDataFromTemplateArgsForCall(i int) (string, *image.Image, bool, bool) {
	fake.getLayerDataFromTemplateMutex.RLock()
	defer fake.getLayerDataFromTemplateMutex.RUnlock()
	argsForCall := fake.getLayerDataFromTemplateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeInterface) GetLayerDataFromTemplateReturns(result1 *secscan.Layer, result2 error) {
	fake.getLayerDataFromTemplateMutex.Lock()
	defer fake.getLayerDataFromTemplateMutex.Unlock()
	fake.GetLayerDataFromTemplateStub = nil
	fake.getLayerDataFromTemplateReturns = struct {
		result1 *secscan.Layer
		result2 error
	}{result1, result2}
}

func (fake *FakeInterface) GetLayerDataFromTemplateReturnsOnCall(i int, result1 *secscan.Layer, result2 error) {
	fake.getLayerDataFromTemplateMutex.Lock()
	defer fake.getLayerDataFromTemplateMutex.Unlock()
	fake.GetLayerDataFromTemplateStub = nil
	if fake.getLayerDataFromTemplateReturnsOnCall == nil {
		fake.getLayerDataFromTemplateReturnsOnCall = make(map[int]struct {
			result1 *secscan.Layer
			result2 error
		})
	}
	fake.getLayerDataFromTemplateReturnsOnCall[i] = struct {
		result1 *secscan.Layer
		result2 error
	}{result1, result2}
}

func (fake *FakeInterface) Wellknown(arg1 string, arg2 string) (secscan.WellknownInterface, error) {
	fake.wellknownMutex.Lock()
	ret, specificReturn := fake.wellknownReturnsOnCall[len(fake.wellknownArgsForCall)]
	fake.wellknownArgsForCall = append(fake.wellknownArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Wellknown", []interface{}{arg1, arg2})
	fake.wellknownMutex.Unlock()
	if fake.WellknownStub != nil {
		return fake.WellknownStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.wellknownReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeInterface) WellknownCallCount() int {
	fake.wellknownMutex.RLock()
	defer fake.wellknownMutex.RUnlock()
	return len(fake.wellknownArgsForCall)
}

func (fake *FakeInterface) WellknownCalls(stub func(string, string) (secscan.WellknownInterface, error)) {
	fake.wellknownMutex.Lock()
	defer fake.wellknownMutex.Unlock()
	fake.WellknownStub = stub
}

func (fake *FakeInterface) WellknownArgsForCall(i int) (string, string) {
	fake.wellknownMutex.RLock()
	defer fake.wellknownMutex.RUnlock()
	argsForCall := fake.wellknownArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeInterface) WellknownReturns(result1 secscan.WellknownInterface, result2 error) {
	fake.wellknownMutex.Lock()
	defer fake.wellknownMutex.Unlock()
	fake.WellknownStub = nil
	fake.wellknownReturns = struct {
		result1 secscan.WellknownInterface
		result2 error
	}{result1, result2}
}

func (fake *FakeInterface) WellknownReturnsOnCall(i int, result1 secscan.WellknownInterface, result2 error) {
	fake.wellknownMutex.Lock()
	defer fake.wellknownMutex.Unlock()
	fake.WellknownStub = nil
	if fake.wellknownReturnsOnCall == nil {
		fake.wellknownReturnsOnCall = make(map[int]struct {
			result1 secscan.WellknownInterface
			result2 error
		})
	}
	fake.wellknownReturnsOnCall[i] = struct {
		result1 secscan.WellknownInterface
		result2 error
	}{result1, result2}
}

func (fake *FakeInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getLayerDataFromTemplateMutex.RLock()
	defer fake.getLayerDataFromTemplateMutex.RUnlock()
	fake.wellknownMutex.RLock()
	defer fake.wellknownMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInterface) recordInvocation(key string, args []interface{}) {
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

var _ secscan.Interface = new(FakeInterface)
