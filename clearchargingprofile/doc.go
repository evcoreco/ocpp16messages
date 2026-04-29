// Package clearchargingprofile implements the OCPP 1.6 ClearChargingProfile
// message pair.
//
// # What It Means
//
// ClearChargingProfile removes one or more charging profiles from a Charge
// Point. The request may filter by profile id, connector id, purpose, or stack
// level; matching profiles are deleted and the Charge Point re-evaluates its
// active limits. The Charge Point replies Accepted if at least one profile was
// removed, or Unknown if no match was found.
//
// # When It Is Used
//
// The Central System sends ClearChargingProfile.req to withdraw smart-charging
// limits: at the end of a transaction where a TxProfile was applied, when a
// demand-response event ends and a default charging profile should be removed,
// or when reconfiguring the profile stack before installing new profiles via
// SetChargingProfile. Sending the request with all optional fields omitted
// removes all profiles on the Charge Point.
//
// # What It Is Not
//
// ClearChargingProfile does not stop a charging transaction; it only removes
// power or current limits. If no lower-stack-level profile remains after
// removal the Charge Point falls back to unmanaged charging. It is not the
// inverse of GetCompositeSchedule; that message queries the computed schedule
// without modifying it.
//
// # Adjacent Concepts
//
//   - setchargingprofile: installs the profiles that this message removes.
//   - getcompositeschedule: queries what the combined active schedule
//     looks like at any point in time, useful for verifying the effect
//     of a clear.
//   - remotestarttransaction: may embed a TxProfile whose lifecycle is bounded
//     by the transaction; ClearChargingProfile removes it explicitly.
package clearchargingprofile
