// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"encoding/json"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GCPSTSServiceAccountUpdateRequest Service account info.
type GCPSTSServiceAccountUpdateRequest struct {
	// Data on your service account.
	Data *GCPSTSServiceAccountUpdateRequestData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewGCPSTSServiceAccountUpdateRequest instantiates a new GCPSTSServiceAccountUpdateRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGCPSTSServiceAccountUpdateRequest() *GCPSTSServiceAccountUpdateRequest {
	this := GCPSTSServiceAccountUpdateRequest{}
	return &this
}

// NewGCPSTSServiceAccountUpdateRequestWithDefaults instantiates a new GCPSTSServiceAccountUpdateRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGCPSTSServiceAccountUpdateRequestWithDefaults() *GCPSTSServiceAccountUpdateRequest {
	this := GCPSTSServiceAccountUpdateRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GCPSTSServiceAccountUpdateRequest) GetData() GCPSTSServiceAccountUpdateRequestData {
	if o == nil || o.Data == nil {
		var ret GCPSTSServiceAccountUpdateRequestData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GCPSTSServiceAccountUpdateRequest) GetDataOk() (*GCPSTSServiceAccountUpdateRequestData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GCPSTSServiceAccountUpdateRequest) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given GCPSTSServiceAccountUpdateRequestData and assigns it to the Data field.
func (o *GCPSTSServiceAccountUpdateRequest) SetData(v GCPSTSServiceAccountUpdateRequestData) {
	o.Data = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GCPSTSServiceAccountUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return json.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return json.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GCPSTSServiceAccountUpdateRequest) UnmarshalJSON(bytes []byte) (err error) {
	raw := map[string]interface{}{}
	all := struct {
		Data *GCPSTSServiceAccountUpdateRequestData `json:"data,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Data = all.Data
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}