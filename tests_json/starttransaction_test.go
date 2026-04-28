package testsjson_test

import (
	"testing"

	"github.com/evcoreco/ocpp16messages/starttransaction"
)

func TestStartTransactionReq_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	req, err := starttransaction.Req(starttransaction.ReqInput{
		ConnectorID:   1,
		IDTag:         "RFID-TAG-12345",
		MeterStart:    1000,
		Timestamp:     "2025-01-15T10:30:00Z",
		ReservationID: nil,
	})
	if err != nil {
		t.Fatalf("starttransaction.Req: %v", err)
	}

	assertAllFieldsValid(t, req)
	roundTripJSON(t, req)
}

func TestStartTransactionConf_JSONRoundTrip(t *testing.T) {
	t.Parallel()

	conf, err := starttransaction.Conf(starttransaction.ConfInput{
		TransactionID: 12345,
		Status:        "Accepted",
		ExpiryDate:    nil,
		ParentIDTag:   nil,
	})
	if err != nil {
		t.Fatalf("starttransaction.Conf: %v", err)
	}

	assertAllFieldsValid(t, conf)
	roundTripJSON(t, conf)
}
