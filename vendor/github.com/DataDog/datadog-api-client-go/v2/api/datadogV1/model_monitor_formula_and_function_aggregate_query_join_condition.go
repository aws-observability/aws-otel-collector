// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorFormulaAndFunctionAggregateQueryJoinCondition Join condition for aggregate augmented queries.
type MonitorFormulaAndFunctionAggregateQueryJoinCondition struct {
	// Attribute from the augment query to join on.
	AugmentAttribute string `json:"augment_attribute"`
	// Attribute from the base query to join on.
	BaseAttribute string `json:"base_attribute"`
	// Join type for aggregate query join conditions.
	JoinType MonitorFormulaAndFunctionAggregateQueryJoinType `json:"join_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewMonitorFormulaAndFunctionAggregateQueryJoinCondition instantiates a new MonitorFormulaAndFunctionAggregateQueryJoinCondition object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorFormulaAndFunctionAggregateQueryJoinCondition(augmentAttribute string, baseAttribute string, joinType MonitorFormulaAndFunctionAggregateQueryJoinType) *MonitorFormulaAndFunctionAggregateQueryJoinCondition {
	this := MonitorFormulaAndFunctionAggregateQueryJoinCondition{}
	this.AugmentAttribute = augmentAttribute
	this.BaseAttribute = baseAttribute
	this.JoinType = joinType
	return &this
}

// NewMonitorFormulaAndFunctionAggregateQueryJoinConditionWithDefaults instantiates a new MonitorFormulaAndFunctionAggregateQueryJoinCondition object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorFormulaAndFunctionAggregateQueryJoinConditionWithDefaults() *MonitorFormulaAndFunctionAggregateQueryJoinCondition {
	this := MonitorFormulaAndFunctionAggregateQueryJoinCondition{}
	return &this
}

// GetAugmentAttribute returns the AugmentAttribute field value.
func (o *MonitorFormulaAndFunctionAggregateQueryJoinCondition) GetAugmentAttribute() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AugmentAttribute
}

// GetAugmentAttributeOk returns a tuple with the AugmentAttribute field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateQueryJoinCondition) GetAugmentAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AugmentAttribute, true
}

// SetAugmentAttribute sets field value.
func (o *MonitorFormulaAndFunctionAggregateQueryJoinCondition) SetAugmentAttribute(v string) {
	o.AugmentAttribute = v
}

// GetBaseAttribute returns the BaseAttribute field value.
func (o *MonitorFormulaAndFunctionAggregateQueryJoinCondition) GetBaseAttribute() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BaseAttribute
}

// GetBaseAttributeOk returns a tuple with the BaseAttribute field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateQueryJoinCondition) GetBaseAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseAttribute, true
}

// SetBaseAttribute sets field value.
func (o *MonitorFormulaAndFunctionAggregateQueryJoinCondition) SetBaseAttribute(v string) {
	o.BaseAttribute = v
}

// GetJoinType returns the JoinType field value.
func (o *MonitorFormulaAndFunctionAggregateQueryJoinCondition) GetJoinType() MonitorFormulaAndFunctionAggregateQueryJoinType {
	if o == nil {
		var ret MonitorFormulaAndFunctionAggregateQueryJoinType
		return ret
	}
	return o.JoinType
}

// GetJoinTypeOk returns a tuple with the JoinType field value
// and a boolean to check if the value has been set.
func (o *MonitorFormulaAndFunctionAggregateQueryJoinCondition) GetJoinTypeOk() (*MonitorFormulaAndFunctionAggregateQueryJoinType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinType, true
}

// SetJoinType sets field value.
func (o *MonitorFormulaAndFunctionAggregateQueryJoinCondition) SetJoinType(v MonitorFormulaAndFunctionAggregateQueryJoinType) {
	o.JoinType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorFormulaAndFunctionAggregateQueryJoinCondition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["augment_attribute"] = o.AugmentAttribute
	toSerialize["base_attribute"] = o.BaseAttribute
	toSerialize["join_type"] = o.JoinType
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorFormulaAndFunctionAggregateQueryJoinCondition) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AugmentAttribute *string                                          `json:"augment_attribute"`
		BaseAttribute    *string                                          `json:"base_attribute"`
		JoinType         *MonitorFormulaAndFunctionAggregateQueryJoinType `json:"join_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AugmentAttribute == nil {
		return fmt.Errorf("required field augment_attribute missing")
	}
	if all.BaseAttribute == nil {
		return fmt.Errorf("required field base_attribute missing")
	}
	if all.JoinType == nil {
		return fmt.Errorf("required field join_type missing")
	}

	hasInvalidField := false
	o.AugmentAttribute = *all.AugmentAttribute
	o.BaseAttribute = *all.BaseAttribute
	if !all.JoinType.IsValid() {
		hasInvalidField = true
	} else {
		o.JoinType = *all.JoinType
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
