// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineElasticsearchDestination The `elasticsearch` destination writes logs to an Elasticsearch cluster.
type ObservabilityPipelineElasticsearchDestination struct {
	// The Elasticsearch API version to use. Set to `auto` to auto-detect.
	ApiVersion *ObservabilityPipelineElasticsearchDestinationApiVersion `json:"api_version,omitempty"`
	// The index to write logs to in Elasticsearch.
	BulkIndex *string `json:"bulk_index,omitempty"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The destination type. The value should always be `elasticsearch`.
	Type ObservabilityPipelineElasticsearchDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineElasticsearchDestination instantiates a new ObservabilityPipelineElasticsearchDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineElasticsearchDestination(id string, inputs []string, typeVar ObservabilityPipelineElasticsearchDestinationType) *ObservabilityPipelineElasticsearchDestination {
	this := ObservabilityPipelineElasticsearchDestination{}
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineElasticsearchDestinationWithDefaults instantiates a new ObservabilityPipelineElasticsearchDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineElasticsearchDestinationWithDefaults() *ObservabilityPipelineElasticsearchDestination {
	this := ObservabilityPipelineElasticsearchDestination{}
	var typeVar ObservabilityPipelineElasticsearchDestinationType = OBSERVABILITYPIPELINEELASTICSEARCHDESTINATIONTYPE_ELASTICSEARCH
	this.Type = typeVar
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ObservabilityPipelineElasticsearchDestination) GetApiVersion() ObservabilityPipelineElasticsearchDestinationApiVersion {
	if o == nil || o.ApiVersion == nil {
		var ret ObservabilityPipelineElasticsearchDestinationApiVersion
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestination) GetApiVersionOk() (*ObservabilityPipelineElasticsearchDestinationApiVersion, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ObservabilityPipelineElasticsearchDestination) HasApiVersion() bool {
	return o != nil && o.ApiVersion != nil
}

// SetApiVersion gets a reference to the given ObservabilityPipelineElasticsearchDestinationApiVersion and assigns it to the ApiVersion field.
func (o *ObservabilityPipelineElasticsearchDestination) SetApiVersion(v ObservabilityPipelineElasticsearchDestinationApiVersion) {
	o.ApiVersion = &v
}

// GetBulkIndex returns the BulkIndex field value if set, zero value otherwise.
func (o *ObservabilityPipelineElasticsearchDestination) GetBulkIndex() string {
	if o == nil || o.BulkIndex == nil {
		var ret string
		return ret
	}
	return *o.BulkIndex
}

// GetBulkIndexOk returns a tuple with the BulkIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestination) GetBulkIndexOk() (*string, bool) {
	if o == nil || o.BulkIndex == nil {
		return nil, false
	}
	return o.BulkIndex, true
}

// HasBulkIndex returns a boolean if a field has been set.
func (o *ObservabilityPipelineElasticsearchDestination) HasBulkIndex() bool {
	return o != nil && o.BulkIndex != nil
}

// SetBulkIndex gets a reference to the given string and assigns it to the BulkIndex field.
func (o *ObservabilityPipelineElasticsearchDestination) SetBulkIndex(v string) {
	o.BulkIndex = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineElasticsearchDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineElasticsearchDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineElasticsearchDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineElasticsearchDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineElasticsearchDestination) GetType() ObservabilityPipelineElasticsearchDestinationType {
	if o == nil {
		var ret ObservabilityPipelineElasticsearchDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestination) GetTypeOk() (*ObservabilityPipelineElasticsearchDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineElasticsearchDestination) SetType(v ObservabilityPipelineElasticsearchDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineElasticsearchDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
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
func (o *ObservabilityPipelineElasticsearchDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiVersion *ObservabilityPipelineElasticsearchDestinationApiVersion `json:"api_version,omitempty"`
		BulkIndex  *string                                                  `json:"bulk_index,omitempty"`
		Id         *string                                                  `json:"id"`
		Inputs     *[]string                                                `json:"inputs"`
		Type       *ObservabilityPipelineElasticsearchDestinationType       `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"api_version", "bulk_index", "id", "inputs", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.ApiVersion != nil && !all.ApiVersion.IsValid() {
		hasInvalidField = true
	} else {
		o.ApiVersion = all.ApiVersion
	}
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
