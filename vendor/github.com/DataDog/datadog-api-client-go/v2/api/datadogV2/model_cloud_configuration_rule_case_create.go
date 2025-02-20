// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudConfigurationRuleCaseCreate Description of signals.
type CloudConfigurationRuleCaseCreate struct {
	// Notification targets for each rule case.
	Notifications []string `json:"notifications,omitempty"`
	// Severity of the Security Signal.
	Status SecurityMonitoringRuleSeverity `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudConfigurationRuleCaseCreate instantiates a new CloudConfigurationRuleCaseCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudConfigurationRuleCaseCreate(status SecurityMonitoringRuleSeverity) *CloudConfigurationRuleCaseCreate {
	this := CloudConfigurationRuleCaseCreate{}
	this.Status = status
	return &this
}

// NewCloudConfigurationRuleCaseCreateWithDefaults instantiates a new CloudConfigurationRuleCaseCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudConfigurationRuleCaseCreateWithDefaults() *CloudConfigurationRuleCaseCreate {
	this := CloudConfigurationRuleCaseCreate{}
	return &this
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *CloudConfigurationRuleCaseCreate) GetNotifications() []string {
	if o == nil || o.Notifications == nil {
		var ret []string
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudConfigurationRuleCaseCreate) GetNotificationsOk() (*[]string, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return &o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *CloudConfigurationRuleCaseCreate) HasNotifications() bool {
	return o != nil && o.Notifications != nil
}

// SetNotifications gets a reference to the given []string and assigns it to the Notifications field.
func (o *CloudConfigurationRuleCaseCreate) SetNotifications(v []string) {
	o.Notifications = v
}

// GetStatus returns the Status field value.
func (o *CloudConfigurationRuleCaseCreate) GetStatus() SecurityMonitoringRuleSeverity {
	if o == nil {
		var ret SecurityMonitoringRuleSeverity
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CloudConfigurationRuleCaseCreate) GetStatusOk() (*SecurityMonitoringRuleSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *CloudConfigurationRuleCaseCreate) SetStatus(v SecurityMonitoringRuleSeverity) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudConfigurationRuleCaseCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudConfigurationRuleCaseCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Notifications []string                        `json:"notifications,omitempty"`
		Status        *SecurityMonitoringRuleSeverity `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"notifications", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Notifications = all.Notifications
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
