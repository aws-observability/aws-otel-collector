// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTypeUpdateAttributes Incident type's attributes for updates.
type IncidentTypeUpdateAttributes struct {
	// Timestamp when the incident type was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A unique identifier that represents the user that created the incident type.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Text that describes the incident type.
	Description *string `json:"description,omitempty"`
	// When true, this incident type will be used as the default type when an incident type is not specified.
	IsDefault *bool `json:"is_default,omitempty"`
	// A unique identifier that represents the user that last modified the incident type.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// Timestamp when the incident type was last modified.
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// The name of the incident type.
	Name *string `json:"name,omitempty"`
	// The string that will be prepended to the incident title across the Datadog app.
	Prefix *string `json:"prefix,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTypeUpdateAttributes instantiates a new IncidentTypeUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTypeUpdateAttributes() *IncidentTypeUpdateAttributes {
	this := IncidentTypeUpdateAttributes{}
	return &this
}

// NewIncidentTypeUpdateAttributesWithDefaults instantiates a new IncidentTypeUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTypeUpdateAttributesWithDefaults() *IncidentTypeUpdateAttributes {
	this := IncidentTypeUpdateAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IncidentTypeUpdateAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeUpdateAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IncidentTypeUpdateAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IncidentTypeUpdateAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *IncidentTypeUpdateAttributes) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeUpdateAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *IncidentTypeUpdateAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *IncidentTypeUpdateAttributes) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IncidentTypeUpdateAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeUpdateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IncidentTypeUpdateAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IncidentTypeUpdateAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *IncidentTypeUpdateAttributes) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeUpdateAttributes) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *IncidentTypeUpdateAttributes) HasIsDefault() bool {
	return o != nil && o.IsDefault != nil
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *IncidentTypeUpdateAttributes) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *IncidentTypeUpdateAttributes) GetLastModifiedBy() string {
	if o == nil || o.LastModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeUpdateAttributes) GetLastModifiedByOk() (*string, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *IncidentTypeUpdateAttributes) HasLastModifiedBy() bool {
	return o != nil && o.LastModifiedBy != nil
}

// SetLastModifiedBy gets a reference to the given string and assigns it to the LastModifiedBy field.
func (o *IncidentTypeUpdateAttributes) SetLastModifiedBy(v string) {
	o.LastModifiedBy = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *IncidentTypeUpdateAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeUpdateAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *IncidentTypeUpdateAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *IncidentTypeUpdateAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IncidentTypeUpdateAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IncidentTypeUpdateAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IncidentTypeUpdateAttributes) SetName(v string) {
	o.Name = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *IncidentTypeUpdateAttributes) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeUpdateAttributes) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *IncidentTypeUpdateAttributes) HasPrefix() bool {
	return o != nil && o.Prefix != nil
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *IncidentTypeUpdateAttributes) SetPrefix(v string) {
	o.Prefix = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTypeUpdateAttributes) MarshalJSON() ([]byte, error) {
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
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.IsDefault != nil {
		toSerialize["is_default"] = o.IsDefault
	}
	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modifiedAt"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modifiedAt"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTypeUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt      *time.Time `json:"createdAt,omitempty"`
		CreatedBy      *string    `json:"createdBy,omitempty"`
		Description    *string    `json:"description,omitempty"`
		IsDefault      *bool      `json:"is_default,omitempty"`
		LastModifiedBy *string    `json:"lastModifiedBy,omitempty"`
		ModifiedAt     *time.Time `json:"modifiedAt,omitempty"`
		Name           *string    `json:"name,omitempty"`
		Prefix         *string    `json:"prefix,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"createdAt", "createdBy", "description", "is_default", "lastModifiedBy", "modifiedAt", "name", "prefix"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.CreatedBy = all.CreatedBy
	o.Description = all.Description
	o.IsDefault = all.IsDefault
	o.LastModifiedBy = all.LastModifiedBy
	o.ModifiedAt = all.ModifiedAt
	o.Name = all.Name
	o.Prefix = all.Prefix

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
