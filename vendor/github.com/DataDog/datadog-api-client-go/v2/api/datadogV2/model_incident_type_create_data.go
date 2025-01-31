// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTypeCreateData Incident type data for a create request.
type IncidentTypeCreateData struct {
	// Incident type's attributes.
	Attributes IncidentTypeAttributes `json:"attributes"`
	// Incident type resource type.
	Type IncidentTypeType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTypeCreateData instantiates a new IncidentTypeCreateData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTypeCreateData(attributes IncidentTypeAttributes, typeVar IncidentTypeType) *IncidentTypeCreateData {
	this := IncidentTypeCreateData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewIncidentTypeCreateDataWithDefaults instantiates a new IncidentTypeCreateData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTypeCreateDataWithDefaults() *IncidentTypeCreateData {
	this := IncidentTypeCreateData{}
	var typeVar IncidentTypeType = INCIDENTTYPETYPE_INCIDENT_TYPES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *IncidentTypeCreateData) GetAttributes() IncidentTypeAttributes {
	if o == nil {
		var ret IncidentTypeAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *IncidentTypeCreateData) GetAttributesOk() (*IncidentTypeAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *IncidentTypeCreateData) SetAttributes(v IncidentTypeAttributes) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *IncidentTypeCreateData) GetType() IncidentTypeType {
	if o == nil {
		var ret IncidentTypeType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IncidentTypeCreateData) GetTypeOk() (*IncidentTypeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *IncidentTypeCreateData) SetType(v IncidentTypeType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTypeCreateData) MarshalJSON() ([]byte, error) {
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
func (o *IncidentTypeCreateData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *IncidentTypeAttributes `json:"attributes"`
		Type       *IncidentTypeType       `json:"type"`
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
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
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
