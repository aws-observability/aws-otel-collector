// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSLogsConfig AWS Logs Collection config.
type AWSLogsConfig struct {
	// Log Autosubscription configuration for Datadog Forwarder Lambda functions. Automatically set up triggers for existing
	// and new logs for some services, ensuring no logs from new resources are missed and saving time spent on manual configuration.
	LambdaForwarder *AWSLambdaForwarderConfig `json:"lambda_forwarder,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSLogsConfig instantiates a new AWSLogsConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSLogsConfig() *AWSLogsConfig {
	this := AWSLogsConfig{}
	return &this
}

// NewAWSLogsConfigWithDefaults instantiates a new AWSLogsConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSLogsConfigWithDefaults() *AWSLogsConfig {
	this := AWSLogsConfig{}
	return &this
}

// GetLambdaForwarder returns the LambdaForwarder field value if set, zero value otherwise.
func (o *AWSLogsConfig) GetLambdaForwarder() AWSLambdaForwarderConfig {
	if o == nil || o.LambdaForwarder == nil {
		var ret AWSLambdaForwarderConfig
		return ret
	}
	return *o.LambdaForwarder
}

// GetLambdaForwarderOk returns a tuple with the LambdaForwarder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSLogsConfig) GetLambdaForwarderOk() (*AWSLambdaForwarderConfig, bool) {
	if o == nil || o.LambdaForwarder == nil {
		return nil, false
	}
	return o.LambdaForwarder, true
}

// HasLambdaForwarder returns a boolean if a field has been set.
func (o *AWSLogsConfig) HasLambdaForwarder() bool {
	return o != nil && o.LambdaForwarder != nil
}

// SetLambdaForwarder gets a reference to the given AWSLambdaForwarderConfig and assigns it to the LambdaForwarder field.
func (o *AWSLogsConfig) SetLambdaForwarder(v AWSLambdaForwarderConfig) {
	o.LambdaForwarder = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSLogsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.LambdaForwarder != nil {
		toSerialize["lambda_forwarder"] = o.LambdaForwarder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSLogsConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		LambdaForwarder *AWSLambdaForwarderConfig `json:"lambda_forwarder,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"lambda_forwarder"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.LambdaForwarder != nil && all.LambdaForwarder.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LambdaForwarder = all.LambdaForwarder

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
