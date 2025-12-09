// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetRequest **Datasets Object Constraints**
// - **Tag limit per dataset**:
//   - Each restricted dataset supports a maximum of 10 key:value pairs per product.
//
// - **Tag key rules per telemetry type**:
//   - Only one tag key or attribute may be used to define access within a single telemetry type.
//   - The same or different tag key may be used across different telemetry types.
//
// - **Tag value uniqueness**:
//   - Tag values must be unique within a single dataset.
//   - A tag value used in one dataset cannot be reused in another dataset of the same telemetry type.
type DatasetRequest struct {
	// Dataset metadata and configurations.
	Attributes DatasetAttributesRequest `json:"attributes"`
	// Resource type, always set to `dataset`.
	Type DatasetType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatasetRequest instantiates a new DatasetRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetRequest(attributes DatasetAttributesRequest, typeVar DatasetType) *DatasetRequest {
	this := DatasetRequest{}
	this.Attributes = attributes
	this.Type = typeVar
	return &this
}

// NewDatasetRequestWithDefaults instantiates a new DatasetRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetRequestWithDefaults() *DatasetRequest {
	this := DatasetRequest{}
	var typeVar DatasetType = DATASETTYPE_DATASET
	this.Type = typeVar
	return &this
}

// GetAttributes returns the Attributes field value.
func (o *DatasetRequest) GetAttributes() DatasetAttributesRequest {
	if o == nil {
		var ret DatasetAttributesRequest
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DatasetRequest) GetAttributesOk() (*DatasetAttributesRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value.
func (o *DatasetRequest) SetAttributes(v DatasetAttributesRequest) {
	o.Attributes = v
}

// GetType returns the Type field value.
func (o *DatasetRequest) GetType() DatasetType {
	if o == nil {
		var ret DatasetType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatasetRequest) GetTypeOk() (*DatasetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DatasetRequest) SetType(v DatasetType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attributes"] = o.Attributes
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *DatasetAttributesRequest `json:"attributes"`
		Type       *DatasetType              `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attributes == nil {
		return fmt.Errorf("required field attributes missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = *all.Attributes
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
