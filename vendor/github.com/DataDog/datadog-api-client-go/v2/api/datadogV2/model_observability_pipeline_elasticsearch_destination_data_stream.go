// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ObservabilityPipelineElasticsearchDestinationDataStream Configuration options for writing to Elasticsearch Data Streams instead of a fixed index.
type ObservabilityPipelineElasticsearchDestinationDataStream struct {
	// When `true`, automatically routes events to the appropriate data stream based on the event content.
	AutoRouting *bool `json:"auto_routing,omitempty"`
	// The data stream dataset. This groups events by their source or application.
	Dataset *string `json:"dataset,omitempty"`
	// The data stream type. This determines how events are categorized within the data stream.
	Dtype *string `json:"dtype,omitempty"`
	// The data stream namespace. This separates events into different environments or domains.
	Namespace *string `json:"namespace,omitempty"`
	// When `true`, synchronizes data stream fields with the Elasticsearch index mapping.
	SyncFields *bool `json:"sync_fields,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewObservabilityPipelineElasticsearchDestinationDataStream instantiates a new ObservabilityPipelineElasticsearchDestinationDataStream object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewObservabilityPipelineElasticsearchDestinationDataStream() *ObservabilityPipelineElasticsearchDestinationDataStream {
	this := ObservabilityPipelineElasticsearchDestinationDataStream{}
	return &this
}

// NewObservabilityPipelineElasticsearchDestinationDataStreamWithDefaults instantiates a new ObservabilityPipelineElasticsearchDestinationDataStream object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewObservabilityPipelineElasticsearchDestinationDataStreamWithDefaults() *ObservabilityPipelineElasticsearchDestinationDataStream {
	this := ObservabilityPipelineElasticsearchDestinationDataStream{}
	return &this
}

// GetAutoRouting returns the AutoRouting field value if set, zero value otherwise.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetAutoRouting() bool {
	if o == nil || o.AutoRouting == nil {
		var ret bool
		return ret
	}
	return *o.AutoRouting
}

// GetAutoRoutingOk returns a tuple with the AutoRouting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetAutoRoutingOk() (*bool, bool) {
	if o == nil || o.AutoRouting == nil {
		return nil, false
	}
	return o.AutoRouting, true
}

// HasAutoRouting returns a boolean if a field has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) HasAutoRouting() bool {
	return o != nil && o.AutoRouting != nil
}

// SetAutoRouting gets a reference to the given bool and assigns it to the AutoRouting field.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) SetAutoRouting(v bool) {
	o.AutoRouting = &v
}

// GetDataset returns the Dataset field value if set, zero value otherwise.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetDataset() string {
	if o == nil || o.Dataset == nil {
		var ret string
		return ret
	}
	return *o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetDatasetOk() (*string, bool) {
	if o == nil || o.Dataset == nil {
		return nil, false
	}
	return o.Dataset, true
}

// HasDataset returns a boolean if a field has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) HasDataset() bool {
	return o != nil && o.Dataset != nil
}

// SetDataset gets a reference to the given string and assigns it to the Dataset field.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) SetDataset(v string) {
	o.Dataset = &v
}

// GetDtype returns the Dtype field value if set, zero value otherwise.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetDtype() string {
	if o == nil || o.Dtype == nil {
		var ret string
		return ret
	}
	return *o.Dtype
}

// GetDtypeOk returns a tuple with the Dtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetDtypeOk() (*string, bool) {
	if o == nil || o.Dtype == nil {
		return nil, false
	}
	return o.Dtype, true
}

// HasDtype returns a boolean if a field has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) HasDtype() bool {
	return o != nil && o.Dtype != nil
}

// SetDtype gets a reference to the given string and assigns it to the Dtype field.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) SetDtype(v string) {
	o.Dtype = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) HasNamespace() bool {
	return o != nil && o.Namespace != nil
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) SetNamespace(v string) {
	o.Namespace = &v
}

// GetSyncFields returns the SyncFields field value if set, zero value otherwise.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetSyncFields() bool {
	if o == nil || o.SyncFields == nil {
		var ret bool
		return ret
	}
	return *o.SyncFields
}

// GetSyncFieldsOk returns a tuple with the SyncFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) GetSyncFieldsOk() (*bool, bool) {
	if o == nil || o.SyncFields == nil {
		return nil, false
	}
	return o.SyncFields, true
}

// HasSyncFields returns a boolean if a field has been set.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) HasSyncFields() bool {
	return o != nil && o.SyncFields != nil
}

// SetSyncFields gets a reference to the given bool and assigns it to the SyncFields field.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) SetSyncFields(v bool) {
	o.SyncFields = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ObservabilityPipelineElasticsearchDestinationDataStream) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AutoRouting != nil {
		toSerialize["auto_routing"] = o.AutoRouting
	}
	if o.Dataset != nil {
		toSerialize["dataset"] = o.Dataset
	}
	if o.Dtype != nil {
		toSerialize["dtype"] = o.Dtype
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.SyncFields != nil {
		toSerialize["sync_fields"] = o.SyncFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ObservabilityPipelineElasticsearchDestinationDataStream) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoRouting *bool   `json:"auto_routing,omitempty"`
		Dataset     *string `json:"dataset,omitempty"`
		Dtype       *string `json:"dtype,omitempty"`
		Namespace   *string `json:"namespace,omitempty"`
		SyncFields  *bool   `json:"sync_fields,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_routing", "dataset", "dtype", "namespace", "sync_fields"})
	} else {
		return err
	}
	o.AutoRouting = all.AutoRouting
	o.Dataset = all.Dataset
	o.Dtype = all.Dtype
	o.Namespace = all.Namespace
	o.SyncFields = all.SyncFields

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
