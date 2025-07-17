// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScheduleUpdateRequestDataRelationships Houses relationships for the schedule update, typically referencing teams.
type ScheduleUpdateRequestDataRelationships struct {
	// Associates teams with this schedule in a data structure.
	Teams *DataRelationshipsTeams `json:"teams,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScheduleUpdateRequestDataRelationships instantiates a new ScheduleUpdateRequestDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScheduleUpdateRequestDataRelationships() *ScheduleUpdateRequestDataRelationships {
	this := ScheduleUpdateRequestDataRelationships{}
	return &this
}

// NewScheduleUpdateRequestDataRelationshipsWithDefaults instantiates a new ScheduleUpdateRequestDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScheduleUpdateRequestDataRelationshipsWithDefaults() *ScheduleUpdateRequestDataRelationships {
	this := ScheduleUpdateRequestDataRelationships{}
	return &this
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ScheduleUpdateRequestDataRelationships) GetTeams() DataRelationshipsTeams {
	if o == nil || o.Teams == nil {
		var ret DataRelationshipsTeams
		return ret
	}
	return *o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleUpdateRequestDataRelationships) GetTeamsOk() (*DataRelationshipsTeams, bool) {
	if o == nil || o.Teams == nil {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ScheduleUpdateRequestDataRelationships) HasTeams() bool {
	return o != nil && o.Teams != nil
}

// SetTeams gets a reference to the given DataRelationshipsTeams and assigns it to the Teams field.
func (o *ScheduleUpdateRequestDataRelationships) SetTeams(v DataRelationshipsTeams) {
	o.Teams = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScheduleUpdateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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
func (o *ScheduleUpdateRequestDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Teams *DataRelationshipsTeams `json:"teams,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"teams"})
	} else {
		return err
	}

	hasInvalidField := false
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
