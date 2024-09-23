// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceDefinitionV2Dot1MSTeams Service owner's Microsoft Teams.
type ServiceDefinitionV2Dot1MSTeams struct {
	// Contact value.
	Contact string `json:"contact"`
	// Contact Microsoft Teams.
	Name *string `json:"name,omitempty"`
	// Contact type.
	Type ServiceDefinitionV2Dot1MSTeamsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceDefinitionV2Dot1MSTeams instantiates a new ServiceDefinitionV2Dot1MSTeams object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceDefinitionV2Dot1MSTeams(contact string, typeVar ServiceDefinitionV2Dot1MSTeamsType) *ServiceDefinitionV2Dot1MSTeams {
	this := ServiceDefinitionV2Dot1MSTeams{}
	this.Contact = contact
	this.Type = typeVar
	return &this
}

// NewServiceDefinitionV2Dot1MSTeamsWithDefaults instantiates a new ServiceDefinitionV2Dot1MSTeams object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceDefinitionV2Dot1MSTeamsWithDefaults() *ServiceDefinitionV2Dot1MSTeams {
	this := ServiceDefinitionV2Dot1MSTeams{}
	return &this
}

// GetContact returns the Contact field value.
func (o *ServiceDefinitionV2Dot1MSTeams) GetContact() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1MSTeams) GetContactOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contact, true
}

// SetContact sets field value.
func (o *ServiceDefinitionV2Dot1MSTeams) SetContact(v string) {
	o.Contact = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceDefinitionV2Dot1MSTeams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1MSTeams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceDefinitionV2Dot1MSTeams) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceDefinitionV2Dot1MSTeams) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value.
func (o *ServiceDefinitionV2Dot1MSTeams) GetType() ServiceDefinitionV2Dot1MSTeamsType {
	if o == nil {
		var ret ServiceDefinitionV2Dot1MSTeamsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceDefinitionV2Dot1MSTeams) GetTypeOk() (*ServiceDefinitionV2Dot1MSTeamsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ServiceDefinitionV2Dot1MSTeams) SetType(v ServiceDefinitionV2Dot1MSTeamsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceDefinitionV2Dot1MSTeams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["contact"] = o.Contact
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceDefinitionV2Dot1MSTeams) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Contact *string                             `json:"contact"`
		Name    *string                             `json:"name,omitempty"`
		Type    *ServiceDefinitionV2Dot1MSTeamsType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Contact == nil {
		return fmt.Errorf("required field contact missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"contact", "name", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Contact = *all.Contact
	o.Name = all.Name
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
