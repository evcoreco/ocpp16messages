package metervalues_test

import (
	"fmt"

	"github.com/evcoreco/ocpp16messages/metervalues"
	types "github.com/evcoreco/ocpp16types"
)

const (
	exampleConnectorID   = 1
	exampleTransactionID = 42
	exampleTimestamp     = "2025-01-02T15:00:00Z"
	exampleValue         = "12500"
	outputError          = "Error:"
	outputConnectorID    = "ConnectorID:"
)

// ExampleReq demonstrates creating a valid MeterValues.req message with
// a single meter value containing one sampled value.
func ExampleReq() {
	input := metervalues.ReqInput{
		ConnectorID:   exampleConnectorID,
		TransactionID: nil,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: exampleTimestamp,
				SampledValue: []types.SampledValueInput{
					{
						Value:     exampleValue,
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
		fmt.Println(outputError, err)

		return
	}

	fmt.Println(outputConnectorID, req.ConnectorID.Value())
	fmt.Println("MeterValue count:", len(req.MeterValue))
	// Output:
	// ConnectorID: 1
	// MeterValue count: 1
}

// ExampleReq_withTransactionID demonstrates creating a MeterValues.req
// message with an associated transaction ID.
func ExampleReq_withTransactionID() {
	transactionId := exampleTransactionID

	input := metervalues.ReqInput{
		ConnectorID:   exampleConnectorID,
		TransactionID: &transactionId,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: exampleTimestamp,
				SampledValue: []types.SampledValueInput{
					{
						Value:     exampleValue,
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
		fmt.Println(outputError, err)

		return
	}

	fmt.Println(outputConnectorID, req.ConnectorID.Value())
	fmt.Println("TransactionID:", req.TransactionID.Value())
	// Output:
	// ConnectorID: 1
	// TransactionID: 42
}

// ExampleReq_withOptionalFields demonstrates creating a MeterValues.req
// message with all optional SampledValue fields populated.
func ExampleReq_withOptionalFields() {
	context := "Sample.Periodic"
	format := "Raw"
	measurand := "Energy.Active.Import.Register"
	phase := "L1"
	location := "Outlet"
	unit := "Wh"

	input := metervalues.ReqInput{
		ConnectorID:   exampleConnectorID,
		TransactionID: nil,
		MeterValue: []types.MeterValueInput{
			{
				Timestamp: exampleTimestamp,
				SampledValue: []types.SampledValueInput{
					{
						Value:     exampleValue,
						Context:   &context,
						Format:    &format,
						Measurand: &measurand,
						Phase:     &phase,
						Location:  &location,
						Unit:      &unit,
					},
				},
			},
		},
	}

	req, err := metervalues.Req(input)
	if err != nil {
		fmt.Println(outputError, err)

		return
	}

	fmt.Println(outputConnectorID, req.ConnectorID.Value())
	fmt.Println("Value:", req.MeterValue[0].SampledValue()[0].Value().Value())
	// Output:
	// ConnectorID: 1
	// Value: 12500
}

// ExampleReq_emptyMeterValue demonstrates the error when MeterValue is empty.
func ExampleReq_emptyMeterValue() {
	input := metervalues.ReqInput{
		ConnectorID:   exampleConnectorID,
		TransactionID: nil,
		MeterValue:    []types.MeterValueInput{},
	}

	_, err := metervalues.Req(input)
	if err != nil {
		fmt.Println("Error: MeterValue cannot be empty")
	}
	// Output:
	// Error: MeterValue cannot be empty
}
