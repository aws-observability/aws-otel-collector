// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GlobalIncidentSettingsAttributesResponse Global incident settings attributes
type GlobalIncidentSettingsAttributesResponse struct {
	// The analytics dashboard ID
	AnalyticsDashboardId string `json:"analytics_dashboard_id"`
	// Timestamp when the settings were created
	Created time.Time `json:"created"`
	// Timestamp when the settings were last modified
	Modified time.Time `json:"modified"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGlobalIncidentSettingsAttributesResponse instantiates a new GlobalIncidentSettingsAttributesResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGlobalIncidentSettingsAttributesResponse(analyticsDashboardId string, created time.Time, modified time.Time) *GlobalIncidentSettingsAttributesResponse {
	this := GlobalIncidentSettingsAttributesResponse{}
	this.AnalyticsDashboardId = analyticsDashboardId
	this.Created = created
	this.Modified = modified
	return &this
}

// NewGlobalIncidentSettingsAttributesResponseWithDefaults instantiates a new GlobalIncidentSettingsAttributesResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGlobalIncidentSettingsAttributesResponseWithDefaults() *GlobalIncidentSettingsAttributesResponse {
	this := GlobalIncidentSettingsAttributesResponse{}
	return &this
}

// GetAnalyticsDashboardId returns the AnalyticsDashboardId field value.
func (o *GlobalIncidentSettingsAttributesResponse) GetAnalyticsDashboardId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AnalyticsDashboardId
}

// GetAnalyticsDashboardIdOk returns a tuple with the AnalyticsDashboardId field value
// and a boolean to check if the value has been set.
func (o *GlobalIncidentSettingsAttributesResponse) GetAnalyticsDashboardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnalyticsDashboardId, true
}

// SetAnalyticsDashboardId sets field value.
func (o *GlobalIncidentSettingsAttributesResponse) SetAnalyticsDashboardId(v string) {
	o.AnalyticsDashboardId = v
}

// GetCreated returns the Created field value.
func (o *GlobalIncidentSettingsAttributesResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *GlobalIncidentSettingsAttributesResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value.
func (o *GlobalIncidentSettingsAttributesResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value.
func (o *GlobalIncidentSettingsAttributesResponse) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *GlobalIncidentSettingsAttributesResponse) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value.
func (o *GlobalIncidentSettingsAttributesResponse) SetModified(v time.Time) {
	o.Modified = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GlobalIncidentSettingsAttributesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["analytics_dashboard_id"] = o.AnalyticsDashboardId
	if o.Created.Nanosecond() == 0 {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["created"] = o.Created.Format("2006-01-02T15:04:05.000Z07:00")
	}
	if o.Modified.Nanosecond() == 0 {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05Z07:00")
	} else {
		toSerialize["modified"] = o.Modified.Format("2006-01-02T15:04:05.000Z07:00")
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GlobalIncidentSettingsAttributesResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AnalyticsDashboardId *string    `json:"analytics_dashboard_id"`
		Created              *time.Time `json:"created"`
		Modified             *time.Time `json:"modified"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AnalyticsDashboardId == nil {
		return fmt.Errorf("required field analytics_dashboard_id missing")
	}
	if all.Created == nil {
		return fmt.Errorf("required field created missing")
	}
	if all.Modified == nil {
		return fmt.Errorf("required field modified missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"analytics_dashboard_id", "created", "modified"})
	} else {
		return err
	}
	o.AnalyticsDashboardId = *all.AnalyticsDashboardId
	o.Created = *all.Created
	o.Modified = *all.Modified

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
