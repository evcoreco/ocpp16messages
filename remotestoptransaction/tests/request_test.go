package remotestoptransaction_test

import (
	"strings"
	"testing"

	rst "github.com/evcoreco/ocpp16messages/remotestoptransaction"
	types "github.com/evcoreco/ocpp16types"
)

const (
	testTransactionIDValid = 12345
	testTransactionIDZero  = 0
	testTransactionIDMax   = 65535
	testTransactionIDOver  = 65536
	testTransactionIDNeg   = -1
	errTransactionID       = "transactionId"
)

func TestReq_Valid(t *testing.T) {
	t.Parallel()

	req, err := rst.Req(rst.ReqInput{TransactionID: testTransactionIDValid})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.TransactionID.Value() != uint16(testTransactionIDValid) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testTransactionIDValid),
			req.TransactionID.Value(),
		)
	}
}

func TestReq_Valid_Zero(t *testing.T) {
	t.Parallel()

	req, err := rst.Req(rst.ReqInput{TransactionID: testTransactionIDZero})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.TransactionID.Value() != uint16(testTransactionIDZero) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testTransactionIDZero),
			req.TransactionID.Value(),
		)
	}
}

func TestReq_Valid_Max(t *testing.T) {
	t.Parallel()

	req, err := rst.Req(rst.ReqInput{TransactionID: testTransactionIDMax})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.TransactionID.Value() != uint16(testTransactionIDMax) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(testTransactionIDMax),
			req.TransactionID.Value(),
		)
	}
}

func TestReq_TransactionIDNegative(t *testing.T) {
	t.Parallel()

	_, err := rst.Req(rst.ReqInput{TransactionID: testTransactionIDNeg})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative transactionId")
	}

	if !strings.Contains(err.Error(), errTransactionID) {
		t.Errorf(types.ErrorWantContains, err, errTransactionID)
	}
}

func TestReq_TransactionIDExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := rst.Req(rst.ReqInput{TransactionID: testTransactionIDOver})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "transactionId exceeds max")
	}

	if !strings.Contains(err.Error(), errTransactionID) {
		t.Errorf(types.ErrorWantContains, err, errTransactionID)
	}
}
