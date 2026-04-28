package setchargingprofile_test

import (
	"fmt"

	"github.com/evcoreco/ocpp16messages/setchargingprofile"
	types "github.com/evcoreco/ocpp16types"
)

const (
	exampleConnectorID     = 1
	exampleProfileId       = 100
	exampleStackLevel      = 0
	exampleStartPeriod     = 0
	exampleLimit           = 32.0
	fmtConnectorID         = "ConnectorID: %d\n"
	fmtChargingProfileID   = "ChargingProfileID: %d\n"
	fmtChargingProfileKind = "ChargingProfileKind: %s\n"
)

// ExampleReq demonstrates creating a valid SetChargingProfile.req message
// with a basic charging profile.
func ExampleReq() {
	req, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorID: exampleConnectorID,
		CsChargingProfiles: types.ChargingProfileInput{
			ChargingProfileID:      exampleProfileId,
			TransactionID:          nil,
			StackLevel:             exampleStackLevel,
			ChargingProfilePurpose: "TxDefaultProfile",
			ChargingProfileKind:    "Absolute",
			RecurrencyKind:         nil,
			ValidFrom:              nil,
			ValidTo:                nil,
			ChargingSchedule: types.ChargingScheduleInput{
				Duration:         nil,
				ChargingRateUnit: "W",
				ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
					{
						StartPeriod:  exampleStartPeriod,
						Limit:        exampleLimit,
						NumberPhases: nil,
					},
				},
				MinChargingRate: nil,
				StartSchedule:   nil,
			},
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Printf(fmtConnectorID, req.ConnectorID.Value())

	profileId := req.CsChargingProfiles.ChargingProfileID().Value()
	fmt.Printf(fmtChargingProfileID, profileId)

	profileKind := req.CsChargingProfiles.ChargingProfileKind().String()
	fmt.Printf(fmtChargingProfileKind, profileKind)
	// Output:
	// ConnectorID: 1
	// ChargingProfileID: 100
	// ChargingProfileKind: Absolute
}

// ExampleReq_chargePointMaxProfile demonstrates creating a
// SetChargingProfile.req message with a ChargePointMaxProfile purpose
// for limiting overall charge point power.
func ExampleReq_chargePointMaxProfile() {
	req, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorID: 0,
		CsChargingProfiles: types.ChargingProfileInput{
			ChargingProfileID:      exampleProfileId,
			TransactionID:          nil,
			StackLevel:             exampleStackLevel,
			ChargingProfilePurpose: "ChargePointMaxProfile",
			ChargingProfileKind:    "Absolute",
			RecurrencyKind:         nil,
			ValidFrom:              nil,
			ValidTo:                nil,
			ChargingSchedule: types.ChargingScheduleInput{
				Duration:         nil,
				ChargingRateUnit: "A",
				ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
					{
						StartPeriod:  exampleStartPeriod,
						Limit:        exampleLimit,
						NumberPhases: nil,
					},
				},
				MinChargingRate: nil,
				StartSchedule:   nil,
			},
		},
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	purpose := req.CsChargingProfiles.ChargingProfilePurpose().String()
	fmt.Println("Purpose:", purpose)
	// Output:
	// Purpose: ChargePointMaxProfile
}

// ExampleReq_invalidConnectorID demonstrates the error returned when
// a negative connector ID is provided.
func ExampleReq_invalidConnectorID() {
	_, err := setchargingprofile.Req(setchargingprofile.ReqInput{
		ConnectorID: -1,
		CsChargingProfiles: types.ChargingProfileInput{
			ChargingProfileID:      exampleProfileId,
			TransactionID:          nil,
			StackLevel:             exampleStackLevel,
			ChargingProfilePurpose: "TxDefaultProfile",
			ChargingProfileKind:    "Absolute",
			RecurrencyKind:         nil,
			ValidFrom:              nil,
			ValidTo:                nil,
			ChargingSchedule: types.ChargingScheduleInput{
				Duration:         nil,
				ChargingRateUnit: "W",
				ChargingSchedulePeriod: []types.ChargingSchedulePeriodInput{
					{
						StartPeriod:  exampleStartPeriod,
						Limit:        exampleLimit,
						NumberPhases: nil,
					},
				},
				MinChargingRate: nil,
				StartSchedule:   nil,
			},
		},
	})
	if err != nil {
		fmt.Println("Error: invalid connector ID")
	}
	// Output:
	// Error: invalid connector ID
}
