// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UserJourneySearchGraphFilter Graph filter for user journey search.
type UserJourneySearchGraphFilter struct {
	// Filter name.
	Name *string `json:"name,omitempty"`
	// Filter operator.
	Operator *string `json:"operator,omitempty"`
	// Target for user journey search.
	Target *UserJourneySearchTarget `json:"target,omitempty"`
	// Filter value.
	Value *int64 `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUserJourneySearchGraphFilter instantiates a new UserJourneySearchGraphFilter object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUserJourneySearchGraphFilter() *UserJourneySearchGraphFilter {
	this := UserJourneySearchGraphFilter{}
	return &this
}

// NewUserJourneySearchGraphFilterWithDefaults instantiates a new UserJourneySearchGraphFilter object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUserJourneySearchGraphFilterWithDefaults() *UserJourneySearchGraphFilter {
	this := UserJourneySearchGraphFilter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserJourneySearchGraphFilter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneySearchGraphFilter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserJourneySearchGraphFilter) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserJourneySearchGraphFilter) SetName(v string) {
	o.Name = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *UserJourneySearchGraphFilter) GetOperator() string {
	if o == nil || o.Operator == nil {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneySearchGraphFilter) GetOperatorOk() (*string, bool) {
	if o == nil || o.Operator == nil {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *UserJourneySearchGraphFilter) HasOperator() bool {
	return o != nil && o.Operator != nil
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *UserJourneySearchGraphFilter) SetOperator(v string) {
	o.Operator = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *UserJourneySearchGraphFilter) GetTarget() UserJourneySearchTarget {
	if o == nil || o.Target == nil {
		var ret UserJourneySearchTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneySearchGraphFilter) GetTargetOk() (*UserJourneySearchTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *UserJourneySearchGraphFilter) HasTarget() bool {
	return o != nil && o.Target != nil
}

// SetTarget gets a reference to the given UserJourneySearchTarget and assigns it to the Target field.
func (o *UserJourneySearchGraphFilter) SetTarget(v UserJourneySearchTarget) {
	o.Target = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UserJourneySearchGraphFilter) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserJourneySearchGraphFilter) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UserJourneySearchGraphFilter) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *UserJourneySearchGraphFilter) SetValue(v int64) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UserJourneySearchGraphFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Operator != nil {
		toSerialize["operator"] = o.Operator
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
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
func (o *UserJourneySearchGraphFilter) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name     *string                  `json:"name,omitempty"`
		Operator *string                  `json:"operator,omitempty"`
		Target   *UserJourneySearchTarget `json:"target,omitempty"`
		Value    *int64                   `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "operator", "target", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.Operator = all.Operator
	if all.Target != nil && all.Target.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Target = all.Target
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
