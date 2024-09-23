// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageCWSHour Cloud Workload Security usage for a given organization for a given hour.
type UsageCWSHour struct {
	// The total number of Cloud Workload Security container hours from the start of the given hour’s month until the given hour.
	CwsContainerCount datadog.NullableInt64 `json:"cws_container_count,omitempty"`
	// The total number of Cloud Workload Security host hours from the start of the given hour’s month until the given hour.
	CwsHostCount datadog.NullableInt64 `json:"cws_host_count,omitempty"`
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageCWSHour instantiates a new UsageCWSHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageCWSHour() *UsageCWSHour {
	this := UsageCWSHour{}
	return &this
}

// NewUsageCWSHourWithDefaults instantiates a new UsageCWSHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageCWSHourWithDefaults() *UsageCWSHour {
	this := UsageCWSHour{}
	return &this
}

// GetCwsContainerCount returns the CwsContainerCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageCWSHour) GetCwsContainerCount() int64 {
	if o == nil || o.CwsContainerCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CwsContainerCount.Get()
}

// GetCwsContainerCountOk returns a tuple with the CwsContainerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageCWSHour) GetCwsContainerCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsContainerCount.Get(), o.CwsContainerCount.IsSet()
}

// HasCwsContainerCount returns a boolean if a field has been set.
func (o *UsageCWSHour) HasCwsContainerCount() bool {
	return o != nil && o.CwsContainerCount.IsSet()
}

// SetCwsContainerCount gets a reference to the given datadog.NullableInt64 and assigns it to the CwsContainerCount field.
func (o *UsageCWSHour) SetCwsContainerCount(v int64) {
	o.CwsContainerCount.Set(&v)
}

// SetCwsContainerCountNil sets the value for CwsContainerCount to be an explicit nil.
func (o *UsageCWSHour) SetCwsContainerCountNil() {
	o.CwsContainerCount.Set(nil)
}

// UnsetCwsContainerCount ensures that no value is present for CwsContainerCount, not even an explicit nil.
func (o *UsageCWSHour) UnsetCwsContainerCount() {
	o.CwsContainerCount.Unset()
}

// GetCwsHostCount returns the CwsHostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageCWSHour) GetCwsHostCount() int64 {
	if o == nil || o.CwsHostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.CwsHostCount.Get()
}

// GetCwsHostCountOk returns a tuple with the CwsHostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageCWSHour) GetCwsHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CwsHostCount.Get(), o.CwsHostCount.IsSet()
}

// HasCwsHostCount returns a boolean if a field has been set.
func (o *UsageCWSHour) HasCwsHostCount() bool {
	return o != nil && o.CwsHostCount.IsSet()
}

// SetCwsHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the CwsHostCount field.
func (o *UsageCWSHour) SetCwsHostCount(v int64) {
	o.CwsHostCount.Set(&v)
}

// SetCwsHostCountNil sets the value for CwsHostCount to be an explicit nil.
func (o *UsageCWSHour) SetCwsHostCountNil() {
	o.CwsHostCount.Set(nil)
}

// UnsetCwsHostCount ensures that no value is present for CwsHostCount, not even an explicit nil.
func (o *UsageCWSHour) UnsetCwsHostCount() {
	o.CwsHostCount.Unset()
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageCWSHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCWSHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageCWSHour) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageCWSHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageCWSHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCWSHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageCWSHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageCWSHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageCWSHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageCWSHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageCWSHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageCWSHour) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageCWSHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CwsContainerCount.IsSet() {
		toSerialize["cws_container_count"] = o.CwsContainerCount.Get()
	}
	if o.CwsHostCount.IsSet() {
		toSerialize["cws_host_count"] = o.CwsHostCount.Get()
	}
	if o.Hour != nil {
		if o.Hour.Nanosecond() == 0 {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
func (o *UsageCWSHour) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CwsContainerCount datadog.NullableInt64 `json:"cws_container_count,omitempty"`
		CwsHostCount      datadog.NullableInt64 `json:"cws_host_count,omitempty"`
		Hour              *time.Time            `json:"hour,omitempty"`
		OrgName           *string               `json:"org_name,omitempty"`
		PublicId          *string               `json:"public_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cws_container_count", "cws_host_count", "hour", "org_name", "public_id"})
	} else {
		return err
	}
	o.CwsContainerCount = all.CwsContainerCount
	o.CwsHostCount = all.CwsHostCount
	o.Hour = all.Hour
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
