// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DegradationDataRelationships The relationships of a degradation.
type DegradationDataRelationships struct {
	// The Datadog user who created the degradation.
	CreatedByUser *DegradationDataRelationshipsCreatedByUser `json:"created_by_user,omitempty"`
	// The Datadog user who last modified the degradation.
	LastModifiedByUser *DegradationDataRelationshipsLastModifiedByUser `json:"last_modified_by_user,omitempty"`
	// The status page the degradation belongs to.
	StatusPage *DegradationDataRelationshipsStatusPage `json:"status_page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDegradationDataRelationships instantiates a new DegradationDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDegradationDataRelationships() *DegradationDataRelationships {
	this := DegradationDataRelationships{}
	return &this
}

// NewDegradationDataRelationshipsWithDefaults instantiates a new DegradationDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDegradationDataRelationshipsWithDefaults() *DegradationDataRelationships {
	this := DegradationDataRelationships{}
	return &this
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise.
func (o *DegradationDataRelationships) GetCreatedByUser() DegradationDataRelationshipsCreatedByUser {
	if o == nil || o.CreatedByUser == nil {
		var ret DegradationDataRelationshipsCreatedByUser
		return ret
	}
	return *o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationDataRelationships) GetCreatedByUserOk() (*DegradationDataRelationshipsCreatedByUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *DegradationDataRelationships) HasCreatedByUser() bool {
	return o != nil && o.CreatedByUser != nil
}

// SetCreatedByUser gets a reference to the given DegradationDataRelationshipsCreatedByUser and assigns it to the CreatedByUser field.
func (o *DegradationDataRelationships) SetCreatedByUser(v DegradationDataRelationshipsCreatedByUser) {
	o.CreatedByUser = &v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise.
func (o *DegradationDataRelationships) GetLastModifiedByUser() DegradationDataRelationshipsLastModifiedByUser {
	if o == nil || o.LastModifiedByUser == nil {
		var ret DegradationDataRelationshipsLastModifiedByUser
		return ret
	}
	return *o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationDataRelationships) GetLastModifiedByUserOk() (*DegradationDataRelationshipsLastModifiedByUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *DegradationDataRelationships) HasLastModifiedByUser() bool {
	return o != nil && o.LastModifiedByUser != nil
}

// SetLastModifiedByUser gets a reference to the given DegradationDataRelationshipsLastModifiedByUser and assigns it to the LastModifiedByUser field.
func (o *DegradationDataRelationships) SetLastModifiedByUser(v DegradationDataRelationshipsLastModifiedByUser) {
	o.LastModifiedByUser = &v
}

// GetStatusPage returns the StatusPage field value if set, zero value otherwise.
func (o *DegradationDataRelationships) GetStatusPage() DegradationDataRelationshipsStatusPage {
	if o == nil || o.StatusPage == nil {
		var ret DegradationDataRelationshipsStatusPage
		return ret
	}
	return *o.StatusPage
}

// GetStatusPageOk returns a tuple with the StatusPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DegradationDataRelationships) GetStatusPageOk() (*DegradationDataRelationshipsStatusPage, bool) {
	if o == nil || o.StatusPage == nil {
		return nil, false
	}
	return o.StatusPage, true
}

// HasStatusPage returns a boolean if a field has been set.
func (o *DegradationDataRelationships) HasStatusPage() bool {
	return o != nil && o.StatusPage != nil
}

// SetStatusPage gets a reference to the given DegradationDataRelationshipsStatusPage and assigns it to the StatusPage field.
func (o *DegradationDataRelationships) SetStatusPage(v DegradationDataRelationshipsStatusPage) {
	o.StatusPage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DegradationDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CreatedByUser != nil {
		toSerialize["created_by_user"] = o.CreatedByUser
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
func (o *DegradationDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CreatedByUser      *DegradationDataRelationshipsCreatedByUser      `json:"created_by_user,omitempty"`
		LastModifiedByUser *DegradationDataRelationshipsLastModifiedByUser `json:"last_modified_by_user,omitempty"`
		StatusPage         *DegradationDataRelationshipsStatusPage         `json:"status_page,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"created_by_user", "last_modified_by_user", "status_page"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CreatedByUser != nil && all.CreatedByUser.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CreatedByUser = all.CreatedByUser
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
