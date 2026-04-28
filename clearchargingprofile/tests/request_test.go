package clearchargingprofile_test

import (
	"strings"
	"testing"

	ccp "github.com/evcoreco/ocpp16messages/clearchargingprofile"
	types "github.com/evcoreco/ocpp16types"
)

const (
	errId                     = "id"
	errConnectorID            = "connectorId"
	errChargingProfilePurpose = "chargingProfilePurpose"
	errStackLevel             = "stackLevel"

	purposeNotNil = "ChargingProfilePurpose should not be nil"

	valueZero       = 0
	valueOne        = 1
	valueTwo        = 2
	valueThree      = 3
	valueFive       = 5
	valueId         = 123
	valueNegative   = -1
	valueExceedsMax = 65536
)

func intPtr(v int) *int {
	return &v
}

func strPtr(v string) *string {
	return &v
}

func TestReq_Valid_NoFields(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Id != nil {
		t.Errorf("Id should be nil, got %v", req.Id)
	}

	if req.ConnectorID != nil {
		t.Errorf("ConnectorID should be nil, got %v", req.ConnectorID)
	}

	if req.ChargingProfilePurpose != nil {
		t.Errorf("ChargingProfilePurpose should be nil, got %v",
			req.ChargingProfilePurpose)
	}

	if req.StackLevel != nil {
		t.Errorf("StackLevel should be nil, got %v", req.StackLevel)
	}
}

func TestReq_Valid_WithId(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueId),
		ConnectorID:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Id == nil {
		t.Fatal("Id should not be nil")
	}

	if req.Id.Value() != valueId {
		t.Errorf(types.ErrorMismatchValue, valueId, req.Id.Value())
	}
}

func TestReq_Valid_WithIdZero(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueZero),
		ConnectorID:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Id == nil {
		t.Fatal("Id should not be nil")
	}

	if req.Id.Value() != valueZero {
		t.Errorf(types.ErrorMismatchValue, valueZero, req.Id.Value())
	}
}

func TestReq_Valid_WithConnectorID(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            intPtr(valueOne),
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID == nil {
		t.Fatal("ConnectorID should not be nil")
	}

	if req.ConnectorID.Value() != valueOne {
		t.Errorf(types.ErrorMismatchValue, valueOne, req.ConnectorID.Value())
	}
}

func TestReq_Valid_WithConnectorIDZero(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            intPtr(valueZero),
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID == nil {
		t.Fatal("ConnectorID should not be nil")
	}

	if req.ConnectorID.Value() != valueZero {
		t.Errorf(types.ErrorMismatchValue, valueZero, req.ConnectorID.Value())
	}
}

func TestReq_Valid_WithPurpose_ChargePointMaxProfile(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: strPtr("ChargePointMaxProfile"),
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ChargingProfilePurpose == nil {
		t.Fatal(purposeNotNil)
	}

	expected := types.ChargePointMaxProfile
	if *req.ChargingProfilePurpose != expected {
		t.Errorf(types.ErrorMismatch, expected, *req.ChargingProfilePurpose)
	}
}

func TestReq_Valid_WithPurpose_TxDefaultProfile(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: strPtr("TxDefaultProfile"),
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ChargingProfilePurpose == nil {
		t.Fatal(purposeNotNil)
	}

	expected := types.TxDefaultProfile
	if *req.ChargingProfilePurpose != expected {
		t.Errorf(types.ErrorMismatch, expected, *req.ChargingProfilePurpose)
	}
}

func TestReq_Valid_WithPurpose_TxProfile(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: strPtr("TxProfile"),
		StackLevel:             nil,
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ChargingProfilePurpose == nil {
		t.Fatal(purposeNotNil)
	}

	expected := types.TxProfile
	if *req.ChargingProfilePurpose != expected {
		t.Errorf(types.ErrorMismatch, expected, *req.ChargingProfilePurpose)
	}
}

func TestReq_Valid_WithStackLevel(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             intPtr(valueFive),
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.StackLevel == nil {
		t.Fatal("StackLevel should not be nil")
	}

	if req.StackLevel.Value() != valueFive {
		t.Errorf(types.ErrorMismatchValue, valueFive, req.StackLevel.Value())
	}
}

func TestReq_Valid_AllFields_Id(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueOne),
		ConnectorID:            intPtr(valueTwo),
		ChargingProfilePurpose: strPtr("TxProfile"),
		StackLevel:             intPtr(valueThree),
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.Id == nil || req.Id.Value() != valueOne {
		t.Error("Id mismatch")
	}
}

