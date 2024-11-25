/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.32.4. DO NOT EDIT.

package model

import mock "github.com/stretchr/testify/mock"

// MockPlcQuery is an autogenerated mock type for the PlcQuery type
type MockPlcQuery struct {
	mock.Mock
}

type MockPlcQuery_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcQuery) EXPECT() *MockPlcQuery_Expecter {
	return &MockPlcQuery_Expecter{mock: &_m.Mock}
}

// GetQueryString provides a mock function with given fields:
func (_m *MockPlcQuery) GetQueryString() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcQuery_GetQueryString_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetQueryString'
type MockPlcQuery_GetQueryString_Call struct {
	*mock.Call
}

// GetQueryString is a helper method to define mock.On call
func (_e *MockPlcQuery_Expecter) GetQueryString() *MockPlcQuery_GetQueryString_Call {
	return &MockPlcQuery_GetQueryString_Call{Call: _e.mock.On("GetQueryString")}
}

func (_c *MockPlcQuery_GetQueryString_Call) Run(run func()) *MockPlcQuery_GetQueryString_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcQuery_GetQueryString_Call) Return(_a0 string) *MockPlcQuery_GetQueryString_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcQuery_GetQueryString_Call) RunAndReturn(run func() string) *MockPlcQuery_GetQueryString_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcQuery creates a new instance of MockPlcQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcQuery {
	mock := &MockPlcQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
