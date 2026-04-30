// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ConditionRequest Condition request payload for targeting rules.
type ConditionRequest struct {
	// The user or request attribute to evaluate.
	Attribute string `json:"attribute"`
	// The operator used in a targeting condition.
	Operator ConditionOperator `json:"operator"`
	// Values used by the selected operator.
	Value []string `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewConditionRequest instantiates a new ConditionRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewConditionRequest(attribute string, operator ConditionOperator, value []string) *ConditionRequest {
	this := ConditionRequest{}
	this.Attribute = attribute
	this.Operator = operator
	this.Value = value
	return &this
}

// NewConditionRequestWithDefaults instantiates a new ConditionRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewConditionRequestWithDefaults() *ConditionRequest {
	this := ConditionRequest{}
	return &this
}

// GetAttribute returns the Attribute field value.
func (o *ConditionRequest) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *ConditionRequest) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value.
func (o *ConditionRequest) SetAttribute(v string) {
	o.Attribute = v
}

// GetOperator returns the Operator field value.
func (o *ConditionRequest) GetOperator() ConditionOperator {
	if o == nil {
		var ret ConditionOperator
		return ret
	}
	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ConditionRequest) GetOperatorOk() (*ConditionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value.
func (o *ConditionRequest) SetOperator(v ConditionOperator) {
	o.Operator = v
}

// GetValue returns the Value field value.
func (o *ConditionRequest) GetValue() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ConditionRequest) GetValueOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *ConditionRequest) SetValue(v []string) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ConditionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attribute"] = o.Attribute
	toSerialize["operator"] = o.Operator
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ConditionRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attribute *string            `json:"attribute"`
		Operator  *ConditionOperator `json:"operator"`
		Value     *[]string          `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attribute == nil {
		return fmt.Errorf("required field attribute missing")
	}
	if all.Operator == nil {
		return fmt.Errorf("required field operator missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attribute", "operator", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Attribute = *all.Attribute
	if !all.Operator.IsValid() {
		hasInvalidField = true
	} else {
		o.Operator = *all.Operator
	}
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
