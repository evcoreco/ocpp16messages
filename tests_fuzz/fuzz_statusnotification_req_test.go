//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	sn "github.com/evcoreco/ocpp16messages/statusnotification"
	types "github.com/evcoreco/ocpp16types"
)

func FuzzStatusNotificationReq(f *testing.F) {
	f.Add(0, "NoError", "Available", false, "", false, "", false, "", false, "")
	f.Add(-1, "NoError", "Available", false, "", false, "", false, "", false, "")
	f.Add(0, "invalid-error", "Available", false, "", false, "", false, "", false, "")
	f.Add(0, "NoError", "invalid-status", false, "", false, "", false, "", false, "")
	f.Add(0, "NoError", "Available", true, "", false, "", false, "", false, "")
	f.Add(0, "NoError", "Available", false, "", true, "bad-time", false, "", false, "")

	f.Fuzz(func(
		t *testing.T,
		connectorId int,
		errorCode string,
		status string,
		hasInfo bool,
		info string,
		hasTimestamp bool,
		timestamp string,
		hasVendorID bool,
		vendorId string,
		hasVendorErrorCode bool,
		vendorErrorCode string,
	) {
		if len(errorCode) > maxFuzzStringLen ||
			len(status) > maxFuzzStringLen ||
			len(info) > maxFuzzStringLen ||
			len(timestamp) > maxFuzzStringLen ||
			len(vendorId) > maxFuzzStringLen ||
			len(vendorErrorCode) > maxFuzzStringLen {
			t.Skip()
		}

		var infoPtr *string
		if hasInfo {
			infoPtr = &info
		}

		var timestampPtr *string
		if hasTimestamp {
			timestampPtr = &timestamp
		}

		var vendorIdPtr *string
		if hasVendorID {
			vendorIdPtr = &vendorId
		}

		var vendorErrorCodePtr *string
		if hasVendorErrorCode {
			vendorErrorCodePtr = &vendorErrorCode
		}

		req, err := sn.Req(sn.ReqInput{
			ConnectorID:     connectorId,
			ErrorCode:       errorCode,
			Status:          status,
			Info:            infoPtr,
			Timestamp:       timestampPtr,
			VendorID:        vendorIdPtr,
			VendorErrorCode: vendorErrorCodePtr,
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

		if connectorId < 0 || connectorId > math.MaxUint16 {
			t.Fatalf("Req succeeded with connectorId=%d", connectorId)
		}

		if req.ConnectorID.Value() != uint16(connectorId) {
			t.Fatalf(
				"ConnectorID = %d, want %d",
				req.ConnectorID.Value(),
				connectorId,
			)
		}

		if !req.ErrorCode.IsValid() {
			t.Fatalf("ErrorCode = %q, want valid", req.ErrorCode.String())
		}

		if !req.Status.IsValid() {
			t.Fatalf("Status = %q, want valid", req.Status.String())
		}

		if hasInfo {
			if req.Info == nil {
				t.Fatal("Info = nil, want non-nil")
			}
			if info == "" {
				t.Fatal("Req succeeded with empty Info")
			}
			if req.Info.String() != info {
				t.Fatalf("Info = %q, want %q", req.Info.String(), info)
			}
		} else if req.Info != nil {
			t.Fatal("Info != nil, want nil")
		}

		if hasTimestamp {
			if req.Timestamp == nil {
				t.Fatal("Timestamp = nil, want non-nil")
			}
			if req.Timestamp.Value().Location() != time.UTC {
				t.Fatalf("Timestamp location = %v, want UTC", req.Timestamp.Value().Location())
			}
		} else if req.Timestamp != nil {
			t.Fatal("Timestamp != nil, want nil")
		}

		if hasVendorID {
			if req.VendorID == nil {
				t.Fatal("VendorID = nil, want non-nil")
			}
			if vendorId == "" {
				t.Fatal("Req succeeded with empty VendorID")
			}
			if req.VendorID.String() != vendorId {
				t.Fatalf("VendorID = %q, want %q", req.VendorID.String(), vendorId)
			}
		} else if req.VendorID != nil {
			t.Fatal("VendorID != nil, want nil")
		}

		if hasVendorErrorCode {
			if req.VendorErrorCode == nil {
				t.Fatal("VendorErrorCode = nil, want non-nil")
			}
			if vendorErrorCode == "" {
				t.Fatal("Req succeeded with empty VendorErrorCode")
			}
			if req.VendorErrorCode.String() != vendorErrorCode {
				t.Fatalf(
					"VendorErrorCode = %q, want %q",
					req.VendorErrorCode.String(),
					vendorErrorCode,
				)
			}
		} else if req.VendorErrorCode != nil {
			t.Fatal("VendorErrorCode != nil, want nil")
		}
	})
}
