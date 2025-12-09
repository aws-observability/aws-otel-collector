// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentNotificationRuleConditionsItems A condition that must be met to trigger the notification rule.
type IncidentNotificationRuleConditionsItems struct {
	// The incident field to evaluate
	Field string `json:"field"`
	// The value(s) to compare against. Multiple values are `ORed` together.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentNotificationRuleConditionsItems instantiates a new IncidentNotificationRuleConditionsItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentNotificationRuleConditionsItems(field string, values []string) *IncidentNotificationRuleConditionsItems {
	this := IncidentNotificationRuleConditionsItems{}
	this.Field = field
	this.Values = values
	return &this
}

// NewIncidentNotificationRuleConditionsItemsWithDefaults instantiates a new IncidentNotificationRuleConditionsItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentNotificationRuleConditionsItemsWithDefaults() *IncidentNotificationRuleConditionsItems {
	this := IncidentNotificationRuleConditionsItems{}
	return &this
}

// GetField returns the Field field value.
func (o *IncidentNotificationRuleConditionsItems) GetField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleConditionsItems) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value.
func (o *IncidentNotificationRuleConditionsItems) SetField(v string) {
	o.Field = v
}

// GetValues returns the Values field value.
func (o *IncidentNotificationRuleConditionsItems) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *IncidentNotificationRuleConditionsItems) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *IncidentNotificationRuleConditionsItems) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentNotificationRuleConditionsItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["field"] = o.Field
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentNotificationRuleConditionsItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Field  *string   `json:"field"`
		Values *[]string `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Field == nil {
		return fmt.Errorf("required field field missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"field", "values"})
	} else {
		return err
	}
	o.Field = *all.Field
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
