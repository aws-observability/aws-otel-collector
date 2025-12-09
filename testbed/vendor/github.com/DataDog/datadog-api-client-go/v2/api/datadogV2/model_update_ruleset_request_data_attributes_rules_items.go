// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateRulesetRequestDataAttributesRulesItems The definition of `UpdateRulesetRequestDataAttributesRulesItems` object.
type UpdateRulesetRequestDataAttributesRulesItems struct {
	// The `items` `enabled`.
	Enabled bool `json:"enabled"`
	// The definition of `UpdateRulesetRequestDataAttributesRulesItemsMapping` object.
	Mapping NullableUpdateRulesetRequestDataAttributesRulesItemsMapping `json:"mapping,omitempty"`
	// The `items` `metadata`.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The `items` `name`.
	Name string `json:"name"`
	// The definition of `UpdateRulesetRequestDataAttributesRulesItemsQuery` object.
	Query NullableUpdateRulesetRequestDataAttributesRulesItemsQuery `json:"query,omitempty"`
	// The definition of `UpdateRulesetRequestDataAttributesRulesItemsReferenceTable` object.
	ReferenceTable NullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable `json:"reference_table,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateRulesetRequestDataAttributesRulesItems instantiates a new UpdateRulesetRequestDataAttributesRulesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateRulesetRequestDataAttributesRulesItems(enabled bool, name string) *UpdateRulesetRequestDataAttributesRulesItems {
	this := UpdateRulesetRequestDataAttributesRulesItems{}
	this.Enabled = enabled
	this.Name = name
	return &this
}

// NewUpdateRulesetRequestDataAttributesRulesItemsWithDefaults instantiates a new UpdateRulesetRequestDataAttributesRulesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateRulesetRequestDataAttributesRulesItemsWithDefaults() *UpdateRulesetRequestDataAttributesRulesItems {
	this := UpdateRulesetRequestDataAttributesRulesItems{}
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *UpdateRulesetRequestDataAttributesRulesItems) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItems) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *UpdateRulesetRequestDataAttributesRulesItems) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMapping returns the Mapping field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRulesetRequestDataAttributesRulesItems) GetMapping() UpdateRulesetRequestDataAttributesRulesItemsMapping {
	if o == nil || o.Mapping.Get() == nil {
		var ret UpdateRulesetRequestDataAttributesRulesItemsMapping
		return ret
	}
	return *o.Mapping.Get()
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UpdateRulesetRequestDataAttributesRulesItems) GetMappingOk() (*UpdateRulesetRequestDataAttributesRulesItemsMapping, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mapping.Get(), o.Mapping.IsSet()
}

// HasMapping returns a boolean if a field has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItems) HasMapping() bool {
	return o != nil && o.Mapping.IsSet()
}

// SetMapping gets a reference to the given NullableUpdateRulesetRequestDataAttributesRulesItemsMapping and assigns it to the Mapping field.
func (o *UpdateRulesetRequestDataAttributesRulesItems) SetMapping(v UpdateRulesetRequestDataAttributesRulesItemsMapping) {
	o.Mapping.Set(&v)
}

// SetMappingNil sets the value for Mapping to be an explicit nil.
func (o *UpdateRulesetRequestDataAttributesRulesItems) SetMappingNil() {
	o.Mapping.Set(nil)
}

// UnsetMapping ensures that no value is present for Mapping, not even an explicit nil.
func (o *UpdateRulesetRequestDataAttributesRulesItems) UnsetMapping() {
	o.Mapping.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRulesetRequestDataAttributesRulesItems) GetMetadata() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UpdateRulesetRequestDataAttributesRulesItems) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItems) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UpdateRulesetRequestDataAttributesRulesItems) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetName returns the Name field value.
func (o *UpdateRulesetRequestDataAttributesRulesItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *UpdateRulesetRequestDataAttributesRulesItems) SetName(v string) {
	o.Name = v
}

// GetQuery returns the Query field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRulesetRequestDataAttributesRulesItems) GetQuery() UpdateRulesetRequestDataAttributesRulesItemsQuery {
	if o == nil || o.Query.Get() == nil {
		var ret UpdateRulesetRequestDataAttributesRulesItemsQuery
		return ret
	}
	return *o.Query.Get()
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UpdateRulesetRequestDataAttributesRulesItems) GetQueryOk() (*UpdateRulesetRequestDataAttributesRulesItemsQuery, bool) {
	if o == nil {
		return nil, false
	}
	return o.Query.Get(), o.Query.IsSet()
}

// HasQuery returns a boolean if a field has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItems) HasQuery() bool {
	return o != nil && o.Query.IsSet()
}

// SetQuery gets a reference to the given NullableUpdateRulesetRequestDataAttributesRulesItemsQuery and assigns it to the Query field.
func (o *UpdateRulesetRequestDataAttributesRulesItems) SetQuery(v UpdateRulesetRequestDataAttributesRulesItemsQuery) {
	o.Query.Set(&v)
}

// SetQueryNil sets the value for Query to be an explicit nil.
func (o *UpdateRulesetRequestDataAttributesRulesItems) SetQueryNil() {
	o.Query.Set(nil)
}

// UnsetQuery ensures that no value is present for Query, not even an explicit nil.
func (o *UpdateRulesetRequestDataAttributesRulesItems) UnsetQuery() {
	o.Query.Unset()
}

// GetReferenceTable returns the ReferenceTable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateRulesetRequestDataAttributesRulesItems) GetReferenceTable() UpdateRulesetRequestDataAttributesRulesItemsReferenceTable {
	if o == nil || o.ReferenceTable.Get() == nil {
		var ret UpdateRulesetRequestDataAttributesRulesItemsReferenceTable
		return ret
	}
	return *o.ReferenceTable.Get()
}

// GetReferenceTableOk returns a tuple with the ReferenceTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *UpdateRulesetRequestDataAttributesRulesItems) GetReferenceTableOk() (*UpdateRulesetRequestDataAttributesRulesItemsReferenceTable, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReferenceTable.Get(), o.ReferenceTable.IsSet()
}

// HasReferenceTable returns a boolean if a field has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItems) HasReferenceTable() bool {
	return o != nil && o.ReferenceTable.IsSet()
}

// SetReferenceTable gets a reference to the given NullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable and assigns it to the ReferenceTable field.
func (o *UpdateRulesetRequestDataAttributesRulesItems) SetReferenceTable(v UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) {
	o.ReferenceTable.Set(&v)
}

// SetReferenceTableNil sets the value for ReferenceTable to be an explicit nil.
func (o *UpdateRulesetRequestDataAttributesRulesItems) SetReferenceTableNil() {
	o.ReferenceTable.Set(nil)
}

// UnsetReferenceTable ensures that no value is present for ReferenceTable, not even an explicit nil.
func (o *UpdateRulesetRequestDataAttributesRulesItems) UnsetReferenceTable() {
	o.ReferenceTable.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateRulesetRequestDataAttributesRulesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	if o.Mapping.IsSet() {
		toSerialize["mapping"] = o.Mapping.Get()
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	if o.Query.IsSet() {
		toSerialize["query"] = o.Query.Get()
	}
	if o.ReferenceTable.IsSet() {
		toSerialize["reference_table"] = o.ReferenceTable.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateRulesetRequestDataAttributesRulesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled        *bool                                                              `json:"enabled"`
		Mapping        NullableUpdateRulesetRequestDataAttributesRulesItemsMapping        `json:"mapping,omitempty"`
		Metadata       map[string]string                                                  `json:"metadata,omitempty"`
		Name           *string                                                            `json:"name"`
		Query          NullableUpdateRulesetRequestDataAttributesRulesItemsQuery          `json:"query,omitempty"`
		ReferenceTable NullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable `json:"reference_table,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "mapping", "metadata", "name", "query", "reference_table"})
	} else {
		return err
	}
	o.Enabled = *all.Enabled
	o.Mapping = all.Mapping
	o.Metadata = all.Metadata
	o.Name = *all.Name
	o.Query = all.Query
	o.ReferenceTable = all.ReferenceTable

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
