package starttransaction_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/evcoreco/ocpp16messages/starttransaction"
	types "github.com/evcoreco/ocpp16types"
)

const (
	statusAccepted     = "Accepted"
	errParentIDTag     = "parentIDTag"
	errExpiryDate      = "expiryDate"
	errTransactionID   = "transactionId"
	testTransactionID  = 12345
	testTransactionID2 = 12346
)

func TestConf_ValidAccepted(t *testing.T) {
	t.Parallel()

	conf, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID,
		Status:        statusAccepted,
		ExpiryDate:    nil,
		ParentIDTag:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	expectedTransactionID := uint16(testTransactionID)
	if conf.TransactionID.Value() != expectedTransactionID {
		t.Errorf(
			types.ErrorMismatch,
			expectedTransactionID,
			conf.TransactionID.Value(),
		)
	}

	if conf.IDTagInfo.Status().String() != statusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			statusAccepted,
			conf.IDTagInfo.Status().String(),
		)
	}
}

func TestConf_ValidBlocked(t *testing.T) {
	t.Parallel()

	conf, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID,
		Status:        "Blocked",
		ExpiryDate:    nil,
		ParentIDTag:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo.Status().String() != "Blocked" {
		t.Errorf(
			types.ErrorMismatch,
			"Blocked",
			conf.IDTagInfo.Status().String(),
		)
	}
}

func TestConf_ValidExpired(t *testing.T) {
	t.Parallel()

	conf, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID,
		Status:        "Expired",
		ExpiryDate:    nil,
		ParentIDTag:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo.Status().String() != "Expired" {
		t.Errorf(
			types.ErrorMismatch,
			"Expired",
			conf.IDTagInfo.Status().String(),
		)
	}
}

func TestConf_ValidInvalid(t *testing.T) {
	t.Parallel()

	conf, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID,
		Status:        "Invalid",
		ExpiryDate:    nil,
		ParentIDTag:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo.Status().String() != "Invalid" {
		t.Errorf(
			types.ErrorMismatch,
			"Invalid",
			conf.IDTagInfo.Status().String(),
		)
	}
}

func TestConf_ValidConcurrentTx(t *testing.T) {
	t.Parallel()

	conf, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID,
		Status:        "ConcurrentTx",
		ExpiryDate:    nil,
		ParentIDTag:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo.Status().String() != "ConcurrentTx" {
		t.Errorf(
			types.ErrorMismatch,
			"ConcurrentTx",
			conf.IDTagInfo.Status().String(),
		)
	}
}

func TestConf_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID,
		Status:        "Unknown",
		ExpiryDate:    nil,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf("Conf() error = %v, want ErrInvalidValue", err)
	}
}

func TestConf_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID,
		Status:        "",
		ExpiryDate:    nil,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf("Conf() error = %v, want ErrInvalidValue", err)
	}
}

func TestConf_TransactionIDNegative(t *testing.T) {
	t.Parallel()

	_, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testValueNegativeOne,
		Status:        statusAccepted,
		ExpiryDate:    nil,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for negative transactionId")
	}

	if !strings.Contains(err.Error(), errTransactionID) {
		t.Errorf(types.ErrorWantContains, err, errTransactionID)
	}
}

func TestConf_WithExpiryDate(t *testing.T) {
	t.Parallel()

	expiryDate := "2025-12-31T23:59:59Z"

	conf, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID,
		Status:        statusAccepted,
		ExpiryDate:    &expiryDate,
		ParentIDTag:   nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo.ExpiryDate() == nil {
		t.Error("Conf() ExpiryDate = nil, want non-nil")
	}
}

func TestConf_WithInvalidExpiryDate(t *testing.T) {
	t.Parallel()

	invalidDate := "not-a-date"

	_, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID,
		Status:        statusAccepted,
		ExpiryDate:    &invalidDate,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid expiry date")
	}

	if !strings.Contains(err.Error(), errExpiryDate) {
		t.Errorf(types.ErrorWantContains, err, errExpiryDate)
	}
}

