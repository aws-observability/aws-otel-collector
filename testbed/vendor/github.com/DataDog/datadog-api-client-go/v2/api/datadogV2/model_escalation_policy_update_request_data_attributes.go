// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyUpdateRequestDataAttributes Defines the attributes that can be updated for an escalation policy, such as description, name, resolution behavior, retries, and steps.
type EscalationPolicyUpdateRequestDataAttributes struct {
	// Specifies the name of the escalation policy.
	Name string `json:"name"`
	// Indicates whether the page is automatically resolved when the policy ends.
	ResolvePageOnPolicyEnd *bool `json:"resolve_page_on_policy_end,omitempty"`
	// Specifies how many times the escalation sequence is retried if there is no response.
	Retries *int64 `json:"retries,omitempty"`
	// A list of escalation steps, each defining assignment, escalation timeout, and targets.
	Steps []EscalationPolicyUpdateRequestDataAttributesStepsItems `json:"steps"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyUpdateRequestDataAttributes instantiates a new EscalationPolicyUpdateRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyUpdateRequestDataAttributes(name string, steps []EscalationPolicyUpdateRequestDataAttributesStepsItems) *EscalationPolicyUpdateRequestDataAttributes {
	this := EscalationPolicyUpdateRequestDataAttributes{}
	this.Name = name
	this.Steps = steps
	return &this
}

// NewEscalationPolicyUpdateRequestDataAttributesWithDefaults instantiates a new EscalationPolicyUpdateRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyUpdateRequestDataAttributesWithDefaults() *EscalationPolicyUpdateRequestDataAttributes {
	this := EscalationPolicyUpdateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value.
func (o *EscalationPolicyUpdateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyUpdateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *EscalationPolicyUpdateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetResolvePageOnPolicyEnd returns the ResolvePageOnPolicyEnd field value if set, zero value otherwise.
func (o *EscalationPolicyUpdateRequestDataAttributes) GetResolvePageOnPolicyEnd() bool {
	if o == nil || o.ResolvePageOnPolicyEnd == nil {
		var ret bool
		return ret
	}
	return *o.ResolvePageOnPolicyEnd
}

// GetResolvePageOnPolicyEndOk returns a tuple with the ResolvePageOnPolicyEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyUpdateRequestDataAttributes) GetResolvePageOnPolicyEndOk() (*bool, bool) {
	if o == nil || o.ResolvePageOnPolicyEnd == nil {
		return nil, false
	}
	return o.ResolvePageOnPolicyEnd, true
}

// HasResolvePageOnPolicyEnd returns a boolean if a field has been set.
func (o *EscalationPolicyUpdateRequestDataAttributes) HasResolvePageOnPolicyEnd() bool {
	return o != nil && o.ResolvePageOnPolicyEnd != nil
}

// SetResolvePageOnPolicyEnd gets a reference to the given bool and assigns it to the ResolvePageOnPolicyEnd field.
func (o *EscalationPolicyUpdateRequestDataAttributes) SetResolvePageOnPolicyEnd(v bool) {
	o.ResolvePageOnPolicyEnd = &v
}

// GetRetries returns the Retries field value if set, zero value otherwise.
func (o *EscalationPolicyUpdateRequestDataAttributes) GetRetries() int64 {
	if o == nil || o.Retries == nil {
		var ret int64
		return ret
	}
	return *o.Retries
}

// GetRetriesOk returns a tuple with the Retries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyUpdateRequestDataAttributes) GetRetriesOk() (*int64, bool) {
	if o == nil || o.Retries == nil {
		return nil, false
	}
	return o.Retries, true
}

// HasRetries returns a boolean if a field has been set.
func (o *EscalationPolicyUpdateRequestDataAttributes) HasRetries() bool {
	return o != nil && o.Retries != nil
}

// SetRetries gets a reference to the given int64 and assigns it to the Retries field.
func (o *EscalationPolicyUpdateRequestDataAttributes) SetRetries(v int64) {
	o.Retries = &v
}

// GetSteps returns the Steps field value.
func (o *EscalationPolicyUpdateRequestDataAttributes) GetSteps() []EscalationPolicyUpdateRequestDataAttributesStepsItems {
	if o == nil {
		var ret []EscalationPolicyUpdateRequestDataAttributesStepsItems
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyUpdateRequestDataAttributes) GetStepsOk() (*[]EscalationPolicyUpdateRequestDataAttributesStepsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Steps, true
}

// SetSteps sets field value.
func (o *EscalationPolicyUpdateRequestDataAttributes) SetSteps(v []EscalationPolicyUpdateRequestDataAttributesStepsItems) {
	o.Steps = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
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
func (o *EscalationPolicyUpdateRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name                   *string                                                  `json:"name"`
		ResolvePageOnPolicyEnd *bool                                                    `json:"resolve_page_on_policy_end,omitempty"`
		Retries                *int64                                                   `json:"retries,omitempty"`
		Steps                  *[]EscalationPolicyUpdateRequestDataAttributesStepsItems `json:"steps"`
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
