package starttransaction_test

import (
	"fmt"

	"github.com/evcoreco/ocpp16messages/starttransaction"
)

const testReservationID = 42

// ExampleReq demonstrates creating a valid StartTransaction.req message
// with all required fields.
func ExampleReq() {
	req, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   1,
		IDTag:         "RFID-TAG-12345",
		MeterStart:    1000,
		Timestamp:     "2025-01-15T10:30:00Z",
		ReservationID: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ConnectorID:", req.ConnectorID.Value())
	fmt.Println("IDTag:", req.IDTag.String())
	fmt.Println("MeterStart:", req.MeterStart.Value())
	// Output:
	// ConnectorID: 1
	// IDTag: RFID-TAG-12345
	// MeterStart: 1000
}

// ExampleReq_withReservation demonstrates creating a StartTransaction.req
// message that includes a reservation ID.
func ExampleReq_withReservation() {
	reservationId := testReservationID

	req, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   2,
		IDTag:         "RFID-TAG-67890",
		MeterStart:    500,
		Timestamp:     "2025-01-15T11:00:00Z",
		ReservationID: &reservationId,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ConnectorID:", req.ConnectorID.Value())
	fmt.Println("IDTag:", req.IDTag.String())
	fmt.Println("HasReservation:", req.ReservationID != nil)
	fmt.Println("ReservationID:", req.ReservationID.Value())
	// Output:
	// ConnectorID: 2
	// IDTag: RFID-TAG-67890
	// HasReservation: true
	// ReservationID: 42
}

// ExampleReq_emptyIDTag demonstrates the error returned when
// an empty ID tag is provided.
func ExampleReq_emptyIDTag() {
	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   1,
		IDTag:         "",
		MeterStart:    1000,
		Timestamp:     "2025-01-15T10:30:00Z",
		ReservationID: nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTag: value cannot be empty
}

// ExampleReq_invalidTimestamp demonstrates the error returned when
// an invalid timestamp is provided.
func ExampleReq_invalidTimestamp() {
	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   1,
		IDTag:         "RFID-TAG-12345",
		MeterStart:    1000,
		Timestamp:     "not-a-timestamp",
		ReservationID: nil,
	})
	if err != nil {
		fmt.Println("Invalid timestamp error")
	}
	// Output:
	// Invalid timestamp error
}

// ExampleReq_multipleErrors demonstrates that all validation errors
// are returned at once, not just the first one encountered.
func ExampleReq_multipleErrors() {
	_, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   -1,
		IDTag:         "",
		MeterStart:    -1,
		Timestamp:     "invalid",
		ReservationID: nil,
	})
	if err != nil {
		fmt.Println("Multiple errors")
	}
	// Output:
	// Multiple errors
}
