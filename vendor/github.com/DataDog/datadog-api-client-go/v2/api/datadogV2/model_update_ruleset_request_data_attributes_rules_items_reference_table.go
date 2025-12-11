// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateRulesetRequestDataAttributesRulesItemsReferenceTable The definition of `UpdateRulesetRequestDataAttributesRulesItemsReferenceTable` object.
type UpdateRulesetRequestDataAttributesRulesItemsReferenceTable struct {
	// The `reference_table` `case_insensitivity`.
	CaseInsensitivity *bool `json:"case_insensitivity,omitempty"`
	// The `reference_table` `field_pairs`.
	FieldPairs []UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems `json:"field_pairs"`
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

// NewUpdateRulesetRequestDataAttributesRulesItemsReferenceTable instantiates a new UpdateRulesetRequestDataAttributesRulesItemsReferenceTable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateRulesetRequestDataAttributesRulesItemsReferenceTable(fieldPairs []UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems, sourceKeys []string, tableName string) *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable {
	this := UpdateRulesetRequestDataAttributesRulesItemsReferenceTable{}
	this.FieldPairs = fieldPairs
	this.SourceKeys = sourceKeys
	this.TableName = tableName
	return &this
}

// NewUpdateRulesetRequestDataAttributesRulesItemsReferenceTableWithDefaults instantiates a new UpdateRulesetRequestDataAttributesRulesItemsReferenceTable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateRulesetRequestDataAttributesRulesItemsReferenceTableWithDefaults() *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable {
	this := UpdateRulesetRequestDataAttributesRulesItemsReferenceTable{}
	return &this
}

// GetCaseInsensitivity returns the CaseInsensitivity field value if set, zero value otherwise.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) GetCaseInsensitivity() bool {
	if o == nil || o.CaseInsensitivity == nil {
		var ret bool
		return ret
	}
	return *o.CaseInsensitivity
}

// GetCaseInsensitivityOk returns a tuple with the CaseInsensitivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) GetCaseInsensitivityOk() (*bool, bool) {
	if o == nil || o.CaseInsensitivity == nil {
		return nil, false
	}
	return o.CaseInsensitivity, true
}

// HasCaseInsensitivity returns a boolean if a field has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) HasCaseInsensitivity() bool {
	return o != nil && o.CaseInsensitivity != nil
}

// SetCaseInsensitivity gets a reference to the given bool and assigns it to the CaseInsensitivity field.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) SetCaseInsensitivity(v bool) {
	o.CaseInsensitivity = &v
}

// GetFieldPairs returns the FieldPairs field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) GetFieldPairs() []UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems {
	if o == nil {
		var ret []UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems
		return ret
	}
	return o.FieldPairs
}

// GetFieldPairsOk returns a tuple with the FieldPairs field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) GetFieldPairsOk() (*[]UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldPairs, true
}

// SetFieldPairs sets field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) SetFieldPairs(v []UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems) {
	o.FieldPairs = v
}

// GetIfNotExists returns the IfNotExists field value if set, zero value otherwise.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) GetIfNotExists() bool {
	if o == nil || o.IfNotExists == nil {
		var ret bool
		return ret
	}
	return *o.IfNotExists
}

// GetIfNotExistsOk returns a tuple with the IfNotExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) GetIfNotExistsOk() (*bool, bool) {
	if o == nil || o.IfNotExists == nil {
		return nil, false
	}
	return o.IfNotExists, true
}

// HasIfNotExists returns a boolean if a field has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) HasIfNotExists() bool {
	return o != nil && o.IfNotExists != nil
}

// SetIfNotExists gets a reference to the given bool and assigns it to the IfNotExists field.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) SetIfNotExists(v bool) {
	o.IfNotExists = &v
}

// GetSourceKeys returns the SourceKeys field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) GetSourceKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SourceKeys
}

// GetSourceKeysOk returns a tuple with the SourceKeys field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) GetSourceKeysOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceKeys, true
}

// SetSourceKeys sets field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) SetSourceKeys(v []string) {
	o.SourceKeys = v
}

// GetTableName returns the TableName field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) SetTableName(v string) {
	o.TableName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) MarshalJSON() ([]byte, error) {
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
func (o *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CaseInsensitivity *bool                                                                        `json:"case_insensitivity,omitempty"`
		FieldPairs        *[]UpdateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems `json:"field_pairs"`
		IfNotExists       *bool                                                                        `json:"if_not_exists,omitempty"`
		SourceKeys        *[]string                                                                    `json:"source_keys"`
		TableName         *string                                                                      `json:"table_name"`
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

// NullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable handles when a null is used for UpdateRulesetRequestDataAttributesRulesItemsReferenceTable.
type NullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable struct {
	value *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable
	isSet bool
}

// Get returns the associated value.
func (v NullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable) Get() *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable) Set(val *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable initializes the struct as if Set has been called.
func NewNullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable(val *UpdateRulesetRequestDataAttributesRulesItemsReferenceTable) *NullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable {
	return &NullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableUpdateRulesetRequestDataAttributesRulesItemsReferenceTable) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
