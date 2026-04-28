package remotestoptransaction_test

import (
	"fmt"

	rst "github.com/evcoreco/ocpp16messages/remotestoptransaction"
)

const (
	exampleTransactionID    = 12345
	exampleTransactionIDNeg = -1
)

// ExampleReq demonstrates creating a valid RemoteStopTransaction.req message.
func ExampleReq() {
	req, err := rst.Req(rst.ReqInput{TransactionID: exampleTransactionID})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("TransactionID:", req.TransactionID.Value())
	// Output:
	// TransactionID: 12345
}

// ExampleReq_negativeTransactionID demonstrates the error returned when
// a negative transactionId is provided.
func ExampleReq_negativeTransactionID() {
	_, err := rst.Req(rst.ReqInput{TransactionID: exampleTransactionIDNeg})
	if err != nil {
		fmt.Println("transactionId: invalid value")
	}
	// Output:
	// transactionId: invalid value
}