func TestConf_WithParentIDTag(t *testing.T) {
	t.Parallel()

	parentTag := "PARENT-TAG-123"

	conf, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID,
		Status:        statusAccepted,
		ExpiryDate:    nil,
		ParentIDTag:   &parentTag,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo.ParentIDTag() == nil {
		t.Error("Conf() ParentIDTag = nil, want non-nil")
	}

	if conf.IDTagInfo.ParentIDTag().String() != parentTag {
		t.Errorf(
			types.ErrorMismatch,
			parentTag,
			conf.IDTagInfo.ParentIDTag().String(),
		)
	}
}

func TestConf_WithParentIDTagTooLong(t *testing.T) {
	t.Parallel()

	longTag := "PARENT-TAG-123456789012345" // 26 chars, max is 20

	_, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID,
		Status:        statusAccepted,
		ExpiryDate:    nil,
		ParentIDTag:   &longTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for parentIDTag too long")
	}

	if !strings.Contains(err.Error(), errParentIDTag) {
		t.Errorf(types.ErrorWantContains, err, errParentIDTag)
	}
}

func TestConf_WithEmptyParentIDTag(t *testing.T) {
	t.Parallel()

	emptyTag := ""

	_, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID,
		Status:        statusAccepted,
		ExpiryDate:    nil,
		ParentIDTag:   &emptyTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty parentIDTag")
	}

	if !strings.Contains(err.Error(), errParentIDTag) {
		t.Errorf(types.ErrorWantContains, err, errParentIDTag)
	}
}

func TestConf_Complete(t *testing.T) {
	t.Parallel()

	expiryDate := "2025-12-31T23:59:59Z"
	parentTag := "PARENT-123"

	conf, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID2,
		Status:        statusAccepted,
		ExpiryDate:    &expiryDate,
		ParentIDTag:   &parentTag,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	expectedTransactionID := uint16(testTransactionID2)
	if conf.TransactionID.Value() != expectedTransactionID {
		t.Errorf(
			types.ErrorMismatch,
			expectedTransactionID,
			conf.TransactionID.Value(),
		)
	}

	if conf.IDTagInfo.Status().String() != statusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			statusAccepted,
			conf.IDTagInfo.Status().String(),
		)
	}

	if conf.IDTagInfo.ExpiryDate() == nil {
		t.Error("Conf() ExpiryDate = nil, want non-nil")
	}

	if conf.IDTagInfo.ParentIDTag() == nil {
		t.Error("Conf() ParentIDTag = nil, want non-nil")
	}

	if conf.IDTagInfo.ParentIDTag().String() != parentTag {
		t.Errorf(
			types.ErrorMismatch,
			parentTag,
			conf.IDTagInfo.ParentIDTag().String(),
		)
	}
}

func TestConf_MultipleErrors(t *testing.T) {
	t.Parallel()

	invalidDate := "not-a-date"
	longTag := "THIS-TAG-IS-WAY-TOO-LONG-FOR-OCPP"

	_, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testValueNegativeOne,
		Status:        "Invalid-Status",
		ExpiryDate:    &invalidDate,
		ParentIDTag:   &longTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errTransactionID) {
		t.Errorf(types.ErrorWantContains, err, errTransactionID)
	}

	if !strings.Contains(errStr, "status") {
		t.Errorf("Conf() error = %v, want error containing 'status'", err)
	}

	if !strings.Contains(errStr, errExpiryDate) {
		t.Errorf(types.ErrorWantContains, err, errExpiryDate)
	}

	if !strings.Contains(errStr, errParentIDTag) {
		t.Errorf(types.ErrorWantContains, err, errParentIDTag)
	}
}

func TestConf_MultipleErrors_StatusAndExpiryDate(t *testing.T) {
	t.Parallel()

	invalidDate := "invalid"

	_, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: testTransactionID,
		Status:        "BadStatus",
		ExpiryDate:    &invalidDate,
		ParentIDTag:   nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid status and expiry")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, "status") {
		t.Errorf("Conf() error = %v, want error containing 'status'", err)
	}

	if !strings.Contains(errStr, errExpiryDate) {
		t.Errorf(types.ErrorWantContains, err, errExpiryDate)
	}
}
