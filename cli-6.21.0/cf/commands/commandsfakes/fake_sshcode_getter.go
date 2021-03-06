// This file was generated by counterfeiter
package commandsfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/commandregistry"
	"github.com/cloudfoundry/cli/cf/commands"
	"github.com/cloudfoundry/cli/cf/flags"
	"github.com/cloudfoundry/cli/cf/requirements"
)

type FakeSSHCodeGetter struct {
	MetaDataStub        func() commandregistry.CommandMetadata
	metaDataMutex       sync.RWMutex
	metaDataArgsForCall []struct{}
	metaDataReturns     struct {
		result1 commandregistry.CommandMetadata
	}
	SetDependencyStub        func(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command
	setDependencyMutex       sync.RWMutex
	setDependencyArgsForCall []struct {
		deps       commandregistry.Dependency
		pluginCall bool
	}
	setDependencyReturns struct {
		result1 commandregistry.Command
	}
	RequirementsStub        func(requirementsFactory requirements.Factory, context flags.FlagContext) []requirements.Requirement
	requirementsMutex       sync.RWMutex
	requirementsArgsForCall []struct {
		requirementsFactory requirements.Factory
		context             flags.FlagContext
	}
	requirementsReturns struct {
		result1 []requirements.Requirement
	}
	ExecuteStub        func(context flags.FlagContext) error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		context flags.FlagContext
	}
	executeReturns struct {
		result1 error
	}
	GetStub        func() (string, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct{}
	getReturns     struct {
		result1 string
		result2 error
	}
}

func (fake *FakeSSHCodeGetter) MetaData() commandregistry.CommandMetadata {
	fake.metaDataMutex.Lock()
	fake.metaDataArgsForCall = append(fake.metaDataArgsForCall, struct{}{})
	fake.metaDataMutex.Unlock()
	if fake.MetaDataStub != nil {
		return fake.MetaDataStub()
	} else {
		return fake.metaDataReturns.result1
	}
}

func (fake *FakeSSHCodeGetter) MetaDataCallCount() int {
	fake.metaDataMutex.RLock()
	defer fake.metaDataMutex.RUnlock()
	return len(fake.metaDataArgsForCall)
}

func (fake *FakeSSHCodeGetter) MetaDataReturns(result1 commandregistry.CommandMetadata) {
	fake.MetaDataStub = nil
	fake.metaDataReturns = struct {
		result1 commandregistry.CommandMetadata
	}{result1}
}

func (fake *FakeSSHCodeGetter) SetDependency(deps commandregistry.Dependency, pluginCall bool) commandregistry.Command {
	fake.setDependencyMutex.Lock()
	fake.setDependencyArgsForCall = append(fake.setDependencyArgsForCall, struct {
		deps       commandregistry.Dependency
		pluginCall bool
	}{deps, pluginCall})
	fake.setDependencyMutex.Unlock()
	if fake.SetDependencyStub != nil {
		return fake.SetDependencyStub(deps, pluginCall)
	} else {
		return fake.setDependencyReturns.result1
	}
}

func (fake *FakeSSHCodeGetter) SetDependencyCallCount() int {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return len(fake.setDependencyArgsForCall)
}

func (fake *FakeSSHCodeGetter) SetDependencyArgsForCall(i int) (commandregistry.Dependency, bool) {
	fake.setDependencyMutex.RLock()
	defer fake.setDependencyMutex.RUnlock()
	return fake.setDependencyArgsForCall[i].deps, fake.setDependencyArgsForCall[i].pluginCall
}

func (fake *FakeSSHCodeGetter) SetDependencyReturns(result1 commandregistry.Command) {
	fake.SetDependencyStub = nil
	fake.setDependencyReturns = struct {
		result1 commandregistry.Command
	}{result1}
}

func (fake *FakeSSHCodeGetter) Requirements(requirementsFactory requirements.Factory, context flags.FlagContext) []requirements.Requirement {
	fake.requirementsMutex.Lock()
	fake.requirementsArgsForCall = append(fake.requirementsArgsForCall, struct {
		requirementsFactory requirements.Factory
		context             flags.FlagContext
	}{requirementsFactory, context})
	fake.requirementsMutex.Unlock()
	if fake.RequirementsStub != nil {
		return fake.RequirementsStub(requirementsFactory, context)
	} else {
		return fake.requirementsReturns.result1
	}
}

func (fake *FakeSSHCodeGetter) RequirementsCallCount() int {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return len(fake.requirementsArgsForCall)
}

func (fake *FakeSSHCodeGetter) RequirementsArgsForCall(i int) (requirements.Factory, flags.FlagContext) {
	fake.requirementsMutex.RLock()
	defer fake.requirementsMutex.RUnlock()
	return fake.requirementsArgsForCall[i].requirementsFactory, fake.requirementsArgsForCall[i].context
}

func (fake *FakeSSHCodeGetter) RequirementsReturns(result1 []requirements.Requirement) {
	fake.RequirementsStub = nil
	fake.requirementsReturns = struct {
		result1 []requirements.Requirement
	}{result1}
}

func (fake *FakeSSHCodeGetter) Execute(context flags.FlagContext) error {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		context flags.FlagContext
	}{context})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(context)
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *FakeSSHCodeGetter) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeSSHCodeGetter) ExecuteArgsForCall(i int) flags.FlagContext {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].context
}

func (fake *FakeSSHCodeGetter) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSSHCodeGetter) Get() (string, error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct{}{})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub()
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *FakeSSHCodeGetter) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeSSHCodeGetter) GetReturns(result1 string, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

var _ commands.SSHCodeGetter = new(FakeSSHCodeGetter)
