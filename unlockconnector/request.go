package unlockconnector

import (
	"fmt"

	types "github.com/evcoreco/ocpp16types"
)

const (
	// connectorIdMinValue is the minimum valid connectorId (must be > 0).
	connectorIdMinValue = 0
)

// ReqInput represents the raw input data for creating an UnlockConnector.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The identifier of the connector to unlock. Must be > 0.
	ConnectorID int
}

// ReqMessage represents an OCPP 1.6 UnlockConnector.req message.
type ReqMessage struct {
	ConnectorID types.Integer
}

// Req creates an UnlockConnector.req message from the given input.
// It validates all fields and returns an error if:
//   - ConnectorID is zero, negative, or exceeds uint16 max value (65535)
//
// Note: ConnectorID must be > 0 because connector 0 refers to the Charge Point
// itself, not a physical connector.
func Req(input ReqInput) (ReqMessage, error) {
	if input.ConnectorID <= connectorIdMinValue {
		return ReqMessage{}, fmt.Errorf(
			"connectorId: %w", types.ErrInvalidValue,
		)
	}

	connectorId, err := types.NewInteger(input.ConnectorID)
	if err != nil {
		return ReqMessage{}, fmt.Errorf("connectorId: %w", err)
	}

	return ReqMessage{
		ConnectorID: connectorId,
	}, nil
}
