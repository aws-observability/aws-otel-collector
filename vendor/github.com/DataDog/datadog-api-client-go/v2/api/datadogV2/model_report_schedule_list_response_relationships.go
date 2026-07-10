// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleListResponseRelationships Relationships for a report schedule in a list response.
type ReportScheduleListResponseRelationships struct {
	// Relationship to the author of the report schedule.
	Author ReportScheduleAuthorRelationship `json:"author"`
	// Relationship to the report target resource.
	Resource *ReportScheduleListResourceRelationship `json:"resource,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportScheduleListResponseRelationships instantiates a new ReportScheduleListResponseRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportScheduleListResponseRelationships(author ReportScheduleAuthorRelationship) *ReportScheduleListResponseRelationships {
	this := ReportScheduleListResponseRelationships{}
	this.Author = author
	return &this
}

// NewReportScheduleListResponseRelationshipsWithDefaults instantiates a new ReportScheduleListResponseRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportScheduleListResponseRelationshipsWithDefaults() *ReportScheduleListResponseRelationships {
	this := ReportScheduleListResponseRelationships{}
	return &this
}

// GetAuthor returns the Author field value.
func (o *ReportScheduleListResponseRelationships) GetAuthor() ReportScheduleAuthorRelationship {
	if o == nil {
		var ret ReportScheduleAuthorRelationship
		return ret
	}
	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponseRelationships) GetAuthorOk() (*ReportScheduleAuthorRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value.
func (o *ReportScheduleListResponseRelationships) SetAuthor(v ReportScheduleAuthorRelationship) {
	o.Author = v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ReportScheduleListResponseRelationships) GetResource() ReportScheduleListResourceRelationship {
	if o == nil || o.Resource == nil {
		var ret ReportScheduleListResourceRelationship
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponseRelationships) GetResourceOk() (*ReportScheduleListResourceRelationship, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ReportScheduleListResponseRelationships) HasResource() bool {
	return o != nil && o.Resource != nil
}

// SetResource gets a reference to the given ReportScheduleListResourceRelationship and assigns it to the Resource field.
func (o *ReportScheduleListResponseRelationships) SetResource(v ReportScheduleListResourceRelationship) {
	o.Resource = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportScheduleListResponseRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["author"] = o.Author
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReportScheduleListResponseRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author   *ReportScheduleAuthorRelationship       `json:"author"`
		Resource *ReportScheduleListResourceRelationship `json:"resource,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Author == nil {
		return fmt.Errorf("required field author missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "resource"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Author.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Author = *all.Author
	if all.Resource != nil && all.Resource.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Resource = all.Resource

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
