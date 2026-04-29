// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableFieldVrlLookup Evaluates a VRL expression to produce the lookup key.
type ObservabilityPipelineEnrichmentTableFieldVrlLookup struct {
	// A VRL expression that returns the value to use as the lookup key.
	Vrl string `json:"vrl"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineEnrichmentTableFieldVrlLookup instantiates a new ObservabilityPipelineEnrichmentTableFieldVrlLookup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineEnrichmentTableFieldVrlLookup(vrl string) *ObservabilityPipelineEnrichmentTableFieldVrlLookup {
	this := ObservabilityPipelineEnrichmentTableFieldVrlLookup{}
	this.Vrl = vrl
	return &this
}

// NewObservabilityPipelineEnrichmentTableFieldVrlLookupWithDefaults instantiates a new ObservabilityPipelineEnrichmentTableFieldVrlLookup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineEnrichmentTableFieldVrlLookupWithDefaults() *ObservabilityPipelineEnrichmentTableFieldVrlLookup {
	this := ObservabilityPipelineEnrichmentTableFieldVrlLookup{}
	return &this
}

// GetVrl returns the Vrl field value.
func (o *ObservabilityPipelineEnrichmentTableFieldVrlLookup) GetVrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Vrl
}

// GetVrlOk returns a tuple with the Vrl field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableFieldVrlLookup) GetVrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vrl, true
}

// SetVrl sets field value.
func (o *ObservabilityPipelineEnrichmentTableFieldVrlLookup) SetVrl(v string) {
	o.Vrl = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineEnrichmentTableFieldVrlLookup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["vrl"] = o.Vrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineEnrichmentTableFieldVrlLookup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Vrl *string `json:"vrl"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Vrl == nil {
		return fmt.Errorf("required field vrl missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"vrl"})
	} else {
		return err
	}
	o.Vrl = *all.Vrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
