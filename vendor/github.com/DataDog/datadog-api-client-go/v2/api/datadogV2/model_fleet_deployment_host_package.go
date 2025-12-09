// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentHostPackage Package version information for a host, showing the initial version before deployment,
// the target version to deploy, and the current version on the host.
type FleetDeploymentHostPackage struct {
	// The current version of the package on the host.
	CurrentVersion *string `json:"current_version,omitempty"`
	// The initial version of the package on the host before the deployment started.
	InitialVersion *string `json:"initial_version,omitempty"`
	// The name of the package.
	PackageName *string `json:"package_name,omitempty"`
	// The target version that the deployment is attempting to install.
	TargetVersion *string `json:"target_version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentHostPackage instantiates a new FleetDeploymentHostPackage object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentHostPackage() *FleetDeploymentHostPackage {
	this := FleetDeploymentHostPackage{}
	return &this
}

// NewFleetDeploymentHostPackageWithDefaults instantiates a new FleetDeploymentHostPackage object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentHostPackageWithDefaults() *FleetDeploymentHostPackage {
	this := FleetDeploymentHostPackage{}
	return &this
}

// GetCurrentVersion returns the CurrentVersion field value if set, zero value otherwise.
func (o *FleetDeploymentHostPackage) GetCurrentVersion() string {
	if o == nil || o.CurrentVersion == nil {
		var ret string
		return ret
	}
	return *o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentHostPackage) GetCurrentVersionOk() (*string, bool) {
	if o == nil || o.CurrentVersion == nil {
		return nil, false
	}
	return o.CurrentVersion, true
}

// HasCurrentVersion returns a boolean if a field has been set.
func (o *FleetDeploymentHostPackage) HasCurrentVersion() bool {
	return o != nil && o.CurrentVersion != nil
}

// SetCurrentVersion gets a reference to the given string and assigns it to the CurrentVersion field.
func (o *FleetDeploymentHostPackage) SetCurrentVersion(v string) {
	o.CurrentVersion = &v
}

// GetInitialVersion returns the InitialVersion field value if set, zero value otherwise.
func (o *FleetDeploymentHostPackage) GetInitialVersion() string {
	if o == nil || o.InitialVersion == nil {
		var ret string
		return ret
	}
	return *o.InitialVersion
}

// GetInitialVersionOk returns a tuple with the InitialVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentHostPackage) GetInitialVersionOk() (*string, bool) {
	if o == nil || o.InitialVersion == nil {
		return nil, false
	}
	return o.InitialVersion, true
}

// HasInitialVersion returns a boolean if a field has been set.
func (o *FleetDeploymentHostPackage) HasInitialVersion() bool {
	return o != nil && o.InitialVersion != nil
}

// SetInitialVersion gets a reference to the given string and assigns it to the InitialVersion field.
func (o *FleetDeploymentHostPackage) SetInitialVersion(v string) {
	o.InitialVersion = &v
}

// GetPackageName returns the PackageName field value if set, zero value otherwise.
func (o *FleetDeploymentHostPackage) GetPackageName() string {
	if o == nil || o.PackageName == nil {
		var ret string
		return ret
	}
	return *o.PackageName
}

// GetPackageNameOk returns a tuple with the PackageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentHostPackage) GetPackageNameOk() (*string, bool) {
	if o == nil || o.PackageName == nil {
		return nil, false
	}
	return o.PackageName, true
}

// HasPackageName returns a boolean if a field has been set.
func (o *FleetDeploymentHostPackage) HasPackageName() bool {
	return o != nil && o.PackageName != nil
}

// SetPackageName gets a reference to the given string and assigns it to the PackageName field.
func (o *FleetDeploymentHostPackage) SetPackageName(v string) {
	o.PackageName = &v
}

// GetTargetVersion returns the TargetVersion field value if set, zero value otherwise.
func (o *FleetDeploymentHostPackage) GetTargetVersion() string {
	if o == nil || o.TargetVersion == nil {
		var ret string
		return ret
	}
	return *o.TargetVersion
}

// GetTargetVersionOk returns a tuple with the TargetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentHostPackage) GetTargetVersionOk() (*string, bool) {
	if o == nil || o.TargetVersion == nil {
		return nil, false
	}
	return o.TargetVersion, true
}

// HasTargetVersion returns a boolean if a field has been set.
func (o *FleetDeploymentHostPackage) HasTargetVersion() bool {
	return o != nil && o.TargetVersion != nil
}

// SetTargetVersion gets a reference to the given string and assigns it to the TargetVersion field.
func (o *FleetDeploymentHostPackage) SetTargetVersion(v string) {
	o.TargetVersion = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentHostPackage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CurrentVersion != nil {
		toSerialize["current_version"] = o.CurrentVersion
	}
	if o.InitialVersion != nil {
		toSerialize["initial_version"] = o.InitialVersion
	}
	if o.PackageName != nil {
		toSerialize["package_name"] = o.PackageName
	}
	if o.TargetVersion != nil {
		toSerialize["target_version"] = o.TargetVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentHostPackage) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CurrentVersion *string `json:"current_version,omitempty"`
		InitialVersion *string `json:"initial_version,omitempty"`
		PackageName    *string `json:"package_name,omitempty"`
		TargetVersion  *string `json:"target_version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"current_version", "initial_version", "package_name", "target_version"})
	} else {
		return err
	}
	o.CurrentVersion = all.CurrentVersion
	o.InitialVersion = all.InitialVersion
	o.PackageName = all.PackageName
	o.TargetVersion = all.TargetVersion

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
