package starttransaction_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/evcoreco/ocpp16messages/starttransaction"
	types "github.com/evcoreco/ocpp16types"
)

const (
	testValidIDTag        = "RFID-TAG-12345"
	testValidTimestamp    = "2025-01-15T10:30:00Z"
	testConnectorIDOne    = 1
	testMeterStart1000    = 1000
	testReservationID42   = 42
	testValueZero         = 0
	testValueNegativeOne  = -1
	errFieldConnectorID   = "connectorId"
	errFieldMeterStart    = "meterStart"
	errFieldTimestamp     = "timestamp"
	errFieldReservationID = "reservationId"
)

func TestReq_Valid(t *testing.T) {
	t.Parallel()

	req, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationID: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != testConnectorIDOne {
		t.Errorf(
			types.ErrorMismatch, testConnectorIDOne, req.ConnectorID.Value(),
		)
	}

	if req.IDTag.String() != testValidIDTag {
		t.Errorf(types.ErrorMismatch, testValidIDTag, req.IDTag.String())
	}

	if req.MeterStart.Value() != testMeterStart1000 {
		t.Errorf(
			types.ErrorMismatch, testMeterStart1000, req.MeterStart.Value(),
		)
	}
}

func TestReq_ValidWithReservation(t *testing.T) {
	t.Parallel()

	reservationId := testReservationID42

	req, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationID: &reservationId,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ReservationID == nil {
		t.Error("Req() ReservationID = nil, want non-nil")
	}

	expectedReservationID := uint16(testReservationID42)
	if req.ReservationID.Value() != expectedReservationID {
		t.Errorf(
			types.ErrorMismatch,
			expectedReservationID,
			req.ReservationID.Value(),
		)
	}
}

func TestReq_ConnectorIDZero(t *testing.T) {
	t.Parallel()

	req, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testValueZero,
		IDTag:         testValidIDTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationID: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != testValueZero {
		t.Errorf(types.ErrorMismatch, testValueZero, req.ConnectorID.Value())
	}
}

func TestReq_ConnectorIDNegative(t *testing.T) {
	t.Parallel()

	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testValueNegativeOne,
		IDTag:         testValidIDTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationID: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for negative connectorId")
	}

	if !strings.Contains(err.Error(), errFieldConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errFieldConnectorID)
	}
}

func TestReq_EmptyIDTag(t *testing.T) {
	t.Parallel()

	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testConnectorIDOne,
		IDTag:         "",
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationID: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty idTag")
	}

	if !errors.Is(err, types.ErrEmptyValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrEmptyValue)
	}
}

func TestReq_IDTagTooLong(t *testing.T) {
	t.Parallel()

	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testConnectorIDOne,
		IDTag:         "RFID-ABC123456789012345", // 23 chars, max is 20
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationID: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for IDTag too long")
	}

	if !strings.Contains(err.Error(), "exceeds maximum length") {
		t.Errorf(
			"Req() error = %v, want 'exceeds maximum length'",
			err,
		)
	}
}

func TestReq_IDTagInvalidCharacters(t *testing.T) {
	t.Parallel()

	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testConnectorIDOne,
		IDTag:         "RFID\x00ABC",
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationID: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for non-printable chars")
	}

	if !strings.Contains(err.Error(), "non-printable ASCII") {
		t.Errorf(
			"Req() error = %v, want 'non-printable ASCII'",
			err,
		)
	}
}

func TestReq_MeterStartZero(t *testing.T) {
	t.Parallel()

	req, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		MeterStart:    testValueZero,
		Timestamp:     testValidTimestamp,
		ReservationID: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.MeterStart.Value() != testValueZero {
		t.Errorf(types.ErrorMismatch, testValueZero, req.MeterStart.Value())
	}
}

func TestReq_MeterStartNegative(t *testing.T) {
	t.Parallel()

	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		MeterStart:    testValueNegativeOne,
		Timestamp:     testValidTimestamp,
		ReservationID: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for negative meterStart")
	}

	if !strings.Contains(err.Error(), errFieldMeterStart) {
		t.Errorf(types.ErrorWantContains, err, errFieldMeterStart)
	}
}

func TestReq_InvalidTimestamp(t *testing.T) {
	t.Parallel()

	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     "not-a-timestamp",
		ReservationID: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid timestamp")
	}

	if !strings.Contains(err.Error(), errFieldTimestamp) {
		t.Errorf(types.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_EmptyTimestamp(t *testing.T) {
	t.Parallel()

	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     "",
		ReservationID: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty timestamp")
	}

	if !strings.Contains(err.Error(), errFieldTimestamp) {
		t.Errorf(types.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_ReservationIDNegative(t *testing.T) {
	t.Parallel()

	reservationId := testValueNegativeOne

	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		MeterStart:    testMeterStart1000,
		Timestamp:     testValidTimestamp,
		ReservationID: &reservationId,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for negative reservationId")
	}

	if !strings.Contains(err.Error(), errFieldReservationID) {
		t.Errorf(types.ErrorWantContains, err, errFieldReservationID)
	}
}

func TestReq_MultipleErrors(t *testing.T) {
	t.Parallel()

	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testValueNegativeOne,
		IDTag:         "",
		MeterStart:    testValueNegativeOne,
		Timestamp:     "invalid",
		ReservationID: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errFieldConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errFieldConnectorID)
	}

	if !strings.Contains(errStr, "idTag") {
		t.Errorf(types.ErrorWantContains, err, "idTag")
	}

	if !strings.Contains(errStr, errFieldMeterStart) {
		t.Errorf(types.ErrorWantContains, err, errFieldMeterStart)
	}

	if !strings.Contains(errStr, errFieldTimestamp) {
		t.Errorf(types.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_MultipleErrorsWithReservation(t *testing.T) {
	t.Parallel()

	reservationId := testValueNegativeOne

	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   testValueNegativeOne,
		IDTag:         "",
		MeterStart:    testValueNegativeOne,
		Timestamp:     "invalid",
		ReservationID: &reservationId,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errFieldConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errFieldConnectorID)
	}

	if !strings.Contains(errStr, errFieldReservationID) {
		t.Errorf(types.ErrorWantContains, err, errFieldReservationID)
	}
}
