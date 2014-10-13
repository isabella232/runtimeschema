// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
)

type FakeStagerBBS struct {
	WatchForCompletedTaskStub        func() (<-chan models.Task, chan<- bool, <-chan error)
	watchForCompletedTaskMutex       sync.RWMutex
	watchForCompletedTaskArgsForCall []struct{}
	watchForCompletedTaskReturns     struct {
		result1 <-chan models.Task
		result2 chan<- bool
		result3 <-chan error
	}
	DesireTaskStub        func(models.Task) error
	desireTaskMutex       sync.RWMutex
	desireTaskArgsForCall []struct {
		arg1 models.Task
	}
	desireTaskReturns struct {
		result1 error
	}
	ResolvingTaskStub        func(taskGuid string) error
	resolvingTaskMutex       sync.RWMutex
	resolvingTaskArgsForCall []struct {
		taskGuid string
	}
	resolvingTaskReturns struct {
		result1 error
	}
	ResolveTaskStub        func(taskGuid string) error
	resolveTaskMutex       sync.RWMutex
	resolveTaskArgsForCall []struct {
		taskGuid string
	}
	resolveTaskReturns struct {
		result1 error
	}
	FailedToResolveTaskStub        func(taskGuid string) error
	failedToResolveTaskMutex       sync.RWMutex
	failedToResolveTaskArgsForCall []struct {
		taskGuid string
	}
	failedToResolveTaskReturns struct {
		result1 error
	}
	GetAvailableFileServerStub        func() (string, error)
	getAvailableFileServerMutex       sync.RWMutex
	getAvailableFileServerArgsForCall []struct{}
	getAvailableFileServerReturns     struct {
		result1 string
		result2 error
	}
}

func (fake *FakeStagerBBS) WatchForCompletedTask() (<-chan models.Task, chan<- bool, <-chan error) {
	fake.watchForCompletedTaskMutex.Lock()
	fake.watchForCompletedTaskArgsForCall = append(fake.watchForCompletedTaskArgsForCall, struct{}{})
	fake.watchForCompletedTaskMutex.Unlock()
	if fake.WatchForCompletedTaskStub != nil {
		return fake.WatchForCompletedTaskStub()
	} else {
		return fake.watchForCompletedTaskReturns.result1, fake.watchForCompletedTaskReturns.result2, fake.watchForCompletedTaskReturns.result3
	}
}

func (fake *FakeStagerBBS) WatchForCompletedTaskCallCount() int {
	fake.watchForCompletedTaskMutex.RLock()
	defer fake.watchForCompletedTaskMutex.RUnlock()
	return len(fake.watchForCompletedTaskArgsForCall)
}

