// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSEventBridgeDeleteRequestAttributes The EventBridge source to be deleted.
type AWSEventBridgeDeleteRequestAttributes struct {
	// AWS Account ID.
	AccountId string `json:"account_id"`
	// The event source name.
	EventGeneratorName string `json:"event_generator_name"`
	// The event source's
	// [AWS region](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints).
	Region string `json:"region"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSEventBridgeDeleteRequestAttributes instantiates a new AWSEventBridgeDeleteRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSEventBridgeDeleteRequestAttributes(accountId string, eventGeneratorName string, region string) *AWSEventBridgeDeleteRequestAttributes {
	this := AWSEventBridgeDeleteRequestAttributes{}
	this.AccountId = accountId
	this.EventGeneratorName = eventGeneratorName
	this.Region = region
	return &this
}

// NewAWSEventBridgeDeleteRequestAttributesWithDefaults instantiates a new AWSEventBridgeDeleteRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSEventBridgeDeleteRequestAttributesWithDefaults() *AWSEventBridgeDeleteRequestAttributes {
	this := AWSEventBridgeDeleteRequestAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *AWSEventBridgeDeleteRequestAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeDeleteRequestAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *AWSEventBridgeDeleteRequestAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetEventGeneratorName returns the EventGeneratorName field value.
func (o *AWSEventBridgeDeleteRequestAttributes) GetEventGeneratorName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EventGeneratorName
}

// GetEventGeneratorNameOk returns a tuple with the EventGeneratorName field value
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeDeleteRequestAttributes) GetEventGeneratorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventGeneratorName, true
}

// SetEventGeneratorName sets field value.
func (o *AWSEventBridgeDeleteRequestAttributes) SetEventGeneratorName(v string) {
	o.EventGeneratorName = v
}

// GetRegion returns the Region field value.
func (o *AWSEventBridgeDeleteRequestAttributes) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeDeleteRequestAttributes) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *AWSEventBridgeDeleteRequestAttributes) SetRegion(v string) {
	o.Region = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSEventBridgeDeleteRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	toSerialize["event_generator_name"] = o.EventGeneratorName
	toSerialize["region"] = o.Region

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSEventBridgeDeleteRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId          *string `json:"account_id"`
		EventGeneratorName *string `json:"event_generator_name"`
		Region             *string `json:"region"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccountId == nil {
		return fmt.Errorf("required field account_id missing")
	}
	if all.EventGeneratorName == nil {
		return fmt.Errorf("required field event_generator_name missing")
	}
	if all.Region == nil {
		return fmt.Errorf("required field region missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "event_generator_name", "region"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.EventGeneratorName = *all.EventGeneratorName
	o.Region = *all.Region

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
