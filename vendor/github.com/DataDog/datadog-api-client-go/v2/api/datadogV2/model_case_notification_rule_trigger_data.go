// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseNotificationRuleTriggerData Trigger data
type CaseNotificationRuleTriggerData struct {
	// Change type (added, removed, changed)
	ChangeType *string `json:"change_type,omitempty"`
	// Field name for attribute value changed trigger
	Field *string `json:"field,omitempty"`
	// Status ID to transition from
	FromStatus *string `json:"from_status,omitempty"`
	// Status name to transition from
	FromStatusName *string `json:"from_status_name,omitempty"`
	// Status ID to transition to
	ToStatus *string `json:"to_status,omitempty"`
	// Status name to transition to
	ToStatusName *string `json:"to_status_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseNotificationRuleTriggerData instantiates a new CaseNotificationRuleTriggerData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseNotificationRuleTriggerData() *CaseNotificationRuleTriggerData {
	this := CaseNotificationRuleTriggerData{}
	return &this
}

// NewCaseNotificationRuleTriggerDataWithDefaults instantiates a new CaseNotificationRuleTriggerData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseNotificationRuleTriggerDataWithDefaults() *CaseNotificationRuleTriggerData {
	this := CaseNotificationRuleTriggerData{}
	return &this
}

// GetChangeType returns the ChangeType field value if set, zero value otherwise.
func (o *CaseNotificationRuleTriggerData) GetChangeType() string {
	if o == nil || o.ChangeType == nil {
		var ret string
		return ret
	}
	return *o.ChangeType
}

// GetChangeTypeOk returns a tuple with the ChangeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleTriggerData) GetChangeTypeOk() (*string, bool) {
	if o == nil || o.ChangeType == nil {
		return nil, false
	}
	return o.ChangeType, true
}

// HasChangeType returns a boolean if a field has been set.
func (o *CaseNotificationRuleTriggerData) HasChangeType() bool {
	return o != nil && o.ChangeType != nil
}

// SetChangeType gets a reference to the given string and assigns it to the ChangeType field.
func (o *CaseNotificationRuleTriggerData) SetChangeType(v string) {
	o.ChangeType = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *CaseNotificationRuleTriggerData) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleTriggerData) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *CaseNotificationRuleTriggerData) HasField() bool {
	return o != nil && o.Field != nil
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *CaseNotificationRuleTriggerData) SetField(v string) {
	o.Field = &v
}

// GetFromStatus returns the FromStatus field value if set, zero value otherwise.
func (o *CaseNotificationRuleTriggerData) GetFromStatus() string {
	if o == nil || o.FromStatus == nil {
		var ret string
		return ret
	}
	return *o.FromStatus
}

// GetFromStatusOk returns a tuple with the FromStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleTriggerData) GetFromStatusOk() (*string, bool) {
	if o == nil || o.FromStatus == nil {
		return nil, false
	}
	return o.FromStatus, true
}

// HasFromStatus returns a boolean if a field has been set.
func (o *CaseNotificationRuleTriggerData) HasFromStatus() bool {
	return o != nil && o.FromStatus != nil
}

// SetFromStatus gets a reference to the given string and assigns it to the FromStatus field.
func (o *CaseNotificationRuleTriggerData) SetFromStatus(v string) {
	o.FromStatus = &v
}

// GetFromStatusName returns the FromStatusName field value if set, zero value otherwise.
func (o *CaseNotificationRuleTriggerData) GetFromStatusName() string {
	if o == nil || o.FromStatusName == nil {
		var ret string
		return ret
	}
	return *o.FromStatusName
}

// GetFromStatusNameOk returns a tuple with the FromStatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleTriggerData) GetFromStatusNameOk() (*string, bool) {
	if o == nil || o.FromStatusName == nil {
		return nil, false
	}
	return o.FromStatusName, true
}

// HasFromStatusName returns a boolean if a field has been set.
func (o *CaseNotificationRuleTriggerData) HasFromStatusName() bool {
	return o != nil && o.FromStatusName != nil
}

// SetFromStatusName gets a reference to the given string and assigns it to the FromStatusName field.
func (o *CaseNotificationRuleTriggerData) SetFromStatusName(v string) {
	o.FromStatusName = &v
}

// GetToStatus returns the ToStatus field value if set, zero value otherwise.
func (o *CaseNotificationRuleTriggerData) GetToStatus() string {
	if o == nil || o.ToStatus == nil {
		var ret string
		return ret
	}
	return *o.ToStatus
}

// GetToStatusOk returns a tuple with the ToStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleTriggerData) GetToStatusOk() (*string, bool) {
	if o == nil || o.ToStatus == nil {
		return nil, false
	}
	return o.ToStatus, true
}

// HasToStatus returns a boolean if a field has been set.
func (o *CaseNotificationRuleTriggerData) HasToStatus() bool {
	return o != nil && o.ToStatus != nil
}

// SetToStatus gets a reference to the given string and assigns it to the ToStatus field.
func (o *CaseNotificationRuleTriggerData) SetToStatus(v string) {
	o.ToStatus = &v
}

// GetToStatusName returns the ToStatusName field value if set, zero value otherwise.
func (o *CaseNotificationRuleTriggerData) GetToStatusName() string {
	if o == nil || o.ToStatusName == nil {
		var ret string
		return ret
	}
	return *o.ToStatusName
}

// GetToStatusNameOk returns a tuple with the ToStatusName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleTriggerData) GetToStatusNameOk() (*string, bool) {
	if o == nil || o.ToStatusName == nil {
		return nil, false
	}
	return o.ToStatusName, true
}

// HasToStatusName returns a boolean if a field has been set.
func (o *CaseNotificationRuleTriggerData) HasToStatusName() bool {
	return o != nil && o.ToStatusName != nil
}

// SetToStatusName gets a reference to the given string and assigns it to the ToStatusName field.
func (o *CaseNotificationRuleTriggerData) SetToStatusName(v string) {
	o.ToStatusName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseNotificationRuleTriggerData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChangeType != nil {
		toSerialize["change_type"] = o.ChangeType
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.FromStatus != nil {
		toSerialize["from_status"] = o.FromStatus
	}
	if o.FromStatusName != nil {
		toSerialize["from_status_name"] = o.FromStatusName
	}
	if o.ToStatus != nil {
		toSerialize["to_status"] = o.ToStatus
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
func (o *CaseNotificationRuleTriggerData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChangeType     *string `json:"change_type,omitempty"`
		Field          *string `json:"field,omitempty"`
		FromStatus     *string `json:"from_status,omitempty"`
		FromStatusName *string `json:"from_status_name,omitempty"`
		ToStatus       *string `json:"to_status,omitempty"`
		ToStatusName   *string `json:"to_status_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"change_type", "field", "from_status", "from_status_name", "to_status", "to_status_name"})
	} else {
		return err
	}
	o.ChangeType = all.ChangeType
	o.Field = all.Field
	o.FromStatus = all.FromStatus
	o.FromStatusName = all.FromStatusName
	o.ToStatus = all.ToStatus
	o.ToStatusName = all.ToStatusName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
