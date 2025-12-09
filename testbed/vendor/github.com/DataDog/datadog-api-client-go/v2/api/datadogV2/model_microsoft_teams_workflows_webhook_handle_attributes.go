// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MicrosoftTeamsWorkflowsWebhookHandleAttributes Workflows Webhook handle attributes.
type MicrosoftTeamsWorkflowsWebhookHandleAttributes struct {
	// Workflows Webhook handle name.
	Name *string `json:"name,omitempty"`
	// Workflows Webhook URL.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMicrosoftTeamsWorkflowsWebhookHandleAttributes instantiates a new MicrosoftTeamsWorkflowsWebhookHandleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMicrosoftTeamsWorkflowsWebhookHandleAttributes() *MicrosoftTeamsWorkflowsWebhookHandleAttributes {
	this := MicrosoftTeamsWorkflowsWebhookHandleAttributes{}
	return &this
}

// NewMicrosoftTeamsWorkflowsWebhookHandleAttributesWithDefaults instantiates a new MicrosoftTeamsWorkflowsWebhookHandleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMicrosoftTeamsWorkflowsWebhookHandleAttributesWithDefaults() *MicrosoftTeamsWorkflowsWebhookHandleAttributes {
	this := MicrosoftTeamsWorkflowsWebhookHandleAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MicrosoftTeamsWorkflowsWebhookHandleAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsWorkflowsWebhookHandleAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftTeamsWorkflowsWebhookHandleAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MicrosoftTeamsWorkflowsWebhookHandleAttributes) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *MicrosoftTeamsWorkflowsWebhookHandleAttributes) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftTeamsWorkflowsWebhookHandleAttributes) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *MicrosoftTeamsWorkflowsWebhookHandleAttributes) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *MicrosoftTeamsWorkflowsWebhookHandleAttributes) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MicrosoftTeamsWorkflowsWebhookHandleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MicrosoftTeamsWorkflowsWebhookHandleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name *string `json:"name,omitempty"`
		Url  *string `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "url"})
	} else {
		return err
	}
	o.Name = all.Name
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
