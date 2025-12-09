// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyCreateRequestDataAttributes Defines the attributes for creating an escalation policy, including its description, name, resolution behavior, retries, and steps.
type EscalationPolicyCreateRequestDataAttributes struct {
	// Specifies the name for the new escalation policy.
	Name string `json:"name"`
	// Indicates whether the page is automatically resolved when the policy ends.
	ResolvePageOnPolicyEnd *bool `json:"resolve_page_on_policy_end,omitempty"`
	// Specifies how many times the escalation sequence is retried if there is no response.
	Retries *int64 `json:"retries,omitempty"`
	// A list of escalation steps, each defining assignment, escalation timeout, and targets for the new policy.
	Steps []EscalationPolicyCreateRequestDataAttributesStepsItems `json:"steps"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyCreateRequestDataAttributes instantiates a new EscalationPolicyCreateRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyCreateRequestDataAttributes(name string, steps []EscalationPolicyCreateRequestDataAttributesStepsItems) *EscalationPolicyCreateRequestDataAttributes {
	this := EscalationPolicyCreateRequestDataAttributes{}
	this.Name = name
	this.Steps = steps
	return &this
}

// NewEscalationPolicyCreateRequestDataAttributesWithDefaults instantiates a new EscalationPolicyCreateRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyCreateRequestDataAttributesWithDefaults() *EscalationPolicyCreateRequestDataAttributes {
	this := EscalationPolicyCreateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value.
func (o *EscalationPolicyCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EscalationPolicyCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetResolvePageOnPolicyEnd returns the ResolvePageOnPolicyEnd field value if set, zero value otherwise.
func (o *EscalationPolicyCreateRequestDataAttributes) GetResolvePageOnPolicyEnd() bool {
	if o == nil || o.ResolvePageOnPolicyEnd == nil {
		var ret bool
		return ret
	}
	return *o.ResolvePageOnPolicyEnd
}

// GetResolvePageOnPolicyEndOk returns a tuple with the ResolvePageOnPolicyEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyCreateRequestDataAttributes) GetResolvePageOnPolicyEndOk() (*bool, bool) {
	if o == nil || o.ResolvePageOnPolicyEnd == nil {
		return nil, false
	}
	return o.ResolvePageOnPolicyEnd, true
}

// HasResolvePageOnPolicyEnd returns a boolean if a field has been set.
func (o *EscalationPolicyCreateRequestDataAttributes) HasResolvePageOnPolicyEnd() bool {
	return o != nil && o.ResolvePageOnPolicyEnd != nil
}

// SetResolvePageOnPolicyEnd gets a reference to the given bool and assigns it to the ResolvePageOnPolicyEnd field.
func (o *EscalationPolicyCreateRequestDataAttributes) SetResolvePageOnPolicyEnd(v bool) {
	o.ResolvePageOnPolicyEnd = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *EscalationPolicyCreateRequestDataAttributes) GetRetries() int64 {
	if o == nil || o.Retries == nil {
		var ret int64
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyCreateRequestDataAttributes) GetRetriesOk() (*int64, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *EscalationPolicyCreateRequestDataAttributes) HasRetries() bool {
	return o != nil && o.Retries != nil
}

// SetRetries gets a reference to the given int64 and assigns it to the Retries field.
func (o *EscalationPolicyCreateRequestDataAttributes) SetRetries(v int64) {
	o.Retries = &v
}

// GetSteps returns the Steps field value.
func (o *EscalationPolicyCreateRequestDataAttributes) GetSteps() []EscalationPolicyCreateRequestDataAttributesStepsItems {
	if o == nil {
		var ret []EscalationPolicyCreateRequestDataAttributesStepsItems
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyCreateRequestDataAttributes) GetStepsOk() (*[]EscalationPolicyCreateRequestDataAttributesStepsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Steps, true
}

// SetSteps sets field value.
func (o *EscalationPolicyCreateRequestDataAttributes) SetSteps(v []EscalationPolicyCreateRequestDataAttributesStepsItems) {
	o.Steps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	if o.ResolvePageOnPolicyEnd != nil {
		toSerialize["resolve_page_on_policy_end"] = o.ResolvePageOnPolicyEnd
	}
	if o.Retries != nil {
		toSerialize["retries"] = o.Retries
	}
	toSerialize["steps"] = o.Steps

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicyCreateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                   *string                                                  `json:"name"`
		ResolvePageOnPolicyEnd *bool                                                    `json:"resolve_page_on_policy_end,omitempty"`
		Retries                *int64                                                   `json:"retries,omitempty"`
		Steps                  *[]EscalationPolicyCreateRequestDataAttributesStepsItems `json:"steps"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Steps == nil {
		return fmt.Errorf("required field steps missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "resolve_page_on_policy_end", "retries", "steps"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.ResolvePageOnPolicyEnd = all.ResolvePageOnPolicyEnd
	o.Retries = all.Retries
	o.Steps = *all.Steps

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
