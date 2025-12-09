// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationTemplateCreateAttributes The attributes for creating a notification template.
type IncidentNotificationTemplateCreateAttributes struct {
	// The category of the notification template.
	Category string `json:"category"`
	// The content body of the notification template.
	Content string `json:"content"`
	// The name of the notification template.
	Name string `json:"name"`
	// The subject line of the notification template.
	Subject string `json:"subject"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNotificationTemplateCreateAttributes instantiates a new IncidentNotificationTemplateCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNotificationTemplateCreateAttributes(category string, content string, name string, subject string) *IncidentNotificationTemplateCreateAttributes {
	this := IncidentNotificationTemplateCreateAttributes{}
	this.Category = category
	this.Content = content
	this.Name = name
	this.Subject = subject
	return &this
}

// NewIncidentNotificationTemplateCreateAttributesWithDefaults instantiates a new IncidentNotificationTemplateCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNotificationTemplateCreateAttributesWithDefaults() *IncidentNotificationTemplateCreateAttributes {
	this := IncidentNotificationTemplateCreateAttributes{}
	return &this
}

// GetCategory returns the Category field value.
func (o *IncidentNotificationTemplateCreateAttributes) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateCreateAttributes) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *IncidentNotificationTemplateCreateAttributes) SetCategory(v string) {
	o.Category = v
}

// GetContent returns the Content field value.
func (o *IncidentNotificationTemplateCreateAttributes) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateCreateAttributes) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *IncidentNotificationTemplateCreateAttributes) SetContent(v string) {
	o.Content = v
}

// GetName returns the Name field value.
func (o *IncidentNotificationTemplateCreateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateCreateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentNotificationTemplateCreateAttributes) SetName(v string) {
	o.Name = v
}

// GetSubject returns the Subject field value.
func (o *IncidentNotificationTemplateCreateAttributes) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateCreateAttributes) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value.
func (o *IncidentNotificationTemplateCreateAttributes) SetSubject(v string) {
	o.Subject = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNotificationTemplateCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["content"] = o.Content
	toSerialize["name"] = o.Name
	toSerialize["subject"] = o.Subject

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentNotificationTemplateCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category *string `json:"category"`
		Content  *string `json:"content"`
		Name     *string `json:"name"`
		Subject  *string `json:"subject"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Content == nil {
		return fmt.Errorf("required field content missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Subject == nil {
		return fmt.Errorf("required field subject missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "content", "name", "subject"})
	} else {
		return err
	}
	o.Category = *all.Category
	o.Content = *all.Content
	o.Name = *all.Name
	o.Subject = *all.Subject

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
