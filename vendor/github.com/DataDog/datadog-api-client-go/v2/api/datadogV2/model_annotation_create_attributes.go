// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AnnotationCreateAttributes Attributes for creating or updating an annotation.
type AnnotationCreateAttributes struct {
	// Color used to render the annotation in the UI.
	Color AnnotationColor `json:"color"`
	// User-defined text attached to the annotation.
	Description string `json:"description"`
	// End time of the annotation in milliseconds since the Unix epoch. Required for `timeRegion` annotations; omit or set to null for `pointInTime` annotations.
	EndTime datadog.NullableInt64 `json:"end_time,omitempty"`
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

// NewAnnotationCreateAttributes instantiates a new AnnotationCreateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnnotationCreateAttributes(color AnnotationColor, description string, pageId string, startTime int64, typeVar AnnotationKind) *AnnotationCreateAttributes {
	this := AnnotationCreateAttributes{}
	this.Color = color
	this.Description = description
	this.PageId = pageId
	this.StartTime = startTime
	this.Type = typeVar
	return &this
}

// NewAnnotationCreateAttributesWithDefaults instantiates a new AnnotationCreateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnnotationCreateAttributesWithDefaults() *AnnotationCreateAttributes {
	this := AnnotationCreateAttributes{}
	return &this
}

// GetColor returns the Color field value.
func (o *AnnotationCreateAttributes) GetColor() AnnotationColor {
	if o == nil {
		var ret AnnotationColor
		return ret
	}
	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *AnnotationCreateAttributes) GetColorOk() (*AnnotationColor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value.
func (o *AnnotationCreateAttributes) SetColor(v AnnotationColor) {
	o.Color = v
}

// GetDescription returns the Description field value.
func (o *AnnotationCreateAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AnnotationCreateAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *AnnotationCreateAttributes) SetDescription(v string) {
	o.Description = v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AnnotationCreateAttributes) GetEndTime() int64 {
	if o == nil || o.EndTime.Get() == nil {
		var ret int64
		return ret
	}
	return *o.EndTime.Get()
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AnnotationCreateAttributes) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndTime.Get(), o.EndTime.IsSet()
}

// HasEndTime returns a boolean if a field has been set.
func (o *AnnotationCreateAttributes) HasEndTime() bool {
	return o != nil && o.EndTime.IsSet()
}

// SetEndTime gets a reference to the given datadog.NullableInt64 and assigns it to the EndTime field.
func (o *AnnotationCreateAttributes) SetEndTime(v int64) {
	o.EndTime.Set(&v)
}

// SetEndTimeNil sets the value for EndTime to be an explicit nil.
func (o *AnnotationCreateAttributes) SetEndTimeNil() {
	o.EndTime.Set(nil)
}

// UnsetEndTime ensures that no value is present for EndTime, not even an explicit nil.
func (o *AnnotationCreateAttributes) UnsetEndTime() {
	o.EndTime.Unset()
}

// GetPageId returns the PageId field value.
func (o *AnnotationCreateAttributes) GetPageId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value
// and a boolean to check if the value has been set.
func (o *AnnotationCreateAttributes) GetPageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageId, true
}

// SetPageId sets field value.
func (o *AnnotationCreateAttributes) SetPageId(v string) {
	o.PageId = v
}

// GetStartTime returns the StartTime field value.
func (o *AnnotationCreateAttributes) GetStartTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *AnnotationCreateAttributes) GetStartTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value.
func (o *AnnotationCreateAttributes) SetStartTime(v int64) {
	o.StartTime = v
}

// GetType returns the Type field value.
func (o *AnnotationCreateAttributes) GetType() AnnotationKind {
	if o == nil {
		var ret AnnotationKind
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AnnotationCreateAttributes) GetTypeOk() (*AnnotationKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AnnotationCreateAttributes) SetType(v AnnotationKind) {
	o.Type = v
}

// GetWidgetIds returns the WidgetIds field value if set, zero value otherwise.
func (o *AnnotationCreateAttributes) GetWidgetIds() []string {
	if o == nil || o.WidgetIds == nil {
		var ret []string
		return ret
	}
	return o.WidgetIds
}

// GetWidgetIdsOk returns a tuple with the WidgetIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnotationCreateAttributes) GetWidgetIdsOk() (*[]string, bool) {
	if o == nil || o.WidgetIds == nil {
		return nil, false
	}
	return &o.WidgetIds, true
}

// HasWidgetIds returns a boolean if a field has been set.
func (o *AnnotationCreateAttributes) HasWidgetIds() bool {
	return o != nil && o.WidgetIds != nil
}

// SetWidgetIds gets a reference to the given []string and assigns it to the WidgetIds field.
func (o *AnnotationCreateAttributes) SetWidgetIds(v []string) {
	o.WidgetIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AnnotationCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["color"] = o.Color
	toSerialize["description"] = o.Description
	if o.EndTime.IsSet() {
		toSerialize["end_time"] = o.EndTime.Get()
	}
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
func (o *AnnotationCreateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Color       *AnnotationColor      `json:"color"`
		Description *string               `json:"description"`
		EndTime     datadog.NullableInt64 `json:"end_time,omitempty"`
		PageId      *string               `json:"page_id"`
		StartTime   *int64                `json:"start_time"`
		Type        *AnnotationKind       `json:"type"`
		WidgetIds   []string              `json:"widget_ids,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Color == nil {
		return fmt.Errorf("required field color missing")
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
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
		datadog.DeleteKeys(additionalProperties, &[]string{"color", "description", "end_time", "page_id", "start_time", "type", "widget_ids"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.Color.IsValid() {
		hasInvalidField = true
	} else {
		o.Color = *all.Color
	}
	o.Description = *all.Description
	o.EndTime = all.EndTime
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
