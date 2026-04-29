// Package firmwarestatusnotification implements the OCPP 1.6
// FirmwareStatusNotification message pair.
//
// # What It Means
//
// FirmwareStatusNotification is the progress report a Charge Point sends to
// the Central System during a firmware update initiated by UpdateFirmware.req.
// Each notification carries a status value — Downloaded, DownloadFailed,
// Downloading, Idle, InstallationFailed, Installing, Installed — allowing the
// Central System to monitor the update lifecycle.
//
// # When It Is Used
//
// The Charge Point sends FirmwareStatusNotification.req after receiving
// UpdateFirmware.req and at each meaningful transition: when the download
// starts, completes, fails, when installation begins, and when it finishes. If
// the Central System triggers a FirmwareStatusNotification via TriggerMessage
// and no update is in progress, the Charge Point reports status Idle. The
// Central System must acknowledge each notification with
// FirmwareStatusNotification.conf.
//
// # What It Is Not
//
// FirmwareStatusNotification is not the command that starts a firmware update;
// that is UpdateFirmware.req. It does not transfer the firmware binary — only
// status. The actual firmware image is downloaded by the Charge Point directly
// from the URL provided in UpdateFirmware.req.
//
// # Adjacent Concepts
//
// - updatefirmware: the request that instructs the Charge Point to download and
//   install firmware; FirmwareStatusNotification reports its progress.
// - diagnosticsstatusnotification: the analogous progress-report message for
//   diagnostics upload operations.
// - triggermessage: can request an on-demand FirmwareStatusNotification to
//   query the current update state.
// - reset: commonly follows a successful firmware installation to bring up the
//   new image.
package firmwarestatusnotification
