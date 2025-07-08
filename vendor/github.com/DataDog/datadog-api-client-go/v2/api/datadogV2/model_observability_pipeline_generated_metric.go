// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineGeneratedMetric Defines a log-based custom metric, including its name, type, filter, value computation strategy,
// and optional grouping fields.
type ObservabilityPipelineGeneratedMetric struct {
	// Optional fields used to group the metric series.
	GroupBy []string `json:"group_by,omitempty"`
	// Datadog filter query to match logs for metric generation.
	Include string `json:"include"`
	// Type of metric to create.
	MetricType ObservabilityPipelineGeneratedMetricMetricType `json:"metric_type"`
	// Name of the custom metric to be created.
	Name string `json:"name"`
	// Specifies how the value of the generated metric is computed.
	Value ObservabilityPipelineMetricValue `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineGeneratedMetric instantiates a new ObservabilityPipelineGeneratedMetric object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineGeneratedMetric(include string, metricType ObservabilityPipelineGeneratedMetricMetricType, name string, value ObservabilityPipelineMetricValue) *ObservabilityPipelineGeneratedMetric {
	this := ObservabilityPipelineGeneratedMetric{}
	this.Include = include
	this.MetricType = metricType
	this.Name = name
	this.Value = value
	return &this
}

// NewObservabilityPipelineGeneratedMetricWithDefaults instantiates a new ObservabilityPipelineGeneratedMetric object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineGeneratedMetricWithDefaults() *ObservabilityPipelineGeneratedMetric {
	this := ObservabilityPipelineGeneratedMetric{}
	return &this
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *ObservabilityPipelineGeneratedMetric) GetGroupBy() []string {
	if o == nil || o.GroupBy == nil {
		var ret []string
		return ret
	}
	return o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGeneratedMetric) GetGroupByOk() (*[]string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return &o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *ObservabilityPipelineGeneratedMetric) HasGroupBy() bool {
	return o != nil && o.GroupBy != nil
}

// SetGroupBy gets a reference to the given []string and assigns it to the GroupBy field.
func (o *ObservabilityPipelineGeneratedMetric) SetGroupBy(v []string) {
	o.GroupBy = v
}

// GetInclude returns the Include field value.
func (o *ObservabilityPipelineGeneratedMetric) GetInclude() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGeneratedMetric) GetIncludeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value.
func (o *ObservabilityPipelineGeneratedMetric) SetInclude(v string) {
	o.Include = v
}

// GetMetricType returns the MetricType field value.
func (o *ObservabilityPipelineGeneratedMetric) GetMetricType() ObservabilityPipelineGeneratedMetricMetricType {
	if o == nil {
		var ret ObservabilityPipelineGeneratedMetricMetricType
		return ret
	}
	return o.MetricType
}

// GetMetricTypeOk returns a tuple with the MetricType field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGeneratedMetric) GetMetricTypeOk() (*ObservabilityPipelineGeneratedMetricMetricType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricType, true
}

// SetMetricType sets field value.
func (o *ObservabilityPipelineGeneratedMetric) SetMetricType(v ObservabilityPipelineGeneratedMetricMetricType) {
	o.MetricType = v
}

// GetName returns the Name field value.
func (o *ObservabilityPipelineGeneratedMetric) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGeneratedMetric) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ObservabilityPipelineGeneratedMetric) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value.
func (o *ObservabilityPipelineGeneratedMetric) GetValue() ObservabilityPipelineMetricValue {
	if o == nil {
		var ret ObservabilityPipelineMetricValue
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineGeneratedMetric) GetValueOk() (*ObservabilityPipelineMetricValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ObservabilityPipelineGeneratedMetric) SetValue(v ObservabilityPipelineMetricValue) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineGeneratedMetric) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.GroupBy != nil {
		toSerialize["group_by"] = o.GroupBy
	}
	toSerialize["include"] = o.Include
	toSerialize["metric_type"] = o.MetricType
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineGeneratedMetric) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		GroupBy    []string                                        `json:"group_by,omitempty"`
		Include    *string                                         `json:"include"`
		MetricType *ObservabilityPipelineGeneratedMetricMetricType `json:"metric_type"`
		Name       *string                                         `json:"name"`
		Value      *ObservabilityPipelineMetricValue               `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Include == nil {
		return fmt.Errorf("required field include missing")
	}
	if all.MetricType == nil {
		return fmt.Errorf("required field metric_type missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"group_by", "include", "metric_type", "name", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.GroupBy = all.GroupBy
	o.Include = *all.Include
	if !all.MetricType.IsValid() {
		hasInvalidField = true
	} else {
		o.MetricType = *all.MetricType
	}
	o.Name = *all.Name
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
