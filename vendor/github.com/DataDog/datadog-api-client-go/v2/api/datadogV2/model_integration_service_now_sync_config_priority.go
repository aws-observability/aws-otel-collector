// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationServiceNowSyncConfigPriority Priority synchronization configuration for ServiceNow integration.
type IntegrationServiceNowSyncConfigPriority struct {
	// Mapping of case priority values to ServiceNow impact values.
	ImpactMapping map[string]string `json:"impact_mapping,omitempty"`
	// The type of synchronization to apply for priority.
	SyncType *string `json:"sync_type,omitempty"`
	// Mapping of case priority values to ServiceNow urgency values.
	UrgencyMapping map[string]string `json:"urgency_mapping,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationServiceNowSyncConfigPriority instantiates a new IntegrationServiceNowSyncConfigPriority object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationServiceNowSyncConfigPriority() *IntegrationServiceNowSyncConfigPriority {
	this := IntegrationServiceNowSyncConfigPriority{}
	return &this
}

// NewIntegrationServiceNowSyncConfigPriorityWithDefaults instantiates a new IntegrationServiceNowSyncConfigPriority object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationServiceNowSyncConfigPriorityWithDefaults() *IntegrationServiceNowSyncConfigPriority {
	this := IntegrationServiceNowSyncConfigPriority{}
	return &this
}

// GetImpactMapping returns the ImpactMapping field value if set, zero value otherwise.
func (o *IntegrationServiceNowSyncConfigPriority) GetImpactMapping() map[string]string {
	if o == nil || o.ImpactMapping == nil {
		var ret map[string]string
		return ret
	}
	return o.ImpactMapping
}

// GetImpactMappingOk returns a tuple with the ImpactMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationServiceNowSyncConfigPriority) GetImpactMappingOk() (*map[string]string, bool) {
	if o == nil || o.ImpactMapping == nil {
		return nil, false
	}
	return &o.ImpactMapping, true
}

// HasImpactMapping returns a boolean if a field has been set.
func (o *IntegrationServiceNowSyncConfigPriority) HasImpactMapping() bool {
	return o != nil && o.ImpactMapping != nil
}

// SetImpactMapping gets a reference to the given map[string]string and assigns it to the ImpactMapping field.
func (o *IntegrationServiceNowSyncConfigPriority) SetImpactMapping(v map[string]string) {
	o.ImpactMapping = v
}

// GetSyncType returns the SyncType field value if set, zero value otherwise.
func (o *IntegrationServiceNowSyncConfigPriority) GetSyncType() string {
	if o == nil || o.SyncType == nil {
		var ret string
		return ret
	}
	return *o.SyncType
}

// GetSyncTypeOk returns a tuple with the SyncType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationServiceNowSyncConfigPriority) GetSyncTypeOk() (*string, bool) {
	if o == nil || o.SyncType == nil {
		return nil, false
	}
	return o.SyncType, true
}

// HasSyncType returns a boolean if a field has been set.
func (o *IntegrationServiceNowSyncConfigPriority) HasSyncType() bool {
	return o != nil && o.SyncType != nil
}

// SetSyncType gets a reference to the given string and assigns it to the SyncType field.
func (o *IntegrationServiceNowSyncConfigPriority) SetSyncType(v string) {
	o.SyncType = &v
}

// GetUrgencyMapping returns the UrgencyMapping field value if set, zero value otherwise.
func (o *IntegrationServiceNowSyncConfigPriority) GetUrgencyMapping() map[string]string {
	if o == nil || o.UrgencyMapping == nil {
		var ret map[string]string
		return ret
	}
	return o.UrgencyMapping
}

// GetUrgencyMappingOk returns a tuple with the UrgencyMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationServiceNowSyncConfigPriority) GetUrgencyMappingOk() (*map[string]string, bool) {
	if o == nil || o.UrgencyMapping == nil {
		return nil, false
	}
	return &o.UrgencyMapping, true
}

// HasUrgencyMapping returns a boolean if a field has been set.
func (o *IntegrationServiceNowSyncConfigPriority) HasUrgencyMapping() bool {
	return o != nil && o.UrgencyMapping != nil
}

// SetUrgencyMapping gets a reference to the given map[string]string and assigns it to the UrgencyMapping field.
func (o *IntegrationServiceNowSyncConfigPriority) SetUrgencyMapping(v map[string]string) {
	o.UrgencyMapping = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationServiceNowSyncConfigPriority) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ImpactMapping != nil {
		toSerialize["impact_mapping"] = o.ImpactMapping
	}
	if o.SyncType != nil {
		toSerialize["sync_type"] = o.SyncType
	}
	if o.UrgencyMapping != nil {
		toSerialize["urgency_mapping"] = o.UrgencyMapping
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationServiceNowSyncConfigPriority) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ImpactMapping  map[string]string `json:"impact_mapping,omitempty"`
		SyncType       *string           `json:"sync_type,omitempty"`
		UrgencyMapping map[string]string `json:"urgency_mapping,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"impact_mapping", "sync_type", "urgency_mapping"})
	} else {
		return err
	}
	o.ImpactMapping = all.ImpactMapping
	o.SyncType = all.SyncType
	o.UrgencyMapping = all.UrgencyMapping

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
