// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GoogleChatOrganizationData Google Chat organization data from a response.
type GoogleChatOrganizationData struct {
	// Google Chat organization attributes.
	Attributes *GoogleChatOrganizationAttributes `json:"attributes,omitempty"`
	// The ID of the Google Chat organization binding.
	Id *string `json:"id,omitempty"`
	// Google Chat organization relationships.
	Relationships *GoogleChatOrganizationRelationships `json:"relationships,omitempty"`
	// Google Chat organization resource type.
	Type *GoogleChatOrganizationType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGoogleChatOrganizationData instantiates a new GoogleChatOrganizationData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGoogleChatOrganizationData() *GoogleChatOrganizationData {
	this := GoogleChatOrganizationData{}
	var typeVar GoogleChatOrganizationType = GOOGLECHATORGANIZATIONTYPE_GOOGLE_CHAT_ORGANIZATION_TYPE
	this.Type = &typeVar
	return &this
}

// NewGoogleChatOrganizationDataWithDefaults instantiates a new GoogleChatOrganizationData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGoogleChatOrganizationDataWithDefaults() *GoogleChatOrganizationData {
	this := GoogleChatOrganizationData{}
	var typeVar GoogleChatOrganizationType = GOOGLECHATORGANIZATIONTYPE_GOOGLE_CHAT_ORGANIZATION_TYPE
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GoogleChatOrganizationData) GetAttributes() GoogleChatOrganizationAttributes {
	if o == nil || o.Attributes == nil {
		var ret GoogleChatOrganizationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatOrganizationData) GetAttributesOk() (*GoogleChatOrganizationAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GoogleChatOrganizationData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given GoogleChatOrganizationAttributes and assigns it to the Attributes field.
func (o *GoogleChatOrganizationData) SetAttributes(v GoogleChatOrganizationAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GoogleChatOrganizationData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatOrganizationData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GoogleChatOrganizationData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GoogleChatOrganizationData) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GoogleChatOrganizationData) GetRelationships() GoogleChatOrganizationRelationships {
	if o == nil || o.Relationships == nil {
		var ret GoogleChatOrganizationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatOrganizationData) GetRelationshipsOk() (*GoogleChatOrganizationRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GoogleChatOrganizationData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given GoogleChatOrganizationRelationships and assigns it to the Relationships field.
func (o *GoogleChatOrganizationData) SetRelationships(v GoogleChatOrganizationRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GoogleChatOrganizationData) GetType() GoogleChatOrganizationType {
	if o == nil || o.Type == nil {
		var ret GoogleChatOrganizationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleChatOrganizationData) GetTypeOk() (*GoogleChatOrganizationType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GoogleChatOrganizationData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given GoogleChatOrganizationType and assigns it to the Type field.
func (o *GoogleChatOrganizationData) SetType(v GoogleChatOrganizationType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GoogleChatOrganizationData) MarshalJSON() ([]byte, error) {
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
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GoogleChatOrganizationData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *GoogleChatOrganizationAttributes    `json:"attributes,omitempty"`
		Id            *string                              `json:"id,omitempty"`
		Relationships *GoogleChatOrganizationRelationships `json:"relationships,omitempty"`
		Type          *GoogleChatOrganizationType          `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
