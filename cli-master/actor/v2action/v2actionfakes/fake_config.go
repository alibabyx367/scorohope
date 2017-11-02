// This file was generated by counterfeiter
package v2actionfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/cli/actor/v2action"
)

type FakeConfig struct {
	PollingIntervalStub        func() time.Duration
	pollingIntervalMutex       sync.RWMutex
	pollingIntervalArgsForCall []struct{}
	pollingIntervalReturns     struct {
		result1 time.Duration
	}
	pollingIntervalReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	SetTargetInformationStub        func(api string, apiVersion string, auth string, minCLIVersion string, doppler string, uaa string, routing string, skipSSLValidation bool)
	setTargetInformationMutex       sync.RWMutex
	setTargetInformationArgsForCall []struct {
		api               string
		apiVersion        string
		auth              string
		minCLIVersion     string
		doppler           string
		uaa               string
		routing           string
		skipSSLValidation bool
	}
	SetTokenInformationStub        func(accessToken string, refreshToken string, sshOAuthClient string)
	setTokenInformationMutex       sync.RWMutex
	setTokenInformationArgsForCall []struct {
		accessToken    string
		refreshToken   string
		sshOAuthClient string
	}
	SkipSSLValidationStub        func() bool
	skipSSLValidationMutex       sync.RWMutex
	skipSSLValidationArgsForCall []struct{}
	skipSSLValidationReturns     struct {
		result1 bool
	}
	skipSSLValidationReturnsOnCall map[int]struct {
		result1 bool
	}
	StagingTimeoutStub        func() time.Duration
	stagingTimeoutMutex       sync.RWMutex
	stagingTimeoutArgsForCall []struct{}
	stagingTimeoutReturns     struct {
		result1 time.Duration
	}
	stagingTimeoutReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	StartupTimeoutStub        func() time.Duration
	startupTimeoutMutex       sync.RWMutex
	startupTimeoutArgsForCall []struct{}
	startupTimeoutReturns     struct {
		result1 time.Duration
	}
	startupTimeoutReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	TargetStub        func() string
	targetMutex       sync.RWMutex
	targetArgsForCall []struct{}
	targetReturns     struct {
		result1 string
	}
	targetReturnsOnCall map[int]struct {
		result1 string
	}
	UnsetOrganizationInformationStub        func()
	unsetOrganizationInformationMutex       sync.RWMutex
	unsetOrganizationInformationArgsForCall []struct{}
	UnsetSpaceInformationStub               func()
	unsetSpaceInformationMutex              sync.RWMutex
	unsetSpaceInformationArgsForCall        []struct{}
	invocations                             map[string][][]interface{}
	invocationsMutex                        sync.RWMutex
}

