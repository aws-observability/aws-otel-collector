// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AutomationRuleTriggerData Additional configuration for the trigger, dependent on the trigger type. For `status_transitioned` triggers, specify `from_status_name` and `to_status_name`. For `attribute_value_changed` triggers, specify `field` and `change_type`.
type AutomationRuleTriggerData struct {
	// The approval outcome to match. Used with `case_review_approved` triggers.
	ApprovalType *string `json:"approval_type,omitempty"`
	// The kind of attribute change to match. Allowed values: `VALUE_ADDED`, `VALUE_DELETED`, `ANY_CHANGES`. Used with `attribute_value_changed` triggers.
	ChangeType *string `json:"change_type,omitempty"`
	// The case attribute field name to monitor for changes. Used with `attribute_value_changed` triggers.
	Field *string `json:"field,omitempty"`
	// The originating status name. Used with `status_transitioned` triggers to match transitions from this status.
	FromStatusName *string `json:"from_status_name,omitempty"`
	// The destination status name. Used with `status_transitioned` triggers to match transitions to this status.
	ToStatusName *string `json:"to_status_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAutomationRuleTriggerData instantiates a new AutomationRuleTriggerData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAutomationRuleTriggerData() *AutomationRuleTriggerData {
	this := AutomationRuleTriggerData{}
	return &this
}

// NewAutomationRuleTriggerDataWithDefaults instantiates a new AutomationRuleTriggerData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAutomationRuleTriggerDataWithDefaults() *AutomationRuleTriggerData {
	this := AutomationRuleTriggerData{}
	return &this
}

// GetApprovalType returns the ApprovalType field value if set, zero value otherwise.
func (o *AutomationRuleTriggerData) GetApprovalType() string {
	if o == nil || o.ApprovalType == nil {
		var ret string
		return ret
	}
	return *o.ApprovalType
}

// GetApprovalTypeOk returns a tuple with the ApprovalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRuleTriggerData) GetApprovalTypeOk() (*string, bool) {
	if o == nil || o.ApprovalType == nil {
		return nil, false
	}
	return o.ApprovalType, true
}

// HasApprovalType returns a boolean if a field has been set.
func (o *AutomationRuleTriggerData) HasApprovalType() bool {
	return o != nil && o.ApprovalType != nil
}

// SetApprovalType gets a reference to the given string and assigns it to the ApprovalType field.
func (o *AutomationRuleTriggerData) SetApprovalType(v string) {
	o.ApprovalType = &v
}

// GetChangeType returns the ChangeType field value if set, zero value otherwise.
func (o *AutomationRuleTriggerData) GetChangeType() string {
	if o == nil || o.ChangeType == nil {
		var ret string
		return ret
	}
	return *o.ChangeType
}

// GetChangeTypeOk returns a tuple with the ChangeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRuleTriggerData) GetChangeTypeOk() (*string, bool) {
	if o == nil || o.ChangeType == nil {
		return nil, false
	}
	return o.ChangeType, true
}

// HasChangeType returns a boolean if a field has been set.
func (o *AutomationRuleTriggerData) HasChangeType() bool {
	return o != nil && o.ChangeType != nil
}

// SetChangeType gets a reference to the given string and assigns it to the ChangeType field.
func (o *AutomationRuleTriggerData) SetChangeType(v string) {
	o.ChangeType = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *AutomationRuleTriggerData) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRuleTriggerData) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *AutomationRuleTriggerData) HasField() bool {
	return o != nil && o.Field != nil
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *AutomationRuleTriggerData) SetField(v string) {
	o.Field = &v
}

// GetFromStatusName returns the FromStatusName field value if set, zero value otherwise.
func (o *AutomationRuleTriggerData) GetFromStatusName() string {
	if o == nil || o.FromStatusName == nil {
		var ret string
		return ret
	}
	return *o.FromStatusName
}

// GetFromStatusNameOk returns a tuple with the FromStatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRuleTriggerData) GetFromStatusNameOk() (*string, bool) {
	if o == nil || o.FromStatusName == nil {
		return nil, false
	}
	return o.FromStatusName, true
}

// HasFromStatusName returns a boolean if a field has been set.
func (o *AutomationRuleTriggerData) HasFromStatusName() bool {
	return o != nil && o.FromStatusName != nil
}

// SetFromStatusName gets a reference to the given string and assigns it to the FromStatusName field.
func (o *AutomationRuleTriggerData) SetFromStatusName(v string) {
	o.FromStatusName = &v
}

// GetToStatusName returns the ToStatusName field value if set, zero value otherwise.
func (o *AutomationRuleTriggerData) GetToStatusName() string {
	if o == nil || o.ToStatusName == nil {
		var ret string
		return ret
	}
	return *o.ToStatusName
}

// GetToStatusNameOk returns a tuple with the ToStatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationRuleTriggerData) GetToStatusNameOk() (*string, bool) {
	if o == nil || o.ToStatusName == nil {
		return nil, false
	}
	return o.ToStatusName, true
}

// HasToStatusName returns a boolean if a field has been set.
func (o *AutomationRuleTriggerData) HasToStatusName() bool {
	return o != nil && o.ToStatusName != nil
}

// SetToStatusName gets a reference to the given string and assigns it to the ToStatusName field.
func (o *AutomationRuleTriggerData) SetToStatusName(v string) {
	o.ToStatusName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AutomationRuleTriggerData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApprovalType != nil {
		toSerialize["approval_type"] = o.ApprovalType
	}
	if o.ChangeType != nil {
		toSerialize["change_type"] = o.ChangeType
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.FromStatusName != nil {
		toSerialize["from_status_name"] = o.FromStatusName
	}
	if o.ToStatusName != nil {
		toSerialize["to_status_name"] = o.ToStatusName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AutomationRuleTriggerData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApprovalType   *string `json:"approval_type,omitempty"`
		ChangeType     *string `json:"change_type,omitempty"`
		Field          *string `json:"field,omitempty"`
		FromStatusName *string `json:"from_status_name,omitempty"`
		ToStatusName   *string `json:"to_status_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"approval_type", "change_type", "field", "from_status_name", "to_status_name"})
	} else {
		return err
	}
	o.ApprovalType = all.ApprovalType
	o.ChangeType = all.ChangeType
	o.Field = all.Field
	o.FromStatusName = all.FromStatusName
	o.ToStatusName = all.ToStatusName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
