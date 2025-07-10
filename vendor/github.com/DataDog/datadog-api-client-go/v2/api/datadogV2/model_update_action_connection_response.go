// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateActionConnectionResponse The response for an updated connection.
type UpdateActionConnectionResponse struct {
	// Data related to the connection.
	Data *ActionConnectionData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateActionConnectionResponse instantiates a new UpdateActionConnectionResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateActionConnectionResponse() *UpdateActionConnectionResponse {
	this := UpdateActionConnectionResponse{}
	return &this
}

// NewUpdateActionConnectionResponseWithDefaults instantiates a new UpdateActionConnectionResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateActionConnectionResponseWithDefaults() *UpdateActionConnectionResponse {
	this := UpdateActionConnectionResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UpdateActionConnectionResponse) GetData() ActionConnectionData {
	if o == nil || o.Data == nil {
		var ret ActionConnectionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateActionConnectionResponse) GetDataOk() (*ActionConnectionData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UpdateActionConnectionResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given ActionConnectionData and assigns it to the Data field.
func (o *UpdateActionConnectionResponse) SetData(v ActionConnectionData) {
	o.Data = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateActionConnectionResponse) MarshalJSON() ([]byte, error) {
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
func (o *UpdateActionConnectionResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *ActionConnectionData `json:"data,omitempty"`
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
