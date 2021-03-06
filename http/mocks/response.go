// Code generated by mockery v1.0.0
package mocks

import errors "github.com/stone-payments/goclienttools/errors"

import mock "github.com/stretchr/testify/mock"
import nethttp "net/http"

// Response is an autogenerated mock type for the Response type
type Response struct {
	mock.Mock
}

// Bytes provides a mock function with given fields:
func (_m *Response) Bytes() []byte {
	ret := _m.Called()

	var r0 []byte
	if rf, ok := ret.Get(0).(func() []byte); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// ClearInternalBuffer provides a mock function with given fields:
func (_m *Response) ClearInternalBuffer() {
	_m.Called()
}

// Close provides a mock function with given fields:
func (_m *Response) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Error provides a mock function with given fields:
func (_m *Response) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Header provides a mock function with given fields:
func (_m *Response) Header() nethttp.Header {
	ret := _m.Called()

	var r0 nethttp.Header
	if rf, ok := ret.Get(0).(func() nethttp.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(nethttp.Header)
		}
	}

	return r0
}

// JSON provides a mock function with given fields: _a0
func (_m *Response) JSON(_a0 interface{}) errors.Error {
	ret := _m.Called(_a0)

	var r0 errors.Error
	if rf, ok := ret.Get(0).(func(interface{}) errors.Error); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(errors.Error)
		}
	}

	return r0
}

// Ok provides a mock function with given fields:
func (_m *Response) Ok() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// RawResponse provides a mock function with given fields:
func (_m *Response) RawResponse() *nethttp.Response {
	ret := _m.Called()

	var r0 *nethttp.Response
	if rf, ok := ret.Get(0).(func() *nethttp.Response); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*nethttp.Response)
		}
	}

	return r0
}

// Read provides a mock function with given fields: _a0
func (_m *Response) Read(_a0 []byte) (int, error) {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatusCode provides a mock function with given fields:
func (_m *Response) StatusCode() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
