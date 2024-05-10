// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SpansAggregateBucketType The spans aggregate bucket type.
type SpansAggregateBucketType string

// List of SpansAggregateBucketType.
const (
	SPANSAGGREGATEBUCKETTYPE_BUCKET SpansAggregateBucketType = "bucket"
)

var allowedSpansAggregateBucketTypeEnumValues = []SpansAggregateBucketType{
	SPANSAGGREGATEBUCKETTYPE_BUCKET,
}

// GetAllowedValues reeturns the list of possible values.
func (v *SpansAggregateBucketType) GetAllowedValues() []SpansAggregateBucketType {
	return allowedSpansAggregateBucketTypeEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *SpansAggregateBucketType) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = SpansAggregateBucketType(value)
	return nil
}

// NewSpansAggregateBucketTypeFromValue returns a pointer to a valid SpansAggregateBucketType
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewSpansAggregateBucketTypeFromValue(v string) (*SpansAggregateBucketType, error) {
	ev := SpansAggregateBucketType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for SpansAggregateBucketType: valid values are %v", v, allowedSpansAggregateBucketTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v SpansAggregateBucketType) IsValid() bool {
	for _, existing := range allowedSpansAggregateBucketTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SpansAggregateBucketType value.
func (v SpansAggregateBucketType) Ptr() *SpansAggregateBucketType {
	return &v
}
