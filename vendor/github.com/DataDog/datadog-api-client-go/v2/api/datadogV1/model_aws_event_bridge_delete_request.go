// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSEventBridgeDeleteRequest An object used to delete an EventBridge source.
type AWSEventBridgeDeleteRequest struct {
	// Your AWS Account ID without dashes.
	AccountId *string `json:"account_id,omitempty"`
	// The event source name.
	EventGeneratorName *string `json:"event_generator_name,omitempty"`
	// The event source's [AWS region](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints).
	Region *string `json:"region,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSEventBridgeDeleteRequest instantiates a new AWSEventBridgeDeleteRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSEventBridgeDeleteRequest() *AWSEventBridgeDeleteRequest {
	this := AWSEventBridgeDeleteRequest{}
	return &this
}

// NewAWSEventBridgeDeleteRequestWithDefaults instantiates a new AWSEventBridgeDeleteRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSEventBridgeDeleteRequestWithDefaults() *AWSEventBridgeDeleteRequest {
	this := AWSEventBridgeDeleteRequest{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *AWSEventBridgeDeleteRequest) GetAccountId() string {
	if o == nil || o.AccountId == nil {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeDeleteRequest) GetAccountIdOk() (*string, bool) {
	if o == nil || o.AccountId == nil {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *AWSEventBridgeDeleteRequest) HasAccountId() bool {
	return o != nil && o.AccountId != nil
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *AWSEventBridgeDeleteRequest) SetAccountId(v string) {
	o.AccountId = &v
}

// GetEventGeneratorName returns the EventGeneratorName field value if set, zero value otherwise.
func (o *AWSEventBridgeDeleteRequest) GetEventGeneratorName() string {
	if o == nil || o.EventGeneratorName == nil {
		var ret string
		return ret
	}
	return *o.EventGeneratorName
}

// GetEventGeneratorNameOk returns a tuple with the EventGeneratorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeDeleteRequest) GetEventGeneratorNameOk() (*string, bool) {
	if o == nil || o.EventGeneratorName == nil {
		return nil, false
	}
	return o.EventGeneratorName, true
}

// HasEventGeneratorName returns a boolean if a field has been set.
func (o *AWSEventBridgeDeleteRequest) HasEventGeneratorName() bool {
	return o != nil && o.EventGeneratorName != nil
}

// SetEventGeneratorName gets a reference to the given string and assigns it to the EventGeneratorName field.
func (o *AWSEventBridgeDeleteRequest) SetEventGeneratorName(v string) {
	o.EventGeneratorName = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AWSEventBridgeDeleteRequest) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeDeleteRequest) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AWSEventBridgeDeleteRequest) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AWSEventBridgeDeleteRequest) SetRegion(v string) {
	o.Region = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSEventBridgeDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccountId != nil {
		toSerialize["account_id"] = o.AccountId
	}
	if o.EventGeneratorName != nil {
		toSerialize["event_generator_name"] = o.EventGeneratorName
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSEventBridgeDeleteRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId          *string `json:"account_id,omitempty"`
		EventGeneratorName *string `json:"event_generator_name,omitempty"`
		Region             *string `json:"region,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "event_generator_name", "region"})
	} else {
		return err
	}
	o.AccountId = all.AccountId
	o.EventGeneratorName = all.EventGeneratorName
	o.Region = all.Region

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
