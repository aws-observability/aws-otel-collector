// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CaseNotificationRuleRecipientData Recipient data
type CaseNotificationRuleRecipientData struct {
	// Slack channel name
	Channel *string `json:"channel,omitempty"`
	// Slack channel ID
	ChannelId *string `json:"channel_id,omitempty"`
	// Microsoft Teams channel name
	ChannelName *string `json:"channel_name,omitempty"`
	// Microsoft Teams connector name
	ConnectorName *string `json:"connector_name,omitempty"`
	// Email address
	Email *string `json:"email,omitempty"`
	// HTTP webhook name
	Name *string `json:"name,omitempty"`
	// PagerDuty service name
	ServiceName *string `json:"service_name,omitempty"`
	// Microsoft Teams team ID
	TeamId *string `json:"team_id,omitempty"`
	// Microsoft Teams team name
	TeamName *string `json:"team_name,omitempty"`
	// Microsoft Teams tenant ID
	TenantId *string `json:"tenant_id,omitempty"`
	// Microsoft Teams tenant name
	TenantName *string `json:"tenant_name,omitempty"`
	// Slack workspace name
	Workspace *string `json:"workspace,omitempty"`
	// Slack workspace ID
	WorkspaceId *string `json:"workspace_id,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCaseNotificationRuleRecipientData instantiates a new CaseNotificationRuleRecipientData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCaseNotificationRuleRecipientData() *CaseNotificationRuleRecipientData {
	this := CaseNotificationRuleRecipientData{}
	return &this
}

// NewCaseNotificationRuleRecipientDataWithDefaults instantiates a new CaseNotificationRuleRecipientData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCaseNotificationRuleRecipientDataWithDefaults() *CaseNotificationRuleRecipientData {
	this := CaseNotificationRuleRecipientData{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipientData) GetChannel() string {
	if o == nil || o.Channel == nil {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipientData) GetChannelOk() (*string, bool) {
	if o == nil || o.Channel == nil {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipientData) HasChannel() bool {
	return o != nil && o.Channel != nil
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *CaseNotificationRuleRecipientData) SetChannel(v string) {
	o.Channel = &v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipientData) GetChannelId() string {
	if o == nil || o.ChannelId == nil {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipientData) GetChannelIdOk() (*string, bool) {
	if o == nil || o.ChannelId == nil {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipientData) HasChannelId() bool {
	return o != nil && o.ChannelId != nil
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *CaseNotificationRuleRecipientData) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetChannelName returns the ChannelName field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipientData) GetChannelName() string {
	if o == nil || o.ChannelName == nil {
		var ret string
		return ret
	}
	return *o.ChannelName
}

// GetChannelNameOk returns a tuple with the ChannelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipientData) GetChannelNameOk() (*string, bool) {
	if o == nil || o.ChannelName == nil {
		return nil, false
	}
	return o.ChannelName, true
}

// HasChannelName returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipientData) HasChannelName() bool {
	return o != nil && o.ChannelName != nil
}

// SetChannelName gets a reference to the given string and assigns it to the ChannelName field.
func (o *CaseNotificationRuleRecipientData) SetChannelName(v string) {
	o.ChannelName = &v
}

// GetConnectorName returns the ConnectorName field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipientData) GetConnectorName() string {
	if o == nil || o.ConnectorName == nil {
		var ret string
		return ret
	}
	return *o.ConnectorName
}

// GetConnectorNameOk returns a tuple with the ConnectorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipientData) GetConnectorNameOk() (*string, bool) {
	if o == nil || o.ConnectorName == nil {
		return nil, false
	}
	return o.ConnectorName, true
}

// HasConnectorName returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipientData) HasConnectorName() bool {
	return o != nil && o.ConnectorName != nil
}

// SetConnectorName gets a reference to the given string and assigns it to the ConnectorName field.
func (o *CaseNotificationRuleRecipientData) SetConnectorName(v string) {
	o.ConnectorName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipientData) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipientData) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipientData) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CaseNotificationRuleRecipientData) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipientData) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipientData) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipientData) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CaseNotificationRuleRecipientData) SetName(v string) {
	o.Name = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipientData) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipientData) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipientData) HasServiceName() bool {
	return o != nil && o.ServiceName != nil
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *CaseNotificationRuleRecipientData) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetTeamId returns the TeamId field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipientData) GetTeamId() string {
	if o == nil || o.TeamId == nil {
		var ret string
		return ret
	}
	return *o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipientData) GetTeamIdOk() (*string, bool) {
	if o == nil || o.TeamId == nil {
		return nil, false
	}
	return o.TeamId, true
}

// HasTeamId returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipientData) HasTeamId() bool {
	return o != nil && o.TeamId != nil
}

// SetTeamId gets a reference to the given string and assigns it to the TeamId field.
func (o *CaseNotificationRuleRecipientData) SetTeamId(v string) {
	o.TeamId = &v
}

// GetTeamName returns the TeamName field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipientData) GetTeamName() string {
	if o == nil || o.TeamName == nil {
		var ret string
		return ret
	}
	return *o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipientData) GetTeamNameOk() (*string, bool) {
	if o == nil || o.TeamName == nil {
		return nil, false
	}
	return o.TeamName, true
}

// HasTeamName returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipientData) HasTeamName() bool {
	return o != nil && o.TeamName != nil
}

// SetTeamName gets a reference to the given string and assigns it to the TeamName field.
func (o *CaseNotificationRuleRecipientData) SetTeamName(v string) {
	o.TeamName = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipientData) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipientData) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipientData) HasTenantId() bool {
	return o != nil && o.TenantId != nil
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *CaseNotificationRuleRecipientData) SetTenantId(v string) {
	o.TenantId = &v
}

// GetTenantName returns the TenantName field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipientData) GetTenantName() string {
	if o == nil || o.TenantName == nil {
		var ret string
		return ret
	}
	return *o.TenantName
}

// GetTenantNameOk returns a tuple with the TenantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipientData) GetTenantNameOk() (*string, bool) {
	if o == nil || o.TenantName == nil {
		return nil, false
	}
	return o.TenantName, true
}

// HasTenantName returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipientData) HasTenantName() bool {
	return o != nil && o.TenantName != nil
}

// SetTenantName gets a reference to the given string and assigns it to the TenantName field.
func (o *CaseNotificationRuleRecipientData) SetTenantName(v string) {
	o.TenantName = &v
}

// GetWorkspace returns the Workspace field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipientData) GetWorkspace() string {
	if o == nil || o.Workspace == nil {
		var ret string
		return ret
	}
	return *o.Workspace
}

// GetWorkspaceOk returns a tuple with the Workspace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipientData) GetWorkspaceOk() (*string, bool) {
	if o == nil || o.Workspace == nil {
		return nil, false
	}
	return o.Workspace, true
}

// HasWorkspace returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipientData) HasWorkspace() bool {
	return o != nil && o.Workspace != nil
}

// SetWorkspace gets a reference to the given string and assigns it to the Workspace field.
func (o *CaseNotificationRuleRecipientData) SetWorkspace(v string) {
	o.Workspace = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *CaseNotificationRuleRecipientData) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaseNotificationRuleRecipientData) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *CaseNotificationRuleRecipientData) HasWorkspaceId() bool {
	return o != nil && o.WorkspaceId != nil
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *CaseNotificationRuleRecipientData) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CaseNotificationRuleRecipientData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Channel != nil {
		toSerialize["channel"] = o.Channel
	}
	if o.ChannelId != nil {
		toSerialize["channel_id"] = o.ChannelId
	}
	if o.ChannelName != nil {
		toSerialize["channel_name"] = o.ChannelName
	}
	if o.ConnectorName != nil {
		toSerialize["connector_name"] = o.ConnectorName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.TeamId != nil {
		toSerialize["team_id"] = o.TeamId
	}
	if o.TeamName != nil {
		toSerialize["team_name"] = o.TeamName
	}
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.TenantName != nil {
		toSerialize["tenant_name"] = o.TenantName
	}
	if o.Workspace != nil {
		toSerialize["workspace"] = o.Workspace
	}
	if o.WorkspaceId != nil {
		toSerialize["workspace_id"] = o.WorkspaceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CaseNotificationRuleRecipientData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Channel       *string `json:"channel,omitempty"`
		ChannelId     *string `json:"channel_id,omitempty"`
		ChannelName   *string `json:"channel_name,omitempty"`
		ConnectorName *string `json:"connector_name,omitempty"`
		Email         *string `json:"email,omitempty"`
		Name          *string `json:"name,omitempty"`
		ServiceName   *string `json:"service_name,omitempty"`
		TeamId        *string `json:"team_id,omitempty"`
		TeamName      *string `json:"team_name,omitempty"`
		TenantId      *string `json:"tenant_id,omitempty"`
		TenantName    *string `json:"tenant_name,omitempty"`
		Workspace     *string `json:"workspace,omitempty"`
		WorkspaceId   *string `json:"workspace_id,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"channel", "channel_id", "channel_name", "connector_name", "email", "name", "service_name", "team_id", "team_name", "tenant_id", "tenant_name", "workspace", "workspace_id"})
	} else {
		return err
	}
	o.Channel = all.Channel
	o.ChannelId = all.ChannelId
	o.ChannelName = all.ChannelName
	o.ConnectorName = all.ConnectorName
	o.Email = all.Email
	o.Name = all.Name
	o.ServiceName = all.ServiceName
	o.TeamId = all.TeamId
	o.TeamName = all.TeamName
	o.TenantId = all.TenantId
	o.TenantName = all.TenantName
	o.Workspace = all.Workspace
	o.WorkspaceId = all.WorkspaceId

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
