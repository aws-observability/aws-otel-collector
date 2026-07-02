// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TagIndexingRuleExemptionAttributes Attributes of a tag indexing rule exemption.
type TagIndexingRuleExemptionAttributes struct {
	// Timestamp when the exemption was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Handle of the user who created the exemption.
	CreatedByHandle *string `json:"created_by_handle,omitempty"`
	// Discriminates between an explicit exemption (`exemption`) and a pre-existing legacy tag configuration acting as an implicit exclusion (`legacy_tag_configuration`).
	Kind *string `json:"kind,omitempty"`
	// The reason the metric is exempt from tag indexing rules.
	Reason *string `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTagIndexingRuleExemptionAttributes instantiates a new TagIndexingRuleExemptionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTagIndexingRuleExemptionAttributes() *TagIndexingRuleExemptionAttributes {
	this := TagIndexingRuleExemptionAttributes{}
	return &this
}

// NewTagIndexingRuleExemptionAttributesWithDefaults instantiates a new TagIndexingRuleExemptionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTagIndexingRuleExemptionAttributesWithDefaults() *TagIndexingRuleExemptionAttributes {
	this := TagIndexingRuleExemptionAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TagIndexingRuleExemptionAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleExemptionAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TagIndexingRuleExemptionAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TagIndexingRuleExemptionAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCreatedByHandle returns the CreatedByHandle field value if set, zero value otherwise.
func (o *TagIndexingRuleExemptionAttributes) GetCreatedByHandle() string {
	if o == nil || o.CreatedByHandle == nil {
		var ret string
		return ret
	}
	return *o.CreatedByHandle
}

// GetCreatedByHandleOk returns a tuple with the CreatedByHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleExemptionAttributes) GetCreatedByHandleOk() (*string, bool) {
	if o == nil || o.CreatedByHandle == nil {
		return nil, false
	}
	return o.CreatedByHandle, true
}

// HasCreatedByHandle returns a boolean if a field has been set.
func (o *TagIndexingRuleExemptionAttributes) HasCreatedByHandle() bool {
	return o != nil && o.CreatedByHandle != nil
}

// SetCreatedByHandle gets a reference to the given string and assigns it to the CreatedByHandle field.
func (o *TagIndexingRuleExemptionAttributes) SetCreatedByHandle(v string) {
	o.CreatedByHandle = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *TagIndexingRuleExemptionAttributes) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleExemptionAttributes) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *TagIndexingRuleExemptionAttributes) HasKind() bool {
	return o != nil && o.Kind != nil
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *TagIndexingRuleExemptionAttributes) SetKind(v string) {
	o.Kind = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *TagIndexingRuleExemptionAttributes) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagIndexingRuleExemptionAttributes) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *TagIndexingRuleExemptionAttributes) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *TagIndexingRuleExemptionAttributes) SetReason(v string) {
	o.Reason = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TagIndexingRuleExemptionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CreatedByHandle != nil {
		toSerialize["created_by_handle"] = o.CreatedByHandle
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TagIndexingRuleExemptionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt       *time.Time `json:"created_at,omitempty"`
		CreatedByHandle *string    `json:"created_by_handle,omitempty"`
		Kind            *string    `json:"kind,omitempty"`
		Reason          *string    `json:"reason,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by_handle", "kind", "reason"})
	} else {
		return err
	}
	o.CreatedAt = all.CreatedAt
	o.CreatedByHandle = all.CreatedByHandle
	o.Kind = all.Kind
	o.Reason = all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
