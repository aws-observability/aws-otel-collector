// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RetentionFilterUpdateAttributes The object describing the configuration of the retention filter to create/update.
type RetentionFilterUpdateAttributes struct {
	// Enable/Disable the retention filter.
	Enabled bool `json:"enabled"`
	// The spans filter. Spans matching this filter will be indexed and stored.
	Filter SpansFilterCreate `json:"filter"`
	// The type of retention filter.
	FilterType RetentionFilterAllType `json:"filter_type"`
	// The name of the retention filter.
	Name string `json:"name"`
	// Sample rate to apply to spans going through this retention filter,
	// a value of 1.0 keeps all spans matching the query.
	Rate float64 `json:"rate"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRetentionFilterUpdateAttributes instantiates a new RetentionFilterUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRetentionFilterUpdateAttributes(enabled bool, filter SpansFilterCreate, filterType RetentionFilterAllType, name string, rate float64) *RetentionFilterUpdateAttributes {
	this := RetentionFilterUpdateAttributes{}
	this.Enabled = enabled
	this.Filter = filter
	this.FilterType = filterType
	this.Name = name
	this.Rate = rate
	return &this
}

// NewRetentionFilterUpdateAttributesWithDefaults instantiates a new RetentionFilterUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRetentionFilterUpdateAttributesWithDefaults() *RetentionFilterUpdateAttributes {
	this := RetentionFilterUpdateAttributes{}
	var filterType RetentionFilterAllType = RETENTIONFILTERALLTYPE_SPANS_SAMPLING_PROCESSOR
	this.FilterType = filterType
	return &this
}

// GetEnabled returns the Enabled field value.
func (o *RetentionFilterUpdateAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RetentionFilterUpdateAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *RetentionFilterUpdateAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFilter returns the Filter field value.
func (o *RetentionFilterUpdateAttributes) GetFilter() SpansFilterCreate {
	if o == nil {
		var ret SpansFilterCreate
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *RetentionFilterUpdateAttributes) GetFilterOk() (*SpansFilterCreate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value.
func (o *RetentionFilterUpdateAttributes) SetFilter(v SpansFilterCreate) {
	o.Filter = v
}

// GetFilterType returns the FilterType field value.
func (o *RetentionFilterUpdateAttributes) GetFilterType() RetentionFilterAllType {
	if o == nil {
		var ret RetentionFilterAllType
		return ret
	}
	return o.FilterType
}

// GetFilterTypeOk returns a tuple with the FilterType field value
// and a boolean to check if the value has been set.
func (o *RetentionFilterUpdateAttributes) GetFilterTypeOk() (*RetentionFilterAllType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterType, true
}

// SetFilterType sets field value.
func (o *RetentionFilterUpdateAttributes) SetFilterType(v RetentionFilterAllType) {
	o.FilterType = v
}

// GetName returns the Name field value.
func (o *RetentionFilterUpdateAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RetentionFilterUpdateAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *RetentionFilterUpdateAttributes) SetName(v string) {
	o.Name = v
}

// GetRate returns the Rate field value.
func (o *RetentionFilterUpdateAttributes) GetRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.Rate
}

// GetRateOk returns a tuple with the Rate field value
// and a boolean to check if the value has been set.
func (o *RetentionFilterUpdateAttributes) GetRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rate, true
}

// SetRate sets field value.
func (o *RetentionFilterUpdateAttributes) SetRate(v float64) {
	o.Rate = v
}

// MarshalJSON serializes the struct using spec logic.
func (o RetentionFilterUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["enabled"] = o.Enabled
	toSerialize["filter"] = o.Filter
	toSerialize["filter_type"] = o.FilterType
	toSerialize["name"] = o.Name
	toSerialize["rate"] = o.Rate

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RetentionFilterUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Enabled    *bool                   `json:"enabled"`
		Filter     *SpansFilterCreate      `json:"filter"`
		FilterType *RetentionFilterAllType `json:"filter_type"`
		Name       *string                 `json:"name"`
		Rate       *float64                `json:"rate"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Filter == nil {
		return fmt.Errorf("required field filter missing")
	}
	if all.FilterType == nil {
		return fmt.Errorf("required field filter_type missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Rate == nil {
		return fmt.Errorf("required field rate missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"enabled", "filter", "filter_type", "name", "rate"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Enabled = *all.Enabled
	if all.Filter.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Filter = *all.Filter
	if !all.FilterType.IsValid() {
		hasInvalidField = true
	} else {
		o.FilterType = *all.FilterType
	}
	o.Name = *all.Name
	o.Rate = *all.Rate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
