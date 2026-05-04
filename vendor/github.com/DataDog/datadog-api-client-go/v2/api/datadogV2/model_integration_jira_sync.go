// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationJiraSync Synchronization configuration for Jira integration.
type IntegrationJiraSync struct {
	// Whether Jira field synchronization is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Field synchronization properties for Jira integration.
	Properties *IntegrationJiraSyncProperties `json:"properties,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationJiraSync instantiates a new IntegrationJiraSync object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationJiraSync() *IntegrationJiraSync {
	this := IntegrationJiraSync{}
	return &this
}

// NewIntegrationJiraSyncWithDefaults instantiates a new IntegrationJiraSync object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationJiraSyncWithDefaults() *IntegrationJiraSync {
	this := IntegrationJiraSync{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IntegrationJiraSync) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraSync) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IntegrationJiraSync) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IntegrationJiraSync) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *IntegrationJiraSync) GetProperties() IntegrationJiraSyncProperties {
	if o == nil || o.Properties == nil {
		var ret IntegrationJiraSyncProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraSync) GetPropertiesOk() (*IntegrationJiraSyncProperties, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *IntegrationJiraSync) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given IntegrationJiraSyncProperties and assigns it to the Properties field.
func (o *IntegrationJiraSync) SetProperties(v IntegrationJiraSyncProperties) {
	o.Properties = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationJiraSync) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationJiraSync) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled    *bool                          `json:"enabled,omitempty"`
		Properties *IntegrationJiraSyncProperties `json:"properties,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "properties"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	if all.Properties != nil && all.Properties.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Properties = all.Properties

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
