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

func FuzzStartTransactionReq(f *testing.F) {
	f.Add(0, "RFID-ABC123", 0, "2025-01-15T10:30:00Z", false, 0)
	f.Add(-1, "RFID-ABC123", 0, "2025-01-15T10:30:00Z", false, 0)
	f.Add(0, "", 0, "2025-01-15T10:30:00Z", false, 0)
	f.Add(0, "RFID-ABC123", -1, "2025-01-15T10:30:00Z", false, 0)
	f.Add(0, "RFID-ABC123", 0, "bad-timestamp", false, 0)
	f.Add(0, "RFID-ABC123", 0, "2025-01-15T10:30:00Z", true, -1)

	f.Fuzz(func(
		t *testing.T,
		connectorId int,
		idTag string,
		meterStart int,
		timestamp string,
		hasReservationID bool,
		reservationId int,
	) {
		if len(idTag) > maxFuzzStringLen || len(timestamp) > maxFuzzStringLen {
			t.Skip()
		}

		var reservationIdPtr *int
		if hasReservationID {
			reservationIdPtr = &reservationId
		}

		req, err := stt.Req(stt.ReqInput{
			ConnectorID:   connectorId,
			IDTag:         idTag,
			MeterStart:    meterStart,
			Timestamp:     timestamp,
			ReservationID: reservationIdPtr,
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

		if got := req.ConnectorID.Value(); got != uint16(connectorId) {
			t.Fatalf("ConnectorID = %d, want %d", got, connectorId)
		}

		if idTag == "" {
			t.Fatal("Req succeeded with empty IDTag")
		}

		if req.IDTag.String() != idTag {
			t.Fatalf("IDTag = %q, want %q", req.IDTag.String(), idTag)
		}

		if meterStart < 0 || meterStart > math.MaxUint16 {
			t.Fatalf("Req succeeded with meterStart=%d", meterStart)
		}

		if got := req.MeterStart.Value(); got != uint16(meterStart) {
			t.Fatalf("MeterStart = %d, want %d", got, meterStart)
		}

		if req.Timestamp.Value().Location() != time.UTC {
			t.Fatalf("Timestamp location = %v, want UTC", req.Timestamp.Value().Location())
		}

		if hasReservationID {
			if req.ReservationID == nil {
				t.Fatal("ReservationID = nil, want non-nil")
			}

			if reservationId < 0 || reservationId > math.MaxUint16 {
				t.Fatalf("Req succeeded with reservationId=%d", reservationId)
			}

			if got := req.ReservationID.Value(); got != uint16(reservationId) {
				t.Fatalf("ReservationID = %d, want %d", got, reservationId)
			}
		} else if req.ReservationID != nil {
			t.Fatal("ReservationID != nil, want nil")
		}
	})
}
