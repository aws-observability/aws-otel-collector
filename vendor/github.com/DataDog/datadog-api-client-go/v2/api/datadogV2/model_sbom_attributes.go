// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SBOMAttributes The JSON:API attributes of the SBOM.
type SBOMAttributes struct {
	// Specifies the format of the BOM. This helps to identify the file as CycloneDX since BOM do not have a filename convention nor does JSON schema support namespaces. This value MUST be `CycloneDX`.
	BomFormat string `json:"bomFormat"`
	// A list of software and hardware components.
	Components []SBOMComponent `json:"components"`
	// List of dependencies between components of the SBOM.
	Dependencies []SBOMComponentDependency `json:"dependencies"`
	// Provides additional information about a BOM.
	Metadata SBOMMetadata `json:"metadata"`
	// Every BOM generated has a unique serial number, even if the contents of the BOM have not changed overt time. The serial number follows [RFC-4122](https://datatracker.ietf.org/doc/html/rfc4122)
	SerialNumber string `json:"serialNumber"`
	// The version of the CycloneDX specification a BOM conforms to.
	SpecVersion SpecVersion `json:"specVersion"`
	// It increments when a BOM is modified. The default value is 1.
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSBOMAttributes instantiates a new SBOMAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSBOMAttributes(bomFormat string, components []SBOMComponent, dependencies []SBOMComponentDependency, metadata SBOMMetadata, serialNumber string, specVersion SpecVersion, version int64) *SBOMAttributes {
	this := SBOMAttributes{}
	this.BomFormat = bomFormat
	this.Components = components
	this.Dependencies = dependencies
	this.Metadata = metadata
	this.SerialNumber = serialNumber
	this.SpecVersion = specVersion
	this.Version = version
	return &this
}

// NewSBOMAttributesWithDefaults instantiates a new SBOMAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSBOMAttributesWithDefaults() *SBOMAttributes {
	this := SBOMAttributes{}
	return &this
}

// GetBomFormat returns the BomFormat field value.
func (o *SBOMAttributes) GetBomFormat() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BomFormat
}

// GetBomFormatOk returns a tuple with the BomFormat field value
// and a boolean to check if the value has been set.
func (o *SBOMAttributes) GetBomFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BomFormat, true
}

// SetBomFormat sets field value.
func (o *SBOMAttributes) SetBomFormat(v string) {
	o.BomFormat = v
}

// GetComponents returns the Components field value.
func (o *SBOMAttributes) GetComponents() []SBOMComponent {
	if o == nil {
		var ret []SBOMComponent
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *SBOMAttributes) GetComponentsOk() (*[]SBOMComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value.
func (o *SBOMAttributes) SetComponents(v []SBOMComponent) {
	o.Components = v
}

// GetDependencies returns the Dependencies field value.
func (o *SBOMAttributes) GetDependencies() []SBOMComponentDependency {
	if o == nil {
		var ret []SBOMComponentDependency
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value
// and a boolean to check if the value has been set.
func (o *SBOMAttributes) GetDependenciesOk() (*[]SBOMComponentDependency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dependencies, true
}

// SetDependencies sets field value.
func (o *SBOMAttributes) SetDependencies(v []SBOMComponentDependency) {
	o.Dependencies = v
}

// GetMetadata returns the Metadata field value.
func (o *SBOMAttributes) GetMetadata() SBOMMetadata {
	if o == nil {
		var ret SBOMMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SBOMAttributes) GetMetadataOk() (*SBOMMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value.
func (o *SBOMAttributes) SetMetadata(v SBOMMetadata) {
	o.Metadata = v
}

// GetSerialNumber returns the SerialNumber field value.
func (o *SBOMAttributes) GetSerialNumber() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value
// and a boolean to check if the value has been set.
func (o *SBOMAttributes) GetSerialNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SerialNumber, true
}

// SetSerialNumber sets field value.
func (o *SBOMAttributes) SetSerialNumber(v string) {
	o.SerialNumber = v
}

// GetSpecVersion returns the SpecVersion field value.
func (o *SBOMAttributes) GetSpecVersion() SpecVersion {
	if o == nil {
		var ret SpecVersion
		return ret
	}
	return o.SpecVersion
}

// GetSpecVersionOk returns a tuple with the SpecVersion field value
// and a boolean to check if the value has been set.
func (o *SBOMAttributes) GetSpecVersionOk() (*SpecVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecVersion, true
}

// SetSpecVersion sets field value.
func (o *SBOMAttributes) SetSpecVersion(v SpecVersion) {
	o.SpecVersion = v
}

// GetVersion returns the Version field value.
func (o *SBOMAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SBOMAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *SBOMAttributes) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SBOMAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["bomFormat"] = o.BomFormat
	toSerialize["components"] = o.Components
	toSerialize["dependencies"] = o.Dependencies
	toSerialize["metadata"] = o.Metadata
	toSerialize["serialNumber"] = o.SerialNumber
	toSerialize["specVersion"] = o.SpecVersion
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SBOMAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BomFormat    *string                    `json:"bomFormat"`
		Components   *[]SBOMComponent           `json:"components"`
		Dependencies *[]SBOMComponentDependency `json:"dependencies"`
		Metadata     *SBOMMetadata              `json:"metadata"`
		SerialNumber *string                    `json:"serialNumber"`
		SpecVersion  *SpecVersion               `json:"specVersion"`
		Version      *int64                     `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BomFormat == nil {
		return fmt.Errorf("required field bomFormat missing")
	}
	if all.Components == nil {
		return fmt.Errorf("required field components missing")
	}
	if all.Dependencies == nil {
		return fmt.Errorf("required field dependencies missing")
	}
	if all.Metadata == nil {
		return fmt.Errorf("required field metadata missing")
	}
	if all.SerialNumber == nil {
		return fmt.Errorf("required field serialNumber missing")
	}
	if all.SpecVersion == nil {
		return fmt.Errorf("required field specVersion missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bomFormat", "components", "dependencies", "metadata", "serialNumber", "specVersion", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BomFormat = *all.BomFormat
	o.Components = *all.Components
	o.Dependencies = *all.Dependencies
	if all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = *all.Metadata
	o.SerialNumber = *all.SerialNumber
	if !all.SpecVersion.IsValid() {
		hasInvalidField = true
	} else {
		o.SpecVersion = *all.SpecVersion
	}
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
