// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleOrderAttributes Attributes for the reorder operation.
type TagIndexingRuleOrderAttributes struct {
	// Ordered list of tag indexing rule UUIDs. The server assigns rule_order 1, 2, … matching position in this list.
	RuleIds []string `json:"rule_ids,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagIndexingRuleOrderAttributes instantiates a new TagIndexingRuleOrderAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagIndexingRuleOrderAttributes() *TagIndexingRuleOrderAttributes {
	this := TagIndexingRuleOrderAttributes{}
	return &this
}

// NewTagIndexingRuleOrderAttributesWithDefaults instantiates a new TagIndexingRuleOrderAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagIndexingRuleOrderAttributesWithDefaults() *TagIndexingRuleOrderAttributes {
	this := TagIndexingRuleOrderAttributes{}
	return &this
}

// GetRuleIds returns the RuleIds field value if set, zero value otherwise.
func (o *TagIndexingRuleOrderAttributes) GetRuleIds() []string {
	if o == nil || o.RuleIds == nil {
		var ret []string
		return ret
	}
	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleOrderAttributes) GetRuleIdsOk() (*[]string, bool) {
	if o == nil || o.RuleIds == nil {
		return nil, false
	}
	return &o.RuleIds, true
}

// HasRuleIds returns a boolean if a field has been set.
func (o *TagIndexingRuleOrderAttributes) HasRuleIds() bool {
	return o != nil && o.RuleIds != nil
}

// SetRuleIds gets a reference to the given []string and assigns it to the RuleIds field.
func (o *TagIndexingRuleOrderAttributes) SetRuleIds(v []string) {
	o.RuleIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagIndexingRuleOrderAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RuleIds != nil {
		toSerialize["rule_ids"] = o.RuleIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagIndexingRuleOrderAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RuleIds []string `json:"rule_ids,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rule_ids"})
	} else {
		return err
	}
	o.RuleIds = all.RuleIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
