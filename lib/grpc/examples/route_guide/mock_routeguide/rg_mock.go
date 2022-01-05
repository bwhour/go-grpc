// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/bwhour/go-grpc/lib/grpc/examples/route_guide/routeguide (interfaces: RouteGuideClient,RouteGuide_RouteChatClient)

package mock_routeguide

import (
	grpc "github.com/bwhour/go-grpc/lib/grpc"
	routeguide "github.com/bwhour/go-grpc/lib/grpc/examples/route_guide/routeguide"
	metadata "github.com/bwhour/go-grpc/lib/grpc/metadata"
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
)

// Mock of RouteGuideClient interface
type MockRouteGuideClient struct {
	ctrl     *gomock.Controller
	recorder *_MockRouteGuideClientRecorder
}

// Recorder for MockRouteGuideClient (not exported)
type _MockRouteGuideClientRecorder struct {
	mock *MockRouteGuideClient
}

func NewMockRouteGuideClient(ctrl *gomock.Controller) *MockRouteGuideClient {
	mock := &MockRouteGuideClient{ctrl: ctrl}
	mock.recorder = &_MockRouteGuideClientRecorder{mock}
	return mock
}

func (_m *MockRouteGuideClient) EXPECT() *_MockRouteGuideClientRecorder {
	return _m.recorder
}

func (_m *MockRouteGuideClient) GetFeature(_param0 context.Context, _param1 *routeguide.Point, _param2 ...grpc.CallOption) (*routeguide.Feature, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetFeature", _s...)
	ret0, _ := ret[0].(*routeguide.Feature)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouteGuideClientRecorder) GetFeature(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFeature", _s...)
}

func (_m *MockRouteGuideClient) ListFeatures(_param0 context.Context, _param1 *routeguide.Rectangle, _param2 ...grpc.CallOption) (routeguide.RouteGuide_ListFeaturesClient, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "ListFeatures", _s...)
	ret0, _ := ret[0].(routeguide.RouteGuide_ListFeaturesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouteGuideClientRecorder) ListFeatures(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListFeatures", _s...)
}

func (_m *MockRouteGuideClient) RecordRoute(_param0 context.Context, _param1 ...grpc.CallOption) (routeguide.RouteGuide_RecordRouteClient, error) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "RecordRoute", _s...)
	ret0, _ := ret[0].(routeguide.RouteGuide_RecordRouteClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouteGuideClientRecorder) RecordRoute(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RecordRoute", _s...)
}

func (_m *MockRouteGuideClient) RouteChat(_param0 context.Context, _param1 ...grpc.CallOption) (routeguide.RouteGuide_RouteChatClient, error) {
	_s := []interface{}{_param0}
	for _, _x := range _param1 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "RouteChat", _s...)
	ret0, _ := ret[0].(routeguide.RouteGuide_RouteChatClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouteGuideClientRecorder) RouteChat(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RouteChat", _s...)
}

// Mock of RouteGuide_RouteChatClient interface
type MockRouteGuide_RouteChatClient struct {
	ctrl     *gomock.Controller
	recorder *_MockRouteGuide_RouteChatClientRecorder
}

// Recorder for MockRouteGuide_RouteChatClient (not exported)
type _MockRouteGuide_RouteChatClientRecorder struct {
	mock *MockRouteGuide_RouteChatClient
}

func NewMockRouteGuide_RouteChatClient(ctrl *gomock.Controller) *MockRouteGuide_RouteChatClient {
	mock := &MockRouteGuide_RouteChatClient{ctrl: ctrl}
	mock.recorder = &_MockRouteGuide_RouteChatClientRecorder{mock}
	return mock
}

func (_m *MockRouteGuide_RouteChatClient) EXPECT() *_MockRouteGuide_RouteChatClientRecorder {
	return _m.recorder
}

func (_m *MockRouteGuide_RouteChatClient) CloseSend() error {
	ret := _m.ctrl.Call(_m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRouteGuide_RouteChatClientRecorder) CloseSend() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CloseSend")
}

func (_m *MockRouteGuide_RouteChatClient) Context() context.Context {
	ret := _m.ctrl.Call(_m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

func (_mr *_MockRouteGuide_RouteChatClientRecorder) Context() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Context")
}

func (_m *MockRouteGuide_RouteChatClient) Header() (metadata.MD, error) {
	ret := _m.ctrl.Call(_m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouteGuide_RouteChatClientRecorder) Header() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Header")
}

func (_m *MockRouteGuide_RouteChatClient) Recv() (*routeguide.RouteNote, error) {
	ret := _m.ctrl.Call(_m, "Recv")
	ret0, _ := ret[0].(*routeguide.RouteNote)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockRouteGuide_RouteChatClientRecorder) Recv() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Recv")
}

func (_m *MockRouteGuide_RouteChatClient) RecvMsg(_param0 interface{}) error {
	ret := _m.ctrl.Call(_m, "RecvMsg", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRouteGuide_RouteChatClientRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RecvMsg", arg0)
}

func (_m *MockRouteGuide_RouteChatClient) Send(_param0 *routeguide.RouteNote) error {
	ret := _m.ctrl.Call(_m, "Send", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRouteGuide_RouteChatClientRecorder) Send(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Send", arg0)
}

func (_m *MockRouteGuide_RouteChatClient) SendMsg(_param0 interface{}) error {
	ret := _m.ctrl.Call(_m, "SendMsg", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockRouteGuide_RouteChatClientRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SendMsg", arg0)
}

func (_m *MockRouteGuide_RouteChatClient) Trailer() metadata.MD {
	ret := _m.ctrl.Call(_m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

func (_mr *_MockRouteGuide_RouteChatClientRecorder) Trailer() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Trailer")
}
