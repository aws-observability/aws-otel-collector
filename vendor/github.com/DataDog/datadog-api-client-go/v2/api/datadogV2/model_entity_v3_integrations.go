// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// EntityV3Integrations A base schema for defining third-party integrations.
type EntityV3Integrations struct {
	// An Opsgenie integration schema.
	Opsgenie *EntityV3DatadogIntegrationOpsgenie `json:"opsgenie,omitempty"`
	// A PagerDuty integration schema.
	Pagerduty *EntityV3DatadogIntegrationPagerduty `json:"pagerduty,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject map[string]interface{} `json:"-"`
}

// NewEntityV3Integrations instantiates a new EntityV3Integrations object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewEntityV3Integrations() *EntityV3Integrations {
	this := EntityV3Integrations{}
	return &this
}

// NewEntityV3IntegrationsWithDefaults instantiates a new EntityV3Integrations object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewEntityV3IntegrationsWithDefaults() *EntityV3Integrations {
	this := EntityV3Integrations{}
	return &this
}

// GetOpsgenie returns the Opsgenie field value if set, zero value otherwise.
func (o *EntityV3Integrations) GetOpsgenie() EntityV3DatadogIntegrationOpsgenie {
	if o == nil || o.Opsgenie == nil {
		var ret EntityV3DatadogIntegrationOpsgenie
		return ret
	}
	return *o.Opsgenie
}

// GetOpsgenieOk returns a tuple with the Opsgenie field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3Integrations) GetOpsgenieOk() (*EntityV3DatadogIntegrationOpsgenie, bool) {
	if o == nil || o.Opsgenie == nil {
		return nil, false
	}
	return o.Opsgenie, true
}

// HasOpsgenie returns a boolean if a field has been set.
func (o *EntityV3Integrations) HasOpsgenie() bool {
	return o != nil && o.Opsgenie != nil
}

// SetOpsgenie gets a reference to the given EntityV3DatadogIntegrationOpsgenie and assigns it to the Opsgenie field.
func (o *EntityV3Integrations) SetOpsgenie(v EntityV3DatadogIntegrationOpsgenie) {
	o.Opsgenie = &v
}

// GetPagerduty returns the Pagerduty field value if set, zero value otherwise.
func (o *EntityV3Integrations) GetPagerduty() EntityV3DatadogIntegrationPagerduty {
	if o == nil || o.Pagerduty == nil {
		var ret EntityV3DatadogIntegrationPagerduty
		return ret
	}
	return *o.Pagerduty
}

// GetPagerdutyOk returns a tuple with the Pagerduty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntityV3Integrations) GetPagerdutyOk() (*EntityV3DatadogIntegrationPagerduty, bool) {
	if o == nil || o.Pagerduty == nil {
		return nil, false
	}
	return o.Pagerduty, true
}

// HasPagerduty returns a boolean if a field has been set.
func (o *EntityV3Integrations) HasPagerduty() bool {
	return o != nil && o.Pagerduty != nil
}

// SetPagerduty gets a reference to the given EntityV3DatadogIntegrationPagerduty and assigns it to the Pagerduty field.
func (o *EntityV3Integrations) SetPagerduty(v EntityV3DatadogIntegrationPagerduty) {
	o.Pagerduty = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o EntityV3Integrations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Opsgenie != nil {
		toSerialize["opsgenie"] = o.Opsgenie
	}
	if o.Pagerduty != nil {
		toSerialize["pagerduty"] = o.Pagerduty
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *EntityV3Integrations) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Opsgenie  *EntityV3DatadogIntegrationOpsgenie  `json:"opsgenie,omitempty"`
		Pagerduty *EntityV3DatadogIntegrationPagerduty `json:"pagerduty,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	hasInvalidField := false
	if all.Opsgenie != nil && all.Opsgenie.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Opsgenie = all.Opsgenie
	if all.Pagerduty != nil && all.Pagerduty.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pagerduty = all.Pagerduty

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
