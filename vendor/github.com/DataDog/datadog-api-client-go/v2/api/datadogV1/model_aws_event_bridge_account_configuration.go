// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSEventBridgeAccountConfiguration The EventBridge configuration for one AWS account.
type AWSEventBridgeAccountConfiguration struct {
	// Your AWS Account ID without dashes.
	AccountId *string `json:"accountId,omitempty"`
	// Array of AWS event sources associated with this account.
	EventHubs []AWSEventBridgeSource `json:"eventHubs,omitempty"`
	// Array of tags (in the form `key:value`) which are added to all hosts
	// and metrics reporting through the main AWS integration.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSEventBridgeAccountConfiguration instantiates a new AWSEventBridgeAccountConfiguration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSEventBridgeAccountConfiguration() *AWSEventBridgeAccountConfiguration {
	this := AWSEventBridgeAccountConfiguration{}
	return &this
}

// NewAWSEventBridgeAccountConfigurationWithDefaults instantiates a new AWSEventBridgeAccountConfiguration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSEventBridgeAccountConfigurationWithDefaults() *AWSEventBridgeAccountConfiguration {
	this := AWSEventBridgeAccountConfiguration{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AWSEventBridgeAccountConfiguration) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeAccountConfiguration) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AWSEventBridgeAccountConfiguration) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AWSEventBridgeAccountConfiguration) SetAccountId(v string) {
	o.AccountId = &v
}

// GetEventHubs returns the EventHubs field value if set, zero value otherwise.
func (o *AWSEventBridgeAccountConfiguration) GetEventHubs() []AWSEventBridgeSource {
	if o == nil || o.EventHubs == nil {
		var ret []AWSEventBridgeSource
		return ret
	}
	return o.EventHubs
}

// GetEventHubsOk returns a tuple with the EventHubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeAccountConfiguration) GetEventHubsOk() (*[]AWSEventBridgeSource, bool) {
	if o == nil || o.EventHubs == nil {
		return nil, false
	}
	return &o.EventHubs, true
}

// HasEventHubs returns a boolean if a field has been set.
func (o *AWSEventBridgeAccountConfiguration) HasEventHubs() bool {
	return o != nil && o.EventHubs != nil
}

// SetEventHubs gets a reference to the given []AWSEventBridgeSource and assigns it to the EventHubs field.
func (o *AWSEventBridgeAccountConfiguration) SetEventHubs(v []AWSEventBridgeSource) {
	o.EventHubs = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AWSEventBridgeAccountConfiguration) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeAccountConfiguration) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AWSEventBridgeAccountConfiguration) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AWSEventBridgeAccountConfiguration) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSEventBridgeAccountConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["accountId"] = o.AccountId
	}
	if o.EventHubs != nil {
		toSerialize["eventHubs"] = o.EventHubs
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSEventBridgeAccountConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId *string                `json:"accountId,omitempty"`
		EventHubs []AWSEventBridgeSource `json:"eventHubs,omitempty"`
		Tags      []string               `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"accountId", "eventHubs", "tags"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.EventHubs = all.EventHubs
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
