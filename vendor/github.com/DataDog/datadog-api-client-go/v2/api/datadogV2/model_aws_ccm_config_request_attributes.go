// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSCcmConfigRequestAttributes AWS CCM Config attributes for Create/Update requests.
type AWSCcmConfigRequestAttributes struct {
	// AWS Cloud Cost Management config.
	CcmConfig AWSCcmConfig `json:"ccm_config"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSCcmConfigRequestAttributes instantiates a new AWSCcmConfigRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSCcmConfigRequestAttributes(ccmConfig AWSCcmConfig) *AWSCcmConfigRequestAttributes {
	this := AWSCcmConfigRequestAttributes{}
	this.CcmConfig = ccmConfig
	return &this
}

// NewAWSCcmConfigRequestAttributesWithDefaults instantiates a new AWSCcmConfigRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSCcmConfigRequestAttributesWithDefaults() *AWSCcmConfigRequestAttributes {
	this := AWSCcmConfigRequestAttributes{}
	return &this
}

// GetCcmConfig returns the CcmConfig field value.
func (o *AWSCcmConfigRequestAttributes) GetCcmConfig() AWSCcmConfig {
	if o == nil {
		var ret AWSCcmConfig
		return ret
	}
	return o.CcmConfig
}

// GetCcmConfigOk returns a tuple with the CcmConfig field value
// and a boolean to check if the value has been set.
func (o *AWSCcmConfigRequestAttributes) GetCcmConfigOk() (*AWSCcmConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CcmConfig, true
}

// SetCcmConfig sets field value.
func (o *AWSCcmConfigRequestAttributes) SetCcmConfig(v AWSCcmConfig) {
	o.CcmConfig = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSCcmConfigRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ccm_config"] = o.CcmConfig

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSCcmConfigRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CcmConfig *AWSCcmConfig `json:"ccm_config"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CcmConfig == nil {
		return fmt.Errorf("required field ccm_config missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ccm_config"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CcmConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CcmConfig = *all.CcmConfig

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
