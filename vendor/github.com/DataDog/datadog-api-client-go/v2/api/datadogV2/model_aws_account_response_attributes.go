// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSAccountResponseAttributes AWS Account response attributes.
type AWSAccountResponseAttributes struct {
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
	// Timestamp of when the account integration was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// AWS Logs Collection config.
	LogsConfig *AWSLogsConfig `json:"logs_config,omitempty"`
	// AWS Metrics Collection config.
	MetricsConfig *AWSMetricsConfig `json:"metrics_config,omitempty"`
	// Timestamp of when the account integration was updated.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// AWS Resources Collection config.
	ResourcesConfig *AWSResourcesConfig `json:"resources_config,omitempty"`
	// AWS Traces Collection config.
	TracesConfig *AWSTracesConfig `json:"traces_config,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSAccountResponseAttributes instantiates a new AWSAccountResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSAccountResponseAttributes(awsAccountId string) *AWSAccountResponseAttributes {
	this := AWSAccountResponseAttributes{}
	this.AwsAccountId = awsAccountId
	return &this
}

// NewAWSAccountResponseAttributesWithDefaults instantiates a new AWSAccountResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSAccountResponseAttributesWithDefaults() *AWSAccountResponseAttributes {
	this := AWSAccountResponseAttributes{}
	return &this
}

// GetAccountTags returns the AccountTags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AWSAccountResponseAttributes) GetAccountTags() []string {
	if o == nil || o.AccountTags.Get() == nil {
		var ret []string
		return ret
	}
	return *o.AccountTags.Get()
}

// GetAccountTagsOk returns a tuple with the AccountTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AWSAccountResponseAttributes) GetAccountTagsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AccountTags.Get(), o.AccountTags.IsSet()
}

// HasAccountTags returns a boolean if a field has been set.
func (o *AWSAccountResponseAttributes) HasAccountTags() bool {
	return o != nil && o.AccountTags.IsSet()
}

// SetAccountTags gets a reference to the given datadog.NullableList[string] and assigns it to the AccountTags field.
func (o *AWSAccountResponseAttributes) SetAccountTags(v []string) {
	o.AccountTags.Set(&v)
}

// SetAccountTagsNil sets the value for AccountTags to be an explicit nil.
func (o *AWSAccountResponseAttributes) SetAccountTagsNil() {
	o.AccountTags.Set(nil)
}

// UnsetAccountTags ensures that no value is present for AccountTags, not even an explicit nil.
func (o *AWSAccountResponseAttributes) UnsetAccountTags() {
	o.AccountTags.Unset()
}

// GetAuthConfig returns the AuthConfig field value if set, zero value otherwise.
func (o *AWSAccountResponseAttributes) GetAuthConfig() AWSAuthConfig {
	if o == nil || o.AuthConfig == nil {
		var ret AWSAuthConfig
		return ret
	}
	return *o.AuthConfig
}

// GetAuthConfigOk returns a tuple with the AuthConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountResponseAttributes) GetAuthConfigOk() (*AWSAuthConfig, bool) {
	if o == nil || o.AuthConfig == nil {
		return nil, false
	}
	return o.AuthConfig, true
}

// HasAuthConfig returns a boolean if a field has been set.
func (o *AWSAccountResponseAttributes) HasAuthConfig() bool {
	return o != nil && o.AuthConfig != nil
}

// SetAuthConfig gets a reference to the given AWSAuthConfig and assigns it to the AuthConfig field.
func (o *AWSAccountResponseAttributes) SetAuthConfig(v AWSAuthConfig) {
	o.AuthConfig = &v
}

// GetAwsAccountId returns the AwsAccountId field value.
func (o *AWSAccountResponseAttributes) GetAwsAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AwsAccountId
}

// GetAwsAccountIdOk returns a tuple with the AwsAccountId field value
// and a boolean to check if the value has been set.
func (o *AWSAccountResponseAttributes) GetAwsAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AwsAccountId, true
}

// SetAwsAccountId sets field value.
func (o *AWSAccountResponseAttributes) SetAwsAccountId(v string) {
	o.AwsAccountId = v
}

// GetAwsPartition returns the AwsPartition field value if set, zero value otherwise.
func (o *AWSAccountResponseAttributes) GetAwsPartition() AWSAccountPartition {
	if o == nil || o.AwsPartition == nil {
		var ret AWSAccountPartition
		return ret
	}
	return *o.AwsPartition
}

// GetAwsPartitionOk returns a tuple with the AwsPartition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountResponseAttributes) GetAwsPartitionOk() (*AWSAccountPartition, bool) {
	if o == nil || o.AwsPartition == nil {
		return nil, false
	}
	return o.AwsPartition, true
}

// HasAwsPartition returns a boolean if a field has been set.
func (o *AWSAccountResponseAttributes) HasAwsPartition() bool {
	return o != nil && o.AwsPartition != nil
}

// SetAwsPartition gets a reference to the given AWSAccountPartition and assigns it to the AwsPartition field.
func (o *AWSAccountResponseAttributes) SetAwsPartition(v AWSAccountPartition) {
	o.AwsPartition = &v
}

// GetAwsRegions returns the AwsRegions field value if set, zero value otherwise.
func (o *AWSAccountResponseAttributes) GetAwsRegions() AWSRegions {
	if o == nil || o.AwsRegions == nil {
		var ret AWSRegions
		return ret
	}
	return *o.AwsRegions
}

// GetAwsRegionsOk returns a tuple with the AwsRegions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountResponseAttributes) GetAwsRegionsOk() (*AWSRegions, bool) {
	if o == nil || o.AwsRegions == nil {
		return nil, false
	}
	return o.AwsRegions, true
}

// HasAwsRegions returns a boolean if a field has been set.
func (o *AWSAccountResponseAttributes) HasAwsRegions() bool {
	return o != nil && o.AwsRegions != nil
}

// SetAwsRegions gets a reference to the given AWSRegions and assigns it to the AwsRegions field.
func (o *AWSAccountResponseAttributes) SetAwsRegions(v AWSRegions) {
	o.AwsRegions = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AWSAccountResponseAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AWSAccountResponseAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AWSAccountResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetLogsConfig returns the LogsConfig field value if set, zero value otherwise.
func (o *AWSAccountResponseAttributes) GetLogsConfig() AWSLogsConfig {
	if o == nil || o.LogsConfig == nil {
		var ret AWSLogsConfig
		return ret
	}
	return *o.LogsConfig
}

// GetLogsConfigOk returns a tuple with the LogsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountResponseAttributes) GetLogsConfigOk() (*AWSLogsConfig, bool) {
	if o == nil || o.LogsConfig == nil {
		return nil, false
	}
	return o.LogsConfig, true
}

// HasLogsConfig returns a boolean if a field has been set.
func (o *AWSAccountResponseAttributes) HasLogsConfig() bool {
	return o != nil && o.LogsConfig != nil
}

// SetLogsConfig gets a reference to the given AWSLogsConfig and assigns it to the LogsConfig field.
func (o *AWSAccountResponseAttributes) SetLogsConfig(v AWSLogsConfig) {
	o.LogsConfig = &v
}

// GetMetricsConfig returns the MetricsConfig field value if set, zero value otherwise.
func (o *AWSAccountResponseAttributes) GetMetricsConfig() AWSMetricsConfig {
	if o == nil || o.MetricsConfig == nil {
		var ret AWSMetricsConfig
		return ret
	}
	return *o.MetricsConfig
}

// GetMetricsConfigOk returns a tuple with the MetricsConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountResponseAttributes) GetMetricsConfigOk() (*AWSMetricsConfig, bool) {
	if o == nil || o.MetricsConfig == nil {
		return nil, false
	}
	return o.MetricsConfig, true
}

// HasMetricsConfig returns a boolean if a field has been set.
func (o *AWSAccountResponseAttributes) HasMetricsConfig() bool {
	return o != nil && o.MetricsConfig != nil
}

// SetMetricsConfig gets a reference to the given AWSMetricsConfig and assigns it to the MetricsConfig field.
func (o *AWSAccountResponseAttributes) SetMetricsConfig(v AWSMetricsConfig) {
	o.MetricsConfig = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *AWSAccountResponseAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountResponseAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *AWSAccountResponseAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *AWSAccountResponseAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetResourcesConfig returns the ResourcesConfig field value if set, zero value otherwise.
func (o *AWSAccountResponseAttributes) GetResourcesConfig() AWSResourcesConfig {
	if o == nil || o.ResourcesConfig == nil {
		var ret AWSResourcesConfig
		return ret
	}
	return *o.ResourcesConfig
}

// GetResourcesConfigOk returns a tuple with the ResourcesConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountResponseAttributes) GetResourcesConfigOk() (*AWSResourcesConfig, bool) {
	if o == nil || o.ResourcesConfig == nil {
		return nil, false
	}
	return o.ResourcesConfig, true
}

// HasResourcesConfig returns a boolean if a field has been set.
func (o *AWSAccountResponseAttributes) HasResourcesConfig() bool {
	return o != nil && o.ResourcesConfig != nil
}

// SetResourcesConfig gets a reference to the given AWSResourcesConfig and assigns it to the ResourcesConfig field.
func (o *AWSAccountResponseAttributes) SetResourcesConfig(v AWSResourcesConfig) {
	o.ResourcesConfig = &v
}

// GetTracesConfig returns the TracesConfig field value if set, zero value otherwise.
func (o *AWSAccountResponseAttributes) GetTracesConfig() AWSTracesConfig {
	if o == nil || o.TracesConfig == nil {
		var ret AWSTracesConfig
		return ret
	}
	return *o.TracesConfig
}

// GetTracesConfigOk returns a tuple with the TracesConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSAccountResponseAttributes) GetTracesConfigOk() (*AWSTracesConfig, bool) {
	if o == nil || o.TracesConfig == nil {
		return nil, false
	}
	return o.TracesConfig, true
}

// HasTracesConfig returns a boolean if a field has been set.
func (o *AWSAccountResponseAttributes) HasTracesConfig() bool {
	return o != nil && o.TracesConfig != nil
}

// SetTracesConfig gets a reference to the given AWSTracesConfig and assigns it to the TracesConfig field.
func (o *AWSAccountResponseAttributes) SetTracesConfig(v AWSTracesConfig) {
	o.TracesConfig = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSAccountResponseAttributes) MarshalJSON() ([]byte, error) {
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
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.LogsConfig != nil {
		toSerialize["logs_config"] = o.LogsConfig
	}
	if o.MetricsConfig != nil {
		toSerialize["metrics_config"] = o.MetricsConfig
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
func (o *AWSAccountResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountTags     datadog.NullableList[string] `json:"account_tags,omitempty"`
		AuthConfig      *AWSAuthConfig               `json:"auth_config,omitempty"`
		AwsAccountId    *string                      `json:"aws_account_id"`
		AwsPartition    *AWSAccountPartition         `json:"aws_partition,omitempty"`
		AwsRegions      *AWSRegions                  `json:"aws_regions,omitempty"`
		CreatedAt       *time.Time                   `json:"created_at,omitempty"`
		LogsConfig      *AWSLogsConfig               `json:"logs_config,omitempty"`
		MetricsConfig   *AWSMetricsConfig            `json:"metrics_config,omitempty"`
		ModifiedAt      *time.Time                   `json:"modified_at,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"account_tags", "auth_config", "aws_account_id", "aws_partition", "aws_regions", "created_at", "logs_config", "metrics_config", "modified_at", "resources_config", "traces_config"})
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
	o.CreatedAt = all.CreatedAt
	if all.LogsConfig != nil && all.LogsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LogsConfig = all.LogsConfig
	if all.MetricsConfig != nil && all.MetricsConfig.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MetricsConfig = all.MetricsConfig
	o.ModifiedAt = all.ModifiedAt
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
