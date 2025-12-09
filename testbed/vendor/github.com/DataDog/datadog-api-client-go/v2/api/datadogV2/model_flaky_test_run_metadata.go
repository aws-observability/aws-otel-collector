// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FlakyTestRunMetadata Metadata about the latest failed test run of the flaky test.
type FlakyTestRunMetadata struct {
	// The duration of the test run in milliseconds.
	DurationMs datadog.NullableInt64 `json:"duration_ms,omitempty"`
	// The error message from the test failure.
	ErrorMessage datadog.NullableString `json:"error_message,omitempty"`
	// The stack trace from the test failure.
	ErrorStack datadog.NullableString `json:"error_stack,omitempty"`
	// The line number where the test ends in the source file.
	SourceEnd datadog.NullableInt64 `json:"source_end,omitempty"`
	// The source file where the test is defined.
	SourceFile datadog.NullableString `json:"source_file,omitempty"`
	// The line number where the test starts in the source file.
	SourceStart datadog.NullableInt64 `json:"source_start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFlakyTestRunMetadata instantiates a new FlakyTestRunMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFlakyTestRunMetadata() *FlakyTestRunMetadata {
	this := FlakyTestRunMetadata{}
	return &this
}

// NewFlakyTestRunMetadataWithDefaults instantiates a new FlakyTestRunMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFlakyTestRunMetadataWithDefaults() *FlakyTestRunMetadata {
	this := FlakyTestRunMetadata{}
	return &this
}

// GetDurationMs returns the DurationMs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestRunMetadata) GetDurationMs() int64 {
	if o == nil || o.DurationMs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.DurationMs.Get()
}

// GetDurationMsOk returns a tuple with the DurationMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestRunMetadata) GetDurationMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DurationMs.Get(), o.DurationMs.IsSet()
}

// HasDurationMs returns a boolean if a field has been set.
func (o *FlakyTestRunMetadata) HasDurationMs() bool {
	return o != nil && o.DurationMs.IsSet()
}

// SetDurationMs gets a reference to the given datadog.NullableInt64 and assigns it to the DurationMs field.
func (o *FlakyTestRunMetadata) SetDurationMs(v int64) {
	o.DurationMs.Set(&v)
}

// SetDurationMsNil sets the value for DurationMs to be an explicit nil.
func (o *FlakyTestRunMetadata) SetDurationMsNil() {
	o.DurationMs.Set(nil)
}

// UnsetDurationMs ensures that no value is present for DurationMs, not even an explicit nil.
func (o *FlakyTestRunMetadata) UnsetDurationMs() {
	o.DurationMs.Unset()
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestRunMetadata) GetErrorMessage() string {
	if o == nil || o.ErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage.Get()
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestRunMetadata) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorMessage.Get(), o.ErrorMessage.IsSet()
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *FlakyTestRunMetadata) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage.IsSet()
}

// SetErrorMessage gets a reference to the given datadog.NullableString and assigns it to the ErrorMessage field.
func (o *FlakyTestRunMetadata) SetErrorMessage(v string) {
	o.ErrorMessage.Set(&v)
}

// SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil.
func (o *FlakyTestRunMetadata) SetErrorMessageNil() {
	o.ErrorMessage.Set(nil)
}

// UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil.
func (o *FlakyTestRunMetadata) UnsetErrorMessage() {
	o.ErrorMessage.Unset()
}

// GetErrorStack returns the ErrorStack field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestRunMetadata) GetErrorStack() string {
	if o == nil || o.ErrorStack.Get() == nil {
		var ret string
		return ret
	}
	return *o.ErrorStack.Get()
}

// GetErrorStackOk returns a tuple with the ErrorStack field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestRunMetadata) GetErrorStackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorStack.Get(), o.ErrorStack.IsSet()
}

// HasErrorStack returns a boolean if a field has been set.
func (o *FlakyTestRunMetadata) HasErrorStack() bool {
	return o != nil && o.ErrorStack.IsSet()
}

// SetErrorStack gets a reference to the given datadog.NullableString and assigns it to the ErrorStack field.
func (o *FlakyTestRunMetadata) SetErrorStack(v string) {
	o.ErrorStack.Set(&v)
}

// SetErrorStackNil sets the value for ErrorStack to be an explicit nil.
func (o *FlakyTestRunMetadata) SetErrorStackNil() {
	o.ErrorStack.Set(nil)
}

// UnsetErrorStack ensures that no value is present for ErrorStack, not even an explicit nil.
func (o *FlakyTestRunMetadata) UnsetErrorStack() {
	o.ErrorStack.Unset()
}

// GetSourceEnd returns the SourceEnd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestRunMetadata) GetSourceEnd() int64 {
	if o == nil || o.SourceEnd.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SourceEnd.Get()
}

// GetSourceEndOk returns a tuple with the SourceEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestRunMetadata) GetSourceEndOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceEnd.Get(), o.SourceEnd.IsSet()
}

// HasSourceEnd returns a boolean if a field has been set.
func (o *FlakyTestRunMetadata) HasSourceEnd() bool {
	return o != nil && o.SourceEnd.IsSet()
}

// SetSourceEnd gets a reference to the given datadog.NullableInt64 and assigns it to the SourceEnd field.
func (o *FlakyTestRunMetadata) SetSourceEnd(v int64) {
	o.SourceEnd.Set(&v)
}

// SetSourceEndNil sets the value for SourceEnd to be an explicit nil.
func (o *FlakyTestRunMetadata) SetSourceEndNil() {
	o.SourceEnd.Set(nil)
}

// UnsetSourceEnd ensures that no value is present for SourceEnd, not even an explicit nil.
func (o *FlakyTestRunMetadata) UnsetSourceEnd() {
	o.SourceEnd.Unset()
}

// GetSourceFile returns the SourceFile field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestRunMetadata) GetSourceFile() string {
	if o == nil || o.SourceFile.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceFile.Get()
}

// GetSourceFileOk returns a tuple with the SourceFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestRunMetadata) GetSourceFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceFile.Get(), o.SourceFile.IsSet()
}

// HasSourceFile returns a boolean if a field has been set.
func (o *FlakyTestRunMetadata) HasSourceFile() bool {
	return o != nil && o.SourceFile.IsSet()
}

// SetSourceFile gets a reference to the given datadog.NullableString and assigns it to the SourceFile field.
func (o *FlakyTestRunMetadata) SetSourceFile(v string) {
	o.SourceFile.Set(&v)
}

// SetSourceFileNil sets the value for SourceFile to be an explicit nil.
func (o *FlakyTestRunMetadata) SetSourceFileNil() {
	o.SourceFile.Set(nil)
}

// UnsetSourceFile ensures that no value is present for SourceFile, not even an explicit nil.
func (o *FlakyTestRunMetadata) UnsetSourceFile() {
	o.SourceFile.Unset()
}

// GetSourceStart returns the SourceStart field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlakyTestRunMetadata) GetSourceStart() int64 {
	if o == nil || o.SourceStart.Get() == nil {
		var ret int64
		return ret
	}
	return *o.SourceStart.Get()
}

// GetSourceStartOk returns a tuple with the SourceStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *FlakyTestRunMetadata) GetSourceStartOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceStart.Get(), o.SourceStart.IsSet()
}

// HasSourceStart returns a boolean if a field has been set.
func (o *FlakyTestRunMetadata) HasSourceStart() bool {
	return o != nil && o.SourceStart.IsSet()
}

// SetSourceStart gets a reference to the given datadog.NullableInt64 and assigns it to the SourceStart field.
func (o *FlakyTestRunMetadata) SetSourceStart(v int64) {
	o.SourceStart.Set(&v)
}

// SetSourceStartNil sets the value for SourceStart to be an explicit nil.
func (o *FlakyTestRunMetadata) SetSourceStartNil() {
	o.SourceStart.Set(nil)
}

// UnsetSourceStart ensures that no value is present for SourceStart, not even an explicit nil.
func (o *FlakyTestRunMetadata) UnsetSourceStart() {
	o.SourceStart.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o FlakyTestRunMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DurationMs.IsSet() {
		toSerialize["duration_ms"] = o.DurationMs.Get()
	}
	if o.ErrorMessage.IsSet() {
		toSerialize["error_message"] = o.ErrorMessage.Get()
	}
	if o.ErrorStack.IsSet() {
		toSerialize["error_stack"] = o.ErrorStack.Get()
	}
	if o.SourceEnd.IsSet() {
		toSerialize["source_end"] = o.SourceEnd.Get()
	}
	if o.SourceFile.IsSet() {
		toSerialize["source_file"] = o.SourceFile.Get()
	}
	if o.SourceStart.IsSet() {
		toSerialize["source_start"] = o.SourceStart.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FlakyTestRunMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DurationMs   datadog.NullableInt64  `json:"duration_ms,omitempty"`
		ErrorMessage datadog.NullableString `json:"error_message,omitempty"`
		ErrorStack   datadog.NullableString `json:"error_stack,omitempty"`
		SourceEnd    datadog.NullableInt64  `json:"source_end,omitempty"`
		SourceFile   datadog.NullableString `json:"source_file,omitempty"`
		SourceStart  datadog.NullableInt64  `json:"source_start,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"duration_ms", "error_message", "error_stack", "source_end", "source_file", "source_start"})
	} else {
		return err
	}
	o.DurationMs = all.DurationMs
	o.ErrorMessage = all.ErrorMessage
	o.ErrorStack = all.ErrorStack
	o.SourceEnd = all.SourceEnd
	o.SourceFile = all.SourceFile
	o.SourceStart = all.SourceStart

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
