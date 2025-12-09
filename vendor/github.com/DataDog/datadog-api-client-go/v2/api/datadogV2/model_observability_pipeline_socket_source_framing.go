// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSocketSourceFraming - Framing method configuration for the socket source.
type ObservabilityPipelineSocketSourceFraming struct {
	ObservabilityPipelineSocketSourceFramingNewlineDelimited   *ObservabilityPipelineSocketSourceFramingNewlineDelimited
	ObservabilityPipelineSocketSourceFramingBytes              *ObservabilityPipelineSocketSourceFramingBytes
	ObservabilityPipelineSocketSourceFramingCharacterDelimited *ObservabilityPipelineSocketSourceFramingCharacterDelimited
	ObservabilityPipelineSocketSourceFramingOctetCounting      *ObservabilityPipelineSocketSourceFramingOctetCounting
	ObservabilityPipelineSocketSourceFramingChunkedGelf        *ObservabilityPipelineSocketSourceFramingChunkedGelf

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineSocketSourceFramingNewlineDelimitedAsObservabilityPipelineSocketSourceFraming is a convenience function that returns ObservabilityPipelineSocketSourceFramingNewlineDelimited wrapped in ObservabilityPipelineSocketSourceFraming.
func ObservabilityPipelineSocketSourceFramingNewlineDelimitedAsObservabilityPipelineSocketSourceFraming(v *ObservabilityPipelineSocketSourceFramingNewlineDelimited) ObservabilityPipelineSocketSourceFraming {
	return ObservabilityPipelineSocketSourceFraming{ObservabilityPipelineSocketSourceFramingNewlineDelimited: v}
}

// ObservabilityPipelineSocketSourceFramingBytesAsObservabilityPipelineSocketSourceFraming is a convenience function that returns ObservabilityPipelineSocketSourceFramingBytes wrapped in ObservabilityPipelineSocketSourceFraming.
func ObservabilityPipelineSocketSourceFramingBytesAsObservabilityPipelineSocketSourceFraming(v *ObservabilityPipelineSocketSourceFramingBytes) ObservabilityPipelineSocketSourceFraming {
	return ObservabilityPipelineSocketSourceFraming{ObservabilityPipelineSocketSourceFramingBytes: v}
}

// ObservabilityPipelineSocketSourceFramingCharacterDelimitedAsObservabilityPipelineSocketSourceFraming is a convenience function that returns ObservabilityPipelineSocketSourceFramingCharacterDelimited wrapped in ObservabilityPipelineSocketSourceFraming.
func ObservabilityPipelineSocketSourceFramingCharacterDelimitedAsObservabilityPipelineSocketSourceFraming(v *ObservabilityPipelineSocketSourceFramingCharacterDelimited) ObservabilityPipelineSocketSourceFraming {
	return ObservabilityPipelineSocketSourceFraming{ObservabilityPipelineSocketSourceFramingCharacterDelimited: v}
}

// ObservabilityPipelineSocketSourceFramingOctetCountingAsObservabilityPipelineSocketSourceFraming is a convenience function that returns ObservabilityPipelineSocketSourceFramingOctetCounting wrapped in ObservabilityPipelineSocketSourceFraming.
func ObservabilityPipelineSocketSourceFramingOctetCountingAsObservabilityPipelineSocketSourceFraming(v *ObservabilityPipelineSocketSourceFramingOctetCounting) ObservabilityPipelineSocketSourceFraming {
	return ObservabilityPipelineSocketSourceFraming{ObservabilityPipelineSocketSourceFramingOctetCounting: v}
}

// ObservabilityPipelineSocketSourceFramingChunkedGelfAsObservabilityPipelineSocketSourceFraming is a convenience function that returns ObservabilityPipelineSocketSourceFramingChunkedGelf wrapped in ObservabilityPipelineSocketSourceFraming.
func ObservabilityPipelineSocketSourceFramingChunkedGelfAsObservabilityPipelineSocketSourceFraming(v *ObservabilityPipelineSocketSourceFramingChunkedGelf) ObservabilityPipelineSocketSourceFraming {
	return ObservabilityPipelineSocketSourceFraming{ObservabilityPipelineSocketSourceFramingChunkedGelf: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineSocketSourceFraming) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineSocketSourceFramingNewlineDelimited
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSocketSourceFramingNewlineDelimited)
	if err == nil {
		if obj.ObservabilityPipelineSocketSourceFramingNewlineDelimited != nil && obj.ObservabilityPipelineSocketSourceFramingNewlineDelimited.UnparsedObject == nil {
			jsonObservabilityPipelineSocketSourceFramingNewlineDelimited, _ := datadog.Marshal(obj.ObservabilityPipelineSocketSourceFramingNewlineDelimited)
			if string(jsonObservabilityPipelineSocketSourceFramingNewlineDelimited) == "{}" { // empty struct
				obj.ObservabilityPipelineSocketSourceFramingNewlineDelimited = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSocketSourceFramingNewlineDelimited = nil
		}
	} else {
		obj.ObservabilityPipelineSocketSourceFramingNewlineDelimited = nil
	}

	// try to unmarshal data into ObservabilityPipelineSocketSourceFramingBytes
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSocketSourceFramingBytes)
	if err == nil {
		if obj.ObservabilityPipelineSocketSourceFramingBytes != nil && obj.ObservabilityPipelineSocketSourceFramingBytes.UnparsedObject == nil {
			jsonObservabilityPipelineSocketSourceFramingBytes, _ := datadog.Marshal(obj.ObservabilityPipelineSocketSourceFramingBytes)
			if string(jsonObservabilityPipelineSocketSourceFramingBytes) == "{}" { // empty struct
				obj.ObservabilityPipelineSocketSourceFramingBytes = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSocketSourceFramingBytes = nil
		}
	} else {
		obj.ObservabilityPipelineSocketSourceFramingBytes = nil
	}

	// try to unmarshal data into ObservabilityPipelineSocketSourceFramingCharacterDelimited
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSocketSourceFramingCharacterDelimited)
	if err == nil {
		if obj.ObservabilityPipelineSocketSourceFramingCharacterDelimited != nil && obj.ObservabilityPipelineSocketSourceFramingCharacterDelimited.UnparsedObject == nil {
			jsonObservabilityPipelineSocketSourceFramingCharacterDelimited, _ := datadog.Marshal(obj.ObservabilityPipelineSocketSourceFramingCharacterDelimited)
			if string(jsonObservabilityPipelineSocketSourceFramingCharacterDelimited) == "{}" { // empty struct
				obj.ObservabilityPipelineSocketSourceFramingCharacterDelimited = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSocketSourceFramingCharacterDelimited = nil
		}
	} else {
		obj.ObservabilityPipelineSocketSourceFramingCharacterDelimited = nil
	}

	// try to unmarshal data into ObservabilityPipelineSocketSourceFramingOctetCounting
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSocketSourceFramingOctetCounting)
	if err == nil {
		if obj.ObservabilityPipelineSocketSourceFramingOctetCounting != nil && obj.ObservabilityPipelineSocketSourceFramingOctetCounting.UnparsedObject == nil {
			jsonObservabilityPipelineSocketSourceFramingOctetCounting, _ := datadog.Marshal(obj.ObservabilityPipelineSocketSourceFramingOctetCounting)
			if string(jsonObservabilityPipelineSocketSourceFramingOctetCounting) == "{}" { // empty struct
				obj.ObservabilityPipelineSocketSourceFramingOctetCounting = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSocketSourceFramingOctetCounting = nil
		}
	} else {
		obj.ObservabilityPipelineSocketSourceFramingOctetCounting = nil
	}

	// try to unmarshal data into ObservabilityPipelineSocketSourceFramingChunkedGelf
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSocketSourceFramingChunkedGelf)
	if err == nil {
		if obj.ObservabilityPipelineSocketSourceFramingChunkedGelf != nil && obj.ObservabilityPipelineSocketSourceFramingChunkedGelf.UnparsedObject == nil {
			jsonObservabilityPipelineSocketSourceFramingChunkedGelf, _ := datadog.Marshal(obj.ObservabilityPipelineSocketSourceFramingChunkedGelf)
			if string(jsonObservabilityPipelineSocketSourceFramingChunkedGelf) == "{}" { // empty struct
				obj.ObservabilityPipelineSocketSourceFramingChunkedGelf = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSocketSourceFramingChunkedGelf = nil
		}
	} else {
		obj.ObservabilityPipelineSocketSourceFramingChunkedGelf = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineSocketSourceFramingNewlineDelimited = nil
		obj.ObservabilityPipelineSocketSourceFramingBytes = nil
		obj.ObservabilityPipelineSocketSourceFramingCharacterDelimited = nil
		obj.ObservabilityPipelineSocketSourceFramingOctetCounting = nil
		obj.ObservabilityPipelineSocketSourceFramingChunkedGelf = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineSocketSourceFraming) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineSocketSourceFramingNewlineDelimited != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSocketSourceFramingNewlineDelimited)
	}

	if obj.ObservabilityPipelineSocketSourceFramingBytes != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSocketSourceFramingBytes)
	}

	if obj.ObservabilityPipelineSocketSourceFramingCharacterDelimited != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSocketSourceFramingCharacterDelimited)
	}

	if obj.ObservabilityPipelineSocketSourceFramingOctetCounting != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSocketSourceFramingOctetCounting)
	}

	if obj.ObservabilityPipelineSocketSourceFramingChunkedGelf != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSocketSourceFramingChunkedGelf)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineSocketSourceFraming) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineSocketSourceFramingNewlineDelimited != nil {
		return obj.ObservabilityPipelineSocketSourceFramingNewlineDelimited
	}

	if obj.ObservabilityPipelineSocketSourceFramingBytes != nil {
		return obj.ObservabilityPipelineSocketSourceFramingBytes
	}

	if obj.ObservabilityPipelineSocketSourceFramingCharacterDelimited != nil {
		return obj.ObservabilityPipelineSocketSourceFramingCharacterDelimited
	}

	if obj.ObservabilityPipelineSocketSourceFramingOctetCounting != nil {
		return obj.ObservabilityPipelineSocketSourceFramingOctetCounting
	}

	if obj.ObservabilityPipelineSocketSourceFramingChunkedGelf != nil {
		return obj.ObservabilityPipelineSocketSourceFramingChunkedGelf
	}

	// all schemas are nil
	return nil
}
