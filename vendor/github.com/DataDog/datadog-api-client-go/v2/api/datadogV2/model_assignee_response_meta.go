// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AssigneeResponseMeta Per-finding warnings and failures produced while processing the bulk assignee request.
type AssigneeResponseMeta struct {
	// Findings that could not be assigned or unassigned.
	Failures []AssignmentResult `json:"failures,omitempty"`
	// Findings for which the assignment succeeded but a non-critical error occurred during processing.
	Warnings []AssignmentResult `json:"warnings,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAssigneeResponseMeta instantiates a new AssigneeResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssigneeResponseMeta() *AssigneeResponseMeta {
	this := AssigneeResponseMeta{}
	return &this
}

// NewAssigneeResponseMetaWithDefaults instantiates a new AssigneeResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssigneeResponseMetaWithDefaults() *AssigneeResponseMeta {
	this := AssigneeResponseMeta{}
	return &this
}

// GetFailures returns the Failures field value if set, zero value otherwise.
func (o *AssigneeResponseMeta) GetFailures() []AssignmentResult {
	if o == nil || o.Failures == nil {
		var ret []AssignmentResult
		return ret
	}
	return o.Failures
}

// GetFailuresOk returns a tuple with the Failures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssigneeResponseMeta) GetFailuresOk() (*[]AssignmentResult, bool) {
	if o == nil || o.Failures == nil {
		return nil, false
	}
	return &o.Failures, true
}

// HasFailures returns a boolean if a field has been set.
func (o *AssigneeResponseMeta) HasFailures() bool {
	return o != nil && o.Failures != nil
}

// SetFailures gets a reference to the given []AssignmentResult and assigns it to the Failures field.
func (o *AssigneeResponseMeta) SetFailures(v []AssignmentResult) {
	o.Failures = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AssigneeResponseMeta) GetWarnings() []AssignmentResult {
	if o == nil || o.Warnings == nil {
		var ret []AssignmentResult
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssigneeResponseMeta) GetWarningsOk() (*[]AssignmentResult, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return &o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AssigneeResponseMeta) HasWarnings() bool {
	return o != nil && o.Warnings != nil
}

// SetWarnings gets a reference to the given []AssignmentResult and assigns it to the Warnings field.
func (o *AssigneeResponseMeta) SetWarnings(v []AssignmentResult) {
	o.Warnings = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssigneeResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Failures != nil {
		toSerialize["failures"] = o.Failures
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AssigneeResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Failures []AssignmentResult `json:"failures,omitempty"`
		Warnings []AssignmentResult `json:"warnings,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"failures", "warnings"})
	} else {
		return err
	}
	o.Failures = all.Failures
	o.Warnings = all.Warnings

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
