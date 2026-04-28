//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"time"

	stp "github.com/evcoreco/ocpp16messages/stoptransaction"
	types "github.com/evcoreco/ocpp16types"
)

func FuzzStopTransactionConf(f *testing.F) {
	f.Add(false, "", false, "", false, "")
	f.Add(true, types.AuthorizationStatusAccepted.String(), false, "", false, "")
	f.Add(true, "invalid-status", false, "", false, "")
	f.Add(false, "", true, "bad-timestamp", false, "")
	f.Add(false, "", false, "", true, "bad\x01")

	f.Fuzz(func(
		t *testing.T,
		hasStatus bool,
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

		var statusPtr *string
		if hasStatus {
			statusPtr = &status
		}

		var expiryDatePtr *string
		if hasExpiryDate {
			expiryDatePtr = &expiryDate
		}

		var parentIDTagPtr *string
		if hasParentIDTag {
			parentIDTagPtr = &parentIDTag
		}

		conf, err := stp.Conf(stp.ConfInput{
			Status:      statusPtr,
			ExpiryDate:  expiryDatePtr,
			ParentIDTag: parentIDTagPtr,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) && !errors.Is(err, types.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if !hasStatus {
			if conf.IDTagInfo != nil {
				t.Fatal("IDTagInfo != nil, want nil when Status is omitted")
			}

			return
		}

		if conf.IDTagInfo == nil {
			t.Fatal("IDTagInfo = nil, want non-nil when Status is provided")
		}

		if !conf.IDTagInfo.Status().IsValid() {
			t.Fatalf("Status = %q, want valid", conf.IDTagInfo.Status().String())
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
		}

		if hasParentIDTag {
			if conf.IDTagInfo.ParentIDTag() == nil {
				t.Fatal("ParentIDTag = nil, want non-nil")
			}
		}
	})
}
