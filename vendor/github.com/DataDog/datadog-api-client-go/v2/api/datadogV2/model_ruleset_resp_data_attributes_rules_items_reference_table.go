// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RulesetRespDataAttributesRulesItemsReferenceTable The definition of `RulesetRespDataAttributesRulesItemsReferenceTable` object.
type RulesetRespDataAttributesRulesItemsReferenceTable struct {
	// The `reference_table` `case_insensitivity`.
	CaseInsensitivity *bool `json:"case_insensitivity,omitempty"`
	// The `reference_table` `field_pairs`.
	FieldPairs []RulesetRespDataAttributesRulesItemsReferenceTableFieldPairsItems `json:"field_pairs"`
	// The `reference_table` `if_not_exists`.
	IfNotExists *bool `json:"if_not_exists,omitempty"`
	// The `reference_table` `source_keys`.
	SourceKeys []string `json:"source_keys"`
	// The `reference_table` `table_name`.
	TableName string `json:"table_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRulesetRespDataAttributesRulesItemsReferenceTable instantiates a new RulesetRespDataAttributesRulesItemsReferenceTable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRulesetRespDataAttributesRulesItemsReferenceTable(fieldPairs []RulesetRespDataAttributesRulesItemsReferenceTableFieldPairsItems, sourceKeys []string, tableName string) *RulesetRespDataAttributesRulesItemsReferenceTable {
	this := RulesetRespDataAttributesRulesItemsReferenceTable{}
	this.FieldPairs = fieldPairs
	this.SourceKeys = sourceKeys
	this.TableName = tableName
	return &this
}

// NewRulesetRespDataAttributesRulesItemsReferenceTableWithDefaults instantiates a new RulesetRespDataAttributesRulesItemsReferenceTable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRulesetRespDataAttributesRulesItemsReferenceTableWithDefaults() *RulesetRespDataAttributesRulesItemsReferenceTable {
	this := RulesetRespDataAttributesRulesItemsReferenceTable{}
	return &this
}

// GetCaseInsensitivity returns the CaseInsensitivity field value if set, zero value otherwise.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) GetCaseInsensitivity() bool {
	if o == nil || o.CaseInsensitivity == nil {
		var ret bool
		return ret
	}
	return *o.CaseInsensitivity
}

// GetCaseInsensitivityOk returns a tuple with the CaseInsensitivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) GetCaseInsensitivityOk() (*bool, bool) {
	if o == nil || o.CaseInsensitivity == nil {
		return nil, false
	}
	return o.CaseInsensitivity, true
}

// HasCaseInsensitivity returns a boolean if a field has been set.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) HasCaseInsensitivity() bool {
	return o != nil && o.CaseInsensitivity != nil
}

// SetCaseInsensitivity gets a reference to the given bool and assigns it to the CaseInsensitivity field.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) SetCaseInsensitivity(v bool) {
	o.CaseInsensitivity = &v
}

// GetFieldPairs returns the FieldPairs field value.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) GetFieldPairs() []RulesetRespDataAttributesRulesItemsReferenceTableFieldPairsItems {
	if o == nil {
		var ret []RulesetRespDataAttributesRulesItemsReferenceTableFieldPairsItems
		return ret
	}
	return o.FieldPairs
}

// GetFieldPairsOk returns a tuple with the FieldPairs field value
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) GetFieldPairsOk() (*[]RulesetRespDataAttributesRulesItemsReferenceTableFieldPairsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldPairs, true
}

// SetFieldPairs sets field value.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) SetFieldPairs(v []RulesetRespDataAttributesRulesItemsReferenceTableFieldPairsItems) {
	o.FieldPairs = v
}

// GetIfNotExists returns the IfNotExists field value if set, zero value otherwise.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) GetIfNotExists() bool {
	if o == nil || o.IfNotExists == nil {
		var ret bool
		return ret
	}
	return *o.IfNotExists
}

// GetIfNotExistsOk returns a tuple with the IfNotExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) GetIfNotExistsOk() (*bool, bool) {
	if o == nil || o.IfNotExists == nil {
		return nil, false
	}
	return o.IfNotExists, true
}

// HasIfNotExists returns a boolean if a field has been set.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) HasIfNotExists() bool {
	return o != nil && o.IfNotExists != nil
}

// SetIfNotExists gets a reference to the given bool and assigns it to the IfNotExists field.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) SetIfNotExists(v bool) {
	o.IfNotExists = &v
}

// GetSourceKeys returns the SourceKeys field value.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) GetSourceKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SourceKeys
}

// GetSourceKeysOk returns a tuple with the SourceKeys field value
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) GetSourceKeysOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceKeys, true
}

// SetSourceKeys sets field value.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) SetSourceKeys(v []string) {
	o.SourceKeys = v
}

// GetTableName returns the TableName field value.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) SetTableName(v string) {
	o.TableName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RulesetRespDataAttributesRulesItemsReferenceTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CaseInsensitivity != nil {
		toSerialize["case_insensitivity"] = o.CaseInsensitivity
	}
	toSerialize["field_pairs"] = o.FieldPairs
	if o.IfNotExists != nil {
		toSerialize["if_not_exists"] = o.IfNotExists
	}
	toSerialize["source_keys"] = o.SourceKeys
	toSerialize["table_name"] = o.TableName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RulesetRespDataAttributesRulesItemsReferenceTable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CaseInsensitivity *bool                                                               `json:"case_insensitivity,omitempty"`
		FieldPairs        *[]RulesetRespDataAttributesRulesItemsReferenceTableFieldPairsItems `json:"field_pairs"`
		IfNotExists       *bool                                                               `json:"if_not_exists,omitempty"`
		SourceKeys        *[]string                                                           `json:"source_keys"`
		TableName         *string                                                             `json:"table_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FieldPairs == nil {
		return fmt.Errorf("required field field_pairs missing")
	}
	if all.SourceKeys == nil {
		return fmt.Errorf("required field source_keys missing")
	}
	if all.TableName == nil {
		return fmt.Errorf("required field table_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"case_insensitivity", "field_pairs", "if_not_exists", "source_keys", "table_name"})
	} else {
		return err
	}
	o.CaseInsensitivity = all.CaseInsensitivity
	o.FieldPairs = *all.FieldPairs
	o.IfNotExists = all.IfNotExists
	o.SourceKeys = *all.SourceKeys
	o.TableName = *all.TableName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableRulesetRespDataAttributesRulesItemsReferenceTable handles when a null is used for RulesetRespDataAttributesRulesItemsReferenceTable.
type NullableRulesetRespDataAttributesRulesItemsReferenceTable struct {
	value *RulesetRespDataAttributesRulesItemsReferenceTable
	isSet bool
}

// Get returns the associated value.
func (v NullableRulesetRespDataAttributesRulesItemsReferenceTable) Get() *RulesetRespDataAttributesRulesItemsReferenceTable {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableRulesetRespDataAttributesRulesItemsReferenceTable) Set(val *RulesetRespDataAttributesRulesItemsReferenceTable) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableRulesetRespDataAttributesRulesItemsReferenceTable) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableRulesetRespDataAttributesRulesItemsReferenceTable) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRulesetRespDataAttributesRulesItemsReferenceTable initializes the struct as if Set has been called.
func NewNullableRulesetRespDataAttributesRulesItemsReferenceTable(val *RulesetRespDataAttributesRulesItemsReferenceTable) *NullableRulesetRespDataAttributesRulesItemsReferenceTable {
	return &NullableRulesetRespDataAttributesRulesItemsReferenceTable{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableRulesetRespDataAttributesRulesItemsReferenceTable) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableRulesetRespDataAttributesRulesItemsReferenceTable) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
