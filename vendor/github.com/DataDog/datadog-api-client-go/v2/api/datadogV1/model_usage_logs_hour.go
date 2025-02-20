// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageLogsHour Hour usage for logs.
type UsageLogsHour struct {
	// Contains the number of billable log bytes ingested.
	BillableIngestedBytes datadog.NullableInt64 `json:"billable_ingested_bytes,omitempty"`
	// The hour for the usage.
	Hour *time.Time `json:"hour,omitempty"`
	// Contains the number of log events indexed.
	IndexedEventsCount datadog.NullableInt64 `json:"indexed_events_count,omitempty"`
	// Contains the number of log bytes ingested.
	IngestedEventsBytes datadog.NullableInt64 `json:"ingested_events_bytes,omitempty"`
	// Contains the number of logs forwarded bytes (data available as of April 1st 2023)
	LogsForwardingEventsBytes datadog.NullableInt64 `json:"logs_forwarding_events_bytes,omitempty"`
	// Contains the number of live log events indexed (data available as of December 1, 2020).
	LogsLiveIndexedCount datadog.NullableInt64 `json:"logs_live_indexed_count,omitempty"`
	// Contains the number of live log bytes ingested (data available as of December 1, 2020).
	LogsLiveIngestedBytes datadog.NullableInt64 `json:"logs_live_ingested_bytes,omitempty"`
	// Contains the number of rehydrated log events indexed (data available as of December 1, 2020).
	LogsRehydratedIndexedCount datadog.NullableInt64 `json:"logs_rehydrated_indexed_count,omitempty"`
	// Contains the number of rehydrated log bytes ingested (data available as of December 1, 2020).
	LogsRehydratedIngestedBytes datadog.NullableInt64 `json:"logs_rehydrated_ingested_bytes,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageLogsHour instantiates a new UsageLogsHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageLogsHour() *UsageLogsHour {
	this := UsageLogsHour{}
	return &this
}

// NewUsageLogsHourWithDefaults instantiates a new UsageLogsHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageLogsHourWithDefaults() *UsageLogsHour {
	this := UsageLogsHour{}
	return &this
}

// GetBillableIngestedBytes returns the BillableIngestedBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLogsHour) GetBillableIngestedBytes() int64 {
	if o == nil || o.BillableIngestedBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BillableIngestedBytes.Get()
}

// GetBillableIngestedBytesOk returns a tuple with the BillableIngestedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLogsHour) GetBillableIngestedBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BillableIngestedBytes.Get(), o.BillableIngestedBytes.IsSet()
}

// HasBillableIngestedBytes returns a boolean if a field has been set.
func (o *UsageLogsHour) HasBillableIngestedBytes() bool {
	return o != nil && o.BillableIngestedBytes.IsSet()
}

// SetBillableIngestedBytes gets a reference to the given datadog.NullableInt64 and assigns it to the BillableIngestedBytes field.
func (o *UsageLogsHour) SetBillableIngestedBytes(v int64) {
	o.BillableIngestedBytes.Set(&v)
}

// SetBillableIngestedBytesNil sets the value for BillableIngestedBytes to be an explicit nil.
func (o *UsageLogsHour) SetBillableIngestedBytesNil() {
	o.BillableIngestedBytes.Set(nil)
}

// UnsetBillableIngestedBytes ensures that no value is present for BillableIngestedBytes, not even an explicit nil.
func (o *UsageLogsHour) UnsetBillableIngestedBytes() {
	o.BillableIngestedBytes.Unset()
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageLogsHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLogsHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageLogsHour) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageLogsHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetIndexedEventsCount returns the IndexedEventsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLogsHour) GetIndexedEventsCount() int64 {
	if o == nil || o.IndexedEventsCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IndexedEventsCount.Get()
}

// GetIndexedEventsCountOk returns a tuple with the IndexedEventsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLogsHour) GetIndexedEventsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IndexedEventsCount.Get(), o.IndexedEventsCount.IsSet()
}

// HasIndexedEventsCount returns a boolean if a field has been set.
func (o *UsageLogsHour) HasIndexedEventsCount() bool {
	return o != nil && o.IndexedEventsCount.IsSet()
}

// SetIndexedEventsCount gets a reference to the given datadog.NullableInt64 and assigns it to the IndexedEventsCount field.
func (o *UsageLogsHour) SetIndexedEventsCount(v int64) {
	o.IndexedEventsCount.Set(&v)
}

// SetIndexedEventsCountNil sets the value for IndexedEventsCount to be an explicit nil.
func (o *UsageLogsHour) SetIndexedEventsCountNil() {
	o.IndexedEventsCount.Set(nil)
}

// UnsetIndexedEventsCount ensures that no value is present for IndexedEventsCount, not even an explicit nil.
func (o *UsageLogsHour) UnsetIndexedEventsCount() {
	o.IndexedEventsCount.Unset()
}

// GetIngestedEventsBytes returns the IngestedEventsBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLogsHour) GetIngestedEventsBytes() int64 {
	if o == nil || o.IngestedEventsBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.IngestedEventsBytes.Get()
}

// GetIngestedEventsBytesOk returns a tuple with the IngestedEventsBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLogsHour) GetIngestedEventsBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.IngestedEventsBytes.Get(), o.IngestedEventsBytes.IsSet()
}

// HasIngestedEventsBytes returns a boolean if a field has been set.
func (o *UsageLogsHour) HasIngestedEventsBytes() bool {
	return o != nil && o.IngestedEventsBytes.IsSet()
}

// SetIngestedEventsBytes gets a reference to the given datadog.NullableInt64 and assigns it to the IngestedEventsBytes field.
func (o *UsageLogsHour) SetIngestedEventsBytes(v int64) {
	o.IngestedEventsBytes.Set(&v)
}

// SetIngestedEventsBytesNil sets the value for IngestedEventsBytes to be an explicit nil.
func (o *UsageLogsHour) SetIngestedEventsBytesNil() {
	o.IngestedEventsBytes.Set(nil)
}

// UnsetIngestedEventsBytes ensures that no value is present for IngestedEventsBytes, not even an explicit nil.
func (o *UsageLogsHour) UnsetIngestedEventsBytes() {
	o.IngestedEventsBytes.Unset()
}

// GetLogsForwardingEventsBytes returns the LogsForwardingEventsBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLogsHour) GetLogsForwardingEventsBytes() int64 {
	if o == nil || o.LogsForwardingEventsBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogsForwardingEventsBytes.Get()
}

// GetLogsForwardingEventsBytesOk returns a tuple with the LogsForwardingEventsBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLogsHour) GetLogsForwardingEventsBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsForwardingEventsBytes.Get(), o.LogsForwardingEventsBytes.IsSet()
}

// HasLogsForwardingEventsBytes returns a boolean if a field has been set.
func (o *UsageLogsHour) HasLogsForwardingEventsBytes() bool {
	return o != nil && o.LogsForwardingEventsBytes.IsSet()
}

// SetLogsForwardingEventsBytes gets a reference to the given datadog.NullableInt64 and assigns it to the LogsForwardingEventsBytes field.
func (o *UsageLogsHour) SetLogsForwardingEventsBytes(v int64) {
	o.LogsForwardingEventsBytes.Set(&v)
}

// SetLogsForwardingEventsBytesNil sets the value for LogsForwardingEventsBytes to be an explicit nil.
func (o *UsageLogsHour) SetLogsForwardingEventsBytesNil() {
	o.LogsForwardingEventsBytes.Set(nil)
}

// UnsetLogsForwardingEventsBytes ensures that no value is present for LogsForwardingEventsBytes, not even an explicit nil.
func (o *UsageLogsHour) UnsetLogsForwardingEventsBytes() {
	o.LogsForwardingEventsBytes.Unset()
}

// GetLogsLiveIndexedCount returns the LogsLiveIndexedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLogsHour) GetLogsLiveIndexedCount() int64 {
	if o == nil || o.LogsLiveIndexedCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogsLiveIndexedCount.Get()
}

// GetLogsLiveIndexedCountOk returns a tuple with the LogsLiveIndexedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLogsHour) GetLogsLiveIndexedCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsLiveIndexedCount.Get(), o.LogsLiveIndexedCount.IsSet()
}

// HasLogsLiveIndexedCount returns a boolean if a field has been set.
func (o *UsageLogsHour) HasLogsLiveIndexedCount() bool {
	return o != nil && o.LogsLiveIndexedCount.IsSet()
}

// SetLogsLiveIndexedCount gets a reference to the given datadog.NullableInt64 and assigns it to the LogsLiveIndexedCount field.
func (o *UsageLogsHour) SetLogsLiveIndexedCount(v int64) {
	o.LogsLiveIndexedCount.Set(&v)
}

// SetLogsLiveIndexedCountNil sets the value for LogsLiveIndexedCount to be an explicit nil.
func (o *UsageLogsHour) SetLogsLiveIndexedCountNil() {
	o.LogsLiveIndexedCount.Set(nil)
}

// UnsetLogsLiveIndexedCount ensures that no value is present for LogsLiveIndexedCount, not even an explicit nil.
func (o *UsageLogsHour) UnsetLogsLiveIndexedCount() {
	o.LogsLiveIndexedCount.Unset()
}

// GetLogsLiveIngestedBytes returns the LogsLiveIngestedBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLogsHour) GetLogsLiveIngestedBytes() int64 {
	if o == nil || o.LogsLiveIngestedBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogsLiveIngestedBytes.Get()
}

// GetLogsLiveIngestedBytesOk returns a tuple with the LogsLiveIngestedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLogsHour) GetLogsLiveIngestedBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsLiveIngestedBytes.Get(), o.LogsLiveIngestedBytes.IsSet()
}

// HasLogsLiveIngestedBytes returns a boolean if a field has been set.
func (o *UsageLogsHour) HasLogsLiveIngestedBytes() bool {
	return o != nil && o.LogsLiveIngestedBytes.IsSet()
}

// SetLogsLiveIngestedBytes gets a reference to the given datadog.NullableInt64 and assigns it to the LogsLiveIngestedBytes field.
func (o *UsageLogsHour) SetLogsLiveIngestedBytes(v int64) {
	o.LogsLiveIngestedBytes.Set(&v)
}

// SetLogsLiveIngestedBytesNil sets the value for LogsLiveIngestedBytes to be an explicit nil.
func (o *UsageLogsHour) SetLogsLiveIngestedBytesNil() {
	o.LogsLiveIngestedBytes.Set(nil)
}

// UnsetLogsLiveIngestedBytes ensures that no value is present for LogsLiveIngestedBytes, not even an explicit nil.
func (o *UsageLogsHour) UnsetLogsLiveIngestedBytes() {
	o.LogsLiveIngestedBytes.Unset()
}

// GetLogsRehydratedIndexedCount returns the LogsRehydratedIndexedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLogsHour) GetLogsRehydratedIndexedCount() int64 {
	if o == nil || o.LogsRehydratedIndexedCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogsRehydratedIndexedCount.Get()
}

// GetLogsRehydratedIndexedCountOk returns a tuple with the LogsRehydratedIndexedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLogsHour) GetLogsRehydratedIndexedCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsRehydratedIndexedCount.Get(), o.LogsRehydratedIndexedCount.IsSet()
}

// HasLogsRehydratedIndexedCount returns a boolean if a field has been set.
func (o *UsageLogsHour) HasLogsRehydratedIndexedCount() bool {
	return o != nil && o.LogsRehydratedIndexedCount.IsSet()
}

// SetLogsRehydratedIndexedCount gets a reference to the given datadog.NullableInt64 and assigns it to the LogsRehydratedIndexedCount field.
func (o *UsageLogsHour) SetLogsRehydratedIndexedCount(v int64) {
	o.LogsRehydratedIndexedCount.Set(&v)
}

// SetLogsRehydratedIndexedCountNil sets the value for LogsRehydratedIndexedCount to be an explicit nil.
func (o *UsageLogsHour) SetLogsRehydratedIndexedCountNil() {
	o.LogsRehydratedIndexedCount.Set(nil)
}

// UnsetLogsRehydratedIndexedCount ensures that no value is present for LogsRehydratedIndexedCount, not even an explicit nil.
func (o *UsageLogsHour) UnsetLogsRehydratedIndexedCount() {
	o.LogsRehydratedIndexedCount.Unset()
}

// GetLogsRehydratedIngestedBytes returns the LogsRehydratedIngestedBytes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageLogsHour) GetLogsRehydratedIngestedBytes() int64 {
	if o == nil || o.LogsRehydratedIngestedBytes.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LogsRehydratedIngestedBytes.Get()
}

// GetLogsRehydratedIngestedBytesOk returns a tuple with the LogsRehydratedIngestedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageLogsHour) GetLogsRehydratedIngestedBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogsRehydratedIngestedBytes.Get(), o.LogsRehydratedIngestedBytes.IsSet()
}

// HasLogsRehydratedIngestedBytes returns a boolean if a field has been set.
func (o *UsageLogsHour) HasLogsRehydratedIngestedBytes() bool {
	return o != nil && o.LogsRehydratedIngestedBytes.IsSet()
}

// SetLogsRehydratedIngestedBytes gets a reference to the given datadog.NullableInt64 and assigns it to the LogsRehydratedIngestedBytes field.
func (o *UsageLogsHour) SetLogsRehydratedIngestedBytes(v int64) {
	o.LogsRehydratedIngestedBytes.Set(&v)
}

// SetLogsRehydratedIngestedBytesNil sets the value for LogsRehydratedIngestedBytes to be an explicit nil.
func (o *UsageLogsHour) SetLogsRehydratedIngestedBytesNil() {
	o.LogsRehydratedIngestedBytes.Set(nil)
}

// UnsetLogsRehydratedIngestedBytes ensures that no value is present for LogsRehydratedIngestedBytes, not even an explicit nil.
func (o *UsageLogsHour) UnsetLogsRehydratedIngestedBytes() {
	o.LogsRehydratedIngestedBytes.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageLogsHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLogsHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageLogsHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageLogsHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageLogsHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageLogsHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageLogsHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageLogsHour) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageLogsHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BillableIngestedBytes.IsSet() {
		toSerialize["billable_ingested_bytes"] = o.BillableIngestedBytes.Get()
	}
	if o.Hour != nil {
		if o.Hour.Nanosecond() == 0 {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["hour"] = o.Hour.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.IndexedEventsCount.IsSet() {
		toSerialize["indexed_events_count"] = o.IndexedEventsCount.Get()
	}
	if o.IngestedEventsBytes.IsSet() {
		toSerialize["ingested_events_bytes"] = o.IngestedEventsBytes.Get()
	}
	if o.LogsForwardingEventsBytes.IsSet() {
		toSerialize["logs_forwarding_events_bytes"] = o.LogsForwardingEventsBytes.Get()
	}
	if o.LogsLiveIndexedCount.IsSet() {
		toSerialize["logs_live_indexed_count"] = o.LogsLiveIndexedCount.Get()
	}
	if o.LogsLiveIngestedBytes.IsSet() {
		toSerialize["logs_live_ingested_bytes"] = o.LogsLiveIngestedBytes.Get()
	}
	if o.LogsRehydratedIndexedCount.IsSet() {
		toSerialize["logs_rehydrated_indexed_count"] = o.LogsRehydratedIndexedCount.Get()
	}
	if o.LogsRehydratedIngestedBytes.IsSet() {
		toSerialize["logs_rehydrated_ingested_bytes"] = o.LogsRehydratedIngestedBytes.Get()
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
func (o *UsageLogsHour) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BillableIngestedBytes       datadog.NullableInt64 `json:"billable_ingested_bytes,omitempty"`
		Hour                        *time.Time            `json:"hour,omitempty"`
		IndexedEventsCount          datadog.NullableInt64 `json:"indexed_events_count,omitempty"`
		IngestedEventsBytes         datadog.NullableInt64 `json:"ingested_events_bytes,omitempty"`
		LogsForwardingEventsBytes   datadog.NullableInt64 `json:"logs_forwarding_events_bytes,omitempty"`
		LogsLiveIndexedCount        datadog.NullableInt64 `json:"logs_live_indexed_count,omitempty"`
		LogsLiveIngestedBytes       datadog.NullableInt64 `json:"logs_live_ingested_bytes,omitempty"`
		LogsRehydratedIndexedCount  datadog.NullableInt64 `json:"logs_rehydrated_indexed_count,omitempty"`
		LogsRehydratedIngestedBytes datadog.NullableInt64 `json:"logs_rehydrated_ingested_bytes,omitempty"`
		OrgName                     *string               `json:"org_name,omitempty"`
		PublicId                    *string               `json:"public_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"billable_ingested_bytes", "hour", "indexed_events_count", "ingested_events_bytes", "logs_forwarding_events_bytes", "logs_live_indexed_count", "logs_live_ingested_bytes", "logs_rehydrated_indexed_count", "logs_rehydrated_ingested_bytes", "org_name", "public_id"})
	} else {
		return err
	}
	o.BillableIngestedBytes = all.BillableIngestedBytes
	o.Hour = all.Hour
	o.IndexedEventsCount = all.IndexedEventsCount
	o.IngestedEventsBytes = all.IngestedEventsBytes
	o.LogsForwardingEventsBytes = all.LogsForwardingEventsBytes
	o.LogsLiveIndexedCount = all.LogsLiveIndexedCount
	o.LogsLiveIngestedBytes = all.LogsLiveIngestedBytes
	o.LogsRehydratedIndexedCount = all.LogsRehydratedIndexedCount
	o.LogsRehydratedIngestedBytes = all.LogsRehydratedIngestedBytes
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
