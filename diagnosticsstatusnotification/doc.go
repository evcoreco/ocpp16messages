// Package diagnosticsstatusnotification implements the OCPP 1.6
// DiagnosticsStatusNotification message pair.
//
// # What It Means
//
// DiagnosticsStatusNotification is the progress report a Charge Point sends
// to the Central System during a diagnostics file upload initiated by
// GetDiagnostics.req. Each notification carries a status value — Idle,
// Uploading, Uploaded, UploadFailed — so the Central System can track whether
// the file transfer completed successfully or needs to be retried.
//
// # When It Is Used
//
// The Charge Point sends DiagnosticsStatusNotification.req after receiving
// GetDiagnostics.req and at each meaningful state transition during the upload.
// If the Central System triggers a DiagnosticsStatusNotification via
// TriggerMessage and the Charge Point is not actively uploading, it sends
// status Idle. The Central System must acknowledge each notification with
// DiagnosticsStatusNotification.conf.
//
// # What It Is Not
//
// DiagnosticsStatusNotification is not the request that starts the diagnostics
// process; that is GetDiagnostics.req. It does not carry the diagnostics data
// itself — only a status. The actual file is uploaded out-of-band to the URL
// supplied in GetDiagnostics.req.
//
// # Adjacent Concepts
//
//   - getdiagnostics: the request that initiates the diagnostics upload and
//     provides the upload URL; DiagnosticsStatusNotification reports
//     its progress.
//   - firmwarestatusnotification: the analogous progress-report message for
//     firmware update operations.
//   - triggermessage: can request an on-demand DiagnosticsStatusNotification to
//     query the current upload state.
package diagnosticsstatusnotification
