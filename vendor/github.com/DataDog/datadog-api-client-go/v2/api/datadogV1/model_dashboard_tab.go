// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DashboardTab Dashboard tab for organizing widgets.
type DashboardTab struct {
	// UUID of the tab.
	Id uuid.UUID `json:"id"`
	// Name of the tab.
	Name string `json:"name"`
	// List of widget IDs belonging to this tab. The backend also accepts positional references in @N format (1-indexed) as a convenience for Terraform and other declarative tools.
	WidgetIds []int64 `json:"widget_ids"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDashboardTab instantiates a new DashboardTab object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDashboardTab(id uuid.UUID, name string, widgetIds []int64) *DashboardTab {
	this := DashboardTab{}
	this.Id = id
	this.Name = name
	this.WidgetIds = widgetIds
	return &this
}

// NewDashboardTabWithDefaults instantiates a new DashboardTab object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDashboardTabWithDefaults() *DashboardTab {
	this := DashboardTab{}
	return &this
}

// GetId returns the Id field value.
func (o *DashboardTab) GetId() uuid.UUID {
	if o == nil {
		var ret uuid.UUID
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DashboardTab) GetIdOk() (*uuid.UUID, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *DashboardTab) SetId(v uuid.UUID) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *DashboardTab) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DashboardTab) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *DashboardTab) SetName(v string) {
	o.Name = v
}

// GetWidgetIds returns the WidgetIds field value.
func (o *DashboardTab) GetWidgetIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.WidgetIds
}

// GetWidgetIdsOk returns a tuple with the WidgetIds field value
// and a boolean to check if the value has been set.
func (o *DashboardTab) GetWidgetIdsOk() (*[]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WidgetIds, true
}

// SetWidgetIds sets field value.
func (o *DashboardTab) SetWidgetIds(v []int64) {
	o.WidgetIds = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DashboardTab) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["widget_ids"] = o.WidgetIds

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DashboardTab) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Id        *uuid.UUID `json:"id"`
		Name      *string    `json:"name"`
		WidgetIds *[]int64   `json:"widget_ids"`
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
	if all.WidgetIds == nil {
		return fmt.Errorf("required field widget_ids missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"id", "name", "widget_ids"})
	} else {
		return err
	}
	o.Id = *all.Id
	o.Name = *all.Name
	o.WidgetIds = *all.WidgetIds

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
