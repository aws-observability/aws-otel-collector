// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageLogsByRetentionHour The number of indexed logs for each hour for a given organization broken down by retention period.
type UsageLogsByRetentionHour struct {
	// Total logs indexed with this retention period during a given hour.
	IndexedEventsCount datadog.NullableInt64 `json:"indexed_events_count,omitempty"`
	// Live logs indexed with this retention period during a given hour.
	LiveIndexedEventsCount datadog.NullableInt64 `json:"live_indexed_events_count,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// Rehydrated logs indexed with this retention period during a given hour.
	RehydratedIndexedEventsCount datadog.NullableInt64 `json:"rehydrated_indexed_events_count,omitempty"`
	// The retention period in days or "custom" for all custom retention usage.
	Retention datadog.NullableString `json:"retention,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageLogsByRetentionHour instantiates a new UsageLogsByRetentionHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageLogsByRetentionHour() *UsageLogsByRetentionHour {
	this := UsageLogsByRetentionHour{}
	return &this
}

// NewUsageLogsByRetentionHourWithDefaults instantiates a new UsageLogsByRetentionHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageLogsByRetentionHourWithDefaults() *UsageLogsByRetentionHour {
	this := UsageLogsByRetentionHour{}
	return &this
}

// GetIndexedEventsCount returns the IndexedEventsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLogsByRetentionHour) GetIndexedEventsCount() int64 {
	if o == nil || o.IndexedEventsCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IndexedEventsCount.Get()
}

// GetIndexedEventsCountOk returns a tuple with the IndexedEventsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLogsByRetentionHour) GetIndexedEventsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexedEventsCount.Get(), o.IndexedEventsCount.IsSet()
}

// HasIndexedEventsCount returns a boolean if a field has been set.
func (o *UsageLogsByRetentionHour) HasIndexedEventsCount() bool {
	return o != nil && o.IndexedEventsCount.IsSet()
}

// SetIndexedEventsCount gets a reference to the given datadog.NullableInt64 and assigns it to the IndexedEventsCount field.
func (o *UsageLogsByRetentionHour) SetIndexedEventsCount(v int64) {
	o.IndexedEventsCount.Set(&v)
}

// SetIndexedEventsCountNil sets the value for IndexedEventsCount to be an explicit nil.
func (o *UsageLogsByRetentionHour) SetIndexedEventsCountNil() {
	o.IndexedEventsCount.Set(nil)
}

// UnsetIndexedEventsCount ensures that no value is present for IndexedEventsCount, not even an explicit nil.
func (o *UsageLogsByRetentionHour) UnsetIndexedEventsCount() {
	o.IndexedEventsCount.Unset()
}

// GetLiveIndexedEventsCount returns the LiveIndexedEventsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLogsByRetentionHour) GetLiveIndexedEventsCount() int64 {
	if o == nil || o.LiveIndexedEventsCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LiveIndexedEventsCount.Get()
}

// GetLiveIndexedEventsCountOk returns a tuple with the LiveIndexedEventsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLogsByRetentionHour) GetLiveIndexedEventsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LiveIndexedEventsCount.Get(), o.LiveIndexedEventsCount.IsSet()
}

// HasLiveIndexedEventsCount returns a boolean if a field has been set.
func (o *UsageLogsByRetentionHour) HasLiveIndexedEventsCount() bool {
	return o != nil && o.LiveIndexedEventsCount.IsSet()
}

// SetLiveIndexedEventsCount gets a reference to the given datadog.NullableInt64 and assigns it to the LiveIndexedEventsCount field.
func (o *UsageLogsByRetentionHour) SetLiveIndexedEventsCount(v int64) {
	o.LiveIndexedEventsCount.Set(&v)
}

// SetLiveIndexedEventsCountNil sets the value for LiveIndexedEventsCount to be an explicit nil.
func (o *UsageLogsByRetentionHour) SetLiveIndexedEventsCountNil() {
	o.LiveIndexedEventsCount.Set(nil)
}

// UnsetLiveIndexedEventsCount ensures that no value is present for LiveIndexedEventsCount, not even an explicit nil.
func (o *UsageLogsByRetentionHour) UnsetLiveIndexedEventsCount() {
	o.LiveIndexedEventsCount.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageLogsByRetentionHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLogsByRetentionHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageLogsByRetentionHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageLogsByRetentionHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageLogsByRetentionHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLogsByRetentionHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageLogsByRetentionHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageLogsByRetentionHour) SetPublicId(v string) {
	o.PublicId = &v
}

