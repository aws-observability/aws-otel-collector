// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseCountResponseAttributes Attributes for the count response, including the total count and optional facet breakdowns.
type CaseCountResponseAttributes struct {
	// List of facet groups, one per field specified in `group_bys`.
	Groups []CaseCountGroup `json:"groups"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseCountResponseAttributes instantiates a new CaseCountResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseCountResponseAttributes(groups []CaseCountGroup) *CaseCountResponseAttributes {
	this := CaseCountResponseAttributes{}
	this.Groups = groups
	return &this
}

// NewCaseCountResponseAttributesWithDefaults instantiates a new CaseCountResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseCountResponseAttributesWithDefaults() *CaseCountResponseAttributes {
	this := CaseCountResponseAttributes{}
	return &this
}

// GetGroups returns the Groups field value.
func (o *CaseCountResponseAttributes) GetGroups() []CaseCountGroup {
	if o == nil {
		var ret []CaseCountGroup
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *CaseCountResponseAttributes) GetGroupsOk() (*[]CaseCountGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Groups, true
}

// SetGroups sets field value.
func (o *CaseCountResponseAttributes) SetGroups(v []CaseCountGroup) {
	o.Groups = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseCountResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["groups"] = o.Groups

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseCountResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Groups *[]CaseCountGroup `json:"groups"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Groups == nil {
		return fmt.Errorf("required field groups missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"groups"})
	} else {
		return err
	}
	o.Groups = *all.Groups

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
