// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateRulesetRequestDataAttributesRulesItemsReferenceTable The definition of `CreateRulesetRequestDataAttributesRulesItemsReferenceTable` object.
type CreateRulesetRequestDataAttributesRulesItemsReferenceTable struct {
	// The `reference_table` `case_insensitivity`.
	CaseInsensitivity *bool `json:"case_insensitivity,omitempty"`
	// The `reference_table` `field_pairs`.
	FieldPairs []CreateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems `json:"field_pairs"`
	// Deprecated. Use `if_tag_exists` instead. The `reference_table` `if_not_exists`.
	// Deprecated
	IfNotExists *bool `json:"if_not_exists,omitempty"`
	// The behavior when the tag already exists.
	IfTagExists *DataAttributesRulesItemsIfTagExists `json:"if_tag_exists,omitempty"`
	// The `reference_table` `source_keys`.
	SourceKeys []string `json:"source_keys"`
	// The `reference_table` `table_name`.
	TableName string `json:"table_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateRulesetRequestDataAttributesRulesItemsReferenceTable instantiates a new CreateRulesetRequestDataAttributesRulesItemsReferenceTable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateRulesetRequestDataAttributesRulesItemsReferenceTable(fieldPairs []CreateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems, sourceKeys []string, tableName string) *CreateRulesetRequestDataAttributesRulesItemsReferenceTable {
	this := CreateRulesetRequestDataAttributesRulesItemsReferenceTable{}
	this.FieldPairs = fieldPairs
	this.SourceKeys = sourceKeys
	this.TableName = tableName
	return &this
}

// NewCreateRulesetRequestDataAttributesRulesItemsReferenceTableWithDefaults instantiates a new CreateRulesetRequestDataAttributesRulesItemsReferenceTable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateRulesetRequestDataAttributesRulesItemsReferenceTableWithDefaults() *CreateRulesetRequestDataAttributesRulesItemsReferenceTable {
	this := CreateRulesetRequestDataAttributesRulesItemsReferenceTable{}
	return &this
}

// GetCaseInsensitivity returns the CaseInsensitivity field value if set, zero value otherwise.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) GetCaseInsensitivity() bool {
	if o == nil || o.CaseInsensitivity == nil {
		var ret bool
		return ret
	}
	return *o.CaseInsensitivity
}

// GetCaseInsensitivityOk returns a tuple with the CaseInsensitivity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) GetCaseInsensitivityOk() (*bool, bool) {
	if o == nil || o.CaseInsensitivity == nil {
		return nil, false
	}
	return o.CaseInsensitivity, true
}

// HasCaseInsensitivity returns a boolean if a field has been set.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) HasCaseInsensitivity() bool {
	return o != nil && o.CaseInsensitivity != nil
}

// SetCaseInsensitivity gets a reference to the given bool and assigns it to the CaseInsensitivity field.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) SetCaseInsensitivity(v bool) {
	o.CaseInsensitivity = &v
}

// GetFieldPairs returns the FieldPairs field value.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) GetFieldPairs() []CreateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems {
	if o == nil {
		var ret []CreateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems
		return ret
	}
	return o.FieldPairs
}

// GetFieldPairsOk returns a tuple with the FieldPairs field value
// and a boolean to check if the value has been set.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) GetFieldPairsOk() (*[]CreateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldPairs, true
}

// SetFieldPairs sets field value.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) SetFieldPairs(v []CreateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems) {
	o.FieldPairs = v
}

// GetIfNotExists returns the IfNotExists field value if set, zero value otherwise.
// Deprecated
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) GetIfNotExists() bool {
	if o == nil || o.IfNotExists == nil {
		var ret bool
		return ret
	}
	return *o.IfNotExists
}

// GetIfNotExistsOk returns a tuple with the IfNotExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) GetIfNotExistsOk() (*bool, bool) {
	if o == nil || o.IfNotExists == nil {
		return nil, false
	}
	return o.IfNotExists, true
}

// HasIfNotExists returns a boolean if a field has been set.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) HasIfNotExists() bool {
	return o != nil && o.IfNotExists != nil
}

// SetIfNotExists gets a reference to the given bool and assigns it to the IfNotExists field.
// Deprecated
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) SetIfNotExists(v bool) {
	o.IfNotExists = &v
}

// GetIfTagExists returns the IfTagExists field value if set, zero value otherwise.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) GetIfTagExists() DataAttributesRulesItemsIfTagExists {
	if o == nil || o.IfTagExists == nil {
		var ret DataAttributesRulesItemsIfTagExists
		return ret
	}
	return *o.IfTagExists
}

// GetIfTagExistsOk returns a tuple with the IfTagExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) GetIfTagExistsOk() (*DataAttributesRulesItemsIfTagExists, bool) {
	if o == nil || o.IfTagExists == nil {
		return nil, false
	}
	return o.IfTagExists, true
}

// HasIfTagExists returns a boolean if a field has been set.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) HasIfTagExists() bool {
	return o != nil && o.IfTagExists != nil
}

// SetIfTagExists gets a reference to the given DataAttributesRulesItemsIfTagExists and assigns it to the IfTagExists field.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) SetIfTagExists(v DataAttributesRulesItemsIfTagExists) {
	o.IfTagExists = &v
}

// GetSourceKeys returns the SourceKeys field value.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) GetSourceKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SourceKeys
}

// GetSourceKeysOk returns a tuple with the SourceKeys field value
// and a boolean to check if the value has been set.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) GetSourceKeysOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceKeys, true
}

// SetSourceKeys sets field value.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) SetSourceKeys(v []string) {
	o.SourceKeys = v
}

// GetTableName returns the TableName field value.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) SetTableName(v string) {
	o.TableName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateRulesetRequestDataAttributesRulesItemsReferenceTable) MarshalJSON() ([]byte, error) {
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
	if o.IfTagExists != nil {
		toSerialize["if_tag_exists"] = o.IfTagExists
	}
	toSerialize["source_keys"] = o.SourceKeys
	toSerialize["table_name"] = o.TableName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CaseInsensitivity *bool                                                                        `json:"case_insensitivity,omitempty"`
		FieldPairs        *[]CreateRulesetRequestDataAttributesRulesItemsReferenceTableFieldPairsItems `json:"field_pairs"`
		IfNotExists       *bool                                                                        `json:"if_not_exists,omitempty"`
		IfTagExists       *DataAttributesRulesItemsIfTagExists                                         `json:"if_tag_exists,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"case_insensitivity", "field_pairs", "if_not_exists", "if_tag_exists", "source_keys", "table_name"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CaseInsensitivity = all.CaseInsensitivity
	o.FieldPairs = *all.FieldPairs
	o.IfNotExists = all.IfNotExists
	if all.IfTagExists != nil && !all.IfTagExists.IsValid() {
		hasInvalidField = true
	} else {
		o.IfTagExists = all.IfTagExists
	}
	o.SourceKeys = *all.SourceKeys
	o.TableName = *all.TableName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableCreateRulesetRequestDataAttributesRulesItemsReferenceTable handles when a null is used for CreateRulesetRequestDataAttributesRulesItemsReferenceTable.
type NullableCreateRulesetRequestDataAttributesRulesItemsReferenceTable struct {
	value *CreateRulesetRequestDataAttributesRulesItemsReferenceTable
	isSet bool
}

// Get returns the associated value.
func (v NullableCreateRulesetRequestDataAttributesRulesItemsReferenceTable) Get() *CreateRulesetRequestDataAttributesRulesItemsReferenceTable {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableCreateRulesetRequestDataAttributesRulesItemsReferenceTable) Set(val *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableCreateRulesetRequestDataAttributesRulesItemsReferenceTable) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableCreateRulesetRequestDataAttributesRulesItemsReferenceTable) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCreateRulesetRequestDataAttributesRulesItemsReferenceTable initializes the struct as if Set has been called.
func NewNullableCreateRulesetRequestDataAttributesRulesItemsReferenceTable(val *CreateRulesetRequestDataAttributesRulesItemsReferenceTable) *NullableCreateRulesetRequestDataAttributesRulesItemsReferenceTable {
	return &NullableCreateRulesetRequestDataAttributesRulesItemsReferenceTable{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableCreateRulesetRequestDataAttributesRulesItemsReferenceTable) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableCreateRulesetRequestDataAttributesRulesItemsReferenceTable) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
