// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamOnCallRespondersDataRelationships Relationship objects linked to a team's on-call responder configuration, including escalations and responders.
type TeamOnCallRespondersDataRelationships struct {
	// Defines the escalation policy steps linked to the team's on-call configuration.
	Escalations *TeamOnCallRespondersDataRelationshipsEscalations `json:"escalations,omitempty"`
	// Defines the list of users assigned as on-call responders for the team.
	Responders *TeamOnCallRespondersDataRelationshipsResponders `json:"responders,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamOnCallRespondersDataRelationships instantiates a new TeamOnCallRespondersDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamOnCallRespondersDataRelationships() *TeamOnCallRespondersDataRelationships {
	this := TeamOnCallRespondersDataRelationships{}
	return &this
}

// NewTeamOnCallRespondersDataRelationshipsWithDefaults instantiates a new TeamOnCallRespondersDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamOnCallRespondersDataRelationshipsWithDefaults() *TeamOnCallRespondersDataRelationships {
	this := TeamOnCallRespondersDataRelationships{}
	return &this
}

// GetEscalations returns the Escalations field value if set, zero value otherwise.
func (o *TeamOnCallRespondersDataRelationships) GetEscalations() TeamOnCallRespondersDataRelationshipsEscalations {
	if o == nil || o.Escalations == nil {
		var ret TeamOnCallRespondersDataRelationshipsEscalations
		return ret
	}
	return *o.Escalations
}

// GetEscalationsOk returns a tuple with the Escalations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamOnCallRespondersDataRelationships) GetEscalationsOk() (*TeamOnCallRespondersDataRelationshipsEscalations, bool) {
	if o == nil || o.Escalations == nil {
		return nil, false
	}
	return o.Escalations, true
}

// HasEscalations returns a boolean if a field has been set.
func (o *TeamOnCallRespondersDataRelationships) HasEscalations() bool {
	return o != nil && o.Escalations != nil
}

// SetEscalations gets a reference to the given TeamOnCallRespondersDataRelationshipsEscalations and assigns it to the Escalations field.
func (o *TeamOnCallRespondersDataRelationships) SetEscalations(v TeamOnCallRespondersDataRelationshipsEscalations) {
	o.Escalations = &v
}

// GetResponders returns the Responders field value if set, zero value otherwise.
func (o *TeamOnCallRespondersDataRelationships) GetResponders() TeamOnCallRespondersDataRelationshipsResponders {
	if o == nil || o.Responders == nil {
		var ret TeamOnCallRespondersDataRelationshipsResponders
		return ret
	}
	return *o.Responders
}

// GetRespondersOk returns a tuple with the Responders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamOnCallRespondersDataRelationships) GetRespondersOk() (*TeamOnCallRespondersDataRelationshipsResponders, bool) {
	if o == nil || o.Responders == nil {
		return nil, false
	}
	return o.Responders, true
}

// HasResponders returns a boolean if a field has been set.
func (o *TeamOnCallRespondersDataRelationships) HasResponders() bool {
	return o != nil && o.Responders != nil
}

// SetResponders gets a reference to the given TeamOnCallRespondersDataRelationshipsResponders and assigns it to the Responders field.
func (o *TeamOnCallRespondersDataRelationships) SetResponders(v TeamOnCallRespondersDataRelationshipsResponders) {
	o.Responders = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamOnCallRespondersDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Escalations != nil {
		toSerialize["escalations"] = o.Escalations
	}
	if o.Responders != nil {
		toSerialize["responders"] = o.Responders
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamOnCallRespondersDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Escalations *TeamOnCallRespondersDataRelationshipsEscalations `json:"escalations,omitempty"`
		Responders  *TeamOnCallRespondersDataRelationshipsResponders  `json:"responders,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"escalations", "responders"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Escalations != nil && all.Escalations.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Escalations = all.Escalations
	if all.Responders != nil && all.Responders.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Responders = all.Responders

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
