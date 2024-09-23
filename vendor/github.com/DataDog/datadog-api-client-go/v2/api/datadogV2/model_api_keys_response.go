// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// APIKeysResponse Response for a list of API keys.
type APIKeysResponse struct {
	// Array of API keys.
	Data []PartialAPIKey `json:"data,omitempty"`
	// Array of objects related to the API key.
	Included []APIKeyResponseIncludedItem `json:"included,omitempty"`
	// Additional information related to api keys response.
	Meta *APIKeysResponseMeta `json:"meta,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAPIKeysResponse instantiates a new APIKeysResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAPIKeysResponse() *APIKeysResponse {
	this := APIKeysResponse{}
	return &this
}

// NewAPIKeysResponseWithDefaults instantiates a new APIKeysResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAPIKeysResponseWithDefaults() *APIKeysResponse {
	this := APIKeysResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *APIKeysResponse) GetData() []PartialAPIKey {
	if o == nil || o.Data == nil {
		var ret []PartialAPIKey
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIKeysResponse) GetDataOk() (*[]PartialAPIKey, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return &o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *APIKeysResponse) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given []PartialAPIKey and assigns it to the Data field.
func (o *APIKeysResponse) SetData(v []PartialAPIKey) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *APIKeysResponse) GetIncluded() []APIKeyResponseIncludedItem {
	if o == nil || o.Included == nil {
		var ret []APIKeyResponseIncludedItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIKeysResponse) GetIncludedOk() (*[]APIKeyResponseIncludedItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return &o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *APIKeysResponse) HasIncluded() bool {
	return o != nil && o.Included != nil
}

// SetIncluded gets a reference to the given []APIKeyResponseIncludedItem and assigns it to the Included field.
func (o *APIKeysResponse) SetIncluded(v []APIKeyResponseIncludedItem) {
	o.Included = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *APIKeysResponse) GetMeta() APIKeysResponseMeta {
	if o == nil || o.Meta == nil {
		var ret APIKeysResponseMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIKeysResponse) GetMetaOk() (*APIKeysResponseMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *APIKeysResponse) HasMeta() bool {
	return o != nil && o.Meta != nil
}

// SetMeta gets a reference to the given APIKeysResponseMeta and assigns it to the Meta field.
func (o *APIKeysResponse) SetMeta(v APIKeysResponseMeta) {
	o.Meta = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o APIKeysResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
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
func (o *APIKeysResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     []PartialAPIKey              `json:"data,omitempty"`
		Included []APIKeyResponseIncludedItem `json:"included,omitempty"`
		Meta     *APIKeysResponseMeta         `json:"meta,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included", "meta"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Data = all.Data
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
