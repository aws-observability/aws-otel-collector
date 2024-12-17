// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityMeta Entity metadata.
type EntityMeta struct {
	// The creation time.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The ingestion source.
	IngestionSource *string `json:"ingestionSource,omitempty"`
	// The modification time.
	ModifiedAt *string `json:"modifiedAt,omitempty"`
	// The origin.
	Origin *string `json:"origin,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityMeta instantiates a new EntityMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityMeta() *EntityMeta {
	this := EntityMeta{}
	return &this
}

// NewEntityMetaWithDefaults instantiates a new EntityMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityMetaWithDefaults() *EntityMeta {
	this := EntityMeta{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EntityMeta) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityMeta) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EntityMeta) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *EntityMeta) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetIngestionSource returns the IngestionSource field value if set, zero value otherwise.
func (o *EntityMeta) GetIngestionSource() string {
	if o == nil || o.IngestionSource == nil {
		var ret string
		return ret
	}
	return *o.IngestionSource
}

// GetIngestionSourceOk returns a tuple with the IngestionSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityMeta) GetIngestionSourceOk() (*string, bool) {
	if o == nil || o.IngestionSource == nil {
		return nil, false
	}
	return o.IngestionSource, true
}

// HasIngestionSource returns a boolean if a field has been set.
func (o *EntityMeta) HasIngestionSource() bool {
	return o != nil && o.IngestionSource != nil
}

// SetIngestionSource gets a reference to the given string and assigns it to the IngestionSource field.
func (o *EntityMeta) SetIngestionSource(v string) {
	o.IngestionSource = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *EntityMeta) GetModifiedAt() string {
	if o == nil || o.ModifiedAt == nil {
		var ret string
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityMeta) GetModifiedAtOk() (*string, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *EntityMeta) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given string and assigns it to the ModifiedAt field.
func (o *EntityMeta) SetModifiedAt(v string) {
	o.ModifiedAt = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *EntityMeta) GetOrigin() string {
	if o == nil || o.Origin == nil {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityMeta) GetOriginOk() (*string, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *EntityMeta) HasOrigin() bool {
	return o != nil && o.Origin != nil
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *EntityMeta) SetOrigin(v string) {
	o.Origin = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.IngestionSource != nil {
		toSerialize["ingestionSource"] = o.IngestionSource
	}
	if o.ModifiedAt != nil {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt       *string `json:"createdAt,omitempty"`
		IngestionSource *string `json:"ingestionSource,omitempty"`
		ModifiedAt      *string `json:"modifiedAt,omitempty"`
		Origin          *string `json:"origin,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"createdAt", "ingestionSource", "modifiedAt", "origin"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.IngestionSource = all.IngestionSource
	o.ModifiedAt = all.ModifiedAt
	o.Origin = all.Origin

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
