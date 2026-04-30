// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestRelationships Relationships of a change request.
type ChangeRequestRelationships struct {
	// Relationship to change request decisions.
	ChangeRequestDecisions ChangeRequestDecisionsRelationship `json:"change_request_decisions"`
	// Relationship to a user.
	CreatedBy ChangeRequestUserRelationship `json:"created_by"`
	// Relationship to a user.
	ModifiedBy ChangeRequestUserRelationship `json:"modified_by"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestRelationships instantiates a new ChangeRequestRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestRelationships(changeRequestDecisions ChangeRequestDecisionsRelationship, createdBy ChangeRequestUserRelationship, modifiedBy ChangeRequestUserRelationship) *ChangeRequestRelationships {
	this := ChangeRequestRelationships{}
	this.ChangeRequestDecisions = changeRequestDecisions
	this.CreatedBy = createdBy
	this.ModifiedBy = modifiedBy
	return &this
}

// NewChangeRequestRelationshipsWithDefaults instantiates a new ChangeRequestRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestRelationshipsWithDefaults() *ChangeRequestRelationships {
	this := ChangeRequestRelationships{}
	return &this
}

// GetChangeRequestDecisions returns the ChangeRequestDecisions field value.
func (o *ChangeRequestRelationships) GetChangeRequestDecisions() ChangeRequestDecisionsRelationship {
	if o == nil {
		var ret ChangeRequestDecisionsRelationship
		return ret
	}
	return o.ChangeRequestDecisions
}

// GetChangeRequestDecisionsOk returns a tuple with the ChangeRequestDecisions field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestRelationships) GetChangeRequestDecisionsOk() (*ChangeRequestDecisionsRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeRequestDecisions, true
}

// SetChangeRequestDecisions sets field value.
func (o *ChangeRequestRelationships) SetChangeRequestDecisions(v ChangeRequestDecisionsRelationship) {
	o.ChangeRequestDecisions = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *ChangeRequestRelationships) GetCreatedBy() ChangeRequestUserRelationship {
	if o == nil {
		var ret ChangeRequestUserRelationship
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestRelationships) GetCreatedByOk() (*ChangeRequestUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *ChangeRequestRelationships) SetCreatedBy(v ChangeRequestUserRelationship) {
	o.CreatedBy = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *ChangeRequestRelationships) GetModifiedBy() ChangeRequestUserRelationship {
	if o == nil {
		var ret ChangeRequestUserRelationship
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestRelationships) GetModifiedByOk() (*ChangeRequestUserRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *ChangeRequestRelationships) SetModifiedBy(v ChangeRequestUserRelationship) {
	o.ModifiedBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["change_request_decisions"] = o.ChangeRequestDecisions
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["modified_by"] = o.ModifiedBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeRequestRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChangeRequestDecisions *ChangeRequestDecisionsRelationship `json:"change_request_decisions"`
		CreatedBy              *ChangeRequestUserRelationship      `json:"created_by"`
		ModifiedBy             *ChangeRequestUserRelationship      `json:"modified_by"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ChangeRequestDecisions == nil {
		return fmt.Errorf("required field change_request_decisions missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"change_request_decisions", "created_by", "modified_by"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ChangeRequestDecisions.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ChangeRequestDecisions = *all.ChangeRequestDecisions
	if all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = *all.CreatedBy
	if all.ModifiedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ModifiedBy = *all.ModifiedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
