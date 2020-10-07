// Code generated by MockGen. DO NOT EDIT.
// Source: host.go

// Package host is a generated GoMock package.
package host

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	gorm "github.com/jinzhu/gorm"
	common "github.com/openshift/assisted-service/internal/common"
	models "github.com/openshift/assisted-service/models"
	logrus "github.com/sirupsen/logrus"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// RegisterHost mocks base method
func (m *MockAPI) RegisterHost(ctx context.Context, h *models.Host) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterHost", ctx, h)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterHost indicates an expected call of RegisterHost
func (mr *MockAPIMockRecorder) RegisterHost(ctx, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterHost", reflect.TypeOf((*MockAPI)(nil).RegisterHost), ctx, h)
}

// HandleInstallationFailure mocks base method
func (m *MockAPI) HandleInstallationFailure(ctx context.Context, h *models.Host) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleInstallationFailure", ctx, h)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleInstallationFailure indicates an expected call of HandleInstallationFailure
func (mr *MockAPIMockRecorder) HandleInstallationFailure(ctx, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleInstallationFailure", reflect.TypeOf((*MockAPI)(nil).HandleInstallationFailure), ctx, h)
}

// GetNextSteps mocks base method
func (m *MockAPI) GetNextSteps(ctx context.Context, host *models.Host) (models.Steps, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextSteps", ctx, host)
	ret0, _ := ret[0].(models.Steps)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNextSteps indicates an expected call of GetNextSteps
func (mr *MockAPIMockRecorder) GetNextSteps(ctx, host interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextSteps", reflect.TypeOf((*MockAPI)(nil).GetNextSteps), ctx, host)
}

// UpdateInstallProgress mocks base method
func (m *MockAPI) UpdateInstallProgress(ctx context.Context, h *models.Host, progress *models.HostProgress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInstallProgress", ctx, h, progress)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInstallProgress indicates an expected call of UpdateInstallProgress
func (mr *MockAPIMockRecorder) UpdateInstallProgress(ctx, h, progress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInstallProgress", reflect.TypeOf((*MockAPI)(nil).UpdateInstallProgress), ctx, h, progress)
}

// RefreshStatus mocks base method
func (m *MockAPI) RefreshStatus(ctx context.Context, h *models.Host, db *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshStatus", ctx, h, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshStatus indicates an expected call of RefreshStatus
func (mr *MockAPIMockRecorder) RefreshStatus(ctx, h, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshStatus", reflect.TypeOf((*MockAPI)(nil).RefreshStatus), ctx, h, db)
}

// SetBootstrap mocks base method
func (m *MockAPI) SetBootstrap(ctx context.Context, h *models.Host, isbootstrap bool, db *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBootstrap", ctx, h, isbootstrap, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBootstrap indicates an expected call of SetBootstrap
func (mr *MockAPIMockRecorder) SetBootstrap(ctx, h, isbootstrap, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBootstrap", reflect.TypeOf((*MockAPI)(nil).SetBootstrap), ctx, h, isbootstrap, db)
}

// UpdateConnectivityReport mocks base method
func (m *MockAPI) UpdateConnectivityReport(ctx context.Context, h *models.Host, connectivityReport string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateConnectivityReport", ctx, h, connectivityReport)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateConnectivityReport indicates an expected call of UpdateConnectivityReport
func (mr *MockAPIMockRecorder) UpdateConnectivityReport(ctx, h, connectivityReport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateConnectivityReport", reflect.TypeOf((*MockAPI)(nil).UpdateConnectivityReport), ctx, h, connectivityReport)
}

// UpdateApiVipConnectivityReport mocks base method
func (m *MockAPI) UpdateApiVipConnectivityReport(ctx context.Context, h *models.Host, connectivityReport string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateApiVipConnectivityReport", ctx, h, connectivityReport)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateApiVipConnectivityReport indicates an expected call of UpdateApiVipConnectivityReport
func (mr *MockAPIMockRecorder) UpdateApiVipConnectivityReport(ctx, h, connectivityReport interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApiVipConnectivityReport", reflect.TypeOf((*MockAPI)(nil).UpdateApiVipConnectivityReport), ctx, h, connectivityReport)
}

// HostMonitoring mocks base method
func (m *MockAPI) HostMonitoring() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HostMonitoring")
}

// HostMonitoring indicates an expected call of HostMonitoring
func (mr *MockAPIMockRecorder) HostMonitoring() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HostMonitoring", reflect.TypeOf((*MockAPI)(nil).HostMonitoring))
}

// UpdateRole mocks base method
func (m *MockAPI) UpdateRole(ctx context.Context, h *models.Host, role models.HostRole, db *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", ctx, h, role, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRole indicates an expected call of UpdateRole
func (mr *MockAPIMockRecorder) UpdateRole(ctx, h, role, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockAPI)(nil).UpdateRole), ctx, h, role, db)
}

// UpdateHostname mocks base method
func (m *MockAPI) UpdateHostname(ctx context.Context, h *models.Host, hostname string, db *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHostname", ctx, h, hostname, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHostname indicates an expected call of UpdateHostname
func (mr *MockAPIMockRecorder) UpdateHostname(ctx, h, hostname, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHostname", reflect.TypeOf((*MockAPI)(nil).UpdateHostname), ctx, h, hostname, db)
}

// CancelInstallation mocks base method
func (m *MockAPI) CancelInstallation(ctx context.Context, h *models.Host, reason string, db *gorm.DB) *common.ApiErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelInstallation", ctx, h, reason, db)
	ret0, _ := ret[0].(*common.ApiErrorResponse)
	return ret0
}

// CancelInstallation indicates an expected call of CancelInstallation
func (mr *MockAPIMockRecorder) CancelInstallation(ctx, h, reason, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelInstallation", reflect.TypeOf((*MockAPI)(nil).CancelInstallation), ctx, h, reason, db)
}

// IsRequireUserActionReset mocks base method
func (m *MockAPI) IsRequireUserActionReset(h *models.Host) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsRequireUserActionReset", h)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsRequireUserActionReset indicates an expected call of IsRequireUserActionReset
func (mr *MockAPIMockRecorder) IsRequireUserActionReset(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsRequireUserActionReset", reflect.TypeOf((*MockAPI)(nil).IsRequireUserActionReset), h)
}

// ResetHost mocks base method
func (m *MockAPI) ResetHost(ctx context.Context, h *models.Host, reason string, db *gorm.DB) *common.ApiErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetHost", ctx, h, reason, db)
	ret0, _ := ret[0].(*common.ApiErrorResponse)
	return ret0
}

// ResetHost indicates an expected call of ResetHost
func (mr *MockAPIMockRecorder) ResetHost(ctx, h, reason, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetHost", reflect.TypeOf((*MockAPI)(nil).ResetHost), ctx, h, reason, db)
}

// ResetPendingUserAction mocks base method
func (m *MockAPI) ResetPendingUserAction(ctx context.Context, h *models.Host, db *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetPendingUserAction", ctx, h, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetPendingUserAction indicates an expected call of ResetPendingUserAction
func (mr *MockAPIMockRecorder) ResetPendingUserAction(ctx, h, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetPendingUserAction", reflect.TypeOf((*MockAPI)(nil).ResetPendingUserAction), ctx, h, db)
}

// DisableHost mocks base method
func (m *MockAPI) DisableHost(ctx context.Context, h *models.Host, db *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableHost", ctx, h, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableHost indicates an expected call of DisableHost
func (mr *MockAPIMockRecorder) DisableHost(ctx, h, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableHost", reflect.TypeOf((*MockAPI)(nil).DisableHost), ctx, h, db)
}

// EnableHost mocks base method
func (m *MockAPI) EnableHost(ctx context.Context, h *models.Host, db *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableHost", ctx, h, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableHost indicates an expected call of EnableHost
func (mr *MockAPIMockRecorder) EnableHost(ctx, h, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableHost", reflect.TypeOf((*MockAPI)(nil).EnableHost), ctx, h, db)
}

// Install mocks base method
func (m *MockAPI) Install(ctx context.Context, h *models.Host, db *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", ctx, h, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install
func (mr *MockAPIMockRecorder) Install(ctx, h, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockAPI)(nil).Install), ctx, h, db)
}

// UpdateInventory mocks base method
func (m *MockAPI) UpdateInventory(ctx context.Context, h *models.Host, inventory string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInventory", ctx, h, inventory)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInventory indicates an expected call of UpdateInventory
func (mr *MockAPIMockRecorder) UpdateInventory(ctx, h, inventory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInventory", reflect.TypeOf((*MockAPI)(nil).UpdateInventory), ctx, h, inventory)
}

// GetStagesByRole mocks base method
func (m *MockAPI) GetStagesByRole(role models.HostRole, isbootstrap bool) []models.HostStage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStagesByRole", role, isbootstrap)
	ret0, _ := ret[0].([]models.HostStage)
	return ret0
}

// GetStagesByRole indicates an expected call of GetStagesByRole
func (mr *MockAPIMockRecorder) GetStagesByRole(role, isbootstrap interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStagesByRole", reflect.TypeOf((*MockAPI)(nil).GetStagesByRole), role, isbootstrap)
}

// IsInstallable mocks base method
func (m *MockAPI) IsInstallable(h *models.Host) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsInstallable", h)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsInstallable indicates an expected call of IsInstallable
func (mr *MockAPIMockRecorder) IsInstallable(h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsInstallable", reflect.TypeOf((*MockAPI)(nil).IsInstallable), h)
}

// PrepareForInstallation mocks base method
func (m *MockAPI) PrepareForInstallation(ctx context.Context, h *models.Host, db *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrepareForInstallation", ctx, h, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// PrepareForInstallation indicates an expected call of PrepareForInstallation
func (mr *MockAPIMockRecorder) PrepareForInstallation(ctx, h, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrepareForInstallation", reflect.TypeOf((*MockAPI)(nil).PrepareForInstallation), ctx, h, db)
}

// AutoAssignRole mocks base method
func (m *MockAPI) AutoAssignRole(ctx context.Context, h *models.Host, db *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AutoAssignRole", ctx, h, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// AutoAssignRole indicates an expected call of AutoAssignRole
func (mr *MockAPIMockRecorder) AutoAssignRole(ctx, h, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AutoAssignRole", reflect.TypeOf((*MockAPI)(nil).AutoAssignRole), ctx, h, db)
}

// IsValidMasterCandidate mocks base method
func (m *MockAPI) IsValidMasterCandidate(h *models.Host, db *gorm.DB, log logrus.FieldLogger) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidMasterCandidate", h, db, log)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsValidMasterCandidate indicates an expected call of IsValidMasterCandidate
func (mr *MockAPIMockRecorder) IsValidMasterCandidate(h, db, log interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidMasterCandidate", reflect.TypeOf((*MockAPI)(nil).IsValidMasterCandidate), h, db, log)
}

// SetUploadLogsAt mocks base method
func (m *MockAPI) SetUploadLogsAt(ctx context.Context, h *models.Host, db *gorm.DB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUploadLogsAt", ctx, h, db)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUploadLogsAt indicates an expected call of SetUploadLogsAt
func (mr *MockAPIMockRecorder) SetUploadLogsAt(ctx, h, db interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUploadLogsAt", reflect.TypeOf((*MockAPI)(nil).SetUploadLogsAt), ctx, h, db)
}

// GetHostRequirements mocks base method
func (m *MockAPI) GetHostRequirements(role models.HostRole) models.HostRequirementsRole {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostRequirements", role)
	ret0, _ := ret[0].(models.HostRequirementsRole)
	return ret0
}

// GetHostRequirements indicates an expected call of GetHostRequirements
func (mr *MockAPIMockRecorder) GetHostRequirements(role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostRequirements", reflect.TypeOf((*MockAPI)(nil).GetHostRequirements), role)
}
