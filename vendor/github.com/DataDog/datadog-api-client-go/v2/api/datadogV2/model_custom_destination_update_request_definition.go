// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationUpdateRequestDefinition The definition of a custom destination.
type CustomDestinationUpdateRequestDefinition struct {
	// The attributes associated with the custom destination.
	Attributes *CustomDestinationUpdateRequestAttributes `json:"attributes,omitempty"`
	// The custom destination ID.
	Id string `json:"id"`
	// The type of the resource. The value should always be `custom_destination`.
	Type CustomDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationUpdateRequestDefinition instantiates a new CustomDestinationUpdateRequestDefinition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationUpdateRequestDefinition(id string, typeVar CustomDestinationType) *CustomDestinationUpdateRequestDefinition {
	this := CustomDestinationUpdateRequestDefinition{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewCustomDestinationUpdateRequestDefinitionWithDefaults instantiates a new CustomDestinationUpdateRequestDefinition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationUpdateRequestDefinitionWithDefaults() *CustomDestinationUpdateRequestDefinition {
	this := CustomDestinationUpdateRequestDefinition{}
	var typeVar CustomDestinationType = CUSTOMDESTINATIONTYPE_CUSTOM_DESTINATION
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomDestinationUpdateRequestDefinition) GetAttributes() CustomDestinationUpdateRequestAttributes {
	if o == nil || o.Attributes == nil {
		var ret CustomDestinationUpdateRequestAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDestinationUpdateRequestDefinition) GetAttributesOk() (*CustomDestinationUpdateRequestAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomDestinationUpdateRequestDefinition) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given CustomDestinationUpdateRequestAttributes and assigns it to the Attributes field.
func (o *CustomDestinationUpdateRequestDefinition) SetAttributes(v CustomDestinationUpdateRequestAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value.
func (o *CustomDestinationUpdateRequestDefinition) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationUpdateRequestDefinition) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CustomDestinationUpdateRequestDefinition) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *CustomDestinationUpdateRequestDefinition) GetType() CustomDestinationType {
	if o == nil {
		var ret CustomDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationUpdateRequestDefinition) GetTypeOk() (*CustomDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationUpdateRequestDefinition) SetType(v CustomDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationUpdateRequestDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationUpdateRequestDefinition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *CustomDestinationUpdateRequestAttributes `json:"attributes,omitempty"`
		Id         *string                                   `json:"id"`
		Type       *CustomDestinationType                    `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
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
