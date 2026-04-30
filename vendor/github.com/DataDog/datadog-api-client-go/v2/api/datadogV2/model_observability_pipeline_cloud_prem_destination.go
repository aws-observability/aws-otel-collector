// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineCloudPremDestination The `cloud_prem` destination sends logs to Datadog CloudPrem.
//
// **Supported pipeline types:** logs
type ObservabilityPipelineCloudPremDestination struct {
	// Name of the environment variable or secret that holds the CloudPrem endpoint URL.
	EndpointUrlKey *string `json:"endpoint_url_key,omitempty"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The destination type. The value should always be `cloud_prem`.
	Type ObservabilityPipelineCloudPremDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineCloudPremDestination instantiates a new ObservabilityPipelineCloudPremDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineCloudPremDestination(id string, inputs []string, typeVar ObservabilityPipelineCloudPremDestinationType) *ObservabilityPipelineCloudPremDestination {
	this := ObservabilityPipelineCloudPremDestination{}
	this.Id = id
	this.Inputs = inputs
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineCloudPremDestinationWithDefaults instantiates a new ObservabilityPipelineCloudPremDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineCloudPremDestinationWithDefaults() *ObservabilityPipelineCloudPremDestination {
	this := ObservabilityPipelineCloudPremDestination{}
	var typeVar ObservabilityPipelineCloudPremDestinationType = OBSERVABILITYPIPELINECLOUDPREMDESTINATIONTYPE_CLOUD_PREM
	this.Type = typeVar
	return &this
}

// GetEndpointUrlKey returns the EndpointUrlKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineCloudPremDestination) GetEndpointUrlKey() string {
	if o == nil || o.EndpointUrlKey == nil {
		var ret string
		return ret
	}
	return *o.EndpointUrlKey
}

// GetEndpointUrlKeyOk returns a tuple with the EndpointUrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCloudPremDestination) GetEndpointUrlKeyOk() (*string, bool) {
	if o == nil || o.EndpointUrlKey == nil {
		return nil, false
	}
	return o.EndpointUrlKey, true
}

// HasEndpointUrlKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineCloudPremDestination) HasEndpointUrlKey() bool {
	return o != nil && o.EndpointUrlKey != nil
}

// SetEndpointUrlKey gets a reference to the given string and assigns it to the EndpointUrlKey field.
func (o *ObservabilityPipelineCloudPremDestination) SetEndpointUrlKey(v string) {
	o.EndpointUrlKey = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineCloudPremDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCloudPremDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineCloudPremDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineCloudPremDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCloudPremDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineCloudPremDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineCloudPremDestination) GetType() ObservabilityPipelineCloudPremDestinationType {
	if o == nil {
		var ret ObservabilityPipelineCloudPremDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineCloudPremDestination) GetTypeOk() (*ObservabilityPipelineCloudPremDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineCloudPremDestination) SetType(v ObservabilityPipelineCloudPremDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineCloudPremDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.EndpointUrlKey != nil {
		toSerialize["endpoint_url_key"] = o.EndpointUrlKey
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
func (o *ObservabilityPipelineCloudPremDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		EndpointUrlKey *string                                        `json:"endpoint_url_key,omitempty"`
		Id             *string                                        `json:"id"`
		Inputs         *[]string                                      `json:"inputs"`
		Type           *ObservabilityPipelineCloudPremDestinationType `json:"type"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"endpoint_url_key", "id", "inputs", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.EndpointUrlKey = all.EndpointUrlKey
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
