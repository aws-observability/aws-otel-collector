// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityFilterVersionEntry A single security filter as it existed at a given configuration version.
type SecurityFilterVersionEntry struct {
	// The list of exclusion filters applied in this security filter.
	ExclusionFilters []SecurityFilterExclusionFilterResponse `json:"exclusion_filters"`
	// The filtered data type.
	FilteredDataType SecurityFilterFilteredDataType `json:"filtered_data_type"`
	// The ID of the security filter.
	Id string `json:"id"`
	// Whether the security filter is the built-in filter.
	IsBuiltin bool `json:"is_builtin"`
	// Whether the security filter is enabled.
	IsEnabled bool `json:"is_enabled"`
	// The name of the security filter.
	Name string `json:"name"`
	// The query of the security filter.
	Query string `json:"query"`
	// The version of this security filter.
	Version int32 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityFilterVersionEntry instantiates a new SecurityFilterVersionEntry object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityFilterVersionEntry(exclusionFilters []SecurityFilterExclusionFilterResponse, filteredDataType SecurityFilterFilteredDataType, id string, isBuiltin bool, isEnabled bool, name string, query string, version int32) *SecurityFilterVersionEntry {
	this := SecurityFilterVersionEntry{}
	this.ExclusionFilters = exclusionFilters
	this.FilteredDataType = filteredDataType
	this.Id = id
	this.IsBuiltin = isBuiltin
	this.IsEnabled = isEnabled
	this.Name = name
	this.Query = query
	this.Version = version
	return &this
}

// NewSecurityFilterVersionEntryWithDefaults instantiates a new SecurityFilterVersionEntry object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityFilterVersionEntryWithDefaults() *SecurityFilterVersionEntry {
	this := SecurityFilterVersionEntry{}
	return &this
}

// GetExclusionFilters returns the ExclusionFilters field value.
func (o *SecurityFilterVersionEntry) GetExclusionFilters() []SecurityFilterExclusionFilterResponse {
	if o == nil {
		var ret []SecurityFilterExclusionFilterResponse
		return ret
	}
	return o.ExclusionFilters
}

// GetExclusionFiltersOk returns a tuple with the ExclusionFilters field value
// and a boolean to check if the value has been set.
func (o *SecurityFilterVersionEntry) GetExclusionFiltersOk() (*[]SecurityFilterExclusionFilterResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExclusionFilters, true
}

// SetExclusionFilters sets field value.
func (o *SecurityFilterVersionEntry) SetExclusionFilters(v []SecurityFilterExclusionFilterResponse) {
	o.ExclusionFilters = v
}

// GetFilteredDataType returns the FilteredDataType field value.
func (o *SecurityFilterVersionEntry) GetFilteredDataType() SecurityFilterFilteredDataType {
	if o == nil {
		var ret SecurityFilterFilteredDataType
		return ret
	}
	return o.FilteredDataType
}

// GetFilteredDataTypeOk returns a tuple with the FilteredDataType field value
// and a boolean to check if the value has been set.
func (o *SecurityFilterVersionEntry) GetFilteredDataTypeOk() (*SecurityFilterFilteredDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilteredDataType, true
}

// SetFilteredDataType sets field value.
func (o *SecurityFilterVersionEntry) SetFilteredDataType(v SecurityFilterFilteredDataType) {
	o.FilteredDataType = v
}

// GetId returns the Id field value.
func (o *SecurityFilterVersionEntry) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SecurityFilterVersionEntry) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *SecurityFilterVersionEntry) SetId(v string) {
	o.Id = v
}

// GetIsBuiltin returns the IsBuiltin field value.
func (o *SecurityFilterVersionEntry) GetIsBuiltin() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsBuiltin
}

// GetIsBuiltinOk returns a tuple with the IsBuiltin field value
// and a boolean to check if the value has been set.
func (o *SecurityFilterVersionEntry) GetIsBuiltinOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBuiltin, true
}

// SetIsBuiltin sets field value.
func (o *SecurityFilterVersionEntry) SetIsBuiltin(v bool) {
	o.IsBuiltin = v
}

// GetIsEnabled returns the IsEnabled field value.
func (o *SecurityFilterVersionEntry) GetIsEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *SecurityFilterVersionEntry) GetIsEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value.
func (o *SecurityFilterVersionEntry) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetName returns the Name field value.
func (o *SecurityFilterVersionEntry) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SecurityFilterVersionEntry) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SecurityFilterVersionEntry) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value.
func (o *SecurityFilterVersionEntry) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *SecurityFilterVersionEntry) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *SecurityFilterVersionEntry) SetQuery(v string) {
	o.Query = v
}

// GetVersion returns the Version field value.
func (o *SecurityFilterVersionEntry) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SecurityFilterVersionEntry) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *SecurityFilterVersionEntry) SetVersion(v int32) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityFilterVersionEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["exclusion_filters"] = o.ExclusionFilters
	toSerialize["filtered_data_type"] = o.FilteredDataType
	toSerialize["id"] = o.Id
	toSerialize["is_builtin"] = o.IsBuiltin
	toSerialize["is_enabled"] = o.IsEnabled
	toSerialize["name"] = o.Name
	toSerialize["query"] = o.Query
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityFilterVersionEntry) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ExclusionFilters *[]SecurityFilterExclusionFilterResponse `json:"exclusion_filters"`
		FilteredDataType *SecurityFilterFilteredDataType          `json:"filtered_data_type"`
		Id               *string                                  `json:"id"`
		IsBuiltin        *bool                                    `json:"is_builtin"`
		IsEnabled        *bool                                    `json:"is_enabled"`
		Name             *string                                  `json:"name"`
		Query            *string                                  `json:"query"`
		Version          *int32                                   `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ExclusionFilters == nil {
		return fmt.Errorf("required field exclusion_filters missing")
	}
	if all.FilteredDataType == nil {
		return fmt.Errorf("required field filtered_data_type missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.IsBuiltin == nil {
		return fmt.Errorf("required field is_builtin missing")
	}
	if all.IsEnabled == nil {
		return fmt.Errorf("required field is_enabled missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exclusion_filters", "filtered_data_type", "id", "is_builtin", "is_enabled", "name", "query", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ExclusionFilters = *all.ExclusionFilters
	if !all.FilteredDataType.IsValid() {
		hasInvalidField = true
	} else {
		o.FilteredDataType = *all.FilteredDataType
	}
	o.Id = *all.Id
	o.IsBuiltin = *all.IsBuiltin
	o.IsEnabled = *all.IsEnabled
	o.Name = *all.Name
	o.Query = *all.Query
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
