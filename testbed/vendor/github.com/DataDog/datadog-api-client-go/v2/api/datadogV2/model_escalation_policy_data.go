// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyData Represents the data for a single escalation policy, including its attributes, ID, relationships, and resource type.
type EscalationPolicyData struct {
	// Defines the main attributes of an escalation policy, such as its name and behavior on policy end.
	Attributes *EscalationPolicyDataAttributes `json:"attributes,omitempty"`
	// Specifies the unique identifier of the escalation policy.
	Id *string `json:"id,omitempty"`
	// Represents the relationships for an escalation policy, including references to steps and teams.
	Relationships *EscalationPolicyDataRelationships `json:"relationships,omitempty"`
	// Indicates that the resource is of type `policies`.
	Type EscalationPolicyDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyData instantiates a new EscalationPolicyData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyData(typeVar EscalationPolicyDataType) *EscalationPolicyData {
	this := EscalationPolicyData{}
	this.Type = typeVar
	return &this
}

// NewEscalationPolicyDataWithDefaults instantiates a new EscalationPolicyData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyDataWithDefaults() *EscalationPolicyData {
	this := EscalationPolicyData{}
	var typeVar EscalationPolicyDataType = ESCALATIONPOLICYDATATYPE_POLICIES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EscalationPolicyData) GetAttributes() EscalationPolicyDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret EscalationPolicyDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyData) GetAttributesOk() (*EscalationPolicyDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EscalationPolicyData) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given EscalationPolicyDataAttributes and assigns it to the Attributes field.
func (o *EscalationPolicyData) SetAttributes(v EscalationPolicyDataAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EscalationPolicyData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EscalationPolicyData) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EscalationPolicyData) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *EscalationPolicyData) GetRelationships() EscalationPolicyDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret EscalationPolicyDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyData) GetRelationshipsOk() (*EscalationPolicyDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *EscalationPolicyData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given EscalationPolicyDataRelationships and assigns it to the Relationships field.
func (o *EscalationPolicyData) SetRelationships(v EscalationPolicyDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *EscalationPolicyData) GetType() EscalationPolicyDataType {
	if o == nil {
		var ret EscalationPolicyDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyData) GetTypeOk() (*EscalationPolicyDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EscalationPolicyData) SetType(v EscalationPolicyDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyData) MarshalJSON() ([]byte, error) {
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
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicyData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *EscalationPolicyDataAttributes    `json:"attributes,omitempty"`
		Id            *string                            `json:"id,omitempty"`
		Relationships *EscalationPolicyDataRelationships `json:"relationships,omitempty"`
		Type          *EscalationPolicyDataType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Relationships != nil && all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = all.Relationships
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
