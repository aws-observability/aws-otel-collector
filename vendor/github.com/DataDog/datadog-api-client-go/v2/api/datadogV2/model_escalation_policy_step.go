// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyStep Represents a single step in an escalation policy, including its attributes, relationships, and resource type.
type EscalationPolicyStep struct {
	// Defines attributes for an escalation policy step, such as assignment strategy and escalation timeout.
	Attributes *EscalationPolicyStepAttributes `json:"attributes,omitempty"`
	// Specifies the unique identifier of this escalation policy step.
	Id *string `json:"id,omitempty"`
	// Represents the relationship of an escalation policy step to its targets.
	Relationships *EscalationPolicyStepRelationships `json:"relationships,omitempty"`
	// Indicates that the resource is of type `steps`.
	Type EscalationPolicyStepType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyStep instantiates a new EscalationPolicyStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyStep(typeVar EscalationPolicyStepType) *EscalationPolicyStep {
	this := EscalationPolicyStep{}
	this.Type = typeVar
	return &this
}

// NewEscalationPolicyStepWithDefaults instantiates a new EscalationPolicyStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyStepWithDefaults() *EscalationPolicyStep {
	this := EscalationPolicyStep{}
	var typeVar EscalationPolicyStepType = ESCALATIONPOLICYSTEPTYPE_STEPS
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *EscalationPolicyStep) GetAttributes() EscalationPolicyStepAttributes {
	if o == nil || o.Attributes == nil {
		var ret EscalationPolicyStepAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyStep) GetAttributesOk() (*EscalationPolicyStepAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *EscalationPolicyStep) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given EscalationPolicyStepAttributes and assigns it to the Attributes field.
func (o *EscalationPolicyStep) SetAttributes(v EscalationPolicyStepAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EscalationPolicyStep) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyStep) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EscalationPolicyStep) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EscalationPolicyStep) SetId(v string) {
	o.Id = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *EscalationPolicyStep) GetRelationships() EscalationPolicyStepRelationships {
	if o == nil || o.Relationships == nil {
		var ret EscalationPolicyStepRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyStep) GetRelationshipsOk() (*EscalationPolicyStepRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *EscalationPolicyStep) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given EscalationPolicyStepRelationships and assigns it to the Relationships field.
func (o *EscalationPolicyStep) SetRelationships(v EscalationPolicyStepRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *EscalationPolicyStep) GetType() EscalationPolicyStepType {
	if o == nil {
		var ret EscalationPolicyStepType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyStep) GetTypeOk() (*EscalationPolicyStepType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EscalationPolicyStep) SetType(v EscalationPolicyStepType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyStep) MarshalJSON() ([]byte, error) {
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
func (o *EscalationPolicyStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *EscalationPolicyStepAttributes    `json:"attributes,omitempty"`
		Id            *string                            `json:"id,omitempty"`
		Relationships *EscalationPolicyStepRelationships `json:"relationships,omitempty"`
		Type          *EscalationPolicyStepType          `json:"type"`
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
