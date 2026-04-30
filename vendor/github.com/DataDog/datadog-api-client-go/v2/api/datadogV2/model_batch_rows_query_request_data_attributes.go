// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BatchRowsQueryRequestDataAttributes Attributes for a batch rows query request.
type BatchRowsQueryRequestDataAttributes struct {
	// List of row identifiers to query from the reference table.
	RowIds []string `json:"row_ids"`
	// Unique identifier of the reference table to query.
	TableId string `json:"table_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBatchRowsQueryRequestDataAttributes instantiates a new BatchRowsQueryRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBatchRowsQueryRequestDataAttributes(rowIds []string, tableId string) *BatchRowsQueryRequestDataAttributes {
	this := BatchRowsQueryRequestDataAttributes{}
	this.RowIds = rowIds
	this.TableId = tableId
	return &this
}

// NewBatchRowsQueryRequestDataAttributesWithDefaults instantiates a new BatchRowsQueryRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBatchRowsQueryRequestDataAttributesWithDefaults() *BatchRowsQueryRequestDataAttributes {
	this := BatchRowsQueryRequestDataAttributes{}
	return &this
}

// GetRowIds returns the RowIds field value.
func (o *BatchRowsQueryRequestDataAttributes) GetRowIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RowIds
}

// GetRowIdsOk returns a tuple with the RowIds field value
// and a boolean to check if the value has been set.
func (o *BatchRowsQueryRequestDataAttributes) GetRowIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RowIds, true
}

// SetRowIds sets field value.
func (o *BatchRowsQueryRequestDataAttributes) SetRowIds(v []string) {
	o.RowIds = v
}

// GetTableId returns the TableId field value.
func (o *BatchRowsQueryRequestDataAttributes) GetTableId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableId
}

// GetTableIdOk returns a tuple with the TableId field value
// and a boolean to check if the value has been set.
func (o *BatchRowsQueryRequestDataAttributes) GetTableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableId, true
}

// SetTableId sets field value.
func (o *BatchRowsQueryRequestDataAttributes) SetTableId(v string) {
	o.TableId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o BatchRowsQueryRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["row_ids"] = o.RowIds
	toSerialize["table_id"] = o.TableId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BatchRowsQueryRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RowIds  *[]string `json:"row_ids"`
		TableId *string   `json:"table_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.RowIds == nil {
		return fmt.Errorf("required field row_ids missing")
	}
	if all.TableId == nil {
		return fmt.Errorf("required field table_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"row_ids", "table_id"})
	} else {
		return err
	}
	o.RowIds = *all.RowIds
	o.TableId = *all.TableId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
