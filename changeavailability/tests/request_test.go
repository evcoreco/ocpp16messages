package changeavailability_test

import (
	"strings"
	"testing"

	ca "github.com/evcoreco/ocpp16messages/changeavailability"
	types "github.com/evcoreco/ocpp16types"
)

const (
	errConnectorID = "connectorId"
	errType        = "type"

	typeOperative   = "Operative"
	typeInoperative = "Inoperative"

	valueZero       = 0
	valuePositive   = 1
	valueMaxUint16  = 65535
	valueExceedsMax = 65536
	valueNegative   = -1
)

func TestReq_Valid_ConnectorZero_Inoperative(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorID: valueZero, Type: typeInoperative}

	req, err := ca.Req(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != valueZero {
		t.Errorf(types.ErrorMismatchValue, valueZero, req.ConnectorID.Value())
	}

	if req.Type != types.AvailabilityTypeInoperative {
		t.Errorf(
			types.ErrorMismatchValue,
			types.AvailabilityTypeInoperative,
			req.Type,
		)
	}
}

func TestReq_Valid_ConnectorOne_Operative(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorID: valuePositive, Type: typeOperative}

	req, err := ca.Req(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != valuePositive {
		t.Errorf(
			types.ErrorMismatchValue, valuePositive, req.ConnectorID.Value(),
		)
	}

	if req.Type != types.AvailabilityTypeOperative {
		t.Errorf(
			types.ErrorMismatchValue,
			types.AvailabilityTypeOperative,
			req.Type,
		)
	}
}

func TestReq_Valid_MaxConnectorID(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorID: valueMaxUint16, Type: typeOperative}

	req, err := ca.Req(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != valueMaxUint16 {
		t.Errorf(
			types.ErrorMismatchValue, valueMaxUint16, req.ConnectorID.Value(),
		)
	}
}

func TestReq_NegativeConnectorID(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorID: valueNegative, Type: typeOperative}

	_, err := ca.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative connector ID")
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}

func TestReq_ExceedsMaxConnectorID(t *testing.T) {
	t.Parallel()

	input := ca.ReqInput{ConnectorID: valueExceedsMax, Type: typeOperative}

	_, err := ca.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "connector ID exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}

func TestReq_InvalidType_Empty(t *testing.T) {
	t.Parallel()

	_, err := ca.Req(ca.ReqInput{ConnectorID: valuePositive, Type: ""})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(types.ErrorWantContains, err, errType)
	}
}

func TestReq_InvalidType_Unknown(t *testing.T) {
	t.Parallel()

	_, err := ca.Req(ca.ReqInput{ConnectorID: valuePositive, Type: "Unknown"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(types.ErrorWantContains, err, errType)
	}
}

func TestReq_InvalidType_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := ca.Req(ca.ReqInput{ConnectorID: valuePositive, Type: "operative"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase type")
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(types.ErrorWantContains, err, errType)
	}
}

func TestReq_MultipleErrors_NegativeConnectorAndInvalidType(t *testing.T) {
	t.Parallel()

	_, err := ca.Req(ca.ReqInput{ConnectorID: valueNegative, Type: "Unknown"})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "multiple invalid fields")
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}

	if !strings.Contains(err.Error(), errType) {
		t.Errorf(types.ErrorWantContains, err, errType)
	}
}
