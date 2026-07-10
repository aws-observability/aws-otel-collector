// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SampleLogGenerationBulkSubscriptionItemMeta Per-item status returned for a bulk subscription request.
type SampleLogGenerationBulkSubscriptionItemMeta struct {
	// A description of the error encountered for this content pack, if the subscription could not be created.
	Error *string `json:"error,omitempty"`
	// The HTTP status code that resulted from creating the subscription for this content pack.
	Status int32 `json:"status"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSampleLogGenerationBulkSubscriptionItemMeta instantiates a new SampleLogGenerationBulkSubscriptionItemMeta object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSampleLogGenerationBulkSubscriptionItemMeta(status int32) *SampleLogGenerationBulkSubscriptionItemMeta {
	this := SampleLogGenerationBulkSubscriptionItemMeta{}
	this.Status = status
	return &this
}

// NewSampleLogGenerationBulkSubscriptionItemMetaWithDefaults instantiates a new SampleLogGenerationBulkSubscriptionItemMeta object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSampleLogGenerationBulkSubscriptionItemMetaWithDefaults() *SampleLogGenerationBulkSubscriptionItemMeta {
	this := SampleLogGenerationBulkSubscriptionItemMeta{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *SampleLogGenerationBulkSubscriptionItemMeta) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationBulkSubscriptionItemMeta) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *SampleLogGenerationBulkSubscriptionItemMeta) HasError() bool {
	return o != nil && o.Error != nil
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *SampleLogGenerationBulkSubscriptionItemMeta) SetError(v string) {
	o.Error = &v
}

// GetStatus returns the Status field value.
func (o *SampleLogGenerationBulkSubscriptionItemMeta) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SampleLogGenerationBulkSubscriptionItemMeta) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *SampleLogGenerationBulkSubscriptionItemMeta) SetStatus(v int32) {
	o.Status = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SampleLogGenerationBulkSubscriptionItemMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	toSerialize["status"] = o.Status

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SampleLogGenerationBulkSubscriptionItemMeta) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Error  *string `json:"error,omitempty"`
		Status *int32  `json:"status"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"error", "status"})
	} else {
		return err
	}
	o.Error = all.Error
	o.Status = *all.Status

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
