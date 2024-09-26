// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsPatchTestOperation A single [JSON Patch](https://jsonpatch.com) operation to perform on the test
type SyntheticsPatchTestOperation struct {
	// The operation to perform
	Op *SyntheticsPatchTestOperationName `json:"op,omitempty"`
	// The path to the value to modify
	Path *string `json:"path,omitempty"`
	// A value to use in a [JSON Patch](https://jsonpatch.com) operation
	Value interface{} `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsPatchTestOperation instantiates a new SyntheticsPatchTestOperation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsPatchTestOperation() *SyntheticsPatchTestOperation {
	this := SyntheticsPatchTestOperation{}
	return &this
}

// NewSyntheticsPatchTestOperationWithDefaults instantiates a new SyntheticsPatchTestOperation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsPatchTestOperationWithDefaults() *SyntheticsPatchTestOperation {
	this := SyntheticsPatchTestOperation{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *SyntheticsPatchTestOperation) GetOp() SyntheticsPatchTestOperationName {
	if o == nil || o.Op == nil {
		var ret SyntheticsPatchTestOperationName
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsPatchTestOperation) GetOpOk() (*SyntheticsPatchTestOperationName, bool) {
	if o == nil || o.Op == nil {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *SyntheticsPatchTestOperation) HasOp() bool {
	return o != nil && o.Op != nil
}

// SetOp gets a reference to the given SyntheticsPatchTestOperationName and assigns it to the Op field.
func (o *SyntheticsPatchTestOperation) SetOp(v SyntheticsPatchTestOperationName) {
	o.Op = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SyntheticsPatchTestOperation) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsPatchTestOperation) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SyntheticsPatchTestOperation) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SyntheticsPatchTestOperation) SetPath(v string) {
	o.Path = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SyntheticsPatchTestOperation) GetValue() interface{} {
	if o == nil || o.Value == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsPatchTestOperation) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SyntheticsPatchTestOperation) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *SyntheticsPatchTestOperation) SetValue(v interface{}) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsPatchTestOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Op != nil {
		toSerialize["op"] = o.Op
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsPatchTestOperation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Op    *SyntheticsPatchTestOperationName `json:"op,omitempty"`
		Path  *string                           `json:"path,omitempty"`
		Value interface{}                       `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"op", "path", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Op != nil && !all.Op.IsValid() {
		hasInvalidField = true
	} else {
		o.Op = all.Op
	}
	o.Path = all.Path
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
