package cancelreservation_test

import (
	"strings"
	"testing"

	cr "github.com/evcoreco/ocpp16messages/cancelreservation"
	types "github.com/evcoreco/ocpp16types"
)

const (
	errReservationID = "reservationId"

	valueZero         = 0
	valuePositive     = 123
	valueMaxUint16    = 65535
	valueExceedsMax   = 65536
	valueNegative     = -1
	valueLargeNegativ = -65536
)

func TestReq_Valid_Zero(t *testing.T) {
	t.Parallel()

	req, err := cr.Req(cr.ReqInput{ReservationID: valueZero})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ReservationID.Value() != valueZero {
		t.Errorf(types.ErrorMismatchValue, valueZero, req.ReservationID.Value())
	}
}

func TestReq_Valid_PositiveValue(t *testing.T) {
	t.Parallel()

	req, err := cr.Req(cr.ReqInput{ReservationID: valuePositive})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ReservationID.Value() != valuePositive {
		t.Errorf(
			types.ErrorMismatchValue,
			valuePositive,
			req.ReservationID.Value(),
		)
	}
}

func TestReq_Valid_MaxValue(t *testing.T) {
	t.Parallel()

	req, err := cr.Req(cr.ReqInput{ReservationID: valueMaxUint16})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ReservationID.Value() != valueMaxUint16 {
		t.Errorf(
			types.ErrorMismatchValue,
			valueMaxUint16,
			req.ReservationID.Value(),
		)
	}
}

func TestReq_NegativeValue(t *testing.T) {
	t.Parallel()

	_, err := cr.Req(cr.ReqInput{ReservationID: valueNegative})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative reservation ID")
	}

	if !strings.Contains(err.Error(), errReservationID) {
		t.Errorf(types.ErrorWantContains, err, errReservationID)
	}
}

func TestReq_ExceedsMaxValue(t *testing.T) {
	t.Parallel()

	_, err := cr.Req(cr.ReqInput{ReservationID: valueExceedsMax})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "reservation ID exceeds max")
	}

	if !strings.Contains(err.Error(), errReservationID) {
		t.Errorf(types.ErrorWantContains, err, errReservationID)
	}
}

func TestReq_LargeNegativeValue(t *testing.T) {
	t.Parallel()

	_, err := cr.Req(cr.ReqInput{ReservationID: valueLargeNegativ})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "large negative reservation ID")
	}

	if !strings.Contains(err.Error(), errReservationID) {
		t.Errorf(types.ErrorWantContains, err, errReservationID)
	}
}
