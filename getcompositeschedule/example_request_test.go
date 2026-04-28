package getcompositeschedule_test

import (
	"fmt"

	"github.com/evcoreco/ocpp16messages/getcompositeschedule"
)

const (
	exampleConnectorIDOne    = 1
	exampleConnectorIDZero   = 0
	exampleDurationThreeHund = 300
	exampleDurationSixHund   = 600
	exampleNegativeValue     = -1

	outConnectorID = "ConnectorID:"
	outDuration    = "Duration:"
)

// ExampleReq demonstrates creating a GetCompositeSchedule.req message
// with required fields only.
func ExampleReq() {
	req, err := getcompositeschedule.Req(getcompositeschedule.ReqInput{
		ConnectorID:      exampleConnectorIDOne,
		Duration:         exampleDurationThreeHund,
		ChargingRateUnit: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outConnectorID, req.ConnectorID.Value())
	fmt.Println(outDuration, req.Duration.Value())
	// Output:
	// ConnectorID: 1
	// Duration: 300
}

// ExampleReq_withChargingRateUnit demonstrates creating a
// GetCompositeSchedule.req message with the optional ChargingRateUnit field.
func ExampleReq_withChargingRateUnit() {
	unit := "W"

	req, err := getcompositeschedule.Req(getcompositeschedule.ReqInput{
		ConnectorID:      exampleConnectorIDOne,
		Duration:         exampleDurationSixHund,
		ChargingRateUnit: &unit,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outConnectorID, req.ConnectorID.Value())
	fmt.Println(outDuration, req.Duration.Value())
	fmt.Println("ChargingRateUnit:", req.ChargingRateUnit.String())
	// Output:
	// ConnectorID: 1
	// Duration: 600
	// ChargingRateUnit: W
}

// ExampleReq_entireChargePoint demonstrates requesting a composite schedule
// for the entire Charge Point by using ConnectorID 0.
func ExampleReq_entireChargePoint() {
	req, err := getcompositeschedule.Req(getcompositeschedule.ReqInput{
		ConnectorID:      exampleConnectorIDZero,
		Duration:         exampleDurationThreeHund,
		ChargingRateUnit: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(outConnectorID, req.ConnectorID.Value())
	fmt.Println(outDuration, req.Duration.Value())
	// Output:
	// ConnectorID: 0
	// Duration: 300
}

// ExampleReq_invalidConnectorID demonstrates the error returned when
// a negative ConnectorID is provided.
func ExampleReq_invalidConnectorID() {
	_, err := getcompositeschedule.Req(getcompositeschedule.ReqInput{
		ConnectorID:      exampleNegativeValue,
		Duration:         exampleDurationThreeHund,
		ChargingRateUnit: nil,
	})
	if err != nil {
		fmt.Println("Error: invalid connector ID")
	}
	// Output:
	// Error: invalid connector ID
}

// ExampleReq_invalidChargingRateUnit demonstrates the error returned when
// an invalid ChargingRateUnit is provided.
func ExampleReq_invalidChargingRateUnit() {
	unit := "X"

	_, err := getcompositeschedule.Req(getcompositeschedule.ReqInput{
		ConnectorID:      exampleConnectorIDOne,
		Duration:         exampleDurationThreeHund,
		ChargingRateUnit: &unit,
	})
	if err != nil {
		fmt.Println("Error: invalid charging rate unit")
	}
	// Output:
	// Error: invalid charging rate unit
}
