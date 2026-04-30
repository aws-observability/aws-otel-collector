// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentUserDefinedFieldListMeta Pagination metadata for the user-defined field list response.
type IncidentUserDefinedFieldListMeta struct {
	// The offset of the current page.
	Offset *int64 `json:"offset,omitempty"`
	// The total number of items in the current page.
	Size *int64 `json:"size,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentUserDefinedFieldListMeta instantiates a new IncidentUserDefinedFieldListMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentUserDefinedFieldListMeta() *IncidentUserDefinedFieldListMeta {
	this := IncidentUserDefinedFieldListMeta{}
	return &this
}

// NewIncidentUserDefinedFieldListMetaWithDefaults instantiates a new IncidentUserDefinedFieldListMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentUserDefinedFieldListMetaWithDefaults() *IncidentUserDefinedFieldListMeta {
	this := IncidentUserDefinedFieldListMeta{}
	return &this
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *IncidentUserDefinedFieldListMeta) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldListMeta) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldListMeta) HasOffset() bool {
	return o != nil && o.Offset != nil
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *IncidentUserDefinedFieldListMeta) SetOffset(v int64) {
	o.Offset = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *IncidentUserDefinedFieldListMeta) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentUserDefinedFieldListMeta) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *IncidentUserDefinedFieldListMeta) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *IncidentUserDefinedFieldListMeta) SetSize(v int64) {
	o.Size = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentUserDefinedFieldListMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentUserDefinedFieldListMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Offset *int64 `json:"offset,omitempty"`
		Size   *int64 `json:"size,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"offset", "size"})
	} else {
		return err
	}
	o.Offset = all.Offset
	o.Size = all.Size

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
