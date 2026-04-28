package reservenow_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/evcoreco/ocpp16messages/reservenow"
	types "github.com/evcoreco/ocpp16types"
)

const (
	testValidIDTag            = "RFID-TAG-12345"
	testValidParentIDTag      = "PARENT-12345"
	testValidExpiryDate       = "2025-01-15T10:00:00Z"
	testReservationIDOne      = 1
	testReservationIDZero     = 0
	testReservationIDMax      = 65535
	testReservationIDOver     = 65536
	testReservationIDNeg      = -1
	testConnectorIDOne        = 1
	testConnectorIDZero       = 0
	testConnectorIDMax        = 65535
	testConnectorIDOver       = 65536
	testConnectorIDNeg        = -1
	errReservationID          = "reservationId"
	errConnectorID            = "connectorId"
	errIDTag                  = "idTag"
	errExpiryDate             = "expiryDate"
	errParentIDTag            = "parentIDTag"
	errExceedsMaxLength       = "exceeds maximum length"
	errNonPrintableASCII      = "non-printable ASCII"
	fieldNameParentIDTag      = "ParentIDTag"
	wantParentIDTagNilMsg     = "ParentIDTag should be nil"
	testIDTagTooLong          = "RFID-ABC123456789012345"
	testParentIDTagTooLong    = "PARENT-1234567890123456"
	testInvalidExpiryDate     = "invalid-date"
	testIDTagWithNullByte     = "RFID\x00ABC"
	testParentTagWithNullByte = "PARENT\x00ABC"
)

func TestReq_Valid_RequiredFieldsOnly(t *testing.T) {
	t.Parallel()

	req, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDOne,
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ReservationID.Value() != uint16(testReservationIDOne) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testReservationIDOne),
			req.ReservationID.Value(),
		)
	}

	if req.ConnectorID.Value() != uint16(testConnectorIDOne) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testConnectorIDOne),
			req.ConnectorID.Value(),
		)
	}

	if req.IDTag.Value() != testValidIDTag {
		t.Errorf(types.ErrorMismatch, testValidIDTag, req.IDTag.Value())
	}

	if req.ParentIDTag != nil {
		t.Errorf(types.ErrorWantNonNil, wantParentIDTagNilMsg)
	}
}

func TestReq_Valid_WithParentIDTag(t *testing.T) {
	t.Parallel()

	parentIDTag := testValidParentIDTag

	req, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDOne,
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   &parentIDTag,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ParentIDTag == nil {
		t.Errorf(types.ErrorWantNonNil, fieldNameParentIDTag)
	}

	if req.ParentIDTag.Value() != testValidParentIDTag {
		t.Errorf(
			types.ErrorMismatch,
			testValidParentIDTag,
			req.ParentIDTag.Value(),
		)
	}
}

func TestReq_Valid_ReservationIDZero(t *testing.T) {
	t.Parallel()

	req, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDZero,
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ReservationID.Value() != uint16(testReservationIDZero) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testReservationIDZero),
			req.ReservationID.Value(),
		)
	}
}

func TestReq_Valid_ReservationIDMax(t *testing.T) {
	t.Parallel()

	req, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDMax,
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ReservationID.Value() != uint16(testReservationIDMax) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testReservationIDMax),
			req.ReservationID.Value(),
		)
	}
}

func TestReq_Valid_ConnectorIDZero(t *testing.T) {
	t.Parallel()

	req, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDOne,
		ConnectorID:   testConnectorIDZero,
		IDTag:         testValidIDTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != uint16(testConnectorIDZero) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testConnectorIDZero),
			req.ConnectorID.Value(),
		)
	}
}

func TestReq_Valid_ConnectorIDMax(t *testing.T) {
	t.Parallel()

	req, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDOne,
		ConnectorID:   testConnectorIDMax,
		IDTag:         testValidIDTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != uint16(testConnectorIDMax) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testConnectorIDMax),
			req.ConnectorID.Value(),
		)
	}
}

