//go:build bench

package benchmark

import (
	"testing"

	stt "github.com/evcoreco/ocpp16messages/starttransaction"
)

var (
	sinkStartTxCustom    stt.ReqMessage
	sinkStartTxPrimitive primitiveStartTransactionReq
)

func BenchmarkStartTransactionReq_Custom(b *testing.B) {
	b.ReportAllocs()

	reservationId := 42
	input := stt.ReqInput{
		ConnectorID:   1,
		IDTag:         "TAG-123",
		MeterStart:    100,
		Timestamp:     benchmarkTimestamp,
		ReservationID: &reservationId,
	}

	for i := 0; i < b.N; i++ {
		message, err := stt.Req(input)
		if err != nil {
			b.Fatal(err)
		}

		sinkStartTxCustom = message
	}
}

func BenchmarkStartTransactionReq_PrimitiveDirect(b *testing.B) {
	b.ReportAllocs()

	reservationId := 42
	input := primitiveStartTransactionReq{
		ConnectorID:   1,
		IDTag:         "TAG-123",
		MeterStart:    100,
		Timestamp:     benchmarkTimestamp,
		ReservationID: &reservationId,
	}

	for i := 0; i < b.N; i++ {
		sinkStartTxPrimitive = input
	}
}

func BenchmarkStartTransactionReq_PrimitiveValidated(b *testing.B) {
	b.ReportAllocs()

	reservationId := 42
	input := primitiveStartTransactionReq{
		ConnectorID:   1,
		IDTag:         "TAG-123",
		MeterStart:    100,
		Timestamp:     benchmarkTimestamp,
		ReservationID: &reservationId,
	}

	for i := 0; i < b.N; i++ {
		if err := validatePrimitiveStartTransactionReq(input); err != nil {
			b.Fatal(err)
		}

		sinkStartTxPrimitive = input
	}
}
