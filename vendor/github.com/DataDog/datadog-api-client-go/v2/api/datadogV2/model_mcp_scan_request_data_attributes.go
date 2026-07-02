// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// McpScanRequestDataAttributes The attributes of an MCP SCA scan request, describing the libraries to scan and their context.
type McpScanRequestDataAttributes struct {
	// The commit hash of the source code being scanned.
	CommitHash string `json:"commit_hash"`
	// The list of libraries to scan for vulnerabilities.
	Libraries []McpScanRequestDataAttributesLibrariesItems `json:"libraries"`
	// The name of the resource (typically the repository or project name) being scanned.
	ResourceName string `json:"resource_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMcpScanRequestDataAttributes instantiates a new McpScanRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMcpScanRequestDataAttributes(commitHash string, libraries []McpScanRequestDataAttributesLibrariesItems, resourceName string) *McpScanRequestDataAttributes {
	this := McpScanRequestDataAttributes{}
	this.CommitHash = commitHash
	this.Libraries = libraries
	this.ResourceName = resourceName
	return &this
}

// NewMcpScanRequestDataAttributesWithDefaults instantiates a new McpScanRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMcpScanRequestDataAttributesWithDefaults() *McpScanRequestDataAttributes {
	this := McpScanRequestDataAttributes{}
	return &this
}

// GetCommitHash returns the CommitHash field value.
func (o *McpScanRequestDataAttributes) GetCommitHash() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.CommitHash
}

// GetCommitHashOk returns a tuple with the CommitHash field value
// and a boolean to check if the value has been set.
func (o *McpScanRequestDataAttributes) GetCommitHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitHash, true
}

// SetCommitHash sets field value.
func (o *McpScanRequestDataAttributes) SetCommitHash(v string) {
	o.CommitHash = v
}

// GetLibraries returns the Libraries field value.
func (o *McpScanRequestDataAttributes) GetLibraries() []McpScanRequestDataAttributesLibrariesItems {
	if o == nil {
		var ret []McpScanRequestDataAttributesLibrariesItems
		return ret
	}
	return o.Libraries
}

// GetLibrariesOk returns a tuple with the Libraries field value
// and a boolean to check if the value has been set.
func (o *McpScanRequestDataAttributes) GetLibrariesOk() (*[]McpScanRequestDataAttributesLibrariesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Libraries, true
}

// SetLibraries sets field value.
func (o *McpScanRequestDataAttributes) SetLibraries(v []McpScanRequestDataAttributesLibrariesItems) {
	o.Libraries = v
}

// GetResourceName returns the ResourceName field value.
func (o *McpScanRequestDataAttributes) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *McpScanRequestDataAttributes) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value.
func (o *McpScanRequestDataAttributes) SetResourceName(v string) {
	o.ResourceName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o McpScanRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["commit_hash"] = o.CommitHash
	toSerialize["libraries"] = o.Libraries
	toSerialize["resource_name"] = o.ResourceName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *McpScanRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CommitHash   *string                                       `json:"commit_hash"`
		Libraries    *[]McpScanRequestDataAttributesLibrariesItems `json:"libraries"`
		ResourceName *string                                       `json:"resource_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.CommitHash == nil {
		return fmt.Errorf("required field commit_hash missing")
	}
	if all.Libraries == nil {
		return fmt.Errorf("required field libraries missing")
	}
	if all.ResourceName == nil {
		return fmt.Errorf("required field resource_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commit_hash", "libraries", "resource_name"})
	} else {
		return err
	}
	o.CommitHash = *all.CommitHash
	o.Libraries = *all.Libraries
	o.ResourceName = *all.ResourceName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
