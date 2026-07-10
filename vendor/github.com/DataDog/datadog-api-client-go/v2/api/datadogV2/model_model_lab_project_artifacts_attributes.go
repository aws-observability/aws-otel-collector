// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabProjectArtifactsAttributes Artifact listing for a Model Lab project.
type ModelLabProjectArtifactsAttributes struct {
	// The list of artifact files associated with the project.
	Files []ModelLabArtifactInfo `json:"files"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabProjectArtifactsAttributes instantiates a new ModelLabProjectArtifactsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabProjectArtifactsAttributes(files []ModelLabArtifactInfo) *ModelLabProjectArtifactsAttributes {
	this := ModelLabProjectArtifactsAttributes{}
	this.Files = files
	return &this
}

// NewModelLabProjectArtifactsAttributesWithDefaults instantiates a new ModelLabProjectArtifactsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabProjectArtifactsAttributesWithDefaults() *ModelLabProjectArtifactsAttributes {
	this := ModelLabProjectArtifactsAttributes{}
	return &this
}

// GetFiles returns the Files field value.
func (o *ModelLabProjectArtifactsAttributes) GetFiles() []ModelLabArtifactInfo {
	if o == nil {
		var ret []ModelLabArtifactInfo
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *ModelLabProjectArtifactsAttributes) GetFilesOk() (*[]ModelLabArtifactInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Files, true
}

// SetFiles sets field value.
func (o *ModelLabProjectArtifactsAttributes) SetFiles(v []ModelLabArtifactInfo) {
	o.Files = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabProjectArtifactsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["files"] = o.Files

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabProjectArtifactsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Files *[]ModelLabArtifactInfo `json:"files"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Files == nil {
		return fmt.Errorf("required field files missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"files"})
	} else {
		return err
	}
	o.Files = *all.Files

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
