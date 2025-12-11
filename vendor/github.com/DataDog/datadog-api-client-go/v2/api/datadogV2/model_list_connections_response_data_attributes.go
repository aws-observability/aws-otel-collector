// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListConnectionsResponseDataAttributes
type ListConnectionsResponseDataAttributes struct {
	//
	Connections []ListConnectionsResponseDataAttributesConnectionsItems `json:"connections,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListConnectionsResponseDataAttributes instantiates a new ListConnectionsResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListConnectionsResponseDataAttributes() *ListConnectionsResponseDataAttributes {
	this := ListConnectionsResponseDataAttributes{}
	return &this
}

// NewListConnectionsResponseDataAttributesWithDefaults instantiates a new ListConnectionsResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListConnectionsResponseDataAttributesWithDefaults() *ListConnectionsResponseDataAttributes {
	this := ListConnectionsResponseDataAttributes{}
	return &this
}

// GetConnections returns the Connections field value if set, zero value otherwise.
func (o *ListConnectionsResponseDataAttributes) GetConnections() []ListConnectionsResponseDataAttributesConnectionsItems {
	if o == nil || o.Connections == nil {
		var ret []ListConnectionsResponseDataAttributesConnectionsItems
		return ret
	}
	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectionsResponseDataAttributes) GetConnectionsOk() (*[]ListConnectionsResponseDataAttributesConnectionsItems, bool) {
	if o == nil || o.Connections == nil {
		return nil, false
	}
	return &o.Connections, true
}

// HasConnections returns a boolean if a field has been set.
func (o *ListConnectionsResponseDataAttributes) HasConnections() bool {
	return o != nil && o.Connections != nil
}

// SetConnections gets a reference to the given []ListConnectionsResponseDataAttributesConnectionsItems and assigns it to the Connections field.
func (o *ListConnectionsResponseDataAttributes) SetConnections(v []ListConnectionsResponseDataAttributesConnectionsItems) {
	o.Connections = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListConnectionsResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListConnectionsResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Connections []ListConnectionsResponseDataAttributesConnectionsItems `json:"connections,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connections"})
	} else {
		return err
	}
	o.Connections = all.Connections

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
