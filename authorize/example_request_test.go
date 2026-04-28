package authorize_test

import (
	"fmt"

	"github.com/evcoreco/ocpp16messages/authorize"
)

// ExampleReq demonstrates creating a valid Authorize.req message
// with a properly formatted ID tag using the ReqInput struct.
func ExampleReq() {
	req, err := authorize.Req(authorize.ReqInput{IDTag: "RFID-TAG-12345"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IDTag:", req.IDTag.String())
	// Output:
	// IDTag: RFID-TAG-12345
}

// ExampleReq_emptyIDTag demonstrates the error returned when
// an empty ID tag is provided.
func ExampleReq_emptyIDTag() {
	_, err := authorize.Req(authorize.ReqInput{IDTag: ""})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// idTag: value cannot be empty
}

// ExampleReq_idTagTooLong demonstrates the error returned when
// the ID tag exceeds the maximum length of 20 characters.
func ExampleReq_idTagTooLong() {
	_, err := authorize.Req(authorize.ReqInput{
		IDTag: "RFID-ABC123456789012345",
	}) // 23 chars
	if err != nil {
		fmt.Println("idTag: Exceeds maximum length")
	}
	// Output:
	// idTag: Exceeds maximum length
}

// ExampleReq_invalidCharacters demonstrates the error returned when
// the ID tag contains non-printable ASCII characters.
func ExampleReq_invalidCharacters() {
	// Contains null byte
	_, err := authorize.Req(authorize.ReqInput{IDTag: "RFID\x00TAG"})
	if err != nil {
		fmt.Println("idTag has non-printable ASCII characters")
	}
	// Output:
	// idTag has non-printable ASCII characters
}

// ExampleReq_shortIDTag demonstrates that short ID tags
// (within the 20 character limit) are valid.
func ExampleReq_shortIDTag() {
	req, err := authorize.Req(authorize.ReqInput{IDTag: "TAG1"})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println("IDTag:", req.IDTag.String())
	// Output:
	// IDTag: TAG1
}
