// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultVariable A variable used or extracted during a test.
type SyntheticsTestResultVariable struct {
	// Error encountered when evaluating the variable.
	Err *string `json:"err,omitempty"`
	// Human-readable error message for variable evaluation.
	ErrorMessage *string `json:"error_message,omitempty"`
	// Example value for the variable.
	Example *string `json:"example,omitempty"`
	// Variable identifier.
	Id *string `json:"id,omitempty"`
	// Variable name.
	Name *string `json:"name,omitempty"`
	// Pattern used to extract the variable.
	Pattern *string `json:"pattern,omitempty"`
	// Whether the variable holds a secure value.
	Secure *bool `json:"secure,omitempty"`
	// Variable type.
	Type *string `json:"type,omitempty"`
	// Evaluated value of the variable.
	Val *string `json:"val,omitempty"`
	// Current value of the variable.
	Value *string `json:"value,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultVariable instantiates a new SyntheticsTestResultVariable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultVariable() *SyntheticsTestResultVariable {
	this := SyntheticsTestResultVariable{}
	return &this
}

// NewSyntheticsTestResultVariableWithDefaults instantiates a new SyntheticsTestResultVariable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultVariableWithDefaults() *SyntheticsTestResultVariable {
	this := SyntheticsTestResultVariable{}
	return &this
}

// GetErr returns the Err field value if set, zero value otherwise.
func (o *SyntheticsTestResultVariable) GetErr() string {
	if o == nil || o.Err == nil {
		var ret string
		return ret
	}
	return *o.Err
}

// GetErrOk returns a tuple with the Err field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVariable) GetErrOk() (*string, bool) {
	if o == nil || o.Err == nil {
		return nil, false
	}
	return o.Err, true
}

// HasErr returns a boolean if a field has been set.
func (o *SyntheticsTestResultVariable) HasErr() bool {
	return o != nil && o.Err != nil
}

// SetErr gets a reference to the given string and assigns it to the Err field.
func (o *SyntheticsTestResultVariable) SetErr(v string) {
	o.Err = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *SyntheticsTestResultVariable) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVariable) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *SyntheticsTestResultVariable) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *SyntheticsTestResultVariable) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *SyntheticsTestResultVariable) GetExample() string {
	if o == nil || o.Example == nil {
		var ret string
		return ret
	}
	return *o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVariable) GetExampleOk() (*string, bool) {
	if o == nil || o.Example == nil {
		return nil, false
	}
	return o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *SyntheticsTestResultVariable) HasExample() bool {
	return o != nil && o.Example != nil
}

// SetExample gets a reference to the given string and assigns it to the Example field.
func (o *SyntheticsTestResultVariable) SetExample(v string) {
	o.Example = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsTestResultVariable) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVariable) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsTestResultVariable) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsTestResultVariable) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTestResultVariable) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVariable) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTestResultVariable) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTestResultVariable) SetName(v string) {
	o.Name = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *SyntheticsTestResultVariable) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVariable) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *SyntheticsTestResultVariable) HasPattern() bool {
	return o != nil && o.Pattern != nil
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *SyntheticsTestResultVariable) SetPattern(v string) {
	o.Pattern = &v
}

// GetSecure returns the Secure field value if set, zero value otherwise.
func (o *SyntheticsTestResultVariable) GetSecure() bool {
	if o == nil || o.Secure == nil {
		var ret bool
		return ret
	}
	return *o.Secure
}

// GetSecureOk returns a tuple with the Secure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVariable) GetSecureOk() (*bool, bool) {
	if o == nil || o.Secure == nil {
		return nil, false
	}
	return o.Secure, true
}

// HasSecure returns a boolean if a field has been set.
func (o *SyntheticsTestResultVariable) HasSecure() bool {
	return o != nil && o.Secure != nil
}

// SetSecure gets a reference to the given bool and assigns it to the Secure field.
func (o *SyntheticsTestResultVariable) SetSecure(v bool) {
	o.Secure = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsTestResultVariable) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVariable) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsTestResultVariable) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SyntheticsTestResultVariable) SetType(v string) {
	o.Type = &v
}

// GetVal returns the Val field value if set, zero value otherwise.
func (o *SyntheticsTestResultVariable) GetVal() string {
	if o == nil || o.Val == nil {
		var ret string
		return ret
	}
	return *o.Val
}

// GetValOk returns a tuple with the Val field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVariable) GetValOk() (*string, bool) {
	if o == nil || o.Val == nil {
		return nil, false
	}
	return o.Val, true
}

// HasVal returns a boolean if a field has been set.
func (o *SyntheticsTestResultVariable) HasVal() bool {
	return o != nil && o.Val != nil
}

// SetVal gets a reference to the given string and assigns it to the Val field.
func (o *SyntheticsTestResultVariable) SetVal(v string) {
	o.Val = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SyntheticsTestResultVariable) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVariable) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SyntheticsTestResultVariable) HasValue() bool {
	return o != nil && o.Value != nil
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SyntheticsTestResultVariable) SetValue(v string) {
	o.Value = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultVariable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Err != nil {
		toSerialize["err"] = o.Err
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.Example != nil {
		toSerialize["example"] = o.Example
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Secure != nil {
		toSerialize["secure"] = o.Secure
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Val != nil {
		toSerialize["val"] = o.Val
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultVariable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Err          *string `json:"err,omitempty"`
		ErrorMessage *string `json:"error_message,omitempty"`
		Example      *string `json:"example,omitempty"`
		Id           *string `json:"id,omitempty"`
		Name         *string `json:"name,omitempty"`
		Pattern      *string `json:"pattern,omitempty"`
		Secure       *bool   `json:"secure,omitempty"`
		Type         *string `json:"type,omitempty"`
		Val          *string `json:"val,omitempty"`
		Value        *string `json:"value,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"err", "error_message", "example", "id", "name", "pattern", "secure", "type", "val", "value"})
	} else {
		return err
	}
	o.Err = all.Err
	o.ErrorMessage = all.ErrorMessage
	o.Example = all.Example
	o.Id = all.Id
	o.Name = all.Name
	o.Pattern = all.Pattern
	o.Secure = all.Secure
	o.Type = all.Type
	o.Val = all.Val
	o.Value = all.Value

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
