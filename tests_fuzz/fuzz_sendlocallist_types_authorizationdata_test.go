//go:build fuzz

package fuzz

import (
	"errors"
	"testing"
	"time"

	types "github.com/evcoreco/ocpp16types"
)

func FuzzNewAuthorizationData(f *testing.F) {
	f.Add("", false, "", false, "", false, "")
	f.Add("RFID-ABC123", false, "", false, "", false, "")
	f.Add(
		"RFID-ABC123",
		true,
		types.AuthorizationStatusAccepted.String(),
		true,
		"2025-01-15T10:30:00Z",
		true,
		"RFID-PARENT",
	)
	f.Add("RFID-ABC123", true, "invalid-status", false, "", false, "")
	f.Add("bad\x01", false, "", false, "", false, "")

	f.Fuzz(func(
		t *testing.T,
		idTag string,
		hasIDTagInfo bool,
		status string,
		hasExpiryDate bool,
		expiryDate string,
		hasParentIDTag bool,
		parentIDTag string,
	) {
		if len(idTag) > maxFuzzStringLen ||
			len(status) > maxFuzzStringLen ||
			len(expiryDate) > maxFuzzStringLen ||
			len(parentIDTag) > maxFuzzStringLen {
			t.Skip()
		}

		var idTagInfoPtr *types.IDTagInfoInput
		if hasIDTagInfo {
			var expiryDatePtr *string
			if hasExpiryDate {
				expiryDatePtr = &expiryDate
			}

			var parentIDTagPtr *string
			if hasParentIDTag {
				parentIDTagPtr = &parentIDTag
			}

			idTagInfoPtr = &types.IDTagInfoInput{
				Status:      status,
				ExpiryDate:  expiryDatePtr,
				ParentIDTag: parentIDTagPtr,
			}
		}

		authData, err := types.NewAuthorizationData(types.AuthorizationDataInput{
			IDTag:     idTag,
			IDTagInfo: idTagInfoPtr,
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

		if idTag == "" {
			t.Fatal("NewAuthorizationData succeeded with empty IDTag")
		}

		if authData.IDTag().String() != idTag {
			t.Fatalf("IDTag = %q, want %q", authData.IDTag().String(), idTag)
		}

		if !hasIDTagInfo {
			if authData.IDTagInfo() != nil {
				t.Fatal("IDTagInfo != nil, want nil")
			}

			return
		}

		if authData.IDTagInfo() == nil {
			t.Fatal("IDTagInfo = nil, want non-nil")
		}

		if !authData.IDTagInfo().Status().IsValid() {
			t.Fatalf("Status = %q, want valid", authData.IDTagInfo().Status().String())
		}

		if hasExpiryDate {
			if authData.IDTagInfo().ExpiryDate() == nil {
				t.Fatal("ExpiryDate = nil, want non-nil")
			}
			if authData.IDTagInfo().ExpiryDate().Value().Location() != time.UTC {
				t.Fatalf(
					"ExpiryDate location = %v, want UTC",
					authData.IDTagInfo().ExpiryDate().Value().Location(),
				)
			}
		} else if authData.IDTagInfo().ExpiryDate() != nil {
			t.Fatal("ExpiryDate != nil, want nil")
		}

		if hasParentIDTag {
			if authData.IDTagInfo().ParentIDTag() == nil {
				t.Fatal("ParentIDTag = nil, want non-nil")
			}
			if authData.IDTagInfo().ParentIDTag().String() != parentIDTag {
				t.Fatalf(
					"ParentIDTag = %q, want %q",
					authData.IDTagInfo().ParentIDTag().String(),
					parentIDTag,
				)
			}
		} else if authData.IDTagInfo().ParentIDTag() != nil {
			t.Fatal("ParentIDTag != nil, want nil")
		}
	})
}
