// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IssueAttributes Object containing the information of an issue.
type IssueAttributes struct {
	// Error message associated with the issue.
	ErrorMessage *string `json:"error_message,omitempty"`
	// Type of the error that matches the issue.
	ErrorType *string `json:"error_type,omitempty"`
	// Path of the file where the issue occurred.
	FilePath *string `json:"file_path,omitempty"`
	// Timestamp of the first seen error in milliseconds since the Unix epoch.
	FirstSeen *int64 `json:"first_seen,omitempty"`
	// The application version (for example, git commit hash) where the issue was first observed.
	FirstSeenVersion *string `json:"first_seen_version,omitempty"`
	// Name of the function where the issue occurred.
	FunctionName *string `json:"function_name,omitempty"`
	// Error is a crash.
	IsCrash *bool `json:"is_crash,omitempty"`
	// Array of programming languages associated with the issue.
	Languages []IssueLanguage `json:"languages,omitempty"`
	// Timestamp of the last seen error in milliseconds since the Unix epoch.
	LastSeen *int64 `json:"last_seen,omitempty"`
	// The application version (for example, git commit hash) where the issue was last observed.
	LastSeenVersion *string `json:"last_seen_version,omitempty"`
	// Platform associated with the issue.
	Platform *IssuePlatform `json:"platform,omitempty"`
	// Service name.
	Service *string `json:"service,omitempty"`
	// State of the issue
	State *IssueState `json:"state,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIssueAttributes instantiates a new IssueAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIssueAttributes() *IssueAttributes {
	this := IssueAttributes{}
	return &this
}

// NewIssueAttributesWithDefaults instantiates a new IssueAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIssueAttributesWithDefaults() *IssueAttributes {
	this := IssueAttributes{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *IssueAttributes) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttributes) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *IssueAttributes) HasErrorMessage() bool {
	return o != nil && o.ErrorMessage != nil
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *IssueAttributes) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *IssueAttributes) GetErrorType() string {
	if o == nil || o.ErrorType == nil {
		var ret string
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttributes) GetErrorTypeOk() (*string, bool) {
	if o == nil || o.ErrorType == nil {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *IssueAttributes) HasErrorType() bool {
	return o != nil && o.ErrorType != nil
}

// SetErrorType gets a reference to the given string and assigns it to the ErrorType field.
func (o *IssueAttributes) SetErrorType(v string) {
	o.ErrorType = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *IssueAttributes) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttributes) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *IssueAttributes) HasFilePath() bool {
	return o != nil && o.FilePath != nil
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *IssueAttributes) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFirstSeen returns the FirstSeen field value if set, zero value otherwise.
func (o *IssueAttributes) GetFirstSeen() int64 {
	if o == nil || o.FirstSeen == nil {
		var ret int64
		return ret
	}
	return *o.FirstSeen
}

// GetFirstSeenOk returns a tuple with the FirstSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttributes) GetFirstSeenOk() (*int64, bool) {
	if o == nil || o.FirstSeen == nil {
		return nil, false
	}
	return o.FirstSeen, true
}

// HasFirstSeen returns a boolean if a field has been set.
func (o *IssueAttributes) HasFirstSeen() bool {
	return o != nil && o.FirstSeen != nil
}

// SetFirstSeen gets a reference to the given int64 and assigns it to the FirstSeen field.
func (o *IssueAttributes) SetFirstSeen(v int64) {
	o.FirstSeen = &v
}

// GetFirstSeenVersion returns the FirstSeenVersion field value if set, zero value otherwise.
func (o *IssueAttributes) GetFirstSeenVersion() string {
	if o == nil || o.FirstSeenVersion == nil {
		var ret string
		return ret
	}
	return *o.FirstSeenVersion
}

// GetFirstSeenVersionOk returns a tuple with the FirstSeenVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttributes) GetFirstSeenVersionOk() (*string, bool) {
	if o == nil || o.FirstSeenVersion == nil {
		return nil, false
	}
	return o.FirstSeenVersion, true
}

// HasFirstSeenVersion returns a boolean if a field has been set.
func (o *IssueAttributes) HasFirstSeenVersion() bool {
	return o != nil && o.FirstSeenVersion != nil
}

// SetFirstSeenVersion gets a reference to the given string and assigns it to the FirstSeenVersion field.
func (o *IssueAttributes) SetFirstSeenVersion(v string) {
	o.FirstSeenVersion = &v
}

// GetFunctionName returns the FunctionName field value if set, zero value otherwise.
func (o *IssueAttributes) GetFunctionName() string {
	if o == nil || o.FunctionName == nil {
		var ret string
		return ret
	}
	return *o.FunctionName
}

// GetFunctionNameOk returns a tuple with the FunctionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttributes) GetFunctionNameOk() (*string, bool) {
	if o == nil || o.FunctionName == nil {
		return nil, false
	}
	return o.FunctionName, true
}

// HasFunctionName returns a boolean if a field has been set.
func (o *IssueAttributes) HasFunctionName() bool {
	return o != nil && o.FunctionName != nil
}

// SetFunctionName gets a reference to the given string and assigns it to the FunctionName field.
func (o *IssueAttributes) SetFunctionName(v string) {
	o.FunctionName = &v
}

// GetIsCrash returns the IsCrash field value if set, zero value otherwise.
func (o *IssueAttributes) GetIsCrash() bool {
	if o == nil || o.IsCrash == nil {
		var ret bool
		return ret
	}
	return *o.IsCrash
}

// GetIsCrashOk returns a tuple with the IsCrash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttributes) GetIsCrashOk() (*bool, bool) {
	if o == nil || o.IsCrash == nil {
		return nil, false
	}
	return o.IsCrash, true
}

// HasIsCrash returns a boolean if a field has been set.
func (o *IssueAttributes) HasIsCrash() bool {
	return o != nil && o.IsCrash != nil
}

// SetIsCrash gets a reference to the given bool and assigns it to the IsCrash field.
func (o *IssueAttributes) SetIsCrash(v bool) {
	o.IsCrash = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *IssueAttributes) GetLanguages() []IssueLanguage {
	if o == nil || o.Languages == nil {
		var ret []IssueLanguage
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttributes) GetLanguagesOk() (*[]IssueLanguage, bool) {
	if o == nil || o.Languages == nil {
		return nil, false
	}
	return &o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *IssueAttributes) HasLanguages() bool {
	return o != nil && o.Languages != nil
}

// SetLanguages gets a reference to the given []IssueLanguage and assigns it to the Languages field.
func (o *IssueAttributes) SetLanguages(v []IssueLanguage) {
	o.Languages = v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *IssueAttributes) GetLastSeen() int64 {
	if o == nil || o.LastSeen == nil {
		var ret int64
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttributes) GetLastSeenOk() (*int64, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *IssueAttributes) HasLastSeen() bool {
	return o != nil && o.LastSeen != nil
}

// SetLastSeen gets a reference to the given int64 and assigns it to the LastSeen field.
func (o *IssueAttributes) SetLastSeen(v int64) {
	o.LastSeen = &v
}

// GetLastSeenVersion returns the LastSeenVersion field value if set, zero value otherwise.
func (o *IssueAttributes) GetLastSeenVersion() string {
	if o == nil || o.LastSeenVersion == nil {
		var ret string
		return ret
	}
	return *o.LastSeenVersion
}

// GetLastSeenVersionOk returns a tuple with the LastSeenVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttributes) GetLastSeenVersionOk() (*string, bool) {
	if o == nil || o.LastSeenVersion == nil {
		return nil, false
	}
	return o.LastSeenVersion, true
}

// HasLastSeenVersion returns a boolean if a field has been set.
func (o *IssueAttributes) HasLastSeenVersion() bool {
	return o != nil && o.LastSeenVersion != nil
}

// SetLastSeenVersion gets a reference to the given string and assigns it to the LastSeenVersion field.
func (o *IssueAttributes) SetLastSeenVersion(v string) {
	o.LastSeenVersion = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *IssueAttributes) GetPlatform() IssuePlatform {
	if o == nil || o.Platform == nil {
		var ret IssuePlatform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttributes) GetPlatformOk() (*IssuePlatform, bool) {
	if o == nil || o.Platform == nil {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *IssueAttributes) HasPlatform() bool {
	return o != nil && o.Platform != nil
}

// SetPlatform gets a reference to the given IssuePlatform and assigns it to the Platform field.
func (o *IssueAttributes) SetPlatform(v IssuePlatform) {
	o.Platform = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *IssueAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *IssueAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *IssueAttributes) SetService(v string) {
	o.Service = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *IssueAttributes) GetState() IssueState {
	if o == nil || o.State == nil {
		var ret IssueState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAttributes) GetStateOk() (*IssueState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *IssueAttributes) HasState() bool {
	return o != nil && o.State != nil
}

// SetState gets a reference to the given IssueState and assigns it to the State field.
func (o *IssueAttributes) SetState(v IssueState) {
	o.State = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IssueAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.ErrorType != nil {
		toSerialize["error_type"] = o.ErrorType
	}
	if o.FilePath != nil {
		toSerialize["file_path"] = o.FilePath
	}
	if o.FirstSeen != nil {
		toSerialize["first_seen"] = o.FirstSeen
	}
	if o.FirstSeenVersion != nil {
		toSerialize["first_seen_version"] = o.FirstSeenVersion
	}
	if o.FunctionName != nil {
		toSerialize["function_name"] = o.FunctionName
	}
	if o.IsCrash != nil {
		toSerialize["is_crash"] = o.IsCrash
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if o.LastSeen != nil {
		toSerialize["last_seen"] = o.LastSeen
	}
	if o.LastSeenVersion != nil {
		toSerialize["last_seen_version"] = o.LastSeenVersion
	}
	if o.Platform != nil {
		toSerialize["platform"] = o.Platform
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IssueAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ErrorMessage     *string         `json:"error_message,omitempty"`
		ErrorType        *string         `json:"error_type,omitempty"`
		FilePath         *string         `json:"file_path,omitempty"`
		FirstSeen        *int64          `json:"first_seen,omitempty"`
		FirstSeenVersion *string         `json:"first_seen_version,omitempty"`
		FunctionName     *string         `json:"function_name,omitempty"`
		IsCrash          *bool           `json:"is_crash,omitempty"`
		Languages        []IssueLanguage `json:"languages,omitempty"`
		LastSeen         *int64          `json:"last_seen,omitempty"`
		LastSeenVersion  *string         `json:"last_seen_version,omitempty"`
		Platform         *IssuePlatform  `json:"platform,omitempty"`
		Service          *string         `json:"service,omitempty"`
		State            *IssueState     `json:"state,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error_message", "error_type", "file_path", "first_seen", "first_seen_version", "function_name", "is_crash", "languages", "last_seen", "last_seen_version", "platform", "service", "state"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ErrorMessage = all.ErrorMessage
	o.ErrorType = all.ErrorType
	o.FilePath = all.FilePath
	o.FirstSeen = all.FirstSeen
	o.FirstSeenVersion = all.FirstSeenVersion
	o.FunctionName = all.FunctionName
	o.IsCrash = all.IsCrash
	o.Languages = all.Languages
	o.LastSeen = all.LastSeen
	o.LastSeenVersion = all.LastSeenVersion
	if all.Platform != nil && !all.Platform.IsValid() {
		hasInvalidField = true
	} else {
		o.Platform = all.Platform
	}
	o.Service = all.Service
	if all.State != nil && !all.State.IsValid() {
		hasInvalidField = true
	} else {
		o.State = all.State
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
