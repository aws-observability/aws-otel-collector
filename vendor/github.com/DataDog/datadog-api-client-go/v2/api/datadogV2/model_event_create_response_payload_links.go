// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EventCreateResponsePayloadLinks Links to the event.
type EventCreateResponsePayloadLinks struct {
	// The URL of the event. This link is only functional when using the default subdomain.
	Self *string `json:"self,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEventCreateResponsePayloadLinks instantiates a new EventCreateResponsePayloadLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEventCreateResponsePayloadLinks() *EventCreateResponsePayloadLinks {
	this := EventCreateResponsePayloadLinks{}
	return &this
}

// NewEventCreateResponsePayloadLinksWithDefaults instantiates a new EventCreateResponsePayloadLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEventCreateResponsePayloadLinksWithDefaults() *EventCreateResponsePayloadLinks {
	this := EventCreateResponsePayloadLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *EventCreateResponsePayloadLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCreateResponsePayloadLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *EventCreateResponsePayloadLinks) HasSelf() bool {
	return o != nil && o.Self != nil
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *EventCreateResponsePayloadLinks) SetSelf(v string) {
	o.Self = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EventCreateResponsePayloadLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EventCreateResponsePayloadLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Self *string `json:"self,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"self"})
	} else {
		return err
	}
	o.Self = all.Self

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
