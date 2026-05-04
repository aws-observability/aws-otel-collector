// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationItem A single annotation on an interaction.
type LLMObsAnnotationItem struct {
	// Timestamp when the annotation was created.
	CreatedAt time.Time `json:"created_at"`
	// Identifier of the user who created the annotation.
	CreatedBy string `json:"created_by"`
	// Unique identifier of the annotation.
	Id string `json:"id"`
	// Identifier of the interaction this annotation belongs to.
	InteractionId string `json:"interaction_id"`
	// The label values for this annotation.
	LabelValues map[string]interface{} `json:"label_values"`
	// Timestamp when the annotation was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// Identifier of the user who last modified the annotation.
	ModifiedBy string `json:"modified_by"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationItem instantiates a new LLMObsAnnotationItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationItem(createdAt time.Time, createdBy string, id string, interactionId string, labelValues map[string]interface{}, modifiedAt time.Time, modifiedBy string) *LLMObsAnnotationItem {
	this := LLMObsAnnotationItem{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Id = id
	this.InteractionId = interactionId
	this.LabelValues = labelValues
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	return &this
}

// NewLLMObsAnnotationItemWithDefaults instantiates a new LLMObsAnnotationItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationItemWithDefaults() *LLMObsAnnotationItem {
	this := LLMObsAnnotationItem{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsAnnotationItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsAnnotationItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *LLMObsAnnotationItem) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationItem) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *LLMObsAnnotationItem) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetId returns the Id field value.
func (o *LLMObsAnnotationItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsAnnotationItem) SetId(v string) {
	o.Id = v
}

// GetInteractionId returns the InteractionId field value.
func (o *LLMObsAnnotationItem) GetInteractionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InteractionId
}

// GetInteractionIdOk returns a tuple with the InteractionId field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationItem) GetInteractionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InteractionId, true
}

// SetInteractionId sets field value.
func (o *LLMObsAnnotationItem) SetInteractionId(v string) {
	o.InteractionId = v
}

// GetLabelValues returns the LabelValues field value.
func (o *LLMObsAnnotationItem) GetLabelValues() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.LabelValues
}

// GetLabelValuesOk returns a tuple with the LabelValues field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationItem) GetLabelValuesOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelValues, true
}

// SetLabelValues sets field value.
func (o *LLMObsAnnotationItem) SetLabelValues(v map[string]interface{}) {
	o.LabelValues = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *LLMObsAnnotationItem) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationItem) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *LLMObsAnnotationItem) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *LLMObsAnnotationItem) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationItem) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *LLMObsAnnotationItem) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["created_by"] = o.CreatedBy
	toSerialize["id"] = o.Id
	toSerialize["interaction_id"] = o.InteractionId
	toSerialize["label_values"] = o.LabelValues
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["modified_by"] = o.ModifiedBy

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt     *time.Time              `json:"created_at"`
		CreatedBy     *string                 `json:"created_by"`
		Id            *string                 `json:"id"`
		InteractionId *string                 `json:"interaction_id"`
		LabelValues   *map[string]interface{} `json:"label_values"`
		ModifiedAt    *time.Time              `json:"modified_at"`
		ModifiedBy    *string                 `json:"modified_by"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CreatedBy == nil {
		return fmt.Errorf("required field created_by missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.InteractionId == nil {
		return fmt.Errorf("required field interaction_id missing")
	}
	if all.LabelValues == nil {
		return fmt.Errorf("required field label_values missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "id", "interaction_id", "label_values", "modified_at", "modified_by"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Id = *all.Id
	o.InteractionId = *all.InteractionId
	o.LabelValues = *all.LabelValues
	o.ModifiedAt = *all.ModifiedAt
	o.ModifiedBy = *all.ModifiedBy

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
