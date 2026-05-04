// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseNotificationRuleTrigger Notification rule trigger
type CaseNotificationRuleTrigger struct {
	// Trigger data
	Data *CaseNotificationRuleTriggerData `json:"data,omitempty"`
	// Type of trigger (CASE_CREATED, STATUS_TRANSITIONED, ATTRIBUTE_VALUE_CHANGED, EVENT_CORRELATION_SIGNAL_CORRELATED)
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseNotificationRuleTrigger instantiates a new CaseNotificationRuleTrigger object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseNotificationRuleTrigger() *CaseNotificationRuleTrigger {
	this := CaseNotificationRuleTrigger{}
	return &this
}

// NewCaseNotificationRuleTriggerWithDefaults instantiates a new CaseNotificationRuleTrigger object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseNotificationRuleTriggerWithDefaults() *CaseNotificationRuleTrigger {
	this := CaseNotificationRuleTrigger{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CaseNotificationRuleTrigger) GetData() CaseNotificationRuleTriggerData {
	if o == nil || o.Data == nil {
		var ret CaseNotificationRuleTriggerData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleTrigger) GetDataOk() (*CaseNotificationRuleTriggerData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CaseNotificationRuleTrigger) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given CaseNotificationRuleTriggerData and assigns it to the Data field.
func (o *CaseNotificationRuleTrigger) SetData(v CaseNotificationRuleTriggerData) {
	o.Data = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CaseNotificationRuleTrigger) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleTrigger) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CaseNotificationRuleTrigger) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CaseNotificationRuleTrigger) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseNotificationRuleTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseNotificationRuleTrigger) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *CaseNotificationRuleTriggerData `json:"data,omitempty"`
		Type *string                          `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
