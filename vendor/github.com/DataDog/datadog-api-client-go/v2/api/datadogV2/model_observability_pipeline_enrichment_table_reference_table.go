// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineEnrichmentTableReferenceTable Uses a Datadog reference table to enrich logs.
type ObservabilityPipelineEnrichmentTableReferenceTable struct {
	// Name of the environment variable or secret that holds the Datadog application key used to access the reference table.
	AppKeyKey *string `json:"app_key_key,omitempty"`
	// List of column names to include from the reference table. If not provided, all columns are included.
	Columns []string `json:"columns,omitempty"`
	// Path to the field in the log event to match against the reference table.
	KeyField string `json:"key_field"`
	// The unique identifier of the reference table.
	TableId string `json:"table_id"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineEnrichmentTableReferenceTable instantiates a new ObservabilityPipelineEnrichmentTableReferenceTable object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineEnrichmentTableReferenceTable(keyField string, tableId string) *ObservabilityPipelineEnrichmentTableReferenceTable {
	this := ObservabilityPipelineEnrichmentTableReferenceTable{}
	this.KeyField = keyField
	this.TableId = tableId
	return &this
}

// NewObservabilityPipelineEnrichmentTableReferenceTableWithDefaults instantiates a new ObservabilityPipelineEnrichmentTableReferenceTable object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineEnrichmentTableReferenceTableWithDefaults() *ObservabilityPipelineEnrichmentTableReferenceTable {
	this := ObservabilityPipelineEnrichmentTableReferenceTable{}
	return &this
}

// GetAppKeyKey returns the AppKeyKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) GetAppKeyKey() string {
	if o == nil || o.AppKeyKey == nil {
		var ret string
		return ret
	}
	return *o.AppKeyKey
}

// GetAppKeyKeyOk returns a tuple with the AppKeyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) GetAppKeyKeyOk() (*string, bool) {
	if o == nil || o.AppKeyKey == nil {
		return nil, false
	}
	return o.AppKeyKey, true
}

// HasAppKeyKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) HasAppKeyKey() bool {
	return o != nil && o.AppKeyKey != nil
}

// SetAppKeyKey gets a reference to the given string and assigns it to the AppKeyKey field.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) SetAppKeyKey(v string) {
	o.AppKeyKey = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) GetColumns() []string {
	if o == nil || o.Columns == nil {
		var ret []string
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) GetColumnsOk() (*[]string, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return &o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) HasColumns() bool {
	return o != nil && o.Columns != nil
}

// SetColumns gets a reference to the given []string and assigns it to the Columns field.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) SetColumns(v []string) {
	o.Columns = v
}

// GetKeyField returns the KeyField field value.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) GetKeyField() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.KeyField
}

// GetKeyFieldOk returns a tuple with the KeyField field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) GetKeyFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyField, true
}

// SetKeyField sets field value.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) SetKeyField(v string) {
	o.KeyField = v
}

// GetTableId returns the TableId field value.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) GetTableId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableId
}

// GetTableIdOk returns a tuple with the TableId field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) GetTableIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableId, true
}

// SetTableId sets field value.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) SetTableId(v string) {
	o.TableId = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineEnrichmentTableReferenceTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AppKeyKey != nil {
		toSerialize["app_key_key"] = o.AppKeyKey
	}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	toSerialize["key_field"] = o.KeyField
	toSerialize["table_id"] = o.TableId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineEnrichmentTableReferenceTable) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppKeyKey *string  `json:"app_key_key,omitempty"`
		Columns   []string `json:"columns,omitempty"`
		KeyField  *string  `json:"key_field"`
		TableId   *string  `json:"table_id"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.KeyField == nil {
		return fmt.Errorf("required field key_field missing")
	}
	if all.TableId == nil {
		return fmt.Errorf("required field table_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"app_key_key", "columns", "key_field", "table_id"})
	} else {
		return err
	}
	o.AppKeyKey = all.AppKeyKey
	o.Columns = all.Columns
	o.KeyField = *all.KeyField
	o.TableId = *all.TableId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
