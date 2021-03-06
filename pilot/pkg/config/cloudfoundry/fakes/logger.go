// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type Logger struct {
	WarnfStub        func(template string, args ...interface{})
	warnfMutex       sync.RWMutex
	warnfArgsForCall []struct {
		template string
		args     []interface{}
	}
	InfofStub        func(template string, args ...interface{})
	infofMutex       sync.RWMutex
	infofArgsForCall []struct {
		template string
		args     []interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Logger) Warnf(template string, args ...interface{}) {
	fake.warnfMutex.Lock()
	fake.warnfArgsForCall = append(fake.warnfArgsForCall, struct {
		template string
		args     []interface{}
	}{template, args})
	fake.recordInvocation("Warnf", []interface{}{template, args})
	fake.warnfMutex.Unlock()
	if fake.WarnfStub != nil {
		fake.WarnfStub(template, args...)
	}
}

func (fake *Logger) WarnfCallCount() int {
	fake.warnfMutex.RLock()
	defer fake.warnfMutex.RUnlock()
	return len(fake.warnfArgsForCall)
}

func (fake *Logger) WarnfArgsForCall(i int) (string, []interface{}) {
	fake.warnfMutex.RLock()
	defer fake.warnfMutex.RUnlock()
	return fake.warnfArgsForCall[i].template, fake.warnfArgsForCall[i].args
}

func (fake *Logger) Infof(template string, args ...interface{}) {
	fake.infofMutex.Lock()
	fake.infofArgsForCall = append(fake.infofArgsForCall, struct {
		template string
		args     []interface{}
	}{template, args})
	fake.recordInvocation("Infof", []interface{}{template, args})
	fake.infofMutex.Unlock()
	if fake.InfofStub != nil {
		fake.InfofStub(template, args...)
	}
}

func (fake *Logger) InfofCallCount() int {
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	return len(fake.infofArgsForCall)
}

func (fake *Logger) InfofArgsForCall(i int) (string, []interface{}) {
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	return fake.infofArgsForCall[i].template, fake.infofArgsForCall[i].args
}

func (fake *Logger) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.warnfMutex.RLock()
	defer fake.warnfMutex.RUnlock()
	fake.infofMutex.RLock()
	defer fake.infofMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Logger) recordInvocation(key string, args []interface{}) {
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
