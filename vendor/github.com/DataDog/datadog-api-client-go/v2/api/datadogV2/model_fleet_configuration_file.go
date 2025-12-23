// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetConfigurationFile A configuration file for an integration.
type FleetConfigurationFile struct {
	// The raw content of the configuration file.
	FileContent *string `json:"file_content,omitempty"`
	// Path to the configuration file.
	FilePath *string `json:"file_path,omitempty"`
	// Name of the configuration file.
	Filename *string `json:"filename,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetConfigurationFile instantiates a new FleetConfigurationFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetConfigurationFile() *FleetConfigurationFile {
	this := FleetConfigurationFile{}
	return &this
}

// NewFleetConfigurationFileWithDefaults instantiates a new FleetConfigurationFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetConfigurationFileWithDefaults() *FleetConfigurationFile {
	this := FleetConfigurationFile{}
	return &this
}

// GetFileContent returns the FileContent field value if set, zero value otherwise.
func (o *FleetConfigurationFile) GetFileContent() string {
	if o == nil || o.FileContent == nil {
		var ret string
		return ret
	}
	return *o.FileContent
}

// GetFileContentOk returns a tuple with the FileContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetConfigurationFile) GetFileContentOk() (*string, bool) {
	if o == nil || o.FileContent == nil {
		return nil, false
	}
	return o.FileContent, true
}

// HasFileContent returns a boolean if a field has been set.
func (o *FleetConfigurationFile) HasFileContent() bool {
	return o != nil && o.FileContent != nil
}

// SetFileContent gets a reference to the given string and assigns it to the FileContent field.
func (o *FleetConfigurationFile) SetFileContent(v string) {
	o.FileContent = &v
}

// GetFilePath returns the FilePath field value if set, zero value otherwise.
func (o *FleetConfigurationFile) GetFilePath() string {
	if o == nil || o.FilePath == nil {
		var ret string
		return ret
	}
	return *o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetConfigurationFile) GetFilePathOk() (*string, bool) {
	if o == nil || o.FilePath == nil {
		return nil, false
	}
	return o.FilePath, true
}

// HasFilePath returns a boolean if a field has been set.
func (o *FleetConfigurationFile) HasFilePath() bool {
	return o != nil && o.FilePath != nil
}

// SetFilePath gets a reference to the given string and assigns it to the FilePath field.
func (o *FleetConfigurationFile) SetFilePath(v string) {
	o.FilePath = &v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *FleetConfigurationFile) GetFilename() string {
	if o == nil || o.Filename == nil {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetConfigurationFile) GetFilenameOk() (*string, bool) {
	if o == nil || o.Filename == nil {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *FleetConfigurationFile) HasFilename() bool {
	return o != nil && o.Filename != nil
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *FleetConfigurationFile) SetFilename(v string) {
	o.Filename = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetConfigurationFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FileContent != nil {
		toSerialize["file_content"] = o.FileContent
	}
	if o.FilePath != nil {
		toSerialize["file_path"] = o.FilePath
	}
	if o.Filename != nil {
		toSerialize["filename"] = o.Filename
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetConfigurationFile) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FileContent *string `json:"file_content,omitempty"`
		FilePath    *string `json:"file_path,omitempty"`
		Filename    *string `json:"filename,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"file_content", "file_path", "filename"})
	} else {
		return err
	}
	o.FileContent = all.FileContent
	o.FilePath = all.FilePath
	o.Filename = all.Filename

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
