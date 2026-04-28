package authorize_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/evcoreco/ocpp16messages/authorize"
	types "github.com/evcoreco/ocpp16types"
)

const (
	StatusAccepted = "Accepted"
	ErrParentIDTag = "parentIDTag"
	ErrExpiryDate  = "expiryDate"
)

func TestConf_ValidAccepted(t *testing.T) {
	t.Parallel()

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      StatusAccepted,
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo.Status().String() != StatusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			StatusAccepted,
			conf.IDTagInfo.Status().String(),
		)
	}
}

func TestConf_ValidBlocked(t *testing.T) {
	t.Parallel()

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Blocked",
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	status := conf.IDTagInfo.Status().String()
	if status != "Blocked" {
		t.Errorf(types.ErrorMismatch, "Blocked", status)
	}
}

func TestConf_ValidExpired(t *testing.T) {
	t.Parallel()

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Expired",
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	status := conf.IDTagInfo.Status().String()
	if status != "Expired" {
		t.Errorf(types.ErrorMismatch, "Expired", status)
	}
}

func TestConf_ValidInvalid(t *testing.T) {
	t.Parallel()

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Invalid",
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	status := conf.IDTagInfo.Status().String()
	if status != "Invalid" {
		t.Errorf(types.ErrorMismatch, "Invalid", status)
	}
}

func TestConf_ValidConcurrentTx(t *testing.T) {
	t.Parallel()

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "ConcurrentTx",
		ExpiryDate:  nil,
		ParentIDTag: nil,
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

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Unknown",
		ExpiryDate:  nil,
		ParentIDTag: nil,
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

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "",
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty status")
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf("Conf() error = %v, want ErrInvalidValue", err)
	}
}

func TestConf_WithExpiryDate(t *testing.T) {
	t.Parallel()

	expiryDate := "2025-12-31T23:59:59Z"

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  &expiryDate,
		ParentIDTag: nil,
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

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  &invalidDate,
		ParentIDTag: nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid expiry date")
	}

	if !strings.Contains(err.Error(), "expiryDate") {
		t.Errorf("Conf() error = %v, want error containing 'expiryDate'", err)
	}
}

func TestConf_WithParentIDTag(t *testing.T) {
	t.Parallel()

	parentTag := "PARENT-TAG-123"

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  nil,
		ParentIDTag: &parentTag,
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

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  nil,
		ParentIDTag: &longTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for parentIDTag too long")
	}

	if !strings.Contains(err.Error(), ErrParentIDTag) {
		t.Errorf(types.ErrorWantContains, err, ErrParentIDTag)
	}
}

func TestConf_WithEmptyParentIDTag(t *testing.T) {
	t.Parallel()

	emptyTag := ""

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  nil,
		ParentIDTag: &emptyTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty parentIDTag")
	}

	if !strings.Contains(err.Error(), ErrParentIDTag) {
		t.Errorf(types.ErrorWantContains, err, ErrParentIDTag)
	}
}

func TestConf_Complete(t *testing.T) {
	t.Parallel()

	expiryDate := "2025-12-31T23:59:59Z"
	parentTag := "PARENT-123"

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  &expiryDate,
		ParentIDTag: &parentTag,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo.Status().String() != "Accepted" {
		t.Errorf(
			types.ErrorMismatch, "Accepted", conf.IDTagInfo.Status().String(),
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

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Invalid-Status",
		ExpiryDate:  &invalidDate,
		ParentIDTag: &longTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for multiple invalid fields")
	}

	// Check that all errors are present
	errStr := err.Error()
	if !strings.Contains(errStr, "status") {
		t.Errorf("Conf() error = %v, want error containing 'status'", err)
	}

	if !strings.Contains(errStr, ErrExpiryDate) {
		t.Errorf(types.ErrorWantContains, err, ErrExpiryDate)
	}

	if !strings.Contains(errStr, ErrParentIDTag) {
		t.Errorf(types.ErrorWantContains, err, ErrParentIDTag)
	}
}

func TestConf_MultipleErrors_StatusAndExpiryDate(t *testing.T) {
	t.Parallel()

	invalidDate := "invalid"

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "BadStatus",
		ExpiryDate:  &invalidDate,
		ParentIDTag: nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid status and expiry")
	}

	errStr := err.Error()
	if !strings.Contains(errStr, "status") {
		t.Errorf("Conf() error = %v, want error containing 'status'", err)
	}

	if !strings.Contains(errStr, ErrExpiryDate) {
		t.Errorf(types.ErrorWantContains, err, ErrExpiryDate)
	}
}
