// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAccountUpdateRequestAttributes The AWS Account Integration Config to be updated.
type AWSAccountUpdateRequestAttributes struct {
	// Tags to apply to all hosts and metrics reporting for this account. Defaults to `[]`.
	AccountTags datadog.NullableList[string] `json:"account_tags,omitempty"`
	// AWS Authentication config.
	AuthConfig *AWSAuthConfig `json:"auth_config,omitempty"`
	// AWS Account ID.
	AwsAccountId string `json:"aws_account_id"`
	// AWS partition your AWS account is scoped to. Defaults to `aws`.
	// See [Partitions](https://docs.aws.amazon.com/whitepapers/latest/aws-fault-isolation-boundaries/partitions.html) in the AWS documentation for more information.
	AwsPartition *AWSAccountPartition `json:"aws_partition,omitempty"`
	// AWS Regions to collect data from. Defaults to `include_all`.
	AwsRegions *AWSRegions `json:"aws_regions,omitempty"`
	// AWS Logs Collection config.
	LogsConfig *AWSLogsConfig `json:"logs_config,omitempty"`
	// AWS Metrics Collection config.
	MetricsConfig *AWSMetricsConfig `json:"metrics_config,omitempty"`
	// AWS Resources Collection config.
	ResourcesConfig *AWSResourcesConfig `json:"resources_config,omitempty"`
	// AWS Traces Collection config.
	TracesConfig *AWSTracesConfig `json:"traces_config,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSAccountUpdateRequestAttributes instantiates a new AWSAccountUpdateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSAccountUpdateRequestAttributes(awsAccountId string) *AWSAccountUpdateRequestAttributes {
	this := AWSAccountUpdateRequestAttributes{}
	this.AwsAccountId = awsAccountId
	return &this
}

// NewAWSAccountUpdateRequestAttributesWithDefaults instantiates a new AWSAccountUpdateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSAccountUpdateRequestAttributesWithDefaults() *AWSAccountUpdateRequestAttributes {
	this := AWSAccountUpdateRequestAttributes{}
	return &this
}

// GetAccountTags returns the AccountTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSAccountUpdateRequestAttributes) GetAccountTags() []string {
	if o == nil || o.AccountTags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.AccountTags.Get()
}

// GetAccountTagsOk returns a tuple with the AccountTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AWSAccountUpdateRequestAttributes) GetAccountTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountTags.Get(), o.AccountTags.IsSet()
}

// HasAccountTags returns a boolean if a field has been set.
func (o *AWSAccountUpdateRequestAttributes) HasAccountTags() bool {
	return o != nil && o.AccountTags.IsSet()
}

// SetAccountTags gets a reference to the given datadog.NullableList[string] and assigns it to the AccountTags field.
func (o *AWSAccountUpdateRequestAttributes) SetAccountTags(v []string) {
	o.AccountTags.Set(&v)
}

// SetAccountTagsNil sets the value for AccountTags to be an explicit nil.
func (o *AWSAccountUpdateRequestAttributes) SetAccountTagsNil() {
	o.AccountTags.Set(nil)
}

// UnsetAccountTags ensures that no value is present for AccountTags, not even an explicit nil.
func (o *AWSAccountUpdateRequestAttributes) UnsetAccountTags() {
	o.AccountTags.Unset()
}

// GetAuthConfig returns the AuthConfig field value if set, zero value otherwise.
func (o *AWSAccountUpdateRequestAttributes) GetAuthConfig() AWSAuthConfig {
	if o == nil || o.AuthConfig == nil {
		var ret AWSAuthConfig
		return ret
	}
	return *o.AuthConfig
}

// GetAuthConfigOk returns a tuple with the AuthConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountUpdateRequestAttributes) GetAuthConfigOk() (*AWSAuthConfig, bool) {
	if o == nil || o.AuthConfig == nil {
		return nil, false
	}
	return o.AuthConfig, true
}

// HasAuthConfig returns a boolean if a field has been set.
func (o *AWSAccountUpdateRequestAttributes) HasAuthConfig() bool {
	return o != nil && o.AuthConfig != nil
}

// SetAuthConfig gets a reference to the given AWSAuthConfig and assigns it to the AuthConfig field.
func (o *AWSAccountUpdateRequestAttributes) SetAuthConfig(v AWSAuthConfig) {
	o.AuthConfig = &v
}

// GetAwsAccountId returns the AwsAccountId field value.
func (o *AWSAccountUpdateRequestAttributes) GetAwsAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value
// and a boolean to check if the value has been set.
func (o *AWSAccountUpdateRequestAttributes) GetAwsAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountId, true
}

// SetAwsAccountId sets field value.
func (o *AWSAccountUpdateRequestAttributes) SetAwsAccountId(v string) {
	o.AwsAccountId = v
}

// GetAwsPartition returns the AwsPartition field value if set, zero value otherwise.
func (o *AWSAccountUpdateRequestAttributes) GetAwsPartition() AWSAccountPartition {
	if o == nil || o.AwsPartition == nil {
		var ret AWSAccountPartition
		return ret
	}
	return *o.AwsPartition
}

// GetAwsPartitionOk returns a tuple with the AwsPartition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountUpdateRequestAttributes) GetAwsPartitionOk() (*AWSAccountPartition, bool) {
	if o == nil || o.AwsPartition == nil {
		return nil, false
	}
	return o.AwsPartition, true
}

// HasAwsPartition returns a boolean if a field has been set.
func (o *AWSAccountUpdateRequestAttributes) HasAwsPartition() bool {
	return o != nil && o.AwsPartition != nil
}

// SetAwsPartition gets a reference to the given AWSAccountPartition and assigns it to the AwsPartition field.
func (o *AWSAccountUpdateRequestAttributes) SetAwsPartition(v AWSAccountPartition) {
	o.AwsPartition = &v
}

// GetAwsRegions returns the AwsRegions field value if set, zero value otherwise.
func (o *AWSAccountUpdateRequestAttributes) GetAwsRegions() AWSRegions {
	if o == nil || o.AwsRegions == nil {
		var ret AWSRegions
		return ret
	}
	return *o.AwsRegions
}

// GetAwsRegionsOk returns a tuple with the AwsRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountUpdateRequestAttributes) GetAwsRegionsOk() (*AWSRegions, bool) {
	if o == nil || o.AwsRegions == nil {
		return nil, false
	}
	return o.AwsRegions, true
}

// HasAwsRegions returns a boolean if a field has been set.
func (o *AWSAccountUpdateRequestAttributes) HasAwsRegions() bool {
	return o != nil && o.AwsRegions != nil
}

// SetAwsRegions gets a reference to the given AWSRegions and assigns it to the AwsRegions field.
func (o *AWSAccountUpdateRequestAttributes) SetAwsRegions(v AWSRegions) {
	o.AwsRegions = &v
}

// GetLogsConfig returns the LogsConfig field value if set, zero value otherwise.
func (o *AWSAccountUpdateRequestAttributes) GetLogsConfig() AWSLogsConfig {
	if o == nil || o.LogsConfig == nil {
		var ret AWSLogsConfig
		return ret
	}
	return *o.LogsConfig
}

// GetLogsConfigOk returns a tuple with the LogsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountUpdateRequestAttributes) GetLogsConfigOk() (*AWSLogsConfig, bool) {
	if o == nil || o.LogsConfig == nil {
		return nil, false
	}
	return o.LogsConfig, true
}

// HasLogsConfig returns a boolean if a field has been set.
func (o *AWSAccountUpdateRequestAttributes) HasLogsConfig() bool {
	return o != nil && o.LogsConfig != nil
}

// SetLogsConfig gets a reference to the given AWSLogsConfig and assigns it to the LogsConfig field.
func (o *AWSAccountUpdateRequestAttributes) SetLogsConfig(v AWSLogsConfig) {
	o.LogsConfig = &v
}

// GetMetricsConfig returns the MetricsConfig field value if set, zero value otherwise.
func (o *AWSAccountUpdateRequestAttributes) GetMetricsConfig() AWSMetricsConfig {
	if o == nil || o.MetricsConfig == nil {
		var ret AWSMetricsConfig
		return ret
	}
	return *o.MetricsConfig
}

// GetMetricsConfigOk returns a tuple with the MetricsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountUpdateRequestAttributes) GetMetricsConfigOk() (*AWSMetricsConfig, bool) {
	if o == nil || o.MetricsConfig == nil {
		return nil, false
	}
	return o.MetricsConfig, true
}

// HasMetricsConfig returns a boolean if a field has been set.
func (o *AWSAccountUpdateRequestAttributes) HasMetricsConfig() bool {
	return o != nil && o.MetricsConfig != nil
}

// SetMetricsConfig gets a reference to the given AWSMetricsConfig and assigns it to the MetricsConfig field.
func (o *AWSAccountUpdateRequestAttributes) SetMetricsConfig(v AWSMetricsConfig) {
	o.MetricsConfig = &v
}

// GetResourcesConfig returns the ResourcesConfig field value if set, zero value otherwise.
func (o *AWSAccountUpdateRequestAttributes) GetResourcesConfig() AWSResourcesConfig {
	if o == nil || o.ResourcesConfig == nil {
		var ret AWSResourcesConfig
		return ret
	}
	return *o.ResourcesConfig
}

// GetResourcesConfigOk returns a tuple with the ResourcesConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountUpdateRequestAttributes) GetResourcesConfigOk() (*AWSResourcesConfig, bool) {
	if o == nil || o.ResourcesConfig == nil {
		return nil, false
	}
	return o.ResourcesConfig, true
}

// HasResourcesConfig returns a boolean if a field has been set.
func (o *AWSAccountUpdateRequestAttributes) HasResourcesConfig() bool {
	return o != nil && o.ResourcesConfig != nil
}

// SetResourcesConfig gets a reference to the given AWSResourcesConfig and assigns it to the ResourcesConfig field.
func (o *AWSAccountUpdateRequestAttributes) SetResourcesConfig(v AWSResourcesConfig) {
	o.ResourcesConfig = &v
}

// GetTracesConfig returns the TracesConfig field value if set, zero value otherwise.
func (o *AWSAccountUpdateRequestAttributes) GetTracesConfig() AWSTracesConfig {
	if o == nil || o.TracesConfig == nil {
		var ret AWSTracesConfig
		return ret
	}
	return *o.TracesConfig
}

// GetTracesConfigOk returns a tuple with the TracesConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountUpdateRequestAttributes) GetTracesConfigOk() (*AWSTracesConfig, bool) {
	if o == nil || o.TracesConfig == nil {
		return nil, false
	}
	return o.TracesConfig, true
}

// HasTracesConfig returns a boolean if a field has been set.
func (o *AWSAccountUpdateRequestAttributes) HasTracesConfig() bool {
	return o != nil && o.TracesConfig != nil
}

// SetTracesConfig gets a reference to the given AWSTracesConfig and assigns it to the TracesConfig field.
func (o *AWSAccountUpdateRequestAttributes) SetTracesConfig(v AWSTracesConfig) {
	o.TracesConfig = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSAccountUpdateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountTags.IsSet() {
		toSerialize["account_tags"] = o.AccountTags.Get()
	}
	if o.AuthConfig != nil {
		toSerialize["auth_config"] = o.AuthConfig
	}
	toSerialize["aws_account_id"] = o.AwsAccountId
	if o.AwsPartition != nil {
		toSerialize["aws_partition"] = o.AwsPartition
	}
	if o.AwsRegions != nil {
		toSerialize["aws_regions"] = o.AwsRegions
	}
	if o.LogsConfig != nil {
		toSerialize["logs_config"] = o.LogsConfig
	}
	if o.MetricsConfig != nil {
		toSerialize["metrics_config"] = o.MetricsConfig
	}
	if o.ResourcesConfig != nil {
		toSerialize["resources_config"] = o.ResourcesConfig
	}
	if o.TracesConfig != nil {
		toSerialize["traces_config"] = o.TracesConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSAccountUpdateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountTags     datadog.NullableList[string] `json:"account_tags,omitempty"`
		AuthConfig      *AWSAuthConfig               `json:"auth_config,omitempty"`
		AwsAccountId    *string                      `json:"aws_account_id"`
		AwsPartition    *AWSAccountPartition         `json:"aws_partition,omitempty"`
		AwsRegions      *AWSRegions                  `json:"aws_regions,omitempty"`
		LogsConfig      *AWSLogsConfig               `json:"logs_config,omitempty"`
		MetricsConfig   *AWSMetricsConfig            `json:"metrics_config,omitempty"`
		ResourcesConfig *AWSResourcesConfig          `json:"resources_config,omitempty"`
		TracesConfig    *AWSTracesConfig             `json:"traces_config,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AwsAccountId == nil {
		return fmt.Errorf("required field aws_account_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_tags", "auth_config", "aws_account_id", "aws_partition", "aws_regions", "logs_config", "metrics_config", "resources_config", "traces_config"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccountTags = all.AccountTags
	o.AuthConfig = all.AuthConfig
	o.AwsAccountId = *all.AwsAccountId
	if all.AwsPartition != nil && !all.AwsPartition.IsValid() {
		hasInvalidField = true
	} else {
		o.AwsPartition = all.AwsPartition
	}
	o.AwsRegions = all.AwsRegions
	if all.LogsConfig != nil && all.LogsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogsConfig = all.LogsConfig
	if all.MetricsConfig != nil && all.MetricsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricsConfig = all.MetricsConfig
	if all.ResourcesConfig != nil && all.ResourcesConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.ResourcesConfig = all.ResourcesConfig
	if all.TracesConfig != nil && all.TracesConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TracesConfig = all.TracesConfig

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
