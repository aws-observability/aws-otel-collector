// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationJiraSyncProperties Field synchronization properties for Jira integration.
type IntegrationJiraSyncProperties struct {
	// Sync property configuration.
	Assignee *SyncProperty `json:"assignee,omitempty"`
	// Sync property configuration.
	Comments *SyncProperty `json:"comments,omitempty"`
	// Map of custom field identifiers to their sync configurations.
	CustomFields map[string]IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties `json:"custom_fields,omitempty"`
	// Sync property configuration.
	Description *SyncProperty `json:"description,omitempty"`
	// Due date synchronization configuration for Jira integration.
	DueDate *IntegrationJiraSyncDueDate `json:"due_date,omitempty"`
	// Sync property with mapping configuration.
	Priority *SyncPropertyWithMapping `json:"priority,omitempty"`
	// Sync property with mapping configuration.
	Status *SyncPropertyWithMapping `json:"status,omitempty"`
	// Sync property configuration.
	Title *SyncProperty `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationJiraSyncProperties instantiates a new IntegrationJiraSyncProperties object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationJiraSyncProperties() *IntegrationJiraSyncProperties {
	this := IntegrationJiraSyncProperties{}
	return &this
}

// NewIntegrationJiraSyncPropertiesWithDefaults instantiates a new IntegrationJiraSyncProperties object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationJiraSyncPropertiesWithDefaults() *IntegrationJiraSyncProperties {
	this := IntegrationJiraSyncProperties{}
	return &this
}

// GetAssignee returns the Assignee field value if set, zero value otherwise.
func (o *IntegrationJiraSyncProperties) GetAssignee() SyncProperty {
	if o == nil || o.Assignee == nil {
		var ret SyncProperty
		return ret
	}
	return *o.Assignee
}

// GetAssigneeOk returns a tuple with the Assignee field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraSyncProperties) GetAssigneeOk() (*SyncProperty, bool) {
	if o == nil || o.Assignee == nil {
		return nil, false
	}
	return o.Assignee, true
}

// HasAssignee returns a boolean if a field has been set.
func (o *IntegrationJiraSyncProperties) HasAssignee() bool {
	return o != nil && o.Assignee != nil
}

// SetAssignee gets a reference to the given SyncProperty and assigns it to the Assignee field.
func (o *IntegrationJiraSyncProperties) SetAssignee(v SyncProperty) {
	o.Assignee = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *IntegrationJiraSyncProperties) GetComments() SyncProperty {
	if o == nil || o.Comments == nil {
		var ret SyncProperty
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraSyncProperties) GetCommentsOk() (*SyncProperty, bool) {
	if o == nil || o.Comments == nil {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *IntegrationJiraSyncProperties) HasComments() bool {
	return o != nil && o.Comments != nil
}

// SetComments gets a reference to the given SyncProperty and assigns it to the Comments field.
func (o *IntegrationJiraSyncProperties) SetComments(v SyncProperty) {
	o.Comments = &v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *IntegrationJiraSyncProperties) GetCustomFields() map[string]IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties {
	if o == nil || o.CustomFields == nil {
		var ret map[string]IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraSyncProperties) GetCustomFieldsOk() (*map[string]IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties, bool) {
	if o == nil || o.CustomFields == nil {
		return nil, false
	}
	return &o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *IntegrationJiraSyncProperties) HasCustomFields() bool {
	return o != nil && o.CustomFields != nil
}

// SetCustomFields gets a reference to the given map[string]IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties and assigns it to the CustomFields field.
func (o *IntegrationJiraSyncProperties) SetCustomFields(v map[string]IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties) {
	o.CustomFields = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IntegrationJiraSyncProperties) GetDescription() SyncProperty {
	if o == nil || o.Description == nil {
		var ret SyncProperty
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraSyncProperties) GetDescriptionOk() (*SyncProperty, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IntegrationJiraSyncProperties) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given SyncProperty and assigns it to the Description field.
func (o *IntegrationJiraSyncProperties) SetDescription(v SyncProperty) {
	o.Description = &v
}

// GetDueDate returns the DueDate field value if set, zero value otherwise.
func (o *IntegrationJiraSyncProperties) GetDueDate() IntegrationJiraSyncDueDate {
	if o == nil || o.DueDate == nil {
		var ret IntegrationJiraSyncDueDate
		return ret
	}
	return *o.DueDate
}

// GetDueDateOk returns a tuple with the DueDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraSyncProperties) GetDueDateOk() (*IntegrationJiraSyncDueDate, bool) {
	if o == nil || o.DueDate == nil {
		return nil, false
	}
	return o.DueDate, true
}

// HasDueDate returns a boolean if a field has been set.
func (o *IntegrationJiraSyncProperties) HasDueDate() bool {
	return o != nil && o.DueDate != nil
}

// SetDueDate gets a reference to the given IntegrationJiraSyncDueDate and assigns it to the DueDate field.
func (o *IntegrationJiraSyncProperties) SetDueDate(v IntegrationJiraSyncDueDate) {
	o.DueDate = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *IntegrationJiraSyncProperties) GetPriority() SyncPropertyWithMapping {
	if o == nil || o.Priority == nil {
		var ret SyncPropertyWithMapping
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraSyncProperties) GetPriorityOk() (*SyncPropertyWithMapping, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *IntegrationJiraSyncProperties) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given SyncPropertyWithMapping and assigns it to the Priority field.
func (o *IntegrationJiraSyncProperties) SetPriority(v SyncPropertyWithMapping) {
	o.Priority = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IntegrationJiraSyncProperties) GetStatus() SyncPropertyWithMapping {
	if o == nil || o.Status == nil {
		var ret SyncPropertyWithMapping
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraSyncProperties) GetStatusOk() (*SyncPropertyWithMapping, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IntegrationJiraSyncProperties) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SyncPropertyWithMapping and assigns it to the Status field.
func (o *IntegrationJiraSyncProperties) SetStatus(v SyncPropertyWithMapping) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *IntegrationJiraSyncProperties) GetTitle() SyncProperty {
	if o == nil || o.Title == nil {
		var ret SyncProperty
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationJiraSyncProperties) GetTitleOk() (*SyncProperty, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *IntegrationJiraSyncProperties) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given SyncProperty and assigns it to the Title field.
func (o *IntegrationJiraSyncProperties) SetTitle(v SyncProperty) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationJiraSyncProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assignee != nil {
		toSerialize["assignee"] = o.Assignee
	}
	if o.Comments != nil {
		toSerialize["comments"] = o.Comments
	}
	if o.CustomFields != nil {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DueDate != nil {
		toSerialize["due_date"] = o.DueDate
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationJiraSyncProperties) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assignee     *SyncProperty                                                            `json:"assignee,omitempty"`
		Comments     *SyncProperty                                                            `json:"comments,omitempty"`
		CustomFields map[string]IntegrationJiraSyncPropertiesCustomFieldsAdditionalProperties `json:"custom_fields,omitempty"`
		Description  *SyncProperty                                                            `json:"description,omitempty"`
		DueDate      *IntegrationJiraSyncDueDate                                              `json:"due_date,omitempty"`
		Priority     *SyncPropertyWithMapping                                                 `json:"priority,omitempty"`
		Status       *SyncPropertyWithMapping                                                 `json:"status,omitempty"`
		Title        *SyncProperty                                                            `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assignee", "comments", "custom_fields", "description", "due_date", "priority", "status", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Assignee != nil && all.Assignee.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Assignee = all.Assignee
	if all.Comments != nil && all.Comments.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Comments = all.Comments
	o.CustomFields = all.CustomFields
	if all.Description != nil && all.Description.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Description = all.Description
	if all.DueDate != nil && all.DueDate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DueDate = all.DueDate
	if all.Priority != nil && all.Priority.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Priority = all.Priority
	if all.Status != nil && all.Status.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Status = all.Status
	if all.Title != nil && all.Title.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
