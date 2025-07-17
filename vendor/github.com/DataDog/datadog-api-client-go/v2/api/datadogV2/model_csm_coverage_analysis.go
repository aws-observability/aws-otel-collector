// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmCoverageAnalysis CSM Coverage Analysis.
type CsmCoverageAnalysis struct {
	// The number of fully configured resources.
	ConfiguredResourcesCount *int64 `json:"configured_resources_count,omitempty"`
	// The coverage percentage.
	Coverage *float64 `json:"coverage,omitempty"`
	// The number of partially configured resources.
	PartiallyConfiguredResourcesCount *int64 `json:"partially_configured_resources_count,omitempty"`
	// The total number of resources.
	TotalResourcesCount *int64 `json:"total_resources_count,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCsmCoverageAnalysis instantiates a new CsmCoverageAnalysis object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCsmCoverageAnalysis() *CsmCoverageAnalysis {
	this := CsmCoverageAnalysis{}
	return &this
}

// NewCsmCoverageAnalysisWithDefaults instantiates a new CsmCoverageAnalysis object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCsmCoverageAnalysisWithDefaults() *CsmCoverageAnalysis {
	this := CsmCoverageAnalysis{}
	return &this
}

// GetConfiguredResourcesCount returns the ConfiguredResourcesCount field value if set, zero value otherwise.
func (o *CsmCoverageAnalysis) GetConfiguredResourcesCount() int64 {
	if o == nil || o.ConfiguredResourcesCount == nil {
		var ret int64
		return ret
	}
	return *o.ConfiguredResourcesCount
}

// GetConfiguredResourcesCountOk returns a tuple with the ConfiguredResourcesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmCoverageAnalysis) GetConfiguredResourcesCountOk() (*int64, bool) {
	if o == nil || o.ConfiguredResourcesCount == nil {
		return nil, false
	}
	return o.ConfiguredResourcesCount, true
}

// HasConfiguredResourcesCount returns a boolean if a field has been set.
func (o *CsmCoverageAnalysis) HasConfiguredResourcesCount() bool {
	return o != nil && o.ConfiguredResourcesCount != nil
}

// SetConfiguredResourcesCount gets a reference to the given int64 and assigns it to the ConfiguredResourcesCount field.
func (o *CsmCoverageAnalysis) SetConfiguredResourcesCount(v int64) {
	o.ConfiguredResourcesCount = &v
}

// GetCoverage returns the Coverage field value if set, zero value otherwise.
func (o *CsmCoverageAnalysis) GetCoverage() float64 {
	if o == nil || o.Coverage == nil {
		var ret float64
		return ret
	}
	return *o.Coverage
}

// GetCoverageOk returns a tuple with the Coverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmCoverageAnalysis) GetCoverageOk() (*float64, bool) {
	if o == nil || o.Coverage == nil {
		return nil, false
	}
	return o.Coverage, true
}

// HasCoverage returns a boolean if a field has been set.
func (o *CsmCoverageAnalysis) HasCoverage() bool {
	return o != nil && o.Coverage != nil
}

// SetCoverage gets a reference to the given float64 and assigns it to the Coverage field.
func (o *CsmCoverageAnalysis) SetCoverage(v float64) {
	o.Coverage = &v
}

// GetPartiallyConfiguredResourcesCount returns the PartiallyConfiguredResourcesCount field value if set, zero value otherwise.
func (o *CsmCoverageAnalysis) GetPartiallyConfiguredResourcesCount() int64 {
	if o == nil || o.PartiallyConfiguredResourcesCount == nil {
		var ret int64
		return ret
	}
	return *o.PartiallyConfiguredResourcesCount
}

// GetPartiallyConfiguredResourcesCountOk returns a tuple with the PartiallyConfiguredResourcesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmCoverageAnalysis) GetPartiallyConfiguredResourcesCountOk() (*int64, bool) {
	if o == nil || o.PartiallyConfiguredResourcesCount == nil {
		return nil, false
	}
	return o.PartiallyConfiguredResourcesCount, true
}

// HasPartiallyConfiguredResourcesCount returns a boolean if a field has been set.
func (o *CsmCoverageAnalysis) HasPartiallyConfiguredResourcesCount() bool {
	return o != nil && o.PartiallyConfiguredResourcesCount != nil
}

// SetPartiallyConfiguredResourcesCount gets a reference to the given int64 and assigns it to the PartiallyConfiguredResourcesCount field.
func (o *CsmCoverageAnalysis) SetPartiallyConfiguredResourcesCount(v int64) {
	o.PartiallyConfiguredResourcesCount = &v
}

// GetTotalResourcesCount returns the TotalResourcesCount field value if set, zero value otherwise.
func (o *CsmCoverageAnalysis) GetTotalResourcesCount() int64 {
	if o == nil || o.TotalResourcesCount == nil {
		var ret int64
		return ret
	}
	return *o.TotalResourcesCount
}

// GetTotalResourcesCountOk returns a tuple with the TotalResourcesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmCoverageAnalysis) GetTotalResourcesCountOk() (*int64, bool) {
	if o == nil || o.TotalResourcesCount == nil {
		return nil, false
	}
	return o.TotalResourcesCount, true
}

// HasTotalResourcesCount returns a boolean if a field has been set.
func (o *CsmCoverageAnalysis) HasTotalResourcesCount() bool {
	return o != nil && o.TotalResourcesCount != nil
}

// SetTotalResourcesCount gets a reference to the given int64 and assigns it to the TotalResourcesCount field.
func (o *CsmCoverageAnalysis) SetTotalResourcesCount(v int64) {
	o.TotalResourcesCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CsmCoverageAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConfiguredResourcesCount != nil {
		toSerialize["configured_resources_count"] = o.ConfiguredResourcesCount
	}
	if o.Coverage != nil {
		toSerialize["coverage"] = o.Coverage
	}
	if o.PartiallyConfiguredResourcesCount != nil {
		toSerialize["partially_configured_resources_count"] = o.PartiallyConfiguredResourcesCount
	}
	if o.TotalResourcesCount != nil {
		toSerialize["total_resources_count"] = o.TotalResourcesCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CsmCoverageAnalysis) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfiguredResourcesCount          *int64   `json:"configured_resources_count,omitempty"`
		Coverage                          *float64 `json:"coverage,omitempty"`
		PartiallyConfiguredResourcesCount *int64   `json:"partially_configured_resources_count,omitempty"`
		TotalResourcesCount               *int64   `json:"total_resources_count,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"configured_resources_count", "coverage", "partially_configured_resources_count", "total_resources_count"})
	} else {
		return err
	}
	o.ConfiguredResourcesCount = all.ConfiguredResourcesCount
	o.Coverage = all.Coverage
	o.PartiallyConfiguredResourcesCount = all.PartiallyConfiguredResourcesCount
	o.TotalResourcesCount = all.TotalResourcesCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
