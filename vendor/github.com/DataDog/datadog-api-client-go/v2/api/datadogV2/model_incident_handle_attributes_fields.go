// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentHandleAttributesFields Dynamic fields associated with the handle
type IncidentHandleAttributesFields struct {
	// Severity levels associated with the handle
	Severity []string `json:"severity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentHandleAttributesFields instantiates a new IncidentHandleAttributesFields object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentHandleAttributesFields() *IncidentHandleAttributesFields {
	this := IncidentHandleAttributesFields{}
	return &this
}

// NewIncidentHandleAttributesFieldsWithDefaults instantiates a new IncidentHandleAttributesFields object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentHandleAttributesFieldsWithDefaults() *IncidentHandleAttributesFields {
	this := IncidentHandleAttributesFields{}
	return &this
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *IncidentHandleAttributesFields) GetSeverity() []string {
	if o == nil || o.Severity == nil {
		var ret []string
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentHandleAttributesFields) GetSeverityOk() (*[]string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return &o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *IncidentHandleAttributesFields) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given []string and assigns it to the Severity field.
func (o *IncidentHandleAttributesFields) SetSeverity(v []string) {
	o.Severity = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentHandleAttributesFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentHandleAttributesFields) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Severity []string `json:"severity,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"severity"})
	} else {
		return err
	}
	o.Severity = all.Severity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
