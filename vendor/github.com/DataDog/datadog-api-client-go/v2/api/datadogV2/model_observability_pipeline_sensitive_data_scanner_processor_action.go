// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorAction - Defines what action to take when sensitive data is matched.
type ObservabilityPipelineSensitiveDataScannerProcessorAction struct {
	ObservabilityPipelineSensitiveDataScannerProcessorActionRedact        *ObservabilityPipelineSensitiveDataScannerProcessorActionRedact
	ObservabilityPipelineSensitiveDataScannerProcessorActionHash          *ObservabilityPipelineSensitiveDataScannerProcessorActionHash
	ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAsObservabilityPipelineSensitiveDataScannerProcessorAction is a convenience function that returns ObservabilityPipelineSensitiveDataScannerProcessorActionRedact wrapped in ObservabilityPipelineSensitiveDataScannerProcessorAction.
func ObservabilityPipelineSensitiveDataScannerProcessorActionRedactAsObservabilityPipelineSensitiveDataScannerProcessorAction(v *ObservabilityPipelineSensitiveDataScannerProcessorActionRedact) ObservabilityPipelineSensitiveDataScannerProcessorAction {
	return ObservabilityPipelineSensitiveDataScannerProcessorAction{ObservabilityPipelineSensitiveDataScannerProcessorActionRedact: v}
}

// ObservabilityPipelineSensitiveDataScannerProcessorActionHashAsObservabilityPipelineSensitiveDataScannerProcessorAction is a convenience function that returns ObservabilityPipelineSensitiveDataScannerProcessorActionHash wrapped in ObservabilityPipelineSensitiveDataScannerProcessorAction.
func ObservabilityPipelineSensitiveDataScannerProcessorActionHashAsObservabilityPipelineSensitiveDataScannerProcessorAction(v *ObservabilityPipelineSensitiveDataScannerProcessorActionHash) ObservabilityPipelineSensitiveDataScannerProcessorAction {
	return ObservabilityPipelineSensitiveDataScannerProcessorAction{ObservabilityPipelineSensitiveDataScannerProcessorActionHash: v}
}

// ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAsObservabilityPipelineSensitiveDataScannerProcessorAction is a convenience function that returns ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact wrapped in ObservabilityPipelineSensitiveDataScannerProcessorAction.
func ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedactAsObservabilityPipelineSensitiveDataScannerProcessorAction(v *ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact) ObservabilityPipelineSensitiveDataScannerProcessorAction {
	return ObservabilityPipelineSensitiveDataScannerProcessorAction{ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineSensitiveDataScannerProcessorAction) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineSensitiveDataScannerProcessorActionRedact
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSensitiveDataScannerProcessorActionRedact)
	if err == nil {
		if obj.ObservabilityPipelineSensitiveDataScannerProcessorActionRedact != nil && obj.ObservabilityPipelineSensitiveDataScannerProcessorActionRedact.UnparsedObject == nil {
			jsonObservabilityPipelineSensitiveDataScannerProcessorActionRedact, _ := datadog.Marshal(obj.ObservabilityPipelineSensitiveDataScannerProcessorActionRedact)
			if string(jsonObservabilityPipelineSensitiveDataScannerProcessorActionRedact) == "{}" { // empty struct
				obj.ObservabilityPipelineSensitiveDataScannerProcessorActionRedact = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSensitiveDataScannerProcessorActionRedact = nil
		}
	} else {
		obj.ObservabilityPipelineSensitiveDataScannerProcessorActionRedact = nil
	}

	// try to unmarshal data into ObservabilityPipelineSensitiveDataScannerProcessorActionHash
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSensitiveDataScannerProcessorActionHash)
	if err == nil {
		if obj.ObservabilityPipelineSensitiveDataScannerProcessorActionHash != nil && obj.ObservabilityPipelineSensitiveDataScannerProcessorActionHash.UnparsedObject == nil {
			jsonObservabilityPipelineSensitiveDataScannerProcessorActionHash, _ := datadog.Marshal(obj.ObservabilityPipelineSensitiveDataScannerProcessorActionHash)
			if string(jsonObservabilityPipelineSensitiveDataScannerProcessorActionHash) == "{}" { // empty struct
				obj.ObservabilityPipelineSensitiveDataScannerProcessorActionHash = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSensitiveDataScannerProcessorActionHash = nil
		}
	} else {
		obj.ObservabilityPipelineSensitiveDataScannerProcessorActionHash = nil
	}

	// try to unmarshal data into ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact)
	if err == nil {
		if obj.ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact != nil && obj.ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact.UnparsedObject == nil {
			jsonObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact, _ := datadog.Marshal(obj.ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact)
			if string(jsonObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact) == "{}" { // empty struct
				obj.ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact = nil
		}
	} else {
		obj.ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineSensitiveDataScannerProcessorActionRedact = nil
		obj.ObservabilityPipelineSensitiveDataScannerProcessorActionHash = nil
		obj.ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineSensitiveDataScannerProcessorAction) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineSensitiveDataScannerProcessorActionRedact != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSensitiveDataScannerProcessorActionRedact)
	}

	if obj.ObservabilityPipelineSensitiveDataScannerProcessorActionHash != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSensitiveDataScannerProcessorActionHash)
	}

	if obj.ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineSensitiveDataScannerProcessorAction) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineSensitiveDataScannerProcessorActionRedact != nil {
		return obj.ObservabilityPipelineSensitiveDataScannerProcessorActionRedact
	}

	if obj.ObservabilityPipelineSensitiveDataScannerProcessorActionHash != nil {
		return obj.ObservabilityPipelineSensitiveDataScannerProcessorActionHash
	}

	if obj.ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact != nil {
		return obj.ObservabilityPipelineSensitiveDataScannerProcessorActionPartialRedact
	}

	// all schemas are nil
	return nil
}
