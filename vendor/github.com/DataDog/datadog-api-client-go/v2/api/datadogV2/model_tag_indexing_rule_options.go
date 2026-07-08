// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleOptions Versioned configuration options for a tag indexing rule.
type TagIndexingRuleOptions struct {
	// Data payload for tag indexing rule options.
	Data *TagIndexingRuleOptionsData `json:"data,omitempty"`
	// Options schema version. Only `1` is supported.
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagIndexingRuleOptions instantiates a new TagIndexingRuleOptions object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagIndexingRuleOptions() *TagIndexingRuleOptions {
	this := TagIndexingRuleOptions{}
	return &this
}

// NewTagIndexingRuleOptionsWithDefaults instantiates a new TagIndexingRuleOptions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagIndexingRuleOptionsWithDefaults() *TagIndexingRuleOptions {
	this := TagIndexingRuleOptions{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TagIndexingRuleOptions) GetData() TagIndexingRuleOptionsData {
	if o == nil || o.Data == nil {
		var ret TagIndexingRuleOptionsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleOptions) GetDataOk() (*TagIndexingRuleOptionsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TagIndexingRuleOptions) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given TagIndexingRuleOptionsData and assigns it to the Data field.
func (o *TagIndexingRuleOptions) SetData(v TagIndexingRuleOptionsData) {
	o.Data = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TagIndexingRuleOptions) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleOptions) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TagIndexingRuleOptions) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *TagIndexingRuleOptions) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagIndexingRuleOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagIndexingRuleOptions) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data    *TagIndexingRuleOptionsData `json:"data,omitempty"`
		Version *int64                      `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
