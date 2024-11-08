// Code generated by MockGen. DO NOT EDIT.
// Source: ./article.go
//
// Generated by this command:
//
//	mockgen -source=./article.go -package=repomocks -destination=mocks/article.mock.go ArticleRepository
//

// Package repomocks is a generated GoMock package.
package repomocks

import (
	context "context"
	reflect "reflect"
	time "time"

	domain "github.com/lazywoo/mercury/internal/article/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockArticleRepository is a mock of ArticleRepository interface.
type MockArticleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockArticleRepositoryMockRecorder
}

// MockArticleRepositoryMockRecorder is the mock recorder for MockArticleRepository.
type MockArticleRepositoryMockRecorder struct {
	mock *MockArticleRepository
}

// NewMockArticleRepository creates a new mock instance.
func NewMockArticleRepository(ctrl *gomock.Controller) *MockArticleRepository {
	mock := &MockArticleRepository{ctrl: ctrl}
	mock.recorder = &MockArticleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleRepository) EXPECT() *MockArticleRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockArticleRepository) Create(ctx context.Context, atcl domain.Article) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, atcl)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockArticleRepositoryMockRecorder) Create(ctx, atcl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArticleRepository)(nil).Create), ctx, atcl)
}

// GetById mocks base method.
func (m *MockArticleRepository) GetById(ctx context.Context, id int64) (domain.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", ctx, id)
	ret0, _ := ret[0].(domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockArticleRepositoryMockRecorder) GetById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockArticleRepository)(nil).GetById), ctx, id)
}

// GetPublishedById mocks base method.
func (m *MockArticleRepository) GetPublishedById(ctx context.Context, id int64) (domain.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublishedById", ctx, id)
	ret0, _ := ret[0].(domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublishedById indicates an expected call of GetPublishedById.
func (mr *MockArticleRepositoryMockRecorder) GetPublishedById(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublishedById", reflect.TypeOf((*MockArticleRepository)(nil).GetPublishedById), ctx, id)
}

// List mocks base method.
func (m *MockArticleRepository) List(ctx context.Context, authorId int64, offset, limit int) ([]domain.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, authorId, offset, limit)
	ret0, _ := ret[0].([]domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockArticleRepositoryMockRecorder) List(ctx, authorId, offset, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockArticleRepository)(nil).List), ctx, authorId, offset, limit)
}

// ListPub mocks base method.
func (m *MockArticleRepository) ListPub(ctx context.Context, utime time.Time, offset, limit int) ([]domain.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPub", ctx, utime, offset, limit)
	ret0, _ := ret[0].([]domain.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPub indicates an expected call of ListPub.
func (mr *MockArticleRepositoryMockRecorder) ListPub(ctx, utime, offset, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPub", reflect.TypeOf((*MockArticleRepository)(nil).ListPub), ctx, utime, offset, limit)
}

// Sync mocks base method.
func (m *MockArticleRepository) Sync(ctx context.Context, atcl domain.Article) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", ctx, atcl)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sync indicates an expected call of Sync.
func (mr *MockArticleRepositoryMockRecorder) Sync(ctx, atcl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockArticleRepository)(nil).Sync), ctx, atcl)
}

// SyncStatus mocks base method.
func (m *MockArticleRepository) SyncStatus(ctx context.Context, id, authorId int64, status domain.ArticleStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncStatus", ctx, id, authorId, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncStatus indicates an expected call of SyncStatus.
func (mr *MockArticleRepositoryMockRecorder) SyncStatus(ctx, id, authorId, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncStatus", reflect.TypeOf((*MockArticleRepository)(nil).SyncStatus), ctx, id, authorId, status)
}

// Update mocks base method.
func (m *MockArticleRepository) Update(ctx context.Context, atcl domain.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, atcl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockArticleRepositoryMockRecorder) Update(ctx, atcl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArticleRepository)(nil).Update), ctx, atcl)
}
