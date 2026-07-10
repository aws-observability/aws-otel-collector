// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineClickhouseDestination The `clickhouse` destination sends log events to a ClickHouse database table over HTTP.
//
// **Supported pipeline types:** logs.
type ObservabilityPipelineClickhouseDestination struct {
	// HTTP Basic Authentication credentials for the ClickHouse destination.
	// When `strategy` is `basic`, provide `username_key` and `password_key` that reference environment variables or secrets containing the credentials.
	Auth *ObservabilityPipelineClickhouseDestinationAuth `json:"auth,omitempty"`
	// Batching configuration for ClickHouse inserts.
	Batch *ObservabilityPipelineClickhouseDestinationBatch `json:"batch,omitempty"`
	// Batch encoding configuration for the ClickHouse destination.
	// Required when `format` is `arrow_stream`. The `codec` field must be set to `arrow_stream`.
	BatchEncoding *ObservabilityPipelineClickhouseDestinationBatchEncoding `json:"batch_encoding,omitempty"`
	// Configuration for buffer settings on destination components.
	Buffer *ObservabilityPipelineBufferOptions `json:"buffer,omitempty"`
	// Compression setting for outbound HTTP requests to ClickHouse.
	// Can be specified as a shorthand string (`"gzip"` or `"none"`) or as an object
	// with an `algorithm` field and an optional `level` (gzip only, 1–9).
	Compression *ObservabilityPipelineClickhouseDestinationCompression `json:"compression,omitempty"`
	// Optional ClickHouse database name. If omitted, the user's default database on the ClickHouse server is used.
	Database *string `json:"database,omitempty"`
	// When `true`, enables flexible DateTime parsing on the ClickHouse server side.
	DateTimeBestEffort *bool `json:"date_time_best_effort,omitempty"`
	// Name of the environment variable or secret that contains the ClickHouse HTTP endpoint URL.
	// Defaults to `DESTINATION_CLICKHOUSE_ENDPOINT_URL` (prefixed with `DD_OP_` at runtime).
	EndpointUrlKey *string `json:"endpoint_url_key,omitempty"`
	// Insert format for events sent to ClickHouse.
	// - `json_each_row`: Maps event fields to columns by name (ClickHouse `JSONEachRow`).
	// - `json_as_object`: Inserts each event into a single `Object('json')` / `JSON` column (ClickHouse `JSONAsObject`).
	// - `json_as_string`: Inserts each event into a single `String`-typed column as raw JSON (ClickHouse `JSONAsString`).
	// - `arrow_stream`: Batches events using Apache Arrow IPC streaming format. Requires `batch_encoding`.
	Format *ObservabilityPipelineClickhouseDestinationFormat `json:"format,omitempty"`
	// The unique identifier for this component.
	Id string `json:"id"`
	// A list of component IDs whose output is used as the `input` for this component.
	Inputs []string `json:"inputs"`
	// When `true`, fields not present in the target table schema are dropped instead of causing insert errors.
	// When unset, the ClickHouse server's own `input_format_skip_unknown_fields` setting applies.
	SkipUnknownFields datadog.NullableBool `json:"skip_unknown_fields,omitempty"`
	// Target ClickHouse table name. Events are inserted into this table.
	Table string `json:"table"`
	// Configuration for enabling TLS encryption between the pipeline component and external services.
	Tls *ObservabilityPipelineTls `json:"tls,omitempty"`
	// The destination type. The value must be `clickhouse`.
	Type ObservabilityPipelineClickhouseDestinationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineClickhouseDestination instantiates a new ObservabilityPipelineClickhouseDestination object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineClickhouseDestination(id string, inputs []string, table string, typeVar ObservabilityPipelineClickhouseDestinationType) *ObservabilityPipelineClickhouseDestination {
	this := ObservabilityPipelineClickhouseDestination{}
	this.Id = id
	this.Inputs = inputs
	this.Table = table
	this.Type = typeVar
	return &this
}

// NewObservabilityPipelineClickhouseDestinationWithDefaults instantiates a new ObservabilityPipelineClickhouseDestination object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineClickhouseDestinationWithDefaults() *ObservabilityPipelineClickhouseDestination {
	this := ObservabilityPipelineClickhouseDestination{}
	var typeVar ObservabilityPipelineClickhouseDestinationType = OBSERVABILITYPIPELINECLICKHOUSEDESTINATIONTYPE_CLICKHOUSE
	this.Type = typeVar
	return &this
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *ObservabilityPipelineClickhouseDestination) GetAuth() ObservabilityPipelineClickhouseDestinationAuth {
	if o == nil || o.Auth == nil {
		var ret ObservabilityPipelineClickhouseDestinationAuth
		return ret
	}
	return *o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetAuthOk() (*ObservabilityPipelineClickhouseDestinationAuth, bool) {
	if o == nil || o.Auth == nil {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *ObservabilityPipelineClickhouseDestination) HasAuth() bool {
	return o != nil && o.Auth != nil
}

// SetAuth gets a reference to the given ObservabilityPipelineClickhouseDestinationAuth and assigns it to the Auth field.
func (o *ObservabilityPipelineClickhouseDestination) SetAuth(v ObservabilityPipelineClickhouseDestinationAuth) {
	o.Auth = &v
}

// GetBatch returns the Batch field value if set, zero value otherwise.
func (o *ObservabilityPipelineClickhouseDestination) GetBatch() ObservabilityPipelineClickhouseDestinationBatch {
	if o == nil || o.Batch == nil {
		var ret ObservabilityPipelineClickhouseDestinationBatch
		return ret
	}
	return *o.Batch
}

// GetBatchOk returns a tuple with the Batch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetBatchOk() (*ObservabilityPipelineClickhouseDestinationBatch, bool) {
	if o == nil || o.Batch == nil {
		return nil, false
	}
	return o.Batch, true
}

// HasBatch returns a boolean if a field has been set.
func (o *ObservabilityPipelineClickhouseDestination) HasBatch() bool {
	return o != nil && o.Batch != nil
}

// SetBatch gets a reference to the given ObservabilityPipelineClickhouseDestinationBatch and assigns it to the Batch field.
func (o *ObservabilityPipelineClickhouseDestination) SetBatch(v ObservabilityPipelineClickhouseDestinationBatch) {
	o.Batch = &v
}

// GetBatchEncoding returns the BatchEncoding field value if set, zero value otherwise.
func (o *ObservabilityPipelineClickhouseDestination) GetBatchEncoding() ObservabilityPipelineClickhouseDestinationBatchEncoding {
	if o == nil || o.BatchEncoding == nil {
		var ret ObservabilityPipelineClickhouseDestinationBatchEncoding
		return ret
	}
	return *o.BatchEncoding
}

// GetBatchEncodingOk returns a tuple with the BatchEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetBatchEncodingOk() (*ObservabilityPipelineClickhouseDestinationBatchEncoding, bool) {
	if o == nil || o.BatchEncoding == nil {
		return nil, false
	}
	return o.BatchEncoding, true
}

