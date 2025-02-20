// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrganizationListResponse Response with the list of organizations.
type OrganizationListResponse struct {
	// Array of organization objects.
	Orgs []Organization `json:"orgs,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrganizationListResponse instantiates a new OrganizationListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrganizationListResponse() *OrganizationListResponse {
	this := OrganizationListResponse{}
	return &this
}

// NewOrganizationListResponseWithDefaults instantiates a new OrganizationListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrganizationListResponseWithDefaults() *OrganizationListResponse {
	this := OrganizationListResponse{}
	return &this
}

// GetOrgs returns the Orgs field value if set, zero value otherwise.
func (o *OrganizationListResponse) GetOrgs() []Organization {
	if o == nil || o.Orgs == nil {
		var ret []Organization
		return ret
	}
	return o.Orgs
}

// GetOrgsOk returns a tuple with the Orgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationListResponse) GetOrgsOk() (*[]Organization, bool) {
	if o == nil || o.Orgs == nil {
		return nil, false
	}
	return &o.Orgs, true
}

// HasOrgs returns a boolean if a field has been set.
func (o *OrganizationListResponse) HasOrgs() bool {
	return o != nil && o.Orgs != nil
}

// SetOrgs gets a reference to the given []Organization and assigns it to the Orgs field.
func (o *OrganizationListResponse) SetOrgs(v []Organization) {
	o.Orgs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrganizationListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Orgs != nil {
		toSerialize["orgs"] = o.Orgs
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrganizationListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Orgs []Organization `json:"orgs,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"orgs"})
	} else {
		return err
	}
	o.Orgs = all.Orgs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
