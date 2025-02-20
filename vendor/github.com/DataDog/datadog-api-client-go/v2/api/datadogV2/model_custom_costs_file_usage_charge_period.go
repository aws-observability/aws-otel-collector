// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CustomCostsFileUsageChargePeriod Usage charge period of a Custom Costs file.
type CustomCostsFileUsageChargePeriod struct {
	// End of the usage of the Custom Costs file.
	End *float64 `json:"end,omitempty"`
	// Start of the usage of the Custom Costs file.
	Start *float64 `json:"start,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCustomCostsFileUsageChargePeriod instantiates a new CustomCostsFileUsageChargePeriod object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCustomCostsFileUsageChargePeriod() *CustomCostsFileUsageChargePeriod {
	this := CustomCostsFileUsageChargePeriod{}
	return &this
}

// NewCustomCostsFileUsageChargePeriodWithDefaults instantiates a new CustomCostsFileUsageChargePeriod object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCustomCostsFileUsageChargePeriodWithDefaults() *CustomCostsFileUsageChargePeriod {
	this := CustomCostsFileUsageChargePeriod{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *CustomCostsFileUsageChargePeriod) GetEnd() float64 {
	if o == nil || o.End == nil {
		var ret float64
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileUsageChargePeriod) GetEndOk() (*float64, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *CustomCostsFileUsageChargePeriod) HasEnd() bool {
	return o != nil && o.End != nil
}

// SetEnd gets a reference to the given float64 and assigns it to the End field.
func (o *CustomCostsFileUsageChargePeriod) SetEnd(v float64) {
	o.End = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *CustomCostsFileUsageChargePeriod) GetStart() float64 {
	if o == nil || o.Start == nil {
		var ret float64
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomCostsFileUsageChargePeriod) GetStartOk() (*float64, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *CustomCostsFileUsageChargePeriod) HasStart() bool {
	return o != nil && o.Start != nil
}

// SetStart gets a reference to the given float64 and assigns it to the Start field.
func (o *CustomCostsFileUsageChargePeriod) SetStart(v float64) {
	o.Start = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CustomCostsFileUsageChargePeriod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CustomCostsFileUsageChargePeriod) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		End   *float64 `json:"end,omitempty"`
		Start *float64 `json:"start,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"end", "start"})
	} else {
		return err
	}
	o.End = all.End
	o.Start = all.Start

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
