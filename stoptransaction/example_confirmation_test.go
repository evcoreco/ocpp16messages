package stoptransaction_test

import (
	"fmt"

	"github.com/evcoreco/ocpp16messages/stoptransaction"
)

const (
	exampleStatusAccepted = "Accepted"
	exampleStatusLabel    = "Status:"
	exampleErrorLabel     = "Error:"
)

// ExampleConf demonstrates creating a StopTransaction.conf with no IDTagInfo.
func ExampleConf() {
	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      nil,
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		fmt.Println(exampleErrorLabel, err)

		return
	}

	if conf.IDTagInfo == nil {
		fmt.Println("IDTagInfo: nil (no authorization info)")
	}
	// Output:
	// IDTagInfo: nil (no authorization info)
}

// ExampleConf_accepted demonstrates creating an accepted response.
func ExampleConf_accepted() {
	status := exampleStatusAccepted

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		fmt.Println(exampleErrorLabel, err)

		return
	}

	fmt.Println(exampleStatusLabel, conf.IDTagInfo.Status().String())
	// Output:
	// Status: Accepted
}

// ExampleConf_withExpiryDate demonstrates including an expiry date.
func ExampleConf_withExpiryDate() {
	status := exampleStatusAccepted
	expiryDate := "2025-12-31T23:59:59Z"

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  &expiryDate,
		ParentIDTag: nil,
	})
	if err != nil {
		fmt.Println(exampleErrorLabel, err)

		return
	}

	fmt.Println(exampleStatusLabel, conf.IDTagInfo.Status().String())
	fmt.Println("Has expiry date:", conf.IDTagInfo.ExpiryDate() != nil)
	// Output:
	// Status: Accepted
	// Has expiry date: true
}

// ExampleConf_withParentIDTag demonstrates including a parent ID tag.
func ExampleConf_withParentIDTag() {
	status := exampleStatusAccepted
	parentIDTag := "PARENT-123"

	conf, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIDTag: &parentIDTag,
	})
	if err != nil {
		fmt.Println(exampleErrorLabel, err)

		return
	}

	fmt.Println(exampleStatusLabel, conf.IDTagInfo.Status().String())
	fmt.Println("ParentIDTag:", conf.IDTagInfo.ParentIDTag().String())
	// Output:
	// Status: Accepted
	// ParentIDTag: PARENT-123
}

// ExampleConf_invalidStatus demonstrates validation error for invalid status.
func ExampleConf_invalidStatus() {
	status := "InvalidStatus"

	_, err := stoptransaction.Conf(stoptransaction.ConfInput{
		Status:      &status,
		ExpiryDate:  nil,
		ParentIDTag: nil,
	})
	if err != nil {
		fmt.Println("Validation failed as expected")
	}
	// Output:
	// Validation failed as expected
}
