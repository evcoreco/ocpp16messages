//go:build fuzz

package fuzz

import (
	"errors"
	"testing"

	dt "github.com/evcoreco/ocpp16messages/datatransfer"
	types "github.com/evcoreco/ocpp16types"
)

func FuzzDataTransferReq(f *testing.F) {
	f.Add("VendorID", false, "", false, "")
	f.Add("", false, "", false, "")
	f.Add("VendorID", true, "MessageID", true, "payload")
	f.Add("VendorID", true, "", false, "")

	f.Fuzz(func(
		t *testing.T,
		vendorId string,
		hasMessageID bool,
		messageId string,
		hasData bool,
		data string,
	) {
		if len(vendorId) > maxFuzzStringLen ||
			len(messageId) > maxFuzzStringLen ||
			len(data) > maxFuzzStringLen {
			t.Skip()
		}

		var messageIdPtr *string
		if hasMessageID {
			messageIdPtr = &messageId
		}

		var dataPtr *string
		if hasData {
			dataPtr = &data
		}

		req, err := dt.Req(dt.ReqInput{
			VendorID:  vendorId,
			MessageID: messageIdPtr,
			Data:      dataPtr,
		})
		if err != nil {
			if !errors.Is(err, types.ErrInvalidValue) && !errors.Is(err, types.ErrEmptyValue) {
				t.Fatalf(
					"error = %v, want wrapping ErrEmptyValue or ErrInvalidValue",
					err,
				)
			}

			return
		}

		if vendorId == "" {
			t.Fatal("Req succeeded with empty VendorID")
		}

		if req.VendorID.String() != vendorId {
			t.Fatalf("VendorID = %q, want %q", req.VendorID.String(), vendorId)
		}

		if hasMessageID {
			if req.MessageID == nil {
				t.Fatal("MessageID = nil, want non-nil")
			}
			if messageId == "" {
				t.Fatal("Req succeeded with empty MessageID")
			}
			if req.MessageID.String() != messageId {
				t.Fatalf(
					"MessageID = %q, want %q",
					req.MessageID.String(),
					messageId,
				)
			}
		} else if req.MessageID != nil {
			t.Fatal("MessageID != nil, want nil")
		}

		if hasData {
			if req.Data == nil {
				t.Fatal("Data = nil, want non-nil")
			}
			if *req.Data != data {
				t.Fatalf("Data = %q, want %q", *req.Data, data)
			}
		} else if req.Data != nil {
			t.Fatal("Data != nil, want nil")
		}
	})
}
