package triggermessage_test

import (
	"fmt"

	"github.com/evcoreco/ocpp16messages/triggermessage"
)

const (
	labelRequestedMessage = "RequestedMessage:"
	labelConnectorID      = "ConnectorID:"
	connectorIdZero       = 0
	connectorIdOne        = 1
)

// ExampleReq demonstrates creating a valid TriggerMessage.req message
// to trigger a Heartbeat message.
func ExampleReq() {
	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorID:      nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelRequestedMessage, req.RequestedMessage.String())
	// Output:
	// RequestedMessage: Heartbeat
}

// ExampleReq_withConnectorID demonstrates creating a TriggerMessage.req
// message with an optional connectorId.
func ExampleReq_withConnectorID() {
	connectorId := connectorIdOne

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "StatusNotification",
		ConnectorID:      &connectorId,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelRequestedMessage, req.RequestedMessage.String())
	fmt.Println(labelConnectorID, req.ConnectorID.Value())
	// Output:
	// RequestedMessage: StatusNotification
	// ConnectorID: 1
}

// ExampleReq_metervalues demonstrates triggering a MeterValues message.
func ExampleReq_metervalues() {
	connectorId := connectorIdZero

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "MeterValues",
		ConnectorID:      &connectorId,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelRequestedMessage, req.RequestedMessage.String())
	fmt.Println(labelConnectorID, req.ConnectorID.Value())
	// Output:
	// RequestedMessage: MeterValues
	// ConnectorID: 0
}

// ExampleReq_invalidMessage demonstrates the error returned when
// an invalid message trigger is provided.
func ExampleReq_invalidMessage() {
	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Unknown",
		ConnectorID:      nil,
	})
	if err != nil {
		fmt.Println("Error: invalid requestedMessage")
	}
	// Output:
	// Error: invalid requestedMessage
}

// ExampleReq_emptyMessage demonstrates the error returned when
// an empty message trigger is provided.
func ExampleReq_emptyMessage() {
	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "",
		ConnectorID:      nil,
	})
	if err != nil {
		fmt.Println("Error: invalid requestedMessage")
	}
	// Output:
	// Error: invalid requestedMessage
}
