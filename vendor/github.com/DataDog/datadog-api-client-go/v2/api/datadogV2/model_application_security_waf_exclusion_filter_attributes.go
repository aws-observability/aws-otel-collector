// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafExclusionFilterAttributes Attributes describing a WAF exclusion filter.
type ApplicationSecurityWafExclusionFilterAttributes struct {
	// A description for the exclusion filter.
	Description *string `json:"description,omitempty"`
	// Indicates whether the exclusion filter is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The event query matched by the legacy exclusion filter. Cannot be created nor updated.
	EventQuery *string `json:"event_query,omitempty"`
	// The client IP addresses matched by the exclusion filter (CIDR notation is supported).
	IpList []string `json:"ip_list,omitempty"`
	// Extra information about the exclusion filter.
	Metadata *ApplicationSecurityWafExclusionFilterMetadata `json:"metadata,omitempty"`
	// The action taken when the exclusion filter matches. When set to `monitor`, security traces are emitted but the requests are not blocked. By default, security traces are not emitted and the requests are not blocked.
	OnMatch *ApplicationSecurityWafExclusionFilterOnMatch `json:"on_match,omitempty"`
	// A list of parameters matched by the exclusion filter in the HTTP query string and HTTP request body. Nested parameters can be matched by joining fields with a dot character.
	Parameters []string `json:"parameters,omitempty"`
	// The HTTP path glob expression matched by the exclusion filter.
	PathGlob *string `json:"path_glob,omitempty"`
	// The WAF rules targeted by the exclusion filter.
	RulesTarget []ApplicationSecurityWafExclusionFilterRulesTarget `json:"rules_target,omitempty"`
	// The services where the exclusion filter is deployed.
	Scope []ApplicationSecurityWafExclusionFilterScope `json:"scope,omitempty"`
	// Generated event search query for traces matching the exclusion filter.
	SearchQuery *string `json:"search_query,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityWafExclusionFilterAttributes instantiates a new ApplicationSecurityWafExclusionFilterAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityWafExclusionFilterAttributes() *ApplicationSecurityWafExclusionFilterAttributes {
	this := ApplicationSecurityWafExclusionFilterAttributes{}
	return &this
}

// NewApplicationSecurityWafExclusionFilterAttributesWithDefaults instantiates a new ApplicationSecurityWafExclusionFilterAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityWafExclusionFilterAttributesWithDefaults() *ApplicationSecurityWafExclusionFilterAttributes {
	this := ApplicationSecurityWafExclusionFilterAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationSecurityWafExclusionFilterAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplicationSecurityWafExclusionFilterAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEventQuery returns the EventQuery field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetEventQuery() string {
	if o == nil || o.EventQuery == nil {
		var ret string
		return ret
	}
	return *o.EventQuery
}

// GetEventQueryOk returns a tuple with the EventQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetEventQueryOk() (*string, bool) {
	if o == nil || o.EventQuery == nil {
		return nil, false
	}
	return o.EventQuery, true
}

// HasEventQuery returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) HasEventQuery() bool {
	return o != nil && o.EventQuery != nil
}

// SetEventQuery gets a reference to the given string and assigns it to the EventQuery field.
func (o *ApplicationSecurityWafExclusionFilterAttributes) SetEventQuery(v string) {
	o.EventQuery = &v
}

// GetIpList returns the IpList field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetIpList() []string {
	if o == nil || o.IpList == nil {
		var ret []string
		return ret
	}
	return o.IpList
}

// GetIpListOk returns a tuple with the IpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetIpListOk() (*[]string, bool) {
	if o == nil || o.IpList == nil {
		return nil, false
	}
	return &o.IpList, true
}

// HasIpList returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) HasIpList() bool {
	return o != nil && o.IpList != nil
}

// SetIpList gets a reference to the given []string and assigns it to the IpList field.
func (o *ApplicationSecurityWafExclusionFilterAttributes) SetIpList(v []string) {
	o.IpList = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetMetadata() ApplicationSecurityWafExclusionFilterMetadata {
	if o == nil || o.Metadata == nil {
		var ret ApplicationSecurityWafExclusionFilterMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetMetadataOk() (*ApplicationSecurityWafExclusionFilterMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given ApplicationSecurityWafExclusionFilterMetadata and assigns it to the Metadata field.
func (o *ApplicationSecurityWafExclusionFilterAttributes) SetMetadata(v ApplicationSecurityWafExclusionFilterMetadata) {
	o.Metadata = &v
}

// GetOnMatch returns the OnMatch field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetOnMatch() ApplicationSecurityWafExclusionFilterOnMatch {
	if o == nil || o.OnMatch == nil {
		var ret ApplicationSecurityWafExclusionFilterOnMatch
		return ret
	}
	return *o.OnMatch
}

// GetOnMatchOk returns a tuple with the OnMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetOnMatchOk() (*ApplicationSecurityWafExclusionFilterOnMatch, bool) {
	if o == nil || o.OnMatch == nil {
		return nil, false
	}
	return o.OnMatch, true
}

// HasOnMatch returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) HasOnMatch() bool {
	return o != nil && o.OnMatch != nil
}

// SetOnMatch gets a reference to the given ApplicationSecurityWafExclusionFilterOnMatch and assigns it to the OnMatch field.
func (o *ApplicationSecurityWafExclusionFilterAttributes) SetOnMatch(v ApplicationSecurityWafExclusionFilterOnMatch) {
	o.OnMatch = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetParameters() []string {
	if o == nil || o.Parameters == nil {
		var ret []string
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetParametersOk() (*[]string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) HasParameters() bool {
	return o != nil && o.Parameters != nil
}

// SetParameters gets a reference to the given []string and assigns it to the Parameters field.
func (o *ApplicationSecurityWafExclusionFilterAttributes) SetParameters(v []string) {
	o.Parameters = v
}

// GetPathGlob returns the PathGlob field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetPathGlob() string {
	if o == nil || o.PathGlob == nil {
		var ret string
		return ret
	}
	return *o.PathGlob
}

// GetPathGlobOk returns a tuple with the PathGlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetPathGlobOk() (*string, bool) {
	if o == nil || o.PathGlob == nil {
		return nil, false
	}
	return o.PathGlob, true
}

// HasPathGlob returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) HasPathGlob() bool {
	return o != nil && o.PathGlob != nil
}

// SetPathGlob gets a reference to the given string and assigns it to the PathGlob field.
func (o *ApplicationSecurityWafExclusionFilterAttributes) SetPathGlob(v string) {
	o.PathGlob = &v
}

// GetRulesTarget returns the RulesTarget field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetRulesTarget() []ApplicationSecurityWafExclusionFilterRulesTarget {
	if o == nil || o.RulesTarget == nil {
		var ret []ApplicationSecurityWafExclusionFilterRulesTarget
		return ret
	}
	return o.RulesTarget
}

// GetRulesTargetOk returns a tuple with the RulesTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetRulesTargetOk() (*[]ApplicationSecurityWafExclusionFilterRulesTarget, bool) {
	if o == nil || o.RulesTarget == nil {
		return nil, false
	}
	return &o.RulesTarget, true
}

// HasRulesTarget returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) HasRulesTarget() bool {
	return o != nil && o.RulesTarget != nil
}

// SetRulesTarget gets a reference to the given []ApplicationSecurityWafExclusionFilterRulesTarget and assigns it to the RulesTarget field.
func (o *ApplicationSecurityWafExclusionFilterAttributes) SetRulesTarget(v []ApplicationSecurityWafExclusionFilterRulesTarget) {
	o.RulesTarget = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetScope() []ApplicationSecurityWafExclusionFilterScope {
	if o == nil || o.Scope == nil {
		var ret []ApplicationSecurityWafExclusionFilterScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetScopeOk() (*[]ApplicationSecurityWafExclusionFilterScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given []ApplicationSecurityWafExclusionFilterScope and assigns it to the Scope field.
func (o *ApplicationSecurityWafExclusionFilterAttributes) SetScope(v []ApplicationSecurityWafExclusionFilterScope) {
	o.Scope = v
}

// GetSearchQuery returns the SearchQuery field value if set, zero value otherwise.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetSearchQuery() string {
	if o == nil || o.SearchQuery == nil {
		var ret string
		return ret
	}
	return *o.SearchQuery
}

// GetSearchQueryOk returns a tuple with the SearchQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) GetSearchQueryOk() (*string, bool) {
	if o == nil || o.SearchQuery == nil {
		return nil, false
	}
	return o.SearchQuery, true
}

// HasSearchQuery returns a boolean if a field has been set.
func (o *ApplicationSecurityWafExclusionFilterAttributes) HasSearchQuery() bool {
	return o != nil && o.SearchQuery != nil
}

// SetSearchQuery gets a reference to the given string and assigns it to the SearchQuery field.
func (o *ApplicationSecurityWafExclusionFilterAttributes) SetSearchQuery(v string) {
	o.SearchQuery = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityWafExclusionFilterAttributes) MarshalJSON() ([]byte, error) {
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
	if o.EventQuery != nil {
		toSerialize["event_query"] = o.EventQuery
	}
	if o.IpList != nil {
		toSerialize["ip_list"] = o.IpList
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.OnMatch != nil {
		toSerialize["on_match"] = o.OnMatch
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.PathGlob != nil {
		toSerialize["path_glob"] = o.PathGlob
	}
	if o.RulesTarget != nil {
		toSerialize["rules_target"] = o.RulesTarget
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.SearchQuery != nil {
		toSerialize["search_query"] = o.SearchQuery
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityWafExclusionFilterAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description *string                                            `json:"description,omitempty"`
		Enabled     *bool                                              `json:"enabled,omitempty"`
		EventQuery  *string                                            `json:"event_query,omitempty"`
		IpList      []string                                           `json:"ip_list,omitempty"`
		Metadata    *ApplicationSecurityWafExclusionFilterMetadata     `json:"metadata,omitempty"`
		OnMatch     *ApplicationSecurityWafExclusionFilterOnMatch      `json:"on_match,omitempty"`
		Parameters  []string                                           `json:"parameters,omitempty"`
		PathGlob    *string                                            `json:"path_glob,omitempty"`
		RulesTarget []ApplicationSecurityWafExclusionFilterRulesTarget `json:"rules_target,omitempty"`
		Scope       []ApplicationSecurityWafExclusionFilterScope       `json:"scope,omitempty"`
		SearchQuery *string                                            `json:"search_query,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "enabled", "event_query", "ip_list", "metadata", "on_match", "parameters", "path_glob", "rules_target", "scope", "search_query"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.Enabled = all.Enabled
	o.EventQuery = all.EventQuery
	o.IpList = all.IpList
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	if all.OnMatch != nil && !all.OnMatch.IsValid() {
		hasInvalidField = true
	} else {
		o.OnMatch = all.OnMatch
	}
	o.Parameters = all.Parameters
	o.PathGlob = all.PathGlob
	o.RulesTarget = all.RulesTarget
	o.Scope = all.Scope
	o.SearchQuery = all.SearchQuery

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
