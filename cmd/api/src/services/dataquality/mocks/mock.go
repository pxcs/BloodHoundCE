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
// Source: github.com/specterops/bloodhound/src/services/dataquality (interfaces: DataQualityData)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	model "github.com/specterops/bloodhound/src/model"
	gomock "go.uber.org/mock/gomock"
)

// MockDataQualityData is a mock of DataQualityData interface.
type MockDataQualityData struct {
	ctrl     *gomock.Controller
	recorder *MockDataQualityDataMockRecorder
}

// MockDataQualityDataMockRecorder is the mock recorder for MockDataQualityData.
type MockDataQualityDataMockRecorder struct {
	mock *MockDataQualityData
}

// NewMockDataQualityData creates a new mock instance.
func NewMockDataQualityData(ctrl *gomock.Controller) *MockDataQualityData {
	mock := &MockDataQualityData{ctrl: ctrl}
	mock.recorder = &MockDataQualityDataMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataQualityData) EXPECT() *MockDataQualityDataMockRecorder {
	return m.recorder
}

// CreateADDataQualityAggregation mocks base method.
func (m *MockDataQualityData) CreateADDataQualityAggregation(arg0 context.Context, arg1 model.ADDataQualityAggregation) (model.ADDataQualityAggregation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateADDataQualityAggregation", arg0, arg1)
	ret0, _ := ret[0].(model.ADDataQualityAggregation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateADDataQualityAggregation indicates an expected call of CreateADDataQualityAggregation.
func (mr *MockDataQualityDataMockRecorder) CreateADDataQualityAggregation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateADDataQualityAggregation", reflect.TypeOf((*MockDataQualityData)(nil).CreateADDataQualityAggregation), arg0, arg1)
}

// CreateADDataQualityStats mocks base method.
func (m *MockDataQualityData) CreateADDataQualityStats(arg0 context.Context, arg1 model.ADDataQualityStats) (model.ADDataQualityStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateADDataQualityStats", arg0, arg1)
	ret0, _ := ret[0].(model.ADDataQualityStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateADDataQualityStats indicates an expected call of CreateADDataQualityStats.
func (mr *MockDataQualityDataMockRecorder) CreateADDataQualityStats(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateADDataQualityStats", reflect.TypeOf((*MockDataQualityData)(nil).CreateADDataQualityStats), arg0, arg1)
}

// CreateAzureDataQualityAggregation mocks base method.
func (m *MockDataQualityData) CreateAzureDataQualityAggregation(arg0 model.AzureDataQualityAggregation) (model.AzureDataQualityAggregation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAzureDataQualityAggregation", arg0)
	ret0, _ := ret[0].(model.AzureDataQualityAggregation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAzureDataQualityAggregation indicates an expected call of CreateAzureDataQualityAggregation.
func (mr *MockDataQualityDataMockRecorder) CreateAzureDataQualityAggregation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAzureDataQualityAggregation", reflect.TypeOf((*MockDataQualityData)(nil).CreateAzureDataQualityAggregation), arg0)
}

// CreateAzureDataQualityStats mocks base method.
func (m *MockDataQualityData) CreateAzureDataQualityStats(arg0 model.AzureDataQualityStats) (model.AzureDataQualityStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAzureDataQualityStats", arg0)
	ret0, _ := ret[0].(model.AzureDataQualityStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAzureDataQualityStats indicates an expected call of CreateAzureDataQualityStats.
func (mr *MockDataQualityDataMockRecorder) CreateAzureDataQualityStats(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAzureDataQualityStats", reflect.TypeOf((*MockDataQualityData)(nil).CreateAzureDataQualityStats), arg0)
}
