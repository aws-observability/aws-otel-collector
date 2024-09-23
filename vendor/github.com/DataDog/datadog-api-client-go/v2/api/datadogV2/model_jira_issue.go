// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraIssue Jira issue attached to case
type JiraIssue struct {
	// Jira issue information
	Result *JiraIssueResult `json:"result,omitempty"`
	// Case status
	Status *Case3rdPartyTicketStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraIssue instantiates a new JiraIssue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraIssue() *JiraIssue {
	this := JiraIssue{}
	return &this
}

// NewJiraIssueWithDefaults instantiates a new JiraIssue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraIssueWithDefaults() *JiraIssue {
	this := JiraIssue{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *JiraIssue) GetResult() JiraIssueResult {
	if o == nil || o.Result == nil {
		var ret JiraIssueResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssue) GetResultOk() (*JiraIssueResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *JiraIssue) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given JiraIssueResult and assigns it to the Result field.
func (o *JiraIssue) SetResult(v JiraIssueResult) {
	o.Result = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *JiraIssue) GetStatus() Case3rdPartyTicketStatus {
	if o == nil || o.Status == nil {
		var ret Case3rdPartyTicketStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraIssue) GetStatusOk() (*Case3rdPartyTicketStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *JiraIssue) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given Case3rdPartyTicketStatus and assigns it to the Status field.
func (o *JiraIssue) SetStatus(v Case3rdPartyTicketStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraIssue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraIssue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Result *JiraIssueResult          `json:"result,omitempty"`
		Status *Case3rdPartyTicketStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"result", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Result != nil && all.Result.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Result = all.Result
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableJiraIssue handles when a null is used for JiraIssue.
type NullableJiraIssue struct {
	value *JiraIssue
	isSet bool
}

// Get returns the associated value.
func (v NullableJiraIssue) Get() *JiraIssue {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableJiraIssue) Set(val *JiraIssue) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableJiraIssue) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableJiraIssue) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableJiraIssue initializes the struct as if Set has been called.
func NewNullableJiraIssue(val *JiraIssue) *NullableJiraIssue {
	return &NullableJiraIssue{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableJiraIssue) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableJiraIssue) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
