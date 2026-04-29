// Package datatransfer implements the OCPP 1.6 DataTransfer message pair.
//
// # What It Means
//
// DataTransfer is the OCPP 1.6 extension point for vendor-specific messages
// that do not fit any of the 27 other defined operations. Either side — Charge
// Point or Central System — may initiate a DataTransfer.req. The message
// carries a vendorId that identifies the vendor-specific implementation, an
// optional messageId for sub-command dispatch, and an optional opaque data
// payload whose format is agreed between sender and receiver.
//
// # When It Is Used
//
// DataTransfer is used when a Charge Point and a Central System need to
// exchange proprietary information: custom display text, vendor diagnostics,
// tariff data, or any feature outside the OCPP 1.6 specification. The vendorId
// should be a reversed DNS name matching the vendor's primary domain to avoid
// collisions. If the recipient does not recognize the vendorId it replies with
// status UnknownVendor; if it recognizes the vendor but not the messageId it
// replies UnknownMessageId.
//
// # What It Is Not
//
// DataTransfer is not a general-purpose RPC layer for re-implementing standard
// OCPP operations. It is not a substitute for ChangeConfiguration or
// GetConfiguration; those cover all standard configuration exchanges. It is
// also not a diagnostic upload mechanism; use GetDiagnostics for file-based
// diagnostics retrieval.
//
// # Adjacent Concepts
//
//   - changeconfiguration / getconfiguration: the standard path for reading and
//     writing named configuration keys, preferred over DataTransfer when the
//     data fits a simple key-value model.
//   - getdiagnostics / diagnosticsstatusnotification: structured diagnostics
//     upload for known OCPP diagnostic use cases.
package datatransfer
