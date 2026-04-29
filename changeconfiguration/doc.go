// Package changeconfiguration implements the OCPP 1.6 ChangeConfiguration
// message pair.
//
// # What It Means
//
// ChangeConfiguration lets the Central System write a single key-value
// configuration parameter on the Charge Point. The Charge Point replies with
// whether the change was accepted immediately, requires a reboot to
// take effect,
// is not supported by that hardware, or was rejected for another reason such as
// an out-of-range value.
//
// # When It Is Used
//
// The Central System sends ChangeConfiguration.req to tune Charge Point
// behaviour: adjusting heartbeat intervals, enabling or disabling local
// authorization, setting the connector phase rotation, or updating any other
// OCPP-defined or vendor-specific configuration key. Because only one key-value
// pair is carried per message, multiple calls are needed to change several
// settings at once.
//
// # What It Is Not
//
// ChangeConfiguration is not a firmware update; it changes runtime parameters,
// not the software image. It is also not GetConfiguration: it writes a value
// rather than reading one. The content and format of keys and values beyond the
// OCPP-defined standard keys are vendor-defined and not validated by this
// library beyond the CiString length and character constraints.
//
// # Adjacent Concepts
//
//   - getconfiguration: reads one or more configuration keys from the Charge
//     Point, complementing this write operation.
//   - updatefirmware: the path to replacing the Charge Point software image
//     rather than adjusting its runtime configuration.
//   - reset: some configuration changes require a reboot; the RebootRequired
//     response status signals this and Reset.req is the follow-up action.
package changeconfiguration
