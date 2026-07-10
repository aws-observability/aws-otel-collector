// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringDatasetVersionFieldChange A single field change between two versions of a dataset.
type SecurityMonitoringDatasetVersionFieldChange struct {
	// The current value of the field, serialized as a JSON value.
	Current interface{} `json:"current"`
	// The name of the field that changed.
	Field string `json:"field"`
	// The previous value of the field, serialized as a JSON value.
	Previous interface{} `json:"previous"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringDatasetVersionFieldChange instantiates a new SecurityMonitoringDatasetVersionFieldChange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringDatasetVersionFieldChange(current interface{}, field string, previous interface{}) *SecurityMonitoringDatasetVersionFieldChange {
	this := SecurityMonitoringDatasetVersionFieldChange{}
	this.Current = current
	this.Field = field
	this.Previous = previous
	return &this
}

// NewSecurityMonitoringDatasetVersionFieldChangeWithDefaults instantiates a new SecurityMonitoringDatasetVersionFieldChange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringDatasetVersionFieldChangeWithDefaults() *SecurityMonitoringDatasetVersionFieldChange {
	this := SecurityMonitoringDatasetVersionFieldChange{}
	return &this
}

// GetCurrent returns the Current field value.
func (o *SecurityMonitoringDatasetVersionFieldChange) GetCurrent() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Current
}

// GetCurrentOk returns a tuple with the Current field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetVersionFieldChange) GetCurrentOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Current, true
}

// SetCurrent sets field value.
func (o *SecurityMonitoringDatasetVersionFieldChange) SetCurrent(v interface{}) {
	o.Current = v
}

// GetField returns the Field field value.
func (o *SecurityMonitoringDatasetVersionFieldChange) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetVersionFieldChange) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *SecurityMonitoringDatasetVersionFieldChange) SetField(v string) {
	o.Field = v
}

// GetPrevious returns the Previous field value.
func (o *SecurityMonitoringDatasetVersionFieldChange) GetPrevious() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringDatasetVersionFieldChange) GetPreviousOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Previous, true
}

// SetPrevious sets field value.
func (o *SecurityMonitoringDatasetVersionFieldChange) SetPrevious(v interface{}) {
	o.Previous = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringDatasetVersionFieldChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["current"] = o.Current
	toSerialize["field"] = o.Field
	toSerialize["previous"] = o.Previous

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringDatasetVersionFieldChange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Current  *interface{} `json:"current"`
		Field    *string      `json:"field"`
		Previous *interface{} `json:"previous"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Current == nil {
		return fmt.Errorf("required field current missing")
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	if all.Previous == nil {
		return fmt.Errorf("required field previous missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"current", "field", "previous"})
	} else {
		return err
	}
	o.Current = *all.Current
	o.Field = *all.Field
	o.Previous = *all.Previous

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
