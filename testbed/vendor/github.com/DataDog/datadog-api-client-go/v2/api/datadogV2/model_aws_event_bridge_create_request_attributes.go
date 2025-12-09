// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSEventBridgeCreateRequestAttributes The EventBridge source to be created.
type AWSEventBridgeCreateRequestAttributes struct {
	// AWS Account ID.
	AccountId string `json:"account_id"`
	// Set to true if Datadog should create the event bus in addition to the event
	// source. Requires the `events:CreateEventBus` permission.
	CreateEventBus *bool `json:"create_event_bus,omitempty"`
	// The given part of the event source name, which is then combined with an
	// assigned suffix to form the full name.
	EventGeneratorName string `json:"event_generator_name"`
	// The event source's
	// [AWS region](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints).
	Region string `json:"region"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSEventBridgeCreateRequestAttributes instantiates a new AWSEventBridgeCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSEventBridgeCreateRequestAttributes(accountId string, eventGeneratorName string, region string) *AWSEventBridgeCreateRequestAttributes {
	this := AWSEventBridgeCreateRequestAttributes{}
	this.AccountId = accountId
	this.EventGeneratorName = eventGeneratorName
	this.Region = region
	return &this
}

// NewAWSEventBridgeCreateRequestAttributesWithDefaults instantiates a new AWSEventBridgeCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSEventBridgeCreateRequestAttributesWithDefaults() *AWSEventBridgeCreateRequestAttributes {
	this := AWSEventBridgeCreateRequestAttributes{}
	return &this
}

// GetAccountId returns the AccountId field value.
func (o *AWSEventBridgeCreateRequestAttributes) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeCreateRequestAttributes) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value.
func (o *AWSEventBridgeCreateRequestAttributes) SetAccountId(v string) {
	o.AccountId = v
}

// GetCreateEventBus returns the CreateEventBus field value if set, zero value otherwise.
func (o *AWSEventBridgeCreateRequestAttributes) GetCreateEventBus() bool {
	if o == nil || o.CreateEventBus == nil {
		var ret bool
		return ret
	}
	return *o.CreateEventBus
}

// GetCreateEventBusOk returns a tuple with the CreateEventBus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeCreateRequestAttributes) GetCreateEventBusOk() (*bool, bool) {
	if o == nil || o.CreateEventBus == nil {
		return nil, false
	}
	return o.CreateEventBus, true
}

// HasCreateEventBus returns a boolean if a field has been set.
func (o *AWSEventBridgeCreateRequestAttributes) HasCreateEventBus() bool {
	return o != nil && o.CreateEventBus != nil
}

// SetCreateEventBus gets a reference to the given bool and assigns it to the CreateEventBus field.
func (o *AWSEventBridgeCreateRequestAttributes) SetCreateEventBus(v bool) {
	o.CreateEventBus = &v
}

// GetEventGeneratorName returns the EventGeneratorName field value.
func (o *AWSEventBridgeCreateRequestAttributes) GetEventGeneratorName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.EventGeneratorName
}

// GetEventGeneratorNameOk returns a tuple with the EventGeneratorName field value
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeCreateRequestAttributes) GetEventGeneratorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventGeneratorName, true
}

// SetEventGeneratorName sets field value.
func (o *AWSEventBridgeCreateRequestAttributes) SetEventGeneratorName(v string) {
	o.EventGeneratorName = v
}

// GetRegion returns the Region field value.
func (o *AWSEventBridgeCreateRequestAttributes) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *AWSEventBridgeCreateRequestAttributes) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value.
func (o *AWSEventBridgeCreateRequestAttributes) SetRegion(v string) {
	o.Region = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSEventBridgeCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["account_id"] = o.AccountId
	if o.CreateEventBus != nil {
		toSerialize["create_event_bus"] = o.CreateEventBus
	}
	toSerialize["event_generator_name"] = o.EventGeneratorName
	toSerialize["region"] = o.Region

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSEventBridgeCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccountId          *string `json:"account_id"`
		CreateEventBus     *bool   `json:"create_event_bus,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"account_id", "create_event_bus", "event_generator_name", "region"})
	} else {
		return err
	}
	o.AccountId = *all.AccountId
	o.CreateEventBus = all.CreateEventBus
	o.EventGeneratorName = *all.EventGeneratorName
	o.Region = *all.Region

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
