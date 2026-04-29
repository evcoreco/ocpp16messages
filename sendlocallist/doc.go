// Package sendlocallist implements the OCPP 1.6 SendLocalList message pair.
//
// # What It Means
//
// SendLocalList pushes an authorization list to a Charge Point so it can
// approve or reject idTags without contacting the Central System. A Full update
// replaces the entire stored list. A Differential update adds, modifies, or
// removes individual entries from the existing list. Both update types carry a
// version number the Charge Point stores alongside the list.
//
// # When It Is Used
//
// The Central System sends SendLocalList.req when connectivity to the Central
// System cannot be guaranteed, when reducing authorization latency for frequent
// users, or during initial deployment of a Charge Point. After a
// VersionMismatch
// or Failed response to a Differential update, the Central System should retry
// with a Full update to restore a known-good state.
//
// # What It Is Not
//
// SendLocalList manages the Local Authorization List, not the Authorization
// Cache. The cache is populated automatically from Authorize, StartTransaction,
// and StopTransaction responses and is cleared with ClearCache. The local list
// is an explicit, versioned dataset. SendLocalList does not fetch the list back
// from the Charge Point; the Central System is the authoritative source and
// OCPP 1.6 has no mechanism to read the list contents out of a Charge Point.
//
// # Adjacent Concepts
//
//   - getlocallistversion: reads the Charge Point's current list version before
//     deciding whether a SendLocalList is needed and which update type to use.
//   - clearcache: clears the Authorization Cache, which is separate from the
//     local list managed by SendLocalList.
//   - authorize: the network round-trip that becomes unnecessary for idTags
//     covered by an up-to-date local list.
package sendlocallist
