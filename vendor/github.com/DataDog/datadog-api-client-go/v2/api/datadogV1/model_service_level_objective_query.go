// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ServiceLevelObjectiveQuery A count-based (metric) SLO query. This field is superseded by `sli_specification` but is retained for backwards compatibility. Note that Datadog only allows the sum by aggregator
// to be used because this will sum up all request counts instead of averaging them, or taking the max or
// min of all of those requests.
type ServiceLevelObjectiveQuery struct {
	// A Datadog metric query for total (valid) events.
	Denominator string `json:"denominator"`
	// A Datadog metric query for good events.
	Numerator string `json:"numerator"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewServiceLevelObjectiveQuery instantiates a new ServiceLevelObjectiveQuery object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewServiceLevelObjectiveQuery(denominator string, numerator string) *ServiceLevelObjectiveQuery {
	this := ServiceLevelObjectiveQuery{}
	this.Denominator = denominator
	this.Numerator = numerator
	return &this
}

// NewServiceLevelObjectiveQueryWithDefaults instantiates a new ServiceLevelObjectiveQuery object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewServiceLevelObjectiveQueryWithDefaults() *ServiceLevelObjectiveQuery {
	this := ServiceLevelObjectiveQuery{}
	return &this
}

// GetDenominator returns the Denominator field value.
func (o *ServiceLevelObjectiveQuery) GetDenominator() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Denominator
}

// GetDenominatorOk returns a tuple with the Denominator field value
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjectiveQuery) GetDenominatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Denominator, true
}

// SetDenominator sets field value.
func (o *ServiceLevelObjectiveQuery) SetDenominator(v string) {
	o.Denominator = v
}

// GetNumerator returns the Numerator field value.
func (o *ServiceLevelObjectiveQuery) GetNumerator() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Numerator
}

// GetNumeratorOk returns a tuple with the Numerator field value
// and a boolean to check if the value has been set.
func (o *ServiceLevelObjectiveQuery) GetNumeratorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Numerator, true
}

// SetNumerator sets field value.
func (o *ServiceLevelObjectiveQuery) SetNumerator(v string) {
	o.Numerator = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ServiceLevelObjectiveQuery) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["denominator"] = o.Denominator
	toSerialize["numerator"] = o.Numerator

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ServiceLevelObjectiveQuery) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Denominator *string `json:"denominator"`
		Numerator   *string `json:"numerator"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Denominator == nil {
		return fmt.Errorf("required field denominator missing")
	}
	if all.Numerator == nil {
		return fmt.Errorf("required field numerator missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"denominator", "numerator"})
	} else {
		return err
	}
	o.Denominator = *all.Denominator
	o.Numerator = *all.Numerator

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
