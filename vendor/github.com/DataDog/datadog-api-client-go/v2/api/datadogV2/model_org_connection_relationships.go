// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrgConnectionRelationships Related organizations and user.
type OrgConnectionRelationships struct {
	// User relationship.
	CreatedBy *OrgConnectionUserRelationship `json:"created_by,omitempty"`
	// Org relationship.
	SinkOrg *OrgConnectionOrgRelationship `json:"sink_org,omitempty"`
	// Org relationship.
	SourceOrg *OrgConnectionOrgRelationship `json:"source_org,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrgConnectionRelationships instantiates a new OrgConnectionRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrgConnectionRelationships() *OrgConnectionRelationships {
	this := OrgConnectionRelationships{}
	return &this
}

// NewOrgConnectionRelationshipsWithDefaults instantiates a new OrgConnectionRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrgConnectionRelationshipsWithDefaults() *OrgConnectionRelationships {
	this := OrgConnectionRelationships{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *OrgConnectionRelationships) GetCreatedBy() OrgConnectionUserRelationship {
	if o == nil || o.CreatedBy == nil {
		var ret OrgConnectionUserRelationship
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgConnectionRelationships) GetCreatedByOk() (*OrgConnectionUserRelationship, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *OrgConnectionRelationships) HasCreatedBy() bool {
	return o != nil && o.CreatedBy != nil
}

// SetCreatedBy gets a reference to the given OrgConnectionUserRelationship and assigns it to the CreatedBy field.
func (o *OrgConnectionRelationships) SetCreatedBy(v OrgConnectionUserRelationship) {
	o.CreatedBy = &v
}

// GetSinkOrg returns the SinkOrg field value if set, zero value otherwise.
func (o *OrgConnectionRelationships) GetSinkOrg() OrgConnectionOrgRelationship {
	if o == nil || o.SinkOrg == nil {
		var ret OrgConnectionOrgRelationship
		return ret
	}
	return *o.SinkOrg
}

// GetSinkOrgOk returns a tuple with the SinkOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgConnectionRelationships) GetSinkOrgOk() (*OrgConnectionOrgRelationship, bool) {
	if o == nil || o.SinkOrg == nil {
		return nil, false
	}
	return o.SinkOrg, true
}

// HasSinkOrg returns a boolean if a field has been set.
func (o *OrgConnectionRelationships) HasSinkOrg() bool {
	return o != nil && o.SinkOrg != nil
}

// SetSinkOrg gets a reference to the given OrgConnectionOrgRelationship and assigns it to the SinkOrg field.
func (o *OrgConnectionRelationships) SetSinkOrg(v OrgConnectionOrgRelationship) {
	o.SinkOrg = &v
}

// GetSourceOrg returns the SourceOrg field value if set, zero value otherwise.
func (o *OrgConnectionRelationships) GetSourceOrg() OrgConnectionOrgRelationship {
	if o == nil || o.SourceOrg == nil {
		var ret OrgConnectionOrgRelationship
		return ret
	}
	return *o.SourceOrg
}

// GetSourceOrgOk returns a tuple with the SourceOrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgConnectionRelationships) GetSourceOrgOk() (*OrgConnectionOrgRelationship, bool) {
	if o == nil || o.SourceOrg == nil {
		return nil, false
	}
	return o.SourceOrg, true
}

// HasSourceOrg returns a boolean if a field has been set.
func (o *OrgConnectionRelationships) HasSourceOrg() bool {
	return o != nil && o.SourceOrg != nil
}

// SetSourceOrg gets a reference to the given OrgConnectionOrgRelationship and assigns it to the SourceOrg field.
func (o *OrgConnectionRelationships) SetSourceOrg(v OrgConnectionOrgRelationship) {
	o.SourceOrg = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrgConnectionRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.SinkOrg != nil {
		toSerialize["sink_org"] = o.SinkOrg
	}
	if o.SourceOrg != nil {
		toSerialize["source_org"] = o.SourceOrg
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrgConnectionRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedBy *OrgConnectionUserRelationship `json:"created_by,omitempty"`
		SinkOrg   *OrgConnectionOrgRelationship  `json:"sink_org,omitempty"`
		SourceOrg *OrgConnectionOrgRelationship  `json:"source_org,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by", "sink_org", "source_org"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedBy != nil && all.CreatedBy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedBy = all.CreatedBy
	if all.SinkOrg != nil && all.SinkOrg.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SinkOrg = all.SinkOrg
	if all.SourceOrg != nil && all.SourceOrg.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SourceOrg = all.SourceOrg

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
