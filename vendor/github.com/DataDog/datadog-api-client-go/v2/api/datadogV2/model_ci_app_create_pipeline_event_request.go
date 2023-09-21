// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
<<<<<<< HEAD
	"encoding/json"
=======
	"github.com/goccy/go-json"
>>>>>>> main

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CIAppCreatePipelineEventRequest Request object.
type CIAppCreatePipelineEventRequest struct {
	// Data of the pipeline event to create.
	Data *CIAppCreatePipelineEventRequestData `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewCIAppCreatePipelineEventRequest instantiates a new CIAppCreatePipelineEventRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCIAppCreatePipelineEventRequest() *CIAppCreatePipelineEventRequest {
	this := CIAppCreatePipelineEventRequest{}
	return &this
}

// NewCIAppCreatePipelineEventRequestWithDefaults instantiates a new CIAppCreatePipelineEventRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCIAppCreatePipelineEventRequestWithDefaults() *CIAppCreatePipelineEventRequest {
	this := CIAppCreatePipelineEventRequest{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CIAppCreatePipelineEventRequest) GetData() CIAppCreatePipelineEventRequestData {
	if o == nil || o.Data == nil {
		var ret CIAppCreatePipelineEventRequestData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CIAppCreatePipelineEventRequest) GetDataOk() (*CIAppCreatePipelineEventRequestData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CIAppCreatePipelineEventRequest) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given CIAppCreatePipelineEventRequestData and assigns it to the Data field.
func (o *CIAppCreatePipelineEventRequest) SetData(v CIAppCreatePipelineEventRequestData) {
	o.Data = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CIAppCreatePipelineEventRequest) MarshalJSON() ([]byte, error) {
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
func (o *CIAppCreatePipelineEventRequest) UnmarshalJSON(bytes []byte) (err error) {
<<<<<<< HEAD
	raw := map[string]interface{}{}
=======
>>>>>>> main
	all := struct {
		Data *CIAppCreatePipelineEventRequestData `json:"data,omitempty"`
	}{}
	if err = json.Unmarshal(bytes, &all); err != nil {
<<<<<<< HEAD
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
		return nil
=======
		return json.Unmarshal(bytes, &o.UnparsedObject)
>>>>>>> main
	}
	additionalProperties := make(map[string]interface{})
	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data"})
	} else {
		return err
	}
<<<<<<< HEAD
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		err = json.Unmarshal(bytes, &raw)
		if err != nil {
			return err
		}
		o.UnparsedObject = raw
	}
	o.Data = all.Data
=======

	hasInvalidField := false
	if all.Data != nil && all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = all.Data

>>>>>>> main
	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

<<<<<<< HEAD
=======
	if hasInvalidField {
		return json.Unmarshal(bytes, &o.UnparsedObject)
	}

>>>>>>> main
	return nil
}
