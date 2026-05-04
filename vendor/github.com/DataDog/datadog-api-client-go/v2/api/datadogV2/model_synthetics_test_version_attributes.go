// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestVersionAttributes Attributes of a specific Synthetic test version.
type SyntheticsTestVersionAttributes struct {
	// Object describing the author of a test version.
	Author *SyntheticsTestVersionAuthor `json:"author,omitempty"`
	// List of metadata describing individual changes in this version.
	// Only returned when the `include_change_metadata` query parameter is `true`.
	ChangeMetadata []SyntheticsTestVersionChangeMetadataItem `json:"change_metadata,omitempty"`
	// The full test configuration at this version.
	Payload map[string]interface{} `json:"payload,omitempty"`
	// Timestamp of when this version was created.
	VersionPayloadCreatedAt *time.Time `json:"version_payload_created_at,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestVersionAttributes instantiates a new SyntheticsTestVersionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestVersionAttributes() *SyntheticsTestVersionAttributes {
	this := SyntheticsTestVersionAttributes{}
	return &this
}

// NewSyntheticsTestVersionAttributesWithDefaults instantiates a new SyntheticsTestVersionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestVersionAttributesWithDefaults() *SyntheticsTestVersionAttributes {
	this := SyntheticsTestVersionAttributes{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *SyntheticsTestVersionAttributes) GetAuthor() SyntheticsTestVersionAuthor {
	if o == nil || o.Author == nil {
		var ret SyntheticsTestVersionAuthor
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionAttributes) GetAuthorOk() (*SyntheticsTestVersionAuthor, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *SyntheticsTestVersionAttributes) HasAuthor() bool {
	return o != nil && o.Author != nil
}

// SetAuthor gets a reference to the given SyntheticsTestVersionAuthor and assigns it to the Author field.
func (o *SyntheticsTestVersionAttributes) SetAuthor(v SyntheticsTestVersionAuthor) {
	o.Author = &v
}

// GetChangeMetadata returns the ChangeMetadata field value if set, zero value otherwise.
func (o *SyntheticsTestVersionAttributes) GetChangeMetadata() []SyntheticsTestVersionChangeMetadataItem {
	if o == nil || o.ChangeMetadata == nil {
		var ret []SyntheticsTestVersionChangeMetadataItem
		return ret
	}
	return o.ChangeMetadata
}

// GetChangeMetadataOk returns a tuple with the ChangeMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionAttributes) GetChangeMetadataOk() (*[]SyntheticsTestVersionChangeMetadataItem, bool) {
	if o == nil || o.ChangeMetadata == nil {
		return nil, false
	}
	return &o.ChangeMetadata, true
}

// HasChangeMetadata returns a boolean if a field has been set.
func (o *SyntheticsTestVersionAttributes) HasChangeMetadata() bool {
	return o != nil && o.ChangeMetadata != nil
}

// SetChangeMetadata gets a reference to the given []SyntheticsTestVersionChangeMetadataItem and assigns it to the ChangeMetadata field.
func (o *SyntheticsTestVersionAttributes) SetChangeMetadata(v []SyntheticsTestVersionChangeMetadataItem) {
	o.ChangeMetadata = v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *SyntheticsTestVersionAttributes) GetPayload() map[string]interface{} {
	if o == nil || o.Payload == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionAttributes) GetPayloadOk() (*map[string]interface{}, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return &o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *SyntheticsTestVersionAttributes) HasPayload() bool {
	return o != nil && o.Payload != nil
}

// SetPayload gets a reference to the given map[string]interface{} and assigns it to the Payload field.
func (o *SyntheticsTestVersionAttributes) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetVersionPayloadCreatedAt returns the VersionPayloadCreatedAt field value if set, zero value otherwise.
func (o *SyntheticsTestVersionAttributes) GetVersionPayloadCreatedAt() time.Time {
	if o == nil || o.VersionPayloadCreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.VersionPayloadCreatedAt
}

// GetVersionPayloadCreatedAtOk returns a tuple with the VersionPayloadCreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestVersionAttributes) GetVersionPayloadCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.VersionPayloadCreatedAt == nil {
		return nil, false
	}
	return o.VersionPayloadCreatedAt, true
}

// HasVersionPayloadCreatedAt returns a boolean if a field has been set.
func (o *SyntheticsTestVersionAttributes) HasVersionPayloadCreatedAt() bool {
	return o != nil && o.VersionPayloadCreatedAt != nil
}

// SetVersionPayloadCreatedAt gets a reference to the given time.Time and assigns it to the VersionPayloadCreatedAt field.
func (o *SyntheticsTestVersionAttributes) SetVersionPayloadCreatedAt(v time.Time) {
	o.VersionPayloadCreatedAt = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestVersionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.ChangeMetadata != nil {
		toSerialize["change_metadata"] = o.ChangeMetadata
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
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
func (o *SyntheticsTestVersionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Author                  *SyntheticsTestVersionAuthor              `json:"author,omitempty"`
		ChangeMetadata          []SyntheticsTestVersionChangeMetadataItem `json:"change_metadata,omitempty"`
		Payload                 map[string]interface{}                    `json:"payload,omitempty"`
		VersionPayloadCreatedAt *time.Time                                `json:"version_payload_created_at,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author", "change_metadata", "payload", "version_payload_created_at"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Author != nil && all.Author.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Author = all.Author
	o.ChangeMetadata = all.ChangeMetadata
	o.Payload = all.Payload
	o.VersionPayloadCreatedAt = all.VersionPayloadCreatedAt

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