func TestReq_Valid_AllFields_ConnectorID(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueOne),
		ConnectorID:            intPtr(valueTwo),
		ChargingProfilePurpose: strPtr("TxProfile"),
		StackLevel:             intPtr(valueThree),
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ConnectorID == nil || req.ConnectorID.Value() != valueTwo {
		t.Error("ConnectorID mismatch")
	}
}

func TestReq_Valid_AllFields_Purpose(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueOne),
		ConnectorID:            intPtr(valueTwo),
		ChargingProfilePurpose: strPtr("TxProfile"),
		StackLevel:             intPtr(valueThree),
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ChargingProfilePurpose == nil {
		t.Fatal(purposeNotNil)
	}

	if *req.ChargingProfilePurpose != types.TxProfile {
		t.Error("ChargingProfilePurpose mismatch")
	}
}

func TestReq_Valid_AllFields_StackLevel(t *testing.T) {
	t.Parallel()

	req, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueOne),
		ConnectorID:            intPtr(valueTwo),
		ChargingProfilePurpose: strPtr("TxProfile"),
		StackLevel:             intPtr(valueThree),
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.StackLevel == nil || req.StackLevel.Value() != valueThree {
		t.Error("StackLevel mismatch")
	}
}

func TestReq_Invalid_NegativeId(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueNegative),
		ConnectorID:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative Id")
	}

	if !strings.Contains(err.Error(), errId) {
		t.Errorf(types.ErrorWantContains, err, errId)
	}
}

func TestReq_Invalid_IdExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueExceedsMax),
		ConnectorID:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "Id exceeds max")
	}

	if !strings.Contains(err.Error(), errId) {
		t.Errorf(types.ErrorWantContains, err, errId)
	}
}

func TestReq_Invalid_NegativeConnectorID(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            intPtr(valueNegative),
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative ConnectorID")
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}

func TestReq_Invalid_ConnectorIDExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            intPtr(valueExceedsMax),
		ChargingProfilePurpose: nil,
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "ConnectorID exceeds max")
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}
}

func TestReq_Invalid_Purpose(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: strPtr("Invalid"),
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid ChargingProfilePurpose")
	}

	if !strings.Contains(err.Error(), errChargingProfilePurpose) {
		t.Errorf(types.ErrorWantContains, err, errChargingProfilePurpose)
	}
}

func TestReq_Invalid_PurposeEmpty(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: strPtr(""),
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty ChargingProfilePurpose")
	}

	if !strings.Contains(err.Error(), errChargingProfilePurpose) {
		t.Errorf(types.ErrorWantContains, err, errChargingProfilePurpose)
	}
}

func TestReq_Invalid_PurposeLowercase(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: strPtr("txprofile"),
		StackLevel:             nil,
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase ChargingProfilePurpose")
	}

	if !strings.Contains(err.Error(), errChargingProfilePurpose) {
		t.Errorf(types.ErrorWantContains, err, errChargingProfilePurpose)
	}
}

func TestReq_Invalid_NegativeStackLevel(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             intPtr(valueNegative),
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative StackLevel")
	}

	if !strings.Contains(err.Error(), errStackLevel) {
		t.Errorf(types.ErrorWantContains, err, errStackLevel)
	}
}

func TestReq_Invalid_StackLevelExceedsMax(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     nil,
		ConnectorID:            nil,
		ChargingProfilePurpose: nil,
		StackLevel:             intPtr(valueExceedsMax),
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "StackLevel exceeds max")
	}

	if !strings.Contains(err.Error(), errStackLevel) {
		t.Errorf(types.ErrorWantContains, err, errStackLevel)
	}
}

func TestReq_Invalid_MultipleErrors(t *testing.T) {
	t.Parallel()

	_, err := ccp.Req(ccp.ReqInput{
		Id:                     intPtr(valueNegative),
		ConnectorID:            intPtr(valueNegative),
		ChargingProfilePurpose: strPtr("Invalid"),
		StackLevel:             intPtr(valueNegative),
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "multiple invalid fields")
	}

	if !strings.Contains(err.Error(), errId) {
		t.Errorf(types.ErrorWantContains, err, errId)
	}

	if !strings.Contains(err.Error(), errConnectorID) {
		t.Errorf(types.ErrorWantContains, err, errConnectorID)
	}

	if !strings.Contains(err.Error(), errChargingProfilePurpose) {
		t.Errorf(types.ErrorWantContains, err, errChargingProfilePurpose)
	}

	if !strings.Contains(err.Error(), errStackLevel) {
		t.Errorf(types.ErrorWantContains, err, errStackLevel)
	}
}
