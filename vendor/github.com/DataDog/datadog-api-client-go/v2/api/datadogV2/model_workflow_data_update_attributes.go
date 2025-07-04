// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowDataUpdateAttributes The definition of `WorkflowDataUpdateAttributes` object.
type WorkflowDataUpdateAttributes struct {
	// When the workflow was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Description of the workflow.
	Description *string `json:"description,omitempty"`
	// Name of the workflow.
	Name *string `json:"name,omitempty"`
	// Set the workflow to published or unpublished. Workflows in an unpublished state will only be executable via manual runs. Automatic triggers such as Schedule will not execute the workflow until it is published.
	Published *bool `json:"published,omitempty"`
	// The spec defines what the workflow does.
	Spec *Spec `json:"spec,omitempty"`
	// Tags of the workflow.
	Tags []string `json:"tags,omitempty"`
	// When the workflow was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// If a Webhook trigger is defined on this workflow, a webhookSecret is required and should be provided here.
	WebhookSecret *string `json:"webhookSecret,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowDataUpdateAttributes instantiates a new WorkflowDataUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowDataUpdateAttributes() *WorkflowDataUpdateAttributes {
	this := WorkflowDataUpdateAttributes{}
	return &this
}

// NewWorkflowDataUpdateAttributesWithDefaults instantiates a new WorkflowDataUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowDataUpdateAttributesWithDefaults() *WorkflowDataUpdateAttributes {
	this := WorkflowDataUpdateAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WorkflowDataUpdateAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataUpdateAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WorkflowDataUpdateAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WorkflowDataUpdateAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowDataUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowDataUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowDataUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowDataUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowDataUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowDataUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *WorkflowDataUpdateAttributes) GetPublished() bool {
	if o == nil || o.Published == nil {
		var ret bool
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataUpdateAttributes) GetPublishedOk() (*bool, bool) {
	if o == nil || o.Published == nil {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *WorkflowDataUpdateAttributes) HasPublished() bool {
	return o != nil && o.Published != nil
}

// SetPublished gets a reference to the given bool and assigns it to the Published field.
func (o *WorkflowDataUpdateAttributes) SetPublished(v bool) {
	o.Published = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *WorkflowDataUpdateAttributes) GetSpec() Spec {
	if o == nil || o.Spec == nil {
		var ret Spec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataUpdateAttributes) GetSpecOk() (*Spec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *WorkflowDataUpdateAttributes) HasSpec() bool {
	return o != nil && o.Spec != nil
}

// SetSpec gets a reference to the given Spec and assigns it to the Spec field.
func (o *WorkflowDataUpdateAttributes) SetSpec(v Spec) {
	o.Spec = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WorkflowDataUpdateAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataUpdateAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WorkflowDataUpdateAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *WorkflowDataUpdateAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WorkflowDataUpdateAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataUpdateAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WorkflowDataUpdateAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WorkflowDataUpdateAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetWebhookSecret returns the WebhookSecret field value if set, zero value otherwise.
func (o *WorkflowDataUpdateAttributes) GetWebhookSecret() string {
	if o == nil || o.WebhookSecret == nil {
		var ret string
		return ret
	}
	return *o.WebhookSecret
}

// GetWebhookSecretOk returns a tuple with the WebhookSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowDataUpdateAttributes) GetWebhookSecretOk() (*string, bool) {
	if o == nil || o.WebhookSecret == nil {
		return nil, false
	}
	return o.WebhookSecret, true
}

// HasWebhookSecret returns a boolean if a field has been set.
func (o *WorkflowDataUpdateAttributes) HasWebhookSecret() bool {
	return o != nil && o.WebhookSecret != nil
}

// SetWebhookSecret gets a reference to the given string and assigns it to the WebhookSecret field.
func (o *WorkflowDataUpdateAttributes) SetWebhookSecret(v string) {
	o.WebhookSecret = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowDataUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Published != nil {
		toSerialize["published"] = o.Published
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdatedAt != nil {
		if o.UpdatedAt.Nanosecond() == 0 {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["updatedAt"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.WebhookSecret != nil {
		toSerialize["webhookSecret"] = o.WebhookSecret
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowDataUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt     *time.Time `json:"createdAt,omitempty"`
		Description   *string    `json:"description,omitempty"`
		Name          *string    `json:"name,omitempty"`
		Published     *bool      `json:"published,omitempty"`
		Spec          *Spec      `json:"spec,omitempty"`
		Tags          []string   `json:"tags,omitempty"`
		UpdatedAt     *time.Time `json:"updatedAt,omitempty"`
		WebhookSecret *string    `json:"webhookSecret,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"createdAt", "description", "name", "published", "spec", "tags", "updatedAt", "webhookSecret"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.Description = all.Description
	o.Name = all.Name
	o.Published = all.Published
	if all.Spec != nil && all.Spec.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Spec = all.Spec
	o.Tags = all.Tags
	o.UpdatedAt = all.UpdatedAt
	o.WebhookSecret = all.WebhookSecret

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
