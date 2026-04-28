package sendlocallist_test

import (
	"strings"
	"testing"

	"github.com/evcoreco/ocpp16messages/sendlocallist"
	types "github.com/evcoreco/ocpp16types"
)

const (
	errListVersion            = "ListVersion"
	errUpdateType             = "UpdateType"
	errLocalAuthorizationList = "localAuthorizationList"
	validIDTag                = "RFID12345"
	validStatus               = "Accepted"
	listVersionZero           = 0
	listVersionOne            = 1
	listVersionNegative       = -1
	expectedLenZero           = 0
	expectedLenOne            = 1
	expectedLenThree          = 3
)

func TestReq_Valid_Full_EmptyList(t *testing.T) {
	t.Parallel()

	req, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionOne,
		LocalAuthorizationList: nil,
		UpdateType:             "Full",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.UpdateType != types.UpdateTypeFull {
		t.Errorf(
			types.ErrorMismatch,
			types.UpdateTypeFull,
			req.UpdateType,
		)
	}
}

func TestReq_Valid_Differential_EmptyList(t *testing.T) {
	t.Parallel()

	req, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionOne,
		LocalAuthorizationList: nil,
		UpdateType:             "Differential",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.UpdateType != types.UpdateTypeDifferential {
		t.Errorf(
			types.ErrorMismatch,
			types.UpdateTypeDifferential,
			req.UpdateType,
		)
	}
}

func TestReq_Valid_ListVersionZero(t *testing.T) {
	t.Parallel()

	req, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionZero,
		LocalAuthorizationList: nil,
		UpdateType:             "Full",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.ListVersion.Value() != listVersionZero {
		t.Errorf(
			types.ErrorMismatchValue,
			listVersionZero,
			req.ListVersion.Value(),
		)
	}
}

func TestReq_Valid_WithAuthorizationList(t *testing.T) {
	t.Parallel()

	req, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion: listVersionOne,
		LocalAuthorizationList: []types.AuthorizationDataInput{
			{
				IDTag: validIDTag,
				IDTagInfo: &types.IDTagInfoInput{
					Status:      validStatus,
					ExpiryDate:  nil,
					ParentIDTag: nil,
				},
			},
		},
		UpdateType: "Full",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if len(req.LocalAuthorizationList) != expectedLenOne {
		t.Errorf(
			types.ErrorMismatchValue,
			expectedLenOne,
			len(req.LocalAuthorizationList),
		)
	}
}

func TestReq_Valid_WithMultipleAuthorizationEntries(t *testing.T) {
	t.Parallel()

	req, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion: listVersionOne,
		LocalAuthorizationList: []types.AuthorizationDataInput{
			{
				IDTag: "RFID001",
				IDTagInfo: &types.IDTagInfoInput{
					Status:      "Accepted",
					ExpiryDate:  nil,
					ParentIDTag: nil,
				},
			},
			{
				IDTag: "RFID002",
				IDTagInfo: &types.IDTagInfoInput{
					Status:      "Blocked",
					ExpiryDate:  nil,
					ParentIDTag: nil,
				},
			},
			{
				IDTag:     "RFID003",
				IDTagInfo: nil,
			},
		},
		UpdateType: "Differential",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if len(req.LocalAuthorizationList) != expectedLenThree {
		t.Errorf(
			types.ErrorMismatchValue,
			expectedLenThree,
			len(req.LocalAuthorizationList),
		)
	}
}

func TestReq_Valid_WithEmptyAuthorizationList(t *testing.T) {
	t.Parallel()

	req, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionOne,
		LocalAuthorizationList: []types.AuthorizationDataInput{},
		UpdateType:             "Full",
	})
	if err != nil {
		t.Errorf(types.ErrorUnexpectedError, err)
	}

	if req.LocalAuthorizationList == nil {
		t.Error("LocalAuthorizationList = nil, want empty slice")
	}

	if len(req.LocalAuthorizationList) != expectedLenZero {
		t.Errorf(
			types.ErrorMismatchValue,
			expectedLenZero,
			len(req.LocalAuthorizationList),
		)
	}
}

func TestReq_InvalidListVersion_Negative(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionNegative,
		LocalAuthorizationList: nil,
		UpdateType:             "Full",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "negative list version")
	}

	if !strings.Contains(err.Error(), errListVersion) {
		t.Errorf(types.ErrorWantContains, err, errListVersion)
	}
}

func TestReq_InvalidUpdateType_Empty(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionOne,
		LocalAuthorizationList: nil,
		UpdateType:             "",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty update type")
	}

	if !strings.Contains(err.Error(), errUpdateType) {
		t.Errorf(types.ErrorWantContains, err, errUpdateType)
	}
}

func TestReq_InvalidUpdateType_Unknown(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionOne,
		LocalAuthorizationList: nil,
		UpdateType:             "Unknown",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "unknown update type")
	}

	if !strings.Contains(err.Error(), errUpdateType) {
		t.Errorf(types.ErrorWantContains, err, errUpdateType)
	}
}

func TestReq_InvalidUpdateType_Lowercase(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionOne,
		LocalAuthorizationList: nil,
		UpdateType:             "full",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "lowercase update type")
	}

	if !strings.Contains(err.Error(), errUpdateType) {
		t.Errorf(types.ErrorWantContains, err, errUpdateType)
	}
}

func TestReq_InvalidAuthorizationList_EmptyIDTag(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion: listVersionOne,
		LocalAuthorizationList: []types.AuthorizationDataInput{
			{IDTag: "", IDTagInfo: nil},
		},
		UpdateType: "Full",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "empty idTag in list")
	}

	if !strings.Contains(err.Error(), errLocalAuthorizationList) {
		t.Errorf(types.ErrorWantContains, err, errLocalAuthorizationList)
	}
}

func TestReq_InvalidAuthorizationList_InvalidStatus(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion: listVersionOne,
		LocalAuthorizationList: []types.AuthorizationDataInput{
			{
				IDTag: validIDTag,
				IDTagInfo: &types.IDTagInfoInput{
					Status:      "InvalidStatus",
					ExpiryDate:  nil,
					ParentIDTag: nil,
				},
			},
		},
		UpdateType: "Full",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "invalid status in list")
	}

	if !strings.Contains(err.Error(), errLocalAuthorizationList) {
		t.Errorf(types.ErrorWantContains, err, errLocalAuthorizationList)
	}
}

func TestReq_MultipleErrors_Accumulated(t *testing.T) {
	t.Parallel()

	_, err := sendlocallist.Req(sendlocallist.ReqInput{
		ListVersion:            listVersionNegative,
		LocalAuthorizationList: nil,
		UpdateType:             "Invalid",
	})
	if err == nil {
		t.Errorf(types.ErrorWantNil, "multiple errors")
	}

	if !strings.Contains(err.Error(), errListVersion) {
		t.Errorf(types.ErrorWantContains, err, errListVersion)
	}

	if !strings.Contains(err.Error(), errUpdateType) {
		t.Errorf(types.ErrorWantContains, err, errUpdateType)
	}
}
