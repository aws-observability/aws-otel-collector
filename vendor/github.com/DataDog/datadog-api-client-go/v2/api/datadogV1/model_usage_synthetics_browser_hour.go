// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSyntheticsBrowserHour Number of Synthetics Browser tests run for each hour for a given organization.
type UsageSyntheticsBrowserHour struct {
	// Contains the number of Synthetics Browser tests run.
	BrowserCheckCallsCount datadog.NullableInt64 `json:"browser_check_calls_count,omitempty"`
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

// NewUsageSyntheticsBrowserHour instantiates a new UsageSyntheticsBrowserHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageSyntheticsBrowserHour() *UsageSyntheticsBrowserHour {
	this := UsageSyntheticsBrowserHour{}
	return &this
}

// NewUsageSyntheticsBrowserHourWithDefaults instantiates a new UsageSyntheticsBrowserHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageSyntheticsBrowserHourWithDefaults() *UsageSyntheticsBrowserHour {
	this := UsageSyntheticsBrowserHour{}
	return &this
}

// GetBrowserCheckCallsCount returns the BrowserCheckCallsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageSyntheticsBrowserHour) GetBrowserCheckCallsCount() int64 {
	if o == nil || o.BrowserCheckCallsCount.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BrowserCheckCallsCount.Get()
}

// GetBrowserCheckCallsCountOk returns a tuple with the BrowserCheckCallsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageSyntheticsBrowserHour) GetBrowserCheckCallsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserCheckCallsCount.Get(), o.BrowserCheckCallsCount.IsSet()
}

// HasBrowserCheckCallsCount returns a boolean if a field has been set.
func (o *UsageSyntheticsBrowserHour) HasBrowserCheckCallsCount() bool {
	return o != nil && o.BrowserCheckCallsCount.IsSet()
}

// SetBrowserCheckCallsCount gets a reference to the given datadog.NullableInt64 and assigns it to the BrowserCheckCallsCount field.
func (o *UsageSyntheticsBrowserHour) SetBrowserCheckCallsCount(v int64) {
	o.BrowserCheckCallsCount.Set(&v)
}

// SetBrowserCheckCallsCountNil sets the value for BrowserCheckCallsCount to be an explicit nil.
func (o *UsageSyntheticsBrowserHour) SetBrowserCheckCallsCountNil() {
	o.BrowserCheckCallsCount.Set(nil)
}

// UnsetBrowserCheckCallsCount ensures that no value is present for BrowserCheckCallsCount, not even an explicit nil.
func (o *UsageSyntheticsBrowserHour) UnsetBrowserCheckCallsCount() {
	o.BrowserCheckCallsCount.Unset()
}

// GetHour returns the Hour field value if set, zero value otherwise.
func (o *UsageSyntheticsBrowserHour) GetHour() time.Time {
	if o == nil || o.Hour == nil {
		var ret time.Time
		return ret
	}
	return *o.Hour
}

// GetHourOk returns a tuple with the Hour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSyntheticsBrowserHour) GetHourOk() (*time.Time, bool) {
	if o == nil || o.Hour == nil {
		return nil, false
	}
	return o.Hour, true
}

// HasHour returns a boolean if a field has been set.
func (o *UsageSyntheticsBrowserHour) HasHour() bool {
	return o != nil && o.Hour != nil
}

// SetHour gets a reference to the given time.Time and assigns it to the Hour field.
func (o *UsageSyntheticsBrowserHour) SetHour(v time.Time) {
	o.Hour = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageSyntheticsBrowserHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSyntheticsBrowserHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageSyntheticsBrowserHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageSyntheticsBrowserHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageSyntheticsBrowserHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSyntheticsBrowserHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageSyntheticsBrowserHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageSyntheticsBrowserHour) SetPublicId(v string) {
	o.PublicId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageSyntheticsBrowserHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BrowserCheckCallsCount.IsSet() {
		toSerialize["browser_check_calls_count"] = o.BrowserCheckCallsCount.Get()
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
func (o *UsageSyntheticsBrowserHour) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BrowserCheckCallsCount datadog.NullableInt64 `json:"browser_check_calls_count,omitempty"`
		Hour                   *time.Time            `json:"hour,omitempty"`
		OrgName                *string               `json:"org_name,omitempty"`
		PublicId               *string               `json:"public_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"browser_check_calls_count", "hour", "org_name", "public_id"})
	} else {
		return err
	}
	o.BrowserCheckCallsCount = all.BrowserCheckCallsCount
	o.Hour = all.Hour
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
