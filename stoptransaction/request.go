package stoptransaction

import (
	"errors"
	"fmt"

	types "github.com/evcoreco/ocpp16types"
)

const (
	errCountZero = 0
)

// ReqInput represents the raw input data for creating a StopTransaction.req
// message. The constructor Req validates all fields automatically.
type ReqInput struct {
	// Required: The transaction ID of the transaction to stop.
	TransactionID int
	// Optional: The identifier that was used to stop the transaction.
	// May be omitted when the Charge Point itself stops the transaction.
	IDTag *string
	// Required: Energy meter reading at the end of the transaction in Wh.
	MeterStop int
	// Required: Timestamp of when the transaction was stopped.
	Timestamp string
	// Optional: The reason for stopping the transaction.
	// May be omitted if the transaction ended normally (Local).
	Reason *string
	// Optional: Transaction-related meter values.
	TransactionData []types.MeterValueInput
}

// ReqMessage represents an OCPP 1.6 StopTransaction.req message.
type ReqMessage struct {
	TransactionID   types.Integer
	IDTag           *types.IDToken
	MeterStop       types.Integer
	Timestamp       types.DateTime
	Reason          *types.Reason
	TransactionData []types.MeterValue
}

// reqValidation holds validated fields during Req construction.
type reqValidation struct {
	transactionId   types.Integer
	idTag           *types.IDToken
	meterStop       types.Integer
	timestamp       types.DateTime
	reason          *types.Reason
	transactionData []types.MeterValue
}

// Req creates a StopTransaction.req message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// Returns an error if:
//   - TransactionID is negative or exceeds uint16 max value (65535)
//   - IDTag (if provided) exceeds 20 characters or contains non-printable ASCII
//   - MeterStop is negative or exceeds uint16 max value (65535)
//   - Timestamp is not a valid RFC3339 formatted date
//   - Reason (if provided) is not a valid Reason value
//   - TransactionData (if provided) contains invalid meter values
func Req(input ReqInput) (ReqMessage, error) {
	validated, errs := validateReqInput(input)

	if len(errs) > errCountZero {
		return ReqMessage{
			TransactionID:   types.Integer{},
			IDTag:           nil,
			MeterStop:       types.Integer{},
			Timestamp:       types.DateTime{},
			Reason:          nil,
			TransactionData: nil,
		}, errors.Join(errs...)
	}

	return ReqMessage{
		TransactionID:   validated.transactionId,
		IDTag:           validated.idTag,
		MeterStop:       validated.meterStop,
		Timestamp:       validated.timestamp,
		Reason:          validated.reason,
		TransactionData: validated.transactionData,
	}, nil
}

// validateReqInput validates all fields in ReqInput.
func validateReqInput(input ReqInput) (reqValidation, []error) {
	var errs []error

	var validated reqValidation

	validated.transactionId, errs = validateTransactionID(
		input.TransactionID,
		errs,
	)
	validated.meterStop, errs = validateMeterStop(input.MeterStop, errs)
	validated.timestamp, errs = validateTimestamp(input.Timestamp, errs)

	if input.IDTag != nil {
		validated.idTag, errs = validateIDTag(*input.IDTag, errs)
	}

	if input.Reason != nil {
		validated.reason, errs = validateReason(*input.Reason, errs)
	}

	if input.TransactionData != nil {
		validated.transactionData, errs = validateTransactionData(
			input.TransactionData,
			errs,
		)
	}

	return validated, errs
}

// validateTransactionID validates the transactionId field.
func validateTransactionID(
	transactionId int,
	errs []error,
) (types.Integer, []error) {
	val, err := types.NewInteger(transactionId)
	if err != nil {
		return types.Integer{}, append(
			errs, fmt.Errorf("transactionId: %w", err),
		)
	}

	return val, errs
}

// validateIDTag validates the idTag field.
func validateIDTag(idTag string, errs []error) (*types.IDToken, []error) {
	ciStr, err := types.NewCiString20Type(idTag)
	if err != nil {
		return nil, append(errs, fmt.Errorf("idTag: %w", err))
	}

	token := types.NewIDToken(ciStr)

	return &token, errs
}

// validateMeterStop validates the meterStop field.
func validateMeterStop(meterStop int, errs []error) (types.Integer, []error) {
	val, err := types.NewInteger(meterStop)
	if err != nil {
		return types.Integer{}, append(errs, fmt.Errorf("meterStop: %w", err))
	}

	return val, errs
}

// validateTimestamp validates the timestamp field.
func validateTimestamp(
	timestamp string, errs []error,
) (types.DateTime, []error) {
	val, err := types.NewDateTime(timestamp)
	if err != nil {
		return types.DateTime{}, append(errs, fmt.Errorf("timestamp: %w", err))
	}

	return val, errs
}

// validateReason validates the reason field.
func validateReason(reason string, errs []error) (*types.Reason, []error) {
	reasonVal := types.Reason(reason)
	if !reasonVal.IsValid() {
		return nil, append(
			errs,
			fmt.Errorf("reason: %w: %s", types.ErrInvalidValue, reason),
		)
	}

	return &reasonVal, errs
}

// validateTransactionData validates the transactionData field.
func validateTransactionData(
	transactionData []types.MeterValueInput,
	errs []error,
) ([]types.MeterValue, []error) {
	const transactionDataEmpty = 0

	if transactionData == nil {
		return nil, errs
	}

	var validValues []types.MeterValue

	for i, mvInput := range transactionData {
		meterValue, err := types.NewMeterValue(mvInput)
		if err != nil {
			errs = append(errs, fmt.Errorf("transactionData[%d]: %w", i, err))
		} else {
			validValues = append(validValues, meterValue)
		}
	}

	if len(transactionData) == transactionDataEmpty {
		return []types.MeterValue{}, errs
	}

	return validValues, errs
}
