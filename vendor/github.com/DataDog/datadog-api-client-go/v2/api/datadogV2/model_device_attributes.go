// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DeviceAttributes The device attributes
type DeviceAttributes struct {
	// The device description
	Description *string `json:"description,omitempty"`
	// The device type
	DeviceType *string `json:"device_type,omitempty"`
	// The device integration
	Integration *string `json:"integration,omitempty"`
	// Count of the device interfaces by status
	InterfaceStatuses *DeviceAttributesInterfaceStatuses `json:"interface_statuses,omitempty"`
	// The device IP address
	IpAddress *string `json:"ip_address,omitempty"`
	// The device location
	Location *string `json:"location,omitempty"`
	// The device model
	Model *string `json:"model,omitempty"`
	// The device name
	Name *string `json:"name,omitempty"`
	// The device OS hostname
	OsHostname *string `json:"os_hostname,omitempty"`
	// The device OS name
	OsName *string `json:"os_name,omitempty"`
	// The device OS version
	OsVersion *string `json:"os_version,omitempty"`
	// The device ping status
	PingStatus *string `json:"ping_status,omitempty"`
	// The device product name
	ProductName *string `json:"product_name,omitempty"`
	// The device serial number
	SerialNumber *string `json:"serial_number,omitempty"`
	// The device SNMP status
	Status *string `json:"status,omitempty"`
	// The device subnet
	Subnet *string `json:"subnet,omitempty"`
	// The device `sys_object_id`
	SysObjectId *string `json:"sys_object_id,omitempty"`
	// The list of device tags
	Tags []string `json:"tags,omitempty"`
	// The device vendor
	Vendor *string `json:"vendor,omitempty"`
	// The device version
	Version *string `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDeviceAttributes instantiates a new DeviceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDeviceAttributes() *DeviceAttributes {
	this := DeviceAttributes{}
	return &this
}

// NewDeviceAttributesWithDefaults instantiates a new DeviceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDeviceAttributesWithDefaults() *DeviceAttributes {
	this := DeviceAttributes{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DeviceAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DeviceAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DeviceAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise.
func (o *DeviceAttributes) GetDeviceType() string {
	if o == nil || o.DeviceType == nil {
		var ret string
		return ret
	}
	return *o.DeviceType
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetDeviceTypeOk() (*string, bool) {
	if o == nil || o.DeviceType == nil {
		return nil, false
	}
	return o.DeviceType, true
}

// HasDeviceType returns a boolean if a field has been set.
func (o *DeviceAttributes) HasDeviceType() bool {
	return o != nil && o.DeviceType != nil
}

// SetDeviceType gets a reference to the given string and assigns it to the DeviceType field.
func (o *DeviceAttributes) SetDeviceType(v string) {
	o.DeviceType = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *DeviceAttributes) GetIntegration() string {
	if o == nil || o.Integration == nil {
		var ret string
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetIntegrationOk() (*string, bool) {
	if o == nil || o.Integration == nil {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *DeviceAttributes) HasIntegration() bool {
	return o != nil && o.Integration != nil
}

// SetIntegration gets a reference to the given string and assigns it to the Integration field.
func (o *DeviceAttributes) SetIntegration(v string) {
	o.Integration = &v
}

// GetInterfaceStatuses returns the InterfaceStatuses field value if set, zero value otherwise.
func (o *DeviceAttributes) GetInterfaceStatuses() DeviceAttributesInterfaceStatuses {
	if o == nil || o.InterfaceStatuses == nil {
		var ret DeviceAttributesInterfaceStatuses
		return ret
	}
	return *o.InterfaceStatuses
}

// GetInterfaceStatusesOk returns a tuple with the InterfaceStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetInterfaceStatusesOk() (*DeviceAttributesInterfaceStatuses, bool) {
	if o == nil || o.InterfaceStatuses == nil {
		return nil, false
	}
	return o.InterfaceStatuses, true
}

// HasInterfaceStatuses returns a boolean if a field has been set.
func (o *DeviceAttributes) HasInterfaceStatuses() bool {
	return o != nil && o.InterfaceStatuses != nil
}

// SetInterfaceStatuses gets a reference to the given DeviceAttributesInterfaceStatuses and assigns it to the InterfaceStatuses field.
func (o *DeviceAttributes) SetInterfaceStatuses(v DeviceAttributesInterfaceStatuses) {
	o.InterfaceStatuses = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *DeviceAttributes) GetIpAddress() string {
	if o == nil || o.IpAddress == nil {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetIpAddressOk() (*string, bool) {
	if o == nil || o.IpAddress == nil {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *DeviceAttributes) HasIpAddress() bool {
	return o != nil && o.IpAddress != nil
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *DeviceAttributes) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *DeviceAttributes) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *DeviceAttributes) HasLocation() bool {
	return o != nil && o.Location != nil
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *DeviceAttributes) SetLocation(v string) {
	o.Location = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *DeviceAttributes) GetModel() string {
	if o == nil || o.Model == nil {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetModelOk() (*string, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *DeviceAttributes) HasModel() bool {
	return o != nil && o.Model != nil
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *DeviceAttributes) SetModel(v string) {
	o.Model = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DeviceAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DeviceAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DeviceAttributes) SetName(v string) {
	o.Name = &v
}

// GetOsHostname returns the OsHostname field value if set, zero value otherwise.
func (o *DeviceAttributes) GetOsHostname() string {
	if o == nil || o.OsHostname == nil {
		var ret string
		return ret
	}
	return *o.OsHostname
}

// GetOsHostnameOk returns a tuple with the OsHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetOsHostnameOk() (*string, bool) {
	if o == nil || o.OsHostname == nil {
		return nil, false
	}
	return o.OsHostname, true
}

// HasOsHostname returns a boolean if a field has been set.
func (o *DeviceAttributes) HasOsHostname() bool {
	return o != nil && o.OsHostname != nil
}

// SetOsHostname gets a reference to the given string and assigns it to the OsHostname field.
func (o *DeviceAttributes) SetOsHostname(v string) {
	o.OsHostname = &v
}

// GetOsName returns the OsName field value if set, zero value otherwise.
func (o *DeviceAttributes) GetOsName() string {
	if o == nil || o.OsName == nil {
		var ret string
		return ret
	}
	return *o.OsName
}

// GetOsNameOk returns a tuple with the OsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetOsNameOk() (*string, bool) {
	if o == nil || o.OsName == nil {
		return nil, false
	}
	return o.OsName, true
}

// HasOsName returns a boolean if a field has been set.
func (o *DeviceAttributes) HasOsName() bool {
	return o != nil && o.OsName != nil
}

// SetOsName gets a reference to the given string and assigns it to the OsName field.
func (o *DeviceAttributes) SetOsName(v string) {
	o.OsName = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *DeviceAttributes) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *DeviceAttributes) HasOsVersion() bool {
	return o != nil && o.OsVersion != nil
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *DeviceAttributes) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetPingStatus returns the PingStatus field value if set, zero value otherwise.
func (o *DeviceAttributes) GetPingStatus() string {
	if o == nil || o.PingStatus == nil {
		var ret string
		return ret
	}
	return *o.PingStatus
}

// GetPingStatusOk returns a tuple with the PingStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetPingStatusOk() (*string, bool) {
	if o == nil || o.PingStatus == nil {
		return nil, false
	}
	return o.PingStatus, true
}

// HasPingStatus returns a boolean if a field has been set.
func (o *DeviceAttributes) HasPingStatus() bool {
	return o != nil && o.PingStatus != nil
}

// SetPingStatus gets a reference to the given string and assigns it to the PingStatus field.
func (o *DeviceAttributes) SetPingStatus(v string) {
	o.PingStatus = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *DeviceAttributes) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *DeviceAttributes) HasProductName() bool {
	return o != nil && o.ProductName != nil
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *DeviceAttributes) SetProductName(v string) {
	o.ProductName = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *DeviceAttributes) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *DeviceAttributes) HasSerialNumber() bool {
	return o != nil && o.SerialNumber != nil
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *DeviceAttributes) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DeviceAttributes) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DeviceAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DeviceAttributes) SetStatus(v string) {
	o.Status = &v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *DeviceAttributes) GetSubnet() string {
	if o == nil || o.Subnet == nil {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetSubnetOk() (*string, bool) {
	if o == nil || o.Subnet == nil {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *DeviceAttributes) HasSubnet() bool {
	return o != nil && o.Subnet != nil
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *DeviceAttributes) SetSubnet(v string) {
	o.Subnet = &v
}

// GetSysObjectId returns the SysObjectId field value if set, zero value otherwise.
func (o *DeviceAttributes) GetSysObjectId() string {
	if o == nil || o.SysObjectId == nil {
		var ret string
		return ret
	}
	return *o.SysObjectId
}

// GetSysObjectIdOk returns a tuple with the SysObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetSysObjectIdOk() (*string, bool) {
	if o == nil || o.SysObjectId == nil {
		return nil, false
	}
	return o.SysObjectId, true
}

// HasSysObjectId returns a boolean if a field has been set.
func (o *DeviceAttributes) HasSysObjectId() bool {
	return o != nil && o.SysObjectId != nil
}

// SetSysObjectId gets a reference to the given string and assigns it to the SysObjectId field.
func (o *DeviceAttributes) SetSysObjectId(v string) {
	o.SysObjectId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DeviceAttributes) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *DeviceAttributes) SetTags(v []string) {
	o.Tags = v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *DeviceAttributes) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *DeviceAttributes) HasVendor() bool {
	return o != nil && o.Vendor != nil
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *DeviceAttributes) SetVendor(v string) {
	o.Vendor = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DeviceAttributes) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAttributes) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DeviceAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *DeviceAttributes) SetVersion(v string) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DeviceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DeviceType != nil {
		toSerialize["device_type"] = o.DeviceType
	}
	if o.Integration != nil {
		toSerialize["integration"] = o.Integration
	}
	if o.InterfaceStatuses != nil {
		toSerialize["interface_statuses"] = o.InterfaceStatuses
	}
	if o.IpAddress != nil {
		toSerialize["ip_address"] = o.IpAddress
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OsHostname != nil {
		toSerialize["os_hostname"] = o.OsHostname
	}
	if o.OsName != nil {
		toSerialize["os_name"] = o.OsName
	}
	if o.OsVersion != nil {
		toSerialize["os_version"] = o.OsVersion
	}
	if o.PingStatus != nil {
		toSerialize["ping_status"] = o.PingStatus
	}
	if o.ProductName != nil {
		toSerialize["product_name"] = o.ProductName
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Subnet != nil {
		toSerialize["subnet"] = o.Subnet
	}
	if o.SysObjectId != nil {
		toSerialize["sys_object_id"] = o.SysObjectId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DeviceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Description       *string                            `json:"description,omitempty"`
		DeviceType        *string                            `json:"device_type,omitempty"`
		Integration       *string                            `json:"integration,omitempty"`
		InterfaceStatuses *DeviceAttributesInterfaceStatuses `json:"interface_statuses,omitempty"`
		IpAddress         *string                            `json:"ip_address,omitempty"`
		Location          *string                            `json:"location,omitempty"`
		Model             *string                            `json:"model,omitempty"`
		Name              *string                            `json:"name,omitempty"`
		OsHostname        *string                            `json:"os_hostname,omitempty"`
		OsName            *string                            `json:"os_name,omitempty"`
		OsVersion         *string                            `json:"os_version,omitempty"`
		PingStatus        *string                            `json:"ping_status,omitempty"`
		ProductName       *string                            `json:"product_name,omitempty"`
		SerialNumber      *string                            `json:"serial_number,omitempty"`
		Status            *string                            `json:"status,omitempty"`
		Subnet            *string                            `json:"subnet,omitempty"`
		SysObjectId       *string                            `json:"sys_object_id,omitempty"`
		Tags              []string                           `json:"tags,omitempty"`
		Vendor            *string                            `json:"vendor,omitempty"`
		Version           *string                            `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"description", "device_type", "integration", "interface_statuses", "ip_address", "location", "model", "name", "os_hostname", "os_name", "os_version", "ping_status", "product_name", "serial_number", "status", "subnet", "sys_object_id", "tags", "vendor", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Description = all.Description
	o.DeviceType = all.DeviceType
	o.Integration = all.Integration
	if all.InterfaceStatuses != nil && all.InterfaceStatuses.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.InterfaceStatuses = all.InterfaceStatuses
	o.IpAddress = all.IpAddress
	o.Location = all.Location
	o.Model = all.Model
	o.Name = all.Name
	o.OsHostname = all.OsHostname
	o.OsName = all.OsName
	o.OsVersion = all.OsVersion
	o.PingStatus = all.PingStatus
	o.ProductName = all.ProductName
	o.SerialNumber = all.SerialNumber
	o.Status = all.Status
	o.Subnet = all.Subnet
	o.SysObjectId = all.SysObjectId
	o.Tags = all.Tags
	o.Vendor = all.Vendor
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
