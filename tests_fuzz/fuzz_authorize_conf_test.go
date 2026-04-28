//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"time"

	"github.com/evcoreco/ocpp16messages/authorize"
	types "github.com/evcoreco/ocpp16types"
)

func FuzzAuthorizeConf(f *testing.F) {
	f.Add(types.AuthorizationStatusAccepted.String(), false, "", false, "")
	f.Add(types.AuthorizationStatusAccepted.String(), true, "2025-01-15T10:30:00Z", false, "")
	f.Add(types.AuthorizationStatusAccepted.String(), false, "", true, "PARENT-123")
	f.Add(types.AuthorizationStatusInvalid.String(), false, "", false, "")
	f.Add(types.AuthorizationStatusAccepted.String(), true, "bad-time", false, "")
	f.Add(types.AuthorizationStatusAccepted.String(), false, "", true, "")

	f.Fuzz(func(
		t *testing.T,
		status string,
		hasExpiryDate bool,
		expiryDate string,
		hasParentIDTag bool,
		parentIDTag string,
	) {
		if len(status) > maxFuzzStringLen ||
			len(expiryDate) > maxFuzzStringLen ||
			len(parentIDTag) > maxFuzzStringLen {
			t.Skip()
		}

		var expiryDatePtr *string
		if hasExpiryDate {
			expiryDatePtr = &expiryDate
		}

		var parentIDTagPtr *string
		if hasParentIDTag {
			parentIDTagPtr = &parentIDTag
		}

		conf, err := authorize.Conf(authorize.ConfInput{
			Status:      status,
			ExpiryDate:  expiryDatePtr,
			ParentIDTag: parentIDTagPtr,
		})
		if err != nil {
			if !errors.Is(err, types.ErrEmptyValue) && !errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if !conf.IDTagInfo.Status().IsValid() {
			t.Fatalf("Status = %q, want valid", conf.IDTagInfo.Status().String())
		}
		if conf.IDTagInfo.Status().String() != status {
			t.Fatalf("Status = %q, want %q", conf.IDTagInfo.Status().String(), status)
		}

		if hasExpiryDate {
			if conf.IDTagInfo.ExpiryDate() == nil {
				t.Fatal("ExpiryDate = nil, want non-nil")
			}
			if conf.IDTagInfo.ExpiryDate().Value().Location() != time.UTC {
				t.Fatalf(
					"ExpiryDate location = %v, want UTC",
					conf.IDTagInfo.ExpiryDate().Value().Location(),
				)
			}
		} else if conf.IDTagInfo.ExpiryDate() != nil {
			t.Fatal("ExpiryDate != nil, want nil")
		}

		if hasParentIDTag {
			if conf.IDTagInfo.ParentIDTag() == nil {
				t.Fatal("ParentIDTag = nil, want non-nil")
			}
			if parentIDTag == "" {
				t.Fatal("Conf succeeded with empty ParentIDTag")
			}
			if conf.IDTagInfo.ParentIDTag().String() != parentIDTag {
				t.Fatalf(
					"ParentIDTag = %q, want %q",
					conf.IDTagInfo.ParentIDTag().String(),
					parentIDTag,
				)
			}
		} else if conf.IDTagInfo.ParentIDTag() != nil {
			t.Fatal("ParentIDTag != nil, want nil")
		}
	})
}
