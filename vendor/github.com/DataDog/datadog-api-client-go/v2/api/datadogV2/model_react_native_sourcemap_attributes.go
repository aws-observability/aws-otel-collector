// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReactNativeSourcemapAttributes Attributes of a React Native source map.
type ReactNativeSourcemapAttributes struct {
	// The build number.
	BuildNumber *string `json:"build_number,omitempty"`
	// The bundle name.
	BundleName *string `json:"bundle_name,omitempty"`
	// The bundle version.
	BundleVersion *string `json:"bundle_version,omitempty"`
	// The timestamp when the source map was created.
	CreatedAt time.Time `json:"created_at"`
	// The debug identifier (UUID format).
	DebugId *string `json:"debug_id,omitempty"`
	// The type of source map.
	Mapkind string `json:"mapkind"`
	// The platform the source map was built for (e.g., `ios`, `android`).
	Platform *string `json:"platform,omitempty"`
	// The service name associated with the source map.
	Service *string `json:"service,omitempty"`
	// The size of the source map file in bytes.
	Size int64 `json:"size"`
	// The version of the service associated with the source map.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReactNativeSourcemapAttributes instantiates a new ReactNativeSourcemapAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReactNativeSourcemapAttributes(createdAt time.Time, mapkind string, size int64) *ReactNativeSourcemapAttributes {
	this := ReactNativeSourcemapAttributes{}
	this.CreatedAt = createdAt
	this.Mapkind = mapkind
	this.Size = size
	return &this
}

// NewReactNativeSourcemapAttributesWithDefaults instantiates a new ReactNativeSourcemapAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReactNativeSourcemapAttributesWithDefaults() *ReactNativeSourcemapAttributes {
	this := ReactNativeSourcemapAttributes{}
	return &this
}

// GetBuildNumber returns the BuildNumber field value if set, zero value otherwise.
func (o *ReactNativeSourcemapAttributes) GetBuildNumber() string {
	if o == nil || o.BuildNumber == nil {
		var ret string
		return ret
	}
	return *o.BuildNumber
}

// GetBuildNumberOk returns a tuple with the BuildNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReactNativeSourcemapAttributes) GetBuildNumberOk() (*string, bool) {
	if o == nil || o.BuildNumber == nil {
		return nil, false
	}
	return o.BuildNumber, true
}

// HasBuildNumber returns a boolean if a field has been set.
func (o *ReactNativeSourcemapAttributes) HasBuildNumber() bool {
	return o != nil && o.BuildNumber != nil
}

// SetBuildNumber gets a reference to the given string and assigns it to the BuildNumber field.
func (o *ReactNativeSourcemapAttributes) SetBuildNumber(v string) {
	o.BuildNumber = &v
}

// GetBundleName returns the BundleName field value if set, zero value otherwise.
func (o *ReactNativeSourcemapAttributes) GetBundleName() string {
	if o == nil || o.BundleName == nil {
		var ret string
		return ret
	}
	return *o.BundleName
}

// GetBundleNameOk returns a tuple with the BundleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReactNativeSourcemapAttributes) GetBundleNameOk() (*string, bool) {
	if o == nil || o.BundleName == nil {
		return nil, false
	}
	return o.BundleName, true
}

// HasBundleName returns a boolean if a field has been set.
func (o *ReactNativeSourcemapAttributes) HasBundleName() bool {
	return o != nil && o.BundleName != nil
}

// SetBundleName gets a reference to the given string and assigns it to the BundleName field.
func (o *ReactNativeSourcemapAttributes) SetBundleName(v string) {
	o.BundleName = &v
}

// GetBundleVersion returns the BundleVersion field value if set, zero value otherwise.
func (o *ReactNativeSourcemapAttributes) GetBundleVersion() string {
	if o == nil || o.BundleVersion == nil {
		var ret string
		return ret
	}
	return *o.BundleVersion
}

// GetBundleVersionOk returns a tuple with the BundleVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReactNativeSourcemapAttributes) GetBundleVersionOk() (*string, bool) {
	if o == nil || o.BundleVersion == nil {
		return nil, false
	}
	return o.BundleVersion, true
}

// HasBundleVersion returns a boolean if a field has been set.
func (o *ReactNativeSourcemapAttributes) HasBundleVersion() bool {
	return o != nil && o.BundleVersion != nil
}

