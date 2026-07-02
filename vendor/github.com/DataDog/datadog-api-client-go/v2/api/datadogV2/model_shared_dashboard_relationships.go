// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SharedDashboardRelationships Relationships of a shared dashboard.
type SharedDashboardRelationships struct {
	// Dashboard associated with the shared dashboard.
	Dashboard SharedDashboardRelationshipDashboard `json:"dashboard"`
	// User who shared the dashboard.
	Sharer SharedDashboardRelationshipSharer `json:"sharer"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSharedDashboardRelationships instantiates a new SharedDashboardRelationships object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSharedDashboardRelationships(dashboard SharedDashboardRelationshipDashboard, sharer SharedDashboardRelationshipSharer) *SharedDashboardRelationships {
	this := SharedDashboardRelationships{}
	this.Dashboard = dashboard
	this.Sharer = sharer
	return &this
}

// NewSharedDashboardRelationshipsWithDefaults instantiates a new SharedDashboardRelationships object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSharedDashboardRelationshipsWithDefaults() *SharedDashboardRelationships {
	this := SharedDashboardRelationships{}
	return &this
}

// GetDashboard returns the Dashboard field value.
func (o *SharedDashboardRelationships) GetDashboard() SharedDashboardRelationshipDashboard {
	if o == nil {
		var ret SharedDashboardRelationshipDashboard
		return ret
	}
	return o.Dashboard
}

// GetDashboardOk returns a tuple with the Dashboard field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardRelationships) GetDashboardOk() (*SharedDashboardRelationshipDashboard, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dashboard, true
}

// SetDashboard sets field value.
func (o *SharedDashboardRelationships) SetDashboard(v SharedDashboardRelationshipDashboard) {
	o.Dashboard = v
}

// GetSharer returns the Sharer field value.
func (o *SharedDashboardRelationships) GetSharer() SharedDashboardRelationshipSharer {
	if o == nil {
		var ret SharedDashboardRelationshipSharer
		return ret
	}
	return o.Sharer
}

// GetSharerOk returns a tuple with the Sharer field value
// and a boolean to check if the value has been set.
func (o *SharedDashboardRelationships) GetSharerOk() (*SharedDashboardRelationshipSharer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sharer, true
}

// SetSharer sets field value.
func (o *SharedDashboardRelationships) SetSharer(v SharedDashboardRelationshipSharer) {
	o.Sharer = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SharedDashboardRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dashboard"] = o.Dashboard
	toSerialize["sharer"] = o.Sharer

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SharedDashboardRelationships) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Dashboard *SharedDashboardRelationshipDashboard `json:"dashboard"`
		Sharer    *SharedDashboardRelationshipSharer    `json:"sharer"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Dashboard == nil {
		return fmt.Errorf("required field dashboard missing")
	}
	if all.Sharer == nil {
		return fmt.Errorf("required field sharer missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dashboard", "sharer"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Dashboard.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Dashboard = *all.Dashboard
	if all.Sharer.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Sharer = *all.Sharer

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
