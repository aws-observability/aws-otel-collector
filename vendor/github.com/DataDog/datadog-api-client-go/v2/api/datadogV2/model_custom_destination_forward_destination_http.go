// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationForwardDestinationHttp The HTTP destination.
type CustomDestinationForwardDestinationHttp struct {
	// Authentication method of the HTTP requests.
	Auth CustomDestinationHttpDestinationAuth `json:"auth"`
	// The destination for which logs will be forwarded to.
	// Must have HTTPS scheme and forwarding back to Datadog is not allowed.
	Endpoint string `json:"endpoint"`
	// Type of the HTTP destination.
	Type CustomDestinationForwardDestinationHttpType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationForwardDestinationHttp instantiates a new CustomDestinationForwardDestinationHttp object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationForwardDestinationHttp(auth CustomDestinationHttpDestinationAuth, endpoint string, typeVar CustomDestinationForwardDestinationHttpType) *CustomDestinationForwardDestinationHttp {
	this := CustomDestinationForwardDestinationHttp{}
	this.Auth = auth
	this.Endpoint = endpoint
	this.Type = typeVar
	return &this
}

// NewCustomDestinationForwardDestinationHttpWithDefaults instantiates a new CustomDestinationForwardDestinationHttp object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationForwardDestinationHttpWithDefaults() *CustomDestinationForwardDestinationHttp {
	this := CustomDestinationForwardDestinationHttp{}
	var typeVar CustomDestinationForwardDestinationHttpType = CUSTOMDESTINATIONFORWARDDESTINATIONHTTPTYPE_HTTP
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value.
func (o *CustomDestinationForwardDestinationHttp) GetAuth() CustomDestinationHttpDestinationAuth {
	if o == nil {
		var ret CustomDestinationHttpDestinationAuth
		return ret
	}
	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationForwardDestinationHttp) GetAuthOk() (*CustomDestinationHttpDestinationAuth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value.
func (o *CustomDestinationForwardDestinationHttp) SetAuth(v CustomDestinationHttpDestinationAuth) {
	o.Auth = v
}

// GetEndpoint returns the Endpoint field value.
func (o *CustomDestinationForwardDestinationHttp) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationForwardDestinationHttp) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value.
func (o *CustomDestinationForwardDestinationHttp) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetType returns the Type field value.
func (o *CustomDestinationForwardDestinationHttp) GetType() CustomDestinationForwardDestinationHttpType {
	if o == nil {
		var ret CustomDestinationForwardDestinationHttpType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationForwardDestinationHttp) GetTypeOk() (*CustomDestinationForwardDestinationHttpType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationForwardDestinationHttp) SetType(v CustomDestinationForwardDestinationHttpType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationForwardDestinationHttp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth"] = o.Auth
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationForwardDestinationHttp) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth     *CustomDestinationHttpDestinationAuth        `json:"auth"`
		Endpoint *string                                      `json:"endpoint"`
		Type     *CustomDestinationForwardDestinationHttpType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Auth == nil {
		return fmt.Errorf("required field auth missing")
	}
	if all.Endpoint == nil {
		return fmt.Errorf("required field endpoint missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "endpoint", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Auth = *all.Auth
	o.Endpoint = *all.Endpoint
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
