// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorNotificationRuleUpdateRequestData Object to update a monitor notification rule.
type MonitorNotificationRuleUpdateRequestData struct {
	// Attributes of the monitor notification rule.
	Attributes MonitorNotificationRuleAttributes `json:"attributes"`
	// The ID of the monitor notification rule.
	Id string `json:"id"`
	// Monitor notification rule resource type.
	Type *MonitorNotificationRuleResourceType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorNotificationRuleUpdateRequestData instantiates a new MonitorNotificationRuleUpdateRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorNotificationRuleUpdateRequestData(attributes MonitorNotificationRuleAttributes, id string) *MonitorNotificationRuleUpdateRequestData {
	this := MonitorNotificationRuleUpdateRequestData{}
	this.Attributes = attributes
	this.Id = id
	var typeVar MonitorNotificationRuleResourceType = MONITORNOTIFICATIONRULERESOURCETYPE_MONITOR_NOTIFICATION_RULE
	this.Type = &typeVar
	return &this
}

// NewMonitorNotificationRuleUpdateRequestDataWithDefaults instantiates a new MonitorNotificationRuleUpdateRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorNotificationRuleUpdateRequestDataWithDefaults() *MonitorNotificationRuleUpdateRequestData {
	this := MonitorNotificationRuleUpdateRequestData{}
	var typeVar MonitorNotificationRuleResourceType = MONITORNOTIFICATIONRULERESOURCETYPE_MONITOR_NOTIFICATION_RULE
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *MonitorNotificationRuleUpdateRequestData) GetAttributes() MonitorNotificationRuleAttributes {
	if o == nil {
		var ret MonitorNotificationRuleAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleUpdateRequestData) GetAttributesOk() (*MonitorNotificationRuleAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *MonitorNotificationRuleUpdateRequestData) SetAttributes(v MonitorNotificationRuleAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *MonitorNotificationRuleUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *MonitorNotificationRuleUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitorNotificationRuleUpdateRequestData) GetType() MonitorNotificationRuleResourceType {
	if o == nil || o.Type == nil {
		var ret MonitorNotificationRuleResourceType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleUpdateRequestData) GetTypeOk() (*MonitorNotificationRuleResourceType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitorNotificationRuleUpdateRequestData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given MonitorNotificationRuleResourceType and assigns it to the Type field.
func (o *MonitorNotificationRuleUpdateRequestData) SetType(v MonitorNotificationRuleResourceType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorNotificationRuleUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorNotificationRuleUpdateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *MonitorNotificationRuleAttributes   `json:"attributes"`
		Id         *string                              `json:"id"`
		Type       *MonitorNotificationRuleResourceType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
