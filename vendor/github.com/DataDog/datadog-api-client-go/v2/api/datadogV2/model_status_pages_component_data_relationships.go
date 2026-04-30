// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPagesComponentDataRelationships The relationships of a component.
type StatusPagesComponentDataRelationships struct {
	// The Datadog user who created the component.
	CreatedByUser *StatusPagesComponentDataRelationshipsCreatedByUser `json:"created_by_user,omitempty"`
	// The group the component belongs to.
	Group *StatusPagesComponentDataRelationshipsGroup `json:"group,omitempty"`
	// The Datadog user who last modified the component.
	LastModifiedByUser *StatusPagesComponentDataRelationshipsLastModifiedByUser `json:"last_modified_by_user,omitempty"`
	// The status page the component belongs to.
	StatusPage *StatusPagesComponentDataRelationshipsStatusPage `json:"status_page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPagesComponentDataRelationships instantiates a new StatusPagesComponentDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPagesComponentDataRelationships() *StatusPagesComponentDataRelationships {
	this := StatusPagesComponentDataRelationships{}
	return &this
}

// NewStatusPagesComponentDataRelationshipsWithDefaults instantiates a new StatusPagesComponentDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPagesComponentDataRelationshipsWithDefaults() *StatusPagesComponentDataRelationships {
	this := StatusPagesComponentDataRelationships{}
	return &this
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *StatusPagesComponentDataRelationships) GetCreatedByUser() StatusPagesComponentDataRelationshipsCreatedByUser {
	if o == nil || o.CreatedByUser == nil {
		var ret StatusPagesComponentDataRelationshipsCreatedByUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentDataRelationships) GetCreatedByUserOk() (*StatusPagesComponentDataRelationshipsCreatedByUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *StatusPagesComponentDataRelationships) HasCreatedByUser() bool {
	return o != nil && o.CreatedByUser != nil
}

// SetCreatedByUser gets a reference to the given StatusPagesComponentDataRelationshipsCreatedByUser and assigns it to the CreatedByUser field.
func (o *StatusPagesComponentDataRelationships) SetCreatedByUser(v StatusPagesComponentDataRelationshipsCreatedByUser) {
	o.CreatedByUser = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *StatusPagesComponentDataRelationships) GetGroup() StatusPagesComponentDataRelationshipsGroup {
	if o == nil || o.Group == nil {
		var ret StatusPagesComponentDataRelationshipsGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentDataRelationships) GetGroupOk() (*StatusPagesComponentDataRelationshipsGroup, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *StatusPagesComponentDataRelationships) HasGroup() bool {
	return o != nil && o.Group != nil
}

// SetGroup gets a reference to the given StatusPagesComponentDataRelationshipsGroup and assigns it to the Group field.
func (o *StatusPagesComponentDataRelationships) SetGroup(v StatusPagesComponentDataRelationshipsGroup) {
	o.Group = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *StatusPagesComponentDataRelationships) GetLastModifiedByUser() StatusPagesComponentDataRelationshipsLastModifiedByUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret StatusPagesComponentDataRelationshipsLastModifiedByUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentDataRelationships) GetLastModifiedByUserOk() (*StatusPagesComponentDataRelationshipsLastModifiedByUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *StatusPagesComponentDataRelationships) HasLastModifiedByUser() bool {
	return o != nil && o.LastModifiedByUser != nil
}

// SetLastModifiedByUser gets a reference to the given StatusPagesComponentDataRelationshipsLastModifiedByUser and assigns it to the LastModifiedByUser field.
func (o *StatusPagesComponentDataRelationships) SetLastModifiedByUser(v StatusPagesComponentDataRelationshipsLastModifiedByUser) {
	o.LastModifiedByUser = &v
}

// GetStatusPage returns the StatusPage field value if set, zero value otherwise.
func (o *StatusPagesComponentDataRelationships) GetStatusPage() StatusPagesComponentDataRelationshipsStatusPage {
	if o == nil || o.StatusPage == nil {
		var ret StatusPagesComponentDataRelationshipsStatusPage
		return ret
	}
	return *o.StatusPage
}

// GetStatusPageOk returns a tuple with the StatusPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentDataRelationships) GetStatusPageOk() (*StatusPagesComponentDataRelationshipsStatusPage, bool) {
	if o == nil || o.StatusPage == nil {
		return nil, false
	}
	return o.StatusPage, true
}

// HasStatusPage returns a boolean if a field has been set.
func (o *StatusPagesComponentDataRelationships) HasStatusPage() bool {
	return o != nil && o.StatusPage != nil
}

// SetStatusPage gets a reference to the given StatusPagesComponentDataRelationshipsStatusPage and assigns it to the StatusPage field.
func (o *StatusPagesComponentDataRelationships) SetStatusPage(v StatusPagesComponentDataRelationshipsStatusPage) {
	o.StatusPage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPagesComponentDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedByUser != nil {
		toSerialize["created_by_user"] = o.CreatedByUser
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.LastModifiedByUser != nil {
		toSerialize["last_modified_by_user"] = o.LastModifiedByUser
	}
	if o.StatusPage != nil {
		toSerialize["status_page"] = o.StatusPage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StatusPagesComponentDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedByUser      *StatusPagesComponentDataRelationshipsCreatedByUser      `json:"created_by_user,omitempty"`
		Group              *StatusPagesComponentDataRelationshipsGroup              `json:"group,omitempty"`
		LastModifiedByUser *StatusPagesComponentDataRelationshipsLastModifiedByUser `json:"last_modified_by_user,omitempty"`
		StatusPage         *StatusPagesComponentDataRelationshipsStatusPage         `json:"status_page,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by_user", "group", "last_modified_by_user", "status_page"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedByUser != nil && all.CreatedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedByUser = all.CreatedByUser
	if all.Group != nil && all.Group.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Group = all.Group
	if all.LastModifiedByUser != nil && all.LastModifiedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastModifiedByUser = all.LastModifiedByUser
	if all.StatusPage != nil && all.StatusPage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.StatusPage = all.StatusPage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
