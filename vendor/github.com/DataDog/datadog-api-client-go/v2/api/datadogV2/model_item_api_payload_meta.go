// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ItemApiPayloadMeta Additional metadata about a collection of datastore items, including pagination and schema information.
type ItemApiPayloadMeta struct {
	// Pagination information for a collection of datastore items.
	Page *ItemApiPayloadMetaPage `json:"page,omitempty"`
	// Schema information about the datastore, including its primary key and field definitions.
	Schema *ItemApiPayloadMetaSchema `json:"schema,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewItemApiPayloadMeta instantiates a new ItemApiPayloadMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewItemApiPayloadMeta() *ItemApiPayloadMeta {
	this := ItemApiPayloadMeta{}
	return &this
}

// NewItemApiPayloadMetaWithDefaults instantiates a new ItemApiPayloadMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewItemApiPayloadMetaWithDefaults() *ItemApiPayloadMeta {
	this := ItemApiPayloadMeta{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ItemApiPayloadMeta) GetPage() ItemApiPayloadMetaPage {
	if o == nil || o.Page == nil {
		var ret ItemApiPayloadMetaPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadMeta) GetPageOk() (*ItemApiPayloadMetaPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ItemApiPayloadMeta) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given ItemApiPayloadMetaPage and assigns it to the Page field.
func (o *ItemApiPayloadMeta) SetPage(v ItemApiPayloadMetaPage) {
	o.Page = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ItemApiPayloadMeta) GetSchema() ItemApiPayloadMetaSchema {
	if o == nil || o.Schema == nil {
		var ret ItemApiPayloadMetaSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemApiPayloadMeta) GetSchemaOk() (*ItemApiPayloadMetaSchema, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ItemApiPayloadMeta) HasSchema() bool {
	return o != nil && o.Schema != nil
}

// SetSchema gets a reference to the given ItemApiPayloadMetaSchema and assigns it to the Schema field.
func (o *ItemApiPayloadMeta) SetSchema(v ItemApiPayloadMetaSchema) {
	o.Schema = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ItemApiPayloadMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
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
func (o *ItemApiPayloadMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Page   *ItemApiPayloadMetaPage   `json:"page,omitempty"`
		Schema *ItemApiPayloadMetaSchema `json:"schema,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"page", "schema"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page
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
