// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonS3GenericCompression - Compression algorithm applied to encoded logs.
type ObservabilityPipelineAmazonS3GenericCompression struct {
	ObservabilityPipelineAmazonS3GenericCompressionZstd   *ObservabilityPipelineAmazonS3GenericCompressionZstd
	ObservabilityPipelineAmazonS3GenericCompressionGzip   *ObservabilityPipelineAmazonS3GenericCompressionGzip
	ObservabilityPipelineAmazonS3GenericCompressionSnappy *ObservabilityPipelineAmazonS3GenericCompressionSnappy

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineAmazonS3GenericCompressionZstdAsObservabilityPipelineAmazonS3GenericCompression is a convenience function that returns ObservabilityPipelineAmazonS3GenericCompressionZstd wrapped in ObservabilityPipelineAmazonS3GenericCompression.
func ObservabilityPipelineAmazonS3GenericCompressionZstdAsObservabilityPipelineAmazonS3GenericCompression(v *ObservabilityPipelineAmazonS3GenericCompressionZstd) ObservabilityPipelineAmazonS3GenericCompression {
	return ObservabilityPipelineAmazonS3GenericCompression{ObservabilityPipelineAmazonS3GenericCompressionZstd: v}
}

// ObservabilityPipelineAmazonS3GenericCompressionGzipAsObservabilityPipelineAmazonS3GenericCompression is a convenience function that returns ObservabilityPipelineAmazonS3GenericCompressionGzip wrapped in ObservabilityPipelineAmazonS3GenericCompression.
func ObservabilityPipelineAmazonS3GenericCompressionGzipAsObservabilityPipelineAmazonS3GenericCompression(v *ObservabilityPipelineAmazonS3GenericCompressionGzip) ObservabilityPipelineAmazonS3GenericCompression {
	return ObservabilityPipelineAmazonS3GenericCompression{ObservabilityPipelineAmazonS3GenericCompressionGzip: v}
}

// ObservabilityPipelineAmazonS3GenericCompressionSnappyAsObservabilityPipelineAmazonS3GenericCompression is a convenience function that returns ObservabilityPipelineAmazonS3GenericCompressionSnappy wrapped in ObservabilityPipelineAmazonS3GenericCompression.
func ObservabilityPipelineAmazonS3GenericCompressionSnappyAsObservabilityPipelineAmazonS3GenericCompression(v *ObservabilityPipelineAmazonS3GenericCompressionSnappy) ObservabilityPipelineAmazonS3GenericCompression {
	return ObservabilityPipelineAmazonS3GenericCompression{ObservabilityPipelineAmazonS3GenericCompressionSnappy: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineAmazonS3GenericCompression) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineAmazonS3GenericCompressionZstd
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAmazonS3GenericCompressionZstd)
	if err == nil {
		if obj.ObservabilityPipelineAmazonS3GenericCompressionZstd != nil && obj.ObservabilityPipelineAmazonS3GenericCompressionZstd.UnparsedObject == nil {
			jsonObservabilityPipelineAmazonS3GenericCompressionZstd, _ := datadog.Marshal(obj.ObservabilityPipelineAmazonS3GenericCompressionZstd)
			if string(jsonObservabilityPipelineAmazonS3GenericCompressionZstd) == "{}" { // empty struct
				obj.ObservabilityPipelineAmazonS3GenericCompressionZstd = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAmazonS3GenericCompressionZstd = nil
		}
	} else {
		obj.ObservabilityPipelineAmazonS3GenericCompressionZstd = nil
	}

	// try to unmarshal data into ObservabilityPipelineAmazonS3GenericCompressionGzip
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAmazonS3GenericCompressionGzip)
	if err == nil {
		if obj.ObservabilityPipelineAmazonS3GenericCompressionGzip != nil && obj.ObservabilityPipelineAmazonS3GenericCompressionGzip.UnparsedObject == nil {
			jsonObservabilityPipelineAmazonS3GenericCompressionGzip, _ := datadog.Marshal(obj.ObservabilityPipelineAmazonS3GenericCompressionGzip)
			if string(jsonObservabilityPipelineAmazonS3GenericCompressionGzip) == "{}" { // empty struct
				obj.ObservabilityPipelineAmazonS3GenericCompressionGzip = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAmazonS3GenericCompressionGzip = nil
		}
	} else {
		obj.ObservabilityPipelineAmazonS3GenericCompressionGzip = nil
	}

	// try to unmarshal data into ObservabilityPipelineAmazonS3GenericCompressionSnappy
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineAmazonS3GenericCompressionSnappy)
	if err == nil {
		if obj.ObservabilityPipelineAmazonS3GenericCompressionSnappy != nil && obj.ObservabilityPipelineAmazonS3GenericCompressionSnappy.UnparsedObject == nil {
			jsonObservabilityPipelineAmazonS3GenericCompressionSnappy, _ := datadog.Marshal(obj.ObservabilityPipelineAmazonS3GenericCompressionSnappy)
			if string(jsonObservabilityPipelineAmazonS3GenericCompressionSnappy) == "{}" { // empty struct
				obj.ObservabilityPipelineAmazonS3GenericCompressionSnappy = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineAmazonS3GenericCompressionSnappy = nil
		}
	} else {
		obj.ObservabilityPipelineAmazonS3GenericCompressionSnappy = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineAmazonS3GenericCompressionZstd = nil
		obj.ObservabilityPipelineAmazonS3GenericCompressionGzip = nil
		obj.ObservabilityPipelineAmazonS3GenericCompressionSnappy = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineAmazonS3GenericCompression) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineAmazonS3GenericCompressionZstd != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAmazonS3GenericCompressionZstd)
	}

	if obj.ObservabilityPipelineAmazonS3GenericCompressionGzip != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAmazonS3GenericCompressionGzip)
	}

	if obj.ObservabilityPipelineAmazonS3GenericCompressionSnappy != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineAmazonS3GenericCompressionSnappy)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineAmazonS3GenericCompression) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineAmazonS3GenericCompressionZstd != nil {
		return obj.ObservabilityPipelineAmazonS3GenericCompressionZstd
	}

	if obj.ObservabilityPipelineAmazonS3GenericCompressionGzip != nil {
		return obj.ObservabilityPipelineAmazonS3GenericCompressionGzip
	}

	if obj.ObservabilityPipelineAmazonS3GenericCompressionSnappy != nil {
		return obj.ObservabilityPipelineAmazonS3GenericCompressionSnappy
	}

	// all schemas are nil
	return nil
}
