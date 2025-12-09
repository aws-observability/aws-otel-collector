// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesDependenciesItems
type ScaRequestDataAttributesDependenciesItems struct {
	//
	Exclusions []string `json:"exclusions,omitempty"`
	//
	Group *string `json:"group,omitempty"`
	//
	IsDev *bool `json:"is_dev,omitempty"`
	//
	IsDirect *bool `json:"is_direct,omitempty"`
	//
	Language *string `json:"language,omitempty"`
	//
	Locations []ScaRequestDataAttributesDependenciesItemsLocationsItems `json:"locations,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	PackageManager *string `json:"package_manager,omitempty"`
	//
	Purl *string `json:"purl,omitempty"`
	//
	ReachableSymbolProperties []ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems `json:"reachable_symbol_properties,omitempty"`
	//
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesDependenciesItems instantiates a new ScaRequestDataAttributesDependenciesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesDependenciesItems() *ScaRequestDataAttributesDependenciesItems {
	this := ScaRequestDataAttributesDependenciesItems{}
	return &this
}

// NewScaRequestDataAttributesDependenciesItemsWithDefaults instantiates a new ScaRequestDataAttributesDependenciesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesDependenciesItemsWithDefaults() *ScaRequestDataAttributesDependenciesItems {
	this := ScaRequestDataAttributesDependenciesItems{}
	return &this
}

// GetExclusions returns the Exclusions field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetExclusions() []string {
	if o == nil || o.Exclusions == nil {
		var ret []string
		return ret
	}
	return o.Exclusions
}

// GetExclusionsOk returns a tuple with the Exclusions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetExclusionsOk() (*[]string, bool) {
	if o == nil || o.Exclusions == nil {
		return nil, false
	}
	return &o.Exclusions, true
}

// HasExclusions returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasExclusions() bool {
	return o != nil && o.Exclusions != nil
}

// SetExclusions gets a reference to the given []string and assigns it to the Exclusions field.
func (o *ScaRequestDataAttributesDependenciesItems) SetExclusions(v []string) {
	o.Exclusions = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasGroup() bool {
	return o != nil && o.Group != nil
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *ScaRequestDataAttributesDependenciesItems) SetGroup(v string) {
	o.Group = &v
}

// GetIsDev returns the IsDev field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetIsDev() bool {
	if o == nil || o.IsDev == nil {
		var ret bool
		return ret
	}
	return *o.IsDev
}

// GetIsDevOk returns a tuple with the IsDev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetIsDevOk() (*bool, bool) {
	if o == nil || o.IsDev == nil {
		return nil, false
	}
	return o.IsDev, true
}

// HasIsDev returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasIsDev() bool {
	return o != nil && o.IsDev != nil
}

// SetIsDev gets a reference to the given bool and assigns it to the IsDev field.
func (o *ScaRequestDataAttributesDependenciesItems) SetIsDev(v bool) {
	o.IsDev = &v
}

// GetIsDirect returns the IsDirect field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetIsDirect() bool {
	if o == nil || o.IsDirect == nil {
		var ret bool
		return ret
	}
	return *o.IsDirect
}

// GetIsDirectOk returns a tuple with the IsDirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetIsDirectOk() (*bool, bool) {
	if o == nil || o.IsDirect == nil {
		return nil, false
	}
	return o.IsDirect, true
}

// HasIsDirect returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasIsDirect() bool {
	return o != nil && o.IsDirect != nil
}

// SetIsDirect gets a reference to the given bool and assigns it to the IsDirect field.
func (o *ScaRequestDataAttributesDependenciesItems) SetIsDirect(v bool) {
	o.IsDirect = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasLanguage() bool {
	return o != nil && o.Language != nil
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ScaRequestDataAttributesDependenciesItems) SetLanguage(v string) {
	o.Language = &v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetLocations() []ScaRequestDataAttributesDependenciesItemsLocationsItems {
	if o == nil || o.Locations == nil {
		var ret []ScaRequestDataAttributesDependenciesItemsLocationsItems
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetLocationsOk() (*[]ScaRequestDataAttributesDependenciesItemsLocationsItems, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return &o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasLocations() bool {
	return o != nil && o.Locations != nil
}

// SetLocations gets a reference to the given []ScaRequestDataAttributesDependenciesItemsLocationsItems and assigns it to the Locations field.
func (o *ScaRequestDataAttributesDependenciesItems) SetLocations(v []ScaRequestDataAttributesDependenciesItemsLocationsItems) {
	o.Locations = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ScaRequestDataAttributesDependenciesItems) SetName(v string) {
	o.Name = &v
}

// GetPackageManager returns the PackageManager field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetPackageManager() string {
	if o == nil || o.PackageManager == nil {
		var ret string
		return ret
	}
	return *o.PackageManager
}

// GetPackageManagerOk returns a tuple with the PackageManager field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetPackageManagerOk() (*string, bool) {
	if o == nil || o.PackageManager == nil {
		return nil, false
	}
	return o.PackageManager, true
}

// HasPackageManager returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasPackageManager() bool {
	return o != nil && o.PackageManager != nil
}

// SetPackageManager gets a reference to the given string and assigns it to the PackageManager field.
func (o *ScaRequestDataAttributesDependenciesItems) SetPackageManager(v string) {
	o.PackageManager = &v
}

// GetPurl returns the Purl field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetPurl() string {
	if o == nil || o.Purl == nil {
		var ret string
		return ret
	}
	return *o.Purl
}

// GetPurlOk returns a tuple with the Purl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetPurlOk() (*string, bool) {
	if o == nil || o.Purl == nil {
		return nil, false
	}
	return o.Purl, true
}

// HasPurl returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasPurl() bool {
	return o != nil && o.Purl != nil
}

// SetPurl gets a reference to the given string and assigns it to the Purl field.
func (o *ScaRequestDataAttributesDependenciesItems) SetPurl(v string) {
	o.Purl = &v
}

// GetReachableSymbolProperties returns the ReachableSymbolProperties field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetReachableSymbolProperties() []ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems {
	if o == nil || o.ReachableSymbolProperties == nil {
		var ret []ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems
		return ret
	}
	return o.ReachableSymbolProperties
}

// GetReachableSymbolPropertiesOk returns a tuple with the ReachableSymbolProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetReachableSymbolPropertiesOk() (*[]ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems, bool) {
	if o == nil || o.ReachableSymbolProperties == nil {
		return nil, false
	}
	return &o.ReachableSymbolProperties, true
}

// HasReachableSymbolProperties returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasReachableSymbolProperties() bool {
	return o != nil && o.ReachableSymbolProperties != nil
}

// SetReachableSymbolProperties gets a reference to the given []ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems and assigns it to the ReachableSymbolProperties field.
func (o *ScaRequestDataAttributesDependenciesItems) SetReachableSymbolProperties(v []ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems) {
	o.ReachableSymbolProperties = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesDependenciesItems) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesDependenciesItems) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesDependenciesItems) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ScaRequestDataAttributesDependenciesItems) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesDependenciesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Exclusions != nil {
		toSerialize["exclusions"] = o.Exclusions
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.IsDev != nil {
		toSerialize["is_dev"] = o.IsDev
	}
	if o.IsDirect != nil {
		toSerialize["is_direct"] = o.IsDirect
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PackageManager != nil {
		toSerialize["package_manager"] = o.PackageManager
	}
	if o.Purl != nil {
		toSerialize["purl"] = o.Purl
	}
	if o.ReachableSymbolProperties != nil {
		toSerialize["reachable_symbol_properties"] = o.ReachableSymbolProperties
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesDependenciesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Exclusions                []string                                                                  `json:"exclusions,omitempty"`
		Group                     *string                                                                   `json:"group,omitempty"`
		IsDev                     *bool                                                                     `json:"is_dev,omitempty"`
		IsDirect                  *bool                                                                     `json:"is_direct,omitempty"`
		Language                  *string                                                                   `json:"language,omitempty"`
		Locations                 []ScaRequestDataAttributesDependenciesItemsLocationsItems                 `json:"locations,omitempty"`
		Name                      *string                                                                   `json:"name,omitempty"`
		PackageManager            *string                                                                   `json:"package_manager,omitempty"`
		Purl                      *string                                                                   `json:"purl,omitempty"`
		ReachableSymbolProperties []ScaRequestDataAttributesDependenciesItemsReachableSymbolPropertiesItems `json:"reachable_symbol_properties,omitempty"`
		Version                   *string                                                                   `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exclusions", "group", "is_dev", "is_direct", "language", "locations", "name", "package_manager", "purl", "reachable_symbol_properties", "version"})
	} else {
		return err
	}
	o.Exclusions = all.Exclusions
	o.Group = all.Group
	o.IsDev = all.IsDev
	o.IsDirect = all.IsDirect
	o.Language = all.Language
	o.Locations = all.Locations
	o.Name = all.Name
	o.PackageManager = all.PackageManager
	o.Purl = all.Purl
	o.ReachableSymbolProperties = all.ReachableSymbolProperties
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
