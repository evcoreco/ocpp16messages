package testsjson_test

import (
	"testing"

	"github.com/evcoreco/ocpp16messages/authorize"
)

func TestAuthorizeReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := authorize.Req(authorize.ReqInput{IDTag: "RFID-TAG-12345"})
	if err != nil {
		t.Fatalf("authorize.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestAuthorizeConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		t.Fatalf("authorize.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
