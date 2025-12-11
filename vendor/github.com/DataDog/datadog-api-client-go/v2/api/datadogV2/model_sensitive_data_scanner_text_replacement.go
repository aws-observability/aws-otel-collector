// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SensitiveDataScannerTextReplacement Object describing how the scanned event will be replaced.
type SensitiveDataScannerTextReplacement struct {
	// Required if type == 'partial_replacement_from_beginning'
	// or 'partial_replacement_from_end'. It must be > 0.
	NumberOfChars *int64 `json:"number_of_chars,omitempty"`
	// Required if type == 'replacement_string'.
	ReplacementString *string `json:"replacement_string,omitempty"`
	// Only valid when type == `replacement_string`. When enabled, matches can be unmasked in logs by users with ‘Data Scanner Unmask’ permission. As a security best practice, avoid masking for highly-sensitive, long-lived data.
	ShouldSaveMatch *bool `json:"should_save_match,omitempty"`
	// Type of the replacement text. None means no replacement.
	// hash means the data will be stubbed. replacement_string means that
	// one can chose a text to replace the data. partial_replacement_from_beginning
	// allows a user to partially replace the data from the beginning, and
	// partial_replacement_from_end on the other hand, allows to replace data from
	// the end.
	Type *SensitiveDataScannerTextReplacementType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSensitiveDataScannerTextReplacement instantiates a new SensitiveDataScannerTextReplacement object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSensitiveDataScannerTextReplacement() *SensitiveDataScannerTextReplacement {
	this := SensitiveDataScannerTextReplacement{}
	var typeVar SensitiveDataScannerTextReplacementType = SENSITIVEDATASCANNERTEXTREPLACEMENTTYPE_NONE
	this.Type = &typeVar
	return &this
}

// NewSensitiveDataScannerTextReplacementWithDefaults instantiates a new SensitiveDataScannerTextReplacement object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSensitiveDataScannerTextReplacementWithDefaults() *SensitiveDataScannerTextReplacement {
	this := SensitiveDataScannerTextReplacement{}
	var typeVar SensitiveDataScannerTextReplacementType = SENSITIVEDATASCANNERTEXTREPLACEMENTTYPE_NONE
	this.Type = &typeVar
	return &this
}

// GetNumberOfChars returns the NumberOfChars field value if set, zero value otherwise.
func (o *SensitiveDataScannerTextReplacement) GetNumberOfChars() int64 {
	if o == nil || o.NumberOfChars == nil {
		var ret int64
		return ret
	}
	return *o.NumberOfChars
}

// GetNumberOfCharsOk returns a tuple with the NumberOfChars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTextReplacement) GetNumberOfCharsOk() (*int64, bool) {
	if o == nil || o.NumberOfChars == nil {
		return nil, false
	}
	return o.NumberOfChars, true
}

// HasNumberOfChars returns a boolean if a field has been set.
func (o *SensitiveDataScannerTextReplacement) HasNumberOfChars() bool {
	return o != nil && o.NumberOfChars != nil
}

// SetNumberOfChars gets a reference to the given int64 and assigns it to the NumberOfChars field.
func (o *SensitiveDataScannerTextReplacement) SetNumberOfChars(v int64) {
	o.NumberOfChars = &v
}

// GetReplacementString returns the ReplacementString field value if set, zero value otherwise.
func (o *SensitiveDataScannerTextReplacement) GetReplacementString() string {
	if o == nil || o.ReplacementString == nil {
		var ret string
		return ret
	}
	return *o.ReplacementString
}

// GetReplacementStringOk returns a tuple with the ReplacementString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTextReplacement) GetReplacementStringOk() (*string, bool) {
	if o == nil || o.ReplacementString == nil {
		return nil, false
	}
	return o.ReplacementString, true
}

// HasReplacementString returns a boolean if a field has been set.
func (o *SensitiveDataScannerTextReplacement) HasReplacementString() bool {
	return o != nil && o.ReplacementString != nil
}

// SetReplacementString gets a reference to the given string and assigns it to the ReplacementString field.
func (o *SensitiveDataScannerTextReplacement) SetReplacementString(v string) {
	o.ReplacementString = &v
}

// GetShouldSaveMatch returns the ShouldSaveMatch field value if set, zero value otherwise.
func (o *SensitiveDataScannerTextReplacement) GetShouldSaveMatch() bool {
	if o == nil || o.ShouldSaveMatch == nil {
		var ret bool
		return ret
	}
	return *o.ShouldSaveMatch
}

// GetShouldSaveMatchOk returns a tuple with the ShouldSaveMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTextReplacement) GetShouldSaveMatchOk() (*bool, bool) {
	if o == nil || o.ShouldSaveMatch == nil {
		return nil, false
	}
	return o.ShouldSaveMatch, true
}

// HasShouldSaveMatch returns a boolean if a field has been set.
func (o *SensitiveDataScannerTextReplacement) HasShouldSaveMatch() bool {
	return o != nil && o.ShouldSaveMatch != nil
}

// SetShouldSaveMatch gets a reference to the given bool and assigns it to the ShouldSaveMatch field.
func (o *SensitiveDataScannerTextReplacement) SetShouldSaveMatch(v bool) {
	o.ShouldSaveMatch = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SensitiveDataScannerTextReplacement) GetType() SensitiveDataScannerTextReplacementType {
	if o == nil || o.Type == nil {
		var ret SensitiveDataScannerTextReplacementType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SensitiveDataScannerTextReplacement) GetTypeOk() (*SensitiveDataScannerTextReplacementType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SensitiveDataScannerTextReplacement) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given SensitiveDataScannerTextReplacementType and assigns it to the Type field.
func (o *SensitiveDataScannerTextReplacement) SetType(v SensitiveDataScannerTextReplacementType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SensitiveDataScannerTextReplacement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.NumberOfChars != nil {
		toSerialize["number_of_chars"] = o.NumberOfChars
	}
	if o.ReplacementString != nil {
		toSerialize["replacement_string"] = o.ReplacementString
	}
	if o.ShouldSaveMatch != nil {
		toSerialize["should_save_match"] = o.ShouldSaveMatch
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SensitiveDataScannerTextReplacement) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		NumberOfChars     *int64                                   `json:"number_of_chars,omitempty"`
		ReplacementString *string                                  `json:"replacement_string,omitempty"`
		ShouldSaveMatch   *bool                                    `json:"should_save_match,omitempty"`
		Type              *SensitiveDataScannerTextReplacementType `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"number_of_chars", "replacement_string", "should_save_match", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.NumberOfChars = all.NumberOfChars
	o.ReplacementString = all.ReplacementString
	o.ShouldSaveMatch = all.ShouldSaveMatch
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
