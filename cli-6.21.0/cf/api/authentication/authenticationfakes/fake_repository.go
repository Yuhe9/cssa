// This file was generated by counterfeiter
package authenticationfakes

import (
	"net/http"
	"sync"

	"github.com/cloudfoundry/cli/cf/api/authentication"
	"github.com/cloudfoundry/cli/cf/configuration/coreconfig"
)

type FakeRepository struct {
	DumpRequestStub        func(*http.Request)
	dumpRequestMutex       sync.RWMutex
	dumpRequestArgsForCall []struct {
		arg1 *http.Request
	}
	DumpResponseStub        func(*http.Response)
	dumpResponseMutex       sync.RWMutex
	dumpResponseArgsForCall []struct {
		arg1 *http.Response
	}
	RefreshAuthTokenStub        func() (updatedToken string, apiErr error)
	refreshAuthTokenMutex       sync.RWMutex
	refreshAuthTokenArgsForCall []struct{}
	refreshAuthTokenReturns     struct {
		result1 string
		result2 error
	}
	AuthenticateStub        func(credentials map[string]string) (apiErr error)
	authenticateMutex       sync.RWMutex
	authenticateArgsForCall []struct {
		credentials map[string]string
	}
	authenticateReturns struct {
		result1 error
	}
	AuthorizeStub        func(token string) (string, error)
	authorizeMutex       sync.RWMutex
	authorizeArgsForCall []struct {
		token string
	}
	authorizeReturns struct {
		result1 string
		result2 error
	}
	GetLoginPromptsAndSaveUAAServerURLStub        func() (map[string]coreconfig.AuthPrompt, error)
	getLoginPromptsAndSaveUAAServerURLMutex       sync.RWMutex
	getLoginPromptsAndSaveUAAServerURLArgsForCall []struct{}
	getLoginPromptsAndSaveUAAServerURLReturns     struct {
		result1 map[string]coreconfig.AuthPrompt
		result2 error
	}
}

func (fake *FakeRepository) DumpRequest(arg1 *http.Request) {
	fake.dumpRequestMutex.Lock()
	fake.dumpRequestArgsForCall = append(fake.dumpRequestArgsForCall, struct {
		arg1 *http.Request
	}{arg1})
	fake.dumpRequestMutex.Unlock()
	if fake.DumpRequestStub != nil {
		fake.DumpRequestStub(arg1)
	}
}

func (fake *FakeRepository) DumpRequestCallCount() int {
	fake.dumpRequestMutex.RLock()
	defer fake.dumpRequestMutex.RUnlock()
	return len(fake.dumpRequestArgsForCall)
}

func (fake *FakeRepository) DumpRequestArgsForCall(i int) *http.Request {
	fake.dumpRequestMutex.RLock()
	defer fake.dumpRequestMutex.RUnlock()
	return fake.dumpRequestArgsForCall[i].arg1
}

func (fake *FakeRepository) DumpResponse(arg1 *http.Response) {
	fake.dumpResponseMutex.Lock()
	fake.dumpResponseArgsForCall = append(fake.dumpResponseArgsForCall, struct {
		arg1 *http.Response
	}{arg1})
	fake.dumpResponseMutex.Unlock()
	if fake.DumpResponseStub != nil {
		fake.DumpResponseStub(arg1)
	}
}

func (fake *FakeRepository) DumpResponseCallCount() int {
	fake.dumpResponseMutex.RLock()
	defer fake.dumpResponseMutex.RUnlock()
	return len(fake.dumpResponseArgsForCall)
}

func (fake *FakeRepository) DumpResponseArgsForCall(i int) *http.Response {
	fake.dumpResponseMutex.RLock()
	defer fake.dumpResponseMutex.RUnlock()
	return fake.dumpResponseArgsForCall[i].arg1
}

func (fake *FakeRepository) RefreshAuthToken() (updatedToken string, apiErr error) {
	fake.refreshAuthTokenMutex.Lock()
	fake.refreshAuthTokenArgsForCall = append(fake.refreshAuthTokenArgsForCall, struct{}{})
	fake.refreshAuthTokenMutex.Unlock()
	if fake.RefreshAuthTokenStub != nil {
		return fake.RefreshAuthTokenStub()
	} else {
		return fake.refreshAuthTokenReturns.result1, fake.refreshAuthTokenReturns.result2
	}
}

