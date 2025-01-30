// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListHistoricalJobsResponse List of historical jobs.
type ListHistoricalJobsResponse struct {
	// Array containing the list of historical jobs.
	Data []HistoricalJobResponseData `json:"data,omitempty"`
	// Metadata about the list of jobs.
	Meta *HistoricalJobListMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListHistoricalJobsResponse instantiates a new ListHistoricalJobsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListHistoricalJobsResponse() *ListHistoricalJobsResponse {
	this := ListHistoricalJobsResponse{}
	return &this
}

// NewListHistoricalJobsResponseWithDefaults instantiates a new ListHistoricalJobsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListHistoricalJobsResponseWithDefaults() *ListHistoricalJobsResponse {
	this := ListHistoricalJobsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListHistoricalJobsResponse) GetData() []HistoricalJobResponseData {
	if o == nil || o.Data == nil {
		var ret []HistoricalJobResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHistoricalJobsResponse) GetDataOk() (*[]HistoricalJobResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListHistoricalJobsResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []HistoricalJobResponseData and assigns it to the Data field.
func (o *ListHistoricalJobsResponse) SetData(v []HistoricalJobResponseData) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListHistoricalJobsResponse) GetMeta() HistoricalJobListMeta {
	if o == nil || o.Meta == nil {
		var ret HistoricalJobListMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHistoricalJobsResponse) GetMetaOk() (*HistoricalJobListMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListHistoricalJobsResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given HistoricalJobListMeta and assigns it to the Meta field.
func (o *ListHistoricalJobsResponse) SetMeta(v HistoricalJobListMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListHistoricalJobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListHistoricalJobsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data []HistoricalJobResponseData `json:"data,omitempty"`
		Meta *HistoricalJobListMeta      `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
