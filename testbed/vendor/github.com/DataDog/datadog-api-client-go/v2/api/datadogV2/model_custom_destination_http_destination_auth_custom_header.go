// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationHttpDestinationAuthCustomHeader Custom header access authentication.
type CustomDestinationHttpDestinationAuthCustomHeader struct {
	// The header name of the authentication.
	HeaderName string `json:"header_name"`
	// The header value of the authentication. This field is not returned by the API.
	HeaderValue string `json:"header_value"`
	// Type of the custom header access authentication.
	Type CustomDestinationHttpDestinationAuthCustomHeaderType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationHttpDestinationAuthCustomHeader instantiates a new CustomDestinationHttpDestinationAuthCustomHeader object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationHttpDestinationAuthCustomHeader(headerName string, headerValue string, typeVar CustomDestinationHttpDestinationAuthCustomHeaderType) *CustomDestinationHttpDestinationAuthCustomHeader {
	this := CustomDestinationHttpDestinationAuthCustomHeader{}
	this.HeaderName = headerName
	this.HeaderValue = headerValue
	this.Type = typeVar
	return &this
}

// NewCustomDestinationHttpDestinationAuthCustomHeaderWithDefaults instantiates a new CustomDestinationHttpDestinationAuthCustomHeader object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationHttpDestinationAuthCustomHeaderWithDefaults() *CustomDestinationHttpDestinationAuthCustomHeader {
	this := CustomDestinationHttpDestinationAuthCustomHeader{}
	var typeVar CustomDestinationHttpDestinationAuthCustomHeaderType = CUSTOMDESTINATIONHTTPDESTINATIONAUTHCUSTOMHEADERTYPE_CUSTOM_HEADER
	this.Type = typeVar
	return &this
}

// GetHeaderName returns the HeaderName field value.
func (o *CustomDestinationHttpDestinationAuthCustomHeader) GetHeaderName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.HeaderName
}

// GetHeaderNameOk returns a tuple with the HeaderName field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationHttpDestinationAuthCustomHeader) GetHeaderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeaderName, true
}

// SetHeaderName sets field value.
func (o *CustomDestinationHttpDestinationAuthCustomHeader) SetHeaderName(v string) {
	o.HeaderName = v
}

// GetHeaderValue returns the HeaderValue field value.
func (o *CustomDestinationHttpDestinationAuthCustomHeader) GetHeaderValue() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.HeaderValue
}

// GetHeaderValueOk returns a tuple with the HeaderValue field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationHttpDestinationAuthCustomHeader) GetHeaderValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeaderValue, true
}

// SetHeaderValue sets field value.
func (o *CustomDestinationHttpDestinationAuthCustomHeader) SetHeaderValue(v string) {
	o.HeaderValue = v
}

// GetType returns the Type field value.
func (o *CustomDestinationHttpDestinationAuthCustomHeader) GetType() CustomDestinationHttpDestinationAuthCustomHeaderType {
	if o == nil {
		var ret CustomDestinationHttpDestinationAuthCustomHeaderType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationHttpDestinationAuthCustomHeader) GetTypeOk() (*CustomDestinationHttpDestinationAuthCustomHeaderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationHttpDestinationAuthCustomHeader) SetType(v CustomDestinationHttpDestinationAuthCustomHeaderType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationHttpDestinationAuthCustomHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["header_name"] = o.HeaderName
	toSerialize["header_value"] = o.HeaderValue
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationHttpDestinationAuthCustomHeader) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HeaderName  *string                                               `json:"header_name"`
		HeaderValue *string                                               `json:"header_value"`
		Type        *CustomDestinationHttpDestinationAuthCustomHeaderType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HeaderName == nil {
		return fmt.Errorf("required field header_name missing")
	}
	if all.HeaderValue == nil {
		return fmt.Errorf("required field header_value missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"header_name", "header_value", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.HeaderName = *all.HeaderName
	o.HeaderValue = *all.HeaderValue
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
