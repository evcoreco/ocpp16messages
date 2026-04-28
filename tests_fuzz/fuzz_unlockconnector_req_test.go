//go:build fuzz

package fuzz

import (
	"errors"
	"math"
	"testing"

	"github.com/evcoreco/ocpp16messages/unlockconnector"
	types "github.com/evcoreco/ocpp16types"
)

func FuzzUnlockConnectorReq(f *testing.F) {
	f.Add(1)
	f.Add(0)
	f.Add(-1)
	f.Add(math.MaxUint16 + 1)

	f.Fuzz(func(t *testing.T, connectorId int) {
		req, err := unlockconnector.Req(unlockconnector.ReqInput{
			ConnectorID: connectorId,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) {
				t.Fatalf("error = %v, want wrapping ErrInvalidValue", err)
			}

			return
		}

		if connectorId <= 0 || connectorId > math.MaxUint16 {
			t.Fatalf("Req succeeded with connectorId=%d", connectorId)
		}

		if got := req.ConnectorID.Value(); got != uint16(connectorId) {
			t.Fatalf("ConnectorID = %d, want %d", got, connectorId)
		}
	})
}
