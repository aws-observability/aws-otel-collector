// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetAstResponseDataAttributes The attributes of the get-AST response, containing the parsed abstract syntax tree.
type GetAstResponseDataAttributes struct {
	// The parsed abstract syntax tree as a JSON object.
	Ast map[string]interface{} `json:"ast"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetAstResponseDataAttributes instantiates a new GetAstResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetAstResponseDataAttributes(ast map[string]interface{}) *GetAstResponseDataAttributes {
	this := GetAstResponseDataAttributes{}
	this.Ast = ast
	return &this
}

// NewGetAstResponseDataAttributesWithDefaults instantiates a new GetAstResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetAstResponseDataAttributesWithDefaults() *GetAstResponseDataAttributes {
	this := GetAstResponseDataAttributes{}
	return &this
}

// GetAst returns the Ast field value.
func (o *GetAstResponseDataAttributes) GetAst() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Ast
}

// GetAstOk returns a tuple with the Ast field value
// and a boolean to check if the value has been set.
func (o *GetAstResponseDataAttributes) GetAstOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ast, true
}

// SetAst sets field value.
func (o *GetAstResponseDataAttributes) SetAst(v map[string]interface{}) {
	o.Ast = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetAstResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["ast"] = o.Ast

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetAstResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ast *map[string]interface{} `json:"ast"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Ast == nil {
		return fmt.Errorf("required field ast missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ast"})
	} else {
		return err
	}
	o.Ast = *all.Ast

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
