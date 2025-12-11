// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureUCConfigPairAttributes Attributes for Azure config pair.
type AzureUCConfigPairAttributes struct {
	// An Azure config.
	Configs []AzureUCConfig `json:"configs"`
	// The ID of the Azure config pair.
	Id *string `json:"id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAzureUCConfigPairAttributes instantiates a new AzureUCConfigPairAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAzureUCConfigPairAttributes(configs []AzureUCConfig) *AzureUCConfigPairAttributes {
	this := AzureUCConfigPairAttributes{}
	this.Configs = configs
	return &this
}

// NewAzureUCConfigPairAttributesWithDefaults instantiates a new AzureUCConfigPairAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAzureUCConfigPairAttributesWithDefaults() *AzureUCConfigPairAttributes {
	this := AzureUCConfigPairAttributes{}
	return &this
}

// GetConfigs returns the Configs field value.
func (o *AzureUCConfigPairAttributes) GetConfigs() []AzureUCConfig {
	if o == nil {
		var ret []AzureUCConfig
		return ret
	}
	return o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfigPairAttributes) GetConfigsOk() (*[]AzureUCConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configs, true
}

// SetConfigs sets field value.
func (o *AzureUCConfigPairAttributes) SetConfigs(v []AzureUCConfig) {
	o.Configs = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AzureUCConfigPairAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureUCConfigPairAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AzureUCConfigPairAttributes) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AzureUCConfigPairAttributes) SetId(v string) {
	o.Id = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AzureUCConfigPairAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["configs"] = o.Configs
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AzureUCConfigPairAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Configs *[]AzureUCConfig `json:"configs"`
		Id      *string          `json:"id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Configs == nil {
		return fmt.Errorf("required field configs missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configs", "id"})
	} else {
		return err
	}
	o.Configs = *all.Configs
	o.Id = all.Id

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
