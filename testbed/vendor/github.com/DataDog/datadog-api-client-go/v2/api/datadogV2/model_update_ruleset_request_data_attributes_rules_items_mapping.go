// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UpdateRulesetRequestDataAttributesRulesItemsMapping The definition of `UpdateRulesetRequestDataAttributesRulesItemsMapping` object.
type UpdateRulesetRequestDataAttributesRulesItemsMapping struct {
	// The `mapping` `destination_key`.
	DestinationKey string `json:"destination_key"`
	// The `mapping` `if_not_exists`.
	IfNotExists bool `json:"if_not_exists"`
	// The `mapping` `source_keys`.
	SourceKeys []string `json:"source_keys"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUpdateRulesetRequestDataAttributesRulesItemsMapping instantiates a new UpdateRulesetRequestDataAttributesRulesItemsMapping object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUpdateRulesetRequestDataAttributesRulesItemsMapping(destinationKey string, ifNotExists bool, sourceKeys []string) *UpdateRulesetRequestDataAttributesRulesItemsMapping {
	this := UpdateRulesetRequestDataAttributesRulesItemsMapping{}
	this.DestinationKey = destinationKey
	this.IfNotExists = ifNotExists
	this.SourceKeys = sourceKeys
	return &this
}

// NewUpdateRulesetRequestDataAttributesRulesItemsMappingWithDefaults instantiates a new UpdateRulesetRequestDataAttributesRulesItemsMapping object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUpdateRulesetRequestDataAttributesRulesItemsMappingWithDefaults() *UpdateRulesetRequestDataAttributesRulesItemsMapping {
	this := UpdateRulesetRequestDataAttributesRulesItemsMapping{}
	return &this
}

// GetDestinationKey returns the DestinationKey field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsMapping) GetDestinationKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DestinationKey
}

// GetDestinationKeyOk returns a tuple with the DestinationKey field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsMapping) GetDestinationKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationKey, true
}

// SetDestinationKey sets field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsMapping) SetDestinationKey(v string) {
	o.DestinationKey = v
}

// GetIfNotExists returns the IfNotExists field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsMapping) GetIfNotExists() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IfNotExists
}

// GetIfNotExistsOk returns a tuple with the IfNotExists field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsMapping) GetIfNotExistsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IfNotExists, true
}

// SetIfNotExists sets field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsMapping) SetIfNotExists(v bool) {
	o.IfNotExists = v
}

// GetSourceKeys returns the SourceKeys field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsMapping) GetSourceKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SourceKeys
}

// GetSourceKeysOk returns a tuple with the SourceKeys field value
// and a boolean to check if the value has been set.
func (o *UpdateRulesetRequestDataAttributesRulesItemsMapping) GetSourceKeysOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceKeys, true
}

// SetSourceKeys sets field value.
func (o *UpdateRulesetRequestDataAttributesRulesItemsMapping) SetSourceKeys(v []string) {
	o.SourceKeys = v
}

// MarshalJSON serializes the struct using spec logic.
func (o UpdateRulesetRequestDataAttributesRulesItemsMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["destination_key"] = o.DestinationKey
	toSerialize["if_not_exists"] = o.IfNotExists
	toSerialize["source_keys"] = o.SourceKeys

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UpdateRulesetRequestDataAttributesRulesItemsMapping) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DestinationKey *string   `json:"destination_key"`
		IfNotExists    *bool     `json:"if_not_exists"`
		SourceKeys     *[]string `json:"source_keys"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DestinationKey == nil {
		return fmt.Errorf("required field destination_key missing")
	}
	if all.IfNotExists == nil {
		return fmt.Errorf("required field if_not_exists missing")
	}
	if all.SourceKeys == nil {
		return fmt.Errorf("required field source_keys missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"destination_key", "if_not_exists", "source_keys"})
	} else {
		return err
	}
	o.DestinationKey = *all.DestinationKey
	o.IfNotExists = *all.IfNotExists
	o.SourceKeys = *all.SourceKeys

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}

// NullableUpdateRulesetRequestDataAttributesRulesItemsMapping handles when a null is used for UpdateRulesetRequestDataAttributesRulesItemsMapping.
type NullableUpdateRulesetRequestDataAttributesRulesItemsMapping struct {
	value *UpdateRulesetRequestDataAttributesRulesItemsMapping
	isSet bool
}

// Get returns the associated value.
func (v NullableUpdateRulesetRequestDataAttributesRulesItemsMapping) Get() *UpdateRulesetRequestDataAttributesRulesItemsMapping {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableUpdateRulesetRequestDataAttributesRulesItemsMapping) Set(val *UpdateRulesetRequestDataAttributesRulesItemsMapping) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableUpdateRulesetRequestDataAttributesRulesItemsMapping) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableUpdateRulesetRequestDataAttributesRulesItemsMapping) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUpdateRulesetRequestDataAttributesRulesItemsMapping initializes the struct as if Set has been called.
func NewNullableUpdateRulesetRequestDataAttributesRulesItemsMapping(val *UpdateRulesetRequestDataAttributesRulesItemsMapping) *NullableUpdateRulesetRequestDataAttributesRulesItemsMapping {
	return &NullableUpdateRulesetRequestDataAttributesRulesItemsMapping{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableUpdateRulesetRequestDataAttributesRulesItemsMapping) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableUpdateRulesetRequestDataAttributesRulesItemsMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
