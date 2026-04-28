//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	rn "github.com/evcoreco/ocpp16messages/reservenow"
	types "github.com/evcoreco/ocpp16types"
)

func FuzzReserveNowReq(f *testing.F) {
	f.Add(1, 0, "RFID-ABC123", "2025-01-15T10:30:00Z", false, "")
	f.Add(-1, 0, "RFID-ABC123", "2025-01-15T10:30:00Z", false, "")
	f.Add(1, -1, "RFID-ABC123", "2025-01-15T10:30:00Z", false, "")
	f.Add(1, 0, "", "2025-01-15T10:30:00Z", false, "")
	f.Add(1, 0, "RFID-ABC123", "bad-time", false, "")
	f.Add(1, 0, "RFID-ABC123", "2025-01-15T10:30:00Z", true, "")

	f.Fuzz(func(
		t *testing.T,
		reservationId int,
		connectorId int,
		idTag string,
		expiryDate string,
		hasParentIDTag bool,
		parentIDTag string,
	) {
		if len(idTag) > maxFuzzStringLen ||
			len(expiryDate) > maxFuzzStringLen ||
			len(parentIDTag) > maxFuzzStringLen {
			t.Skip()
		}

		var parentIDTagPtr *string
		if hasParentIDTag {
			parentIDTagPtr = &parentIDTag
		}

		req, err := rn.Req(rn.ReqInput{
			ReservationID: reservationId,
			ConnectorID:   connectorId,
			IDTag:         idTag,
			ExpiryDate:    expiryDate,
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

		if reservationId < 0 || reservationId > math.MaxUint16 {
			t.Fatalf("Req succeeded with reservationId=%d", reservationId)
		}

		if connectorId < 0 || connectorId > math.MaxUint16 {
			t.Fatalf("Req succeeded with connectorId=%d", connectorId)
		}

		if idTag == "" {
			t.Fatal("Req succeeded with empty IDTag")
		}

		if req.ReservationID.Value() != uint16(reservationId) {
			t.Fatalf(
				"ReservationID = %d, want %d",
				req.ReservationID.Value(),
				reservationId,
			)
		}

		if req.ConnectorID.Value() != uint16(connectorId) {
			t.Fatalf("ConnectorID = %d, want %d", req.ConnectorID.Value(), connectorId)
		}

		if req.IDTag.String() != idTag {
			t.Fatalf("IDTag = %q, want %q", req.IDTag.String(), idTag)
		}

		if req.ExpiryDate.Value().Location() != time.UTC {
			t.Fatalf(
				"ExpiryDate location = %v, want UTC",
				req.ExpiryDate.Value().Location(),
			)
		}

		if hasParentIDTag {
			if req.ParentIDTag == nil {
				t.Fatal("ParentIDTag = nil, want non-nil")
			}
			if parentIDTag == "" {
				t.Fatal("Req succeeded with empty ParentIDTag")
			}
			if req.ParentIDTag.String() != parentIDTag {
				t.Fatalf(
					"ParentIDTag = %q, want %q",
					req.ParentIDTag.String(),
					parentIDTag,
				)
			}
		} else if req.ParentIDTag != nil {
			t.Fatal("ParentIDTag != nil, want nil")
		}
	})
}