// SetBundleVersion gets a reference to the given string and assigns it to the BundleVersion field.
func (o *ReactNativeSourcemapAttributes) SetBundleVersion(v string) {
	o.BundleVersion = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *ReactNativeSourcemapAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ReactNativeSourcemapAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *ReactNativeSourcemapAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDebugId returns the DebugId field value if set, zero value otherwise.
func (o *ReactNativeSourcemapAttributes) GetDebugId() string {
	if o == nil || o.DebugId == nil {
		var ret string
		return ret
	}
	return *o.DebugId
}

// GetDebugIdOk returns a tuple with the DebugId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReactNativeSourcemapAttributes) GetDebugIdOk() (*string, bool) {
	if o == nil || o.DebugId == nil {
		return nil, false
	}
	return o.DebugId, true
}

// HasDebugId returns a boolean if a field has been set.
func (o *ReactNativeSourcemapAttributes) HasDebugId() bool {
	return o != nil && o.DebugId != nil
}

// SetDebugId gets a reference to the given string and assigns it to the DebugId field.
func (o *ReactNativeSourcemapAttributes) SetDebugId(v string) {
	o.DebugId = &v
}

// GetMapkind returns the Mapkind field value.
func (o *ReactNativeSourcemapAttributes) GetMapkind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mapkind
}

// GetMapkindOk returns a tuple with the Mapkind field value
// and a boolean to check if the value has been set.
func (o *ReactNativeSourcemapAttributes) GetMapkindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mapkind, true
}

// SetMapkind sets field value.
func (o *ReactNativeSourcemapAttributes) SetMapkind(v string) {
	o.Mapkind = v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *ReactNativeSourcemapAttributes) GetPlatform() string {
	if o == nil || o.Platform == nil {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReactNativeSourcemapAttributes) GetPlatformOk() (*string, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *ReactNativeSourcemapAttributes) HasPlatform() bool {
	return o != nil && o.Platform != nil
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *ReactNativeSourcemapAttributes) SetPlatform(v string) {
	o.Platform = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ReactNativeSourcemapAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReactNativeSourcemapAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ReactNativeSourcemapAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ReactNativeSourcemapAttributes) SetService(v string) {
	o.Service = &v
}

// GetSize returns the Size field value.
func (o *ReactNativeSourcemapAttributes) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ReactNativeSourcemapAttributes) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *ReactNativeSourcemapAttributes) SetSize(v int64) {
	o.Size = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ReactNativeSourcemapAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReactNativeSourcemapAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ReactNativeSourcemapAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ReactNativeSourcemapAttributes) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReactNativeSourcemapAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BuildNumber != nil {
		toSerialize["build_number"] = o.BuildNumber
	}
	if o.BundleName != nil {
		toSerialize["bundle_name"] = o.BundleName
	}
	if o.BundleVersion != nil {
		toSerialize["bundle_version"] = o.BundleVersion
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.DebugId != nil {
		toSerialize["debug_id"] = o.DebugId
	}
	toSerialize["mapkind"] = o.Mapkind
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	toSerialize["size"] = o.Size
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReactNativeSourcemapAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BuildNumber   *string    `json:"build_number,omitempty"`
		BundleName    *string    `json:"bundle_name,omitempty"`
		BundleVersion *string    `json:"bundle_version,omitempty"`
		CreatedAt     *time.Time `json:"created_at"`
		DebugId       *string    `json:"debug_id,omitempty"`
		Mapkind       *string    `json:"mapkind"`
		Platform      *string    `json:"platform,omitempty"`
		Service       *string    `json:"service,omitempty"`
		Size          *int64     `json:"size"`
		Version       *string    `json:"version,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"build_number", "bundle_name", "bundle_version", "created_at", "debug_id", "mapkind", "platform", "service", "size", "version"})
	} else {
		return err
	}
	o.BuildNumber = all.BuildNumber
	o.BundleName = all.BundleName
	o.BundleVersion = all.BundleVersion
	o.CreatedAt = *all.CreatedAt
	o.DebugId = all.DebugId
	o.Mapkind = *all.Mapkind
	o.Platform = all.Platform
	o.Service = all.Service
	o.Size = *all.Size
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
