// Package updatefirmware implements the OCPP 1.6 UpdateFirmware message pair.
//
// # What It Means
//
// UpdateFirmware instructs a Charge Point to download a firmware image from a
// given URL and install it. The request specifies the earliest time the Charge
// Point may begin the download and optional retry parameters. The Charge Point
// replies immediately with an empty UpdateFirmware.conf (no status), then
// carries out the download and installation independently, reporting progress
// via FirmwareStatusNotification.req messages.
//
// # When It Is Used
//
// The Central System sends UpdateFirmware.req to deploy a security patch, a
// feature release, or a compliance fix to one or more Charge Points. The
// retrieve date lets operators schedule updates during off-peak hours. During
// installation the Charge Point should set idle connectors to Unavailable to
// prevent new sessions from starting while the update is in progress. A reboot
// after installation is recommended to verify the new image before reporting
// the Installed status.
//
// # What It Is Not
//
// UpdateFirmware is not a configuration change; it replaces the software image,
// not runtime parameters. It is not a synchronous operation; the Central System
// does not receive a completion status in the response — it learns the outcome
// through FirmwareStatusNotification messages. The firmware file format and
// integrity verification method are vendor-defined and outside the OCPP
// specification.
//
// # Adjacent Concepts
//
// - firmwarestatusnotification: the progress reports the Charge Point sends
//   throughout the download and installation triggered by this message.
// - reset: commonly sent after a firmware installation when the Charge Point
//   needs to reboot to activate the new image.
// - changeconfiguration: the lighter-weight path for changing runtime behaviour
//   without replacing the firmware.
// - getdiagnostics: the parallel operation for pulling data out of a Charge
//   Point rather than pushing software into it.
package updatefirmware