// GetRehydratedIndexedEventsCount returns the RehydratedIndexedEventsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLogsByRetentionHour) GetRehydratedIndexedEventsCount() int64 {
	if o == nil || o.RehydratedIndexedEventsCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RehydratedIndexedEventsCount.Get()
}

// GetRehydratedIndexedEventsCountOk returns a tuple with the RehydratedIndexedEventsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLogsByRetentionHour) GetRehydratedIndexedEventsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RehydratedIndexedEventsCount.Get(), o.RehydratedIndexedEventsCount.IsSet()
}

// HasRehydratedIndexedEventsCount returns a boolean if a field has been set.
func (o *UsageLogsByRetentionHour) HasRehydratedIndexedEventsCount() bool {
	return o != nil && o.RehydratedIndexedEventsCount.IsSet()
}

// SetRehydratedIndexedEventsCount gets a reference to the given datadog.NullableInt64 and assigns it to the RehydratedIndexedEventsCount field.
func (o *UsageLogsByRetentionHour) SetRehydratedIndexedEventsCount(v int64) {
	o.RehydratedIndexedEventsCount.Set(&v)
}

// SetRehydratedIndexedEventsCountNil sets the value for RehydratedIndexedEventsCount to be an explicit nil.
func (o *UsageLogsByRetentionHour) SetRehydratedIndexedEventsCountNil() {
	o.RehydratedIndexedEventsCount.Set(nil)
}

// UnsetRehydratedIndexedEventsCount ensures that no value is present for RehydratedIndexedEventsCount, not even an explicit nil.
func (o *UsageLogsByRetentionHour) UnsetRehydratedIndexedEventsCount() {
	o.RehydratedIndexedEventsCount.Unset()
}

// GetRetention returns the Retention field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLogsByRetentionHour) GetRetention() string {
	if o == nil || o.Retention.Get() == nil {
		var ret string
		return ret
	}
	return *o.Retention.Get()
}

// GetRetentionOk returns a tuple with the Retention field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLogsByRetentionHour) GetRetentionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Retention.Get(), o.Retention.IsSet()
}

// HasRetention returns a boolean if a field has been set.
func (o *UsageLogsByRetentionHour) HasRetention() bool {
	return o != nil && o.Retention.IsSet()
}

// SetRetention gets a reference to the given datadog.NullableString and assigns it to the Retention field.
func (o *UsageLogsByRetentionHour) SetRetention(v string) {
	o.Retention.Set(&v)
}

// SetRetentionNil sets the value for Retention to be an explicit nil.
func (o *UsageLogsByRetentionHour) SetRetentionNil() {
	o.Retention.Set(nil)
}

// UnsetRetention ensures that no value is present for Retention, not even an explicit nil.
func (o *UsageLogsByRetentionHour) UnsetRetention() {
	o.Retention.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageLogsByRetentionHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.IndexedEventsCount.IsSet() {
		toSerialize["indexed_events_count"] = o.IndexedEventsCount.Get()
	}
	if o.LiveIndexedEventsCount.IsSet() {
		toSerialize["live_indexed_events_count"] = o.LiveIndexedEventsCount.Get()
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.RehydratedIndexedEventsCount.IsSet() {
		toSerialize["rehydrated_indexed_events_count"] = o.RehydratedIndexedEventsCount.Get()
	}
	if o.Retention.IsSet() {
		toSerialize["retention"] = o.Retention.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageLogsByRetentionHour) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		IndexedEventsCount           datadog.NullableInt64  `json:"indexed_events_count,omitempty"`
		LiveIndexedEventsCount       datadog.NullableInt64  `json:"live_indexed_events_count,omitempty"`
		OrgName                      *string                `json:"org_name,omitempty"`
		PublicId                     *string                `json:"public_id,omitempty"`
		RehydratedIndexedEventsCount datadog.NullableInt64  `json:"rehydrated_indexed_events_count,omitempty"`
		Retention                    datadog.NullableString `json:"retention,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"indexed_events_count", "live_indexed_events_count", "org_name", "public_id", "rehydrated_indexed_events_count", "retention"})
	} else {
		return err
	}
	o.IndexedEventsCount = all.IndexedEventsCount
	o.LiveIndexedEventsCount = all.LiveIndexedEventsCount
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId
	o.RehydratedIndexedEventsCount = all.RehydratedIndexedEventsCount
	o.Retention = all.Retention

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
