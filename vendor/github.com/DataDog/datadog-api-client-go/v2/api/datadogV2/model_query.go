// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Query The definition of `Query` object.
type Query struct {
	// The `Query` `events`.
	Events []AppBuilderEvent `json:"events,omitempty"`
	// The `Query` `id`.
	Id string `json:"id"`
	// The `Query` `name`.
	Name string `json:"name"`
	// The `Query` `properties`.
	Properties interface{} `json:"properties,omitempty"`
	// The definition of `QueryType` object.
	Type QueryType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewQuery instantiates a new Query object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQuery(id string, name string, typeVar QueryType) *Query {
	this := Query{}
	this.Id = id
	this.Name = name
	this.Type = typeVar
	return &this
}

// NewQueryWithDefaults instantiates a new Query object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryWithDefaults() *Query {
	this := Query{}
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *Query) GetEvents() []AppBuilderEvent {
	if o == nil || o.Events == nil {
		var ret []AppBuilderEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetEventsOk() (*[]AppBuilderEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *Query) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given []AppBuilderEvent and assigns it to the Events field.
func (o *Query) SetEvents(v []AppBuilderEvent) {
	o.Events = v
}

// GetId returns the Id field value.
func (o *Query) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Query) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *Query) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *Query) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Query) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *Query) SetName(v string) {
	o.Name = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Query) GetProperties() interface{} {
	if o == nil || o.Properties == nil {
		var ret interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Query) GetPropertiesOk() (*interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return &o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Query) HasProperties() bool {
	return o != nil && o.Properties != nil
}

// SetProperties gets a reference to the given interface{} and assigns it to the Properties field.
func (o *Query) SetProperties(v interface{}) {
	o.Properties = v
}

// GetType returns the Type field value.
func (o *Query) GetType() QueryType {
	if o == nil {
		var ret QueryType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Query) GetTypeOk() (*QueryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Query) SetType(v QueryType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Query) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Query) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Events     []AppBuilderEvent `json:"events,omitempty"`
		Id         *string           `json:"id"`
		Name       *string           `json:"name"`
		Properties interface{}       `json:"properties,omitempty"`
		Type       *QueryType        `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"events", "id", "name", "properties", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Events = all.Events
	o.Id = *all.Id
	o.Name = *all.Name
	o.Properties = all.Properties
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