// HasBatchEncoding returns a boolean if a field has been set.
func (o *ObservabilityPipelineClickhouseDestination) HasBatchEncoding() bool {
	return o != nil && o.BatchEncoding != nil
}

// SetBatchEncoding gets a reference to the given ObservabilityPipelineClickhouseDestinationBatchEncoding and assigns it to the BatchEncoding field.
func (o *ObservabilityPipelineClickhouseDestination) SetBatchEncoding(v ObservabilityPipelineClickhouseDestinationBatchEncoding) {
	o.BatchEncoding = &v
}

// GetBuffer returns the Buffer field value if set, zero value otherwise.
func (o *ObservabilityPipelineClickhouseDestination) GetBuffer() ObservabilityPipelineBufferOptions {
	if o == nil || o.Buffer == nil {
		var ret ObservabilityPipelineBufferOptions
		return ret
	}
	return *o.Buffer
}

// GetBufferOk returns a tuple with the Buffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetBufferOk() (*ObservabilityPipelineBufferOptions, bool) {
	if o == nil || o.Buffer == nil {
		return nil, false
	}
	return o.Buffer, true
}

// HasBuffer returns a boolean if a field has been set.
func (o *ObservabilityPipelineClickhouseDestination) HasBuffer() bool {
	return o != nil && o.Buffer != nil
}

