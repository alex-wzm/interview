// Code generated by MockGen. DO NOT EDIT.
// Source: interview_grpc.pb.go

// Package mock_interview is a generated GoMock package.
package mock_interview

import (
	context "context"
	interview "interview-client/internal/api/interview"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockInterviewServiceClient is a mock of InterviewServiceClient interface.
type MockInterviewServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockInterviewServiceClientMockRecorder
}

// MockInterviewServiceClientMockRecorder is the mock recorder for MockInterviewServiceClient.
type MockInterviewServiceClientMockRecorder struct {
	mock *MockInterviewServiceClient
}

// NewMockInterviewServiceClient creates a new mock instance.
func NewMockInterviewServiceClient(ctrl *gomock.Controller) *MockInterviewServiceClient {
	mock := &MockInterviewServiceClient{ctrl: ctrl}
	mock.recorder = &MockInterviewServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterviewServiceClient) EXPECT() *MockInterviewServiceClientMockRecorder {
	return m.recorder
}

// HelloWorld mocks base method.
func (m *MockInterviewServiceClient) HelloWorld(ctx context.Context, in *interview.HelloWorldRequest, opts ...grpc.CallOption) (*interview.HelloWorldResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HelloWorld", varargs...)
	ret0, _ := ret[0].(*interview.HelloWorldResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HelloWorld indicates an expected call of HelloWorld.
func (mr *MockInterviewServiceClientMockRecorder) HelloWorld(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloWorld", reflect.TypeOf((*MockInterviewServiceClient)(nil).HelloWorld), varargs...)
}

// MockInterviewServiceServer is a mock of InterviewServiceServer interface.
type MockInterviewServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockInterviewServiceServerMockRecorder
}

// MockInterviewServiceServerMockRecorder is the mock recorder for MockInterviewServiceServer.
type MockInterviewServiceServerMockRecorder struct {
	mock *MockInterviewServiceServer
}

// NewMockInterviewServiceServer creates a new mock instance.
func NewMockInterviewServiceServer(ctrl *gomock.Controller) *MockInterviewServiceServer {
	mock := &MockInterviewServiceServer{ctrl: ctrl}
	mock.recorder = &MockInterviewServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterviewServiceServer) EXPECT() *MockInterviewServiceServerMockRecorder {
	return m.recorder
}

// HelloWorld mocks base method.
func (m *MockInterviewServiceServer) HelloWorld(arg0 context.Context, arg1 *interview.HelloWorldRequest) (*interview.HelloWorldResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HelloWorld", arg0, arg1)
	ret0, _ := ret[0].(*interview.HelloWorldResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HelloWorld indicates an expected call of HelloWorld.
func (mr *MockInterviewServiceServerMockRecorder) HelloWorld(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HelloWorld", reflect.TypeOf((*MockInterviewServiceServer)(nil).HelloWorld), arg0, arg1)
}

// mustEmbedUnimplementedInterviewServiceServer mocks base method.
func (m *MockInterviewServiceServer) mustEmbedUnimplementedInterviewServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedInterviewServiceServer")
}

// mustEmbedUnimplementedInterviewServiceServer indicates an expected call of mustEmbedUnimplementedInterviewServiceServer.
func (mr *MockInterviewServiceServerMockRecorder) mustEmbedUnimplementedInterviewServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedInterviewServiceServer", reflect.TypeOf((*MockInterviewServiceServer)(nil).mustEmbedUnimplementedInterviewServiceServer))
}

// MockUnsafeInterviewServiceServer is a mock of UnsafeInterviewServiceServer interface.
type MockUnsafeInterviewServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeInterviewServiceServerMockRecorder
}

// MockUnsafeInterviewServiceServerMockRecorder is the mock recorder for MockUnsafeInterviewServiceServer.
type MockUnsafeInterviewServiceServerMockRecorder struct {
	mock *MockUnsafeInterviewServiceServer
}

// NewMockUnsafeInterviewServiceServer creates a new mock instance.
func NewMockUnsafeInterviewServiceServer(ctrl *gomock.Controller) *MockUnsafeInterviewServiceServer {
	mock := &MockUnsafeInterviewServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeInterviewServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeInterviewServiceServer) EXPECT() *MockUnsafeInterviewServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedInterviewServiceServer mocks base method.
func (m *MockUnsafeInterviewServiceServer) mustEmbedUnimplementedInterviewServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedInterviewServiceServer")
}

// mustEmbedUnimplementedInterviewServiceServer indicates an expected call of mustEmbedUnimplementedInterviewServiceServer.
func (mr *MockUnsafeInterviewServiceServerMockRecorder) mustEmbedUnimplementedInterviewServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedInterviewServiceServer", reflect.TypeOf((*MockUnsafeInterviewServiceServer)(nil).mustEmbedUnimplementedInterviewServiceServer))
}
