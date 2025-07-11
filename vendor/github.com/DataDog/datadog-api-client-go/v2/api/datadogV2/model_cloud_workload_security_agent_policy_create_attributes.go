// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentPolicyCreateAttributes Create a new Cloud Workload Security Agent policy
type CloudWorkloadSecurityAgentPolicyCreateAttributes struct {
	// The description of the policy
	Description *string `json:"description,omitempty"`
	// Whether the policy is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// The host tags defining where this policy is deployed
	HostTags []string `json:"hostTags,omitempty"`
	// The host tags defining where this policy is deployed, the inner values are linked with AND, the outer values are linked with OR
	HostTagsLists [][]string `json:"hostTagsLists,omitempty"`
	// The name of the policy
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentPolicyCreateAttributes instantiates a new CloudWorkloadSecurityAgentPolicyCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentPolicyCreateAttributes(name string) *CloudWorkloadSecurityAgentPolicyCreateAttributes {
	this := CloudWorkloadSecurityAgentPolicyCreateAttributes{}
	this.Name = name
	return &this
}

// NewCloudWorkloadSecurityAgentPolicyCreateAttributesWithDefaults instantiates a new CloudWorkloadSecurityAgentPolicyCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentPolicyCreateAttributesWithDefaults() *CloudWorkloadSecurityAgentPolicyCreateAttributes {
	this := CloudWorkloadSecurityAgentPolicyCreateAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHostTags returns the HostTags field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) GetHostTags() []string {
	if o == nil || o.HostTags == nil {
		var ret []string
		return ret
	}
	return o.HostTags
}

// GetHostTagsOk returns a tuple with the HostTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) GetHostTagsOk() (*[]string, bool) {
	if o == nil || o.HostTags == nil {
		return nil, false
	}
	return &o.HostTags, true
}

// HasHostTags returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) HasHostTags() bool {
	return o != nil && o.HostTags != nil
}

// SetHostTags gets a reference to the given []string and assigns it to the HostTags field.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) SetHostTags(v []string) {
	o.HostTags = v
}

// GetHostTagsLists returns the HostTagsLists field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) GetHostTagsLists() [][]string {
	if o == nil || o.HostTagsLists == nil {
		var ret [][]string
		return ret
	}
	return o.HostTagsLists
}

// GetHostTagsListsOk returns a tuple with the HostTagsLists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) GetHostTagsListsOk() (*[][]string, bool) {
	if o == nil || o.HostTagsLists == nil {
		return nil, false
	}
	return &o.HostTagsLists, true
}

// HasHostTagsLists returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) HasHostTagsLists() bool {
	return o != nil && o.HostTagsLists != nil
}

// SetHostTagsLists gets a reference to the given [][]string and assigns it to the HostTagsLists field.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) SetHostTagsLists(v [][]string) {
	o.HostTagsLists = v
}

// GetName returns the Name field value.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentPolicyCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.HostTags != nil {
		toSerialize["hostTags"] = o.HostTags
	}
	if o.HostTagsLists != nil {
		toSerialize["hostTagsLists"] = o.HostTagsLists
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudWorkloadSecurityAgentPolicyCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description   *string    `json:"description,omitempty"`
		Enabled       *bool      `json:"enabled,omitempty"`
		HostTags      []string   `json:"hostTags,omitempty"`
		HostTagsLists [][]string `json:"hostTagsLists,omitempty"`
		Name          *string    `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "enabled", "hostTags", "hostTagsLists", "name"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Enabled = all.Enabled
	o.HostTags = all.HostTags
	o.HostTagsLists = all.HostTagsLists
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
