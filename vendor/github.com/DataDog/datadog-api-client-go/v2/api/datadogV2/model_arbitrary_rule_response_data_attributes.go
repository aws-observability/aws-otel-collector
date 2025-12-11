// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ArbitraryRuleResponseDataAttributes The definition of `ArbitraryRuleResponseDataAttributes` object.
type ArbitraryRuleResponseDataAttributes struct {
	// The `attributes` `costs_to_allocate`.
	CostsToAllocate []ArbitraryRuleResponseDataAttributesCostsToAllocateItems `json:"costs_to_allocate"`
	// The `attributes` `created`.
	Created time.Time `json:"created"`
	// The `attributes` `enabled`.
	Enabled bool `json:"enabled"`
	// The `attributes` `last_modified_user_uuid`.
	LastModifiedUserUuid string `json:"last_modified_user_uuid"`
	// The `attributes` `order_id`.
	OrderId int64 `json:"order_id"`
	// The `attributes` `processing_status`.
	ProcessingStatus *string `json:"processing_status,omitempty"`
	// The `attributes` `provider`.
	Provider []string `json:"provider"`
	// The `attributes` `rejected`.
	Rejected *bool `json:"rejected,omitempty"`
	// The `attributes` `rule_name`.
	RuleName string `json:"rule_name"`
	// The definition of `ArbitraryRuleResponseDataAttributesStrategy` object.
	Strategy ArbitraryRuleResponseDataAttributesStrategy `json:"strategy"`
	// The `attributes` `type`.
	Type string `json:"type"`
	// The `attributes` `updated`.
	Updated time.Time `json:"updated"`
	// The `attributes` `version`.
	Version int32 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewArbitraryRuleResponseDataAttributes instantiates a new ArbitraryRuleResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewArbitraryRuleResponseDataAttributes(costsToAllocate []ArbitraryRuleResponseDataAttributesCostsToAllocateItems, created time.Time, enabled bool, lastModifiedUserUuid string, orderId int64, provider []string, ruleName string, strategy ArbitraryRuleResponseDataAttributesStrategy, typeVar string, updated time.Time, version int32) *ArbitraryRuleResponseDataAttributes {
	this := ArbitraryRuleResponseDataAttributes{}
	this.CostsToAllocate = costsToAllocate
	this.Created = created
	this.Enabled = enabled
	this.LastModifiedUserUuid = lastModifiedUserUuid
	this.OrderId = orderId
	this.Provider = provider
	this.RuleName = ruleName
	this.Strategy = strategy
	this.Type = typeVar
	this.Updated = updated
	this.Version = version
	return &this
}

// NewArbitraryRuleResponseDataAttributesWithDefaults instantiates a new ArbitraryRuleResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewArbitraryRuleResponseDataAttributesWithDefaults() *ArbitraryRuleResponseDataAttributes {
	this := ArbitraryRuleResponseDataAttributes{}
	return &this
}

// GetCostsToAllocate returns the CostsToAllocate field value.
func (o *ArbitraryRuleResponseDataAttributes) GetCostsToAllocate() []ArbitraryRuleResponseDataAttributesCostsToAllocateItems {
	if o == nil {
		var ret []ArbitraryRuleResponseDataAttributesCostsToAllocateItems
		return ret
	}
	return o.CostsToAllocate
}

// GetCostsToAllocateOk returns a tuple with the CostsToAllocate field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributes) GetCostsToAllocateOk() (*[]ArbitraryRuleResponseDataAttributesCostsToAllocateItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CostsToAllocate, true
}

// SetCostsToAllocate sets field value.
func (o *ArbitraryRuleResponseDataAttributes) SetCostsToAllocate(v []ArbitraryRuleResponseDataAttributesCostsToAllocateItems) {
	o.CostsToAllocate = v
}

// GetCreated returns the Created field value.
func (o *ArbitraryRuleResponseDataAttributes) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributes) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *ArbitraryRuleResponseDataAttributes) SetCreated(v time.Time) {
	o.Created = v
}

// GetEnabled returns the Enabled field value.
func (o *ArbitraryRuleResponseDataAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ArbitraryRuleResponseDataAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetLastModifiedUserUuid returns the LastModifiedUserUuid field value.
func (o *ArbitraryRuleResponseDataAttributes) GetLastModifiedUserUuid() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LastModifiedUserUuid
}

// GetLastModifiedUserUuidOk returns a tuple with the LastModifiedUserUuid field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributes) GetLastModifiedUserUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedUserUuid, true
}

// SetLastModifiedUserUuid sets field value.
func (o *ArbitraryRuleResponseDataAttributes) SetLastModifiedUserUuid(v string) {
	o.LastModifiedUserUuid = v
}

// GetOrderId returns the OrderId field value.
func (o *ArbitraryRuleResponseDataAttributes) GetOrderId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributes) GetOrderIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value.
func (o *ArbitraryRuleResponseDataAttributes) SetOrderId(v int64) {
	o.OrderId = v
}

// GetProcessingStatus returns the ProcessingStatus field value if set, zero value otherwise.
func (o *ArbitraryRuleResponseDataAttributes) GetProcessingStatus() string {
	if o == nil || o.ProcessingStatus == nil {
		var ret string
		return ret
	}
	return *o.ProcessingStatus
}

// GetProcessingStatusOk returns a tuple with the ProcessingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributes) GetProcessingStatusOk() (*string, bool) {
	if o == nil || o.ProcessingStatus == nil {
		return nil, false
	}
	return o.ProcessingStatus, true
}

// HasProcessingStatus returns a boolean if a field has been set.
func (o *ArbitraryRuleResponseDataAttributes) HasProcessingStatus() bool {
	return o != nil && o.ProcessingStatus != nil
}

