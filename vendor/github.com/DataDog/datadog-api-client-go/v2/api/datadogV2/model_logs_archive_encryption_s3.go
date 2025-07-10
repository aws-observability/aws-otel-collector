// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LogsArchiveEncryptionS3 The S3 encryption settings.
type LogsArchiveEncryptionS3 struct {
	// An Amazon Resource Name (ARN) used to identify an AWS KMS key.
	Key *string `json:"key,omitempty"`
	// Type of S3 encryption for a destination.
	Type LogsArchiveEncryptionS3Type `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLogsArchiveEncryptionS3 instantiates a new LogsArchiveEncryptionS3 object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLogsArchiveEncryptionS3(typeVar LogsArchiveEncryptionS3Type) *LogsArchiveEncryptionS3 {
	this := LogsArchiveEncryptionS3{}
	this.Type = typeVar
	return &this
}

// NewLogsArchiveEncryptionS3WithDefaults instantiates a new LogsArchiveEncryptionS3 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLogsArchiveEncryptionS3WithDefaults() *LogsArchiveEncryptionS3 {
	this := LogsArchiveEncryptionS3{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *LogsArchiveEncryptionS3) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsArchiveEncryptionS3) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *LogsArchiveEncryptionS3) HasKey() bool {
	return o != nil && o.Key != nil
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *LogsArchiveEncryptionS3) SetKey(v string) {
	o.Key = &v
}

// GetType returns the Type field value.
func (o *LogsArchiveEncryptionS3) GetType() LogsArchiveEncryptionS3Type {
	if o == nil {
		var ret LogsArchiveEncryptionS3Type
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LogsArchiveEncryptionS3) GetTypeOk() (*LogsArchiveEncryptionS3Type, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *LogsArchiveEncryptionS3) SetType(v LogsArchiveEncryptionS3Type) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LogsArchiveEncryptionS3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LogsArchiveEncryptionS3) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Key  *string                      `json:"key,omitempty"`
		Type *LogsArchiveEncryptionS3Type `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"key", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Key = all.Key
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
