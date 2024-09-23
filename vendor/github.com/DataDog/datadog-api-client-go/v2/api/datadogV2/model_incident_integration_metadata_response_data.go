// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentIntegrationMetadataResponseData Incident integration metadata from a response.
type IncidentIntegrationMetadataResponseData struct {
	// Incident integration metadata's attributes for a create request.
	Attributes *IncidentIntegrationMetadataAttributes `json:"attributes,omitempty"`
	// The incident integration metadata's ID.
	Id string `json:"id"`
	// The incident's integration relationships from a response.
	Relationships *IncidentIntegrationRelationships `json:"relationships,omitempty"`
	// Integration metadata resource type.
	Type IncidentIntegrationMetadataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentIntegrationMetadataResponseData instantiates a new IncidentIntegrationMetadataResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentIntegrationMetadataResponseData(id string, typeVar IncidentIntegrationMetadataType) *IncidentIntegrationMetadataResponseData {
	this := IncidentIntegrationMetadataResponseData{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewIncidentIntegrationMetadataResponseDataWithDefaults instantiates a new IncidentIntegrationMetadataResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentIntegrationMetadataResponseDataWithDefaults() *IncidentIntegrationMetadataResponseData {
	this := IncidentIntegrationMetadataResponseData{}
	var typeVar IncidentIntegrationMetadataType = INCIDENTINTEGRATIONMETADATATYPE_INCIDENT_INTEGRATIONS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *IncidentIntegrationMetadataResponseData) GetAttributes() IncidentIntegrationMetadataAttributes {
	if o == nil || o.Attributes == nil {
		var ret IncidentIntegrationMetadataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentIntegrationMetadataResponseData) GetAttributesOk() (*IncidentIntegrationMetadataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *IncidentIntegrationMetadataResponseData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given IncidentIntegrationMetadataAttributes and assigns it to the Attributes field.
func (o *IncidentIntegrationMetadataResponseData) SetAttributes(v IncidentIntegrationMetadataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *IncidentIntegrationMetadataResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IncidentIntegrationMetadataResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *IncidentIntegrationMetadataResponseData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *IncidentIntegrationMetadataResponseData) GetRelationships() IncidentIntegrationRelationships {
	if o == nil || o.Relationships == nil {
		var ret IncidentIntegrationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentIntegrationMetadataResponseData) GetRelationshipsOk() (*IncidentIntegrationRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *IncidentIntegrationMetadataResponseData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given IncidentIntegrationRelationships and assigns it to the Relationships field.
func (o *IncidentIntegrationMetadataResponseData) SetRelationships(v IncidentIntegrationRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *IncidentIntegrationMetadataResponseData) GetType() IncidentIntegrationMetadataType {
	if o == nil {
		var ret IncidentIntegrationMetadataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentIntegrationMetadataResponseData) GetTypeOk() (*IncidentIntegrationMetadataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentIntegrationMetadataResponseData) SetType(v IncidentIntegrationMetadataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentIntegrationMetadataResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentIntegrationMetadataResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *IncidentIntegrationMetadataAttributes `json:"attributes,omitempty"`
		Id            *string                                `json:"id"`
		Relationships *IncidentIntegrationRelationships      `json:"relationships,omitempty"`
		Type          *IncidentIntegrationMetadataType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
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
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = *all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
