// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryCostUpsertRequestDataAttributes The definition of `ArbitraryCostUpsertRequestDataAttributes` object.
type ArbitraryCostUpsertRequestDataAttributes struct {
	// The `attributes` `costs_to_allocate`.
	CostsToAllocate []ArbitraryCostUpsertRequestDataAttributesCostsToAllocateItems `json:"costs_to_allocate"`
	// The `attributes` `enabled`.
	Enabled *bool `json:"enabled,omitempty"`
	// The `attributes` `order_id`.
	OrderId *int64 `json:"order_id,omitempty"`
	// The `attributes` `provider`.
	Provider []string `json:"provider"`
	// The `attributes` `rejected`.
	Rejected *bool `json:"rejected,omitempty"`
	// The `attributes` `rule_name`.
	RuleName string `json:"rule_name"`
	// The definition of `ArbitraryCostUpsertRequestDataAttributesStrategy` object.
	Strategy ArbitraryCostUpsertRequestDataAttributesStrategy `json:"strategy"`
	// The `attributes` `type`.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArbitraryCostUpsertRequestDataAttributes instantiates a new ArbitraryCostUpsertRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArbitraryCostUpsertRequestDataAttributes(costsToAllocate []ArbitraryCostUpsertRequestDataAttributesCostsToAllocateItems, provider []string, ruleName string, strategy ArbitraryCostUpsertRequestDataAttributesStrategy, typeVar string) *ArbitraryCostUpsertRequestDataAttributes {
	this := ArbitraryCostUpsertRequestDataAttributes{}
	this.CostsToAllocate = costsToAllocate
	this.Provider = provider
	this.RuleName = ruleName
	this.Strategy = strategy
	this.Type = typeVar
	return &this
}

// NewArbitraryCostUpsertRequestDataAttributesWithDefaults instantiates a new ArbitraryCostUpsertRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArbitraryCostUpsertRequestDataAttributesWithDefaults() *ArbitraryCostUpsertRequestDataAttributes {
	this := ArbitraryCostUpsertRequestDataAttributes{}
	return &this
}

// GetCostsToAllocate returns the CostsToAllocate field value.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetCostsToAllocate() []ArbitraryCostUpsertRequestDataAttributesCostsToAllocateItems {
	if o == nil {
		var ret []ArbitraryCostUpsertRequestDataAttributesCostsToAllocateItems
		return ret
	}
	return o.CostsToAllocate
}

// GetCostsToAllocateOk returns a tuple with the CostsToAllocate field value
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetCostsToAllocateOk() (*[]ArbitraryCostUpsertRequestDataAttributesCostsToAllocateItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostsToAllocate, true
}

// SetCostsToAllocate sets field value.
func (o *ArbitraryCostUpsertRequestDataAttributes) SetCostsToAllocate(v []ArbitraryCostUpsertRequestDataAttributesCostsToAllocateItems) {
	o.CostsToAllocate = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ArbitraryCostUpsertRequestDataAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ArbitraryCostUpsertRequestDataAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetOrderId() int64 {
	if o == nil || o.OrderId == nil {
		var ret int64
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetOrderIdOk() (*int64, bool) {
	if o == nil || o.OrderId == nil {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *ArbitraryCostUpsertRequestDataAttributes) HasOrderId() bool {
	return o != nil && o.OrderId != nil
}

// SetOrderId gets a reference to the given int64 and assigns it to the OrderId field.
func (o *ArbitraryCostUpsertRequestDataAttributes) SetOrderId(v int64) {
	o.OrderId = &v
}

// GetProvider returns the Provider field value.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetProvider() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetProviderOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *ArbitraryCostUpsertRequestDataAttributes) SetProvider(v []string) {
	o.Provider = v
}

// GetRejected returns the Rejected field value if set, zero value otherwise.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetRejected() bool {
	if o == nil || o.Rejected == nil {
		var ret bool
		return ret
	}
	return *o.Rejected
}

// GetRejectedOk returns a tuple with the Rejected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetRejectedOk() (*bool, bool) {
	if o == nil || o.Rejected == nil {
		return nil, false
	}
	return o.Rejected, true
}

// HasRejected returns a boolean if a field has been set.
func (o *ArbitraryCostUpsertRequestDataAttributes) HasRejected() bool {
	return o != nil && o.Rejected != nil
}

// SetRejected gets a reference to the given bool and assigns it to the Rejected field.
func (o *ArbitraryCostUpsertRequestDataAttributes) SetRejected(v bool) {
	o.Rejected = &v
}

// GetRuleName returns the RuleName field value.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetRuleName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetRuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleName, true
}

// SetRuleName sets field value.
func (o *ArbitraryCostUpsertRequestDataAttributes) SetRuleName(v string) {
	o.RuleName = v
}

// GetStrategy returns the Strategy field value.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetStrategy() ArbitraryCostUpsertRequestDataAttributesStrategy {
	if o == nil {
		var ret ArbitraryCostUpsertRequestDataAttributesStrategy
		return ret
	}
	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetStrategyOk() (*ArbitraryCostUpsertRequestDataAttributesStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value.
func (o *ArbitraryCostUpsertRequestDataAttributes) SetStrategy(v ArbitraryCostUpsertRequestDataAttributesStrategy) {
	o.Strategy = v
}

// GetType returns the Type field value.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArbitraryCostUpsertRequestDataAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ArbitraryCostUpsertRequestDataAttributes) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArbitraryCostUpsertRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["costs_to_allocate"] = o.CostsToAllocate
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.OrderId != nil {
		toSerialize["order_id"] = o.OrderId
	}
	toSerialize["provider"] = o.Provider
	if o.Rejected != nil {
		toSerialize["rejected"] = o.Rejected
	}
	toSerialize["rule_name"] = o.RuleName
	toSerialize["strategy"] = o.Strategy
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArbitraryCostUpsertRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CostsToAllocate *[]ArbitraryCostUpsertRequestDataAttributesCostsToAllocateItems `json:"costs_to_allocate"`
		Enabled         *bool                                                           `json:"enabled,omitempty"`
		OrderId         *int64                                                          `json:"order_id,omitempty"`
		Provider        *[]string                                                       `json:"provider"`
		Rejected        *bool                                                           `json:"rejected,omitempty"`
		RuleName        *string                                                         `json:"rule_name"`
		Strategy        *ArbitraryCostUpsertRequestDataAttributesStrategy               `json:"strategy"`
		Type            *string                                                         `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CostsToAllocate == nil {
		return fmt.Errorf("required field costs_to_allocate missing")
	}
	if all.Provider == nil {
		return fmt.Errorf("required field provider missing")
	}
	if all.RuleName == nil {
		return fmt.Errorf("required field rule_name missing")
	}
	if all.Strategy == nil {
		return fmt.Errorf("required field strategy missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"costs_to_allocate", "enabled", "order_id", "provider", "rejected", "rule_name", "strategy", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CostsToAllocate = *all.CostsToAllocate
	o.Enabled = all.Enabled
	o.OrderId = all.OrderId
	o.Provider = *all.Provider
	o.Rejected = all.Rejected
	o.RuleName = *all.RuleName
	if all.Strategy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Strategy = *all.Strategy
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
