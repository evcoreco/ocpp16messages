// Package cancelreservation implements the OCPP 1.6 CancelReservation
// message pair.
//
// # What It Means
//
// CancelReservation lets the Central System release a connector reservation
// before it expires or before the reserved idTag arrives. The Charge Point
// looks up the given reservationId; if it exists the reservation is removed
// and the connector returns to its normal available state.
//
// # When It Is Used
//
// The Central System sends CancelReservation.req when a user cancels a booking
// through the operator portal or mobile app, or when a back-end process
// determines the reservation is no longer needed. The Charge Point replies
// Accepted if a matching reservation was found and removed, or Rejected if no
// reservation with that id exists.
//
// # What It Is Not
//
// CancelReservation is not a way to stop an ongoing charging transaction. Once
// a transaction has started the reservation has already been consumed; use
// RemoteStopTransaction to end the session. It is also not related to connector
// availability in general; use ChangeAvailability for that.
//
// # Adjacent Concepts
//
//   - reservenow: the counterpart that creates the reservation being cancelled.
//   - remotestoptransaction: stops a charging transaction that started
//     after the reservation was consumed.
//   - statusnotification: the Charge Point sends this after the connector
//     transitions back to Available following a cancellation.
package cancelreservation
