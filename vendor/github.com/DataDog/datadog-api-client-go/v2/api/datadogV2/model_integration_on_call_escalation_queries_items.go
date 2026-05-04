// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationOnCallEscalationQueriesItems An On-Call escalation query entry used to route cases to on-call responders.
type IntegrationOnCallEscalationQueriesItems struct {
	// Whether this escalation query is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Unique identifier of the escalation query.
	Id *string `json:"id,omitempty"`
	// The query used to match cases for escalation.
	Query *string `json:"query,omitempty"`
	// The target recipient for an On-Call escalation query.
	Target *IntegrationOnCallEscalationQueriesItemsTarget `json:"target,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationOnCallEscalationQueriesItems instantiates a new IntegrationOnCallEscalationQueriesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationOnCallEscalationQueriesItems() *IntegrationOnCallEscalationQueriesItems {
	this := IntegrationOnCallEscalationQueriesItems{}
	return &this
}

// NewIntegrationOnCallEscalationQueriesItemsWithDefaults instantiates a new IntegrationOnCallEscalationQueriesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationOnCallEscalationQueriesItemsWithDefaults() *IntegrationOnCallEscalationQueriesItems {
	this := IntegrationOnCallEscalationQueriesItems{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IntegrationOnCallEscalationQueriesItems) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOnCallEscalationQueriesItems) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IntegrationOnCallEscalationQueriesItems) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IntegrationOnCallEscalationQueriesItems) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegrationOnCallEscalationQueriesItems) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOnCallEscalationQueriesItems) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegrationOnCallEscalationQueriesItems) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IntegrationOnCallEscalationQueriesItems) SetId(v string) {
	o.Id = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *IntegrationOnCallEscalationQueriesItems) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOnCallEscalationQueriesItems) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *IntegrationOnCallEscalationQueriesItems) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *IntegrationOnCallEscalationQueriesItems) SetQuery(v string) {
	o.Query = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *IntegrationOnCallEscalationQueriesItems) GetTarget() IntegrationOnCallEscalationQueriesItemsTarget {
	if o == nil || o.Target == nil {
		var ret IntegrationOnCallEscalationQueriesItemsTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOnCallEscalationQueriesItems) GetTargetOk() (*IntegrationOnCallEscalationQueriesItemsTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *IntegrationOnCallEscalationQueriesItems) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given IntegrationOnCallEscalationQueriesItemsTarget and assigns it to the Target field.
func (o *IntegrationOnCallEscalationQueriesItems) SetTarget(v IntegrationOnCallEscalationQueriesItemsTarget) {
	o.Target = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationOnCallEscalationQueriesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationOnCallEscalationQueriesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled *bool                                          `json:"enabled,omitempty"`
		Id      *string                                        `json:"id,omitempty"`
		Query   *string                                        `json:"query,omitempty"`
		Target  *IntegrationOnCallEscalationQueriesItemsTarget `json:"target,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "id", "query", "target"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = all.Enabled
	o.Id = all.Id
	o.Query = all.Query
	if all.Target != nil && all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = all.Target

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
