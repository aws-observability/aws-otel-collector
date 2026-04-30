// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JsonPatchOperation A JSON Patch operation as per RFC 6902.
type JsonPatchOperation struct {
	// The operation to perform.
	Op JsonPatchOperationOp `json:"op"`
	// A JSON Pointer path (e.g., "/name", "/value/secure").
	Path string `json:"path"`
	// The value to use for the operation (not applicable for "remove" and "test" operations).
	Value interface{} `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJsonPatchOperation instantiates a new JsonPatchOperation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJsonPatchOperation(op JsonPatchOperationOp, path string) *JsonPatchOperation {
	this := JsonPatchOperation{}
	this.Op = op
	this.Path = path
	return &this
}

// NewJsonPatchOperationWithDefaults instantiates a new JsonPatchOperation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJsonPatchOperationWithDefaults() *JsonPatchOperation {
	this := JsonPatchOperation{}
	return &this
}

// GetOp returns the Op field value.
func (o *JsonPatchOperation) GetOp() JsonPatchOperationOp {
	if o == nil {
		var ret JsonPatchOperationOp
		return ret
	}
	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *JsonPatchOperation) GetOpOk() (*JsonPatchOperationOp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value.
func (o *JsonPatchOperation) SetOp(v JsonPatchOperationOp) {
	o.Op = v
}

// GetPath returns the Path field value.
func (o *JsonPatchOperation) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *JsonPatchOperation) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value.
func (o *JsonPatchOperation) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *JsonPatchOperation) GetValue() interface{} {
	if o == nil || o.Value == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JsonPatchOperation) GetValueOk() (*interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *JsonPatchOperation) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *JsonPatchOperation) SetValue(v interface{}) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o JsonPatchOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JsonPatchOperation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Op    *JsonPatchOperationOp `json:"op"`
		Path  *string               `json:"path"`
		Value interface{}           `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Op == nil {
		return fmt.Errorf("required field op missing")
	}
	if all.Path == nil {
		return fmt.Errorf("required field path missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"op", "path", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Op.IsValid() {
		hasInvalidField = true
	} else {
		o.Op = *all.Op
	}
	o.Path = *all.Path
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
