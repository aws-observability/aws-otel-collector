// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultFileRef Reference to a file attached to a Synthetic test request.
type SyntheticsTestResultFileRef struct {
	// Storage bucket key where the file is stored.
	BucketKey *string `json:"bucket_key,omitempty"`
	// Encoding of the file contents.
	Encoding *string `json:"encoding,omitempty"`
	// File name.
	Name *string `json:"name,omitempty"`
	// File size in bytes.
	Size *int64 `json:"size,omitempty"`
	// File MIME type.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultFileRef instantiates a new SyntheticsTestResultFileRef object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultFileRef() *SyntheticsTestResultFileRef {
	this := SyntheticsTestResultFileRef{}
	return &this
}

// NewSyntheticsTestResultFileRefWithDefaults instantiates a new SyntheticsTestResultFileRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultFileRefWithDefaults() *SyntheticsTestResultFileRef {
	this := SyntheticsTestResultFileRef{}
	return &this
}

// GetBucketKey returns the BucketKey field value if set, zero value otherwise.
func (o *SyntheticsTestResultFileRef) GetBucketKey() string {
	if o == nil || o.BucketKey == nil {
		var ret string
		return ret
	}
	return *o.BucketKey
}

// GetBucketKeyOk returns a tuple with the BucketKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultFileRef) GetBucketKeyOk() (*string, bool) {
	if o == nil || o.BucketKey == nil {
		return nil, false
	}
	return o.BucketKey, true
}

// HasBucketKey returns a boolean if a field has been set.
func (o *SyntheticsTestResultFileRef) HasBucketKey() bool {
	return o != nil && o.BucketKey != nil
}

// SetBucketKey gets a reference to the given string and assigns it to the BucketKey field.
func (o *SyntheticsTestResultFileRef) SetBucketKey(v string) {
	o.BucketKey = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *SyntheticsTestResultFileRef) GetEncoding() string {
	if o == nil || o.Encoding == nil {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultFileRef) GetEncodingOk() (*string, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *SyntheticsTestResultFileRef) HasEncoding() bool {
	return o != nil && o.Encoding != nil
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *SyntheticsTestResultFileRef) SetEncoding(v string) {
	o.Encoding = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTestResultFileRef) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultFileRef) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTestResultFileRef) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTestResultFileRef) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SyntheticsTestResultFileRef) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultFileRef) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SyntheticsTestResultFileRef) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *SyntheticsTestResultFileRef) SetSize(v int64) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsTestResultFileRef) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultFileRef) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsTestResultFileRef) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SyntheticsTestResultFileRef) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultFileRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BucketKey != nil {
		toSerialize["bucket_key"] = o.BucketKey
	}
	if o.Encoding != nil {
		toSerialize["encoding"] = o.Encoding
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
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
func (o *SyntheticsTestResultFileRef) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BucketKey *string `json:"bucket_key,omitempty"`
		Encoding  *string `json:"encoding,omitempty"`
		Name      *string `json:"name,omitempty"`
		Size      *int64  `json:"size,omitempty"`
		Type      *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bucket_key", "encoding", "name", "size", "type"})
	} else {
		return err
	}
	o.BucketKey = all.BucketKey
	o.Encoding = all.Encoding
	o.Name = all.Name
	o.Size = all.Size
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
