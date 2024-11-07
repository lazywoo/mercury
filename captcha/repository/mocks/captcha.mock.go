// Code generated by MockGen. DO NOT EDIT.
// Source: ./captcha.go
//
// Generated by this command:
//
//	mockgen -source=./captcha.go -package=repomocks -destination=mocks/captcha.mock.go CaptchaRepository
//

// Package repomocks is a generated GoMock package.
package repomocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCaptchaRepository is a mock of CaptchaRepository interface.
type MockCaptchaRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCaptchaRepositoryMockRecorder
}

// MockCaptchaRepositoryMockRecorder is the mock recorder for MockCaptchaRepository.
type MockCaptchaRepositoryMockRecorder struct {
	mock *MockCaptchaRepository
}

// NewMockCaptchaRepository creates a new mock instance.
func NewMockCaptchaRepository(ctrl *gomock.Controller) *MockCaptchaRepository {
	mock := &MockCaptchaRepository{ctrl: ctrl}
	mock.recorder = &MockCaptchaRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCaptchaRepository) EXPECT() *MockCaptchaRepositoryMockRecorder {
	return m.recorder
}

// Store mocks base method.
func (m *MockCaptchaRepository) Store(ctx context.Context, biz, phone, code string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", ctx, biz, phone, code)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockCaptchaRepositoryMockRecorder) Store(ctx, biz, phone, code any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockCaptchaRepository)(nil).Store), ctx, biz, phone, code)
}

// Verify mocks base method.
func (m *MockCaptchaRepository) Verify(ctx context.Context, biz, phone, inputCaptcha string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verify", ctx, biz, phone, inputCaptcha)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify.
func (mr *MockCaptchaRepositoryMockRecorder) Verify(ctx, biz, phone, inputCaptcha any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockCaptchaRepository)(nil).Verify), ctx, biz, phone, inputCaptcha)
}
