// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// QueryValueWidgetComparison A change indicator that compares the current value to a historical period.
type QueryValueWidgetComparison struct {
	// Color-coding direction: `increase_better` (green on rise), `decrease_better` (green on drop), or `neutral` (no color).
	Directionality *QueryValueWidgetComparisonDirectionality `json:"directionality,omitempty"`
	// The comparison period. Use a preset `type` value or set `type` to `custom_timeframe` and provide `custom_timeframe` with explicit millisecond epoch bounds.
	Duration ComparisonDuration `json:"duration"`
	// How the delta is expressed: `absolute` (raw difference), `relative` (percentage), or `both`.
	Type *QueryValueWidgetComparisonType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewQueryValueWidgetComparison instantiates a new QueryValueWidgetComparison object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewQueryValueWidgetComparison(duration ComparisonDuration) *QueryValueWidgetComparison {
	this := QueryValueWidgetComparison{}
	var directionality QueryValueWidgetComparisonDirectionality = QUERYVALUEWIDGETCOMPARISONDIRECTIONALITY_NEUTRAL
	this.Directionality = &directionality
	this.Duration = duration
	var typeVar QueryValueWidgetComparisonType = QUERYVALUEWIDGETCOMPARISONTYPE_ABSOLUTE
	this.Type = &typeVar
	return &this
}

// NewQueryValueWidgetComparisonWithDefaults instantiates a new QueryValueWidgetComparison object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewQueryValueWidgetComparisonWithDefaults() *QueryValueWidgetComparison {
	this := QueryValueWidgetComparison{}
	var directionality QueryValueWidgetComparisonDirectionality = QUERYVALUEWIDGETCOMPARISONDIRECTIONALITY_NEUTRAL
	this.Directionality = &directionality
	var typeVar QueryValueWidgetComparisonType = QUERYVALUEWIDGETCOMPARISONTYPE_ABSOLUTE
	this.Type = &typeVar
	return &this
}

// GetDirectionality returns the Directionality field value if set, zero value otherwise.
func (o *QueryValueWidgetComparison) GetDirectionality() QueryValueWidgetComparisonDirectionality {
	if o == nil || o.Directionality == nil {
		var ret QueryValueWidgetComparisonDirectionality
		return ret
	}
	return *o.Directionality
}

// GetDirectionalityOk returns a tuple with the Directionality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueWidgetComparison) GetDirectionalityOk() (*QueryValueWidgetComparisonDirectionality, bool) {
	if o == nil || o.Directionality == nil {
		return nil, false
	}
	return o.Directionality, true
}

// HasDirectionality returns a boolean if a field has been set.
func (o *QueryValueWidgetComparison) HasDirectionality() bool {
	return o != nil && o.Directionality != nil
}

// SetDirectionality gets a reference to the given QueryValueWidgetComparisonDirectionality and assigns it to the Directionality field.
func (o *QueryValueWidgetComparison) SetDirectionality(v QueryValueWidgetComparisonDirectionality) {
	o.Directionality = &v
}

// GetDuration returns the Duration field value.
func (o *QueryValueWidgetComparison) GetDuration() ComparisonDuration {
	if o == nil {
		var ret ComparisonDuration
		return ret
	}
	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *QueryValueWidgetComparison) GetDurationOk() (*ComparisonDuration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value.
func (o *QueryValueWidgetComparison) SetDuration(v ComparisonDuration) {
	o.Duration = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *QueryValueWidgetComparison) GetType() QueryValueWidgetComparisonType {
	if o == nil || o.Type == nil {
		var ret QueryValueWidgetComparisonType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryValueWidgetComparison) GetTypeOk() (*QueryValueWidgetComparisonType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *QueryValueWidgetComparison) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given QueryValueWidgetComparisonType and assigns it to the Type field.
func (o *QueryValueWidgetComparison) SetType(v QueryValueWidgetComparisonType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o QueryValueWidgetComparison) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Directionality != nil {
		toSerialize["directionality"] = o.Directionality
	}
	toSerialize["duration"] = o.Duration
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *QueryValueWidgetComparison) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Directionality *QueryValueWidgetComparisonDirectionality `json:"directionality,omitempty"`
		Duration       *ComparisonDuration                       `json:"duration"`
		Type           *QueryValueWidgetComparisonType           `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Duration == nil {
		return fmt.Errorf("required field duration missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"directionality", "duration", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Directionality != nil && !all.Directionality.IsValid() {
		hasInvalidField = true
	} else {
		o.Directionality = all.Directionality
	}
	if all.Duration.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Duration = *all.Duration
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
