// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentHost A host that is part of a deployment with its current status.
type FleetDeploymentHost struct {
	// Error message if the deployment failed on this host.
	Error *string `json:"error,omitempty"`
	// The hostname of the agent.
	Hostname *string `json:"hostname,omitempty"`
	// Current deployment status for this specific host.
	Status *string `json:"status,omitempty"`
	// List of packages and their versions currently installed on this host.
	Versions []FleetDeploymentHostPackage `json:"versions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentHost instantiates a new FleetDeploymentHost object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentHost() *FleetDeploymentHost {
	this := FleetDeploymentHost{}
	return &this
}

// NewFleetDeploymentHostWithDefaults instantiates a new FleetDeploymentHost object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentHostWithDefaults() *FleetDeploymentHost {
	this := FleetDeploymentHost{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *FleetDeploymentHost) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentHost) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *FleetDeploymentHost) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *FleetDeploymentHost) SetError(v string) {
	o.Error = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *FleetDeploymentHost) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentHost) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *FleetDeploymentHost) HasHostname() bool {
	return o != nil && o.Hostname != nil
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *FleetDeploymentHost) SetHostname(v string) {
	o.Hostname = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FleetDeploymentHost) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentHost) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FleetDeploymentHost) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FleetDeploymentHost) SetStatus(v string) {
	o.Status = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *FleetDeploymentHost) GetVersions() []FleetDeploymentHostPackage {
	if o == nil || o.Versions == nil {
		var ret []FleetDeploymentHostPackage
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentHost) GetVersionsOk() (*[]FleetDeploymentHostPackage, bool) {
	if o == nil || o.Versions == nil {
		return nil, false
	}
	return &o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *FleetDeploymentHost) HasVersions() bool {
	return o != nil && o.Versions != nil
}

// SetVersions gets a reference to the given []FleetDeploymentHostPackage and assigns it to the Versions field.
func (o *FleetDeploymentHost) SetVersions(v []FleetDeploymentHostPackage) {
	o.Versions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentHost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Hostname != nil {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Versions != nil {
		toSerialize["versions"] = o.Versions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentHost) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Error    *string                      `json:"error,omitempty"`
		Hostname *string                      `json:"hostname,omitempty"`
		Status   *string                      `json:"status,omitempty"`
		Versions []FleetDeploymentHostPackage `json:"versions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error", "hostname", "status", "versions"})
	} else {
		return err
	}
	o.Error = all.Error
	o.Hostname = all.Hostname
	o.Status = all.Status
	o.Versions = all.Versions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
