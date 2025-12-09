// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatasetResponse **Datasets Object Constraints**
// - **Tag Limit per Dataset**:
//   - Each restricted dataset supports a maximum of 10 key:value pairs per product.
//
// - **Tag Key Rules per Telemetry Type**:
//   - Only one tag key or attribute may be used to define access within a single telemetry type.
//   - The same or different tag key may be used across different telemetry types.
//
// - **Tag Value Uniqueness**:
//   - Tag values must be unique within a single dataset.
//   - A tag value used in one dataset cannot be reused in another dataset of the same telemetry type.
type DatasetResponse struct {
	// Dataset metadata and configuration(s).
	Attributes *DatasetAttributesResponse `json:"attributes,omitempty"`
	// Unique identifier for the dataset.
	Id *string `json:"id,omitempty"`
	// Resource type, always set to `dataset`.
	Type *DatasetType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatasetResponse instantiates a new DatasetResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatasetResponse() *DatasetResponse {
	this := DatasetResponse{}
	var typeVar DatasetType = DATASETTYPE_DATASET
	this.Type = &typeVar
	return &this
}

// NewDatasetResponseWithDefaults instantiates a new DatasetResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatasetResponseWithDefaults() *DatasetResponse {
	this := DatasetResponse{}
	var typeVar DatasetType = DATASETTYPE_DATASET
	this.Type = &typeVar
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *DatasetResponse) GetAttributes() DatasetAttributesResponse {
	if o == nil || o.Attributes == nil {
		var ret DatasetAttributesResponse
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetResponse) GetAttributesOk() (*DatasetAttributesResponse, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *DatasetResponse) HasAttributes() bool {
	return o != nil && o.Attributes != nil
}

// SetAttributes gets a reference to the given DatasetAttributesResponse and assigns it to the Attributes field.
func (o *DatasetResponse) SetAttributes(v DatasetAttributesResponse) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DatasetResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DatasetResponse) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DatasetResponse) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DatasetResponse) GetType() DatasetType {
	if o == nil || o.Type == nil {
		var ret DatasetType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetResponse) GetTypeOk() (*DatasetType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DatasetResponse) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given DatasetType and assigns it to the Type field.
func (o *DatasetResponse) SetType(v DatasetType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatasetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatasetResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attributes *DatasetAttributesResponse `json:"attributes,omitempty"`
		Id         *string                    `json:"id,omitempty"`
		Type       *DatasetType               `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attributes", "id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attributes != nil && all.Attributes.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attributes = all.Attributes
	o.Id = all.Id
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
