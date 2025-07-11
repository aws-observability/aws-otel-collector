// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomConnectionAttributes The custom connection attributes.
type CustomConnectionAttributes struct {
	// The name of the custom connection.
	Name *string `json:"name,omitempty"`
	// Information about the Private Action Runner used by the custom connection, if the custom connection is associated with a Private Action Runner.
	OnPremRunner *CustomConnectionAttributesOnPremRunner `json:"onPremRunner,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomConnectionAttributes instantiates a new CustomConnectionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomConnectionAttributes() *CustomConnectionAttributes {
	this := CustomConnectionAttributes{}
	return &this
}

// NewCustomConnectionAttributesWithDefaults instantiates a new CustomConnectionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomConnectionAttributesWithDefaults() *CustomConnectionAttributes {
	this := CustomConnectionAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CustomConnectionAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomConnectionAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CustomConnectionAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CustomConnectionAttributes) SetName(v string) {
	o.Name = &v
}

// GetOnPremRunner returns the OnPremRunner field value if set, zero value otherwise.
func (o *CustomConnectionAttributes) GetOnPremRunner() CustomConnectionAttributesOnPremRunner {
	if o == nil || o.OnPremRunner == nil {
		var ret CustomConnectionAttributesOnPremRunner
		return ret
	}
	return *o.OnPremRunner
}

// GetOnPremRunnerOk returns a tuple with the OnPremRunner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomConnectionAttributes) GetOnPremRunnerOk() (*CustomConnectionAttributesOnPremRunner, bool) {
	if o == nil || o.OnPremRunner == nil {
		return nil, false
	}
	return o.OnPremRunner, true
}

// HasOnPremRunner returns a boolean if a field has been set.
func (o *CustomConnectionAttributes) HasOnPremRunner() bool {
	return o != nil && o.OnPremRunner != nil
}

// SetOnPremRunner gets a reference to the given CustomConnectionAttributesOnPremRunner and assigns it to the OnPremRunner field.
func (o *CustomConnectionAttributes) SetOnPremRunner(v CustomConnectionAttributesOnPremRunner) {
	o.OnPremRunner = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomConnectionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OnPremRunner != nil {
		toSerialize["onPremRunner"] = o.OnPremRunner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomConnectionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string                                 `json:"name,omitempty"`
		OnPremRunner *CustomConnectionAttributesOnPremRunner `json:"onPremRunner,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "onPremRunner"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	if all.OnPremRunner != nil && all.OnPremRunner.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.OnPremRunner = all.OnPremRunner

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
