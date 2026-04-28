package metervalues

import (
	"errors"
	"fmt"

	types "github.com/evcoreco/ocpp16types"
)

// ReqInput represents the raw input data for creating a MeterValues.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The ID of the connector for which meter values are reported.
	// Use 0 for the entire Charge Point.
	ConnectorID int
	// Optional: The transaction ID for which meter values are reported.
	TransactionID *int
	// Required: One or more meter value sets.
	MeterValue []types.MeterValueInput
}

// ReqMessage represents an OCPP 1.6 MeterValues.req message.
type ReqMessage struct {
	ConnectorID   types.Integer
	TransactionID *types.Integer
	MeterValue    []types.MeterValue
}

// reqValidation holds validated fields during construction.
type reqValidation struct {
	connectorId   types.Integer
	transactionId *types.Integer
	meterValue    []types.MeterValue
}

// Req creates a MeterValues.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - ConnectorID is negative
//   - MeterValue is empty or contains invalid entries
//   - TransactionID is provided but invalid
func Req(input ReqInput) (ReqMessage, error) {
	validated, errs := validateReqInput(input)

	if errs != nil {
		return ReqMessage{
			ConnectorID:   types.Integer{},
			TransactionID: nil,
			MeterValue:    nil,
		}, errors.Join(errs...)
	}

	return ReqMessage{
		ConnectorID:   validated.connectorId,
		TransactionID: validated.transactionId,
		MeterValue:    validated.meterValue,
	}, nil
}

func validateReqInput(input ReqInput) (reqValidation, []error) {
	var errs []error

	var validated reqValidation

	validated.connectorId, errs = validateReqConnectorID(
		input.ConnectorID,
		errs,
	)

	if input.TransactionID != nil {
		validated.transactionId, errs = validateReqTransactionID(
			*input.TransactionID,
			errs,
		)
	}

	validated.meterValue, errs = validateReqMeterValues(input.MeterValue, errs)

	return validated, errs
}

func validateReqConnectorID(
	connectorId int,
	errs []error,
) (types.Integer, []error) {
	intVal, err := types.NewInteger(connectorId)
	if err != nil {
		return types.Integer{}, append(
			errs,
			fmt.Errorf(types.ErrorFieldFormat, "ConnectorID", err),
		)
	}

	return intVal, errs
}

func validateReqTransactionID(
	transactionId int,
	errs []error,
) (*types.Integer, []error) {
	intVal, err := types.NewInteger(transactionId)
	if err != nil {
		return nil, append(
			errs,
			fmt.Errorf(types.ErrorFieldFormat, "TransactionID", err),
		)
	}

	return &intVal, errs
}

const metervaluesLenZero = 0

func validateReqMeterValues(
	metervalues []types.MeterValueInput,
	errs []error,
) ([]types.MeterValue, []error) {
	if len(metervalues) == metervaluesLenZero {
		return nil, append(
			errs,
			fmt.Errorf(
				types.ErrorFieldFormat, "MeterValue", types.ErrEmptyValue,
			),
		)
	}

	var validValues []types.MeterValue

	for i, mvInput := range metervalues {
		meterValue, err := types.NewMeterValue(mvInput)
		if err != nil {
			errs = append(errs, fmt.Errorf("meterValue[%d]: %w", i, err))
		} else {
			validValues = append(validValues, meterValue)
		}
	}

	return validValues, errs
}
