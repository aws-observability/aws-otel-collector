// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AiCustomRulesetRequestAttributes Attributes for creating an AI custom ruleset.
type AiCustomRulesetRequestAttributes struct {
	// Base64-encoded full description of the ruleset.
	Description string `json:"description"`
	// The ruleset name.
	Name string `json:"name"`
	// Base64-encoded short description of the ruleset.
	ShortDescription string `json:"short_description"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAiCustomRulesetRequestAttributes instantiates a new AiCustomRulesetRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAiCustomRulesetRequestAttributes(description string, name string, shortDescription string) *AiCustomRulesetRequestAttributes {
	this := AiCustomRulesetRequestAttributes{}
	this.Description = description
	this.Name = name
	this.ShortDescription = shortDescription
	return &this
}

// NewAiCustomRulesetRequestAttributesWithDefaults instantiates a new AiCustomRulesetRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAiCustomRulesetRequestAttributesWithDefaults() *AiCustomRulesetRequestAttributes {
	this := AiCustomRulesetRequestAttributes{}
	return &this
}

// GetDescription returns the Description field value.
func (o *AiCustomRulesetRequestAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AiCustomRulesetRequestAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *AiCustomRulesetRequestAttributes) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value.
func (o *AiCustomRulesetRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AiCustomRulesetRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *AiCustomRulesetRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetShortDescription returns the ShortDescription field value.
func (o *AiCustomRulesetRequestAttributes) GetShortDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value
// and a boolean to check if the value has been set.
func (o *AiCustomRulesetRequestAttributes) GetShortDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortDescription, true
}

// SetShortDescription sets field value.
func (o *AiCustomRulesetRequestAttributes) SetShortDescription(v string) {
	o.ShortDescription = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AiCustomRulesetRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["short_description"] = o.ShortDescription

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AiCustomRulesetRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description      *string `json:"description"`
		Name             *string `json:"name"`
		ShortDescription *string `json:"short_description"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ShortDescription == nil {
		return fmt.Errorf("required field short_description missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "name", "short_description"})
	} else {
		return err
	}
	o.Description = *all.Description
	o.Name = *all.Name
	o.ShortDescription = *all.ShortDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
