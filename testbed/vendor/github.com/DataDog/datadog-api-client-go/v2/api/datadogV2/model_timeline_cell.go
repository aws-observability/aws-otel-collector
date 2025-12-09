// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TimelineCell timeline cell
type TimelineCell struct {
	// author of the timeline cell
	Author *TimelineCellAuthor `json:"author,omitempty"`
	// timeline cell content
	CellContent *TimelineCellContent `json:"cell_content,omitempty"`
	// Timestamp of when the cell was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Timestamp of when the cell was deleted
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// Timestamp of when the cell was last modified
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// Timeline cell content type
	Type *TimelineCellType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTimelineCell instantiates a new TimelineCell object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTimelineCell() *TimelineCell {
	this := TimelineCell{}
	return &this
}

// NewTimelineCellWithDefaults instantiates a new TimelineCell object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTimelineCellWithDefaults() *TimelineCell {
	this := TimelineCell{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *TimelineCell) GetAuthor() TimelineCellAuthor {
	if o == nil || o.Author == nil {
		var ret TimelineCellAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCell) GetAuthorOk() (*TimelineCellAuthor, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *TimelineCell) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given TimelineCellAuthor and assigns it to the Author field.
func (o *TimelineCell) SetAuthor(v TimelineCellAuthor) {
	o.Author = &v
}

// GetCellContent returns the CellContent field value if set, zero value otherwise.
func (o *TimelineCell) GetCellContent() TimelineCellContent {
	if o == nil || o.CellContent == nil {
		var ret TimelineCellContent
		return ret
	}
	return *o.CellContent
}

// GetCellContentOk returns a tuple with the CellContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCell) GetCellContentOk() (*TimelineCellContent, bool) {
	if o == nil || o.CellContent == nil {
		return nil, false
	}
	return o.CellContent, true
}

// HasCellContent returns a boolean if a field has been set.
func (o *TimelineCell) HasCellContent() bool {
	return o != nil && o.CellContent != nil
}

// SetCellContent gets a reference to the given TimelineCellContent and assigns it to the CellContent field.
func (o *TimelineCell) SetCellContent(v TimelineCellContent) {
	o.CellContent = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TimelineCell) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCell) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TimelineCell) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *TimelineCell) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *TimelineCell) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCell) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *TimelineCell) HasDeletedAt() bool {
	return o != nil && o.DeletedAt != nil
}

// SetDeletedAt gets a reference to the given time.Time and assigns it to the DeletedAt field.
func (o *TimelineCell) SetDeletedAt(v time.Time) {
	o.DeletedAt = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *TimelineCell) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCell) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *TimelineCell) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *TimelineCell) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TimelineCell) GetType() TimelineCellType {
	if o == nil || o.Type == nil {
		var ret TimelineCellType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimelineCell) GetTypeOk() (*TimelineCellType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TimelineCell) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given TimelineCellType and assigns it to the Type field.
func (o *TimelineCell) SetType(v TimelineCellType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TimelineCell) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.CellContent != nil {
		toSerialize["cell_content"] = o.CellContent
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.DeletedAt != nil {
		if o.DeletedAt.Nanosecond() == 0 {
			toSerialize["deleted_at"] = o.DeletedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["deleted_at"] = o.DeletedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TimelineCell) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author      *TimelineCellAuthor  `json:"author,omitempty"`
		CellContent *TimelineCellContent `json:"cell_content,omitempty"`
		CreatedAt   *time.Time           `json:"created_at,omitempty"`
		DeletedAt   *time.Time           `json:"deleted_at,omitempty"`
		ModifiedAt  *time.Time           `json:"modified_at,omitempty"`
		Type        *TimelineCellType    `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "cell_content", "created_at", "deleted_at", "modified_at", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Author = all.Author
	o.CellContent = all.CellContent
	o.CreatedAt = all.CreatedAt
	o.DeletedAt = all.DeletedAt
	o.ModifiedAt = all.ModifiedAt
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
