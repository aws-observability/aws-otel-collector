// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BatchUpsertRowsRequestDataAttributes Attributes containing row data values for row creation or update operations.
type BatchUpsertRowsRequestDataAttributes struct {
	// Key-value pairs representing row data, where keys are field names from the schema.
	Values map[string]interface{} `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBatchUpsertRowsRequestDataAttributes instantiates a new BatchUpsertRowsRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBatchUpsertRowsRequestDataAttributes(values map[string]interface{}) *BatchUpsertRowsRequestDataAttributes {
	this := BatchUpsertRowsRequestDataAttributes{}
	this.Values = values
	return &this
}

// NewBatchUpsertRowsRequestDataAttributesWithDefaults instantiates a new BatchUpsertRowsRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBatchUpsertRowsRequestDataAttributesWithDefaults() *BatchUpsertRowsRequestDataAttributes {
	this := BatchUpsertRowsRequestDataAttributes{}
	return &this
}

// GetValues returns the Values field value.
func (o *BatchUpsertRowsRequestDataAttributes) GetValues() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *BatchUpsertRowsRequestDataAttributes) GetValuesOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *BatchUpsertRowsRequestDataAttributes) SetValues(v map[string]interface{}) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BatchUpsertRowsRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BatchUpsertRowsRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Values *map[string]interface{} `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"values"})
	} else {
		return err
	}
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
