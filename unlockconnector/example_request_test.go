package unlockconnector_test

import (
	"fmt"

	"github.com/evcoreco/ocpp16messages/unlockconnector"
)

const (
	exampleConnectorID     = 1
	exampleConnectorIDZero = 0
)

// ExampleReq demonstrates creating a valid UnlockConnector.req message.
func ExampleReq() {
	req, err := unlockconnector.Req(unlockconnector.ReqInput{
		ConnectorID: exampleConnectorID,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("ConnectorID:", req.ConnectorID.Value())
	// Output:
	// ConnectorID: 1
}

// ExampleReq_zeroConnectorID demonstrates the error returned when
// connectorId is zero (connector 0 refers to the Charge Point itself).
func ExampleReq_zeroConnectorID() {
	_, err := unlockconnector.Req(unlockconnector.ReqInput{
		ConnectorID: exampleConnectorIDZero,
	})
	if err != nil {
		fmt.Println("connectorId: invalid value")
	}
	// Output:
	// connectorId: invalid value
}
