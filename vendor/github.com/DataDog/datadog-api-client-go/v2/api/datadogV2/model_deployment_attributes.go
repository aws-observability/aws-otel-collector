// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeploymentAttributes The attributes object containing the version ID of the published app.
type DeploymentAttributes struct {
	// The version ID of the app that was published. For an unpublished app, this is always the nil UUID (`00000000-0000-0000-0000-000000000000`).
	AppVersionId *uuid.UUID `json:"app_version_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeploymentAttributes instantiates a new DeploymentAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeploymentAttributes() *DeploymentAttributes {
	this := DeploymentAttributes{}
	return &this
}

// NewDeploymentAttributesWithDefaults instantiates a new DeploymentAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeploymentAttributesWithDefaults() *DeploymentAttributes {
	this := DeploymentAttributes{}
	return &this
}

// GetAppVersionId returns the AppVersionId field value if set, zero value otherwise.
func (o *DeploymentAttributes) GetAppVersionId() uuid.UUID {
	if o == nil || o.AppVersionId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.AppVersionId
}

// GetAppVersionIdOk returns a tuple with the AppVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeploymentAttributes) GetAppVersionIdOk() (*uuid.UUID, bool) {
	if o == nil || o.AppVersionId == nil {
		return nil, false
	}
	return o.AppVersionId, true
}

// HasAppVersionId returns a boolean if a field has been set.
func (o *DeploymentAttributes) HasAppVersionId() bool {
	return o != nil && o.AppVersionId != nil
}

// SetAppVersionId gets a reference to the given uuid.UUID and assigns it to the AppVersionId field.
func (o *DeploymentAttributes) SetAppVersionId(v uuid.UUID) {
	o.AppVersionId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeploymentAttributes) MarshalJSON() ([]byte, error) {
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
func (o *DeploymentAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppVersionId *uuid.UUID `json:"app_version_id,omitempty"`
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
