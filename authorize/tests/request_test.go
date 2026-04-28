package authorize_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/evcoreco/ocpp16messages/authorize"
	types "github.com/evcoreco/ocpp16types"
)

const testValidIDTag = "RFID-TAG-12345"

func TestReq_Valid(t *testing.T) {
	t.Parallel()

	req, err := authorize.Req(authorize.ReqInput{IDTag: testValidIDTag})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.IDTag.String() != testValidIDTag {
		t.Errorf(types.ErrorMismatch, testValidIDTag, req.IDTag.String())
	}
}

func TestReq_EmptyIDTag(t *testing.T) {
	t.Parallel()

	_, err := authorize.Req(authorize.ReqInput{IDTag: ""})
	if err == nil {
		t.Error("Req() error = nil, want error for empty idTag")
	}

	if !errors.Is(err, types.ErrEmptyValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrEmptyValue)
	}
}

func TestReq_IDTagTooLong(t *testing.T) {
	t.Parallel()

	// 23 chars, max is 20
	_, err := authorize.Req(authorize.ReqInput{
		IDTag: "RFID-ABC123456789012345",
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

func TestReq_InvalidCharacters(t *testing.T) {
	t.Parallel()

	// Contains null byte
	_, err := authorize.Req(authorize.ReqInput{IDTag: "RFID\x00ABC"})
	if err == nil {
		t.Error(
			"Req() error = nil, want error for non-printable chars",
		)
	}

	if !strings.Contains(err.Error(), "non-printable ASCII") {
		t.Errorf(
			"Req() error = %v, want 'non-printable ASCII'",
			err,
		)
	}
}
