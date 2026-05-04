// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ChangeRequestDecisionResponseAttributes Attributes of a change request decision in a response.
type ChangeRequestDecisionResponseAttributes struct {
	// The status of a change request decision.
	ChangeRequestStatus ChangeRequestDecisionStatusType `json:"change_request_status"`
	// Timestamp of when the decision was made.
	DecidedAt time.Time `json:"decided_at"`
	// The reason for the decision.
	DecisionReason string `json:"decision_reason"`
	// Timestamp of when the decision was deleted.
	DeletedAt time.Time `json:"deleted_at"`
	// The reason for requesting the decision.
	RequestReason string `json:"request_reason"`
	// Timestamp of when the decision was requested.
	RequestedAt time.Time `json:"requested_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewChangeRequestDecisionResponseAttributes instantiates a new ChangeRequestDecisionResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewChangeRequestDecisionResponseAttributes(changeRequestStatus ChangeRequestDecisionStatusType, decidedAt time.Time, decisionReason string, deletedAt time.Time, requestReason string, requestedAt time.Time) *ChangeRequestDecisionResponseAttributes {
	this := ChangeRequestDecisionResponseAttributes{}
	this.ChangeRequestStatus = changeRequestStatus
	this.DecidedAt = decidedAt
	this.DecisionReason = decisionReason
	this.DeletedAt = deletedAt
	this.RequestReason = requestReason
	this.RequestedAt = requestedAt
	return &this
}

// NewChangeRequestDecisionResponseAttributesWithDefaults instantiates a new ChangeRequestDecisionResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewChangeRequestDecisionResponseAttributesWithDefaults() *ChangeRequestDecisionResponseAttributes {
	this := ChangeRequestDecisionResponseAttributes{}
	return &this
}

// GetChangeRequestStatus returns the ChangeRequestStatus field value.
func (o *ChangeRequestDecisionResponseAttributes) GetChangeRequestStatus() ChangeRequestDecisionStatusType {
	if o == nil {
		var ret ChangeRequestDecisionStatusType
		return ret
	}
	return o.ChangeRequestStatus
}

// GetChangeRequestStatusOk returns a tuple with the ChangeRequestStatus field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestDecisionResponseAttributes) GetChangeRequestStatusOk() (*ChangeRequestDecisionStatusType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeRequestStatus, true
}

// SetChangeRequestStatus sets field value.
func (o *ChangeRequestDecisionResponseAttributes) SetChangeRequestStatus(v ChangeRequestDecisionStatusType) {
	o.ChangeRequestStatus = v
}

// GetDecidedAt returns the DecidedAt field value.
func (o *ChangeRequestDecisionResponseAttributes) GetDecidedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.DecidedAt
}

// GetDecidedAtOk returns a tuple with the DecidedAt field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestDecisionResponseAttributes) GetDecidedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DecidedAt, true
}

// SetDecidedAt sets field value.
func (o *ChangeRequestDecisionResponseAttributes) SetDecidedAt(v time.Time) {
	o.DecidedAt = v
}

// GetDecisionReason returns the DecisionReason field value.
func (o *ChangeRequestDecisionResponseAttributes) GetDecisionReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DecisionReason
}

// GetDecisionReasonOk returns a tuple with the DecisionReason field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestDecisionResponseAttributes) GetDecisionReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DecisionReason, true
}

// SetDecisionReason sets field value.
func (o *ChangeRequestDecisionResponseAttributes) SetDecisionReason(v string) {
	o.DecisionReason = v
}

// GetDeletedAt returns the DeletedAt field value.
func (o *ChangeRequestDecisionResponseAttributes) GetDeletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestDecisionResponseAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value.
func (o *ChangeRequestDecisionResponseAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt = v
}

// GetRequestReason returns the RequestReason field value.
func (o *ChangeRequestDecisionResponseAttributes) GetRequestReason() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RequestReason
}

// GetRequestReasonOk returns a tuple with the RequestReason field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestDecisionResponseAttributes) GetRequestReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestReason, true
}

// SetRequestReason sets field value.
func (o *ChangeRequestDecisionResponseAttributes) SetRequestReason(v string) {
	o.RequestReason = v
}

// GetRequestedAt returns the RequestedAt field value.
func (o *ChangeRequestDecisionResponseAttributes) GetRequestedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value
// and a boolean to check if the value has been set.
func (o *ChangeRequestDecisionResponseAttributes) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestedAt, true
}

// SetRequestedAt sets field value.
func (o *ChangeRequestDecisionResponseAttributes) SetRequestedAt(v time.Time) {
	o.RequestedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ChangeRequestDecisionResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["change_request_status"] = o.ChangeRequestStatus
	if o.DecidedAt.Nanosecond() == 0 {
		toSerialize["decided_at"] = o.DecidedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["decided_at"] = o.DecidedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["decision_reason"] = o.DecisionReason
	if o.DeletedAt.Nanosecond() == 0 {
		toSerialize["deleted_at"] = o.DeletedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["deleted_at"] = o.DeletedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["request_reason"] = o.RequestReason
	if o.RequestedAt.Nanosecond() == 0 {
		toSerialize["requested_at"] = o.RequestedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["requested_at"] = o.RequestedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ChangeRequestDecisionResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChangeRequestStatus *ChangeRequestDecisionStatusType `json:"change_request_status"`
		DecidedAt           *time.Time                       `json:"decided_at"`
		DecisionReason      *string                          `json:"decision_reason"`
		DeletedAt           *time.Time                       `json:"deleted_at"`
		RequestReason       *string                          `json:"request_reason"`
		RequestedAt         *time.Time                       `json:"requested_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ChangeRequestStatus == nil {
		return fmt.Errorf("required field change_request_status missing")
	}
	if all.DecidedAt == nil {
		return fmt.Errorf("required field decided_at missing")
	}
	if all.DecisionReason == nil {
		return fmt.Errorf("required field decision_reason missing")
	}
	if all.DeletedAt == nil {
		return fmt.Errorf("required field deleted_at missing")
	}
	if all.RequestReason == nil {
		return fmt.Errorf("required field request_reason missing")
	}
	if all.RequestedAt == nil {
		return fmt.Errorf("required field requested_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"change_request_status", "decided_at", "decision_reason", "deleted_at", "request_reason", "requested_at"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.ChangeRequestStatus.IsValid() {
		hasInvalidField = true
	} else {
		o.ChangeRequestStatus = *all.ChangeRequestStatus
	}
	o.DecidedAt = *all.DecidedAt
	o.DecisionReason = *all.DecisionReason
	o.DeletedAt = *all.DeletedAt
	o.RequestReason = *all.RequestReason
	o.RequestedAt = *all.RequestedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
