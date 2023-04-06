// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/numaproj/numaflow-go/pkg/apis/proto/function/v1 (interfaces: UserDefinedFunctionClient,UserDefinedFunction_ReduceFnClient)

// Package funcmock is a generated GoMock package.
package funcmock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/numaproj/numaflow-go/pkg/apis/proto/function/v1"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockUserDefinedFunctionClient is a mock of UserDefinedFunctionClient interface.
type MockUserDefinedFunctionClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserDefinedFunctionClientMockRecorder
}

// MockUserDefinedFunctionClientMockRecorder is the mock recorder for MockUserDefinedFunctionClient.
type MockUserDefinedFunctionClientMockRecorder struct {
	mock *MockUserDefinedFunctionClient
}

// NewMockUserDefinedFunctionClient creates a new mock instance.
func NewMockUserDefinedFunctionClient(ctrl *gomock.Controller) *MockUserDefinedFunctionClient {
	mock := &MockUserDefinedFunctionClient{ctrl: ctrl}
	mock.recorder = &MockUserDefinedFunctionClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDefinedFunctionClient) EXPECT() *MockUserDefinedFunctionClientMockRecorder {
	return m.recorder
}

// IsReady mocks base method.
func (m *MockUserDefinedFunctionClient) IsReady(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*v1.ReadyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsReady", varargs...)
	ret0, _ := ret[0].(*v1.ReadyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsReady indicates an expected call of IsReady.
func (mr *MockUserDefinedFunctionClientMockRecorder) IsReady(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsReady", reflect.TypeOf((*MockUserDefinedFunctionClient)(nil).IsReady), varargs...)
}

// MapFn mocks base method.
func (m *MockUserDefinedFunctionClient) MapFn(arg0 context.Context, arg1 *v1.DatumRequest, arg2 ...grpc.CallOption) (*v1.DatumResponseList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MapFn", varargs...)
	ret0, _ := ret[0].(*v1.DatumResponseList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MapFn indicates an expected call of MapFn.
func (mr *MockUserDefinedFunctionClientMockRecorder) MapFn(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapFn", reflect.TypeOf((*MockUserDefinedFunctionClient)(nil).MapFn), varargs...)
}

// MapTFn mocks base method.
func (m *MockUserDefinedFunctionClient) MapTFn(arg0 context.Context, arg1 *v1.DatumRequest, arg2 ...grpc.CallOption) (*v1.DatumResponseList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MapTFn", varargs...)
	ret0, _ := ret[0].(*v1.DatumResponseList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MapTFn indicates an expected call of MapTFn.
func (mr *MockUserDefinedFunctionClientMockRecorder) MapTFn(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapTFn", reflect.TypeOf((*MockUserDefinedFunctionClient)(nil).MapTFn), varargs...)
}

// ReduceFn mocks base method.
func (m *MockUserDefinedFunctionClient) ReduceFn(arg0 context.Context, arg1 ...grpc.CallOption) (v1.UserDefinedFunction_ReduceFnClient, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReduceFn", varargs...)
	ret0, _ := ret[0].(v1.UserDefinedFunction_ReduceFnClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReduceFn indicates an expected call of ReduceFn.
func (mr *MockUserDefinedFunctionClientMockRecorder) ReduceFn(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReduceFn", reflect.TypeOf((*MockUserDefinedFunctionClient)(nil).ReduceFn), varargs...)
}

// MockUserDefinedFunction_ReduceFnClient is a mock of UserDefinedFunction_ReduceFnClient interface.
type MockUserDefinedFunction_ReduceFnClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserDefinedFunction_ReduceFnClientMockRecorder
}

// MockUserDefinedFunction_ReduceFnClientMockRecorder is the mock recorder for MockUserDefinedFunction_ReduceFnClient.
type MockUserDefinedFunction_ReduceFnClientMockRecorder struct {
	mock *MockUserDefinedFunction_ReduceFnClient
}

// NewMockUserDefinedFunction_ReduceFnClient creates a new mock instance.
func NewMockUserDefinedFunction_ReduceFnClient(ctrl *gomock.Controller) *MockUserDefinedFunction_ReduceFnClient {
	mock := &MockUserDefinedFunction_ReduceFnClient{ctrl: ctrl}
	mock.recorder = &MockUserDefinedFunction_ReduceFnClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDefinedFunction_ReduceFnClient) EXPECT() *MockUserDefinedFunction_ReduceFnClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method.
func (m *MockUserDefinedFunction_ReduceFnClient) CloseSend() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend.
func (mr *MockUserDefinedFunction_ReduceFnClientMockRecorder) CloseSend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockUserDefinedFunction_ReduceFnClient)(nil).CloseSend))
}

// Context mocks base method.
func (m *MockUserDefinedFunction_ReduceFnClient) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockUserDefinedFunction_ReduceFnClientMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockUserDefinedFunction_ReduceFnClient)(nil).Context))
}

// Header mocks base method.
func (m *MockUserDefinedFunction_ReduceFnClient) Header() (metadata.MD, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header.
func (mr *MockUserDefinedFunction_ReduceFnClientMockRecorder) Header() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockUserDefinedFunction_ReduceFnClient)(nil).Header))
}

// Recv mocks base method.
func (m *MockUserDefinedFunction_ReduceFnClient) Recv() (*v1.DatumResponseList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*v1.DatumResponseList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv.
func (mr *MockUserDefinedFunction_ReduceFnClientMockRecorder) Recv() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockUserDefinedFunction_ReduceFnClient)(nil).Recv))
}

// RecvMsg mocks base method.
func (m *MockUserDefinedFunction_ReduceFnClient) RecvMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg.
func (mr *MockUserDefinedFunction_ReduceFnClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockUserDefinedFunction_ReduceFnClient)(nil).RecvMsg), arg0)
}

// Send mocks base method.
func (m *MockUserDefinedFunction_ReduceFnClient) Send(arg0 *v1.DatumRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
func (mr *MockUserDefinedFunction_ReduceFnClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockUserDefinedFunction_ReduceFnClient)(nil).Send), arg0)
}

// SendMsg mocks base method.
func (m *MockUserDefinedFunction_ReduceFnClient) SendMsg(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg.
func (mr *MockUserDefinedFunction_ReduceFnClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockUserDefinedFunction_ReduceFnClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method.
func (m *MockUserDefinedFunction_ReduceFnClient) Trailer() metadata.MD {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer.
func (mr *MockUserDefinedFunction_ReduceFnClientMockRecorder) Trailer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockUserDefinedFunction_ReduceFnClient)(nil).Trailer))
}