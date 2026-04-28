//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	stt "github.com/evcoreco/ocpp16messages/starttransaction"
	types "github.com/evcoreco/ocpp16types"
)

func FuzzStartTransactionConf(f *testing.F) {
	f.Add(1, types.AuthorizationStatusAccepted.String(), false, "", false, "")
	f.Add(-1, types.AuthorizationStatusAccepted.String(), false, "", false, "")
	f.Add(1, "invalid-status", false, "", false, "")
	f.Add(1, types.AuthorizationStatusAccepted.String(), true, "bad-timestamp", false, "")
	f.Add(1, types.AuthorizationStatusAccepted.String(), false, "", true, "")

	f.Fuzz(func(
		t *testing.T,
		transactionId int,
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

		conf, err := stt.Conf(stt.ConfInput{
			TransactionID: transactionId,
			Status:        status,
			ExpiryDate:    expiryDatePtr,
			ParentIDTag:   parentIDTagPtr,
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

		if transactionId < 0 || transactionId > math.MaxUint16 {
			t.Fatalf("Conf succeeded with transactionId=%d", transactionId)
		}

		if got := conf.TransactionID.Value(); got != uint16(transactionId) {
			t.Fatalf("TransactionID = %d, want %d", got, transactionId)
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
		} else if conf.IDTagInfo.ExpiryDate() != nil {
			t.Fatal("ExpiryDate != nil, want nil")
		}

		if hasParentIDTag {
			if conf.IDTagInfo.ParentIDTag() == nil {
				t.Fatal("ParentIDTag = nil, want non-nil")
			}
		} else if conf.IDTagInfo.ParentIDTag() != nil {
			t.Fatal("ParentIDTag != nil, want nil")
		}
	})
}
