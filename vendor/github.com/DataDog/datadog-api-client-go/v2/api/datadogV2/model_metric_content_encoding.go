// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricContentEncoding HTTP header used to compress the media-type.
type MetricContentEncoding string

// List of MetricContentEncoding.
const (
	METRICCONTENTENCODING_DEFLATE MetricContentEncoding = "deflate"
	METRICCONTENTENCODING_ZSTD1   MetricContentEncoding = "zstd1"
	METRICCONTENTENCODING_GZIP    MetricContentEncoding = "gzip"
)

var allowedMetricContentEncodingEnumValues = []MetricContentEncoding{
	METRICCONTENTENCODING_DEFLATE,
	METRICCONTENTENCODING_ZSTD1,
	METRICCONTENTENCODING_GZIP,
}

// GetAllowedValues reeturns the list of possible values.
func (v *MetricContentEncoding) GetAllowedValues() []MetricContentEncoding {
	return allowedMetricContentEncodingEnumValues
}

// UnmarshalJSON deserializes the given payload.
func (v *MetricContentEncoding) UnmarshalJSON(src []byte) error {
	var value string
	err := datadog.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	*v = MetricContentEncoding(value)
	return nil
}

// NewMetricContentEncodingFromValue returns a pointer to a valid MetricContentEncoding
// for the value passed as argument, or an error if the value passed is not allowed by the enum.
func NewMetricContentEncodingFromValue(v string) (*MetricContentEncoding, error) {
	ev := MetricContentEncoding(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for MetricContentEncoding: valid values are %v", v, allowedMetricContentEncodingEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise.
func (v MetricContentEncoding) IsValid() bool {
	for _, existing := range allowedMetricContentEncodingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricContentEncoding value.
func (v MetricContentEncoding) Ptr() *MetricContentEncoding {
	return &v
}
