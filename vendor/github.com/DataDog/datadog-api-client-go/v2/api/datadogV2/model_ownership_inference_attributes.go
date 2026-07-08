// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OwnershipInferenceAttributes The attributes of a single ownership inference.
type OwnershipInferenceAttributes struct {
	// A checksum that uniquely identifies the current state of the inference. Required when submitting feedback.
	Checksum string `json:"checksum"`
	// The confidence score of the inference, expressed as a numeric string with up to four decimal places.
	Confidence string `json:"confidence"`
	// The time when the inference was created.
	CreatedAt time.Time `json:"created_at"`
	// The list of evidence versions associated with an inference.
	EvidenceVersions datadog.NullableList[map[string]interface{}] `json:"evidence_versions"`
	// A human-readable explanation of how the inference was produced.
	Explanation string `json:"explanation"`
	// The owner type for an ownership inference.
	OwnerType OwnershipOwnerType `json:"owner_type"`
	// The primary contact reference for the inferred owner, formatted as `ref:handle/<owner_handle>`.
	PrimaryContactRef datadog.NullableString `json:"primary_contact_ref,omitempty"`
	// The list of sources backing an ownership inference. Empty when the inference status is not whitelisted to expose sources.
	Sources []map[string]interface{} `json:"sources"`
	// The lifecycle status of an ownership inference.
	Status OwnershipInferenceStatus `json:"status"`
	// The time when the inference was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOwnershipInferenceAttributes instantiates a new OwnershipInferenceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOwnershipInferenceAttributes(checksum string, confidence string, createdAt time.Time, evidenceVersions datadog.NullableList[map[string]interface{}], explanation string, ownerType OwnershipOwnerType, sources []map[string]interface{}, status OwnershipInferenceStatus, updatedAt time.Time) *OwnershipInferenceAttributes {
	this := OwnershipInferenceAttributes{}
	this.Checksum = checksum
	this.Confidence = confidence
	this.CreatedAt = createdAt
	this.EvidenceVersions = evidenceVersions
	this.Explanation = explanation
	this.OwnerType = ownerType
	this.Sources = sources
	this.Status = status
	this.UpdatedAt = updatedAt
	return &this
}

// NewOwnershipInferenceAttributesWithDefaults instantiates a new OwnershipInferenceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOwnershipInferenceAttributesWithDefaults() *OwnershipInferenceAttributes {
	this := OwnershipInferenceAttributes{}
	return &this
}

// GetChecksum returns the Checksum field value.
func (o *OwnershipInferenceAttributes) GetChecksum() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Checksum
}

// GetChecksumOk returns a tuple with the Checksum field value
// and a boolean to check if the value has been set.
func (o *OwnershipInferenceAttributes) GetChecksumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Checksum, true
}

// SetChecksum sets field value.
func (o *OwnershipInferenceAttributes) SetChecksum(v string) {
	o.Checksum = v
}

// GetConfidence returns the Confidence field value.
func (o *OwnershipInferenceAttributes) GetConfidence() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *OwnershipInferenceAttributes) GetConfidenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value.
func (o *OwnershipInferenceAttributes) SetConfidence(v string) {
	o.Confidence = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *OwnershipInferenceAttributes) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *OwnershipInferenceAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *OwnershipInferenceAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetEvidenceVersions returns the EvidenceVersions field value.
// If the value is explicit nil, the zero value for []map[string]interface{} will be returned.
func (o *OwnershipInferenceAttributes) GetEvidenceVersions() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.EvidenceVersions.Get()
}

// GetEvidenceVersionsOk returns a tuple with the EvidenceVersions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OwnershipInferenceAttributes) GetEvidenceVersionsOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.EvidenceVersions.Get(), o.EvidenceVersions.IsSet()
}

// SetEvidenceVersions sets field value.
func (o *OwnershipInferenceAttributes) SetEvidenceVersions(v []map[string]interface{}) {
	o.EvidenceVersions.Set(&v)
}

// GetExplanation returns the Explanation field value.
func (o *OwnershipInferenceAttributes) GetExplanation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Explanation
}

// GetExplanationOk returns a tuple with the Explanation field value
// and a boolean to check if the value has been set.
func (o *OwnershipInferenceAttributes) GetExplanationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Explanation, true
}

// SetExplanation sets field value.
func (o *OwnershipInferenceAttributes) SetExplanation(v string) {
	o.Explanation = v
}

