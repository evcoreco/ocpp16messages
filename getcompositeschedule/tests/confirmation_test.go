package getcompositeschedule_test

import (
	"strings"
	"testing"

	gcs "github.com/evcoreco/ocpp16messages/getcompositeschedule"
	types "github.com/evcoreco/ocpp16types"
)

const (
	errStatus           = "status"
	errConnectorIDConf  = "connectorId"
	errScheduleStart    = "scheduleStart"
	errChargingSchedule = "chargingSchedule"

	validTimestamp    = "2025-01-15T10:00:00Z"
	invalidTimestamp  = "invalid-timestamp"
	durationConfValue = 3600
	limitConfValue    = 32.0

	connectorIdNotNil      = "ConnectorID should not be nil"
	scheduleStartNotNil    = "ScheduleStart should not be nil"
	chargingScheduleNotNil = "ChargingSchedule should not be nil"
)

func intPtr(v int) *int {
	return &v
}

func TestConf_Valid_AcceptedOnly(t *testing.T) {
	t.Parallel()

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorID:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.GetCompositeScheduleStatusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			types.GetCompositeScheduleStatusAccepted,
			conf.Status,
		)
	}

	if conf.ConnectorID != nil {
		t.Errorf("ConnectorID should be nil, got %v", conf.ConnectorID)
	}

	if conf.ScheduleStart != nil {
		t.Errorf("ScheduleStart should be nil, got %v", conf.ScheduleStart)
	}

	if conf.ChargingSchedule != nil {
		t.Errorf(
			"ChargingSchedule should be nil, got %v",
			conf.ChargingSchedule,
		)
	}
}

func TestConf_Valid_RejectedOnly(t *testing.T) {
	t.Parallel()

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:           "Rejected",
		ConnectorID:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.GetCompositeScheduleStatusRejected {
		t.Errorf(
			types.ErrorMismatch,
			types.GetCompositeScheduleStatusRejected,
			conf.Status,
		)
	}
}

func TestConf_Valid_WithConnectorID(t *testing.T) {
	t.Parallel()

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorID:      intPtr(valueOne),
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.ConnectorID == nil {
		t.Fatal(connectorIdNotNil)
	}

	if conf.ConnectorID.Value() != valueOne {
		t.Errorf(types.ErrorMismatchValue, valueOne, conf.ConnectorID.Value())
	}
}

func TestConf_Valid_WithConnectorIDZero(t *testing.T) {
	t.Parallel()

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorID:      intPtr(valueZero),
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.ConnectorID == nil {
		t.Fatal(connectorIdNotNil)
	}

	if conf.ConnectorID.Value() != valueZero {
		t.Errorf(types.ErrorMismatchValue, valueZero, conf.ConnectorID.Value())
	}
}

func TestConf_Valid_WithScheduleStart(t *testing.T) {
	t.Parallel()

	scheduleStart := validTimestamp

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorID:      nil,
		ScheduleStart:    &scheduleStart,
		ChargingSchedule: nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.ScheduleStart == nil {
		t.Fatal(scheduleStartNotNil)
	}
}

