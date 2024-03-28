// Copyright 2023 Specter Ops, Inc.
//
// Licensed under the Apache License, Version 2.0
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/specterops/bloodhound/src/daemons/datapipe (interfaces: Tasker)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	model "github.com/specterops/bloodhound/src/model"
	gomock "go.uber.org/mock/gomock"
)

// MockTasker is a mock of Tasker interface.
type MockTasker struct {
	ctrl     *gomock.Controller
	recorder *MockTaskerMockRecorder
}

// MockTaskerMockRecorder is the mock recorder for MockTasker.
type MockTaskerMockRecorder struct {
	mock *MockTasker
}

// NewMockTasker creates a new mock instance.
func NewMockTasker(ctrl *gomock.Controller) *MockTasker {
	mock := &MockTasker{ctrl: ctrl}
	mock.recorder = &MockTaskerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTasker) EXPECT() *MockTaskerMockRecorder {
	return m.recorder
}

// GetStatus mocks base method.
func (m *MockTasker) GetStatus() model.DatapipeStatusWrapper {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus")
	ret0, _ := ret[0].(model.DatapipeStatusWrapper)
	return ret0
}

// GetStatus indicates an expected call of GetStatus.
func (mr *MockTaskerMockRecorder) GetStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockTasker)(nil).GetStatus))
}

// RequestAnalysis mocks base method.
func (m *MockTasker) RequestAnalysis() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RequestAnalysis")
}

// RequestAnalysis indicates an expected call of RequestAnalysis.
func (mr *MockTaskerMockRecorder) RequestAnalysis() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestAnalysis", reflect.TypeOf((*MockTasker)(nil).RequestAnalysis))
}

// RequestDeletion mocks base method.
func (m *MockTasker) RequestDeletion() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RequestDeletion")
}

// RequestDeletion indicates an expected call of RequestDeletion.
func (mr *MockTaskerMockRecorder) RequestDeletion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestDeletion", reflect.TypeOf((*MockTasker)(nil).RequestDeletion))
}
