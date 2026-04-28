//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	tm "github.com/evcoreco/ocpp16messages/triggermessage"
	types "github.com/evcoreco/ocpp16types"
)

func FuzzTriggerMessageReq(f *testing.F) {
	f.Add(types.MessageTriggerBootNotification.String(), false, 0)
	f.Add(types.MessageTriggerHeartbeat.String(), true, 1)
	f.Add("invalid-trigger", false, 0)
	f.Add(types.MessageTriggerMeterValues.String(), true, -1)

	f.Fuzz(func(
		t *testing.T,
		requestedMessage string,
		hasConnectorID bool,
		connectorId int,
	) {
		if len(requestedMessage) > maxFuzzStringLen {
			t.Skip()
		}

		var connectorIdPtr *int
		if hasConnectorID {
			connectorIdPtr = &connectorId
		}

		req, err := tm.Req(tm.ReqInput{
			RequestedMessage: requestedMessage,
			ConnectorID:      connectorIdPtr,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if !req.RequestedMessage.IsValid() {
			t.Fatalf("RequestedMessage = %q, want valid", req.RequestedMessage.String())
		}

		if hasConnectorID {
			if req.ConnectorID == nil {
				t.Fatal("ConnectorID = nil, want non-nil")
			}
			if connectorId < 0 || connectorId > math.MaxUint16 {
				t.Fatalf("Req succeeded with connectorId=%d", connectorId)
			}
			if got := req.ConnectorID.Value(); got != uint16(connectorId) {
				t.Fatalf("ConnectorID = %d, want %d", got, connectorId)
			}
		} else if req.ConnectorID != nil {
			t.Fatal("ConnectorID != nil, want nil")
		}
	})
}
