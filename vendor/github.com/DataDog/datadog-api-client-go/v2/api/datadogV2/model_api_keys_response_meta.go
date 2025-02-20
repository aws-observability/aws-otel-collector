// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// APIKeysResponseMeta Additional information related to api keys response.
type APIKeysResponseMeta struct {
	// Max allowed number of API keys.
	MaxAllowed *int64 `json:"max_allowed,omitempty"`
	// Additional information related to the API keys response.
	Page *APIKeysResponseMetaPage `json:"page,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAPIKeysResponseMeta instantiates a new APIKeysResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAPIKeysResponseMeta() *APIKeysResponseMeta {
	this := APIKeysResponseMeta{}
	return &this
}

// NewAPIKeysResponseMetaWithDefaults instantiates a new APIKeysResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAPIKeysResponseMetaWithDefaults() *APIKeysResponseMeta {
	this := APIKeysResponseMeta{}
	return &this
}

// GetMaxAllowed returns the MaxAllowed field value if set, zero value otherwise.
func (o *APIKeysResponseMeta) GetMaxAllowed() int64 {
	if o == nil || o.MaxAllowed == nil {
		var ret int64
		return ret
	}
	return *o.MaxAllowed
}

// GetMaxAllowedOk returns a tuple with the MaxAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIKeysResponseMeta) GetMaxAllowedOk() (*int64, bool) {
	if o == nil || o.MaxAllowed == nil {
		return nil, false
	}
	return o.MaxAllowed, true
}

// HasMaxAllowed returns a boolean if a field has been set.
func (o *APIKeysResponseMeta) HasMaxAllowed() bool {
	return o != nil && o.MaxAllowed != nil
}

// SetMaxAllowed gets a reference to the given int64 and assigns it to the MaxAllowed field.
func (o *APIKeysResponseMeta) SetMaxAllowed(v int64) {
	o.MaxAllowed = &v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *APIKeysResponseMeta) GetPage() APIKeysResponseMetaPage {
	if o == nil || o.Page == nil {
		var ret APIKeysResponseMetaPage
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIKeysResponseMeta) GetPageOk() (*APIKeysResponseMetaPage, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *APIKeysResponseMeta) HasPage() bool {
	return o != nil && o.Page != nil
}

// SetPage gets a reference to the given APIKeysResponseMetaPage and assigns it to the Page field.
func (o *APIKeysResponseMeta) SetPage(v APIKeysResponseMetaPage) {
	o.Page = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o APIKeysResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.MaxAllowed != nil {
		toSerialize["max_allowed"] = o.MaxAllowed
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *APIKeysResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		MaxAllowed *int64                   `json:"max_allowed,omitempty"`
		Page       *APIKeysResponseMetaPage `json:"page,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"max_allowed", "page"})
	} else {
		return err
	}

	hasInvalidField := false
	o.MaxAllowed = all.MaxAllowed
	if all.Page != nil && all.Page.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Page = all.Page

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
