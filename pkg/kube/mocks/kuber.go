// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/jenkins-x/jx/pkg/kube (interfaces: Kuber)

package kube_test

import (
	pegomock "github.com/petergtz/pegomock"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	api "k8s.io/client-go/tools/clientcmd/api"
	"reflect"
)

type MockKuber struct {
	fail func(message string, callerSkip ...int)
}

func NewMockKuber() *MockKuber {
	return &MockKuber{fail: pegomock.GlobalFailHandler}
}

func (mock *MockKuber) LoadConfig() (*api.Config, *clientcmd.PathOptions, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockKuber().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("LoadConfig", params, []reflect.Type{reflect.TypeOf((**api.Config)(nil)).Elem(), reflect.TypeOf((**clientcmd.PathOptions)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *api.Config
	var ret1 *clientcmd.PathOptions
	var ret2 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*api.Config)
		}
		if result[1] != nil {
			ret1 = result[1].(*clientcmd.PathOptions)
		}
		if result[2] != nil {
			ret2 = result[2].(error)
		}
	}
	return ret0, ret1, ret2
}

func (mock *MockKuber) UpdateConfig(_param0 string, _param1 string, _param2 string, _param3 string, _param4 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockKuber().")
	}
	params := []pegomock.Param{_param0, _param1, _param2, _param3, _param4}
	result := pegomock.GetGenericMockFrom(mock).Invoke("UpdateConfig", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockKuber) VerifyWasCalledOnce() *VerifierKuber {
	return &VerifierKuber{mock, pegomock.Times(1), nil}
}

func (mock *MockKuber) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierKuber {
	return &VerifierKuber{mock, invocationCountMatcher, nil}
}

func (mock *MockKuber) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierKuber {
	return &VerifierKuber{mock, invocationCountMatcher, inOrderContext}
}

type VerifierKuber struct {
	mock                   *MockKuber
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
}

func (verifier *VerifierKuber) LoadConfig() *Kuber_LoadConfig_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "LoadConfig", params)
	return &Kuber_LoadConfig_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Kuber_LoadConfig_OngoingVerification struct {
	mock              *MockKuber
	methodInvocations []pegomock.MethodInvocation
}

func (c *Kuber_LoadConfig_OngoingVerification) GetCapturedArguments() {
}

func (c *Kuber_LoadConfig_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierKuber) UpdateConfig(_param0 string, _param1 string, _param2 string, _param3 string, _param4 string) *Kuber_UpdateConfig_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2, _param3, _param4}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "UpdateConfig", params)
	return &Kuber_UpdateConfig_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type Kuber_UpdateConfig_OngoingVerification struct {
	mock              *MockKuber
	methodInvocations []pegomock.MethodInvocation
}

func (c *Kuber_UpdateConfig_OngoingVerification) GetCapturedArguments() (string, string, string, string, string) {
	_param0, _param1, _param2, _param3, _param4 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1], _param3[len(_param3)-1], _param4[len(_param4)-1]
}

func (c *Kuber_UpdateConfig_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []string, _param3 []string, _param4 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(params[0]))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(params[1]))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]string, len(params[2]))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
		_param3 = make([]string, len(params[3]))
		for u, param := range params[3] {
			_param3[u] = param.(string)
		}
		_param4 = make([]string, len(params[4]))
		for u, param := range params[4] {
			_param4[u] = param.(string)
		}
	}
	return
}
