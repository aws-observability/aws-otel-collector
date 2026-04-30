// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedFieldListResponse Response containing a list of incident user-defined fields.
type IncidentUserDefinedFieldListResponse struct {
	// An array of user-defined field objects.
	Data []IncidentUserDefinedFieldResponseData `json:"data"`
	// Pagination metadata for the user-defined field list response.
	Meta IncidentUserDefinedFieldListMeta `json:"meta"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedFieldListResponse instantiates a new IncidentUserDefinedFieldListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedFieldListResponse(data []IncidentUserDefinedFieldResponseData, meta IncidentUserDefinedFieldListMeta) *IncidentUserDefinedFieldListResponse {
	this := IncidentUserDefinedFieldListResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewIncidentUserDefinedFieldListResponseWithDefaults instantiates a new IncidentUserDefinedFieldListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedFieldListResponseWithDefaults() *IncidentUserDefinedFieldListResponse {
	this := IncidentUserDefinedFieldListResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentUserDefinedFieldListResponse) GetData() []IncidentUserDefinedFieldResponseData {
	if o == nil {
		var ret []IncidentUserDefinedFieldResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldListResponse) GetDataOk() (*[]IncidentUserDefinedFieldResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentUserDefinedFieldListResponse) SetData(v []IncidentUserDefinedFieldResponseData) {
	o.Data = v
}

// GetMeta returns the Meta field value.
func (o *IncidentUserDefinedFieldListResponse) GetMeta() IncidentUserDefinedFieldListMeta {
	if o == nil {
		var ret IncidentUserDefinedFieldListMeta
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldListResponse) GetMetaOk() (*IncidentUserDefinedFieldListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value.
func (o *IncidentUserDefinedFieldListResponse) SetMeta(v IncidentUserDefinedFieldListMeta) {
	o.Meta = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedFieldListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedFieldListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data *[]IncidentUserDefinedFieldResponseData `json:"data"`
		Meta *IncidentUserDefinedFieldListMeta       `json:"meta"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Meta == nil {
		return fmt.Errorf("required field meta missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	if all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = *all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
