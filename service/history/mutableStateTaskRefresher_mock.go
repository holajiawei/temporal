// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: mutableStateTaskRefresher.go

// Package history is a generated GoMock package.
package history

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockmutableStateTaskRefresher is a mock of mutableStateTaskRefresher interface.
type MockmutableStateTaskRefresher struct {
	ctrl     *gomock.Controller
	recorder *MockmutableStateTaskRefresherMockRecorder
}

// MockmutableStateTaskRefresherMockRecorder is the mock recorder for MockmutableStateTaskRefresher.
type MockmutableStateTaskRefresherMockRecorder struct {
	mock *MockmutableStateTaskRefresher
}

// NewMockmutableStateTaskRefresher creates a new mock instance.
func NewMockmutableStateTaskRefresher(ctrl *gomock.Controller) *MockmutableStateTaskRefresher {
	mock := &MockmutableStateTaskRefresher{ctrl: ctrl}
	mock.recorder = &MockmutableStateTaskRefresherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmutableStateTaskRefresher) EXPECT() *MockmutableStateTaskRefresherMockRecorder {
	return m.recorder
}

// refreshTasks mocks base method.
func (m *MockmutableStateTaskRefresher) refreshTasks(now time.Time, mutableState mutableState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "refreshTasks", now, mutableState)
	ret0, _ := ret[0].(error)
	return ret0
}

// refreshTasks indicates an expected call of refreshTasks.
func (mr *MockmutableStateTaskRefresherMockRecorder) refreshTasks(now, mutableState interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "refreshTasks", reflect.TypeOf((*MockmutableStateTaskRefresher)(nil).refreshTasks), now, mutableState)
}
