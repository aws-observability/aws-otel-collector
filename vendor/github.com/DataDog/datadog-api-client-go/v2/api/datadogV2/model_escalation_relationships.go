// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EscalationRelationships Contains the relationships of an escalation object, including its responders.
type EscalationRelationships struct {
	// Lists the users involved in a specific step of the escalation policy.
	Responders *EscalationRelationshipsResponders `json:"responders,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEscalationRelationships instantiates a new EscalationRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEscalationRelationships() *EscalationRelationships {
	this := EscalationRelationships{}
	return &this
}

// NewEscalationRelationshipsWithDefaults instantiates a new EscalationRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEscalationRelationshipsWithDefaults() *EscalationRelationships {
	this := EscalationRelationships{}
	return &this
}

// GetResponders returns the Responders field value if set, zero value otherwise.
func (o *EscalationRelationships) GetResponders() EscalationRelationshipsResponders {
	if o == nil || o.Responders == nil {
		var ret EscalationRelationshipsResponders
		return ret
	}
	return *o.Responders
}

// GetRespondersOk returns a tuple with the Responders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EscalationRelationships) GetRespondersOk() (*EscalationRelationshipsResponders, bool) {
	if o == nil || o.Responders == nil {
		return nil, false
	}
	return o.Responders, true
}

// HasResponders returns a boolean if a field has been set.
func (o *EscalationRelationships) HasResponders() bool {
	return o != nil && o.Responders != nil
}

// SetResponders gets a reference to the given EscalationRelationshipsResponders and assigns it to the Responders field.
func (o *EscalationRelationships) SetResponders(v EscalationRelationshipsResponders) {
	o.Responders = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EscalationRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Responders != nil {
		toSerialize["responders"] = o.Responders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EscalationRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Responders *EscalationRelationshipsResponders `json:"responders,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"responders"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Responders != nil && all.Responders.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Responders = all.Responders

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
