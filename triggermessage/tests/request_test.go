package triggermessage_test

import (
	"strings"
	"testing"

	"github.com/evcoreco/ocpp16messages/triggermessage"
	types "github.com/evcoreco/ocpp16types"
)

const (
	errRequestedMessage = "requestedMessage"
	errConnectorID      = "connectorId"
	fieldConnectorID    = "ConnectorID"
	connectorIdZero     = 0
	connectorIdOne      = 1
	connectorIdNegative = -1
	connectorIdMax      = 65535
	connectorIdOverflow = 65536
)

func TestReq_Valid_BootNotification(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "BootNotification",
		ConnectorID:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.RequestedMessage != types.MessageTriggerBootNotification {
		t.Errorf(
			types.ErrorMismatch,
			types.MessageTriggerBootNotification,
			req.RequestedMessage,
		)
	}
}

func TestReq_Valid_DiagnosticsStatusNotification(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "DiagnosticsStatusNotification",
		ConnectorID:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.RequestedMessage !=
		types.MessageTriggerDiagnosticsStatusNotification {
		t.Errorf(
			types.ErrorMismatch,
			types.MessageTriggerDiagnosticsStatusNotification,
			req.RequestedMessage,
		)
	}
}

func TestReq_Valid_FirmwareStatusNotification(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "FirmwareStatusNotification",
		ConnectorID:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.RequestedMessage != types.MessageTriggerFirmwareStatusNotification {
		t.Errorf(
			types.ErrorMismatch,
			types.MessageTriggerFirmwareStatusNotification,
			req.RequestedMessage,
		)
	}
}

func TestReq_Valid_Heartbeat(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorID:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.RequestedMessage != types.MessageTriggerHeartbeat {
		t.Errorf(
			types.ErrorMismatch,
			types.MessageTriggerHeartbeat,
			req.RequestedMessage,
		)
	}
}

func TestReq_Valid_MeterValues(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "MeterValues",
		ConnectorID:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.RequestedMessage != types.MessageTriggerMeterValues {
		t.Errorf(
			types.ErrorMismatch,
			types.MessageTriggerMeterValues,
			req.RequestedMessage,
		)
	}
}

func TestReq_Valid_StatusNotification(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "StatusNotification",
		ConnectorID:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.RequestedMessage != types.MessageTriggerStatusNotification {
		t.Errorf(
			types.ErrorMismatch,
			types.MessageTriggerStatusNotification,
			req.RequestedMessage,
		)
	}
}

func TestReq_Valid_WithConnectorIDZero(t *testing.T) {
	t.Parallel()

	connectorId := connectorIdZero

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "StatusNotification",
		ConnectorID:      &connectorId,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID == nil {
		t.Errorf(types.ErrorWantNonNil, fieldConnectorID)
	}

	if req.ConnectorID.Value() != uint16(connectorIdZero) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(connectorIdZero),
			req.ConnectorID.Value(),
		)
	}
}

func TestReq_Valid_WithConnectorIDOne(t *testing.T) {
	t.Parallel()

	connectorId := connectorIdOne

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "MeterValues",
		ConnectorID:      &connectorId,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID == nil {
		t.Errorf(types.ErrorWantNonNil, fieldConnectorID)
	}

	if req.ConnectorID.Value() != uint16(connectorIdOne) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(connectorIdOne),
			req.ConnectorID.Value(),
		)
	}
}

func TestReq_Valid_WithConnectorIDMax(t *testing.T) {
	t.Parallel()

	connectorId := connectorIdMax

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorID:      &connectorId,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID == nil {
		t.Errorf(types.ErrorWantNonNil, fieldConnectorID)
	}

	if req.ConnectorID.Value() != uint16(connectorIdMax) {
		t.Errorf(
			types.ErrorMismatchValue,
			uint16(connectorIdMax),
			req.ConnectorID.Value(),
		)
	}
}

func TestReq_Valid_WithoutConnectorID(t *testing.T) {
	t.Parallel()

	req, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorID:      nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID != nil {
		t.Errorf("ConnectorID = %v, want nil", req.ConnectorID)
	}
}

func TestReq_EmptyRequestedMessage(t *testing.T) {
	t.Parallel()

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "",
		ConnectorID:      nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty requestedMessage")
	}

	if !strings.Contains(err.Error(), errRequestedMessage) {
		t.Errorf(types.ErrorWantContains, err, errRequestedMessage)
	}
}

func TestReq_InvalidRequestedMessage_Unknown(t *testing.T) {
	t.Parallel()

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Unknown",
		ConnectorID:      nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown requestedMessage")
	}

	if !strings.Contains(err.Error(), errRequestedMessage) {
		t.Errorf(types.ErrorWantContains, err, errRequestedMessage)
	}
}

func TestReq_InvalidRequestedMessage_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "heartbeat",
		ConnectorID:      nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase requestedMessage")
	}

	if !strings.Contains(err.Error(), errRequestedMessage) {
		t.Errorf(types.ErrorWantContains, err, errRequestedMessage)
	}
}

func TestReq_InvalidRequestedMessage_StartTransaction(t *testing.T) {
	t.Parallel()

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "StartTransaction",
		ConnectorID:      nil,
	})
	if err == nil {
		t.Errorf(
			types.ErrorWantNil,
			"StartTransaction (not valid per OCPP 1.6)",
		)
	}

	if !strings.Contains(err.Error(), errRequestedMessage) {
		t.Errorf(types.ErrorWantContains, err, errRequestedMessage)
	}
}

func TestReq_InvalidConnectorID_Negative(t *testing.T) {
	t.Parallel()

	connectorId := connectorIdNegative

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorID:      &connectorId,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative connectorId")
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}

func TestReq_InvalidConnectorID_Overflow(t *testing.T) {
	t.Parallel()

	connectorId := connectorIdOverflow

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Heartbeat",
		ConnectorID:      &connectorId,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "connectorId overflow")
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}

func TestReq_MultipleErrors_InvalidMessageAndConnectorID(t *testing.T) {
	t.Parallel()

	connectorId := connectorIdNegative

	_, err := triggermessage.Req(triggermessage.ReqInput{
		RequestedMessage: "Unknown",
		ConnectorID:      &connectorId,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid message and connectorId")
	}

	if !strings.Contains(err.Error(), errRequestedMessage) {
		t.Errorf(types.ErrorWantContains, err, errRequestedMessage)
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}
