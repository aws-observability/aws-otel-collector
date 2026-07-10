// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipFeedbackResultAttributes The attributes of an ownership feedback result.
type OwnershipFeedbackResultAttributes struct {
	// The feedback action to apply to an inference.
	Action OwnershipFeedbackAction `json:"action"`
	// The checksum of the inference after the feedback was applied.
	Checksum string `json:"checksum"`
	// The lifecycle status of an ownership inference.
	NewStatus OwnershipInferenceStatus `json:"new_status"`
	// The owner type for an ownership inference.
	OwnerType OwnershipOwnerType `json:"owner_type"`
	// The lifecycle status of an ownership inference.
	PreviousStatus OwnershipInferenceStatus `json:"previous_status"`
	// The primary contact reference for the inferred owner after the feedback was applied, formatted as `ref:handle/<owner_handle>`.
	PrimaryContactRef datadog.NullableString `json:"primary_contact_ref,omitempty"`
	// The time when the inference was updated by the feedback.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOwnershipFeedbackResultAttributes instantiates a new OwnershipFeedbackResultAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOwnershipFeedbackResultAttributes(action OwnershipFeedbackAction, checksum string, newStatus OwnershipInferenceStatus, ownerType OwnershipOwnerType, previousStatus OwnershipInferenceStatus, updatedAt time.Time) *OwnershipFeedbackResultAttributes {
	this := OwnershipFeedbackResultAttributes{}
	this.Action = action
	this.Checksum = checksum
	this.NewStatus = newStatus
	this.OwnerType = ownerType
	this.PreviousStatus = previousStatus
	this.UpdatedAt = updatedAt
	return &this
}

// NewOwnershipFeedbackResultAttributesWithDefaults instantiates a new OwnershipFeedbackResultAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOwnershipFeedbackResultAttributesWithDefaults() *OwnershipFeedbackResultAttributes {
	this := OwnershipFeedbackResultAttributes{}
	return &this
}

// GetAction returns the Action field value.
func (o *OwnershipFeedbackResultAttributes) GetAction() OwnershipFeedbackAction {
	if o == nil {
		var ret OwnershipFeedbackAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *OwnershipFeedbackResultAttributes) GetActionOk() (*OwnershipFeedbackAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *OwnershipFeedbackResultAttributes) SetAction(v OwnershipFeedbackAction) {
	o.Action = v
}

// GetChecksum returns the Checksum field value.
func (o *OwnershipFeedbackResultAttributes) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *OwnershipFeedbackResultAttributes) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value.
func (o *OwnershipFeedbackResultAttributes) SetChecksum(v string) {
	o.Checksum = v
}

// GetNewStatus returns the NewStatus field value.
func (o *OwnershipFeedbackResultAttributes) GetNewStatus() OwnershipInferenceStatus {
	if o == nil {
		var ret OwnershipInferenceStatus
		return ret
	}
	return o.NewStatus
}

// GetNewStatusOk returns a tuple with the NewStatus field value
// and a boolean to check if the value has been set.
func (o *OwnershipFeedbackResultAttributes) GetNewStatusOk() (*OwnershipInferenceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewStatus, true
}

// SetNewStatus sets field value.
func (o *OwnershipFeedbackResultAttributes) SetNewStatus(v OwnershipInferenceStatus) {
	o.NewStatus = v
}

