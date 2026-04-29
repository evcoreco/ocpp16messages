// Package getlocallistversion implements the OCPP 1.6 GetLocalListVersion message pair.
//
// # What It Means
//
// GetLocalListVersion asks a Charge Point to report the version number of its
// current Local Authorization List. The Central System uses this version number
// to decide whether a SendLocalList is necessary and whether to send a Full or
// Differential update. Version 0 means the list is empty; version -1 means the
// Charge Point does not support local authorization lists at all.
//
// # When It Is Used
//
// The Central System sends GetLocalListVersion.req before initiating a list
// synchronization: to confirm the Charge Point's list is up to date, to detect
// drift after a network outage, or to bootstrap a freshly connected Charge
// Point. If the reported version matches the expected version no update is
// needed; if it is lower, the Central System decides whether a Differential or
// Full update is more efficient.
//
// # What It Is Not
//
// GetLocalListVersion does not return the list contents — only its version
// number. Fetching or displaying the actual idTag entries is not possible via
// OCPP 1.6; the Central System is the authoritative source of the list. This
// message is also not related to the Authorization Cache, which is managed
// through ClearCache.
//
// # Adjacent Concepts
//
// - sendlocallist: pushes a new or updated Local Authorization List to the
//   Charge Point; GetLocalListVersion determines whether this is needed.
// - clearcache: clears the Authorization Cache, a separate authorization
//   mechanism not affected by the local list version.
// - authorize: the network round-trip that bypasses the local list; a populated
//   and up-to-date local list reduces the need for Authorize.req.
package getlocallistversion
