package remotestoptransaction

import (
	"fmt"

	types "github.com/evcoreco/ocpp16types"
)

// ReqInput represents the raw input data for creating a
// RemoteStopTransaction.req message.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The identifier of the transaction to stop.
	TransactionID int
}

// ReqMessage represents an OCPP 1.6 RemoteStopTransaction.req message.
type ReqMessage struct {
	TransactionID types.Integer
}

// Req creates a RemoteStopTransaction.req message from the given input.
// It validates all fields and returns an error if:
//   - TransactionID is negative or exceeds uint16 max value (65535)
func Req(input ReqInput) (ReqMessage, error) {
	transactionId, err := types.NewInteger(input.TransactionID)
	if err != nil {
		return ReqMessage{}, fmt.Errorf("transactionId: %w", err)
	}

	return ReqMessage{
		TransactionID: transactionId,
	}, nil
}
