// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampleLogGenerationSubscriptionAttributes The attributes describing a sample log generation subscription.
type SampleLogGenerationSubscriptionAttributes struct {
	// The identifier of the Cloud SIEM content pack the subscription targets.
	ContentPackId string `json:"content_pack_id"`
	// The time at which the subscription was created.
	CreatedAt time.Time `json:"created_at"`
	// The time at which the subscription expires and stops generating logs.
	ExpiresAt time.Time `json:"expires_at"`
	// Whether the subscription is currently active and generating logs.
	IsActive bool `json:"is_active"`
	// The status of the subscription.
	Status SampleLogGenerationSubscriptionStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSampleLogGenerationSubscriptionAttributes instantiates a new SampleLogGenerationSubscriptionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSampleLogGenerationSubscriptionAttributes(contentPackId string, createdAt time.Time, expiresAt time.Time, isActive bool, status SampleLogGenerationSubscriptionStatus) *SampleLogGenerationSubscriptionAttributes {
	this := SampleLogGenerationSubscriptionAttributes{}
	this.ContentPackId = contentPackId
	this.CreatedAt = createdAt
	this.ExpiresAt = expiresAt
	this.IsActive = isActive
	this.Status = status
	return &this
}

// NewSampleLogGenerationSubscriptionAttributesWithDefaults instantiates a new SampleLogGenerationSubscriptionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSampleLogGenerationSubscriptionAttributesWithDefaults() *SampleLogGenerationSubscriptionAttributes {
	this := SampleLogGenerationSubscriptionAttributes{}
	return &this
}

// GetContentPackId returns the ContentPackId field value.
func (o *SampleLogGenerationSubscriptionAttributes) GetContentPackId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ContentPackId
}

// GetContentPackIdOk returns a tuple with the ContentPackId field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationSubscriptionAttributes) GetContentPackIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentPackId, true
}

// SetContentPackId sets field value.
func (o *SampleLogGenerationSubscriptionAttributes) SetContentPackId(v string) {
	o.ContentPackId = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *SampleLogGenerationSubscriptionAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationSubscriptionAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *SampleLogGenerationSubscriptionAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetExpiresAt returns the ExpiresAt field value.
func (o *SampleLogGenerationSubscriptionAttributes) GetExpiresAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationSubscriptionAttributes) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresAt, true
}

// SetExpiresAt sets field value.
func (o *SampleLogGenerationSubscriptionAttributes) SetExpiresAt(v time.Time) {
	o.ExpiresAt = v
}

// GetIsActive returns the IsActive field value.
func (o *SampleLogGenerationSubscriptionAttributes) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationSubscriptionAttributes) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value.
func (o *SampleLogGenerationSubscriptionAttributes) SetIsActive(v bool) {
	o.IsActive = v
}

// GetStatus returns the Status field value.
func (o *SampleLogGenerationSubscriptionAttributes) GetStatus() SampleLogGenerationSubscriptionStatus {
	if o == nil {
		var ret SampleLogGenerationSubscriptionStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationSubscriptionAttributes) GetStatusOk() (*SampleLogGenerationSubscriptionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *SampleLogGenerationSubscriptionAttributes) SetStatus(v SampleLogGenerationSubscriptionStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SampleLogGenerationSubscriptionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["content_pack_id"] = o.ContentPackId
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.ExpiresAt.Nanosecond() == 0 {
		toSerialize["expires_at"] = o.ExpiresAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["expires_at"] = o.ExpiresAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["is_active"] = o.IsActive
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SampleLogGenerationSubscriptionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContentPackId *string                                `json:"content_pack_id"`
		CreatedAt     *time.Time                             `json:"created_at"`
		ExpiresAt     *time.Time                             `json:"expires_at"`
		IsActive      *bool                                  `json:"is_active"`
		Status        *SampleLogGenerationSubscriptionStatus `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ContentPackId == nil {
		return fmt.Errorf("required field content_pack_id missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.ExpiresAt == nil {
		return fmt.Errorf("required field expires_at missing")
	}
	if all.IsActive == nil {
		return fmt.Errorf("required field is_active missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"content_pack_id", "created_at", "expires_at", "is_active", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ContentPackId = *all.ContentPackId
	o.CreatedAt = *all.CreatedAt
	o.ExpiresAt = *all.ExpiresAt
	o.IsActive = *all.IsActive
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
