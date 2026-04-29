// Package stoptransaction implements the OCPP 1.6 StopTransaction message pair.
//
// # What It Means
//
// StopTransaction is the Charge Point's formal notification that a charging
// session has ended. It carries the transactionId, the closing meter reading,
// the stop timestamp, and an optional stop reason. It may include an optional
// TransactionData array of MeterValues sampled during the session. The Central
// System may return idTag authorization information in the response, which the
// Charge Point should use to update its Authorization Cache.
//
// # When It Is Used
//
// The Charge Point sends StopTransaction.req whenever a session ends: when the
// EV driver presents a token to stop, when a RemoteStopTransaction is processed,
// after a Reset, or when configured stop conditions occur such as
// StopTransactionOnEVSideDisconnect. The idTag may be omitted when the Charge
// Point itself stops the transaction (for example, during a reset). If the stop
// reason is a normal local action the reason field may be omitted and is assumed
// to be Local; for all other reasons it should be set explicitly. The Central
// System cannot prevent a transaction from stopping; it must always respond.
//
// # What It Is Not
//
// StopTransaction is not a command; it is always initiated by the Charge Point.
// The Central System cannot withhold StopTransaction.conf to keep a session
// open. StopTransaction does not unlock the connector — that is a side effect
// of the stop process governed by the UnlockConnectorOnEVSideDisconnect
// configuration key. It is not the same as RemoteStopTransaction, which is the
// Central System command that triggers a StopTransaction.
//
// # Adjacent Concepts
//
// - starttransaction: the counterpart that opened the session this message
//   closes, carrying the transactionId used here.
// - remotestoptransaction: the Central System command that causes the Charge
//   Point to send StopTransaction.req.
// - metervalues: the periodic readings during the session; TransactionData in
//   StopTransaction reuses the same MeterValue structure.
// - unlockconnector: releases the cable retention lock; related but separate
//   from the session stop.
// - types.MeterValue: the shared type used in TransactionData.
package stoptransaction
