// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SBOMComponent Software or hardware component.
type SBOMComponent struct {
	// An optional identifier that can be used to reference the component elsewhere in the BOM.
	BomRef *string `json:"bom-ref,omitempty"`
	// The software licenses of the SBOM component.
	Licenses []SBOMComponentLicense `json:"licenses,omitempty"`
	// The name of the component. This will often be a shortened, single name of the component.
	Name string `json:"name"`
	// The custom properties of the component of the SBOM.
	Properties []SBOMComponentProperty `json:"properties,omitempty"`
	// Specifies the package-url (purl). The purl, if specified, MUST be valid and conform to the [specification](https://github.com/package-url/purl-spec).
	Purl *string `json:"purl,omitempty"`
	// The supplier of the component.
	Supplier SBOMComponentSupplier `json:"supplier"`
	// The SBOM component type
	Type SBOMComponentType `json:"type"`
	// The component version.
	Version string `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSBOMComponent instantiates a new SBOMComponent object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSBOMComponent(name string, supplier SBOMComponentSupplier, typeVar SBOMComponentType, version string) *SBOMComponent {
	this := SBOMComponent{}
	this.Name = name
	this.Supplier = supplier
	this.Type = typeVar
	this.Version = version
	return &this
}

// NewSBOMComponentWithDefaults instantiates a new SBOMComponent object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSBOMComponentWithDefaults() *SBOMComponent {
	this := SBOMComponent{}
	return &this
}

// GetBomRef returns the BomRef field value if set, zero value otherwise.
func (o *SBOMComponent) GetBomRef() string {
	if o == nil || o.BomRef == nil {
		var ret string
		return ret
	}
	return *o.BomRef
}

// GetBomRefOk returns a tuple with the BomRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBOMComponent) GetBomRefOk() (*string, bool) {
	if o == nil || o.BomRef == nil {
		return nil, false
	}
	return o.BomRef, true
}

// HasBomRef returns a boolean if a field has been set.
func (o *SBOMComponent) HasBomRef() bool {
	return o != nil && o.BomRef != nil
}

// SetBomRef gets a reference to the given string and assigns it to the BomRef field.
func (o *SBOMComponent) SetBomRef(v string) {
	o.BomRef = &v
}

// GetLicenses returns the Licenses field value if set, zero value otherwise.
func (o *SBOMComponent) GetLicenses() []SBOMComponentLicense {
	if o == nil || o.Licenses == nil {
		var ret []SBOMComponentLicense
		return ret
	}
	return o.Licenses
}

// GetLicensesOk returns a tuple with the Licenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBOMComponent) GetLicensesOk() (*[]SBOMComponentLicense, bool) {
	if o == nil || o.Licenses == nil {
		return nil, false
	}
	return &o.Licenses, true
}

// HasLicenses returns a boolean if a field has been set.
func (o *SBOMComponent) HasLicenses() bool {
	return o != nil && o.Licenses != nil
}

// SetLicenses gets a reference to the given []SBOMComponentLicense and assigns it to the Licenses field.
func (o *SBOMComponent) SetLicenses(v []SBOMComponentLicense) {
	o.Licenses = v
}

// GetName returns the Name field value.
func (o *SBOMComponent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SBOMComponent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *SBOMComponent) SetName(v string) {
	o.Name = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *SBOMComponent) GetProperties() []SBOMComponentProperty {
	if o == nil || o.Properties == nil {
		var ret []SBOMComponentProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBOMComponent) GetPropertiesOk() (*[]SBOMComponentProperty, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *SBOMComponent) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given []SBOMComponentProperty and assigns it to the Properties field.
func (o *SBOMComponent) SetProperties(v []SBOMComponentProperty) {
	o.Properties = v
}

// GetPurl returns the Purl field value if set, zero value otherwise.
func (o *SBOMComponent) GetPurl() string {
	if o == nil || o.Purl == nil {
		var ret string
		return ret
	}
	return *o.Purl
}

// GetPurlOk returns a tuple with the Purl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBOMComponent) GetPurlOk() (*string, bool) {
	if o == nil || o.Purl == nil {
		return nil, false
	}
	return o.Purl, true
}

// HasPurl returns a boolean if a field has been set.
func (o *SBOMComponent) HasPurl() bool {
	return o != nil && o.Purl != nil
}

// SetPurl gets a reference to the given string and assigns it to the Purl field.
func (o *SBOMComponent) SetPurl(v string) {
	o.Purl = &v
}

// GetSupplier returns the Supplier field value.
func (o *SBOMComponent) GetSupplier() SBOMComponentSupplier {
	if o == nil {
		var ret SBOMComponentSupplier
		return ret
	}
	return o.Supplier
}

// GetSupplierOk returns a tuple with the Supplier field value
// and a boolean to check if the value has been set.
func (o *SBOMComponent) GetSupplierOk() (*SBOMComponentSupplier, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supplier, true
}

// SetSupplier sets field value.
func (o *SBOMComponent) SetSupplier(v SBOMComponentSupplier) {
	o.Supplier = v
}

// GetType returns the Type field value.
func (o *SBOMComponent) GetType() SBOMComponentType {
	if o == nil {
		var ret SBOMComponentType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SBOMComponent) GetTypeOk() (*SBOMComponentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *SBOMComponent) SetType(v SBOMComponentType) {
	o.Type = v
}

// GetVersion returns the Version field value.
func (o *SBOMComponent) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SBOMComponent) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *SBOMComponent) SetVersion(v string) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SBOMComponent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BomRef != nil {
		toSerialize["bom-ref"] = o.BomRef
	}
	if o.Licenses != nil {
		toSerialize["licenses"] = o.Licenses
	}
	toSerialize["name"] = o.Name
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Purl != nil {
		toSerialize["purl"] = o.Purl
	}
	toSerialize["supplier"] = o.Supplier
	toSerialize["type"] = o.Type
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SBOMComponent) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BomRef     *string                 `json:"bom-ref,omitempty"`
		Licenses   []SBOMComponentLicense  `json:"licenses,omitempty"`
		Name       *string                 `json:"name"`
		Properties []SBOMComponentProperty `json:"properties,omitempty"`
		Purl       *string                 `json:"purl,omitempty"`
		Supplier   *SBOMComponentSupplier  `json:"supplier"`
		Type       *SBOMComponentType      `json:"type"`
		Version    *string                 `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Supplier == nil {
		return fmt.Errorf("required field supplier missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bom-ref", "licenses", "name", "properties", "purl", "supplier", "type", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BomRef = all.BomRef
	o.Licenses = all.Licenses
	o.Name = *all.Name
	o.Properties = all.Properties
	o.Purl = all.Purl
	if all.Supplier.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Supplier = *all.Supplier
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
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
