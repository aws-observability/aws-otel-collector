// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContentEncoding HTTP header used to compress the media-type.
type ContentEncoding string

// List of ContentEncoding.
const (
	CONTENTENCODING_GZIP    ContentEncoding = "gzip"
	CONTENTENCODING_DEFLATE ContentEncoding = "deflate"
)

var allowedContentEncodingEnumValues = []ContentEncoding{
	CONTENTENCODING_GZIP,
	CONTENTENCODING_DEFLATE,
}

// GetAllowedValues reeturns the list of possible values.
func (v *ContentEncoding) GetAllowedValues() []ContentEncoding {
	return allowedContentEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *ContentEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = ContentEncoding(value)
	return nil
}

// NewContentEncodingFromValue returns a pointer to a valid ContentEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewContentEncodingFromValue(v string) (*ContentEncoding, error) {
	ev := ContentEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ContentEncoding: valid values are %v", v, allowedContentEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v ContentEncoding) IsValid() bool {
	for _, existing := range allowedContentEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContentEncoding value.
func (v ContentEncoding) Ptr() *ContentEncoding {
	return &v
}
