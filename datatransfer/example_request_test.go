package datatransfer_test

import (
	"fmt"

	"github.com/evcoreco/ocpp16messages/datatransfer"
)

const labelVendorID = "VendorID:"

// ExampleReq demonstrates creating a valid DataTransfer.req message
// with only the required vendorId field.
func ExampleReq() {
	req, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  "com.example.vendor",
		MessageID: nil,
		Data:      nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelVendorID, req.VendorID.Value())
	// Output:
	// VendorID: com.example.vendor
}

// ExampleReq_withMessageID demonstrates creating a DataTransfer.req message
// with an optional messageId.
func ExampleReq_withMessageID() {
	messageId := "CustomMessage"

	req, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  "com.example.vendor",
		MessageID: &messageId,
		Data:      nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelVendorID, req.VendorID.Value())
	fmt.Println("MessageID:", req.MessageID.Value())
	// Output:
	// VendorID: com.example.vendor
	// MessageID: CustomMessage
}

// ExampleReq_withData demonstrates creating a DataTransfer.req message
// with an optional data payload.
func ExampleReq_withData() {
	data := `{"temperature": 25.5, "humidity": 60}`

	req, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  "com.example.vendor",
		MessageID: nil,
		Data:      &data,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelVendorID, req.VendorID.Value())
	fmt.Println("HasData:", req.Data != nil)
	// Output:
	// VendorID: com.example.vendor
	// HasData: true
}

// ExampleReq_complete demonstrates creating a complete DataTransfer.req
// message with all fields populated.
func ExampleReq_complete() {
	messageId := "GetTemperature"
	data := `{"sensorId": "temp-001"}`

	req, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  "com.example.vendor",
		MessageID: &messageId,
		Data:      &data,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelVendorID, req.VendorID.Value())
	fmt.Println("MessageID:", req.MessageID.Value())
	fmt.Println("Data:", *req.Data)
	// Output:
	// VendorID: com.example.vendor
	// MessageID: GetTemperature
	// Data: {"sensorId": "temp-001"}
}

// ExampleReq_emptyVendorID demonstrates the error returned when
// an empty vendorId is provided.
func ExampleReq_emptyVendorID() {
	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  "",
		MessageID: nil,
		Data:      nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// vendorId: value cannot be empty
}

// ExampleReq_invalidVendorIDChars demonstrates the error returned when
// the vendorId contains non-printable ASCII characters.
func ExampleReq_invalidVendorIDChars() {
	_, err := datatransfer.Req(datatransfer.ReqInput{
		VendorID:  "vendor\x00id",
		MessageID: nil,
		Data:      nil,
	})
	if err != nil {
		fmt.Println("vendorId has non-printable ASCII characters")
	}
	// Output:
	// vendorId has non-printable ASCII characters
}
