// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageProfilingHour The number of profiled hosts for each hour for a given organization.
type UsageProfilingHour struct {
	// Contains the total number of profiled Azure app services reporting during a given hour.
	AasCount datadog.NullableInt64 `json:"aas_count,omitempty"`
	// Get average number of container agents for that hour.
	AvgContainerAgentCount datadog.NullableInt64 `json:"avg_container_agent_count,omitempty"`
	// Contains the total number of profiled hosts reporting during a given hour.
	HostCount datadog.NullableInt64 `json:"host_count,omitempty"`
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

// NewUsageProfilingHour instantiates a new UsageProfilingHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageProfilingHour() *UsageProfilingHour {
	this := UsageProfilingHour{}
	return &this
}

// NewUsageProfilingHourWithDefaults instantiates a new UsageProfilingHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageProfilingHourWithDefaults() *UsageProfilingHour {
	this := UsageProfilingHour{}
	return &this
}

// GetAasCount returns the AasCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageProfilingHour) GetAasCount() int64 {
	if o == nil || o.AasCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AasCount.Get()
}

// GetAasCountOk returns a tuple with the AasCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageProfilingHour) GetAasCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AasCount.Get(), o.AasCount.IsSet()
}

// HasAasCount returns a boolean if a field has been set.
func (o *UsageProfilingHour) HasAasCount() bool {
	return o != nil && o.AasCount.IsSet()
}

// SetAasCount gets a reference to the given datadog.NullableInt64 and assigns it to the AasCount field.
func (o *UsageProfilingHour) SetAasCount(v int64) {
	o.AasCount.Set(&v)
}

// SetAasCountNil sets the value for AasCount to be an explicit nil.
func (o *UsageProfilingHour) SetAasCountNil() {
	o.AasCount.Set(nil)
}

// UnsetAasCount ensures that no value is present for AasCount, not even an explicit nil.
func (o *UsageProfilingHour) UnsetAasCount() {
	o.AasCount.Unset()
}

// GetAvgContainerAgentCount returns the AvgContainerAgentCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageProfilingHour) GetAvgContainerAgentCount() int64 {
	if o == nil || o.AvgContainerAgentCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.AvgContainerAgentCount.Get()
}

// GetAvgContainerAgentCountOk returns a tuple with the AvgContainerAgentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageProfilingHour) GetAvgContainerAgentCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvgContainerAgentCount.Get(), o.AvgContainerAgentCount.IsSet()
}

// HasAvgContainerAgentCount returns a boolean if a field has been set.
func (o *UsageProfilingHour) HasAvgContainerAgentCount() bool {
	return o != nil && o.AvgContainerAgentCount.IsSet()
}

// SetAvgContainerAgentCount gets a reference to the given datadog.NullableInt64 and assigns it to the AvgContainerAgentCount field.
func (o *UsageProfilingHour) SetAvgContainerAgentCount(v int64) {
	o.AvgContainerAgentCount.Set(&v)
}

// SetAvgContainerAgentCountNil sets the value for AvgContainerAgentCount to be an explicit nil.
func (o *UsageProfilingHour) SetAvgContainerAgentCountNil() {
	o.AvgContainerAgentCount.Set(nil)
}

// UnsetAvgContainerAgentCount ensures that no value is present for AvgContainerAgentCount, not even an explicit nil.
func (o *UsageProfilingHour) UnsetAvgContainerAgentCount() {
	o.AvgContainerAgentCount.Unset()
}

// GetHostCount returns the HostCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageProfilingHour) GetHostCount() int64 {
	if o == nil || o.HostCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.HostCount.Get()
}

// GetHostCountOk returns a tuple with the HostCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageProfilingHour) GetHostCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HostCount.Get(), o.HostCount.IsSet()
}

// HasHostCount returns a boolean if a field has been set.
func (o *UsageProfilingHour) HasHostCount() bool {
	return o != nil && o.HostCount.IsSet()
}

// SetHostCount gets a reference to the given datadog.NullableInt64 and assigns it to the HostCount field.
func (o *UsageProfilingHour) SetHostCount(v int64) {
	o.HostCount.Set(&v)
}

// SetHostCountNil sets the value for HostCount to be an explicit nil.
func (o *UsageProfilingHour) SetHostCountNil() {
	o.HostCount.Set(nil)
}

// UnsetHostCount ensures that no value is present for HostCount, not even an explicit nil.
func (o *UsageProfilingHour) UnsetHostCount() {
	o.HostCount.Unset()
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageProfilingHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageProfilingHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageProfilingHour) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageProfilingHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageProfilingHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageProfilingHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageProfilingHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageProfilingHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageProfilingHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageProfilingHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageProfilingHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageProfilingHour) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageProfilingHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AasCount.IsSet() {
		toSerialize["aas_count"] = o.AasCount.Get()
	}
	if o.AvgContainerAgentCount.IsSet() {
		toSerialize["avg_container_agent_count"] = o.AvgContainerAgentCount.Get()
	}
	if o.HostCount.IsSet() {
		toSerialize["host_count"] = o.HostCount.Get()
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
func (o *UsageProfilingHour) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AasCount               datadog.NullableInt64 `json:"aas_count,omitempty"`
		AvgContainerAgentCount datadog.NullableInt64 `json:"avg_container_agent_count,omitempty"`
		HostCount              datadog.NullableInt64 `json:"host_count,omitempty"`
		Hour                   *time.Time            `json:"hour,omitempty"`
		OrgName                *string               `json:"org_name,omitempty"`
		PublicId               *string               `json:"public_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aas_count", "avg_container_agent_count", "host_count", "hour", "org_name", "public_id"})
	} else {
		return err
	}
	o.AasCount = all.AasCount
	o.AvgContainerAgentCount = all.AvgContainerAgentCount
	o.HostCount = all.HostCount
	o.Hour = all.Hour
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
