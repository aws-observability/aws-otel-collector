// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssuesSearchResultAttributes Object containing the information of a search result.
type IssuesSearchResultAttributes struct {
	// Count of sessions impacted by the issue over the queried time window.
	ImpactedSessions *int64 `json:"impacted_sessions,omitempty"`
	// Count of users impacted by the issue over the queried time window.
	ImpactedUsers *int64 `json:"impacted_users,omitempty"`
	// Total count of errors that match the issue over the queried time window.
	TotalCount *int64 `json:"total_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssuesSearchResultAttributes instantiates a new IssuesSearchResultAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssuesSearchResultAttributes() *IssuesSearchResultAttributes {
	this := IssuesSearchResultAttributes{}
	return &this
}

// NewIssuesSearchResultAttributesWithDefaults instantiates a new IssuesSearchResultAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssuesSearchResultAttributesWithDefaults() *IssuesSearchResultAttributes {
	this := IssuesSearchResultAttributes{}
	return &this
}

// GetImpactedSessions returns the ImpactedSessions field value if set, zero value otherwise.
func (o *IssuesSearchResultAttributes) GetImpactedSessions() int64 {
	if o == nil || o.ImpactedSessions == nil {
		var ret int64
		return ret
	}
	return *o.ImpactedSessions
}

// GetImpactedSessionsOk returns a tuple with the ImpactedSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesSearchResultAttributes) GetImpactedSessionsOk() (*int64, bool) {
	if o == nil || o.ImpactedSessions == nil {
		return nil, false
	}
	return o.ImpactedSessions, true
}

// HasImpactedSessions returns a boolean if a field has been set.
func (o *IssuesSearchResultAttributes) HasImpactedSessions() bool {
	return o != nil && o.ImpactedSessions != nil
}

// SetImpactedSessions gets a reference to the given int64 and assigns it to the ImpactedSessions field.
func (o *IssuesSearchResultAttributes) SetImpactedSessions(v int64) {
	o.ImpactedSessions = &v
}

// GetImpactedUsers returns the ImpactedUsers field value if set, zero value otherwise.
func (o *IssuesSearchResultAttributes) GetImpactedUsers() int64 {
	if o == nil || o.ImpactedUsers == nil {
		var ret int64
		return ret
	}
	return *o.ImpactedUsers
}

// GetImpactedUsersOk returns a tuple with the ImpactedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesSearchResultAttributes) GetImpactedUsersOk() (*int64, bool) {
	if o == nil || o.ImpactedUsers == nil {
		return nil, false
	}
	return o.ImpactedUsers, true
}

// HasImpactedUsers returns a boolean if a field has been set.
func (o *IssuesSearchResultAttributes) HasImpactedUsers() bool {
	return o != nil && o.ImpactedUsers != nil
}

// SetImpactedUsers gets a reference to the given int64 and assigns it to the ImpactedUsers field.
func (o *IssuesSearchResultAttributes) SetImpactedUsers(v int64) {
	o.ImpactedUsers = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *IssuesSearchResultAttributes) GetTotalCount() int64 {
	if o == nil || o.TotalCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssuesSearchResultAttributes) GetTotalCountOk() (*int64, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *IssuesSearchResultAttributes) HasTotalCount() bool {
	return o != nil && o.TotalCount != nil
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *IssuesSearchResultAttributes) SetTotalCount(v int64) {
	o.TotalCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssuesSearchResultAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ImpactedSessions != nil {
		toSerialize["impacted_sessions"] = o.ImpactedSessions
	}
	if o.ImpactedUsers != nil {
		toSerialize["impacted_users"] = o.ImpactedUsers
	}
	if o.TotalCount != nil {
		toSerialize["total_count"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssuesSearchResultAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ImpactedSessions *int64 `json:"impacted_sessions,omitempty"`
		ImpactedUsers    *int64 `json:"impacted_users,omitempty"`
		TotalCount       *int64 `json:"total_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"impacted_sessions", "impacted_users", "total_count"})
	} else {
		return err
	}
	o.ImpactedSessions = all.ImpactedSessions
	o.ImpactedUsers = all.ImpactedUsers
	o.TotalCount = all.TotalCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
