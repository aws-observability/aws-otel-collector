// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityContextEntityAttributes The attributes of an entity context entry, grouping all the historical revisions of the entity.
type EntityContextEntityAttributes struct {
	// The historical revisions of the entity, ordered chronologically.
	Revisions []EntityContextRevision `json:"revisions"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityContextEntityAttributes instantiates a new EntityContextEntityAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityContextEntityAttributes(revisions []EntityContextRevision) *EntityContextEntityAttributes {
	this := EntityContextEntityAttributes{}
	this.Revisions = revisions
	return &this
}

// NewEntityContextEntityAttributesWithDefaults instantiates a new EntityContextEntityAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityContextEntityAttributesWithDefaults() *EntityContextEntityAttributes {
	this := EntityContextEntityAttributes{}
	return &this
}

// GetRevisions returns the Revisions field value.
func (o *EntityContextEntityAttributes) GetRevisions() []EntityContextRevision {
	if o == nil {
		var ret []EntityContextRevision
		return ret
	}
	return o.Revisions
}

// GetRevisionsOk returns a tuple with the Revisions field value
// and a boolean to check if the value has been set.
func (o *EntityContextEntityAttributes) GetRevisionsOk() (*[]EntityContextRevision, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revisions, true
}

// SetRevisions sets field value.
func (o *EntityContextEntityAttributes) SetRevisions(v []EntityContextRevision) {
	o.Revisions = v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityContextEntityAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["revisions"] = o.Revisions

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityContextEntityAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Revisions *[]EntityContextRevision `json:"revisions"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Revisions == nil {
		return fmt.Errorf("required field revisions missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"revisions"})
	} else {
		return err
	}
	o.Revisions = *all.Revisions

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
