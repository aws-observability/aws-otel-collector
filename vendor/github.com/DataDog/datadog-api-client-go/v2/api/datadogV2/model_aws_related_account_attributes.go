// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AWSRelatedAccountAttributes Attributes for an AWS related account.
type AWSRelatedAccountAttributes struct {
	// Whether or not the AWS account has a Datadog integration.
	HasDatadogIntegration *bool `json:"has_datadog_integration,omitempty"`
	// The name of the AWS account.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewAWSRelatedAccountAttributes instantiates a new AWSRelatedAccountAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAWSRelatedAccountAttributes() *AWSRelatedAccountAttributes {
	this := AWSRelatedAccountAttributes{}
	return &this
}

// NewAWSRelatedAccountAttributesWithDefaults instantiates a new AWSRelatedAccountAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAWSRelatedAccountAttributesWithDefaults() *AWSRelatedAccountAttributes {
	this := AWSRelatedAccountAttributes{}
	return &this
}

// GetHasDatadogIntegration returns the HasDatadogIntegration field value if set, zero value otherwise.
func (o *AWSRelatedAccountAttributes) GetHasDatadogIntegration() bool {
	if o == nil || o.HasDatadogIntegration == nil {
		var ret bool
		return ret
	}
	return *o.HasDatadogIntegration
}

// GetHasDatadogIntegrationOk returns a tuple with the HasDatadogIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSRelatedAccountAttributes) GetHasDatadogIntegrationOk() (*bool, bool) {
	if o == nil || o.HasDatadogIntegration == nil {
		return nil, false
	}
	return o.HasDatadogIntegration, true
}

// HasHasDatadogIntegration returns a boolean if a field has been set.
func (o *AWSRelatedAccountAttributes) HasHasDatadogIntegration() bool {
	return o != nil && o.HasDatadogIntegration != nil
}

// SetHasDatadogIntegration gets a reference to the given bool and assigns it to the HasDatadogIntegration field.
func (o *AWSRelatedAccountAttributes) SetHasDatadogIntegration(v bool) {
	o.HasDatadogIntegration = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AWSRelatedAccountAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AWSRelatedAccountAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AWSRelatedAccountAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AWSRelatedAccountAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o AWSRelatedAccountAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.HasDatadogIntegration != nil {
		toSerialize["has_datadog_integration"] = o.HasDatadogIntegration
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AWSRelatedAccountAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HasDatadogIntegration *bool   `json:"has_datadog_integration,omitempty"`
		Name                  *string `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"has_datadog_integration", "name"})
	} else {
		return err
	}
	o.HasDatadogIntegration = all.HasDatadogIntegration
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
