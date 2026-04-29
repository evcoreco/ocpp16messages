// Package reset implements the OCPP 1.6 Reset message pair.
//
// # What It Means
//
// Reset instructs a Charge Point to reboot. A Soft reset gracefully stops any
// ongoing transactions, sends StopTransaction.req for each, and then restarts
// the application software. A Hard reset restarts all hardware immediately;
// graceful transaction termination is not required, but the Charge Point should
// send StopTransaction.req for previously active transactions after the reboot
// completes and BootNotification.conf is accepted. Certain states — such as a
// connector set to Unavailable — persist across both types of reset.
//
// # When It Is Used
//
// The Central System sends Reset.req for planned maintenance, after a
// ChangeConfiguration.req that returned RebootRequired, or as a recovery action
// when a Charge Point is unresponsive. The Charge Point replies with whether it
// will attempt the reset before carrying it out. Hard reset is a last resort;
// queued messages may be lost.
//
// # What It Is Not
//
// Reset is not a way to stop a single transaction; use
// RemoteStopTransaction for that. It is not a configuration change; it
// does not alter any setting on its own. A Soft reset is not guaranteed
// to complete instantly; it waits for active transactions to close,
// which may take time.
//
// # Adjacent Concepts
//
//   - bootnotification: sent by the Charge Point after it comes back online
//     following a reset.
//   - stoptransaction: the Charge Point sends this for each active transaction
//     before (Soft) or after (Hard) the reboot.
//   - changeconfiguration: the operation that most commonly triggers a
//     Reset when the response status is RebootRequired.
//   - updatefirmware: another operation that may require a reboot as
//     part of the firmware installation process.
package reset
