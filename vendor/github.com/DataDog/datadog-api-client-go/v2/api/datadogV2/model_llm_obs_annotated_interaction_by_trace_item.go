// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotatedInteractionByTraceItem An annotated interaction returned by the cross-queue lookup, including the source queue metadata.
type LLMObsAnnotatedInteractionByTraceItem struct {
	// List of annotations for this interaction.
	Annotations []LLMObsAnnotationItem `json:"annotations"`
	// Upstream entity identifier (trace ID, session ID, or deterministic display_block ID).
	ContentId string `json:"content_id"`
	// Timestamp when the interaction was added to the queue.
	CreatedAt time.Time `json:"created_at"`
	// List of content blocks that make up a `display_block` interaction.
	// Must contain at least one block.
	DisplayBlock []LLMObsContentBlock `json:"display_block,omitempty"`
	// Unique identifier of the interaction.
	Id string `json:"id"`
	// Timestamp when the interaction was last updated.
	ModifiedAt time.Time `json:"modified_at"`
	// Identifier of the annotation queue this interaction belongs to.
	QueueId string `json:"queue_id"`
	// Name of the annotation queue this interaction belongs to.
	QueueName string `json:"queue_name"`
	// Type of an annotated interaction.
	Type LLMObsAnyInteractionType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotatedInteractionByTraceItem instantiates a new LLMObsAnnotatedInteractionByTraceItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotatedInteractionByTraceItem(annotations []LLMObsAnnotationItem, contentId string, createdAt time.Time, id string, modifiedAt time.Time, queueId string, queueName string, typeVar LLMObsAnyInteractionType) *LLMObsAnnotatedInteractionByTraceItem {
	this := LLMObsAnnotatedInteractionByTraceItem{}
	this.Annotations = annotations
	this.ContentId = contentId
	this.CreatedAt = createdAt
	this.Id = id
	this.ModifiedAt = modifiedAt
	this.QueueId = queueId
	this.QueueName = queueName
	this.Type = typeVar
	return &this
}

// NewLLMObsAnnotatedInteractionByTraceItemWithDefaults instantiates a new LLMObsAnnotatedInteractionByTraceItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotatedInteractionByTraceItemWithDefaults() *LLMObsAnnotatedInteractionByTraceItem {
	this := LLMObsAnnotatedInteractionByTraceItem{}
	return &this
}

// GetAnnotations returns the Annotations field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetAnnotations() []LLMObsAnnotationItem {
	if o == nil {
		var ret []LLMObsAnnotationItem
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetAnnotationsOk() (*[]LLMObsAnnotationItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// SetAnnotations sets field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) SetAnnotations(v []LLMObsAnnotationItem) {
	o.Annotations = v
}

// GetContentId returns the ContentId field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetContentId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ContentId
}

// GetContentIdOk returns a tuple with the ContentId field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetContentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContentId, true
}

// SetContentId sets field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) SetContentId(v string) {
	o.ContentId = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDisplayBlock returns the DisplayBlock field value if set, zero value otherwise.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetDisplayBlock() []LLMObsContentBlock {
	if o == nil || o.DisplayBlock == nil {
		var ret []LLMObsContentBlock
		return ret
	}
	return o.DisplayBlock
}

// GetDisplayBlockOk returns a tuple with the DisplayBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetDisplayBlockOk() (*[]LLMObsContentBlock, bool) {
	if o == nil || o.DisplayBlock == nil {
		return nil, false
	}
	return &o.DisplayBlock, true
}

// HasDisplayBlock returns a boolean if a field has been set.
func (o *LLMObsAnnotatedInteractionByTraceItem) HasDisplayBlock() bool {
	return o != nil && o.DisplayBlock != nil
}

// SetDisplayBlock gets a reference to the given []LLMObsContentBlock and assigns it to the DisplayBlock field.
func (o *LLMObsAnnotatedInteractionByTraceItem) SetDisplayBlock(v []LLMObsContentBlock) {
	o.DisplayBlock = v
}

// GetId returns the Id field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) SetId(v string) {
	o.Id = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetQueueId returns the QueueId field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetQueueId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueueId
}

// GetQueueIdOk returns a tuple with the QueueId field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetQueueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueueId, true
}

// SetQueueId sets field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) SetQueueId(v string) {
	o.QueueId = v
}

// GetQueueName returns the QueueName field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetQueueName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.QueueName
}

// GetQueueNameOk returns a tuple with the QueueName field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetQueueNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueueName, true
}

// SetQueueName sets field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) SetQueueName(v string) {
	o.QueueName = v
}

// GetType returns the Type field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetType() LLMObsAnyInteractionType {
	if o == nil {
		var ret LLMObsAnyInteractionType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotatedInteractionByTraceItem) GetTypeOk() (*LLMObsAnyInteractionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsAnnotatedInteractionByTraceItem) SetType(v LLMObsAnyInteractionType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotatedInteractionByTraceItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["annotations"] = o.Annotations
	toSerialize["content_id"] = o.ContentId
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.DisplayBlock != nil {
		toSerialize["display_block"] = o.DisplayBlock
	}
	toSerialize["id"] = o.Id
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["queue_id"] = o.QueueId
	toSerialize["queue_name"] = o.QueueName
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotatedInteractionByTraceItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Annotations  *[]LLMObsAnnotationItem   `json:"annotations"`
		ContentId    *string                   `json:"content_id"`
		CreatedAt    *time.Time                `json:"created_at"`
		DisplayBlock []LLMObsContentBlock      `json:"display_block,omitempty"`
		Id           *string                   `json:"id"`
		ModifiedAt   *time.Time                `json:"modified_at"`
		QueueId      *string                   `json:"queue_id"`
		QueueName    *string                   `json:"queue_name"`
		Type         *LLMObsAnyInteractionType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Annotations == nil {
		return fmt.Errorf("required field annotations missing")
	}
	if all.ContentId == nil {
		return fmt.Errorf("required field content_id missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.QueueId == nil {
		return fmt.Errorf("required field queue_id missing")
	}
	if all.QueueName == nil {
		return fmt.Errorf("required field queue_name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotations", "content_id", "created_at", "display_block", "id", "modified_at", "queue_id", "queue_name", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Annotations = *all.Annotations
	o.ContentId = *all.ContentId
	o.CreatedAt = *all.CreatedAt
	o.DisplayBlock = all.DisplayBlock
	o.Id = *all.Id
	o.ModifiedAt = *all.ModifiedAt
	o.QueueId = *all.QueueId
	o.QueueName = *all.QueueName
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
