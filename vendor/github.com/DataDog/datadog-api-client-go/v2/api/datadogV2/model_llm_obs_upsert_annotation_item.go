// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsUpsertAnnotationItem A single annotation to create or update. The annotation is matched by
// `interaction_id` and the requesting user's identity.
type LLMObsUpsertAnnotationItem struct {
	// ID of the interaction to annotate.
	InteractionId string `json:"interaction_id"`
	// Label values for this annotation. Each entry references a label schema by ID
	// and provides the corresponding value validated against the schema type constraints.
	LabelValues []LLMObsAnnotationLabelValue `json:"label_values"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsUpsertAnnotationItem instantiates a new LLMObsUpsertAnnotationItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsUpsertAnnotationItem(interactionId string, labelValues []LLMObsAnnotationLabelValue) *LLMObsUpsertAnnotationItem {
	this := LLMObsUpsertAnnotationItem{}
	this.InteractionId = interactionId
	this.LabelValues = labelValues
	return &this
}

// NewLLMObsUpsertAnnotationItemWithDefaults instantiates a new LLMObsUpsertAnnotationItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsUpsertAnnotationItemWithDefaults() *LLMObsUpsertAnnotationItem {
	this := LLMObsUpsertAnnotationItem{}
	return &this
}

// GetInteractionId returns the InteractionId field value.
func (o *LLMObsUpsertAnnotationItem) GetInteractionId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InteractionId
}

// GetInteractionIdOk returns a tuple with the InteractionId field value
// and a boolean to check if the value has been set.
func (o *LLMObsUpsertAnnotationItem) GetInteractionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InteractionId, true
}

// SetInteractionId sets field value.
func (o *LLMObsUpsertAnnotationItem) SetInteractionId(v string) {
	o.InteractionId = v
}

// GetLabelValues returns the LabelValues field value.
func (o *LLMObsUpsertAnnotationItem) GetLabelValues() []LLMObsAnnotationLabelValue {
	if o == nil {
		var ret []LLMObsAnnotationLabelValue
		return ret
	}
	return o.LabelValues
}

// GetLabelValuesOk returns a tuple with the LabelValues field value
// and a boolean to check if the value has been set.
func (o *LLMObsUpsertAnnotationItem) GetLabelValuesOk() (*[]LLMObsAnnotationLabelValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelValues, true
}

// SetLabelValues sets field value.
func (o *LLMObsUpsertAnnotationItem) SetLabelValues(v []LLMObsAnnotationLabelValue) {
	o.LabelValues = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsUpsertAnnotationItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["interaction_id"] = o.InteractionId
	toSerialize["label_values"] = o.LabelValues

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsUpsertAnnotationItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InteractionId *string                       `json:"interaction_id"`
		LabelValues   *[]LLMObsAnnotationLabelValue `json:"label_values"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.InteractionId == nil {
		return fmt.Errorf("required field interaction_id missing")
	}
	if all.LabelValues == nil {
		return fmt.Errorf("required field label_values missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"interaction_id", "label_values"})
	} else {
		return err
	}
	o.InteractionId = *all.InteractionId
	o.LabelValues = *all.LabelValues

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