func (fake *FakeConfig) PollingInterval() time.Duration {
	fake.pollingIntervalMutex.Lock()
	ret, specificReturn := fake.pollingIntervalReturnsOnCall[len(fake.pollingIntervalArgsForCall)]
	fake.pollingIntervalArgsForCall = append(fake.pollingIntervalArgsForCall, struct{}{})
	fake.recordInvocation("PollingInterval", []interface{}{})
	fake.pollingIntervalMutex.Unlock()
	if fake.PollingIntervalStub != nil {
		return fake.PollingIntervalStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.pollingIntervalReturns.result1
}

func (fake *FakeConfig) PollingIntervalCallCount() int {
	fake.pollingIntervalMutex.RLock()
	defer fake.pollingIntervalMutex.RUnlock()
	return len(fake.pollingIntervalArgsForCall)
}

func (fake *FakeConfig) PollingIntervalReturns(result1 time.Duration) {
	fake.PollingIntervalStub = nil
	fake.pollingIntervalReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) PollingIntervalReturnsOnCall(i int, result1 time.Duration) {
	fake.PollingIntervalStub = nil
	if fake.pollingIntervalReturnsOnCall == nil {
		fake.pollingIntervalReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.pollingIntervalReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) SetTargetInformation(api string, apiVersion string, auth string, minCLIVersion string, doppler string, uaa string, routing string, skipSSLValidation bool) {
	fake.setTargetInformationMutex.Lock()
	fake.setTargetInformationArgsForCall = append(fake.setTargetInformationArgsForCall, struct {
		api               string
		apiVersion        string
		auth              string
		minCLIVersion     string
		doppler           string
		uaa               string
		routing           string
		skipSSLValidation bool
	}{api, apiVersion, auth, minCLIVersion, doppler, uaa, routing, skipSSLValidation})
	fake.recordInvocation("SetTargetInformation", []interface{}{api, apiVersion, auth, minCLIVersion, doppler, uaa, routing, skipSSLValidation})
	fake.setTargetInformationMutex.Unlock()
	if fake.SetTargetInformationStub != nil {
		fake.SetTargetInformationStub(api, apiVersion, auth, minCLIVersion, doppler, uaa, routing, skipSSLValidation)
	}
}

func (fake *FakeConfig) SetTargetInformationCallCount() int {
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	return len(fake.setTargetInformationArgsForCall)
}

func (fake *FakeConfig) SetTargetInformationArgsForCall(i int) (string, string, string, string, string, string, string, bool) {
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	return fake.setTargetInformationArgsForCall[i].api, fake.setTargetInformationArgsForCall[i].apiVersion, fake.setTargetInformationArgsForCall[i].auth, fake.setTargetInformationArgsForCall[i].minCLIVersion, fake.setTargetInformationArgsForCall[i].doppler, fake.setTargetInformationArgsForCall[i].uaa, fake.setTargetInformationArgsForCall[i].routing, fake.setTargetInformationArgsForCall[i].skipSSLValidation
}

func (fake *FakeConfig) SetTokenInformation(accessToken string, refreshToken string, sshOAuthClient string) {
	fake.setTokenInformationMutex.Lock()
	fake.setTokenInformationArgsForCall = append(fake.setTokenInformationArgsForCall, struct {
		accessToken    string
		refreshToken   string
		sshOAuthClient string
	}{accessToken, refreshToken, sshOAuthClient})
	fake.recordInvocation("SetTokenInformation", []interface{}{accessToken, refreshToken, sshOAuthClient})
	fake.setTokenInformationMutex.Unlock()
	if fake.SetTokenInformationStub != nil {
		fake.SetTokenInformationStub(accessToken, refreshToken, sshOAuthClient)
	}
}

func (fake *FakeConfig) SetTokenInformationCallCount() int {
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	return len(fake.setTokenInformationArgsForCall)
}

func (fake *FakeConfig) SetTokenInformationArgsForCall(i int) (string, string, string) {
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	return fake.setTokenInformationArgsForCall[i].accessToken, fake.setTokenInformationArgsForCall[i].refreshToken, fake.setTokenInformationArgsForCall[i].sshOAuthClient
}

func (fake *FakeConfig) SkipSSLValidation() bool {
	fake.skipSSLValidationMutex.Lock()
	ret, specificReturn := fake.skipSSLValidationReturnsOnCall[len(fake.skipSSLValidationArgsForCall)]
	fake.skipSSLValidationArgsForCall = append(fake.skipSSLValidationArgsForCall, struct{}{})
	fake.recordInvocation("SkipSSLValidation", []interface{}{})
	fake.skipSSLValidationMutex.Unlock()
	if fake.SkipSSLValidationStub != nil {
		return fake.SkipSSLValidationStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.skipSSLValidationReturns.result1
}

func (fake *FakeConfig) SkipSSLValidationCallCount() int {
	fake.skipSSLValidationMutex.RLock()
	defer fake.skipSSLValidationMutex.RUnlock()
	return len(fake.skipSSLValidationArgsForCall)
}

func (fake *FakeConfig) SkipSSLValidationReturns(result1 bool) {
	fake.SkipSSLValidationStub = nil
	fake.skipSSLValidationReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConfig) SkipSSLValidationReturnsOnCall(i int, result1 bool) {
	fake.SkipSSLValidationStub = nil
	if fake.skipSSLValidationReturnsOnCall == nil {
		fake.skipSSLValidationReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.skipSSLValidationReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeConfig) StagingTimeout() time.Duration {
	fake.stagingTimeoutMutex.Lock()
	ret, specificReturn := fake.stagingTimeoutReturnsOnCall[len(fake.stagingTimeoutArgsForCall)]
	fake.stagingTimeoutArgsForCall = append(fake.stagingTimeoutArgsForCall, struct{}{})
	fake.recordInvocation("StagingTimeout", []interface{}{})
	fake.stagingTimeoutMutex.Unlock()
	if fake.StagingTimeoutStub != nil {
		return fake.StagingTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stagingTimeoutReturns.result1
}

func (fake *FakeConfig) StagingTimeoutCallCount() int {
	fake.stagingTimeoutMutex.RLock()
	defer fake.stagingTimeoutMutex.RUnlock()
	return len(fake.stagingTimeoutArgsForCall)
}

func (fake *FakeConfig) StagingTimeoutReturns(result1 time.Duration) {
	fake.StagingTimeoutStub = nil
	fake.stagingTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) StagingTimeoutReturnsOnCall(i int, result1 time.Duration) {
	fake.StagingTimeoutStub = nil
	if fake.stagingTimeoutReturnsOnCall == nil {
		fake.stagingTimeoutReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.stagingTimeoutReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) StartupTimeout() time.Duration {
	fake.startupTimeoutMutex.Lock()
	ret, specificReturn := fake.startupTimeoutReturnsOnCall[len(fake.startupTimeoutArgsForCall)]
	fake.startupTimeoutArgsForCall = append(fake.startupTimeoutArgsForCall, struct{}{})
	fake.recordInvocation("StartupTimeout", []interface{}{})
	fake.startupTimeoutMutex.Unlock()
	if fake.StartupTimeoutStub != nil {
		return fake.StartupTimeoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.startupTimeoutReturns.result1
}

func (fake *FakeConfig) StartupTimeoutCallCount() int {
	fake.startupTimeoutMutex.RLock()
	defer fake.startupTimeoutMutex.RUnlock()
	return len(fake.startupTimeoutArgsForCall)
}

func (fake *FakeConfig) StartupTimeoutReturns(result1 time.Duration) {
	fake.StartupTimeoutStub = nil
	fake.startupTimeoutReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) StartupTimeoutReturnsOnCall(i int, result1 time.Duration) {
	fake.StartupTimeoutStub = nil
	if fake.startupTimeoutReturnsOnCall == nil {
		fake.startupTimeoutReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.startupTimeoutReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) Target() string {
	fake.targetMutex.Lock()
	ret, specificReturn := fake.targetReturnsOnCall[len(fake.targetArgsForCall)]
	fake.targetArgsForCall = append(fake.targetArgsForCall, struct{}{})
	fake.recordInvocation("Target", []interface{}{})
	fake.targetMutex.Unlock()
	if fake.TargetStub != nil {
		return fake.TargetStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.targetReturns.result1
}

func (fake *FakeConfig) TargetCallCount() int {
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	return len(fake.targetArgsForCall)
}

func (fake *FakeConfig) TargetReturns(result1 string) {
	fake.TargetStub = nil
	fake.targetReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) TargetReturnsOnCall(i int, result1 string) {
	fake.TargetStub = nil
	if fake.targetReturnsOnCall == nil {
		fake.targetReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.targetReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) UnsetOrganizationInformation() {
	fake.unsetOrganizationInformationMutex.Lock()
	fake.unsetOrganizationInformationArgsForCall = append(fake.unsetOrganizationInformationArgsForCall, struct{}{})
	fake.recordInvocation("UnsetOrganizationInformation", []interface{}{})
	fake.unsetOrganizationInformationMutex.Unlock()
	if fake.UnsetOrganizationInformationStub != nil {
		fake.UnsetOrganizationInformationStub()
	}
}

func (fake *FakeConfig) UnsetOrganizationInformationCallCount() int {
	fake.unsetOrganizationInformationMutex.RLock()
	defer fake.unsetOrganizationInformationMutex.RUnlock()
	return len(fake.unsetOrganizationInformationArgsForCall)
}

func (fake *FakeConfig) UnsetSpaceInformation() {
	fake.unsetSpaceInformationMutex.Lock()
	fake.unsetSpaceInformationArgsForCall = append(fake.unsetSpaceInformationArgsForCall, struct{}{})
	fake.recordInvocation("UnsetSpaceInformation", []interface{}{})
	fake.unsetSpaceInformationMutex.Unlock()
	if fake.UnsetSpaceInformationStub != nil {
		fake.UnsetSpaceInformationStub()
	}
}

func (fake *FakeConfig) UnsetSpaceInformationCallCount() int {
	fake.unsetSpaceInformationMutex.RLock()
	defer fake.unsetSpaceInformationMutex.RUnlock()
	return len(fake.unsetSpaceInformationArgsForCall)
}

func (fake *FakeConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.pollingIntervalMutex.RLock()
	defer fake.pollingIntervalMutex.RUnlock()
	fake.setTargetInformationMutex.RLock()
	defer fake.setTargetInformationMutex.RUnlock()
	fake.setTokenInformationMutex.RLock()
	defer fake.setTokenInformationMutex.RUnlock()
	fake.skipSSLValidationMutex.RLock()
	defer fake.skipSSLValidationMutex.RUnlock()
	fake.stagingTimeoutMutex.RLock()
	defer fake.stagingTimeoutMutex.RUnlock()
	fake.startupTimeoutMutex.RLock()
	defer fake.startupTimeoutMutex.RUnlock()
	fake.targetMutex.RLock()
	defer fake.targetMutex.RUnlock()
	fake.unsetOrganizationInformationMutex.RLock()
	defer fake.unsetOrganizationInformationMutex.RUnlock()
	fake.unsetSpaceInformationMutex.RLock()
	defer fake.unsetSpaceInformationMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeConfig) recordInvocation(key string, args []interface{}) {
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

var _ v2action.Config = new(FakeConfig)
