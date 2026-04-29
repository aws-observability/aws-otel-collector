// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomDestinationForwardDestinationSplunk The Splunk HTTP Event Collector (HEC) destination.
type CustomDestinationForwardDestinationSplunk struct {
	// Access token of the Splunk HTTP Event Collector. This field is not returned by the API.
	AccessToken string `json:"access_token"`
	// The destination for which logs will be forwarded to.
	// Must have HTTPS scheme and forwarding back to Datadog is not allowed.
	Endpoint string `json:"endpoint"`
	// The Splunk sourcetype for the events sent to this Splunk destination.
	//
	// If the field is absent from the request and no sourcetype has been previously set on this destination, the default sourcetype `_json` is used.
	// On update, if the field is absent from the request but a sourcetype was previously set, the previous value is kept.
	// If set to `null`, the sourcetype field is omitted from the forwarded event entirely.
	// Otherwise, the provided string value is used as the sourcetype.
	Sourcetype datadog.NullableString `json:"sourcetype,omitempty"`
	// Type of the Splunk HTTP Event Collector (HEC) destination.
	Type CustomDestinationForwardDestinationSplunkType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomDestinationForwardDestinationSplunk instantiates a new CustomDestinationForwardDestinationSplunk object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomDestinationForwardDestinationSplunk(accessToken string, endpoint string, typeVar CustomDestinationForwardDestinationSplunkType) *CustomDestinationForwardDestinationSplunk {
	this := CustomDestinationForwardDestinationSplunk{}
	this.AccessToken = accessToken
	this.Endpoint = endpoint
	this.Type = typeVar
	return &this
}

// NewCustomDestinationForwardDestinationSplunkWithDefaults instantiates a new CustomDestinationForwardDestinationSplunk object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomDestinationForwardDestinationSplunkWithDefaults() *CustomDestinationForwardDestinationSplunk {
	this := CustomDestinationForwardDestinationSplunk{}
	var typeVar CustomDestinationForwardDestinationSplunkType = CUSTOMDESTINATIONFORWARDDESTINATIONSPLUNKTYPE_SPLUNK_HEC
	this.Type = typeVar
	return &this
}

// GetAccessToken returns the AccessToken field value.
func (o *CustomDestinationForwardDestinationSplunk) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationForwardDestinationSplunk) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value.
func (o *CustomDestinationForwardDestinationSplunk) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetEndpoint returns the Endpoint field value.
func (o *CustomDestinationForwardDestinationSplunk) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationForwardDestinationSplunk) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value.
func (o *CustomDestinationForwardDestinationSplunk) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetSourcetype returns the Sourcetype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomDestinationForwardDestinationSplunk) GetSourcetype() string {
	if o == nil || o.Sourcetype.Get() == nil {
		var ret string
		return ret
	}
	return *o.Sourcetype.Get()
}

// GetSourcetypeOk returns a tuple with the Sourcetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CustomDestinationForwardDestinationSplunk) GetSourcetypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sourcetype.Get(), o.Sourcetype.IsSet()
}

// HasSourcetype returns a boolean if a field has been set.
func (o *CustomDestinationForwardDestinationSplunk) HasSourcetype() bool {
	return o != nil && o.Sourcetype.IsSet()
}

// SetSourcetype gets a reference to the given datadog.NullableString and assigns it to the Sourcetype field.
func (o *CustomDestinationForwardDestinationSplunk) SetSourcetype(v string) {
	o.Sourcetype.Set(&v)
}

// SetSourcetypeNil sets the value for Sourcetype to be an explicit nil.
func (o *CustomDestinationForwardDestinationSplunk) SetSourcetypeNil() {
	o.Sourcetype.Set(nil)
}

// UnsetSourcetype ensures that no value is present for Sourcetype, not even an explicit nil.
func (o *CustomDestinationForwardDestinationSplunk) UnsetSourcetype() {
	o.Sourcetype.Unset()
}

// GetType returns the Type field value.
func (o *CustomDestinationForwardDestinationSplunk) GetType() CustomDestinationForwardDestinationSplunkType {
	if o == nil {
		var ret CustomDestinationForwardDestinationSplunkType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CustomDestinationForwardDestinationSplunk) GetTypeOk() (*CustomDestinationForwardDestinationSplunkType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CustomDestinationForwardDestinationSplunk) SetType(v CustomDestinationForwardDestinationSplunkType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomDestinationForwardDestinationSplunk) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["endpoint"] = o.Endpoint
	if o.Sourcetype.IsSet() {
		toSerialize["sourcetype"] = o.Sourcetype.Get()
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomDestinationForwardDestinationSplunk) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AccessToken *string                                        `json:"access_token"`
		Endpoint    *string                                        `json:"endpoint"`
		Sourcetype  datadog.NullableString                         `json:"sourcetype,omitempty"`
		Type        *CustomDestinationForwardDestinationSplunkType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AccessToken == nil {
		return fmt.Errorf("required field access_token missing")
	}
	if all.Endpoint == nil {
		return fmt.Errorf("required field endpoint missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"access_token", "endpoint", "sourcetype", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AccessToken = *all.AccessToken
	o.Endpoint = *all.Endpoint
	o.Sourcetype = all.Sourcetype
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
