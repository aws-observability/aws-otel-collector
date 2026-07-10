// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsAnnotationLabelValueResponse A single label value entry in an annotation response.
// In addition to the submitted fields, the server populates `type` and
// `name_when_saved` to mirror the schema state at the time the annotation
// was created — these help clients display values correctly when the schema
// has since changed.
type LLMObsAnnotationLabelValueResponse struct {
	// Assessment result for a label value.
	Assessment *LLMObsAnnotationAssessment `json:"assessment,omitempty"`
	// ID of the label schema this value corresponds to.
	LabelSchemaId string `json:"label_schema_id"`
	// Name of the label schema at the time the annotation was created.
	NameWhenSaved *string `json:"name_when_saved,omitempty"`
	// Free text reasoning for this label value.
	Reasoning *string `json:"reasoning,omitempty"`
	// Type of a label in an annotation queue label schema.
	Type *LLMObsLabelSchemaType `json:"type,omitempty"`
	// The value for this label. Must comply with the label schema type constraints.
	Value LLMObsAnnotationLabelValueValue `json:"value"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsAnnotationLabelValueResponse instantiates a new LLMObsAnnotationLabelValueResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsAnnotationLabelValueResponse(labelSchemaId string, value LLMObsAnnotationLabelValueValue) *LLMObsAnnotationLabelValueResponse {
	this := LLMObsAnnotationLabelValueResponse{}
	this.LabelSchemaId = labelSchemaId
	this.Value = value
	return &this
}

// NewLLMObsAnnotationLabelValueResponseWithDefaults instantiates a new LLMObsAnnotationLabelValueResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsAnnotationLabelValueResponseWithDefaults() *LLMObsAnnotationLabelValueResponse {
	this := LLMObsAnnotationLabelValueResponse{}
	return &this
}

// GetAssessment returns the Assessment field value if set, zero value otherwise.
func (o *LLMObsAnnotationLabelValueResponse) GetAssessment() LLMObsAnnotationAssessment {
	if o == nil || o.Assessment == nil {
		var ret LLMObsAnnotationAssessment
		return ret
	}
	return *o.Assessment
}

// GetAssessmentOk returns a tuple with the Assessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationLabelValueResponse) GetAssessmentOk() (*LLMObsAnnotationAssessment, bool) {
	if o == nil || o.Assessment == nil {
		return nil, false
	}
	return o.Assessment, true
}

// HasAssessment returns a boolean if a field has been set.
func (o *LLMObsAnnotationLabelValueResponse) HasAssessment() bool {
	return o != nil && o.Assessment != nil
}

// SetAssessment gets a reference to the given LLMObsAnnotationAssessment and assigns it to the Assessment field.
func (o *LLMObsAnnotationLabelValueResponse) SetAssessment(v LLMObsAnnotationAssessment) {
	o.Assessment = &v
}

// GetLabelSchemaId returns the LabelSchemaId field value.
func (o *LLMObsAnnotationLabelValueResponse) GetLabelSchemaId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LabelSchemaId
}

// GetLabelSchemaIdOk returns a tuple with the LabelSchemaId field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationLabelValueResponse) GetLabelSchemaIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelSchemaId, true
}

// SetLabelSchemaId sets field value.
func (o *LLMObsAnnotationLabelValueResponse) SetLabelSchemaId(v string) {
	o.LabelSchemaId = v
}

// GetNameWhenSaved returns the NameWhenSaved field value if set, zero value otherwise.
func (o *LLMObsAnnotationLabelValueResponse) GetNameWhenSaved() string {
	if o == nil || o.NameWhenSaved == nil {
		var ret string
		return ret
	}
	return *o.NameWhenSaved
}

// GetNameWhenSavedOk returns a tuple with the NameWhenSaved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationLabelValueResponse) GetNameWhenSavedOk() (*string, bool) {
	if o == nil || o.NameWhenSaved == nil {
		return nil, false
	}
	return o.NameWhenSaved, true
}

// HasNameWhenSaved returns a boolean if a field has been set.
func (o *LLMObsAnnotationLabelValueResponse) HasNameWhenSaved() bool {
	return o != nil && o.NameWhenSaved != nil
}

// SetNameWhenSaved gets a reference to the given string and assigns it to the NameWhenSaved field.
func (o *LLMObsAnnotationLabelValueResponse) SetNameWhenSaved(v string) {
	o.NameWhenSaved = &v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *LLMObsAnnotationLabelValueResponse) GetReasoning() string {
	if o == nil || o.Reasoning == nil {
		var ret string
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationLabelValueResponse) GetReasoningOk() (*string, bool) {
	if o == nil || o.Reasoning == nil {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *LLMObsAnnotationLabelValueResponse) HasReasoning() bool {
	return o != nil && o.Reasoning != nil
}

// SetReasoning gets a reference to the given string and assigns it to the Reasoning field.
func (o *LLMObsAnnotationLabelValueResponse) SetReasoning(v string) {
	o.Reasoning = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *LLMObsAnnotationLabelValueResponse) GetType() LLMObsLabelSchemaType {
	if o == nil || o.Type == nil {
		var ret LLMObsLabelSchemaType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationLabelValueResponse) GetTypeOk() (*LLMObsLabelSchemaType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *LLMObsAnnotationLabelValueResponse) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given LLMObsLabelSchemaType and assigns it to the Type field.
func (o *LLMObsAnnotationLabelValueResponse) SetType(v LLMObsLabelSchemaType) {
	o.Type = &v
}

// GetValue returns the Value field value.
func (o *LLMObsAnnotationLabelValueResponse) GetValue() LLMObsAnnotationLabelValueValue {
	if o == nil {
		var ret LLMObsAnnotationLabelValueValue
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *LLMObsAnnotationLabelValueResponse) GetValueOk() (*LLMObsAnnotationLabelValueValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value.
func (o *LLMObsAnnotationLabelValueResponse) SetValue(v LLMObsAnnotationLabelValueValue) {
	o.Value = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsAnnotationLabelValueResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assessment != nil {
		toSerialize["assessment"] = o.Assessment
	}
	toSerialize["label_schema_id"] = o.LabelSchemaId
	if o.NameWhenSaved != nil {
		toSerialize["name_when_saved"] = o.NameWhenSaved
	}
	if o.Reasoning != nil {
		toSerialize["reasoning"] = o.Reasoning
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsAnnotationLabelValueResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assessment    *LLMObsAnnotationAssessment      `json:"assessment,omitempty"`
		LabelSchemaId *string                          `json:"label_schema_id"`
		NameWhenSaved *string                          `json:"name_when_saved,omitempty"`
		Reasoning     *string                          `json:"reasoning,omitempty"`
		Type          *LLMObsLabelSchemaType           `json:"type,omitempty"`
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
		datadog.DeleteKeys(additionalProperties, &[]string{"assessment", "label_schema_id", "name_when_saved", "reasoning", "type", "value"})
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
	o.NameWhenSaved = all.NameWhenSaved
	o.Reasoning = all.Reasoning
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	o.Value = *all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
