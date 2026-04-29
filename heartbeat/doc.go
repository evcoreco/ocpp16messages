// Package heartbeat implements the OCPP 1.6 Heartbeat message pair.
//
// # What It Means
//
// Heartbeat is the periodic signal a Charge Point sends to tell the Central
// System it is still online. The Central System replies with its current UTC
// timestamp, which the Charge Point should use to keep its internal clock
// synchronized.
//
// # When It Is Used
//
// A Charge Point sends Heartbeat.req at the interval specified by the
// HeartbeatInterval configuration key, set by the Central System in
// BootNotification.conf or via ChangeConfiguration. Any other OCPP message sent
// within the heartbeat window counts as proof of connectivity, so a Heartbeat
// is skipped if another PDU was already sent. When using JSON over WebSocket,
// heartbeats are not mandatory for connection keep-alive, but at least one per
// 24 hours is recommended to maintain clock synchronization.
//
// # What It Is Not
//
// Heartbeat is not the initial registration message; that is BootNotification.
// It does not carry any transaction or authorization information. It is not a
// replacement for WebSocket ping/pong frames, which operate at the transport
// layer independently of OCPP.
//
// # Adjacent Concepts
//
//   - bootnotification: establishes the connection and provides the initial
//     heartbeat interval; Heartbeat runs on that cadence afterwards.
//   - changeconfiguration: the mechanism for adjusting the HeartbeatInterval
//     configuration key.
//   - triggermessage: the Central System may use this to request an immediate
//     Heartbeat outside the normal interval.
package heartbeat
