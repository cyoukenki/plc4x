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

package bacnetip

import mock "github.com/stretchr/testify/mock"

// mock_ApplicationServiceElement is an autogenerated mock type for the _ApplicationServiceElement type
type mock_ApplicationServiceElement struct {
	mock.Mock
}

type mock_ApplicationServiceElement_Expecter struct {
	mock *mock.Mock
}

func (_m *mock_ApplicationServiceElement) EXPECT() *mock_ApplicationServiceElement_Expecter {
	return &mock_ApplicationServiceElement_Expecter{mock: &_m.Mock}
}

// Confirmation provides a mock function with given fields: pdu
func (_m *mock_ApplicationServiceElement) Confirmation(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_ApplicationServiceElement_Confirmation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Confirmation'
type mock_ApplicationServiceElement_Confirmation_Call struct {
	*mock.Call
}

// Confirmation is a helper method to define mock.On call
//   - pdu _PDU
func (_e *mock_ApplicationServiceElement_Expecter) Confirmation(pdu interface{}) *mock_ApplicationServiceElement_Confirmation_Call {
	return &mock_ApplicationServiceElement_Confirmation_Call{Call: _e.mock.On("Confirmation", pdu)}
}

func (_c *mock_ApplicationServiceElement_Confirmation_Call) Run(run func(pdu _PDU)) *mock_ApplicationServiceElement_Confirmation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *mock_ApplicationServiceElement_Confirmation_Call) Return(_a0 error) *mock_ApplicationServiceElement_Confirmation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_ApplicationServiceElement_Confirmation_Call) RunAndReturn(run func(_PDU) error) *mock_ApplicationServiceElement_Confirmation_Call {
	_c.Call.Return(run)
	return _c
}

// Indication provides a mock function with given fields: pdu
func (_m *mock_ApplicationServiceElement) Indication(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_ApplicationServiceElement_Indication_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Indication'
type mock_ApplicationServiceElement_Indication_Call struct {
	*mock.Call
}

// Indication is a helper method to define mock.On call
//   - pdu _PDU
func (_e *mock_ApplicationServiceElement_Expecter) Indication(pdu interface{}) *mock_ApplicationServiceElement_Indication_Call {
	return &mock_ApplicationServiceElement_Indication_Call{Call: _e.mock.On("Indication", pdu)}
}

func (_c *mock_ApplicationServiceElement_Indication_Call) Run(run func(pdu _PDU)) *mock_ApplicationServiceElement_Indication_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *mock_ApplicationServiceElement_Indication_Call) Return(_a0 error) *mock_ApplicationServiceElement_Indication_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_ApplicationServiceElement_Indication_Call) RunAndReturn(run func(_PDU) error) *mock_ApplicationServiceElement_Indication_Call {
	_c.Call.Return(run)
	return _c
}

// Request provides a mock function with given fields: pdu
func (_m *mock_ApplicationServiceElement) Request(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_ApplicationServiceElement_Request_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Request'
type mock_ApplicationServiceElement_Request_Call struct {
	*mock.Call
}

// Request is a helper method to define mock.On call
//   - pdu _PDU
func (_e *mock_ApplicationServiceElement_Expecter) Request(pdu interface{}) *mock_ApplicationServiceElement_Request_Call {
	return &mock_ApplicationServiceElement_Request_Call{Call: _e.mock.On("Request", pdu)}
}

func (_c *mock_ApplicationServiceElement_Request_Call) Run(run func(pdu _PDU)) *mock_ApplicationServiceElement_Request_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *mock_ApplicationServiceElement_Request_Call) Return(_a0 error) *mock_ApplicationServiceElement_Request_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_ApplicationServiceElement_Request_Call) RunAndReturn(run func(_PDU) error) *mock_ApplicationServiceElement_Request_Call {
	_c.Call.Return(run)
	return _c
}

// Response provides a mock function with given fields: pdu
func (_m *mock_ApplicationServiceElement) Response(pdu _PDU) error {
	ret := _m.Called(pdu)

	var r0 error
	if rf, ok := ret.Get(0).(func(_PDU) error); ok {
		r0 = rf(pdu)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// mock_ApplicationServiceElement_Response_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Response'
type mock_ApplicationServiceElement_Response_Call struct {
	*mock.Call
}

// Response is a helper method to define mock.On call
//   - pdu _PDU
func (_e *mock_ApplicationServiceElement_Expecter) Response(pdu interface{}) *mock_ApplicationServiceElement_Response_Call {
	return &mock_ApplicationServiceElement_Response_Call{Call: _e.mock.On("Response", pdu)}
}

func (_c *mock_ApplicationServiceElement_Response_Call) Run(run func(pdu _PDU)) *mock_ApplicationServiceElement_Response_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_PDU))
	})
	return _c
}

func (_c *mock_ApplicationServiceElement_Response_Call) Return(_a0 error) *mock_ApplicationServiceElement_Response_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *mock_ApplicationServiceElement_Response_Call) RunAndReturn(run func(_PDU) error) *mock_ApplicationServiceElement_Response_Call {
	_c.Call.Return(run)
	return _c
}

// _setElementService provides a mock function with given fields: elementService
func (_m *mock_ApplicationServiceElement) _setElementService(elementService _ServiceAccessPoint) {
	_m.Called(elementService)
}

// mock_ApplicationServiceElement__setElementService_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method '_setElementService'
type mock_ApplicationServiceElement__setElementService_Call struct {
	*mock.Call
}

// _setElementService is a helper method to define mock.On call
//   - elementService _ServiceAccessPoint
func (_e *mock_ApplicationServiceElement_Expecter) _setElementService(elementService interface{}) *mock_ApplicationServiceElement__setElementService_Call {
	return &mock_ApplicationServiceElement__setElementService_Call{Call: _e.mock.On("_setElementService", elementService)}
}

func (_c *mock_ApplicationServiceElement__setElementService_Call) Run(run func(elementService _ServiceAccessPoint)) *mock_ApplicationServiceElement__setElementService_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(_ServiceAccessPoint))
	})
	return _c
}

func (_c *mock_ApplicationServiceElement__setElementService_Call) Return() *mock_ApplicationServiceElement__setElementService_Call {
	_c.Call.Return()
	return _c
}

func (_c *mock_ApplicationServiceElement__setElementService_Call) RunAndReturn(run func(_ServiceAccessPoint)) *mock_ApplicationServiceElement__setElementService_Call {
	_c.Call.Return(run)
	return _c
}

// newMock_ApplicationServiceElement creates a new instance of mock_ApplicationServiceElement. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMock_ApplicationServiceElement(t interface {
	mock.TestingT
	Cleanup(func())
}) *mock_ApplicationServiceElement {
	mock := &mock_ApplicationServiceElement{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
