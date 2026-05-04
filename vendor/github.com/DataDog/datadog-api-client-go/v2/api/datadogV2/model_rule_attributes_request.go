// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RuleAttributesRequest Attributes for creating or updating a rule. Server-managed fields (created_at, modified_at, custom) are excluded.
type RuleAttributesRequest struct {
	// Explanation of the rule.
	Description *string `json:"description,omitempty"`
	// If enabled, the rule is calculated as part of the score.
	Enabled *bool `json:"enabled,omitempty"`
	// The maturity level of the rule (1, 2, or 3).
	Level *int32 `json:"level,omitempty"`
	// Name of the rule.
	Name *string `json:"name,omitempty"`
	// Owner of the rule.
	Owner *string `json:"owner,omitempty"`
	// A query to filter which entities this rule applies to.
	ScopeQuery *string `json:"scope_query,omitempty"`
	// The scorecard name to which this rule must belong.
	ScorecardName *string `json:"scorecard_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRuleAttributesRequest instantiates a new RuleAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRuleAttributesRequest() *RuleAttributesRequest {
	this := RuleAttributesRequest{}
	return &this
}

// NewRuleAttributesRequestWithDefaults instantiates a new RuleAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRuleAttributesRequestWithDefaults() *RuleAttributesRequest {
	this := RuleAttributesRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RuleAttributesRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributesRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RuleAttributesRequest) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RuleAttributesRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RuleAttributesRequest) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributesRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RuleAttributesRequest) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RuleAttributesRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *RuleAttributesRequest) GetLevel() int32 {
	if o == nil || o.Level == nil {
		var ret int32
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributesRequest) GetLevelOk() (*int32, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *RuleAttributesRequest) HasLevel() bool {
	return o != nil && o.Level != nil
}

// SetLevel gets a reference to the given int32 and assigns it to the Level field.
func (o *RuleAttributesRequest) SetLevel(v int32) {
	o.Level = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RuleAttributesRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributesRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RuleAttributesRequest) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RuleAttributesRequest) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *RuleAttributesRequest) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributesRequest) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *RuleAttributesRequest) HasOwner() bool {
	return o != nil && o.Owner != nil
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *RuleAttributesRequest) SetOwner(v string) {
	o.Owner = &v
}

// GetScopeQuery returns the ScopeQuery field value if set, zero value otherwise.
func (o *RuleAttributesRequest) GetScopeQuery() string {
	if o == nil || o.ScopeQuery == nil {
		var ret string
		return ret
	}
	return *o.ScopeQuery
}

// GetScopeQueryOk returns a tuple with the ScopeQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributesRequest) GetScopeQueryOk() (*string, bool) {
	if o == nil || o.ScopeQuery == nil {
		return nil, false
	}
	return o.ScopeQuery, true
}

// HasScopeQuery returns a boolean if a field has been set.
func (o *RuleAttributesRequest) HasScopeQuery() bool {
	return o != nil && o.ScopeQuery != nil
}

// SetScopeQuery gets a reference to the given string and assigns it to the ScopeQuery field.
func (o *RuleAttributesRequest) SetScopeQuery(v string) {
	o.ScopeQuery = &v
}

// GetScorecardName returns the ScorecardName field value if set, zero value otherwise.
func (o *RuleAttributesRequest) GetScorecardName() string {
	if o == nil || o.ScorecardName == nil {
		var ret string
		return ret
	}
	return *o.ScorecardName
}

// GetScorecardNameOk returns a tuple with the ScorecardName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleAttributesRequest) GetScorecardNameOk() (*string, bool) {
	if o == nil || o.ScorecardName == nil {
		return nil, false
	}
	return o.ScorecardName, true
}

// HasScorecardName returns a boolean if a field has been set.
func (o *RuleAttributesRequest) HasScorecardName() bool {
	return o != nil && o.ScorecardName != nil
}

// SetScorecardName gets a reference to the given string and assigns it to the ScorecardName field.
func (o *RuleAttributesRequest) SetScorecardName(v string) {
	o.ScorecardName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RuleAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.ScopeQuery != nil {
		toSerialize["scope_query"] = o.ScopeQuery
	}
	if o.ScorecardName != nil {
		toSerialize["scorecard_name"] = o.ScorecardName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RuleAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description   *string `json:"description,omitempty"`
		Enabled       *bool   `json:"enabled,omitempty"`
		Level         *int32  `json:"level,omitempty"`
		Name          *string `json:"name,omitempty"`
		Owner         *string `json:"owner,omitempty"`
		ScopeQuery    *string `json:"scope_query,omitempty"`
		ScorecardName *string `json:"scorecard_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "enabled", "level", "name", "owner", "scope_query", "scorecard_name"})
	} else {
		return err
	}
	o.Description = all.Description
	o.Enabled = all.Enabled
	o.Level = all.Level
	o.Name = all.Name
	o.Owner = all.Owner
	o.ScopeQuery = all.ScopeQuery
	o.ScorecardName = all.ScorecardName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
