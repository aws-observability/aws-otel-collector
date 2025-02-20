// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BillingDimensionsMappingBodyItemAttributes Mapping of billing dimensions to endpoint keys.
type BillingDimensionsMappingBodyItemAttributes struct {
	// List of supported endpoints with their keys mapped to the billing_dimension.
	Endpoints []BillingDimensionsMappingBodyItemAttributesEndpointsItems `json:"endpoints,omitempty"`
	// Label used for the billing dimension in the Plan & Usage charts.
	InAppLabel *string `json:"in_app_label,omitempty"`
	// Month in ISO-8601 format, UTC, and precise to the second: `[YYYY-MM-DDThh:mm:ss]`.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBillingDimensionsMappingBodyItemAttributes instantiates a new BillingDimensionsMappingBodyItemAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBillingDimensionsMappingBodyItemAttributes() *BillingDimensionsMappingBodyItemAttributes {
	this := BillingDimensionsMappingBodyItemAttributes{}
	return &this
}

// NewBillingDimensionsMappingBodyItemAttributesWithDefaults instantiates a new BillingDimensionsMappingBodyItemAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBillingDimensionsMappingBodyItemAttributesWithDefaults() *BillingDimensionsMappingBodyItemAttributes {
	this := BillingDimensionsMappingBodyItemAttributes{}
	return &this
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *BillingDimensionsMappingBodyItemAttributes) GetEndpoints() []BillingDimensionsMappingBodyItemAttributesEndpointsItems {
	if o == nil || o.Endpoints == nil {
		var ret []BillingDimensionsMappingBodyItemAttributesEndpointsItems
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingDimensionsMappingBodyItemAttributes) GetEndpointsOk() (*[]BillingDimensionsMappingBodyItemAttributesEndpointsItems, bool) {
	if o == nil || o.Endpoints == nil {
		return nil, false
	}
	return &o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *BillingDimensionsMappingBodyItemAttributes) HasEndpoints() bool {
	return o != nil && o.Endpoints != nil
}

// SetEndpoints gets a reference to the given []BillingDimensionsMappingBodyItemAttributesEndpointsItems and assigns it to the Endpoints field.
func (o *BillingDimensionsMappingBodyItemAttributes) SetEndpoints(v []BillingDimensionsMappingBodyItemAttributesEndpointsItems) {
	o.Endpoints = v
}

// GetInAppLabel returns the InAppLabel field value if set, zero value otherwise.
func (o *BillingDimensionsMappingBodyItemAttributes) GetInAppLabel() string {
	if o == nil || o.InAppLabel == nil {
		var ret string
		return ret
	}
	return *o.InAppLabel
}

// GetInAppLabelOk returns a tuple with the InAppLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingDimensionsMappingBodyItemAttributes) GetInAppLabelOk() (*string, bool) {
	if o == nil || o.InAppLabel == nil {
		return nil, false
	}
	return o.InAppLabel, true
}

// HasInAppLabel returns a boolean if a field has been set.
func (o *BillingDimensionsMappingBodyItemAttributes) HasInAppLabel() bool {
	return o != nil && o.InAppLabel != nil
}

// SetInAppLabel gets a reference to the given string and assigns it to the InAppLabel field.
func (o *BillingDimensionsMappingBodyItemAttributes) SetInAppLabel(v string) {
	o.InAppLabel = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BillingDimensionsMappingBodyItemAttributes) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingDimensionsMappingBodyItemAttributes) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BillingDimensionsMappingBodyItemAttributes) HasTimestamp() bool {
	return o != nil && o.Timestamp != nil
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *BillingDimensionsMappingBodyItemAttributes) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BillingDimensionsMappingBodyItemAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Endpoints != nil {
		toSerialize["endpoints"] = o.Endpoints
	}
	if o.InAppLabel != nil {
		toSerialize["in_app_label"] = o.InAppLabel
	}
	if o.Timestamp != nil {
		if o.Timestamp.Nanosecond() == 0 {
			toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["timestamp"] = o.Timestamp.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *BillingDimensionsMappingBodyItemAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Endpoints  []BillingDimensionsMappingBodyItemAttributesEndpointsItems `json:"endpoints,omitempty"`
		InAppLabel *string                                                    `json:"in_app_label,omitempty"`
		Timestamp  *time.Time                                                 `json:"timestamp,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"endpoints", "in_app_label", "timestamp"})
	} else {
		return err
	}
	o.Endpoints = all.Endpoints
	o.InAppLabel = all.InAppLabel
	o.Timestamp = all.Timestamp

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
