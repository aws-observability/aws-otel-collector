// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSResourcesConfig AWS Resources Collection config.
type AWSResourcesConfig struct {
	// Enable Cloud Security Management to scan AWS resources for vulnerabilities, misconfigurations, identity risks, and compliance violations. Defaults to `false`. Requires `extended_collection` to be set to `true`.
	CloudSecurityPostureManagementCollection *bool `json:"cloud_security_posture_management_collection,omitempty"`
	// Whether Datadog collects additional attributes and configuration information about the resources in your AWS account. Defaults to `true`. Required for `cloud_security_posture_management_collection`.
	ExtendedCollection *bool `json:"extended_collection,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAWSResourcesConfig instantiates a new AWSResourcesConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSResourcesConfig() *AWSResourcesConfig {
	this := AWSResourcesConfig{}
	return &this
}

// NewAWSResourcesConfigWithDefaults instantiates a new AWSResourcesConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSResourcesConfigWithDefaults() *AWSResourcesConfig {
	this := AWSResourcesConfig{}
	return &this
}

// GetCloudSecurityPostureManagementCollection returns the CloudSecurityPostureManagementCollection field value if set, zero value otherwise.
func (o *AWSResourcesConfig) GetCloudSecurityPostureManagementCollection() bool {
	if o == nil || o.CloudSecurityPostureManagementCollection == nil {
		var ret bool
		return ret
	}
	return *o.CloudSecurityPostureManagementCollection
}

// GetCloudSecurityPostureManagementCollectionOk returns a tuple with the CloudSecurityPostureManagementCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSResourcesConfig) GetCloudSecurityPostureManagementCollectionOk() (*bool, bool) {
	if o == nil || o.CloudSecurityPostureManagementCollection == nil {
		return nil, false
	}
	return o.CloudSecurityPostureManagementCollection, true
}

// HasCloudSecurityPostureManagementCollection returns a boolean if a field has been set.
func (o *AWSResourcesConfig) HasCloudSecurityPostureManagementCollection() bool {
	return o != nil && o.CloudSecurityPostureManagementCollection != nil
}

// SetCloudSecurityPostureManagementCollection gets a reference to the given bool and assigns it to the CloudSecurityPostureManagementCollection field.
func (o *AWSResourcesConfig) SetCloudSecurityPostureManagementCollection(v bool) {
	o.CloudSecurityPostureManagementCollection = &v
}

// GetExtendedCollection returns the ExtendedCollection field value if set, zero value otherwise.
func (o *AWSResourcesConfig) GetExtendedCollection() bool {
	if o == nil || o.ExtendedCollection == nil {
		var ret bool
		return ret
	}
	return *o.ExtendedCollection
}

// GetExtendedCollectionOk returns a tuple with the ExtendedCollection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSResourcesConfig) GetExtendedCollectionOk() (*bool, bool) {
	if o == nil || o.ExtendedCollection == nil {
		return nil, false
	}
	return o.ExtendedCollection, true
}

// HasExtendedCollection returns a boolean if a field has been set.
func (o *AWSResourcesConfig) HasExtendedCollection() bool {
	return o != nil && o.ExtendedCollection != nil
}

// SetExtendedCollection gets a reference to the given bool and assigns it to the ExtendedCollection field.
func (o *AWSResourcesConfig) SetExtendedCollection(v bool) {
	o.ExtendedCollection = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSResourcesConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CloudSecurityPostureManagementCollection != nil {
		toSerialize["cloud_security_posture_management_collection"] = o.CloudSecurityPostureManagementCollection
	}
	if o.ExtendedCollection != nil {
		toSerialize["extended_collection"] = o.ExtendedCollection
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSResourcesConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CloudSecurityPostureManagementCollection *bool `json:"cloud_security_posture_management_collection,omitempty"`
		ExtendedCollection                       *bool `json:"extended_collection,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cloud_security_posture_management_collection", "extended_collection"})
	} else {
		return err
	}
	o.CloudSecurityPostureManagementCollection = all.CloudSecurityPostureManagementCollection
	o.ExtendedCollection = all.ExtendedCollection

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
