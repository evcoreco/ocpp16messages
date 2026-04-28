package unlockconnector_test

import (
	"strings"
	"testing"

	uc "github.com/evcoreco/ocpp16messages/unlockconnector"
	types "github.com/evcoreco/ocpp16types"
)

const (
	testConnectorIDValid = 1
	testConnectorIDTwo   = 2
	testConnectorIDMax   = 65535
	testConnectorIDZero  = 0
	testConnectorIDOver  = 65536
	testConnectorIDNeg   = -1
	errConnectorID       = "connectorId"
)

func TestReq_Valid(t *testing.T) {
	t.Parallel()

	req, err := uc.Req(uc.ReqInput{ConnectorID: testConnectorIDValid})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != uint16(testConnectorIDValid) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testConnectorIDValid),
			req.ConnectorID.Value(),
		)
	}
}

func TestReq_Valid_ConnectorTwo(t *testing.T) {
	t.Parallel()

	req, err := uc.Req(uc.ReqInput{ConnectorID: testConnectorIDTwo})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != uint16(testConnectorIDTwo) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testConnectorIDTwo),
			req.ConnectorID.Value(),
		)
	}
}

func TestReq_Valid_Max(t *testing.T) {
	t.Parallel()

	req, err := uc.Req(uc.ReqInput{ConnectorID: testConnectorIDMax})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != uint16(testConnectorIDMax) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testConnectorIDMax),
			req.ConnectorID.Value(),
		)
	}
}

func TestReq_Invalid_ConnectorIDZero(t *testing.T) {
	t.Parallel()

	_, err := uc.Req(uc.ReqInput{ConnectorID: testConnectorIDZero})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "connectorId zero")
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}

func TestReq_Invalid_ConnectorIDNegative(t *testing.T) {
	t.Parallel()

	_, err := uc.Req(uc.ReqInput{ConnectorID: testConnectorIDNeg})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative connectorId")
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}

func TestReq_Invalid_ConnectorIDExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := uc.Req(uc.ReqInput{ConnectorID: testConnectorIDOver})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "connectorId exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}
