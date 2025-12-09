// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationTemplateUpdateAttributes The attributes to update on a notification template.
type IncidentNotificationTemplateUpdateAttributes struct {
	// The category of the notification template.
	Category *string `json:"category,omitempty"`
	// The content body of the notification template.
	Content *string `json:"content,omitempty"`
	// The name of the notification template.
	Name *string `json:"name,omitempty"`
	// The subject line of the notification template.
	Subject *string `json:"subject,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNotificationTemplateUpdateAttributes instantiates a new IncidentNotificationTemplateUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNotificationTemplateUpdateAttributes() *IncidentNotificationTemplateUpdateAttributes {
	this := IncidentNotificationTemplateUpdateAttributes{}
	return &this
}

// NewIncidentNotificationTemplateUpdateAttributesWithDefaults instantiates a new IncidentNotificationTemplateUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNotificationTemplateUpdateAttributesWithDefaults() *IncidentNotificationTemplateUpdateAttributes {
	this := IncidentNotificationTemplateUpdateAttributes{}
	return &this
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *IncidentNotificationTemplateUpdateAttributes) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateUpdateAttributes) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *IncidentNotificationTemplateUpdateAttributes) HasCategory() bool {
	return o != nil && o.Category != nil
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *IncidentNotificationTemplateUpdateAttributes) SetCategory(v string) {
	o.Category = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *IncidentNotificationTemplateUpdateAttributes) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateUpdateAttributes) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *IncidentNotificationTemplateUpdateAttributes) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *IncidentNotificationTemplateUpdateAttributes) SetContent(v string) {
	o.Content = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IncidentNotificationTemplateUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IncidentNotificationTemplateUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IncidentNotificationTemplateUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *IncidentNotificationTemplateUpdateAttributes) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentNotificationTemplateUpdateAttributes) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *IncidentNotificationTemplateUpdateAttributes) HasSubject() bool {
	return o != nil && o.Subject != nil
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *IncidentNotificationTemplateUpdateAttributes) SetSubject(v string) {
	o.Subject = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNotificationTemplateUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentNotificationTemplateUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category *string `json:"category,omitempty"`
		Content  *string `json:"content,omitempty"`
		Name     *string `json:"name,omitempty"`
		Subject  *string `json:"subject,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "content", "name", "subject"})
	} else {
		return err
	}
	o.Category = all.Category
	o.Content = all.Content
	o.Name = all.Name
	o.Subject = all.Subject

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
