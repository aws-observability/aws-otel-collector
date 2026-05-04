// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestVersionDiffPatchDiff Object describing a single text diff operation.
type SyntheticsTestVersionDiffPatchDiff struct {
	// The text that was changed.
	ChangeText *string `json:"change_text,omitempty"`
	// The diff operation applied.
	Operation *string `json:"operation,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestVersionDiffPatchDiff instantiates a new SyntheticsTestVersionDiffPatchDiff object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestVersionDiffPatchDiff() *SyntheticsTestVersionDiffPatchDiff {
	this := SyntheticsTestVersionDiffPatchDiff{}
	return &this
}

// NewSyntheticsTestVersionDiffPatchDiffWithDefaults instantiates a new SyntheticsTestVersionDiffPatchDiff object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestVersionDiffPatchDiffWithDefaults() *SyntheticsTestVersionDiffPatchDiff {
	this := SyntheticsTestVersionDiffPatchDiff{}
	return &this
}

// GetChangeText returns the ChangeText field value if set, zero value otherwise.
func (o *SyntheticsTestVersionDiffPatchDiff) GetChangeText() string {
	if o == nil || o.ChangeText == nil {
		var ret string
		return ret
	}
	return *o.ChangeText
}

// GetChangeTextOk returns a tuple with the ChangeText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionDiffPatchDiff) GetChangeTextOk() (*string, bool) {
	if o == nil || o.ChangeText == nil {
		return nil, false
	}
	return o.ChangeText, true
}

// HasChangeText returns a boolean if a field has been set.
func (o *SyntheticsTestVersionDiffPatchDiff) HasChangeText() bool {
	return o != nil && o.ChangeText != nil
}

// SetChangeText gets a reference to the given string and assigns it to the ChangeText field.
func (o *SyntheticsTestVersionDiffPatchDiff) SetChangeText(v string) {
	o.ChangeText = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *SyntheticsTestVersionDiffPatchDiff) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionDiffPatchDiff) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *SyntheticsTestVersionDiffPatchDiff) HasOperation() bool {
	return o != nil && o.Operation != nil
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *SyntheticsTestVersionDiffPatchDiff) SetOperation(v string) {
	o.Operation = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestVersionDiffPatchDiff) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ChangeText != nil {
		toSerialize["change_text"] = o.ChangeText
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestVersionDiffPatchDiff) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ChangeText *string `json:"change_text,omitempty"`
		Operation  *string `json:"operation,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"change_text", "operation"})
	} else {
		return err
	}
	o.ChangeText = all.ChangeText
	o.Operation = all.Operation

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
