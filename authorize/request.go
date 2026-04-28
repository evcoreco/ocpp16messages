package authorize

import (
	"errors"
	"fmt"

	types "github.com/evcoreco/ocpp16types"
)

const (
	// errCountZero is the empty error count.
	errCountZero = 0
)

// ReqInput represents the raw input data for creating an Authorize.req message.
// The constructor Req validates all fields automatically.
type ReqInput struct {
	IDTag string
}

// ReqMessage represents an OCPP 1.6 Authorize.req message.
type ReqMessage struct {
	IDTag types.IDToken
}

// Req creates an Authorize.req message from the given input.
// It validates all fields automatically and returns an error if:
//   - IDTag is empty
//   - IDTag exceeds 20 characters
//   - IDTag contains non-printable ASCII characters
func Req(input ReqInput) (ReqMessage, error) {
	var errs []error

	str, err := types.NewCiString20Type(input.IDTag)
	if err != nil {
		errs = append(errs, fmt.Errorf("idTag: %w", err))
	}

	if len(errs) > errCountZero {
		return ReqMessage{}, errors.Join(errs...)
	}

	idToken := types.NewIDToken(str)

	return ReqMessage{IDTag: idToken}, nil
}
