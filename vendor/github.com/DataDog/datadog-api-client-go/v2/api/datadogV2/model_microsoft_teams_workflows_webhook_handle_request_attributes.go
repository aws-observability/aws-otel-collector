// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes Workflows Webhook handle attributes.
type MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes struct {
	// Workflows Webhook handle name.
	Name string `json:"name"`
	// Workflows Webhook URL.
	Url string `json:"url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMicrosoftTeamsWorkflowsWebhookHandleRequestAttributes instantiates a new MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMicrosoftTeamsWorkflowsWebhookHandleRequestAttributes(name string, url string) *MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes {
	this := MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes{}
	this.Name = name
	this.Url = url
	return &this
}

// NewMicrosoftTeamsWorkflowsWebhookHandleRequestAttributesWithDefaults instantiates a new MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMicrosoftTeamsWorkflowsWebhookHandleRequestAttributesWithDefaults() *MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes {
	this := MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes{}
	return &this
}

// GetName returns the Name field value.
func (o *MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value.
func (o *MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value.
func (o *MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes) SetUrl(v string) {
	o.Url = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["name"] = o.Name
	toSerialize["url"] = o.Url

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MicrosoftTeamsWorkflowsWebhookHandleRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string `json:"name"`
		Url  *string `json:"url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Url == nil {
		return fmt.Errorf("required field url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "url"})
	} else {
		return err
	}
	o.Name = *all.Name
	o.Url = *all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
