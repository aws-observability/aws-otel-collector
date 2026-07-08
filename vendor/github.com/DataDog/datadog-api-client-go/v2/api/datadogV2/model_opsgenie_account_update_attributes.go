// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpsgenieAccountUpdateAttributes The Opsgenie account attributes for an update request.
type OpsgenieAccountUpdateAttributes struct {
	// The Opsgenie API key for your Opsgenie account.
	ApiKey *string `json:"api_key,omitempty"`
	// The region for the Opsgenie service.
	Region *OpsgenieServiceRegionType `json:"region,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsgenieAccountUpdateAttributes instantiates a new OpsgenieAccountUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsgenieAccountUpdateAttributes() *OpsgenieAccountUpdateAttributes {
	this := OpsgenieAccountUpdateAttributes{}
	return &this
}

// NewOpsgenieAccountUpdateAttributesWithDefaults instantiates a new OpsgenieAccountUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsgenieAccountUpdateAttributesWithDefaults() *OpsgenieAccountUpdateAttributes {
	this := OpsgenieAccountUpdateAttributes{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *OpsgenieAccountUpdateAttributes) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsgenieAccountUpdateAttributes) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *OpsgenieAccountUpdateAttributes) HasApiKey() bool {
	return o != nil && o.ApiKey != nil
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *OpsgenieAccountUpdateAttributes) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *OpsgenieAccountUpdateAttributes) GetRegion() OpsgenieServiceRegionType {
	if o == nil || o.Region == nil {
		var ret OpsgenieServiceRegionType
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsgenieAccountUpdateAttributes) GetRegionOk() (*OpsgenieServiceRegionType, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *OpsgenieAccountUpdateAttributes) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given OpsgenieServiceRegionType and assigns it to the Region field.
func (o *OpsgenieAccountUpdateAttributes) SetRegion(v OpsgenieServiceRegionType) {
	o.Region = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsgenieAccountUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApiKey != nil {
		toSerialize["api_key"] = o.ApiKey
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsgenieAccountUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey *string                    `json:"api_key,omitempty"`
		Region *OpsgenieServiceRegionType `json:"region,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key", "region"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApiKey = all.ApiKey
	if all.Region != nil && !all.Region.IsValid() {
		hasInvalidField = true
	} else {
		o.Region = all.Region
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
