// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GovernanceControlUpdateAttributes The attributes of a governance control that can be updated. Only the attributes present in the request are modified.
type GovernanceControlUpdateAttributes struct {
	// How often detections should be evaluated for the control.
	DetectionFrequency *string `json:"detection_frequency,omitempty"`
	// A free-form map of parameter names to their configured values.
	DetectionParameters map[string]interface{} `json:"detection_parameters,omitempty"`
	// A free-form map of parameter names to their configured values.
	MitigationParameters map[string]interface{} `json:"mitigation_parameters,omitempty"`
	// The mitigation type to configure for the control.
	MitigationType *string `json:"mitigation_type,omitempty"`
	// A new human-readable name for the control.
	Name *string `json:"name,omitempty"`
	// The notification frequency to configure for the control.
	NotificationFrequency *string `json:"notification_frequency,omitempty"`
	// A free-form map of parameter names to their configured values.
	NotificationParameters map[string]interface{} `json:"notification_parameters,omitempty"`
	// The notification type to configure for the control.
	NotificationType *string `json:"notification_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGovernanceControlUpdateAttributes instantiates a new GovernanceControlUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGovernanceControlUpdateAttributes() *GovernanceControlUpdateAttributes {
	this := GovernanceControlUpdateAttributes{}
	return &this
}

// NewGovernanceControlUpdateAttributesWithDefaults instantiates a new GovernanceControlUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGovernanceControlUpdateAttributesWithDefaults() *GovernanceControlUpdateAttributes {
	this := GovernanceControlUpdateAttributes{}
	return &this
}

// GetDetectionFrequency returns the DetectionFrequency field value if set, zero value otherwise.
func (o *GovernanceControlUpdateAttributes) GetDetectionFrequency() string {
	if o == nil || o.DetectionFrequency == nil {
		var ret string
		return ret
	}
	return *o.DetectionFrequency
}

// GetDetectionFrequencyOk returns a tuple with the DetectionFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlUpdateAttributes) GetDetectionFrequencyOk() (*string, bool) {
	if o == nil || o.DetectionFrequency == nil {
		return nil, false
	}
	return o.DetectionFrequency, true
}

// HasDetectionFrequency returns a boolean if a field has been set.
func (o *GovernanceControlUpdateAttributes) HasDetectionFrequency() bool {
	return o != nil && o.DetectionFrequency != nil
}

// SetDetectionFrequency gets a reference to the given string and assigns it to the DetectionFrequency field.
func (o *GovernanceControlUpdateAttributes) SetDetectionFrequency(v string) {
	o.DetectionFrequency = &v
}

// GetDetectionParameters returns the DetectionParameters field value if set, zero value otherwise.
func (o *GovernanceControlUpdateAttributes) GetDetectionParameters() map[string]interface{} {
	if o == nil || o.DetectionParameters == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.DetectionParameters
}

// GetDetectionParametersOk returns a tuple with the DetectionParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlUpdateAttributes) GetDetectionParametersOk() (*map[string]interface{}, bool) {
	if o == nil || o.DetectionParameters == nil {
		return nil, false
	}
	return &o.DetectionParameters, true
}

// HasDetectionParameters returns a boolean if a field has been set.
func (o *GovernanceControlUpdateAttributes) HasDetectionParameters() bool {
	return o != nil && o.DetectionParameters != nil
}

// SetDetectionParameters gets a reference to the given map[string]interface{} and assigns it to the DetectionParameters field.
func (o *GovernanceControlUpdateAttributes) SetDetectionParameters(v map[string]interface{}) {
	o.DetectionParameters = v
}

// GetMitigationParameters returns the MitigationParameters field value if set, zero value otherwise.
func (o *GovernanceControlUpdateAttributes) GetMitigationParameters() map[string]interface{} {
	if o == nil || o.MitigationParameters == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.MitigationParameters
}

// GetMitigationParametersOk returns a tuple with the MitigationParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlUpdateAttributes) GetMitigationParametersOk() (*map[string]interface{}, bool) {
	if o == nil || o.MitigationParameters == nil {
		return nil, false
	}
	return &o.MitigationParameters, true
}

// HasMitigationParameters returns a boolean if a field has been set.
func (o *GovernanceControlUpdateAttributes) HasMitigationParameters() bool {
	return o != nil && o.MitigationParameters != nil
}

// SetMitigationParameters gets a reference to the given map[string]interface{} and assigns it to the MitigationParameters field.
func (o *GovernanceControlUpdateAttributes) SetMitigationParameters(v map[string]interface{}) {
	o.MitigationParameters = v
}

// GetMitigationType returns the MitigationType field value if set, zero value otherwise.
func (o *GovernanceControlUpdateAttributes) GetMitigationType() string {
	if o == nil || o.MitigationType == nil {
		var ret string
		return ret
	}
	return *o.MitigationType
}

// GetMitigationTypeOk returns a tuple with the MitigationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlUpdateAttributes) GetMitigationTypeOk() (*string, bool) {
	if o == nil || o.MitigationType == nil {
		return nil, false
	}
	return o.MitigationType, true
}

// HasMitigationType returns a boolean if a field has been set.
func (o *GovernanceControlUpdateAttributes) HasMitigationType() bool {
	return o != nil && o.MitigationType != nil
}

// SetMitigationType gets a reference to the given string and assigns it to the MitigationType field.
func (o *GovernanceControlUpdateAttributes) SetMitigationType(v string) {
	o.MitigationType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GovernanceControlUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GovernanceControlUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GovernanceControlUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetNotificationFrequency returns the NotificationFrequency field value if set, zero value otherwise.
func (o *GovernanceControlUpdateAttributes) GetNotificationFrequency() string {
	if o == nil || o.NotificationFrequency == nil {
		var ret string
		return ret
	}
	return *o.NotificationFrequency
}

// GetNotificationFrequencyOk returns a tuple with the NotificationFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlUpdateAttributes) GetNotificationFrequencyOk() (*string, bool) {
	if o == nil || o.NotificationFrequency == nil {
		return nil, false
	}
	return o.NotificationFrequency, true
}

// HasNotificationFrequency returns a boolean if a field has been set.
func (o *GovernanceControlUpdateAttributes) HasNotificationFrequency() bool {
	return o != nil && o.NotificationFrequency != nil
}

// SetNotificationFrequency gets a reference to the given string and assigns it to the NotificationFrequency field.
func (o *GovernanceControlUpdateAttributes) SetNotificationFrequency(v string) {
	o.NotificationFrequency = &v
}

// GetNotificationParameters returns the NotificationParameters field value if set, zero value otherwise.
func (o *GovernanceControlUpdateAttributes) GetNotificationParameters() map[string]interface{} {
	if o == nil || o.NotificationParameters == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.NotificationParameters
}

// GetNotificationParametersOk returns a tuple with the NotificationParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlUpdateAttributes) GetNotificationParametersOk() (*map[string]interface{}, bool) {
	if o == nil || o.NotificationParameters == nil {
		return nil, false
	}
	return &o.NotificationParameters, true
}

// HasNotificationParameters returns a boolean if a field has been set.
func (o *GovernanceControlUpdateAttributes) HasNotificationParameters() bool {
	return o != nil && o.NotificationParameters != nil
}

// SetNotificationParameters gets a reference to the given map[string]interface{} and assigns it to the NotificationParameters field.
func (o *GovernanceControlUpdateAttributes) SetNotificationParameters(v map[string]interface{}) {
	o.NotificationParameters = v
}

// GetNotificationType returns the NotificationType field value if set, zero value otherwise.
func (o *GovernanceControlUpdateAttributes) GetNotificationType() string {
	if o == nil || o.NotificationType == nil {
		var ret string
		return ret
	}
	return *o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GovernanceControlUpdateAttributes) GetNotificationTypeOk() (*string, bool) {
	if o == nil || o.NotificationType == nil {
		return nil, false
	}
	return o.NotificationType, true
}

// HasNotificationType returns a boolean if a field has been set.
func (o *GovernanceControlUpdateAttributes) HasNotificationType() bool {
	return o != nil && o.NotificationType != nil
}

// SetNotificationType gets a reference to the given string and assigns it to the NotificationType field.
func (o *GovernanceControlUpdateAttributes) SetNotificationType(v string) {
	o.NotificationType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GovernanceControlUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DetectionFrequency != nil {
		toSerialize["detection_frequency"] = o.DetectionFrequency
	}
	if o.DetectionParameters != nil {
		toSerialize["detection_parameters"] = o.DetectionParameters
	}
	if o.MitigationParameters != nil {
		toSerialize["mitigation_parameters"] = o.MitigationParameters
	}
	if o.MitigationType != nil {
		toSerialize["mitigation_type"] = o.MitigationType
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NotificationFrequency != nil {
		toSerialize["notification_frequency"] = o.NotificationFrequency
	}
	if o.NotificationParameters != nil {
		toSerialize["notification_parameters"] = o.NotificationParameters
	}
	if o.NotificationType != nil {
		toSerialize["notification_type"] = o.NotificationType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GovernanceControlUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DetectionFrequency     *string                `json:"detection_frequency,omitempty"`
		DetectionParameters    map[string]interface{} `json:"detection_parameters,omitempty"`
		MitigationParameters   map[string]interface{} `json:"mitigation_parameters,omitempty"`
		MitigationType         *string                `json:"mitigation_type,omitempty"`
		Name                   *string                `json:"name,omitempty"`
		NotificationFrequency  *string                `json:"notification_frequency,omitempty"`
		NotificationParameters map[string]interface{} `json:"notification_parameters,omitempty"`
		NotificationType       *string                `json:"notification_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"detection_frequency", "detection_parameters", "mitigation_parameters", "mitigation_type", "name", "notification_frequency", "notification_parameters", "notification_type"})
	} else {
		return err
	}
	o.DetectionFrequency = all.DetectionFrequency
	o.DetectionParameters = all.DetectionParameters
	o.MitigationParameters = all.MitigationParameters
	o.MitigationType = all.MitigationType
	o.Name = all.Name
	o.NotificationFrequency = all.NotificationFrequency
	o.NotificationParameters = all.NotificationParameters
	o.NotificationType = all.NotificationType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
