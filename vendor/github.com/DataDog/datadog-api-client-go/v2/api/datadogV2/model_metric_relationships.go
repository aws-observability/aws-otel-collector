// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricRelationships Relationships for a metric.
type MetricRelationships struct {
	// Relationship to a metric volume included in the response.
	MetricVolumes *MetricVolumesRelationship `json:"metric_volumes,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricRelationships instantiates a new MetricRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricRelationships() *MetricRelationships {
	this := MetricRelationships{}
	return &this
}

// NewMetricRelationshipsWithDefaults instantiates a new MetricRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricRelationshipsWithDefaults() *MetricRelationships {
	this := MetricRelationships{}
	return &this
}

// GetMetricVolumes returns the MetricVolumes field value if set, zero value otherwise.
func (o *MetricRelationships) GetMetricVolumes() MetricVolumesRelationship {
	if o == nil || o.MetricVolumes == nil {
		var ret MetricVolumesRelationship
		return ret
	}
	return *o.MetricVolumes
}

// GetMetricVolumesOk returns a tuple with the MetricVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricRelationships) GetMetricVolumesOk() (*MetricVolumesRelationship, bool) {
	if o == nil || o.MetricVolumes == nil {
		return nil, false
	}
	return o.MetricVolumes, true
}

// HasMetricVolumes returns a boolean if a field has been set.
func (o *MetricRelationships) HasMetricVolumes() bool {
	return o != nil && o.MetricVolumes != nil
}

// SetMetricVolumes gets a reference to the given MetricVolumesRelationship and assigns it to the MetricVolumes field.
func (o *MetricRelationships) SetMetricVolumes(v MetricVolumesRelationship) {
	o.MetricVolumes = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MetricVolumes != nil {
		toSerialize["metric_volumes"] = o.MetricVolumes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MetricVolumes *MetricVolumesRelationship `json:"metric_volumes,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"metric_volumes"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.MetricVolumes != nil && all.MetricVolumes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricVolumes = all.MetricVolumes

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
