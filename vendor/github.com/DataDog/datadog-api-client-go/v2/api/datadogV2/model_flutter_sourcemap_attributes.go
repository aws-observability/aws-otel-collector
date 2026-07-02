// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlutterSourcemapAttributes Attributes of a Flutter symbol file.
type FlutterSourcemapAttributes struct {
	// The target CPU architecture.
	Arch *string `json:"arch,omitempty"`
	// The timestamp when the symbol file was created.
	CreatedAt time.Time `json:"created_at"`
	// The type of source map.
	Mapkind string `json:"mapkind"`
	// The service name associated with the symbol file.
	Service *string `json:"service,omitempty"`
	// The size of the symbol file in bytes.
	Size int64 `json:"size"`
	// The build variant.
	Variant *string `json:"variant,omitempty"`
	// The version of the service associated with the symbol file.
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFlutterSourcemapAttributes instantiates a new FlutterSourcemapAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFlutterSourcemapAttributes(createdAt time.Time, mapkind string, size int64) *FlutterSourcemapAttributes {
	this := FlutterSourcemapAttributes{}
	this.CreatedAt = createdAt
	this.Mapkind = mapkind
	this.Size = size
	return &this
}

// NewFlutterSourcemapAttributesWithDefaults instantiates a new FlutterSourcemapAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFlutterSourcemapAttributesWithDefaults() *FlutterSourcemapAttributes {
	this := FlutterSourcemapAttributes{}
	return &this
}

// GetArch returns the Arch field value if set, zero value otherwise.
func (o *FlutterSourcemapAttributes) GetArch() string {
	if o == nil || o.Arch == nil {
		var ret string
		return ret
	}
	return *o.Arch
}

// GetArchOk returns a tuple with the Arch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlutterSourcemapAttributes) GetArchOk() (*string, bool) {
	if o == nil || o.Arch == nil {
		return nil, false
	}
	return o.Arch, true
}

// HasArch returns a boolean if a field has been set.
func (o *FlutterSourcemapAttributes) HasArch() bool {
	return o != nil && o.Arch != nil
}

// SetArch gets a reference to the given string and assigns it to the Arch field.
func (o *FlutterSourcemapAttributes) SetArch(v string) {
	o.Arch = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *FlutterSourcemapAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *FlutterSourcemapAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *FlutterSourcemapAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetMapkind returns the Mapkind field value.
func (o *FlutterSourcemapAttributes) GetMapkind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mapkind
}

// GetMapkindOk returns a tuple with the Mapkind field value
// and a boolean to check if the value has been set.
func (o *FlutterSourcemapAttributes) GetMapkindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mapkind, true
}

// SetMapkind sets field value.
func (o *FlutterSourcemapAttributes) SetMapkind(v string) {
	o.Mapkind = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *FlutterSourcemapAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlutterSourcemapAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *FlutterSourcemapAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *FlutterSourcemapAttributes) SetService(v string) {
	o.Service = &v
}

// GetSize returns the Size field value.
func (o *FlutterSourcemapAttributes) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *FlutterSourcemapAttributes) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *FlutterSourcemapAttributes) SetSize(v int64) {
	o.Size = v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *FlutterSourcemapAttributes) GetVariant() string {
	if o == nil || o.Variant == nil {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlutterSourcemapAttributes) GetVariantOk() (*string, bool) {
	if o == nil || o.Variant == nil {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *FlutterSourcemapAttributes) HasVariant() bool {
	return o != nil && o.Variant != nil
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *FlutterSourcemapAttributes) SetVariant(v string) {
	o.Variant = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *FlutterSourcemapAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlutterSourcemapAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *FlutterSourcemapAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *FlutterSourcemapAttributes) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FlutterSourcemapAttributes) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FlutterSourcemapAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Arch      *string    `json:"arch,omitempty"`
		CreatedAt *time.Time `json:"created_at"`
		Mapkind   *string    `json:"mapkind"`
		Service   *string    `json:"service,omitempty"`
		Size      *int64     `json:"size"`
		Variant   *string    `json:"variant,omitempty"`
		Version   *string    `json:"version,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"arch", "created_at", "mapkind", "service", "size", "variant", "version"})
	} else {
		return err
	}
	o.Arch = all.Arch
	o.CreatedAt = *all.CreatedAt
	o.Mapkind = *all.Mapkind
	o.Service = all.Service
	o.Size = *all.Size
	o.Variant = all.Variant
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
