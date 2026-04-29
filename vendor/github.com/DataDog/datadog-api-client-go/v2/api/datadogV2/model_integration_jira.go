// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationJira Jira integration settings.
type IntegrationJira struct {
	// Auto-creation settings for Jira issues from cases.
	AutoCreation *IntegrationJiraAutoCreation `json:"auto_creation,omitempty"`
	// Whether Jira integration is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Metadata for connecting a case management project to a Jira project.
	Metadata *IntegrationJiraMetadata `json:"metadata,omitempty"`
	// Synchronization configuration for Jira integration.
	Sync *IntegrationJiraSync `json:"sync,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationJira instantiates a new IntegrationJira object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationJira() *IntegrationJira {
	this := IntegrationJira{}
	return &this
}

// NewIntegrationJiraWithDefaults instantiates a new IntegrationJira object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationJiraWithDefaults() *IntegrationJira {
	this := IntegrationJira{}
	return &this
}

// GetAutoCreation returns the AutoCreation field value if set, zero value otherwise.
func (o *IntegrationJira) GetAutoCreation() IntegrationJiraAutoCreation {
	if o == nil || o.AutoCreation == nil {
		var ret IntegrationJiraAutoCreation
		return ret
	}
	return *o.AutoCreation
}

// GetAutoCreationOk returns a tuple with the AutoCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJira) GetAutoCreationOk() (*IntegrationJiraAutoCreation, bool) {
	if o == nil || o.AutoCreation == nil {
		return nil, false
	}
	return o.AutoCreation, true
}

// HasAutoCreation returns a boolean if a field has been set.
func (o *IntegrationJira) HasAutoCreation() bool {
	return o != nil && o.AutoCreation != nil
}

// SetAutoCreation gets a reference to the given IntegrationJiraAutoCreation and assigns it to the AutoCreation field.
func (o *IntegrationJira) SetAutoCreation(v IntegrationJiraAutoCreation) {
	o.AutoCreation = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IntegrationJira) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJira) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IntegrationJira) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IntegrationJira) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IntegrationJira) GetMetadata() IntegrationJiraMetadata {
	if o == nil || o.Metadata == nil {
		var ret IntegrationJiraMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJira) GetMetadataOk() (*IntegrationJiraMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IntegrationJira) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given IntegrationJiraMetadata and assigns it to the Metadata field.
func (o *IntegrationJira) SetMetadata(v IntegrationJiraMetadata) {
	o.Metadata = &v
}

// GetSync returns the Sync field value if set, zero value otherwise.
func (o *IntegrationJira) GetSync() IntegrationJiraSync {
	if o == nil || o.Sync == nil {
		var ret IntegrationJiraSync
		return ret
	}
	return *o.Sync
}

// GetSyncOk returns a tuple with the Sync field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJira) GetSyncOk() (*IntegrationJiraSync, bool) {
	if o == nil || o.Sync == nil {
		return nil, false
	}
	return o.Sync, true
}

// HasSync returns a boolean if a field has been set.
func (o *IntegrationJira) HasSync() bool {
	return o != nil && o.Sync != nil
}

// SetSync gets a reference to the given IntegrationJiraSync and assigns it to the Sync field.
func (o *IntegrationJira) SetSync(v IntegrationJiraSync) {
	o.Sync = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationJira) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutoCreation != nil {
		toSerialize["auto_creation"] = o.AutoCreation
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Sync != nil {
		toSerialize["sync"] = o.Sync
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationJira) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoCreation *IntegrationJiraAutoCreation `json:"auto_creation,omitempty"`
		Enabled      *bool                        `json:"enabled,omitempty"`
		Metadata     *IntegrationJiraMetadata     `json:"metadata,omitempty"`
		Sync         *IntegrationJiraSync         `json:"sync,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_creation", "enabled", "metadata", "sync"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AutoCreation != nil && all.AutoCreation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AutoCreation = all.AutoCreation
	o.Enabled = all.Enabled
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	if all.Sync != nil && all.Sync.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sync = all.Sync

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
