// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineAmazonOpenSearchDestinationAuth Authentication settings for the Amazon OpenSearch destination.
// The `strategy` field determines whether basic or AWS-based authentication is used.
type ObservabilityPipelineAmazonOpenSearchDestinationAuth struct {
	// The ARN of the role to assume (used with `aws` strategy).
	AssumeRole *string `json:"assume_role,omitempty"`
	// AWS region
	AwsRegion *string `json:"aws_region,omitempty"`
	// External ID for the assumed role (used with `aws` strategy).
	ExternalId *string `json:"external_id,omitempty"`
	// Session name for the assumed role (used with `aws` strategy).
	SessionName *string `json:"session_name,omitempty"`
	// The authentication strategy to use.
	Strategy ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy `json:"strategy"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineAmazonOpenSearchDestinationAuth instantiates a new ObservabilityPipelineAmazonOpenSearchDestinationAuth object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineAmazonOpenSearchDestinationAuth(strategy ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy) *ObservabilityPipelineAmazonOpenSearchDestinationAuth {
	this := ObservabilityPipelineAmazonOpenSearchDestinationAuth{}
	this.Strategy = strategy
	return &this
}

// NewObservabilityPipelineAmazonOpenSearchDestinationAuthWithDefaults instantiates a new ObservabilityPipelineAmazonOpenSearchDestinationAuth object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineAmazonOpenSearchDestinationAuthWithDefaults() *ObservabilityPipelineAmazonOpenSearchDestinationAuth {
	this := ObservabilityPipelineAmazonOpenSearchDestinationAuth{}
	return &this
}

// GetAssumeRole returns the AssumeRole field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) GetAssumeRole() string {
	if o == nil || o.AssumeRole == nil {
		var ret string
		return ret
	}
	return *o.AssumeRole
}

// GetAssumeRoleOk returns a tuple with the AssumeRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) GetAssumeRoleOk() (*string, bool) {
	if o == nil || o.AssumeRole == nil {
		return nil, false
	}
	return o.AssumeRole, true
}

// HasAssumeRole returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) HasAssumeRole() bool {
	return o != nil && o.AssumeRole != nil
}

// SetAssumeRole gets a reference to the given string and assigns it to the AssumeRole field.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) SetAssumeRole(v string) {
	o.AssumeRole = &v
}

// GetAwsRegion returns the AwsRegion field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) GetAwsRegion() string {
	if o == nil || o.AwsRegion == nil {
		var ret string
		return ret
	}
	return *o.AwsRegion
}

// GetAwsRegionOk returns a tuple with the AwsRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) GetAwsRegionOk() (*string, bool) {
	if o == nil || o.AwsRegion == nil {
		return nil, false
	}
	return o.AwsRegion, true
}

// HasAwsRegion returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) HasAwsRegion() bool {
	return o != nil && o.AwsRegion != nil
}

// SetAwsRegion gets a reference to the given string and assigns it to the AwsRegion field.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) SetAwsRegion(v string) {
	o.AwsRegion = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) HasExternalId() bool {
	return o != nil && o.ExternalId != nil
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetSessionName returns the SessionName field value if set, zero value otherwise.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) GetSessionName() string {
	if o == nil || o.SessionName == nil {
		var ret string
		return ret
	}
	return *o.SessionName
}

// GetSessionNameOk returns a tuple with the SessionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) GetSessionNameOk() (*string, bool) {
	if o == nil || o.SessionName == nil {
		return nil, false
	}
	return o.SessionName, true
}

// HasSessionName returns a boolean if a field has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) HasSessionName() bool {
	return o != nil && o.SessionName != nil
}

// SetSessionName gets a reference to the given string and assigns it to the SessionName field.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) SetSessionName(v string) {
	o.SessionName = &v
}

// GetStrategy returns the Strategy field value.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) GetStrategy() ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy {
	if o == nil {
		var ret ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy
		return ret
	}
	return o.Strategy
}

// GetStrategyOk returns a tuple with the Strategy field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) GetStrategyOk() (*ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Strategy, true
}

// SetStrategy sets field value.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) SetStrategy(v ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy) {
	o.Strategy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineAmazonOpenSearchDestinationAuth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssumeRole != nil {
		toSerialize["assume_role"] = o.AssumeRole
	}
	if o.AwsRegion != nil {
		toSerialize["aws_region"] = o.AwsRegion
	}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.SessionName != nil {
		toSerialize["session_name"] = o.SessionName
	}
	toSerialize["strategy"] = o.Strategy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineAmazonOpenSearchDestinationAuth) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssumeRole  *string                                                       `json:"assume_role,omitempty"`
		AwsRegion   *string                                                       `json:"aws_region,omitempty"`
		ExternalId  *string                                                       `json:"external_id,omitempty"`
		SessionName *string                                                       `json:"session_name,omitempty"`
		Strategy    *ObservabilityPipelineAmazonOpenSearchDestinationAuthStrategy `json:"strategy"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Strategy == nil {
		return fmt.Errorf("required field strategy missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assume_role", "aws_region", "external_id", "session_name", "strategy"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AssumeRole = all.AssumeRole
	o.AwsRegion = all.AwsRegion
	o.ExternalId = all.ExternalId
	o.SessionName = all.SessionName
	if !all.Strategy.IsValid() {
		hasInvalidField = true
	} else {
		o.Strategy = *all.Strategy
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
