package reservenow

import (
	"errors"
	"fmt"

	types "github.com/evcoreco/ocpp16types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// ReqInput represents the raw input data for creating a ReserveNow.req message.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: Unique identifier for this reservation.
	ReservationID int
	// Required: The connector to reserve. 0 means reserve any available.
	ConnectorID int
	// Required: The identifier to use for the reservation.
	IDTag string
	// Required: The date and time when the reservation expires (RFC3339).
	ExpiryDate string
	// Optional: The parent identifier to use for authorization.
	ParentIDTag *string
}

// ReqMessage represents an OCPP 1.6 ReserveNow.req message.
type ReqMessage struct {
	ReservationID types.Integer
	ConnectorID   types.Integer
	IDTag         types.CiString20Type
	ExpiryDate    types.DateTime
	ParentIDTag   *types.CiString20Type
}

// Req creates a ReserveNow.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - ReservationID is negative or exceeds uint16 max value (65535)
//   - ConnectorID is negative or exceeds uint16 max value (65535)
//   - IDTag is empty
//   - IDTag exceeds 20 characters
//   - IDTag contains non-printable ASCII characters
//   - ExpiryDate is not a valid RFC3339 timestamp
//   - ParentIDTag (if provided) exceeds 20 characters
//   - ParentIDTag (if provided) contains non-printable ASCII characters
func Req(input ReqInput) (ReqMessage, error) {
	var errs []error

	reservationId, err := types.NewInteger(input.ReservationID)
	if err != nil {
		errs = append(errs, fmt.Errorf("reservationId: %w", err))
	}

	connectorId, err := types.NewInteger(input.ConnectorID)
	if err != nil {
		errs = append(errs, fmt.Errorf("connectorId: %w", err))
	}

	idTag, err := types.NewCiString20Type(input.IDTag)
	if err != nil {
		errs = append(errs, fmt.Errorf("idTag: %w", err))
	}

	expiryDate, err := types.NewDateTime(input.ExpiryDate)
	if err != nil {
		errs = append(errs, fmt.Errorf("expiryDate: %w", err))
	}

	var parentIDTag *types.CiString20Type

	if input.ParentIDTag != nil {
		parentIDTag, errs = validateParentIDTag(*input.ParentIDTag, errs)
	}

	if len(errs) > errCountZero {
		return ReqMessage{}, errors.Join(errs...)
	}

	return ReqMessage{
		ReservationID: reservationId,
		ConnectorID:   connectorId,
		IDTag:         idTag,
		ExpiryDate:    expiryDate,
		ParentIDTag:   parentIDTag,
	}, nil
}

// validateParentIDTag validates the optional parentIDTag field.
func validateParentIDTag(
	parentIDTag string,
	errs []error,
) (*types.CiString20Type, []error) {
	val, err := types.NewCiString20Type(parentIDTag)
	if err != nil {
		return nil, append(errs, fmt.Errorf("parentIDTag: %w", err))
	}

	return &val, errs
}
