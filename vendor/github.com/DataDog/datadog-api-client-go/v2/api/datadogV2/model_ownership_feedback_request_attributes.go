// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipFeedbackRequestAttributes The attributes of an ownership feedback request.
type OwnershipFeedbackRequestAttributes struct {
	// The feedback action to apply to an inference.
	Action OwnershipFeedbackAction `json:"action"`
	// The handle of the actor submitting the feedback.
	ActorHandle string `json:"actor_handle"`
	// The type of actor submitting the feedback, for example `user` or `service`.
	ActorType string `json:"actor_type"`
	// The corrected owner handle. Required when `action` is `correct`.
	CorrectedOwnerHandle datadog.NullableString `json:"corrected_owner_handle,omitempty"`
	// The corrected owner type. Required when `action` is `correct`.
	CorrectedOwnerType datadog.NullableString `json:"corrected_owner_type,omitempty"`
	// The checksum of the inference being acted upon. Must match the current inference checksum or the request returns a conflict.
	InferenceChecksum string `json:"inference_checksum"`
	// An optional free-form reason explaining the feedback.
	Reason datadog.NullableString `json:"reason,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOwnershipFeedbackRequestAttributes instantiates a new OwnershipFeedbackRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOwnershipFeedbackRequestAttributes(action OwnershipFeedbackAction, actorHandle string, actorType string, inferenceChecksum string) *OwnershipFeedbackRequestAttributes {
	this := OwnershipFeedbackRequestAttributes{}
	this.Action = action
	this.ActorHandle = actorHandle
	this.ActorType = actorType
	this.InferenceChecksum = inferenceChecksum
	return &this
}

// NewOwnershipFeedbackRequestAttributesWithDefaults instantiates a new OwnershipFeedbackRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOwnershipFeedbackRequestAttributesWithDefaults() *OwnershipFeedbackRequestAttributes {
	this := OwnershipFeedbackRequestAttributes{}
	return &this
}

// GetAction returns the Action field value.
func (o *OwnershipFeedbackRequestAttributes) GetAction() OwnershipFeedbackAction {
	if o == nil {
		var ret OwnershipFeedbackAction
		return ret
	}
	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *OwnershipFeedbackRequestAttributes) GetActionOk() (*OwnershipFeedbackAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value.
func (o *OwnershipFeedbackRequestAttributes) SetAction(v OwnershipFeedbackAction) {
	o.Action = v
}

// GetActorHandle returns the ActorHandle field value.
func (o *OwnershipFeedbackRequestAttributes) GetActorHandle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActorHandle
}

// GetActorHandleOk returns a tuple with the ActorHandle field value
// and a boolean to check if the value has been set.
func (o *OwnershipFeedbackRequestAttributes) GetActorHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActorHandle, true
}

// SetActorHandle sets field value.
func (o *OwnershipFeedbackRequestAttributes) SetActorHandle(v string) {
	o.ActorHandle = v
}

// GetActorType returns the ActorType field value.
func (o *OwnershipFeedbackRequestAttributes) GetActorType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ActorType
}

// GetActorTypeOk returns a tuple with the ActorType field value
// and a boolean to check if the value has been set.
func (o *OwnershipFeedbackRequestAttributes) GetActorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActorType, true
}

// SetActorType sets field value.
func (o *OwnershipFeedbackRequestAttributes) SetActorType(v string) {
	o.ActorType = v
}

// GetCorrectedOwnerHandle returns the CorrectedOwnerHandle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OwnershipFeedbackRequestAttributes) GetCorrectedOwnerHandle() string {
	if o == nil || o.CorrectedOwnerHandle.Get() == nil {
		var ret string
		return ret
	}
	return *o.CorrectedOwnerHandle.Get()
}

// GetCorrectedOwnerHandleOk returns a tuple with the CorrectedOwnerHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OwnershipFeedbackRequestAttributes) GetCorrectedOwnerHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CorrectedOwnerHandle.Get(), o.CorrectedOwnerHandle.IsSet()
}

// HasCorrectedOwnerHandle returns a boolean if a field has been set.
func (o *OwnershipFeedbackRequestAttributes) HasCorrectedOwnerHandle() bool {
	return o != nil && o.CorrectedOwnerHandle.IsSet()
}

// SetCorrectedOwnerHandle gets a reference to the given datadog.NullableString and assigns it to the CorrectedOwnerHandle field.
func (o *OwnershipFeedbackRequestAttributes) SetCorrectedOwnerHandle(v string) {
	o.CorrectedOwnerHandle.Set(&v)
}

// SetCorrectedOwnerHandleNil sets the value for CorrectedOwnerHandle to be an explicit nil.
func (o *OwnershipFeedbackRequestAttributes) SetCorrectedOwnerHandleNil() {
	o.CorrectedOwnerHandle.Set(nil)
}

// UnsetCorrectedOwnerHandle ensures that no value is present for CorrectedOwnerHandle, not even an explicit nil.
func (o *OwnershipFeedbackRequestAttributes) UnsetCorrectedOwnerHandle() {
	o.CorrectedOwnerHandle.Unset()
}

// GetCorrectedOwnerType returns the CorrectedOwnerType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OwnershipFeedbackRequestAttributes) GetCorrectedOwnerType() string {
	if o == nil || o.CorrectedOwnerType.Get() == nil {
		var ret string
		return ret
	}
	return *o.CorrectedOwnerType.Get()
}

// GetCorrectedOwnerTypeOk returns a tuple with the CorrectedOwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OwnershipFeedbackRequestAttributes) GetCorrectedOwnerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CorrectedOwnerType.Get(), o.CorrectedOwnerType.IsSet()
}

// HasCorrectedOwnerType returns a boolean if a field has been set.
func (o *OwnershipFeedbackRequestAttributes) HasCorrectedOwnerType() bool {
	return o != nil && o.CorrectedOwnerType.IsSet()
}

// SetCorrectedOwnerType gets a reference to the given datadog.NullableString and assigns it to the CorrectedOwnerType field.
func (o *OwnershipFeedbackRequestAttributes) SetCorrectedOwnerType(v string) {
	o.CorrectedOwnerType.Set(&v)
}

// SetCorrectedOwnerTypeNil sets the value for CorrectedOwnerType to be an explicit nil.
func (o *OwnershipFeedbackRequestAttributes) SetCorrectedOwnerTypeNil() {
	o.CorrectedOwnerType.Set(nil)
}

// UnsetCorrectedOwnerType ensures that no value is present for CorrectedOwnerType, not even an explicit nil.
func (o *OwnershipFeedbackRequestAttributes) UnsetCorrectedOwnerType() {
	o.CorrectedOwnerType.Unset()
}

// GetInferenceChecksum returns the InferenceChecksum field value.
func (o *OwnershipFeedbackRequestAttributes) GetInferenceChecksum() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InferenceChecksum
}

// GetInferenceChecksumOk returns a tuple with the InferenceChecksum field value
// and a boolean to check if the value has been set.
func (o *OwnershipFeedbackRequestAttributes) GetInferenceChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InferenceChecksum, true
}

// SetInferenceChecksum sets field value.
func (o *OwnershipFeedbackRequestAttributes) SetInferenceChecksum(v string) {
	o.InferenceChecksum = v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OwnershipFeedbackRequestAttributes) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OwnershipFeedbackRequestAttributes) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *OwnershipFeedbackRequestAttributes) HasReason() bool {
	return o != nil && o.Reason.IsSet()
}

// SetReason gets a reference to the given datadog.NullableString and assigns it to the Reason field.
func (o *OwnershipFeedbackRequestAttributes) SetReason(v string) {
	o.Reason.Set(&v)
}

// SetReasonNil sets the value for Reason to be an explicit nil.
func (o *OwnershipFeedbackRequestAttributes) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil.
func (o *OwnershipFeedbackRequestAttributes) UnsetReason() {
	o.Reason.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o OwnershipFeedbackRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["action"] = o.Action
	toSerialize["actor_handle"] = o.ActorHandle
	toSerialize["actor_type"] = o.ActorType
	if o.CorrectedOwnerHandle.IsSet() {
		toSerialize["corrected_owner_handle"] = o.CorrectedOwnerHandle.Get()
	}
	if o.CorrectedOwnerType.IsSet() {
		toSerialize["corrected_owner_type"] = o.CorrectedOwnerType.Get()
	}
	toSerialize["inference_checksum"] = o.InferenceChecksum
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OwnershipFeedbackRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action               *OwnershipFeedbackAction `json:"action"`
		ActorHandle          *string                  `json:"actor_handle"`
		ActorType            *string                  `json:"actor_type"`
		CorrectedOwnerHandle datadog.NullableString   `json:"corrected_owner_handle,omitempty"`
		CorrectedOwnerType   datadog.NullableString   `json:"corrected_owner_type,omitempty"`
		InferenceChecksum    *string                  `json:"inference_checksum"`
		Reason               datadog.NullableString   `json:"reason,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Action == nil {
		return fmt.Errorf("required field action missing")
	}
	if all.ActorHandle == nil {
		return fmt.Errorf("required field actor_handle missing")
	}
	if all.ActorType == nil {
		return fmt.Errorf("required field actor_type missing")
	}
	if all.InferenceChecksum == nil {
		return fmt.Errorf("required field inference_checksum missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "actor_handle", "actor_type", "corrected_owner_handle", "corrected_owner_type", "inference_checksum", "reason"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Action.IsValid() {
		hasInvalidField = true
	} else {
		o.Action = *all.Action
	}
	o.ActorHandle = *all.ActorHandle
	o.ActorType = *all.ActorType
	o.CorrectedOwnerHandle = all.CorrectedOwnerHandle
	o.CorrectedOwnerType = all.CorrectedOwnerType
	o.InferenceChecksum = *all.InferenceChecksum
	o.Reason = all.Reason

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
