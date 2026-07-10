// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsPatternsTopic A topic discovered by an LLM Observability patterns run.
type LLMObsPatternsTopic struct {
	// Timestamp when the topic was created.
	CreatedAt time.Time `json:"created_at"`
	// Description of the topic.
	Description string `json:"description"`
	// Timestamp when the topic was first seen.
	FirstSeenAt time.Time `json:"first_seen_at"`
	// Level of the topic in the hierarchy. Level 0 is a leaf topic.
	HierarchyLevel int64 `json:"hierarchy_level"`
	// Unique identifier of the topic.
	Id string `json:"id"`
	// Whether the topic has been validated.
	IsValidated bool `json:"is_validated"`
	// Name of the topic.
	Name string `json:"name"`
	// Identifier of the parent topic. Empty for top-level topics.
	ParentTopicId string `json:"parent_topic_id"`
	// Number of data points assigned to the topic.
	PointCount int64 `json:"point_count"`
	// Identifier of the run that produced the topic.
	RunId string `json:"run_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsPatternsTopic instantiates a new LLMObsPatternsTopic object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsPatternsTopic(createdAt time.Time, description string, firstSeenAt time.Time, hierarchyLevel int64, id string, isValidated bool, name string, parentTopicId string, pointCount int64, runId string) *LLMObsPatternsTopic {
	this := LLMObsPatternsTopic{}
	this.CreatedAt = createdAt
	this.Description = description
	this.FirstSeenAt = firstSeenAt
	this.HierarchyLevel = hierarchyLevel
	this.Id = id
	this.IsValidated = isValidated
	this.Name = name
	this.ParentTopicId = parentTopicId
	this.PointCount = pointCount
	this.RunId = runId
	return &this
}

// NewLLMObsPatternsTopicWithDefaults instantiates a new LLMObsPatternsTopic object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsPatternsTopicWithDefaults() *LLMObsPatternsTopic {
	this := LLMObsPatternsTopic{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsPatternsTopic) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopic) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsPatternsTopic) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value.
func (o *LLMObsPatternsTopic) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopic) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *LLMObsPatternsTopic) SetDescription(v string) {
	o.Description = v
}

// GetFirstSeenAt returns the FirstSeenAt field value.
func (o *LLMObsPatternsTopic) GetFirstSeenAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.FirstSeenAt
}

// GetFirstSeenAtOk returns a tuple with the FirstSeenAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopic) GetFirstSeenAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstSeenAt, true
}

// SetFirstSeenAt sets field value.
func (o *LLMObsPatternsTopic) SetFirstSeenAt(v time.Time) {
	o.FirstSeenAt = v
}

// GetHierarchyLevel returns the HierarchyLevel field value.
func (o *LLMObsPatternsTopic) GetHierarchyLevel() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.HierarchyLevel
}

// GetHierarchyLevelOk returns a tuple with the HierarchyLevel field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopic) GetHierarchyLevelOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HierarchyLevel, true
}

// SetHierarchyLevel sets field value.
func (o *LLMObsPatternsTopic) SetHierarchyLevel(v int64) {
	o.HierarchyLevel = v
}

// GetId returns the Id field value.
func (o *LLMObsPatternsTopic) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopic) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *LLMObsPatternsTopic) SetId(v string) {
	o.Id = v
}

// GetIsValidated returns the IsValidated field value.
func (o *LLMObsPatternsTopic) GetIsValidated() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IsValidated
}

// GetIsValidatedOk returns a tuple with the IsValidated field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopic) GetIsValidatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsValidated, true
}

// SetIsValidated sets field value.
func (o *LLMObsPatternsTopic) SetIsValidated(v bool) {
	o.IsValidated = v
}

// GetName returns the Name field value.
func (o *LLMObsPatternsTopic) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopic) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsPatternsTopic) SetName(v string) {
	o.Name = v
}

// GetParentTopicId returns the ParentTopicId field value.
func (o *LLMObsPatternsTopic) GetParentTopicId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ParentTopicId
}

// GetParentTopicIdOk returns a tuple with the ParentTopicId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopic) GetParentTopicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentTopicId, true
}

// SetParentTopicId sets field value.
func (o *LLMObsPatternsTopic) SetParentTopicId(v string) {
	o.ParentTopicId = v
}

// GetPointCount returns the PointCount field value.
func (o *LLMObsPatternsTopic) GetPointCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.PointCount
}

// GetPointCountOk returns a tuple with the PointCount field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopic) GetPointCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PointCount, true
}

// SetPointCount sets field value.
func (o *LLMObsPatternsTopic) SetPointCount(v int64) {
	o.PointCount = v
}

// GetRunId returns the RunId field value.
func (o *LLMObsPatternsTopic) GetRunId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value
// and a boolean to check if the value has been set.
func (o *LLMObsPatternsTopic) GetRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RunId, true
}

// SetRunId sets field value.
func (o *LLMObsPatternsTopic) SetRunId(v string) {
	o.RunId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsPatternsTopic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["description"] = o.Description
	if o.FirstSeenAt.Nanosecond() == 0 {
		toSerialize["first_seen_at"] = o.FirstSeenAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["first_seen_at"] = o.FirstSeenAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["hierarchy_level"] = o.HierarchyLevel
	toSerialize["id"] = o.Id
	toSerialize["is_validated"] = o.IsValidated
	toSerialize["name"] = o.Name
	toSerialize["parent_topic_id"] = o.ParentTopicId
	toSerialize["point_count"] = o.PointCount
	toSerialize["run_id"] = o.RunId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsPatternsTopic) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt      *time.Time `json:"created_at"`
		Description    *string    `json:"description"`
		FirstSeenAt    *time.Time `json:"first_seen_at"`
		HierarchyLevel *int64     `json:"hierarchy_level"`
		Id             *string    `json:"id"`
		IsValidated    *bool      `json:"is_validated"`
		Name           *string    `json:"name"`
		ParentTopicId  *string    `json:"parent_topic_id"`
		PointCount     *int64     `json:"point_count"`
		RunId          *string    `json:"run_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.FirstSeenAt == nil {
		return fmt.Errorf("required field first_seen_at missing")
	}
	if all.HierarchyLevel == nil {
		return fmt.Errorf("required field hierarchy_level missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.IsValidated == nil {
		return fmt.Errorf("required field is_validated missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.ParentTopicId == nil {
		return fmt.Errorf("required field parent_topic_id missing")
	}
	if all.PointCount == nil {
		return fmt.Errorf("required field point_count missing")
	}
	if all.RunId == nil {
		return fmt.Errorf("required field run_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "description", "first_seen_at", "hierarchy_level", "id", "is_validated", "name", "parent_topic_id", "point_count", "run_id"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.Description = *all.Description
	o.FirstSeenAt = *all.FirstSeenAt
	o.HierarchyLevel = *all.HierarchyLevel
	o.Id = *all.Id
	o.IsValidated = *all.IsValidated
	o.Name = *all.Name
	o.ParentTopicId = *all.ParentTopicId
	o.PointCount = *all.PointCount
	o.RunId = *all.RunId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
