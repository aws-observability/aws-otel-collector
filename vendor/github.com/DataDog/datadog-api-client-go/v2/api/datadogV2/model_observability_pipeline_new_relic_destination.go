// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineNewRelicDestination The `new_relic` destination sends logs to the New Relic platform.
type ObservabilityPipelineNewRelicDestination struct {
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The New Relic region.
	Region ObservabilityPipelineNewRelicDestinationRegion `json:"region"`
	// The destination type. The value should always be `new_relic`.
	Type ObservabilityPipelineNewRelicDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineNewRelicDestination instantiates a new ObservabilityPipelineNewRelicDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineNewRelicDestination(id string, inputs []string, region ObservabilityPipelineNewRelicDestinationRegion, typeVar ObservabilityPipelineNewRelicDestinationType) *ObservabilityPipelineNewRelicDestination {
	this := ObservabilityPipelineNewRelicDestination{}
	this.Id = id
	this.Inputs = inputs
	this.Region = region
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineNewRelicDestinationWithDefaults instantiates a new ObservabilityPipelineNewRelicDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineNewRelicDestinationWithDefaults() *ObservabilityPipelineNewRelicDestination {
	this := ObservabilityPipelineNewRelicDestination{}
	var typeVar ObservabilityPipelineNewRelicDestinationType = OBSERVABILITYPIPELINENEWRELICDESTINATIONTYPE_NEW_RELIC
	this.Type = typeVar
	return &this
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineNewRelicDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineNewRelicDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineNewRelicDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineNewRelicDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineNewRelicDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineNewRelicDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetRegion returns the Region field value.
func (o *ObservabilityPipelineNewRelicDestination) GetRegion() ObservabilityPipelineNewRelicDestinationRegion {
	if o == nil {
		var ret ObservabilityPipelineNewRelicDestinationRegion
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineNewRelicDestination) GetRegionOk() (*ObservabilityPipelineNewRelicDestinationRegion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *ObservabilityPipelineNewRelicDestination) SetRegion(v ObservabilityPipelineNewRelicDestinationRegion) {
	o.Region = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineNewRelicDestination) GetType() ObservabilityPipelineNewRelicDestinationType {
	if o == nil {
		var ret ObservabilityPipelineNewRelicDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineNewRelicDestination) GetTypeOk() (*ObservabilityPipelineNewRelicDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineNewRelicDestination) SetType(v ObservabilityPipelineNewRelicDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineNewRelicDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	toSerialize["region"] = o.Region
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineNewRelicDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id     *string                                         `json:"id"`
		Inputs *[]string                                       `json:"inputs"`
		Region *ObservabilityPipelineNewRelicDestinationRegion `json:"region"`
		Type   *ObservabilityPipelineNewRelicDestinationType   `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "inputs", "region", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	if !all.Region.IsValid() {
		hasInvalidField = true
	} else {
		o.Region = *all.Region
	}
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
