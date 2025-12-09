// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineSplunkHecDestination The `splunk_hec` destination forwards logs to Splunk using the HTTP Event Collector (HEC).
type ObservabilityPipelineSplunkHecDestination struct {
	// If `true`, Splunk tries to extract timestamps from incoming log events.
	// If `false`, Splunk assigns the time the event was received.
	AutoExtractTimestamp *bool `json:"auto_extract_timestamp,omitempty"`
	// Encoding format for log events.
	Encoding *ObservabilityPipelineSplunkHecDestinationEncoding `json:"encoding,omitempty"`
	// The unique identifier for this component. Used to reference this component in other parts of the pipeline (e.g., as input to downstream components).
	Id string `json:"id"`
	// Optional name of the Splunk index where logs are written.
	Index *string `json:"index,omitempty"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The Splunk sourcetype to assign to log events.
	Sourcetype *string `json:"sourcetype,omitempty"`
	// The destination type. Always `splunk_hec`.
	Type ObservabilityPipelineSplunkHecDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineSplunkHecDestination instantiates a new ObservabilityPipelineSplunkHecDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineSplunkHecDestination(id string, inputs []string, typeVar ObservabilityPipelineSplunkHecDestinationType) *ObservabilityPipelineSplunkHecDestination {
	this := ObservabilityPipelineSplunkHecDestination{}
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineSplunkHecDestinationWithDefaults instantiates a new ObservabilityPipelineSplunkHecDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineSplunkHecDestinationWithDefaults() *ObservabilityPipelineSplunkHecDestination {
	this := ObservabilityPipelineSplunkHecDestination{}
	var typeVar ObservabilityPipelineSplunkHecDestinationType = OBSERVABILITYPIPELINESPLUNKHECDESTINATIONTYPE_SPLUNK_HEC
	this.Type = typeVar
	return &this
}

// GetAutoExtractTimestamp returns the AutoExtractTimestamp field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecDestination) GetAutoExtractTimestamp() bool {
	if o == nil || o.AutoExtractTimestamp == nil {
		var ret bool
		return ret
	}
	return *o.AutoExtractTimestamp
}

// GetAutoExtractTimestampOk returns a tuple with the AutoExtractTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecDestination) GetAutoExtractTimestampOk() (*bool, bool) {
	if o == nil || o.AutoExtractTimestamp == nil {
		return nil, false
	}
	return o.AutoExtractTimestamp, true
}

// HasAutoExtractTimestamp returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecDestination) HasAutoExtractTimestamp() bool {
	return o != nil && o.AutoExtractTimestamp != nil
}

// SetAutoExtractTimestamp gets a reference to the given bool and assigns it to the AutoExtractTimestamp field.
func (o *ObservabilityPipelineSplunkHecDestination) SetAutoExtractTimestamp(v bool) {
	o.AutoExtractTimestamp = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecDestination) GetEncoding() ObservabilityPipelineSplunkHecDestinationEncoding {
	if o == nil || o.Encoding == nil {
		var ret ObservabilityPipelineSplunkHecDestinationEncoding
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecDestination) GetEncodingOk() (*ObservabilityPipelineSplunkHecDestinationEncoding, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecDestination) HasEncoding() bool {
	return o != nil && o.Encoding != nil
}

// SetEncoding gets a reference to the given ObservabilityPipelineSplunkHecDestinationEncoding and assigns it to the Encoding field.
func (o *ObservabilityPipelineSplunkHecDestination) SetEncoding(v ObservabilityPipelineSplunkHecDestinationEncoding) {
	o.Encoding = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineSplunkHecDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineSplunkHecDestination) SetId(v string) {
	o.Id = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecDestination) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecDestination) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecDestination) HasIndex() bool {
	return o != nil && o.Index != nil
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *ObservabilityPipelineSplunkHecDestination) SetIndex(v string) {
	o.Index = &v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineSplunkHecDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineSplunkHecDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetSourcetype returns the Sourcetype field value if set, zero value otherwise.
func (o *ObservabilityPipelineSplunkHecDestination) GetSourcetype() string {
	if o == nil || o.Sourcetype == nil {
		var ret string
		return ret
	}
	return *o.Sourcetype
}

// GetSourcetypeOk returns a tuple with the Sourcetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecDestination) GetSourcetypeOk() (*string, bool) {
	if o == nil || o.Sourcetype == nil {
		return nil, false
	}
	return o.Sourcetype, true
}

// HasSourcetype returns a boolean if a field has been set.
func (o *ObservabilityPipelineSplunkHecDestination) HasSourcetype() bool {
	return o != nil && o.Sourcetype != nil
}

// SetSourcetype gets a reference to the given string and assigns it to the Sourcetype field.
func (o *ObservabilityPipelineSplunkHecDestination) SetSourcetype(v string) {
	o.Sourcetype = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineSplunkHecDestination) GetType() ObservabilityPipelineSplunkHecDestinationType {
	if o == nil {
		var ret ObservabilityPipelineSplunkHecDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineSplunkHecDestination) GetTypeOk() (*ObservabilityPipelineSplunkHecDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineSplunkHecDestination) SetType(v ObservabilityPipelineSplunkHecDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineSplunkHecDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutoExtractTimestamp != nil {
		toSerialize["auto_extract_timestamp"] = o.AutoExtractTimestamp
	}
	if o.Encoding != nil {
		toSerialize["encoding"] = o.Encoding
	}
	toSerialize["id"] = o.Id
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	toSerialize["inputs"] = o.Inputs
	if o.Sourcetype != nil {
		toSerialize["sourcetype"] = o.Sourcetype
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineSplunkHecDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoExtractTimestamp *bool                                              `json:"auto_extract_timestamp,omitempty"`
		Encoding             *ObservabilityPipelineSplunkHecDestinationEncoding `json:"encoding,omitempty"`
		Id                   *string                                            `json:"id"`
		Index                *string                                            `json:"index,omitempty"`
		Inputs               *[]string                                          `json:"inputs"`
		Sourcetype           *string                                            `json:"sourcetype,omitempty"`
		Type                 *ObservabilityPipelineSplunkHecDestinationType     `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_extract_timestamp", "encoding", "id", "index", "inputs", "sourcetype", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AutoExtractTimestamp = all.AutoExtractTimestamp
	if all.Encoding != nil && !all.Encoding.IsValid() {
		hasInvalidField = true
	} else {
		o.Encoding = all.Encoding
	}
	o.Id = *all.Id
	o.Index = all.Index
	o.Inputs = *all.Inputs
	o.Sourcetype = all.Sourcetype
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
