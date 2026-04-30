// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationServiceNow ServiceNow integration settings.
type IntegrationServiceNow struct {
	// Assignment group.
	AssignmentGroup *string `json:"assignment_group,omitempty"`
	// Auto-creation settings for ServiceNow incidents from cases.
	AutoCreation *IntegrationServiceNowAutoCreation `json:"auto_creation,omitempty"`
	// Whether ServiceNow integration is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// ServiceNow instance name.
	InstanceName *string `json:"instance_name,omitempty"`
	// Synchronization configuration for ServiceNow integration.
	SyncConfig *IntegrationServiceNowSyncConfig `json:"sync_config,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationServiceNow instantiates a new IntegrationServiceNow object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationServiceNow() *IntegrationServiceNow {
	this := IntegrationServiceNow{}
	return &this
}

// NewIntegrationServiceNowWithDefaults instantiates a new IntegrationServiceNow object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationServiceNowWithDefaults() *IntegrationServiceNow {
	this := IntegrationServiceNow{}
	return &this
}

// GetAssignmentGroup returns the AssignmentGroup field value if set, zero value otherwise.
func (o *IntegrationServiceNow) GetAssignmentGroup() string {
	if o == nil || o.AssignmentGroup == nil {
		var ret string
		return ret
	}
	return *o.AssignmentGroup
}

// GetAssignmentGroupOk returns a tuple with the AssignmentGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationServiceNow) GetAssignmentGroupOk() (*string, bool) {
	if o == nil || o.AssignmentGroup == nil {
		return nil, false
	}
	return o.AssignmentGroup, true
}

// HasAssignmentGroup returns a boolean if a field has been set.
func (o *IntegrationServiceNow) HasAssignmentGroup() bool {
	return o != nil && o.AssignmentGroup != nil
}

// SetAssignmentGroup gets a reference to the given string and assigns it to the AssignmentGroup field.
func (o *IntegrationServiceNow) SetAssignmentGroup(v string) {
	o.AssignmentGroup = &v
}

// GetAutoCreation returns the AutoCreation field value if set, zero value otherwise.
func (o *IntegrationServiceNow) GetAutoCreation() IntegrationServiceNowAutoCreation {
	if o == nil || o.AutoCreation == nil {
		var ret IntegrationServiceNowAutoCreation
		return ret
	}
	return *o.AutoCreation
}

// GetAutoCreationOk returns a tuple with the AutoCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationServiceNow) GetAutoCreationOk() (*IntegrationServiceNowAutoCreation, bool) {
	if o == nil || o.AutoCreation == nil {
		return nil, false
	}
	return o.AutoCreation, true
}

// HasAutoCreation returns a boolean if a field has been set.
func (o *IntegrationServiceNow) HasAutoCreation() bool {
	return o != nil && o.AutoCreation != nil
}

// SetAutoCreation gets a reference to the given IntegrationServiceNowAutoCreation and assigns it to the AutoCreation field.
func (o *IntegrationServiceNow) SetAutoCreation(v IntegrationServiceNowAutoCreation) {
	o.AutoCreation = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IntegrationServiceNow) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationServiceNow) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IntegrationServiceNow) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IntegrationServiceNow) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *IntegrationServiceNow) GetInstanceName() string {
	if o == nil || o.InstanceName == nil {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationServiceNow) GetInstanceNameOk() (*string, bool) {
	if o == nil || o.InstanceName == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *IntegrationServiceNow) HasInstanceName() bool {
	return o != nil && o.InstanceName != nil
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *IntegrationServiceNow) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetSyncConfig returns the SyncConfig field value if set, zero value otherwise.
func (o *IntegrationServiceNow) GetSyncConfig() IntegrationServiceNowSyncConfig {
	if o == nil || o.SyncConfig == nil {
		var ret IntegrationServiceNowSyncConfig
		return ret
	}
	return *o.SyncConfig
}

// GetSyncConfigOk returns a tuple with the SyncConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationServiceNow) GetSyncConfigOk() (*IntegrationServiceNowSyncConfig, bool) {
	if o == nil || o.SyncConfig == nil {
		return nil, false
	}
	return o.SyncConfig, true
}

// HasSyncConfig returns a boolean if a field has been set.
func (o *IntegrationServiceNow) HasSyncConfig() bool {
	return o != nil && o.SyncConfig != nil
}

// SetSyncConfig gets a reference to the given IntegrationServiceNowSyncConfig and assigns it to the SyncConfig field.
func (o *IntegrationServiceNow) SetSyncConfig(v IntegrationServiceNowSyncConfig) {
	o.SyncConfig = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationServiceNow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssignmentGroup != nil {
		toSerialize["assignment_group"] = o.AssignmentGroup
	}
	if o.AutoCreation != nil {
		toSerialize["auto_creation"] = o.AutoCreation
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.InstanceName != nil {
		toSerialize["instance_name"] = o.InstanceName
	}
	if o.SyncConfig != nil {
		toSerialize["sync_config"] = o.SyncConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationServiceNow) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssignmentGroup *string                            `json:"assignment_group,omitempty"`
		AutoCreation    *IntegrationServiceNowAutoCreation `json:"auto_creation,omitempty"`
		Enabled         *bool                              `json:"enabled,omitempty"`
		InstanceName    *string                            `json:"instance_name,omitempty"`
		SyncConfig      *IntegrationServiceNowSyncConfig   `json:"sync_config,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignment_group", "auto_creation", "enabled", "instance_name", "sync_config"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AssignmentGroup = all.AssignmentGroup
	if all.AutoCreation != nil && all.AutoCreation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AutoCreation = all.AutoCreation
	o.Enabled = all.Enabled
	o.InstanceName = all.InstanceName
	if all.SyncConfig != nil && all.SyncConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SyncConfig = all.SyncConfig

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
