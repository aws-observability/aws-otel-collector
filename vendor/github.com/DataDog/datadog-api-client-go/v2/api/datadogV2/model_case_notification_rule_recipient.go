// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseNotificationRuleRecipient Notification rule recipient
type CaseNotificationRuleRecipient struct {
	// Recipient data
	Data *CaseNotificationRuleRecipientData `json:"data,omitempty"`
	// Type of recipient (SLACK_CHANNEL, EMAIL, HTTP, PAGERDUTY_SERVICE, MS_TEAMS_CHANNEL)
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseNotificationRuleRecipient instantiates a new CaseNotificationRuleRecipient object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseNotificationRuleRecipient() *CaseNotificationRuleRecipient {
	this := CaseNotificationRuleRecipient{}
	return &this
}

// NewCaseNotificationRuleRecipientWithDefaults instantiates a new CaseNotificationRuleRecipient object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseNotificationRuleRecipientWithDefaults() *CaseNotificationRuleRecipient {
	this := CaseNotificationRuleRecipient{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipient) GetData() CaseNotificationRuleRecipientData {
	if o == nil || o.Data == nil {
		var ret CaseNotificationRuleRecipientData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipient) GetDataOk() (*CaseNotificationRuleRecipientData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipient) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given CaseNotificationRuleRecipientData and assigns it to the Data field.
func (o *CaseNotificationRuleRecipient) SetData(v CaseNotificationRuleRecipientData) {
	o.Data = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipient) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipient) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipient) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CaseNotificationRuleRecipient) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseNotificationRuleRecipient) MarshalJSON() ([]byte, error) {
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
func (o *CaseNotificationRuleRecipient) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *CaseNotificationRuleRecipientData `json:"data,omitempty"`
		Type *string                            `json:"type,omitempty"`
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
