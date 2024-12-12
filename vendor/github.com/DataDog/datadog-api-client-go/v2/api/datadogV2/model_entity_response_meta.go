// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityResponseMeta Entity metadata.
type EntityResponseMeta struct {
	// Total entities count.
	Count *int64 `json:"count,omitempty"`
	// Total included data count.
	IncludeCount *int64 `json:"includeCount,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewEntityResponseMeta instantiates a new EntityResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityResponseMeta() *EntityResponseMeta {
	this := EntityResponseMeta{}
	return &this
}

// NewEntityResponseMetaWithDefaults instantiates a new EntityResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityResponseMetaWithDefaults() *EntityResponseMeta {
	this := EntityResponseMeta{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *EntityResponseMeta) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseMeta) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *EntityResponseMeta) HasCount() bool {
	return o != nil && o.Count != nil
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *EntityResponseMeta) SetCount(v int64) {
	o.Count = &v
}

// GetIncludeCount returns the IncludeCount field value if set, zero value otherwise.
func (o *EntityResponseMeta) GetIncludeCount() int64 {
	if o == nil || o.IncludeCount == nil {
		var ret int64
		return ret
	}
	return *o.IncludeCount
}

// GetIncludeCountOk returns a tuple with the IncludeCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityResponseMeta) GetIncludeCountOk() (*int64, bool) {
	if o == nil || o.IncludeCount == nil {
		return nil, false
	}
	return o.IncludeCount, true
}

// HasIncludeCount returns a boolean if a field has been set.
func (o *EntityResponseMeta) HasIncludeCount() bool {
	return o != nil && o.IncludeCount != nil
}

// SetIncludeCount gets a reference to the given int64 and assigns it to the IncludeCount field.
func (o *EntityResponseMeta) SetIncludeCount(v int64) {
	o.IncludeCount = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.IncludeCount != nil {
		toSerialize["includeCount"] = o.IncludeCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Count        *int64 `json:"count,omitempty"`
		IncludeCount *int64 `json:"includeCount,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"count", "includeCount"})
	} else {
		return err
	}
	o.Count = all.Count
	o.IncludeCount = all.IncludeCount

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
