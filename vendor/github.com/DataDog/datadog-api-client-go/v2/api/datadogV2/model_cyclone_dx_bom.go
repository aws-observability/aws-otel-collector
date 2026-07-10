// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CycloneDXBom A CycloneDX 1.5 Bill of Materials (BOM) document containing vulnerability data.
type CycloneDXBom struct {
	// The BOM format identifier. Must be `CycloneDX`.
	BomFormat string `json:"bomFormat"`
	// The list of scanned software components. Cannot be empty.
	Components []CycloneDXComponent `json:"components"`
	// Metadata about the BOM, including the scanned asset and the scanner tool.
	Metadata CycloneDXMetadata `json:"metadata"`
	// The CycloneDX specification version. Must be `1.5`.
	SpecVersion string `json:"specVersion"`
	// The version number of the BOM document.
	Version *int64 `json:"version,omitempty"`
	// The list of detected vulnerabilities. Cannot be empty.
	Vulnerabilities []CycloneDXVulnerability `json:"vulnerabilities"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCycloneDXBom instantiates a new CycloneDXBom object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCycloneDXBom(bomFormat string, components []CycloneDXComponent, metadata CycloneDXMetadata, specVersion string, vulnerabilities []CycloneDXVulnerability) *CycloneDXBom {
	this := CycloneDXBom{}
	this.BomFormat = bomFormat
	this.Components = components
	this.Metadata = metadata
	this.SpecVersion = specVersion
	this.Vulnerabilities = vulnerabilities
	return &this
}

// NewCycloneDXBomWithDefaults instantiates a new CycloneDXBom object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCycloneDXBomWithDefaults() *CycloneDXBom {
	this := CycloneDXBom{}
	return &this
}

// GetBomFormat returns the BomFormat field value.
func (o *CycloneDXBom) GetBomFormat() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BomFormat
}

// GetBomFormatOk returns a tuple with the BomFormat field value
// and a boolean to check if the value has been set.
func (o *CycloneDXBom) GetBomFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BomFormat, true
}

// SetBomFormat sets field value.
func (o *CycloneDXBom) SetBomFormat(v string) {
	o.BomFormat = v
}

// GetComponents returns the Components field value.
func (o *CycloneDXBom) GetComponents() []CycloneDXComponent {
	if o == nil {
		var ret []CycloneDXComponent
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
func (o *CycloneDXBom) GetComponentsOk() (*[]CycloneDXComponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Components, true
}

// SetComponents sets field value.
func (o *CycloneDXBom) SetComponents(v []CycloneDXComponent) {
	o.Components = v
}

// GetMetadata returns the Metadata field value.
func (o *CycloneDXBom) GetMetadata() CycloneDXMetadata {
	if o == nil {
		var ret CycloneDXMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *CycloneDXBom) GetMetadataOk() (*CycloneDXMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value.
func (o *CycloneDXBom) SetMetadata(v CycloneDXMetadata) {
	o.Metadata = v
}

// GetSpecVersion returns the SpecVersion field value.
func (o *CycloneDXBom) GetSpecVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SpecVersion
}

// GetSpecVersionOk returns a tuple with the SpecVersion field value
// and a boolean to check if the value has been set.
func (o *CycloneDXBom) GetSpecVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecVersion, true
}

// SetSpecVersion sets field value.
func (o *CycloneDXBom) SetSpecVersion(v string) {
	o.SpecVersion = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CycloneDXBom) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CycloneDXBom) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CycloneDXBom) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *CycloneDXBom) SetVersion(v int64) {
	o.Version = &v
}

// GetVulnerabilities returns the Vulnerabilities field value.
func (o *CycloneDXBom) GetVulnerabilities() []CycloneDXVulnerability {
	if o == nil {
		var ret []CycloneDXVulnerability
		return ret
	}
	return o.Vulnerabilities
}

// GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field value
// and a boolean to check if the value has been set.
func (o *CycloneDXBom) GetVulnerabilitiesOk() (*[]CycloneDXVulnerability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vulnerabilities, true
}

// SetVulnerabilities sets field value.
func (o *CycloneDXBom) SetVulnerabilities(v []CycloneDXVulnerability) {
	o.Vulnerabilities = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CycloneDXBom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["bomFormat"] = o.BomFormat
	toSerialize["components"] = o.Components
	toSerialize["metadata"] = o.Metadata
	toSerialize["specVersion"] = o.SpecVersion
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	toSerialize["vulnerabilities"] = o.Vulnerabilities

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CycloneDXBom) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BomFormat       *string                   `json:"bomFormat"`
		Components      *[]CycloneDXComponent     `json:"components"`
		Metadata        *CycloneDXMetadata        `json:"metadata"`
		SpecVersion     *string                   `json:"specVersion"`
		Version         *int64                    `json:"version,omitempty"`
		Vulnerabilities *[]CycloneDXVulnerability `json:"vulnerabilities"`
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
	if all.Metadata == nil {
		return fmt.Errorf("required field metadata missing")
	}
	if all.SpecVersion == nil {
		return fmt.Errorf("required field specVersion missing")
	}
	if all.Vulnerabilities == nil {
		return fmt.Errorf("required field vulnerabilities missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bomFormat", "components", "metadata", "specVersion", "version", "vulnerabilities"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BomFormat = *all.BomFormat
	o.Components = *all.Components
	if all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = *all.Metadata
	o.SpecVersion = *all.SpecVersion
	o.Version = all.Version
	o.Vulnerabilities = *all.Vulnerabilities

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
