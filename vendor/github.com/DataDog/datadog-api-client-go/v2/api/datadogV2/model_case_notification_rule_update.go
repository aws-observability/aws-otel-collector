// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseNotificationRuleUpdate Notification rule update
type CaseNotificationRuleUpdate struct {
	// Notification rule attributes
	Attributes *CaseNotificationRuleAttributes `json:"attributes,omitempty"`
	// Notification rule resource type
	Type CaseNotificationRuleResourceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseNotificationRuleUpdate instantiates a new CaseNotificationRuleUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseNotificationRuleUpdate(typeVar CaseNotificationRuleResourceType) *CaseNotificationRuleUpdate {
	this := CaseNotificationRuleUpdate{}
	this.Type = typeVar
	return &this
}

// NewCaseNotificationRuleUpdateWithDefaults instantiates a new CaseNotificationRuleUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseNotificationRuleUpdateWithDefaults() *CaseNotificationRuleUpdate {
	this := CaseNotificationRuleUpdate{}
	var typeVar CaseNotificationRuleResourceType = CASENOTIFICATIONRULERESOURCETYPE_NOTIFICATION_RULE
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CaseNotificationRuleUpdate) GetAttributes() CaseNotificationRuleAttributes {
	if o == nil || o.Attributes == nil {
		var ret CaseNotificationRuleAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleUpdate) GetAttributesOk() (*CaseNotificationRuleAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CaseNotificationRuleUpdate) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CaseNotificationRuleAttributes and assigns it to the Attributes field.
func (o *CaseNotificationRuleUpdate) SetAttributes(v CaseNotificationRuleAttributes) {
	o.Attributes = &v
}

// GetType returns the Type field value.
func (o *CaseNotificationRuleUpdate) GetType() CaseNotificationRuleResourceType {
	if o == nil {
		var ret CaseNotificationRuleResourceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleUpdate) GetTypeOk() (*CaseNotificationRuleResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CaseNotificationRuleUpdate) SetType(v CaseNotificationRuleResourceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseNotificationRuleUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseNotificationRuleUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CaseNotificationRuleAttributes   `json:"attributes,omitempty"`
		Type       *CaseNotificationRuleResourceType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
