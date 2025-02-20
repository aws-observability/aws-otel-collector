// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SlackIntegrationChannel The Slack channel configuration.
type SlackIntegrationChannel struct {
	// Configuration options for what is shown in an alert event message.
	Display *SlackIntegrationChannelDisplay `json:"display,omitempty"`
	// Your channel name.
	Name *string `json:"name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSlackIntegrationChannel instantiates a new SlackIntegrationChannel object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSlackIntegrationChannel() *SlackIntegrationChannel {
	this := SlackIntegrationChannel{}
	return &this
}

// NewSlackIntegrationChannelWithDefaults instantiates a new SlackIntegrationChannel object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSlackIntegrationChannelWithDefaults() *SlackIntegrationChannel {
	this := SlackIntegrationChannel{}
	return &this
}

// GetDisplay returns the Display field value if set, zero value otherwise.
func (o *SlackIntegrationChannel) GetDisplay() SlackIntegrationChannelDisplay {
	if o == nil || o.Display == nil {
		var ret SlackIntegrationChannelDisplay
		return ret
	}
	return *o.Display
}

// GetDisplayOk returns a tuple with the Display field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegrationChannel) GetDisplayOk() (*SlackIntegrationChannelDisplay, bool) {
	if o == nil || o.Display == nil {
		return nil, false
	}
	return o.Display, true
}

// HasDisplay returns a boolean if a field has been set.
func (o *SlackIntegrationChannel) HasDisplay() bool {
	return o != nil && o.Display != nil
}

// SetDisplay gets a reference to the given SlackIntegrationChannelDisplay and assigns it to the Display field.
func (o *SlackIntegrationChannel) SetDisplay(v SlackIntegrationChannelDisplay) {
	o.Display = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SlackIntegrationChannel) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SlackIntegrationChannel) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SlackIntegrationChannel) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SlackIntegrationChannel) SetName(v string) {
	o.Name = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SlackIntegrationChannel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Display != nil {
		toSerialize["display"] = o.Display
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SlackIntegrationChannel) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Display *SlackIntegrationChannelDisplay `json:"display,omitempty"`
		Name    *string                         `json:"name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display", "name"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Display != nil && all.Display.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Display = all.Display
	o.Name = all.Name

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
