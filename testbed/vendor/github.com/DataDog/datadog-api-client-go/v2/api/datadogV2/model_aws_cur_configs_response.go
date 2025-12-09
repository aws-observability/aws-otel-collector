// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AwsCURConfigsResponse List of AWS CUR configs.
type AwsCURConfigsResponse struct {
	// An AWS CUR config.
	Data []AwsCURConfig `json:"data"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAwsCURConfigsResponse instantiates a new AwsCURConfigsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAwsCURConfigsResponse(data []AwsCURConfig) *AwsCURConfigsResponse {
	this := AwsCURConfigsResponse{}
	this.Data = data
	return &this
}

// NewAwsCURConfigsResponseWithDefaults instantiates a new AwsCURConfigsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAwsCURConfigsResponseWithDefaults() *AwsCURConfigsResponse {
	this := AwsCURConfigsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *AwsCURConfigsResponse) GetData() []AwsCURConfig {
	if o == nil {
		var ret []AwsCURConfig
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AwsCURConfigsResponse) GetDataOk() (*[]AwsCURConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *AwsCURConfigsResponse) SetData(v []AwsCURConfig) {
	o.Data = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AwsCURConfigsResponse) MarshalJSON() ([]byte, error) {
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
func (o *AwsCURConfigsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]AwsCURConfig `json:"data"`
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
