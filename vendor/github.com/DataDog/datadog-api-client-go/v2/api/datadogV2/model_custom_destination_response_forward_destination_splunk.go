// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationResponseForwardDestinationSplunk The Splunk HTTP Event Collector (HEC) destination.
type CustomDestinationResponseForwardDestinationSplunk struct {
	// The destination for which logs will be forwarded to.
	// Must have HTTPS scheme and forwarding back to Datadog is not allowed.
	Endpoint string `json:"endpoint"`
	// Type of the Splunk HTTP Event Collector (HEC) destination.
	Type CustomDestinationResponseForwardDestinationSplunkType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationResponseForwardDestinationSplunk instantiates a new CustomDestinationResponseForwardDestinationSplunk object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationResponseForwardDestinationSplunk(endpoint string, typeVar CustomDestinationResponseForwardDestinationSplunkType) *CustomDestinationResponseForwardDestinationSplunk {
	this := CustomDestinationResponseForwardDestinationSplunk{}
	this.Endpoint = endpoint
	this.Type = typeVar
	return &this
}

// NewCustomDestinationResponseForwardDestinationSplunkWithDefaults instantiates a new CustomDestinationResponseForwardDestinationSplunk object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationResponseForwardDestinationSplunkWithDefaults() *CustomDestinationResponseForwardDestinationSplunk {
	this := CustomDestinationResponseForwardDestinationSplunk{}
	var typeVar CustomDestinationResponseForwardDestinationSplunkType = CUSTOMDESTINATIONRESPONSEFORWARDDESTINATIONSPLUNKTYPE_SPLUNK_HEC
	this.Type = typeVar
	return &this
}

// GetEndpoint returns the Endpoint field value.
func (o *CustomDestinationResponseForwardDestinationSplunk) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationSplunk) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value.
func (o *CustomDestinationResponseForwardDestinationSplunk) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetType returns the Type field value.
func (o *CustomDestinationResponseForwardDestinationSplunk) GetType() CustomDestinationResponseForwardDestinationSplunkType {
	if o == nil {
		var ret CustomDestinationResponseForwardDestinationSplunkType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationResponseForwardDestinationSplunk) GetTypeOk() (*CustomDestinationResponseForwardDestinationSplunkType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationResponseForwardDestinationSplunk) SetType(v CustomDestinationResponseForwardDestinationSplunkType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationResponseForwardDestinationSplunk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationResponseForwardDestinationSplunk) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Endpoint *string                                                `json:"endpoint"`
		Type     *CustomDestinationResponseForwardDestinationSplunkType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Endpoint == nil {
		return fmt.Errorf("required field endpoint missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"endpoint", "type"})
	} else {
		return err
	}

	hasInvalidField := false
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
