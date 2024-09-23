// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricAssetResponseData Metric assets response data.
type MetricAssetResponseData struct {
	// The metric name for this resource.
	Id string `json:"id"`
	// Relationships to assets related to the metric.
	Relationships *MetricAssetResponseRelationships `json:"relationships,omitempty"`
	// The metric resource type.
	Type MetricType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricAssetResponseData instantiates a new MetricAssetResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricAssetResponseData(id string, typeVar MetricType) *MetricAssetResponseData {
	this := MetricAssetResponseData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewMetricAssetResponseDataWithDefaults instantiates a new MetricAssetResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricAssetResponseDataWithDefaults() *MetricAssetResponseData {
	this := MetricAssetResponseData{}
	var typeVar MetricType = METRICTYPE_METRICS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *MetricAssetResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MetricAssetResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *MetricAssetResponseData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *MetricAssetResponseData) GetRelationships() MetricAssetResponseRelationships {
	if o == nil || o.Relationships == nil {
		var ret MetricAssetResponseRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAssetResponseData) GetRelationshipsOk() (*MetricAssetResponseRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *MetricAssetResponseData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given MetricAssetResponseRelationships and assigns it to the Relationships field.
func (o *MetricAssetResponseData) SetRelationships(v MetricAssetResponseRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *MetricAssetResponseData) GetType() MetricType {
	if o == nil {
		var ret MetricType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetricAssetResponseData) GetTypeOk() (*MetricType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *MetricAssetResponseData) SetType(v MetricType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricAssetResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricAssetResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id            *string                           `json:"id"`
		Relationships *MetricAssetResponseRelationships `json:"relationships,omitempty"`
		Type          *MetricType                       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
