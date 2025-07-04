// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleDataRelationships Groups the relationships for a schedule object, referencing layers and teams.
type ScheduleDataRelationships struct {
	// Associates layers with this schedule in a data structure.
	Layers *ScheduleDataRelationshipsLayers `json:"layers,omitempty"`
	// Associates teams with this schedule in a data structure.
	Teams *DataRelationshipsTeams `json:"teams,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleDataRelationships instantiates a new ScheduleDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleDataRelationships() *ScheduleDataRelationships {
	this := ScheduleDataRelationships{}
	return &this
}

// NewScheduleDataRelationshipsWithDefaults instantiates a new ScheduleDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleDataRelationshipsWithDefaults() *ScheduleDataRelationships {
	this := ScheduleDataRelationships{}
	return &this
}

// GetLayers returns the Layers field value if set, zero value otherwise.
func (o *ScheduleDataRelationships) GetLayers() ScheduleDataRelationshipsLayers {
	if o == nil || o.Layers == nil {
		var ret ScheduleDataRelationshipsLayers
		return ret
	}
	return *o.Layers
}

// GetLayersOk returns a tuple with the Layers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleDataRelationships) GetLayersOk() (*ScheduleDataRelationshipsLayers, bool) {
	if o == nil || o.Layers == nil {
		return nil, false
	}
	return o.Layers, true
}

// HasLayers returns a boolean if a field has been set.
func (o *ScheduleDataRelationships) HasLayers() bool {
	return o != nil && o.Layers != nil
}

// SetLayers gets a reference to the given ScheduleDataRelationshipsLayers and assigns it to the Layers field.
func (o *ScheduleDataRelationships) SetLayers(v ScheduleDataRelationshipsLayers) {
	o.Layers = &v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ScheduleDataRelationships) GetTeams() DataRelationshipsTeams {
	if o == nil || o.Teams == nil {
		var ret DataRelationshipsTeams
		return ret
	}
	return *o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleDataRelationships) GetTeamsOk() (*DataRelationshipsTeams, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ScheduleDataRelationships) HasTeams() bool {
	return o != nil && o.Teams != nil
}

// SetTeams gets a reference to the given DataRelationshipsTeams and assigns it to the Teams field.
func (o *ScheduleDataRelationships) SetTeams(v DataRelationshipsTeams) {
	o.Teams = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Layers != nil {
		toSerialize["layers"] = o.Layers
	}
	if o.Teams != nil {
		toSerialize["teams"] = o.Teams
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScheduleDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Layers *ScheduleDataRelationshipsLayers `json:"layers,omitempty"`
		Teams  *DataRelationshipsTeams          `json:"teams,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"layers", "teams"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Layers != nil && all.Layers.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Layers = all.Layers
	if all.Teams != nil && all.Teams.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Teams = all.Teams

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
