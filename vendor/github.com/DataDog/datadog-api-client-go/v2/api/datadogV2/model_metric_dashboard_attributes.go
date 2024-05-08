// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MetricDashboardAttributes Attributes related to the dashboard, including title and popularity.
type MetricDashboardAttributes struct {
	// Value from 0 to 5 that ranks popularity of the dashboard.
	Popularity *float64 `json:"popularity,omitempty"`
	// Title of the asset.
	Title *string `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{}
}

// NewMetricDashboardAttributes instantiates a new MetricDashboardAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMetricDashboardAttributes() *MetricDashboardAttributes {
	this := MetricDashboardAttributes{}
	return &this
}

// NewMetricDashboardAttributesWithDefaults instantiates a new MetricDashboardAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMetricDashboardAttributesWithDefaults() *MetricDashboardAttributes {
	this := MetricDashboardAttributes{}
	return &this
}

// GetPopularity returns the Popularity field value if set, zero value otherwise.
func (o *MetricDashboardAttributes) GetPopularity() float64 {
	if o == nil || o.Popularity == nil {
		var ret float64
		return ret
	}
	return *o.Popularity
}

// GetPopularityOk returns a tuple with the Popularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDashboardAttributes) GetPopularityOk() (*float64, bool) {
	if o == nil || o.Popularity == nil {
		return nil, false
	}
	return o.Popularity, true
}

// HasPopularity returns a boolean if a field has been set.
func (o *MetricDashboardAttributes) HasPopularity() bool {
	return o != nil && o.Popularity != nil
}

// SetPopularity gets a reference to the given float64 and assigns it to the Popularity field.
func (o *MetricDashboardAttributes) SetPopularity(v float64) {
	o.Popularity = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MetricDashboardAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDashboardAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MetricDashboardAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MetricDashboardAttributes) SetTitle(v string) {
	o.Title = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MetricDashboardAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Popularity != nil {
		toSerialize["popularity"] = o.Popularity
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MetricDashboardAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Popularity *float64 `json:"popularity,omitempty"`
		Title      *string  `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"popularity", "title"})
	} else {
		return err
	}
	o.Popularity = all.Popularity
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
