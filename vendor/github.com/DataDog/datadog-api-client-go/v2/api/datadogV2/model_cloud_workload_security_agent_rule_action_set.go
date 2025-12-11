// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudWorkloadSecurityAgentRuleActionSet The set action applied on the scope matching the rule
type CloudWorkloadSecurityAgentRuleActionSet struct {
	// Whether the value should be appended to the field.
	Append *bool `json:"append,omitempty"`
	// The default value of the set action
	DefaultValue *string `json:"default_value,omitempty"`
	// The expression of the set action.
	Expression *string `json:"expression,omitempty"`
	// The field of the set action
	Field *string `json:"field,omitempty"`
	// Whether the value should be inherited.
	Inherited *bool `json:"inherited,omitempty"`
	// The name of the set action
	Name *string `json:"name,omitempty"`
	// The scope of the set action.
	Scope *string `json:"scope,omitempty"`
	// The size of the set action.
	Size *int64 `json:"size,omitempty"`
	// The time to live of the set action.
	Ttl *int64 `json:"ttl,omitempty"`
	// The value of the set action
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudWorkloadSecurityAgentRuleActionSet instantiates a new CloudWorkloadSecurityAgentRuleActionSet object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudWorkloadSecurityAgentRuleActionSet() *CloudWorkloadSecurityAgentRuleActionSet {
	this := CloudWorkloadSecurityAgentRuleActionSet{}
	return &this
}

// NewCloudWorkloadSecurityAgentRuleActionSetWithDefaults instantiates a new CloudWorkloadSecurityAgentRuleActionSet object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudWorkloadSecurityAgentRuleActionSetWithDefaults() *CloudWorkloadSecurityAgentRuleActionSet {
	this := CloudWorkloadSecurityAgentRuleActionSet{}
	return &this
}

// GetAppend returns the Append field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetAppend() bool {
	if o == nil || o.Append == nil {
		var ret bool
		return ret
	}
	return *o.Append
}

// GetAppendOk returns a tuple with the Append field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetAppendOk() (*bool, bool) {
	if o == nil || o.Append == nil {
		return nil, false
	}
	return o.Append, true
}

// HasAppend returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) HasAppend() bool {
	return o != nil && o.Append != nil
}

// SetAppend gets a reference to the given bool and assigns it to the Append field.
func (o *CloudWorkloadSecurityAgentRuleActionSet) SetAppend(v bool) {
	o.Append = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetDefaultValueOk() (*string, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) HasDefaultValue() bool {
	return o != nil && o.DefaultValue != nil
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *CloudWorkloadSecurityAgentRuleActionSet) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) HasExpression() bool {
	return o != nil && o.Expression != nil
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *CloudWorkloadSecurityAgentRuleActionSet) SetExpression(v string) {
	o.Expression = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) HasField() bool {
	return o != nil && o.Field != nil
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *CloudWorkloadSecurityAgentRuleActionSet) SetField(v string) {
	o.Field = &v
}

// GetInherited returns the Inherited field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetInherited() bool {
	if o == nil || o.Inherited == nil {
		var ret bool
		return ret
	}
	return *o.Inherited
}

// GetInheritedOk returns a tuple with the Inherited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetInheritedOk() (*bool, bool) {
	if o == nil || o.Inherited == nil {
		return nil, false
	}
	return o.Inherited, true
}

// HasInherited returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) HasInherited() bool {
	return o != nil && o.Inherited != nil
}

// SetInherited gets a reference to the given bool and assigns it to the Inherited field.
func (o *CloudWorkloadSecurityAgentRuleActionSet) SetInherited(v bool) {
	o.Inherited = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CloudWorkloadSecurityAgentRuleActionSet) SetName(v string) {
	o.Name = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *CloudWorkloadSecurityAgentRuleActionSet) SetScope(v string) {
	o.Scope = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *CloudWorkloadSecurityAgentRuleActionSet) SetSize(v int64) {
	o.Size = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetTtl() int64 {
	if o == nil || o.Ttl == nil {
		var ret int64
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetTtlOk() (*int64, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) HasTtl() bool {
	return o != nil && o.Ttl != nil
}

// SetTtl gets a reference to the given int64 and assigns it to the Ttl field.
func (o *CloudWorkloadSecurityAgentRuleActionSet) SetTtl(v int64) {
	o.Ttl = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CloudWorkloadSecurityAgentRuleActionSet) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *CloudWorkloadSecurityAgentRuleActionSet) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudWorkloadSecurityAgentRuleActionSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Append != nil {
		toSerialize["append"] = o.Append
	}
	if o.DefaultValue != nil {
		toSerialize["default_value"] = o.DefaultValue
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	if o.Inherited != nil {
		toSerialize["inherited"] = o.Inherited
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
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
func (o *CloudWorkloadSecurityAgentRuleActionSet) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Append       *bool   `json:"append,omitempty"`
		DefaultValue *string `json:"default_value,omitempty"`
		Expression   *string `json:"expression,omitempty"`
		Field        *string `json:"field,omitempty"`
		Inherited    *bool   `json:"inherited,omitempty"`
		Name         *string `json:"name,omitempty"`
		Scope        *string `json:"scope,omitempty"`
		Size         *int64  `json:"size,omitempty"`
		Ttl          *int64  `json:"ttl,omitempty"`
		Value        *string `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"append", "default_value", "expression", "field", "inherited", "name", "scope", "size", "ttl", "value"})
	} else {
		return err
	}
	o.Append = all.Append
	o.DefaultValue = all.DefaultValue
	o.Expression = all.Expression
	o.Field = all.Field
	o.Inherited = all.Inherited
	o.Name = all.Name
	o.Scope = all.Scope
	o.Size = all.Size
	o.Ttl = all.Ttl
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
