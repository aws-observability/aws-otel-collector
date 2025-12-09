// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageLambdaHour Number of Lambda functions and sum of the invocations of all Lambda functions
// for each hour for a given organization.
type UsageLambdaHour struct {
	// Contains the number of different functions for each region and AWS account.
	FuncCount datadog.NullableInt64 `json:"func_count,omitempty"`
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// Contains the sum of invocations of all functions.
	InvocationsSum datadog.NullableInt64 `json:"invocations_sum,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageLambdaHour instantiates a new UsageLambdaHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageLambdaHour() *UsageLambdaHour {
	this := UsageLambdaHour{}
	return &this
}

// NewUsageLambdaHourWithDefaults instantiates a new UsageLambdaHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageLambdaHourWithDefaults() *UsageLambdaHour {
	this := UsageLambdaHour{}
	return &this
}

// GetFuncCount returns the FuncCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLambdaHour) GetFuncCount() int64 {
	if o == nil || o.FuncCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.FuncCount.Get()
}

// GetFuncCountOk returns a tuple with the FuncCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLambdaHour) GetFuncCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.FuncCount.Get(), o.FuncCount.IsSet()
}

// HasFuncCount returns a boolean if a field has been set.
func (o *UsageLambdaHour) HasFuncCount() bool {
	return o != nil && o.FuncCount.IsSet()
}

// SetFuncCount gets a reference to the given datadog.NullableInt64 and assigns it to the FuncCount field.
func (o *UsageLambdaHour) SetFuncCount(v int64) {
	o.FuncCount.Set(&v)
}

// SetFuncCountNil sets the value for FuncCount to be an explicit nil.
func (o *UsageLambdaHour) SetFuncCountNil() {
	o.FuncCount.Set(nil)
}

// UnsetFuncCount ensures that no value is present for FuncCount, not even an explicit nil.
func (o *UsageLambdaHour) UnsetFuncCount() {
	o.FuncCount.Unset()
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageLambdaHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLambdaHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageLambdaHour) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageLambdaHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetInvocationsSum returns the InvocationsSum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLambdaHour) GetInvocationsSum() int64 {
	if o == nil || o.InvocationsSum.Get() == nil {
		var ret int64
		return ret
	}
	return *o.InvocationsSum.Get()
}

// GetInvocationsSumOk returns a tuple with the InvocationsSum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLambdaHour) GetInvocationsSumOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvocationsSum.Get(), o.InvocationsSum.IsSet()
}

// HasInvocationsSum returns a boolean if a field has been set.
func (o *UsageLambdaHour) HasInvocationsSum() bool {
	return o != nil && o.InvocationsSum.IsSet()
}

// SetInvocationsSum gets a reference to the given datadog.NullableInt64 and assigns it to the InvocationsSum field.
func (o *UsageLambdaHour) SetInvocationsSum(v int64) {
	o.InvocationsSum.Set(&v)
}

// SetInvocationsSumNil sets the value for InvocationsSum to be an explicit nil.
func (o *UsageLambdaHour) SetInvocationsSumNil() {
	o.InvocationsSum.Set(nil)
}

// UnsetInvocationsSum ensures that no value is present for InvocationsSum, not even an explicit nil.
func (o *UsageLambdaHour) UnsetInvocationsSum() {
	o.InvocationsSum.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageLambdaHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLambdaHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageLambdaHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageLambdaHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageLambdaHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLambdaHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageLambdaHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageLambdaHour) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageLambdaHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FuncCount.IsSet() {
		toSerialize["func_count"] = o.FuncCount.Get()
	}
	if o.Hour != nil {
		if o.Hour.Nanosecond() == 0 {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.InvocationsSum.IsSet() {
		toSerialize["invocations_sum"] = o.InvocationsSum.Get()
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageLambdaHour) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FuncCount      datadog.NullableInt64 `json:"func_count,omitempty"`
		Hour           *time.Time            `json:"hour,omitempty"`
		InvocationsSum datadog.NullableInt64 `json:"invocations_sum,omitempty"`
		OrgName        *string               `json:"org_name,omitempty"`
		PublicId       *string               `json:"public_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"func_count", "hour", "invocations_sum", "org_name", "public_id"})
	} else {
		return err
	}
	o.FuncCount = all.FuncCount
	o.Hour = all.Hour
	o.InvocationsSum = all.InvocationsSum
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
