// Package starttransaction implements the OCPP 1.6 StartTransaction
// message pair.
//
// # What It Means
//
// StartTransaction is the Charge Point's formal notification that a charging
// session has opened. It carries the connector id, the idTag that triggered the
// session, the opening meter reading, and the timestamp. The Central System
// assigns a transactionId and returns the idTag's authorization status. If the
// session ends a reservation the StartTransaction.req must include the
// reservationId.
//
// # When It Is Used
//
// The Charge Point sends StartTransaction.req as soon as a session is
// established — whether triggered by a physical token presentation, a
// RemoteStartTransaction, or a reservation being consumed. The Central System
// must always respond; withholding the response causes the Charge Point to
// retry according to its transaction-related error handling rules. The Central
// System should re-verify the idTag in StartTransaction.req because the Charge
// Point may have authorized it locally with stale cached data.
//
// # What It Is Not
//
// StartTransaction is not the authorization step; that is Authorize.req, which
// may have occurred before this message or may be implied by the
// RemoteStartTransaction flow. It does not carry periodic meter readings; those
// are reported via MeterValues.req. StartTransaction is not a command from the
// Central System; it is always initiated by the Charge Point.
//
// # Adjacent Concepts
//
//   - stoptransaction: the counterpart that closes the session opened here.
//   - authorize: may precede StartTransaction to validate the idTag,
//     though it is not always required.
//   - metervalues: carries the meter samples between the opening reading in
//     StartTransaction and the closing reading in StopTransaction.
//   - remotestarttransaction: a Central System-initiated trigger that
//     results in the Charge Point sending StartTransaction.req.
//   - reservenow: a reservation that is consumed when this transaction
//     opens must include the reservationId in StartTransaction.req.
package starttransaction
