// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListAppsResponseDataItemsRelationships The app's publication information.
type ListAppsResponseDataItemsRelationships struct {
	// Information pointing to the app's publication status.
	Deployment *DeploymentRelationship `json:"deployment,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListAppsResponseDataItemsRelationships instantiates a new ListAppsResponseDataItemsRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListAppsResponseDataItemsRelationships() *ListAppsResponseDataItemsRelationships {
	this := ListAppsResponseDataItemsRelationships{}
	return &this
}

// NewListAppsResponseDataItemsRelationshipsWithDefaults instantiates a new ListAppsResponseDataItemsRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListAppsResponseDataItemsRelationshipsWithDefaults() *ListAppsResponseDataItemsRelationships {
	this := ListAppsResponseDataItemsRelationships{}
	return &this
}

// GetDeployment returns the Deployment field value if set, zero value otherwise.
func (o *ListAppsResponseDataItemsRelationships) GetDeployment() DeploymentRelationship {
	if o == nil || o.Deployment == nil {
		var ret DeploymentRelationship
		return ret
	}
	return *o.Deployment
}

// GetDeploymentOk returns a tuple with the Deployment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppsResponseDataItemsRelationships) GetDeploymentOk() (*DeploymentRelationship, bool) {
	if o == nil || o.Deployment == nil {
		return nil, false
	}
	return o.Deployment, true
}

// HasDeployment returns a boolean if a field has been set.
func (o *ListAppsResponseDataItemsRelationships) HasDeployment() bool {
	return o != nil && o.Deployment != nil
}

// SetDeployment gets a reference to the given DeploymentRelationship and assigns it to the Deployment field.
func (o *ListAppsResponseDataItemsRelationships) SetDeployment(v DeploymentRelationship) {
	o.Deployment = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListAppsResponseDataItemsRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Deployment != nil {
		toSerialize["deployment"] = o.Deployment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListAppsResponseDataItemsRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Deployment *DeploymentRelationship `json:"deployment,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"deployment"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Deployment != nil && all.Deployment.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Deployment = all.Deployment

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
