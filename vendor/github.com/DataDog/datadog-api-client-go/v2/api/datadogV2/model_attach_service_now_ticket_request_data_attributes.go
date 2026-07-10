// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AttachServiceNowTicketRequestDataAttributes Attributes of the ServiceNow ticket to attach security findings to.
type AttachServiceNowTicketRequestDataAttributes struct {
	// URL of the ServiceNow incident to attach security findings to. Must be a service-now.com URL pointing to an incident record.
	ServicenowTicketUrl string `json:"servicenow_ticket_url"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAttachServiceNowTicketRequestDataAttributes instantiates a new AttachServiceNowTicketRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAttachServiceNowTicketRequestDataAttributes(servicenowTicketUrl string) *AttachServiceNowTicketRequestDataAttributes {
	this := AttachServiceNowTicketRequestDataAttributes{}
	this.ServicenowTicketUrl = servicenowTicketUrl
	return &this
}

// NewAttachServiceNowTicketRequestDataAttributesWithDefaults instantiates a new AttachServiceNowTicketRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAttachServiceNowTicketRequestDataAttributesWithDefaults() *AttachServiceNowTicketRequestDataAttributes {
	this := AttachServiceNowTicketRequestDataAttributes{}
	return &this
}

// GetServicenowTicketUrl returns the ServicenowTicketUrl field value.
func (o *AttachServiceNowTicketRequestDataAttributes) GetServicenowTicketUrl() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ServicenowTicketUrl
}

// GetServicenowTicketUrlOk returns a tuple with the ServicenowTicketUrl field value
// and a boolean to check if the value has been set.
func (o *AttachServiceNowTicketRequestDataAttributes) GetServicenowTicketUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServicenowTicketUrl, true
}

// SetServicenowTicketUrl sets field value.
func (o *AttachServiceNowTicketRequestDataAttributes) SetServicenowTicketUrl(v string) {
	o.ServicenowTicketUrl = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AttachServiceNowTicketRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["servicenow_ticket_url"] = o.ServicenowTicketUrl

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AttachServiceNowTicketRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ServicenowTicketUrl *string `json:"servicenow_ticket_url"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ServicenowTicketUrl == nil {
		return fmt.Errorf("required field servicenow_ticket_url missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"servicenow_ticket_url"})
	} else {
		return err
	}
	o.ServicenowTicketUrl = *all.ServicenowTicketUrl

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
