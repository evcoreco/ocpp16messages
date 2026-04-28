package authorize_test

import (
	"fmt"

	"github.com/evcoreco/ocpp16messages/authorize"
)

const labelStatus = "Status:"

// ExampleConf demonstrates creating a valid Authorize.conf message
// with an Accepted status.
func ExampleConf() {
	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.IDTagInfo.Status().String())
	// Output:
	// Status: Accepted
}

// ExampleConf_blocked demonstrates creating an Authorize.conf message
// with a Blocked status.
func ExampleConf_blocked() {
	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Blocked",
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.IDTagInfo.Status().String())
	// Output:
	// Status: Blocked
}

// ExampleConf_withExpiryDate demonstrates creating an Authorize.conf message
// with an expiry date.
func ExampleConf_withExpiryDate() {
	expiryDate := "2025-12-31T23:59:59Z"

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  &expiryDate,
		ParentIDTag: nil,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.IDTagInfo.Status().String())
	fmt.Println("HasExpiryDate:", conf.IDTagInfo.ExpiryDate() != nil)
	// Output:
	// Status: Accepted
	// HasExpiryDate: true
}

// ExampleConf_withParentIDTag demonstrates creating an Authorize.conf message
// with a parent ID tag.
func ExampleConf_withParentIDTag() {
	parentTag := "PARENT-123"

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  nil,
		ParentIDTag: &parentTag,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.IDTagInfo.Status().String())
	fmt.Println("ParentIDTag:", conf.IDTagInfo.ParentIDTag().String())
	// Output:
	// Status: Accepted
	// ParentIDTag: PARENT-123
}

// ExampleConf_complete demonstrates creating a complete Authorize.conf message
// with all optional fields populated.
func ExampleConf_complete() {
	expiryDate := "2025-12-31T23:59:59Z"
	parentTag := "PARENT-123"

	conf, err := authorize.Conf(authorize.ConfInput{
		Status:      "Accepted",
		ExpiryDate:  &expiryDate,
		ParentIDTag: &parentTag,
	})
	if err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(labelStatus, conf.IDTagInfo.Status().String())
	fmt.Println("HasExpiryDate:", conf.IDTagInfo.ExpiryDate() != nil)
	fmt.Println("ParentIDTag:", conf.IDTagInfo.ParentIDTag().String())
	// Output:
	// Status: Accepted
	// HasExpiryDate: true
	// ParentIDTag: PARENT-123
}

// ExampleConf_invalidStatus demonstrates the error returned when
// an invalid status is provided.
func ExampleConf_invalidStatus() {
	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Unknown",
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// status: NewIDTagInfo: AuthorizationStatus: invalid value
}

// ExampleConf_multipleErrors demonstrates that all validation errors
// are returned at once, not just the first one encountered.
func ExampleConf_multipleErrors() {
	invalidDate := "not-a-date"
	longTag := "THIS-TAG-IS-WAY-TOO-LONG-FOR-OCPP"

	_, err := authorize.Conf(authorize.ConfInput{
		Status:      "Invalid-Status",
		ExpiryDate:  &invalidDate,
		ParentIDTag: &longTag,
	})
	if err != nil {
		fmt.Println("Multiple errors")
	}
	// Output:
	// Multiple errors
}
