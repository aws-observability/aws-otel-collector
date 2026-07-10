// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsLabelSchema Schema definition for a single label in an annotation queue.
type LLMObsLabelSchema struct {
	// Description of the label.
	Description *string `json:"description,omitempty"`
	// Whether this label includes an assessment field.
	HasAssessment *bool `json:"has_assessment,omitempty"`
	// Whether this label includes a reasoning field.
	HasReasoning *bool `json:"has_reasoning,omitempty"`
	// Unique identifier of the label schema. Assigned by the server if not provided.
	Id *string `json:"id,omitempty"`
	// Whether the boolean label represents an assessment. Requires `has_assessment` to be true.
	IsAssessment *bool `json:"is_assessment,omitempty"`
	// Whether score values must be integers. Applicable to score-type labels.
	IsInteger *bool `json:"is_integer,omitempty"`
	// Whether this label is required for an annotation.
	IsRequired *bool `json:"is_required,omitempty"`
	// Maximum value for score-type labels.
	Max *float64 `json:"max,omitempty"`
	// Minimum value for score-type labels.
	Min *float64 `json:"min,omitempty"`
	// Name of the label. Must match the pattern `^[a-zA-Z0-9_-]+$` and be unique within the queue.
	Name string `json:"name"`
	// Type of a label in an annotation queue label schema.
	Type LLMObsLabelSchemaType `json:"type"`
	// Allowed values for categorical-type labels. Must contain at least one non-empty, unique value.
	Values []string `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsLabelSchema instantiates a new LLMObsLabelSchema object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsLabelSchema(name string, typeVar LLMObsLabelSchemaType) *LLMObsLabelSchema {
	this := LLMObsLabelSchema{}
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewLLMObsLabelSchemaWithDefaults instantiates a new LLMObsLabelSchema object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsLabelSchemaWithDefaults() *LLMObsLabelSchema {
	this := LLMObsLabelSchema{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *LLMObsLabelSchema) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsLabelSchema) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *LLMObsLabelSchema) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *LLMObsLabelSchema) SetDescription(v string) {
	o.Description = &v
}

// GetHasAssessment returns the HasAssessment field value if set, zero value otherwise.
func (o *LLMObsLabelSchema) GetHasAssessment() bool {
	if o == nil || o.HasAssessment == nil {
		var ret bool
		return ret
	}
	return *o.HasAssessment
}

// GetHasAssessmentOk returns a tuple with the HasAssessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsLabelSchema) GetHasAssessmentOk() (*bool, bool) {
	if o == nil || o.HasAssessment == nil {
		return nil, false
	}
	return o.HasAssessment, true
}

// HasHasAssessment returns a boolean if a field has been set.
func (o *LLMObsLabelSchema) HasHasAssessment() bool {
	return o != nil && o.HasAssessment != nil
}

// SetHasAssessment gets a reference to the given bool and assigns it to the HasAssessment field.
func (o *LLMObsLabelSchema) SetHasAssessment(v bool) {
	o.HasAssessment = &v
}

// GetHasReasoning returns the HasReasoning field value if set, zero value otherwise.
func (o *LLMObsLabelSchema) GetHasReasoning() bool {
	if o == nil || o.HasReasoning == nil {
		var ret bool
		return ret
	}
	return *o.HasReasoning
}

// GetHasReasoningOk returns a tuple with the HasReasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsLabelSchema) GetHasReasoningOk() (*bool, bool) {
	if o == nil || o.HasReasoning == nil {
		return nil, false
	}
	return o.HasReasoning, true
}

// HasHasReasoning returns a boolean if a field has been set.
func (o *LLMObsLabelSchema) HasHasReasoning() bool {
	return o != nil && o.HasReasoning != nil
}

// SetHasReasoning gets a reference to the given bool and assigns it to the HasReasoning field.
func (o *LLMObsLabelSchema) SetHasReasoning(v bool) {
	o.HasReasoning = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LLMObsLabelSchema) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsLabelSchema) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LLMObsLabelSchema) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LLMObsLabelSchema) SetId(v string) {
	o.Id = &v
}

// GetIsAssessment returns the IsAssessment field value if set, zero value otherwise.
func (o *LLMObsLabelSchema) GetIsAssessment() bool {
	if o == nil || o.IsAssessment == nil {
		var ret bool
		return ret
	}
	return *o.IsAssessment
}

// GetIsAssessmentOk returns a tuple with the IsAssessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsLabelSchema) GetIsAssessmentOk() (*bool, bool) {
	if o == nil || o.IsAssessment == nil {
		return nil, false
	}
	return o.IsAssessment, true
}

// HasIsAssessment returns a boolean if a field has been set.
func (o *LLMObsLabelSchema) HasIsAssessment() bool {
	return o != nil && o.IsAssessment != nil
}

// SetIsAssessment gets a reference to the given bool and assigns it to the IsAssessment field.
func (o *LLMObsLabelSchema) SetIsAssessment(v bool) {
	o.IsAssessment = &v
}

// GetIsInteger returns the IsInteger field value if set, zero value otherwise.
func (o *LLMObsLabelSchema) GetIsInteger() bool {
	if o == nil || o.IsInteger == nil {
		var ret bool
		return ret
	}
	return *o.IsInteger
}

// GetIsIntegerOk returns a tuple with the IsInteger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsLabelSchema) GetIsIntegerOk() (*bool, bool) {
	if o == nil || o.IsInteger == nil {
		return nil, false
	}
	return o.IsInteger, true
}

// HasIsInteger returns a boolean if a field has been set.
func (o *LLMObsLabelSchema) HasIsInteger() bool {
	return o != nil && o.IsInteger != nil
}

// SetIsInteger gets a reference to the given bool and assigns it to the IsInteger field.
func (o *LLMObsLabelSchema) SetIsInteger(v bool) {
	o.IsInteger = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *LLMObsLabelSchema) GetIsRequired() bool {
	if o == nil || o.IsRequired == nil {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsLabelSchema) GetIsRequiredOk() (*bool, bool) {
	if o == nil || o.IsRequired == nil {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *LLMObsLabelSchema) HasIsRequired() bool {
	return o != nil && o.IsRequired != nil
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *LLMObsLabelSchema) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *LLMObsLabelSchema) GetMax() float64 {
	if o == nil || o.Max == nil {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsLabelSchema) GetMaxOk() (*float64, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *LLMObsLabelSchema) HasMax() bool {
	return o != nil && o.Max != nil
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *LLMObsLabelSchema) SetMax(v float64) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *LLMObsLabelSchema) GetMin() float64 {
	if o == nil || o.Min == nil {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsLabelSchema) GetMinOk() (*float64, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *LLMObsLabelSchema) HasMin() bool {
	return o != nil && o.Min != nil
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *LLMObsLabelSchema) SetMin(v float64) {
	o.Min = &v
}

// GetName returns the Name field value.
func (o *LLMObsLabelSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LLMObsLabelSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *LLMObsLabelSchema) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value.
func (o *LLMObsLabelSchema) GetType() LLMObsLabelSchemaType {
	if o == nil {
		var ret LLMObsLabelSchemaType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LLMObsLabelSchema) GetTypeOk() (*LLMObsLabelSchemaType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LLMObsLabelSchema) SetType(v LLMObsLabelSchemaType) {
	o.Type = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *LLMObsLabelSchema) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LLMObsLabelSchema) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *LLMObsLabelSchema) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *LLMObsLabelSchema) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsLabelSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.HasAssessment != nil {
		toSerialize["has_assessment"] = o.HasAssessment
	}
	if o.HasReasoning != nil {
		toSerialize["has_reasoning"] = o.HasReasoning
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsAssessment != nil {
		toSerialize["is_assessment"] = o.IsAssessment
	}
	if o.IsInteger != nil {
		toSerialize["is_integer"] = o.IsInteger
	}
	if o.IsRequired != nil {
		toSerialize["is_required"] = o.IsRequired
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsLabelSchema) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description   *string                `json:"description,omitempty"`
		HasAssessment *bool                  `json:"has_assessment,omitempty"`
		HasReasoning  *bool                  `json:"has_reasoning,omitempty"`
		Id            *string                `json:"id,omitempty"`
		IsAssessment  *bool                  `json:"is_assessment,omitempty"`
		IsInteger     *bool                  `json:"is_integer,omitempty"`
		IsRequired    *bool                  `json:"is_required,omitempty"`
		Max           *float64               `json:"max,omitempty"`
		Min           *float64               `json:"min,omitempty"`
		Name          *string                `json:"name"`
		Type          *LLMObsLabelSchemaType `json:"type"`
		Values        []string               `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "has_assessment", "has_reasoning", "id", "is_assessment", "is_integer", "is_required", "max", "min", "name", "type", "values"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.HasAssessment = all.HasAssessment
	o.HasReasoning = all.HasReasoning
	o.Id = all.Id
	o.IsAssessment = all.IsAssessment
	o.IsInteger = all.IsInteger
	o.IsRequired = all.IsRequired
	o.Max = all.Max
	o.Min = all.Min
	o.Name = *all.Name
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
