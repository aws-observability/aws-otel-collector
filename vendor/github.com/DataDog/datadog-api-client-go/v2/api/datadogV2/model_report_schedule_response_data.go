// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleResponseData The JSON:API data object representing a report schedule.
type ReportScheduleResponseData struct {
	// The configuration and derived state of a report schedule.
	Attributes ReportScheduleResponseAttributes `json:"attributes"`
	// The unique identifier of the report schedule.
	Id string `json:"id"`
	// Relationships for the report schedule.
	Relationships ReportScheduleResponseRelationships `json:"relationships"`
	// JSON:API resource type for report schedules.
	Type ReportScheduleType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportScheduleResponseData instantiates a new ReportScheduleResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportScheduleResponseData(attributes ReportScheduleResponseAttributes, id string, relationships ReportScheduleResponseRelationships, typeVar ReportScheduleType) *ReportScheduleResponseData {
	this := ReportScheduleResponseData{}
	this.Attributes = attributes
	this.Id = id
	this.Relationships = relationships
	this.Type = typeVar
	return &this
}

// NewReportScheduleResponseDataWithDefaults instantiates a new ReportScheduleResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportScheduleResponseDataWithDefaults() *ReportScheduleResponseData {
	this := ReportScheduleResponseData{}
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *ReportScheduleResponseData) GetAttributes() ReportScheduleResponseAttributes {
	if o == nil {
		var ret ReportScheduleResponseAttributes
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseData) GetAttributesOk() (*ReportScheduleResponseAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *ReportScheduleResponseData) SetAttributes(v ReportScheduleResponseAttributes) {
	o.Attributes = v
}

// GetId returns the Id field value.
func (o *ReportScheduleResponseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ReportScheduleResponseData) SetId(v string) {
	o.Id = v
}

// GetRelationships returns the Relationships field value.
func (o *ReportScheduleResponseData) GetRelationships() ReportScheduleResponseRelationships {
	if o == nil {
		var ret ReportScheduleResponseRelationships
		return ret
	}
	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseData) GetRelationshipsOk() (*ReportScheduleResponseRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value.
func (o *ReportScheduleResponseData) SetRelationships(v ReportScheduleResponseRelationships) {
	o.Relationships = v
}

// GetType returns the Type field value.
func (o *ReportScheduleResponseData) GetType() ReportScheduleType {
	if o == nil {
		var ret ReportScheduleType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseData) GetTypeOk() (*ReportScheduleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ReportScheduleResponseData) SetType(v ReportScheduleType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportScheduleResponseData) MarshalJSON() ([]byte, error) {
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
func (o *ReportScheduleResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes    *ReportScheduleResponseAttributes    `json:"attributes"`
		Id            *string                              `json:"id"`
		Relationships *ReportScheduleResponseRelationships `json:"relationships"`
		Type          *ReportScheduleType                  `json:"type"`
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
