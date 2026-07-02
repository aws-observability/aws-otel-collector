// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityFilterVersionAttributes The attributes describing a single security filter configuration version.
type SecurityFilterVersionAttributes struct {
	// The Unix timestamp in milliseconds at which this configuration version was applied.
	Date int64 `json:"date"`
	// The set of security filters at this configuration version.
	Filters []SecurityFilterVersionEntry `json:"filters"`
	// The configuration version number.
	Version int32 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityFilterVersionAttributes instantiates a new SecurityFilterVersionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityFilterVersionAttributes(date int64, filters []SecurityFilterVersionEntry, version int32) *SecurityFilterVersionAttributes {
	this := SecurityFilterVersionAttributes{}
	this.Date = date
	this.Filters = filters
	this.Version = version
	return &this
}

// NewSecurityFilterVersionAttributesWithDefaults instantiates a new SecurityFilterVersionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityFilterVersionAttributesWithDefaults() *SecurityFilterVersionAttributes {
	this := SecurityFilterVersionAttributes{}
	return &this
}

// GetDate returns the Date field value.
func (o *SecurityFilterVersionAttributes) GetDate() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *SecurityFilterVersionAttributes) GetDateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value.
func (o *SecurityFilterVersionAttributes) SetDate(v int64) {
	o.Date = v
}

// GetFilters returns the Filters field value.
func (o *SecurityFilterVersionAttributes) GetFilters() []SecurityFilterVersionEntry {
	if o == nil {
		var ret []SecurityFilterVersionEntry
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *SecurityFilterVersionAttributes) GetFiltersOk() (*[]SecurityFilterVersionEntry, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value.
func (o *SecurityFilterVersionAttributes) SetFilters(v []SecurityFilterVersionEntry) {
	o.Filters = v
}

// GetVersion returns the Version field value.
func (o *SecurityFilterVersionAttributes) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SecurityFilterVersionAttributes) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *SecurityFilterVersionAttributes) SetVersion(v int32) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityFilterVersionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["date"] = o.Date
	toSerialize["filters"] = o.Filters
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityFilterVersionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Date    *int64                        `json:"date"`
		Filters *[]SecurityFilterVersionEntry `json:"filters"`
		Version *int32                        `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Date == nil {
		return fmt.Errorf("required field date missing")
	}
	if all.Filters == nil {
		return fmt.Errorf("required field filters missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"date", "filters", "version"})
	} else {
		return err
	}
	o.Date = *all.Date
	o.Filters = *all.Filters
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
