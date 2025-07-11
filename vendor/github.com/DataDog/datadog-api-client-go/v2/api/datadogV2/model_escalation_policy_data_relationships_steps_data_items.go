// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationPolicyDataRelationshipsStepsDataItems Defines a relationship to a single step within an escalation policy. Contains the step's `id` and `type`.
type EscalationPolicyDataRelationshipsStepsDataItems struct {
	// Specifies the unique identifier for the step resource.
	Id string `json:"id"`
	// Indicates that the resource is of type `steps`.
	Type EscalationPolicyDataRelationshipsStepsDataItemsType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationPolicyDataRelationshipsStepsDataItems instantiates a new EscalationPolicyDataRelationshipsStepsDataItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationPolicyDataRelationshipsStepsDataItems(id string, typeVar EscalationPolicyDataRelationshipsStepsDataItemsType) *EscalationPolicyDataRelationshipsStepsDataItems {
	this := EscalationPolicyDataRelationshipsStepsDataItems{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewEscalationPolicyDataRelationshipsStepsDataItemsWithDefaults instantiates a new EscalationPolicyDataRelationshipsStepsDataItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationPolicyDataRelationshipsStepsDataItemsWithDefaults() *EscalationPolicyDataRelationshipsStepsDataItems {
	this := EscalationPolicyDataRelationshipsStepsDataItems{}
	var typeVar EscalationPolicyDataRelationshipsStepsDataItemsType = ESCALATIONPOLICYDATARELATIONSHIPSSTEPSDATAITEMSTYPE_STEPS
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *EscalationPolicyDataRelationshipsStepsDataItems) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyDataRelationshipsStepsDataItems) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *EscalationPolicyDataRelationshipsStepsDataItems) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *EscalationPolicyDataRelationshipsStepsDataItems) GetType() EscalationPolicyDataRelationshipsStepsDataItemsType {
	if o == nil {
		var ret EscalationPolicyDataRelationshipsStepsDataItemsType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EscalationPolicyDataRelationshipsStepsDataItems) GetTypeOk() (*EscalationPolicyDataRelationshipsStepsDataItemsType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *EscalationPolicyDataRelationshipsStepsDataItems) SetType(v EscalationPolicyDataRelationshipsStepsDataItemsType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationPolicyDataRelationshipsStepsDataItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationPolicyDataRelationshipsStepsDataItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string                                              `json:"id"`
		Type *EscalationPolicyDataRelationshipsStepsDataItemsType `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
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
