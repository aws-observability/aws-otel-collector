// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketDestinationFraming - Framing method configuration.
type ObservabilityPipelineSocketDestinationFraming struct {
	ObservabilityPipelineSocketDestinationFramingNewlineDelimited   *ObservabilityPipelineSocketDestinationFramingNewlineDelimited
	ObservabilityPipelineSocketDestinationFramingBytes              *ObservabilityPipelineSocketDestinationFramingBytes
	ObservabilityPipelineSocketDestinationFramingCharacterDelimited *ObservabilityPipelineSocketDestinationFramingCharacterDelimited

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineSocketDestinationFramingNewlineDelimitedAsObservabilityPipelineSocketDestinationFraming is a convenience function that returns ObservabilityPipelineSocketDestinationFramingNewlineDelimited wrapped in ObservabilityPipelineSocketDestinationFraming.
func ObservabilityPipelineSocketDestinationFramingNewlineDelimitedAsObservabilityPipelineSocketDestinationFraming(v *ObservabilityPipelineSocketDestinationFramingNewlineDelimited) ObservabilityPipelineSocketDestinationFraming {
	return ObservabilityPipelineSocketDestinationFraming{ObservabilityPipelineSocketDestinationFramingNewlineDelimited: v}
}

// ObservabilityPipelineSocketDestinationFramingBytesAsObservabilityPipelineSocketDestinationFraming is a convenience function that returns ObservabilityPipelineSocketDestinationFramingBytes wrapped in ObservabilityPipelineSocketDestinationFraming.
func ObservabilityPipelineSocketDestinationFramingBytesAsObservabilityPipelineSocketDestinationFraming(v *ObservabilityPipelineSocketDestinationFramingBytes) ObservabilityPipelineSocketDestinationFraming {
	return ObservabilityPipelineSocketDestinationFraming{ObservabilityPipelineSocketDestinationFramingBytes: v}
}

// ObservabilityPipelineSocketDestinationFramingCharacterDelimitedAsObservabilityPipelineSocketDestinationFraming is a convenience function that returns ObservabilityPipelineSocketDestinationFramingCharacterDelimited wrapped in ObservabilityPipelineSocketDestinationFraming.
func ObservabilityPipelineSocketDestinationFramingCharacterDelimitedAsObservabilityPipelineSocketDestinationFraming(v *ObservabilityPipelineSocketDestinationFramingCharacterDelimited) ObservabilityPipelineSocketDestinationFraming {
	return ObservabilityPipelineSocketDestinationFraming{ObservabilityPipelineSocketDestinationFramingCharacterDelimited: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineSocketDestinationFraming) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineSocketDestinationFramingNewlineDelimited
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSocketDestinationFramingNewlineDelimited)
	if err == nil {
		if obj.ObservabilityPipelineSocketDestinationFramingNewlineDelimited != nil && obj.ObservabilityPipelineSocketDestinationFramingNewlineDelimited.UnparsedObject == nil {
			jsonObservabilityPipelineSocketDestinationFramingNewlineDelimited, _ := datadog.Marshal(obj.ObservabilityPipelineSocketDestinationFramingNewlineDelimited)
			if string(jsonObservabilityPipelineSocketDestinationFramingNewlineDelimited) == "{}" { // empty struct
				obj.ObservabilityPipelineSocketDestinationFramingNewlineDelimited = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSocketDestinationFramingNewlineDelimited = nil
		}
	} else {
		obj.ObservabilityPipelineSocketDestinationFramingNewlineDelimited = nil
	}

	// try to unmarshal data into ObservabilityPipelineSocketDestinationFramingBytes
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSocketDestinationFramingBytes)
	if err == nil {
		if obj.ObservabilityPipelineSocketDestinationFramingBytes != nil && obj.ObservabilityPipelineSocketDestinationFramingBytes.UnparsedObject == nil {
			jsonObservabilityPipelineSocketDestinationFramingBytes, _ := datadog.Marshal(obj.ObservabilityPipelineSocketDestinationFramingBytes)
			if string(jsonObservabilityPipelineSocketDestinationFramingBytes) == "{}" { // empty struct
				obj.ObservabilityPipelineSocketDestinationFramingBytes = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSocketDestinationFramingBytes = nil
		}
	} else {
		obj.ObservabilityPipelineSocketDestinationFramingBytes = nil
	}

	// try to unmarshal data into ObservabilityPipelineSocketDestinationFramingCharacterDelimited
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSocketDestinationFramingCharacterDelimited)
	if err == nil {
		if obj.ObservabilityPipelineSocketDestinationFramingCharacterDelimited != nil && obj.ObservabilityPipelineSocketDestinationFramingCharacterDelimited.UnparsedObject == nil {
			jsonObservabilityPipelineSocketDestinationFramingCharacterDelimited, _ := datadog.Marshal(obj.ObservabilityPipelineSocketDestinationFramingCharacterDelimited)
			if string(jsonObservabilityPipelineSocketDestinationFramingCharacterDelimited) == "{}" { // empty struct
				obj.ObservabilityPipelineSocketDestinationFramingCharacterDelimited = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSocketDestinationFramingCharacterDelimited = nil
		}
	} else {
		obj.ObservabilityPipelineSocketDestinationFramingCharacterDelimited = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineSocketDestinationFramingNewlineDelimited = nil
		obj.ObservabilityPipelineSocketDestinationFramingBytes = nil
		obj.ObservabilityPipelineSocketDestinationFramingCharacterDelimited = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineSocketDestinationFraming) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineSocketDestinationFramingNewlineDelimited != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSocketDestinationFramingNewlineDelimited)
	}

	if obj.ObservabilityPipelineSocketDestinationFramingBytes != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSocketDestinationFramingBytes)
	}

	if obj.ObservabilityPipelineSocketDestinationFramingCharacterDelimited != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSocketDestinationFramingCharacterDelimited)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineSocketDestinationFraming) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineSocketDestinationFramingNewlineDelimited != nil {
		return obj.ObservabilityPipelineSocketDestinationFramingNewlineDelimited
	}

	if obj.ObservabilityPipelineSocketDestinationFramingBytes != nil {
		return obj.ObservabilityPipelineSocketDestinationFramingBytes
	}

	if obj.ObservabilityPipelineSocketDestinationFramingCharacterDelimited != nil {
		return obj.ObservabilityPipelineSocketDestinationFramingCharacterDelimited
	}

	// all schemas are nil
	return nil
}
