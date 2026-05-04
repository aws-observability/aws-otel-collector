// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JiraAccountAttributes Attributes of a Jira account
type JiraAccountAttributes struct {
	// The consumer key for the Jira account
	ConsumerKey string `json:"consumer_key"`
	// The URL of the Jira instance
	InstanceUrl string `json:"instance_url"`
	// Timestamp of the last webhook received
	LastWebhookTimestamp *time.Time `json:"last_webhook_timestamp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJiraAccountAttributes instantiates a new JiraAccountAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJiraAccountAttributes(consumerKey string, instanceUrl string) *JiraAccountAttributes {
	this := JiraAccountAttributes{}
	this.ConsumerKey = consumerKey
	this.InstanceUrl = instanceUrl
	return &this
}

// NewJiraAccountAttributesWithDefaults instantiates a new JiraAccountAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJiraAccountAttributesWithDefaults() *JiraAccountAttributes {
	this := JiraAccountAttributes{}
	return &this
}

// GetConsumerKey returns the ConsumerKey field value.
func (o *JiraAccountAttributes) GetConsumerKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ConsumerKey
}

// GetConsumerKeyOk returns a tuple with the ConsumerKey field value
// and a boolean to check if the value has been set.
func (o *JiraAccountAttributes) GetConsumerKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerKey, true
}

// SetConsumerKey sets field value.
func (o *JiraAccountAttributes) SetConsumerKey(v string) {
	o.ConsumerKey = v
}

// GetInstanceUrl returns the InstanceUrl field value.
func (o *JiraAccountAttributes) GetInstanceUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InstanceUrl
}

// GetInstanceUrlOk returns a tuple with the InstanceUrl field value
// and a boolean to check if the value has been set.
func (o *JiraAccountAttributes) GetInstanceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceUrl, true
}

// SetInstanceUrl sets field value.
func (o *JiraAccountAttributes) SetInstanceUrl(v string) {
	o.InstanceUrl = v
}

// GetLastWebhookTimestamp returns the LastWebhookTimestamp field value if set, zero value otherwise.
func (o *JiraAccountAttributes) GetLastWebhookTimestamp() time.Time {
	if o == nil || o.LastWebhookTimestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.LastWebhookTimestamp
}

// GetLastWebhookTimestampOk returns a tuple with the LastWebhookTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JiraAccountAttributes) GetLastWebhookTimestampOk() (*time.Time, bool) {
	if o == nil || o.LastWebhookTimestamp == nil {
		return nil, false
	}
	return o.LastWebhookTimestamp, true
}

// HasLastWebhookTimestamp returns a boolean if a field has been set.
func (o *JiraAccountAttributes) HasLastWebhookTimestamp() bool {
	return o != nil && o.LastWebhookTimestamp != nil
}

// SetLastWebhookTimestamp gets a reference to the given time.Time and assigns it to the LastWebhookTimestamp field.
func (o *JiraAccountAttributes) SetLastWebhookTimestamp(v time.Time) {
	o.LastWebhookTimestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o JiraAccountAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["consumer_key"] = o.ConsumerKey
	toSerialize["instance_url"] = o.InstanceUrl
	if o.LastWebhookTimestamp != nil {
		if o.LastWebhookTimestamp.Nanosecond() == 0 {
			toSerialize["last_webhook_timestamp"] = o.LastWebhookTimestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_webhook_timestamp"] = o.LastWebhookTimestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JiraAccountAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConsumerKey          *string    `json:"consumer_key"`
		InstanceUrl          *string    `json:"instance_url"`
		LastWebhookTimestamp *time.Time `json:"last_webhook_timestamp,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ConsumerKey == nil {
		return fmt.Errorf("required field consumer_key missing")
	}
	if all.InstanceUrl == nil {
		return fmt.Errorf("required field instance_url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"consumer_key", "instance_url", "last_webhook_timestamp"})
	} else {
		return err
	}
	o.ConsumerKey = *all.ConsumerKey
	o.InstanceUrl = *all.InstanceUrl
	o.LastWebhookTimestamp = all.LastWebhookTimestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
