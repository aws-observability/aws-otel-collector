// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsMobileStepParams The parameters of a mobile step.
type SyntheticsMobileStepParams struct {
	// Type of assertion to apply in an API test.
	Check *SyntheticsCheckType `json:"check,omitempty"`
	// Number of milliseconds to wait between inputs in a `typeText` step type.
	Delay *int64 `json:"delay,omitempty"`
	// The direction of the scroll for a `scrollToElement` step type.
	Direction *SyntheticsMobileStepParamsDirection `json:"direction,omitempty"`
	// Information about the element used for a step.
	Element *SyntheticsMobileStepParamsElement `json:"element,omitempty"`
	// Boolean to change the state of the wifi for a `toggleWiFi` step type.
	Enabled *bool `json:"enabled,omitempty"`
	// Maximum number of scrolls to do for a `scrollToElement` step type.
	MaxScrolls *int64 `json:"maxScrolls,omitempty"`
	// List of positions for the `flick` step type. The maximum is 10 flicks per step
	Positions []SyntheticsMobileStepParamsPositionsItems `json:"positions,omitempty"`
	// Public ID of the test to be played as part of a `playSubTest` step type.
	SubtestPublicId *string `json:"subtestPublicId,omitempty"`
	// Values used in the step for in multiple step types.
	Value *SyntheticsMobileStepParamsValue `json:"value,omitempty"`
	// Variable object for `extractVariable` step type.
	Variable *SyntheticsMobileStepParamsVariable `json:"variable,omitempty"`
	// Boolean to indicate if `Enter` should be pressed at the end of the `typeText` step type.
	WithEnter *bool `json:"withEnter,omitempty"`
	// Amount to scroll by on the `x` axis for a `scroll` step type.
	X *float64 `json:"x,omitempty"`
	// Amount to scroll by on the `y` axis for a `scroll` step type.
	Y *float64 `json:"y,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsMobileStepParams instantiates a new SyntheticsMobileStepParams object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsMobileStepParams() *SyntheticsMobileStepParams {
	this := SyntheticsMobileStepParams{}
	return &this
}

// NewSyntheticsMobileStepParamsWithDefaults instantiates a new SyntheticsMobileStepParams object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsMobileStepParamsWithDefaults() *SyntheticsMobileStepParams {
	this := SyntheticsMobileStepParams{}
	return &this
}

// GetCheck returns the Check field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetCheck() SyntheticsCheckType {
	if o == nil || o.Check == nil {
		var ret SyntheticsCheckType
		return ret
	}
	return *o.Check
}

// GetCheckOk returns a tuple with the Check field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetCheckOk() (*SyntheticsCheckType, bool) {
	if o == nil || o.Check == nil {
		return nil, false
	}
	return o.Check, true
}

// HasCheck returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasCheck() bool {
	return o != nil && o.Check != nil
}

// SetCheck gets a reference to the given SyntheticsCheckType and assigns it to the Check field.
func (o *SyntheticsMobileStepParams) SetCheck(v SyntheticsCheckType) {
	o.Check = &v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetDelay() int64 {
	if o == nil || o.Delay == nil {
		var ret int64
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetDelayOk() (*int64, bool) {
	if o == nil || o.Delay == nil {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasDelay() bool {
	return o != nil && o.Delay != nil
}

// SetDelay gets a reference to the given int64 and assigns it to the Delay field.
func (o *SyntheticsMobileStepParams) SetDelay(v int64) {
	o.Delay = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetDirection() SyntheticsMobileStepParamsDirection {
	if o == nil || o.Direction == nil {
		var ret SyntheticsMobileStepParamsDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetDirectionOk() (*SyntheticsMobileStepParamsDirection, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasDirection() bool {
	return o != nil && o.Direction != nil
}

// SetDirection gets a reference to the given SyntheticsMobileStepParamsDirection and assigns it to the Direction field.
func (o *SyntheticsMobileStepParams) SetDirection(v SyntheticsMobileStepParamsDirection) {
	o.Direction = &v
}

// GetElement returns the Element field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetElement() SyntheticsMobileStepParamsElement {
	if o == nil || o.Element == nil {
		var ret SyntheticsMobileStepParamsElement
		return ret
	}
	return *o.Element
}

// GetElementOk returns a tuple with the Element field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetElementOk() (*SyntheticsMobileStepParamsElement, bool) {
	if o == nil || o.Element == nil {
		return nil, false
	}
	return o.Element, true
}

// HasElement returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasElement() bool {
	return o != nil && o.Element != nil
}

// SetElement gets a reference to the given SyntheticsMobileStepParamsElement and assigns it to the Element field.
func (o *SyntheticsMobileStepParams) SetElement(v SyntheticsMobileStepParamsElement) {
	o.Element = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SyntheticsMobileStepParams) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMaxScrolls returns the MaxScrolls field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetMaxScrolls() int64 {
	if o == nil || o.MaxScrolls == nil {
		var ret int64
		return ret
	}
	return *o.MaxScrolls
}

// GetMaxScrollsOk returns a tuple with the MaxScrolls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetMaxScrollsOk() (*int64, bool) {
	if o == nil || o.MaxScrolls == nil {
		return nil, false
	}
	return o.MaxScrolls, true
}

// HasMaxScrolls returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasMaxScrolls() bool {
	return o != nil && o.MaxScrolls != nil
}

// SetMaxScrolls gets a reference to the given int64 and assigns it to the MaxScrolls field.
func (o *SyntheticsMobileStepParams) SetMaxScrolls(v int64) {
	o.MaxScrolls = &v
}

// GetPositions returns the Positions field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetPositions() []SyntheticsMobileStepParamsPositionsItems {
	if o == nil || o.Positions == nil {
		var ret []SyntheticsMobileStepParamsPositionsItems
		return ret
	}
	return o.Positions
}

// GetPositionsOk returns a tuple with the Positions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetPositionsOk() (*[]SyntheticsMobileStepParamsPositionsItems, bool) {
	if o == nil || o.Positions == nil {
		return nil, false
	}
	return &o.Positions, true
}

// HasPositions returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasPositions() bool {
	return o != nil && o.Positions != nil
}

// SetPositions gets a reference to the given []SyntheticsMobileStepParamsPositionsItems and assigns it to the Positions field.
func (o *SyntheticsMobileStepParams) SetPositions(v []SyntheticsMobileStepParamsPositionsItems) {
	o.Positions = v
}

// GetSubtestPublicId returns the SubtestPublicId field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetSubtestPublicId() string {
	if o == nil || o.SubtestPublicId == nil {
		var ret string
		return ret
	}
	return *o.SubtestPublicId
}

// GetSubtestPublicIdOk returns a tuple with the SubtestPublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetSubtestPublicIdOk() (*string, bool) {
	if o == nil || o.SubtestPublicId == nil {
		return nil, false
	}
	return o.SubtestPublicId, true
}

// HasSubtestPublicId returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasSubtestPublicId() bool {
	return o != nil && o.SubtestPublicId != nil
}

// SetSubtestPublicId gets a reference to the given string and assigns it to the SubtestPublicId field.
func (o *SyntheticsMobileStepParams) SetSubtestPublicId(v string) {
	o.SubtestPublicId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetValue() SyntheticsMobileStepParamsValue {
	if o == nil || o.Value == nil {
		var ret SyntheticsMobileStepParamsValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetValueOk() (*SyntheticsMobileStepParamsValue, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given SyntheticsMobileStepParamsValue and assigns it to the Value field.
func (o *SyntheticsMobileStepParams) SetValue(v SyntheticsMobileStepParamsValue) {
	o.Value = &v
}

// GetVariable returns the Variable field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetVariable() SyntheticsMobileStepParamsVariable {
	if o == nil || o.Variable == nil {
		var ret SyntheticsMobileStepParamsVariable
		return ret
	}
	return *o.Variable
}

// GetVariableOk returns a tuple with the Variable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetVariableOk() (*SyntheticsMobileStepParamsVariable, bool) {
	if o == nil || o.Variable == nil {
		return nil, false
	}
	return o.Variable, true
}

// HasVariable returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasVariable() bool {
	return o != nil && o.Variable != nil
}

// SetVariable gets a reference to the given SyntheticsMobileStepParamsVariable and assigns it to the Variable field.
func (o *SyntheticsMobileStepParams) SetVariable(v SyntheticsMobileStepParamsVariable) {
	o.Variable = &v
}

// GetWithEnter returns the WithEnter field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetWithEnter() bool {
	if o == nil || o.WithEnter == nil {
		var ret bool
		return ret
	}
	return *o.WithEnter
}

// GetWithEnterOk returns a tuple with the WithEnter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetWithEnterOk() (*bool, bool) {
	if o == nil || o.WithEnter == nil {
		return nil, false
	}
	return o.WithEnter, true
}

// HasWithEnter returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasWithEnter() bool {
	return o != nil && o.WithEnter != nil
}

// SetWithEnter gets a reference to the given bool and assigns it to the WithEnter field.
func (o *SyntheticsMobileStepParams) SetWithEnter(v bool) {
	o.WithEnter = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetX() float64 {
	if o == nil || o.X == nil {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetXOk() (*float64, bool) {
	if o == nil || o.X == nil {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasX() bool {
	return o != nil && o.X != nil
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *SyntheticsMobileStepParams) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *SyntheticsMobileStepParams) GetY() float64 {
	if o == nil || o.Y == nil {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsMobileStepParams) GetYOk() (*float64, bool) {
	if o == nil || o.Y == nil {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *SyntheticsMobileStepParams) HasY() bool {
	return o != nil && o.Y != nil
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *SyntheticsMobileStepParams) SetY(v float64) {
	o.Y = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsMobileStepParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Check != nil {
		toSerialize["check"] = o.Check
	}
	if o.Delay != nil {
		toSerialize["delay"] = o.Delay
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.Element != nil {
		toSerialize["element"] = o.Element
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.MaxScrolls != nil {
		toSerialize["maxScrolls"] = o.MaxScrolls
	}
	if o.Positions != nil {
		toSerialize["positions"] = o.Positions
	}
	if o.SubtestPublicId != nil {
		toSerialize["subtestPublicId"] = o.SubtestPublicId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Variable != nil {
		toSerialize["variable"] = o.Variable
	}
	if o.WithEnter != nil {
		toSerialize["withEnter"] = o.WithEnter
	}
	if o.X != nil {
		toSerialize["x"] = o.X
	}
	if o.Y != nil {
		toSerialize["y"] = o.Y
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsMobileStepParams) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Check           *SyntheticsCheckType                       `json:"check,omitempty"`
		Delay           *int64                                     `json:"delay,omitempty"`
		Direction       *SyntheticsMobileStepParamsDirection       `json:"direction,omitempty"`
		Element         *SyntheticsMobileStepParamsElement         `json:"element,omitempty"`
		Enabled         *bool                                      `json:"enabled,omitempty"`
		MaxScrolls      *int64                                     `json:"maxScrolls,omitempty"`
		Positions       []SyntheticsMobileStepParamsPositionsItems `json:"positions,omitempty"`
		SubtestPublicId *string                                    `json:"subtestPublicId,omitempty"`
		Value           *SyntheticsMobileStepParamsValue           `json:"value,omitempty"`
		Variable        *SyntheticsMobileStepParamsVariable        `json:"variable,omitempty"`
		WithEnter       *bool                                      `json:"withEnter,omitempty"`
		X               *float64                                   `json:"x,omitempty"`
		Y               *float64                                   `json:"y,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"check", "delay", "direction", "element", "enabled", "maxScrolls", "positions", "subtestPublicId", "value", "variable", "withEnter", "x", "y"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Check != nil && !all.Check.IsValid() {
		hasInvalidField = true
	} else {
		o.Check = all.Check
	}
	o.Delay = all.Delay
	if all.Direction != nil && !all.Direction.IsValid() {
		hasInvalidField = true
	} else {
		o.Direction = all.Direction
	}
	if all.Element != nil && all.Element.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Element = all.Element
	o.Enabled = all.Enabled
	o.MaxScrolls = all.MaxScrolls
	o.Positions = all.Positions
	o.SubtestPublicId = all.SubtestPublicId
	o.Value = all.Value
	if all.Variable != nil && all.Variable.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Variable = all.Variable
	o.WithEnter = all.WithEnter
	o.X = all.X
	o.Y = all.Y

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