// GetOwnerType returns the OwnerType field value.
func (o *OwnershipInferenceAttributes) GetOwnerType() OwnershipOwnerType {
	if o == nil {
		var ret OwnershipOwnerType
		return ret
	}
	return o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value
// and a boolean to check if the value has been set.
func (o *OwnershipInferenceAttributes) GetOwnerTypeOk() (*OwnershipOwnerType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerType, true
}

// SetOwnerType sets field value.
func (o *OwnershipInferenceAttributes) SetOwnerType(v OwnershipOwnerType) {
	o.OwnerType = v
}

// GetPrimaryContactRef returns the PrimaryContactRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OwnershipInferenceAttributes) GetPrimaryContactRef() string {
	if o == nil || o.PrimaryContactRef.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrimaryContactRef.Get()
}

// GetPrimaryContactRefOk returns a tuple with the PrimaryContactRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OwnershipInferenceAttributes) GetPrimaryContactRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimaryContactRef.Get(), o.PrimaryContactRef.IsSet()
}

// HasPrimaryContactRef returns a boolean if a field has been set.
func (o *OwnershipInferenceAttributes) HasPrimaryContactRef() bool {
	return o != nil && o.PrimaryContactRef.IsSet()
}

// SetPrimaryContactRef gets a reference to the given datadog.NullableString and assigns it to the PrimaryContactRef field.
func (o *OwnershipInferenceAttributes) SetPrimaryContactRef(v string) {
	o.PrimaryContactRef.Set(&v)
}

// SetPrimaryContactRefNil sets the value for PrimaryContactRef to be an explicit nil.
func (o *OwnershipInferenceAttributes) SetPrimaryContactRefNil() {
	o.PrimaryContactRef.Set(nil)
}

// UnsetPrimaryContactRef ensures that no value is present for PrimaryContactRef, not even an explicit nil.
func (o *OwnershipInferenceAttributes) UnsetPrimaryContactRef() {
	o.PrimaryContactRef.Unset()
}

// GetSources returns the Sources field value.
func (o *OwnershipInferenceAttributes) GetSources() []map[string]interface{} {
	if o == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *OwnershipInferenceAttributes) GetSourcesOk() (*[]map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *OwnershipInferenceAttributes) SetSources(v []map[string]interface{}) {
	o.Sources = v
}

// GetStatus returns the Status field value.
func (o *OwnershipInferenceAttributes) GetStatus() OwnershipInferenceStatus {
	if o == nil {
		var ret OwnershipInferenceStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OwnershipInferenceAttributes) GetStatusOk() (*OwnershipInferenceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *OwnershipInferenceAttributes) SetStatus(v OwnershipInferenceStatus) {
	o.Status = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *OwnershipInferenceAttributes) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OwnershipInferenceAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *OwnershipInferenceAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OwnershipInferenceAttributes) MarshalJSON() ([]byte, error) {
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
	toSerialize["owner_type"] = o.OwnerType
	if o.PrimaryContactRef.IsSet() {
		toSerialize["primary_contact_ref"] = o.PrimaryContactRef.Get()
	}
	toSerialize["sources"] = o.Sources
	toSerialize["status"] = o.Status
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
func (o *OwnershipInferenceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Checksum          *string                                      `json:"checksum"`
		Confidence        *string                                      `json:"confidence"`
		CreatedAt         *time.Time                                   `json:"created_at"`
		EvidenceVersions  datadog.NullableList[map[string]interface{}] `json:"evidence_versions"`
		Explanation       *string                                      `json:"explanation"`
		OwnerType         *OwnershipOwnerType                          `json:"owner_type"`
		PrimaryContactRef datadog.NullableString                       `json:"primary_contact_ref,omitempty"`
		Sources           *[]map[string]interface{}                    `json:"sources"`
		Status            *OwnershipInferenceStatus                    `json:"status"`
		UpdatedAt         *time.Time                                   `json:"updated_at"`
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
	if all.OwnerType == nil {
		return fmt.Errorf("required field owner_type missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"checksum", "confidence", "created_at", "evidence_versions", "explanation", "owner_type", "primary_contact_ref", "sources", "status", "updated_at"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Checksum = *all.Checksum
	o.Confidence = *all.Confidence
	o.CreatedAt = *all.CreatedAt
	o.EvidenceVersions = all.EvidenceVersions
	o.Explanation = *all.Explanation
	if !all.OwnerType.IsValid() {
		hasInvalidField = true
	} else {
		o.OwnerType = *all.OwnerType
	}
	o.PrimaryContactRef = all.PrimaryContactRef
	o.Sources = *all.Sources
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
