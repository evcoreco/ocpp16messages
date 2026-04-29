// Package statusnotification implements the OCPP 1.6 StatusNotification message pair.
//
// # What It Means
//
// StatusNotification is the Charge Point's report of a connector or system
// state change to the Central System. It carries the connector id, the new
// status, an error code, and optional informational text. The Central System
// acknowledges with an empty StatusNotification.conf. Using connectorId 0
// reports the status of the Charge Point as a whole rather than any individual
// connector.
//
// # When It Is Used
//
// The Charge Point sends StatusNotification.req on every meaningful status
// transition: when a connector becomes Available, when an EV plugs in
// (Preparing), during active energy delivery (Charging), when charging is
// paused by the EV or the EVSE (SuspendedEV, SuspendedEVSE), when the session
// is finishing, when a fault is detected, and when the connector is set
// Unavailable via ChangeAvailability. It is also sent after a reservation
// expires to signal the connector returning to Available.
//
// # What It Is Not
//
// StatusNotification is a push notification from the Charge Point; it is not a
// query. The Central System cannot ask "what is your status?" directly through
// StatusNotification — it uses TriggerMessage to request one on demand.
// StatusNotification does not start or stop transactions; it only reports state
// transitions that may be caused by transactions. The "EVSE" terminology in
// this message is used for forward compatibility with OCPP 2.x and does not
// imply a different physical unit.
//
// # Adjacent Concepts
//
// - triggermessage: the Central System sends this to request an immediate
//   StatusNotification for a specific connector or all connectors.
// - changeavailability: a Central System command that causes the Charge Point
//   to send StatusNotification after the availability change takes effect.
// - starttransaction / stoptransaction: transaction lifecycle events that drive
//   status transitions from Available → Preparing → Charging → Finishing →
//   Available.
// - reservenow: a reservation expiry or consumption triggers a StatusNotification
//   to report the resulting connector state.
package statusnotification
