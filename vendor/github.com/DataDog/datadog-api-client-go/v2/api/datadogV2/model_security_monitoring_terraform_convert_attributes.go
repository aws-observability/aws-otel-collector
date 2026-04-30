// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringTerraformConvertAttributes Attributes for the convert request.
type SecurityMonitoringTerraformConvertAttributes struct {
	// The resource attributes as a JSON object, matching the structure returned by the corresponding Datadog API (for example, the attributes of a suppression rule).
	ResourceJson map[string]interface{} `json:"resource_json"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringTerraformConvertAttributes instantiates a new SecurityMonitoringTerraformConvertAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringTerraformConvertAttributes(resourceJson map[string]interface{}) *SecurityMonitoringTerraformConvertAttributes {
	this := SecurityMonitoringTerraformConvertAttributes{}
	this.ResourceJson = resourceJson
	return &this
}

// NewSecurityMonitoringTerraformConvertAttributesWithDefaults instantiates a new SecurityMonitoringTerraformConvertAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringTerraformConvertAttributesWithDefaults() *SecurityMonitoringTerraformConvertAttributes {
	this := SecurityMonitoringTerraformConvertAttributes{}
	return &this
}

// GetResourceJson returns the ResourceJson field value.
func (o *SecurityMonitoringTerraformConvertAttributes) GetResourceJson() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.ResourceJson
}

// GetResourceJsonOk returns a tuple with the ResourceJson field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringTerraformConvertAttributes) GetResourceJsonOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceJson, true
}

// SetResourceJson sets field value.
func (o *SecurityMonitoringTerraformConvertAttributes) SetResourceJson(v map[string]interface{}) {
	o.ResourceJson = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringTerraformConvertAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["resource_json"] = o.ResourceJson

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringTerraformConvertAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ResourceJson *map[string]interface{} `json:"resource_json"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ResourceJson == nil {
		return fmt.Errorf("required field resource_json missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"resource_json"})
	} else {
		return err
	}
	o.ResourceJson = *all.ResourceJson

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
