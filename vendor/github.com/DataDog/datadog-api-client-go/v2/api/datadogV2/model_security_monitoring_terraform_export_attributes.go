// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringTerraformExportAttributes Attributes of the Terraform export response.
type SecurityMonitoringTerraformExportAttributes struct {
	// The Terraform configuration for the resource.
	Output *string `json:"output,omitempty"`
	// The ID of the exported resource.
	ResourceId string `json:"resource_id"`
	// The Terraform resource type name.
	TypeName string `json:"type_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringTerraformExportAttributes instantiates a new SecurityMonitoringTerraformExportAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringTerraformExportAttributes(resourceId string, typeName string) *SecurityMonitoringTerraformExportAttributes {
	this := SecurityMonitoringTerraformExportAttributes{}
	this.ResourceId = resourceId
	this.TypeName = typeName
	return &this
}

// NewSecurityMonitoringTerraformExportAttributesWithDefaults instantiates a new SecurityMonitoringTerraformExportAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringTerraformExportAttributesWithDefaults() *SecurityMonitoringTerraformExportAttributes {
	this := SecurityMonitoringTerraformExportAttributes{}
	return &this
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *SecurityMonitoringTerraformExportAttributes) GetOutput() string {
	if o == nil || o.Output == nil {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringTerraformExportAttributes) GetOutputOk() (*string, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *SecurityMonitoringTerraformExportAttributes) HasOutput() bool {
	return o != nil && o.Output != nil
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *SecurityMonitoringTerraformExportAttributes) SetOutput(v string) {
	o.Output = &v
}

// GetResourceId returns the ResourceId field value.
func (o *SecurityMonitoringTerraformExportAttributes) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringTerraformExportAttributes) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value.
func (o *SecurityMonitoringTerraformExportAttributes) SetResourceId(v string) {
	o.ResourceId = v
}

// GetTypeName returns the TypeName field value.
func (o *SecurityMonitoringTerraformExportAttributes) GetTypeName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringTerraformExportAttributes) GetTypeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeName, true
}

// SetTypeName sets field value.
func (o *SecurityMonitoringTerraformExportAttributes) SetTypeName(v string) {
	o.TypeName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringTerraformExportAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["type_name"] = o.TypeName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringTerraformExportAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Output     *string `json:"output,omitempty"`
		ResourceId *string `json:"resource_id"`
		TypeName   *string `json:"type_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ResourceId == nil {
		return fmt.Errorf("required field resource_id missing")
	}
	if all.TypeName == nil {
		return fmt.Errorf("required field type_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"output", "resource_id", "type_name"})
	} else {
		return err
	}
	o.Output = all.Output
	o.ResourceId = *all.ResourceId
	o.TypeName = *all.TypeName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
