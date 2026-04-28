package stoptransaction_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/evcoreco/ocpp16messages/stoptransaction"
	types "github.com/evcoreco/ocpp16types"
)

const (
	statusAccepted       = "Accepted"
	statusBlocked        = "Blocked"
	statusExpired        = "Expired"
	statusInvalid        = "Invalid"
	statusConcurrentTx   = "ConcurrentTx"
	errConfStatus        = "status"
	errConfExpiryDate    = "expiryDate"
	errConfParentId      = "parentIDTag"
	errConfIDTagInfoNil  = "Conf() IDTagInfo = nil, want non-nil"
	errConfExpiryDateNil = "Conf() ExpiryDate = nil, want non-nil"
	errConfParentIdNil   = "Conf() ParentIDTag = nil, want non-nil"
)

func TestConf_NoIDTagInfo(t *testing.T) {
	t.Parallel()

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      nil,
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo != nil {
		t.Error("Conf() IDTagInfo != nil, want nil for empty input")
	}
}

func TestConf_ValidAccepted(t *testing.T) {
	t.Parallel()

	status := statusAccepted

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo == nil {
		t.Error(errConfIDTagInfoNil)

		return
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

	status := statusBlocked

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo == nil {
		t.Error(errConfIDTagInfoNil)

		return
	}

	if conf.IDTagInfo.Status().String() != statusBlocked {
		t.Errorf(
			types.ErrorMismatch,
			statusBlocked,
			conf.IDTagInfo.Status().String(),
		)
	}
}

func TestConf_ValidExpired(t *testing.T) {
	t.Parallel()

	status := statusExpired

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo == nil {
		t.Error(errConfIDTagInfoNil)

		return
	}

	if conf.IDTagInfo.Status().String() != statusExpired {
		t.Errorf(
			types.ErrorMismatch,
			statusExpired,
			conf.IDTagInfo.Status().String(),
		)
	}
}

func TestConf_ValidInvalid(t *testing.T) {
	t.Parallel()

	status := statusInvalid

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo == nil {
		t.Error(errConfIDTagInfoNil)

		return
	}

	if conf.IDTagInfo.Status().String() != statusInvalid {
		t.Errorf(
			types.ErrorMismatch,
			statusInvalid,
			conf.IDTagInfo.Status().String(),
		)
	}
}

func TestConf_ValidConcurrentTx(t *testing.T) {
	t.Parallel()

	status := statusConcurrentTx

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo == nil {
		t.Error(errConfIDTagInfoNil)

		return
	}

	if conf.IDTagInfo.Status().String() != statusConcurrentTx {
		t.Errorf(
			types.ErrorMismatch,
			statusConcurrentTx,
			conf.IDTagInfo.Status().String(),
		)
	}
}

func TestConf_InvalidStatus(t *testing.T) {
	t.Parallel()

	status := "Unknown"

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
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

	status := ""

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
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

	status := statusAccepted
	expiryDate := "2025-12-31T23:59:59Z"

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  &expiryDate,
		ParentIDTag: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo == nil {
		t.Error(errConfIDTagInfoNil)

		return
	}

	if conf.IDTagInfo.ExpiryDate() == nil {
		t.Error(errConfExpiryDateNil)
	}
}

func TestConf_WithInvalidExpiryDate(t *testing.T) {
	t.Parallel()

	status := statusAccepted
	invalidDate := "not-a-date"

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  &invalidDate,
		ParentIDTag: nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid expiry date")
	}

	if !strings.Contains(err.Error(), errConfExpiryDate) {
		t.Errorf(types.ErrorWantContains, err, errConfExpiryDate)
	}
}

func TestConf_WithParentIDTag(t *testing.T) {
	t.Parallel()

	status := statusAccepted
	parentTag := "PARENT-TAG-123"

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIDTag: &parentTag,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo == nil {
		t.Error(errConfIDTagInfoNil)

		return
	}

	if conf.IDTagInfo.ParentIDTag() == nil {
		t.Error(errConfParentIdNil)
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

	status := statusAccepted
	longTag := "PARENT-TAG-123456789012345" // 26 chars, max is 20

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIDTag: &longTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for parentIDTag too long")
	}

	if !strings.Contains(err.Error(), errConfParentId) {
		t.Errorf(types.ErrorWantContains, err, errConfParentId)
	}
}

func TestConf_WithEmptyParentIDTag(t *testing.T) {
	t.Parallel()

	status := statusAccepted
	emptyTag := ""

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIDTag: &emptyTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for empty parentIDTag")
	}

	if !strings.Contains(err.Error(), errConfParentId) {
		t.Errorf(types.ErrorWantContains, err, errConfParentId)
	}
}

func TestConf_Complete(t *testing.T) {
	t.Parallel()

	status := statusAccepted
	expiryDate := "2025-12-31T23:59:59Z"
	parentTag := "PARENT-123"

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  &expiryDate,
		ParentIDTag: &parentTag,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.IDTagInfo == nil {
		t.Error(errConfIDTagInfoNil)

		return
	}

	if conf.IDTagInfo.Status().String() != statusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			statusAccepted,
			conf.IDTagInfo.Status().String(),
		)
	}

	if conf.IDTagInfo.ExpiryDate() == nil {
		t.Error(errConfExpiryDateNil)
	}

	if conf.IDTagInfo.ParentIDTag() == nil {
		t.Error(errConfParentIdNil)
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

	invalidStatus := "Invalid-Status"
	invalidDate := "not-a-date"
	longTag := "THIS-TAG-IS-WAY-TOO-LONG-FOR-OCPP"

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &invalidStatus,
		ExpiryDate:  &invalidDate,
		ParentIDTag: &longTag,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errConfStatus) {
		t.Errorf(types.ErrorWantContains, err, errConfStatus)
	}

	if !strings.Contains(errStr, errConfExpiryDate) {
		t.Errorf(types.ErrorWantContains, err, errConfExpiryDate)
	}

	if !strings.Contains(errStr, errConfParentId) {
		t.Errorf(types.ErrorWantContains, err, errConfParentId)
	}
}

func TestConf_MultipleErrors_StatusAndExpiryDate(t *testing.T) {
	t.Parallel()

	invalidStatus := "BadStatus"
	invalidDate := "invalid"

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &invalidStatus,
		ExpiryDate:  &invalidDate,
		ParentIDTag: nil,
	})
	if err == nil {
		t.Error("Conf() error = nil, want error for invalid status and expiry")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errConfStatus) {
		t.Errorf(types.ErrorWantContains, err, errConfStatus)
	}

	if !strings.Contains(errStr, errConfExpiryDate) {
		t.Errorf(types.ErrorWantContains, err, errConfExpiryDate)
	}
}
