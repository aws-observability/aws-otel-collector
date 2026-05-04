// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPagesComponentGroupRelationships The relationships of a component group.
type StatusPagesComponentGroupRelationships struct {
	// The Datadog user who created the component group.
	CreatedByUser *StatusPagesComponentGroupRelationshipsCreatedByUser `json:"created_by_user,omitempty"`
	// The group the component group belongs to.
	Group *StatusPagesComponentGroupRelationshipsGroup `json:"group,omitempty"`
	// The Datadog user who last modified the component group.
	LastModifiedByUser *StatusPagesComponentGroupRelationshipsLastModifiedByUser `json:"last_modified_by_user,omitempty"`
	// The status page the component group belongs to.
	StatusPage *StatusPagesComponentGroupRelationshipsStatusPage `json:"status_page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPagesComponentGroupRelationships instantiates a new StatusPagesComponentGroupRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPagesComponentGroupRelationships() *StatusPagesComponentGroupRelationships {
	this := StatusPagesComponentGroupRelationships{}
	return &this
}

// NewStatusPagesComponentGroupRelationshipsWithDefaults instantiates a new StatusPagesComponentGroupRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPagesComponentGroupRelationshipsWithDefaults() *StatusPagesComponentGroupRelationships {
	this := StatusPagesComponentGroupRelationships{}
	return &this
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *StatusPagesComponentGroupRelationships) GetCreatedByUser() StatusPagesComponentGroupRelationshipsCreatedByUser {
	if o == nil || o.CreatedByUser == nil {
		var ret StatusPagesComponentGroupRelationshipsCreatedByUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentGroupRelationships) GetCreatedByUserOk() (*StatusPagesComponentGroupRelationshipsCreatedByUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *StatusPagesComponentGroupRelationships) HasCreatedByUser() bool {
	return o != nil && o.CreatedByUser != nil
}

// SetCreatedByUser gets a reference to the given StatusPagesComponentGroupRelationshipsCreatedByUser and assigns it to the CreatedByUser field.
func (o *StatusPagesComponentGroupRelationships) SetCreatedByUser(v StatusPagesComponentGroupRelationshipsCreatedByUser) {
	o.CreatedByUser = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *StatusPagesComponentGroupRelationships) GetGroup() StatusPagesComponentGroupRelationshipsGroup {
	if o == nil || o.Group == nil {
		var ret StatusPagesComponentGroupRelationshipsGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentGroupRelationships) GetGroupOk() (*StatusPagesComponentGroupRelationshipsGroup, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *StatusPagesComponentGroupRelationships) HasGroup() bool {
	return o != nil && o.Group != nil
}

// SetGroup gets a reference to the given StatusPagesComponentGroupRelationshipsGroup and assigns it to the Group field.
func (o *StatusPagesComponentGroupRelationships) SetGroup(v StatusPagesComponentGroupRelationshipsGroup) {
	o.Group = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *StatusPagesComponentGroupRelationships) GetLastModifiedByUser() StatusPagesComponentGroupRelationshipsLastModifiedByUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret StatusPagesComponentGroupRelationshipsLastModifiedByUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentGroupRelationships) GetLastModifiedByUserOk() (*StatusPagesComponentGroupRelationshipsLastModifiedByUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *StatusPagesComponentGroupRelationships) HasLastModifiedByUser() bool {
	return o != nil && o.LastModifiedByUser != nil
}

// SetLastModifiedByUser gets a reference to the given StatusPagesComponentGroupRelationshipsLastModifiedByUser and assigns it to the LastModifiedByUser field.
func (o *StatusPagesComponentGroupRelationships) SetLastModifiedByUser(v StatusPagesComponentGroupRelationshipsLastModifiedByUser) {
	o.LastModifiedByUser = &v
}

// GetStatusPage returns the StatusPage field value if set, zero value otherwise.
func (o *StatusPagesComponentGroupRelationships) GetStatusPage() StatusPagesComponentGroupRelationshipsStatusPage {
	if o == nil || o.StatusPage == nil {
		var ret StatusPagesComponentGroupRelationshipsStatusPage
		return ret
	}
	return *o.StatusPage
}

// GetStatusPageOk returns a tuple with the StatusPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPagesComponentGroupRelationships) GetStatusPageOk() (*StatusPagesComponentGroupRelationshipsStatusPage, bool) {
	if o == nil || o.StatusPage == nil {
		return nil, false
	}
	return o.StatusPage, true
}

// HasStatusPage returns a boolean if a field has been set.
func (o *StatusPagesComponentGroupRelationships) HasStatusPage() bool {
	return o != nil && o.StatusPage != nil
}

// SetStatusPage gets a reference to the given StatusPagesComponentGroupRelationshipsStatusPage and assigns it to the StatusPage field.
func (o *StatusPagesComponentGroupRelationships) SetStatusPage(v StatusPagesComponentGroupRelationshipsStatusPage) {
	o.StatusPage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPagesComponentGroupRelationships) MarshalJSON() ([]byte, error) {
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
func (o *StatusPagesComponentGroupRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedByUser      *StatusPagesComponentGroupRelationshipsCreatedByUser      `json:"created_by_user,omitempty"`
		Group              *StatusPagesComponentGroupRelationshipsGroup              `json:"group,omitempty"`
		LastModifiedByUser *StatusPagesComponentGroupRelationshipsLastModifiedByUser `json:"last_modified_by_user,omitempty"`
		StatusPage         *StatusPagesComponentGroupRelationshipsStatusPage         `json:"status_page,omitempty"`
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
