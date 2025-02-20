// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseHttpDestinationAuthCustomHeader Custom header access authentication.
type CustomDestinationResponseHttpDestinationAuthCustomHeader struct {
	// The header name of the authentication.
	HeaderName string `json:"header_name"`
	// Type of the custom header access authentication.
	Type CustomDestinationResponseHttpDestinationAuthCustomHeaderType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationResponseHttpDestinationAuthCustomHeader instantiates a new CustomDestinationResponseHttpDestinationAuthCustomHeader object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationResponseHttpDestinationAuthCustomHeader(headerName string, typeVar CustomDestinationResponseHttpDestinationAuthCustomHeaderType) *CustomDestinationResponseHttpDestinationAuthCustomHeader {
	this := CustomDestinationResponseHttpDestinationAuthCustomHeader{}
	this.HeaderName = headerName
	this.Type = typeVar
	return &this
}

// NewCustomDestinationResponseHttpDestinationAuthCustomHeaderWithDefaults instantiates a new CustomDestinationResponseHttpDestinationAuthCustomHeader object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationResponseHttpDestinationAuthCustomHeaderWithDefaults() *CustomDestinationResponseHttpDestinationAuthCustomHeader {
	this := CustomDestinationResponseHttpDestinationAuthCustomHeader{}
	var typeVar CustomDestinationResponseHttpDestinationAuthCustomHeaderType = CUSTOMDESTINATIONRESPONSEHTTPDESTINATIONAUTHCUSTOMHEADERTYPE_CUSTOM_HEADER
	this.Type = typeVar
	return &this
}

// GetHeaderName returns the HeaderName field value.
func (o *CustomDestinationResponseHttpDestinationAuthCustomHeader) GetHeaderName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.HeaderName
}

// GetHeaderNameOk returns a tuple with the HeaderName field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseHttpDestinationAuthCustomHeader) GetHeaderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HeaderName, true
}

// SetHeaderName sets field value.
func (o *CustomDestinationResponseHttpDestinationAuthCustomHeader) SetHeaderName(v string) {
	o.HeaderName = v
}

// GetType returns the Type field value.
func (o *CustomDestinationResponseHttpDestinationAuthCustomHeader) GetType() CustomDestinationResponseHttpDestinationAuthCustomHeaderType {
	if o == nil {
		var ret CustomDestinationResponseHttpDestinationAuthCustomHeaderType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseHttpDestinationAuthCustomHeader) GetTypeOk() (*CustomDestinationResponseHttpDestinationAuthCustomHeaderType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationResponseHttpDestinationAuthCustomHeader) SetType(v CustomDestinationResponseHttpDestinationAuthCustomHeaderType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationResponseHttpDestinationAuthCustomHeader) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["header_name"] = o.HeaderName
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationResponseHttpDestinationAuthCustomHeader) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		HeaderName *string                                                       `json:"header_name"`
		Type       *CustomDestinationResponseHttpDestinationAuthCustomHeaderType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.HeaderName == nil {
		return fmt.Errorf("required field header_name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"header_name", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.HeaderName = *all.HeaderName
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
