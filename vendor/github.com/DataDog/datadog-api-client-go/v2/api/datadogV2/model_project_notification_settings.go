// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProjectNotificationSettings Project notification settings.
type ProjectNotificationSettings struct {
	// Notification destinations (1=email, 2=slack, 3=in-app).
	Destinations []int32 `json:"destinations,omitempty"`
	// Whether notifications are enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Whether to send a notification when a case is assigned.
	NotifyOnCaseAssignment *bool `json:"notify_on_case_assignment,omitempty"`
	// Whether to send a notification when a case is closed.
	NotifyOnCaseClosed *bool `json:"notify_on_case_closed,omitempty"`
	// Whether to send a notification when a comment is added to a case.
	NotifyOnCaseComment *bool `json:"notify_on_case_comment,omitempty"`
	// Whether to send a notification when a user is mentioned in a case comment.
	NotifyOnCaseCommentMention *bool `json:"notify_on_case_comment_mention,omitempty"`
	// Whether to send a notification when a case's priority changes.
	NotifyOnCasePriorityChange *bool `json:"notify_on_case_priority_change,omitempty"`
	// Whether to send a notification when a case's status changes.
	NotifyOnCaseStatusChange *bool `json:"notify_on_case_status_change,omitempty"`
	// Whether to send a notification when a case is unassigned.
	NotifyOnCaseUnassignment *bool `json:"notify_on_case_unassignment,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProjectNotificationSettings instantiates a new ProjectNotificationSettings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProjectNotificationSettings() *ProjectNotificationSettings {
	this := ProjectNotificationSettings{}
	return &this
}

// NewProjectNotificationSettingsWithDefaults instantiates a new ProjectNotificationSettings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectNotificationSettingsWithDefaults() *ProjectNotificationSettings {
	this := ProjectNotificationSettings{}
	return &this
}

// GetDestinations returns the Destinations field value if set, zero value otherwise.
func (o *ProjectNotificationSettings) GetDestinations() []int32 {
	if o == nil || o.Destinations == nil {
		var ret []int32
		return ret
	}
	return o.Destinations
}

// GetDestinationsOk returns a tuple with the Destinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectNotificationSettings) GetDestinationsOk() (*[]int32, bool) {
	if o == nil || o.Destinations == nil {
		return nil, false
	}
	return &o.Destinations, true
}

// HasDestinations returns a boolean if a field has been set.
func (o *ProjectNotificationSettings) HasDestinations() bool {
	return o != nil && o.Destinations != nil
}

// SetDestinations gets a reference to the given []int32 and assigns it to the Destinations field.
func (o *ProjectNotificationSettings) SetDestinations(v []int32) {
	o.Destinations = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ProjectNotificationSettings) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectNotificationSettings) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ProjectNotificationSettings) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ProjectNotificationSettings) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetNotifyOnCaseAssignment returns the NotifyOnCaseAssignment field value if set, zero value otherwise.
func (o *ProjectNotificationSettings) GetNotifyOnCaseAssignment() bool {
	if o == nil || o.NotifyOnCaseAssignment == nil {
		var ret bool
		return ret
	}
	return *o.NotifyOnCaseAssignment
}

// GetNotifyOnCaseAssignmentOk returns a tuple with the NotifyOnCaseAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectNotificationSettings) GetNotifyOnCaseAssignmentOk() (*bool, bool) {
	if o == nil || o.NotifyOnCaseAssignment == nil {
		return nil, false
	}
	return o.NotifyOnCaseAssignment, true
}

// HasNotifyOnCaseAssignment returns a boolean if a field has been set.
func (o *ProjectNotificationSettings) HasNotifyOnCaseAssignment() bool {
	return o != nil && o.NotifyOnCaseAssignment != nil
}

// SetNotifyOnCaseAssignment gets a reference to the given bool and assigns it to the NotifyOnCaseAssignment field.
func (o *ProjectNotificationSettings) SetNotifyOnCaseAssignment(v bool) {
	o.NotifyOnCaseAssignment = &v
}

// GetNotifyOnCaseClosed returns the NotifyOnCaseClosed field value if set, zero value otherwise.
func (o *ProjectNotificationSettings) GetNotifyOnCaseClosed() bool {
	if o == nil || o.NotifyOnCaseClosed == nil {
		var ret bool
		return ret
	}
	return *o.NotifyOnCaseClosed
}

// GetNotifyOnCaseClosedOk returns a tuple with the NotifyOnCaseClosed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectNotificationSettings) GetNotifyOnCaseClosedOk() (*bool, bool) {
	if o == nil || o.NotifyOnCaseClosed == nil {
		return nil, false
	}
	return o.NotifyOnCaseClosed, true
}

// HasNotifyOnCaseClosed returns a boolean if a field has been set.
func (o *ProjectNotificationSettings) HasNotifyOnCaseClosed() bool {
	return o != nil && o.NotifyOnCaseClosed != nil
}

// SetNotifyOnCaseClosed gets a reference to the given bool and assigns it to the NotifyOnCaseClosed field.
func (o *ProjectNotificationSettings) SetNotifyOnCaseClosed(v bool) {
	o.NotifyOnCaseClosed = &v
}

// GetNotifyOnCaseComment returns the NotifyOnCaseComment field value if set, zero value otherwise.
func (o *ProjectNotificationSettings) GetNotifyOnCaseComment() bool {
	if o == nil || o.NotifyOnCaseComment == nil {
		var ret bool
		return ret
	}
	return *o.NotifyOnCaseComment
}

// GetNotifyOnCaseCommentOk returns a tuple with the NotifyOnCaseComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectNotificationSettings) GetNotifyOnCaseCommentOk() (*bool, bool) {
	if o == nil || o.NotifyOnCaseComment == nil {
		return nil, false
	}
	return o.NotifyOnCaseComment, true
}

// HasNotifyOnCaseComment returns a boolean if a field has been set.
func (o *ProjectNotificationSettings) HasNotifyOnCaseComment() bool {
	return o != nil && o.NotifyOnCaseComment != nil
}

// SetNotifyOnCaseComment gets a reference to the given bool and assigns it to the NotifyOnCaseComment field.
func (o *ProjectNotificationSettings) SetNotifyOnCaseComment(v bool) {
	o.NotifyOnCaseComment = &v
}

// GetNotifyOnCaseCommentMention returns the NotifyOnCaseCommentMention field value if set, zero value otherwise.
func (o *ProjectNotificationSettings) GetNotifyOnCaseCommentMention() bool {
	if o == nil || o.NotifyOnCaseCommentMention == nil {
		var ret bool
		return ret
	}
	return *o.NotifyOnCaseCommentMention
}

// GetNotifyOnCaseCommentMentionOk returns a tuple with the NotifyOnCaseCommentMention field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectNotificationSettings) GetNotifyOnCaseCommentMentionOk() (*bool, bool) {
	if o == nil || o.NotifyOnCaseCommentMention == nil {
		return nil, false
	}
	return o.NotifyOnCaseCommentMention, true
}

// HasNotifyOnCaseCommentMention returns a boolean if a field has been set.
func (o *ProjectNotificationSettings) HasNotifyOnCaseCommentMention() bool {
	return o != nil && o.NotifyOnCaseCommentMention != nil
}

// SetNotifyOnCaseCommentMention gets a reference to the given bool and assigns it to the NotifyOnCaseCommentMention field.
func (o *ProjectNotificationSettings) SetNotifyOnCaseCommentMention(v bool) {
	o.NotifyOnCaseCommentMention = &v
}

// GetNotifyOnCasePriorityChange returns the NotifyOnCasePriorityChange field value if set, zero value otherwise.
func (o *ProjectNotificationSettings) GetNotifyOnCasePriorityChange() bool {
	if o == nil || o.NotifyOnCasePriorityChange == nil {
		var ret bool
		return ret
	}
	return *o.NotifyOnCasePriorityChange
}

// GetNotifyOnCasePriorityChangeOk returns a tuple with the NotifyOnCasePriorityChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectNotificationSettings) GetNotifyOnCasePriorityChangeOk() (*bool, bool) {
	if o == nil || o.NotifyOnCasePriorityChange == nil {
		return nil, false
	}
	return o.NotifyOnCasePriorityChange, true
}

// HasNotifyOnCasePriorityChange returns a boolean if a field has been set.
func (o *ProjectNotificationSettings) HasNotifyOnCasePriorityChange() bool {
	return o != nil && o.NotifyOnCasePriorityChange != nil
}

// SetNotifyOnCasePriorityChange gets a reference to the given bool and assigns it to the NotifyOnCasePriorityChange field.
func (o *ProjectNotificationSettings) SetNotifyOnCasePriorityChange(v bool) {
	o.NotifyOnCasePriorityChange = &v
}

// GetNotifyOnCaseStatusChange returns the NotifyOnCaseStatusChange field value if set, zero value otherwise.
func (o *ProjectNotificationSettings) GetNotifyOnCaseStatusChange() bool {
	if o == nil || o.NotifyOnCaseStatusChange == nil {
		var ret bool
		return ret
	}
	return *o.NotifyOnCaseStatusChange
}

// GetNotifyOnCaseStatusChangeOk returns a tuple with the NotifyOnCaseStatusChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectNotificationSettings) GetNotifyOnCaseStatusChangeOk() (*bool, bool) {
	if o == nil || o.NotifyOnCaseStatusChange == nil {
		return nil, false
	}
	return o.NotifyOnCaseStatusChange, true
}

// HasNotifyOnCaseStatusChange returns a boolean if a field has been set.
func (o *ProjectNotificationSettings) HasNotifyOnCaseStatusChange() bool {
	return o != nil && o.NotifyOnCaseStatusChange != nil
}

// SetNotifyOnCaseStatusChange gets a reference to the given bool and assigns it to the NotifyOnCaseStatusChange field.
func (o *ProjectNotificationSettings) SetNotifyOnCaseStatusChange(v bool) {
	o.NotifyOnCaseStatusChange = &v
}

// GetNotifyOnCaseUnassignment returns the NotifyOnCaseUnassignment field value if set, zero value otherwise.
func (o *ProjectNotificationSettings) GetNotifyOnCaseUnassignment() bool {
	if o == nil || o.NotifyOnCaseUnassignment == nil {
		var ret bool
		return ret
	}
	return *o.NotifyOnCaseUnassignment
}

// GetNotifyOnCaseUnassignmentOk returns a tuple with the NotifyOnCaseUnassignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectNotificationSettings) GetNotifyOnCaseUnassignmentOk() (*bool, bool) {
	if o == nil || o.NotifyOnCaseUnassignment == nil {
		return nil, false
	}
	return o.NotifyOnCaseUnassignment, true
}

// HasNotifyOnCaseUnassignment returns a boolean if a field has been set.
func (o *ProjectNotificationSettings) HasNotifyOnCaseUnassignment() bool {
	return o != nil && o.NotifyOnCaseUnassignment != nil
}

// SetNotifyOnCaseUnassignment gets a reference to the given bool and assigns it to the NotifyOnCaseUnassignment field.
func (o *ProjectNotificationSettings) SetNotifyOnCaseUnassignment(v bool) {
	o.NotifyOnCaseUnassignment = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProjectNotificationSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Destinations != nil {
		toSerialize["destinations"] = o.Destinations
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.NotifyOnCaseAssignment != nil {
		toSerialize["notify_on_case_assignment"] = o.NotifyOnCaseAssignment
	}
	if o.NotifyOnCaseClosed != nil {
		toSerialize["notify_on_case_closed"] = o.NotifyOnCaseClosed
	}
	if o.NotifyOnCaseComment != nil {
		toSerialize["notify_on_case_comment"] = o.NotifyOnCaseComment
	}
	if o.NotifyOnCaseCommentMention != nil {
		toSerialize["notify_on_case_comment_mention"] = o.NotifyOnCaseCommentMention
	}
	if o.NotifyOnCasePriorityChange != nil {
		toSerialize["notify_on_case_priority_change"] = o.NotifyOnCasePriorityChange
	}
	if o.NotifyOnCaseStatusChange != nil {
		toSerialize["notify_on_case_status_change"] = o.NotifyOnCaseStatusChange
	}
	if o.NotifyOnCaseUnassignment != nil {
		toSerialize["notify_on_case_unassignment"] = o.NotifyOnCaseUnassignment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProjectNotificationSettings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Destinations               []int32 `json:"destinations,omitempty"`
		Enabled                    *bool   `json:"enabled,omitempty"`
		NotifyOnCaseAssignment     *bool   `json:"notify_on_case_assignment,omitempty"`
		NotifyOnCaseClosed         *bool   `json:"notify_on_case_closed,omitempty"`
		NotifyOnCaseComment        *bool   `json:"notify_on_case_comment,omitempty"`
		NotifyOnCaseCommentMention *bool   `json:"notify_on_case_comment_mention,omitempty"`
		NotifyOnCasePriorityChange *bool   `json:"notify_on_case_priority_change,omitempty"`
		NotifyOnCaseStatusChange   *bool   `json:"notify_on_case_status_change,omitempty"`
		NotifyOnCaseUnassignment   *bool   `json:"notify_on_case_unassignment,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"destinations", "enabled", "notify_on_case_assignment", "notify_on_case_closed", "notify_on_case_comment", "notify_on_case_comment_mention", "notify_on_case_priority_change", "notify_on_case_status_change", "notify_on_case_unassignment"})
	} else {
		return err
	}
	o.Destinations = all.Destinations
	o.Enabled = all.Enabled
	o.NotifyOnCaseAssignment = all.NotifyOnCaseAssignment
	o.NotifyOnCaseClosed = all.NotifyOnCaseClosed
	o.NotifyOnCaseComment = all.NotifyOnCaseComment
	o.NotifyOnCaseCommentMention = all.NotifyOnCaseCommentMention
	o.NotifyOnCasePriorityChange = all.NotifyOnCasePriorityChange
	o.NotifyOnCaseStatusChange = all.NotifyOnCaseStatusChange
	o.NotifyOnCaseUnassignment = all.NotifyOnCaseUnassignment

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
