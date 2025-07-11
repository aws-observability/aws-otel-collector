// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GithubWebhookTriggerWrapper Schema for a GitHub webhook-based trigger.
type GithubWebhookTriggerWrapper struct {
	// Trigger a workflow from a GitHub webhook. To trigger a workflow from GitHub, you must set a `webhookSecret`. In your GitHub Webhook Settings, set the Payload URL to "base_url"/api/v2/workflows/"workflow_id"/webhook?orgId="org_id", select application/json for the content type, and be highly recommend enabling SSL verification for security. The workflow must be published.
	GithubWebhookTrigger GithubWebhookTrigger `json:"githubWebhookTrigger"`
	// A list of steps that run first after a trigger fires.
	StartStepNames []string `json:"startStepNames,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGithubWebhookTriggerWrapper instantiates a new GithubWebhookTriggerWrapper object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGithubWebhookTriggerWrapper(githubWebhookTrigger GithubWebhookTrigger) *GithubWebhookTriggerWrapper {
	this := GithubWebhookTriggerWrapper{}
	this.GithubWebhookTrigger = githubWebhookTrigger
	return &this
}

// NewGithubWebhookTriggerWrapperWithDefaults instantiates a new GithubWebhookTriggerWrapper object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGithubWebhookTriggerWrapperWithDefaults() *GithubWebhookTriggerWrapper {
	this := GithubWebhookTriggerWrapper{}
	return &this
}

// GetGithubWebhookTrigger returns the GithubWebhookTrigger field value.
func (o *GithubWebhookTriggerWrapper) GetGithubWebhookTrigger() GithubWebhookTrigger {
	if o == nil {
		var ret GithubWebhookTrigger
		return ret
	}
	return o.GithubWebhookTrigger
}

// GetGithubWebhookTriggerOk returns a tuple with the GithubWebhookTrigger field value
// and a boolean to check if the value has been set.
func (o *GithubWebhookTriggerWrapper) GetGithubWebhookTriggerOk() (*GithubWebhookTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GithubWebhookTrigger, true
}

// SetGithubWebhookTrigger sets field value.
func (o *GithubWebhookTriggerWrapper) SetGithubWebhookTrigger(v GithubWebhookTrigger) {
	o.GithubWebhookTrigger = v
}

// GetStartStepNames returns the StartStepNames field value if set, zero value otherwise.
func (o *GithubWebhookTriggerWrapper) GetStartStepNames() []string {
	if o == nil || o.StartStepNames == nil {
		var ret []string
		return ret
	}
	return o.StartStepNames
}

// GetStartStepNamesOk returns a tuple with the StartStepNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GithubWebhookTriggerWrapper) GetStartStepNamesOk() (*[]string, bool) {
	if o == nil || o.StartStepNames == nil {
		return nil, false
	}
	return &o.StartStepNames, true
}

// HasStartStepNames returns a boolean if a field has been set.
func (o *GithubWebhookTriggerWrapper) HasStartStepNames() bool {
	return o != nil && o.StartStepNames != nil
}

// SetStartStepNames gets a reference to the given []string and assigns it to the StartStepNames field.
func (o *GithubWebhookTriggerWrapper) SetStartStepNames(v []string) {
	o.StartStepNames = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GithubWebhookTriggerWrapper) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["githubWebhookTrigger"] = o.GithubWebhookTrigger
	if o.StartStepNames != nil {
		toSerialize["startStepNames"] = o.StartStepNames
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GithubWebhookTriggerWrapper) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GithubWebhookTrigger *GithubWebhookTrigger `json:"githubWebhookTrigger"`
		StartStepNames       []string              `json:"startStepNames,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.GithubWebhookTrigger == nil {
		return fmt.Errorf("required field githubWebhookTrigger missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"githubWebhookTrigger", "startStepNames"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.GithubWebhookTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GithubWebhookTrigger = *all.GithubWebhookTrigger
	o.StartStepNames = all.StartStepNames

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
