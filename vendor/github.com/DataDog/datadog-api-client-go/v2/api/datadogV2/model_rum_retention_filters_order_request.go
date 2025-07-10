// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumRetentionFiltersOrderRequest The list of RUM retention filter IDs along with their corresponding type to reorder.
// All retention filter IDs should be included in the list created for a RUM application.
type RumRetentionFiltersOrderRequest struct {
	// A list of RUM retention filter IDs along with type.
	Data []RumRetentionFiltersOrderData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumRetentionFiltersOrderRequest instantiates a new RumRetentionFiltersOrderRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumRetentionFiltersOrderRequest() *RumRetentionFiltersOrderRequest {
	this := RumRetentionFiltersOrderRequest{}
	return &this
}

// NewRumRetentionFiltersOrderRequestWithDefaults instantiates a new RumRetentionFiltersOrderRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumRetentionFiltersOrderRequestWithDefaults() *RumRetentionFiltersOrderRequest {
	this := RumRetentionFiltersOrderRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RumRetentionFiltersOrderRequest) GetData() []RumRetentionFiltersOrderData {
	if o == nil || o.Data == nil {
		var ret []RumRetentionFiltersOrderData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumRetentionFiltersOrderRequest) GetDataOk() (*[]RumRetentionFiltersOrderData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RumRetentionFiltersOrderRequest) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []RumRetentionFiltersOrderData and assigns it to the Data field.
func (o *RumRetentionFiltersOrderRequest) SetData(v []RumRetentionFiltersOrderData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumRetentionFiltersOrderRequest) MarshalJSON() ([]byte, error) {
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
func (o *RumRetentionFiltersOrderRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []RumRetentionFiltersOrderData `json:"data,omitempty"`
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
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
