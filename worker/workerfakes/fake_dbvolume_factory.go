// This file was generated by counterfeiter
package workerfakes

import (
	"sync"

	"github.com/concourse/atc/dbng"
	"github.com/concourse/atc/worker"
)

type FakeDBVolumeFactory struct {
	CreateContainerVolumeStub        func(*dbng.Team, *dbng.Worker, *dbng.CreatingContainer, string) (dbng.CreatingVolume, error)
	createContainerVolumeMutex       sync.RWMutex
	createContainerVolumeArgsForCall []struct {
		arg1 *dbng.Team
		arg2 *dbng.Worker
		arg3 *dbng.CreatingContainer
		arg4 string
	}
	createContainerVolumeReturns struct {
		result1 dbng.CreatingVolume
		result2 error
	}
	FindContainerVolumeStub        func(*dbng.Team, *dbng.Worker, *dbng.CreatingContainer, string) (dbng.CreatingVolume, dbng.CreatedVolume, error)
	findContainerVolumeMutex       sync.RWMutex
	findContainerVolumeArgsForCall []struct {
		arg1 *dbng.Team
		arg2 *dbng.Worker
		arg3 *dbng.CreatingContainer
		arg4 string
	}
	findContainerVolumeReturns struct {
		result1 dbng.CreatingVolume
		result2 dbng.CreatedVolume
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDBVolumeFactory) CreateContainerVolume(arg1 *dbng.Team, arg2 *dbng.Worker, arg3 *dbng.CreatingContainer, arg4 string) (dbng.CreatingVolume, error) {
	fake.createContainerVolumeMutex.Lock()
	fake.createContainerVolumeArgsForCall = append(fake.createContainerVolumeArgsForCall, struct {
		arg1 *dbng.Team
		arg2 *dbng.Worker
		arg3 *dbng.CreatingContainer
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("CreateContainerVolume", []interface{}{arg1, arg2, arg3, arg4})
	fake.createContainerVolumeMutex.Unlock()
	if fake.CreateContainerVolumeStub != nil {
		return fake.CreateContainerVolumeStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.createContainerVolumeReturns.result1, fake.createContainerVolumeReturns.result2
	}
}

func (fake *FakeDBVolumeFactory) CreateContainerVolumeCallCount() int {
	fake.createContainerVolumeMutex.RLock()
	defer fake.createContainerVolumeMutex.RUnlock()
	return len(fake.createContainerVolumeArgsForCall)
}

func (fake *FakeDBVolumeFactory) CreateContainerVolumeArgsForCall(i int) (*dbng.Team, *dbng.Worker, *dbng.CreatingContainer, string) {
	fake.createContainerVolumeMutex.RLock()
	defer fake.createContainerVolumeMutex.RUnlock()
	return fake.createContainerVolumeArgsForCall[i].arg1, fake.createContainerVolumeArgsForCall[i].arg2, fake.createContainerVolumeArgsForCall[i].arg3, fake.createContainerVolumeArgsForCall[i].arg4
}

func (fake *FakeDBVolumeFactory) CreateContainerVolumeReturns(result1 dbng.CreatingVolume, result2 error) {
	fake.CreateContainerVolumeStub = nil
	fake.createContainerVolumeReturns = struct {
		result1 dbng.CreatingVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeDBVolumeFactory) FindContainerVolume(arg1 *dbng.Team, arg2 *dbng.Worker, arg3 *dbng.CreatingContainer, arg4 string) (dbng.CreatingVolume, dbng.CreatedVolume, error) {
	fake.findContainerVolumeMutex.Lock()
	fake.findContainerVolumeArgsForCall = append(fake.findContainerVolumeArgsForCall, struct {
		arg1 *dbng.Team
		arg2 *dbng.Worker
		arg3 *dbng.CreatingContainer
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("FindContainerVolume", []interface{}{arg1, arg2, arg3, arg4})
	fake.findContainerVolumeMutex.Unlock()
	if fake.FindContainerVolumeStub != nil {
		return fake.FindContainerVolumeStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.findContainerVolumeReturns.result1, fake.findContainerVolumeReturns.result2, fake.findContainerVolumeReturns.result3
	}
}

func (fake *FakeDBVolumeFactory) FindContainerVolumeCallCount() int {
	fake.findContainerVolumeMutex.RLock()
	defer fake.findContainerVolumeMutex.RUnlock()
	return len(fake.findContainerVolumeArgsForCall)
}

func (fake *FakeDBVolumeFactory) FindContainerVolumeArgsForCall(i int) (*dbng.Team, *dbng.Worker, *dbng.CreatingContainer, string) {
	fake.findContainerVolumeMutex.RLock()
	defer fake.findContainerVolumeMutex.RUnlock()
	return fake.findContainerVolumeArgsForCall[i].arg1, fake.findContainerVolumeArgsForCall[i].arg2, fake.findContainerVolumeArgsForCall[i].arg3, fake.findContainerVolumeArgsForCall[i].arg4
}

func (fake *FakeDBVolumeFactory) FindContainerVolumeReturns(result1 dbng.CreatingVolume, result2 dbng.CreatedVolume, result3 error) {
	fake.FindContainerVolumeStub = nil
	fake.findContainerVolumeReturns = struct {
		result1 dbng.CreatingVolume
		result2 dbng.CreatedVolume
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeDBVolumeFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createContainerVolumeMutex.RLock()
	defer fake.createContainerVolumeMutex.RUnlock()
	fake.findContainerVolumeMutex.RLock()
	defer fake.findContainerVolumeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDBVolumeFactory) recordInvocation(key string, args []interface{}) {
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

var _ worker.DBVolumeFactory = new(FakeDBVolumeFactory)