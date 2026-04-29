// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DORADeploymentObject A DORA deployment event.
type DORADeploymentObject struct {
	// The attributes of the deployment event.
	Attributes *DORADeploymentObjectAttributes `json:"attributes,omitempty"`
	// The ID of the deployment event.
	Id *string `json:"id,omitempty"`
	// JSON:API type for DORA deployment events.
	Type *DORADeploymentType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDORADeploymentObject instantiates a new DORADeploymentObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDORADeploymentObject() *DORADeploymentObject {
	this := DORADeploymentObject{}
	var typeVar DORADeploymentType = DORADEPLOYMENTTYPE_DORA_DEPLOYMENT
	this.Type = &typeVar
	return &this
}

// NewDORADeploymentObjectWithDefaults instantiates a new DORADeploymentObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDORADeploymentObjectWithDefaults() *DORADeploymentObject {
	this := DORADeploymentObject{}
	var typeVar DORADeploymentType = DORADEPLOYMENTTYPE_DORA_DEPLOYMENT
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DORADeploymentObject) GetAttributes() DORADeploymentObjectAttributes {
	if o == nil || o.Attributes == nil {
		var ret DORADeploymentObjectAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObject) GetAttributesOk() (*DORADeploymentObjectAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DORADeploymentObject) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given DORADeploymentObjectAttributes and assigns it to the Attributes field.
func (o *DORADeploymentObject) SetAttributes(v DORADeploymentObjectAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DORADeploymentObject) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObject) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DORADeploymentObject) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DORADeploymentObject) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DORADeploymentObject) GetType() DORADeploymentType {
	if o == nil || o.Type == nil {
		var ret DORADeploymentType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DORADeploymentObject) GetTypeOk() (*DORADeploymentType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DORADeploymentObject) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given DORADeploymentType and assigns it to the Type field.
func (o *DORADeploymentObject) SetType(v DORADeploymentType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DORADeploymentObject) MarshalJSON() ([]byte, error) {
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
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DORADeploymentObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *DORADeploymentObjectAttributes `json:"attributes,omitempty"`
		Id         *string                         `json:"id,omitempty"`
		Type       *DORADeploymentType             `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
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
