// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampleLogGenerationBulkSubscriptionAttributes The attributes for creating sample log generation subscriptions for multiple content packs.
type SampleLogGenerationBulkSubscriptionAttributes struct {
	// The identifiers of the Cloud SIEM content packs to subscribe to. At most five content packs can be requested in a single call.
	ContentPackIds []string `json:"content_pack_ids"`
	// How long the subscription should remain active before expiring.
	Duration *SampleLogGenerationDuration `json:"duration,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSampleLogGenerationBulkSubscriptionAttributes instantiates a new SampleLogGenerationBulkSubscriptionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSampleLogGenerationBulkSubscriptionAttributes(contentPackIds []string) *SampleLogGenerationBulkSubscriptionAttributes {
	this := SampleLogGenerationBulkSubscriptionAttributes{}
	this.ContentPackIds = contentPackIds
	var duration SampleLogGenerationDuration = SAMPLELOGGENERATIONDURATION_THREE_DAYS
	this.Duration = &duration
	return &this
}

// NewSampleLogGenerationBulkSubscriptionAttributesWithDefaults instantiates a new SampleLogGenerationBulkSubscriptionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSampleLogGenerationBulkSubscriptionAttributesWithDefaults() *SampleLogGenerationBulkSubscriptionAttributes {
	this := SampleLogGenerationBulkSubscriptionAttributes{}
	var duration SampleLogGenerationDuration = SAMPLELOGGENERATIONDURATION_THREE_DAYS
	this.Duration = &duration
	return &this
}

// GetContentPackIds returns the ContentPackIds field value.
func (o *SampleLogGenerationBulkSubscriptionAttributes) GetContentPackIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ContentPackIds
}

// GetContentPackIdsOk returns a tuple with the ContentPackIds field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationBulkSubscriptionAttributes) GetContentPackIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentPackIds, true
}

// SetContentPackIds sets field value.
func (o *SampleLogGenerationBulkSubscriptionAttributes) SetContentPackIds(v []string) {
	o.ContentPackIds = v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SampleLogGenerationBulkSubscriptionAttributes) GetDuration() SampleLogGenerationDuration {
	if o == nil || o.Duration == nil {
		var ret SampleLogGenerationDuration
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationBulkSubscriptionAttributes) GetDurationOk() (*SampleLogGenerationDuration, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SampleLogGenerationBulkSubscriptionAttributes) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given SampleLogGenerationDuration and assigns it to the Duration field.
func (o *SampleLogGenerationBulkSubscriptionAttributes) SetDuration(v SampleLogGenerationDuration) {
	o.Duration = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SampleLogGenerationBulkSubscriptionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["content_pack_ids"] = o.ContentPackIds
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SampleLogGenerationBulkSubscriptionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContentPackIds *[]string                    `json:"content_pack_ids"`
		Duration       *SampleLogGenerationDuration `json:"duration,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ContentPackIds == nil {
		return fmt.Errorf("required field content_pack_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content_pack_ids", "duration"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ContentPackIds = *all.ContentPackIds
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
