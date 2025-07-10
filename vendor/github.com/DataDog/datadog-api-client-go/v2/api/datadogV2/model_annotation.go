// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Annotation A list of annotations used in the workflow. These are like sticky notes for your workflow!
type Annotation struct {
	// The definition of `AnnotationDisplay` object.
	Display AnnotationDisplay `json:"display"`
	// The `Annotation` `id`.
	Id string `json:"id"`
	// The definition of `AnnotationMarkdownTextAnnotation` object.
	MarkdownTextAnnotation AnnotationMarkdownTextAnnotation `json:"markdownTextAnnotation"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAnnotation instantiates a new Annotation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAnnotation(display AnnotationDisplay, id string, markdownTextAnnotation AnnotationMarkdownTextAnnotation) *Annotation {
	this := Annotation{}
	this.Display = display
	this.Id = id
	this.MarkdownTextAnnotation = markdownTextAnnotation
	return &this
}

// NewAnnotationWithDefaults instantiates a new Annotation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAnnotationWithDefaults() *Annotation {
	this := Annotation{}
	return &this
}

// GetDisplay returns the Display field value.
func (o *Annotation) GetDisplay() AnnotationDisplay {
	if o == nil {
		var ret AnnotationDisplay
		return ret
	}
	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetDisplayOk() (*AnnotationDisplay, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value.
func (o *Annotation) SetDisplay(v AnnotationDisplay) {
	o.Display = v
}

// GetId returns the Id field value.
func (o *Annotation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *Annotation) SetId(v string) {
	o.Id = v
}

// GetMarkdownTextAnnotation returns the MarkdownTextAnnotation field value.
func (o *Annotation) GetMarkdownTextAnnotation() AnnotationMarkdownTextAnnotation {
	if o == nil {
		var ret AnnotationMarkdownTextAnnotation
		return ret
	}
	return o.MarkdownTextAnnotation
}

// GetMarkdownTextAnnotationOk returns a tuple with the MarkdownTextAnnotation field value
// and a boolean to check if the value has been set.
func (o *Annotation) GetMarkdownTextAnnotationOk() (*AnnotationMarkdownTextAnnotation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarkdownTextAnnotation, true
}

// SetMarkdownTextAnnotation sets field value.
func (o *Annotation) SetMarkdownTextAnnotation(v AnnotationMarkdownTextAnnotation) {
	o.MarkdownTextAnnotation = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Annotation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["display"] = o.Display
	toSerialize["id"] = o.Id
	toSerialize["markdownTextAnnotation"] = o.MarkdownTextAnnotation

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Annotation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Display                *AnnotationDisplay                `json:"display"`
		Id                     *string                           `json:"id"`
		MarkdownTextAnnotation *AnnotationMarkdownTextAnnotation `json:"markdownTextAnnotation"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Display == nil {
		return fmt.Errorf("required field display missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.MarkdownTextAnnotation == nil {
		return fmt.Errorf("required field markdownTextAnnotation missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display", "id", "markdownTextAnnotation"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Display.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Display = *all.Display
	o.Id = *all.Id
	if all.MarkdownTextAnnotation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MarkdownTextAnnotation = *all.MarkdownTextAnnotation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
