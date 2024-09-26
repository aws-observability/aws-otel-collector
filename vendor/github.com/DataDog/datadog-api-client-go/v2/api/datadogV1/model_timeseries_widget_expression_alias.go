// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimeseriesWidgetExpressionAlias Define an expression alias.
type TimeseriesWidgetExpressionAlias struct {
	// Expression alias.
	AliasName *string `json:"alias_name,omitempty"`
	// Expression name.
	Expression string `json:"expression"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimeseriesWidgetExpressionAlias instantiates a new TimeseriesWidgetExpressionAlias object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimeseriesWidgetExpressionAlias(expression string) *TimeseriesWidgetExpressionAlias {
	this := TimeseriesWidgetExpressionAlias{}
	this.Expression = expression
	return &this
}

// NewTimeseriesWidgetExpressionAliasWithDefaults instantiates a new TimeseriesWidgetExpressionAlias object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimeseriesWidgetExpressionAliasWithDefaults() *TimeseriesWidgetExpressionAlias {
	this := TimeseriesWidgetExpressionAlias{}
	return &this
}

// GetAliasName returns the AliasName field value if set, zero value otherwise.
func (o *TimeseriesWidgetExpressionAlias) GetAliasName() string {
	if o == nil || o.AliasName == nil {
		var ret string
		return ret
	}
	return *o.AliasName
}

// GetAliasNameOk returns a tuple with the AliasName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetExpressionAlias) GetAliasNameOk() (*string, bool) {
	if o == nil || o.AliasName == nil {
		return nil, false
	}
	return o.AliasName, true
}

// HasAliasName returns a boolean if a field has been set.
func (o *TimeseriesWidgetExpressionAlias) HasAliasName() bool {
	return o != nil && o.AliasName != nil
}

// SetAliasName gets a reference to the given string and assigns it to the AliasName field.
func (o *TimeseriesWidgetExpressionAlias) SetAliasName(v string) {
	o.AliasName = &v
}

// GetExpression returns the Expression field value.
func (o *TimeseriesWidgetExpressionAlias) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetExpressionAlias) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value.
func (o *TimeseriesWidgetExpressionAlias) SetExpression(v string) {
	o.Expression = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimeseriesWidgetExpressionAlias) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AliasName != nil {
		toSerialize["alias_name"] = o.AliasName
	}
	toSerialize["expression"] = o.Expression

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimeseriesWidgetExpressionAlias) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AliasName  *string `json:"alias_name,omitempty"`
		Expression *string `json:"expression"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Expression == nil {
		return fmt.Errorf("required field expression missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alias_name", "expression"})
	} else {
		return err
	}
	o.AliasName = all.AliasName
	o.Expression = *all.Expression

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
