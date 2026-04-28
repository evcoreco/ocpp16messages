package stoptransaction_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/evcoreco/ocpp16messages/stoptransaction"
	types "github.com/evcoreco/ocpp16types"
)

const (
	testValidIDTag          = "RFID-TAG-12345"
	testValidTimestamp      = "2025-01-15T10:30:00Z"
	testTransactionID12345  = 12345
	testMeterStop5000       = 5000
	testValueZero           = 0
	testValueNegativeOne    = -1
	testTxDataLenZero       = 0
	testExpectedTxDataLen   = 1
	errFieldTransactionID   = "transactionId"
	errFieldMeterStop       = "meterStop"
	errFieldTimestamp       = "timestamp"
	errFieldIDTag           = "idTag"
	errFieldReason          = "reason"
	errReqReasonNil         = "Req() Reason = nil, want non-nil"
	errReqIDTagNil          = "Req() IDTag = nil, want non-nil"
	errMsgExceedsMaxLen     = "exceeds maximum length"
	errMsgWantExceedsMaxLen = "Req() error = %v, want 'exceeds maximum length'"
	errMsgNonPrintable      = "non-printable ASCII"
	errMsgWantNonPrintable  = "Req() error = %v, want 'non-printable ASCII'"
	errMsgInvalidTxData     = "transactionData[0]"
)

func TestReq_Valid(t *testing.T) {
	t.Parallel()

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	expectedTransactionID := uint16(testTransactionID12345)
	if req.TransactionID.Value() != expectedTransactionID {
		t.Errorf(
			types.ErrorMismatch,
			expectedTransactionID,
			req.TransactionID.Value(),
		)
	}

	expectedMeterStop := uint16(testMeterStop5000)
	if req.MeterStop.Value() != expectedMeterStop {
		t.Errorf(types.ErrorMismatch, expectedMeterStop, req.MeterStop.Value())
	}
}

func TestReq_ValidWithIDTag(t *testing.T) {
	t.Parallel()

	idTag := testValidIDTag

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           &idTag,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.IDTag == nil {
		t.Error(errReqIDTagNil)
	}

	if req.IDTag.String() != testValidIDTag {
		t.Errorf(types.ErrorMismatch, testValidIDTag, req.IDTag.String())
	}
}

func TestReq_ValidWithReasonLocal(t *testing.T) {
	t.Parallel()

	reason := "Local"

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          &reason,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Reason == nil {
		t.Error(errReqReasonNil)
	}

	if req.Reason.String() != reason {
		t.Errorf(types.ErrorMismatch, reason, req.Reason.String())
	}
}

func TestReq_ValidWithReasonRemote(t *testing.T) {
	t.Parallel()

	reason := "Remote"

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          &reason,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Reason == nil {
		t.Error(errReqReasonNil)
	}

	if req.Reason.String() != reason {
		t.Errorf(types.ErrorMismatch, reason, req.Reason.String())
	}
}

func TestReq_ValidWithReasonEVDisconnected(t *testing.T) {
	t.Parallel()

	reason := "EVDisconnected"

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          &reason,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Reason == nil {
		t.Error(errReqReasonNil)
	}

	if req.Reason.String() != reason {
		t.Errorf(types.ErrorMismatch, reason, req.Reason.String())
	}
}

func TestReq_ValidWithTransactionData(t *testing.T) {
	t.Parallel()

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID: testTransactionID12345,
		IDTag:         nil,
		MeterStop:     testMeterStop5000,
		Timestamp:     testValidTimestamp,
		Reason:        nil,
		TransactionData: []types.MeterValueInput{
			{
				Timestamp: testValidTimestamp,
				SampledValue: []types.SampledValueInput{
					{
						Value:     "5000",
						Context:   nil,
						Format:    nil,
						Measurand: nil,
						Phase:     nil,
						Location:  nil,
						Unit:      nil,
					},
				},
			},
		},
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if len(req.TransactionData) != testExpectedTxDataLen {
		t.Errorf(
			types.ErrorMismatch,
			testExpectedTxDataLen,
			len(req.TransactionData),
		)
	}
}

func TestReq_ValidWithEmptyTransactionData(t *testing.T) {
	t.Parallel()

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: []types.MeterValueInput{},
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.TransactionData == nil {
		t.Error("TransactionData = nil, want empty slice")
	}

	if len(req.TransactionData) != testTxDataLenZero {
		t.Errorf(
			types.ErrorMismatch, testTxDataLenZero, len(req.TransactionData),
		)
	}
}

func TestReq_TransactionIDZero(t *testing.T) {
	t.Parallel()

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testValueZero,
		IDTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.TransactionID.Value() != testValueZero {
		t.Errorf(types.ErrorMismatch, testValueZero, req.TransactionID.Value())
	}
}

func TestReq_TransactionIDNegative(t *testing.T) {
	t.Parallel()

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testValueNegativeOne,
		IDTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for negative transactionId")
	}

	if !strings.Contains(err.Error(), errFieldTransactionID) {
		t.Errorf(types.ErrorWantContains, err, errFieldTransactionID)
	}
}

func TestReq_MeterStopZero(t *testing.T) {
	t.Parallel()

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           nil,
		MeterStop:       testValueZero,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.MeterStop.Value() != testValueZero {
		t.Errorf(types.ErrorMismatch, testValueZero, req.MeterStop.Value())
	}
}

