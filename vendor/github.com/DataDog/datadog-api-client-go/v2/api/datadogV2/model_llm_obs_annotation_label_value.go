// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationLabelValue A single label value entry in an annotation.
// The `value` type must match the label schema type:
// - `score`: a number within the schema `min`/`max` range (integer if `is_integer` is `true`).
// - `categorical`: a string that is one of the schema `values`.
// - `boolean`: `true` or `false`.
// - `text`: any non-empty string.
type LLMObsAnnotationLabelValue struct {
	// Assessment result for a label value.
	Assessment *LLMObsAnnotationAssessment `json:"assessment,omitempty"`
	// ID of the label schema this value corresponds to.
	LabelSchemaId string `json:"label_schema_id"`
	// Free text reasoning for this label value.
	Reasoning *string `json:"reasoning,omitempty"`
	// The value for this label. Must comply with the label schema type constraints.
	Value LLMObsAnnotationLabelValueValue `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationLabelValue instantiates a new LLMObsAnnotationLabelValue object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationLabelValue(labelSchemaId string, value LLMObsAnnotationLabelValueValue) *LLMObsAnnotationLabelValue {
	this := LLMObsAnnotationLabelValue{}
	this.LabelSchemaId = labelSchemaId
	this.Value = value
	return &this
}

// NewLLMObsAnnotationLabelValueWithDefaults instantiates a new LLMObsAnnotationLabelValue object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationLabelValueWithDefaults() *LLMObsAnnotationLabelValue {
	this := LLMObsAnnotationLabelValue{}
	return &this
}

// GetAssessment returns the Assessment field value if set, zero value otherwise.
func (o *LLMObsAnnotationLabelValue) GetAssessment() LLMObsAnnotationAssessment {
	if o == nil || o.Assessment == nil {
		var ret LLMObsAnnotationAssessment
		return ret
	}
	return *o.Assessment
}

// GetAssessmentOk returns a tuple with the Assessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationLabelValue) GetAssessmentOk() (*LLMObsAnnotationAssessment, bool) {
	if o == nil || o.Assessment == nil {
		return nil, false
	}
	return o.Assessment, true
}

// HasAssessment returns a boolean if a field has been set.
func (o *LLMObsAnnotationLabelValue) HasAssessment() bool {
	return o != nil && o.Assessment != nil
}

// SetAssessment gets a reference to the given LLMObsAnnotationAssessment and assigns it to the Assessment field.
func (o *LLMObsAnnotationLabelValue) SetAssessment(v LLMObsAnnotationAssessment) {
	o.Assessment = &v
}

// GetLabelSchemaId returns the LabelSchemaId field value.
func (o *LLMObsAnnotationLabelValue) GetLabelSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LabelSchemaId
}

// GetLabelSchemaIdOk returns a tuple with the LabelSchemaId field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationLabelValue) GetLabelSchemaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelSchemaId, true
}

// SetLabelSchemaId sets field value.
func (o *LLMObsAnnotationLabelValue) SetLabelSchemaId(v string) {
	o.LabelSchemaId = v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *LLMObsAnnotationLabelValue) GetReasoning() string {
	if o == nil || o.Reasoning == nil {
		var ret string
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationLabelValue) GetReasoningOk() (*string, bool) {
	if o == nil || o.Reasoning == nil {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *LLMObsAnnotationLabelValue) HasReasoning() bool {
	return o != nil && o.Reasoning != nil
}

// SetReasoning gets a reference to the given string and assigns it to the Reasoning field.
func (o *LLMObsAnnotationLabelValue) SetReasoning(v string) {
	o.Reasoning = &v
}

// GetValue returns the Value field value.
func (o *LLMObsAnnotationLabelValue) GetValue() LLMObsAnnotationLabelValueValue {
	if o == nil {
		var ret LLMObsAnnotationLabelValueValue
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationLabelValue) GetValueOk() (*LLMObsAnnotationLabelValueValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *LLMObsAnnotationLabelValue) SetValue(v LLMObsAnnotationLabelValueValue) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationLabelValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assessment != nil {
		toSerialize["assessment"] = o.Assessment
	}
	toSerialize["label_schema_id"] = o.LabelSchemaId
	if o.Reasoning != nil {
		toSerialize["reasoning"] = o.Reasoning
	}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationLabelValue) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assessment    *LLMObsAnnotationAssessment      `json:"assessment,omitempty"`
		LabelSchemaId *string                          `json:"label_schema_id"`
		Reasoning     *string                          `json:"reasoning,omitempty"`
		Value         *LLMObsAnnotationLabelValueValue `json:"value"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.LabelSchemaId == nil {
		return fmt.Errorf("required field label_schema_id missing")
	}
	if all.Value == nil {
		return fmt.Errorf("required field value missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assessment", "label_schema_id", "reasoning", "value"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Assessment != nil && !all.Assessment.IsValid() {
		hasInvalidField = true
	} else {
		o.Assessment = all.Assessment
	}
	o.LabelSchemaId = *all.LabelSchemaId
	o.Reasoning = all.Reasoning
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
