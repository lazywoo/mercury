// Code generated by MockGen. DO NOT EDIT.
// Source: ./interactive.go
//
// Generated by this command:
//
//	mockgen -source=./interactive.go -package=cachemocks -destination=mocks/interactive.mock.go InteractiveCache
//

// Package cachemocks is a generated GoMock package.
package cachemocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/lazywoo/mercury/internal/interactive/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockInteractiveCache is a mock of InteractiveCache interface.
type MockInteractiveCache struct {
	ctrl     *gomock.Controller
	recorder *MockInteractiveCacheMockRecorder
}

// MockInteractiveCacheMockRecorder is the mock recorder for MockInteractiveCache.
type MockInteractiveCacheMockRecorder struct {
	mock *MockInteractiveCache
}

// NewMockInteractiveCache creates a new mock instance.
func NewMockInteractiveCache(ctrl *gomock.Controller) *MockInteractiveCache {
	mock := &MockInteractiveCache{ctrl: ctrl}
	mock.recorder = &MockInteractiveCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInteractiveCache) EXPECT() *MockInteractiveCacheMockRecorder {
	return m.recorder
}

// BatchIncrReadCntIfPresent mocks base method.
func (m *MockInteractiveCache) BatchIncrReadCntIfPresent(ctx context.Context, biz string, bizIds []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchIncrReadCntIfPresent", ctx, biz, bizIds)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchIncrReadCntIfPresent indicates an expected call of BatchIncrReadCntIfPresent.
func (mr *MockInteractiveCacheMockRecorder) BatchIncrReadCntIfPresent(ctx, biz, bizIds any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchIncrReadCntIfPresent", reflect.TypeOf((*MockInteractiveCache)(nil).BatchIncrReadCntIfPresent), ctx, biz, bizIds)
}

// DecrFavoriteCntIfPresent mocks base method.
func (m *MockInteractiveCache) DecrFavoriteCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrFavoriteCntIfPresent", ctx, biz, bizId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecrFavoriteCntIfPresent indicates an expected call of DecrFavoriteCntIfPresent.
func (mr *MockInteractiveCacheMockRecorder) DecrFavoriteCntIfPresent(ctx, biz, bizId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrFavoriteCntIfPresent", reflect.TypeOf((*MockInteractiveCache)(nil).DecrFavoriteCntIfPresent), ctx, biz, bizId)
}

// DecrLikeCntIfPresent mocks base method.
func (m *MockInteractiveCache) DecrLikeCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecrLikeCntIfPresent", ctx, biz, bizId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecrLikeCntIfPresent indicates an expected call of DecrLikeCntIfPresent.
func (mr *MockInteractiveCacheMockRecorder) DecrLikeCntIfPresent(ctx, biz, bizId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrLikeCntIfPresent", reflect.TypeOf((*MockInteractiveCache)(nil).DecrLikeCntIfPresent), ctx, biz, bizId)
}

// Get mocks base method.
func (m *MockInteractiveCache) Get(ctx context.Context, biz string, bizId int64) (domain.Interactive, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, biz, bizId)
	ret0, _ := ret[0].(domain.Interactive)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInteractiveCacheMockRecorder) Get(ctx, biz, bizId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInteractiveCache)(nil).Get), ctx, biz, bizId)
}

// IncrFavoriteCntIfPresent mocks base method.
func (m *MockInteractiveCache) IncrFavoriteCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrFavoriteCntIfPresent", ctx, biz, bizId)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrFavoriteCntIfPresent indicates an expected call of IncrFavoriteCntIfPresent.
func (mr *MockInteractiveCacheMockRecorder) IncrFavoriteCntIfPresent(ctx, biz, bizId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrFavoriteCntIfPresent", reflect.TypeOf((*MockInteractiveCache)(nil).IncrFavoriteCntIfPresent), ctx, biz, bizId)
}

// IncrLikeCntIfPresent mocks base method.
func (m *MockInteractiveCache) IncrLikeCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrLikeCntIfPresent", ctx, biz, bizId)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrLikeCntIfPresent indicates an expected call of IncrLikeCntIfPresent.
func (mr *MockInteractiveCacheMockRecorder) IncrLikeCntIfPresent(ctx, biz, bizId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrLikeCntIfPresent", reflect.TypeOf((*MockInteractiveCache)(nil).IncrLikeCntIfPresent), ctx, biz, bizId)
}

// IncrReadCntIfPresent mocks base method.
func (m *MockInteractiveCache) IncrReadCntIfPresent(ctx context.Context, biz string, bizId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrReadCntIfPresent", ctx, biz, bizId)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrReadCntIfPresent indicates an expected call of IncrReadCntIfPresent.
func (mr *MockInteractiveCacheMockRecorder) IncrReadCntIfPresent(ctx, biz, bizId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrReadCntIfPresent", reflect.TypeOf((*MockInteractiveCache)(nil).IncrReadCntIfPresent), ctx, biz, bizId)
}

// Set mocks base method.
func (m *MockInteractiveCache) Set(ctx context.Context, biz string, bizId int64, intr domain.Interactive) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, biz, bizId, intr)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockInteractiveCacheMockRecorder) Set(ctx, biz, bizId, intr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockInteractiveCache)(nil).Set), ctx, biz, bizId, intr)
}