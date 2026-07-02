// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ComparisonDuration The comparison period. Use a preset `type` value or set `type` to `custom_timeframe` and provide `custom_timeframe` with explicit millisecond epoch bounds.
type ComparisonDuration struct {
	// Fixed time range for a `custom_timeframe` comparison.
	CustomTimeframe *ComparisonCustomTimeframe `json:"custom_timeframe,omitempty"`
	// The comparison window type.
	Type ComparisonDurationType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewComparisonDuration instantiates a new ComparisonDuration object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewComparisonDuration(typeVar ComparisonDurationType) *ComparisonDuration {
	this := ComparisonDuration{}
	this.Type = typeVar
	return &this
}

// NewComparisonDurationWithDefaults instantiates a new ComparisonDuration object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewComparisonDurationWithDefaults() *ComparisonDuration {
	this := ComparisonDuration{}
	return &this
}

// GetCustomTimeframe returns the CustomTimeframe field value if set, zero value otherwise.
func (o *ComparisonDuration) GetCustomTimeframe() ComparisonCustomTimeframe {
	if o == nil || o.CustomTimeframe == nil {
		var ret ComparisonCustomTimeframe
		return ret
	}
	return *o.CustomTimeframe
}

// GetCustomTimeframeOk returns a tuple with the CustomTimeframe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComparisonDuration) GetCustomTimeframeOk() (*ComparisonCustomTimeframe, bool) {
	if o == nil || o.CustomTimeframe == nil {
		return nil, false
	}
	return o.CustomTimeframe, true
}

// HasCustomTimeframe returns a boolean if a field has been set.
func (o *ComparisonDuration) HasCustomTimeframe() bool {
	return o != nil && o.CustomTimeframe != nil
}

// SetCustomTimeframe gets a reference to the given ComparisonCustomTimeframe and assigns it to the CustomTimeframe field.
func (o *ComparisonDuration) SetCustomTimeframe(v ComparisonCustomTimeframe) {
	o.CustomTimeframe = &v
}

// GetType returns the Type field value.
func (o *ComparisonDuration) GetType() ComparisonDurationType {
	if o == nil {
		var ret ComparisonDurationType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ComparisonDuration) GetTypeOk() (*ComparisonDurationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *ComparisonDuration) SetType(v ComparisonDurationType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ComparisonDuration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CustomTimeframe != nil {
		toSerialize["custom_timeframe"] = o.CustomTimeframe
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ComparisonDuration) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CustomTimeframe *ComparisonCustomTimeframe `json:"custom_timeframe,omitempty"`
		Type            *ComparisonDurationType    `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"custom_timeframe", "type"})
	} else {
		return err
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

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
