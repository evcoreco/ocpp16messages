// Package bootnotification implements the OCPP 1.6 BootNotification
// message pair.
//
// # What It Means
//
// BootNotification is the registration handshake a Charge Point performs with
// the Central System immediately after establishing a connection. The Charge
// Point reports its hardware identity (vendor, model, serial numbers, firmware
// version, etc.) and the Central System replies with an acceptance status,
// the current time, and the heartbeat interval the Charge Point should use.
//
// # When It Is Used
//
// A Charge Point sends BootNotification.req on every boot and every reconnect.
// It must not send any other OCPP message until the Central System has
// responded. If the response status is Accepted the Charge Point may proceed
// normally, synchronize its clock to the provided timestamp, and adopt the
// returned heartbeat interval. If the status is Pending the communication
// channel stays open; the Central System may send configuration requests, but
// the Charge Point must not initiate any requests other than those triggered
// via TriggerMessage. If the status is Rejected the Charge Point must wait for
// the returned interval before retrying, and must not respond to any Central
// System-initiated messages in the interim.
//
// # What It Is Not
//
// BootNotification is not a heartbeat. It does not prove the Charge Point is
// still connected during a session; that is the role of Heartbeat.req.
// BootNotification is also not an authorization step: it identifies the
// hardware but does not authorize any idTag for charging.
//
// # Adjacent Concepts
//
//   - heartbeat: the periodic keep-alive sent once the Charge Point is
//     accepted.
//   - triggermessage: used by the Central System to request a new
//     BootNotification while the Charge Point is in Pending state.
//   - reset: a soft or hard reboot that causes the Charge Point to send a new
//     BootNotification after coming back online.
//   - changeconfiguration / getconfiguration: configuration exchanges
//     the Central System may initiate while in Pending state.
package bootnotification
