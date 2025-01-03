// Code generated by MockGen. DO NOT EDIT.
// Source: interactive.go
//
// Generated by this command:
//
//	mockgen -source=interactive.go -package=svcmocks -destination=mocks/interactive.mock.go InteractiveService
//

// Package svcmocks is a generated GoMock package.
package svcmocks

import (
	context "context"
	reflect "reflect"

	domain "github.com/lazywoo/mercury/internal/interactive/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockInteractiveService is a mock of InteractiveService interface.
type MockInteractiveService struct {
	ctrl     *gomock.Controller
	recorder *MockInteractiveServiceMockRecorder
}

// MockInteractiveServiceMockRecorder is the mock recorder for MockInteractiveService.
type MockInteractiveServiceMockRecorder struct {
	mock *MockInteractiveService
}

// NewMockInteractiveService creates a new mock instance.
func NewMockInteractiveService(ctrl *gomock.Controller) *MockInteractiveService {
	mock := &MockInteractiveService{ctrl: ctrl}
	mock.recorder = &MockInteractiveServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInteractiveService) EXPECT() *MockInteractiveServiceMockRecorder {
	return m.recorder
}

// CancelFavorite mocks base method.
func (m *MockInteractiveService) CancelFavorite(ctx context.Context, biz string, bizId, uid, fid int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelFavorite", ctx, biz, bizId, uid, fid)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelFavorite indicates an expected call of CancelFavorite.
func (mr *MockInteractiveServiceMockRecorder) CancelFavorite(ctx, biz, bizId, uid, fid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelFavorite", reflect.TypeOf((*MockInteractiveService)(nil).CancelFavorite), ctx, biz, bizId, uid, fid)
}

// CancelLike mocks base method.
func (m *MockInteractiveService) CancelLike(ctx context.Context, biz string, bizId, uid int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelLike", ctx, biz, bizId, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelLike indicates an expected call of CancelLike.
func (mr *MockInteractiveServiceMockRecorder) CancelLike(ctx, biz, bizId, uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelLike", reflect.TypeOf((*MockInteractiveService)(nil).CancelLike), ctx, biz, bizId, uid)
}

// Favorite mocks base method.
func (m *MockInteractiveService) Favorite(ctx context.Context, biz string, bizId, uid, fid int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Favorite", ctx, biz, bizId, uid, fid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Favorite indicates an expected call of Favorite.
func (mr *MockInteractiveServiceMockRecorder) Favorite(ctx, biz, bizId, uid, fid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Favorite", reflect.TypeOf((*MockInteractiveService)(nil).Favorite), ctx, biz, bizId, uid, fid)
}

// Get mocks base method.
func (m *MockInteractiveService) Get(ctx context.Context, biz string, bizId, uid int64) (domain.Interactive, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, biz, bizId, uid)
	ret0, _ := ret[0].(domain.Interactive)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInteractiveServiceMockRecorder) Get(ctx, biz, bizId, uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInteractiveService)(nil).Get), ctx, biz, bizId, uid)
}

// GetByIds mocks base method.
func (m *MockInteractiveService) GetByIds(ctx context.Context, biz string, bizIds []int64) (map[int64]domain.Interactive, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByIds", ctx, biz, bizIds)
	ret0, _ := ret[0].(map[int64]domain.Interactive)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIds indicates an expected call of GetByIds.
func (mr *MockInteractiveServiceMockRecorder) GetByIds(ctx, biz, bizIds any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByIds", reflect.TypeOf((*MockInteractiveService)(nil).GetByIds), ctx, biz, bizIds)
}

// IncrReadCnt mocks base method.
func (m *MockInteractiveService) IncrReadCnt(ctx context.Context, biz string, bizId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrReadCnt", ctx, biz, bizId)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrReadCnt indicates an expected call of IncrReadCnt.
func (mr *MockInteractiveServiceMockRecorder) IncrReadCnt(ctx, biz, bizId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrReadCnt", reflect.TypeOf((*MockInteractiveService)(nil).IncrReadCnt), ctx, biz, bizId)
}

// Like mocks base method.
func (m *MockInteractiveService) Like(ctx context.Context, biz string, bizId, uid int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Like", ctx, biz, bizId, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Like indicates an expected call of Like.
func (mr *MockInteractiveServiceMockRecorder) Like(ctx, biz, bizId, uid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Like", reflect.TypeOf((*MockInteractiveService)(nil).Like), ctx, biz, bizId, uid)
}
