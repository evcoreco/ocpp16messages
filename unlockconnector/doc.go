// Package unlockconnector implements the OCPP 1.6 UnlockConnector message pair.
//
// # What It Means
//
// UnlockConnector instructs a Charge Point to release the cable retention lock
// on a specific connector. It is a mechanical action: the connector latch is
// disengaged so the EV driver can remove the cable. The Charge Point replies
// with whether the unlock succeeded, failed, or is not supported on that
// connector.
//
// # When It Is Used
//
// The Central System sends UnlockConnector.req when an EV driver reports that
// the cable is stuck and cannot be removed — typically due to a connector
// malfunction or a locking mechanism fault. If a transaction is in progress on
// the target connector the Charge Point finishes and records the transaction
// first before unlocking, following the same rules as StopTransaction. The
// Charge Point then attempts to disengage the lock and reports the result.
//
// # What It Is Not
//
// UnlockConnector is not a way to stop a running charging session. Use
// RemoteStopTransaction to end a session; the connector will be unlocked as
// part of that process. UnlockConnector only affects the cable retention lock
// on the connector itself — it does not unlock a station enclosure door,
// access panel, or any other physical lock on the station. It is also not a
// safety emergency stop.
//
// # Adjacent Concepts
//
// - remotestoptransaction: the correct path for ending a session remotely;
//   connector unlock happens automatically as part of the transaction stop.
// - stoptransaction: the message the Charge Point sends to record the session
//   end before unlocking, if a transaction was running.
// - statusnotification: the Charge Point may send this after the connector
//   state changes as a result of the unlock.
package unlockconnector
