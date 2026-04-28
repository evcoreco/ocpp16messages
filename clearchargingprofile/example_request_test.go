package clearchargingprofile_test

import (
	"fmt"

	"github.com/evcoreco/ocpp16messages/clearchargingprofile"
)

const (
	exampleIdValue         = 123
	exampleConnectorID     = 1
	exampleStackLevelZero  = 0
	exampleIdOne           = 1
	exampleConnectorIDTwo  = 2
	exampleStackLevelThree = 3
	exampleNegativeId      = -1
)

// ExampleReq demonstrates creating a ClearChargingProfile.req message
// with no optional fields (clears all charging profiles).
func ExampleReq() {
	_, err := clearchargingprofile.Req(clearchargingprofile.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ClearChargingProfile.req created successfully")
	// Output:
	// ClearChargingProfile.req created successfully
}

// ExampleReq_withId demonstrates creating a ClearChargingProfile.req message
// with a specific charging profile ID.
func ExampleReq_withId() {
	profileId := exampleIdValue

	req, err := clearchargingprofile.Req(clearchargingprofile.ReqInput{
		Id:                     &profileId,
		ConnectorID:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Id:", req.Id.Value())
	// Output:
	// Id: 123
}

// ExampleReq_withConnectorID demonstrates creating a ClearChargingProfile.req
// message with a connector ID.
func ExampleReq_withConnectorID() {
	connectorId := exampleConnectorID

	req, err := clearchargingprofile.Req(clearchargingprofile.ReqInput{
		Id:                     nil,
		ConnectorID:            &connectorId,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ConnectorID:", req.ConnectorID.Value())
	// Output:
	// ConnectorID: 1
}

// ExampleReq_withPurpose demonstrates creating a ClearChargingProfile.req
// message with a charging profile purpose.
func ExampleReq_withPurpose() {
	purpose := "TxProfile"

	req, err := clearchargingprofile.Req(clearchargingprofile.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: &purpose,
		StackLevel:             nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Purpose:", req.ChargingProfilePurpose.String())
	// Output:
	// Purpose: TxProfile
}

// ExampleReq_withStackLevel demonstrates creating a ClearChargingProfile.req
// message with a stack level.
func ExampleReq_withStackLevel() {
	stackLevel := exampleStackLevelZero

	req, err := clearchargingprofile.Req(clearchargingprofile.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             &stackLevel,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("StackLevel:", req.StackLevel.Value())
	// Output:
	// StackLevel: 0
}

// ExampleReq_withAllFields demonstrates creating a ClearChargingProfile.req
// message with all optional fields.
func ExampleReq_withAllFields() {
	profileId := exampleIdOne
	connectorId := exampleConnectorIDTwo
	purpose := "TxDefaultProfile"
	stackLevel := exampleStackLevelThree

	req, err := clearchargingprofile.Req(clearchargingprofile.ReqInput{
		Id:                     &profileId,
		ConnectorID:            &connectorId,
		ChargingProfilePurpose: &purpose,
		StackLevel:             &stackLevel,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("Id:", req.Id.Value())
	fmt.Println("ConnectorID:", req.ConnectorID.Value())
	fmt.Println("Purpose:", req.ChargingProfilePurpose.String())
	fmt.Println("StackLevel:", req.StackLevel.Value())
	// Output:
	// Id: 1
	// ConnectorID: 2
	// Purpose: TxDefaultProfile
	// StackLevel: 3
}

// ExampleReq_invalidPurpose demonstrates the error returned when
// an invalid charging profile purpose is provided.
func ExampleReq_invalidPurpose() {
	purpose := "Invalid"

	_, err := clearchargingprofile.Req(clearchargingprofile.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: &purpose,
		StackLevel:             nil,
	})
	if err != nil {
		fmt.Println("Error: invalid charging profile purpose")
	}
	// Output:
	// Error: invalid charging profile purpose
}

// ExampleReq_negativeId demonstrates the error returned when
// a negative ID is provided.
func ExampleReq_negativeId() {
	profileId := exampleNegativeId

	_, err := clearchargingprofile.Req(clearchargingprofile.ReqInput{
		Id:                     &profileId,
		ConnectorID:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		fmt.Println("Error: invalid ID")
	}
	// Output:
	// Error: invalid ID
}