// SetBuffer gets a reference to the given ObservabilityPipelineBufferOptions and assigns it to the Buffer field.
func (o *ObservabilityPipelineClickhouseDestination) SetBuffer(v ObservabilityPipelineBufferOptions) {
	o.Buffer = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *ObservabilityPipelineClickhouseDestination) GetCompression() ObservabilityPipelineClickhouseDestinationCompression {
	if o == nil || o.Compression == nil {
		var ret ObservabilityPipelineClickhouseDestinationCompression
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetCompressionOk() (*ObservabilityPipelineClickhouseDestinationCompression, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *ObservabilityPipelineClickhouseDestination) HasCompression() bool {
	return o != nil && o.Compression != nil
}

// SetCompression gets a reference to the given ObservabilityPipelineClickhouseDestinationCompression and assigns it to the Compression field.
func (o *ObservabilityPipelineClickhouseDestination) SetCompression(v ObservabilityPipelineClickhouseDestinationCompression) {
	o.Compression = &v
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *ObservabilityPipelineClickhouseDestination) GetDatabase() string {
	if o == nil || o.Database == nil {
		var ret string
		return ret
	}
	return *o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetDatabaseOk() (*string, bool) {
	if o == nil || o.Database == nil {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *ObservabilityPipelineClickhouseDestination) HasDatabase() bool {
	return o != nil && o.Database != nil
}

// SetDatabase gets a reference to the given string and assigns it to the Database field.
func (o *ObservabilityPipelineClickhouseDestination) SetDatabase(v string) {
	o.Database = &v
}

// GetDateTimeBestEffort returns the DateTimeBestEffort field value if set, zero value otherwise.
func (o *ObservabilityPipelineClickhouseDestination) GetDateTimeBestEffort() bool {
	if o == nil || o.DateTimeBestEffort == nil {
		var ret bool
		return ret
	}
	return *o.DateTimeBestEffort
}

// GetDateTimeBestEffortOk returns a tuple with the DateTimeBestEffort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetDateTimeBestEffortOk() (*bool, bool) {
	if o == nil || o.DateTimeBestEffort == nil {
		return nil, false
	}
	return o.DateTimeBestEffort, true
}

// HasDateTimeBestEffort returns a boolean if a field has been set.
func (o *ObservabilityPipelineClickhouseDestination) HasDateTimeBestEffort() bool {
	return o != nil && o.DateTimeBestEffort != nil
}

// SetDateTimeBestEffort gets a reference to the given bool and assigns it to the DateTimeBestEffort field.
func (o *ObservabilityPipelineClickhouseDestination) SetDateTimeBestEffort(v bool) {
	o.DateTimeBestEffort = &v
}

// GetEndpointUrlKey returns the EndpointUrlKey field value if set, zero value otherwise.
func (o *ObservabilityPipelineClickhouseDestination) GetEndpointUrlKey() string {
	if o == nil || o.EndpointUrlKey == nil {
		var ret string
		return ret
	}
	return *o.EndpointUrlKey
}

// GetEndpointUrlKeyOk returns a tuple with the EndpointUrlKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetEndpointUrlKeyOk() (*string, bool) {
	if o == nil || o.EndpointUrlKey == nil {
		return nil, false
	}
	return o.EndpointUrlKey, true
}

// HasEndpointUrlKey returns a boolean if a field has been set.
func (o *ObservabilityPipelineClickhouseDestination) HasEndpointUrlKey() bool {
	return o != nil && o.EndpointUrlKey != nil
}

// SetEndpointUrlKey gets a reference to the given string and assigns it to the EndpointUrlKey field.
func (o *ObservabilityPipelineClickhouseDestination) SetEndpointUrlKey(v string) {
	o.EndpointUrlKey = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ObservabilityPipelineClickhouseDestination) GetFormat() ObservabilityPipelineClickhouseDestinationFormat {
	if o == nil || o.Format == nil {
		var ret ObservabilityPipelineClickhouseDestinationFormat
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetFormatOk() (*ObservabilityPipelineClickhouseDestinationFormat, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ObservabilityPipelineClickhouseDestination) HasFormat() bool {
	return o != nil && o.Format != nil
}

// SetFormat gets a reference to the given ObservabilityPipelineClickhouseDestinationFormat and assigns it to the Format field.
func (o *ObservabilityPipelineClickhouseDestination) SetFormat(v ObservabilityPipelineClickhouseDestinationFormat) {
	o.Format = &v
}

// GetId returns the Id field value.
func (o *ObservabilityPipelineClickhouseDestination) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *ObservabilityPipelineClickhouseDestination) SetId(v string) {
	o.Id = v
}

// GetInputs returns the Inputs field value.
func (o *ObservabilityPipelineClickhouseDestination) GetInputs() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetInputsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// SetInputs sets field value.
func (o *ObservabilityPipelineClickhouseDestination) SetInputs(v []string) {
	o.Inputs = v
}

// GetSkipUnknownFields returns the SkipUnknownFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ObservabilityPipelineClickhouseDestination) GetSkipUnknownFields() bool {
	if o == nil || o.SkipUnknownFields.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SkipUnknownFields.Get()
}

// GetSkipUnknownFieldsOk returns a tuple with the SkipUnknownFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ObservabilityPipelineClickhouseDestination) GetSkipUnknownFieldsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SkipUnknownFields.Get(), o.SkipUnknownFields.IsSet()
}

// HasSkipUnknownFields returns a boolean if a field has been set.
func (o *ObservabilityPipelineClickhouseDestination) HasSkipUnknownFields() bool {
	return o != nil && o.SkipUnknownFields.IsSet()
}

// SetSkipUnknownFields gets a reference to the given datadog.NullableBool and assigns it to the SkipUnknownFields field.
func (o *ObservabilityPipelineClickhouseDestination) SetSkipUnknownFields(v bool) {
	o.SkipUnknownFields.Set(&v)
}

// SetSkipUnknownFieldsNil sets the value for SkipUnknownFields to be an explicit nil.
func (o *ObservabilityPipelineClickhouseDestination) SetSkipUnknownFieldsNil() {
	o.SkipUnknownFields.Set(nil)
}

// UnsetSkipUnknownFields ensures that no value is present for SkipUnknownFields, not even an explicit nil.
func (o *ObservabilityPipelineClickhouseDestination) UnsetSkipUnknownFields() {
	o.SkipUnknownFields.Unset()
}

// GetTable returns the Table field value.
func (o *ObservabilityPipelineClickhouseDestination) GetTable() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Table
}

