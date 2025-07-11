// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineOpenSearchDestination The `opensearch` destination writes logs to an OpenSearch cluster.
type ObservabilityPipelineOpenSearchDestination struct {
	// The index to write logs to.
	BulkIndex *string `json:"bulk_index,omitempty"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The destination type. The value should always be `opensearch`.
	Type ObservabilityPipelineOpenSearchDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineOpenSearchDestination instantiates a new ObservabilityPipelineOpenSearchDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineOpenSearchDestination(id string, inputs []string, typeVar ObservabilityPipelineOpenSearchDestinationType) *ObservabilityPipelineOpenSearchDestination {
	this := ObservabilityPipelineOpenSearchDestination{}
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineOpenSearchDestinationWithDefaults instantiates a new ObservabilityPipelineOpenSearchDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineOpenSearchDestinationWithDefaults() *ObservabilityPipelineOpenSearchDestination {
	this := ObservabilityPipelineOpenSearchDestination{}
	var typeVar ObservabilityPipelineOpenSearchDestinationType = OBSERVABILITYPIPELINEOPENSEARCHDESTINATIONTYPE_OPENSEARCH
	this.Type = typeVar
	return &this
}

// GetBulkIndex returns the BulkIndex field value if set, zero value otherwise.
func (o *ObservabilityPipelineOpenSearchDestination) GetBulkIndex() string {
	if o == nil || o.BulkIndex == nil {
		var ret string
		return ret
	}
	return *o.BulkIndex
}

// GetBulkIndexOk returns a tuple with the BulkIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpenSearchDestination) GetBulkIndexOk() (*string, bool) {
	if o == nil || o.BulkIndex == nil {
		return nil, false
	}
	return o.BulkIndex, true
}

// HasBulkIndex returns a boolean if a field has been set.
func (o *ObservabilityPipelineOpenSearchDestination) HasBulkIndex() bool {
	return o != nil && o.BulkIndex != nil
}

// SetBulkIndex gets a reference to the given string and assigns it to the BulkIndex field.
func (o *ObservabilityPipelineOpenSearchDestination) SetBulkIndex(v string) {
	o.BulkIndex = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineOpenSearchDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpenSearchDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineOpenSearchDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineOpenSearchDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpenSearchDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineOpenSearchDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineOpenSearchDestination) GetType() ObservabilityPipelineOpenSearchDestinationType {
	if o == nil {
		var ret ObservabilityPipelineOpenSearchDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineOpenSearchDestination) GetTypeOk() (*ObservabilityPipelineOpenSearchDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineOpenSearchDestination) SetType(v ObservabilityPipelineOpenSearchDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineOpenSearchDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BulkIndex != nil {
		toSerialize["bulk_index"] = o.BulkIndex
	}
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineOpenSearchDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BulkIndex *string                                         `json:"bulk_index,omitempty"`
		Id        *string                                         `json:"id"`
		Inputs    *[]string                                       `json:"inputs"`
		Type      *ObservabilityPipelineOpenSearchDestinationType `json:"type"`
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
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bulk_index", "id", "inputs", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.BulkIndex = all.BulkIndex
	o.Id = *all.Id
	o.Inputs = *all.Inputs
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
