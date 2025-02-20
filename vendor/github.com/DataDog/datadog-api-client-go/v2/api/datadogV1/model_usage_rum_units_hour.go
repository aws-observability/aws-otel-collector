// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageRumUnitsHour Number of RUM Units used for each hour for a given organization (data available as of November 1, 2021).
type UsageRumUnitsHour struct {
	// The number of browser RUM units.
	BrowserRumUnits datadog.NullableInt64 `json:"browser_rum_units,omitempty"`
	// The number of mobile RUM units.
	MobileRumUnits datadog.NullableInt64 `json:"mobile_rum_units,omitempty"`
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// Total RUM units across mobile and browser RUM.
	RumUnits datadog.NullableInt64 `json:"rum_units,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageRumUnitsHour instantiates a new UsageRumUnitsHour object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageRumUnitsHour() *UsageRumUnitsHour {
	this := UsageRumUnitsHour{}
	return &this
}

// NewUsageRumUnitsHourWithDefaults instantiates a new UsageRumUnitsHour object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageRumUnitsHourWithDefaults() *UsageRumUnitsHour {
	this := UsageRumUnitsHour{}
	return &this
}

// GetBrowserRumUnits returns the BrowserRumUnits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageRumUnitsHour) GetBrowserRumUnits() int64 {
	if o == nil || o.BrowserRumUnits.Get() == nil {
		var ret int64
		return ret
	}
	return *o.BrowserRumUnits.Get()
}

// GetBrowserRumUnitsOk returns a tuple with the BrowserRumUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageRumUnitsHour) GetBrowserRumUnitsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BrowserRumUnits.Get(), o.BrowserRumUnits.IsSet()
}

// HasBrowserRumUnits returns a boolean if a field has been set.
func (o *UsageRumUnitsHour) HasBrowserRumUnits() bool {
	return o != nil && o.BrowserRumUnits.IsSet()
}

// SetBrowserRumUnits gets a reference to the given datadog.NullableInt64 and assigns it to the BrowserRumUnits field.
func (o *UsageRumUnitsHour) SetBrowserRumUnits(v int64) {
	o.BrowserRumUnits.Set(&v)
}

// SetBrowserRumUnitsNil sets the value for BrowserRumUnits to be an explicit nil.
func (o *UsageRumUnitsHour) SetBrowserRumUnitsNil() {
	o.BrowserRumUnits.Set(nil)
}

// UnsetBrowserRumUnits ensures that no value is present for BrowserRumUnits, not even an explicit nil.
func (o *UsageRumUnitsHour) UnsetBrowserRumUnits() {
	o.BrowserRumUnits.Unset()
}

// GetMobileRumUnits returns the MobileRumUnits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageRumUnitsHour) GetMobileRumUnits() int64 {
	if o == nil || o.MobileRumUnits.Get() == nil {
		var ret int64
		return ret
	}
	return *o.MobileRumUnits.Get()
}

// GetMobileRumUnitsOk returns a tuple with the MobileRumUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageRumUnitsHour) GetMobileRumUnitsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MobileRumUnits.Get(), o.MobileRumUnits.IsSet()
}

// HasMobileRumUnits returns a boolean if a field has been set.
func (o *UsageRumUnitsHour) HasMobileRumUnits() bool {
	return o != nil && o.MobileRumUnits.IsSet()
}

// SetMobileRumUnits gets a reference to the given datadog.NullableInt64 and assigns it to the MobileRumUnits field.
func (o *UsageRumUnitsHour) SetMobileRumUnits(v int64) {
	o.MobileRumUnits.Set(&v)
}

// SetMobileRumUnitsNil sets the value for MobileRumUnits to be an explicit nil.
func (o *UsageRumUnitsHour) SetMobileRumUnitsNil() {
	o.MobileRumUnits.Set(nil)
}

// UnsetMobileRumUnits ensures that no value is present for MobileRumUnits, not even an explicit nil.
func (o *UsageRumUnitsHour) UnsetMobileRumUnits() {
	o.MobileRumUnits.Unset()
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageRumUnitsHour) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageRumUnitsHour) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageRumUnitsHour) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageRumUnitsHour) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageRumUnitsHour) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageRumUnitsHour) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageRumUnitsHour) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageRumUnitsHour) SetPublicId(v string) {
	o.PublicId = &v
}

// GetRumUnits returns the RumUnits field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UsageRumUnitsHour) GetRumUnits() int64 {
	if o == nil || o.RumUnits.Get() == nil {
		var ret int64
		return ret
	}
	return *o.RumUnits.Get()
}

// GetRumUnitsOk returns a tuple with the RumUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UsageRumUnitsHour) GetRumUnitsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RumUnits.Get(), o.RumUnits.IsSet()
}

// HasRumUnits returns a boolean if a field has been set.
func (o *UsageRumUnitsHour) HasRumUnits() bool {
	return o != nil && o.RumUnits.IsSet()
}

// SetRumUnits gets a reference to the given datadog.NullableInt64 and assigns it to the RumUnits field.
func (o *UsageRumUnitsHour) SetRumUnits(v int64) {
	o.RumUnits.Set(&v)
}

// SetRumUnitsNil sets the value for RumUnits to be an explicit nil.
func (o *UsageRumUnitsHour) SetRumUnitsNil() {
	o.RumUnits.Set(nil)
}

// UnsetRumUnits ensures that no value is present for RumUnits, not even an explicit nil.
func (o *UsageRumUnitsHour) UnsetRumUnits() {
	o.RumUnits.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageRumUnitsHour) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BrowserRumUnits.IsSet() {
		toSerialize["browser_rum_units"] = o.BrowserRumUnits.Get()
	}
	if o.MobileRumUnits.IsSet() {
		toSerialize["mobile_rum_units"] = o.MobileRumUnits.Get()
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.RumUnits.IsSet() {
		toSerialize["rum_units"] = o.RumUnits.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageRumUnitsHour) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BrowserRumUnits datadog.NullableInt64 `json:"browser_rum_units,omitempty"`
		MobileRumUnits  datadog.NullableInt64 `json:"mobile_rum_units,omitempty"`
		OrgName         *string               `json:"org_name,omitempty"`
		PublicId        *string               `json:"public_id,omitempty"`
		RumUnits        datadog.NullableInt64 `json:"rum_units,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"browser_rum_units", "mobile_rum_units", "org_name", "public_id", "rum_units"})
	} else {
		return err
	}
	o.BrowserRumUnits = all.BrowserRumUnits
	o.MobileRumUnits = all.MobileRumUnits
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId
	o.RumUnits = all.RumUnits

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
