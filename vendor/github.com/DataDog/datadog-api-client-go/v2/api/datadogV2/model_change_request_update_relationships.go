// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestUpdateRelationships Relationships for updating a change request.
type ChangeRequestUpdateRelationships struct {
	// Relationship to change request decisions.
	ChangeRequestDecisions *ChangeRequestDecisionsRelationship `json:"change_request_decisions,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestUpdateRelationships instantiates a new ChangeRequestUpdateRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestUpdateRelationships() *ChangeRequestUpdateRelationships {
	this := ChangeRequestUpdateRelationships{}
	return &this
}

// NewChangeRequestUpdateRelationshipsWithDefaults instantiates a new ChangeRequestUpdateRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestUpdateRelationshipsWithDefaults() *ChangeRequestUpdateRelationships {
	this := ChangeRequestUpdateRelationships{}
	return &this
}

// GetChangeRequestDecisions returns the ChangeRequestDecisions field value if set, zero value otherwise.
func (o *ChangeRequestUpdateRelationships) GetChangeRequestDecisions() ChangeRequestDecisionsRelationship {
	if o == nil || o.ChangeRequestDecisions == nil {
		var ret ChangeRequestDecisionsRelationship
		return ret
	}
	return *o.ChangeRequestDecisions
}

// GetChangeRequestDecisionsOk returns a tuple with the ChangeRequestDecisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeRequestUpdateRelationships) GetChangeRequestDecisionsOk() (*ChangeRequestDecisionsRelationship, bool) {
	if o == nil || o.ChangeRequestDecisions == nil {
		return nil, false
	}
	return o.ChangeRequestDecisions, true
}

// HasChangeRequestDecisions returns a boolean if a field has been set.
func (o *ChangeRequestUpdateRelationships) HasChangeRequestDecisions() bool {
	return o != nil && o.ChangeRequestDecisions != nil
}

// SetChangeRequestDecisions gets a reference to the given ChangeRequestDecisionsRelationship and assigns it to the ChangeRequestDecisions field.
func (o *ChangeRequestUpdateRelationships) SetChangeRequestDecisions(v ChangeRequestDecisionsRelationship) {
	o.ChangeRequestDecisions = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestUpdateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChangeRequestDecisions != nil {
		toSerialize["change_request_decisions"] = o.ChangeRequestDecisions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeRequestUpdateRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChangeRequestDecisions *ChangeRequestDecisionsRelationship `json:"change_request_decisions,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"change_request_decisions"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ChangeRequestDecisions != nil && all.ChangeRequestDecisions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ChangeRequestDecisions = all.ChangeRequestDecisions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
