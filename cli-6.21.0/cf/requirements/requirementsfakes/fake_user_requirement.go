// This file was generated by counterfeiter
package requirementsfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/requirements"
)

type FakeUserRequirement struct {
	ExecuteStub        func() error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct{}
	executeReturns     struct {
		result1 error
	}
	GetUserStub        func() models.UserFields
	getUserMutex       sync.RWMutex
	getUserArgsForCall []struct{}
	getUserReturns     struct {
		result1 models.UserFields
	}
}

func (fake *FakeUserRequirement) Execute() error {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct{}{})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub()
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *FakeUserRequirement) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeUserRequirement) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRequirement) GetUser() models.UserFields {
	fake.getUserMutex.Lock()
	fake.getUserArgsForCall = append(fake.getUserArgsForCall, struct{}{})
	fake.getUserMutex.Unlock()
	if fake.GetUserStub != nil {
		return fake.GetUserStub()
	} else {
		return fake.getUserReturns.result1
	}
}

func (fake *FakeUserRequirement) GetUserCallCount() int {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	return len(fake.getUserArgsForCall)
}

func (fake *FakeUserRequirement) GetUserReturns(result1 models.UserFields) {
	fake.GetUserStub = nil
	fake.getUserReturns = struct {
		result1 models.UserFields
	}{result1}
}

var _ requirements.UserRequirement = new(FakeUserRequirement)
