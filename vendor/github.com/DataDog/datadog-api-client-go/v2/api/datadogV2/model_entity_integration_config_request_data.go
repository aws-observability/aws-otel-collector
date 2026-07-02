// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityIntegrationConfigRequestData JSON:API resource object used in a request to create or update an entity integration configuration.
type EntityIntegrationConfigRequestData struct {
	// Attributes used to create or update an entity integration configuration.
	Attributes EntityIntegrationConfigRequestAttributes `json:"attributes"`
	// JSON:API resource type for the entity integration configuration create or update request. Always `entity_integration_config_requests`.
	Type EntityIntegrationConfigRequestType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityIntegrationConfigRequestData instantiates a new EntityIntegrationConfigRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityIntegrationConfigRequestData(attributes EntityIntegrationConfigRequestAttributes, typeVar EntityIntegrationConfigRequestType) *EntityIntegrationConfigRequestData {
	this := EntityIntegrationConfigRequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewEntityIntegrationConfigRequestDataWithDefaults instantiates a new EntityIntegrationConfigRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityIntegrationConfigRequestDataWithDefaults() *EntityIntegrationConfigRequestData {
	this := EntityIntegrationConfigRequestData{}
	var typeVar EntityIntegrationConfigRequestType = ENTITYINTEGRATIONCONFIGREQUESTTYPE_ENTITY_INTEGRATION_CONFIG_REQUESTS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *EntityIntegrationConfigRequestData) GetAttributes() EntityIntegrationConfigRequestAttributes {
	if o == nil {
		var ret EntityIntegrationConfigRequestAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *EntityIntegrationConfigRequestData) GetAttributesOk() (*EntityIntegrationConfigRequestAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *EntityIntegrationConfigRequestData) SetAttributes(v EntityIntegrationConfigRequestAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *EntityIntegrationConfigRequestData) GetType() EntityIntegrationConfigRequestType {
	if o == nil {
		var ret EntityIntegrationConfigRequestType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EntityIntegrationConfigRequestData) GetTypeOk() (*EntityIntegrationConfigRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EntityIntegrationConfigRequestData) SetType(v EntityIntegrationConfigRequestType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityIntegrationConfigRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityIntegrationConfigRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *EntityIntegrationConfigRequestAttributes `json:"attributes"`
		Type       *EntityIntegrationConfigRequestType       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
