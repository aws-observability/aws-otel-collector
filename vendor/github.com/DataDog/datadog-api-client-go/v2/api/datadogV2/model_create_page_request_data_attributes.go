// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreatePageRequestDataAttributes Details about the On-Call Page you want to create.
type CreatePageRequestDataAttributes struct {
	// A short summary of the issue or context.
	Description *string `json:"description,omitempty"`
	// Tags to help categorize or filter the page.
	Tags []string `json:"tags,omitempty"`
	// Information about the target to notify (such as a team or user).
	Target CreatePageRequestDataAttributesTarget `json:"target"`
	// The title of the page.
	Title string `json:"title"`
	// On-Call Page urgency level.
	Urgency PageUrgency `json:"urgency"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreatePageRequestDataAttributes instantiates a new CreatePageRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreatePageRequestDataAttributes(target CreatePageRequestDataAttributesTarget, title string, urgency PageUrgency) *CreatePageRequestDataAttributes {
	this := CreatePageRequestDataAttributes{}
	this.Target = target
	this.Title = title
	this.Urgency = urgency
	return &this
}

// NewCreatePageRequestDataAttributesWithDefaults instantiates a new CreatePageRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreatePageRequestDataAttributesWithDefaults() *CreatePageRequestDataAttributes {
	this := CreatePageRequestDataAttributes{}
	var urgency PageUrgency = PAGEURGENCY_HIGH
	this.Urgency = urgency
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreatePageRequestDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePageRequestDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreatePageRequestDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreatePageRequestDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CreatePageRequestDataAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePageRequestDataAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CreatePageRequestDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *CreatePageRequestDataAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetTarget returns the Target field value.
func (o *CreatePageRequestDataAttributes) GetTarget() CreatePageRequestDataAttributesTarget {
	if o == nil {
		var ret CreatePageRequestDataAttributesTarget
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *CreatePageRequestDataAttributes) GetTargetOk() (*CreatePageRequestDataAttributesTarget, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value.
func (o *CreatePageRequestDataAttributes) SetTarget(v CreatePageRequestDataAttributesTarget) {
	o.Target = v
}

// GetTitle returns the Title field value.
func (o *CreatePageRequestDataAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreatePageRequestDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *CreatePageRequestDataAttributes) SetTitle(v string) {
	o.Title = v
}

// GetUrgency returns the Urgency field value.
func (o *CreatePageRequestDataAttributes) GetUrgency() PageUrgency {
	if o == nil {
		var ret PageUrgency
		return ret
	}
	return o.Urgency
}

// GetUrgencyOk returns a tuple with the Urgency field value
// and a boolean to check if the value has been set.
func (o *CreatePageRequestDataAttributes) GetUrgencyOk() (*PageUrgency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Urgency, true
}

// SetUrgency sets field value.
func (o *CreatePageRequestDataAttributes) SetUrgency(v PageUrgency) {
	o.Urgency = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreatePageRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["target"] = o.Target
	toSerialize["title"] = o.Title
	toSerialize["urgency"] = o.Urgency

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreatePageRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                                `json:"description,omitempty"`
		Tags        []string                               `json:"tags,omitempty"`
		Target      *CreatePageRequestDataAttributesTarget `json:"target"`
		Title       *string                                `json:"title"`
		Urgency     *PageUrgency                           `json:"urgency"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Target == nil {
		return fmt.Errorf("required field target missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Urgency == nil {
		return fmt.Errorf("required field urgency missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "tags", "target", "title", "urgency"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Tags = all.Tags
	if all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = *all.Target
	o.Title = *all.Title
	if !all.Urgency.IsValid() {
		hasInvalidField = true
	} else {
		o.Urgency = *all.Urgency
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
