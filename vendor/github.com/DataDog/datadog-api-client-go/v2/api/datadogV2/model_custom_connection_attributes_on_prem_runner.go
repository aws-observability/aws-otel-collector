// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomConnectionAttributesOnPremRunner Information about the Private Action Runner used by the custom connection, if the custom connection is associated with a Private Action Runner.
type CustomConnectionAttributesOnPremRunner struct {
	// The Private Action Runner ID.
	Id *string `json:"id,omitempty"`
	// The URL of the Private Action Runner.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomConnectionAttributesOnPremRunner instantiates a new CustomConnectionAttributesOnPremRunner object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomConnectionAttributesOnPremRunner() *CustomConnectionAttributesOnPremRunner {
	this := CustomConnectionAttributesOnPremRunner{}
	return &this
}

// NewCustomConnectionAttributesOnPremRunnerWithDefaults instantiates a new CustomConnectionAttributesOnPremRunner object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomConnectionAttributesOnPremRunnerWithDefaults() *CustomConnectionAttributesOnPremRunner {
	this := CustomConnectionAttributesOnPremRunner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomConnectionAttributesOnPremRunner) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomConnectionAttributesOnPremRunner) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomConnectionAttributesOnPremRunner) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomConnectionAttributesOnPremRunner) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CustomConnectionAttributesOnPremRunner) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomConnectionAttributesOnPremRunner) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CustomConnectionAttributesOnPremRunner) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CustomConnectionAttributesOnPremRunner) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomConnectionAttributesOnPremRunner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
func (o *CustomConnectionAttributesOnPremRunner) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id  *string `json:"id,omitempty"`
		Url *string `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "url"})
	} else {
		return err
	}
	o.Id = all.Id
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
