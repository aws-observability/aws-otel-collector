// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WebhooksAuthMethodAttributes Attributes of a webhooks auth method.
type WebhooksAuthMethodAttributes struct {
	// Authentication protocol used by the auth method.
	Protocol *WebhooksAuthMethodProtocol `json:"protocol,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWebhooksAuthMethodAttributes instantiates a new WebhooksAuthMethodAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWebhooksAuthMethodAttributes() *WebhooksAuthMethodAttributes {
	this := WebhooksAuthMethodAttributes{}
	return &this
}

// NewWebhooksAuthMethodAttributesWithDefaults instantiates a new WebhooksAuthMethodAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWebhooksAuthMethodAttributesWithDefaults() *WebhooksAuthMethodAttributes {
	this := WebhooksAuthMethodAttributes{}
	return &this
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *WebhooksAuthMethodAttributes) GetProtocol() WebhooksAuthMethodProtocol {
	if o == nil || o.Protocol == nil {
		var ret WebhooksAuthMethodProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhooksAuthMethodAttributes) GetProtocolOk() (*WebhooksAuthMethodProtocol, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *WebhooksAuthMethodAttributes) HasProtocol() bool {
	return o != nil && o.Protocol != nil
}

// SetProtocol gets a reference to the given WebhooksAuthMethodProtocol and assigns it to the Protocol field.
func (o *WebhooksAuthMethodAttributes) SetProtocol(v WebhooksAuthMethodProtocol) {
	o.Protocol = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WebhooksAuthMethodAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WebhooksAuthMethodAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Protocol *WebhooksAuthMethodProtocol `json:"protocol,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"protocol"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Protocol != nil && !all.Protocol.IsValid() {
		hasInvalidField = true
	} else {
		o.Protocol = all.Protocol
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
