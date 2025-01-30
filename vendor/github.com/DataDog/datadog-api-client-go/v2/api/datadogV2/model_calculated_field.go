// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CalculatedField Calculated field.
type CalculatedField struct {
	// Expression.
	Expression string `json:"expression"`
	// Field name.
	Name string `json:"name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCalculatedField instantiates a new CalculatedField object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCalculatedField(expression string, name string) *CalculatedField {
	this := CalculatedField{}
	this.Expression = expression
	this.Name = name
	return &this
}

// NewCalculatedFieldWithDefaults instantiates a new CalculatedField object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCalculatedFieldWithDefaults() *CalculatedField {
	this := CalculatedField{}
	return &this
}

// GetExpression returns the Expression field value.
func (o *CalculatedField) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *CalculatedField) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value.
func (o *CalculatedField) SetExpression(v string) {
	o.Expression = v
}

// GetName returns the Name field value.
func (o *CalculatedField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CalculatedField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CalculatedField) SetName(v string) {
	o.Name = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CalculatedField) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["expression"] = o.Expression
	toSerialize["name"] = o.Name

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CalculatedField) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Expression *string `json:"expression"`
		Name       *string `json:"name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Expression == nil {
		return fmt.Errorf("required field expression missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"expression", "name"})
	} else {
		return err
	}
	o.Expression = *all.Expression
	o.Name = *all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
