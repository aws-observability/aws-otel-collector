// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureUCConfigPatchRequestAttributes Attributes for Azure config Patch Request.
type AzureUCConfigPatchRequestAttributes struct {
	// Whether or not the Cloud Cost Management account is enabled.
	IsEnabled bool `json:"is_enabled"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAzureUCConfigPatchRequestAttributes instantiates a new AzureUCConfigPatchRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAzureUCConfigPatchRequestAttributes(isEnabled bool) *AzureUCConfigPatchRequestAttributes {
	this := AzureUCConfigPatchRequestAttributes{}
	this.IsEnabled = isEnabled
	return &this
}

// NewAzureUCConfigPatchRequestAttributesWithDefaults instantiates a new AzureUCConfigPatchRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAzureUCConfigPatchRequestAttributesWithDefaults() *AzureUCConfigPatchRequestAttributes {
	this := AzureUCConfigPatchRequestAttributes{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value.
func (o *AzureUCConfigPatchRequestAttributes) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *AzureUCConfigPatchRequestAttributes) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value.
func (o *AzureUCConfigPatchRequestAttributes) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AzureUCConfigPatchRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["is_enabled"] = o.IsEnabled

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AzureUCConfigPatchRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IsEnabled *bool `json:"is_enabled"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsEnabled == nil {
		return fmt.Errorf("required field is_enabled missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"is_enabled"})
	} else {
		return err
	}
	o.IsEnabled = *all.IsEnabled

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
