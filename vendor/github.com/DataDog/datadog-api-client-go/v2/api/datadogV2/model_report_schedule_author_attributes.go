// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleAuthorAttributes Attributes of the report author.
type ReportScheduleAuthorAttributes struct {
	// The email address of the report author, or `null` if unavailable.
	Email datadog.NullableString `json:"email"`
	// The display name of the report author, or `null` if unavailable.
	Name datadog.NullableString `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportScheduleAuthorAttributes instantiates a new ReportScheduleAuthorAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportScheduleAuthorAttributes(email datadog.NullableString, name datadog.NullableString) *ReportScheduleAuthorAttributes {
	this := ReportScheduleAuthorAttributes{}
	this.Email = email
	this.Name = name
	return &this
}

// NewReportScheduleAuthorAttributesWithDefaults instantiates a new ReportScheduleAuthorAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportScheduleAuthorAttributesWithDefaults() *ReportScheduleAuthorAttributes {
	this := ReportScheduleAuthorAttributes{}
	return &this
}

// GetEmail returns the Email field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ReportScheduleAuthorAttributes) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleAuthorAttributes) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// SetEmail sets field value.
func (o *ReportScheduleAuthorAttributes) SetEmail(v string) {
	o.Email.Set(&v)
}

// GetName returns the Name field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ReportScheduleAuthorAttributes) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleAuthorAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value.
func (o *ReportScheduleAuthorAttributes) SetName(v string) {
	o.Name.Set(&v)
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportScheduleAuthorAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["email"] = o.Email.Get()
	toSerialize["name"] = o.Name.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReportScheduleAuthorAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Email datadog.NullableString `json:"email"`
		Name  datadog.NullableString `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Email.IsSet() {
		return fmt.Errorf("required field email missing")
	}
	if !all.Name.IsSet() {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"email", "name"})
	} else {
		return err
	}
	o.Email = all.Email
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
