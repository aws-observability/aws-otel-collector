// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IncidentTodoListResponse Response with a list of incident todos.
type IncidentTodoListResponse struct {
	// An array of incident todos.
	Data []IncidentTodoResponseData `json:"data"`
	// Included related resources that the user requested.
	Included []IncidentTodoResponseIncludedItem `json:"included,omitempty"`
	// The metadata object containing pagination metadata.
	Meta *IncidentResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIncidentTodoListResponse instantiates a new IncidentTodoListResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIncidentTodoListResponse(data []IncidentTodoResponseData) *IncidentTodoListResponse {
	this := IncidentTodoListResponse{}
	this.Data = data
	return &this
}

// NewIncidentTodoListResponseWithDefaults instantiates a new IncidentTodoListResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIncidentTodoListResponseWithDefaults() *IncidentTodoListResponse {
	this := IncidentTodoListResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *IncidentTodoListResponse) GetData() []IncidentTodoResponseData {
	if o == nil {
		var ret []IncidentTodoResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IncidentTodoListResponse) GetDataOk() (*[]IncidentTodoResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *IncidentTodoListResponse) SetData(v []IncidentTodoResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncidentTodoListResponse) GetIncluded() []IncidentTodoResponseIncludedItem {
	if o == nil || o.Included == nil {
		var ret []IncidentTodoResponseIncludedItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTodoListResponse) GetIncludedOk() (*[]IncidentTodoResponseIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncidentTodoListResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []IncidentTodoResponseIncludedItem and assigns it to the Included field.
func (o *IncidentTodoListResponse) SetIncluded(v []IncidentTodoResponseIncludedItem) {
	o.Included = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *IncidentTodoListResponse) GetMeta() IncidentResponseMeta {
	if o == nil || o.Meta == nil {
		var ret IncidentResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncidentTodoListResponse) GetMetaOk() (*IncidentResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *IncidentTodoListResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given IncidentResponseMeta and assigns it to the Meta field.
func (o *IncidentTodoListResponse) SetMeta(v IncidentResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o IncidentTodoListResponse) MarshalJSON() ([]byte, error) {
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
func (o *IncidentTodoListResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *[]IncidentTodoResponseData        `json:"data"`
		Included []IncidentTodoResponseIncludedItem `json:"included,omitempty"`
		Meta     *IncidentResponseMeta              `json:"meta,omitempty"`
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
