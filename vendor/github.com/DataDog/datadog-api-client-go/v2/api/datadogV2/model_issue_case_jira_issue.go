// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueCaseJiraIssue Jira issue of the case.
type IssueCaseJiraIssue struct {
	// Contains the identifiers and URL for a successfully created Jira issue.
	Result *IssueCaseJiraIssueResult `json:"result,omitempty"`
	// Creation status of the Jira issue.
	Status *string `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueCaseJiraIssue instantiates a new IssueCaseJiraIssue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueCaseJiraIssue() *IssueCaseJiraIssue {
	this := IssueCaseJiraIssue{}
	return &this
}

// NewIssueCaseJiraIssueWithDefaults instantiates a new IssueCaseJiraIssue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueCaseJiraIssueWithDefaults() *IssueCaseJiraIssue {
	this := IssueCaseJiraIssue{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *IssueCaseJiraIssue) GetResult() IssueCaseJiraIssueResult {
	if o == nil || o.Result == nil {
		var ret IssueCaseJiraIssueResult
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseJiraIssue) GetResultOk() (*IssueCaseJiraIssueResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *IssueCaseJiraIssue) HasResult() bool {
	return o != nil && o.Result != nil
}

// SetResult gets a reference to the given IssueCaseJiraIssueResult and assigns it to the Result field.
func (o *IssueCaseJiraIssue) SetResult(v IssueCaseJiraIssueResult) {
	o.Result = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IssueCaseJiraIssue) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCaseJiraIssue) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IssueCaseJiraIssue) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IssueCaseJiraIssue) SetStatus(v string) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueCaseJiraIssue) MarshalJSON() ([]byte, error) {
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
func (o *IssueCaseJiraIssue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Result *IssueCaseJiraIssueResult `json:"result,omitempty"`
		Status *string                   `json:"status,omitempty"`
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
	o.Status = all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
