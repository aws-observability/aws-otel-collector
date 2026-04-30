// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableFieldEventLookup Looks up a value from a field path in the log event.
type ObservabilityPipelineEnrichmentTableFieldEventLookup struct {
	// The path to the field in the log event to use as the lookup key.
	Event string `json:"event"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineEnrichmentTableFieldEventLookup instantiates a new ObservabilityPipelineEnrichmentTableFieldEventLookup object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineEnrichmentTableFieldEventLookup(event string) *ObservabilityPipelineEnrichmentTableFieldEventLookup {
	this := ObservabilityPipelineEnrichmentTableFieldEventLookup{}
	this.Event = event
	return &this
}

// NewObservabilityPipelineEnrichmentTableFieldEventLookupWithDefaults instantiates a new ObservabilityPipelineEnrichmentTableFieldEventLookup object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineEnrichmentTableFieldEventLookupWithDefaults() *ObservabilityPipelineEnrichmentTableFieldEventLookup {
	this := ObservabilityPipelineEnrichmentTableFieldEventLookup{}
	return &this
}

// GetEvent returns the Event field value.
func (o *ObservabilityPipelineEnrichmentTableFieldEventLookup) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableFieldEventLookup) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value.
func (o *ObservabilityPipelineEnrichmentTableFieldEventLookup) SetEvent(v string) {
	o.Event = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineEnrichmentTableFieldEventLookup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["event"] = o.Event

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineEnrichmentTableFieldEventLookup) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Event *string `json:"event"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Event == nil {
		return fmt.Errorf("required field event missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"event"})
	} else {
		return err
	}
	o.Event = *all.Event

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
