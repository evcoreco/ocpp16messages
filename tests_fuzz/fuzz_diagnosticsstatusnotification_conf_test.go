//go:build fuzz

package fuzz

import (
	"testing"

	"github.com/evcoreco/ocpp16messages/diagnosticsstatusnotification"
)

func FuzzDiagnosticsStatusNotificationConf(f *testing.F) {
	f.Add(uint8(0))

	f.Fuzz(func(t *testing.T, _ uint8) {
		_, err := diagnosticsstatusnotification.Conf(
			diagnosticsstatusnotification.ConfInput{},
		)
		if err != nil {
			t.Fatalf("error = %v, want nil", err)
		}
	})
}
