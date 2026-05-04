// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IntegrationOnCallEscalationQueriesItemsTarget The target recipient for an On-Call escalation query.
type IntegrationOnCallEscalationQueriesItemsTarget struct {
	// Whether to use dynamic team paging for escalation.
	DynamicTeamPaging *bool `json:"dynamic_team_paging,omitempty"`
	// The identifier of the team to escalate to.
	TeamId *string `json:"team_id,omitempty"`
	// The identifier of the user to escalate to.
	UserId *string `json:"user_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIntegrationOnCallEscalationQueriesItemsTarget instantiates a new IntegrationOnCallEscalationQueriesItemsTarget object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIntegrationOnCallEscalationQueriesItemsTarget() *IntegrationOnCallEscalationQueriesItemsTarget {
	this := IntegrationOnCallEscalationQueriesItemsTarget{}
	return &this
}

// NewIntegrationOnCallEscalationQueriesItemsTargetWithDefaults instantiates a new IntegrationOnCallEscalationQueriesItemsTarget object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIntegrationOnCallEscalationQueriesItemsTargetWithDefaults() *IntegrationOnCallEscalationQueriesItemsTarget {
	this := IntegrationOnCallEscalationQueriesItemsTarget{}
	return &this
}

// GetDynamicTeamPaging returns the DynamicTeamPaging field value if set, zero value otherwise.
func (o *IntegrationOnCallEscalationQueriesItemsTarget) GetDynamicTeamPaging() bool {
	if o == nil || o.DynamicTeamPaging == nil {
		var ret bool
		return ret
	}
	return *o.DynamicTeamPaging
}

// GetDynamicTeamPagingOk returns a tuple with the DynamicTeamPaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOnCallEscalationQueriesItemsTarget) GetDynamicTeamPagingOk() (*bool, bool) {
	if o == nil || o.DynamicTeamPaging == nil {
		return nil, false
	}
	return o.DynamicTeamPaging, true
}

// HasDynamicTeamPaging returns a boolean if a field has been set.
func (o *IntegrationOnCallEscalationQueriesItemsTarget) HasDynamicTeamPaging() bool {
	return o != nil && o.DynamicTeamPaging != nil
}

// SetDynamicTeamPaging gets a reference to the given bool and assigns it to the DynamicTeamPaging field.
func (o *IntegrationOnCallEscalationQueriesItemsTarget) SetDynamicTeamPaging(v bool) {
	o.DynamicTeamPaging = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *IntegrationOnCallEscalationQueriesItemsTarget) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOnCallEscalationQueriesItemsTarget) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *IntegrationOnCallEscalationQueriesItemsTarget) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *IntegrationOnCallEscalationQueriesItemsTarget) SetTeamId(v string) {
	o.TeamId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *IntegrationOnCallEscalationQueriesItemsTarget) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationOnCallEscalationQueriesItemsTarget) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *IntegrationOnCallEscalationQueriesItemsTarget) HasUserId() bool {
	return o != nil && o.UserId != nil
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *IntegrationOnCallEscalationQueriesItemsTarget) SetUserId(v string) {
	o.UserId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IntegrationOnCallEscalationQueriesItemsTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DynamicTeamPaging != nil {
		toSerialize["dynamic_team_paging"] = o.DynamicTeamPaging
	}
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IntegrationOnCallEscalationQueriesItemsTarget) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DynamicTeamPaging *bool   `json:"dynamic_team_paging,omitempty"`
		TeamId            *string `json:"team_id,omitempty"`
		UserId            *string `json:"user_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dynamic_team_paging", "team_id", "user_id"})
	} else {
		return err
	}
	o.DynamicTeamPaging = all.DynamicTeamPaging
	o.TeamId = all.TeamId
	o.UserId = all.UserId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
