// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentAttachmentLinkAttributes The attributes object for a link attachment.
type IncidentAttachmentLinkAttributes struct {
	// The link attachment.
	Attachment IncidentAttachmentLinkAttributesAttachmentObject `json:"attachment"`
	// The type of link attachment attributes.
	AttachmentType IncidentAttachmentLinkAttachmentType `json:"attachment_type"`
	// Timestamp when the incident attachment link was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentAttachmentLinkAttributes instantiates a new IncidentAttachmentLinkAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentAttachmentLinkAttributes(attachment IncidentAttachmentLinkAttributesAttachmentObject, attachmentType IncidentAttachmentLinkAttachmentType) *IncidentAttachmentLinkAttributes {
	this := IncidentAttachmentLinkAttributes{}
	this.Attachment = attachment
	this.AttachmentType = attachmentType
	return &this
}

// NewIncidentAttachmentLinkAttributesWithDefaults instantiates a new IncidentAttachmentLinkAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentAttachmentLinkAttributesWithDefaults() *IncidentAttachmentLinkAttributes {
	this := IncidentAttachmentLinkAttributes{}
	var attachmentType IncidentAttachmentLinkAttachmentType = INCIDENTATTACHMENTLINKATTACHMENTTYPE_LINK
	this.AttachmentType = attachmentType
	return &this
}

// GetAttachment returns the Attachment field value.
func (o *IncidentAttachmentLinkAttributes) GetAttachment() IncidentAttachmentLinkAttributesAttachmentObject {
	if o == nil {
		var ret IncidentAttachmentLinkAttributesAttachmentObject
		return ret
	}
	return o.Attachment
}

// GetAttachmentOk returns a tuple with the Attachment field value
// and a boolean to check if the value has been set.
func (o *IncidentAttachmentLinkAttributes) GetAttachmentOk() (*IncidentAttachmentLinkAttributesAttachmentObject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attachment, true
}

// SetAttachment sets field value.
func (o *IncidentAttachmentLinkAttributes) SetAttachment(v IncidentAttachmentLinkAttributesAttachmentObject) {
	o.Attachment = v
}

// GetAttachmentType returns the AttachmentType field value.
func (o *IncidentAttachmentLinkAttributes) GetAttachmentType() IncidentAttachmentLinkAttachmentType {
	if o == nil {
		var ret IncidentAttachmentLinkAttachmentType
		return ret
	}
	return o.AttachmentType
}

// GetAttachmentTypeOk returns a tuple with the AttachmentType field value
// and a boolean to check if the value has been set.
func (o *IncidentAttachmentLinkAttributes) GetAttachmentTypeOk() (*IncidentAttachmentLinkAttachmentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttachmentType, true
}

// SetAttachmentType sets field value.
func (o *IncidentAttachmentLinkAttributes) SetAttachmentType(v IncidentAttachmentLinkAttachmentType) {
	o.AttachmentType = v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *IncidentAttachmentLinkAttributes) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentAttachmentLinkAttributes) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *IncidentAttachmentLinkAttributes) HasModified() bool {
	return o != nil && o.Modified != nil
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *IncidentAttachmentLinkAttributes) SetModified(v time.Time) {
	o.Modified = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentAttachmentLinkAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["attachment"] = o.Attachment
	toSerialize["attachment_type"] = o.AttachmentType
	if o.Modified != nil {
		if o.Modified.Nanosecond() == 0 {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IncidentAttachmentLinkAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Attachment     *IncidentAttachmentLinkAttributesAttachmentObject `json:"attachment"`
		AttachmentType *IncidentAttachmentLinkAttachmentType             `json:"attachment_type"`
		Modified       *time.Time                                        `json:"modified,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Attachment == nil {
		return fmt.Errorf("required field attachment missing")
	}
	if all.AttachmentType == nil {
		return fmt.Errorf("required field attachment_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"attachment", "attachment_type", "modified"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Attachment.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Attachment = *all.Attachment
	if !all.AttachmentType.IsValid() {
		hasInvalidField = true
	} else {
		o.AttachmentType = *all.AttachmentType
	}
	o.Modified = all.Modified

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
