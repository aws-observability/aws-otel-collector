// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BatchRowsQueryResponseDataRelationships Relationships of the batch rows query response data.
type BatchRowsQueryResponseDataRelationships struct {
	// Relationship data containing the list of matching rows.
	Rows *BatchRowsQueryResponseDataRelationshipsRows `json:"rows,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBatchRowsQueryResponseDataRelationships instantiates a new BatchRowsQueryResponseDataRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBatchRowsQueryResponseDataRelationships() *BatchRowsQueryResponseDataRelationships {
	this := BatchRowsQueryResponseDataRelationships{}
	return &this
}

// NewBatchRowsQueryResponseDataRelationshipsWithDefaults instantiates a new BatchRowsQueryResponseDataRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBatchRowsQueryResponseDataRelationshipsWithDefaults() *BatchRowsQueryResponseDataRelationships {
	this := BatchRowsQueryResponseDataRelationships{}
	return &this
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *BatchRowsQueryResponseDataRelationships) GetRows() BatchRowsQueryResponseDataRelationshipsRows {
	if o == nil || o.Rows == nil {
		var ret BatchRowsQueryResponseDataRelationshipsRows
		return ret
	}
	return *o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchRowsQueryResponseDataRelationships) GetRowsOk() (*BatchRowsQueryResponseDataRelationshipsRows, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *BatchRowsQueryResponseDataRelationships) HasRows() bool {
	return o != nil && o.Rows != nil
}

// SetRows gets a reference to the given BatchRowsQueryResponseDataRelationshipsRows and assigns it to the Rows field.
func (o *BatchRowsQueryResponseDataRelationships) SetRows(v BatchRowsQueryResponseDataRelationshipsRows) {
	o.Rows = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BatchRowsQueryResponseDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BatchRowsQueryResponseDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Rows *BatchRowsQueryResponseDataRelationshipsRows `json:"rows,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"rows"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Rows != nil && all.Rows.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Rows = all.Rows

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
