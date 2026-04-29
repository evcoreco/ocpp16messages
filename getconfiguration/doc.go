// Package getconfiguration implements the OCPP 1.6 GetConfiguration message pair.
//
// # What It Means
//
// GetConfiguration lets the Central System read one or more configuration
// parameters from a Charge Point. The request carries an optional list of key
// names; the Charge Point returns the current value and read-only status for
// each recognized key, and lists any unrecognized keys separately. Sending the
// request with no keys causes the Charge Point to return all of its
// configuration settings.
//
// # When It Is Used
//
// The Central System sends GetConfiguration.req to audit Charge Point settings
// after deployment, to verify that a previous ChangeConfiguration.req was
// applied correctly, or to discover what configuration keys a particular
// hardware model supports. The maximum number of keys that can be requested in
// a single call may be capped by the Charge Point's GetConfigurationMaxKeys
// setting.
//
// # What It Is Not
//
// GetConfiguration is a read-only operation; it does not write any values.
// Writing values is done via ChangeConfiguration. It does not return
// diagnostics data or meter readings; use GetDiagnostics and MeterValues for
// those respectively.
//
// # Adjacent Concepts
//
// - changeconfiguration: the write counterpart that sets a configuration key
//   to a new value.
// - getconfiguration/types.KeyValue: the type that carries each key name, its
//   current value, and its read-only flag in the response.
// - datatransfer: an alternative path for vendor-specific data exchange when
//   the information does not map to a named key-value pair.
package getconfiguration
