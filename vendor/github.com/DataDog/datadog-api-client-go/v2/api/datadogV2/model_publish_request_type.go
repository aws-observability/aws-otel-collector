// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PublishRequestType The publish-request resource type.
type PublishRequestType string

// List of PublishRequestType.
const (
	PUBLISHREQUESTTYPE_PUBLISHREQUEST PublishRequestType = "publishRequest"
)

var allowedPublishRequestTypeEnumValues = []PublishRequestType{
	PUBLISHREQUESTTYPE_PUBLISHREQUEST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *PublishRequestType) GetAllowedValues() []PublishRequestType {
	return allowedPublishRequestTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *PublishRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = PublishRequestType(value)
	return nil
}

// NewPublishRequestTypeFromValue returns a pointer to a valid PublishRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewPublishRequestTypeFromValue(v string) (*PublishRequestType, error) {
	ev := PublishRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for PublishRequestType: valid values are %v", v, allowedPublishRequestTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v PublishRequestType) IsValid() bool {
	for _, existing := range allowedPublishRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PublishRequestType value.
func (v PublishRequestType) Ptr() *PublishRequestType {
	return &v
}
