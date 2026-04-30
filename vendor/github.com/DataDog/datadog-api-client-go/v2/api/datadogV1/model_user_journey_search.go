// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserJourneySearch User journey search configuration.
type UserJourneySearch struct {
	// Expression string.
	Expression string `json:"expression"`
	// Filters for user journey search.
	Filters *UserJourneySearchFilters `json:"filters,omitempty"`
	// Join keys for user journey queries.
	JoinKeys *UserJourneyJoinKeys `json:"join_keys,omitempty"`
	// Node objects mapping.
	NodeObjects map[string]ProductAnalyticsBaseQuery `json:"node_objects"`
	// Step aliases mapping.
	StepAliases map[string]string `json:"step_aliases,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewUserJourneySearch instantiates a new UserJourneySearch object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserJourneySearch(expression string, nodeObjects map[string]ProductAnalyticsBaseQuery) *UserJourneySearch {
	this := UserJourneySearch{}
	this.Expression = expression
	this.NodeObjects = nodeObjects
	return &this
}

// NewUserJourneySearchWithDefaults instantiates a new UserJourneySearch object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserJourneySearchWithDefaults() *UserJourneySearch {
	this := UserJourneySearch{}
	return &this
}

// GetExpression returns the Expression field value.
func (o *UserJourneySearch) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *UserJourneySearch) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value.
func (o *UserJourneySearch) SetExpression(v string) {
	o.Expression = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *UserJourneySearch) GetFilters() UserJourneySearchFilters {
	if o == nil || o.Filters == nil {
		var ret UserJourneySearchFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneySearch) GetFiltersOk() (*UserJourneySearchFilters, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *UserJourneySearch) HasFilters() bool {
	return o != nil && o.Filters != nil
}

// SetFilters gets a reference to the given UserJourneySearchFilters and assigns it to the Filters field.
func (o *UserJourneySearch) SetFilters(v UserJourneySearchFilters) {
	o.Filters = &v
}

// GetJoinKeys returns the JoinKeys field value if set, zero value otherwise.
func (o *UserJourneySearch) GetJoinKeys() UserJourneyJoinKeys {
	if o == nil || o.JoinKeys == nil {
		var ret UserJourneyJoinKeys
		return ret
	}
	return *o.JoinKeys
}

// GetJoinKeysOk returns a tuple with the JoinKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneySearch) GetJoinKeysOk() (*UserJourneyJoinKeys, bool) {
	if o == nil || o.JoinKeys == nil {
		return nil, false
	}
	return o.JoinKeys, true
}

// HasJoinKeys returns a boolean if a field has been set.
func (o *UserJourneySearch) HasJoinKeys() bool {
	return o != nil && o.JoinKeys != nil
}

// SetJoinKeys gets a reference to the given UserJourneyJoinKeys and assigns it to the JoinKeys field.
func (o *UserJourneySearch) SetJoinKeys(v UserJourneyJoinKeys) {
	o.JoinKeys = &v
}

// GetNodeObjects returns the NodeObjects field value.
func (o *UserJourneySearch) GetNodeObjects() map[string]ProductAnalyticsBaseQuery {
	if o == nil {
		var ret map[string]ProductAnalyticsBaseQuery
		return ret
	}
	return o.NodeObjects
}

// GetNodeObjectsOk returns a tuple with the NodeObjects field value
// and a boolean to check if the value has been set.
func (o *UserJourneySearch) GetNodeObjectsOk() (*map[string]ProductAnalyticsBaseQuery, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeObjects, true
}

// SetNodeObjects sets field value.
func (o *UserJourneySearch) SetNodeObjects(v map[string]ProductAnalyticsBaseQuery) {
	o.NodeObjects = v
}

// GetStepAliases returns the StepAliases field value if set, zero value otherwise.
func (o *UserJourneySearch) GetStepAliases() map[string]string {
	if o == nil || o.StepAliases == nil {
		var ret map[string]string
		return ret
	}
	return o.StepAliases
}

// GetStepAliasesOk returns a tuple with the StepAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneySearch) GetStepAliasesOk() (*map[string]string, bool) {
	if o == nil || o.StepAliases == nil {
		return nil, false
	}
	return &o.StepAliases, true
}

// HasStepAliases returns a boolean if a field has been set.
func (o *UserJourneySearch) HasStepAliases() bool {
	return o != nil && o.StepAliases != nil
}

// SetStepAliases gets a reference to the given map[string]string and assigns it to the StepAliases field.
func (o *UserJourneySearch) SetStepAliases(v map[string]string) {
	o.StepAliases = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserJourneySearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["expression"] = o.Expression
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}
	if o.JoinKeys != nil {
		toSerialize["join_keys"] = o.JoinKeys
	}
	toSerialize["node_objects"] = o.NodeObjects
	if o.StepAliases != nil {
		toSerialize["step_aliases"] = o.StepAliases
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UserJourneySearch) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Expression  *string                               `json:"expression"`
		Filters     *UserJourneySearchFilters             `json:"filters,omitempty"`
		JoinKeys    *UserJourneyJoinKeys                  `json:"join_keys,omitempty"`
		NodeObjects *map[string]ProductAnalyticsBaseQuery `json:"node_objects"`
		StepAliases map[string]string                     `json:"step_aliases,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Expression == nil {
		return fmt.Errorf("required field expression missing")
	}
	if all.NodeObjects == nil {
		return fmt.Errorf("required field node_objects missing")
	}

	hasInvalidField := false
	o.Expression = *all.Expression
	if all.Filters != nil && all.Filters.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filters = all.Filters
	if all.JoinKeys != nil && all.JoinKeys.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.JoinKeys = all.JoinKeys
	o.NodeObjects = *all.NodeObjects
	o.StepAliases = all.StepAliases

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
