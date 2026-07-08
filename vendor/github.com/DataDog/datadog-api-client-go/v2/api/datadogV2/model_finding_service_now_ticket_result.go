// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// FindingServiceNowTicketResult Result of the ServiceNow ticket creation or attachment.
type FindingServiceNowTicketResult struct {
	// ServiceNow instance name extracted from the ticket URL.
	InstanceName *string `json:"instance_name,omitempty"`
	// Unique identifier of the ServiceNow incident record.
	SysId *string `json:"sys_id,omitempty"`
	// Direct link to the ServiceNow incident record.
	SysTargetLink *string `json:"sys_target_link,omitempty"`
	// Unique identifier of the target ServiceNow record.
	SysTargetSysId *string `json:"sys_target_sys_id,omitempty"`
	// ServiceNow table containing the incident record.
	TableName *string `json:"table_name,omitempty"`
	// URL of the ServiceNow incident record.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewFindingServiceNowTicketResult instantiates a new FindingServiceNowTicketResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewFindingServiceNowTicketResult() *FindingServiceNowTicketResult {
	this := FindingServiceNowTicketResult{}
	return &this
}

// NewFindingServiceNowTicketResultWithDefaults instantiates a new FindingServiceNowTicketResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewFindingServiceNowTicketResultWithDefaults() *FindingServiceNowTicketResult {
	this := FindingServiceNowTicketResult{}
	return &this
}

// GetInstanceName returns the InstanceName field value if set, zero value otherwise.
func (o *FindingServiceNowTicketResult) GetInstanceName() string {
	if o == nil || o.InstanceName == nil {
		var ret string
		return ret
	}
	return *o.InstanceName
}

// GetInstanceNameOk returns a tuple with the InstanceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingServiceNowTicketResult) GetInstanceNameOk() (*string, bool) {
	if o == nil || o.InstanceName == nil {
		return nil, false
	}
	return o.InstanceName, true
}

// HasInstanceName returns a boolean if a field has been set.
func (o *FindingServiceNowTicketResult) HasInstanceName() bool {
	return o != nil && o.InstanceName != nil
}

// SetInstanceName gets a reference to the given string and assigns it to the InstanceName field.
func (o *FindingServiceNowTicketResult) SetInstanceName(v string) {
	o.InstanceName = &v
}

// GetSysId returns the SysId field value if set, zero value otherwise.
func (o *FindingServiceNowTicketResult) GetSysId() string {
	if o == nil || o.SysId == nil {
		var ret string
		return ret
	}
	return *o.SysId
}

// GetSysIdOk returns a tuple with the SysId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingServiceNowTicketResult) GetSysIdOk() (*string, bool) {
	if o == nil || o.SysId == nil {
		return nil, false
	}
	return o.SysId, true
}

// HasSysId returns a boolean if a field has been set.
func (o *FindingServiceNowTicketResult) HasSysId() bool {
	return o != nil && o.SysId != nil
}

// SetSysId gets a reference to the given string and assigns it to the SysId field.
func (o *FindingServiceNowTicketResult) SetSysId(v string) {
	o.SysId = &v
}

// GetSysTargetLink returns the SysTargetLink field value if set, zero value otherwise.
func (o *FindingServiceNowTicketResult) GetSysTargetLink() string {
	if o == nil || o.SysTargetLink == nil {
		var ret string
		return ret
	}
	return *o.SysTargetLink
}

// GetSysTargetLinkOk returns a tuple with the SysTargetLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingServiceNowTicketResult) GetSysTargetLinkOk() (*string, bool) {
	if o == nil || o.SysTargetLink == nil {
		return nil, false
	}
	return o.SysTargetLink, true
}

// HasSysTargetLink returns a boolean if a field has been set.
func (o *FindingServiceNowTicketResult) HasSysTargetLink() bool {
	return o != nil && o.SysTargetLink != nil
}

// SetSysTargetLink gets a reference to the given string and assigns it to the SysTargetLink field.
func (o *FindingServiceNowTicketResult) SetSysTargetLink(v string) {
	o.SysTargetLink = &v
}

// GetSysTargetSysId returns the SysTargetSysId field value if set, zero value otherwise.
func (o *FindingServiceNowTicketResult) GetSysTargetSysId() string {
	if o == nil || o.SysTargetSysId == nil {
		var ret string
		return ret
	}
	return *o.SysTargetSysId
}

// GetSysTargetSysIdOk returns a tuple with the SysTargetSysId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingServiceNowTicketResult) GetSysTargetSysIdOk() (*string, bool) {
	if o == nil || o.SysTargetSysId == nil {
		return nil, false
	}
	return o.SysTargetSysId, true
}

// HasSysTargetSysId returns a boolean if a field has been set.
func (o *FindingServiceNowTicketResult) HasSysTargetSysId() bool {
	return o != nil && o.SysTargetSysId != nil
}

// SetSysTargetSysId gets a reference to the given string and assigns it to the SysTargetSysId field.
func (o *FindingServiceNowTicketResult) SetSysTargetSysId(v string) {
	o.SysTargetSysId = &v
}

// GetTableName returns the TableName field value if set, zero value otherwise.
func (o *FindingServiceNowTicketResult) GetTableName() string {
	if o == nil || o.TableName == nil {
		var ret string
		return ret
	}
	return *o.TableName
}

// GetTableNameOk returns a tuple with the TableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingServiceNowTicketResult) GetTableNameOk() (*string, bool) {
	if o == nil || o.TableName == nil {
		return nil, false
	}
	return o.TableName, true
}

// HasTableName returns a boolean if a field has been set.
func (o *FindingServiceNowTicketResult) HasTableName() bool {
	return o != nil && o.TableName != nil
}

// SetTableName gets a reference to the given string and assigns it to the TableName field.
func (o *FindingServiceNowTicketResult) SetTableName(v string) {
	o.TableName = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *FindingServiceNowTicketResult) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindingServiceNowTicketResult) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *FindingServiceNowTicketResult) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *FindingServiceNowTicketResult) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o FindingServiceNowTicketResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.InstanceName != nil {
		toSerialize["instance_name"] = o.InstanceName
	}
	if o.SysId != nil {
		toSerialize["sys_id"] = o.SysId
	}
	if o.SysTargetLink != nil {
		toSerialize["sys_target_link"] = o.SysTargetLink
	}
	if o.SysTargetSysId != nil {
		toSerialize["sys_target_sys_id"] = o.SysTargetSysId
	}
	if o.TableName != nil {
		toSerialize["table_name"] = o.TableName
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *FindingServiceNowTicketResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceName   *string `json:"instance_name,omitempty"`
		SysId          *string `json:"sys_id,omitempty"`
		SysTargetLink  *string `json:"sys_target_link,omitempty"`
		SysTargetSysId *string `json:"sys_target_sys_id,omitempty"`
		TableName      *string `json:"table_name,omitempty"`
		Url            *string `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"instance_name", "sys_id", "sys_target_link", "sys_target_sys_id", "table_name", "url"})
	} else {
		return err
	}
	o.InstanceName = all.InstanceName
	o.SysId = all.SysId
	o.SysTargetLink = all.SysTargetLink
	o.SysTargetSysId = all.SysTargetSysId
	o.TableName = all.TableName
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
