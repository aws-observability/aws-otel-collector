// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomCostListResponseMeta Meta for the response from the List Custom Costs endpoints.
type CustomCostListResponseMeta struct {
	// Number of Custom Costs files returned by the List Custom Costs endpoint
	TotalFilteredCount *int64 `json:"total_filtered_count,omitempty"`
	// Version of Custom Costs file
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomCostListResponseMeta instantiates a new CustomCostListResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomCostListResponseMeta() *CustomCostListResponseMeta {
	this := CustomCostListResponseMeta{}
	return &this
}

// NewCustomCostListResponseMetaWithDefaults instantiates a new CustomCostListResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomCostListResponseMetaWithDefaults() *CustomCostListResponseMeta {
	this := CustomCostListResponseMeta{}
	return &this
}

// GetTotalFilteredCount returns the TotalFilteredCount field value if set, zero value otherwise.
func (o *CustomCostListResponseMeta) GetTotalFilteredCount() int64 {
	if o == nil || o.TotalFilteredCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalFilteredCount
}

// GetTotalFilteredCountOk returns a tuple with the TotalFilteredCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostListResponseMeta) GetTotalFilteredCountOk() (*int64, bool) {
	if o == nil || o.TotalFilteredCount == nil {
		return nil, false
	}
	return o.TotalFilteredCount, true
}

// HasTotalFilteredCount returns a boolean if a field has been set.
func (o *CustomCostListResponseMeta) HasTotalFilteredCount() bool {
	return o != nil && o.TotalFilteredCount != nil
}

// SetTotalFilteredCount gets a reference to the given int64 and assigns it to the TotalFilteredCount field.
func (o *CustomCostListResponseMeta) SetTotalFilteredCount(v int64) {
	o.TotalFilteredCount = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CustomCostListResponseMeta) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostListResponseMeta) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CustomCostListResponseMeta) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *CustomCostListResponseMeta) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomCostListResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.TotalFilteredCount != nil {
		toSerialize["total_filtered_count"] = o.TotalFilteredCount
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
func (o *CustomCostListResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		TotalFilteredCount *int64  `json:"total_filtered_count,omitempty"`
		Version            *string `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"total_filtered_count", "version"})
	} else {
		return err
	}
	o.TotalFilteredCount = all.TotalFilteredCount
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
