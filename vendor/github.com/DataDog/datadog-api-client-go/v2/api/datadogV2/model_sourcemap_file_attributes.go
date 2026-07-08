// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SourcemapFileAttributes Attributes of a JavaScript source map file.
type SourcemapFileAttributes struct {
	// The name of the minified JavaScript file.
	File string `json:"file"`
	// The Base64 VLQ encoded string that maps positions in the minified
	// file to positions in the original source files.
	Mappings string `json:"mappings"`
	// List of character counts for each line in the minified file.
	MinifiedLineLengths []int64 `json:"minifiedLineLengths"`
	// List of symbol names referenced in the mappings.
	Names []interface{} `json:"names"`
	// The root path prepended to source file paths.
	SourceRoot string `json:"sourceRoot"`
	// List of original source file paths.
	Sources []string `json:"sources"`
	// List of original source file contents corresponding to the paths in `sources`.
	SourcesContent []string `json:"sourcesContent"`
	// The version of the source map format (typically 3).
	Version int64 `json:"version"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSourcemapFileAttributes instantiates a new SourcemapFileAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSourcemapFileAttributes(file string, mappings string, minifiedLineLengths []int64, names []interface{}, sourceRoot string, sources []string, sourcesContent []string, version int64) *SourcemapFileAttributes {
	this := SourcemapFileAttributes{}
	this.File = file
	this.Mappings = mappings
	this.MinifiedLineLengths = minifiedLineLengths
	this.Names = names
	this.SourceRoot = sourceRoot
	this.Sources = sources
	this.SourcesContent = sourcesContent
	this.Version = version
	return &this
}

// NewSourcemapFileAttributesWithDefaults instantiates a new SourcemapFileAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSourcemapFileAttributesWithDefaults() *SourcemapFileAttributes {
	this := SourcemapFileAttributes{}
	return &this
}

// GetFile returns the File field value.
func (o *SourcemapFileAttributes) GetFile() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *SourcemapFileAttributes) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value.
func (o *SourcemapFileAttributes) SetFile(v string) {
	o.File = v
}

// GetMappings returns the Mappings field value.
func (o *SourcemapFileAttributes) GetMappings() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Mappings
}

// GetMappingsOk returns a tuple with the Mappings field value
// and a boolean to check if the value has been set.
func (o *SourcemapFileAttributes) GetMappingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mappings, true
}

// SetMappings sets field value.
func (o *SourcemapFileAttributes) SetMappings(v string) {
	o.Mappings = v
}

// GetMinifiedLineLengths returns the MinifiedLineLengths field value.
func (o *SourcemapFileAttributes) GetMinifiedLineLengths() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.MinifiedLineLengths
}

// GetMinifiedLineLengthsOk returns a tuple with the MinifiedLineLengths field value
// and a boolean to check if the value has been set.
func (o *SourcemapFileAttributes) GetMinifiedLineLengthsOk() (*[]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinifiedLineLengths, true
}

// SetMinifiedLineLengths sets field value.
func (o *SourcemapFileAttributes) SetMinifiedLineLengths(v []int64) {
	o.MinifiedLineLengths = v
}

// GetNames returns the Names field value.
func (o *SourcemapFileAttributes) GetNames() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value
// and a boolean to check if the value has been set.
func (o *SourcemapFileAttributes) GetNamesOk() (*[]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Names, true
}

// SetNames sets field value.
func (o *SourcemapFileAttributes) SetNames(v []interface{}) {
	o.Names = v
}

// GetSourceRoot returns the SourceRoot field value.
func (o *SourcemapFileAttributes) GetSourceRoot() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceRoot
}

// GetSourceRootOk returns a tuple with the SourceRoot field value
// and a boolean to check if the value has been set.
func (o *SourcemapFileAttributes) GetSourceRootOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceRoot, true
}

// SetSourceRoot sets field value.
func (o *SourcemapFileAttributes) SetSourceRoot(v string) {
	o.SourceRoot = v
}

// GetSources returns the Sources field value.
func (o *SourcemapFileAttributes) GetSources() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value
// and a boolean to check if the value has been set.
func (o *SourcemapFileAttributes) GetSourcesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sources, true
}

// SetSources sets field value.
func (o *SourcemapFileAttributes) SetSources(v []string) {
	o.Sources = v
}

// GetSourcesContent returns the SourcesContent field value.
func (o *SourcemapFileAttributes) GetSourcesContent() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SourcesContent
}

// GetSourcesContentOk returns a tuple with the SourcesContent field value
// and a boolean to check if the value has been set.
func (o *SourcemapFileAttributes) GetSourcesContentOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourcesContent, true
}

// SetSourcesContent sets field value.
func (o *SourcemapFileAttributes) SetSourcesContent(v []string) {
	o.SourcesContent = v
}

// GetVersion returns the Version field value.
func (o *SourcemapFileAttributes) GetVersion() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *SourcemapFileAttributes) GetVersionOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value.
func (o *SourcemapFileAttributes) SetVersion(v int64) {
	o.Version = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SourcemapFileAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["file"] = o.File
	toSerialize["mappings"] = o.Mappings
	toSerialize["minifiedLineLengths"] = o.MinifiedLineLengths
	toSerialize["names"] = o.Names
	toSerialize["sourceRoot"] = o.SourceRoot
	toSerialize["sources"] = o.Sources
	toSerialize["sourcesContent"] = o.SourcesContent
	toSerialize["version"] = o.Version

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SourcemapFileAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		File                *string        `json:"file"`
		Mappings            *string        `json:"mappings"`
		MinifiedLineLengths *[]int64       `json:"minifiedLineLengths"`
		Names               *[]interface{} `json:"names"`
		SourceRoot          *string        `json:"sourceRoot"`
		Sources             *[]string      `json:"sources"`
		SourcesContent      *[]string      `json:"sourcesContent"`
		Version             *int64         `json:"version"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.File == nil {
		return fmt.Errorf("required field file missing")
	}
	if all.Mappings == nil {
		return fmt.Errorf("required field mappings missing")
	}
	if all.MinifiedLineLengths == nil {
		return fmt.Errorf("required field minifiedLineLengths missing")
	}
	if all.Names == nil {
		return fmt.Errorf("required field names missing")
	}
	if all.SourceRoot == nil {
		return fmt.Errorf("required field sourceRoot missing")
	}
	if all.Sources == nil {
		return fmt.Errorf("required field sources missing")
	}
	if all.SourcesContent == nil {
		return fmt.Errorf("required field sourcesContent missing")
	}
	if all.Version == nil {
		return fmt.Errorf("required field version missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"file", "mappings", "minifiedLineLengths", "names", "sourceRoot", "sources", "sourcesContent", "version"})
	} else {
		return err
	}
	o.File = *all.File
	o.Mappings = *all.Mappings
	o.MinifiedLineLengths = *all.MinifiedLineLengths
	o.Names = *all.Names
	o.SourceRoot = *all.SourceRoot
	o.Sources = *all.Sources
	o.SourcesContent = *all.SourcesContent
	o.Version = *all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
