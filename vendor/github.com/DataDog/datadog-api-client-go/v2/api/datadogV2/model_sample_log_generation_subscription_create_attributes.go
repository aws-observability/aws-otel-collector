// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampleLogGenerationSubscriptionCreateAttributes The attributes for creating a sample log generation subscription.
type SampleLogGenerationSubscriptionCreateAttributes struct {
	// The identifier of the Cloud SIEM content pack to subscribe to.
	ContentPackId string `json:"content_pack_id"`
	// How long the subscription should remain active before expiring.
	Duration *SampleLogGenerationDuration `json:"duration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSampleLogGenerationSubscriptionCreateAttributes instantiates a new SampleLogGenerationSubscriptionCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSampleLogGenerationSubscriptionCreateAttributes(contentPackId string) *SampleLogGenerationSubscriptionCreateAttributes {
	this := SampleLogGenerationSubscriptionCreateAttributes{}
	this.ContentPackId = contentPackId
	var duration SampleLogGenerationDuration = SAMPLELOGGENERATIONDURATION_THREE_DAYS
	this.Duration = &duration
	return &this
}

// NewSampleLogGenerationSubscriptionCreateAttributesWithDefaults instantiates a new SampleLogGenerationSubscriptionCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSampleLogGenerationSubscriptionCreateAttributesWithDefaults() *SampleLogGenerationSubscriptionCreateAttributes {
	this := SampleLogGenerationSubscriptionCreateAttributes{}
	var duration SampleLogGenerationDuration = SAMPLELOGGENERATIONDURATION_THREE_DAYS
	this.Duration = &duration
	return &this
}

// GetContentPackId returns the ContentPackId field value.
func (o *SampleLogGenerationSubscriptionCreateAttributes) GetContentPackId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ContentPackId
}

// GetContentPackIdOk returns a tuple with the ContentPackId field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationSubscriptionCreateAttributes) GetContentPackIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentPackId, true
}

// SetContentPackId sets field value.
func (o *SampleLogGenerationSubscriptionCreateAttributes) SetContentPackId(v string) {
	o.ContentPackId = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SampleLogGenerationSubscriptionCreateAttributes) GetDuration() SampleLogGenerationDuration {
	if o == nil || o.Duration == nil {
		var ret SampleLogGenerationDuration
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationSubscriptionCreateAttributes) GetDurationOk() (*SampleLogGenerationDuration, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SampleLogGenerationSubscriptionCreateAttributes) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given SampleLogGenerationDuration and assigns it to the Duration field.
func (o *SampleLogGenerationSubscriptionCreateAttributes) SetDuration(v SampleLogGenerationDuration) {
	o.Duration = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SampleLogGenerationSubscriptionCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["content_pack_id"] = o.ContentPackId
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SampleLogGenerationSubscriptionCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContentPackId *string                      `json:"content_pack_id"`
		Duration      *SampleLogGenerationDuration `json:"duration,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ContentPackId == nil {
		return fmt.Errorf("required field content_pack_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content_pack_id", "duration"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ContentPackId = *all.ContentPackId
	if all.Duration != nil && !all.Duration.IsValid() {
		hasInvalidField = true
	} else {
		o.Duration = all.Duration
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
