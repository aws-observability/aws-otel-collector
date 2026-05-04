// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OnCallNotificationRuleRelationships Relationship object for creating a notification rule
type OnCallNotificationRuleRelationships struct {
	// Relationship object for creating a notification rule
	Channel *OnCallNotificationRuleChannelRelationship `json:"channel,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOnCallNotificationRuleRelationships instantiates a new OnCallNotificationRuleRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOnCallNotificationRuleRelationships() *OnCallNotificationRuleRelationships {
	this := OnCallNotificationRuleRelationships{}
	return &this
}

// NewOnCallNotificationRuleRelationshipsWithDefaults instantiates a new OnCallNotificationRuleRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOnCallNotificationRuleRelationshipsWithDefaults() *OnCallNotificationRuleRelationships {
	this := OnCallNotificationRuleRelationships{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *OnCallNotificationRuleRelationships) GetChannel() OnCallNotificationRuleChannelRelationship {
	if o == nil || o.Channel == nil {
		var ret OnCallNotificationRuleChannelRelationship
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OnCallNotificationRuleRelationships) GetChannelOk() (*OnCallNotificationRuleChannelRelationship, bool) {
	if o == nil || o.Channel == nil {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *OnCallNotificationRuleRelationships) HasChannel() bool {
	return o != nil && o.Channel != nil
}

// SetChannel gets a reference to the given OnCallNotificationRuleChannelRelationship and assigns it to the Channel field.
func (o *OnCallNotificationRuleRelationships) SetChannel(v OnCallNotificationRuleChannelRelationship) {
	o.Channel = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OnCallNotificationRuleRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Channel != nil {
		toSerialize["channel"] = o.Channel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OnCallNotificationRuleRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Channel *OnCallNotificationRuleChannelRelationship `json:"channel,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channel"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Channel != nil && all.Channel.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Channel = all.Channel

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
