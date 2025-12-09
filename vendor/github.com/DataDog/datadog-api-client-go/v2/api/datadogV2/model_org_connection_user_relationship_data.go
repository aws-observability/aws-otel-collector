// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgConnectionUserRelationshipData The data for a user relationship.
type OrgConnectionUserRelationshipData struct {
	// User UUID.
	Id *string `json:"id,omitempty"`
	// User name.
	Name *string `json:"name,omitempty"`
	// The type of the user relationship.
	Type *OrgConnectionUserRelationshipDataType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgConnectionUserRelationshipData instantiates a new OrgConnectionUserRelationshipData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgConnectionUserRelationshipData() *OrgConnectionUserRelationshipData {
	this := OrgConnectionUserRelationshipData{}
	return &this
}

// NewOrgConnectionUserRelationshipDataWithDefaults instantiates a new OrgConnectionUserRelationshipData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgConnectionUserRelationshipDataWithDefaults() *OrgConnectionUserRelationshipData {
	this := OrgConnectionUserRelationshipData{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrgConnectionUserRelationshipData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgConnectionUserRelationshipData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrgConnectionUserRelationshipData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrgConnectionUserRelationshipData) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *OrgConnectionUserRelationshipData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgConnectionUserRelationshipData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *OrgConnectionUserRelationshipData) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *OrgConnectionUserRelationshipData) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrgConnectionUserRelationshipData) GetType() OrgConnectionUserRelationshipDataType {
	if o == nil || o.Type == nil {
		var ret OrgConnectionUserRelationshipDataType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgConnectionUserRelationshipData) GetTypeOk() (*OrgConnectionUserRelationshipDataType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrgConnectionUserRelationshipData) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given OrgConnectionUserRelationshipDataType and assigns it to the Type field.
func (o *OrgConnectionUserRelationshipData) SetType(v OrgConnectionUserRelationshipDataType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgConnectionUserRelationshipData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
func (o *OrgConnectionUserRelationshipData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                                `json:"id,omitempty"`
		Name *string                                `json:"name,omitempty"`
		Type *OrgConnectionUserRelationshipDataType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "name", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = all.Id
	o.Name = all.Name
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
