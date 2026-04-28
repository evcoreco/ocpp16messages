//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	"github.com/evcoreco/ocpp16messages/cancelreservation"
	types "github.com/evcoreco/ocpp16types"
)

func FuzzCancelReservationReq(f *testing.F) {
	f.Add(1)
	f.Add(0)
	f.Add(-1)
	f.Add(math.MaxUint16 + 1)

	f.Fuzz(func(t *testing.T, reservationId int) {
		req, err := cancelreservation.Req(cancelreservation.ReqInput{
			ReservationID: reservationId,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if reservationId < 0 || reservationId > math.MaxUint16 {
			t.Fatalf("Req succeeded with reservationId=%d", reservationId)
		}

		if got := req.ReservationID.Value(); got != uint16(reservationId) {
			t.Fatalf("ReservationID = %d, want %d", got, reservationId)
		}
	})
}
