// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityRelationships Entity relationships.
type EntityRelationships struct {
	// Entity to incidents relationship.
	Incidents *EntityToIncidents `json:"incidents,omitempty"`
	// Entity to oncalls relationship.
	Oncall *EntityToOncalls `json:"oncall,omitempty"`
	// Entity to raw schema relationship.
	RawSchema *EntityToRawSchema `json:"rawSchema,omitempty"`
	// Entity to related entities relationship.
	RelatedEntities *EntityToRelatedEntities `json:"relatedEntities,omitempty"`
	// Entity to detail schema relationship.
	Schema *EntityToSchema `json:"schema,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityRelationships instantiates a new EntityRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityRelationships() *EntityRelationships {
	this := EntityRelationships{}
	return &this
}

// NewEntityRelationshipsWithDefaults instantiates a new EntityRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityRelationshipsWithDefaults() *EntityRelationships {
	this := EntityRelationships{}
	return &this
}

// GetIncidents returns the Incidents field value if set, zero value otherwise.
func (o *EntityRelationships) GetIncidents() EntityToIncidents {
	if o == nil || o.Incidents == nil {
		var ret EntityToIncidents
		return ret
	}
	return *o.Incidents
}

// GetIncidentsOk returns a tuple with the Incidents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRelationships) GetIncidentsOk() (*EntityToIncidents, bool) {
	if o == nil || o.Incidents == nil {
		return nil, false
	}
	return o.Incidents, true
}

// HasIncidents returns a boolean if a field has been set.
func (o *EntityRelationships) HasIncidents() bool {
	return o != nil && o.Incidents != nil
}

// SetIncidents gets a reference to the given EntityToIncidents and assigns it to the Incidents field.
func (o *EntityRelationships) SetIncidents(v EntityToIncidents) {
	o.Incidents = &v
}

// GetOncall returns the Oncall field value if set, zero value otherwise.
func (o *EntityRelationships) GetOncall() EntityToOncalls {
	if o == nil || o.Oncall == nil {
		var ret EntityToOncalls
		return ret
	}
	return *o.Oncall
}

// GetOncallOk returns a tuple with the Oncall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRelationships) GetOncallOk() (*EntityToOncalls, bool) {
	if o == nil || o.Oncall == nil {
		return nil, false
	}
	return o.Oncall, true
}

// HasOncall returns a boolean if a field has been set.
func (o *EntityRelationships) HasOncall() bool {
	return o != nil && o.Oncall != nil
}

// SetOncall gets a reference to the given EntityToOncalls and assigns it to the Oncall field.
func (o *EntityRelationships) SetOncall(v EntityToOncalls) {
	o.Oncall = &v
}

// GetRawSchema returns the RawSchema field value if set, zero value otherwise.
func (o *EntityRelationships) GetRawSchema() EntityToRawSchema {
	if o == nil || o.RawSchema == nil {
		var ret EntityToRawSchema
		return ret
	}
	return *o.RawSchema
}

// GetRawSchemaOk returns a tuple with the RawSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRelationships) GetRawSchemaOk() (*EntityToRawSchema, bool) {
	if o == nil || o.RawSchema == nil {
		return nil, false
	}
	return o.RawSchema, true
}

// HasRawSchema returns a boolean if a field has been set.
func (o *EntityRelationships) HasRawSchema() bool {
	return o != nil && o.RawSchema != nil
}

// SetRawSchema gets a reference to the given EntityToRawSchema and assigns it to the RawSchema field.
func (o *EntityRelationships) SetRawSchema(v EntityToRawSchema) {
	o.RawSchema = &v
}

// GetRelatedEntities returns the RelatedEntities field value if set, zero value otherwise.
func (o *EntityRelationships) GetRelatedEntities() EntityToRelatedEntities {
	if o == nil || o.RelatedEntities == nil {
		var ret EntityToRelatedEntities
		return ret
	}
	return *o.RelatedEntities
}

// GetRelatedEntitiesOk returns a tuple with the RelatedEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRelationships) GetRelatedEntitiesOk() (*EntityToRelatedEntities, bool) {
	if o == nil || o.RelatedEntities == nil {
		return nil, false
	}
	return o.RelatedEntities, true
}

// HasRelatedEntities returns a boolean if a field has been set.
func (o *EntityRelationships) HasRelatedEntities() bool {
	return o != nil && o.RelatedEntities != nil
}

// SetRelatedEntities gets a reference to the given EntityToRelatedEntities and assigns it to the RelatedEntities field.
func (o *EntityRelationships) SetRelatedEntities(v EntityToRelatedEntities) {
	o.RelatedEntities = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *EntityRelationships) GetSchema() EntityToSchema {
	if o == nil || o.Schema == nil {
		var ret EntityToSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityRelationships) GetSchemaOk() (*EntityToSchema, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *EntityRelationships) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given EntityToSchema and assigns it to the Schema field.
func (o *EntityRelationships) SetSchema(v EntityToSchema) {
	o.Schema = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Incidents != nil {
		toSerialize["incidents"] = o.Incidents
	}
	if o.Oncall != nil {
		toSerialize["oncall"] = o.Oncall
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
func (o *EntityRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Incidents       *EntityToIncidents       `json:"incidents,omitempty"`
		Oncall          *EntityToOncalls         `json:"oncall,omitempty"`
		RawSchema       *EntityToRawSchema       `json:"rawSchema,omitempty"`
		RelatedEntities *EntityToRelatedEntities `json:"relatedEntities,omitempty"`
		Schema          *EntityToSchema          `json:"schema,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"incidents", "oncall", "rawSchema", "relatedEntities", "schema"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Incidents != nil && all.Incidents.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Incidents = all.Incidents
	if all.Oncall != nil && all.Oncall.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Oncall = all.Oncall
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
