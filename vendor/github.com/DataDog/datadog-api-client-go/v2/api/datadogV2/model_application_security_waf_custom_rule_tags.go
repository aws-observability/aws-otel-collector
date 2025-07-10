// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleTags Tags associated with the WAF Custom Rule. The concatenation of category and type will form the security
// activity field associated with the traces.
type ApplicationSecurityWafCustomRuleTags struct {
	// The category of the WAF Rule, can be either `business_logic`, `attack_attempt` or `security_response`.
	Category ApplicationSecurityWafCustomRuleTagsCategory `json:"category"`
	// The type of the WAF rule, associated with the category will form the security activity.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]string      `json:"-"`
}

// NewApplicationSecurityWafCustomRuleTags instantiates a new ApplicationSecurityWafCustomRuleTags object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityWafCustomRuleTags(category ApplicationSecurityWafCustomRuleTagsCategory, typeVar string) *ApplicationSecurityWafCustomRuleTags {
	this := ApplicationSecurityWafCustomRuleTags{}
	this.Category = category
	this.Type = typeVar
	return &this
}

// NewApplicationSecurityWafCustomRuleTagsWithDefaults instantiates a new ApplicationSecurityWafCustomRuleTags object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityWafCustomRuleTagsWithDefaults() *ApplicationSecurityWafCustomRuleTags {
	this := ApplicationSecurityWafCustomRuleTags{}
	return &this
}

// GetCategory returns the Category field value.
func (o *ApplicationSecurityWafCustomRuleTags) GetCategory() ApplicationSecurityWafCustomRuleTagsCategory {
	if o == nil {
		var ret ApplicationSecurityWafCustomRuleTagsCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleTags) GetCategoryOk() (*ApplicationSecurityWafCustomRuleTagsCategory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value.
func (o *ApplicationSecurityWafCustomRuleTags) SetCategory(v ApplicationSecurityWafCustomRuleTagsCategory) {
	o.Category = v
}

// GetType returns the Type field value.
func (o *ApplicationSecurityWafCustomRuleTags) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleTags) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ApplicationSecurityWafCustomRuleTags) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityWafCustomRuleTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["category"] = o.Category
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityWafCustomRuleTags) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Category *ApplicationSecurityWafCustomRuleTagsCategory `json:"category"`
		Type     *string                                       `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Category == nil {
		return fmt.Errorf("required field category missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]string)
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"category", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Category.IsValid() {
		hasInvalidField = true
	} else {
		o.Category = *all.Category
	}
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
