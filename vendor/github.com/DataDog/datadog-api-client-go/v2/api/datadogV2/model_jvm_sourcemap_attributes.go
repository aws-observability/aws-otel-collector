// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JVMSourcemapAttributes Attributes of a JVM mapping file.
type JVMSourcemapAttributes struct {
	// The build identifier (UUID format).
	BuildId *string `json:"build_id,omitempty"`
	// The timestamp when the mapping file was created.
	CreatedAt time.Time `json:"created_at"`
	// The type of source map.
	Mapkind string `json:"mapkind"`
	// The service name associated with the mapping file.
	Service *string `json:"service,omitempty"`
	// The size of the mapping file in bytes.
	Size int64 `json:"size"`
	// The build variant (e.g., `release`, `debug`).
	Variant *string `json:"variant,omitempty"`
	// The version of the service associated with the mapping file.
	Version *string `json:"version,omitempty"`
	// The version code.
	VersionCode *string `json:"version_code,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJVMSourcemapAttributes instantiates a new JVMSourcemapAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJVMSourcemapAttributes(createdAt time.Time, mapkind string, size int64) *JVMSourcemapAttributes {
	this := JVMSourcemapAttributes{}
	this.CreatedAt = createdAt
	this.Mapkind = mapkind
	this.Size = size
	return &this
}

// NewJVMSourcemapAttributesWithDefaults instantiates a new JVMSourcemapAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJVMSourcemapAttributesWithDefaults() *JVMSourcemapAttributes {
	this := JVMSourcemapAttributes{}
	return &this
}

// GetBuildId returns the BuildId field value if set, zero value otherwise.
func (o *JVMSourcemapAttributes) GetBuildId() string {
	if o == nil || o.BuildId == nil {
		var ret string
		return ret
	}
	return *o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JVMSourcemapAttributes) GetBuildIdOk() (*string, bool) {
	if o == nil || o.BuildId == nil {
		return nil, false
	}
	return o.BuildId, true
}

// HasBuildId returns a boolean if a field has been set.
func (o *JVMSourcemapAttributes) HasBuildId() bool {
	return o != nil && o.BuildId != nil
}

// SetBuildId gets a reference to the given string and assigns it to the BuildId field.
func (o *JVMSourcemapAttributes) SetBuildId(v string) {
	o.BuildId = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *JVMSourcemapAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *JVMSourcemapAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *JVMSourcemapAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetMapkind returns the Mapkind field value.
func (o *JVMSourcemapAttributes) GetMapkind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mapkind
}

// GetMapkindOk returns a tuple with the Mapkind field value
// and a boolean to check if the value has been set.
func (o *JVMSourcemapAttributes) GetMapkindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mapkind, true
}

// SetMapkind sets field value.
func (o *JVMSourcemapAttributes) SetMapkind(v string) {
	o.Mapkind = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *JVMSourcemapAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JVMSourcemapAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *JVMSourcemapAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *JVMSourcemapAttributes) SetService(v string) {
	o.Service = &v
}

// GetSize returns the Size field value.
func (o *JVMSourcemapAttributes) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *JVMSourcemapAttributes) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *JVMSourcemapAttributes) SetSize(v int64) {
	o.Size = v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *JVMSourcemapAttributes) GetVariant() string {
	if o == nil || o.Variant == nil {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JVMSourcemapAttributes) GetVariantOk() (*string, bool) {
	if o == nil || o.Variant == nil {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *JVMSourcemapAttributes) HasVariant() bool {
	return o != nil && o.Variant != nil
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *JVMSourcemapAttributes) SetVariant(v string) {
	o.Variant = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *JVMSourcemapAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JVMSourcemapAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *JVMSourcemapAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *JVMSourcemapAttributes) SetVersion(v string) {
	o.Version = &v
}

// GetVersionCode returns the VersionCode field value if set, zero value otherwise.
func (o *JVMSourcemapAttributes) GetVersionCode() string {
	if o == nil || o.VersionCode == nil {
		var ret string
		return ret
	}
	return *o.VersionCode
}

// GetVersionCodeOk returns a tuple with the VersionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JVMSourcemapAttributes) GetVersionCodeOk() (*string, bool) {
	if o == nil || o.VersionCode == nil {
		return nil, false
	}
	return o.VersionCode, true
}

// HasVersionCode returns a boolean if a field has been set.
func (o *JVMSourcemapAttributes) HasVersionCode() bool {
	return o != nil && o.VersionCode != nil
}

// SetVersionCode gets a reference to the given string and assigns it to the VersionCode field.
func (o *JVMSourcemapAttributes) SetVersionCode(v string) {
	o.VersionCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o JVMSourcemapAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BuildId != nil {
		toSerialize["build_id"] = o.BuildId
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
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
func (o *JVMSourcemapAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BuildId     *string    `json:"build_id,omitempty"`
		CreatedAt   *time.Time `json:"created_at"`
		Mapkind     *string    `json:"mapkind"`
		Service     *string    `json:"service,omitempty"`
		Size        *int64     `json:"size"`
		Variant     *string    `json:"variant,omitempty"`
		Version     *string    `json:"version,omitempty"`
		VersionCode *string    `json:"version_code,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"build_id", "created_at", "mapkind", "service", "size", "variant", "version", "version_code"})
	} else {
		return err
	}
	o.BuildId = all.BuildId
	o.CreatedAt = *all.CreatedAt
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
