// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateRulesetRequestDataAttributesRulesItemsQuery The definition of `UpdateRulesetRequestDataAttributesRulesItemsQuery` object.
type UpdateRulesetRequestDataAttributesRulesItemsQuery struct {
	// The definition of `UpdateRulesetRequestDataAttributesRulesItemsQueryAddition` object.
	Addition NullableUpdateRulesetRequestDataAttributesRulesItemsQueryAddition `json:"addition"`
	// The `query` `case_insensitivity`.
	CaseInsensitivity *bool `json:"case_insensitivity,omitempty"`
	// Deprecated. Use `if_tag_exists` instead. The `query` `if_not_exists`.
	// Deprecated
	IfNotExists *bool `json:"if_not_exists,omitempty"`
	// The behavior when the tag already exists.
	IfTagExists *DataAttributesRulesItemsIfTagExists `json:"if_tag_exists,omitempty"`
	// The `query` `query`.
	Query string `json:"query"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateRulesetRequestDataAttributesRulesItemsQuery instantiates a new UpdateRulesetRequestDataAttributesRulesItemsQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateRulesetRequestDataAttributesRulesItemsQuery(addition NullableUpdateRulesetRequestDataAttributesRulesItemsQueryAddition, query string) *UpdateRulesetRequestDataAttributesRulesItemsQuery {
	this := UpdateRulesetRequestDataAttributesRulesItemsQuery{}
	this.Addition = addition
	this.Query = query
	return &this
}

// NewUpdateRulesetRequestDataAttributesRulesItemsQueryWithDefaults instantiates a new UpdateRulesetRequestDataAttributesRulesItemsQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateRulesetRequestDataAttributesRulesItemsQueryWithDefaults() *UpdateRulesetRequestDataAttributesRulesItemsQuery {
	this := UpdateRulesetRequestDataAttributesRulesItemsQuery{}
	return &this
}

// GetAddition returns the Addition field value.
// If the value is explicit nil, the zero value for UpdateRulesetRequestDataAttributesRulesItemsQueryAddition will be returned.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) GetAddition() UpdateRulesetRequestDataAttributesRulesItemsQueryAddition {
	if o == nil || o.Addition.Get() == nil {
		var ret UpdateRulesetRequestDataAttributesRulesItemsQueryAddition
		return ret
	}
	return *o.Addition.Get()
}

// GetAdditionOk returns a tuple with the Addition field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) GetAdditionOk() (*UpdateRulesetRequestDataAttributesRulesItemsQueryAddition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addition.Get(), o.Addition.IsSet()
}

// SetAddition sets field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) SetAddition(v UpdateRulesetRequestDataAttributesRulesItemsQueryAddition) {
	o.Addition.Set(&v)
}

// GetCaseInsensitivity returns the CaseInsensitivity field value if set, zero value otherwise.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) GetCaseInsensitivity() bool {
	if o == nil || o.CaseInsensitivity == nil {
		var ret bool
		return ret
	}
	return *o.CaseInsensitivity
}

// GetCaseInsensitivityOk returns a tuple with the CaseInsensitivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) GetCaseInsensitivityOk() (*bool, bool) {
	if o == nil || o.CaseInsensitivity == nil {
		return nil, false
	}
	return o.CaseInsensitivity, true
}

// HasCaseInsensitivity returns a boolean if a field has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) HasCaseInsensitivity() bool {
	return o != nil && o.CaseInsensitivity != nil
}

// SetCaseInsensitivity gets a reference to the given bool and assigns it to the CaseInsensitivity field.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) SetCaseInsensitivity(v bool) {
	o.CaseInsensitivity = &v
}

// GetIfNotExists returns the IfNotExists field value if set, zero value otherwise.
// Deprecated
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) GetIfNotExists() bool {
	if o == nil || o.IfNotExists == nil {
		var ret bool
		return ret
	}
	return *o.IfNotExists
}

// GetIfNotExistsOk returns a tuple with the IfNotExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) GetIfNotExistsOk() (*bool, bool) {
	if o == nil || o.IfNotExists == nil {
		return nil, false
	}
	return o.IfNotExists, true
}

// HasIfNotExists returns a boolean if a field has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) HasIfNotExists() bool {
	return o != nil && o.IfNotExists != nil
}

// SetIfNotExists gets a reference to the given bool and assigns it to the IfNotExists field.
// Deprecated
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) SetIfNotExists(v bool) {
	o.IfNotExists = &v
}

// GetIfTagExists returns the IfTagExists field value if set, zero value otherwise.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) GetIfTagExists() DataAttributesRulesItemsIfTagExists {
	if o == nil || o.IfTagExists == nil {
		var ret DataAttributesRulesItemsIfTagExists
		return ret
	}
	return *o.IfTagExists
}

// GetIfTagExistsOk returns a tuple with the IfTagExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) GetIfTagExistsOk() (*DataAttributesRulesItemsIfTagExists, bool) {
	if o == nil || o.IfTagExists == nil {
		return nil, false
	}
	return o.IfTagExists, true
}

// HasIfTagExists returns a boolean if a field has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) HasIfTagExists() bool {
	return o != nil && o.IfTagExists != nil
}

// SetIfTagExists gets a reference to the given DataAttributesRulesItemsIfTagExists and assigns it to the IfTagExists field.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) SetIfTagExists(v DataAttributesRulesItemsIfTagExists) {
	o.IfTagExists = &v
}

// GetQuery returns the Query field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) SetQuery(v string) {
	o.Query = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateRulesetRequestDataAttributesRulesItemsQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["addition"] = o.Addition.Get()
	if o.CaseInsensitivity != nil {
		toSerialize["case_insensitivity"] = o.CaseInsensitivity
	}
	if o.IfNotExists != nil {
		toSerialize["if_not_exists"] = o.IfNotExists
	}
	if o.IfTagExists != nil {
		toSerialize["if_tag_exists"] = o.IfTagExists
	}
	toSerialize["query"] = o.Query

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateRulesetRequestDataAttributesRulesItemsQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Addition          NullableUpdateRulesetRequestDataAttributesRulesItemsQueryAddition `json:"addition"`
		CaseInsensitivity *bool                                                             `json:"case_insensitivity,omitempty"`
		IfNotExists       *bool                                                             `json:"if_not_exists,omitempty"`
		IfTagExists       *DataAttributesRulesItemsIfTagExists                              `json:"if_tag_exists,omitempty"`
		Query             *string                                                           `json:"query"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if !all.Addition.IsSet() {
		return fmt.Errorf("required field addition missing")
	}
	if all.Query == nil {
		return fmt.Errorf("required field query missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"addition", "case_insensitivity", "if_not_exists", "if_tag_exists", "query"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Addition = all.Addition
	o.CaseInsensitivity = all.CaseInsensitivity
	o.IfNotExists = all.IfNotExists
	if all.IfTagExists != nil && !all.IfTagExists.IsValid() {
		hasInvalidField = true
	} else {
		o.IfTagExists = all.IfTagExists
	}
	o.Query = *all.Query

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableUpdateRulesetRequestDataAttributesRulesItemsQuery handles when a null is used for UpdateRulesetRequestDataAttributesRulesItemsQuery.
type NullableUpdateRulesetRequestDataAttributesRulesItemsQuery struct {
	value *UpdateRulesetRequestDataAttributesRulesItemsQuery
	isSet bool
}

// Get returns the associated value.
func (v NullableUpdateRulesetRequestDataAttributesRulesItemsQuery) Get() *UpdateRulesetRequestDataAttributesRulesItemsQuery {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableUpdateRulesetRequestDataAttributesRulesItemsQuery) Set(val *UpdateRulesetRequestDataAttributesRulesItemsQuery) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableUpdateRulesetRequestDataAttributesRulesItemsQuery) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableUpdateRulesetRequestDataAttributesRulesItemsQuery) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUpdateRulesetRequestDataAttributesRulesItemsQuery initializes the struct as if Set has been called.
func NewNullableUpdateRulesetRequestDataAttributesRulesItemsQuery(val *UpdateRulesetRequestDataAttributesRulesItemsQuery) *NullableUpdateRulesetRequestDataAttributesRulesItemsQuery {
	return &NullableUpdateRulesetRequestDataAttributesRulesItemsQuery{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableUpdateRulesetRequestDataAttributesRulesItemsQuery) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableUpdateRulesetRequestDataAttributesRulesItemsQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
