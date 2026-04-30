// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GlobalIncidentSettingsAttributesRequest Global incident settings attributes
type GlobalIncidentSettingsAttributesRequest struct {
	// The analytics dashboard ID
	AnalyticsDashboardId *string `json:"analytics_dashboard_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGlobalIncidentSettingsAttributesRequest instantiates a new GlobalIncidentSettingsAttributesRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGlobalIncidentSettingsAttributesRequest() *GlobalIncidentSettingsAttributesRequest {
	this := GlobalIncidentSettingsAttributesRequest{}
	return &this
}

// NewGlobalIncidentSettingsAttributesRequestWithDefaults instantiates a new GlobalIncidentSettingsAttributesRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGlobalIncidentSettingsAttributesRequestWithDefaults() *GlobalIncidentSettingsAttributesRequest {
	this := GlobalIncidentSettingsAttributesRequest{}
	return &this
}

// GetAnalyticsDashboardId returns the AnalyticsDashboardId field value if set, zero value otherwise.
func (o *GlobalIncidentSettingsAttributesRequest) GetAnalyticsDashboardId() string {
	if o == nil || o.AnalyticsDashboardId == nil {
		var ret string
		return ret
	}
	return *o.AnalyticsDashboardId
}

// GetAnalyticsDashboardIdOk returns a tuple with the AnalyticsDashboardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIncidentSettingsAttributesRequest) GetAnalyticsDashboardIdOk() (*string, bool) {
	if o == nil || o.AnalyticsDashboardId == nil {
		return nil, false
	}
	return o.AnalyticsDashboardId, true
}

// HasAnalyticsDashboardId returns a boolean if a field has been set.
func (o *GlobalIncidentSettingsAttributesRequest) HasAnalyticsDashboardId() bool {
	return o != nil && o.AnalyticsDashboardId != nil
}

// SetAnalyticsDashboardId gets a reference to the given string and assigns it to the AnalyticsDashboardId field.
func (o *GlobalIncidentSettingsAttributesRequest) SetAnalyticsDashboardId(v string) {
	o.AnalyticsDashboardId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o GlobalIncidentSettingsAttributesRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AnalyticsDashboardId != nil {
		toSerialize["analytics_dashboard_id"] = o.AnalyticsDashboardId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GlobalIncidentSettingsAttributesRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnalyticsDashboardId *string `json:"analytics_dashboard_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"analytics_dashboard_id"})
	} else {
		return err
	}
	o.AnalyticsDashboardId = all.AnalyticsDashboardId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
