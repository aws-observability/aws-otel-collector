// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDatadogLogsDestination The `datadog_logs` destination forwards logs to Datadog Log Management.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineDatadogLogsDestination struct {
	// Configuration for buffer settings on destination components.
	Buffer *ObservabilityPipelineBufferOptions `json:"buffer,omitempty"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// A list of routing rules that forward matching logs to Datadog using dedicated API keys.
	Routes []ObservabilityPipelineDatadogLogsDestinationRoute `json:"routes,omitempty"`
	// The destination type. The value should always be `datadog_logs`.
	Type ObservabilityPipelineDatadogLogsDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineDatadogLogsDestination instantiates a new ObservabilityPipelineDatadogLogsDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineDatadogLogsDestination(id string, inputs []string, typeVar ObservabilityPipelineDatadogLogsDestinationType) *ObservabilityPipelineDatadogLogsDestination {
	this := ObservabilityPipelineDatadogLogsDestination{}
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineDatadogLogsDestinationWithDefaults instantiates a new ObservabilityPipelineDatadogLogsDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineDatadogLogsDestinationWithDefaults() *ObservabilityPipelineDatadogLogsDestination {
	this := ObservabilityPipelineDatadogLogsDestination{}
	var typeVar ObservabilityPipelineDatadogLogsDestinationType = OBSERVABILITYPIPELINEDATADOGLOGSDESTINATIONTYPE_DATADOG_LOGS
	this.Type = typeVar
	return &this
}

// GetBuffer returns the Buffer field value if set, zero value otherwise.
func (o *ObservabilityPipelineDatadogLogsDestination) GetBuffer() ObservabilityPipelineBufferOptions {
	if o == nil || o.Buffer == nil {
		var ret ObservabilityPipelineBufferOptions
		return ret
	}
	return *o.Buffer
}

// GetBufferOk returns a tuple with the Buffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogLogsDestination) GetBufferOk() (*ObservabilityPipelineBufferOptions, bool) {
	if o == nil || o.Buffer == nil {
		return nil, false
	}
	return o.Buffer, true
}

// HasBuffer returns a boolean if a field has been set.
func (o *ObservabilityPipelineDatadogLogsDestination) HasBuffer() bool {
	return o != nil && o.Buffer != nil
}

// SetBuffer gets a reference to the given ObservabilityPipelineBufferOptions and assigns it to the Buffer field.
func (o *ObservabilityPipelineDatadogLogsDestination) SetBuffer(v ObservabilityPipelineBufferOptions) {
	o.Buffer = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineDatadogLogsDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogLogsDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineDatadogLogsDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineDatadogLogsDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogLogsDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineDatadogLogsDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetRoutes returns the Routes field value if set, zero value otherwise.
func (o *ObservabilityPipelineDatadogLogsDestination) GetRoutes() []ObservabilityPipelineDatadogLogsDestinationRoute {
	if o == nil || o.Routes == nil {
		var ret []ObservabilityPipelineDatadogLogsDestinationRoute
		return ret
	}
	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogLogsDestination) GetRoutesOk() (*[]ObservabilityPipelineDatadogLogsDestinationRoute, bool) {
	if o == nil || o.Routes == nil {
		return nil, false
	}
	return &o.Routes, true
}

// HasRoutes returns a boolean if a field has been set.
func (o *ObservabilityPipelineDatadogLogsDestination) HasRoutes() bool {
	return o != nil && o.Routes != nil
}

// SetRoutes gets a reference to the given []ObservabilityPipelineDatadogLogsDestinationRoute and assigns it to the Routes field.
func (o *ObservabilityPipelineDatadogLogsDestination) SetRoutes(v []ObservabilityPipelineDatadogLogsDestinationRoute) {
	o.Routes = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineDatadogLogsDestination) GetType() ObservabilityPipelineDatadogLogsDestinationType {
	if o == nil {
		var ret ObservabilityPipelineDatadogLogsDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatadogLogsDestination) GetTypeOk() (*ObservabilityPipelineDatadogLogsDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineDatadogLogsDestination) SetType(v ObservabilityPipelineDatadogLogsDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineDatadogLogsDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Buffer != nil {
		toSerialize["buffer"] = o.Buffer
	}
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	if o.Routes != nil {
		toSerialize["routes"] = o.Routes
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineDatadogLogsDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Buffer *ObservabilityPipelineBufferOptions                `json:"buffer,omitempty"`
		Id     *string                                            `json:"id"`
		Inputs *[]string                                          `json:"inputs"`
		Routes []ObservabilityPipelineDatadogLogsDestinationRoute `json:"routes,omitempty"`
		Type   *ObservabilityPipelineDatadogLogsDestinationType   `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"buffer", "id", "inputs", "routes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Buffer = all.Buffer
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	o.Routes = all.Routes
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
