package reservenow_test

import (
	"fmt"

	"github.com/evcoreco/ocpp16messages/reservenow"
)

const (
	exampleReservationID = 1
	exampleConnectorID   = 1
	exampleIDTag         = "RFID-TAG-12345"
	exampleExpiryDate    = "2025-01-15T10:00:00Z"
	exampleParentIDTag   = "PARENT-12345"
)

// ExampleReq demonstrates creating a valid ReserveNow.req message with only
// the required fields.
func ExampleReq() {
	req, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: exampleReservationID,
		ConnectorID:   exampleConnectorID,
		IDTag:         exampleIDTag,
		ExpiryDate:    exampleExpiryDate,
		ParentIDTag:   nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ReservationID:", req.ReservationID.Value())
	fmt.Println("ConnectorID:", req.ConnectorID.Value())
	fmt.Println("IDTag:", req.IDTag.Value())
	// Output:
	// ReservationID: 1
	// ConnectorID: 1
	// IDTag: RFID-TAG-12345
}

// ExampleReq_withParentIDTag demonstrates creating a ReserveNow.req message
// with all fields including the optional parentIDTag.
func ExampleReq_withParentIDTag() {
	parentIDTag := exampleParentIDTag

	req, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: exampleReservationID,
		ConnectorID:   exampleConnectorID,
		IDTag:         exampleIDTag,
		ExpiryDate:    exampleExpiryDate,
		ParentIDTag:   &parentIDTag,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IDTag:", req.IDTag.Value())
	fmt.Println("ParentIDTag:", req.ParentIDTag.Value())
	// Output:
	// IDTag: RFID-TAG-12345
	// ParentIDTag: PARENT-12345
}

// ExampleReq_emptyIDTag demonstrates the error returned when an empty idTag
// is provided.
func ExampleReq_emptyIDTag() {
	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: exampleReservationID,
		ConnectorID:   exampleConnectorID,
		IDTag:         "",
		ExpiryDate:    exampleExpiryDate,
		ParentIDTag:   nil,
	})
	if err != nil {
		fmt.Println("idTag: value cannot be empty")
	}
	// Output:
	// idTag: value cannot be empty
}

// ExampleReq_invalidExpiryDate demonstrates the error returned when an invalid
// expiry date is provided.
func ExampleReq_invalidExpiryDate() {
	_, err := reservenow.Req(reservenow.ReqInput{
		ReservationID: exampleReservationID,
		ConnectorID:   exampleConnectorID,
		IDTag:         exampleIDTag,
		ExpiryDate:    "invalid-date",
		ParentIDTag:   nil,
	})
	if err != nil {
		fmt.Println("expiryDate: invalid value")
	}
	// Output:
	// expiryDate: invalid value
}
