// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// McpScanRequestDataAttributesLibrariesItems A library declaration to include in the dependency scan.
type McpScanRequestDataAttributesLibrariesItems struct {
	// The list of dependency PURLs to exclude when resolving transitive dependencies for this library.
	Exclusions []string `json:"exclusions,omitempty"`
	// Whether this library is a development-only dependency.
	IsDev bool `json:"is_dev"`
	// Whether this library is a direct (rather than transitive) dependency.
	IsDirect bool `json:"is_direct"`
	// The package manager that produced this library entry (for example, `npm`, `pip`, `nuget`).
	PackageManager string `json:"package_manager"`
	// The Package URL (PURL) uniquely identifying the library and its version.
	Purl string `json:"purl"`
	// The list of target framework identifiers associated with the library.
	TargetFrameworks []string `json:"target_frameworks,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMcpScanRequestDataAttributesLibrariesItems instantiates a new McpScanRequestDataAttributesLibrariesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMcpScanRequestDataAttributesLibrariesItems(isDev bool, isDirect bool, packageManager string, purl string) *McpScanRequestDataAttributesLibrariesItems {
	this := McpScanRequestDataAttributesLibrariesItems{}
	this.IsDev = isDev
	this.IsDirect = isDirect
	this.PackageManager = packageManager
	this.Purl = purl
	return &this
}

// NewMcpScanRequestDataAttributesLibrariesItemsWithDefaults instantiates a new McpScanRequestDataAttributesLibrariesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMcpScanRequestDataAttributesLibrariesItemsWithDefaults() *McpScanRequestDataAttributesLibrariesItems {
	this := McpScanRequestDataAttributesLibrariesItems{}
	return &this
}

// GetExclusions returns the Exclusions field value if set, zero value otherwise.
func (o *McpScanRequestDataAttributesLibrariesItems) GetExclusions() []string {
	if o == nil || o.Exclusions == nil {
		var ret []string
		return ret
	}
	return o.Exclusions
}

// GetExclusionsOk returns a tuple with the Exclusions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpScanRequestDataAttributesLibrariesItems) GetExclusionsOk() (*[]string, bool) {
	if o == nil || o.Exclusions == nil {
		return nil, false
	}
	return &o.Exclusions, true
}

// HasExclusions returns a boolean if a field has been set.
func (o *McpScanRequestDataAttributesLibrariesItems) HasExclusions() bool {
	return o != nil && o.Exclusions != nil
}

// SetExclusions gets a reference to the given []string and assigns it to the Exclusions field.
func (o *McpScanRequestDataAttributesLibrariesItems) SetExclusions(v []string) {
	o.Exclusions = v
}

// GetIsDev returns the IsDev field value.
func (o *McpScanRequestDataAttributesLibrariesItems) GetIsDev() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDev
}

// GetIsDevOk returns a tuple with the IsDev field value
// and a boolean to check if the value has been set.
func (o *McpScanRequestDataAttributesLibrariesItems) GetIsDevOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDev, true
}

// SetIsDev sets field value.
func (o *McpScanRequestDataAttributesLibrariesItems) SetIsDev(v bool) {
	o.IsDev = v
}

// GetIsDirect returns the IsDirect field value.
func (o *McpScanRequestDataAttributesLibrariesItems) GetIsDirect() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsDirect
}

// GetIsDirectOk returns a tuple with the IsDirect field value
// and a boolean to check if the value has been set.
func (o *McpScanRequestDataAttributesLibrariesItems) GetIsDirectOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDirect, true
}

// SetIsDirect sets field value.
func (o *McpScanRequestDataAttributesLibrariesItems) SetIsDirect(v bool) {
	o.IsDirect = v
}

// GetPackageManager returns the PackageManager field value.
func (o *McpScanRequestDataAttributesLibrariesItems) GetPackageManager() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PackageManager
}

// GetPackageManagerOk returns a tuple with the PackageManager field value
// and a boolean to check if the value has been set.
func (o *McpScanRequestDataAttributesLibrariesItems) GetPackageManagerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PackageManager, true
}

// SetPackageManager sets field value.
func (o *McpScanRequestDataAttributesLibrariesItems) SetPackageManager(v string) {
	o.PackageManager = v
}

// GetPurl returns the Purl field value.
func (o *McpScanRequestDataAttributesLibrariesItems) GetPurl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Purl
}

// GetPurlOk returns a tuple with the Purl field value
// and a boolean to check if the value has been set.
func (o *McpScanRequestDataAttributesLibrariesItems) GetPurlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Purl, true
}

// SetPurl sets field value.
func (o *McpScanRequestDataAttributesLibrariesItems) SetPurl(v string) {
	o.Purl = v
}

// GetTargetFrameworks returns the TargetFrameworks field value if set, zero value otherwise.
func (o *McpScanRequestDataAttributesLibrariesItems) GetTargetFrameworks() []string {
	if o == nil || o.TargetFrameworks == nil {
		var ret []string
		return ret
	}
	return o.TargetFrameworks
}

// GetTargetFrameworksOk returns a tuple with the TargetFrameworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *McpScanRequestDataAttributesLibrariesItems) GetTargetFrameworksOk() (*[]string, bool) {
	if o == nil || o.TargetFrameworks == nil {
		return nil, false
	}
	return &o.TargetFrameworks, true
}

// HasTargetFrameworks returns a boolean if a field has been set.
func (o *McpScanRequestDataAttributesLibrariesItems) HasTargetFrameworks() bool {
	return o != nil && o.TargetFrameworks != nil
}

// SetTargetFrameworks gets a reference to the given []string and assigns it to the TargetFrameworks field.
func (o *McpScanRequestDataAttributesLibrariesItems) SetTargetFrameworks(v []string) {
	o.TargetFrameworks = v
}

// MarshalJSON serializes the struct using spec logic.
func (o McpScanRequestDataAttributesLibrariesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Exclusions != nil {
		toSerialize["exclusions"] = o.Exclusions
	}
	toSerialize["is_dev"] = o.IsDev
	toSerialize["is_direct"] = o.IsDirect
	toSerialize["package_manager"] = o.PackageManager
	toSerialize["purl"] = o.Purl
	if o.TargetFrameworks != nil {
		toSerialize["target_frameworks"] = o.TargetFrameworks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *McpScanRequestDataAttributesLibrariesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Exclusions       []string `json:"exclusions,omitempty"`
		IsDev            *bool    `json:"is_dev"`
		IsDirect         *bool    `json:"is_direct"`
		PackageManager   *string  `json:"package_manager"`
		Purl             *string  `json:"purl"`
		TargetFrameworks []string `json:"target_frameworks,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.IsDev == nil {
		return fmt.Errorf("required field is_dev missing")
	}
	if all.IsDirect == nil {
		return fmt.Errorf("required field is_direct missing")
	}
	if all.PackageManager == nil {
		return fmt.Errorf("required field package_manager missing")
	}
	if all.Purl == nil {
		return fmt.Errorf("required field purl missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"exclusions", "is_dev", "is_direct", "package_manager", "purl", "target_frameworks"})
	} else {
		return err
	}
	o.Exclusions = all.Exclusions
	o.IsDev = *all.IsDev
	o.IsDirect = *all.IsDirect
	o.PackageManager = *all.PackageManager
	o.Purl = *all.Purl
	o.TargetFrameworks = all.TargetFrameworks

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
