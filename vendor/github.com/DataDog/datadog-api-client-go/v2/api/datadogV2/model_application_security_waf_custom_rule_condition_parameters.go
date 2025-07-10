// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleConditionParameters The scope of the WAF custom rule.
type ApplicationSecurityWafCustomRuleConditionParameters struct {
	// Identifier of a list of data from the denylist. Can only be used as substitution from the list parameter.
	Data *string `json:"data,omitempty"`
	// List of inputs on which at least one should match with the given operator.
	Inputs []ApplicationSecurityWafCustomRuleConditionInput `json:"inputs"`
	// List of value to use with the condition. Only used with the phrase_match, !phrase_match, exact_match and
	// !exact_match operator.
	List []string `json:"list,omitempty"`
	// Options for the operator of this condition.
	Options *ApplicationSecurityWafCustomRuleConditionOptions `json:"options,omitempty"`
	// Regex to use with the condition. Only used with match_regex and !match_regex operator.
	Regex *string `json:"regex,omitempty"`
	// Store the captured value in the specified tag name. Only used with the capture_data operator.
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityWafCustomRuleConditionParameters instantiates a new ApplicationSecurityWafCustomRuleConditionParameters object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityWafCustomRuleConditionParameters(inputs []ApplicationSecurityWafCustomRuleConditionInput) *ApplicationSecurityWafCustomRuleConditionParameters {
	this := ApplicationSecurityWafCustomRuleConditionParameters{}
	this.Inputs = inputs
	return &this
}

// NewApplicationSecurityWafCustomRuleConditionParametersWithDefaults instantiates a new ApplicationSecurityWafCustomRuleConditionParameters object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityWafCustomRuleConditionParametersWithDefaults() *ApplicationSecurityWafCustomRuleConditionParameters {
	this := ApplicationSecurityWafCustomRuleConditionParameters{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) SetData(v string) {
	o.Data = &v
}

// GetInputs returns the Inputs field value.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) GetInputs() []ApplicationSecurityWafCustomRuleConditionInput {
	if o == nil {
		var ret []ApplicationSecurityWafCustomRuleConditionInput
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) GetInputsOk() (*[]ApplicationSecurityWafCustomRuleConditionInput, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) SetInputs(v []ApplicationSecurityWafCustomRuleConditionInput) {
	o.Inputs = v
}

// GetList returns the List field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) GetList() []string {
	if o == nil || o.List == nil {
		var ret []string
		return ret
	}
	return o.List
}

// GetListOk returns a tuple with the List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) GetListOk() (*[]string, bool) {
	if o == nil || o.List == nil {
		return nil, false
	}
	return &o.List, true
}

// HasList returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) HasList() bool {
	return o != nil && o.List != nil
}

// SetList gets a reference to the given []string and assigns it to the List field.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) SetList(v []string) {
	o.List = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) GetOptions() ApplicationSecurityWafCustomRuleConditionOptions {
	if o == nil || o.Options == nil {
		var ret ApplicationSecurityWafCustomRuleConditionOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) GetOptionsOk() (*ApplicationSecurityWafCustomRuleConditionOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) HasOptions() bool {
	return o != nil && o.Options != nil
}

// SetOptions gets a reference to the given ApplicationSecurityWafCustomRuleConditionOptions and assigns it to the Options field.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) SetOptions(v ApplicationSecurityWafCustomRuleConditionOptions) {
	o.Options = &v
}

// GetRegex returns the Regex field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) GetRegex() string {
	if o == nil || o.Regex == nil {
		var ret string
		return ret
	}
	return *o.Regex
}

// GetRegexOk returns a tuple with the Regex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) GetRegexOk() (*string, bool) {
	if o == nil || o.Regex == nil {
		return nil, false
	}
	return o.Regex, true
}

// HasRegex returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) HasRegex() bool {
	return o != nil && o.Regex != nil
}

// SetRegex gets a reference to the given string and assigns it to the Regex field.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) SetRegex(v string) {
	o.Regex = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityWafCustomRuleConditionParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	toSerialize["inputs"] = o.Inputs
	if o.List != nil {
		toSerialize["list"] = o.List
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Regex != nil {
		toSerialize["regex"] = o.Regex
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityWafCustomRuleConditionParameters) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data    *string                                           `json:"data,omitempty"`
		Inputs  *[]ApplicationSecurityWafCustomRuleConditionInput `json:"inputs"`
		List    []string                                          `json:"list,omitempty"`
		Options *ApplicationSecurityWafCustomRuleConditionOptions `json:"options,omitempty"`
		Regex   *string                                           `json:"regex,omitempty"`
		Value   *string                                           `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "inputs", "list", "options", "regex", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	o.Inputs = *all.Inputs
	o.List = all.List
	if all.Options != nil && all.Options.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Options = all.Options
	o.Regex = all.Regex
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
