// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ELFSourcemapAttributes Attributes of an ELF symbol file.
type ELFSourcemapAttributes struct {
	// The target CPU architecture.
	Arch *string `json:"arch,omitempty"`
	// The timestamp when the symbol file was created.
	CreatedAt time.Time `json:"created_at"`
	// The SHA256 hash of the ELF file.
	FileHash *string `json:"file_hash,omitempty"`
	// The ELF file name.
	FileName *string `json:"file_name,omitempty"`
	// The GNU build ID (UUID format).
	GnuBuildId *string `json:"gnu_build_id,omitempty"`
	// The Go build ID (UUID format).
	GoBuildId *string `json:"go_build_id,omitempty"`
	// The type of source map.
	Mapkind string `json:"mapkind"`
	// The origin of the ELF file.
	Origin *string `json:"origin,omitempty"`
	// The version of the origin package.
	OriginVersion *string `json:"origin_version,omitempty"`
	// The size of the ELF file in bytes.
	Size int64 `json:"size"`
	// The source of the debug symbols.
	SymbolSource *string `json:"symbol_source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewELFSourcemapAttributes instantiates a new ELFSourcemapAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewELFSourcemapAttributes(createdAt time.Time, mapkind string, size int64) *ELFSourcemapAttributes {
	this := ELFSourcemapAttributes{}
	this.CreatedAt = createdAt
	this.Mapkind = mapkind
	this.Size = size
	return &this
}

// NewELFSourcemapAttributesWithDefaults instantiates a new ELFSourcemapAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewELFSourcemapAttributesWithDefaults() *ELFSourcemapAttributes {
	this := ELFSourcemapAttributes{}
	return &this
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *ELFSourcemapAttributes) GetArch() string {
	if o == nil || o.Arch == nil {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ELFSourcemapAttributes) GetArchOk() (*string, bool) {
	if o == nil || o.Arch == nil {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *ELFSourcemapAttributes) HasArch() bool {
	return o != nil && o.Arch != nil
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *ELFSourcemapAttributes) SetArch(v string) {
	o.Arch = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ELFSourcemapAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ELFSourcemapAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ELFSourcemapAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetFileHash returns the FileHash field value if set, zero value otherwise.
func (o *ELFSourcemapAttributes) GetFileHash() string {
	if o == nil || o.FileHash == nil {
		var ret string
		return ret
	}
	return *o.FileHash
}

// GetFileHashOk returns a tuple with the FileHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ELFSourcemapAttributes) GetFileHashOk() (*string, bool) {
	if o == nil || o.FileHash == nil {
		return nil, false
	}
	return o.FileHash, true
}

// HasFileHash returns a boolean if a field has been set.
func (o *ELFSourcemapAttributes) HasFileHash() bool {
	return o != nil && o.FileHash != nil
}

// SetFileHash gets a reference to the given string and assigns it to the FileHash field.
func (o *ELFSourcemapAttributes) SetFileHash(v string) {
	o.FileHash = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *ELFSourcemapAttributes) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ELFSourcemapAttributes) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *ELFSourcemapAttributes) HasFileName() bool {
	return o != nil && o.FileName != nil
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *ELFSourcemapAttributes) SetFileName(v string) {
	o.FileName = &v
}

// GetGnuBuildId returns the GnuBuildId field value if set, zero value otherwise.
func (o *ELFSourcemapAttributes) GetGnuBuildId() string {
	if o == nil || o.GnuBuildId == nil {
		var ret string
		return ret
	}
	return *o.GnuBuildId
}

// GetGnuBuildIdOk returns a tuple with the GnuBuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ELFSourcemapAttributes) GetGnuBuildIdOk() (*string, bool) {
	if o == nil || o.GnuBuildId == nil {
		return nil, false
	}
	return o.GnuBuildId, true
}

// HasGnuBuildId returns a boolean if a field has been set.
func (o *ELFSourcemapAttributes) HasGnuBuildId() bool {
	return o != nil && o.GnuBuildId != nil
}

// SetGnuBuildId gets a reference to the given string and assigns it to the GnuBuildId field.
func (o *ELFSourcemapAttributes) SetGnuBuildId(v string) {
	o.GnuBuildId = &v
}

// GetGoBuildId returns the GoBuildId field value if set, zero value otherwise.
func (o *ELFSourcemapAttributes) GetGoBuildId() string {
	if o == nil || o.GoBuildId == nil {
		var ret string
		return ret
	}
	return *o.GoBuildId
}

// GetGoBuildIdOk returns a tuple with the GoBuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ELFSourcemapAttributes) GetGoBuildIdOk() (*string, bool) {
	if o == nil || o.GoBuildId == nil {
		return nil, false
	}
	return o.GoBuildId, true
}

// HasGoBuildId returns a boolean if a field has been set.
func (o *ELFSourcemapAttributes) HasGoBuildId() bool {
	return o != nil && o.GoBuildId != nil
}

// SetGoBuildId gets a reference to the given string and assigns it to the GoBuildId field.
func (o *ELFSourcemapAttributes) SetGoBuildId(v string) {
	o.GoBuildId = &v
}

// GetMapkind returns the Mapkind field value.
func (o *ELFSourcemapAttributes) GetMapkind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mapkind
}

// GetMapkindOk returns a tuple with the Mapkind field value
// and a boolean to check if the value has been set.
func (o *ELFSourcemapAttributes) GetMapkindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mapkind, true
}

// SetMapkind sets field value.
func (o *ELFSourcemapAttributes) SetMapkind(v string) {
	o.Mapkind = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *ELFSourcemapAttributes) GetOrigin() string {
	if o == nil || o.Origin == nil {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ELFSourcemapAttributes) GetOriginOk() (*string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *ELFSourcemapAttributes) HasOrigin() bool {
	return o != nil && o.Origin != nil
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *ELFSourcemapAttributes) SetOrigin(v string) {
	o.Origin = &v
}

// GetOriginVersion returns the OriginVersion field value if set, zero value otherwise.
func (o *ELFSourcemapAttributes) GetOriginVersion() string {
	if o == nil || o.OriginVersion == nil {
		var ret string
		return ret
	}
	return *o.OriginVersion
}

// GetOriginVersionOk returns a tuple with the OriginVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ELFSourcemapAttributes) GetOriginVersionOk() (*string, bool) {
	if o == nil || o.OriginVersion == nil {
		return nil, false
	}
	return o.OriginVersion, true
}

// HasOriginVersion returns a boolean if a field has been set.
func (o *ELFSourcemapAttributes) HasOriginVersion() bool {
	return o != nil && o.OriginVersion != nil
}

// SetOriginVersion gets a reference to the given string and assigns it to the OriginVersion field.
func (o *ELFSourcemapAttributes) SetOriginVersion(v string) {
	o.OriginVersion = &v
}

// GetSize returns the Size field value.
func (o *ELFSourcemapAttributes) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ELFSourcemapAttributes) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *ELFSourcemapAttributes) SetSize(v int64) {
	o.Size = v
}

// GetSymbolSource returns the SymbolSource field value if set, zero value otherwise.
func (o *ELFSourcemapAttributes) GetSymbolSource() string {
	if o == nil || o.SymbolSource == nil {
		var ret string
		return ret
	}
	return *o.SymbolSource
}

// GetSymbolSourceOk returns a tuple with the SymbolSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ELFSourcemapAttributes) GetSymbolSourceOk() (*string, bool) {
	if o == nil || o.SymbolSource == nil {
		return nil, false
	}
	return o.SymbolSource, true
}

// HasSymbolSource returns a boolean if a field has been set.
func (o *ELFSourcemapAttributes) HasSymbolSource() bool {
	return o != nil && o.SymbolSource != nil
}

// SetSymbolSource gets a reference to the given string and assigns it to the SymbolSource field.
func (o *ELFSourcemapAttributes) SetSymbolSource(v string) {
	o.SymbolSource = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ELFSourcemapAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Arch != nil {
		toSerialize["arch"] = o.Arch
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.FileHash != nil {
		toSerialize["file_hash"] = o.FileHash
	}
	if o.FileName != nil {
		toSerialize["file_name"] = o.FileName
	}
	if o.GnuBuildId != nil {
		toSerialize["gnu_build_id"] = o.GnuBuildId
	}
	if o.GoBuildId != nil {
		toSerialize["go_build_id"] = o.GoBuildId
	}
	toSerialize["mapkind"] = o.Mapkind
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if o.OriginVersion != nil {
		toSerialize["origin_version"] = o.OriginVersion
	}
	toSerialize["size"] = o.Size
	if o.SymbolSource != nil {
		toSerialize["symbol_source"] = o.SymbolSource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ELFSourcemapAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arch          *string    `json:"arch,omitempty"`
		CreatedAt     *time.Time `json:"created_at"`
		FileHash      *string    `json:"file_hash,omitempty"`
		FileName      *string    `json:"file_name,omitempty"`
		GnuBuildId    *string    `json:"gnu_build_id,omitempty"`
		GoBuildId     *string    `json:"go_build_id,omitempty"`
		Mapkind       *string    `json:"mapkind"`
		Origin        *string    `json:"origin,omitempty"`
		OriginVersion *string    `json:"origin_version,omitempty"`
		Size          *int64     `json:"size"`
		SymbolSource  *string    `json:"symbol_source,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Mapkind == nil {
		return fmt.Errorf("required field mapkind missing")
	}
	if all.Size == nil {
		return fmt.Errorf("required field size missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"arch", "created_at", "file_hash", "file_name", "gnu_build_id", "go_build_id", "mapkind", "origin", "origin_version", "size", "symbol_source"})
	} else {
		return err
	}
	o.Arch = all.Arch
	o.CreatedAt = *all.CreatedAt
	o.FileHash = all.FileHash
	o.FileName = all.FileName
	o.GnuBuildId = all.GnuBuildId
	o.GoBuildId = all.GoBuildId
	o.Mapkind = *all.Mapkind
	o.Origin = all.Origin
	o.OriginVersion = all.OriginVersion
	o.Size = *all.Size
	o.SymbolSource = all.SymbolSource

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
