// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentIncludedAttributes The definition of `DeploymentIncludedAttributes` object.
type DeploymentIncludedAttributes struct {
	// The `attributes` `app_version_id`.
	AppVersionId *string `json:"app_version_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentIncludedAttributes instantiates a new DeploymentIncludedAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentIncludedAttributes() *DeploymentIncludedAttributes {
	this := DeploymentIncludedAttributes{}
	return &this
}

// NewDeploymentIncludedAttributesWithDefaults instantiates a new DeploymentIncludedAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentIncludedAttributesWithDefaults() *DeploymentIncludedAttributes {
	this := DeploymentIncludedAttributes{}
	return &this
}

// GetAppVersionId returns the AppVersionId field value if set, zero value otherwise.
func (o *DeploymentIncludedAttributes) GetAppVersionId() string {
	if o == nil || o.AppVersionId == nil {
		var ret string
		return ret
	}
	return *o.AppVersionId
}

// GetAppVersionIdOk returns a tuple with the AppVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentIncludedAttributes) GetAppVersionIdOk() (*string, bool) {
	if o == nil || o.AppVersionId == nil {
		return nil, false
	}
	return o.AppVersionId, true
}

// HasAppVersionId returns a boolean if a field has been set.
func (o *DeploymentIncludedAttributes) HasAppVersionId() bool {
	return o != nil && o.AppVersionId != nil
}

// SetAppVersionId gets a reference to the given string and assigns it to the AppVersionId field.
func (o *DeploymentIncludedAttributes) SetAppVersionId(v string) {
	o.AppVersionId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentIncludedAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AppVersionId != nil {
		toSerialize["app_version_id"] = o.AppVersionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeploymentIncludedAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppVersionId *string `json:"app_version_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"app_version_id"})
	} else {
		return err
	}
	o.AppVersionId = all.AppVersionId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
