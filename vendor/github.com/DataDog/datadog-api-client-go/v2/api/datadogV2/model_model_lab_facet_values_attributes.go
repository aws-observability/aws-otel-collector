// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabFacetValuesAttributes Available values for a specific facet key.
type ModelLabFacetValuesAttributes struct {
	// The name of the facet.
	FacetName string `json:"facet_name"`
	// The type of the facet.
	FacetType string `json:"facet_type"`
	// The ranges for each metric statistic.
	MetricStatRanges []ModelLabMetricStatRange `json:"metric_stat_ranges,omitempty"`
	// The numeric range of values for a facet.
	NumericRange *ModelLabNumericRange `json:"numeric_range,omitempty"`
	// The list of available string values for this facet.
	Values []string `json:"values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabFacetValuesAttributes instantiates a new ModelLabFacetValuesAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabFacetValuesAttributes(facetName string, facetType string, values []string) *ModelLabFacetValuesAttributes {
	this := ModelLabFacetValuesAttributes{}
	this.FacetName = facetName
	this.FacetType = facetType
	this.Values = values
	return &this
}

// NewModelLabFacetValuesAttributesWithDefaults instantiates a new ModelLabFacetValuesAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabFacetValuesAttributesWithDefaults() *ModelLabFacetValuesAttributes {
	this := ModelLabFacetValuesAttributes{}
	return &this
}

// GetFacetName returns the FacetName field value.
func (o *ModelLabFacetValuesAttributes) GetFacetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FacetName
}

// GetFacetNameOk returns a tuple with the FacetName field value
// and a boolean to check if the value has been set.
func (o *ModelLabFacetValuesAttributes) GetFacetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FacetName, true
}

// SetFacetName sets field value.
func (o *ModelLabFacetValuesAttributes) SetFacetName(v string) {
	o.FacetName = v
}

// GetFacetType returns the FacetType field value.
func (o *ModelLabFacetValuesAttributes) GetFacetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FacetType
}

// GetFacetTypeOk returns a tuple with the FacetType field value
// and a boolean to check if the value has been set.
func (o *ModelLabFacetValuesAttributes) GetFacetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FacetType, true
}

// SetFacetType sets field value.
func (o *ModelLabFacetValuesAttributes) SetFacetType(v string) {
	o.FacetType = v
}

// GetMetricStatRanges returns the MetricStatRanges field value if set, zero value otherwise.
func (o *ModelLabFacetValuesAttributes) GetMetricStatRanges() []ModelLabMetricStatRange {
	if o == nil || o.MetricStatRanges == nil {
		var ret []ModelLabMetricStatRange
		return ret
	}
	return o.MetricStatRanges
}

// GetMetricStatRangesOk returns a tuple with the MetricStatRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLabFacetValuesAttributes) GetMetricStatRangesOk() (*[]ModelLabMetricStatRange, bool) {
	if o == nil || o.MetricStatRanges == nil {
		return nil, false
	}
	return &o.MetricStatRanges, true
}

// HasMetricStatRanges returns a boolean if a field has been set.
func (o *ModelLabFacetValuesAttributes) HasMetricStatRanges() bool {
	return o != nil && o.MetricStatRanges != nil
}

// SetMetricStatRanges gets a reference to the given []ModelLabMetricStatRange and assigns it to the MetricStatRanges field.
func (o *ModelLabFacetValuesAttributes) SetMetricStatRanges(v []ModelLabMetricStatRange) {
	o.MetricStatRanges = v
}

// GetNumericRange returns the NumericRange field value if set, zero value otherwise.
func (o *ModelLabFacetValuesAttributes) GetNumericRange() ModelLabNumericRange {
	if o == nil || o.NumericRange == nil {
		var ret ModelLabNumericRange
		return ret
	}
	return *o.NumericRange
}

// GetNumericRangeOk returns a tuple with the NumericRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLabFacetValuesAttributes) GetNumericRangeOk() (*ModelLabNumericRange, bool) {
	if o == nil || o.NumericRange == nil {
		return nil, false
	}
	return o.NumericRange, true
}

// HasNumericRange returns a boolean if a field has been set.
func (o *ModelLabFacetValuesAttributes) HasNumericRange() bool {
	return o != nil && o.NumericRange != nil
}

// SetNumericRange gets a reference to the given ModelLabNumericRange and assigns it to the NumericRange field.
func (o *ModelLabFacetValuesAttributes) SetNumericRange(v ModelLabNumericRange) {
	o.NumericRange = &v
}

// GetValues returns the Values field value.
func (o *ModelLabFacetValuesAttributes) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ModelLabFacetValuesAttributes) GetValuesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value.
func (o *ModelLabFacetValuesAttributes) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabFacetValuesAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["facet_name"] = o.FacetName
	toSerialize["facet_type"] = o.FacetType
	if o.MetricStatRanges != nil {
		toSerialize["metric_stat_ranges"] = o.MetricStatRanges
	}
	if o.NumericRange != nil {
		toSerialize["numeric_range"] = o.NumericRange
	}
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabFacetValuesAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FacetName        *string                   `json:"facet_name"`
		FacetType        *string                   `json:"facet_type"`
		MetricStatRanges []ModelLabMetricStatRange `json:"metric_stat_ranges,omitempty"`
		NumericRange     *ModelLabNumericRange     `json:"numeric_range,omitempty"`
		Values           *[]string                 `json:"values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FacetName == nil {
		return fmt.Errorf("required field facet_name missing")
	}
	if all.FacetType == nil {
		return fmt.Errorf("required field facet_type missing")
	}
	if all.Values == nil {
		return fmt.Errorf("required field values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"facet_name", "facet_type", "metric_stat_ranges", "numeric_range", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	o.FacetName = *all.FacetName
	o.FacetType = *all.FacetType
	o.MetricStatRanges = all.MetricStatRanges
	if all.NumericRange != nil && all.NumericRange.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.NumericRange = all.NumericRange
	o.Values = *all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
