// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TeamNotificationRuleAttributes Team notification rule attributes
type TeamNotificationRuleAttributes struct {
	// Email notification settings for the team
	Email *TeamNotificationRuleAttributesEmail `json:"email,omitempty"`
	// MS Teams notification settings for the team
	MsTeams *TeamNotificationRuleAttributesMsTeams `json:"ms_teams,omitempty"`
	// PagerDuty notification settings for the team
	Pagerduty *TeamNotificationRuleAttributesPagerduty `json:"pagerduty,omitempty"`
	// Slack notification settings for the team
	Slack *TeamNotificationRuleAttributesSlack `json:"slack,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTeamNotificationRuleAttributes instantiates a new TeamNotificationRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTeamNotificationRuleAttributes() *TeamNotificationRuleAttributes {
	this := TeamNotificationRuleAttributes{}
	return &this
}

// NewTeamNotificationRuleAttributesWithDefaults instantiates a new TeamNotificationRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTeamNotificationRuleAttributesWithDefaults() *TeamNotificationRuleAttributes {
	this := TeamNotificationRuleAttributes{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *TeamNotificationRuleAttributes) GetEmail() TeamNotificationRuleAttributesEmail {
	if o == nil || o.Email == nil {
		var ret TeamNotificationRuleAttributesEmail
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRuleAttributes) GetEmailOk() (*TeamNotificationRuleAttributesEmail, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *TeamNotificationRuleAttributes) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given TeamNotificationRuleAttributesEmail and assigns it to the Email field.
func (o *TeamNotificationRuleAttributes) SetEmail(v TeamNotificationRuleAttributesEmail) {
	o.Email = &v
}

// GetMsTeams returns the MsTeams field value if set, zero value otherwise.
func (o *TeamNotificationRuleAttributes) GetMsTeams() TeamNotificationRuleAttributesMsTeams {
	if o == nil || o.MsTeams == nil {
		var ret TeamNotificationRuleAttributesMsTeams
		return ret
	}
	return *o.MsTeams
}

// GetMsTeamsOk returns a tuple with the MsTeams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRuleAttributes) GetMsTeamsOk() (*TeamNotificationRuleAttributesMsTeams, bool) {
	if o == nil || o.MsTeams == nil {
		return nil, false
	}
	return o.MsTeams, true
}

// HasMsTeams returns a boolean if a field has been set.
func (o *TeamNotificationRuleAttributes) HasMsTeams() bool {
	return o != nil && o.MsTeams != nil
}

// SetMsTeams gets a reference to the given TeamNotificationRuleAttributesMsTeams and assigns it to the MsTeams field.
func (o *TeamNotificationRuleAttributes) SetMsTeams(v TeamNotificationRuleAttributesMsTeams) {
	o.MsTeams = &v
}

// GetPagerduty returns the Pagerduty field value if set, zero value otherwise.
func (o *TeamNotificationRuleAttributes) GetPagerduty() TeamNotificationRuleAttributesPagerduty {
	if o == nil || o.Pagerduty == nil {
		var ret TeamNotificationRuleAttributesPagerduty
		return ret
	}
	return *o.Pagerduty
}

// GetPagerdutyOk returns a tuple with the Pagerduty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRuleAttributes) GetPagerdutyOk() (*TeamNotificationRuleAttributesPagerduty, bool) {
	if o == nil || o.Pagerduty == nil {
		return nil, false
	}
	return o.Pagerduty, true
}

// HasPagerduty returns a boolean if a field has been set.
func (o *TeamNotificationRuleAttributes) HasPagerduty() bool {
	return o != nil && o.Pagerduty != nil
}

// SetPagerduty gets a reference to the given TeamNotificationRuleAttributesPagerduty and assigns it to the Pagerduty field.
func (o *TeamNotificationRuleAttributes) SetPagerduty(v TeamNotificationRuleAttributesPagerduty) {
	o.Pagerduty = &v
}

// GetSlack returns the Slack field value if set, zero value otherwise.
func (o *TeamNotificationRuleAttributes) GetSlack() TeamNotificationRuleAttributesSlack {
	if o == nil || o.Slack == nil {
		var ret TeamNotificationRuleAttributesSlack
		return ret
	}
	return *o.Slack
}

// GetSlackOk returns a tuple with the Slack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TeamNotificationRuleAttributes) GetSlackOk() (*TeamNotificationRuleAttributesSlack, bool) {
	if o == nil || o.Slack == nil {
		return nil, false
	}
	return o.Slack, true
}

// HasSlack returns a boolean if a field has been set.
func (o *TeamNotificationRuleAttributes) HasSlack() bool {
	return o != nil && o.Slack != nil
}

// SetSlack gets a reference to the given TeamNotificationRuleAttributesSlack and assigns it to the Slack field.
func (o *TeamNotificationRuleAttributes) SetSlack(v TeamNotificationRuleAttributesSlack) {
	o.Slack = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TeamNotificationRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.MsTeams != nil {
		toSerialize["ms_teams"] = o.MsTeams
	}
	if o.Pagerduty != nil {
		toSerialize["pagerduty"] = o.Pagerduty
	}
	if o.Slack != nil {
		toSerialize["slack"] = o.Slack
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TeamNotificationRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Email     *TeamNotificationRuleAttributesEmail     `json:"email,omitempty"`
		MsTeams   *TeamNotificationRuleAttributesMsTeams   `json:"ms_teams,omitempty"`
		Pagerduty *TeamNotificationRuleAttributesPagerduty `json:"pagerduty,omitempty"`
		Slack     *TeamNotificationRuleAttributesSlack     `json:"slack,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"email", "ms_teams", "pagerduty", "slack"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Email != nil && all.Email.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Email = all.Email
	if all.MsTeams != nil && all.MsTeams.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MsTeams = all.MsTeams
	if all.Pagerduty != nil && all.Pagerduty.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Pagerduty = all.Pagerduty
	if all.Slack != nil && all.Slack.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Slack = all.Slack

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
