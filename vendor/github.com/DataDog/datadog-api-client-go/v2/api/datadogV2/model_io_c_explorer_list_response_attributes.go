// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IoCExplorerListResponseAttributes Attributes of the IoC Explorer list response.
type IoCExplorerListResponseAttributes struct {
	// List of indicators of compromise.
	Data []IoCIndicator `json:"data,omitempty"`
	// Response metadata.
	Metadata *IoCExplorerListResponseMetadata `json:"metadata,omitempty"`
	// Pagination information.
	Paging *IoCExplorerListResponsePaging `json:"paging,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIoCExplorerListResponseAttributes instantiates a new IoCExplorerListResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIoCExplorerListResponseAttributes() *IoCExplorerListResponseAttributes {
	this := IoCExplorerListResponseAttributes{}
	return &this
}

// NewIoCExplorerListResponseAttributesWithDefaults instantiates a new IoCExplorerListResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIoCExplorerListResponseAttributesWithDefaults() *IoCExplorerListResponseAttributes {
	this := IoCExplorerListResponseAttributes{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IoCExplorerListResponseAttributes) GetData() []IoCIndicator {
	if o == nil || o.Data == nil {
		var ret []IoCIndicator
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCExplorerListResponseAttributes) GetDataOk() (*[]IoCIndicator, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IoCExplorerListResponseAttributes) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []IoCIndicator and assigns it to the Data field.
func (o *IoCExplorerListResponseAttributes) SetData(v []IoCIndicator) {
	o.Data = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IoCExplorerListResponseAttributes) GetMetadata() IoCExplorerListResponseMetadata {
	if o == nil || o.Metadata == nil {
		var ret IoCExplorerListResponseMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCExplorerListResponseAttributes) GetMetadataOk() (*IoCExplorerListResponseMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IoCExplorerListResponseAttributes) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given IoCExplorerListResponseMetadata and assigns it to the Metadata field.
func (o *IoCExplorerListResponseAttributes) SetMetadata(v IoCExplorerListResponseMetadata) {
	o.Metadata = &v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *IoCExplorerListResponseAttributes) GetPaging() IoCExplorerListResponsePaging {
	if o == nil || o.Paging == nil {
		var ret IoCExplorerListResponsePaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCExplorerListResponseAttributes) GetPagingOk() (*IoCExplorerListResponsePaging, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *IoCExplorerListResponseAttributes) HasPaging() bool {
	return o != nil && o.Paging != nil
}

// SetPaging gets a reference to the given IoCExplorerListResponsePaging and assigns it to the Paging field.
func (o *IoCExplorerListResponseAttributes) SetPaging(v IoCExplorerListResponsePaging) {
	o.Paging = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IoCExplorerListResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IoCExplorerListResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     []IoCIndicator                   `json:"data,omitempty"`
		Metadata *IoCExplorerListResponseMetadata `json:"metadata,omitempty"`
		Paging   *IoCExplorerListResponsePaging   `json:"paging,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "metadata", "paging"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	if all.Paging != nil && all.Paging.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Paging = all.Paging

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
