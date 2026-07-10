// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IL2CPPSourcemapAttributes Attributes of an IL2CPP mapping file.
type IL2CPPSourcemapAttributes struct {
	// The build identifier (UUID format).
	BuildId *string `json:"build_id,omitempty"`
	// The timestamp when the mapping file was created.
	CreatedAt time.Time `json:"created_at"`
	// The type of source map.
	Mapkind string `json:"mapkind"`
	// The size of the mapping file in bytes.
	Size int64 `json:"size"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIL2CPPSourcemapAttributes instantiates a new IL2CPPSourcemapAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIL2CPPSourcemapAttributes(createdAt time.Time, mapkind string, size int64) *IL2CPPSourcemapAttributes {
	this := IL2CPPSourcemapAttributes{}
	this.CreatedAt = createdAt
	this.Mapkind = mapkind
	this.Size = size
	return &this
}

// NewIL2CPPSourcemapAttributesWithDefaults instantiates a new IL2CPPSourcemapAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIL2CPPSourcemapAttributesWithDefaults() *IL2CPPSourcemapAttributes {
	this := IL2CPPSourcemapAttributes{}
	return &this
}

// GetBuildId returns the BuildId field value if set, zero value otherwise.
func (o *IL2CPPSourcemapAttributes) GetBuildId() string {
	if o == nil || o.BuildId == nil {
		var ret string
		return ret
	}
	return *o.BuildId
}

// GetBuildIdOk returns a tuple with the BuildId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IL2CPPSourcemapAttributes) GetBuildIdOk() (*string, bool) {
	if o == nil || o.BuildId == nil {
		return nil, false
	}
	return o.BuildId, true
}

// HasBuildId returns a boolean if a field has been set.
func (o *IL2CPPSourcemapAttributes) HasBuildId() bool {
	return o != nil && o.BuildId != nil
}

// SetBuildId gets a reference to the given string and assigns it to the BuildId field.
func (o *IL2CPPSourcemapAttributes) SetBuildId(v string) {
	o.BuildId = &v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *IL2CPPSourcemapAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IL2CPPSourcemapAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *IL2CPPSourcemapAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetMapkind returns the Mapkind field value.
func (o *IL2CPPSourcemapAttributes) GetMapkind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mapkind
}

// GetMapkindOk returns a tuple with the Mapkind field value
// and a boolean to check if the value has been set.
func (o *IL2CPPSourcemapAttributes) GetMapkindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mapkind, true
}

// SetMapkind sets field value.
func (o *IL2CPPSourcemapAttributes) SetMapkind(v string) {
	o.Mapkind = v
}

// GetSize returns the Size field value.
func (o *IL2CPPSourcemapAttributes) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *IL2CPPSourcemapAttributes) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *IL2CPPSourcemapAttributes) SetSize(v int64) {
	o.Size = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IL2CPPSourcemapAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["size"] = o.Size

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IL2CPPSourcemapAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BuildId   *string    `json:"build_id,omitempty"`
		CreatedAt *time.Time `json:"created_at"`
		Mapkind   *string    `json:"mapkind"`
		Size      *int64     `json:"size"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"build_id", "created_at", "mapkind", "size"})
	} else {
		return err
	}
	o.BuildId = all.BuildId
	o.CreatedAt = *all.CreatedAt
	o.Mapkind = *all.Mapkind
	o.Size = *all.Size

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