func TestReq_ReservationIDNegative(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDNeg,
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative reservationId")
	}

	if !strings.Contains(err.Error(), errReservationID) {
		t.Errorf(types.ErrorWantContains, err, errReservationID)
	}
}

func TestReq_ReservationIDExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDOver,
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "reservationId exceeds max")
	}

	if !strings.Contains(err.Error(), errReservationID) {
		t.Errorf(types.ErrorWantContains, err, errReservationID)
	}
}

func TestReq_ConnectorIDNegative(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDOne,
		ConnectorID:   testConnectorIDNeg,
		IDTag:         testValidIDTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative connectorId")
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}

func TestReq_ConnectorIDExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDOne,
		ConnectorID:   testConnectorIDOver,
		IDTag:         testValidIDTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "connectorId exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}

func TestReq_EmptyIDTag(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDOne,
		ConnectorID:   testConnectorIDOne,
		IDTag:         "",
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty idTag")
	}

	if !errors.Is(err, types.ErrEmptyValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrEmptyValue)
	}
}

func TestReq_IDTagTooLong(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDOne,
		ConnectorID:   testConnectorIDOne,
		IDTag:         testIDTagTooLong,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "IDTag too long")
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(types.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestReq_IDTagInvalidCharacters(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDOne,
		ConnectorID:   testConnectorIDOne,
		IDTag:         testIDTagWithNullByte,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "non-printable chars in idTag")
	}

	if !strings.Contains(err.Error(), errNonPrintableASCII) {
		t.Errorf(types.ErrorWantContains, err, errNonPrintableASCII)
	}
}

func TestReq_InvalidExpiryDate(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDOne,
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		ExpiryDate:    testInvalidExpiryDate,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid expiryDate")
	}

	if !strings.Contains(err.Error(), errExpiryDate) {
		t.Errorf(types.ErrorWantContains, err, errExpiryDate)
	}
}

func TestReq_ParentIDTagTooLong(t *testing.T) {
	t.Parallel()

	parentIDTag := testParentIDTagTooLong

	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDOne,
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   &parentIDTag,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "ParentIDTag too long")
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(types.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestReq_ParentIDTagInvalidCharacters(t *testing.T) {
	t.Parallel()

	parentIDTag := testParentTagWithNullByte

	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDOne,
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   &parentIDTag,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "non-printable chars in parentIDTag")
	}

	if !strings.Contains(err.Error(), errNonPrintableASCII) {
		t.Errorf(types.ErrorWantContains, err, errNonPrintableASCII)
	}
}

func TestReq_MultipleErrors_AllFieldsInvalid(t *testing.T) {
	t.Parallel()

	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDNeg,
		ConnectorID:   testConnectorIDNeg,
		IDTag:         "",
		ExpiryDate:    testInvalidExpiryDate,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "all fields invalid")
	}

	if !strings.Contains(err.Error(), errReservationID) {
		t.Errorf(types.ErrorWantContains, err, errReservationID)
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}

	if !strings.Contains(err.Error(), errIDTag) {
		t.Errorf(types.ErrorWantContains, err, errIDTag)
	}

	if !strings.Contains(err.Error(), errExpiryDate) {
		t.Errorf(types.ErrorWantContains, err, errExpiryDate)
	}
}

func TestReq_MultipleErrors_WithInvalidParentIDTag(t *testing.T) {
	t.Parallel()

	parentIDTag := testParentIDTagTooLong

	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: testReservationIDNeg,
		ConnectorID:   testConnectorIDOne,
		IDTag:         testValidIDTag,
		ExpiryDate:    testValidExpiryDate,
		ParentIDTag:   &parentIDTag,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "reservationId and parentIDTag invalid")
	}

	if !strings.Contains(err.Error(), errReservationID) {
		t.Errorf(types.ErrorWantContains, err, errReservationID)
	}

	if !strings.Contains(err.Error(), errParentIDTag) {
		t.Errorf(types.ErrorWantContains, err, errParentIDTag)
	}
}