// GetOwnerType returns the OwnerType field value.
func (o *OwnershipFeedbackResultAttributes) GetOwnerType() OwnershipOwnerType {
	if o == nil {
		var ret OwnershipOwnerType
		return ret
	}
	return o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value
// and a boolean to check if the value has been set.
func (o *OwnershipFeedbackResultAttributes) GetOwnerTypeOk() (*OwnershipOwnerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerType, true
}

// SetOwnerType sets field value.
func (o *OwnershipFeedbackResultAttributes) SetOwnerType(v OwnershipOwnerType) {
	o.OwnerType = v
}

// GetPreviousStatus returns the PreviousStatus field value.
func (o *OwnershipFeedbackResultAttributes) GetPreviousStatus() OwnershipInferenceStatus {
	if o == nil {
		var ret OwnershipInferenceStatus
		return ret
	}
	return o.PreviousStatus
}

// GetPreviousStatusOk returns a tuple with the PreviousStatus field value
// and a boolean to check if the value has been set.
func (o *OwnershipFeedbackResultAttributes) GetPreviousStatusOk() (*OwnershipInferenceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreviousStatus, true
}

// SetPreviousStatus sets field value.
func (o *OwnershipFeedbackResultAttributes) SetPreviousStatus(v OwnershipInferenceStatus) {
	o.PreviousStatus = v
}

// GetPrimaryContactRef returns the PrimaryContactRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OwnershipFeedbackResultAttributes) GetPrimaryContactRef() string {
	if o == nil || o.PrimaryContactRef.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrimaryContactRef.Get()
}

// GetPrimaryContactRefOk returns a tuple with the PrimaryContactRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OwnershipFeedbackResultAttributes) GetPrimaryContactRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryContactRef.Get(), o.PrimaryContactRef.IsSet()
}

// HasPrimaryContactRef returns a boolean if a field has been set.
func (o *OwnershipFeedbackResultAttributes) HasPrimaryContactRef() bool {
	return o != nil && o.PrimaryContactRef.IsSet()
}

// SetPrimaryContactRef gets a reference to the given datadog.NullableString and assigns it to the PrimaryContactRef field.
func (o *OwnershipFeedbackResultAttributes) SetPrimaryContactRef(v string) {
	o.PrimaryContactRef.Set(&v)
}

// SetPrimaryContactRefNil sets the value for PrimaryContactRef to be an explicit nil.
func (o *OwnershipFeedbackResultAttributes) SetPrimaryContactRefNil() {
	o.PrimaryContactRef.Set(nil)
}

// UnsetPrimaryContactRef ensures that no value is present for PrimaryContactRef, not even an explicit nil.
func (o *OwnershipFeedbackResultAttributes) UnsetPrimaryContactRef() {
	o.PrimaryContactRef.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *OwnershipFeedbackResultAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OwnershipFeedbackResultAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *OwnershipFeedbackResultAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OwnershipFeedbackResultAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["checksum"] = o.Checksum
	toSerialize["new_status"] = o.NewStatus
	toSerialize["owner_type"] = o.OwnerType
	toSerialize["previous_status"] = o.PreviousStatus
	if o.PrimaryContactRef.IsSet() {
		toSerialize["primary_contact_ref"] = o.PrimaryContactRef.Get()
	}
	if o.UpdatedAt.Nanosecond() == 0 {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["updated_at"] = o.UpdatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OwnershipFeedbackResultAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action            *OwnershipFeedbackAction  `json:"action"`
		Checksum          *string                   `json:"checksum"`
		NewStatus         *OwnershipInferenceStatus `json:"new_status"`
		OwnerType         *OwnershipOwnerType       `json:"owner_type"`
		PreviousStatus    *OwnershipInferenceStatus `json:"previous_status"`
		PrimaryContactRef datadog.NullableString    `json:"primary_contact_ref,omitempty"`
		UpdatedAt         *time.Time                `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.Checksum == nil {
		return fmt.Errorf("required field checksum missing")
	}
	if all.NewStatus == nil {
		return fmt.Errorf("required field new_status missing")
	}
	if all.OwnerType == nil {
		return fmt.Errorf("required field owner_type missing")
	}
	if all.PreviousStatus == nil {
		return fmt.Errorf("required field previous_status missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "checksum", "new_status", "owner_type", "previous_status", "primary_contact_ref", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	o.Checksum = *all.Checksum
	if !all.NewStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.NewStatus = *all.NewStatus
	}
	if !all.OwnerType.IsValid() {
		hasInvalidField = true
	} else {
		o.OwnerType = *all.OwnerType
	}
	if !all.PreviousStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.PreviousStatus = *all.PreviousStatus
	}
	o.PrimaryContactRef = all.PrimaryContactRef
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