func TestReq_MeterStopNegative(t *testing.T) {
	t.Parallel()

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           nil,
		MeterStop:       testValueNegativeOne,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for negative meterStop")
	}

	if !strings.Contains(err.Error(), errFieldMeterStop) {
		t.Errorf(types.ErrorWantContains, err, errFieldMeterStop)
	}
}

func TestReq_InvalidTimestamp(t *testing.T) {
	t.Parallel()

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       "not-a-timestamp",
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid timestamp")
	}

	if !strings.Contains(err.Error(), errFieldTimestamp) {
		t.Errorf(types.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_EmptyTimestamp(t *testing.T) {
	t.Parallel()

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       "",
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty timestamp")
	}

	if !strings.Contains(err.Error(), errFieldTimestamp) {
		t.Errorf(types.ErrorWantContains, err, errFieldTimestamp)
	}
}

func TestReq_EmptyIDTag(t *testing.T) {
	t.Parallel()

	emptyTag := ""

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           &emptyTag,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for empty idTag")
	}

	if !errors.Is(err, types.ErrEmptyValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrEmptyValue)
	}
}

func TestReq_IDTagTooLong(t *testing.T) {
	t.Parallel()

	longTag := "RFID-ABC123456789012345" // 23 chars, max is 20

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           &longTag,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for IDTag too long")
	}

	if !strings.Contains(err.Error(), errMsgExceedsMaxLen) {
		t.Errorf(errMsgWantExceedsMaxLen, err)
	}
}

func TestReq_IDTagInvalidCharacters(t *testing.T) {
	t.Parallel()

	invalidTag := "RFID\x00ABC"

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           &invalidTag,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for non-printable chars")
	}

	if !strings.Contains(err.Error(), errMsgNonPrintable) {
		t.Errorf(errMsgWantNonPrintable, err)
	}
}

func TestReq_InvalidReason(t *testing.T) {
	t.Parallel()

	invalidReason := "InvalidReason"

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           nil,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          &invalidReason,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid reason")
	}

	if !strings.Contains(err.Error(), errFieldReason) {
		t.Errorf(types.ErrorWantContains, err, errFieldReason)
	}

	if !errors.Is(err, types.ErrInvalidValue) {
		t.Errorf(types.ErrorWrapping, err, types.ErrInvalidValue)
	}
}

func TestReq_InvalidTransactionData(t *testing.T) {
	t.Parallel()

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID: testTransactionID12345,
		IDTag:         nil,
		MeterStop:     testMeterStop5000,
		Timestamp:     testValidTimestamp,
		Reason:        nil,
		TransactionData: []types.MeterValueInput{
			{
				Timestamp:    "invalid-timestamp",
				SampledValue: []types.SampledValueInput{},
			},
		},
	})
	if err == nil {
		t.Error("Req() error = nil, want error for invalid transactionData")
	}

	if !strings.Contains(err.Error(), errMsgInvalidTxData) {
		t.Errorf(types.ErrorWantContains, err, errMsgInvalidTxData)
	}
}

func TestReq_MultipleErrors(t *testing.T) {
	t.Parallel()

	emptyTag := ""
	invalidReason := "BadReason"

	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testValueNegativeOne,
		IDTag:           &emptyTag,
		MeterStop:       testValueNegativeOne,
		Timestamp:       "invalid",
		Reason:          &invalidReason,
		TransactionData: nil,
	})
	if err == nil {
		t.Error("Req() error = nil, want error for multiple invalid fields")
	}

	errStr := err.Error()

	if !strings.Contains(errStr, errFieldTransactionID) {
		t.Errorf(types.ErrorWantContains, err, errFieldTransactionID)
	}

	if !strings.Contains(errStr, errFieldIDTag) {
		t.Errorf(types.ErrorWantContains, err, errFieldIDTag)
	}

	if !strings.Contains(errStr, errFieldMeterStop) {
		t.Errorf(types.ErrorWantContains, err, errFieldMeterStop)
	}

	if !strings.Contains(errStr, errFieldTimestamp) {
		t.Errorf(types.ErrorWantContains, err, errFieldTimestamp)
	}

	if !strings.Contains(errStr, errFieldReason) {
		t.Errorf(types.ErrorWantContains, err, errFieldReason)
	}
}

func TestReq_Complete(t *testing.T) {
	t.Parallel()

	idTag := testValidIDTag
	reason := "Local"

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   testTransactionID12345,
		IDTag:           &idTag,
		MeterStop:       testMeterStop5000,
		Timestamp:       testValidTimestamp,
		Reason:          &reason,
		TransactionData: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	expectedTxId := uint16(testTransactionID12345)
	if req.TransactionID.Value() != expectedTxId {
		t.Errorf(types.ErrorMismatch, expectedTxId, req.TransactionID.Value())
	}

	if req.IDTag == nil {
		t.Error(errReqIDTagNil)
	}

	expectedMeterStop := uint16(testMeterStop5000)
	if req.MeterStop.Value() != expectedMeterStop {
		t.Errorf(types.ErrorMismatch, expectedMeterStop, req.MeterStop.Value())
	}

	if req.Reason == nil {
		t.Error(errReqReasonNil)
	}
}
