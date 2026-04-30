// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationServiceNowSyncConfig139772721534496 Field-level synchronization properties for ServiceNow integration.
type IntegrationServiceNowSyncConfig139772721534496 struct {
	// Sync property configuration.
	Comments *SyncProperty `json:"comments,omitempty"`
	// Priority synchronization configuration for ServiceNow integration.
	Priority *IntegrationServiceNowSyncConfigPriority `json:"priority,omitempty"`
	// Sync property with mapping configuration.
	Status *SyncPropertyWithMapping `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationServiceNowSyncConfig139772721534496 instantiates a new IntegrationServiceNowSyncConfig139772721534496 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationServiceNowSyncConfig139772721534496() *IntegrationServiceNowSyncConfig139772721534496 {
	this := IntegrationServiceNowSyncConfig139772721534496{}
	return &this
}

// NewIntegrationServiceNowSyncConfig139772721534496WithDefaults instantiates a new IntegrationServiceNowSyncConfig139772721534496 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationServiceNowSyncConfig139772721534496WithDefaults() *IntegrationServiceNowSyncConfig139772721534496 {
	this := IntegrationServiceNowSyncConfig139772721534496{}
	return &this
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IntegrationServiceNowSyncConfig139772721534496) GetComments() SyncProperty {
	if o == nil || o.Comments == nil {
		var ret SyncProperty
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationServiceNowSyncConfig139772721534496) GetCommentsOk() (*SyncProperty, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IntegrationServiceNowSyncConfig139772721534496) HasComments() bool {
	return o != nil && o.Comments != nil
}

// SetComments gets a reference to the given SyncProperty and assigns it to the Comments field.
func (o *IntegrationServiceNowSyncConfig139772721534496) SetComments(v SyncProperty) {
	o.Comments = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *IntegrationServiceNowSyncConfig139772721534496) GetPriority() IntegrationServiceNowSyncConfigPriority {
	if o == nil || o.Priority == nil {
		var ret IntegrationServiceNowSyncConfigPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationServiceNowSyncConfig139772721534496) GetPriorityOk() (*IntegrationServiceNowSyncConfigPriority, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *IntegrationServiceNowSyncConfig139772721534496) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given IntegrationServiceNowSyncConfigPriority and assigns it to the Priority field.
func (o *IntegrationServiceNowSyncConfig139772721534496) SetPriority(v IntegrationServiceNowSyncConfigPriority) {
	o.Priority = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IntegrationServiceNowSyncConfig139772721534496) GetStatus() SyncPropertyWithMapping {
	if o == nil || o.Status == nil {
		var ret SyncPropertyWithMapping
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationServiceNowSyncConfig139772721534496) GetStatusOk() (*SyncPropertyWithMapping, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IntegrationServiceNowSyncConfig139772721534496) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SyncPropertyWithMapping and assigns it to the Status field.
func (o *IntegrationServiceNowSyncConfig139772721534496) SetStatus(v SyncPropertyWithMapping) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationServiceNowSyncConfig139772721534496) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationServiceNowSyncConfig139772721534496) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Comments *SyncProperty                            `json:"comments,omitempty"`
		Priority *IntegrationServiceNowSyncConfigPriority `json:"priority,omitempty"`
		Status   *SyncPropertyWithMapping                 `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"comments", "priority", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Comments != nil && all.Comments.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Comments = all.Comments
	if all.Priority != nil && all.Priority.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Priority = all.Priority
	if all.Status != nil && all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
