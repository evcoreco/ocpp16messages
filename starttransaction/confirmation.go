package starttransaction

import (
	"errors"
	"fmt"

	types "github.com/evcoreco/ocpp16types"
)

// ConfInput represents the raw input data for creating a StartTransaction.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Required: The transaction identifier assigned by the Central System.
	TransactionID int
	// Required: AuthorizationStatus value.
	Status string
	// Optional: RFC3339 formatted expiry date for the authorization.
	ExpiryDate *string
	// Optional: Parent ID tag (max 20 chars).
	ParentIDTag *string
}

// ConfMessage represents an OCPP 1.6 StartTransaction.conf message.
type ConfMessage struct {
	TransactionID types.Integer
	IDTagInfo     types.IDTagInfo
}

// confValidation holds validated fields during Conf construction.
type confValidation struct {
	transactionId types.Integer
	idTagInfo     types.IDTagInfo
	expiryDate    types.DateTime
	parentIDToken types.IDToken
}

// Conf creates a StartTransaction.conf message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// This allows callers to see all validation issues at once rather than one at
// a time. Returns an error if:
//   - TransactionID is negative or exceeds uint16 max value (65535)
//   - Status is not a valid AuthorizationStatus value
//   - ExpiryDate (if provided) is not a valid RFC3339 date
//   - ParentIDTag (if provided) exceeds 20 characters or contains invalid chars
func Conf(input ConfInput) (ConfMessage, error) {
	validated, errs := validateConfInput(input)

	if len(errs) > errCountZero {
		return ConfMessage{}, errors.Join(errs...)
	}

	return buildConfMessage(input, validated), nil
}

// validateConfInput validates all fields in ConfInput and returns validated
// values along with any errors.
func validateConfInput(input ConfInput) (confValidation, []error) {
	var errs []error

	var validated confValidation

	// Validate transactionId (required)
	validated.transactionId, errs = validateTransactionID(
		input.TransactionID,
		errs,
	)

	// Validate status (required)
	validated.idTagInfo, errs = validateStatus(input.Status, errs)

	// Validate expiryDate (optional)
	if input.ExpiryDate != nil {
		validated.expiryDate, errs = validateExpiryDate(*input.ExpiryDate, errs)
	}

	// Validate parentIDTag (optional)
	if input.ParentIDTag != nil {
		validated.parentIDToken, errs = validateParentIDTag(
			*input.ParentIDTag,
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

// validateStatus validates the status field and returns the IDTagInfo.
func validateStatus(status string, errs []error) (types.IDTagInfo, []error) {
	info, err := types.NewIDTagInfo(types.AuthorizationStatus(status))
	if err != nil {
		return types.IDTagInfo{}, append(errs, fmt.Errorf("status: %w", err))
	}

	return info, errs
}

// validateExpiryDate validates the expiry date string and returns the DateTime.
func validateExpiryDate(date string, errs []error) (types.DateTime, []error) {
	expiryDate, err := types.NewDateTime(date)
	if err != nil {
		return types.DateTime{}, append(errs, fmt.Errorf("expiryDate: %w", err))
	}

	return expiryDate, errs
}

// validateParentIDTag validates the parent ID tag and returns the token.
func validateParentIDTag(tag string, errs []error) (types.IDToken, []error) {
	ciStr, err := types.NewCiString20Type(tag)
	if err != nil {
		return types.IDToken{}, append(errs, fmt.Errorf("parentIDTag: %w", err))
	}

	return types.NewIDToken(ciStr), errs
}

// buildConfMessage constructs the final ConfMessage with validated fields.
func buildConfMessage(input ConfInput, validated confValidation) ConfMessage {
	idTagInfo := validated.idTagInfo

	if input.ExpiryDate != nil {
		idTagInfo = idTagInfo.WithExpiryDate(validated.expiryDate)
	}

	if input.ParentIDTag != nil {
		idTagInfo = idTagInfo.WithParentIDTag(validated.parentIDToken)
	}

	return ConfMessage{
		TransactionID: validated.transactionId,
		IDTagInfo:     idTagInfo,
	}
}
