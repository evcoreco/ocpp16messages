//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	rst "github.com/evcoreco/ocpp16messages/remotestarttransaction"
	types "github.com/evcoreco/ocpp16types"
)

func FuzzRemoteStartTransactionReq(f *testing.F) {
	f.Add("RFID-ABC123", false, 0)
	f.Add("", false, 0)
	f.Add("RFID-ABC123", true, 1)
	f.Add("RFID-ABC123", true, -1)

	f.Fuzz(func(t *testing.T, idTag string, hasConnectorID bool, connectorId int) {
		if len(idTag) > maxFuzzStringLen {
			t.Skip()
		}

		var connectorIdPtr *int
		if hasConnectorID {
			connectorIdPtr = &connectorId
		}

		req, err := rst.Req(rst.ReqInput{
			IDTag:       idTag,
			ConnectorID: connectorIdPtr,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) && !errors.Is(err, types.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if idTag == "" {
			t.Fatal("Req succeeded with empty IDTag")
		}

		if req.IDTag.String() != idTag {
			t.Fatalf("IDTag = %q, want %q", req.IDTag.String(), idTag)
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
