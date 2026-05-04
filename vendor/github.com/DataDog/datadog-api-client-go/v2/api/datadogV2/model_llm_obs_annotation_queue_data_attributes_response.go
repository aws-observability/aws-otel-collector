// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationQueueDataAttributesResponse Attributes of an LLM Observability annotation queue.
type LLMObsAnnotationQueueDataAttributesResponse struct {
	// Timestamp when the queue was created.
	CreatedAt time.Time `json:"created_at"`
	// Identifier of the user who created the queue.
	CreatedBy string `json:"created_by"`
	// Description of the annotation queue.
	Description string `json:"description"`
	// Timestamp when the queue was last modified.
	ModifiedAt time.Time `json:"modified_at"`
	// Identifier of the user who last modified the queue.
	ModifiedBy string `json:"modified_by"`
	// Name of the annotation queue.
	Name string `json:"name"`
	// Identifier of the user who owns the queue.
	OwnedBy string `json:"owned_by"`
	// Identifier of the project this queue belongs to.
	ProjectId string `json:"project_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationQueueDataAttributesResponse instantiates a new LLMObsAnnotationQueueDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationQueueDataAttributesResponse(createdAt time.Time, createdBy string, description string, modifiedAt time.Time, modifiedBy string, name string, ownedBy string, projectId string) *LLMObsAnnotationQueueDataAttributesResponse {
	this := LLMObsAnnotationQueueDataAttributesResponse{}
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.Description = description
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Name = name
	this.OwnedBy = ownedBy
	this.ProjectId = projectId
	return &this
}

// NewLLMObsAnnotationQueueDataAttributesResponseWithDefaults instantiates a new LLMObsAnnotationQueueDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationQueueDataAttributesResponseWithDefaults() *LLMObsAnnotationQueueDataAttributesResponse {
	this := LLMObsAnnotationQueueDataAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetDescription returns the Description field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetDescription(v string) {
	o.Description = v
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetModifiedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetName returns the Name field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetOwnedBy returns the OwnedBy field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetOwnedBy() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.OwnedBy
}

// GetOwnedByOk returns a tuple with the OwnedBy field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetOwnedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnedBy, true
}

// SetOwnedBy sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetOwnedBy(v string) {
	o.OwnedBy = v
}

// GetProjectId returns the ProjectId field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationQueueDataAttributesResponse) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value.
func (o *LLMObsAnnotationQueueDataAttributesResponse) SetProjectId(v string) {
	o.ProjectId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationQueueDataAttributesResponse) MarshalJSON() ([]byte, error) {
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
	toSerialize["description"] = o.Description
	if o.ModifiedAt.Nanosecond() == 0 {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["modified_by"] = o.ModifiedBy
	toSerialize["name"] = o.Name
	toSerialize["owned_by"] = o.OwnedBy
	toSerialize["project_id"] = o.ProjectId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationQueueDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt   *time.Time `json:"created_at"`
		CreatedBy   *string    `json:"created_by"`
		Description *string    `json:"description"`
		ModifiedAt  *time.Time `json:"modified_at"`
		ModifiedBy  *string    `json:"modified_by"`
		Name        *string    `json:"name"`
		OwnedBy     *string    `json:"owned_by"`
		ProjectId   *string    `json:"project_id"`
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
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.ModifiedBy == nil {
		return fmt.Errorf("required field modified_by missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.OwnedBy == nil {
		return fmt.Errorf("required field owned_by missing")
	}
	if all.ProjectId == nil {
		return fmt.Errorf("required field project_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "created_by", "description", "modified_at", "modified_by", "name", "owned_by", "project_id"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.CreatedBy = *all.CreatedBy
	o.Description = *all.Description
	o.ModifiedAt = *all.ModifiedAt
	o.ModifiedBy = *all.ModifiedBy
	o.Name = *all.Name
	o.OwnedBy = *all.OwnedBy
	o.ProjectId = *all.ProjectId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
