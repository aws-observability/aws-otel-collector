// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentHandleResponse Response payload for a single incident handle, including the handle data and related resources.
type IncidentHandleResponse struct {
	// Data object representing an incident handle in a response.
	Data IncidentHandleDataResponse `json:"data"`
	// Included related resources
	Included []IncidentHandleIncludedItemResponse `json:"included,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentHandleResponse instantiates a new IncidentHandleResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentHandleResponse(data IncidentHandleDataResponse) *IncidentHandleResponse {
	this := IncidentHandleResponse{}
	this.Data = data
	return &this
}

// NewIncidentHandleResponseWithDefaults instantiates a new IncidentHandleResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentHandleResponseWithDefaults() *IncidentHandleResponse {
	this := IncidentHandleResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentHandleResponse) GetData() IncidentHandleDataResponse {
	if o == nil {
		var ret IncidentHandleDataResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentHandleResponse) GetDataOk() (*IncidentHandleDataResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentHandleResponse) SetData(v IncidentHandleDataResponse) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentHandleResponse) GetIncluded() []IncidentHandleIncludedItemResponse {
	if o == nil || o.Included == nil {
		var ret []IncidentHandleIncludedItemResponse
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentHandleResponse) GetIncludedOk() (*[]IncidentHandleIncludedItemResponse, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentHandleResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []IncidentHandleIncludedItemResponse and assigns it to the Included field.
func (o *IncidentHandleResponse) SetIncluded(v []IncidentHandleIncludedItemResponse) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentHandleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentHandleResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *IncidentHandleDataResponse          `json:"data"`
		Included []IncidentHandleIncludedItemResponse `json:"included,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data
	o.Included = all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
