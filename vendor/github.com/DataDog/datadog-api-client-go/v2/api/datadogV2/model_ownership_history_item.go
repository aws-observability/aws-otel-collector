// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipHistoryItem A single ownership inference history entry.
type OwnershipHistoryItem struct {
	// A checksum identifying the state of the inference at this point in time.
	Checksum string `json:"checksum"`
	// The confidence score of the inference, expressed as a numeric string with up to four decimal places.
	Confidence string `json:"confidence"`
	// The time this history entry was created.
	CreatedAt time.Time `json:"created_at"`
	// The list of evidence versions associated with an inference.
	EvidenceVersions datadog.NullableList[map[string]interface{}] `json:"evidence_versions"`
	// A human-readable explanation of how the inference was produced.
	Explanation string `json:"explanation"`
	// The time when this inference failed, if applicable.
	FailedAt datadog.NullableTime `json:"failed_at,omitempty"`
	// The reason why this inference failed, if applicable.
	FailureReason datadog.NullableString `json:"failure_reason,omitempty"`
	// The unique identifier of the history entry.
	Id int64 `json:"id"`
	// The owner type for an ownership inference.
	OwnerType OwnershipOwnerType `json:"owner_type"`
	// The primary contact reference for the inferred owner, formatted as `ref:handle/<owner_handle>`.
	PrimaryContactRef datadog.NullableString `json:"primary_contact_ref,omitempty"`
	// The identifier of the resource that the inference applies to.
	ResourceId string `json:"resource_id"`
	// The scheduled retry time for a failed inference, if applicable.
	RetrySchedule datadog.NullableTime `json:"retry_schedule,omitempty"`
	// The list of sources backing an ownership inference. Empty when the inference status is not whitelisted to expose sources.
	Sources []map[string]interface{} `json:"sources"`
	// The lifecycle status of an ownership inference.
	Status OwnershipInferenceStatus `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOwnershipHistoryItem instantiates a new OwnershipHistoryItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOwnershipHistoryItem(checksum string, confidence string, createdAt time.Time, evidenceVersions datadog.NullableList[map[string]interface{}], explanation string, id int64, ownerType OwnershipOwnerType, resourceId string, sources []map[string]interface{}, status OwnershipInferenceStatus) *OwnershipHistoryItem {
	this := OwnershipHistoryItem{}
	this.Checksum = checksum
	this.Confidence = confidence
	this.CreatedAt = createdAt
	this.EvidenceVersions = evidenceVersions
	this.Explanation = explanation
	this.Id = id
	this.OwnerType = ownerType
	this.ResourceId = resourceId
	this.Sources = sources
	this.Status = status
	return &this
}

// NewOwnershipHistoryItemWithDefaults instantiates a new OwnershipHistoryItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOwnershipHistoryItemWithDefaults() *OwnershipHistoryItem {
	this := OwnershipHistoryItem{}
	return &this
}

// GetChecksum returns the Checksum field value.
func (o *OwnershipHistoryItem) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *OwnershipHistoryItem) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value.
func (o *OwnershipHistoryItem) SetChecksum(v string) {
	o.Checksum = v
}

// GetConfidence returns the Confidence field value.
func (o *OwnershipHistoryItem) GetConfidence() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *OwnershipHistoryItem) GetConfidenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value.
func (o *OwnershipHistoryItem) SetConfidence(v string) {
	o.Confidence = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *OwnershipHistoryItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OwnershipHistoryItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *OwnershipHistoryItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEvidenceVersions returns the EvidenceVersions field value.
// If the value is explicit nil, the zero value for []map[string]interface{} will be returned.
func (o *OwnershipHistoryItem) GetEvidenceVersions() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.EvidenceVersions.Get()
}

// GetEvidenceVersionsOk returns a tuple with the EvidenceVersions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OwnershipHistoryItem) GetEvidenceVersionsOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.EvidenceVersions.Get(), o.EvidenceVersions.IsSet()
}

// SetEvidenceVersions sets field value.
func (o *OwnershipHistoryItem) SetEvidenceVersions(v []map[string]interface{}) {
	o.EvidenceVersions.Set(&v)
}

// GetExplanation returns the Explanation field value.
func (o *OwnershipHistoryItem) GetExplanation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Explanation
}

// GetExplanationOk returns a tuple with the Explanation field value
// and a boolean to check if the value has been set.
func (o *OwnershipHistoryItem) GetExplanationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Explanation, true
}

// SetExplanation sets field value.
func (o *OwnershipHistoryItem) SetExplanation(v string) {
	o.Explanation = v
}

// GetFailedAt returns the FailedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OwnershipHistoryItem) GetFailedAt() time.Time {
	if o == nil || o.FailedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.FailedAt.Get()
}

// GetFailedAtOk returns a tuple with the FailedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OwnershipHistoryItem) GetFailedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailedAt.Get(), o.FailedAt.IsSet()
}

// HasFailedAt returns a boolean if a field has been set.
func (o *OwnershipHistoryItem) HasFailedAt() bool {
	return o != nil && o.FailedAt.IsSet()
}

// SetFailedAt gets a reference to the given datadog.NullableTime and assigns it to the FailedAt field.
func (o *OwnershipHistoryItem) SetFailedAt(v time.Time) {
	o.FailedAt.Set(&v)
}

// SetFailedAtNil sets the value for FailedAt to be an explicit nil.
func (o *OwnershipHistoryItem) SetFailedAtNil() {
	o.FailedAt.Set(nil)
}

// UnsetFailedAt ensures that no value is present for FailedAt, not even an explicit nil.
func (o *OwnershipHistoryItem) UnsetFailedAt() {
	o.FailedAt.Unset()
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OwnershipHistoryItem) GetFailureReason() string {
	if o == nil || o.FailureReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OwnershipHistoryItem) GetFailureReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// HasFailureReason returns a boolean if a field has been set.
func (o *OwnershipHistoryItem) HasFailureReason() bool {
	return o != nil && o.FailureReason.IsSet()
}

// SetFailureReason gets a reference to the given datadog.NullableString and assigns it to the FailureReason field.
func (o *OwnershipHistoryItem) SetFailureReason(v string) {
	o.FailureReason.Set(&v)
}

// SetFailureReasonNil sets the value for FailureReason to be an explicit nil.
func (o *OwnershipHistoryItem) SetFailureReasonNil() {
	o.FailureReason.Set(nil)
}

// UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil.
func (o *OwnershipHistoryItem) UnsetFailureReason() {
	o.FailureReason.Unset()
}

// GetId returns the Id field value.
func (o *OwnershipHistoryItem) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OwnershipHistoryItem) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *OwnershipHistoryItem) SetId(v int64) {
	o.Id = v
}

// GetOwnerType returns the OwnerType field value.
func (o *OwnershipHistoryItem) GetOwnerType() OwnershipOwnerType {
	if o == nil {
		var ret OwnershipOwnerType
		return ret
	}
	return o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value
// and a boolean to check if the value has been set.
func (o *OwnershipHistoryItem) GetOwnerTypeOk() (*OwnershipOwnerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerType, true
}

// SetOwnerType sets field value.
func (o *OwnershipHistoryItem) SetOwnerType(v OwnershipOwnerType) {
	o.OwnerType = v
}

// GetPrimaryContactRef returns the PrimaryContactRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OwnershipHistoryItem) GetPrimaryContactRef() string {
	if o == nil || o.PrimaryContactRef.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrimaryContactRef.Get()
}

// GetPrimaryContactRefOk returns a tuple with the PrimaryContactRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OwnershipHistoryItem) GetPrimaryContactRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryContactRef.Get(), o.PrimaryContactRef.IsSet()
}

// HasPrimaryContactRef returns a boolean if a field has been set.
func (o *OwnershipHistoryItem) HasPrimaryContactRef() bool {
	return o != nil && o.PrimaryContactRef.IsSet()
}

// SetPrimaryContactRef gets a reference to the given datadog.NullableString and assigns it to the PrimaryContactRef field.
func (o *OwnershipHistoryItem) SetPrimaryContactRef(v string) {
	o.PrimaryContactRef.Set(&v)
}

// SetPrimaryContactRefNil sets the value for PrimaryContactRef to be an explicit nil.
func (o *OwnershipHistoryItem) SetPrimaryContactRefNil() {
	o.PrimaryContactRef.Set(nil)
}

// UnsetPrimaryContactRef ensures that no value is present for PrimaryContactRef, not even an explicit nil.
func (o *OwnershipHistoryItem) UnsetPrimaryContactRef() {
	o.PrimaryContactRef.Unset()
}

// GetResourceId returns the ResourceId field value.
func (o *OwnershipHistoryItem) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *OwnershipHistoryItem) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value.
func (o *OwnershipHistoryItem) SetResourceId(v string) {
	o.ResourceId = v
}

// GetRetrySchedule returns the RetrySchedule field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OwnershipHistoryItem) GetRetrySchedule() time.Time {
	if o == nil || o.RetrySchedule.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RetrySchedule.Get()
}

// GetRetryScheduleOk returns a tuple with the RetrySchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OwnershipHistoryItem) GetRetryScheduleOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetrySchedule.Get(), o.RetrySchedule.IsSet()
}

// HasRetrySchedule returns a boolean if a field has been set.
func (o *OwnershipHistoryItem) HasRetrySchedule() bool {
	return o != nil && o.RetrySchedule.IsSet()
}

// SetRetrySchedule gets a reference to the given datadog.NullableTime and assigns it to the RetrySchedule field.
func (o *OwnershipHistoryItem) SetRetrySchedule(v time.Time) {
	o.RetrySchedule.Set(&v)
}

// SetRetryScheduleNil sets the value for RetrySchedule to be an explicit nil.
func (o *OwnershipHistoryItem) SetRetryScheduleNil() {
	o.RetrySchedule.Set(nil)
}

// UnsetRetrySchedule ensures that no value is present for RetrySchedule, not even an explicit nil.
func (o *OwnershipHistoryItem) UnsetRetrySchedule() {
	o.RetrySchedule.Unset()
}

// GetSources returns the Sources field value.
func (o *OwnershipHistoryItem) GetSources() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *OwnershipHistoryItem) GetSourcesOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *OwnershipHistoryItem) SetSources(v []map[string]interface{}) {
	o.Sources = v
}

// GetStatus returns the Status field value.
func (o *OwnershipHistoryItem) GetStatus() OwnershipInferenceStatus {
	if o == nil {
		var ret OwnershipInferenceStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OwnershipHistoryItem) GetStatusOk() (*OwnershipInferenceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *OwnershipHistoryItem) SetStatus(v OwnershipInferenceStatus) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OwnershipHistoryItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["checksum"] = o.Checksum
	toSerialize["confidence"] = o.Confidence
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["evidence_versions"] = o.EvidenceVersions.Get()
	toSerialize["explanation"] = o.Explanation
	if o.FailedAt.IsSet() {
		toSerialize["failed_at"] = o.FailedAt.Get()
	}
	if o.FailureReason.IsSet() {
		toSerialize["failure_reason"] = o.FailureReason.Get()
	}
	toSerialize["id"] = o.Id
	toSerialize["owner_type"] = o.OwnerType
	if o.PrimaryContactRef.IsSet() {
		toSerialize["primary_contact_ref"] = o.PrimaryContactRef.Get()
	}
	toSerialize["resource_id"] = o.ResourceId
	if o.RetrySchedule.IsSet() {
		toSerialize["retry_schedule"] = o.RetrySchedule.Get()
	}
	toSerialize["sources"] = o.Sources
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OwnershipHistoryItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Checksum          *string                                      `json:"checksum"`
		Confidence        *string                                      `json:"confidence"`
		CreatedAt         *time.Time                                   `json:"created_at"`
		EvidenceVersions  datadog.NullableList[map[string]interface{}] `json:"evidence_versions"`
		Explanation       *string                                      `json:"explanation"`
		FailedAt          datadog.NullableTime                         `json:"failed_at,omitempty"`
		FailureReason     datadog.NullableString                       `json:"failure_reason,omitempty"`
		Id                *int64                                       `json:"id"`
		OwnerType         *OwnershipOwnerType                          `json:"owner_type"`
		PrimaryContactRef datadog.NullableString                       `json:"primary_contact_ref,omitempty"`
		ResourceId        *string                                      `json:"resource_id"`
		RetrySchedule     datadog.NullableTime                         `json:"retry_schedule,omitempty"`
		Sources           *[]map[string]interface{}                    `json:"sources"`
		Status            *OwnershipInferenceStatus                    `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Checksum == nil {
		return fmt.Errorf("required field checksum missing")
	}
	if all.Confidence == nil {
		return fmt.Errorf("required field confidence missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if !all.EvidenceVersions.IsSet() {
		return fmt.Errorf("required field evidence_versions missing")
	}
	if all.Explanation == nil {
		return fmt.Errorf("required field explanation missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.OwnerType == nil {
		return fmt.Errorf("required field owner_type missing")
	}
	if all.ResourceId == nil {
		return fmt.Errorf("required field resource_id missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"checksum", "confidence", "created_at", "evidence_versions", "explanation", "failed_at", "failure_reason", "id", "owner_type", "primary_contact_ref", "resource_id", "retry_schedule", "sources", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Checksum = *all.Checksum
	o.Confidence = *all.Confidence
	o.CreatedAt = *all.CreatedAt
	o.EvidenceVersions = all.EvidenceVersions
	o.Explanation = *all.Explanation
	o.FailedAt = all.FailedAt
	o.FailureReason = all.FailureReason
	o.Id = *all.Id
	if !all.OwnerType.IsValid() {
		hasInvalidField = true
	} else {
		o.OwnerType = *all.OwnerType
	}
	o.PrimaryContactRef = all.PrimaryContactRef
	o.ResourceId = *all.ResourceId
	o.RetrySchedule = all.RetrySchedule
	o.Sources = *all.Sources
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
