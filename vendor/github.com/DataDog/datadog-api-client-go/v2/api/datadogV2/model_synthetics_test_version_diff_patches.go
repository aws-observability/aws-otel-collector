// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestVersionDiffPatches Object describing a patch in the diff.
type SyntheticsTestVersionDiffPatches struct {
	// List of individual diff operations.
	Diffs []SyntheticsTestVersionDiffPatchDiff `json:"diffs,omitempty"`
	// Length of the original text segment.
	Length1 *int64 `json:"length1,omitempty"`
	// Length of the modified text segment.
	Length2 *int64 `json:"length2,omitempty"`
	// Start position in the original text.
	Start1 *int64 `json:"start1,omitempty"`
	// Start position in the modified text.
	Start2 *int64 `json:"start2,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestVersionDiffPatches instantiates a new SyntheticsTestVersionDiffPatches object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestVersionDiffPatches() *SyntheticsTestVersionDiffPatches {
	this := SyntheticsTestVersionDiffPatches{}
	return &this
}

// NewSyntheticsTestVersionDiffPatchesWithDefaults instantiates a new SyntheticsTestVersionDiffPatches object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestVersionDiffPatchesWithDefaults() *SyntheticsTestVersionDiffPatches {
	this := SyntheticsTestVersionDiffPatches{}
	return &this
}

// GetDiffs returns the Diffs field value if set, zero value otherwise.
func (o *SyntheticsTestVersionDiffPatches) GetDiffs() []SyntheticsTestVersionDiffPatchDiff {
	if o == nil || o.Diffs == nil {
		var ret []SyntheticsTestVersionDiffPatchDiff
		return ret
	}
	return o.Diffs
}

// GetDiffsOk returns a tuple with the Diffs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionDiffPatches) GetDiffsOk() (*[]SyntheticsTestVersionDiffPatchDiff, bool) {
	if o == nil || o.Diffs == nil {
		return nil, false
	}
	return &o.Diffs, true
}

// HasDiffs returns a boolean if a field has been set.
func (o *SyntheticsTestVersionDiffPatches) HasDiffs() bool {
	return o != nil && o.Diffs != nil
}

// SetDiffs gets a reference to the given []SyntheticsTestVersionDiffPatchDiff and assigns it to the Diffs field.
func (o *SyntheticsTestVersionDiffPatches) SetDiffs(v []SyntheticsTestVersionDiffPatchDiff) {
	o.Diffs = v
}

// GetLength1 returns the Length1 field value if set, zero value otherwise.
func (o *SyntheticsTestVersionDiffPatches) GetLength1() int64 {
	if o == nil || o.Length1 == nil {
		var ret int64
		return ret
	}
	return *o.Length1
}

// GetLength1Ok returns a tuple with the Length1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionDiffPatches) GetLength1Ok() (*int64, bool) {
	if o == nil || o.Length1 == nil {
		return nil, false
	}
	return o.Length1, true
}

// HasLength1 returns a boolean if a field has been set.
func (o *SyntheticsTestVersionDiffPatches) HasLength1() bool {
	return o != nil && o.Length1 != nil
}

// SetLength1 gets a reference to the given int64 and assigns it to the Length1 field.
func (o *SyntheticsTestVersionDiffPatches) SetLength1(v int64) {
	o.Length1 = &v
}

// GetLength2 returns the Length2 field value if set, zero value otherwise.
func (o *SyntheticsTestVersionDiffPatches) GetLength2() int64 {
	if o == nil || o.Length2 == nil {
		var ret int64
		return ret
	}
	return *o.Length2
}

// GetLength2Ok returns a tuple with the Length2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionDiffPatches) GetLength2Ok() (*int64, bool) {
	if o == nil || o.Length2 == nil {
		return nil, false
	}
	return o.Length2, true
}

// HasLength2 returns a boolean if a field has been set.
func (o *SyntheticsTestVersionDiffPatches) HasLength2() bool {
	return o != nil && o.Length2 != nil
}

// SetLength2 gets a reference to the given int64 and assigns it to the Length2 field.
func (o *SyntheticsTestVersionDiffPatches) SetLength2(v int64) {
	o.Length2 = &v
}

// GetStart1 returns the Start1 field value if set, zero value otherwise.
func (o *SyntheticsTestVersionDiffPatches) GetStart1() int64 {
	if o == nil || o.Start1 == nil {
		var ret int64
		return ret
	}
	return *o.Start1
}

// GetStart1Ok returns a tuple with the Start1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionDiffPatches) GetStart1Ok() (*int64, bool) {
	if o == nil || o.Start1 == nil {
		return nil, false
	}
	return o.Start1, true
}

// HasStart1 returns a boolean if a field has been set.
func (o *SyntheticsTestVersionDiffPatches) HasStart1() bool {
	return o != nil && o.Start1 != nil
}

// SetStart1 gets a reference to the given int64 and assigns it to the Start1 field.
func (o *SyntheticsTestVersionDiffPatches) SetStart1(v int64) {
	o.Start1 = &v
}

// GetStart2 returns the Start2 field value if set, zero value otherwise.
func (o *SyntheticsTestVersionDiffPatches) GetStart2() int64 {
	if o == nil || o.Start2 == nil {
		var ret int64
		return ret
	}
	return *o.Start2
}

// GetStart2Ok returns a tuple with the Start2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionDiffPatches) GetStart2Ok() (*int64, bool) {
	if o == nil || o.Start2 == nil {
		return nil, false
	}
	return o.Start2, true
}

// HasStart2 returns a boolean if a field has been set.
func (o *SyntheticsTestVersionDiffPatches) HasStart2() bool {
	return o != nil && o.Start2 != nil
}

// SetStart2 gets a reference to the given int64 and assigns it to the Start2 field.
func (o *SyntheticsTestVersionDiffPatches) SetStart2(v int64) {
	o.Start2 = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestVersionDiffPatches) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Diffs != nil {
		toSerialize["diffs"] = o.Diffs
	}
	if o.Length1 != nil {
		toSerialize["length1"] = o.Length1
	}
	if o.Length2 != nil {
		toSerialize["length2"] = o.Length2
	}
	if o.Start1 != nil {
		toSerialize["start1"] = o.Start1
	}
	if o.Start2 != nil {
		toSerialize["start2"] = o.Start2
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestVersionDiffPatches) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Diffs   []SyntheticsTestVersionDiffPatchDiff `json:"diffs,omitempty"`
		Length1 *int64                               `json:"length1,omitempty"`
		Length2 *int64                               `json:"length2,omitempty"`
		Start1  *int64                               `json:"start1,omitempty"`
		Start2  *int64                               `json:"start2,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"diffs", "length1", "length2", "start1", "start2"})
	} else {
		return err
	}
	o.Diffs = all.Diffs
	o.Length1 = all.Length1
	o.Length2 = all.Length2
	o.Start1 = all.Start1
	o.Start2 = all.Start2

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
