// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestRestrictionPolicyBinding Objects describing the binding used for a mobile test.
type SyntheticsTestRestrictionPolicyBinding struct {
	// List of principals for a mobile test binding.
	Principals []string `json:"principals,omitempty"`
	// The type of relation for the binding.
	Relation *SyntheticsTestRestrictionPolicyBindingRelation `json:"relation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestRestrictionPolicyBinding instantiates a new SyntheticsTestRestrictionPolicyBinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestRestrictionPolicyBinding() *SyntheticsTestRestrictionPolicyBinding {
	this := SyntheticsTestRestrictionPolicyBinding{}
	return &this
}

// NewSyntheticsTestRestrictionPolicyBindingWithDefaults instantiates a new SyntheticsTestRestrictionPolicyBinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestRestrictionPolicyBindingWithDefaults() *SyntheticsTestRestrictionPolicyBinding {
	this := SyntheticsTestRestrictionPolicyBinding{}
	return &this
}

// GetPrincipals returns the Principals field value if set, zero value otherwise.
func (o *SyntheticsTestRestrictionPolicyBinding) GetPrincipals() []string {
	if o == nil || o.Principals == nil {
		var ret []string
		return ret
	}
	return o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRestrictionPolicyBinding) GetPrincipalsOk() (*[]string, bool) {
	if o == nil || o.Principals == nil {
		return nil, false
	}
	return &o.Principals, true
}

// HasPrincipals returns a boolean if a field has been set.
func (o *SyntheticsTestRestrictionPolicyBinding) HasPrincipals() bool {
	return o != nil && o.Principals != nil
}

// SetPrincipals gets a reference to the given []string and assigns it to the Principals field.
func (o *SyntheticsTestRestrictionPolicyBinding) SetPrincipals(v []string) {
	o.Principals = v
}

// GetRelation returns the Relation field value if set, zero value otherwise.
func (o *SyntheticsTestRestrictionPolicyBinding) GetRelation() SyntheticsTestRestrictionPolicyBindingRelation {
	if o == nil || o.Relation == nil {
		var ret SyntheticsTestRestrictionPolicyBindingRelation
		return ret
	}
	return *o.Relation
}

// GetRelationOk returns a tuple with the Relation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRestrictionPolicyBinding) GetRelationOk() (*SyntheticsTestRestrictionPolicyBindingRelation, bool) {
	if o == nil || o.Relation == nil {
		return nil, false
	}
	return o.Relation, true
}

// HasRelation returns a boolean if a field has been set.
func (o *SyntheticsTestRestrictionPolicyBinding) HasRelation() bool {
	return o != nil && o.Relation != nil
}

// SetRelation gets a reference to the given SyntheticsTestRestrictionPolicyBindingRelation and assigns it to the Relation field.
func (o *SyntheticsTestRestrictionPolicyBinding) SetRelation(v SyntheticsTestRestrictionPolicyBindingRelation) {
	o.Relation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestRestrictionPolicyBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Principals != nil {
		toSerialize["principals"] = o.Principals
	}
	if o.Relation != nil {
		toSerialize["relation"] = o.Relation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestRestrictionPolicyBinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Principals []string                                        `json:"principals,omitempty"`
		Relation   *SyntheticsTestRestrictionPolicyBindingRelation `json:"relation,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"principals", "relation"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Principals = all.Principals
	if all.Relation != nil && !all.Relation.IsValid() {
		hasInvalidField = true
	} else {
		o.Relation = all.Relation
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
