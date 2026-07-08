// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetRecordTagOperations Explicit tag operations for updating records. Operations are applied in order, Remove then Add then Set. `set` is the final override; if specified, the result of `remove` and `add` is discarded.
type LLMObsDatasetRecordTagOperations struct {
	// List of tag strings.
	Add []string `json:"add,omitempty"`
	// List of tag strings.
	Remove []string `json:"remove,omitempty"`
	// List of tag strings.
	Set []string `json:"set,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetRecordTagOperations instantiates a new LLMObsDatasetRecordTagOperations object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetRecordTagOperations() *LLMObsDatasetRecordTagOperations {
	this := LLMObsDatasetRecordTagOperations{}
	return &this
}

// NewLLMObsDatasetRecordTagOperationsWithDefaults instantiates a new LLMObsDatasetRecordTagOperations object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetRecordTagOperationsWithDefaults() *LLMObsDatasetRecordTagOperations {
	this := LLMObsDatasetRecordTagOperations{}
	return &this
}

// GetAdd returns the Add field value if set, zero value otherwise.
func (o *LLMObsDatasetRecordTagOperations) GetAdd() []string {
	if o == nil || o.Add == nil {
		var ret []string
		return ret
	}
	return o.Add
}

// GetAddOk returns a tuple with the Add field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordTagOperations) GetAddOk() (*[]string, bool) {
	if o == nil || o.Add == nil {
		return nil, false
	}
	return &o.Add, true
}

// HasAdd returns a boolean if a field has been set.
func (o *LLMObsDatasetRecordTagOperations) HasAdd() bool {
	return o != nil && o.Add != nil
}

// SetAdd gets a reference to the given []string and assigns it to the Add field.
func (o *LLMObsDatasetRecordTagOperations) SetAdd(v []string) {
	o.Add = v
}

// GetRemove returns the Remove field value if set, zero value otherwise.
func (o *LLMObsDatasetRecordTagOperations) GetRemove() []string {
	if o == nil || o.Remove == nil {
		var ret []string
		return ret
	}
	return o.Remove
}

// GetRemoveOk returns a tuple with the Remove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordTagOperations) GetRemoveOk() (*[]string, bool) {
	if o == nil || o.Remove == nil {
		return nil, false
	}
	return &o.Remove, true
}

// HasRemove returns a boolean if a field has been set.
func (o *LLMObsDatasetRecordTagOperations) HasRemove() bool {
	return o != nil && o.Remove != nil
}

// SetRemove gets a reference to the given []string and assigns it to the Remove field.
func (o *LLMObsDatasetRecordTagOperations) SetRemove(v []string) {
	o.Remove = v
}

// GetSet returns the Set field value if set, zero value otherwise.
func (o *LLMObsDatasetRecordTagOperations) GetSet() []string {
	if o == nil || o.Set == nil {
		var ret []string
		return ret
	}
	return o.Set
}

// GetSetOk returns a tuple with the Set field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetRecordTagOperations) GetSetOk() (*[]string, bool) {
	if o == nil || o.Set == nil {
		return nil, false
	}
	return &o.Set, true
}

// HasSet returns a boolean if a field has been set.
func (o *LLMObsDatasetRecordTagOperations) HasSet() bool {
	return o != nil && o.Set != nil
}

// SetSet gets a reference to the given []string and assigns it to the Set field.
func (o *LLMObsDatasetRecordTagOperations) SetSet(v []string) {
	o.Set = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetRecordTagOperations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Add != nil {
		toSerialize["add"] = o.Add
	}
	if o.Remove != nil {
		toSerialize["remove"] = o.Remove
	}
	if o.Set != nil {
		toSerialize["set"] = o.Set
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetRecordTagOperations) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Add    []string `json:"add,omitempty"`
		Remove []string `json:"remove,omitempty"`
		Set    []string `json:"set,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"add", "remove", "set"})
	} else {
		return err
	}
	o.Add = all.Add
	o.Remove = all.Remove
	o.Set = all.Set

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
