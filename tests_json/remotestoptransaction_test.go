package testsjson_test

import (
	"testing"

	rst "github.com/evcoreco/ocpp16messages/remotestoptransaction"
)

func TestRemoteStopTransactionReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := rst.Req(rst.ReqInput{TransactionID: 12345})
	if err != nil {
		t.Fatalf("remotestoptransaction.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestRemoteStopTransactionConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := rst.Conf(rst.ConfInput{Status: "Accepted"})
	if err != nil {
		t.Fatalf("remotestoptransaction.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
