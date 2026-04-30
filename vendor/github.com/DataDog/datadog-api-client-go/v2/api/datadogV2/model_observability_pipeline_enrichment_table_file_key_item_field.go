// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableFileKeyItemField - Specifies the source of the key value used for enrichment table lookups.
// Can be a plain field path string or an object specifying `event`, `vrl`, or `secret`.
type ObservabilityPipelineEnrichmentTableFileKeyItemField struct {
	ObservabilityPipelineEnrichmentTableFieldStringPath   *string
	ObservabilityPipelineEnrichmentTableFieldEventLookup  *ObservabilityPipelineEnrichmentTableFieldEventLookup
	ObservabilityPipelineEnrichmentTableFieldVrlLookup    *ObservabilityPipelineEnrichmentTableFieldVrlLookup
	ObservabilityPipelineEnrichmentTableFieldSecretLookup *ObservabilityPipelineEnrichmentTableFieldSecretLookup

	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject interface{}
}

// ObservabilityPipelineEnrichmentTableFieldStringPathAsObservabilityPipelineEnrichmentTableFileKeyItemField is a convenience function that returns string wrapped in ObservabilityPipelineEnrichmentTableFileKeyItemField.
func ObservabilityPipelineEnrichmentTableFieldStringPathAsObservabilityPipelineEnrichmentTableFileKeyItemField(v *string) ObservabilityPipelineEnrichmentTableFileKeyItemField {
	return ObservabilityPipelineEnrichmentTableFileKeyItemField{ObservabilityPipelineEnrichmentTableFieldStringPath: v}
}

// ObservabilityPipelineEnrichmentTableFieldEventLookupAsObservabilityPipelineEnrichmentTableFileKeyItemField is a convenience function that returns ObservabilityPipelineEnrichmentTableFieldEventLookup wrapped in ObservabilityPipelineEnrichmentTableFileKeyItemField.
func ObservabilityPipelineEnrichmentTableFieldEventLookupAsObservabilityPipelineEnrichmentTableFileKeyItemField(v *ObservabilityPipelineEnrichmentTableFieldEventLookup) ObservabilityPipelineEnrichmentTableFileKeyItemField {
	return ObservabilityPipelineEnrichmentTableFileKeyItemField{ObservabilityPipelineEnrichmentTableFieldEventLookup: v}
}

// ObservabilityPipelineEnrichmentTableFieldVrlLookupAsObservabilityPipelineEnrichmentTableFileKeyItemField is a convenience function that returns ObservabilityPipelineEnrichmentTableFieldVrlLookup wrapped in ObservabilityPipelineEnrichmentTableFileKeyItemField.
func ObservabilityPipelineEnrichmentTableFieldVrlLookupAsObservabilityPipelineEnrichmentTableFileKeyItemField(v *ObservabilityPipelineEnrichmentTableFieldVrlLookup) ObservabilityPipelineEnrichmentTableFileKeyItemField {
	return ObservabilityPipelineEnrichmentTableFileKeyItemField{ObservabilityPipelineEnrichmentTableFieldVrlLookup: v}
}

// ObservabilityPipelineEnrichmentTableFieldSecretLookupAsObservabilityPipelineEnrichmentTableFileKeyItemField is a convenience function that returns ObservabilityPipelineEnrichmentTableFieldSecretLookup wrapped in ObservabilityPipelineEnrichmentTableFileKeyItemField.
func ObservabilityPipelineEnrichmentTableFieldSecretLookupAsObservabilityPipelineEnrichmentTableFileKeyItemField(v *ObservabilityPipelineEnrichmentTableFieldSecretLookup) ObservabilityPipelineEnrichmentTableFileKeyItemField {
	return ObservabilityPipelineEnrichmentTableFileKeyItemField{ObservabilityPipelineEnrichmentTableFieldSecretLookup: v}
}

