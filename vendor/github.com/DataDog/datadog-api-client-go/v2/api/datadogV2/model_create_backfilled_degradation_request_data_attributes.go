// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateBackfilledDegradationRequestDataAttributes The supported attributes for creating a backfilled degradation.
type CreateBackfilledDegradationRequestDataAttributes struct {
	// The title of the backfilled degradation.
	Title string `json:"title"`
	// The list of status updates describing the timeline of the degradation.
	Updates []CreateBackfilledDegradationRequestDataAttributesUpdatesItems `json:"updates"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateBackfilledDegradationRequestDataAttributes instantiates a new CreateBackfilledDegradationRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateBackfilledDegradationRequestDataAttributes(title string, updates []CreateBackfilledDegradationRequestDataAttributesUpdatesItems) *CreateBackfilledDegradationRequestDataAttributes {
	this := CreateBackfilledDegradationRequestDataAttributes{}
	this.Title = title
	this.Updates = updates
	return &this
}

// NewCreateBackfilledDegradationRequestDataAttributesWithDefaults instantiates a new CreateBackfilledDegradationRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateBackfilledDegradationRequestDataAttributesWithDefaults() *CreateBackfilledDegradationRequestDataAttributes {
	this := CreateBackfilledDegradationRequestDataAttributes{}
	return &this
}

// GetTitle returns the Title field value.
func (o *CreateBackfilledDegradationRequestDataAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateBackfilledDegradationRequestDataAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *CreateBackfilledDegradationRequestDataAttributes) SetTitle(v string) {
	o.Title = v
}

// GetUpdates returns the Updates field value.
func (o *CreateBackfilledDegradationRequestDataAttributes) GetUpdates() []CreateBackfilledDegradationRequestDataAttributesUpdatesItems {
	if o == nil {
		var ret []CreateBackfilledDegradationRequestDataAttributesUpdatesItems
		return ret
	}
	return o.Updates
}

// GetUpdatesOk returns a tuple with the Updates field value
// and a boolean to check if the value has been set.
func (o *CreateBackfilledDegradationRequestDataAttributes) GetUpdatesOk() (*[]CreateBackfilledDegradationRequestDataAttributesUpdatesItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updates, true
}

// SetUpdates sets field value.
func (o *CreateBackfilledDegradationRequestDataAttributes) SetUpdates(v []CreateBackfilledDegradationRequestDataAttributesUpdatesItems) {
	o.Updates = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateBackfilledDegradationRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["title"] = o.Title
	toSerialize["updates"] = o.Updates

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateBackfilledDegradationRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Title   *string                                                         `json:"title"`
		Updates *[]CreateBackfilledDegradationRequestDataAttributesUpdatesItems `json:"updates"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	if all.Updates == nil {
		return fmt.Errorf("required field updates missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"title", "updates"})
	} else {
		return err
	}
	o.Title = *all.Title
	o.Updates = *all.Updates

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