// SetProcessingStatus gets a reference to the given string and assigns it to the ProcessingStatus field.
func (o *ArbitraryRuleResponseDataAttributes) SetProcessingStatus(v string) {
	o.ProcessingStatus = &v
}

// GetProvider returns the Provider field value.
func (o *ArbitraryRuleResponseDataAttributes) GetProvider() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributes) GetProviderOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value.
func (o *ArbitraryRuleResponseDataAttributes) SetProvider(v []string) {
	o.Provider = v
}

// GetRejected returns the Rejected field value if set, zero value otherwise.
func (o *ArbitraryRuleResponseDataAttributes) GetRejected() bool {
	if o == nil || o.Rejected == nil {
		var ret bool
		return ret
	}
	return *o.Rejected
}

// GetRejectedOk returns a tuple with the Rejected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributes) GetRejectedOk() (*bool, bool) {
	if o == nil || o.Rejected == nil {
		return nil, false
	}
	return o.Rejected, true
}

// HasRejected returns a boolean if a field has been set.
func (o *ArbitraryRuleResponseDataAttributes) HasRejected() bool {
	return o != nil && o.Rejected != nil
}

// SetRejected gets a reference to the given bool and assigns it to the Rejected field.
func (o *ArbitraryRuleResponseDataAttributes) SetRejected(v bool) {
	o.Rejected = &v
}

// GetRuleName returns the RuleName field value.
func (o *ArbitraryRuleResponseDataAttributes) GetRuleName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributes) GetRuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleName, true
}

// SetRuleName sets field value.
func (o *ArbitraryRuleResponseDataAttributes) SetRuleName(v string) {
	o.RuleName = v
}

// GetStrategy returns the Strategy field value.
func (o *ArbitraryRuleResponseDataAttributes) GetStrategy() ArbitraryRuleResponseDataAttributesStrategy {
	if o == nil {
		var ret ArbitraryRuleResponseDataAttributesStrategy
		return ret
	}
	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributes) GetStrategyOk() (*ArbitraryRuleResponseDataAttributesStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value.
func (o *ArbitraryRuleResponseDataAttributes) SetStrategy(v ArbitraryRuleResponseDataAttributesStrategy) {
	o.Strategy = v
}

// GetType returns the Type field value.
func (o *ArbitraryRuleResponseDataAttributes) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributes) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ArbitraryRuleResponseDataAttributes) SetType(v string) {
	o.Type = v
}

// GetUpdated returns the Updated field value.
func (o *ArbitraryRuleResponseDataAttributes) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributes) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value.
func (o *ArbitraryRuleResponseDataAttributes) SetUpdated(v time.Time) {
	o.Updated = v
}

// GetVersion returns the Version field value.
func (o *ArbitraryRuleResponseDataAttributes) GetVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ArbitraryRuleResponseDataAttributes) GetVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *ArbitraryRuleResponseDataAttributes) SetVersion(v int32) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ArbitraryRuleResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["costs_to_allocate"] = o.CostsToAllocate
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["last_modified_user_uuid"] = o.LastModifiedUserUuid
	toSerialize["order_id"] = o.OrderId
	if o.ProcessingStatus != nil {
		toSerialize["processing_status"] = o.ProcessingStatus
	}
	toSerialize["provider"] = o.Provider
	if o.Rejected != nil {
		toSerialize["rejected"] = o.Rejected
	}
	toSerialize["rule_name"] = o.RuleName
	toSerialize["strategy"] = o.Strategy
	toSerialize["type"] = o.Type
	if o.Updated.Nanosecond() == 0 {
		toSerialize["updated"] = o.Updated.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated"] = o.Updated.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ArbitraryRuleResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CostsToAllocate      *[]ArbitraryRuleResponseDataAttributesCostsToAllocateItems `json:"costs_to_allocate"`
		Created              *time.Time                                                 `json:"created"`
		Enabled              *bool                                                      `json:"enabled"`
		LastModifiedUserUuid *string                                                    `json:"last_modified_user_uuid"`
		OrderId              *int64                                                     `json:"order_id"`
		ProcessingStatus     *string                                                    `json:"processing_status,omitempty"`
		Provider             *[]string                                                  `json:"provider"`
		Rejected             *bool                                                      `json:"rejected,omitempty"`
		RuleName             *string                                                    `json:"rule_name"`
		Strategy             *ArbitraryRuleResponseDataAttributesStrategy               `json:"strategy"`
		Type                 *string                                                    `json:"type"`
		Updated              *time.Time                                                 `json:"updated"`
		Version              *int32                                                     `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CostsToAllocate == nil {
		return fmt.Errorf("required field costs_to_allocate missing")
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.LastModifiedUserUuid == nil {
		return fmt.Errorf("required field last_modified_user_uuid missing")
	}
	if all.OrderId == nil {
		return fmt.Errorf("required field order_id missing")
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
	if all.Updated == nil {
		return fmt.Errorf("required field updated missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"costs_to_allocate", "created", "enabled", "last_modified_user_uuid", "order_id", "processing_status", "provider", "rejected", "rule_name", "strategy", "type", "updated", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CostsToAllocate = *all.CostsToAllocate
	o.Created = *all.Created
	o.Enabled = *all.Enabled
	o.LastModifiedUserUuid = *all.LastModifiedUserUuid
	o.OrderId = *all.OrderId
	o.ProcessingStatus = all.ProcessingStatus
	o.Provider = *all.Provider
	o.Rejected = all.Rejected
	o.RuleName = *all.RuleName
	if all.Strategy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Strategy = *all.Strategy
	o.Type = *all.Type
	o.Updated = *all.Updated
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
