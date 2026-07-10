// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsTriggerResponseAttributes Attributes of an LLM Observability patterns trigger response.
type LLMObsPatternsTriggerResponseAttributes struct {
	// The ID of the patterns configuration that was run.
	ConfigId string `json:"config_id"`
	// The ID of the patterns run that was started.
	RunId string `json:"run_id"`
	// Status of the patterns run.
	Status string `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsTriggerResponseAttributes instantiates a new LLMObsPatternsTriggerResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsTriggerResponseAttributes(configId string, runId string, status string) *LLMObsPatternsTriggerResponseAttributes {
	this := LLMObsPatternsTriggerResponseAttributes{}
	this.ConfigId = configId
	this.RunId = runId
	this.Status = status
	return &this
}

// NewLLMObsPatternsTriggerResponseAttributesWithDefaults instantiates a new LLMObsPatternsTriggerResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsTriggerResponseAttributesWithDefaults() *LLMObsPatternsTriggerResponseAttributes {
	this := LLMObsPatternsTriggerResponseAttributes{}
	return &this
}

// GetConfigId returns the ConfigId field value.
func (o *LLMObsPatternsTriggerResponseAttributes) GetConfigId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConfigId
}

// GetConfigIdOk returns a tuple with the ConfigId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTriggerResponseAttributes) GetConfigIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConfigId, true
}

// SetConfigId sets field value.
func (o *LLMObsPatternsTriggerResponseAttributes) SetConfigId(v string) {
	o.ConfigId = v
}

// GetRunId returns the RunId field value.
func (o *LLMObsPatternsTriggerResponseAttributes) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTriggerResponseAttributes) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *LLMObsPatternsTriggerResponseAttributes) SetRunId(v string) {
	o.RunId = v
}

// GetStatus returns the Status field value.
func (o *LLMObsPatternsTriggerResponseAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTriggerResponseAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *LLMObsPatternsTriggerResponseAttributes) SetStatus(v string) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsTriggerResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["config_id"] = o.ConfigId
	toSerialize["run_id"] = o.RunId
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsTriggerResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConfigId *string `json:"config_id"`
		RunId    *string `json:"run_id"`
		Status   *string `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConfigId == nil {
		return fmt.Errorf("required field config_id missing")
	}
	if all.RunId == nil {
		return fmt.Errorf("required field run_id missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"config_id", "run_id", "status"})
	} else {
		return err
	}
	o.ConfigId = *all.ConfigId
	o.RunId = *all.RunId
	o.Status = *all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