func (fake *FakeStagerBBS) WatchForCompletedTaskReturns(result1 <-chan models.Task, result2 chan<- bool, result3 <-chan error) {
	fake.WatchForCompletedTaskStub = nil
	fake.watchForCompletedTaskReturns = struct {
		result1 <-chan models.Task
		result2 chan<- bool
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeStagerBBS) DesireTask(arg1 models.Task) error {
	fake.desireTaskMutex.Lock()
	fake.desireTaskArgsForCall = append(fake.desireTaskArgsForCall, struct {
		arg1 models.Task
	}{arg1})
	fake.desireTaskMutex.Unlock()
	if fake.DesireTaskStub != nil {
		return fake.DesireTaskStub(arg1)
	} else {
		return fake.desireTaskReturns.result1
	}
}

func (fake *FakeStagerBBS) DesireTaskCallCount() int {
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	return len(fake.desireTaskArgsForCall)
}

func (fake *FakeStagerBBS) DesireTaskArgsForCall(i int) models.Task {
	fake.desireTaskMutex.RLock()
	defer fake.desireTaskMutex.RUnlock()
	return fake.desireTaskArgsForCall[i].arg1
}

func (fake *FakeStagerBBS) DesireTaskReturns(result1 error) {
	fake.DesireTaskStub = nil
	fake.desireTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStagerBBS) ResolvingTask(taskGuid string) error {
	fake.resolvingTaskMutex.Lock()
	fake.resolvingTaskArgsForCall = append(fake.resolvingTaskArgsForCall, struct {
		taskGuid string
	}{taskGuid})
	fake.resolvingTaskMutex.Unlock()
	if fake.ResolvingTaskStub != nil {
		return fake.ResolvingTaskStub(taskGuid)
	} else {
		return fake.resolvingTaskReturns.result1
	}
}

func (fake *FakeStagerBBS) ResolvingTaskCallCount() int {
	fake.resolvingTaskMutex.RLock()
	defer fake.resolvingTaskMutex.RUnlock()
	return len(fake.resolvingTaskArgsForCall)
}

func (fake *FakeStagerBBS) ResolvingTaskArgsForCall(i int) string {
	fake.resolvingTaskMutex.RLock()
	defer fake.resolvingTaskMutex.RUnlock()
	return fake.resolvingTaskArgsForCall[i].taskGuid
}

func (fake *FakeStagerBBS) ResolvingTaskReturns(result1 error) {
	fake.ResolvingTaskStub = nil
	fake.resolvingTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStagerBBS) ResolveTask(taskGuid string) error {
	fake.resolveTaskMutex.Lock()
	fake.resolveTaskArgsForCall = append(fake.resolveTaskArgsForCall, struct {
		taskGuid string
	}{taskGuid})
	fake.resolveTaskMutex.Unlock()
	if fake.ResolveTaskStub != nil {
		return fake.ResolveTaskStub(taskGuid)
	} else {
		return fake.resolveTaskReturns.result1
	}
}

func (fake *FakeStagerBBS) ResolveTaskCallCount() int {
	fake.resolveTaskMutex.RLock()
	defer fake.resolveTaskMutex.RUnlock()
	return len(fake.resolveTaskArgsForCall)
}

func (fake *FakeStagerBBS) ResolveTaskArgsForCall(i int) string {
	fake.resolveTaskMutex.RLock()
	defer fake.resolveTaskMutex.RUnlock()
	return fake.resolveTaskArgsForCall[i].taskGuid
}

func (fake *FakeStagerBBS) ResolveTaskReturns(result1 error) {
	fake.ResolveTaskStub = nil
	fake.resolveTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStagerBBS) FailedToResolveTask(taskGuid string) error {
	fake.failedToResolveTaskMutex.Lock()
	fake.failedToResolveTaskArgsForCall = append(fake.failedToResolveTaskArgsForCall, struct {
		taskGuid string
	}{taskGuid})
	fake.failedToResolveTaskMutex.Unlock()
	if fake.FailedToResolveTaskStub != nil {
		return fake.FailedToResolveTaskStub(taskGuid)
	} else {
		return fake.failedToResolveTaskReturns.result1
	}
}

func (fake *FakeStagerBBS) FailedToResolveTaskCallCount() int {
	fake.failedToResolveTaskMutex.RLock()
	defer fake.failedToResolveTaskMutex.RUnlock()
	return len(fake.failedToResolveTaskArgsForCall)
}

func (fake *FakeStagerBBS) FailedToResolveTaskArgsForCall(i int) string {
	fake.failedToResolveTaskMutex.RLock()
	defer fake.failedToResolveTaskMutex.RUnlock()
	return fake.failedToResolveTaskArgsForCall[i].taskGuid
}

func (fake *FakeStagerBBS) FailedToResolveTaskReturns(result1 error) {
	fake.FailedToResolveTaskStub = nil
	fake.failedToResolveTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStagerBBS) GetAvailableFileServer() (string, error) {
	fake.getAvailableFileServerMutex.Lock()
	fake.getAvailableFileServerArgsForCall = append(fake.getAvailableFileServerArgsForCall, struct{}{})
	fake.getAvailableFileServerMutex.Unlock()
	if fake.GetAvailableFileServerStub != nil {
		return fake.GetAvailableFileServerStub()
	} else {
		return fake.getAvailableFileServerReturns.result1, fake.getAvailableFileServerReturns.result2
	}
}

func (fake *FakeStagerBBS) GetAvailableFileServerCallCount() int {
	fake.getAvailableFileServerMutex.RLock()
	defer fake.getAvailableFileServerMutex.RUnlock()
	return len(fake.getAvailableFileServerArgsForCall)
}

func (fake *FakeStagerBBS) GetAvailableFileServerReturns(result1 string, result2 error) {
	fake.GetAvailableFileServerStub = nil
	fake.getAvailableFileServerReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

var _ bbs.StagerBBS = new(FakeStagerBBS)
