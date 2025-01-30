// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumMetricCompute The compute rule to compute the rum-based metric.
type RumMetricCompute struct {
	// The type of aggregation to use.
	AggregationType RumMetricComputeAggregationType `json:"aggregation_type"`
	// Toggle to include or exclude percentile aggregations for distribution metrics.
	// Only present when `aggregation_type` is `distribution`.
	IncludePercentiles *bool `json:"include_percentiles,omitempty"`
	// The path to the value the rum-based metric will aggregate on.
	// Only present when `aggregation_type` is `distribution`.
	Path *string `json:"path,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumMetricCompute instantiates a new RumMetricCompute object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumMetricCompute(aggregationType RumMetricComputeAggregationType) *RumMetricCompute {
	this := RumMetricCompute{}
	this.AggregationType = aggregationType
	return &this
}

// NewRumMetricComputeWithDefaults instantiates a new RumMetricCompute object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumMetricComputeWithDefaults() *RumMetricCompute {
	this := RumMetricCompute{}
	return &this
}

// GetAggregationType returns the AggregationType field value.
func (o *RumMetricCompute) GetAggregationType() RumMetricComputeAggregationType {
	if o == nil {
		var ret RumMetricComputeAggregationType
		return ret
	}
	return o.AggregationType
}

// GetAggregationTypeOk returns a tuple with the AggregationType field value
// and a boolean to check if the value has been set.
func (o *RumMetricCompute) GetAggregationTypeOk() (*RumMetricComputeAggregationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AggregationType, true
}

// SetAggregationType sets field value.
func (o *RumMetricCompute) SetAggregationType(v RumMetricComputeAggregationType) {
	o.AggregationType = v
}

// GetIncludePercentiles returns the IncludePercentiles field value if set, zero value otherwise.
func (o *RumMetricCompute) GetIncludePercentiles() bool {
	if o == nil || o.IncludePercentiles == nil {
		var ret bool
		return ret
	}
	return *o.IncludePercentiles
}

// GetIncludePercentilesOk returns a tuple with the IncludePercentiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumMetricCompute) GetIncludePercentilesOk() (*bool, bool) {
	if o == nil || o.IncludePercentiles == nil {
		return nil, false
	}
	return o.IncludePercentiles, true
}

// HasIncludePercentiles returns a boolean if a field has been set.
func (o *RumMetricCompute) HasIncludePercentiles() bool {
	return o != nil && o.IncludePercentiles != nil
}

// SetIncludePercentiles gets a reference to the given bool and assigns it to the IncludePercentiles field.
func (o *RumMetricCompute) SetIncludePercentiles(v bool) {
	o.IncludePercentiles = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *RumMetricCompute) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumMetricCompute) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *RumMetricCompute) HasPath() bool {
	return o != nil && o.Path != nil
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *RumMetricCompute) SetPath(v string) {
	o.Path = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumMetricCompute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["aggregation_type"] = o.AggregationType
	if o.IncludePercentiles != nil {
		toSerialize["include_percentiles"] = o.IncludePercentiles
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RumMetricCompute) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AggregationType    *RumMetricComputeAggregationType `json:"aggregation_type"`
		IncludePercentiles *bool                            `json:"include_percentiles,omitempty"`
		Path               *string                          `json:"path,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AggregationType == nil {
		return fmt.Errorf("required field aggregation_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aggregation_type", "include_percentiles", "path"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.AggregationType.IsValid() {
		hasInvalidField = true
	} else {
		o.AggregationType = *all.AggregationType
	}
	o.IncludePercentiles = all.IncludePercentiles
	o.Path = all.Path

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
