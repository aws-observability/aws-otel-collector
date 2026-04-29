// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsResponseMeta Metadata for a Product Analytics query response.
type ProductAnalyticsResponseMeta struct {
	// Unique identifier for the request, used for multi-step query continuation.
	RequestId *string `json:"request_id,omitempty"`
	// The execution status of a Product Analytics query.
	Status *ProductAnalyticsResponseMetaStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsResponseMeta instantiates a new ProductAnalyticsResponseMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsResponseMeta() *ProductAnalyticsResponseMeta {
	this := ProductAnalyticsResponseMeta{}
	return &this
}

// NewProductAnalyticsResponseMetaWithDefaults instantiates a new ProductAnalyticsResponseMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsResponseMetaWithDefaults() *ProductAnalyticsResponseMeta {
	this := ProductAnalyticsResponseMeta{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ProductAnalyticsResponseMeta) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsResponseMeta) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ProductAnalyticsResponseMeta) HasRequestId() bool {
	return o != nil && o.RequestId != nil
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ProductAnalyticsResponseMeta) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ProductAnalyticsResponseMeta) GetStatus() ProductAnalyticsResponseMetaStatus {
	if o == nil || o.Status == nil {
		var ret ProductAnalyticsResponseMetaStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsResponseMeta) GetStatusOk() (*ProductAnalyticsResponseMetaStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ProductAnalyticsResponseMeta) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given ProductAnalyticsResponseMetaStatus and assigns it to the Status field.
func (o *ProductAnalyticsResponseMeta) SetStatus(v ProductAnalyticsResponseMetaStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsResponseMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RequestId *string                             `json:"request_id,omitempty"`
		Status    *ProductAnalyticsResponseMetaStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"request_id", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.RequestId = all.RequestId
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
