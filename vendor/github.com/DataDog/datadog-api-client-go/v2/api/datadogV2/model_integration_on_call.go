// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationOnCall On-Call integration settings.
type IntegrationOnCall struct {
	// Whether to auto-assign on-call.
	AutoAssignOnCall *bool `json:"auto_assign_on_call,omitempty"`
	// Whether On-Call integration is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of escalation queries for routing cases to on-call responders.
	EscalationQueries []IntegrationOnCallEscalationQueriesItems `json:"escalation_queries,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationOnCall instantiates a new IntegrationOnCall object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationOnCall() *IntegrationOnCall {
	this := IntegrationOnCall{}
	return &this
}

// NewIntegrationOnCallWithDefaults instantiates a new IntegrationOnCall object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationOnCallWithDefaults() *IntegrationOnCall {
	this := IntegrationOnCall{}
	return &this
}

// GetAutoAssignOnCall returns the AutoAssignOnCall field value if set, zero value otherwise.
func (o *IntegrationOnCall) GetAutoAssignOnCall() bool {
	if o == nil || o.AutoAssignOnCall == nil {
		var ret bool
		return ret
	}
	return *o.AutoAssignOnCall
}

// GetAutoAssignOnCallOk returns a tuple with the AutoAssignOnCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOnCall) GetAutoAssignOnCallOk() (*bool, bool) {
	if o == nil || o.AutoAssignOnCall == nil {
		return nil, false
	}
	return o.AutoAssignOnCall, true
}

// HasAutoAssignOnCall returns a boolean if a field has been set.
func (o *IntegrationOnCall) HasAutoAssignOnCall() bool {
	return o != nil && o.AutoAssignOnCall != nil
}

// SetAutoAssignOnCall gets a reference to the given bool and assigns it to the AutoAssignOnCall field.
func (o *IntegrationOnCall) SetAutoAssignOnCall(v bool) {
	o.AutoAssignOnCall = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *IntegrationOnCall) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOnCall) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *IntegrationOnCall) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *IntegrationOnCall) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEscalationQueries returns the EscalationQueries field value if set, zero value otherwise.
func (o *IntegrationOnCall) GetEscalationQueries() []IntegrationOnCallEscalationQueriesItems {
	if o == nil || o.EscalationQueries == nil {
		var ret []IntegrationOnCallEscalationQueriesItems
		return ret
	}
	return o.EscalationQueries
}

// GetEscalationQueriesOk returns a tuple with the EscalationQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOnCall) GetEscalationQueriesOk() (*[]IntegrationOnCallEscalationQueriesItems, bool) {
	if o == nil || o.EscalationQueries == nil {
		return nil, false
	}
	return &o.EscalationQueries, true
}

// HasEscalationQueries returns a boolean if a field has been set.
func (o *IntegrationOnCall) HasEscalationQueries() bool {
	return o != nil && o.EscalationQueries != nil
}

// SetEscalationQueries gets a reference to the given []IntegrationOnCallEscalationQueriesItems and assigns it to the EscalationQueries field.
func (o *IntegrationOnCall) SetEscalationQueries(v []IntegrationOnCallEscalationQueriesItems) {
	o.EscalationQueries = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationOnCall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutoAssignOnCall != nil {
		toSerialize["auto_assign_on_call"] = o.AutoAssignOnCall
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.EscalationQueries != nil {
		toSerialize["escalation_queries"] = o.EscalationQueries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationOnCall) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoAssignOnCall  *bool                                     `json:"auto_assign_on_call,omitempty"`
		Enabled           *bool                                     `json:"enabled,omitempty"`
		EscalationQueries []IntegrationOnCallEscalationQueriesItems `json:"escalation_queries,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_assign_on_call", "enabled", "escalation_queries"})
	} else {
		return err
	}
	o.AutoAssignOnCall = all.AutoAssignOnCall
	o.Enabled = all.Enabled
	o.EscalationQueries = all.EscalationQueries

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
