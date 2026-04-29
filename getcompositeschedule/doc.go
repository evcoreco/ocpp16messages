// Package getcompositeschedule implements the OCPP 1.6 GetCompositeSchedule message pair.
//
// # What It Means
//
// GetCompositeSchedule asks a Charge Point to compute and return the combined
// charging schedule that results from merging all active charging profiles for
// a given connector over a requested duration. The Charge Point merges profiles
// across stack levels and purposes and responds with a single flattened schedule
// showing the power or current limit at every point in the requested window.
//
// # When It Is Used
//
// The Central System sends GetCompositeSchedule.req to preview what limits a
// Charge Point will enforce: before setting a new profile, to verify that
// profiles were applied correctly, or during demand-response monitoring. The
// schedule covers the interval from the moment the request is received up to
// that moment plus the requested duration. A connectorId of 0 returns the
// aggregate expected consumption of the entire Charge Point rather than a
// single connector.
//
// # What It Is Not
//
// GetCompositeSchedule is a read-only query; it does not modify any profile.
// The returned schedule is a snapshot computed at request time and may change
// as new profiles are installed or removed. It is not a transaction command and
// does not start or stop charging.
//
// # Adjacent Concepts
//
// - setchargingprofile: installs the profiles whose combined effect
//   GetCompositeSchedule computes and returns.
// - clearchargingprofile: removes profiles from the stack, altering the
//   composite result.
// - metervalues: the real-time energy measurements that reflect how the
//   composite schedule is being enforced in practice.
package getcompositeschedule
