// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ItemApiPayloadArray A collection of datastore items with pagination and schema metadata.
type ItemApiPayloadArray struct {
	// An array of datastore items with their content and metadata.
	Data []ItemApiPayloadData `json:"data"`
	// Additional metadata about a collection of datastore items, including pagination and schema information.
	Meta *ItemApiPayloadMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewItemApiPayloadArray instantiates a new ItemApiPayloadArray object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewItemApiPayloadArray(data []ItemApiPayloadData) *ItemApiPayloadArray {
	this := ItemApiPayloadArray{}
	this.Data = data
	return &this
}

// NewItemApiPayloadArrayWithDefaults instantiates a new ItemApiPayloadArray object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewItemApiPayloadArrayWithDefaults() *ItemApiPayloadArray {
	this := ItemApiPayloadArray{}
	return &this
}

// GetData returns the Data field value.
func (o *ItemApiPayloadArray) GetData() []ItemApiPayloadData {
	if o == nil {
		var ret []ItemApiPayloadData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadArray) GetDataOk() (*[]ItemApiPayloadData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ItemApiPayloadArray) SetData(v []ItemApiPayloadData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ItemApiPayloadArray) GetMeta() ItemApiPayloadMeta {
	if o == nil || o.Meta == nil {
		var ret ItemApiPayloadMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadArray) GetMetaOk() (*ItemApiPayloadMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ItemApiPayloadArray) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given ItemApiPayloadMeta and assigns it to the Meta field.
func (o *ItemApiPayloadArray) SetMeta(v ItemApiPayloadMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ItemApiPayloadArray) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ItemApiPayloadArray) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]ItemApiPayloadData `json:"data"`
		Meta *ItemApiPayloadMeta   `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
