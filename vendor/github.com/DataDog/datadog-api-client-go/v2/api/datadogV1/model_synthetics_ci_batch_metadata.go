// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsCIBatchMetadata Metadata for the Synthetic tests run.
type SyntheticsCIBatchMetadata struct {
	// Description of the CI provider.
	Ci *SyntheticsCIBatchMetadataCI `json:"ci,omitempty"`
	// Git information.
	Git *SyntheticsCIBatchMetadataGit `json:"git,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsCIBatchMetadata instantiates a new SyntheticsCIBatchMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsCIBatchMetadata() *SyntheticsCIBatchMetadata {
	this := SyntheticsCIBatchMetadata{}
	return &this
}

// NewSyntheticsCIBatchMetadataWithDefaults instantiates a new SyntheticsCIBatchMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsCIBatchMetadataWithDefaults() *SyntheticsCIBatchMetadata {
	this := SyntheticsCIBatchMetadata{}
	return &this
}

// GetCi returns the Ci field value if set, zero value otherwise.
func (o *SyntheticsCIBatchMetadata) GetCi() SyntheticsCIBatchMetadataCI {
	if o == nil || o.Ci == nil {
		var ret SyntheticsCIBatchMetadataCI
		return ret
	}
	return *o.Ci
}

// GetCiOk returns a tuple with the Ci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCIBatchMetadata) GetCiOk() (*SyntheticsCIBatchMetadataCI, bool) {
	if o == nil || o.Ci == nil {
		return nil, false
	}
	return o.Ci, true
}

// HasCi returns a boolean if a field has been set.
func (o *SyntheticsCIBatchMetadata) HasCi() bool {
	return o != nil && o.Ci != nil
}

// SetCi gets a reference to the given SyntheticsCIBatchMetadataCI and assigns it to the Ci field.
func (o *SyntheticsCIBatchMetadata) SetCi(v SyntheticsCIBatchMetadataCI) {
	o.Ci = &v
}

// GetGit returns the Git field value if set, zero value otherwise.
func (o *SyntheticsCIBatchMetadata) GetGit() SyntheticsCIBatchMetadataGit {
	if o == nil || o.Git == nil {
		var ret SyntheticsCIBatchMetadataGit
		return ret
	}
	return *o.Git
}

// GetGitOk returns a tuple with the Git field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCIBatchMetadata) GetGitOk() (*SyntheticsCIBatchMetadataGit, bool) {
	if o == nil || o.Git == nil {
		return nil, false
	}
	return o.Git, true
}

// HasGit returns a boolean if a field has been set.
func (o *SyntheticsCIBatchMetadata) HasGit() bool {
	return o != nil && o.Git != nil
}

// SetGit gets a reference to the given SyntheticsCIBatchMetadataGit and assigns it to the Git field.
func (o *SyntheticsCIBatchMetadata) SetGit(v SyntheticsCIBatchMetadataGit) {
	o.Git = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsCIBatchMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Ci != nil {
		toSerialize["ci"] = o.Ci
	}
	if o.Git != nil {
		toSerialize["git"] = o.Git
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsCIBatchMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ci  *SyntheticsCIBatchMetadataCI  `json:"ci,omitempty"`
		Git *SyntheticsCIBatchMetadataGit `json:"git,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ci", "git"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Ci != nil && all.Ci.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Ci = all.Ci
	if all.Git != nil && all.Git.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Git = all.Git

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
