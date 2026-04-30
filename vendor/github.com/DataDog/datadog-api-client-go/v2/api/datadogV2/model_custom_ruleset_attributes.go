// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomRulesetAttributes Attributes of a custom ruleset, including its name, description, and rules.
type CustomRulesetAttributes struct {
	// Creation timestamp
	CreatedAt time.Time `json:"created_at"`
	// Creator identifier
	CreatedBy string `json:"created_by"`
	// Base64-encoded full description
	Description string `json:"description"`
	// Ruleset name
	Name string `json:"name"`
	// Rules in the ruleset
	Rules datadog.NullableList[CustomRule] `json:"rules"`
	// Base64-encoded short description
	ShortDescription string `json:"short_description"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomRulesetAttributes instantiates a new CustomRulesetAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomRulesetAttributes(createdAt time.Time, createdBy string, description string, name string, rules datadog.NullableList[CustomRule], shortDescription string) *CustomRulesetAttributes {
	this := CustomRulesetAttributes{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Description = description
	this.Name = name
	this.Rules = rules
	this.ShortDescription = shortDescription
	return &this
}

// NewCustomRulesetAttributesWithDefaults instantiates a new CustomRulesetAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomRulesetAttributesWithDefaults() *CustomRulesetAttributes {
	this := CustomRulesetAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *CustomRulesetAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CustomRulesetAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *CustomRulesetAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *CustomRulesetAttributes) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *CustomRulesetAttributes) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *CustomRulesetAttributes) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetDescription returns the Description field value.
func (o *CustomRulesetAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CustomRulesetAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *CustomRulesetAttributes) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value.
func (o *CustomRulesetAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CustomRulesetAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CustomRulesetAttributes) SetName(v string) {
	o.Name = v
}

// GetRules returns the Rules field value.
// If the value is explicit nil, the zero value for []CustomRule will be returned.
func (o *CustomRulesetAttributes) GetRules() []CustomRule {
	if o == nil {
		var ret []CustomRule
		return ret
	}
	return *o.Rules.Get()
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomRulesetAttributes) GetRulesOk() (*[]CustomRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules.Get(), o.Rules.IsSet()
}

// SetRules sets field value.
func (o *CustomRulesetAttributes) SetRules(v []CustomRule) {
	o.Rules.Set(&v)
}

// GetShortDescription returns the ShortDescription field value.
func (o *CustomRulesetAttributes) GetShortDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ShortDescription
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value
// and a boolean to check if the value has been set.
func (o *CustomRulesetAttributes) GetShortDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortDescription, true
}

// SetShortDescription sets field value.
func (o *CustomRulesetAttributes) SetShortDescription(v string) {
	o.ShortDescription = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomRulesetAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["rules"] = o.Rules.Get()
	toSerialize["short_description"] = o.ShortDescription

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomRulesetAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt        *time.Time                       `json:"created_at"`
		CreatedBy        *string                          `json:"created_by"`
		Description      *string                          `json:"description"`
		Name             *string                          `json:"name"`
		Rules            datadog.NullableList[CustomRule] `json:"rules"`
		ShortDescription *string                          `json:"short_description"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if !all.Rules.IsSet() {
		return fmt.Errorf("required field rules missing")
	}
	if all.ShortDescription == nil {
		return fmt.Errorf("required field short_description missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "description", "name", "rules", "short_description"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Description = *all.Description
	o.Name = *all.Name
	o.Rules = all.Rules
	o.ShortDescription = *all.ShortDescription

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
