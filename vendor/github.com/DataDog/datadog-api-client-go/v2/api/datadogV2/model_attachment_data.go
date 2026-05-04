// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AttachmentData Attachment data from a response.
type AttachmentData struct {
	// The attachment's attributes.
	Attributes AttachmentDataAttributes `json:"attributes"`
	// The unique identifier of the attachment.
	Id string `json:"id"`
	// The attachment's resource relationships.
	Relationships AttachmentDataRelationships `json:"relationships"`
	// The incident attachment resource type.
	Type IncidentAttachmentType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAttachmentData instantiates a new AttachmentData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAttachmentData(attributes AttachmentDataAttributes, id string, relationships AttachmentDataRelationships, typeVar IncidentAttachmentType) *AttachmentData {
	this := AttachmentData{}
	this.Attributes = attributes
	this.Id = id
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewAttachmentDataWithDefaults instantiates a new AttachmentData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAttachmentDataWithDefaults() *AttachmentData {
	this := AttachmentData{}
	var typeVar IncidentAttachmentType = INCIDENTATTACHMENTTYPE_INCIDENT_ATTACHMENTS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *AttachmentData) GetAttributes() AttachmentDataAttributes {
	if o == nil {
		var ret AttachmentDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AttachmentData) GetAttributesOk() (*AttachmentDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *AttachmentData) SetAttributes(v AttachmentDataAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *AttachmentData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttachmentData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *AttachmentData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value.
func (o *AttachmentData) GetRelationships() AttachmentDataRelationships {
	if o == nil {
		var ret AttachmentDataRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AttachmentData) GetRelationshipsOk() (*AttachmentDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *AttachmentData) SetRelationships(v AttachmentDataRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *AttachmentData) GetType() IncidentAttachmentType {
	if o == nil {
		var ret IncidentAttachmentType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AttachmentData) GetTypeOk() (*IncidentAttachmentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AttachmentData) SetType(v IncidentAttachmentType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AttachmentData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AttachmentData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *AttachmentDataAttributes    `json:"attributes"`
		Id            *string                      `json:"id"`
		Relationships *AttachmentDataRelationships `json:"relationships"`
		Type          *IncidentAttachmentType      `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = *all.Relationships
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
