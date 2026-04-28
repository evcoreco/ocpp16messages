//go:build bench

package benchmark

import (
	"testing"

	types "github.com/evcoreco/ocpp16types"
)

func BenchmarkNewChargingSchedulePeriod_WithNumberPhases(b *testing.B) {
	b.ReportAllocs()

	phases := 3
	input := types.ChargingSchedulePeriodInput{
		StartPeriod:  0,
		Limit:        16,
		NumberPhases: &phases,
	}

	for i := 0; i < b.N; i++ {
		if _, err := types.NewChargingSchedulePeriod(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetConfigurationNewKeyValue_ValueNil(b *testing.B) {
	b.ReportAllocs()

	input := types.KeyValueInput{
		Key:      "HeartbeatInterval",
		Readonly: false,
		Value:    nil,
	}

	for i := 0; i < b.N; i++ {
		if _, err := types.NewKeyValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGetConfigurationNewKeyValue_ValueSet(b *testing.B) {
	b.ReportAllocs()

	value := "60"
	input := types.KeyValueInput{
		Key:      "HeartbeatInterval",
		Readonly: false,
		Value:    &value,
	}

	for i := 0; i < b.N; i++ {
		if _, err := types.NewKeyValue(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSendLocalListNewAuthorizationData_WithIDTagInfo(b *testing.B) {
	b.ReportAllocs()

	expiry := sampleTimestamp
	parentIDTag := "PARENT-1"

	input := types.AuthorizationDataInput{
		IDTag: "TAG-1",
		IDTagInfo: &types.IDTagInfoInput{
			Status:      types.AuthorizationStatusAccepted.String(),
			ExpiryDate:  &expiry,
			ParentIDTag: &parentIDTag,
		},
	}

	for i := 0; i < b.N; i++ {
		if _, err := types.NewAuthorizationData(input); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSetChargingProfileNewChargingProfile(b *testing.B) {
	b.ReportAllocs()

	duration := 60
	scheduleStart := sampleTimestamp
	minChargingRate := 0.0
	periods := []types.ChargingSchedulePeriodInput{{StartPeriod: 0, Limit: 16}}

	scheduleInput := types.ChargingScheduleInput{
		Duration:               &duration,
		ChargingRateUnit:       types.ChargingRateUnitWatts.String(),
		ChargingSchedulePeriod: periods,
		MinChargingRate:        &minChargingRate,
		StartSchedule:          &scheduleStart,
	}

	input := types.ChargingProfileInput{
		ChargingProfileID:      1,
		TransactionID:          nil,
		StackLevel:             0,
		ChargingProfilePurpose: types.TxProfile.String(),
		ChargingProfileKind:    types.ChargingProfileKindAbsolute.String(),
		RecurrencyKind:         nil,
		ValidFrom:              nil,
		ValidTo:                nil,
		ChargingSchedule:       scheduleInput,
	}

	for i := 0; i < b.N; i++ {
		if _, err := types.NewChargingProfile(input); err != nil {
			b.Fatal(err)
		}
	}
}