// UnmarshalJSON turns data into one of the pointers in the struct.
func (obj *ObservabilityPipelineEnrichmentTableFileKeyItemField) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ObservabilityPipelineEnrichmentTableFieldStringPath
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineEnrichmentTableFieldStringPath)
	if err == nil {
		if obj.ObservabilityPipelineEnrichmentTableFieldStringPath != nil {
			jsonObservabilityPipelineEnrichmentTableFieldStringPath, _ := datadog.Marshal(obj.ObservabilityPipelineEnrichmentTableFieldStringPath)
			if string(jsonObservabilityPipelineEnrichmentTableFieldStringPath) == "{}" { // empty struct
				obj.ObservabilityPipelineEnrichmentTableFieldStringPath = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineEnrichmentTableFieldStringPath = nil
		}
	} else {
		obj.ObservabilityPipelineEnrichmentTableFieldStringPath = nil
	}

	// try to unmarshal data into ObservabilityPipelineEnrichmentTableFieldEventLookup
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineEnrichmentTableFieldEventLookup)
	if err == nil {
		if obj.ObservabilityPipelineEnrichmentTableFieldEventLookup != nil && obj.ObservabilityPipelineEnrichmentTableFieldEventLookup.UnparsedObject == nil {
			jsonObservabilityPipelineEnrichmentTableFieldEventLookup, _ := datadog.Marshal(obj.ObservabilityPipelineEnrichmentTableFieldEventLookup)
			if string(jsonObservabilityPipelineEnrichmentTableFieldEventLookup) == "{}" { // empty struct
				obj.ObservabilityPipelineEnrichmentTableFieldEventLookup = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineEnrichmentTableFieldEventLookup = nil
		}
	} else {
		obj.ObservabilityPipelineEnrichmentTableFieldEventLookup = nil
	}

	// try to unmarshal data into ObservabilityPipelineEnrichmentTableFieldVrlLookup
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineEnrichmentTableFieldVrlLookup)
	if err == nil {
		if obj.ObservabilityPipelineEnrichmentTableFieldVrlLookup != nil && obj.ObservabilityPipelineEnrichmentTableFieldVrlLookup.UnparsedObject == nil {
			jsonObservabilityPipelineEnrichmentTableFieldVrlLookup, _ := datadog.Marshal(obj.ObservabilityPipelineEnrichmentTableFieldVrlLookup)
			if string(jsonObservabilityPipelineEnrichmentTableFieldVrlLookup) == "{}" { // empty struct
				obj.ObservabilityPipelineEnrichmentTableFieldVrlLookup = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineEnrichmentTableFieldVrlLookup = nil
		}
	} else {
		obj.ObservabilityPipelineEnrichmentTableFieldVrlLookup = nil
	}

	// try to unmarshal data into ObservabilityPipelineEnrichmentTableFieldSecretLookup
	err = datadog.Unmarshal(data, &obj.ObservabilityPipelineEnrichmentTableFieldSecretLookup)
	if err == nil {
		if obj.ObservabilityPipelineEnrichmentTableFieldSecretLookup != nil && obj.ObservabilityPipelineEnrichmentTableFieldSecretLookup.UnparsedObject == nil {
			jsonObservabilityPipelineEnrichmentTableFieldSecretLookup, _ := datadog.Marshal(obj.ObservabilityPipelineEnrichmentTableFieldSecretLookup)
			if string(jsonObservabilityPipelineEnrichmentTableFieldSecretLookup) == "{}" { // empty struct
				obj.ObservabilityPipelineEnrichmentTableFieldSecretLookup = nil
			} else {
				match++
			}
		} else {
			obj.ObservabilityPipelineEnrichmentTableFieldSecretLookup = nil
		}
	} else {
		obj.ObservabilityPipelineEnrichmentTableFieldSecretLookup = nil
	}

	if match != 1 { // more than 1 match
		// reset to nil
		obj.ObservabilityPipelineEnrichmentTableFieldStringPath = nil
		obj.ObservabilityPipelineEnrichmentTableFieldEventLookup = nil
		obj.ObservabilityPipelineEnrichmentTableFieldVrlLookup = nil
		obj.ObservabilityPipelineEnrichmentTableFieldSecretLookup = nil
		return datadog.Unmarshal(data, &obj.UnparsedObject)
	}
	return nil // exactly one match
}

// MarshalJSON turns data from the first non-nil pointers in the struct to JSON.
func (obj ObservabilityPipelineEnrichmentTableFileKeyItemField) MarshalJSON() ([]byte, error) {
	if obj.ObservabilityPipelineEnrichmentTableFieldStringPath != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineEnrichmentTableFieldStringPath)
	}

	if obj.ObservabilityPipelineEnrichmentTableFieldEventLookup != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineEnrichmentTableFieldEventLookup)
	}

	if obj.ObservabilityPipelineEnrichmentTableFieldVrlLookup != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineEnrichmentTableFieldVrlLookup)
	}

	if obj.ObservabilityPipelineEnrichmentTableFieldSecretLookup != nil {
		return datadog.Marshal(&obj.ObservabilityPipelineEnrichmentTableFieldSecretLookup)
	}

	if obj.UnparsedObject != nil {
		return datadog.Marshal(obj.UnparsedObject)
	}
	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns the actual instance.
func (obj *ObservabilityPipelineEnrichmentTableFileKeyItemField) GetActualInstance() interface{} {
	if obj.ObservabilityPipelineEnrichmentTableFieldStringPath != nil {
		return obj.ObservabilityPipelineEnrichmentTableFieldStringPath
	}

	if obj.ObservabilityPipelineEnrichmentTableFieldEventLookup != nil {
		return obj.ObservabilityPipelineEnrichmentTableFieldEventLookup
	}

	if obj.ObservabilityPipelineEnrichmentTableFieldVrlLookup != nil {
		return obj.ObservabilityPipelineEnrichmentTableFieldVrlLookup
	}

	if obj.ObservabilityPipelineEnrichmentTableFieldSecretLookup != nil {
		return obj.ObservabilityPipelineEnrichmentTableFieldSecretLookup
	}

	// all schemas are nil
	return nil
}
