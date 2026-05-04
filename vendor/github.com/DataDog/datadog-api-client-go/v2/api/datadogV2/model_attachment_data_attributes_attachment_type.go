// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AttachmentDataAttributesAttachmentType The type of the attachment.
type AttachmentDataAttributesAttachmentType string

// List of AttachmentDataAttributesAttachmentType.
const (
	ATTACHMENTDATAATTRIBUTESATTACHMENTTYPE_POSTMORTEM AttachmentDataAttributesAttachmentType = "postmortem"
	ATTACHMENTDATAATTRIBUTESATTACHMENTTYPE_LINK       AttachmentDataAttributesAttachmentType = "link"
)

var allowedAttachmentDataAttributesAttachmentTypeEnumValues = []AttachmentDataAttributesAttachmentType{
	ATTACHMENTDATAATTRIBUTESATTACHMENTTYPE_POSTMORTEM,
	ATTACHMENTDATAATTRIBUTESATTACHMENTTYPE_LINK,
}

// GetAllowedValues reeturns the list of possible values.
func (v *AttachmentDataAttributesAttachmentType) GetAllowedValues() []AttachmentDataAttributesAttachmentType {
	return allowedAttachmentDataAttributesAttachmentTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *AttachmentDataAttributesAttachmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = AttachmentDataAttributesAttachmentType(value)
	return nil
}

// NewAttachmentDataAttributesAttachmentTypeFromValue returns a pointer to a valid AttachmentDataAttributesAttachmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewAttachmentDataAttributesAttachmentTypeFromValue(v string) (*AttachmentDataAttributesAttachmentType, error) {
	ev := AttachmentDataAttributesAttachmentType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for AttachmentDataAttributesAttachmentType: valid values are %v", v, allowedAttachmentDataAttributesAttachmentTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v AttachmentDataAttributesAttachmentType) IsValid() bool {
	for _, existing := range allowedAttachmentDataAttributesAttachmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AttachmentDataAttributesAttachmentType value.
func (v AttachmentDataAttributesAttachmentType) Ptr() *AttachmentDataAttributesAttachmentType {
	return &v
}
