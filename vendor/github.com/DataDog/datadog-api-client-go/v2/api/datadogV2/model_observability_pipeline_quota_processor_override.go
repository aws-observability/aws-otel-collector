// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineQuotaProcessorOverride Defines a custom quota limit that applies to specific log events based on matching field values.
type ObservabilityPipelineQuotaProcessorOverride struct {
	// A list of field matchers used to apply a specific override. If an event matches all listed key-value pairs, the corresponding override limit is enforced.
	Fields []ObservabilityPipelineFieldValue `json:"fields"`
	// The maximum amount of data or number of events allowed before the quota is enforced. Can be specified in bytes or events.
	Limit ObservabilityPipelineQuotaProcessorLimit `json:"limit"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineQuotaProcessorOverride instantiates a new ObservabilityPipelineQuotaProcessorOverride object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineQuotaProcessorOverride(fields []ObservabilityPipelineFieldValue, limit ObservabilityPipelineQuotaProcessorLimit) *ObservabilityPipelineQuotaProcessorOverride {
	this := ObservabilityPipelineQuotaProcessorOverride{}
	this.Fields = fields
	this.Limit = limit
	return &this
}

// NewObservabilityPipelineQuotaProcessorOverrideWithDefaults instantiates a new ObservabilityPipelineQuotaProcessorOverride object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineQuotaProcessorOverrideWithDefaults() *ObservabilityPipelineQuotaProcessorOverride {
	this := ObservabilityPipelineQuotaProcessorOverride{}
	return &this
}

// GetFields returns the Fields field value.
func (o *ObservabilityPipelineQuotaProcessorOverride) GetFields() []ObservabilityPipelineFieldValue {
	if o == nil {
		var ret []ObservabilityPipelineFieldValue
		return ret
	}
	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessorOverride) GetFieldsOk() (*[]ObservabilityPipelineFieldValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fields, true
}

// SetFields sets field value.
func (o *ObservabilityPipelineQuotaProcessorOverride) SetFields(v []ObservabilityPipelineFieldValue) {
	o.Fields = v
}

// GetLimit returns the Limit field value.
func (o *ObservabilityPipelineQuotaProcessorOverride) GetLimit() ObservabilityPipelineQuotaProcessorLimit {
	if o == nil {
		var ret ObservabilityPipelineQuotaProcessorLimit
		return ret
	}
	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineQuotaProcessorOverride) GetLimitOk() (*ObservabilityPipelineQuotaProcessorLimit, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value.
func (o *ObservabilityPipelineQuotaProcessorOverride) SetLimit(v ObservabilityPipelineQuotaProcessorLimit) {
	o.Limit = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineQuotaProcessorOverride) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["fields"] = o.Fields
	toSerialize["limit"] = o.Limit

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineQuotaProcessorOverride) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Fields *[]ObservabilityPipelineFieldValue        `json:"fields"`
		Limit  *ObservabilityPipelineQuotaProcessorLimit `json:"limit"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Fields == nil {
		return fmt.Errorf("required field fields missing")
	}
	if all.Limit == nil {
		return fmt.Errorf("required field limit missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"fields", "limit"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Fields = *all.Fields
	if all.Limit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Limit = *all.Limit

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
