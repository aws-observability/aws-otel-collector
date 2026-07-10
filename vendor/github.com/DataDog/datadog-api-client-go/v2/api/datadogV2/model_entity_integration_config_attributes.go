// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityIntegrationConfigAttributes The organization ID, integration identifier, and integration-specific configuration payload for an entity integration configuration.
type EntityIntegrationConfigAttributes struct {
	// Integration-specific configuration payload. The shape of this object depends on the integration identified by the path parameter. For `github`, the object must contain an `enabled_repos` array. For `jira`, it must contain an `enabled_projects` array. For `pagerduty`, it must contain an `accounts` array.
	Config map[string]interface{} `json:"config"`
	// The identifier of the integration this configuration applies to (for example, `github`, `jira`, or `pagerduty`).
	IntegrationId string `json:"integration_id"`
	// The Datadog organization identifier that owns this configuration.
	OrgId int64 `json:"org_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityIntegrationConfigAttributes instantiates a new EntityIntegrationConfigAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityIntegrationConfigAttributes(config map[string]interface{}, integrationId string, orgId int64) *EntityIntegrationConfigAttributes {
	this := EntityIntegrationConfigAttributes{}
	this.Config = config
	this.IntegrationId = integrationId
	this.OrgId = orgId
	return &this
}

// NewEntityIntegrationConfigAttributesWithDefaults instantiates a new EntityIntegrationConfigAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityIntegrationConfigAttributesWithDefaults() *EntityIntegrationConfigAttributes {
	this := EntityIntegrationConfigAttributes{}
	return &this
}

// GetConfig returns the Config field value.
func (o *EntityIntegrationConfigAttributes) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *EntityIntegrationConfigAttributes) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value.
func (o *EntityIntegrationConfigAttributes) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetIntegrationId returns the IntegrationId field value.
func (o *EntityIntegrationConfigAttributes) GetIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value
// and a boolean to check if the value has been set.
func (o *EntityIntegrationConfigAttributes) GetIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IntegrationId, true
}

// SetIntegrationId sets field value.
func (o *EntityIntegrationConfigAttributes) SetIntegrationId(v string) {
	o.IntegrationId = v
}

// GetOrgId returns the OrgId field value.
func (o *EntityIntegrationConfigAttributes) GetOrgId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *EntityIntegrationConfigAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value.
func (o *EntityIntegrationConfigAttributes) SetOrgId(v int64) {
	o.OrgId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityIntegrationConfigAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["config"] = o.Config
	toSerialize["integration_id"] = o.IntegrationId
	toSerialize["org_id"] = o.OrgId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityIntegrationConfigAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Config        *map[string]interface{} `json:"config"`
		IntegrationId *string                 `json:"integration_id"`
		OrgId         *int64                  `json:"org_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Config == nil {
		return fmt.Errorf("required field config missing")
	}
	if all.IntegrationId == nil {
		return fmt.Errorf("required field integration_id missing")
	}
	if all.OrgId == nil {
		return fmt.Errorf("required field org_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config", "integration_id", "org_id"})
	} else {
		return err
	}
	o.Config = *all.Config
	o.IntegrationId = *all.IntegrationId
	o.OrgId = *all.OrgId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
