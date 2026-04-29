// Package remotestoptransaction implements the OCPP 1.6
// RemoteStopTransaction message pair.
//
// # What It Means
//
// RemoteStopTransaction lets the Central System ask a Charge Point to end an
// ongoing charging session identified by its transactionId. The Charge Point
// treats the request as equivalent to a local stop action, sends
// StopTransaction.req to record the session end, and unlocks the connector if
// applicable.
//
// # When It Is Used
//
// The Central System sends RemoteStopTransaction.req when a user ends a session
// through a mobile app or web portal, when a billing system determines the
// session budget is exhausted, or when an operator needs to free a connector
// remotely. The Charge Point replies Accepted if the transactionId is active on
// one of its connectors, or Rejected if it is not.
//
// # What It Is Not
//
// RemoteStopTransaction is not the same as UnlockConnector.
// RemoteStopTransaction
// ends the charging session and then unlocks the cable as a side effect.
// UnlockConnector only releases the cable retention mechanism without stopping
// the session, and is intended for stuck-cable recovery. RemoteStopTransaction
// is also not a forced disconnect at the hardware level; if the Charge Point
// cannot stop gracefully it will still send StopTransaction.req.
//
// # Adjacent Concepts
//
//   - stoptransaction: the Charge Point-initiated message that formally records
//     the session end; always follows a successful RemoteStopTransaction.
//   - unlockconnector: releases the cable retention lock without stopping the
//     session — for mechanical failures, not remote session management.
//   - remotestarttransaction: the counterpart for starting a session remotely.
//   - starttransaction: paired with stoptransaction to bracket the session the
//     remote stop ends.
package remotestoptransaction
