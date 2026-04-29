// Package authorize implements the OCPP 1.6 Authorize message pair.
//
// # What It Means
//
// Authorize is the network round-trip a Charge Point performs to ask the
// Central System whether a given idTag is allowed to charge. The Central
// System replies with an authorization status and, optionally, a parent idTag
// and expiry timestamp bundled in an IdTagInfo.
//
// # When It Is Used
//
// A Charge Point sends Authorize.req when it cannot confirm authorization
// locally — that is, when the idTag is absent from the Local Authorization
// List and not found in the Authorization Cache with a valid, non-expired
// entry. It is also sent when stopping a transaction if the idTag presented
// to stop the session differs from the one that started it. If the idTag is
// resolved locally, sending Authorize.req is optional.
//
// # What It Is Not
//
// Authorize is not the only path to authorization. A Charge Point that has a
// Local Authorization List or a populated Authorization Cache may approve or
// reject an idTag without contacting the Central System at all. Authorize is
// also not a transaction control message: it does not start or stop a session;
// that is the role of StartTransaction and StopTransaction.
//
// # Adjacent Concepts
//
// - starttransaction: sent after a successful authorization to open a session.
// - stoptransaction: sent when a session ends; may trigger an Authorize if
//   the stopping idTag differs from the starting one.
// - sendlocallist / getlocallistversion: manage the Local Authorization List
//   that can make Authorize.req unnecessary.
// - types.IdTagInfo, types.IdToken: the shared types carried in responses.
package authorize
