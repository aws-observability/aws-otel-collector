// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricMetadata Metadata for the metric.
type MetricMetadata struct {
	// Metric origin information.
	Origin *MetricOrigin `json:"origin,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricMetadata instantiates a new MetricMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricMetadata() *MetricMetadata {
	this := MetricMetadata{}
	return &this
}

// NewMetricMetadataWithDefaults instantiates a new MetricMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricMetadataWithDefaults() *MetricMetadata {
	this := MetricMetadata{}
	return &this
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *MetricMetadata) GetOrigin() MetricOrigin {
	if o == nil || o.Origin == nil {
		var ret MetricOrigin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMetadata) GetOriginOk() (*MetricOrigin, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *MetricMetadata) HasOrigin() bool {
	return o != nil && o.Origin != nil
}

// SetOrigin gets a reference to the given MetricOrigin and assigns it to the Origin field.
func (o *MetricMetadata) SetOrigin(v MetricOrigin) {
	o.Origin = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Origin *MetricOrigin `json:"origin,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"origin"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Origin != nil && all.Origin.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Origin = all.Origin

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
