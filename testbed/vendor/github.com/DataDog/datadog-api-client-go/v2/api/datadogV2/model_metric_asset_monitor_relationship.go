// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricAssetMonitorRelationship An object of type `monitor` that can be referenced in the `included` data.
type MetricAssetMonitorRelationship struct {
	// The related monitor's ID.
	Id *string `json:"id,omitempty"`
	// Monitor resource type.
	Type *MetricMonitorType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMetricAssetMonitorRelationship instantiates a new MetricAssetMonitorRelationship object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricAssetMonitorRelationship() *MetricAssetMonitorRelationship {
	this := MetricAssetMonitorRelationship{}
	return &this
}

// NewMetricAssetMonitorRelationshipWithDefaults instantiates a new MetricAssetMonitorRelationship object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricAssetMonitorRelationshipWithDefaults() *MetricAssetMonitorRelationship {
	this := MetricAssetMonitorRelationship{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricAssetMonitorRelationship) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAssetMonitorRelationship) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricAssetMonitorRelationship) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetricAssetMonitorRelationship) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetricAssetMonitorRelationship) GetType() MetricMonitorType {
	if o == nil || o.Type == nil {
		var ret MetricMonitorType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricAssetMonitorRelationship) GetTypeOk() (*MetricMonitorType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetricAssetMonitorRelationship) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given MetricMonitorType and assigns it to the Type field.
func (o *MetricAssetMonitorRelationship) SetType(v MetricMonitorType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricAssetMonitorRelationship) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
func (o *MetricAssetMonitorRelationship) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string            `json:"id,omitempty"`
		Type *MetricMonitorType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
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
