// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionPointsContentEncoding HTTP header used to compress the media-type.
type DistributionPointsContentEncoding string

// List of DistributionPointsContentEncoding.
const (
	DISTRIBUTIONPOINTSCONTENTENCODING_DEFLATE DistributionPointsContentEncoding = "deflate"
)

var allowedDistributionPointsContentEncodingEnumValues = []DistributionPointsContentEncoding{
	DISTRIBUTIONPOINTSCONTENTENCODING_DEFLATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *DistributionPointsContentEncoding) GetAllowedValues() []DistributionPointsContentEncoding {
	return allowedDistributionPointsContentEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *DistributionPointsContentEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = DistributionPointsContentEncoding(value)
	return nil
}

// NewDistributionPointsContentEncodingFromValue returns a pointer to a valid DistributionPointsContentEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewDistributionPointsContentEncodingFromValue(v string) (*DistributionPointsContentEncoding, error) {
	ev := DistributionPointsContentEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DistributionPointsContentEncoding: valid values are %v", v, allowedDistributionPointsContentEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v DistributionPointsContentEncoding) IsValid() bool {
	for _, existing := range allowedDistributionPointsContentEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistributionPointsContentEncoding value.
func (v DistributionPointsContentEncoding) Ptr() *DistributionPointsContentEncoding {
	return &v
}
