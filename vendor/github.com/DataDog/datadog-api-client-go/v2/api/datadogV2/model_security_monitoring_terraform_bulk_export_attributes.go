// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringTerraformBulkExportAttributes Attributes for the bulk export request.
type SecurityMonitoringTerraformBulkExportAttributes struct {
	// The list of resource IDs to export. Maximum 1000 items.
	ResourceIds []string `json:"resource_ids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringTerraformBulkExportAttributes instantiates a new SecurityMonitoringTerraformBulkExportAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringTerraformBulkExportAttributes(resourceIds []string) *SecurityMonitoringTerraformBulkExportAttributes {
	this := SecurityMonitoringTerraformBulkExportAttributes{}
	this.ResourceIds = resourceIds
	return &this
}

// NewSecurityMonitoringTerraformBulkExportAttributesWithDefaults instantiates a new SecurityMonitoringTerraformBulkExportAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringTerraformBulkExportAttributesWithDefaults() *SecurityMonitoringTerraformBulkExportAttributes {
	this := SecurityMonitoringTerraformBulkExportAttributes{}
	return &this
}

// GetResourceIds returns the ResourceIds field value.
func (o *SecurityMonitoringTerraformBulkExportAttributes) GetResourceIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ResourceIds
}

// GetResourceIdsOk returns a tuple with the ResourceIds field value
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringTerraformBulkExportAttributes) GetResourceIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceIds, true
}

// SetResourceIds sets field value.
func (o *SecurityMonitoringTerraformBulkExportAttributes) SetResourceIds(v []string) {
	o.ResourceIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringTerraformBulkExportAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["resource_ids"] = o.ResourceIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringTerraformBulkExportAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ResourceIds *[]string `json:"resource_ids"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ResourceIds == nil {
		return fmt.Errorf("required field resource_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"resource_ids"})
	} else {
		return err
	}
	o.ResourceIds = *all.ResourceIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
