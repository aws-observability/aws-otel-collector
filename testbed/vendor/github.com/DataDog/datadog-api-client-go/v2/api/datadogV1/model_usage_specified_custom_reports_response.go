// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageSpecifiedCustomReportsResponse Returns available specified custom reports.
type UsageSpecifiedCustomReportsResponse struct {
	// Response containing date and type for specified custom reports.
	Data *UsageSpecifiedCustomReportsData `json:"data,omitempty"`
	// The object containing document metadata.
	Meta *UsageSpecifiedCustomReportsMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageSpecifiedCustomReportsResponse instantiates a new UsageSpecifiedCustomReportsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageSpecifiedCustomReportsResponse() *UsageSpecifiedCustomReportsResponse {
	this := UsageSpecifiedCustomReportsResponse{}
	return &this
}

// NewUsageSpecifiedCustomReportsResponseWithDefaults instantiates a new UsageSpecifiedCustomReportsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageSpecifiedCustomReportsResponseWithDefaults() *UsageSpecifiedCustomReportsResponse {
	this := UsageSpecifiedCustomReportsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UsageSpecifiedCustomReportsResponse) GetData() UsageSpecifiedCustomReportsData {
	if o == nil || o.Data == nil {
		var ret UsageSpecifiedCustomReportsData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSpecifiedCustomReportsResponse) GetDataOk() (*UsageSpecifiedCustomReportsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UsageSpecifiedCustomReportsResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given UsageSpecifiedCustomReportsData and assigns it to the Data field.
func (o *UsageSpecifiedCustomReportsResponse) SetData(v UsageSpecifiedCustomReportsData) {
	o.Data = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UsageSpecifiedCustomReportsResponse) GetMeta() UsageSpecifiedCustomReportsMeta {
	if o == nil || o.Meta == nil {
		var ret UsageSpecifiedCustomReportsMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageSpecifiedCustomReportsResponse) GetMetaOk() (*UsageSpecifiedCustomReportsMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UsageSpecifiedCustomReportsResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given UsageSpecifiedCustomReportsMeta and assigns it to the Meta field.
func (o *UsageSpecifiedCustomReportsResponse) SetMeta(v UsageSpecifiedCustomReportsMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageSpecifiedCustomReportsResponse) MarshalJSON() ([]byte, error) {
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
func (o *UsageSpecifiedCustomReportsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *UsageSpecifiedCustomReportsData `json:"data,omitempty"`
		Meta *UsageSpecifiedCustomReportsMeta `json:"meta,omitempty"`
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
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
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
