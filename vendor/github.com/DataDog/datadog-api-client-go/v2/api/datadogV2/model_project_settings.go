// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProjectSettings Project settings.
type ProjectSettings struct {
	// Auto-close inactive cases settings.
	AutoCloseInactiveCases *AutoCloseInactiveCases `json:"auto_close_inactive_cases,omitempty"`
	// Auto-transition assigned cases settings.
	AutoTransitionAssignedCases *AutoTransitionAssignedCases `json:"auto_transition_assigned_cases,omitempty"`
	// Incident integration settings.
	IntegrationIncident *IntegrationIncident `json:"integration_incident,omitempty"`
	// Jira integration settings.
	IntegrationJira *IntegrationJira `json:"integration_jira,omitempty"`
	// Monitor integration settings.
	IntegrationMonitor *IntegrationMonitor `json:"integration_monitor,omitempty"`
	// On-Call integration settings.
	IntegrationOnCall *IntegrationOnCall `json:"integration_on_call,omitempty"`
	// ServiceNow integration settings.
	IntegrationServiceNow *IntegrationServiceNow `json:"integration_service_now,omitempty"`
	// Project notification settings.
	Notification *ProjectNotificationSettings `json:"notification,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProjectSettings instantiates a new ProjectSettings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProjectSettings() *ProjectSettings {
	this := ProjectSettings{}
	return &this
}

// NewProjectSettingsWithDefaults instantiates a new ProjectSettings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProjectSettingsWithDefaults() *ProjectSettings {
	this := ProjectSettings{}
	return &this
}

// GetAutoCloseInactiveCases returns the AutoCloseInactiveCases field value if set, zero value otherwise.
func (o *ProjectSettings) GetAutoCloseInactiveCases() AutoCloseInactiveCases {
	if o == nil || o.AutoCloseInactiveCases == nil {
		var ret AutoCloseInactiveCases
		return ret
	}
	return *o.AutoCloseInactiveCases
}

// GetAutoCloseInactiveCasesOk returns a tuple with the AutoCloseInactiveCases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSettings) GetAutoCloseInactiveCasesOk() (*AutoCloseInactiveCases, bool) {
	if o == nil || o.AutoCloseInactiveCases == nil {
		return nil, false
	}
	return o.AutoCloseInactiveCases, true
}

// HasAutoCloseInactiveCases returns a boolean if a field has been set.
func (o *ProjectSettings) HasAutoCloseInactiveCases() bool {
	return o != nil && o.AutoCloseInactiveCases != nil
}

// SetAutoCloseInactiveCases gets a reference to the given AutoCloseInactiveCases and assigns it to the AutoCloseInactiveCases field.
func (o *ProjectSettings) SetAutoCloseInactiveCases(v AutoCloseInactiveCases) {
	o.AutoCloseInactiveCases = &v
}

// GetAutoTransitionAssignedCases returns the AutoTransitionAssignedCases field value if set, zero value otherwise.
func (o *ProjectSettings) GetAutoTransitionAssignedCases() AutoTransitionAssignedCases {
	if o == nil || o.AutoTransitionAssignedCases == nil {
		var ret AutoTransitionAssignedCases
		return ret
	}
	return *o.AutoTransitionAssignedCases
}

// GetAutoTransitionAssignedCasesOk returns a tuple with the AutoTransitionAssignedCases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSettings) GetAutoTransitionAssignedCasesOk() (*AutoTransitionAssignedCases, bool) {
	if o == nil || o.AutoTransitionAssignedCases == nil {
		return nil, false
	}
	return o.AutoTransitionAssignedCases, true
}

// HasAutoTransitionAssignedCases returns a boolean if a field has been set.
func (o *ProjectSettings) HasAutoTransitionAssignedCases() bool {
	return o != nil && o.AutoTransitionAssignedCases != nil
}

// SetAutoTransitionAssignedCases gets a reference to the given AutoTransitionAssignedCases and assigns it to the AutoTransitionAssignedCases field.
func (o *ProjectSettings) SetAutoTransitionAssignedCases(v AutoTransitionAssignedCases) {
	o.AutoTransitionAssignedCases = &v
}

// GetIntegrationIncident returns the IntegrationIncident field value if set, zero value otherwise.
func (o *ProjectSettings) GetIntegrationIncident() IntegrationIncident {
	if o == nil || o.IntegrationIncident == nil {
		var ret IntegrationIncident
		return ret
	}
	return *o.IntegrationIncident
}

// GetIntegrationIncidentOk returns a tuple with the IntegrationIncident field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSettings) GetIntegrationIncidentOk() (*IntegrationIncident, bool) {
	if o == nil || o.IntegrationIncident == nil {
		return nil, false
	}
	return o.IntegrationIncident, true
}

// HasIntegrationIncident returns a boolean if a field has been set.
func (o *ProjectSettings) HasIntegrationIncident() bool {
	return o != nil && o.IntegrationIncident != nil
}

// SetIntegrationIncident gets a reference to the given IntegrationIncident and assigns it to the IntegrationIncident field.
func (o *ProjectSettings) SetIntegrationIncident(v IntegrationIncident) {
	o.IntegrationIncident = &v
}

// GetIntegrationJira returns the IntegrationJira field value if set, zero value otherwise.
func (o *ProjectSettings) GetIntegrationJira() IntegrationJira {
	if o == nil || o.IntegrationJira == nil {
		var ret IntegrationJira
		return ret
	}
	return *o.IntegrationJira
}

// GetIntegrationJiraOk returns a tuple with the IntegrationJira field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSettings) GetIntegrationJiraOk() (*IntegrationJira, bool) {
	if o == nil || o.IntegrationJira == nil {
		return nil, false
	}
	return o.IntegrationJira, true
}

// HasIntegrationJira returns a boolean if a field has been set.
func (o *ProjectSettings) HasIntegrationJira() bool {
	return o != nil && o.IntegrationJira != nil
}

// SetIntegrationJira gets a reference to the given IntegrationJira and assigns it to the IntegrationJira field.
func (o *ProjectSettings) SetIntegrationJira(v IntegrationJira) {
	o.IntegrationJira = &v
}

// GetIntegrationMonitor returns the IntegrationMonitor field value if set, zero value otherwise.
func (o *ProjectSettings) GetIntegrationMonitor() IntegrationMonitor {
	if o == nil || o.IntegrationMonitor == nil {
		var ret IntegrationMonitor
		return ret
	}
	return *o.IntegrationMonitor
}

// GetIntegrationMonitorOk returns a tuple with the IntegrationMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSettings) GetIntegrationMonitorOk() (*IntegrationMonitor, bool) {
	if o == nil || o.IntegrationMonitor == nil {
		return nil, false
	}
	return o.IntegrationMonitor, true
}

// HasIntegrationMonitor returns a boolean if a field has been set.
func (o *ProjectSettings) HasIntegrationMonitor() bool {
	return o != nil && o.IntegrationMonitor != nil
}

// SetIntegrationMonitor gets a reference to the given IntegrationMonitor and assigns it to the IntegrationMonitor field.
func (o *ProjectSettings) SetIntegrationMonitor(v IntegrationMonitor) {
	o.IntegrationMonitor = &v
}

// GetIntegrationOnCall returns the IntegrationOnCall field value if set, zero value otherwise.
func (o *ProjectSettings) GetIntegrationOnCall() IntegrationOnCall {
	if o == nil || o.IntegrationOnCall == nil {
		var ret IntegrationOnCall
		return ret
	}
	return *o.IntegrationOnCall
}

// GetIntegrationOnCallOk returns a tuple with the IntegrationOnCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSettings) GetIntegrationOnCallOk() (*IntegrationOnCall, bool) {
	if o == nil || o.IntegrationOnCall == nil {
		return nil, false
	}
	return o.IntegrationOnCall, true
}

// HasIntegrationOnCall returns a boolean if a field has been set.
func (o *ProjectSettings) HasIntegrationOnCall() bool {
	return o != nil && o.IntegrationOnCall != nil
}

// SetIntegrationOnCall gets a reference to the given IntegrationOnCall and assigns it to the IntegrationOnCall field.
func (o *ProjectSettings) SetIntegrationOnCall(v IntegrationOnCall) {
	o.IntegrationOnCall = &v
}

// GetIntegrationServiceNow returns the IntegrationServiceNow field value if set, zero value otherwise.
func (o *ProjectSettings) GetIntegrationServiceNow() IntegrationServiceNow {
	if o == nil || o.IntegrationServiceNow == nil {
		var ret IntegrationServiceNow
		return ret
	}
	return *o.IntegrationServiceNow
}

// GetIntegrationServiceNowOk returns a tuple with the IntegrationServiceNow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSettings) GetIntegrationServiceNowOk() (*IntegrationServiceNow, bool) {
	if o == nil || o.IntegrationServiceNow == nil {
		return nil, false
	}
	return o.IntegrationServiceNow, true
}

// HasIntegrationServiceNow returns a boolean if a field has been set.
func (o *ProjectSettings) HasIntegrationServiceNow() bool {
	return o != nil && o.IntegrationServiceNow != nil
}

// SetIntegrationServiceNow gets a reference to the given IntegrationServiceNow and assigns it to the IntegrationServiceNow field.
func (o *ProjectSettings) SetIntegrationServiceNow(v IntegrationServiceNow) {
	o.IntegrationServiceNow = &v
}

// GetNotification returns the Notification field value if set, zero value otherwise.
func (o *ProjectSettings) GetNotification() ProjectNotificationSettings {
	if o == nil || o.Notification == nil {
		var ret ProjectNotificationSettings
		return ret
	}
	return *o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectSettings) GetNotificationOk() (*ProjectNotificationSettings, bool) {
	if o == nil || o.Notification == nil {
		return nil, false
	}
	return o.Notification, true
}

// HasNotification returns a boolean if a field has been set.
func (o *ProjectSettings) HasNotification() bool {
	return o != nil && o.Notification != nil
}

// SetNotification gets a reference to the given ProjectNotificationSettings and assigns it to the Notification field.
func (o *ProjectSettings) SetNotification(v ProjectNotificationSettings) {
	o.Notification = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProjectSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutoCloseInactiveCases != nil {
		toSerialize["auto_close_inactive_cases"] = o.AutoCloseInactiveCases
	}
	if o.AutoTransitionAssignedCases != nil {
		toSerialize["auto_transition_assigned_cases"] = o.AutoTransitionAssignedCases
	}
	if o.IntegrationIncident != nil {
		toSerialize["integration_incident"] = o.IntegrationIncident
	}
	if o.IntegrationJira != nil {
		toSerialize["integration_jira"] = o.IntegrationJira
	}
	if o.IntegrationMonitor != nil {
		toSerialize["integration_monitor"] = o.IntegrationMonitor
	}
	if o.IntegrationOnCall != nil {
		toSerialize["integration_on_call"] = o.IntegrationOnCall
	}
	if o.IntegrationServiceNow != nil {
		toSerialize["integration_service_now"] = o.IntegrationServiceNow
	}
	if o.Notification != nil {
		toSerialize["notification"] = o.Notification
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProjectSettings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoCloseInactiveCases      *AutoCloseInactiveCases      `json:"auto_close_inactive_cases,omitempty"`
		AutoTransitionAssignedCases *AutoTransitionAssignedCases `json:"auto_transition_assigned_cases,omitempty"`
		IntegrationIncident         *IntegrationIncident         `json:"integration_incident,omitempty"`
		IntegrationJira             *IntegrationJira             `json:"integration_jira,omitempty"`
		IntegrationMonitor          *IntegrationMonitor          `json:"integration_monitor,omitempty"`
		IntegrationOnCall           *IntegrationOnCall           `json:"integration_on_call,omitempty"`
		IntegrationServiceNow       *IntegrationServiceNow       `json:"integration_service_now,omitempty"`
		Notification                *ProjectNotificationSettings `json:"notification,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_close_inactive_cases", "auto_transition_assigned_cases", "integration_incident", "integration_jira", "integration_monitor", "integration_on_call", "integration_service_now", "notification"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AutoCloseInactiveCases != nil && all.AutoCloseInactiveCases.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AutoCloseInactiveCases = all.AutoCloseInactiveCases
	if all.AutoTransitionAssignedCases != nil && all.AutoTransitionAssignedCases.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AutoTransitionAssignedCases = all.AutoTransitionAssignedCases
	if all.IntegrationIncident != nil && all.IntegrationIncident.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IntegrationIncident = all.IntegrationIncident
	if all.IntegrationJira != nil && all.IntegrationJira.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IntegrationJira = all.IntegrationJira
	if all.IntegrationMonitor != nil && all.IntegrationMonitor.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IntegrationMonitor = all.IntegrationMonitor
	if all.IntegrationOnCall != nil && all.IntegrationOnCall.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IntegrationOnCall = all.IntegrationOnCall
	if all.IntegrationServiceNow != nil && all.IntegrationServiceNow.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.IntegrationServiceNow = all.IntegrationServiceNow
	if all.Notification != nil && all.Notification.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Notification = all.Notification

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
