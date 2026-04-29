// Package changeavailability implements the OCPP 1.6 ChangeAvailability
// message pair.
//
// # What It Means
//
// ChangeAvailability lets the Central System set a connector — or the entire
// Charge Point — to either Operative (available for charging) or Inoperative
// (not available for charging). The Charge Point replies with whether the
// change was applied immediately, scheduled for after the current transaction,
// or rejected.
//
// # When It Is Used
//
// The Central System sends ChangeAvailability.req for planned maintenance
// windows, load-management policies, or operator-driven out-of-service
// actions. If a transaction is running on the target connector the Charge Point
// responds Scheduled, meaning it will apply the change once the transaction
// finishes. The Charge Point notifies the Central System of the resulting
// status transition via StatusNotification.req after the change takes effect.
//
// # What It Is Not
//
// ChangeAvailability is not a transaction control command. It does not stop a
// running session; use RemoteStopTransaction for that. Setting a connector to
// Inoperative also does not clear its charging profiles or its Authorization
// Cache; use ClearChargingProfile and ClearCache for those.
//
// # Adjacent Concepts
//
//   - statusnotification: the Charge Point sends this after the availability
//     change takes effect to report the new connector state.
//   - remotestoptransaction: stops a transaction before a planned availability
//     change when immediate deactivation is needed.
//   - reset: another path to taking a Charge Point offline that survives across
//     reboots for persistent Unavailable states.
package changeavailability
