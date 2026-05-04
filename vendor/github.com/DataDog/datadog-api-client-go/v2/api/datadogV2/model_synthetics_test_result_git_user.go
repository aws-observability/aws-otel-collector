// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultGitUser A Git user (author or committer).
type SyntheticsTestResultGitUser struct {
	// Timestamp of the commit action for this user.
	Date *string `json:"date,omitempty"`
	// Email address of the Git user.
	Email *string `json:"email,omitempty"`
	// Name of the Git user.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultGitUser instantiates a new SyntheticsTestResultGitUser object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultGitUser() *SyntheticsTestResultGitUser {
	this := SyntheticsTestResultGitUser{}
	return &this
}

// NewSyntheticsTestResultGitUserWithDefaults instantiates a new SyntheticsTestResultGitUser object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultGitUserWithDefaults() *SyntheticsTestResultGitUser {
	this := SyntheticsTestResultGitUser{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *SyntheticsTestResultGitUser) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultGitUser) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *SyntheticsTestResultGitUser) HasDate() bool {
	return o != nil && o.Date != nil
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *SyntheticsTestResultGitUser) SetDate(v string) {
	o.Date = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SyntheticsTestResultGitUser) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultGitUser) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SyntheticsTestResultGitUser) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SyntheticsTestResultGitUser) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTestResultGitUser) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultGitUser) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTestResultGitUser) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTestResultGitUser) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultGitUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultGitUser) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Date  *string `json:"date,omitempty"`
		Email *string `json:"email,omitempty"`
		Name  *string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"date", "email", "name"})
	} else {
		return err
	}
	o.Date = all.Date
	o.Email = all.Email
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
