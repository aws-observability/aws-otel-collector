// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStepParamsElement Information about the element used for a step.
type SyntheticsMobileStepParamsElement struct {
	// Context of the element.
	Context *string `json:"context,omitempty"`
	// Type of the context that the element is in.
	ContextType *SyntheticsMobileStepParamsElementContextType `json:"contextType,omitempty"`
	// Description of the element.
	ElementDescription *string `json:"elementDescription,omitempty"`
	// Multi-locator to find the element.
	MultiLocator interface{} `json:"multiLocator,omitempty"`
	// Position of the action relative to the element.
	RelativePosition *SyntheticsMobileStepParamsElementRelativePosition `json:"relativePosition,omitempty"`
	// Text content of the element.
	TextContent *string `json:"textContent,omitempty"`
	// User locator to find the element.
	UserLocator *SyntheticsMobileStepParamsElementUserLocator `json:"userLocator,omitempty"`
	// Name of the view of the element.
	ViewName *string `json:"viewName,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileStepParamsElement instantiates a new SyntheticsMobileStepParamsElement object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileStepParamsElement() *SyntheticsMobileStepParamsElement {
	this := SyntheticsMobileStepParamsElement{}
	return &this
}

// NewSyntheticsMobileStepParamsElementWithDefaults instantiates a new SyntheticsMobileStepParamsElement object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileStepParamsElementWithDefaults() *SyntheticsMobileStepParamsElement {
	this := SyntheticsMobileStepParamsElement{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElement) GetContext() string {
	if o == nil || o.Context == nil {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElement) GetContextOk() (*string, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElement) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *SyntheticsMobileStepParamsElement) SetContext(v string) {
	o.Context = &v
}

// GetContextType returns the ContextType field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElement) GetContextType() SyntheticsMobileStepParamsElementContextType {
	if o == nil || o.ContextType == nil {
		var ret SyntheticsMobileStepParamsElementContextType
		return ret
	}
	return *o.ContextType
}

// GetContextTypeOk returns a tuple with the ContextType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElement) GetContextTypeOk() (*SyntheticsMobileStepParamsElementContextType, bool) {
	if o == nil || o.ContextType == nil {
		return nil, false
	}
	return o.ContextType, true
}

// HasContextType returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElement) HasContextType() bool {
	return o != nil && o.ContextType != nil
}

// SetContextType gets a reference to the given SyntheticsMobileStepParamsElementContextType and assigns it to the ContextType field.
func (o *SyntheticsMobileStepParamsElement) SetContextType(v SyntheticsMobileStepParamsElementContextType) {
	o.ContextType = &v
}

// GetElementDescription returns the ElementDescription field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElement) GetElementDescription() string {
	if o == nil || o.ElementDescription == nil {
		var ret string
		return ret
	}
	return *o.ElementDescription
}

// GetElementDescriptionOk returns a tuple with the ElementDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElement) GetElementDescriptionOk() (*string, bool) {
	if o == nil || o.ElementDescription == nil {
		return nil, false
	}
	return o.ElementDescription, true
}

// HasElementDescription returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElement) HasElementDescription() bool {
	return o != nil && o.ElementDescription != nil
}

// SetElementDescription gets a reference to the given string and assigns it to the ElementDescription field.
func (o *SyntheticsMobileStepParamsElement) SetElementDescription(v string) {
	o.ElementDescription = &v
}

// GetMultiLocator returns the MultiLocator field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElement) GetMultiLocator() interface{} {
	if o == nil || o.MultiLocator == nil {
		var ret interface{}
		return ret
	}
	return o.MultiLocator
}

// GetMultiLocatorOk returns a tuple with the MultiLocator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElement) GetMultiLocatorOk() (*interface{}, bool) {
	if o == nil || o.MultiLocator == nil {
		return nil, false
	}
	return &o.MultiLocator, true
}

// HasMultiLocator returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElement) HasMultiLocator() bool {
	return o != nil && o.MultiLocator != nil
}

// SetMultiLocator gets a reference to the given interface{} and assigns it to the MultiLocator field.
func (o *SyntheticsMobileStepParamsElement) SetMultiLocator(v interface{}) {
	o.MultiLocator = v
}

// GetRelativePosition returns the RelativePosition field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElement) GetRelativePosition() SyntheticsMobileStepParamsElementRelativePosition {
	if o == nil || o.RelativePosition == nil {
		var ret SyntheticsMobileStepParamsElementRelativePosition
		return ret
	}
	return *o.RelativePosition
}

// GetRelativePositionOk returns a tuple with the RelativePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElement) GetRelativePositionOk() (*SyntheticsMobileStepParamsElementRelativePosition, bool) {
	if o == nil || o.RelativePosition == nil {
		return nil, false
	}
	return o.RelativePosition, true
}

// HasRelativePosition returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElement) HasRelativePosition() bool {
	return o != nil && o.RelativePosition != nil
}

// SetRelativePosition gets a reference to the given SyntheticsMobileStepParamsElementRelativePosition and assigns it to the RelativePosition field.
func (o *SyntheticsMobileStepParamsElement) SetRelativePosition(v SyntheticsMobileStepParamsElementRelativePosition) {
	o.RelativePosition = &v
}

// GetTextContent returns the TextContent field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElement) GetTextContent() string {
	if o == nil || o.TextContent == nil {
		var ret string
		return ret
	}
	return *o.TextContent
}

// GetTextContentOk returns a tuple with the TextContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElement) GetTextContentOk() (*string, bool) {
	if o == nil || o.TextContent == nil {
		return nil, false
	}
	return o.TextContent, true
}

// HasTextContent returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElement) HasTextContent() bool {
	return o != nil && o.TextContent != nil
}

// SetTextContent gets a reference to the given string and assigns it to the TextContent field.
func (o *SyntheticsMobileStepParamsElement) SetTextContent(v string) {
	o.TextContent = &v
}

// GetUserLocator returns the UserLocator field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElement) GetUserLocator() SyntheticsMobileStepParamsElementUserLocator {
	if o == nil || o.UserLocator == nil {
		var ret SyntheticsMobileStepParamsElementUserLocator
		return ret
	}
	return *o.UserLocator
}

// GetUserLocatorOk returns a tuple with the UserLocator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElement) GetUserLocatorOk() (*SyntheticsMobileStepParamsElementUserLocator, bool) {
	if o == nil || o.UserLocator == nil {
		return nil, false
	}
	return o.UserLocator, true
}

// HasUserLocator returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElement) HasUserLocator() bool {
	return o != nil && o.UserLocator != nil
}

// SetUserLocator gets a reference to the given SyntheticsMobileStepParamsElementUserLocator and assigns it to the UserLocator field.
func (o *SyntheticsMobileStepParamsElement) SetUserLocator(v SyntheticsMobileStepParamsElementUserLocator) {
	o.UserLocator = &v
}

// GetViewName returns the ViewName field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParamsElement) GetViewName() string {
	if o == nil || o.ViewName == nil {
		var ret string
		return ret
	}
	return *o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParamsElement) GetViewNameOk() (*string, bool) {
	if o == nil || o.ViewName == nil {
		return nil, false
	}
	return o.ViewName, true
}

// HasViewName returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParamsElement) HasViewName() bool {
	return o != nil && o.ViewName != nil
}

// SetViewName gets a reference to the given string and assigns it to the ViewName field.
func (o *SyntheticsMobileStepParamsElement) SetViewName(v string) {
	o.ViewName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileStepParamsElement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	if o.ContextType != nil {
		toSerialize["contextType"] = o.ContextType
	}
	if o.ElementDescription != nil {
		toSerialize["elementDescription"] = o.ElementDescription
	}
	if o.MultiLocator != nil {
		toSerialize["multiLocator"] = o.MultiLocator
	}
	if o.RelativePosition != nil {
		toSerialize["relativePosition"] = o.RelativePosition
	}
	if o.TextContent != nil {
		toSerialize["textContent"] = o.TextContent
	}
	if o.UserLocator != nil {
		toSerialize["userLocator"] = o.UserLocator
	}
	if o.ViewName != nil {
		toSerialize["viewName"] = o.ViewName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileStepParamsElement) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Context            *string                                            `json:"context,omitempty"`
		ContextType        *SyntheticsMobileStepParamsElementContextType      `json:"contextType,omitempty"`
		ElementDescription *string                                            `json:"elementDescription,omitempty"`
		MultiLocator       interface{}                                        `json:"multiLocator,omitempty"`
		RelativePosition   *SyntheticsMobileStepParamsElementRelativePosition `json:"relativePosition,omitempty"`
		TextContent        *string                                            `json:"textContent,omitempty"`
		UserLocator        *SyntheticsMobileStepParamsElementUserLocator      `json:"userLocator,omitempty"`
		ViewName           *string                                            `json:"viewName,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"context", "contextType", "elementDescription", "multiLocator", "relativePosition", "textContent", "userLocator", "viewName"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Context = all.Context
	if all.ContextType != nil && !all.ContextType.IsValid() {
		hasInvalidField = true
	} else {
		o.ContextType = all.ContextType
	}
	o.ElementDescription = all.ElementDescription
	o.MultiLocator = all.MultiLocator
	if all.RelativePosition != nil && all.RelativePosition.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RelativePosition = all.RelativePosition
	o.TextContent = all.TextContent
	if all.UserLocator != nil && all.UserLocator.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.UserLocator = all.UserLocator
	o.ViewName = all.ViewName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
