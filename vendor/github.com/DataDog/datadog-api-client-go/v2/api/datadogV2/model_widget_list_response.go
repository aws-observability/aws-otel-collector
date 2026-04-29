// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// WidgetListResponse Response containing a list of widgets.
type WidgetListResponse struct {
	// List of widget resources.
	Data []WidgetData `json:"data"`
	// Array of user resources related to the widgets.
	Included []WidgetIncludedUser `json:"included,omitempty"`
	// Metadata about the search results.
	Meta *WidgetSearchMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewWidgetListResponse instantiates a new WidgetListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewWidgetListResponse(data []WidgetData) *WidgetListResponse {
	this := WidgetListResponse{}
	this.Data = data
	return &this
}

// NewWidgetListResponseWithDefaults instantiates a new WidgetListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewWidgetListResponseWithDefaults() *WidgetListResponse {
	this := WidgetListResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *WidgetListResponse) GetData() []WidgetData {
	if o == nil {
		var ret []WidgetData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *WidgetListResponse) GetDataOk() (*[]WidgetData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *WidgetListResponse) SetData(v []WidgetData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *WidgetListResponse) GetIncluded() []WidgetIncludedUser {
	if o == nil || o.Included == nil {
		var ret []WidgetIncludedUser
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetListResponse) GetIncludedOk() (*[]WidgetIncludedUser, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *WidgetListResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []WidgetIncludedUser and assigns it to the Included field.
func (o *WidgetListResponse) SetIncluded(v []WidgetIncludedUser) {
	o.Included = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *WidgetListResponse) GetMeta() WidgetSearchMeta {
	if o == nil || o.Meta == nil {
		var ret WidgetSearchMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WidgetListResponse) GetMetaOk() (*WidgetSearchMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *WidgetListResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given WidgetSearchMeta and assigns it to the Meta field.
func (o *WidgetListResponse) SetMeta(v WidgetSearchMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o WidgetListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *WidgetListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]WidgetData        `json:"data"`
		Included []WidgetIncludedUser `json:"included,omitempty"`
		Meta     *WidgetSearchMeta    `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = *all.Data
	o.Included = all.Included
	if all.Meta != nil && all.Meta.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Meta = all.Meta

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
