package metervalues_test

import (
	"strings"
	"testing"

	"github.com/evcoreco/ocpp16messages/metervalues"
	types "github.com/evcoreco/ocpp16types"
)

const (
	validConnectorID     = 1
	validTimestampReq    = "2025-01-02T15:00:00Z"
	validValueReq        = "100"
	connectorIdZero      = 0
	validTransactionID   = 123
	expectedMeterCount1  = 1
	expectedMeterCount2  = 2
	fieldConnectorID     = "ConnectorID"
	fieldTransactionID   = "TransactionID"
	fieldMeterValue      = "MeterValue"
	negativeConnectorID  = -1
	invalidTransactionID = -1
)

// Helper function to create int pointer.
func intPtr(i int) *int {
	return &i
}

func TestReq_ValidMinimalInput(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorID:   validConnectorID,
		TransactionID: nil,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []types.SampledValueInput{
					{
						Value:     validValueReq,
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
	}

	req, err := metervalues.Req(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != validConnectorID {
		t.Errorf(
			types.ErrorMismatchValue,
			validConnectorID,
			req.ConnectorID.Value(),
		)
	}

	if len(req.MeterValue) != expectedMeterCount1 {
		t.Errorf(
			types.ErrorMismatchValue,
			expectedMeterCount1,
			len(req.MeterValue),
		)
	}
}

func TestReq_ValidWithTransactionID(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorID:   validConnectorID,
		TransactionID: intPtr(validTransactionID),
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []types.SampledValueInput{
					{
						Value:     validValueReq,
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
	}

	req, err := metervalues.Req(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.TransactionID == nil {
		t.Errorf(types.ErrorWantNonNil, "TransactionID")
	}

	if req.TransactionID.Value() != validTransactionID {
		t.Errorf(
			types.ErrorMismatchValue,
			validTransactionID,
			req.TransactionID.Value(),
		)
	}
}

func TestReq_ValidConnectorIDZero(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorID:   connectorIdZero,
		TransactionID: nil,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []types.SampledValueInput{
					{
						Value:     validValueReq,
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
	}

	req, err := metervalues.Req(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID.Value() != connectorIdZero {
		t.Errorf(
			types.ErrorMismatchValue,
			connectorIdZero,
			req.ConnectorID.Value(),
		)
	}
}

func TestReq_ValidMultipleMeterValues(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorID:   validConnectorID,
		TransactionID: nil,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: "2025-01-02T15:00:00Z",
				SampledValue: []types.SampledValueInput{
					{
						Value:     "100",
						Context:   nil,
						Format:    nil,
						Measurand: nil,
						Phase:     nil,
						Location:  nil,
						Unit:      nil,
					},
				},
			},
			{
				Timestamp: "2025-01-02T15:05:00Z",
				SampledValue: []types.SampledValueInput{
					{
						Value:     "150",
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
	}

	req, err := metervalues.Req(input)
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if len(req.MeterValue) != expectedMeterCount2 {
		t.Errorf(
			types.ErrorMismatchValue,
			expectedMeterCount2,
			len(req.MeterValue),
		)
	}
}

func TestReq_NegativeConnectorID(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorID:   negativeConnectorID,
		TransactionID: nil,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []types.SampledValueInput{
					{
						Value:     validValueReq,
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
	}

	_, err := metervalues.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative connector id")
	}

	if !strings.Contains(err.Error(), fieldConnectorID) {
		t.Errorf(types.ErrorWantContains, err, fieldConnectorID)
	}
}

func TestReq_NegativeTransactionID(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorID:   validConnectorID,
		TransactionID: intPtr(invalidTransactionID),
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: validTimestampReq,
				SampledValue: []types.SampledValueInput{
					{
						Value:     validValueReq,
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
	}

	_, err := metervalues.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative transaction id")
	}

	if !strings.Contains(err.Error(), fieldTransactionID) {
		t.Errorf(types.ErrorWantContains, err, fieldTransactionID)
	}
}

func TestReq_EmptyMeterValue(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorID:   validConnectorID,
		TransactionID: nil,
		MeterValue:    []types.MeterValueInput{},
	}

	_, err := metervalues.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty meter value")
	}

	if !strings.Contains(err.Error(), fieldMeterValue) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterValue)
	}
}

func TestReq_NilMeterValue(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorID:   validConnectorID,
		TransactionID: nil,
		MeterValue:    nil,
	}

	_, err := metervalues.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "nil meter value")
	}

	if !strings.Contains(err.Error(), fieldMeterValue) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterValue)
	}
}

func TestReq_InvalidMeterValue(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorID:   validConnectorID,
		TransactionID: nil,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: "invalid-timestamp",
				SampledValue: []types.SampledValueInput{
					{
						Value:     validValueReq,
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
	}

	_, err := metervalues.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid meter value")
	}

	if !strings.Contains(err.Error(), "meterValue[0]") {
		t.Errorf(types.ErrorWantContains, err, "meterValue[0]")
	}
}

func TestReq_MultipleErrors(t *testing.T) {
	t.Parallel()

	input := metervalues.ReqInput{
		ConnectorID:   negativeConnectorID,
		TransactionID: intPtr(invalidTransactionID),
		MeterValue:    nil,
	}

	_, err := metervalues.Req(input)
	if err == nil {
		t.Errorf(types.ErrorWantNil, "multiple errors")
	}

	if !strings.Contains(err.Error(), fieldConnectorID) {
		t.Errorf(types.ErrorWantContains, err, fieldConnectorID)
	}

	if !strings.Contains(err.Error(), fieldTransactionID) {
		t.Errorf(types.ErrorWantContains, err, fieldTransactionID)
	}

	if !strings.Contains(err.Error(), fieldMeterValue) {
		t.Errorf(types.ErrorWantContains, err, fieldMeterValue)
	}
}
