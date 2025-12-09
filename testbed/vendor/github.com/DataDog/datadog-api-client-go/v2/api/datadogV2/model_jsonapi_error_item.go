// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// JSONAPIErrorItem API error response body
type JSONAPIErrorItem struct {
	// A human-readable explanation specific to this occurrence of the error.
	Detail *string `json:"detail,omitempty"`
	// Non-standard meta-information about the error
	Meta map[string]interface{} `json:"meta,omitempty"`
	// References to the source of the error.
	Source *JSONAPIErrorItemSource `json:"source,omitempty"`
	// Status code of the response.
	Status *string `json:"status,omitempty"`
	// Short human-readable summary of the error.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewJSONAPIErrorItem instantiates a new JSONAPIErrorItem object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewJSONAPIErrorItem() *JSONAPIErrorItem {
	this := JSONAPIErrorItem{}
	return &this
}

// NewJSONAPIErrorItemWithDefaults instantiates a new JSONAPIErrorItem object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewJSONAPIErrorItemWithDefaults() *JSONAPIErrorItem {
	this := JSONAPIErrorItem{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *JSONAPIErrorItem) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONAPIErrorItem) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *JSONAPIErrorItem) HasDetail() bool {
	return o != nil && o.Detail != nil
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *JSONAPIErrorItem) SetDetail(v string) {
	o.Detail = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *JSONAPIErrorItem) GetMeta() map[string]interface{} {
	if o == nil || o.Meta == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONAPIErrorItem) GetMetaOk() (*map[string]interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return &o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *JSONAPIErrorItem) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given map[string]interface{} and assigns it to the Meta field.
func (o *JSONAPIErrorItem) SetMeta(v map[string]interface{}) {
	o.Meta = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *JSONAPIErrorItem) GetSource() JSONAPIErrorItemSource {
	if o == nil || o.Source == nil {
		var ret JSONAPIErrorItemSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONAPIErrorItem) GetSourceOk() (*JSONAPIErrorItemSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *JSONAPIErrorItem) HasSource() bool {
	return o != nil && o.Source != nil
}

// SetSource gets a reference to the given JSONAPIErrorItemSource and assigns it to the Source field.
func (o *JSONAPIErrorItem) SetSource(v JSONAPIErrorItemSource) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *JSONAPIErrorItem) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONAPIErrorItem) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *JSONAPIErrorItem) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *JSONAPIErrorItem) SetStatus(v string) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *JSONAPIErrorItem) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JSONAPIErrorItem) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *JSONAPIErrorItem) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *JSONAPIErrorItem) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o JSONAPIErrorItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *JSONAPIErrorItem) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Detail *string                 `json:"detail,omitempty"`
		Meta   map[string]interface{}  `json:"meta,omitempty"`
		Source *JSONAPIErrorItemSource `json:"source,omitempty"`
		Status *string                 `json:"status,omitempty"`
		Title  *string                 `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"detail", "meta", "source", "status", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Detail = all.Detail
	o.Meta = all.Meta
	if all.Source != nil && all.Source.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Source = all.Source
	o.Status = all.Status
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
