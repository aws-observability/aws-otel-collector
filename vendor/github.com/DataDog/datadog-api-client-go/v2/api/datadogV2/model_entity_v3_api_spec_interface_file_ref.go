// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3APISpecInterfaceFileRef The definition of `EntityV3APISpecInterfaceFileRef` object.
type EntityV3APISpecInterfaceFileRef struct {
	// The reference to the API definition file.
	FileRef *string `json:"fileRef,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3APISpecInterfaceFileRef instantiates a new EntityV3APISpecInterfaceFileRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3APISpecInterfaceFileRef() *EntityV3APISpecInterfaceFileRef {
	this := EntityV3APISpecInterfaceFileRef{}
	return &this
}

// NewEntityV3APISpecInterfaceFileRefWithDefaults instantiates a new EntityV3APISpecInterfaceFileRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3APISpecInterfaceFileRefWithDefaults() *EntityV3APISpecInterfaceFileRef {
	this := EntityV3APISpecInterfaceFileRef{}
	return &this
}

// GetFileRef returns the FileRef field value if set, zero value otherwise.
func (o *EntityV3APISpecInterfaceFileRef) GetFileRef() string {
	if o == nil || o.FileRef == nil {
		var ret string
		return ret
	}
	return *o.FileRef
}

// GetFileRefOk returns a tuple with the FileRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3APISpecInterfaceFileRef) GetFileRefOk() (*string, bool) {
	if o == nil || o.FileRef == nil {
		return nil, false
	}
	return o.FileRef, true
}

// HasFileRef returns a boolean if a field has been set.
func (o *EntityV3APISpecInterfaceFileRef) HasFileRef() bool {
	return o != nil && o.FileRef != nil
}

// SetFileRef gets a reference to the given string and assigns it to the FileRef field.
func (o *EntityV3APISpecInterfaceFileRef) SetFileRef(v string) {
	o.FileRef = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3APISpecInterfaceFileRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.FileRef != nil {
		toSerialize["fileRef"] = o.FileRef
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3APISpecInterfaceFileRef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		FileRef *string `json:"fileRef,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	o.FileRef = all.FileRef

	return nil
}
