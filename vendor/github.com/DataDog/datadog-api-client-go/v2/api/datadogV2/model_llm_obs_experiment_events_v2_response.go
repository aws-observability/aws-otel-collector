// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsExperimentEventsV2Response Response for listing experiment events (v2/v3). Returns spans and summary metrics in a single resource.
type LLMObsExperimentEventsV2Response struct {
	// JSON:API data object for an experiment events response.
	Data LLMObsExperimentEventsV2DataResponse `json:"data"`
	// Pagination cursor metadata.
	Meta *LLMObsCursorMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsExperimentEventsV2Response instantiates a new LLMObsExperimentEventsV2Response object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsExperimentEventsV2Response(data LLMObsExperimentEventsV2DataResponse) *LLMObsExperimentEventsV2Response {
	this := LLMObsExperimentEventsV2Response{}
	this.Data = data
	return &this
}

// NewLLMObsExperimentEventsV2ResponseWithDefaults instantiates a new LLMObsExperimentEventsV2Response object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsExperimentEventsV2ResponseWithDefaults() *LLMObsExperimentEventsV2Response {
	this := LLMObsExperimentEventsV2Response{}
	return &this
}

// GetData returns the Data field value.
func (o *LLMObsExperimentEventsV2Response) GetData() LLMObsExperimentEventsV2DataResponse {
	if o == nil {
		var ret LLMObsExperimentEventsV2DataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEventsV2Response) GetDataOk() (*LLMObsExperimentEventsV2DataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *LLMObsExperimentEventsV2Response) SetData(v LLMObsExperimentEventsV2DataResponse) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *LLMObsExperimentEventsV2Response) GetMeta() LLMObsCursorMeta {
	if o == nil || o.Meta == nil {
		var ret LLMObsCursorMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsExperimentEventsV2Response) GetMetaOk() (*LLMObsCursorMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *LLMObsExperimentEventsV2Response) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given LLMObsCursorMeta and assigns it to the Meta field.
func (o *LLMObsExperimentEventsV2Response) SetMeta(v LLMObsCursorMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsExperimentEventsV2Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsExperimentEventsV2Response) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *LLMObsExperimentEventsV2DataResponse `json:"data"`
		Meta *LLMObsCursorMeta                     `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
