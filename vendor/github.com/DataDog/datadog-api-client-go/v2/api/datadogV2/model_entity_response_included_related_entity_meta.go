// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseIncludedRelatedEntityMeta Included related entity meta.
type EntityResponseIncludedRelatedEntityMeta struct {
	// Entity creation time.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Entity relation defined by.
	DefinedBy *string `json:"defined_by,omitempty"`
	// Entity modification time.
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// Entity relation source.
	Source *string `json:"source,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityResponseIncludedRelatedEntityMeta instantiates a new EntityResponseIncludedRelatedEntityMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityResponseIncludedRelatedEntityMeta() *EntityResponseIncludedRelatedEntityMeta {
	this := EntityResponseIncludedRelatedEntityMeta{}
	return &this
}

// NewEntityResponseIncludedRelatedEntityMetaWithDefaults instantiates a new EntityResponseIncludedRelatedEntityMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityResponseIncludedRelatedEntityMetaWithDefaults() *EntityResponseIncludedRelatedEntityMeta {
	this := EntityResponseIncludedRelatedEntityMeta{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedEntityMeta) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedEntityMeta) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedEntityMeta) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EntityResponseIncludedRelatedEntityMeta) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDefinedBy returns the DefinedBy field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedEntityMeta) GetDefinedBy() string {
	if o == nil || o.DefinedBy == nil {
		var ret string
		return ret
	}
	return *o.DefinedBy
}

// GetDefinedByOk returns a tuple with the DefinedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedEntityMeta) GetDefinedByOk() (*string, bool) {
	if o == nil || o.DefinedBy == nil {
		return nil, false
	}
	return o.DefinedBy, true
}

// HasDefinedBy returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedEntityMeta) HasDefinedBy() bool {
	return o != nil && o.DefinedBy != nil
}

// SetDefinedBy gets a reference to the given string and assigns it to the DefinedBy field.
func (o *EntityResponseIncludedRelatedEntityMeta) SetDefinedBy(v string) {
	o.DefinedBy = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedEntityMeta) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedEntityMeta) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedEntityMeta) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *EntityResponseIncludedRelatedEntityMeta) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *EntityResponseIncludedRelatedEntityMeta) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseIncludedRelatedEntityMeta) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *EntityResponseIncludedRelatedEntityMeta) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *EntityResponseIncludedRelatedEntityMeta) SetSource(v string) {
	o.Source = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityResponseIncludedRelatedEntityMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["createdAt"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DefinedBy != nil {
		toSerialize["defined_by"] = o.DefinedBy
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modifiedAt"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modifiedAt"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityResponseIncludedRelatedEntityMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt  *time.Time `json:"createdAt,omitempty"`
		DefinedBy  *string    `json:"defined_by,omitempty"`
		ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
		Source     *string    `json:"source,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"createdAt", "defined_by", "modifiedAt", "source"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.DefinedBy = all.DefinedBy
	o.ModifiedAt = all.ModifiedAt
	o.Source = all.Source

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
