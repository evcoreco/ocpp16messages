//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"
	"time"

	mv "github.com/evcoreco/ocpp16messages/metervalues"
	types "github.com/evcoreco/ocpp16types"
)

func FuzzMeterValuesReq(f *testing.F) {
	f.Add(0, false, 0, false, "", "")
	f.Add(1, false, 0, true, "2025-01-15T10:30:00Z", "100")
	f.Add(-1, false, 0, true, "2025-01-15T10:30:00Z", "100")
	f.Add(1, true, -1, true, "2025-01-15T10:30:00Z", "100")
	f.Add(1, true, 1, true, "bad-timestamp", "100")
	f.Add(1, true, 1, true, "2025-01-15T10:30:00Z", "")

	f.Fuzz(func(
		t *testing.T,
		connectorId int,
		hasTransactionID bool,
		transactionId int,
		hasMeterValue bool,
		timestamp string,
		value string,
	) {
		if len(timestamp) > maxFuzzStringLen || len(value) > maxFuzzStringLen {
			t.Skip()
		}

		var transactionIdPtr *int
		if hasTransactionID {
			transactionIdPtr = &transactionId
		}

		var metervalues []types.MeterValueInput
		if hasMeterValue {
			metervalues = []types.MeterValueInput{
				{
					Timestamp: timestamp,
					SampledValue: []types.SampledValueInput{
						{Value: value},
					},
				},
			}
		}

		req, err := mv.Req(mv.ReqInput{
			ConnectorID:   connectorId,
			TransactionID: transactionIdPtr,
			MeterValue:    metervalues,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) && !errors.Is(err, types.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if connectorId < 0 || connectorId > math.MaxUint16 {
			t.Fatalf("Req succeeded with connectorId=%d", connectorId)
		}

		if got := req.ConnectorID.Value(); got != uint16(connectorId) {
			t.Fatalf("ConnectorID = %d, want %d", got, connectorId)
		}

		if hasTransactionID {
			if req.TransactionID == nil {
				t.Fatal("TransactionID = nil, want non-nil")
			}

			if transactionId < 0 || transactionId > math.MaxUint16 {
				t.Fatalf("Req succeeded with transactionId=%d", transactionId)
			}

			if got := req.TransactionID.Value(); got != uint16(transactionId) {
				t.Fatalf("TransactionID = %d, want %d", got, transactionId)
			}
		} else if req.TransactionID != nil {
			t.Fatal("TransactionID != nil, want nil")
		}

		if len(req.MeterValue) == 0 {
			t.Fatal("MeterValue is empty, want at least one")
		}

		if req.MeterValue[0].Timestamp().Value().Location() != time.UTC {
			t.Fatalf(
				"MeterValue[0].Timestamp location = %v, want UTC",
				req.MeterValue[0].Timestamp().Value().Location(),
			)
		}

		if len(req.MeterValue[0].SampledValue()) == 0 {
			t.Fatal("SampledValue is empty, want at least one")
		}
	})
}
