package statusnotification_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/evcoreco/ocpp16messages/statusnotification"
	types "github.com/evcoreco/ocpp16types"
)

const (
	testConnectorIDOne   = 1
	testConnectorIDZero  = 0
	testValueNegativeOne = -1
	testStatusAvailable  = "Available"
	testStatusCharging   = "Charging"
	testErrorCodeNoError = "NoError"
	errFieldConnectorID  = "connectorId"
	errFieldErrorCode    = "errorCode"
	errFieldStatus       = "status"
	errFieldInfo         = "info"
	errFieldTimestamp    = "timestamp"
	errWantInvalidValue  = "Req() error = %v, want ErrInvalidValue"
)

func TestReq_Valid(t *testing.T) {
	t.Parallel()

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorID:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != testConnectorIDOne {
		t.Errorf(
			types.ErrorMismatch, testConnectorIDOne, req.ConnectorID.Value(),
		)
	}

	if req.ErrorCode.String() != testErrorCodeNoError {
		t.Errorf(
			types.ErrorMismatch, testErrorCodeNoError, req.ErrorCode.String(),
		)
	}

	if req.Status.String() != testStatusAvailable {
		t.Errorf(types.ErrorMismatch, testStatusAvailable, req.Status.String())
	}
}

func TestReq_ValidConnectorIDZero(t *testing.T) {
	t.Parallel()

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDZero,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorID:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != testConnectorIDZero {
		t.Errorf(
			types.ErrorMismatch, testConnectorIDZero, req.ConnectorID.Value(),
		)
	}
}

func TestReq_ValidCharging(t *testing.T) {
	t.Parallel()

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusCharging,
		Info:            nil,
		Timestamp:       nil,
		VendorID:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Status.String() != testStatusCharging {
		t.Errorf(types.ErrorMismatch, testStatusCharging, req.Status.String())
	}
}

func TestReq_ValidWithInfo(t *testing.T) {
	t.Parallel()

	info := "Ground fault detected"

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDOne,
		ErrorCode:       "GroundFailure",
		Status:          "Faulted",
		Info:            &info,
		Timestamp:       nil,
		VendorID:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Info == nil {
		t.Error("Req() Info = nil, want non-nil")
	}

	if req.Info.Value() != info {
		t.Errorf(types.ErrorMismatch, info, req.Info.Value())
	}
}

func TestReq_ValidWithTimestamp(t *testing.T) {
	t.Parallel()

	timestamp := "2025-01-15T10:30:00Z"

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       &timestamp,
		VendorID:        nil,
		VendorErrorCode: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Timestamp == nil {
		t.Error("Req() Timestamp = nil, want non-nil")
	}
}

func TestReq_ValidWithVendorFields(t *testing.T) {
	t.Parallel()

	vendorId := "VendorX"
	vendorErrorCode := "V001"

	req, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorID:        &vendorId,
		VendorErrorCode: &vendorErrorCode,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.VendorID == nil {
		t.Error("Req() VendorID = nil, want non-nil")
	}

	if req.VendorErrorCode == nil {
		t.Error("Req() VendorErrorCode = nil, want non-nil")
	}

	if req.VendorID.Value() != vendorId {
		t.Errorf(types.ErrorMismatch, vendorId, req.VendorID.Value())
	}
}

func TestReq_ConnectorIDNegative(t *testing.T) {
	t.Parallel()

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testValueNegativeOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorID:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for negative connectorId")
	}

	if !strings.Contains(err.Error(), errFieldConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errFieldConnectorID)
	}
}

func TestReq_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          "InvalidStatus",
		Info:            nil,
		Timestamp:       nil,
		VendorID:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(errWantInvalidValue, err)
	}

	if !strings.Contains(err.Error(), errFieldStatus) {
		t.Errorf(types.ErrorWantContains, err, errFieldStatus)
	}
}

func TestReq_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          "",
		Info:            nil,
		Timestamp:       nil,
		VendorID:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(errWantInvalidValue, err)
	}
}

func TestReq_InvalidErrorCode(t *testing.T) {
	t.Parallel()

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDOne,
		ErrorCode:       "InvalidCode",
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorID:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid errorCode")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(errWantInvalidValue, err)
	}

	if !strings.Contains(err.Error(), errFieldErrorCode) {
		t.Errorf(types.ErrorWantContains, err, errFieldErrorCode)
	}
}

func TestReq_EmptyErrorCode(t *testing.T) {
	t.Parallel()

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDOne,
		ErrorCode:       "",
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorID:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty errorCode")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(errWantInvalidValue, err)
	}
}

