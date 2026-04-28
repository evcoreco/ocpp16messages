package remotestarttransaction_test

import (
	"fmt"

	rst "github.com/evcoreco/ocpp16messages/remotestarttransaction"
)

const (
	connectorIdOne        = 1
	connectorIdNegative   = -1
	testExampleValidIDTag = "RFID-TAG-12345"
)

// ExampleReq demonstrates creating a valid RemoteStartTransaction.req message
// with only the required idTag field.
func ExampleReq() {
	req, err := rst.Req(rst.ReqInput{
		IDTag:       testExampleValidIDTag,
		ConnectorID: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IDTag:", req.IDTag.Value())
	// Output:
	// IDTag: RFID-TAG-12345
}

// ExampleReq_withConnectorID demonstrates creating a RemoteStartTransaction.req
// message with both idTag and connectorId.
func ExampleReq_withConnectorID() {
	connectorId := connectorIdOne

	req, err := rst.Req(rst.ReqInput{
		IDTag:       testExampleValidIDTag,
		ConnectorID: &connectorId,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IDTag:", req.IDTag.Value())
	fmt.Println("ConnectorID:", req.ConnectorID.Value())
	// Output:
	// IDTag: RFID-TAG-12345
	// ConnectorID: 1
}

// ExampleReq_emptyIDTag demonstrates the error returned when
// an empty idTag is provided.
func ExampleReq_emptyIDTag() {
	_, err := rst.Req(rst.ReqInput{IDTag: "", ConnectorID: nil})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTag: value cannot be empty
}

// ExampleReq_idTagTooLong demonstrates the error returned when
// the idTag exceeds the maximum length of 20 characters.
func ExampleReq_idTagTooLong() {
	// 23 chars, max is 20
	_, err := rst.Req(rst.ReqInput{
		IDTag:       "RFID-ABC123456789012345",
		ConnectorID: nil,
	})
	if err != nil {
		fmt.Println("idTag: exceeds maximum length")
	}
	// Output:
	// idTag: exceeds maximum length
}

// ExampleReq_invalidConnectorID demonstrates the error returned when
// the connectorId is negative.
func ExampleReq_invalidConnectorID() {
	connectorId := connectorIdNegative

	_, err := rst.Req(rst.ReqInput{
		IDTag:       testExampleValidIDTag,
		ConnectorID: &connectorId,
	})
	if err != nil {
		fmt.Println("connectorId: invalid value")
	}
	// Output:
	// connectorId: invalid value
}
