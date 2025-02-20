// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestRequestBodyFile Object describing a file to be used as part of the request in the test.
type SyntheticsTestRequestBodyFile struct {
	// Bucket key of the file.
	BucketKey *string `json:"bucketKey,omitempty"`
	// Content of the file.
	Content *string `json:"content,omitempty"`
	// Name of the file.
	Name *string `json:"name,omitempty"`
	// Original name of the file.
	OriginalFileName *string `json:"originalFileName,omitempty"`
	// Size of the file.
	Size *int64 `json:"size,omitempty"`
	// Type of the file.
	Type *string `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestRequestBodyFile instantiates a new SyntheticsTestRequestBodyFile object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestRequestBodyFile() *SyntheticsTestRequestBodyFile {
	this := SyntheticsTestRequestBodyFile{}
	return &this
}

// NewSyntheticsTestRequestBodyFileWithDefaults instantiates a new SyntheticsTestRequestBodyFile object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestRequestBodyFileWithDefaults() *SyntheticsTestRequestBodyFile {
	this := SyntheticsTestRequestBodyFile{}
	return &this
}

// GetBucketKey returns the BucketKey field value if set, zero value otherwise.
func (o *SyntheticsTestRequestBodyFile) GetBucketKey() string {
	if o == nil || o.BucketKey == nil {
		var ret string
		return ret
	}
	return *o.BucketKey
}

// GetBucketKeyOk returns a tuple with the BucketKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequestBodyFile) GetBucketKeyOk() (*string, bool) {
	if o == nil || o.BucketKey == nil {
		return nil, false
	}
	return o.BucketKey, true
}

// HasBucketKey returns a boolean if a field has been set.
func (o *SyntheticsTestRequestBodyFile) HasBucketKey() bool {
	return o != nil && o.BucketKey != nil
}

// SetBucketKey gets a reference to the given string and assigns it to the BucketKey field.
func (o *SyntheticsTestRequestBodyFile) SetBucketKey(v string) {
	o.BucketKey = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *SyntheticsTestRequestBodyFile) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequestBodyFile) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *SyntheticsTestRequestBodyFile) HasContent() bool {
	return o != nil && o.Content != nil
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *SyntheticsTestRequestBodyFile) SetContent(v string) {
	o.Content = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SyntheticsTestRequestBodyFile) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequestBodyFile) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SyntheticsTestRequestBodyFile) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SyntheticsTestRequestBodyFile) SetName(v string) {
	o.Name = &v
}

// GetOriginalFileName returns the OriginalFileName field value if set, zero value otherwise.
func (o *SyntheticsTestRequestBodyFile) GetOriginalFileName() string {
	if o == nil || o.OriginalFileName == nil {
		var ret string
		return ret
	}
	return *o.OriginalFileName
}

// GetOriginalFileNameOk returns a tuple with the OriginalFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequestBodyFile) GetOriginalFileNameOk() (*string, bool) {
	if o == nil || o.OriginalFileName == nil {
		return nil, false
	}
	return o.OriginalFileName, true
}

// HasOriginalFileName returns a boolean if a field has been set.
func (o *SyntheticsTestRequestBodyFile) HasOriginalFileName() bool {
	return o != nil && o.OriginalFileName != nil
}

// SetOriginalFileName gets a reference to the given string and assigns it to the OriginalFileName field.
func (o *SyntheticsTestRequestBodyFile) SetOriginalFileName(v string) {
	o.OriginalFileName = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *SyntheticsTestRequestBodyFile) GetSize() int64 {
	if o == nil || o.Size == nil {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequestBodyFile) GetSizeOk() (*int64, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *SyntheticsTestRequestBodyFile) HasSize() bool {
	return o != nil && o.Size != nil
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *SyntheticsTestRequestBodyFile) SetSize(v int64) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsTestRequestBodyFile) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequestBodyFile) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsTestRequestBodyFile) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SyntheticsTestRequestBodyFile) SetType(v string) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestRequestBodyFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BucketKey != nil {
		toSerialize["bucketKey"] = o.BucketKey
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OriginalFileName != nil {
		toSerialize["originalFileName"] = o.OriginalFileName
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
func (o *SyntheticsTestRequestBodyFile) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BucketKey        *string `json:"bucketKey,omitempty"`
		Content          *string `json:"content,omitempty"`
		Name             *string `json:"name,omitempty"`
		OriginalFileName *string `json:"originalFileName,omitempty"`
		Size             *int64  `json:"size,omitempty"`
		Type             *string `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bucketKey", "content", "name", "originalFileName", "size", "type"})
	} else {
		return err
	}
	o.BucketKey = all.BucketKey
	o.Content = all.Content
	o.Name = all.Name
	o.OriginalFileName = all.OriginalFileName
	o.Size = all.Size
	o.Type = all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