// GetTableOk returns a tuple with the Table field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetTableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Table, true
}

// SetTable sets field value.
func (o *ObservabilityPipelineClickhouseDestination) SetTable(v string) {
	o.Table = v
}

// GetTls returns the Tls field value if set, zero value otherwise.
func (o *ObservabilityPipelineClickhouseDestination) GetTls() ObservabilityPipelineTls {
	if o == nil || o.Tls == nil {
		var ret ObservabilityPipelineTls
		return ret
	}
	return *o.Tls
}

// GetTlsOk returns a tuple with the Tls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetTlsOk() (*ObservabilityPipelineTls, bool) {
	if o == nil || o.Tls == nil {
		return nil, false
	}
	return o.Tls, true
}

// HasTls returns a boolean if a field has been set.
func (o *ObservabilityPipelineClickhouseDestination) HasTls() bool {
	return o != nil && o.Tls != nil
}

// SetTls gets a reference to the given ObservabilityPipelineTls and assigns it to the Tls field.
func (o *ObservabilityPipelineClickhouseDestination) SetTls(v ObservabilityPipelineTls) {
	o.Tls = &v
}

// GetType returns the Type field value.
func (o *ObservabilityPipelineClickhouseDestination) GetType() ObservabilityPipelineClickhouseDestinationType {
	if o == nil {
		var ret ObservabilityPipelineClickhouseDestinationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineClickhouseDestination) GetTypeOk() (*ObservabilityPipelineClickhouseDestinationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ObservabilityPipelineClickhouseDestination) SetType(v ObservabilityPipelineClickhouseDestinationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineClickhouseDestination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Auth != nil {
		toSerialize["auth"] = o.Auth
	}
	if o.Batch != nil {
		toSerialize["batch"] = o.Batch
	}
	if o.BatchEncoding != nil {
		toSerialize["batch_encoding"] = o.BatchEncoding
	}
	if o.Buffer != nil {
		toSerialize["buffer"] = o.Buffer
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	if o.Database != nil {
		toSerialize["database"] = o.Database
	}
	if o.DateTimeBestEffort != nil {
		toSerialize["date_time_best_effort"] = o.DateTimeBestEffort
	}
	if o.EndpointUrlKey != nil {
		toSerialize["endpoint_url_key"] = o.EndpointUrlKey
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	toSerialize["id"] = o.Id
	toSerialize["inputs"] = o.Inputs
	if o.SkipUnknownFields.IsSet() {
		toSerialize["skip_unknown_fields"] = o.SkipUnknownFields.Get()
	}
	toSerialize["table"] = o.Table
	if o.Tls != nil {
		toSerialize["tls"] = o.Tls
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineClickhouseDestination) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Auth               *ObservabilityPipelineClickhouseDestinationAuth          `json:"auth,omitempty"`
		Batch              *ObservabilityPipelineClickhouseDestinationBatch         `json:"batch,omitempty"`
		BatchEncoding      *ObservabilityPipelineClickhouseDestinationBatchEncoding `json:"batch_encoding,omitempty"`
		Buffer             *ObservabilityPipelineBufferOptions                      `json:"buffer,omitempty"`
		Compression        *ObservabilityPipelineClickhouseDestinationCompression   `json:"compression,omitempty"`
		Database           *string                                                  `json:"database,omitempty"`
		DateTimeBestEffort *bool                                                    `json:"date_time_best_effort,omitempty"`
		EndpointUrlKey     *string                                                  `json:"endpoint_url_key,omitempty"`
		Format             *ObservabilityPipelineClickhouseDestinationFormat        `json:"format,omitempty"`
		Id                 *string                                                  `json:"id"`
		Inputs             *[]string                                                `json:"inputs"`
		SkipUnknownFields  datadog.NullableBool                                     `json:"skip_unknown_fields,omitempty"`
		Table              *string                                                  `json:"table"`
		Tls                *ObservabilityPipelineTls                                `json:"tls,omitempty"`
		Type               *ObservabilityPipelineClickhouseDestinationType          `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Inputs == nil {
		return fmt.Errorf("required field inputs missing")
	}
	if all.Table == nil {
		return fmt.Errorf("required field table missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auth", "batch", "batch_encoding", "buffer", "compression", "database", "date_time_best_effort", "endpoint_url_key", "format", "id", "inputs", "skip_unknown_fields", "table", "tls", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Auth != nil && all.Auth.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Auth = all.Auth
	if all.Batch != nil && all.Batch.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Batch = all.Batch
	if all.BatchEncoding != nil && all.BatchEncoding.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BatchEncoding = all.BatchEncoding
	o.Buffer = all.Buffer
	o.Compression = all.Compression
	o.Database = all.Database
	o.DateTimeBestEffort = all.DateTimeBestEffort
	o.EndpointUrlKey = all.EndpointUrlKey
	if all.Format != nil && !all.Format.IsValid() {
		hasInvalidField = true
	} else {
		o.Format = all.Format
	}
	o.Id = *all.Id
	o.Inputs = *all.Inputs
	o.SkipUnknownFields = all.SkipUnknownFields
	o.Table = *all.Table
	if all.Tls != nil && all.Tls.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tls = all.Tls
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
