// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TriggerAttributes The trigger definition for starting an investigation.
type TriggerAttributes struct {
	// Attributes for a monitor alert trigger.
	MonitorAlertTrigger MonitorAlertTriggerAttributes `json:"monitor_alert_trigger"`
	// The type of trigger for the investigation.
	Type TriggerType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTriggerAttributes instantiates a new TriggerAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTriggerAttributes(monitorAlertTrigger MonitorAlertTriggerAttributes, typeVar TriggerType) *TriggerAttributes {
	this := TriggerAttributes{}
	this.MonitorAlertTrigger = monitorAlertTrigger
	this.Type = typeVar
	return &this
}

// NewTriggerAttributesWithDefaults instantiates a new TriggerAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTriggerAttributesWithDefaults() *TriggerAttributes {
	this := TriggerAttributes{}
	return &this
}

// GetMonitorAlertTrigger returns the MonitorAlertTrigger field value.
func (o *TriggerAttributes) GetMonitorAlertTrigger() MonitorAlertTriggerAttributes {
	if o == nil {
		var ret MonitorAlertTriggerAttributes
		return ret
	}
	return o.MonitorAlertTrigger
}

// GetMonitorAlertTriggerOk returns a tuple with the MonitorAlertTrigger field value
// and a boolean to check if the value has been set.
func (o *TriggerAttributes) GetMonitorAlertTriggerOk() (*MonitorAlertTriggerAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitorAlertTrigger, true
}

// SetMonitorAlertTrigger sets field value.
func (o *TriggerAttributes) SetMonitorAlertTrigger(v MonitorAlertTriggerAttributes) {
	o.MonitorAlertTrigger = v
}

// GetType returns the Type field value.
func (o *TriggerAttributes) GetType() TriggerType {
	if o == nil {
		var ret TriggerType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TriggerAttributes) GetTypeOk() (*TriggerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *TriggerAttributes) SetType(v TriggerType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TriggerAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["monitor_alert_trigger"] = o.MonitorAlertTrigger
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TriggerAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MonitorAlertTrigger *MonitorAlertTriggerAttributes `json:"monitor_alert_trigger"`
		Type                *TriggerType                   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.MonitorAlertTrigger == nil {
		return fmt.Errorf("required field monitor_alert_trigger missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"monitor_alert_trigger", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.MonitorAlertTrigger.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MonitorAlertTrigger = *all.MonitorAlertTrigger
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
