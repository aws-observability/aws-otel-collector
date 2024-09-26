// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsPipelinesOrder Object containing the ordered list of pipeline IDs.
type LogsPipelinesOrder struct {
	// Ordered Array of `<PIPELINE_ID>` strings, the order of pipeline IDs in the array
	// define the overall Pipelines order for Datadog.
	PipelineIds []string `json:"pipeline_ids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsPipelinesOrder instantiates a new LogsPipelinesOrder object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsPipelinesOrder(pipelineIds []string) *LogsPipelinesOrder {
	this := LogsPipelinesOrder{}
	this.PipelineIds = pipelineIds
	return &this
}

// NewLogsPipelinesOrderWithDefaults instantiates a new LogsPipelinesOrder object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsPipelinesOrderWithDefaults() *LogsPipelinesOrder {
	this := LogsPipelinesOrder{}
	return &this
}

// GetPipelineIds returns the PipelineIds field value.
func (o *LogsPipelinesOrder) GetPipelineIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PipelineIds
}

// GetPipelineIdsOk returns a tuple with the PipelineIds field value
// and a boolean to check if the value has been set.
func (o *LogsPipelinesOrder) GetPipelineIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PipelineIds, true
}

// SetPipelineIds sets field value.
func (o *LogsPipelinesOrder) SetPipelineIds(v []string) {
	o.PipelineIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsPipelinesOrder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["pipeline_ids"] = o.PipelineIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsPipelinesOrder) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PipelineIds *[]string `json:"pipeline_ids"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PipelineIds == nil {
		return fmt.Errorf("required field pipeline_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"pipeline_ids"})
	} else {
		return err
	}
	o.PipelineIds = *all.PipelineIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
