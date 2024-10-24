// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/server/server.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	storager "github.com/GlebZigert/url_shortener.git/internal/storager"
	gomock "github.com/golang/mock/gomock"
)

// MocksrvConfig is a mock of srvConfig interface.
type MocksrvConfig struct {
	ctrl     *gomock.Controller
	recorder *MocksrvConfigMockRecorder
}

// MocksrvConfigMockRecorder is the mock recorder for MocksrvConfig.
type MocksrvConfigMockRecorder struct {
	mock *MocksrvConfig
}

// NewMocksrvConfig creates a new mock instance.
func NewMocksrvConfig(ctrl *gomock.Controller) *MocksrvConfig {
	mock := &MocksrvConfig{ctrl: ctrl}
	mock.recorder = &MocksrvConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksrvConfig) EXPECT() *MocksrvConfigMockRecorder {
	return m.recorder
}

// GetBaseURL mocks base method.
func (m *MocksrvConfig) GetBaseURL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBaseURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetBaseURL indicates an expected call of GetBaseURL.
func (mr *MocksrvConfigMockRecorder) GetBaseURL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBaseURL", reflect.TypeOf((*MocksrvConfig)(nil).GetBaseURL))
}

// GetDatabaseDSN mocks base method.
func (m *MocksrvConfig) GetDatabaseDSN() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDatabaseDSN")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDatabaseDSN indicates an expected call of GetDatabaseDSN.
func (mr *MocksrvConfigMockRecorder) GetDatabaseDSN() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDatabaseDSN", reflect.TypeOf((*MocksrvConfig)(nil).GetDatabaseDSN))
}

// GetFileStoragePath mocks base method.
func (m *MocksrvConfig) GetFileStoragePath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileStoragePath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFileStoragePath indicates an expected call of GetFileStoragePath.
func (mr *MocksrvConfigMockRecorder) GetFileStoragePath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileStoragePath", reflect.TypeOf((*MocksrvConfig)(nil).GetFileStoragePath))
}

// GetFlagLogLevel mocks base method.
func (m *MocksrvConfig) GetFlagLogLevel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFlagLogLevel")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetFlagLogLevel indicates an expected call of GetFlagLogLevel.
func (mr *MocksrvConfigMockRecorder) GetFlagLogLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFlagLogLevel", reflect.TypeOf((*MocksrvConfig)(nil).GetFlagLogLevel))
}

// GetNumWorkers mocks base method.
func (m *MocksrvConfig) GetNumWorkers() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumWorkers")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetNumWorkers indicates an expected call of GetNumWorkers.
func (mr *MocksrvConfigMockRecorder) GetNumWorkers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumWorkers", reflect.TypeOf((*MocksrvConfig)(nil).GetNumWorkers))
}

// GetRunAddr mocks base method.
func (m *MocksrvConfig) GetRunAddr() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunAddr")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRunAddr indicates an expected call of GetRunAddr.
func (mr *MocksrvConfigMockRecorder) GetRunAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunAddr", reflect.TypeOf((*MocksrvConfig)(nil).GetRunAddr))
}

// GetSECRETKEY mocks base method.
func (m *MocksrvConfig) GetSECRETKEY() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSECRETKEY")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSECRETKEY indicates an expected call of GetSECRETKEY.
func (mr *MocksrvConfigMockRecorder) GetSECRETKEY() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSECRETKEY", reflect.TypeOf((*MocksrvConfig)(nil).GetSECRETKEY))
}

// GetTOKENEXP mocks base method.
func (m *MocksrvConfig) GetTOKENEXP() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTOKENEXP")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetTOKENEXP indicates an expected call of GetTOKENEXP.
func (mr *MocksrvConfigMockRecorder) GetTOKENEXP() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTOKENEXP", reflect.TypeOf((*MocksrvConfig)(nil).GetTOKENEXP))
}

// MocksrvMiddleware is a mock of srvMiddleware interface.
type MocksrvMiddleware struct {
	ctrl     *gomock.Controller
	recorder *MocksrvMiddlewareMockRecorder
}

// MocksrvMiddlewareMockRecorder is the mock recorder for MocksrvMiddleware.
type MocksrvMiddlewareMockRecorder struct {
	mock *MocksrvMiddleware
}

// NewMocksrvMiddleware creates a new mock instance.
func NewMocksrvMiddleware(ctrl *gomock.Controller) *MocksrvMiddleware {
	mock := &MocksrvMiddleware{ctrl: ctrl}
	mock.recorder = &MocksrvMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksrvMiddleware) EXPECT() *MocksrvMiddlewareMockRecorder {
	return m.recorder
}

