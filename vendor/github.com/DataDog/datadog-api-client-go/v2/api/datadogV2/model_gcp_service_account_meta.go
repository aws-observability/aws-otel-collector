// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPServiceAccountMeta Additional information related to your service account.
type GCPServiceAccountMeta struct {
	// The current list of projects accessible from your service account.
	AccessibleProjects []string `json:"accessible_projects,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGCPServiceAccountMeta instantiates a new GCPServiceAccountMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPServiceAccountMeta() *GCPServiceAccountMeta {
	this := GCPServiceAccountMeta{}
	return &this
}

// NewGCPServiceAccountMetaWithDefaults instantiates a new GCPServiceAccountMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPServiceAccountMetaWithDefaults() *GCPServiceAccountMeta {
	this := GCPServiceAccountMeta{}
	return &this
}

// GetAccessibleProjects returns the AccessibleProjects field value if set, zero value otherwise.
func (o *GCPServiceAccountMeta) GetAccessibleProjects() []string {
	if o == nil || o.AccessibleProjects == nil {
		var ret []string
		return ret
	}
	return o.AccessibleProjects
}

// GetAccessibleProjectsOk returns a tuple with the AccessibleProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPServiceAccountMeta) GetAccessibleProjectsOk() (*[]string, bool) {
	if o == nil || o.AccessibleProjects == nil {
		return nil, false
	}
	return &o.AccessibleProjects, true
}

// HasAccessibleProjects returns a boolean if a field has been set.
func (o *GCPServiceAccountMeta) HasAccessibleProjects() bool {
	return o != nil && o.AccessibleProjects != nil
}

// SetAccessibleProjects gets a reference to the given []string and assigns it to the AccessibleProjects field.
func (o *GCPServiceAccountMeta) SetAccessibleProjects(v []string) {
	o.AccessibleProjects = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPServiceAccountMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AccessibleProjects != nil {
		toSerialize["accessible_projects"] = o.AccessibleProjects
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GCPServiceAccountMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessibleProjects []string `json:"accessible_projects,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"accessible_projects"})
	} else {
		return err
	}
	o.AccessibleProjects = all.AccessibleProjects

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