func TestConf_Valid_WithChargingSchedule(t *testing.T) {
	t.Parallel()

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:        "Accepted",
		ConnectorID:   nil,
		ScheduleStart: nil,
		ChargingSchedule: &types.ChargingScheduleInput{
			Duration:         nil,
			ChargingRateUnit: "W",
			ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
				{
					StartPeriod:  valueZero,
					Limit:        limitConfValue,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.ChargingSchedule == nil {
		t.Fatal(chargingScheduleNotNil)
	}

	if conf.ChargingSchedule.ChargingRateUnit() != types.ChargingRateUnitWatts {
		t.Errorf(
			types.ErrorMismatch,
			types.ChargingRateUnitWatts,
			conf.ChargingSchedule.ChargingRateUnit(),
		)
	}
}

func TestConf_Valid_WithAllFields(t *testing.T) {
	t.Parallel()

	scheduleStart := validTimestamp
	duration := durationConfValue

	conf, err := gcs.Conf(gcs.ConfInput{
		Status:        "Accepted",
		ConnectorID:   intPtr(valueOne),
		ScheduleStart: &scheduleStart,
		ChargingSchedule: &types.ChargingScheduleInput{
			Duration:         &duration,
			ChargingRateUnit: "A",
			ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
				{
					StartPeriod:  valueZero,
					Limit:        limitConfValue,
					NumberPhases: nil,
				},
			},
			MinChargingRate: nil,
			StartSchedule:   nil,
		},
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if conf.Status != types.GetCompositeScheduleStatusAccepted {
		t.Errorf(
			types.ErrorMismatch,
			types.GetCompositeScheduleStatusAccepted,
			conf.Status,
		)
	}

	if conf.ConnectorID == nil {
		t.Fatal(connectorIdNotNil)
	}

	if conf.ScheduleStart == nil {
		t.Fatal(scheduleStartNotNil)
	}

	if conf.ChargingSchedule == nil {
		t.Fatal(chargingScheduleNotNil)
	}
}

func TestConf_Invalid_EmptyStatus(t *testing.T) {
	t.Parallel()

	_, err := gcs.Conf(gcs.ConfInput{
		Status:           "",
		ConnectorID:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_Invalid_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := gcs.Conf(gcs.ConfInput{
		Status:           "Invalid",
		ConnectorID:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_Invalid_LowercaseStatus(t *testing.T) {
	t.Parallel()

	_, err := gcs.Conf(gcs.ConfInput{
		Status:           "accepted",
		ConnectorID:      nil,
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase status")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}
}

func TestConf_Invalid_NegativeConnectorID(t *testing.T) {
	t.Parallel()

	_, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorID:      intPtr(valueNegative),
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative ConnectorID")
	}

	if !strings.Contains(err.Error(), errConnectorIDConf) {
		t.Errorf(types.ErrorWantContains, err, errConnectorIDConf)
	}
}

func TestConf_Invalid_ConnectorIDExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorID:      intPtr(valueExceedsMax),
		ScheduleStart:    nil,
		ChargingSchedule: nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "ConnectorID exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorIDConf) {
		t.Errorf(types.ErrorWantContains, err, errConnectorIDConf)
	}
}

func TestConf_Invalid_InvalidScheduleStart(t *testing.T) {
	t.Parallel()

	scheduleStart := invalidTimestamp

	_, err := gcs.Conf(gcs.ConfInput{
		Status:           "Accepted",
		ConnectorID:      nil,
		ScheduleStart:    &scheduleStart,
		ChargingSchedule: nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid scheduleStart")
	}

	if !strings.Contains(err.Error(), errScheduleStart) {
		t.Errorf(types.ErrorWantContains, err, errScheduleStart)
	}
}

func TestConf_Invalid_InvalidChargingSchedule(t *testing.T) {
	t.Parallel()

	_, err := gcs.Conf(gcs.ConfInput{
		Status:        "Accepted",
		ConnectorID:   nil,
		ScheduleStart: nil,
		ChargingSchedule: &types.ChargingScheduleInput{
			Duration:               nil,
			ChargingRateUnit:       "X",
			ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{},
			MinChargingRate:        nil,
			StartSchedule:          nil,
		},
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid chargingSchedule")
	}

	if !strings.Contains(err.Error(), errChargingSchedule) {
		t.Errorf(types.ErrorWantContains, err, errChargingSchedule)
	}
}

func TestConf_Invalid_MultipleErrors(t *testing.T) {
	t.Parallel()

	scheduleStart := invalidTimestamp

	_, err := gcs.Conf(gcs.ConfInput{
		Status:        "Invalid",
		ConnectorID:   intPtr(valueNegative),
		ScheduleStart: &scheduleStart,
		ChargingSchedule: &types.ChargingScheduleInput{
			Duration:               nil,
			ChargingRateUnit:       "X",
			ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{},
			MinChargingRate:        nil,
			StartSchedule:          nil,
		},
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "multiple invalid fields")
	}

	if !strings.Contains(err.Error(), errStatus) {
		t.Errorf(types.ErrorWantContains, err, errStatus)
	}

	if !strings.Contains(err.Error(), errConnectorIDConf) {
		t.Errorf(types.ErrorWantContains, err, errConnectorIDConf)
	}

	if !strings.Contains(err.Error(), errScheduleStart) {
		t.Errorf(types.ErrorWantContains, err, errScheduleStart)
	}

	if !strings.Contains(err.Error(), errChargingSchedule) {
		t.Errorf(types.ErrorWantContains, err, errChargingSchedule)
	}
}
