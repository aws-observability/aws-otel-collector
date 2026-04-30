// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetIoCIndicatorResponseAttributes Attributes of the get indicator response.
type GetIoCIndicatorResponseAttributes struct {
	// An indicator of compromise with extended context from your environment.
	Data *IoCIndicatorDetailed `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetIoCIndicatorResponseAttributes instantiates a new GetIoCIndicatorResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetIoCIndicatorResponseAttributes() *GetIoCIndicatorResponseAttributes {
	this := GetIoCIndicatorResponseAttributes{}
	return &this
}

// NewGetIoCIndicatorResponseAttributesWithDefaults instantiates a new GetIoCIndicatorResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetIoCIndicatorResponseAttributesWithDefaults() *GetIoCIndicatorResponseAttributes {
	this := GetIoCIndicatorResponseAttributes{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetIoCIndicatorResponseAttributes) GetData() IoCIndicatorDetailed {
	if o == nil || o.Data == nil {
		var ret IoCIndicatorDetailed
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIoCIndicatorResponseAttributes) GetDataOk() (*IoCIndicatorDetailed, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetIoCIndicatorResponseAttributes) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given IoCIndicatorDetailed and assigns it to the Data field.
func (o *GetIoCIndicatorResponseAttributes) SetData(v IoCIndicatorDetailed) {
	o.Data = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetIoCIndicatorResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetIoCIndicatorResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *IoCIndicatorDetailed `json:"data,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