// Auth mocks base method.
func (m *MocksrvMiddleware) Auth(h http.Handler) http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Auth", h)
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// Auth indicates an expected call of Auth.
func (mr *MocksrvMiddlewareMockRecorder) Auth(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Auth", reflect.TypeOf((*MocksrvMiddleware)(nil).Auth), h)
}

// CheckUID mocks base method.
func (m *MocksrvMiddleware) CheckUID(ctx context.Context) (int, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUID", ctx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CheckUID indicates an expected call of CheckUID.
func (mr *MocksrvMiddlewareMockRecorder) CheckUID(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUID", reflect.TypeOf((*MocksrvMiddleware)(nil).CheckUID), ctx)
}

// ErrHandler mocks base method.
func (m *MocksrvMiddleware) ErrHandler(f http.Handler) http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ErrHandler", f)
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// ErrHandler indicates an expected call of ErrHandler.
func (mr *MocksrvMiddlewareMockRecorder) ErrHandler(f interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrHandler", reflect.TypeOf((*MocksrvMiddleware)(nil).ErrHandler), f)
}

// Log mocks base method.
func (m *MocksrvMiddleware) Log(h http.Handler) http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Log", h)
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// Log indicates an expected call of Log.
func (mr *MocksrvMiddlewareMockRecorder) Log(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MocksrvMiddleware)(nil).Log), h)
}

// MocksrvLogger is a mock of srvLogger interface.
type MocksrvLogger struct {
	ctrl     *gomock.Controller
	recorder *MocksrvLoggerMockRecorder
}

// MocksrvLoggerMockRecorder is the mock recorder for MocksrvLogger.
type MocksrvLoggerMockRecorder struct {
	mock *MocksrvLogger
}

// NewMocksrvLogger creates a new mock instance.
func NewMocksrvLogger(ctrl *gomock.Controller) *MocksrvLogger {
	mock := &MocksrvLogger{ctrl: ctrl}
	mock.recorder = &MocksrvLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksrvLogger) EXPECT() *MocksrvLoggerMockRecorder {
	return m.recorder
}

// Error mocks base method.
func (m *MocksrvLogger) Error(msg string, fields map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Error", msg, fields)
}

// Error indicates an expected call of Error.
func (mr *MocksrvLoggerMockRecorder) Error(msg, fields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MocksrvLogger)(nil).Error), msg, fields)
}

// Info mocks base method.
func (m *MocksrvLogger) Info(msg string, fields map[string]interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Info", msg, fields)
}

// Info indicates an expected call of Info.
func (mr *MocksrvLoggerMockRecorder) Info(msg, fields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MocksrvLogger)(nil).Info), msg, fields)
}

// MocksrvService is a mock of srvService interface.
type MocksrvService struct {
	ctrl     *gomock.Controller
	recorder *MocksrvServiceMockRecorder
}

// MocksrvServiceMockRecorder is the mock recorder for MocksrvService.
type MocksrvServiceMockRecorder struct {
	mock *MocksrvService
}

// NewMocksrvService creates a new mock instance.
func NewMocksrvService(ctrl *gomock.Controller) *MocksrvService {
	mock := &MocksrvService{ctrl: ctrl}
	mock.recorder = &MocksrvServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocksrvService) EXPECT() *MocksrvServiceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MocksrvService) Delete(shorts []string, uid int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", shorts, uid)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MocksrvServiceMockRecorder) Delete(shorts, uid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MocksrvService)(nil).Delete), shorts, uid)
}

// GetAll mocks base method.
func (m *MocksrvService) GetAll() *[]*storager.Shorten {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(*[]*storager.Shorten)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MocksrvServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MocksrvService)(nil).GetAll))
}

// Origin mocks base method.
func (m *MocksrvService) Origin(short string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Origin", short)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Origin indicates an expected call of Origin.
func (mr *MocksrvServiceMockRecorder) Origin(short interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Origin", reflect.TypeOf((*MocksrvService)(nil).Origin), short)
}

// Short mocks base method.
func (m *MocksrvService) Short(oririn string, uuid int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Short", oririn, uuid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Short indicates an expected call of Short.
func (mr *MocksrvServiceMockRecorder) Short(oririn, uuid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Short", reflect.TypeOf((*MocksrvService)(nil).Short), oririn, uuid)
}
