// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorNotificationRuleRelationships All relationships associated with monitor notification rule.
type MonitorNotificationRuleRelationships struct {
	// The user who created the monitor notification rule.
	CreatedBy *MonitorNotificationRuleRelationshipsCreatedBy `json:"created_by,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorNotificationRuleRelationships instantiates a new MonitorNotificationRuleRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorNotificationRuleRelationships() *MonitorNotificationRuleRelationships {
	this := MonitorNotificationRuleRelationships{}
	return &this
}

// NewMonitorNotificationRuleRelationshipsWithDefaults instantiates a new MonitorNotificationRuleRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorNotificationRuleRelationshipsWithDefaults() *MonitorNotificationRuleRelationships {
	this := MonitorNotificationRuleRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *MonitorNotificationRuleRelationships) GetCreatedBy() MonitorNotificationRuleRelationshipsCreatedBy {
	if o == nil || o.CreatedBy == nil {
		var ret MonitorNotificationRuleRelationshipsCreatedBy
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorNotificationRuleRelationships) GetCreatedByOk() (*MonitorNotificationRuleRelationshipsCreatedBy, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MonitorNotificationRuleRelationships) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given MonitorNotificationRuleRelationshipsCreatedBy and assigns it to the CreatedBy field.
func (o *MonitorNotificationRuleRelationships) SetCreatedBy(v MonitorNotificationRuleRelationshipsCreatedBy) {
	o.CreatedBy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorNotificationRuleRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorNotificationRuleRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedBy *MonitorNotificationRuleRelationshipsCreatedBy `json:"created_by,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedBy != nil && all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = all.CreatedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
