// Package reservenow implements the OCPP 1.6 ReserveNow message pair.
//
// # What It Means
//
// ReserveNow lets the Central System hold a connector for a specific idTag
// until an expiry time. While the reservation is active the Charge Point
// refuses other idTags on that connector. The reserved idTag (and its parent
// idTag, if any) is the only token that can start a session on the reserved
// connector before the expiry. When a matching session starts, the reservation
// is consumed and the reservationId is included in StartTransaction.req.
//
// # When It Is Used
//
// The Central System sends ReserveNow.req when a user books a charging spot in
// advance via a mobile app or operator portal. A connectorId of 0,
// combined with the ReserveConnectorZeroSupported configuration key set
// to true, reserves any one connector at the station rather than a
// specific one. If the given
// reservationId matches an existing reservation that reservation is replaced.
// If the connector is occupied, faulted, unavailable, or the Charge Point is
// configured to reject reservations, the appropriate status is returned.
//
// # What It Is Not
//
// ReserveNow does not start a charging session; it only holds a connector.
// The session still begins through the normal StartTransaction flow when the
// user arrives. A reservation is not a payment guarantee and does not interact
// with billing. ReserveNow is not a way to block a connector indefinitely; the
// expiry date enforces a time limit after which the connector is released
// automatically.
//
// # Adjacent Concepts
//
//   - cancelreservation: releases a reservation before it expires or before the
//     user arrives.
//   - starttransaction: consumes the reservation when the session opens; must
//     include the reservationId.
//   - authorize: the Charge Point may send this to validate the
//     reserved idTag or parent idTag while the reservation is being
//     set up.
//   - statusnotification: sent after an expired reservation to report the
//     connector returning to Available.
package reservenow
