// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnnotationAttributes Attributes of an annotation returned in a response.
type AnnotationAttributes struct {
	// Identifier of the user who created the annotation.
	AuthorId string `json:"author_id"`
	// Color used to render the annotation in the UI.
	Color AnnotationColor `json:"color"`
	// Creation time of the annotation in milliseconds since the Unix epoch.
	CreatedAt int64 `json:"created_at"`
	// User-defined text attached to the annotation.
	Description string `json:"description"`
	// End time of the annotation in milliseconds since the Unix epoch. Null for `pointInTime` annotations.
	EndTime datadog.NullableInt64 `json:"end_time"`
	// Last modification time of the annotation in milliseconds since the Unix epoch.
	ModifiedAt int64 `json:"modified_at"`
	// ID of the page the annotation belongs to, prefixed with the page type and joined by a colon
	// (for example, `dashboard:abc-def-xyz` or `notebook:1234567890`).
	PageId string `json:"page_id"`
	// Start time of the annotation in milliseconds since the Unix epoch.
	StartTime int64 `json:"start_time"`
	// Kind of annotation. `pointInTime` annotations mark a single moment in time,
	// while `timeRegion` annotations span a window of time and require an `end_time`.
	Type AnnotationKind `json:"type"`
	// IDs of widgets the annotation is associated with. When empty or omitted, the annotation applies to the whole page.
	WidgetIds []string `json:"widget_ids,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnnotationAttributes instantiates a new AnnotationAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnnotationAttributes(authorId string, color AnnotationColor, createdAt int64, description string, endTime datadog.NullableInt64, modifiedAt int64, pageId string, startTime int64, typeVar AnnotationKind) *AnnotationAttributes {
	this := AnnotationAttributes{}
	this.AuthorId = authorId
	this.Color = color
	this.CreatedAt = createdAt
	this.Description = description
	this.EndTime = endTime
	this.ModifiedAt = modifiedAt
	this.PageId = pageId
	this.StartTime = startTime
	this.Type = typeVar
	return &this
}

// NewAnnotationAttributesWithDefaults instantiates a new AnnotationAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnnotationAttributesWithDefaults() *AnnotationAttributes {
	this := AnnotationAttributes{}
	return &this
}

// GetAuthorId returns the AuthorId field value.
func (o *AnnotationAttributes) GetAuthorId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value
// and a boolean to check if the value has been set.
func (o *AnnotationAttributes) GetAuthorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthorId, true
}

// SetAuthorId sets field value.
func (o *AnnotationAttributes) SetAuthorId(v string) {
	o.AuthorId = v
}

// GetColor returns the Color field value.
func (o *AnnotationAttributes) GetColor() AnnotationColor {
	if o == nil {
		var ret AnnotationColor
		return ret
	}
	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *AnnotationAttributes) GetColorOk() (*AnnotationColor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value.
func (o *AnnotationAttributes) SetColor(v AnnotationColor) {
	o.Color = v
}

// GetCreatedAt returns the CreatedAt field value.
func (o *AnnotationAttributes) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AnnotationAttributes) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value.
func (o *AnnotationAttributes) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value.
func (o *AnnotationAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AnnotationAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *AnnotationAttributes) SetDescription(v string) {
	o.Description = v
}

// GetEndTime returns the EndTime field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *AnnotationAttributes) GetEndTime() int64 {
	if o == nil || o.EndTime.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AnnotationAttributes) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// SetEndTime sets field value.
func (o *AnnotationAttributes) SetEndTime(v int64) {
	o.EndTime.Set(&v)
}

// GetModifiedAt returns the ModifiedAt field value.
func (o *AnnotationAttributes) GetModifiedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *AnnotationAttributes) GetModifiedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value.
func (o *AnnotationAttributes) SetModifiedAt(v int64) {
	o.ModifiedAt = v
}

// GetPageId returns the PageId field value.
func (o *AnnotationAttributes) GetPageId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value
// and a boolean to check if the value has been set.
func (o *AnnotationAttributes) GetPageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageId, true
}

// SetPageId sets field value.
func (o *AnnotationAttributes) SetPageId(v string) {
	o.PageId = v
}

// GetStartTime returns the StartTime field value.
func (o *AnnotationAttributes) GetStartTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *AnnotationAttributes) GetStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value.
func (o *AnnotationAttributes) SetStartTime(v int64) {
	o.StartTime = v
}

// GetType returns the Type field value.
func (o *AnnotationAttributes) GetType() AnnotationKind {
	if o == nil {
		var ret AnnotationKind
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AnnotationAttributes) GetTypeOk() (*AnnotationKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AnnotationAttributes) SetType(v AnnotationKind) {
	o.Type = v
}

// GetWidgetIds returns the WidgetIds field value if set, zero value otherwise.
func (o *AnnotationAttributes) GetWidgetIds() []string {
	if o == nil || o.WidgetIds == nil {
		var ret []string
		return ret
	}
	return o.WidgetIds
}

// GetWidgetIdsOk returns a tuple with the WidgetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationAttributes) GetWidgetIdsOk() (*[]string, bool) {
	if o == nil || o.WidgetIds == nil {
		return nil, false
	}
	return &o.WidgetIds, true
}

// HasWidgetIds returns a boolean if a field has been set.
func (o *AnnotationAttributes) HasWidgetIds() bool {
	return o != nil && o.WidgetIds != nil
}

// SetWidgetIds gets a reference to the given []string and assigns it to the WidgetIds field.
func (o *AnnotationAttributes) SetWidgetIds(v []string) {
	o.WidgetIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnnotationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["author_id"] = o.AuthorId
	toSerialize["color"] = o.Color
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["description"] = o.Description
	toSerialize["end_time"] = o.EndTime.Get()
	toSerialize["modified_at"] = o.ModifiedAt
	toSerialize["page_id"] = o.PageId
	toSerialize["start_time"] = o.StartTime
	toSerialize["type"] = o.Type
	if o.WidgetIds != nil {
		toSerialize["widget_ids"] = o.WidgetIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AnnotationAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthorId    *string               `json:"author_id"`
		Color       *AnnotationColor      `json:"color"`
		CreatedAt   *int64                `json:"created_at"`
		Description *string               `json:"description"`
		EndTime     datadog.NullableInt64 `json:"end_time"`
		ModifiedAt  *int64                `json:"modified_at"`
		PageId      *string               `json:"page_id"`
		StartTime   *int64                `json:"start_time"`
		Type        *AnnotationKind       `json:"type"`
		WidgetIds   []string              `json:"widget_ids,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AuthorId == nil {
		return fmt.Errorf("required field author_id missing")
	}
	if all.Color == nil {
		return fmt.Errorf("required field color missing")
	}
	if all.CreatedAt == nil {
		return fmt.Errorf("required field created_at missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if !all.EndTime.IsSet() {
		return fmt.Errorf("required field end_time missing")
	}
	if all.ModifiedAt == nil {
		return fmt.Errorf("required field modified_at missing")
	}
	if all.PageId == nil {
		return fmt.Errorf("required field page_id missing")
	}
	if all.StartTime == nil {
		return fmt.Errorf("required field start_time missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author_id", "color", "created_at", "description", "end_time", "modified_at", "page_id", "start_time", "type", "widget_ids"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AuthorId = *all.AuthorId
	if !all.Color.IsValid() {
		hasInvalidField = true
	} else {
		o.Color = *all.Color
	}
	o.CreatedAt = *all.CreatedAt
	o.Description = *all.Description
	o.EndTime = all.EndTime
	o.ModifiedAt = *all.ModifiedAt
	o.PageId = *all.PageId
	o.StartTime = *all.StartTime
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.WidgetIds = all.WidgetIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
