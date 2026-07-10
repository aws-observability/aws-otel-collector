// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleResponseRelationships Relationships for the report schedule.
type ReportScheduleResponseRelationships struct {
	// Relationship to the author of the report schedule.
	Author ReportScheduleAuthorRelationship `json:"author"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportScheduleResponseRelationships instantiates a new ReportScheduleResponseRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportScheduleResponseRelationships(author ReportScheduleAuthorRelationship) *ReportScheduleResponseRelationships {
	this := ReportScheduleResponseRelationships{}
	this.Author = author
	return &this
}

// NewReportScheduleResponseRelationshipsWithDefaults instantiates a new ReportScheduleResponseRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportScheduleResponseRelationshipsWithDefaults() *ReportScheduleResponseRelationships {
	this := ReportScheduleResponseRelationships{}
	return &this
}

// GetAuthor returns the Author field value.
func (o *ReportScheduleResponseRelationships) GetAuthor() ReportScheduleAuthorRelationship {
	if o == nil {
		var ret ReportScheduleAuthorRelationship
		return ret
	}
	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseRelationships) GetAuthorOk() (*ReportScheduleAuthorRelationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value.
func (o *ReportScheduleResponseRelationships) SetAuthor(v ReportScheduleAuthorRelationship) {
	o.Author = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportScheduleResponseRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["author"] = o.Author

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReportScheduleResponseRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author *ReportScheduleAuthorRelationship `json:"author"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Author == nil {
		return fmt.Errorf("required field author missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Author.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Author = *all.Author

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
