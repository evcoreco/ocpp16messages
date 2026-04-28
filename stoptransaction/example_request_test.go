package stoptransaction_test

import (
	"fmt"

	"github.com/evcoreco/ocpp16messages/stoptransaction"
	types "github.com/evcoreco/ocpp16types"
)

const (
	exampleTxId           = 12345
	exampleMeterStop      = 50000
	exampleTimestamp      = "2025-01-15T10:30:00Z"
	exampleReqErrorLabel  = "Error:"
	exampleTxIdLabel      = "TransactionID:"
	exampleMeterStopLabel = "MeterStop:"
)

// ExampleReq demonstrates creating a basic StopTransaction.req message.
func ExampleReq() {
	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   exampleTxId,
		IDTag:           nil,
		MeterStop:       exampleMeterStop,
		Timestamp:       exampleTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		fmt.Println(exampleReqErrorLabel, err)

		return
	}

	fmt.Println(exampleTxIdLabel, req.TransactionID.Value())
	fmt.Println(exampleMeterStopLabel, req.MeterStop.Value())
	// Output:
	// TransactionID: 12345
	// MeterStop: 50000
}

// ExampleReq_withIDTag demonstrates stopping a transaction with an ID tag.
func ExampleReq_withIDTag() {
	idTag := "RFID-ABC123"

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   exampleTxId,
		IDTag:           &idTag,
		MeterStop:       exampleMeterStop,
		Timestamp:       exampleTimestamp,
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		fmt.Println(exampleReqErrorLabel, err)

		return
	}

	fmt.Println(exampleTxIdLabel, req.TransactionID.Value())
	fmt.Println("IDTag:", req.IDTag.String())
	// Output:
	// TransactionID: 12345
	// IDTag: RFID-ABC123
}

// ExampleReq_withReason demonstrates stopping a transaction with a reason.
func ExampleReq_withReason() {
	reason := "Remote"

	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   exampleTxId,
		IDTag:           nil,
		MeterStop:       exampleMeterStop,
		Timestamp:       exampleTimestamp,
		Reason:          &reason,
		TransactionData: nil,
	})
	if err != nil {
		fmt.Println(exampleReqErrorLabel, err)

		return
	}

	fmt.Println(exampleTxIdLabel, req.TransactionID.Value())
	fmt.Println("Reason:", req.Reason.String())
	// Output:
	// TransactionID: 12345
	// Reason: Remote
}

// ExampleReq_withTransactionData shows including meter values in the request.
func ExampleReq_withTransactionData() {
	req, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID: exampleTxId,
		IDTag:         nil,
		MeterStop:     exampleMeterStop,
		Timestamp:     exampleTimestamp,
		Reason:        nil,
		TransactionData: []types.MeterValueInput{
			{
				Timestamp: exampleTimestamp,
				SampledValue: []types.SampledValueInput{
					{
						Value:     "50000",
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
		fmt.Println(exampleReqErrorLabel, err)

		return
	}

	fmt.Println(exampleTxIdLabel, req.TransactionID.Value())
	fmt.Println("TransactionData count:", len(req.TransactionData))
	// Output:
	// TransactionID: 12345
	// TransactionData count: 1
}

// ExampleReq_invalidTimestamp demonstrates validation error for bad timestamp.
func ExampleReq_invalidTimestamp() {
	_, err := stoptransaction.Req(stoptransaction.ReqInput{
		TransactionID:   exampleTxId,
		IDTag:           nil,
		MeterStop:       exampleMeterStop,
		Timestamp:       "invalid-timestamp",
		Reason:          nil,
		TransactionData: nil,
	})
	if err != nil {
		fmt.Println("Validation failed as expected")
	}
	// Output:
	// Validation failed as expected
}
