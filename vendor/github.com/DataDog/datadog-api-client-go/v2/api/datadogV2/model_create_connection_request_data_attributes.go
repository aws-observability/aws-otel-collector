// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateConnectionRequestDataAttributes
type CreateConnectionRequestDataAttributes struct {
	//
	Fields []CreateConnectionRequestDataAttributesFieldsItems `json:"fields,omitempty"`
	//
	JoinAttribute string `json:"join_attribute"`
	//
	JoinType string `json:"join_type"`
	//
	Metadata map[string]string `json:"metadata,omitempty"`
	//
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateConnectionRequestDataAttributes instantiates a new CreateConnectionRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateConnectionRequestDataAttributes(joinAttribute string, joinType string, typeVar string) *CreateConnectionRequestDataAttributes {
	this := CreateConnectionRequestDataAttributes{}
	this.JoinAttribute = joinAttribute
	this.JoinType = joinType
	this.Type = typeVar
	return &this
}

// NewCreateConnectionRequestDataAttributesWithDefaults instantiates a new CreateConnectionRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateConnectionRequestDataAttributesWithDefaults() *CreateConnectionRequestDataAttributes {
	this := CreateConnectionRequestDataAttributes{}
	return &this
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *CreateConnectionRequestDataAttributes) GetFields() []CreateConnectionRequestDataAttributesFieldsItems {
	if o == nil || o.Fields == nil {
		var ret []CreateConnectionRequestDataAttributesFieldsItems
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributes) GetFieldsOk() (*[]CreateConnectionRequestDataAttributesFieldsItems, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return &o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *CreateConnectionRequestDataAttributes) HasFields() bool {
	return o != nil && o.Fields != nil
}

// SetFields gets a reference to the given []CreateConnectionRequestDataAttributesFieldsItems and assigns it to the Fields field.
func (o *CreateConnectionRequestDataAttributes) SetFields(v []CreateConnectionRequestDataAttributesFieldsItems) {
	o.Fields = v
}

// GetJoinAttribute returns the JoinAttribute field value.
func (o *CreateConnectionRequestDataAttributes) GetJoinAttribute() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JoinAttribute
}

// GetJoinAttributeOk returns a tuple with the JoinAttribute field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributes) GetJoinAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinAttribute, true
}

// SetJoinAttribute sets field value.
func (o *CreateConnectionRequestDataAttributes) SetJoinAttribute(v string) {
	o.JoinAttribute = v
}

// GetJoinType returns the JoinType field value.
func (o *CreateConnectionRequestDataAttributes) GetJoinType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.JoinType
}

// GetJoinTypeOk returns a tuple with the JoinType field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributes) GetJoinTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinType, true
}

// SetJoinType sets field value.
func (o *CreateConnectionRequestDataAttributes) SetJoinType(v string) {
	o.JoinType = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateConnectionRequestDataAttributes) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributes) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateConnectionRequestDataAttributes) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateConnectionRequestDataAttributes) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetType returns the Type field value.
func (o *CreateConnectionRequestDataAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionRequestDataAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateConnectionRequestDataAttributes) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateConnectionRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	toSerialize["join_attribute"] = o.JoinAttribute
	toSerialize["join_type"] = o.JoinType
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateConnectionRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields        []CreateConnectionRequestDataAttributesFieldsItems `json:"fields,omitempty"`
		JoinAttribute *string                                            `json:"join_attribute"`
		JoinType      *string                                            `json:"join_type"`
		Metadata      map[string]string                                  `json:"metadata,omitempty"`
		Type          *string                                            `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.JoinAttribute == nil {
		return fmt.Errorf("required field join_attribute missing")
	}
	if all.JoinType == nil {
		return fmt.Errorf("required field join_type missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields", "join_attribute", "join_type", "metadata", "type"})
	} else {
		return err
	}
	o.Fields = all.Fields
	o.JoinAttribute = *all.JoinAttribute
	o.JoinType = *all.JoinType
	o.Metadata = all.Metadata
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
