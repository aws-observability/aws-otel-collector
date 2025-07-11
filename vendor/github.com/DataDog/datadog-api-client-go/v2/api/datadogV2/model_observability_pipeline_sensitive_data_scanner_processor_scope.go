// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSensitiveDataScannerProcessorScope - Determines which parts of the log the pattern-matching rule should be applied to.
type ObservabilityPipelineSensitiveDataScannerProcessorScope struct {
	ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude *ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude
	ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude *ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude
	ObservabilityPipelineSensitiveDataScannerProcessorScopeAll     *ObservabilityPipelineSensitiveDataScannerProcessorScopeAll

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeAsObservabilityPipelineSensitiveDataScannerProcessorScope is a convenience function that returns ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude wrapped in ObservabilityPipelineSensitiveDataScannerProcessorScope.
func ObservabilityPipelineSensitiveDataScannerProcessorScopeIncludeAsObservabilityPipelineSensitiveDataScannerProcessorScope(v *ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude) ObservabilityPipelineSensitiveDataScannerProcessorScope {
	return ObservabilityPipelineSensitiveDataScannerProcessorScope{ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude: v}
}

// ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeAsObservabilityPipelineSensitiveDataScannerProcessorScope is a convenience function that returns ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude wrapped in ObservabilityPipelineSensitiveDataScannerProcessorScope.
func ObservabilityPipelineSensitiveDataScannerProcessorScopeExcludeAsObservabilityPipelineSensitiveDataScannerProcessorScope(v *ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude) ObservabilityPipelineSensitiveDataScannerProcessorScope {
	return ObservabilityPipelineSensitiveDataScannerProcessorScope{ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude: v}
}

// ObservabilityPipelineSensitiveDataScannerProcessorScopeAllAsObservabilityPipelineSensitiveDataScannerProcessorScope is a convenience function that returns ObservabilityPipelineSensitiveDataScannerProcessorScopeAll wrapped in ObservabilityPipelineSensitiveDataScannerProcessorScope.
func ObservabilityPipelineSensitiveDataScannerProcessorScopeAllAsObservabilityPipelineSensitiveDataScannerProcessorScope(v *ObservabilityPipelineSensitiveDataScannerProcessorScopeAll) ObservabilityPipelineSensitiveDataScannerProcessorScope {
	return ObservabilityPipelineSensitiveDataScannerProcessorScope{ObservabilityPipelineSensitiveDataScannerProcessorScopeAll: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineSensitiveDataScannerProcessorScope) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude)
	if err == nil {
		if obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude != nil && obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude.UnparsedObject == nil {
			jsonObservabilityPipelineSensitiveDataScannerProcessorScopeInclude, _ := datadog.Marshal(obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude)
			if string(jsonObservabilityPipelineSensitiveDataScannerProcessorScopeInclude) == "{}" { // empty struct
				obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude = nil
		}
	} else {
		obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude = nil
	}

	// try to unmarshal data into ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude)
	if err == nil {
		if obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude != nil && obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude.UnparsedObject == nil {
			jsonObservabilityPipelineSensitiveDataScannerProcessorScopeExclude, _ := datadog.Marshal(obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude)
			if string(jsonObservabilityPipelineSensitiveDataScannerProcessorScopeExclude) == "{}" { // empty struct
				obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude = nil
		}
	} else {
		obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude = nil
	}

	// try to unmarshal data into ObservabilityPipelineSensitiveDataScannerProcessorScopeAll
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeAll)
	if err == nil {
		if obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeAll != nil && obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeAll.UnparsedObject == nil {
			jsonObservabilityPipelineSensitiveDataScannerProcessorScopeAll, _ := datadog.Marshal(obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeAll)
			if string(jsonObservabilityPipelineSensitiveDataScannerProcessorScopeAll) == "{}" { // empty struct
				obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeAll = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeAll = nil
		}
	} else {
		obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeAll = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude = nil
		obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude = nil
		obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeAll = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineSensitiveDataScannerProcessorScope) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude)
	}

	if obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude)
	}

	if obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeAll != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeAll)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineSensitiveDataScannerProcessorScope) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude != nil {
		return obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeInclude
	}

	if obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude != nil {
		return obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeExclude
	}

	if obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeAll != nil {
		return obj.ObservabilityPipelineSensitiveDataScannerProcessorScopeAll
	}

	// all schemas are nil
	return nil
}
