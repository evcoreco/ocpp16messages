package datatransfer_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/evcoreco/ocpp16messages/datatransfer"
	types "github.com/evcoreco/ocpp16types"
)

const (
	testValidVendorID   = "com.example.vendor"
	testValidMessageID  = "CustomMessage"
	testValidData       = `{"key": "value"}`
	errFieldVendorID    = "vendorId"
	errFieldMessageID   = "messageId"
	errExceedsMaxLength = "exceeds maximum length"
	vendorIdMaxPlusOne  = 256
	messageIdMaxPlusOne = 51
)

func TestReq_Valid_VendorIDOnly(t *testing.T) {
	t.Parallel()

	req, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  testValidVendorID,
		MessageID: nil,
		Data:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.VendorID.Value() != testValidVendorID {
		t.Errorf(types.ErrorMismatch, testValidVendorID, req.VendorID.Value())
	}

	if req.MessageID != nil {
		t.Error("Req() MessageID != nil, want nil")
	}

	if req.Data != nil {
		t.Error("Req() Data != nil, want nil")
	}
}

func TestReq_Valid_WithMessageID(t *testing.T) {
	t.Parallel()

	messageId := testValidMessageID

	req, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  testValidVendorID,
		MessageID: &messageId,
		Data:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.MessageID == nil {
		t.Error("Req() MessageID = nil, want non-nil")

		return
	}

	if req.MessageID.Value() != testValidMessageID {
		t.Errorf(types.ErrorMismatch, testValidMessageID, req.MessageID.Value())
	}
}

func TestReq_Valid_WithData(t *testing.T) {
	t.Parallel()

	data := testValidData

	req, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  testValidVendorID,
		MessageID: nil,
		Data:      &data,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Data == nil {
		t.Error("Req() Data = nil, want non-nil")

		return
	}

	if *req.Data != testValidData {
		t.Errorf(types.ErrorMismatch, testValidData, *req.Data)
	}
}

func TestReq_Valid_Complete(t *testing.T) {
	t.Parallel()

	messageId := testValidMessageID
	data := testValidData

	req, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  testValidVendorID,
		MessageID: &messageId,
		Data:      &data,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.VendorID.Value() != testValidVendorID {
		t.Errorf(types.ErrorMismatch, testValidVendorID, req.VendorID.Value())
	}

	if req.MessageID == nil {
		t.Error("Req() MessageID = nil, want non-nil")

		return
	}

	if req.MessageID.Value() != testValidMessageID {
		t.Errorf(types.ErrorMismatch, testValidMessageID, req.MessageID.Value())
	}

	if req.Data == nil {
		t.Error("Req() Data = nil, want non-nil")

		return
	}

	if *req.Data != testValidData {
		t.Errorf(types.ErrorMismatch, testValidData, *req.Data)
	}
}

func TestReq_EmptyVendorID(t *testing.T) {
	t.Parallel()

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  "",
		MessageID: nil,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty vendorId")
	}

	if !errors.Is(err, types.ErrEmptyValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrEmptyValue)
	}
}

func TestReq_VendorIDTooLong(t *testing.T) {
	t.Parallel()

	// Create a string longer than 255 characters
	longVendorID := strings.Repeat("a", vendorIdMaxPlusOne)

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  longVendorID,
		MessageID: nil,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for vendorId too long")
	}

	if !strings.Contains(err.Error(), errFieldVendorID) {
		t.Errorf(types.ErrorWantContains, err, errFieldVendorID)
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(types.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestReq_VendorIDInvalidChars(t *testing.T) {
	t.Parallel()

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  "vendor\x00id",
		MessageID: nil,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid chars in vendorId")
	}

	if !strings.Contains(err.Error(), errFieldVendorID) {
		t.Errorf(types.ErrorWantContains, err, errFieldVendorID)
	}
}

func TestReq_EmptyMessageID(t *testing.T) {
	t.Parallel()

	emptyMessageID := ""

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  testValidVendorID,
		MessageID: &emptyMessageID,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty messageId")
	}

	if !strings.Contains(err.Error(), errFieldMessageID) {
		t.Errorf(types.ErrorWantContains, err, errFieldMessageID)
	}
}

func TestReq_MessageIDTooLong(t *testing.T) {
	t.Parallel()

	// Create a string longer than 50 characters
	longMessageID := strings.Repeat("m", messageIdMaxPlusOne)

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  testValidVendorID,
		MessageID: &longMessageID,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for messageId too long")
	}

	if !strings.Contains(err.Error(), errFieldMessageID) {
		t.Errorf(types.ErrorWantContains, err, errFieldMessageID)
	}

	if !strings.Contains(err.Error(), errExceedsMaxLength) {
		t.Errorf(types.ErrorWantContains, err, errExceedsMaxLength)
	}
}

func TestReq_MessageIDInvalidChars(t *testing.T) {
	t.Parallel()

	invalidMessageID := "msg\x00id"

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  testValidVendorID,
		MessageID: &invalidMessageID,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid chars in messageId")
	}

	if !strings.Contains(err.Error(), errFieldMessageID) {
		t.Errorf(types.ErrorWantContains, err, errFieldMessageID)
	}
}

func TestReq_MultipleErrors(t *testing.T) {
	t.Parallel()

	invalidMessageID := ""

	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  "",
		MessageID: &invalidMessageID,
		Data:      nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()
	if !strings.Contains(errStr, errFieldVendorID) {
		t.Errorf(types.ErrorWantContains, err, errFieldVendorID)
	}

	if !strings.Contains(errStr, errFieldMessageID) {
		t.Errorf(types.ErrorWantContains, err, errFieldMessageID)
	}
}
