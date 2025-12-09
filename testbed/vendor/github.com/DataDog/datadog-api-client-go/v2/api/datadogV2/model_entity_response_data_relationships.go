// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseDataRelationships
type EntityResponseDataRelationships struct {
	//
	Incidents *EntityResponseDataRelationshipsIncidents `json:"incidents,omitempty"`
	//
	Oncalls *EntityResponseDataRelationshipsOncalls `json:"oncalls,omitempty"`
	//
	RawSchema *EntityResponseDataRelationshipsRawSchema `json:"rawSchema,omitempty"`
	//
	RelatedEntities *EntityResponseDataRelationshipsRelatedEntities `json:"relatedEntities,omitempty"`
	//
	Schema *EntityResponseDataRelationshipsSchema `json:"schema,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityResponseDataRelationships instantiates a new EntityResponseDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityResponseDataRelationships() *EntityResponseDataRelationships {
	this := EntityResponseDataRelationships{}
	return &this
}

// NewEntityResponseDataRelationshipsWithDefaults instantiates a new EntityResponseDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityResponseDataRelationshipsWithDefaults() *EntityResponseDataRelationships {
	this := EntityResponseDataRelationships{}
	return &this
}

// GetIncidents returns the Incidents field value if set, zero value otherwise.
func (o *EntityResponseDataRelationships) GetIncidents() EntityResponseDataRelationshipsIncidents {
	if o == nil || o.Incidents == nil {
		var ret EntityResponseDataRelationshipsIncidents
		return ret
	}
	return *o.Incidents
}

// GetIncidentsOk returns a tuple with the Incidents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseDataRelationships) GetIncidentsOk() (*EntityResponseDataRelationshipsIncidents, bool) {
	if o == nil || o.Incidents == nil {
		return nil, false
	}
	return o.Incidents, true
}

// HasIncidents returns a boolean if a field has been set.
func (o *EntityResponseDataRelationships) HasIncidents() bool {
	return o != nil && o.Incidents != nil
}

// SetIncidents gets a reference to the given EntityResponseDataRelationshipsIncidents and assigns it to the Incidents field.
func (o *EntityResponseDataRelationships) SetIncidents(v EntityResponseDataRelationshipsIncidents) {
	o.Incidents = &v
}

// GetOncalls returns the Oncalls field value if set, zero value otherwise.
func (o *EntityResponseDataRelationships) GetOncalls() EntityResponseDataRelationshipsOncalls {
	if o == nil || o.Oncalls == nil {
		var ret EntityResponseDataRelationshipsOncalls
		return ret
	}
	return *o.Oncalls
}

// GetOncallsOk returns a tuple with the Oncalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseDataRelationships) GetOncallsOk() (*EntityResponseDataRelationshipsOncalls, bool) {
	if o == nil || o.Oncalls == nil {
		return nil, false
	}
	return o.Oncalls, true
}

// HasOncalls returns a boolean if a field has been set.
func (o *EntityResponseDataRelationships) HasOncalls() bool {
	return o != nil && o.Oncalls != nil
}

// SetOncalls gets a reference to the given EntityResponseDataRelationshipsOncalls and assigns it to the Oncalls field.
func (o *EntityResponseDataRelationships) SetOncalls(v EntityResponseDataRelationshipsOncalls) {
	o.Oncalls = &v
}

// GetRawSchema returns the RawSchema field value if set, zero value otherwise.
func (o *EntityResponseDataRelationships) GetRawSchema() EntityResponseDataRelationshipsRawSchema {
	if o == nil || o.RawSchema == nil {
		var ret EntityResponseDataRelationshipsRawSchema
		return ret
	}
	return *o.RawSchema
}

// GetRawSchemaOk returns a tuple with the RawSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseDataRelationships) GetRawSchemaOk() (*EntityResponseDataRelationshipsRawSchema, bool) {
	if o == nil || o.RawSchema == nil {
		return nil, false
	}
	return o.RawSchema, true
}

// HasRawSchema returns a boolean if a field has been set.
func (o *EntityResponseDataRelationships) HasRawSchema() bool {
	return o != nil && o.RawSchema != nil
}

// SetRawSchema gets a reference to the given EntityResponseDataRelationshipsRawSchema and assigns it to the RawSchema field.
func (o *EntityResponseDataRelationships) SetRawSchema(v EntityResponseDataRelationshipsRawSchema) {
	o.RawSchema = &v
}

// GetRelatedEntities returns the RelatedEntities field value if set, zero value otherwise.
func (o *EntityResponseDataRelationships) GetRelatedEntities() EntityResponseDataRelationshipsRelatedEntities {
	if o == nil || o.RelatedEntities == nil {
		var ret EntityResponseDataRelationshipsRelatedEntities
		return ret
	}
	return *o.RelatedEntities
}

// GetRelatedEntitiesOk returns a tuple with the RelatedEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseDataRelationships) GetRelatedEntitiesOk() (*EntityResponseDataRelationshipsRelatedEntities, bool) {
	if o == nil || o.RelatedEntities == nil {
		return nil, false
	}
	return o.RelatedEntities, true
}

// HasRelatedEntities returns a boolean if a field has been set.
func (o *EntityResponseDataRelationships) HasRelatedEntities() bool {
	return o != nil && o.RelatedEntities != nil
}

// SetRelatedEntities gets a reference to the given EntityResponseDataRelationshipsRelatedEntities and assigns it to the RelatedEntities field.
func (o *EntityResponseDataRelationships) SetRelatedEntities(v EntityResponseDataRelationshipsRelatedEntities) {
	o.RelatedEntities = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *EntityResponseDataRelationships) GetSchema() EntityResponseDataRelationshipsSchema {
	if o == nil || o.Schema == nil {
		var ret EntityResponseDataRelationshipsSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseDataRelationships) GetSchemaOk() (*EntityResponseDataRelationshipsSchema, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *EntityResponseDataRelationships) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given EntityResponseDataRelationshipsSchema and assigns it to the Schema field.
func (o *EntityResponseDataRelationships) SetSchema(v EntityResponseDataRelationshipsSchema) {
	o.Schema = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Incidents != nil {
		toSerialize["incidents"] = o.Incidents
	}
	if o.Oncalls != nil {
		toSerialize["oncalls"] = o.Oncalls
	}
	if o.RawSchema != nil {
		toSerialize["rawSchema"] = o.RawSchema
	}
	if o.RelatedEntities != nil {
		toSerialize["relatedEntities"] = o.RelatedEntities
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityResponseDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Incidents       *EntityResponseDataRelationshipsIncidents       `json:"incidents,omitempty"`
		Oncalls         *EntityResponseDataRelationshipsOncalls         `json:"oncalls,omitempty"`
		RawSchema       *EntityResponseDataRelationshipsRawSchema       `json:"rawSchema,omitempty"`
		RelatedEntities *EntityResponseDataRelationshipsRelatedEntities `json:"relatedEntities,omitempty"`
		Schema          *EntityResponseDataRelationshipsSchema          `json:"schema,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incidents", "oncalls", "rawSchema", "relatedEntities", "schema"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Incidents != nil && all.Incidents.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Incidents = all.Incidents
	if all.Oncalls != nil && all.Oncalls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Oncalls = all.Oncalls
	if all.RawSchema != nil && all.RawSchema.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RawSchema = all.RawSchema
	if all.RelatedEntities != nil && all.RelatedEntities.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.RelatedEntities = all.RelatedEntities
	if all.Schema != nil && all.Schema.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Schema = all.Schema

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
