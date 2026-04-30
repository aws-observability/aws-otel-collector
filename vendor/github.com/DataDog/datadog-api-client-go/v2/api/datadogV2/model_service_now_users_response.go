// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowUsersResponse Response containing ServiceNow users
type ServiceNowUsersResponse struct {
	// Array of ServiceNow user data objects
	Data []ServiceNowUserData `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceNowUsersResponse instantiates a new ServiceNowUsersResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceNowUsersResponse(data []ServiceNowUserData) *ServiceNowUsersResponse {
	this := ServiceNowUsersResponse{}
	this.Data = data
	return &this
}

// NewServiceNowUsersResponseWithDefaults instantiates a new ServiceNowUsersResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceNowUsersResponseWithDefaults() *ServiceNowUsersResponse {
	this := ServiceNowUsersResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ServiceNowUsersResponse) GetData() []ServiceNowUserData {
	if o == nil {
		var ret []ServiceNowUserData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ServiceNowUsersResponse) GetDataOk() (*[]ServiceNowUserData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ServiceNowUsersResponse) SetData(v []ServiceNowUserData) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceNowUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceNowUsersResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]ServiceNowUserData `json:"data"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	o.Data = *all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
