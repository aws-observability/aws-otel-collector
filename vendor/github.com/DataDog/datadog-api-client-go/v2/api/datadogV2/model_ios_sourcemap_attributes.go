// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IOSSourcemapAttributes Attributes of an iOS dSYM source map.
type IOSSourcemapAttributes struct {
	// The timestamp when the source map was created.
	CreatedAt time.Time `json:"created_at"`
	// The type of source map.
	Mapkind string `json:"mapkind"`
	// The size of the dSYM file in bytes.
	Size int64 `json:"size"`
	// The UUID(s) associated with the dSYM file.
	Uuids *string `json:"uuids,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIOSSourcemapAttributes instantiates a new IOSSourcemapAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIOSSourcemapAttributes(createdAt time.Time, mapkind string, size int64) *IOSSourcemapAttributes {
	this := IOSSourcemapAttributes{}
	this.CreatedAt = createdAt
	this.Mapkind = mapkind
	this.Size = size
	return &this
}

// NewIOSSourcemapAttributesWithDefaults instantiates a new IOSSourcemapAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIOSSourcemapAttributesWithDefaults() *IOSSourcemapAttributes {
	this := IOSSourcemapAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *IOSSourcemapAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IOSSourcemapAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *IOSSourcemapAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetMapkind returns the Mapkind field value.
func (o *IOSSourcemapAttributes) GetMapkind() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mapkind
}

// GetMapkindOk returns a tuple with the Mapkind field value
// and a boolean to check if the value has been set.
func (o *IOSSourcemapAttributes) GetMapkindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mapkind, true
}

// SetMapkind sets field value.
func (o *IOSSourcemapAttributes) SetMapkind(v string) {
	o.Mapkind = v
}

// GetSize returns the Size field value.
func (o *IOSSourcemapAttributes) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *IOSSourcemapAttributes) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value.
func (o *IOSSourcemapAttributes) SetSize(v int64) {
	o.Size = v
}

// GetUuids returns the Uuids field value if set, zero value otherwise.
func (o *IOSSourcemapAttributes) GetUuids() string {
	if o == nil || o.Uuids == nil {
		var ret string
		return ret
	}
	return *o.Uuids
}

// GetUuidsOk returns a tuple with the Uuids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IOSSourcemapAttributes) GetUuidsOk() (*string, bool) {
	if o == nil || o.Uuids == nil {
		return nil, false
	}
	return o.Uuids, true
}

// HasUuids returns a boolean if a field has been set.
func (o *IOSSourcemapAttributes) HasUuids() bool {
	return o != nil && o.Uuids != nil
}

// SetUuids gets a reference to the given string and assigns it to the Uuids field.
func (o *IOSSourcemapAttributes) SetUuids(v string) {
	o.Uuids = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IOSSourcemapAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["mapkind"] = o.Mapkind
	toSerialize["size"] = o.Size
	if o.Uuids != nil {
		toSerialize["uuids"] = o.Uuids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IOSSourcemapAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt *time.Time `json:"created_at"`
		Mapkind   *string    `json:"mapkind"`
		Size      *int64     `json:"size"`
		Uuids     *string    `json:"uuids,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "mapkind", "size", "uuids"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.Mapkind = *all.Mapkind
	o.Size = *all.Size
	o.Uuids = all.Uuids

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
