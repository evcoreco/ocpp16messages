// Package ocpp16messages provides typed constructors and validation for all
// 28 OCPP 1.6 request and response messages used in EV charging networks.
//
// # What It Means
//
// This library implements the full OCPP 1.6 message surface as Go types.
// Every OCPP operation has its own sub-package (e.g., authorize, heartbeat,
// metervalues) exposing a Req() constructor, a Conf() constructor, and the
// types specific to that operation. Shared types — CiString variants,
// DateTime, Integer, IdTagInfo, ChargingSchedule, MeterValue — live in the
// types sub-package and are reused across messages.
//
// All constructors validate their input at call time and return errors rather
// than panicking. Types are immutable and safe for concurrent reads once built.
//
// # When It Is Used
//
// Use this library when building a Charge Point (EVSE) or a Central System
// (CSMS/backend) in Go that must produce or consume OCPP 1.6 JSON messages.
// It covers message construction and field validation; it is not responsible
// for transport, session management, or WebSocket framing.
//
// # What It Is Not
//
// This is not a full OCPP stack. It does not handle WebSocket connections,
// OCPP call/result/error framing, routing between multiple Charge Points, or
// any OCPP 2.x messages. It also does not implement OCPP 1.6 SOAP; only the
// JSON profile is in scope.
//
// # Adjacent Concepts
//
// Shared validated types are in the types sub-package:
//
//	import "github.com/evcoreco/ocpp16messages/types"
//
// Each OCPP operation lives in its own sub-package. For example:
//
//	import "github.com/evcoreco/ocpp16messages/authorize"
//	import "github.com/evcoreco/ocpp16messages/bootnotification"
//	import "github.com/evcoreco/ocpp16messages/metervalues"
//
// See also: github.com/evcoreco/ocpp16types for the standalone type library.
//
// Example:
//
//	req, err := authorize.Req(authorize.ReqInput{
//		IDTag: "RFID-ABC123",
//	})
//	if err != nil {
//		// handle validation error
//	}
//	fmt.Println(req.IDTag.String())
package ocpp16messages
