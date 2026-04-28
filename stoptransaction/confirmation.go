package stoptransaction

import (
	"errors"
	"fmt"

	types "github.com/evcoreco/ocpp16types"
)

// ConfInput represents the raw input data for creating a StopTransaction.conf
// message. The constructor Conf validates all fields automatically.
type ConfInput struct {
	// Optional: AuthorizationStatus value.
	// If provided, indicates the status of the idTag used to stop the
	// transaction. May be absent if no idTag was included in the request.
	Status *string
	// Optional: RFC3339 formatted expiry date for the authorization.
	ExpiryDate *string
	// Optional: Parent ID tag (max 20 chars).
	ParentIDTag *string
}

// ConfMessage represents an OCPP 1.6 StopTransaction.conf message.
type ConfMessage struct {
	IDTagInfo *types.IDTagInfo
}

// confValidation holds validated fields during Conf construction.
type confValidation struct {
	idTagInfo     types.IDTagInfo
	expiryDate    types.DateTime
	parentIDToken types.IDToken
	hasIDTagInfo  bool
}

// Conf creates a StopTransaction.conf message from the given input.
// It validates all fields and accumulates all errors, returning them together.
// This allows callers to see all validation issues at once rather than one at
// a time. Returns an error if:
//   - Status (if provided) is not a valid AuthorizationStatus value
//   - ExpiryDate (if provided) is not a valid RFC3339 date
//   - ParentIDTag (if provided) exceeds 20 characters or contains invalid chars
//
// Note: IDTagInfo is optional in StopTransaction.conf. If no Status is
// provided, the IDTagInfo field in the response will be nil.
func Conf(input ConfInput) (ConfMessage, error) {
	validated, errs := validateConfInput(input)

	if len(errs) > errCountZero {
		return ConfMessage{IDTagInfo: nil}, errors.Join(errs...)
	}

	return buildConfMessage(input, validated), nil
}

// validateConfInput validates all fields in ConfInput.
func validateConfInput(input ConfInput) (confValidation, []error) {
	var errs []error

	var validated confValidation

	if input.Status != nil {
		validated.idTagInfo, errs = validateStatus(*input.Status, errs)
		validated.hasIDTagInfo = true
	}

	if input.ExpiryDate != nil {
		validated.expiryDate, errs = validateExpiryDate(*input.ExpiryDate, errs)
	}

	if input.ParentIDTag != nil {
		validated.parentIDToken, errs = validateParentIDTag(
			*input.ParentIDTag,
			errs,
		)
	}

	return validated, errs
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
	if !validated.hasIDTagInfo {
		return ConfMessage{IDTagInfo: nil}
	}

	idTagInfo := validated.idTagInfo

	if input.ExpiryDate != nil {
		idTagInfo = idTagInfo.WithExpiryDate(validated.expiryDate)
	}

	if input.ParentIDTag != nil {
		idTagInfo = idTagInfo.WithParentIDTag(validated.parentIDToken)
	}

	return ConfMessage{IDTagInfo: &idTagInfo}
}
