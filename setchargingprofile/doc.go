// Package setchargingprofile implements the OCPP 1.6 SetChargingProfile
// message pair.
//
// # What It Means
//
// SetChargingProfile installs a power or current schedule on a connector or on
// the Charge Point as a whole. A profile consists of a purpose (TxProfile,
// TxDefaultProfile, ChargePointMaxProfile), a stack level for priority
// ordering,
// and a ChargingSchedule that defines one or more time-slotted periods each
// with a maximum charge rate. The Charge Point merges all active profiles by
// stack level and purpose to compute the effective limit at any instant.
//
// # When It Is Used
//
// The Central System sends SetChargingProfile.req in four situations:
// immediately after StartTransaction to attach a per-transaction profile,
// embedded inside RemoteStartTransaction to pre-attach a profile before the
// session begins, during an active transaction to update the running limit, or
// outside any transaction to set a default profile for a connector or for the
// whole Charge Point. The first period's StartSchedule must always be 0. If a
// profile with the same chargingProfileId or the same stackLevel and purpose
// already exists it is replaced; otherwise it is added to the stack.
//
// # What It Is Not
//
// SetChargingProfile does not start or stop a charging session. A TxProfile
// cannot be set without an active transaction on the target connector. The
// returned schedule from GetCompositeSchedule is indicative only; actual
// delivered energy may differ due to EV behaviour and local constraints.
// SetChargingProfile does not interact with the Authorization Cache or the
// Local Authorization List.
//
// # Adjacent Concepts
//
//   - clearchargingprofile: removes profiles installed by this message.
//   - getcompositeschedule: queries the merged schedule that results from all
//     installed profiles on a connector.
//   - remotestarttransaction: can embed a TxProfile so the limit is active from
//     the first moment of the new session.
//   - starttransaction: the event after which a TxProfile can be applied with
//     a separate SetChargingProfile.req.
//   - types.ChargingSchedule, setchargingprofile/types.ChargingProfile: the
//     shared types that carry the schedule data.
package setchargingprofile
