// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RecommendationsFilterRequestSortItems A single sort clause applied to the cost recommendations result set.
type RecommendationsFilterRequestSortItems struct {
	// Field to sort by (for example, `potential_daily_savings.amount`).
	Expression *string `json:"expression,omitempty"`
	// Sort direction, either `ASC` or `DESC`.
	Order *string `json:"order,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRecommendationsFilterRequestSortItems instantiates a new RecommendationsFilterRequestSortItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRecommendationsFilterRequestSortItems() *RecommendationsFilterRequestSortItems {
	this := RecommendationsFilterRequestSortItems{}
	return &this
}

// NewRecommendationsFilterRequestSortItemsWithDefaults instantiates a new RecommendationsFilterRequestSortItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRecommendationsFilterRequestSortItemsWithDefaults() *RecommendationsFilterRequestSortItems {
	this := RecommendationsFilterRequestSortItems{}
	return &this
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *RecommendationsFilterRequestSortItems) GetExpression() string {
	if o == nil || o.Expression == nil {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsFilterRequestSortItems) GetExpressionOk() (*string, bool) {
	if o == nil || o.Expression == nil {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *RecommendationsFilterRequestSortItems) HasExpression() bool {
	return o != nil && o.Expression != nil
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *RecommendationsFilterRequestSortItems) SetExpression(v string) {
	o.Expression = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *RecommendationsFilterRequestSortItems) GetOrder() string {
	if o == nil || o.Order == nil {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsFilterRequestSortItems) GetOrderOk() (*string, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *RecommendationsFilterRequestSortItems) HasOrder() bool {
	return o != nil && o.Order != nil
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *RecommendationsFilterRequestSortItems) SetOrder(v string) {
	o.Order = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RecommendationsFilterRequestSortItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Expression != nil {
		toSerialize["expression"] = o.Expression
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RecommendationsFilterRequestSortItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Expression *string `json:"expression,omitempty"`
		Order      *string `json:"order,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"expression", "order"})
	} else {
		return err
	}
	o.Expression = all.Expression
	o.Order = all.Order

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
