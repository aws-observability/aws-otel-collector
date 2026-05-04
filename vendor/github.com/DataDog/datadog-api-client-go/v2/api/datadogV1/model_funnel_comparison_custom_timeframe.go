// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelComparisonCustomTimeframe Custom timeframe for funnel comparison.
type FunnelComparisonCustomTimeframe struct {
	// Start of the custom timeframe.
	From float64 `json:"from"`
	// End of the custom timeframe.
	To float64 `json:"to"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFunnelComparisonCustomTimeframe instantiates a new FunnelComparisonCustomTimeframe object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFunnelComparisonCustomTimeframe(from float64, to float64) *FunnelComparisonCustomTimeframe {
	this := FunnelComparisonCustomTimeframe{}
	this.From = from
	this.To = to
	return &this
}

// NewFunnelComparisonCustomTimeframeWithDefaults instantiates a new FunnelComparisonCustomTimeframe object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFunnelComparisonCustomTimeframeWithDefaults() *FunnelComparisonCustomTimeframe {
	this := FunnelComparisonCustomTimeframe{}
	return &this
}

// GetFrom returns the From field value.
func (o *FunnelComparisonCustomTimeframe) GetFrom() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *FunnelComparisonCustomTimeframe) GetFromOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value.
func (o *FunnelComparisonCustomTimeframe) SetFrom(v float64) {
	o.From = v
}

// GetTo returns the To field value.
func (o *FunnelComparisonCustomTimeframe) GetTo() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *FunnelComparisonCustomTimeframe) GetToOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value.
func (o *FunnelComparisonCustomTimeframe) SetTo(v float64) {
	o.To = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FunnelComparisonCustomTimeframe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FunnelComparisonCustomTimeframe) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		From *float64 `json:"from"`
		To   *float64 `json:"to"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.From == nil {
		return fmt.Errorf("required field from missing")
	}
	if all.To == nil {
		return fmt.Errorf("required field to missing")
	}
	o.From = *all.From
	o.To = *all.To

	return nil
}
