// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentResponseMeta Metadata for a single deployment response, including pagination information for hosts.
type FleetDeploymentResponseMeta struct {
	// Pagination details for the list of hosts in a deployment.
	Hosts *FleetDeploymentHostsPage `json:"hosts,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentResponseMeta instantiates a new FleetDeploymentResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentResponseMeta() *FleetDeploymentResponseMeta {
	this := FleetDeploymentResponseMeta{}
	return &this
}

// NewFleetDeploymentResponseMetaWithDefaults instantiates a new FleetDeploymentResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentResponseMetaWithDefaults() *FleetDeploymentResponseMeta {
	this := FleetDeploymentResponseMeta{}
	return &this
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *FleetDeploymentResponseMeta) GetHosts() FleetDeploymentHostsPage {
	if o == nil || o.Hosts == nil {
		var ret FleetDeploymentHostsPage
		return ret
	}
	return *o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentResponseMeta) GetHostsOk() (*FleetDeploymentHostsPage, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *FleetDeploymentResponseMeta) HasHosts() bool {
	return o != nil && o.Hosts != nil
}

// SetHosts gets a reference to the given FleetDeploymentHostsPage and assigns it to the Hosts field.
func (o *FleetDeploymentResponseMeta) SetHosts(v FleetDeploymentHostsPage) {
	o.Hosts = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Hosts *FleetDeploymentHostsPage `json:"hosts,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"hosts"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Hosts != nil && all.Hosts.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Hosts = all.Hosts

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
