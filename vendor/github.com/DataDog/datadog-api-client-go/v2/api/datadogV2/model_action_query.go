// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionQuery An action query. This query type is used to trigger an action, such as sending a HTTP request.
type ActionQuery struct {
	// Events to listen for downstream of the action query.
	Events []AppBuilderEvent `json:"events,omitempty"`
	// The ID of the action query.
	Id uuid.UUID `json:"id"`
	// A unique identifier for this action query. This name is also used to access the query's result throughout the app.
	Name string `json:"name"`
	// The properties of the action query.
	Properties ActionQueryProperties `json:"properties"`
	// The action query type.
	Type ActionQueryType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActionQuery instantiates a new ActionQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionQuery(id uuid.UUID, name string, properties ActionQueryProperties, typeVar ActionQueryType) *ActionQuery {
	this := ActionQuery{}
	this.Id = id
	this.Name = name
	this.Properties = properties
	this.Type = typeVar
	return &this
}

// NewActionQueryWithDefaults instantiates a new ActionQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionQueryWithDefaults() *ActionQuery {
	this := ActionQuery{}
	var typeVar ActionQueryType = ACTIONQUERYTYPE_ACTION
	this.Type = typeVar
	return &this
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *ActionQuery) GetEvents() []AppBuilderEvent {
	if o == nil || o.Events == nil {
		var ret []AppBuilderEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionQuery) GetEventsOk() (*[]AppBuilderEvent, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return &o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *ActionQuery) HasEvents() bool {
	return o != nil && o.Events != nil
}

// SetEvents gets a reference to the given []AppBuilderEvent and assigns it to the Events field.
func (o *ActionQuery) SetEvents(v []AppBuilderEvent) {
	o.Events = v
}

// GetId returns the Id field value.
func (o *ActionQuery) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ActionQuery) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ActionQuery) SetId(v uuid.UUID) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *ActionQuery) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ActionQuery) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ActionQuery) SetName(v string) {
	o.Name = v
}

// GetProperties returns the Properties field value.
func (o *ActionQuery) GetProperties() ActionQueryProperties {
	if o == nil {
		var ret ActionQueryProperties
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ActionQuery) GetPropertiesOk() (*ActionQueryProperties, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value.
func (o *ActionQuery) SetProperties(v ActionQueryProperties) {
	o.Properties = v
}

// GetType returns the Type field value.
func (o *ActionQuery) GetType() ActionQueryType {
	if o == nil {
		var ret ActionQueryType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActionQuery) GetTypeOk() (*ActionQueryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ActionQuery) SetType(v ActionQueryType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["properties"] = o.Properties
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActionQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Events     []AppBuilderEvent      `json:"events,omitempty"`
		Id         *uuid.UUID             `json:"id"`
		Name       *string                `json:"name"`
		Properties *ActionQueryProperties `json:"properties"`
		Type       *ActionQueryType       `json:"type"`
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
	if all.Properties == nil {
		return fmt.Errorf("required field properties missing")
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
	if all.Properties.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Properties = *all.Properties
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
