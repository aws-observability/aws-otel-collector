// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchAttachmentRequestDataAttributes The attributes for updating an attachment.
type PatchAttachmentRequestDataAttributes struct {
	// The updated attachment object.
	Attachment *PatchAttachmentRequestDataAttributesAttachment `json:"attachment,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchAttachmentRequestDataAttributes instantiates a new PatchAttachmentRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchAttachmentRequestDataAttributes() *PatchAttachmentRequestDataAttributes {
	this := PatchAttachmentRequestDataAttributes{}
	return &this
}

// NewPatchAttachmentRequestDataAttributesWithDefaults instantiates a new PatchAttachmentRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchAttachmentRequestDataAttributesWithDefaults() *PatchAttachmentRequestDataAttributes {
	this := PatchAttachmentRequestDataAttributes{}
	return &this
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *PatchAttachmentRequestDataAttributes) GetAttachment() PatchAttachmentRequestDataAttributesAttachment {
	if o == nil || o.Attachment == nil {
		var ret PatchAttachmentRequestDataAttributesAttachment
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchAttachmentRequestDataAttributes) GetAttachmentOk() (*PatchAttachmentRequestDataAttributesAttachment, bool) {
	if o == nil || o.Attachment == nil {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *PatchAttachmentRequestDataAttributes) HasAttachment() bool {
	return o != nil && o.Attachment != nil
}

// SetAttachment gets a reference to the given PatchAttachmentRequestDataAttributesAttachment and assigns it to the Attachment field.
func (o *PatchAttachmentRequestDataAttributes) SetAttachment(v PatchAttachmentRequestDataAttributesAttachment) {
	o.Attachment = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchAttachmentRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attachment != nil {
		toSerialize["attachment"] = o.Attachment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *PatchAttachmentRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attachment *PatchAttachmentRequestDataAttributesAttachment `json:"attachment,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attachment"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attachment != nil && all.Attachment.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attachment = all.Attachment

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
