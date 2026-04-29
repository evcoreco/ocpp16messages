// Package clearcache implements the OCPP 1.6 ClearCache message pair.
//
// # What It Means
//
// ClearCache instructs a Charge Point to erase its Authorization Cache — the
// local store of idTag entries previously returned by the Central System in
// Authorize.conf, StartTransaction.conf, and StopTransaction.conf responses.
// The Charge Point replies with whether the cache was successfully cleared.
//
// # When It Is Used
//
// The Central System sends ClearCache.req when cached authorization data has
// become stale or incorrect: for example, after a batch revocation of idTags,
// a change in the authorization backend, or a security incident. Clearing the
// cache forces the Charge Point to re-validate any future idTag against the
// Central System or the Local Authorization List rather than serving a
// potentially outdated cached result.
//
// # What It Is Not
//
// ClearCache targets only the Authorization Cache, not the Local Authorization
// List. The Local Authorization List is a separate, explicitly managed list
// controlled through SendLocalList and GetLocalListVersion. Clearing the cache
// does not remove charging profiles, connector states, or any other Charge
// Point data.
//
// # Adjacent Concepts
//
// - sendlocallist: manages the Local Authorization List, the sibling
//   authorization mechanism that is not affected by ClearCache.
// - getlocallistversion: checks the version of the Local Authorization List.
// - authorize: the network round-trip that populates the Authorization Cache;
//   after a ClearCache, the next unknown idTag will trigger Authorize.req again.
package clearcache
