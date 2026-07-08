// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CostOrchestrator A container orchestrator detected in Cloud Cost Management data.
type CostOrchestrator struct {
	// The orchestrator name, for example `kubernetes` or `ecs`.
	Id string `json:"id"`
	// Type of the Cloud Cost Management orchestrator resource.
	Type CostOrchestratorType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCostOrchestrator instantiates a new CostOrchestrator object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCostOrchestrator(id string, typeVar CostOrchestratorType) *CostOrchestrator {
	this := CostOrchestrator{}
	this.Id = id
	this.Type = typeVar
	return &this
}

// NewCostOrchestratorWithDefaults instantiates a new CostOrchestrator object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCostOrchestratorWithDefaults() *CostOrchestrator {
	this := CostOrchestrator{}
	var typeVar CostOrchestratorType = COSTORCHESTRATORTYPE_COST_ORCHESTRATOR
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *CostOrchestrator) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CostOrchestrator) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CostOrchestrator) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value.
func (o *CostOrchestrator) GetType() CostOrchestratorType {
	if o == nil {
		var ret CostOrchestratorType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CostOrchestrator) GetTypeOk() (*CostOrchestratorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CostOrchestrator) SetType(v CostOrchestratorType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CostOrchestrator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CostOrchestrator) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id   *string               `json:"id"`
		Type *CostOrchestratorType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
