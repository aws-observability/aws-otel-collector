// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DegradationUpdateDataRelationships Relationships of a degradation update resource.
type DegradationUpdateDataRelationships struct {
	// A user relationship of a degradation update.
	CreatedByUser *DegradationUpdateDataRelationshipsUser `json:"created_by_user,omitempty"`
	// The degradation relationship of a degradation update.
	Degradation *DegradationUpdateDataRelationshipsDegradation `json:"degradation,omitempty"`
	// A user relationship of a degradation update.
	DeletedByUser *DegradationUpdateDataRelationshipsUser `json:"deleted_by_user,omitempty"`
	// A user relationship of a degradation update.
	LastModifiedByUser *DegradationUpdateDataRelationshipsUser `json:"last_modified_by_user,omitempty"`
	// The status page relationship of a degradation update.
	StatusPage *DegradationUpdateDataRelationshipsStatusPage `json:"status_page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDegradationUpdateDataRelationships instantiates a new DegradationUpdateDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDegradationUpdateDataRelationships() *DegradationUpdateDataRelationships {
	this := DegradationUpdateDataRelationships{}
	return &this
}

// NewDegradationUpdateDataRelationshipsWithDefaults instantiates a new DegradationUpdateDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDegradationUpdateDataRelationshipsWithDefaults() *DegradationUpdateDataRelationships {
	this := DegradationUpdateDataRelationships{}
	return &this
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *DegradationUpdateDataRelationships) GetCreatedByUser() DegradationUpdateDataRelationshipsUser {
	if o == nil || o.CreatedByUser == nil {
		var ret DegradationUpdateDataRelationshipsUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationUpdateDataRelationships) GetCreatedByUserOk() (*DegradationUpdateDataRelationshipsUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *DegradationUpdateDataRelationships) HasCreatedByUser() bool {
	return o != nil && o.CreatedByUser != nil
}

// SetCreatedByUser gets a reference to the given DegradationUpdateDataRelationshipsUser and assigns it to the CreatedByUser field.
func (o *DegradationUpdateDataRelationships) SetCreatedByUser(v DegradationUpdateDataRelationshipsUser) {
	o.CreatedByUser = &v
}

// GetDegradation returns the Degradation field value if set, zero value otherwise.
func (o *DegradationUpdateDataRelationships) GetDegradation() DegradationUpdateDataRelationshipsDegradation {
	if o == nil || o.Degradation == nil {
		var ret DegradationUpdateDataRelationshipsDegradation
		return ret
	}
	return *o.Degradation
}

// GetDegradationOk returns a tuple with the Degradation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationUpdateDataRelationships) GetDegradationOk() (*DegradationUpdateDataRelationshipsDegradation, bool) {
	if o == nil || o.Degradation == nil {
		return nil, false
	}
	return o.Degradation, true
}

// HasDegradation returns a boolean if a field has been set.
func (o *DegradationUpdateDataRelationships) HasDegradation() bool {
	return o != nil && o.Degradation != nil
}

// SetDegradation gets a reference to the given DegradationUpdateDataRelationshipsDegradation and assigns it to the Degradation field.
func (o *DegradationUpdateDataRelationships) SetDegradation(v DegradationUpdateDataRelationshipsDegradation) {
	o.Degradation = &v
}

// GetDeletedByUser returns the DeletedByUser field value if set, zero value otherwise.
func (o *DegradationUpdateDataRelationships) GetDeletedByUser() DegradationUpdateDataRelationshipsUser {
	if o == nil || o.DeletedByUser == nil {
		var ret DegradationUpdateDataRelationshipsUser
		return ret
	}
	return *o.DeletedByUser
}

// GetDeletedByUserOk returns a tuple with the DeletedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationUpdateDataRelationships) GetDeletedByUserOk() (*DegradationUpdateDataRelationshipsUser, bool) {
	if o == nil || o.DeletedByUser == nil {
		return nil, false
	}
	return o.DeletedByUser, true
}

// HasDeletedByUser returns a boolean if a field has been set.
func (o *DegradationUpdateDataRelationships) HasDeletedByUser() bool {
	return o != nil && o.DeletedByUser != nil
}

// SetDeletedByUser gets a reference to the given DegradationUpdateDataRelationshipsUser and assigns it to the DeletedByUser field.
func (o *DegradationUpdateDataRelationships) SetDeletedByUser(v DegradationUpdateDataRelationshipsUser) {
	o.DeletedByUser = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *DegradationUpdateDataRelationships) GetLastModifiedByUser() DegradationUpdateDataRelationshipsUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret DegradationUpdateDataRelationshipsUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationUpdateDataRelationships) GetLastModifiedByUserOk() (*DegradationUpdateDataRelationshipsUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *DegradationUpdateDataRelationships) HasLastModifiedByUser() bool {
	return o != nil && o.LastModifiedByUser != nil
}

// SetLastModifiedByUser gets a reference to the given DegradationUpdateDataRelationshipsUser and assigns it to the LastModifiedByUser field.
func (o *DegradationUpdateDataRelationships) SetLastModifiedByUser(v DegradationUpdateDataRelationshipsUser) {
	o.LastModifiedByUser = &v
}

// GetStatusPage returns the StatusPage field value if set, zero value otherwise.
func (o *DegradationUpdateDataRelationships) GetStatusPage() DegradationUpdateDataRelationshipsStatusPage {
	if o == nil || o.StatusPage == nil {
		var ret DegradationUpdateDataRelationshipsStatusPage
		return ret
	}
	return *o.StatusPage
}

// GetStatusPageOk returns a tuple with the StatusPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationUpdateDataRelationships) GetStatusPageOk() (*DegradationUpdateDataRelationshipsStatusPage, bool) {
	if o == nil || o.StatusPage == nil {
		return nil, false
	}
	return o.StatusPage, true
}

// HasStatusPage returns a boolean if a field has been set.
func (o *DegradationUpdateDataRelationships) HasStatusPage() bool {
	return o != nil && o.StatusPage != nil
}

// SetStatusPage gets a reference to the given DegradationUpdateDataRelationshipsStatusPage and assigns it to the StatusPage field.
func (o *DegradationUpdateDataRelationships) SetStatusPage(v DegradationUpdateDataRelationshipsStatusPage) {
	o.StatusPage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DegradationUpdateDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedByUser != nil {
		toSerialize["created_by_user"] = o.CreatedByUser
	}
	if o.Degradation != nil {
		toSerialize["degradation"] = o.Degradation
	}
	if o.DeletedByUser != nil {
		toSerialize["deleted_by_user"] = o.DeletedByUser
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
func (o *DegradationUpdateDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedByUser      *DegradationUpdateDataRelationshipsUser        `json:"created_by_user,omitempty"`
		Degradation        *DegradationUpdateDataRelationshipsDegradation `json:"degradation,omitempty"`
		DeletedByUser      *DegradationUpdateDataRelationshipsUser        `json:"deleted_by_user,omitempty"`
		LastModifiedByUser *DegradationUpdateDataRelationshipsUser        `json:"last_modified_by_user,omitempty"`
		StatusPage         *DegradationUpdateDataRelationshipsStatusPage  `json:"status_page,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by_user", "degradation", "deleted_by_user", "last_modified_by_user", "status_page"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedByUser != nil && all.CreatedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedByUser = all.CreatedByUser
	if all.Degradation != nil && all.Degradation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Degradation = all.Degradation
	if all.DeletedByUser != nil && all.DeletedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DeletedByUser = all.DeletedByUser
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
