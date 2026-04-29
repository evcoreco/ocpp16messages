// Package getdiagnostics implements the OCPP 1.6 GetDiagnostics message pair.
//
// # What It Means
//
// GetDiagnostics instructs a Charge Point to collect diagnostic information
// and upload it as a single file to a URL provided by the Central System. The
// request may include optional start and end timestamps to scope the data. The
// Charge Point replies with the name of the file it will upload, or with no
// filename if no diagnostic data is available.
//
// # When It Is Used
//
// The Central System sends GetDiagnostics.req when troubleshooting a Charge
// Point: investigating transaction failures, connectivity issues, hardware
// faults, or unusual behavior. The format of the uploaded file is not
// prescribed by OCPP; it is defined by the Charge Point vendor. Once the
// upload begins the Charge Point keeps the Central System informed via
// DiagnosticsStatusNotification.req messages.
//
// # What It Is Not
//
// GetDiagnostics is not a real-time data stream. It produces a one-shot file
// upload and is not intended for continuous monitoring. It is not the same as
// GetConfiguration, which reads named runtime parameters rather than a raw
// diagnostics artifact. The diagnostics file content is opaque to OCPP.
//
// # Adjacent Concepts
//
//   - diagnosticsstatusnotification: the progress reports the Charge
//     Point sends during the upload triggered by GetDiagnostics.req.
//   - getconfiguration: reads structured, named configuration values in
//     real time without requiring a file upload.
//   - updatefirmware: the parallel mechanism for pushing a file to a
//     Charge Point rather than pulling one from it.
package getdiagnostics
