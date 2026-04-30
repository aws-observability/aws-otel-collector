// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgGroupPolicyConfigData An org group policy config resource.
type OrgGroupPolicyConfigData struct {
	// Attributes of an org group policy config.
	Attributes OrgGroupPolicyConfigAttributes `json:"attributes"`
	// The identifier of the policy config (uses the config name).
	Id string `json:"id"`
	// Org group policy configs resource type.
	Type OrgGroupPolicyConfigType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgGroupPolicyConfigData instantiates a new OrgGroupPolicyConfigData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgGroupPolicyConfigData(attributes OrgGroupPolicyConfigAttributes, id string, typeVar OrgGroupPolicyConfigType) *OrgGroupPolicyConfigData {
	this := OrgGroupPolicyConfigData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewOrgGroupPolicyConfigDataWithDefaults instantiates a new OrgGroupPolicyConfigData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgGroupPolicyConfigDataWithDefaults() *OrgGroupPolicyConfigData {
	this := OrgGroupPolicyConfigData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *OrgGroupPolicyConfigData) GetAttributes() OrgGroupPolicyConfigAttributes {
	if o == nil {
		var ret OrgGroupPolicyConfigAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyConfigData) GetAttributesOk() (*OrgGroupPolicyConfigAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *OrgGroupPolicyConfigData) SetAttributes(v OrgGroupPolicyConfigAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *OrgGroupPolicyConfigData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyConfigData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *OrgGroupPolicyConfigData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *OrgGroupPolicyConfigData) GetType() OrgGroupPolicyConfigType {
	if o == nil {
		var ret OrgGroupPolicyConfigType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OrgGroupPolicyConfigData) GetTypeOk() (*OrgGroupPolicyConfigType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *OrgGroupPolicyConfigData) SetType(v OrgGroupPolicyConfigType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgGroupPolicyConfigData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgGroupPolicyConfigData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *OrgGroupPolicyConfigAttributes `json:"attributes"`
		Id         *string                         `json:"id"`
		Type       *OrgGroupPolicyConfigType       `json:"type"`
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
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
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
