// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ModelLabRunArtifactsAttributes Artifact listing for a Model Lab run.
type ModelLabRunArtifactsAttributes struct {
	// The list of artifact files and directories.
	Files []ModelLabArtifactObjectInfo `json:"files"`
	// The path of the run's artifacts relative to the project's artifact root.
	PathInProject string `json:"path_in_project"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewModelLabRunArtifactsAttributes instantiates a new ModelLabRunArtifactsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewModelLabRunArtifactsAttributes(files []ModelLabArtifactObjectInfo, pathInProject string) *ModelLabRunArtifactsAttributes {
	this := ModelLabRunArtifactsAttributes{}
	this.Files = files
	this.PathInProject = pathInProject
	return &this
}

// NewModelLabRunArtifactsAttributesWithDefaults instantiates a new ModelLabRunArtifactsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewModelLabRunArtifactsAttributesWithDefaults() *ModelLabRunArtifactsAttributes {
	this := ModelLabRunArtifactsAttributes{}
	return &this
}

// GetFiles returns the Files field value.
func (o *ModelLabRunArtifactsAttributes) GetFiles() []ModelLabArtifactObjectInfo {
	if o == nil {
		var ret []ModelLabArtifactObjectInfo
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunArtifactsAttributes) GetFilesOk() (*[]ModelLabArtifactObjectInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Files, true
}

// SetFiles sets field value.
func (o *ModelLabRunArtifactsAttributes) SetFiles(v []ModelLabArtifactObjectInfo) {
	o.Files = v
}

// GetPathInProject returns the PathInProject field value.
func (o *ModelLabRunArtifactsAttributes) GetPathInProject() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PathInProject
}

// GetPathInProjectOk returns a tuple with the PathInProject field value
// and a boolean to check if the value has been set.
func (o *ModelLabRunArtifactsAttributes) GetPathInProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PathInProject, true
}

// SetPathInProject sets field value.
func (o *ModelLabRunArtifactsAttributes) SetPathInProject(v string) {
	o.PathInProject = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ModelLabRunArtifactsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["files"] = o.Files
	toSerialize["path_in_project"] = o.PathInProject

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ModelLabRunArtifactsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Files         *[]ModelLabArtifactObjectInfo `json:"files"`
		PathInProject *string                       `json:"path_in_project"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Files == nil {
		return fmt.Errorf("required field files missing")
	}
	if all.PathInProject == nil {
		return fmt.Errorf("required field path_in_project missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"files", "path_in_project"})
	} else {
		return err
	}
	o.Files = *all.Files
	o.PathInProject = *all.PathInProject

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
