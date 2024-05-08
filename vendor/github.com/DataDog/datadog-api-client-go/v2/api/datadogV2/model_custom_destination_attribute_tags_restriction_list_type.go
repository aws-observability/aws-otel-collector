// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationAttributeTagsRestrictionListType How `forward_tags_restriction_list` parameter should be interpreted.
// If `ALLOW_LIST`, then only tags whose keys on the forwarded logs match the ones on the restriction list
// are forwarded.
//
// `BLOCK_LIST` works the opposite way. It does not forward the tags matching the ones on the list.
type CustomDestinationAttributeTagsRestrictionListType string

// List of CustomDestinationAttributeTagsRestrictionListType.
const (
	CUSTOMDESTINATIONATTRIBUTETAGSRESTRICTIONLISTTYPE_ALLOW_LIST CustomDestinationAttributeTagsRestrictionListType = "ALLOW_LIST"
	CUSTOMDESTINATIONATTRIBUTETAGSRESTRICTIONLISTTYPE_BLOCK_LIST CustomDestinationAttributeTagsRestrictionListType = "BLOCK_LIST"
)

var allowedCustomDestinationAttributeTagsRestrictionListTypeEnumValues = []CustomDestinationAttributeTagsRestrictionListType{
	CUSTOMDESTINATIONATTRIBUTETAGSRESTRICTIONLISTTYPE_ALLOW_LIST,
	CUSTOMDESTINATIONATTRIBUTETAGSRESTRICTIONLISTTYPE_BLOCK_LIST,
}

// GetAllowedValues reeturns the list of possible values.
func (v *CustomDestinationAttributeTagsRestrictionListType) GetAllowedValues() []CustomDestinationAttributeTagsRestrictionListType {
	return allowedCustomDestinationAttributeTagsRestrictionListTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *CustomDestinationAttributeTagsRestrictionListType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = CustomDestinationAttributeTagsRestrictionListType(value)
	return nil
}

// NewCustomDestinationAttributeTagsRestrictionListTypeFromValue returns a pointer to a valid CustomDestinationAttributeTagsRestrictionListType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewCustomDestinationAttributeTagsRestrictionListTypeFromValue(v string) (*CustomDestinationAttributeTagsRestrictionListType, error) {
	ev := CustomDestinationAttributeTagsRestrictionListType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for CustomDestinationAttributeTagsRestrictionListType: valid values are %v", v, allowedCustomDestinationAttributeTagsRestrictionListTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v CustomDestinationAttributeTagsRestrictionListType) IsValid() bool {
	for _, existing := range allowedCustomDestinationAttributeTagsRestrictionListTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomDestinationAttributeTagsRestrictionListType value.
func (v CustomDestinationAttributeTagsRestrictionListType) Ptr() *CustomDestinationAttributeTagsRestrictionListType {
	return &v
}
