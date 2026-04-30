// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TargetingRule Targeting rule details.
type TargetingRule struct {
	// Conditions evaluated by this targeting rule.
	Conditions []Condition `json:"conditions"`
	// The timestamp when the targeting rule was created.
	CreatedAt time.Time `json:"created_at"`
	// The unique identifier of the targeting rule.
	Id uuid.UUID `json:"id"`
	// The timestamp when the targeting rule was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTargetingRule instantiates a new TargetingRule object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTargetingRule(conditions []Condition, createdAt time.Time, id uuid.UUID, updatedAt time.Time) *TargetingRule {
	this := TargetingRule{}
	this.Conditions = conditions
	this.CreatedAt = createdAt
	this.Id = id
	this.UpdatedAt = updatedAt
	return &this
}

// NewTargetingRuleWithDefaults instantiates a new TargetingRule object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTargetingRuleWithDefaults() *TargetingRule {
	this := TargetingRule{}
	return &this
}

// GetConditions returns the Conditions field value.
func (o *TargetingRule) GetConditions() []Condition {
	if o == nil {
		var ret []Condition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *TargetingRule) GetConditionsOk() (*[]Condition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// SetConditions sets field value.
func (o *TargetingRule) SetConditions(v []Condition) {
	o.Conditions = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *TargetingRule) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TargetingRule) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *TargetingRule) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetId returns the Id field value.
func (o *TargetingRule) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TargetingRule) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *TargetingRule) SetId(v uuid.UUID) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *TargetingRule) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TargetingRule) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *TargetingRule) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o TargetingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["conditions"] = o.Conditions
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["id"] = o.Id
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TargetingRule) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Conditions *[]Condition `json:"conditions"`
		CreatedAt  *time.Time   `json:"created_at"`
		Id         *uuid.UUID   `json:"id"`
		UpdatedAt  *time.Time   `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Conditions == nil {
		return fmt.Errorf("required field conditions missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"conditions", "created_at", "id", "updated_at"})
	} else {
		return err
	}
	o.Conditions = *all.Conditions
	o.CreatedAt = *all.CreatedAt
	o.Id = *all.Id
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
