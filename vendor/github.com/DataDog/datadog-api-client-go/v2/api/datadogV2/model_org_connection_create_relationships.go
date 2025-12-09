// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgConnectionCreateRelationships Relationships for org connection creation.
type OrgConnectionCreateRelationships struct {
	// Org relationship.
	SinkOrg OrgConnectionOrgRelationship `json:"sink_org"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgConnectionCreateRelationships instantiates a new OrgConnectionCreateRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgConnectionCreateRelationships(sinkOrg OrgConnectionOrgRelationship) *OrgConnectionCreateRelationships {
	this := OrgConnectionCreateRelationships{}
	this.SinkOrg = sinkOrg
	return &this
}

// NewOrgConnectionCreateRelationshipsWithDefaults instantiates a new OrgConnectionCreateRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgConnectionCreateRelationshipsWithDefaults() *OrgConnectionCreateRelationships {
	this := OrgConnectionCreateRelationships{}
	return &this
}

// GetSinkOrg returns the SinkOrg field value.
func (o *OrgConnectionCreateRelationships) GetSinkOrg() OrgConnectionOrgRelationship {
	if o == nil {
		var ret OrgConnectionOrgRelationship
		return ret
	}
	return o.SinkOrg
}

// GetSinkOrgOk returns a tuple with the SinkOrg field value
// and a boolean to check if the value has been set.
func (o *OrgConnectionCreateRelationships) GetSinkOrgOk() (*OrgConnectionOrgRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SinkOrg, true
}

// SetSinkOrg sets field value.
func (o *OrgConnectionCreateRelationships) SetSinkOrg(v OrgConnectionOrgRelationship) {
	o.SinkOrg = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgConnectionCreateRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["sink_org"] = o.SinkOrg

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgConnectionCreateRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		SinkOrg *OrgConnectionOrgRelationship `json:"sink_org"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.SinkOrg == nil {
		return fmt.Errorf("required field sink_org missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"sink_org"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.SinkOrg.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SinkOrg = *all.SinkOrg

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
