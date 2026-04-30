// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FunnelComparisonDuration Comparison time configuration for funnel widgets.
type FunnelComparisonDuration struct {
	// Custom timeframe for funnel comparison.
	CustomTimeframe *FunnelComparisonCustomTimeframe `json:"custom_timeframe,omitempty"`
	// Type of comparison duration.
	Type FunnelComparisonDurationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewFunnelComparisonDuration instantiates a new FunnelComparisonDuration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFunnelComparisonDuration(typeVar FunnelComparisonDurationType) *FunnelComparisonDuration {
	this := FunnelComparisonDuration{}
	this.Type = typeVar
	return &this
}

// NewFunnelComparisonDurationWithDefaults instantiates a new FunnelComparisonDuration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFunnelComparisonDurationWithDefaults() *FunnelComparisonDuration {
	this := FunnelComparisonDuration{}
	return &this
}

// GetCustomTimeframe returns the CustomTimeframe field value if set, zero value otherwise.
func (o *FunnelComparisonDuration) GetCustomTimeframe() FunnelComparisonCustomTimeframe {
	if o == nil || o.CustomTimeframe == nil {
		var ret FunnelComparisonCustomTimeframe
		return ret
	}
	return *o.CustomTimeframe
}

// GetCustomTimeframeOk returns a tuple with the CustomTimeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FunnelComparisonDuration) GetCustomTimeframeOk() (*FunnelComparisonCustomTimeframe, bool) {
	if o == nil || o.CustomTimeframe == nil {
		return nil, false
	}
	return o.CustomTimeframe, true
}

// HasCustomTimeframe returns a boolean if a field has been set.
func (o *FunnelComparisonDuration) HasCustomTimeframe() bool {
	return o != nil && o.CustomTimeframe != nil
}

// SetCustomTimeframe gets a reference to the given FunnelComparisonCustomTimeframe and assigns it to the CustomTimeframe field.
func (o *FunnelComparisonDuration) SetCustomTimeframe(v FunnelComparisonCustomTimeframe) {
	o.CustomTimeframe = &v
}

// GetType returns the Type field value.
func (o *FunnelComparisonDuration) GetType() FunnelComparisonDurationType {
	if o == nil {
		var ret FunnelComparisonDurationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FunnelComparisonDuration) GetTypeOk() (*FunnelComparisonDurationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *FunnelComparisonDuration) SetType(v FunnelComparisonDurationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o FunnelComparisonDuration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomTimeframe != nil {
		toSerialize["custom_timeframe"] = o.CustomTimeframe
	}
	toSerialize["type"] = o.Type
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FunnelComparisonDuration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomTimeframe *FunnelComparisonCustomTimeframe `json:"custom_timeframe,omitempty"`
		Type            *FunnelComparisonDurationType    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}

	hasInvalidField := false
	if all.CustomTimeframe != nil && all.CustomTimeframe.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CustomTimeframe = all.CustomTimeframe
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
