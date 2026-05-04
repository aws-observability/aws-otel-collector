// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestVersionHistoryMeta Pagination metadata for a version history response.
type SyntheticsTestVersionHistoryMeta struct {
	// The version number to use as the `last_version_number` query parameter
	// to fetch the next page. `null` indicates there are no more pages.
	NextLastVersionNumber datadog.NullableInt64 `json:"next_last_version_number,omitempty"`
	// The number of days that version history is retained.
	RetentionPeriodInDays *int64 `json:"retention_period_in_days,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestVersionHistoryMeta instantiates a new SyntheticsTestVersionHistoryMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestVersionHistoryMeta() *SyntheticsTestVersionHistoryMeta {
	this := SyntheticsTestVersionHistoryMeta{}
	return &this
}

// NewSyntheticsTestVersionHistoryMetaWithDefaults instantiates a new SyntheticsTestVersionHistoryMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestVersionHistoryMetaWithDefaults() *SyntheticsTestVersionHistoryMeta {
	this := SyntheticsTestVersionHistoryMeta{}
	return &this
}

// GetNextLastVersionNumber returns the NextLastVersionNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyntheticsTestVersionHistoryMeta) GetNextLastVersionNumber() int64 {
	if o == nil || o.NextLastVersionNumber.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextLastVersionNumber.Get()
}

// GetNextLastVersionNumberOk returns a tuple with the NextLastVersionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SyntheticsTestVersionHistoryMeta) GetNextLastVersionNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextLastVersionNumber.Get(), o.NextLastVersionNumber.IsSet()
}

// HasNextLastVersionNumber returns a boolean if a field has been set.
func (o *SyntheticsTestVersionHistoryMeta) HasNextLastVersionNumber() bool {
	return o != nil && o.NextLastVersionNumber.IsSet()
}

// SetNextLastVersionNumber gets a reference to the given datadog.NullableInt64 and assigns it to the NextLastVersionNumber field.
func (o *SyntheticsTestVersionHistoryMeta) SetNextLastVersionNumber(v int64) {
	o.NextLastVersionNumber.Set(&v)
}

// SetNextLastVersionNumberNil sets the value for NextLastVersionNumber to be an explicit nil.
func (o *SyntheticsTestVersionHistoryMeta) SetNextLastVersionNumberNil() {
	o.NextLastVersionNumber.Set(nil)
}

// UnsetNextLastVersionNumber ensures that no value is present for NextLastVersionNumber, not even an explicit nil.
func (o *SyntheticsTestVersionHistoryMeta) UnsetNextLastVersionNumber() {
	o.NextLastVersionNumber.Unset()
}

// GetRetentionPeriodInDays returns the RetentionPeriodInDays field value if set, zero value otherwise.
func (o *SyntheticsTestVersionHistoryMeta) GetRetentionPeriodInDays() int64 {
	if o == nil || o.RetentionPeriodInDays == nil {
		var ret int64
		return ret
	}
	return *o.RetentionPeriodInDays
}

// GetRetentionPeriodInDaysOk returns a tuple with the RetentionPeriodInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionHistoryMeta) GetRetentionPeriodInDaysOk() (*int64, bool) {
	if o == nil || o.RetentionPeriodInDays == nil {
		return nil, false
	}
	return o.RetentionPeriodInDays, true
}

// HasRetentionPeriodInDays returns a boolean if a field has been set.
func (o *SyntheticsTestVersionHistoryMeta) HasRetentionPeriodInDays() bool {
	return o != nil && o.RetentionPeriodInDays != nil
}

// SetRetentionPeriodInDays gets a reference to the given int64 and assigns it to the RetentionPeriodInDays field.
func (o *SyntheticsTestVersionHistoryMeta) SetRetentionPeriodInDays(v int64) {
	o.RetentionPeriodInDays = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestVersionHistoryMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NextLastVersionNumber.IsSet() {
		toSerialize["next_last_version_number"] = o.NextLastVersionNumber.Get()
	}
	if o.RetentionPeriodInDays != nil {
		toSerialize["retention_period_in_days"] = o.RetentionPeriodInDays
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestVersionHistoryMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NextLastVersionNumber datadog.NullableInt64 `json:"next_last_version_number,omitempty"`
		RetentionPeriodInDays *int64                `json:"retention_period_in_days,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"next_last_version_number", "retention_period_in_days"})
	} else {
		return err
	}
	o.NextLastVersionNumber = all.NextLastVersionNumber
	o.RetentionPeriodInDays = all.RetentionPeriodInDays

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
