// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestVersionActionMetadata Object containing metadata about a change action.
type SyntheticsTestVersionActionMetadata struct {
	// The value of the property after the change.
	AfterValue interface{} `json:"after_value,omitempty"`
	// The value of the property before the change.
	BeforeValue interface{} `json:"before_value,omitempty"`
	// List of diff patches for text changes.
	DiffPatches []SyntheticsTestVersionDiffPatches `json:"diff_patches,omitempty"`
	// The dot-separated path of the property that was changed.
	PropertyPath *string `json:"property_path,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestVersionActionMetadata instantiates a new SyntheticsTestVersionActionMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestVersionActionMetadata() *SyntheticsTestVersionActionMetadata {
	this := SyntheticsTestVersionActionMetadata{}
	return &this
}

// NewSyntheticsTestVersionActionMetadataWithDefaults instantiates a new SyntheticsTestVersionActionMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestVersionActionMetadataWithDefaults() *SyntheticsTestVersionActionMetadata {
	this := SyntheticsTestVersionActionMetadata{}
	return &this
}

// GetAfterValue returns the AfterValue field value if set, zero value otherwise.
func (o *SyntheticsTestVersionActionMetadata) GetAfterValue() interface{} {
	if o == nil || o.AfterValue == nil {
		var ret interface{}
		return ret
	}
	return o.AfterValue
}

// GetAfterValueOk returns a tuple with the AfterValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionActionMetadata) GetAfterValueOk() (*interface{}, bool) {
	if o == nil || o.AfterValue == nil {
		return nil, false
	}
	return &o.AfterValue, true
}

// HasAfterValue returns a boolean if a field has been set.
func (o *SyntheticsTestVersionActionMetadata) HasAfterValue() bool {
	return o != nil && o.AfterValue != nil
}

// SetAfterValue gets a reference to the given interface{} and assigns it to the AfterValue field.
func (o *SyntheticsTestVersionActionMetadata) SetAfterValue(v interface{}) {
	o.AfterValue = v
}

// GetBeforeValue returns the BeforeValue field value if set, zero value otherwise.
func (o *SyntheticsTestVersionActionMetadata) GetBeforeValue() interface{} {
	if o == nil || o.BeforeValue == nil {
		var ret interface{}
		return ret
	}
	return o.BeforeValue
}

// GetBeforeValueOk returns a tuple with the BeforeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionActionMetadata) GetBeforeValueOk() (*interface{}, bool) {
	if o == nil || o.BeforeValue == nil {
		return nil, false
	}
	return &o.BeforeValue, true
}

// HasBeforeValue returns a boolean if a field has been set.
func (o *SyntheticsTestVersionActionMetadata) HasBeforeValue() bool {
	return o != nil && o.BeforeValue != nil
}

// SetBeforeValue gets a reference to the given interface{} and assigns it to the BeforeValue field.
func (o *SyntheticsTestVersionActionMetadata) SetBeforeValue(v interface{}) {
	o.BeforeValue = v
}

// GetDiffPatches returns the DiffPatches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SyntheticsTestVersionActionMetadata) GetDiffPatches() []SyntheticsTestVersionDiffPatches {
	if o == nil {
		var ret []SyntheticsTestVersionDiffPatches
		return ret
	}
	return o.DiffPatches
}

// GetDiffPatchesOk returns a tuple with the DiffPatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *SyntheticsTestVersionActionMetadata) GetDiffPatchesOk() (*[]SyntheticsTestVersionDiffPatches, bool) {
	if o == nil || o.DiffPatches == nil {
		return nil, false
	}
	return &o.DiffPatches, true
}

// HasDiffPatches returns a boolean if a field has been set.
func (o *SyntheticsTestVersionActionMetadata) HasDiffPatches() bool {
	return o != nil && o.DiffPatches != nil
}

// SetDiffPatches gets a reference to the given []SyntheticsTestVersionDiffPatches and assigns it to the DiffPatches field.
func (o *SyntheticsTestVersionActionMetadata) SetDiffPatches(v []SyntheticsTestVersionDiffPatches) {
	o.DiffPatches = v
}

// GetPropertyPath returns the PropertyPath field value if set, zero value otherwise.
func (o *SyntheticsTestVersionActionMetadata) GetPropertyPath() string {
	if o == nil || o.PropertyPath == nil {
		var ret string
		return ret
	}
	return *o.PropertyPath
}

// GetPropertyPathOk returns a tuple with the PropertyPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionActionMetadata) GetPropertyPathOk() (*string, bool) {
	if o == nil || o.PropertyPath == nil {
		return nil, false
	}
	return o.PropertyPath, true
}

// HasPropertyPath returns a boolean if a field has been set.
func (o *SyntheticsTestVersionActionMetadata) HasPropertyPath() bool {
	return o != nil && o.PropertyPath != nil
}

// SetPropertyPath gets a reference to the given string and assigns it to the PropertyPath field.
func (o *SyntheticsTestVersionActionMetadata) SetPropertyPath(v string) {
	o.PropertyPath = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestVersionActionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AfterValue != nil {
		toSerialize["after_value"] = o.AfterValue
	}
	if o.BeforeValue != nil {
		toSerialize["before_value"] = o.BeforeValue
	}
	if o.DiffPatches != nil {
		toSerialize["diff_patches"] = o.DiffPatches
	}
	if o.PropertyPath != nil {
		toSerialize["property_path"] = o.PropertyPath
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestVersionActionMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AfterValue   interface{}                        `json:"after_value,omitempty"`
		BeforeValue  interface{}                        `json:"before_value,omitempty"`
		DiffPatches  []SyntheticsTestVersionDiffPatches `json:"diff_patches,omitempty"`
		PropertyPath *string                            `json:"property_path,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"after_value", "before_value", "diff_patches", "property_path"})
	} else {
		return err
	}
	o.AfterValue = all.AfterValue
	o.BeforeValue = all.BeforeValue
	o.DiffPatches = all.DiffPatches
	o.PropertyPath = all.PropertyPath

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
