// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestVersionChangeAttributes Attributes of a version change record.
type SyntheticsTestVersionChangeAttributes struct {
	// UUID of the user who created this version.
	AuthorUuid *string `json:"author_uuid,omitempty"`
	// List of metadata describing individual changes in this version.
	ChangeMetadata []SyntheticsTestVersionChangeMetadataItem `json:"change_metadata,omitempty"`
	// The sequential version number.
	VersionNumber *int64 `json:"version_number,omitempty"`
	// Timestamp of when this version was created.
	VersionPayloadCreatedAt *time.Time `json:"version_payload_created_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestVersionChangeAttributes instantiates a new SyntheticsTestVersionChangeAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestVersionChangeAttributes() *SyntheticsTestVersionChangeAttributes {
	this := SyntheticsTestVersionChangeAttributes{}
	return &this
}

// NewSyntheticsTestVersionChangeAttributesWithDefaults instantiates a new SyntheticsTestVersionChangeAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestVersionChangeAttributesWithDefaults() *SyntheticsTestVersionChangeAttributes {
	this := SyntheticsTestVersionChangeAttributes{}
	return &this
}

// GetAuthorUuid returns the AuthorUuid field value if set, zero value otherwise.
func (o *SyntheticsTestVersionChangeAttributes) GetAuthorUuid() string {
	if o == nil || o.AuthorUuid == nil {
		var ret string
		return ret
	}
	return *o.AuthorUuid
}

// GetAuthorUuidOk returns a tuple with the AuthorUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionChangeAttributes) GetAuthorUuidOk() (*string, bool) {
	if o == nil || o.AuthorUuid == nil {
		return nil, false
	}
	return o.AuthorUuid, true
}

// HasAuthorUuid returns a boolean if a field has been set.
func (o *SyntheticsTestVersionChangeAttributes) HasAuthorUuid() bool {
	return o != nil && o.AuthorUuid != nil
}

// SetAuthorUuid gets a reference to the given string and assigns it to the AuthorUuid field.
func (o *SyntheticsTestVersionChangeAttributes) SetAuthorUuid(v string) {
	o.AuthorUuid = &v
}

// GetChangeMetadata returns the ChangeMetadata field value if set, zero value otherwise.
func (o *SyntheticsTestVersionChangeAttributes) GetChangeMetadata() []SyntheticsTestVersionChangeMetadataItem {
	if o == nil || o.ChangeMetadata == nil {
		var ret []SyntheticsTestVersionChangeMetadataItem
		return ret
	}
	return o.ChangeMetadata
}

// GetChangeMetadataOk returns a tuple with the ChangeMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionChangeAttributes) GetChangeMetadataOk() (*[]SyntheticsTestVersionChangeMetadataItem, bool) {
	if o == nil || o.ChangeMetadata == nil {
		return nil, false
	}
	return &o.ChangeMetadata, true
}

// HasChangeMetadata returns a boolean if a field has been set.
func (o *SyntheticsTestVersionChangeAttributes) HasChangeMetadata() bool {
	return o != nil && o.ChangeMetadata != nil
}

// SetChangeMetadata gets a reference to the given []SyntheticsTestVersionChangeMetadataItem and assigns it to the ChangeMetadata field.
func (o *SyntheticsTestVersionChangeAttributes) SetChangeMetadata(v []SyntheticsTestVersionChangeMetadataItem) {
	o.ChangeMetadata = v
}

// GetVersionNumber returns the VersionNumber field value if set, zero value otherwise.
func (o *SyntheticsTestVersionChangeAttributes) GetVersionNumber() int64 {
	if o == nil || o.VersionNumber == nil {
		var ret int64
		return ret
	}
	return *o.VersionNumber
}

// GetVersionNumberOk returns a tuple with the VersionNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionChangeAttributes) GetVersionNumberOk() (*int64, bool) {
	if o == nil || o.VersionNumber == nil {
		return nil, false
	}
	return o.VersionNumber, true
}

// HasVersionNumber returns a boolean if a field has been set.
func (o *SyntheticsTestVersionChangeAttributes) HasVersionNumber() bool {
	return o != nil && o.VersionNumber != nil
}

// SetVersionNumber gets a reference to the given int64 and assigns it to the VersionNumber field.
func (o *SyntheticsTestVersionChangeAttributes) SetVersionNumber(v int64) {
	o.VersionNumber = &v
}

// GetVersionPayloadCreatedAt returns the VersionPayloadCreatedAt field value if set, zero value otherwise.
func (o *SyntheticsTestVersionChangeAttributes) GetVersionPayloadCreatedAt() time.Time {
	if o == nil || o.VersionPayloadCreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.VersionPayloadCreatedAt
}

// GetVersionPayloadCreatedAtOk returns a tuple with the VersionPayloadCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionChangeAttributes) GetVersionPayloadCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.VersionPayloadCreatedAt == nil {
		return nil, false
	}
	return o.VersionPayloadCreatedAt, true
}

// HasVersionPayloadCreatedAt returns a boolean if a field has been set.
func (o *SyntheticsTestVersionChangeAttributes) HasVersionPayloadCreatedAt() bool {
	return o != nil && o.VersionPayloadCreatedAt != nil
}

// SetVersionPayloadCreatedAt gets a reference to the given time.Time and assigns it to the VersionPayloadCreatedAt field.
func (o *SyntheticsTestVersionChangeAttributes) SetVersionPayloadCreatedAt(v time.Time) {
	o.VersionPayloadCreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestVersionChangeAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AuthorUuid != nil {
		toSerialize["author_uuid"] = o.AuthorUuid
	}
	if o.ChangeMetadata != nil {
		toSerialize["change_metadata"] = o.ChangeMetadata
	}
	if o.VersionNumber != nil {
		toSerialize["version_number"] = o.VersionNumber
	}
	if o.VersionPayloadCreatedAt != nil {
		if o.VersionPayloadCreatedAt.Nanosecond() == 0 {
			toSerialize["version_payload_created_at"] = o.VersionPayloadCreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["version_payload_created_at"] = o.VersionPayloadCreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestVersionChangeAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthorUuid              *string                                   `json:"author_uuid,omitempty"`
		ChangeMetadata          []SyntheticsTestVersionChangeMetadataItem `json:"change_metadata,omitempty"`
		VersionNumber           *int64                                    `json:"version_number,omitempty"`
		VersionPayloadCreatedAt *time.Time                                `json:"version_payload_created_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author_uuid", "change_metadata", "version_number", "version_payload_created_at"})
	} else {
		return err
	}
	o.AuthorUuid = all.AuthorUuid
	o.ChangeMetadata = all.ChangeMetadata
	o.VersionNumber = all.VersionNumber
	o.VersionPayloadCreatedAt = all.VersionPayloadCreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
