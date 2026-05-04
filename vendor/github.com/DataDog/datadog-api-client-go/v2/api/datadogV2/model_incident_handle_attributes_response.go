// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentHandleAttributesResponse Incident handle attributes for responses
type IncidentHandleAttributesResponse struct {
	// Timestamp when the handle was created
	CreatedAt time.Time `json:"created_at"`
	// Dynamic fields associated with the handle
	Fields IncidentHandleAttributesFields `json:"fields"`
	// Timestamp when the handle was last modified
	ModifiedAt time.Time `json:"modified_at"`
	// The handle name
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentHandleAttributesResponse instantiates a new IncidentHandleAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentHandleAttributesResponse(createdAt time.Time, fields IncidentHandleAttributesFields, modifiedAt time.Time, name string) *IncidentHandleAttributesResponse {
	this := IncidentHandleAttributesResponse{}
	this.CreatedAt = createdAt
	this.Fields = fields
	this.ModifiedAt = modifiedAt
	this.Name = name
	return &this
}

// NewIncidentHandleAttributesResponseWithDefaults instantiates a new IncidentHandleAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentHandleAttributesResponseWithDefaults() *IncidentHandleAttributesResponse {
	this := IncidentHandleAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *IncidentHandleAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *IncidentHandleAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetFields returns the Fields field value.
func (o *IncidentHandleAttributesResponse) GetFields() IncidentHandleAttributesFields {
	if o == nil {
		var ret IncidentHandleAttributesFields
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleAttributesResponse) GetFieldsOk() (*IncidentHandleAttributesFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value.
func (o *IncidentHandleAttributesResponse) SetFields(v IncidentHandleAttributesFields) {
	o.Fields = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *IncidentHandleAttributesResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleAttributesResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *IncidentHandleAttributesResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetName returns the Name field value.
func (o *IncidentHandleAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *IncidentHandleAttributesResponse) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentHandleAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["fields"] = o.Fields
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentHandleAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt  *time.Time                      `json:"created_at"`
		Fields     *IncidentHandleAttributesFields `json:"fields"`
		ModifiedAt *time.Time                      `json:"modified_at"`
		Name       *string                         `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Fields == nil {
		return fmt.Errorf("required field fields missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "fields", "modified_at", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CreatedAt = *all.CreatedAt
	if all.Fields.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Fields = *all.Fields
	o.ModifiedAt = *all.ModifiedAt
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