func TestReq_InvalidTimestamp(t *testing.T) {
	t.Parallel()

	invalidTimestamp := "not-a-timestamp"

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       &invalidTimestamp,
		VendorID:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid timestamp")
	}

	if !strings.Contains(err.Error(), errFieldTimestamp) {
		t.Errorf(types.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_InfoTooLong(t *testing.T) {
	t.Parallel()

	// 51 chars, max is 50
	longInfo := "This info string is way too long for the OCPP spec!"

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            &longInfo,
		Timestamp:       nil,
		VendorID:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for info too long")
	}

	if !strings.Contains(err.Error(), errFieldInfo) {
		t.Errorf(types.ErrorWantContains, err, errFieldInfo)
	}
}

func TestReq_MultipleErrors(t *testing.T) {
	t.Parallel()

	invalidTimestamp := "invalid"

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testValueNegativeOne,
		ErrorCode:       "InvalidCode",
		Status:          "InvalidStatus",
		Info:            nil,
		Timestamp:       &invalidTimestamp,
		VendorID:        nil,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errFieldConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errFieldConnectorID)
	}

	if !strings.Contains(errStr, errFieldErrorCode) {
		t.Errorf(types.ErrorWantContains, err, errFieldErrorCode)
	}

	if !strings.Contains(errStr, errFieldStatus) {
		t.Errorf(types.ErrorWantContains, err, errFieldStatus)
	}

	if !strings.Contains(errStr, errFieldTimestamp) {
		t.Errorf(types.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_AllStatuses(t *testing.T) {
	t.Parallel()

	statuses := []string{
		"Available",
		"Preparing",
		"Charging",
		"SuspendedEV",
		"SuspendedEVSE",
		"Finishing",
		"Reserved",
		"Unavailable",
		"Faulted",
	}

	for _, status := range statuses {
		req, err := statusnotification.Req(statusnotification.ReqInput{
			ConnectorID:     testConnectorIDOne,
			ErrorCode:       testErrorCodeNoError,
			Status:          status,
			Info:            nil,
			Timestamp:       nil,
			VendorID:        nil,
			VendorErrorCode: nil,
		})
		if err != nil {
			t.Errorf("Req() error = %v for status %s", err, status)
		}

		if req.Status.String() != status {
			t.Errorf(types.ErrorMismatch, status, req.Status.String())
		}
	}
}

func TestReq_AllErrorCodes(t *testing.T) {
	t.Parallel()

	errorCodes := []string{
		"ConnectorLockFailure",
		"EVCommunicationError",
		"GroundFailure",
		"HighTemperature",
		"InternalError",
		"LocalListConflict",
		"NoError",
		"OtherError",
		"OverCurrentFailure",
		"OverVoltage",
		"PowerMeterFailure",
		"PowerSwitchFailure",
		"ReaderFailure",
		"ResetFailure",
		"UnderVoltage",
		"WeakSignal",
	}

	for _, errorCode := range errorCodes {
		req, err := statusnotification.Req(statusnotification.ReqInput{
			ConnectorID:     testConnectorIDOne,
			ErrorCode:       errorCode,
			Status:          testStatusAvailable,
			Info:            nil,
			Timestamp:       nil,
			VendorID:        nil,
			VendorErrorCode: nil,
		})
		if err != nil {
			t.Errorf("Req() error = %v for errorCode %s", err, errorCode)
		}

		if req.ErrorCode.String() != errorCode {
			t.Errorf(types.ErrorMismatch, errorCode, req.ErrorCode.String())
		}
	}
}

func TestReq_InvalidVendorID(t *testing.T) {
	t.Parallel()

	// VendorID must be valid CiString255, test with invalid ASCII
	invalidVendorID := "Vendor\x00Id"

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorID:        &invalidVendorID,
		VendorErrorCode: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid vendorId")
	}

	if !strings.Contains(err.Error(), "vendorId") {
		t.Errorf(types.ErrorWantContains, err, "vendorId")
	}
}

func TestReq_InvalidVendorErrorCode(t *testing.T) {
	t.Parallel()

	// VendorErrorCode must be valid CiString50, test with invalid ASCII
	invalidVendorErrorCode := "Error\x00Code"

	_, err := statusnotification.Req(statusnotification.ReqInput{
		ConnectorID:     testConnectorIDOne,
		ErrorCode:       testErrorCodeNoError,
		Status:          testStatusAvailable,
		Info:            nil,
		Timestamp:       nil,
		VendorID:        nil,
		VendorErrorCode: &invalidVendorErrorCode,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid vendorErrorCode")
	}

	if !strings.Contains(err.Error(), "vendorErrorCode") {
		t.Errorf(types.ErrorWantContains, err, "vendorErrorCode")
	}
}
