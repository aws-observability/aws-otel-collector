// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryUsersRequestDataAttributesSort
type QueryUsersRequestDataAttributesSort struct {
	//
	Field *string `json:"field,omitempty"`
	//
	Order *string `json:"order,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewQueryUsersRequestDataAttributesSort instantiates a new QueryUsersRequestDataAttributesSort object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryUsersRequestDataAttributesSort() *QueryUsersRequestDataAttributesSort {
	this := QueryUsersRequestDataAttributesSort{}
	return &this
}

// NewQueryUsersRequestDataAttributesSortWithDefaults instantiates a new QueryUsersRequestDataAttributesSort object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryUsersRequestDataAttributesSortWithDefaults() *QueryUsersRequestDataAttributesSort {
	this := QueryUsersRequestDataAttributesSort{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *QueryUsersRequestDataAttributesSort) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUsersRequestDataAttributesSort) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *QueryUsersRequestDataAttributesSort) HasField() bool {
	return o != nil && o.Field != nil
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *QueryUsersRequestDataAttributesSort) SetField(v string) {
	o.Field = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *QueryUsersRequestDataAttributesSort) GetOrder() string {
	if o == nil || o.Order == nil {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryUsersRequestDataAttributesSort) GetOrderOk() (*string, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *QueryUsersRequestDataAttributesSort) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *QueryUsersRequestDataAttributesSort) SetOrder(v string) {
	o.Order = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryUsersRequestDataAttributesSort) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryUsersRequestDataAttributesSort) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Field *string `json:"field,omitempty"`
		Order *string `json:"order,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"field", "order"})
	} else {
		return err
	}
	o.Field = all.Field
	o.Order = all.Order

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
