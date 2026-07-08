// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleListResponseData The JSON:API data object representing a report schedule in a list response.
type ReportScheduleListResponseData struct {
	// The configuration and derived state of a report schedule in a list response.
	Attributes ReportScheduleListResponseAttributes `json:"attributes"`
	// The unique identifier of the report schedule.
	Id string `json:"id"`
	// Relationships for a report schedule in a list response.
	Relationships ReportScheduleListResponseRelationships `json:"relationships"`
	// JSON:API resource type for report schedules.
	Type ReportScheduleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportScheduleListResponseData instantiates a new ReportScheduleListResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportScheduleListResponseData(attributes ReportScheduleListResponseAttributes, id string, relationships ReportScheduleListResponseRelationships, typeVar ReportScheduleType) *ReportScheduleListResponseData {
	this := ReportScheduleListResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewReportScheduleListResponseDataWithDefaults instantiates a new ReportScheduleListResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportScheduleListResponseDataWithDefaults() *ReportScheduleListResponseData {
	this := ReportScheduleListResponseData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ReportScheduleListResponseData) GetAttributes() ReportScheduleListResponseAttributes {
	if o == nil {
		var ret ReportScheduleListResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponseData) GetAttributesOk() (*ReportScheduleListResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ReportScheduleListResponseData) SetAttributes(v ReportScheduleListResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *ReportScheduleListResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ReportScheduleListResponseData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value.
func (o *ReportScheduleListResponseData) GetRelationships() ReportScheduleListResponseRelationships {
	if o == nil {
		var ret ReportScheduleListResponseRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponseData) GetRelationshipsOk() (*ReportScheduleListResponseRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *ReportScheduleListResponseData) SetRelationships(v ReportScheduleListResponseRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *ReportScheduleListResponseData) GetType() ReportScheduleType {
	if o == nil {
		var ret ReportScheduleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleListResponseData) GetTypeOk() (*ReportScheduleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ReportScheduleListResponseData) SetType(v ReportScheduleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportScheduleListResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["id"] = o.Id
	toSerialize["relationships"] = o.Relationships
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReportScheduleListResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *ReportScheduleListResponseAttributes    `json:"attributes"`
		Id            *string                                  `json:"id"`
		Relationships *ReportScheduleListResponseRelationships `json:"relationships"`
		Type          *ReportScheduleType                      `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Relationships == nil {
		return fmt.Errorf("required field relationships missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "relationships", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
	o.Id = *all.Id
	if all.Relationships.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Relationships = *all.Relationships
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
