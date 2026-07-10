// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineDatabricksZerobusDestination The `databricks_zerobus` destination sends logs to Databricks using the Zerobus ingestion API, streaming data directly into your Databricks Lakehouse.
//
// **Supported pipeline types:** Logs, rehydration
type ObservabilityPipelineDatabricksZerobusDestination struct {
	// OAuth credentials for authenticating with the Databricks Zerobus ingestion API.
	Auth ObservabilityPipelineDatabricksZerobusDestinationAuth `json:"auth"`
	// Configuration for buffer settings on destination components.
	Buffer *ObservabilityPipelineBufferOptions `json:"buffer,omitempty"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// Name of the environment variable or the secret identifier that references the Databricks Zerobus ingestion endpoint, which is used to stream data directly into your Databricks Lakehouse.
	IngestionEndpointKey *string `json:"ingestion_endpoint_key,omitempty"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// The fully qualified name of your target Databricks table. Make sure this table already exists in your Databricks workspace before deploying.
	TableName string `json:"table_name"`
	// The destination type. The value must be `databricks_zerobus`.
	Type ObservabilityPipelineDatabricksZerobusDestinationType `json:"type"`
	// Name of the environment variable or the secret identifier that references your Databricks workspace URL, which is used to communicate with the Unity Catalog API.
	UnityCatalogEndpointKey *string `json:"unity_catalog_endpoint_key,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineDatabricksZerobusDestination instantiates a new ObservabilityPipelineDatabricksZerobusDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineDatabricksZerobusDestination(auth ObservabilityPipelineDatabricksZerobusDestinationAuth, id string, inputs []string, tableName string, typeVar ObservabilityPipelineDatabricksZerobusDestinationType) *ObservabilityPipelineDatabricksZerobusDestination {
	this := ObservabilityPipelineDatabricksZerobusDestination{}
	this.Auth = auth
	this.Id = id
	this.Inputs = inputs
	this.TableName = tableName
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineDatabricksZerobusDestinationWithDefaults instantiates a new ObservabilityPipelineDatabricksZerobusDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineDatabricksZerobusDestinationWithDefaults() *ObservabilityPipelineDatabricksZerobusDestination {
	this := ObservabilityPipelineDatabricksZerobusDestination{}
	var typeVar ObservabilityPipelineDatabricksZerobusDestinationType = OBSERVABILITYPIPELINEDATABRICKSZEROBUSDESTINATIONTYPE_DATABRICKS_ZEROBUS
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetAuth() ObservabilityPipelineDatabricksZerobusDestinationAuth {
	if o == nil {
		var ret ObservabilityPipelineDatabricksZerobusDestinationAuth
		return ret
	}
	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetAuthOk() (*ObservabilityPipelineDatabricksZerobusDestinationAuth, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Auth, true
}

// SetAuth sets field value.
func (o *ObservabilityPipelineDatabricksZerobusDestination) SetAuth(v ObservabilityPipelineDatabricksZerobusDestinationAuth) {
	o.Auth = v
}

// GetBuffer returns the Buffer field value if set, zero value otherwise.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetBuffer() ObservabilityPipelineBufferOptions {
	if o == nil || o.Buffer == nil {
		var ret ObservabilityPipelineBufferOptions
		return ret
	}
	return *o.Buffer
}

// GetBufferOk returns a tuple with the Buffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetBufferOk() (*ObservabilityPipelineBufferOptions, bool) {
	if o == nil || o.Buffer == nil {
		return nil, false
	}
	return o.Buffer, true
}

// HasBuffer returns a boolean if a field has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestination) HasBuffer() bool {
	return o != nil && o.Buffer != nil
}

// SetBuffer gets a reference to the given ObservabilityPipelineBufferOptions and assigns it to the Buffer field.
func (o *ObservabilityPipelineDatabricksZerobusDestination) SetBuffer(v ObservabilityPipelineBufferOptions) {
	o.Buffer = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineDatabricksZerobusDestination) SetId(v string) {
	o.Id = v
}

// GetIngestionEndpointKey returns the IngestionEndpointKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetIngestionEndpointKey() string {
	if o == nil || o.IngestionEndpointKey == nil {
		var ret string
		return ret
	}
	return *o.IngestionEndpointKey
}

// GetIngestionEndpointKeyOk returns a tuple with the IngestionEndpointKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetIngestionEndpointKeyOk() (*string, bool) {
	if o == nil || o.IngestionEndpointKey == nil {
		return nil, false
	}
	return o.IngestionEndpointKey, true
}

// HasIngestionEndpointKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestination) HasIngestionEndpointKey() bool {
	return o != nil && o.IngestionEndpointKey != nil
}

// SetIngestionEndpointKey gets a reference to the given string and assigns it to the IngestionEndpointKey field.
func (o *ObservabilityPipelineDatabricksZerobusDestination) SetIngestionEndpointKey(v string) {
	o.IngestionEndpointKey = &v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineDatabricksZerobusDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetTableName returns the TableName field value.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetTableName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetTableNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TableName, true
}

// SetTableName sets field value.
func (o *ObservabilityPipelineDatabricksZerobusDestination) SetTableName(v string) {
	o.TableName = v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetType() ObservabilityPipelineDatabricksZerobusDestinationType {
	if o == nil {
		var ret ObservabilityPipelineDatabricksZerobusDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetTypeOk() (*ObservabilityPipelineDatabricksZerobusDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineDatabricksZerobusDestination) SetType(v ObservabilityPipelineDatabricksZerobusDestinationType) {
	o.Type = v
}

// GetUnityCatalogEndpointKey returns the UnityCatalogEndpointKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetUnityCatalogEndpointKey() string {
	if o == nil || o.UnityCatalogEndpointKey == nil {
		var ret string
		return ret
	}
	return *o.UnityCatalogEndpointKey
}

// GetUnityCatalogEndpointKeyOk returns a tuple with the UnityCatalogEndpointKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestination) GetUnityCatalogEndpointKeyOk() (*string, bool) {
	if o == nil || o.UnityCatalogEndpointKey == nil {
		return nil, false
	}
	return o.UnityCatalogEndpointKey, true
}

// HasUnityCatalogEndpointKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineDatabricksZerobusDestination) HasUnityCatalogEndpointKey() bool {
	return o != nil && o.UnityCatalogEndpointKey != nil
}

// SetUnityCatalogEndpointKey gets a reference to the given string and assigns it to the UnityCatalogEndpointKey field.
func (o *ObservabilityPipelineDatabricksZerobusDestination) SetUnityCatalogEndpointKey(v string) {
	o.UnityCatalogEndpointKey = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineDatabricksZerobusDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auth"] = o.Auth
	if o.Buffer != nil {
		toSerialize["buffer"] = o.Buffer
	}
	toSerialize["id"] = o.Id
	if o.IngestionEndpointKey != nil {
		toSerialize["ingestion_endpoint_key"] = o.IngestionEndpointKey
	}
	toSerialize["inputs"] = o.Inputs
	toSerialize["table_name"] = o.TableName
	toSerialize["type"] = o.Type
	if o.UnityCatalogEndpointKey != nil {
		toSerialize["unity_catalog_endpoint_key"] = o.UnityCatalogEndpointKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineDatabricksZerobusDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth                    *ObservabilityPipelineDatabricksZerobusDestinationAuth `json:"auth"`
		Buffer                  *ObservabilityPipelineBufferOptions                    `json:"buffer,omitempty"`
		Id                      *string                                                `json:"id"`
		IngestionEndpointKey    *string                                                `json:"ingestion_endpoint_key,omitempty"`
		Inputs                  *[]string                                              `json:"inputs"`
		TableName               *string                                                `json:"table_name"`
		Type                    *ObservabilityPipelineDatabricksZerobusDestinationType `json:"type"`
		UnityCatalogEndpointKey *string                                                `json:"unity_catalog_endpoint_key,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Auth == nil {
		return fmt.Errorf("required field auth missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.TableName == nil {
		return fmt.Errorf("required field table_name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "buffer", "id", "ingestion_endpoint_key", "inputs", "table_name", "type", "unity_catalog_endpoint_key"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = *all.Auth
	o.Buffer = all.Buffer
	o.Id = *all.Id
	o.IngestionEndpointKey = all.IngestionEndpointKey
	o.Inputs = *all.Inputs
	o.TableName = *all.TableName
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	o.UnityCatalogEndpointKey = all.UnityCatalogEndpointKey

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
