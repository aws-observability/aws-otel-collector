// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSLambdaForwarderConfig Log Autosubscription configuration for Datadog Forwarder Lambda functions.
// Automatically set up triggers for existing and new logs for some services,
// ensuring no logs from new resources are missed and saving time spent on manual configuration.
type AWSLambdaForwarderConfig struct {
	// List of Datadog Lambda Log Forwarder ARNs in your AWS account. Defaults to `[]`.
	Lambdas []string `json:"lambdas,omitempty"`
	// Log source configuration.
	LogSourceConfig *AWSLambdaForwarderConfigLogSourceConfig `json:"log_source_config,omitempty"`
	// List of service IDs set to enable automatic log collection.
	// Discover the list of available services with the
	// [Get list of AWS log ready services](https://docs.datadoghq.com/api/latest/aws-logs-integration/#get-list-of-aws-log-ready-services)
	// endpoint.
	Sources []string `json:"sources,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSLambdaForwarderConfig instantiates a new AWSLambdaForwarderConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSLambdaForwarderConfig() *AWSLambdaForwarderConfig {
	this := AWSLambdaForwarderConfig{}
	return &this
}

// NewAWSLambdaForwarderConfigWithDefaults instantiates a new AWSLambdaForwarderConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSLambdaForwarderConfigWithDefaults() *AWSLambdaForwarderConfig {
	this := AWSLambdaForwarderConfig{}
	return &this
}

// GetLambdas returns the Lambdas field value if set, zero value otherwise.
func (o *AWSLambdaForwarderConfig) GetLambdas() []string {
	if o == nil || o.Lambdas == nil {
		var ret []string
		return ret
	}
	return o.Lambdas
}

// GetLambdasOk returns a tuple with the Lambdas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSLambdaForwarderConfig) GetLambdasOk() (*[]string, bool) {
	if o == nil || o.Lambdas == nil {
		return nil, false
	}
	return &o.Lambdas, true
}

// HasLambdas returns a boolean if a field has been set.
func (o *AWSLambdaForwarderConfig) HasLambdas() bool {
	return o != nil && o.Lambdas != nil
}

// SetLambdas gets a reference to the given []string and assigns it to the Lambdas field.
func (o *AWSLambdaForwarderConfig) SetLambdas(v []string) {
	o.Lambdas = v
}

// GetLogSourceConfig returns the LogSourceConfig field value if set, zero value otherwise.
func (o *AWSLambdaForwarderConfig) GetLogSourceConfig() AWSLambdaForwarderConfigLogSourceConfig {
	if o == nil || o.LogSourceConfig == nil {
		var ret AWSLambdaForwarderConfigLogSourceConfig
		return ret
	}
	return *o.LogSourceConfig
}

// GetLogSourceConfigOk returns a tuple with the LogSourceConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSLambdaForwarderConfig) GetLogSourceConfigOk() (*AWSLambdaForwarderConfigLogSourceConfig, bool) {
	if o == nil || o.LogSourceConfig == nil {
		return nil, false
	}
	return o.LogSourceConfig, true
}

// HasLogSourceConfig returns a boolean if a field has been set.
func (o *AWSLambdaForwarderConfig) HasLogSourceConfig() bool {
	return o != nil && o.LogSourceConfig != nil
}

// SetLogSourceConfig gets a reference to the given AWSLambdaForwarderConfigLogSourceConfig and assigns it to the LogSourceConfig field.
func (o *AWSLambdaForwarderConfig) SetLogSourceConfig(v AWSLambdaForwarderConfigLogSourceConfig) {
	o.LogSourceConfig = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *AWSLambdaForwarderConfig) GetSources() []string {
	if o == nil || o.Sources == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSLambdaForwarderConfig) GetSourcesOk() (*[]string, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return &o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *AWSLambdaForwarderConfig) HasSources() bool {
	return o != nil && o.Sources != nil
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *AWSLambdaForwarderConfig) SetSources(v []string) {
	o.Sources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSLambdaForwarderConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Lambdas != nil {
		toSerialize["lambdas"] = o.Lambdas
	}
	if o.LogSourceConfig != nil {
		toSerialize["log_source_config"] = o.LogSourceConfig
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSLambdaForwarderConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Lambdas         []string                                 `json:"lambdas,omitempty"`
		LogSourceConfig *AWSLambdaForwarderConfigLogSourceConfig `json:"log_source_config,omitempty"`
		Sources         []string                                 `json:"sources,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"lambdas", "log_source_config", "sources"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Lambdas = all.Lambdas
	if all.LogSourceConfig != nil && all.LogSourceConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogSourceConfig = all.LogSourceConfig
	o.Sources = all.Sources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
