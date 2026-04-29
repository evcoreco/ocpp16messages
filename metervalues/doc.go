// Package metervalues implements the OCPP 1.6 MeterValues message pair.
//
// # What It Means
//
// MeterValues carries electrical measurements — energy, power, current,
// voltage, temperature, and more — from a Charge Point connector to the
// Central System. Each request contains one or more MeterValue elements, each
// timestamped and containing one or more SampledValue readings that describe
// the measurement type, phase, unit, location, and optional context and format.
//
// # When It Is Used
//
// A Charge Point sends MeterValues.req on a configurable periodic interval
// (controlled by MeterValueSampleInterval and MeterValuesSampledData), at the
// start and end of a transaction (Clock-Aligned samples), and when the Central
// System requests it via TriggerMessage. connectorId 0 refers to the Charge
// Point's main energy meter rather than any individual connector. transactionId
// is optional and may be omitted when no transaction is in progress.
//
// The default interpretation of a SampledValue with no optional fields is a
// register reading of active import energy in Wh. Two special measurands —
// Current.Offered and Power.Offered — report the maximum current or power the
// Charge Point is currently offering to the EV for smart charging purposes.
//
// # What It Is Not
//
// MeterValues is not a transaction control message; it does not start, stop, or
// authorize a session. It is not the only source of energy data: StartTransaction
// and StopTransaction also carry meter readings at the transaction boundaries.
// The format field ("SignedData") is marked experimental in OCPP 1.6 and may be
// deprecated in future versions.
//
// # Adjacent Concepts
//
// - starttransaction / stoptransaction: carry meter readings at session
//   boundaries; MeterValues fills in the periodic samples between them.
// - setchargingprofile: sets the power or current limits that Current.Offered
//   and Power.Offered measurands reflect.
// - changeconfiguration: adjusts MeterValueSampleInterval and
//   MeterValuesSampledData to control when and what is reported.
// - triggermessage: can request an immediate MeterValues report with the most
//   recent sample set.
// - types.MeterValue, types.SampledValue: the shared types that this package
//   uses for the measurement payload.
package metervalues
