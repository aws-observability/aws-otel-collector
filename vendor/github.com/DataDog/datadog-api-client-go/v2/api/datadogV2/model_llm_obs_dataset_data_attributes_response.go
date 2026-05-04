// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetDataAttributesResponse Attributes of an LLM Observability dataset.
type LLMObsDatasetDataAttributesResponse struct {
	// Timestamp when the dataset was created.
	CreatedAt time.Time `json:"created_at"`
	// Current version number of the dataset.
	CurrentVersion int64 `json:"current_version"`
	// Description of the dataset.
	Description datadog.NullableString `json:"description"`
	// Arbitrary metadata associated with the dataset.
	Metadata map[string]interface{} `json:"metadata"`
	// Name of the dataset.
	Name string `json:"name"`
	// Timestamp when the dataset was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetDataAttributesResponse instantiates a new LLMObsDatasetDataAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetDataAttributesResponse(createdAt time.Time, currentVersion int64, description datadog.NullableString, metadata map[string]interface{}, name string, updatedAt time.Time) *LLMObsDatasetDataAttributesResponse {
	this := LLMObsDatasetDataAttributesResponse{}
	this.CreatedAt = createdAt
	this.CurrentVersion = currentVersion
	this.Description = description
	this.Metadata = metadata
	this.Name = name
	this.UpdatedAt = updatedAt
	return &this
}

// NewLLMObsDatasetDataAttributesResponseWithDefaults instantiates a new LLMObsDatasetDataAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetDataAttributesResponseWithDefaults() *LLMObsDatasetDataAttributesResponse {
	this := LLMObsDatasetDataAttributesResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value.
func (o *LLMObsDatasetDataAttributesResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetDataAttributesResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *LLMObsDatasetDataAttributesResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCurrentVersion returns the CurrentVersion field value.
func (o *LLMObsDatasetDataAttributesResponse) GetCurrentVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CurrentVersion
}

// GetCurrentVersionOk returns a tuple with the CurrentVersion field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetDataAttributesResponse) GetCurrentVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentVersion, true
}

// SetCurrentVersion sets field value.
func (o *LLMObsDatasetDataAttributesResponse) SetCurrentVersion(v int64) {
	o.CurrentVersion = v
}

// GetDescription returns the Description field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *LLMObsDatasetDataAttributesResponse) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetDataAttributesResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value.
func (o *LLMObsDatasetDataAttributesResponse) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetMetadata returns the Metadata field value.
// If the value is explicit nil, the zero value for map[string]interface{} will be returned.
func (o *LLMObsDatasetDataAttributesResponse) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetDataAttributesResponse) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value.
func (o *LLMObsDatasetDataAttributesResponse) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetName returns the Name field value.
func (o *LLMObsDatasetDataAttributesResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetDataAttributesResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsDatasetDataAttributesResponse) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value.
func (o *LLMObsDatasetDataAttributesResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetDataAttributesResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value.
func (o *LLMObsDatasetDataAttributesResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetDataAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedAt.Nanosecond() == 0 {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
	}
	toSerialize["current_version"] = o.CurrentVersion
	toSerialize["description"] = o.Description.Get()
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
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
func (o *LLMObsDatasetDataAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedAt      *time.Time             `json:"created_at"`
		CurrentVersion *int64                 `json:"current_version"`
		Description    datadog.NullableString `json:"description"`
		Metadata       map[string]interface{} `json:"metadata"`
		Name           *string                `json:"name"`
		UpdatedAt      *time.Time             `json:"updated_at"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.CurrentVersion == nil {
		return fmt.Errorf("required field current_version missing")
	}
	if !all.Description.IsSet() {
		return fmt.Errorf("required field description missing")
	}
	if all.Metadata == nil {
		return fmt.Errorf("required field metadata missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.UpdatedAt == nil {
		return fmt.Errorf("required field updated_at missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_at", "current_version", "description", "metadata", "name", "updated_at"})
	} else {
		return err
	}
	o.CreatedAt = *all.CreatedAt
	o.CurrentVersion = *all.CurrentVersion
	o.Description = all.Description
	o.Metadata = all.Metadata
	o.Name = *all.Name
	o.UpdatedAt = *all.UpdatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
