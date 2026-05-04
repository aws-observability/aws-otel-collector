// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetClustersResponseDataAttributes Attributes of the fleet clusters response containing the list of clusters.
type FleetClustersResponseDataAttributes struct {
	// Array of clusters matching the query criteria.
	Clusters []FleetClusterAttributes `json:"clusters,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetClustersResponseDataAttributes instantiates a new FleetClustersResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetClustersResponseDataAttributes() *FleetClustersResponseDataAttributes {
	this := FleetClustersResponseDataAttributes{}
	return &this
}

// NewFleetClustersResponseDataAttributesWithDefaults instantiates a new FleetClustersResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetClustersResponseDataAttributesWithDefaults() *FleetClustersResponseDataAttributes {
	this := FleetClustersResponseDataAttributes{}
	return &this
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *FleetClustersResponseDataAttributes) GetClusters() []FleetClusterAttributes {
	if o == nil || o.Clusters == nil {
		var ret []FleetClusterAttributes
		return ret
	}
	return o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetClustersResponseDataAttributes) GetClustersOk() (*[]FleetClusterAttributes, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return &o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *FleetClustersResponseDataAttributes) HasClusters() bool {
	return o != nil && o.Clusters != nil
}

// SetClusters gets a reference to the given []FleetClusterAttributes and assigns it to the Clusters field.
func (o *FleetClustersResponseDataAttributes) SetClusters(v []FleetClusterAttributes) {
	o.Clusters = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetClustersResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetClustersResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Clusters []FleetClusterAttributes `json:"clusters,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"clusters"})
	} else {
		return err
	}
	o.Clusters = all.Clusters

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
