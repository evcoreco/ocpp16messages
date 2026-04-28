package remotestarttransaction

import (
	"errors"
	"fmt"

	types "github.com/evcoreco/ocpp16types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// ReqInput represents the raw input data for creating a
// RemoteStartTransaction.req message.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The identifier to use for starting the transaction.
	IDTag string
	// Optional: The connector on which to start the transaction.
	// If not provided, the Charge Point will choose an available connector.
	ConnectorID *int
}

// ReqMessage represents an OCPP 1.6 RemoteStartTransaction.req message.
type ReqMessage struct {
	IDTag       types.CiString20Type
	ConnectorID *types.Integer
}

// Req creates a RemoteStartTransaction.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - IDTag is empty
//   - IDTag exceeds 20 characters
//   - IDTag contains non-printable ASCII characters
//   - ConnectorID (if provided) is negative or exceeds uint16 max value (65535)
func Req(input ReqInput) (ReqMessage, error) {
	var errs []error

	idTag, err := types.NewCiString20Type(input.IDTag)
	if err != nil {
		errs = append(errs, fmt.Errorf("idTag: %w", err))
	}

	var connectorId *types.Integer

	if input.ConnectorID != nil {
		connectorId, errs = validateConnectorID(*input.ConnectorID, errs)
	}

	if len(errs) > errCountZero {
		return ReqMessage{}, errors.Join(errs...)
	}

	return ReqMessage{
		IDTag:       idTag,
		ConnectorID: connectorId,
	}, nil
}

// validateConnectorID validates the connectorId field.
func validateConnectorID(
	connectorId int,
	errs []error,
) (*types.Integer, []error) {
	val, err := types.NewInteger(connectorId)
	if err != nil {
		return nil, append(errs, fmt.Errorf("connectorId: %w", err))
	}

	return &val, errs
}
