// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RestrictionPolicyBinding Specifies which principals are associated with a relation.
type RestrictionPolicyBinding struct {
	// An array of principals. A principal is a subject or group of subjects.
	// Each principal is formatted as `type:id`. Supported types: `role`, `team`, `user`, and `org`.
	// The org ID can be obtained through the api/v2/current_user API.
	// The user principal type accepts service account IDs.
	Principals []string `json:"principals"`
	// The role/level of access.
	Relation string `json:"relation"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRestrictionPolicyBinding instantiates a new RestrictionPolicyBinding object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRestrictionPolicyBinding(principals []string, relation string) *RestrictionPolicyBinding {
	this := RestrictionPolicyBinding{}
	this.Principals = principals
	this.Relation = relation
	return &this
}

// NewRestrictionPolicyBindingWithDefaults instantiates a new RestrictionPolicyBinding object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRestrictionPolicyBindingWithDefaults() *RestrictionPolicyBinding {
	this := RestrictionPolicyBinding{}
	return &this
}

// GetPrincipals returns the Principals field value.
func (o *RestrictionPolicyBinding) GetPrincipals() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Principals
}

// GetPrincipalsOk returns a tuple with the Principals field value
// and a boolean to check if the value has been set.
func (o *RestrictionPolicyBinding) GetPrincipalsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principals, true
}

// SetPrincipals sets field value.
func (o *RestrictionPolicyBinding) SetPrincipals(v []string) {
	o.Principals = v
}

// GetRelation returns the Relation field value.
func (o *RestrictionPolicyBinding) GetRelation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Relation
}

// GetRelationOk returns a tuple with the Relation field value
// and a boolean to check if the value has been set.
func (o *RestrictionPolicyBinding) GetRelationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relation, true
}

// SetRelation sets field value.
func (o *RestrictionPolicyBinding) SetRelation(v string) {
	o.Relation = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RestrictionPolicyBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["principals"] = o.Principals
	toSerialize["relation"] = o.Relation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RestrictionPolicyBinding) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Principals *[]string `json:"principals"`
		Relation   *string   `json:"relation"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Principals == nil {
		return fmt.Errorf("required field principals missing")
	}
	if all.Relation == nil {
		return fmt.Errorf("required field relation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"principals", "relation"})
	} else {
		return err
	}
	o.Principals = *all.Principals
	o.Relation = *all.Relation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
