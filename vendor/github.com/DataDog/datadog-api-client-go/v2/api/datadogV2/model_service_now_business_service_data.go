// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceNowBusinessServiceData Data object for a ServiceNow business service
type ServiceNowBusinessServiceData struct {
	// Attributes of a ServiceNow business service
	Attributes ServiceNowBusinessServiceAttributes `json:"attributes"`
	// Unique identifier for the ServiceNow business service
	Id uuid.UUID `json:"id"`
	// Type identifier for ServiceNow business service resources
	Type ServiceNowBusinessServiceType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceNowBusinessServiceData instantiates a new ServiceNowBusinessServiceData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceNowBusinessServiceData(attributes ServiceNowBusinessServiceAttributes, id uuid.UUID, typeVar ServiceNowBusinessServiceType) *ServiceNowBusinessServiceData {
	this := ServiceNowBusinessServiceData{}
	this.Attributes = attributes
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewServiceNowBusinessServiceDataWithDefaults instantiates a new ServiceNowBusinessServiceData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceNowBusinessServiceDataWithDefaults() *ServiceNowBusinessServiceData {
	this := ServiceNowBusinessServiceData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ServiceNowBusinessServiceData) GetAttributes() ServiceNowBusinessServiceAttributes {
	if o == nil {
		var ret ServiceNowBusinessServiceAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ServiceNowBusinessServiceData) GetAttributesOk() (*ServiceNowBusinessServiceAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ServiceNowBusinessServiceData) SetAttributes(v ServiceNowBusinessServiceAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *ServiceNowBusinessServiceData) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ServiceNowBusinessServiceData) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ServiceNowBusinessServiceData) SetId(v uuid.UUID) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *ServiceNowBusinessServiceData) GetType() ServiceNowBusinessServiceType {
	if o == nil {
		var ret ServiceNowBusinessServiceType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceNowBusinessServiceData) GetTypeOk() (*ServiceNowBusinessServiceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ServiceNowBusinessServiceData) SetType(v ServiceNowBusinessServiceType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceNowBusinessServiceData) MarshalJSON() ([]byte, error) {
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
func (o *ServiceNowBusinessServiceData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *ServiceNowBusinessServiceAttributes `json:"attributes"`
		Id         *uuid.UUID                           `json:"id"`
		Type       *ServiceNowBusinessServiceType       `json:"type"`
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
