// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyCreateRequestData Represents the data for creating an escalation policy, including its attributes, relationships, and resource type.
type EscalationPolicyCreateRequestData struct {
	// Defines the attributes for creating an escalation policy, including its description, name, resolution behavior, retries, and steps.
	Attributes EscalationPolicyCreateRequestDataAttributes `json:"attributes"`
	// Represents relationships in an escalation policy creation request, including references to teams.
	Relationships *EscalationPolicyCreateRequestDataRelationships `json:"relationships,omitempty"`
	// Indicates that the resource is of type `policies`.
	Type EscalationPolicyCreateRequestDataType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyCreateRequestData instantiates a new EscalationPolicyCreateRequestData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyCreateRequestData(attributes EscalationPolicyCreateRequestDataAttributes, typeVar EscalationPolicyCreateRequestDataType) *EscalationPolicyCreateRequestData {
	this := EscalationPolicyCreateRequestData{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewEscalationPolicyCreateRequestDataWithDefaults instantiates a new EscalationPolicyCreateRequestData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyCreateRequestDataWithDefaults() *EscalationPolicyCreateRequestData {
	this := EscalationPolicyCreateRequestData{}
	var typeVar EscalationPolicyCreateRequestDataType = ESCALATIONPOLICYCREATEREQUESTDATATYPE_POLICIES
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *EscalationPolicyCreateRequestData) GetAttributes() EscalationPolicyCreateRequestDataAttributes {
	if o == nil {
		var ret EscalationPolicyCreateRequestDataAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyCreateRequestData) GetAttributesOk() (*EscalationPolicyCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *EscalationPolicyCreateRequestData) SetAttributes(v EscalationPolicyCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *EscalationPolicyCreateRequestData) GetRelationships() EscalationPolicyCreateRequestDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret EscalationPolicyCreateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationPolicyCreateRequestData) GetRelationshipsOk() (*EscalationPolicyCreateRequestDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *EscalationPolicyCreateRequestData) HasRelationships() bool {
	return o != nil && o.Relationships != nil
}

// SetRelationships gets a reference to the given EscalationPolicyCreateRequestDataRelationships and assigns it to the Relationships field.
func (o *EscalationPolicyCreateRequestData) SetRelationships(v EscalationPolicyCreateRequestDataRelationships) {
	o.Relationships = &v
}

// GetType returns the Type field value.
func (o *EscalationPolicyCreateRequestData) GetType() EscalationPolicyCreateRequestDataType {
	if o == nil {
		var ret EscalationPolicyCreateRequestDataType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyCreateRequestData) GetTypeOk() (*EscalationPolicyCreateRequestDataType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EscalationPolicyCreateRequestData) SetType(v EscalationPolicyCreateRequestDataType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
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
func (o *EscalationPolicyCreateRequestData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *EscalationPolicyCreateRequestDataAttributes    `json:"attributes"`
		Relationships *EscalationPolicyCreateRequestDataRelationships `json:"relationships,omitempty"`
		Type          *EscalationPolicyCreateRequestDataType          `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
