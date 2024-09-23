// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentAttachmentUpdateData A single incident attachment.
type IncidentAttachmentUpdateData struct {
	// Incident attachment attributes.
	Attributes *IncidentAttachmentUpdateAttributes `json:"attributes,omitempty"`
	// A unique identifier that represents the incident attachment.
	Id *string `json:"id,omitempty"`
	// The incident attachment resource type.
	Type IncidentAttachmentType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentAttachmentUpdateData instantiates a new IncidentAttachmentUpdateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentAttachmentUpdateData(typeVar IncidentAttachmentType) *IncidentAttachmentUpdateData {
	this := IncidentAttachmentUpdateData{}
	this.Type = typeVar
	return &this
}

// NewIncidentAttachmentUpdateDataWithDefaults instantiates a new IncidentAttachmentUpdateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentAttachmentUpdateDataWithDefaults() *IncidentAttachmentUpdateData {
	this := IncidentAttachmentUpdateData{}
	var typeVar IncidentAttachmentType = INCIDENTATTACHMENTTYPE_INCIDENT_ATTACHMENTS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentAttachmentUpdateData) GetAttributes() IncidentAttachmentUpdateAttributes {
	if o == nil || o.Attributes == nil {
		var ret IncidentAttachmentUpdateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAttachmentUpdateData) GetAttributesOk() (*IncidentAttachmentUpdateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentAttachmentUpdateData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given IncidentAttachmentUpdateAttributes and assigns it to the Attributes field.
func (o *IncidentAttachmentUpdateData) SetAttributes(v IncidentAttachmentUpdateAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IncidentAttachmentUpdateData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAttachmentUpdateData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IncidentAttachmentUpdateData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IncidentAttachmentUpdateData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value.
func (o *IncidentAttachmentUpdateData) GetType() IncidentAttachmentType {
	if o == nil {
		var ret IncidentAttachmentType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentAttachmentUpdateData) GetTypeOk() (*IncidentAttachmentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentAttachmentUpdateData) SetType(v IncidentAttachmentType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentAttachmentUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentAttachmentUpdateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *IncidentAttachmentUpdateAttributes `json:"attributes,omitempty"`
		Id         *string                             `json:"id,omitempty"`
		Type       *IncidentAttachmentType             `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Attributes = all.Attributes
	o.Id = all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
