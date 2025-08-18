// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AssetAttributes The JSON:API attributes of the asset.
type AssetAttributes struct {
	// Asset architecture.
	Arch *string `json:"arch,omitempty"`
	// List of environments where the asset is deployed.
	Environments []string `json:"environments"`
	// Asset name.
	Name string `json:"name"`
	// Asset operating system.
	OperatingSystem *AssetOperatingSystem `json:"operating_system,omitempty"`
	// Asset risks.
	Risks AssetRisks `json:"risks"`
	// List of teams that own the asset.
	Teams []string `json:"teams,omitempty"`
	// The asset type
	Type AssetType `json:"type"`
	// Asset version.
	Version *AssetVersion `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAssetAttributes instantiates a new AssetAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAssetAttributes(environments []string, name string, risks AssetRisks, typeVar AssetType) *AssetAttributes {
	this := AssetAttributes{}
	this.Environments = environments
	this.Name = name
	this.Risks = risks
	this.Type = typeVar
	return &this
}

// NewAssetAttributesWithDefaults instantiates a new AssetAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAssetAttributesWithDefaults() *AssetAttributes {
	this := AssetAttributes{}
	return &this
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *AssetAttributes) GetArch() string {
	if o == nil || o.Arch == nil {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAttributes) GetArchOk() (*string, bool) {
	if o == nil || o.Arch == nil {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *AssetAttributes) HasArch() bool {
	return o != nil && o.Arch != nil
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *AssetAttributes) SetArch(v string) {
	o.Arch = &v
}

// GetEnvironments returns the Environments field value.
func (o *AssetAttributes) GetEnvironments() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value
// and a boolean to check if the value has been set.
func (o *AssetAttributes) GetEnvironmentsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environments, true
}

// SetEnvironments sets field value.
func (o *AssetAttributes) SetEnvironments(v []string) {
	o.Environments = v
}

// GetName returns the Name field value.
func (o *AssetAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AssetAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AssetAttributes) SetName(v string) {
	o.Name = v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *AssetAttributes) GetOperatingSystem() AssetOperatingSystem {
	if o == nil || o.OperatingSystem == nil {
		var ret AssetOperatingSystem
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAttributes) GetOperatingSystemOk() (*AssetOperatingSystem, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *AssetAttributes) HasOperatingSystem() bool {
	return o != nil && o.OperatingSystem != nil
}

// SetOperatingSystem gets a reference to the given AssetOperatingSystem and assigns it to the OperatingSystem field.
func (o *AssetAttributes) SetOperatingSystem(v AssetOperatingSystem) {
	o.OperatingSystem = &v
}

// GetRisks returns the Risks field value.
func (o *AssetAttributes) GetRisks() AssetRisks {
	if o == nil {
		var ret AssetRisks
		return ret
	}
	return o.Risks
}

// GetRisksOk returns a tuple with the Risks field value
// and a boolean to check if the value has been set.
func (o *AssetAttributes) GetRisksOk() (*AssetRisks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Risks, true
}

// SetRisks sets field value.
func (o *AssetAttributes) SetRisks(v AssetRisks) {
	o.Risks = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *AssetAttributes) GetTeams() []string {
	if o == nil || o.Teams == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAttributes) GetTeamsOk() (*[]string, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return &o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *AssetAttributes) HasTeams() bool {
	return o != nil && o.Teams != nil
}

// SetTeams gets a reference to the given []string and assigns it to the Teams field.
func (o *AssetAttributes) SetTeams(v []string) {
	o.Teams = v
}

// GetType returns the Type field value.
func (o *AssetAttributes) GetType() AssetType {
	if o == nil {
		var ret AssetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AssetAttributes) GetTypeOk() (*AssetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AssetAttributes) SetType(v AssetType) {
	o.Type = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AssetAttributes) GetVersion() AssetVersion {
	if o == nil || o.Version == nil {
		var ret AssetVersion
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAttributes) GetVersionOk() (*AssetVersion, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AssetAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given AssetVersion and assigns it to the Version field.
func (o *AssetAttributes) SetVersion(v AssetVersion) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AssetAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Arch != nil {
		toSerialize["arch"] = o.Arch
	}
	toSerialize["environments"] = o.Environments
	toSerialize["name"] = o.Name
	if o.OperatingSystem != nil {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	toSerialize["risks"] = o.Risks
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}
	toSerialize["type"] = o.Type
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AssetAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arch            *string               `json:"arch,omitempty"`
		Environments    *[]string             `json:"environments"`
		Name            *string               `json:"name"`
		OperatingSystem *AssetOperatingSystem `json:"operating_system,omitempty"`
		Risks           *AssetRisks           `json:"risks"`
		Teams           []string              `json:"teams,omitempty"`
		Type            *AssetType            `json:"type"`
		Version         *AssetVersion         `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Environments == nil {
		return fmt.Errorf("required field environments missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Risks == nil {
		return fmt.Errorf("required field risks missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arch", "environments", "name", "operating_system", "risks", "teams", "type", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Arch = all.Arch
	o.Environments = *all.Environments
	o.Name = *all.Name
	if all.OperatingSystem != nil && all.OperatingSystem.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OperatingSystem = all.OperatingSystem
	if all.Risks.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Risks = *all.Risks
	o.Teams = all.Teams
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if all.Version != nil && all.Version.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
