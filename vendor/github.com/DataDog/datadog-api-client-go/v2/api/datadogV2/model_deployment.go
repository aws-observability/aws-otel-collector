// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Deployment The version of the app that was published.
type Deployment struct {
	// The attributes object containing the version ID of the published app.
	Attributes *DeploymentAttributes `json:"attributes,omitempty"`
	// The deployment ID.
	Id *uuid.UUID `json:"id,omitempty"`
	// Metadata object containing the publication creation information.
	Meta *DeploymentMetadata `json:"meta,omitempty"`
	// The deployment type.
	Type *AppDeploymentType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeployment instantiates a new Deployment object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeployment() *Deployment {
	this := Deployment{}
	var typeVar AppDeploymentType = APPDEPLOYMENTTYPE_DEPLOYMENT
	this.Type = &typeVar
	return &this
}

// NewDeploymentWithDefaults instantiates a new Deployment object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentWithDefaults() *Deployment {
	this := Deployment{}
	var typeVar AppDeploymentType = APPDEPLOYMENTTYPE_DEPLOYMENT
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Deployment) GetAttributes() DeploymentAttributes {
	if o == nil || o.Attributes == nil {
		var ret DeploymentAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetAttributesOk() (*DeploymentAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Deployment) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given DeploymentAttributes and assigns it to the Attributes field.
func (o *Deployment) SetAttributes(v DeploymentAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Deployment) GetId() uuid.UUID {
	if o == nil || o.Id == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetIdOk() (*uuid.UUID, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Deployment) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given uuid.UUID and assigns it to the Id field.
func (o *Deployment) SetId(v uuid.UUID) {
	o.Id = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Deployment) GetMeta() DeploymentMetadata {
	if o == nil || o.Meta == nil {
		var ret DeploymentMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetMetaOk() (*DeploymentMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Deployment) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given DeploymentMetadata and assigns it to the Meta field.
func (o *Deployment) SetMeta(v DeploymentMetadata) {
	o.Meta = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Deployment) GetType() AppDeploymentType {
	if o == nil || o.Type == nil {
		var ret AppDeploymentType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetTypeOk() (*AppDeploymentType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Deployment) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given AppDeploymentType and assigns it to the Type field.
func (o *Deployment) SetType(v AppDeploymentType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Deployment) MarshalJSON() ([]byte, error) {
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
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
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
func (o *Deployment) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *DeploymentAttributes `json:"attributes,omitempty"`
		Id         *uuid.UUID            `json:"id,omitempty"`
		Meta       *DeploymentMetadata   `json:"meta,omitempty"`
		Type       *AppDeploymentType    `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "meta", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta
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
