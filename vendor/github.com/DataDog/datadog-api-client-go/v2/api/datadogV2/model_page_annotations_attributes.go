// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PageAnnotationsAttributes Attributes of the annotations on a page.
type PageAnnotationsAttributes struct {
	// Map of annotation UUID to annotation object, keyed by annotation ID.
	Annotations map[string]AnnotationInPage `json:"annotations"`
	// List of annotation IDs that apply to the entire page rather than a specific widget.
	GlobalAnnotations []uuid.UUID `json:"global_annotations"`
	// Map from widget ID to the list of annotation IDs displayed on that widget.
	WidgetMapping map[string][]uuid.UUID `json:"widget_mapping"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPageAnnotationsAttributes instantiates a new PageAnnotationsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPageAnnotationsAttributes(annotations map[string]AnnotationInPage, globalAnnotations []uuid.UUID, widgetMapping map[string][]uuid.UUID) *PageAnnotationsAttributes {
	this := PageAnnotationsAttributes{}
	this.Annotations = annotations
	this.GlobalAnnotations = globalAnnotations
	this.WidgetMapping = widgetMapping
	return &this
}

// NewPageAnnotationsAttributesWithDefaults instantiates a new PageAnnotationsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPageAnnotationsAttributesWithDefaults() *PageAnnotationsAttributes {
	this := PageAnnotationsAttributes{}
	return &this
}

// GetAnnotations returns the Annotations field value.
func (o *PageAnnotationsAttributes) GetAnnotations() map[string]AnnotationInPage {
	if o == nil {
		var ret map[string]AnnotationInPage
		return ret
	}
	return o.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value
// and a boolean to check if the value has been set.
func (o *PageAnnotationsAttributes) GetAnnotationsOk() (*map[string]AnnotationInPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Annotations, true
}

// SetAnnotations sets field value.
func (o *PageAnnotationsAttributes) SetAnnotations(v map[string]AnnotationInPage) {
	o.Annotations = v
}

// GetGlobalAnnotations returns the GlobalAnnotations field value.
func (o *PageAnnotationsAttributes) GetGlobalAnnotations() []uuid.UUID {
	if o == nil {
		var ret []uuid.UUID
		return ret
	}
	return o.GlobalAnnotations
}

// GetGlobalAnnotationsOk returns a tuple with the GlobalAnnotations field value
// and a boolean to check if the value has been set.
func (o *PageAnnotationsAttributes) GetGlobalAnnotationsOk() (*[]uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalAnnotations, true
}

// SetGlobalAnnotations sets field value.
func (o *PageAnnotationsAttributes) SetGlobalAnnotations(v []uuid.UUID) {
	o.GlobalAnnotations = v
}

// GetWidgetMapping returns the WidgetMapping field value.
func (o *PageAnnotationsAttributes) GetWidgetMapping() map[string][]uuid.UUID {
	if o == nil {
		var ret map[string][]uuid.UUID
		return ret
	}
	return o.WidgetMapping
}

// GetWidgetMappingOk returns a tuple with the WidgetMapping field value
// and a boolean to check if the value has been set.
func (o *PageAnnotationsAttributes) GetWidgetMappingOk() (*map[string][]uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WidgetMapping, true
}

// SetWidgetMapping sets field value.
func (o *PageAnnotationsAttributes) SetWidgetMapping(v map[string][]uuid.UUID) {
	o.WidgetMapping = v
}

// MarshalJSON serializes the struct using spec logic.
func (o PageAnnotationsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["annotations"] = o.Annotations
	toSerialize["global_annotations"] = o.GlobalAnnotations
	toSerialize["widget_mapping"] = o.WidgetMapping

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PageAnnotationsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Annotations       *map[string]AnnotationInPage `json:"annotations"`
		GlobalAnnotations *[]uuid.UUID                 `json:"global_annotations"`
		WidgetMapping     *map[string][]uuid.UUID      `json:"widget_mapping"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Annotations == nil {
		return fmt.Errorf("required field annotations missing")
	}
	if all.GlobalAnnotations == nil {
		return fmt.Errorf("required field global_annotations missing")
	}
	if all.WidgetMapping == nil {
		return fmt.Errorf("required field widget_mapping missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"annotations", "global_annotations", "widget_mapping"})
	} else {
		return err
	}
	o.Annotations = *all.Annotations
	o.GlobalAnnotations = *all.GlobalAnnotations
	o.WidgetMapping = *all.WidgetMapping

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
