// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationTemplateAttributes The notification template's attributes.
type IncidentNotificationTemplateAttributes struct {
	// The category of the notification template.
	Category string `json:"category"`
	// The content body of the notification template.
	Content string `json:"content"`
	// Timestamp when the notification template was created.
	Created time.Time `json:"created"`
	// Timestamp when the notification template was last modified.
	Modified time.Time `json:"modified"`
	// The name of the notification template.
	Name string `json:"name"`
	// The subject line of the notification template.
	Subject string `json:"subject"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNotificationTemplateAttributes instantiates a new IncidentNotificationTemplateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNotificationTemplateAttributes(category string, content string, created time.Time, modified time.Time, name string, subject string) *IncidentNotificationTemplateAttributes {
	this := IncidentNotificationTemplateAttributes{}
	this.Category = category
	this.Content = content
	this.Created = created
	this.Modified = modified
	this.Name = name
	this.Subject = subject
	return &this
}

// NewIncidentNotificationTemplateAttributesWithDefaults instantiates a new IncidentNotificationTemplateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNotificationTemplateAttributesWithDefaults() *IncidentNotificationTemplateAttributes {
	this := IncidentNotificationTemplateAttributes{}
	return &this
}

// GetCategory returns the Category field value.
func (o *IncidentNotificationTemplateAttributes) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateAttributes) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *IncidentNotificationTemplateAttributes) SetCategory(v string) {
	o.Category = v
}

// GetContent returns the Content field value.
func (o *IncidentNotificationTemplateAttributes) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateAttributes) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value.
func (o *IncidentNotificationTemplateAttributes) SetContent(v string) {
	o.Content = v
}

// GetCreated returns the Created field value.
func (o *IncidentNotificationTemplateAttributes) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *IncidentNotificationTemplateAttributes) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value.
func (o *IncidentNotificationTemplateAttributes) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateAttributes) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *IncidentNotificationTemplateAttributes) SetModified(v time.Time) {
	o.Modified = v
}

// GetName returns the Name field value.
func (o *IncidentNotificationTemplateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentNotificationTemplateAttributes) SetName(v string) {
	o.Name = v
}

// GetSubject returns the Subject field value.
func (o *IncidentNotificationTemplateAttributes) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateAttributes) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value.
func (o *IncidentNotificationTemplateAttributes) SetSubject(v string) {
	o.Subject = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNotificationTemplateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["content"] = o.Content
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name
	toSerialize["subject"] = o.Subject

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentNotificationTemplateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category *string    `json:"category"`
		Content  *string    `json:"content"`
		Created  *time.Time `json:"created"`
		Modified *time.Time `json:"modified"`
		Name     *string    `json:"name"`
		Subject  *string    `json:"subject"`
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
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Subject == nil {
		return fmt.Errorf("required field subject missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "content", "created", "modified", "name", "subject"})
	} else {
		return err
	}
	o.Category = *all.Category
	o.Content = *all.Content
	o.Created = *all.Created
	o.Modified = *all.Modified
	o.Name = *all.Name
	o.Subject = *all.Subject

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
