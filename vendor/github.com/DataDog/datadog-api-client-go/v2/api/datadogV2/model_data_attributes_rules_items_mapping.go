// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataAttributesRulesItemsMapping The definition of `DataAttributesRulesItemsMapping` object.
type DataAttributesRulesItemsMapping struct {
	// The `mapping` `destination_key`.
	DestinationKey string `json:"destination_key"`
	// Deprecated. Use `if_tag_exists` instead. The `mapping` `if_not_exists`.
	// Deprecated
	IfNotExists *bool `json:"if_not_exists,omitempty"`
	// The behavior when the tag already exists.
	IfTagExists *DataAttributesRulesItemsIfTagExists `json:"if_tag_exists,omitempty"`
	// The `mapping` `source_keys`.
	SourceKeys []string `json:"source_keys"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataAttributesRulesItemsMapping instantiates a new DataAttributesRulesItemsMapping object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataAttributesRulesItemsMapping(destinationKey string, sourceKeys []string) *DataAttributesRulesItemsMapping {
	this := DataAttributesRulesItemsMapping{}
	this.DestinationKey = destinationKey
	this.SourceKeys = sourceKeys
	return &this
}

// NewDataAttributesRulesItemsMappingWithDefaults instantiates a new DataAttributesRulesItemsMapping object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataAttributesRulesItemsMappingWithDefaults() *DataAttributesRulesItemsMapping {
	this := DataAttributesRulesItemsMapping{}
	return &this
}

// GetDestinationKey returns the DestinationKey field value.
func (o *DataAttributesRulesItemsMapping) GetDestinationKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DestinationKey
}

// GetDestinationKeyOk returns a tuple with the DestinationKey field value
// and a boolean to check if the value has been set.
func (o *DataAttributesRulesItemsMapping) GetDestinationKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationKey, true
}

// SetDestinationKey sets field value.
func (o *DataAttributesRulesItemsMapping) SetDestinationKey(v string) {
	o.DestinationKey = v
}

// GetIfNotExists returns the IfNotExists field value if set, zero value otherwise.
// Deprecated
func (o *DataAttributesRulesItemsMapping) GetIfNotExists() bool {
	if o == nil || o.IfNotExists == nil {
		var ret bool
		return ret
	}
	return *o.IfNotExists
}

// GetIfNotExistsOk returns a tuple with the IfNotExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DataAttributesRulesItemsMapping) GetIfNotExistsOk() (*bool, bool) {
	if o == nil || o.IfNotExists == nil {
		return nil, false
	}
	return o.IfNotExists, true
}

// HasIfNotExists returns a boolean if a field has been set.
func (o *DataAttributesRulesItemsMapping) HasIfNotExists() bool {
	return o != nil && o.IfNotExists != nil
}

// SetIfNotExists gets a reference to the given bool and assigns it to the IfNotExists field.
// Deprecated
func (o *DataAttributesRulesItemsMapping) SetIfNotExists(v bool) {
	o.IfNotExists = &v
}

// GetIfTagExists returns the IfTagExists field value if set, zero value otherwise.
func (o *DataAttributesRulesItemsMapping) GetIfTagExists() DataAttributesRulesItemsIfTagExists {
	if o == nil || o.IfTagExists == nil {
		var ret DataAttributesRulesItemsIfTagExists
		return ret
	}
	return *o.IfTagExists
}

// GetIfTagExistsOk returns a tuple with the IfTagExists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataAttributesRulesItemsMapping) GetIfTagExistsOk() (*DataAttributesRulesItemsIfTagExists, bool) {
	if o == nil || o.IfTagExists == nil {
		return nil, false
	}
	return o.IfTagExists, true
}

// HasIfTagExists returns a boolean if a field has been set.
func (o *DataAttributesRulesItemsMapping) HasIfTagExists() bool {
	return o != nil && o.IfTagExists != nil
}

// SetIfTagExists gets a reference to the given DataAttributesRulesItemsIfTagExists and assigns it to the IfTagExists field.
func (o *DataAttributesRulesItemsMapping) SetIfTagExists(v DataAttributesRulesItemsIfTagExists) {
	o.IfTagExists = &v
}

// GetSourceKeys returns the SourceKeys field value.
func (o *DataAttributesRulesItemsMapping) GetSourceKeys() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SourceKeys
}

// GetSourceKeysOk returns a tuple with the SourceKeys field value
// and a boolean to check if the value has been set.
func (o *DataAttributesRulesItemsMapping) GetSourceKeysOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceKeys, true
}

// SetSourceKeys sets field value.
func (o *DataAttributesRulesItemsMapping) SetSourceKeys(v []string) {
	o.SourceKeys = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataAttributesRulesItemsMapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["destination_key"] = o.DestinationKey
	if o.IfNotExists != nil {
		toSerialize["if_not_exists"] = o.IfNotExists
	}
	if o.IfTagExists != nil {
		toSerialize["if_tag_exists"] = o.IfTagExists
	}
	toSerialize["source_keys"] = o.SourceKeys

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataAttributesRulesItemsMapping) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DestinationKey *string                              `json:"destination_key"`
		IfNotExists    *bool                                `json:"if_not_exists,omitempty"`
		IfTagExists    *DataAttributesRulesItemsIfTagExists `json:"if_tag_exists,omitempty"`
		SourceKeys     *[]string                            `json:"source_keys"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DestinationKey == nil {
		return fmt.Errorf("required field destination_key missing")
	}
	if all.SourceKeys == nil {
		return fmt.Errorf("required field source_keys missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"destination_key", "if_not_exists", "if_tag_exists", "source_keys"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DestinationKey = *all.DestinationKey
	o.IfNotExists = all.IfNotExists
	if all.IfTagExists != nil && !all.IfTagExists.IsValid() {
		hasInvalidField = true
	} else {
		o.IfTagExists = all.IfTagExists
	}
	o.SourceKeys = *all.SourceKeys

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}

// NullableDataAttributesRulesItemsMapping handles when a null is used for DataAttributesRulesItemsMapping.
type NullableDataAttributesRulesItemsMapping struct {
	value *DataAttributesRulesItemsMapping
	isSet bool
}

// Get returns the associated value.
func (v NullableDataAttributesRulesItemsMapping) Get() *DataAttributesRulesItemsMapping {
	return v.value
}

// Set changes the value and indicates it's been called.
func (v *NullableDataAttributesRulesItemsMapping) Set(val *DataAttributesRulesItemsMapping) {
	v.value = val
	v.isSet = true
}

// IsSet returns whether Set has been called.
func (v NullableDataAttributesRulesItemsMapping) IsSet() bool {
	return v.isSet
}

// Unset sets the value to nil and resets the set flag/
func (v *NullableDataAttributesRulesItemsMapping) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDataAttributesRulesItemsMapping initializes the struct as if Set has been called.
func NewNullableDataAttributesRulesItemsMapping(val *DataAttributesRulesItemsMapping) *NullableDataAttributesRulesItemsMapping {
	return &NullableDataAttributesRulesItemsMapping{value: val, isSet: true}
}

// MarshalJSON serializes the associated value.
func (v NullableDataAttributesRulesItemsMapping) MarshalJSON() ([]byte, error) {
	return datadog.Marshal(v.value)
}

// UnmarshalJSON deserializes the payload and sets the flag as if Set has been called.
func (v *NullableDataAttributesRulesItemsMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true

	// this object is nullable so check if the payload is null or empty string
	if string(src) == "" || string(src) == "{}" {
		return nil
	}

	return datadog.Unmarshal(src, &v.value)
}
