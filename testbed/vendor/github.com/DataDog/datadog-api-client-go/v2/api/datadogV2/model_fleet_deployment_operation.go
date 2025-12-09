// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FleetDeploymentOperation A single configuration file operation to perform on the target hosts.
type FleetDeploymentOperation struct {
	// Type of file operation to perform on the target configuration file.
	// - `merge-patch`: Merges the provided patch data with the existing configuration file.
	//   Creates the file if it doesn't exist.
	// - `delete`: Removes the specified configuration file from the target hosts.
	FileOp FleetDeploymentFileOp `json:"file_op"`
	// Absolute path to the target configuration file on the host.
	FilePath string `json:"file_path"`
	// Patch data in JSON format to apply to the configuration file.
	// When using `merge-patch`, this object is merged with the existing configuration,
	// allowing you to add, update, or override specific fields without replacing the entire file.
	// The structure must match the target configuration file format (for example, YAML structure
	// for Datadog Agent config). Not applicable when using the `delete` operation.
	Patch map[string]interface{} `json:"patch,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFleetDeploymentOperation instantiates a new FleetDeploymentOperation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFleetDeploymentOperation(fileOp FleetDeploymentFileOp, filePath string) *FleetDeploymentOperation {
	this := FleetDeploymentOperation{}
	this.FileOp = fileOp
	this.FilePath = filePath
	return &this
}

// NewFleetDeploymentOperationWithDefaults instantiates a new FleetDeploymentOperation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFleetDeploymentOperationWithDefaults() *FleetDeploymentOperation {
	this := FleetDeploymentOperation{}
	return &this
}

// GetFileOp returns the FileOp field value.
func (o *FleetDeploymentOperation) GetFileOp() FleetDeploymentFileOp {
	if o == nil {
		var ret FleetDeploymentFileOp
		return ret
	}
	return o.FileOp
}

// GetFileOpOk returns a tuple with the FileOp field value
// and a boolean to check if the value has been set.
func (o *FleetDeploymentOperation) GetFileOpOk() (*FleetDeploymentFileOp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileOp, true
}

// SetFileOp sets field value.
func (o *FleetDeploymentOperation) SetFileOp(v FleetDeploymentFileOp) {
	o.FileOp = v
}

// GetFilePath returns the FilePath field value.
func (o *FleetDeploymentOperation) GetFilePath() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FilePath
}

// GetFilePathOk returns a tuple with the FilePath field value
// and a boolean to check if the value has been set.
func (o *FleetDeploymentOperation) GetFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilePath, true
}

// SetFilePath sets field value.
func (o *FleetDeploymentOperation) SetFilePath(v string) {
	o.FilePath = v
}

// GetPatch returns the Patch field value if set, zero value otherwise.
func (o *FleetDeploymentOperation) GetPatch() map[string]interface{} {
	if o == nil || o.Patch == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Patch
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FleetDeploymentOperation) GetPatchOk() (*map[string]interface{}, bool) {
	if o == nil || o.Patch == nil {
		return nil, false
	}
	return &o.Patch, true
}

// HasPatch returns a boolean if a field has been set.
func (o *FleetDeploymentOperation) HasPatch() bool {
	return o != nil && o.Patch != nil
}

// SetPatch gets a reference to the given map[string]interface{} and assigns it to the Patch field.
func (o *FleetDeploymentOperation) SetPatch(v map[string]interface{}) {
	o.Patch = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FleetDeploymentOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["file_op"] = o.FileOp
	toSerialize["file_path"] = o.FilePath
	if o.Patch != nil {
		toSerialize["patch"] = o.Patch
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FleetDeploymentOperation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FileOp   *FleetDeploymentFileOp `json:"file_op"`
		FilePath *string                `json:"file_path"`
		Patch    map[string]interface{} `json:"patch,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.FileOp == nil {
		return fmt.Errorf("required field file_op missing")
	}
	if all.FilePath == nil {
		return fmt.Errorf("required field file_path missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"file_op", "file_path", "patch"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.FileOp.IsValid() {
		hasInvalidField = true
	} else {
		o.FileOp = *all.FileOp
	}
	o.FilePath = *all.FilePath
	o.Patch = all.Patch

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
