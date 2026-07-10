// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpsgenieAccountCreateAttributes The Opsgenie account attributes for a create request.
type OpsgenieAccountCreateAttributes struct {
	// The Opsgenie API key for your Opsgenie account.
	ApiKey string `json:"api_key"`
	// The region for the Opsgenie service.
	Region OpsgenieServiceRegionType `json:"region"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsgenieAccountCreateAttributes instantiates a new OpsgenieAccountCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsgenieAccountCreateAttributes(apiKey string, region OpsgenieServiceRegionType) *OpsgenieAccountCreateAttributes {
	this := OpsgenieAccountCreateAttributes{}
	this.ApiKey = apiKey
	this.Region = region
	return &this
}

// NewOpsgenieAccountCreateAttributesWithDefaults instantiates a new OpsgenieAccountCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsgenieAccountCreateAttributesWithDefaults() *OpsgenieAccountCreateAttributes {
	this := OpsgenieAccountCreateAttributes{}
	return &this
}

// GetApiKey returns the ApiKey field value.
func (o *OpsgenieAccountCreateAttributes) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *OpsgenieAccountCreateAttributes) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value.
func (o *OpsgenieAccountCreateAttributes) SetApiKey(v string) {
	o.ApiKey = v
}

// GetRegion returns the Region field value.
func (o *OpsgenieAccountCreateAttributes) GetRegion() OpsgenieServiceRegionType {
	if o == nil {
		var ret OpsgenieServiceRegionType
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *OpsgenieAccountCreateAttributes) GetRegionOk() (*OpsgenieServiceRegionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *OpsgenieAccountCreateAttributes) SetRegion(v OpsgenieServiceRegionType) {
	o.Region = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsgenieAccountCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["api_key"] = o.ApiKey
	toSerialize["region"] = o.Region

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsgenieAccountCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey *string                    `json:"api_key"`
		Region *OpsgenieServiceRegionType `json:"region"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiKey == nil {
		return fmt.Errorf("required field api_key missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key", "region"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApiKey = *all.ApiKey
	if !all.Region.IsValid() {
		hasInvalidField = true
	} else {
		o.Region = *all.Region
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
