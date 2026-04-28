package remotestarttransaction_test

import (
	"errors"
	"strings"
	"testing"

	rst "github.com/evcoreco/ocpp16messages/remotestarttransaction"
	types "github.com/evcoreco/ocpp16types"
)

const (
	testValidIDTag        = "RFID-TAG-12345"
	testConnectorIDOne    = 1
	testConnectorIDZero   = 0
	testConnectorIDMax    = 65535
	testConnectorIDOver   = 65536
	testConnectorIDNeg    = -1
	errIDTag              = "idTag"
	errConnectorID        = "connectorId"
	errExceedsMaxLength   = "exceeds maximum length"
	errNonPrintableASCII  = "non-printable ASCII"
	fieldNameConnectorID  = "ConnectorID"
	wantConnectorIDNilMsg = "ConnectorID should be nil"
)

func TestReq_Valid_IDTagOnly(t *testing.T) {
	t.Parallel()

	req, err := rst.Req(rst.ReqInput{
		IDTag:       testValidIDTag,
		ConnectorID: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.IDTag.Value() != testValidIDTag {
		t.Errorf(types.ErrorMismatch, testValidIDTag, req.IDTag.Value())
	}

	if req.ConnectorID != nil {
		t.Errorf(types.ErrorWantNonNil, wantConnectorIDNilMsg)
	}
}

func TestReq_Valid_WithConnectorID(t *testing.T) {
	t.Parallel()

	connectorId := testConnectorIDOne

	req, err := rst.Req(rst.ReqInput{
		IDTag:       testValidIDTag,
		ConnectorID: &connectorId,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.IDTag.Value() != testValidIDTag {
		t.Errorf(types.ErrorMismatch, testValidIDTag, req.IDTag.Value())
	}

	if req.ConnectorID == nil {
		t.Errorf(types.ErrorWantNonNil, fieldNameConnectorID)
	}

	if req.ConnectorID.Value() != uint16(testConnectorIDOne) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testConnectorIDOne),
			req.ConnectorID.Value(),
		)
	}
}

func TestReq_Valid_ConnectorIDZero(t *testing.T) {
	t.Parallel()

	connectorId := testConnectorIDZero

	req, err := rst.Req(rst.ReqInput{
		IDTag:       testValidIDTag,
		ConnectorID: &connectorId,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID == nil {
		t.Errorf(types.ErrorWantNonNil, fieldNameConnectorID)
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

	connectorId := testConnectorIDMax

	req, err := rst.Req(rst.ReqInput{
		IDTag:       testValidIDTag,
		ConnectorID: &connectorId,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID == nil {
		t.Errorf(types.ErrorWantNonNil, fieldNameConnectorID)
	}

	if req.ConnectorID.Value() != uint16(testConnectorIDMax) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testConnectorIDMax),
			req.ConnectorID.Value(),
		)
	}
}

func TestReq_EmptyIDTag(t *testing.T) {
	t.Parallel()

	_, err := rst.Req(rst.ReqInput{IDTag: "", ConnectorID: nil})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty idTag")
	}

	if !errors.Is(err, types.ErrEmptyValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrEmptyValue)
	}
}

func TestReq_IDTagTooLong(t *testing.T) {
	t.Parallel()

	// 23 chars, max is 20
	_, err := rst.Req(rst.ReqInput{
		IDTag:       "RFID-ABC123456789012345",
		ConnectorID: nil,
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

	// Contains null byte
	_, err := rst.Req(rst.ReqInput{IDTag: "RFID\x00ABC", ConnectorID: nil})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "non-printable chars in idTag")
	}

	if !strings.Contains(err.Error(), errNonPrintableASCII) {
		t.Errorf(types.ErrorWantContains, err, errNonPrintableASCII)
	}
}

func TestReq_ConnectorIDNegative(t *testing.T) {
	t.Parallel()

	connectorId := testConnectorIDNeg

	_, err := rst.Req(rst.ReqInput{
		IDTag:       testValidIDTag,
		ConnectorID: &connectorId,
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

	connectorId := testConnectorIDOver

	_, err := rst.Req(rst.ReqInput{
		IDTag:       testValidIDTag,
		ConnectorID: &connectorId,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "connectorId exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}

func TestReq_MultipleErrors_EmptyIDTagAndInvalidConnectorID(t *testing.T) {
	t.Parallel()

	connectorId := testConnectorIDNeg

	_, err := rst.Req(rst.ReqInput{
		IDTag:       "",
		ConnectorID: &connectorId,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty idTag and negative connectorId")
	}

	// Should contain both errors
	if !strings.Contains(err.Error(), errIDTag) {
		t.Errorf(types.ErrorWantContains, err, errIDTag)
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}
