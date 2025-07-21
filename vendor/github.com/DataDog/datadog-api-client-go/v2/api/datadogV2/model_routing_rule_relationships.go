// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RoutingRuleRelationships Specifies relationships for a routing rule, linking to associated policy resources.
type RoutingRuleRelationships struct {
	// Defines the relationship that links a routing rule to a policy.
	Policy *RoutingRuleRelationshipsPolicy `json:"policy,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRoutingRuleRelationships instantiates a new RoutingRuleRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRoutingRuleRelationships() *RoutingRuleRelationships {
	this := RoutingRuleRelationships{}
	return &this
}

// NewRoutingRuleRelationshipsWithDefaults instantiates a new RoutingRuleRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRoutingRuleRelationshipsWithDefaults() *RoutingRuleRelationships {
	this := RoutingRuleRelationships{}
	return &this
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *RoutingRuleRelationships) GetPolicy() RoutingRuleRelationshipsPolicy {
	if o == nil || o.Policy == nil {
		var ret RoutingRuleRelationshipsPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingRuleRelationships) GetPolicyOk() (*RoutingRuleRelationshipsPolicy, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *RoutingRuleRelationships) HasPolicy() bool {
	return o != nil && o.Policy != nil
}

// SetPolicy gets a reference to the given RoutingRuleRelationshipsPolicy and assigns it to the Policy field.
func (o *RoutingRuleRelationships) SetPolicy(v RoutingRuleRelationshipsPolicy) {
	o.Policy = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RoutingRuleRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RoutingRuleRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Policy *RoutingRuleRelationshipsPolicy `json:"policy,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"policy"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Policy != nil && all.Policy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Policy = all.Policy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
