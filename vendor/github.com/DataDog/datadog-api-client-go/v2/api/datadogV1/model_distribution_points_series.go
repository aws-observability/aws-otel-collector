// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DistributionPointsSeries A distribution points metric to submit to Datadog.
type DistributionPointsSeries struct {
	// The name of the host that produced the distribution point metric.
	Host *string `json:"host,omitempty"`
	// The name of the distribution points metric.
	Metric string `json:"metric"`
	// Points relating to the distribution point metric. All points must be tuples with timestamp and a list of values (cannot be a string). Timestamps should be in POSIX time in seconds.
	Points [][]DistributionPointItem `json:"points"`
	// A list of tags associated with the distribution point metric.
	Tags []string `json:"tags,omitempty"`
	// The type of the distribution point.
	Type *DistributionPointsType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDistributionPointsSeries instantiates a new DistributionPointsSeries object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDistributionPointsSeries(metric string, points [][]DistributionPointItem) *DistributionPointsSeries {
	this := DistributionPointsSeries{}
	this.Metric = metric
	this.Points = points
	var typeVar DistributionPointsType = DISTRIBUTIONPOINTSTYPE_DISTRIBUTION
	this.Type = &typeVar
	return &this
}

// NewDistributionPointsSeriesWithDefaults instantiates a new DistributionPointsSeries object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDistributionPointsSeriesWithDefaults() *DistributionPointsSeries {
	this := DistributionPointsSeries{}
	var typeVar DistributionPointsType = DISTRIBUTIONPOINTSTYPE_DISTRIBUTION
	this.Type = &typeVar
	return &this
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *DistributionPointsSeries) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionPointsSeries) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *DistributionPointsSeries) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *DistributionPointsSeries) SetHost(v string) {
	o.Host = &v
}

// GetMetric returns the Metric field value.
func (o *DistributionPointsSeries) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *DistributionPointsSeries) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value.
func (o *DistributionPointsSeries) SetMetric(v string) {
	o.Metric = v
}

// GetPoints returns the Points field value.
func (o *DistributionPointsSeries) GetPoints() [][]DistributionPointItem {
	if o == nil {
		var ret [][]DistributionPointItem
		return ret
	}
	return o.Points
}

// GetPointsOk returns a tuple with the Points field value
// and a boolean to check if the value has been set.
func (o *DistributionPointsSeries) GetPointsOk() (*[][]DistributionPointItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Points, true
}

// SetPoints sets field value.
func (o *DistributionPointsSeries) SetPoints(v [][]DistributionPointItem) {
	o.Points = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DistributionPointsSeries) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionPointsSeries) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DistributionPointsSeries) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DistributionPointsSeries) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DistributionPointsSeries) GetType() DistributionPointsType {
	if o == nil || o.Type == nil {
		var ret DistributionPointsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionPointsSeries) GetTypeOk() (*DistributionPointsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DistributionPointsSeries) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given DistributionPointsType and assigns it to the Type field.
func (o *DistributionPointsSeries) SetType(v DistributionPointsType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DistributionPointsSeries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	toSerialize["metric"] = o.Metric
	toSerialize["points"] = o.Points
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DistributionPointsSeries) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Host   *string                    `json:"host,omitempty"`
		Metric *string                    `json:"metric"`
		Points *[][]DistributionPointItem `json:"points"`
		Tags   []string                   `json:"tags,omitempty"`
		Type   *DistributionPointsType    `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Metric == nil {
		return fmt.Errorf("required field metric missing")
	}
	if all.Points == nil {
		return fmt.Errorf("required field points missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"host", "metric", "points", "tags", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Host = all.Host
	o.Metric = *all.Metric
	o.Points = *all.Points
	o.Tags = all.Tags
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
