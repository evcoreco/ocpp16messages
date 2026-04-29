// Package remotestarttransaction implements the OCPP 1.6 RemoteStartTransaction message pair.
//
// # What It Means
//
// RemoteStartTransaction lets the Central System ask a Charge Point to begin a
// charging session on behalf of a user who is not physically present at the
// station. The Central System provides the idTag the session should be
// associated with, optionally a specific connectorId, and optionally a
// TxProfile charging profile to apply from the moment the session opens.
//
// # When It Is Used
//
// The Central System sends RemoteStartTransaction.req when a user initiates
// charging through a mobile app, a web portal, or an SMS gateway. The Charge
// Point responds immediately to say whether it will attempt the start (Accepted
// or Rejected), then proceeds independently: if AuthorizeRemoteTxRequests is
// true it validates the idTag through the normal authorization path before
// opening the session; if false it starts the session and lets StartTransaction
// confirm authorization with the Central System. The Charge Point then sends
// StartTransaction.req once the session actually begins.
//
// # What It Is Not
//
// An Accepted response does not guarantee the transaction will start; it means
// the Charge Point will try. A connector fault, a failed authorization, or a
// missing EV may still prevent the session from opening. RemoteStartTransaction
// is not a substitute for the full StartTransaction flow; StartTransaction.req
// is still sent and must be responded to. It is also not the only way to start
// a session: a user presenting a physical token at the Charge Point follows the
// same StartTransaction path without a RemoteStartTransaction.
//
// # Adjacent Concepts
//
// - starttransaction: the Charge Point-initiated message that formally opens
//   the session; always follows a successful RemoteStartTransaction attempt.
// - remotestoptransaction: the counterpart for ending a session remotely.
// - setchargingprofile: the standalone way to attach a TxProfile to an already
//   running transaction; RemoteStartTransaction can embed this in one step.
// - authorize: the authorization round-trip the Charge Point performs when
//   AuthorizeRemoteTxRequests is true.
package remotestarttransaction
