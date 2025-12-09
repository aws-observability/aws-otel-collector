// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributes
type ScaRequestDataAttributes struct {
	//
	Commit *ScaRequestDataAttributesCommit `json:"commit,omitempty"`
	//
	Dependencies []ScaRequestDataAttributesDependenciesItems `json:"dependencies,omitempty"`
	//
	Env *string `json:"env,omitempty"`
	//
	Files []ScaRequestDataAttributesFilesItems `json:"files,omitempty"`
	//
	Relations []ScaRequestDataAttributesRelationsItems `json:"relations,omitempty"`
	//
	Repository *ScaRequestDataAttributesRepository `json:"repository,omitempty"`
	//
	Service *string `json:"service,omitempty"`
	//
	Tags map[string]string `json:"tags,omitempty"`
	//
	Vulnerabilities []ScaRequestDataAttributesVulnerabilitiesItems `json:"vulnerabilities,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributes instantiates a new ScaRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributes() *ScaRequestDataAttributes {
	this := ScaRequestDataAttributes{}
	return &this
}

// NewScaRequestDataAttributesWithDefaults instantiates a new ScaRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesWithDefaults() *ScaRequestDataAttributes {
	this := ScaRequestDataAttributes{}
	return &this
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *ScaRequestDataAttributes) GetCommit() ScaRequestDataAttributesCommit {
	if o == nil || o.Commit == nil {
		var ret ScaRequestDataAttributesCommit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributes) GetCommitOk() (*ScaRequestDataAttributesCommit, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *ScaRequestDataAttributes) HasCommit() bool {
	return o != nil && o.Commit != nil
}

// SetCommit gets a reference to the given ScaRequestDataAttributesCommit and assigns it to the Commit field.
func (o *ScaRequestDataAttributes) SetCommit(v ScaRequestDataAttributesCommit) {
	o.Commit = &v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *ScaRequestDataAttributes) GetDependencies() []ScaRequestDataAttributesDependenciesItems {
	if o == nil || o.Dependencies == nil {
		var ret []ScaRequestDataAttributesDependenciesItems
		return ret
	}
	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributes) GetDependenciesOk() (*[]ScaRequestDataAttributesDependenciesItems, bool) {
	if o == nil || o.Dependencies == nil {
		return nil, false
	}
	return &o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *ScaRequestDataAttributes) HasDependencies() bool {
	return o != nil && o.Dependencies != nil
}

// SetDependencies gets a reference to the given []ScaRequestDataAttributesDependenciesItems and assigns it to the Dependencies field.
func (o *ScaRequestDataAttributes) SetDependencies(v []ScaRequestDataAttributesDependenciesItems) {
	o.Dependencies = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *ScaRequestDataAttributes) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributes) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *ScaRequestDataAttributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *ScaRequestDataAttributes) SetEnv(v string) {
	o.Env = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *ScaRequestDataAttributes) GetFiles() []ScaRequestDataAttributesFilesItems {
	if o == nil || o.Files == nil {
		var ret []ScaRequestDataAttributesFilesItems
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributes) GetFilesOk() (*[]ScaRequestDataAttributesFilesItems, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return &o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *ScaRequestDataAttributes) HasFiles() bool {
	return o != nil && o.Files != nil
}

// SetFiles gets a reference to the given []ScaRequestDataAttributesFilesItems and assigns it to the Files field.
func (o *ScaRequestDataAttributes) SetFiles(v []ScaRequestDataAttributesFilesItems) {
	o.Files = v
}

// GetRelations returns the Relations field value if set, zero value otherwise.
func (o *ScaRequestDataAttributes) GetRelations() []ScaRequestDataAttributesRelationsItems {
	if o == nil || o.Relations == nil {
		var ret []ScaRequestDataAttributesRelationsItems
		return ret
	}
	return o.Relations
}

// GetRelationsOk returns a tuple with the Relations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributes) GetRelationsOk() (*[]ScaRequestDataAttributesRelationsItems, bool) {
	if o == nil || o.Relations == nil {
		return nil, false
	}
	return &o.Relations, true
}

// HasRelations returns a boolean if a field has been set.
func (o *ScaRequestDataAttributes) HasRelations() bool {
	return o != nil && o.Relations != nil
}

// SetRelations gets a reference to the given []ScaRequestDataAttributesRelationsItems and assigns it to the Relations field.
func (o *ScaRequestDataAttributes) SetRelations(v []ScaRequestDataAttributesRelationsItems) {
	o.Relations = v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ScaRequestDataAttributes) GetRepository() ScaRequestDataAttributesRepository {
	if o == nil || o.Repository == nil {
		var ret ScaRequestDataAttributesRepository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributes) GetRepositoryOk() (*ScaRequestDataAttributesRepository, bool) {
	if o == nil || o.Repository == nil {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ScaRequestDataAttributes) HasRepository() bool {
	return o != nil && o.Repository != nil
}

// SetRepository gets a reference to the given ScaRequestDataAttributesRepository and assigns it to the Repository field.
func (o *ScaRequestDataAttributes) SetRepository(v ScaRequestDataAttributesRepository) {
	o.Repository = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ScaRequestDataAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ScaRequestDataAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ScaRequestDataAttributes) SetService(v string) {
	o.Service = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ScaRequestDataAttributes) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributes) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ScaRequestDataAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *ScaRequestDataAttributes) SetTags(v map[string]string) {
	o.Tags = v
}

// GetVulnerabilities returns the Vulnerabilities field value if set, zero value otherwise.
func (o *ScaRequestDataAttributes) GetVulnerabilities() []ScaRequestDataAttributesVulnerabilitiesItems {
	if o == nil || o.Vulnerabilities == nil {
		var ret []ScaRequestDataAttributesVulnerabilitiesItems
		return ret
	}
	return o.Vulnerabilities
}

// GetVulnerabilitiesOk returns a tuple with the Vulnerabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributes) GetVulnerabilitiesOk() (*[]ScaRequestDataAttributesVulnerabilitiesItems, bool) {
	if o == nil || o.Vulnerabilities == nil {
		return nil, false
	}
	return &o.Vulnerabilities, true
}

// HasVulnerabilities returns a boolean if a field has been set.
func (o *ScaRequestDataAttributes) HasVulnerabilities() bool {
	return o != nil && o.Vulnerabilities != nil
}

// SetVulnerabilities gets a reference to the given []ScaRequestDataAttributesVulnerabilitiesItems and assigns it to the Vulnerabilities field.
func (o *ScaRequestDataAttributes) SetVulnerabilities(v []ScaRequestDataAttributesVulnerabilitiesItems) {
	o.Vulnerabilities = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	if o.Dependencies != nil {
		toSerialize["dependencies"] = o.Dependencies
	}
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.Relations != nil {
		toSerialize["relations"] = o.Relations
	}
	if o.Repository != nil {
		toSerialize["repository"] = o.Repository
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Vulnerabilities != nil {
		toSerialize["vulnerabilities"] = o.Vulnerabilities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Commit          *ScaRequestDataAttributesCommit                `json:"commit,omitempty"`
		Dependencies    []ScaRequestDataAttributesDependenciesItems    `json:"dependencies,omitempty"`
		Env             *string                                        `json:"env,omitempty"`
		Files           []ScaRequestDataAttributesFilesItems           `json:"files,omitempty"`
		Relations       []ScaRequestDataAttributesRelationsItems       `json:"relations,omitempty"`
		Repository      *ScaRequestDataAttributesRepository            `json:"repository,omitempty"`
		Service         *string                                        `json:"service,omitempty"`
		Tags            map[string]string                              `json:"tags,omitempty"`
		Vulnerabilities []ScaRequestDataAttributesVulnerabilitiesItems `json:"vulnerabilities,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"commit", "dependencies", "env", "files", "relations", "repository", "service", "tags", "vulnerabilities"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Commit != nil && all.Commit.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Commit = all.Commit
	o.Dependencies = all.Dependencies
	o.Env = all.Env
	o.Files = all.Files
	o.Relations = all.Relations
	if all.Repository != nil && all.Repository.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Repository = all.Repository
	o.Service = all.Service
	o.Tags = all.Tags
	o.Vulnerabilities = all.Vulnerabilities

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
