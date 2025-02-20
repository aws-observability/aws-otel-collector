// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3DatadogCodeLocationItem Code location item.
type EntityV3DatadogCodeLocationItem struct {
	// The paths (glob) to the source code of the service.
	Paths []string `json:"paths,omitempty"`
	// The repository path of the source code of the entity.
	RepositoryUrl *string `json:"repositoryURL,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3DatadogCodeLocationItem instantiates a new EntityV3DatadogCodeLocationItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3DatadogCodeLocationItem() *EntityV3DatadogCodeLocationItem {
	this := EntityV3DatadogCodeLocationItem{}
	return &this
}

// NewEntityV3DatadogCodeLocationItemWithDefaults instantiates a new EntityV3DatadogCodeLocationItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3DatadogCodeLocationItemWithDefaults() *EntityV3DatadogCodeLocationItem {
	this := EntityV3DatadogCodeLocationItem{}
	return &this
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *EntityV3DatadogCodeLocationItem) GetPaths() []string {
	if o == nil || o.Paths == nil {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3DatadogCodeLocationItem) GetPathsOk() (*[]string, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return &o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *EntityV3DatadogCodeLocationItem) HasPaths() bool {
	return o != nil && o.Paths != nil
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *EntityV3DatadogCodeLocationItem) SetPaths(v []string) {
	o.Paths = v
}

// GetRepositoryUrl returns the RepositoryUrl field value if set, zero value otherwise.
func (o *EntityV3DatadogCodeLocationItem) GetRepositoryUrl() string {
	if o == nil || o.RepositoryUrl == nil {
		var ret string
		return ret
	}
	return *o.RepositoryUrl
}

// GetRepositoryUrlOk returns a tuple with the RepositoryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3DatadogCodeLocationItem) GetRepositoryUrlOk() (*string, bool) {
	if o == nil || o.RepositoryUrl == nil {
		return nil, false
	}
	return o.RepositoryUrl, true
}

// HasRepositoryUrl returns a boolean if a field has been set.
func (o *EntityV3DatadogCodeLocationItem) HasRepositoryUrl() bool {
	return o != nil && o.RepositoryUrl != nil
}

// SetRepositoryUrl gets a reference to the given string and assigns it to the RepositoryUrl field.
func (o *EntityV3DatadogCodeLocationItem) SetRepositoryUrl(v string) {
	o.RepositoryUrl = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3DatadogCodeLocationItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}
	if o.RepositoryUrl != nil {
		toSerialize["repositoryURL"] = o.RepositoryUrl
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3DatadogCodeLocationItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Paths         []string `json:"paths,omitempty"`
		RepositoryUrl *string  `json:"repositoryURL,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.Paths = all.Paths
	o.RepositoryUrl = all.RepositoryUrl

	return nil
}
