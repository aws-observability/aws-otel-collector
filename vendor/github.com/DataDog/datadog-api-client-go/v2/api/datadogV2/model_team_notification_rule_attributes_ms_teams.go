// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamNotificationRuleAttributesMsTeams MS Teams notification settings for the team
type TeamNotificationRuleAttributesMsTeams struct {
	// Handle for MS Teams
	ConnectorName *string `json:"connector_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamNotificationRuleAttributesMsTeams instantiates a new TeamNotificationRuleAttributesMsTeams object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamNotificationRuleAttributesMsTeams() *TeamNotificationRuleAttributesMsTeams {
	this := TeamNotificationRuleAttributesMsTeams{}
	return &this
}

// NewTeamNotificationRuleAttributesMsTeamsWithDefaults instantiates a new TeamNotificationRuleAttributesMsTeams object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamNotificationRuleAttributesMsTeamsWithDefaults() *TeamNotificationRuleAttributesMsTeams {
	this := TeamNotificationRuleAttributesMsTeams{}
	return &this
}

// GetConnectorName returns the ConnectorName field value if set, zero value otherwise.
func (o *TeamNotificationRuleAttributesMsTeams) GetConnectorName() string {
	if o == nil || o.ConnectorName == nil {
		var ret string
		return ret
	}
	return *o.ConnectorName
}

// GetConnectorNameOk returns a tuple with the ConnectorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRuleAttributesMsTeams) GetConnectorNameOk() (*string, bool) {
	if o == nil || o.ConnectorName == nil {
		return nil, false
	}
	return o.ConnectorName, true
}

// HasConnectorName returns a boolean if a field has been set.
func (o *TeamNotificationRuleAttributesMsTeams) HasConnectorName() bool {
	return o != nil && o.ConnectorName != nil
}

// SetConnectorName gets a reference to the given string and assigns it to the ConnectorName field.
func (o *TeamNotificationRuleAttributesMsTeams) SetConnectorName(v string) {
	o.ConnectorName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamNotificationRuleAttributesMsTeams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ConnectorName != nil {
		toSerialize["connector_name"] = o.ConnectorName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamNotificationRuleAttributesMsTeams) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ConnectorName *string `json:"connector_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"connector_name"})
	} else {
		return err
	}
	o.ConnectorName = all.ConnectorName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
