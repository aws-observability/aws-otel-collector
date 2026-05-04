// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultTurnStep A step executed during a goal-based browser test turn.
type SyntheticsTestResultTurnStep struct {
	// Storage bucket keys for artifacts produced during a step or test.
	BucketKeys *SyntheticsTestResultBucketKeys `json:"bucket_keys,omitempty"`
	// Browser step configuration for this turn step.
	Config map[string]interface{} `json:"config,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultTurnStep instantiates a new SyntheticsTestResultTurnStep object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultTurnStep() *SyntheticsTestResultTurnStep {
	this := SyntheticsTestResultTurnStep{}
	return &this
}

// NewSyntheticsTestResultTurnStepWithDefaults instantiates a new SyntheticsTestResultTurnStep object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultTurnStepWithDefaults() *SyntheticsTestResultTurnStep {
	this := SyntheticsTestResultTurnStep{}
	return &this
}

// GetBucketKeys returns the BucketKeys field value if set, zero value otherwise.
func (o *SyntheticsTestResultTurnStep) GetBucketKeys() SyntheticsTestResultBucketKeys {
	if o == nil || o.BucketKeys == nil {
		var ret SyntheticsTestResultBucketKeys
		return ret
	}
	return *o.BucketKeys
}

// GetBucketKeysOk returns a tuple with the BucketKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTurnStep) GetBucketKeysOk() (*SyntheticsTestResultBucketKeys, bool) {
	if o == nil || o.BucketKeys == nil {
		return nil, false
	}
	return o.BucketKeys, true
}

// HasBucketKeys returns a boolean if a field has been set.
func (o *SyntheticsTestResultTurnStep) HasBucketKeys() bool {
	return o != nil && o.BucketKeys != nil
}

// SetBucketKeys gets a reference to the given SyntheticsTestResultBucketKeys and assigns it to the BucketKeys field.
func (o *SyntheticsTestResultTurnStep) SetBucketKeys(v SyntheticsTestResultBucketKeys) {
	o.BucketKeys = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *SyntheticsTestResultTurnStep) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultTurnStep) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return &o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *SyntheticsTestResultTurnStep) HasConfig() bool {
	return o != nil && o.Config != nil
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *SyntheticsTestResultTurnStep) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultTurnStep) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BucketKeys != nil {
		toSerialize["bucket_keys"] = o.BucketKeys
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultTurnStep) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BucketKeys *SyntheticsTestResultBucketKeys `json:"bucket_keys,omitempty"`
		Config     map[string]interface{}          `json:"config,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bucket_keys", "config"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.BucketKeys != nil && all.BucketKeys.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BucketKeys = all.BucketKeys
	o.Config = all.Config

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
