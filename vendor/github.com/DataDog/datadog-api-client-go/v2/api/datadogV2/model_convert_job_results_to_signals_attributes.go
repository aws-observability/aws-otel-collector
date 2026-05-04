// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConvertJobResultsToSignalsAttributes Attributes for converting historical job results to signals.
type ConvertJobResultsToSignalsAttributes struct {
	// Request ID.
	Id *string `json:"id,omitempty"`
	// Job result IDs.
	JobResultIds []string `json:"jobResultIds"`
	// Notifications sent.
	Notifications []string `json:"notifications"`
	// Message of generated signals.
	SignalMessage string `json:"signalMessage"`
	// Severity of the Security Signal.
	SignalSeverity SecurityMonitoringRuleSeverity `json:"signalSeverity"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConvertJobResultsToSignalsAttributes instantiates a new ConvertJobResultsToSignalsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConvertJobResultsToSignalsAttributes(jobResultIds []string, notifications []string, signalMessage string, signalSeverity SecurityMonitoringRuleSeverity) *ConvertJobResultsToSignalsAttributes {
	this := ConvertJobResultsToSignalsAttributes{}
	this.JobResultIds = jobResultIds
	this.Notifications = notifications
	this.SignalMessage = signalMessage
	this.SignalSeverity = signalSeverity
	return &this
}

// NewConvertJobResultsToSignalsAttributesWithDefaults instantiates a new ConvertJobResultsToSignalsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConvertJobResultsToSignalsAttributesWithDefaults() *ConvertJobResultsToSignalsAttributes {
	this := ConvertJobResultsToSignalsAttributes{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConvertJobResultsToSignalsAttributes) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvertJobResultsToSignalsAttributes) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConvertJobResultsToSignalsAttributes) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConvertJobResultsToSignalsAttributes) SetId(v string) {
	o.Id = &v
}

// GetJobResultIds returns the JobResultIds field value.
func (o *ConvertJobResultsToSignalsAttributes) GetJobResultIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.JobResultIds
}

// GetJobResultIdsOk returns a tuple with the JobResultIds field value
// and a boolean to check if the value has been set.
func (o *ConvertJobResultsToSignalsAttributes) GetJobResultIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobResultIds, true
}

// SetJobResultIds sets field value.
func (o *ConvertJobResultsToSignalsAttributes) SetJobResultIds(v []string) {
	o.JobResultIds = v
}

// GetNotifications returns the Notifications field value.
func (o *ConvertJobResultsToSignalsAttributes) GetNotifications() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value
// and a boolean to check if the value has been set.
func (o *ConvertJobResultsToSignalsAttributes) GetNotificationsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notifications, true
}

// SetNotifications sets field value.
func (o *ConvertJobResultsToSignalsAttributes) SetNotifications(v []string) {
	o.Notifications = v
}

// GetSignalMessage returns the SignalMessage field value.
func (o *ConvertJobResultsToSignalsAttributes) GetSignalMessage() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SignalMessage
}

// GetSignalMessageOk returns a tuple with the SignalMessage field value
// and a boolean to check if the value has been set.
func (o *ConvertJobResultsToSignalsAttributes) GetSignalMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalMessage, true
}

// SetSignalMessage sets field value.
func (o *ConvertJobResultsToSignalsAttributes) SetSignalMessage(v string) {
	o.SignalMessage = v
}

// GetSignalSeverity returns the SignalSeverity field value.
func (o *ConvertJobResultsToSignalsAttributes) GetSignalSeverity() SecurityMonitoringRuleSeverity {
	if o == nil {
		var ret SecurityMonitoringRuleSeverity
		return ret
	}
	return o.SignalSeverity
}

// GetSignalSeverityOk returns a tuple with the SignalSeverity field value
// and a boolean to check if the value has been set.
func (o *ConvertJobResultsToSignalsAttributes) GetSignalSeverityOk() (*SecurityMonitoringRuleSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalSeverity, true
}

// SetSignalSeverity sets field value.
func (o *ConvertJobResultsToSignalsAttributes) SetSignalSeverity(v SecurityMonitoringRuleSeverity) {
	o.SignalSeverity = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConvertJobResultsToSignalsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["jobResultIds"] = o.JobResultIds
	toSerialize["notifications"] = o.Notifications
	toSerialize["signalMessage"] = o.SignalMessage
	toSerialize["signalSeverity"] = o.SignalSeverity

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConvertJobResultsToSignalsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id             *string                         `json:"id,omitempty"`
		JobResultIds   *[]string                       `json:"jobResultIds"`
		Notifications  *[]string                       `json:"notifications"`
		SignalMessage  *string                         `json:"signalMessage"`
		SignalSeverity *SecurityMonitoringRuleSeverity `json:"signalSeverity"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JobResultIds == nil {
		return fmt.Errorf("required field jobResultIds missing")
	}
	if all.Notifications == nil {
		return fmt.Errorf("required field notifications missing")
	}
	if all.SignalMessage == nil {
		return fmt.Errorf("required field signalMessage missing")
	}
	if all.SignalSeverity == nil {
		return fmt.Errorf("required field signalSeverity missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "jobResultIds", "notifications", "signalMessage", "signalSeverity"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.JobResultIds = *all.JobResultIds
	o.Notifications = *all.Notifications
	o.SignalMessage = *all.SignalMessage
	if !all.SignalSeverity.IsValid() {
		hasInvalidField = true
	} else {
		o.SignalSeverity = *all.SignalSeverity
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
