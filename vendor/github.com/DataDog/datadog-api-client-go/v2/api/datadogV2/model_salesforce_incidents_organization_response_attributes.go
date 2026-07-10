// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SalesforceIncidentsOrganizationResponseAttributes Attributes of a Salesforce organization connected to the Datadog Salesforce integration.
type SalesforceIncidentsOrganizationResponseAttributes struct {
	// The Salesforce instance URL used to call this organization's APIs.
	InstanceUrl *string `json:"instance_url,omitempty"`
	// Human-readable name of the Salesforce organization.
	Name *string `json:"name,omitempty"`
	// The Salesforce organization identifier (15- or 18-character Salesforce org ID).
	SfdcOrgId *string `json:"sfdc_org_id,omitempty"`
	// The Salesforce organization type (for example, `Production` or `Sandbox`).
	SfdcOrgType *string `json:"sfdc_org_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSalesforceIncidentsOrganizationResponseAttributes instantiates a new SalesforceIncidentsOrganizationResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSalesforceIncidentsOrganizationResponseAttributes() *SalesforceIncidentsOrganizationResponseAttributes {
	this := SalesforceIncidentsOrganizationResponseAttributes{}
	return &this
}

// NewSalesforceIncidentsOrganizationResponseAttributesWithDefaults instantiates a new SalesforceIncidentsOrganizationResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSalesforceIncidentsOrganizationResponseAttributesWithDefaults() *SalesforceIncidentsOrganizationResponseAttributes {
	this := SalesforceIncidentsOrganizationResponseAttributes{}
	return &this
}

// GetInstanceUrl returns the InstanceUrl field value if set, zero value otherwise.
func (o *SalesforceIncidentsOrganizationResponseAttributes) GetInstanceUrl() string {
	if o == nil || o.InstanceUrl == nil {
		var ret string
		return ret
	}
	return *o.InstanceUrl
}

// GetInstanceUrlOk returns a tuple with the InstanceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsOrganizationResponseAttributes) GetInstanceUrlOk() (*string, bool) {
	if o == nil || o.InstanceUrl == nil {
		return nil, false
	}
	return o.InstanceUrl, true
}

// HasInstanceUrl returns a boolean if a field has been set.
func (o *SalesforceIncidentsOrganizationResponseAttributes) HasInstanceUrl() bool {
	return o != nil && o.InstanceUrl != nil
}

// SetInstanceUrl gets a reference to the given string and assigns it to the InstanceUrl field.
func (o *SalesforceIncidentsOrganizationResponseAttributes) SetInstanceUrl(v string) {
	o.InstanceUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SalesforceIncidentsOrganizationResponseAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsOrganizationResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SalesforceIncidentsOrganizationResponseAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SalesforceIncidentsOrganizationResponseAttributes) SetName(v string) {
	o.Name = &v
}

// GetSfdcOrgId returns the SfdcOrgId field value if set, zero value otherwise.
func (o *SalesforceIncidentsOrganizationResponseAttributes) GetSfdcOrgId() string {
	if o == nil || o.SfdcOrgId == nil {
		var ret string
		return ret
	}
	return *o.SfdcOrgId
}

// GetSfdcOrgIdOk returns a tuple with the SfdcOrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsOrganizationResponseAttributes) GetSfdcOrgIdOk() (*string, bool) {
	if o == nil || o.SfdcOrgId == nil {
		return nil, false
	}
	return o.SfdcOrgId, true
}

// HasSfdcOrgId returns a boolean if a field has been set.
func (o *SalesforceIncidentsOrganizationResponseAttributes) HasSfdcOrgId() bool {
	return o != nil && o.SfdcOrgId != nil
}

// SetSfdcOrgId gets a reference to the given string and assigns it to the SfdcOrgId field.
func (o *SalesforceIncidentsOrganizationResponseAttributes) SetSfdcOrgId(v string) {
	o.SfdcOrgId = &v
}

// GetSfdcOrgType returns the SfdcOrgType field value if set, zero value otherwise.
func (o *SalesforceIncidentsOrganizationResponseAttributes) GetSfdcOrgType() string {
	if o == nil || o.SfdcOrgType == nil {
		var ret string
		return ret
	}
	return *o.SfdcOrgType
}

// GetSfdcOrgTypeOk returns a tuple with the SfdcOrgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SalesforceIncidentsOrganizationResponseAttributes) GetSfdcOrgTypeOk() (*string, bool) {
	if o == nil || o.SfdcOrgType == nil {
		return nil, false
	}
	return o.SfdcOrgType, true
}

// HasSfdcOrgType returns a boolean if a field has been set.
func (o *SalesforceIncidentsOrganizationResponseAttributes) HasSfdcOrgType() bool {
	return o != nil && o.SfdcOrgType != nil
}

// SetSfdcOrgType gets a reference to the given string and assigns it to the SfdcOrgType field.
func (o *SalesforceIncidentsOrganizationResponseAttributes) SetSfdcOrgType(v string) {
	o.SfdcOrgType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SalesforceIncidentsOrganizationResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.InstanceUrl != nil {
		toSerialize["instance_url"] = o.InstanceUrl
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SfdcOrgId != nil {
		toSerialize["sfdc_org_id"] = o.SfdcOrgId
	}
	if o.SfdcOrgType != nil {
		toSerialize["sfdc_org_type"] = o.SfdcOrgType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SalesforceIncidentsOrganizationResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		InstanceUrl *string `json:"instance_url,omitempty"`
		Name        *string `json:"name,omitempty"`
		SfdcOrgId   *string `json:"sfdc_org_id,omitempty"`
		SfdcOrgType *string `json:"sfdc_org_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"instance_url", "name", "sfdc_org_id", "sfdc_org_type"})
	} else {
		return err
	}
	o.InstanceUrl = all.InstanceUrl
	o.Name = all.Name
	o.SfdcOrgId = all.SfdcOrgId
	o.SfdcOrgType = all.SfdcOrgType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
