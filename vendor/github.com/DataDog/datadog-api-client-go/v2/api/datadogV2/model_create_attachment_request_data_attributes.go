// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateAttachmentRequestDataAttributes The attributes for creating an attachment.
type CreateAttachmentRequestDataAttributes struct {
	// The attachment object for creating an attachment.
	Attachment *CreateAttachmentRequestDataAttributesAttachment `json:"attachment,omitempty"`
	// The type of the attachment.
	AttachmentType *AttachmentDataAttributesAttachmentType `json:"attachment_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateAttachmentRequestDataAttributes instantiates a new CreateAttachmentRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateAttachmentRequestDataAttributes() *CreateAttachmentRequestDataAttributes {
	this := CreateAttachmentRequestDataAttributes{}
	return &this
}

// NewCreateAttachmentRequestDataAttributesWithDefaults instantiates a new CreateAttachmentRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateAttachmentRequestDataAttributesWithDefaults() *CreateAttachmentRequestDataAttributes {
	this := CreateAttachmentRequestDataAttributes{}
	return &this
}

// GetAttachment returns the Attachment field value if set, zero value otherwise.
func (o *CreateAttachmentRequestDataAttributes) GetAttachment() CreateAttachmentRequestDataAttributesAttachment {
	if o == nil || o.Attachment == nil {
		var ret CreateAttachmentRequestDataAttributesAttachment
		return ret
	}
	return *o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAttachmentRequestDataAttributes) GetAttachmentOk() (*CreateAttachmentRequestDataAttributesAttachment, bool) {
	if o == nil || o.Attachment == nil {
		return nil, false
	}
	return o.Attachment, true
}

// HasAttachment returns a boolean if a field has been set.
func (o *CreateAttachmentRequestDataAttributes) HasAttachment() bool {
	return o != nil && o.Attachment != nil
}

// SetAttachment gets a reference to the given CreateAttachmentRequestDataAttributesAttachment and assigns it to the Attachment field.
func (o *CreateAttachmentRequestDataAttributes) SetAttachment(v CreateAttachmentRequestDataAttributesAttachment) {
	o.Attachment = &v
}

// GetAttachmentType returns the AttachmentType field value if set, zero value otherwise.
func (o *CreateAttachmentRequestDataAttributes) GetAttachmentType() AttachmentDataAttributesAttachmentType {
	if o == nil || o.AttachmentType == nil {
		var ret AttachmentDataAttributesAttachmentType
		return ret
	}
	return *o.AttachmentType
}

// GetAttachmentTypeOk returns a tuple with the AttachmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAttachmentRequestDataAttributes) GetAttachmentTypeOk() (*AttachmentDataAttributesAttachmentType, bool) {
	if o == nil || o.AttachmentType == nil {
		return nil, false
	}
	return o.AttachmentType, true
}

// HasAttachmentType returns a boolean if a field has been set.
func (o *CreateAttachmentRequestDataAttributes) HasAttachmentType() bool {
	return o != nil && o.AttachmentType != nil
}

// SetAttachmentType gets a reference to the given AttachmentDataAttributesAttachmentType and assigns it to the AttachmentType field.
func (o *CreateAttachmentRequestDataAttributes) SetAttachmentType(v AttachmentDataAttributesAttachmentType) {
	o.AttachmentType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateAttachmentRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Attachment != nil {
		toSerialize["attachment"] = o.Attachment
	}
	if o.AttachmentType != nil {
		toSerialize["attachment_type"] = o.AttachmentType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateAttachmentRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attachment     *CreateAttachmentRequestDataAttributesAttachment `json:"attachment,omitempty"`
		AttachmentType *AttachmentDataAttributesAttachmentType          `json:"attachment_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attachment", "attachment_type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attachment != nil && all.Attachment.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attachment = all.Attachment
	if all.AttachmentType != nil && !all.AttachmentType.IsValid() {
		hasInvalidField = true
	} else {
		o.AttachmentType = all.AttachmentType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
