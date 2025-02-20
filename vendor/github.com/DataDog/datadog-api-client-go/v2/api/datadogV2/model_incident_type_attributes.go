// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTypeAttributes Incident type's attributes.
type IncidentTypeAttributes struct {
	// Timestamp when the incident type was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A unique identifier that represents the user that created the incident type.
	CreatedBy *string `json:"createdBy,omitempty"`
	// Text that describes the incident type.
	Description *string `json:"description,omitempty"`
	// If true, this incident type will be used as the default incident type if a type is not specified during the creation of incident resources.
	IsDefault *bool `json:"is_default,omitempty"`
	// A unique identifier that represents the user that last modified the incident type.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	// Timestamp when the incident type was last modified.
	ModifiedAt *time.Time `json:"modifiedAt,omitempty"`
	// The name of the incident type.
	Name string `json:"name"`
	// The string that will be prepended to the incident title across the Datadog app.
	Prefix *string `json:"prefix,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTypeAttributes instantiates a new IncidentTypeAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTypeAttributes(name string) *IncidentTypeAttributes {
	this := IncidentTypeAttributes{}
	var isDefault bool = false
	this.IsDefault = &isDefault
	this.Name = name
	return &this
}

// NewIncidentTypeAttributesWithDefaults instantiates a new IncidentTypeAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTypeAttributesWithDefaults() *IncidentTypeAttributes {
	this := IncidentTypeAttributes{}
	var isDefault bool = false
	this.IsDefault = &isDefault
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IncidentTypeAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IncidentTypeAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *IncidentTypeAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *IncidentTypeAttributes) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *IncidentTypeAttributes) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *IncidentTypeAttributes) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IncidentTypeAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IncidentTypeAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IncidentTypeAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *IncidentTypeAttributes) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeAttributes) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *IncidentTypeAttributes) HasIsDefault() bool {
	return o != nil && o.IsDefault != nil
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *IncidentTypeAttributes) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *IncidentTypeAttributes) GetLastModifiedBy() string {
	if o == nil || o.LastModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeAttributes) GetLastModifiedByOk() (*string, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *IncidentTypeAttributes) HasLastModifiedBy() bool {
	return o != nil && o.LastModifiedBy != nil
}

// SetLastModifiedBy gets a reference to the given string and assigns it to the LastModifiedBy field.
func (o *IncidentTypeAttributes) SetLastModifiedBy(v string) {
	o.LastModifiedBy = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *IncidentTypeAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *IncidentTypeAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *IncidentTypeAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value.
func (o *IncidentTypeAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentTypeAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentTypeAttributes) SetName(v string) {
	o.Name = v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *IncidentTypeAttributes) GetPrefix() string {
	if o == nil || o.Prefix == nil {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTypeAttributes) GetPrefixOk() (*string, bool) {
	if o == nil || o.Prefix == nil {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *IncidentTypeAttributes) HasPrefix() bool {
	return o != nil && o.Prefix != nil
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *IncidentTypeAttributes) SetPrefix(v string) {
	o.Prefix = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTypeAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["name"] = o.Name
	if o.Prefix != nil {
		toSerialize["prefix"] = o.Prefix
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentTypeAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt      *time.Time `json:"createdAt,omitempty"`
		CreatedBy      *string    `json:"createdBy,omitempty"`
		Description    *string    `json:"description,omitempty"`
		IsDefault      *bool      `json:"is_default,omitempty"`
		LastModifiedBy *string    `json:"lastModifiedBy,omitempty"`
		ModifiedAt     *time.Time `json:"modifiedAt,omitempty"`
		Name           *string    `json:"name"`
		Prefix         *string    `json:"prefix,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
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
	o.Name = *all.Name
	o.Prefix = all.Prefix

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
