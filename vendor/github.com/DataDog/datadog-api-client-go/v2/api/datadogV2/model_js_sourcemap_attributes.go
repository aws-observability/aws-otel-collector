// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JSSourcemapAttributes Attributes of a JavaScript source map.
type JSSourcemapAttributes struct {
	// The absolute path to the minified JavaScript file.
	AbsolutePath *string `json:"absolute_path,omitempty"`
	// The path to the source map in blob storage.
	BlobStorageSourcemapPath *string `json:"blob_storage_sourcemap_path,omitempty"`
	// The build identifier.
	BuildId *string `json:"build_id,omitempty"`
	// The timestamp when the source map was created.
	CreatedAt time.Time `json:"created_at"`
	// The domain associated with the source map.
	Domain *string `json:"domain,omitempty"`
	// The file name of the minified JavaScript file.
	FileName *string `json:"file_name,omitempty"`
	// The type of source map.
	Mapkind string `json:"mapkind"`
	// The service name associated with the source map.
	Service *string `json:"service,omitempty"`
	// The size of the source map file in bytes.
	Size int64 `json:"size"`
	// The source map variant.
	Variant *string `json:"variant,omitempty"`
	// The version of the service associated with the source map.
	Version *string `json:"version,omitempty"`
	// The version code.
	VersionCode *string `json:"version_code,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJSSourcemapAttributes instantiates a new JSSourcemapAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJSSourcemapAttributes(createdAt time.Time, mapkind string, size int64) *JSSourcemapAttributes {
	this := JSSourcemapAttributes{}
	this.CreatedAt = createdAt
	this.Mapkind = mapkind
	this.Size = size
	return &this
}

// NewJSSourcemapAttributesWithDefaults instantiates a new JSSourcemapAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJSSourcemapAttributesWithDefaults() *JSSourcemapAttributes {
	this := JSSourcemapAttributes{}
	return &this
}

// GetAbsolutePath returns the AbsolutePath field value if set, zero value otherwise.
func (o *JSSourcemapAttributes) GetAbsolutePath() string {
	if o == nil || o.AbsolutePath == nil {
		var ret string
		return ret
	}
	return *o.AbsolutePath
}

// GetAbsolutePathOk returns a tuple with the AbsolutePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSSourcemapAttributes) GetAbsolutePathOk() (*string, bool) {
	if o == nil || o.AbsolutePath == nil {
		return nil, false
	}
	return o.AbsolutePath, true
}

// HasAbsolutePath returns a boolean if a field has been set.
func (o *JSSourcemapAttributes) HasAbsolutePath() bool {
	return o != nil && o.AbsolutePath != nil
}

// SetAbsolutePath gets a reference to the given string and assigns it to the AbsolutePath field.
func (o *JSSourcemapAttributes) SetAbsolutePath(v string) {
	o.AbsolutePath = &v
}

// GetBlobStorageSourcemapPath returns the BlobStorageSourcemapPath field value if set, zero value otherwise.
func (o *JSSourcemapAttributes) GetBlobStorageSourcemapPath() string {
	if o == nil || o.BlobStorageSourcemapPath == nil {
		var ret string
		return ret
	}
	return *o.BlobStorageSourcemapPath
}

// GetBlobStorageSourcemapPathOk returns a tuple with the BlobStorageSourcemapPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSSourcemapAttributes) GetBlobStorageSourcemapPathOk() (*string, bool) {
	if o == nil || o.BlobStorageSourcemapPath == nil {
		return nil, false
	}
	return o.BlobStorageSourcemapPath, true
}

// HasBlobStorageSourcemapPath returns a boolean if a field has been set.
func (o *JSSourcemapAttributes) HasBlobStorageSourcemapPath() bool {
	return o != nil && o.BlobStorageSourcemapPath != nil
}

// SetBlobStorageSourcemapPath gets a reference to the given string and assigns it to the BlobStorageSourcemapPath field.
func (o *JSSourcemapAttributes) SetBlobStorageSourcemapPath(v string) {
	o.BlobStorageSourcemapPath = &v
}

// GetBuildId returns the BuildId field value if set, zero value otherwise.
func (o *JSSourcemapAttributes) GetBuildId() string {
	if o == nil || o.BuildId == nil {
		var ret string
		return ret
	}
	return *o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSSourcemapAttributes) GetBuildIdOk() (*string, bool) {
	if o == nil || o.BuildId == nil {
		return nil, false
	}
	return o.BuildId, true
}

// HasBuildId returns a boolean if a field has been set.
func (o *JSSourcemapAttributes) HasBuildId() bool {
	return o != nil && o.BuildId != nil
}

// SetBuildId gets a reference to the given string and assigns it to the BuildId field.
func (o *JSSourcemapAttributes) SetBuildId(v string) {
	o.BuildId = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *JSSourcemapAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *JSSourcemapAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *JSSourcemapAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *JSSourcemapAttributes) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSSourcemapAttributes) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *JSSourcemapAttributes) HasDomain() bool {
	return o != nil && o.Domain != nil
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *JSSourcemapAttributes) SetDomain(v string) {
	o.Domain = &v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *JSSourcemapAttributes) GetFileName() string {
	if o == nil || o.FileName == nil {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSSourcemapAttributes) GetFileNameOk() (*string, bool) {
	if o == nil || o.FileName == nil {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *JSSourcemapAttributes) HasFileName() bool {
	return o != nil && o.FileName != nil
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *JSSourcemapAttributes) SetFileName(v string) {
	o.FileName = &v
}

// GetMapkind returns the Mapkind field value.
func (o *JSSourcemapAttributes) GetMapkind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mapkind
}

// GetMapkindOk returns a tuple with the Mapkind field value
// and a boolean to check if the value has been set.
func (o *JSSourcemapAttributes) GetMapkindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mapkind, true
}

// SetMapkind sets field value.
func (o *JSSourcemapAttributes) SetMapkind(v string) {
	o.Mapkind = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *JSSourcemapAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSSourcemapAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *JSSourcemapAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *JSSourcemapAttributes) SetService(v string) {
	o.Service = &v
}

// GetSize returns the Size field value.
func (o *JSSourcemapAttributes) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *JSSourcemapAttributes) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *JSSourcemapAttributes) SetSize(v int64) {
	o.Size = v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *JSSourcemapAttributes) GetVariant() string {
	if o == nil || o.Variant == nil {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSSourcemapAttributes) GetVariantOk() (*string, bool) {
	if o == nil || o.Variant == nil {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *JSSourcemapAttributes) HasVariant() bool {
	return o != nil && o.Variant != nil
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *JSSourcemapAttributes) SetVariant(v string) {
	o.Variant = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *JSSourcemapAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSSourcemapAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *JSSourcemapAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *JSSourcemapAttributes) SetVersion(v string) {
	o.Version = &v
}

// GetVersionCode returns the VersionCode field value if set, zero value otherwise.
func (o *JSSourcemapAttributes) GetVersionCode() string {
	if o == nil || o.VersionCode == nil {
		var ret string
		return ret
	}
	return *o.VersionCode
}

// GetVersionCodeOk returns a tuple with the VersionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSSourcemapAttributes) GetVersionCodeOk() (*string, bool) {
	if o == nil || o.VersionCode == nil {
		return nil, false
	}
	return o.VersionCode, true
}

// HasVersionCode returns a boolean if a field has been set.
func (o *JSSourcemapAttributes) HasVersionCode() bool {
	return o != nil && o.VersionCode != nil
}

// SetVersionCode gets a reference to the given string and assigns it to the VersionCode field.
func (o *JSSourcemapAttributes) SetVersionCode(v string) {
	o.VersionCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o JSSourcemapAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AbsolutePath != nil {
		toSerialize["absolute_path"] = o.AbsolutePath
	}
	if o.BlobStorageSourcemapPath != nil {
		toSerialize["blob_storage_sourcemap_path"] = o.BlobStorageSourcemapPath
	}
	if o.BuildId != nil {
		toSerialize["build_id"] = o.BuildId
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.FileName != nil {
		toSerialize["file_name"] = o.FileName
	}
	toSerialize["mapkind"] = o.Mapkind
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	toSerialize["size"] = o.Size
	if o.Variant != nil {
		toSerialize["variant"] = o.Variant
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.VersionCode != nil {
		toSerialize["version_code"] = o.VersionCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JSSourcemapAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AbsolutePath             *string    `json:"absolute_path,omitempty"`
		BlobStorageSourcemapPath *string    `json:"blob_storage_sourcemap_path,omitempty"`
		BuildId                  *string    `json:"build_id,omitempty"`
		CreatedAt                *time.Time `json:"created_at"`
		Domain                   *string    `json:"domain,omitempty"`
		FileName                 *string    `json:"file_name,omitempty"`
		Mapkind                  *string    `json:"mapkind"`
		Service                  *string    `json:"service,omitempty"`
		Size                     *int64     `json:"size"`
		Variant                  *string    `json:"variant,omitempty"`
		Version                  *string    `json:"version,omitempty"`
		VersionCode              *string    `json:"version_code,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"absolute_path", "blob_storage_sourcemap_path", "build_id", "created_at", "domain", "file_name", "mapkind", "service", "size", "variant", "version", "version_code"})
	} else {
		return err
	}
	o.AbsolutePath = all.AbsolutePath
	o.BlobStorageSourcemapPath = all.BlobStorageSourcemapPath
	o.BuildId = all.BuildId
	o.CreatedAt = *all.CreatedAt
	o.Domain = all.Domain
	o.FileName = all.FileName
	o.Mapkind = *all.Mapkind
	o.Service = all.Service
	o.Size = *all.Size
	o.Variant = all.Variant
	o.Version = all.Version
	o.VersionCode = all.VersionCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
