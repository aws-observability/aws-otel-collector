// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrganizationResponse Response with an organization.
type OrganizationResponse struct {
	// Create, edit, and manage organizations.
	Org *Organization `json:"org,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrganizationResponse instantiates a new OrganizationResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrganizationResponse() *OrganizationResponse {
	this := OrganizationResponse{}
	return &this
}

// NewOrganizationResponseWithDefaults instantiates a new OrganizationResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrganizationResponseWithDefaults() *OrganizationResponse {
	this := OrganizationResponse{}
	return &this
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *OrganizationResponse) GetOrg() Organization {
	if o == nil || o.Org == nil {
		var ret Organization
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationResponse) GetOrgOk() (*Organization, bool) {
	if o == nil || o.Org == nil {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *OrganizationResponse) HasOrg() bool {
	return o != nil && o.Org != nil
}

// SetOrg gets a reference to the given Organization and assigns it to the Org field.
func (o *OrganizationResponse) SetOrg(v Organization) {
	o.Org = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrganizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Org != nil {
		toSerialize["org"] = o.Org
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrganizationResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Org *Organization `json:"org,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"org"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Org != nil && all.Org.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Org = all.Org

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
