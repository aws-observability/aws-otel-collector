// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GlobalVariableJsonPatchRequestDataAttributes Attributes for a JSON Patch request on a Synthetic global variable.
type GlobalVariableJsonPatchRequestDataAttributes struct {
	// JSON Patch operations following RFC 6902.
	JsonPatch []JsonPatchOperation `json:"json_patch,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGlobalVariableJsonPatchRequestDataAttributes instantiates a new GlobalVariableJsonPatchRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGlobalVariableJsonPatchRequestDataAttributes() *GlobalVariableJsonPatchRequestDataAttributes {
	this := GlobalVariableJsonPatchRequestDataAttributes{}
	return &this
}

// NewGlobalVariableJsonPatchRequestDataAttributesWithDefaults instantiates a new GlobalVariableJsonPatchRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGlobalVariableJsonPatchRequestDataAttributesWithDefaults() *GlobalVariableJsonPatchRequestDataAttributes {
	this := GlobalVariableJsonPatchRequestDataAttributes{}
	return &this
}

// GetJsonPatch returns the JsonPatch field value if set, zero value otherwise.
func (o *GlobalVariableJsonPatchRequestDataAttributes) GetJsonPatch() []JsonPatchOperation {
	if o == nil || o.JsonPatch == nil {
		var ret []JsonPatchOperation
		return ret
	}
	return o.JsonPatch
}

// GetJsonPatchOk returns a tuple with the JsonPatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalVariableJsonPatchRequestDataAttributes) GetJsonPatchOk() (*[]JsonPatchOperation, bool) {
	if o == nil || o.JsonPatch == nil {
		return nil, false
	}
	return &o.JsonPatch, true
}

// HasJsonPatch returns a boolean if a field has been set.
func (o *GlobalVariableJsonPatchRequestDataAttributes) HasJsonPatch() bool {
	return o != nil && o.JsonPatch != nil
}

// SetJsonPatch gets a reference to the given []JsonPatchOperation and assigns it to the JsonPatch field.
func (o *GlobalVariableJsonPatchRequestDataAttributes) SetJsonPatch(v []JsonPatchOperation) {
	o.JsonPatch = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GlobalVariableJsonPatchRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.JsonPatch != nil {
		toSerialize["json_patch"] = o.JsonPatch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GlobalVariableJsonPatchRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		JsonPatch []JsonPatchOperation `json:"json_patch,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"json_patch"})
	} else {
		return err
	}
	o.JsonPatch = all.JsonPatch

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
