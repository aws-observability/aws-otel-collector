// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ListInvestigationsResponseDataAttributes Attributes of an investigation list item.
type ListInvestigationsResponseDataAttributes struct {
	// The current status of the investigation.
	Status string `json:"status"`
	// The title of the investigation.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewListInvestigationsResponseDataAttributes instantiates a new ListInvestigationsResponseDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewListInvestigationsResponseDataAttributes(status string, title string) *ListInvestigationsResponseDataAttributes {
	this := ListInvestigationsResponseDataAttributes{}
	this.Status = status
	this.Title = title
	return &this
}

// NewListInvestigationsResponseDataAttributesWithDefaults instantiates a new ListInvestigationsResponseDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewListInvestigationsResponseDataAttributesWithDefaults() *ListInvestigationsResponseDataAttributes {
	this := ListInvestigationsResponseDataAttributes{}
	return &this
}

// GetStatus returns the Status field value.
func (o *ListInvestigationsResponseDataAttributes) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ListInvestigationsResponseDataAttributes) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ListInvestigationsResponseDataAttributes) SetStatus(v string) {
	o.Status = v
}

// GetTitle returns the Title field value.
func (o *ListInvestigationsResponseDataAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ListInvestigationsResponseDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *ListInvestigationsResponseDataAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ListInvestigationsResponseDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ListInvestigationsResponseDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status *string `json:"status"`
		Title  *string `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"status", "title"})
	} else {
		return err
	}
	o.Status = *all.Status
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