func (fake *FakeRepository) RefreshAuthTokenCallCount() int {
	fake.refreshAuthTokenMutex.RLock()
	defer fake.refreshAuthTokenMutex.RUnlock()
	return len(fake.refreshAuthTokenArgsForCall)
}

func (fake *FakeRepository) RefreshAuthTokenReturns(result1 string, result2 error) {
	fake.RefreshAuthTokenStub = nil
	fake.refreshAuthTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Authenticate(credentials map[string]string) (apiErr error) {
	fake.authenticateMutex.Lock()
	fake.authenticateArgsForCall = append(fake.authenticateArgsForCall, struct {
		credentials map[string]string
	}{credentials})
	fake.authenticateMutex.Unlock()
	if fake.AuthenticateStub != nil {
		return fake.AuthenticateStub(credentials)
	} else {
		return fake.authenticateReturns.result1
	}
}

func (fake *FakeRepository) AuthenticateCallCount() int {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return len(fake.authenticateArgsForCall)
}

func (fake *FakeRepository) AuthenticateArgsForCall(i int) map[string]string {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return fake.authenticateArgsForCall[i].credentials
}

func (fake *FakeRepository) AuthenticateReturns(result1 error) {
	fake.AuthenticateStub = nil
	fake.authenticateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRepository) Authorize(token string) (string, error) {
	fake.authorizeMutex.Lock()
	fake.authorizeArgsForCall = append(fake.authorizeArgsForCall, struct {
		token string
	}{token})
	fake.authorizeMutex.Unlock()
	if fake.AuthorizeStub != nil {
		return fake.AuthorizeStub(token)
	} else {
		return fake.authorizeReturns.result1, fake.authorizeReturns.result2
	}
}

func (fake *FakeRepository) AuthorizeCallCount() int {
	fake.authorizeMutex.RLock()
	defer fake.authorizeMutex.RUnlock()
	return len(fake.authorizeArgsForCall)
}

func (fake *FakeRepository) AuthorizeArgsForCall(i int) string {
	fake.authorizeMutex.RLock()
	defer fake.authorizeMutex.RUnlock()
	return fake.authorizeArgsForCall[i].token
}

func (fake *FakeRepository) AuthorizeReturns(result1 string, result2 error) {
	fake.AuthorizeStub = nil
	fake.authorizeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) GetLoginPromptsAndSaveUAAServerURL() (map[string]coreconfig.AuthPrompt, error) {
	fake.getLoginPromptsAndSaveUAAServerURLMutex.Lock()
	fake.getLoginPromptsAndSaveUAAServerURLArgsForCall = append(fake.getLoginPromptsAndSaveUAAServerURLArgsForCall, struct{}{})
	fake.getLoginPromptsAndSaveUAAServerURLMutex.Unlock()
	if fake.GetLoginPromptsAndSaveUAAServerURLStub != nil {
		return fake.GetLoginPromptsAndSaveUAAServerURLStub()
	} else {
		return fake.getLoginPromptsAndSaveUAAServerURLReturns.result1, fake.getLoginPromptsAndSaveUAAServerURLReturns.result2
	}
}

func (fake *FakeRepository) GetLoginPromptsAndSaveUAAServerURLCallCount() int {
	fake.getLoginPromptsAndSaveUAAServerURLMutex.RLock()
	defer fake.getLoginPromptsAndSaveUAAServerURLMutex.RUnlock()
	return len(fake.getLoginPromptsAndSaveUAAServerURLArgsForCall)
}

func (fake *FakeRepository) GetLoginPromptsAndSaveUAAServerURLReturns(result1 map[string]coreconfig.AuthPrompt, result2 error) {
	fake.GetLoginPromptsAndSaveUAAServerURLStub = nil
	fake.getLoginPromptsAndSaveUAAServerURLReturns = struct {
		result1 map[string]coreconfig.AuthPrompt
		result2 error
	}{result1, result2}
}

var _ authentication.Repository = new(FakeRepository)
