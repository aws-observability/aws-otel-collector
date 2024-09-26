// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActiveBillingDimensionsAttributes List of active billing dimensions.
type ActiveBillingDimensionsAttributes struct {
	// Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]`.
	Month *time.Time `json:"month,omitempty"`
	// List of active billing dimensions. Example: `[infra_host, apm_host, serverless_infra]`.
	Values []string `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActiveBillingDimensionsAttributes instantiates a new ActiveBillingDimensionsAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActiveBillingDimensionsAttributes() *ActiveBillingDimensionsAttributes {
	this := ActiveBillingDimensionsAttributes{}
	return &this
}

// NewActiveBillingDimensionsAttributesWithDefaults instantiates a new ActiveBillingDimensionsAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActiveBillingDimensionsAttributesWithDefaults() *ActiveBillingDimensionsAttributes {
	this := ActiveBillingDimensionsAttributes{}
	return &this
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *ActiveBillingDimensionsAttributes) GetMonth() time.Time {
	if o == nil || o.Month == nil {
		var ret time.Time
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveBillingDimensionsAttributes) GetMonthOk() (*time.Time, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *ActiveBillingDimensionsAttributes) HasMonth() bool {
	return o != nil && o.Month != nil
}

// SetMonth gets a reference to the given time.Time and assigns it to the Month field.
func (o *ActiveBillingDimensionsAttributes) SetMonth(v time.Time) {
	o.Month = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ActiveBillingDimensionsAttributes) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveBillingDimensionsAttributes) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ActiveBillingDimensionsAttributes) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *ActiveBillingDimensionsAttributes) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActiveBillingDimensionsAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Month != nil {
		if o.Month.Nanosecond() == 0 {
			toSerialize["month"] = o.Month.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["month"] = o.Month.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActiveBillingDimensionsAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Month  *time.Time `json:"month,omitempty"`
		Values []string   `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"month", "values"})
	} else {
		return err
	}
	o.Month = all.Month
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
