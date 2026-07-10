// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WorkflowListItemAttributes Attributes of a workflow returned in a list response.
type WorkflowListItemAttributes struct {
	// When the workflow was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Description of the workflow.
	Description *string `json:"description,omitempty"`
	// Name of the workflow.
	Name string `json:"name"`
	// Whether the workflow is published. Unpublished workflows can only be run manually. Automatic triggers such as Schedule do not fire until the workflow is published.
	Published *bool `json:"published,omitempty"`
	// The spec defines what the workflow does.
	Spec *Spec `json:"spec,omitempty"`
	// Tags of the workflow.
	Tags []string `json:"tags,omitempty"`
	// When the workflow was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWorkflowListItemAttributes instantiates a new WorkflowListItemAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWorkflowListItemAttributes(name string) *WorkflowListItemAttributes {
	this := WorkflowListItemAttributes{}
	this.Name = name
	return &this
}

// NewWorkflowListItemAttributesWithDefaults instantiates a new WorkflowListItemAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWorkflowListItemAttributesWithDefaults() *WorkflowListItemAttributes {
	this := WorkflowListItemAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *WorkflowListItemAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowListItemAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *WorkflowListItemAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *WorkflowListItemAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowListItemAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowListItemAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowListItemAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowListItemAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value.
func (o *WorkflowListItemAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkflowListItemAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *WorkflowListItemAttributes) SetName(v string) {
	o.Name = v
}

// GetPublished returns the Published field value if set, zero value otherwise.
func (o *WorkflowListItemAttributes) GetPublished() bool {
	if o == nil || o.Published == nil {
		var ret bool
		return ret
	}
	return *o.Published
}

// GetPublishedOk returns a tuple with the Published field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowListItemAttributes) GetPublishedOk() (*bool, bool) {
	if o == nil || o.Published == nil {
		return nil, false
	}
	return o.Published, true
}

// HasPublished returns a boolean if a field has been set.
func (o *WorkflowListItemAttributes) HasPublished() bool {
	return o != nil && o.Published != nil
}

// SetPublished gets a reference to the given bool and assigns it to the Published field.
func (o *WorkflowListItemAttributes) SetPublished(v bool) {
	o.Published = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *WorkflowListItemAttributes) GetSpec() Spec {
	if o == nil || o.Spec == nil {
		var ret Spec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowListItemAttributes) GetSpecOk() (*Spec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *WorkflowListItemAttributes) HasSpec() bool {
	return o != nil && o.Spec != nil
}

// SetSpec gets a reference to the given Spec and assigns it to the Spec field.
func (o *WorkflowListItemAttributes) SetSpec(v Spec) {
	o.Spec = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WorkflowListItemAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowListItemAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WorkflowListItemAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *WorkflowListItemAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WorkflowListItemAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowListItemAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WorkflowListItemAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *WorkflowListItemAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WorkflowListItemAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["name"] = o.Name
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WorkflowListItemAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time `json:"createdAt,omitempty"`
		Description *string    `json:"description,omitempty"`
		Name        *string    `json:"name"`
		Published   *bool      `json:"published,omitempty"`
		Spec        *Spec      `json:"spec,omitempty"`
		Tags        []string   `json:"tags,omitempty"`
		UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"createdAt", "description", "name", "published", "spec", "tags", "updatedAt"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = all.CreatedAt
	o.Description = all.Description
	o.Name = *all.Name
	o.Published = all.Published
	if all.Spec != nil && all.Spec.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Spec = all.Spec
	o.Tags = all.Tags
	o.UpdatedAt = all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
