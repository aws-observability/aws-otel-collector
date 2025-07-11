// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// InterfaceAttributes The interface attributes
type InterfaceAttributes struct {
	// The interface alias
	Alias *string `json:"alias,omitempty"`
	// The interface description
	Description *string `json:"description,omitempty"`
	// The interface index
	Index *int64 `json:"index,omitempty"`
	// The interface IP addresses
	IpAddresses []string `json:"ip_addresses,omitempty"`
	// The interface MAC address
	MacAddress *string `json:"mac_address,omitempty"`
	// The interface name
	Name *string `json:"name,omitempty"`
	// The interface status
	Status *InterfaceAttributesStatus `json:"status,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewInterfaceAttributes instantiates a new InterfaceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewInterfaceAttributes() *InterfaceAttributes {
	this := InterfaceAttributes{}
	return &this
}

// NewInterfaceAttributesWithDefaults instantiates a new InterfaceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewInterfaceAttributesWithDefaults() *InterfaceAttributes {
	this := InterfaceAttributes{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *InterfaceAttributes) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceAttributes) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *InterfaceAttributes) HasAlias() bool {
	return o != nil && o.Alias != nil
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *InterfaceAttributes) SetAlias(v string) {
	o.Alias = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *InterfaceAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *InterfaceAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *InterfaceAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *InterfaceAttributes) GetIndex() int64 {
	if o == nil || o.Index == nil {
		var ret int64
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceAttributes) GetIndexOk() (*int64, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *InterfaceAttributes) HasIndex() bool {
	return o != nil && o.Index != nil
}

// SetIndex gets a reference to the given int64 and assigns it to the Index field.
func (o *InterfaceAttributes) SetIndex(v int64) {
	o.Index = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *InterfaceAttributes) GetIpAddresses() []string {
	if o == nil || o.IpAddresses == nil {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceAttributes) GetIpAddressesOk() (*[]string, bool) {
	if o == nil || o.IpAddresses == nil {
		return nil, false
	}
	return &o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *InterfaceAttributes) HasIpAddresses() bool {
	return o != nil && o.IpAddresses != nil
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *InterfaceAttributes) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetMacAddress returns the MacAddress field value if set, zero value otherwise.
func (o *InterfaceAttributes) GetMacAddress() string {
	if o == nil || o.MacAddress == nil {
		var ret string
		return ret
	}
	return *o.MacAddress
}

// GetMacAddressOk returns a tuple with the MacAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceAttributes) GetMacAddressOk() (*string, bool) {
	if o == nil || o.MacAddress == nil {
		return nil, false
	}
	return o.MacAddress, true
}

// HasMacAddress returns a boolean if a field has been set.
func (o *InterfaceAttributes) HasMacAddress() bool {
	return o != nil && o.MacAddress != nil
}

// SetMacAddress gets a reference to the given string and assigns it to the MacAddress field.
func (o *InterfaceAttributes) SetMacAddress(v string) {
	o.MacAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InterfaceAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InterfaceAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InterfaceAttributes) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InterfaceAttributes) GetStatus() InterfaceAttributesStatus {
	if o == nil || o.Status == nil {
		var ret InterfaceAttributesStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceAttributes) GetStatusOk() (*InterfaceAttributesStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InterfaceAttributes) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given InterfaceAttributesStatus and assigns it to the Status field.
func (o *InterfaceAttributes) SetStatus(v InterfaceAttributesStatus) {
	o.Status = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o InterfaceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.IpAddresses != nil {
		toSerialize["ip_addresses"] = o.IpAddresses
	}
	if o.MacAddress != nil {
		toSerialize["mac_address"] = o.MacAddress
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *InterfaceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Alias       *string                    `json:"alias,omitempty"`
		Description *string                    `json:"description,omitempty"`
		Index       *int64                     `json:"index,omitempty"`
		IpAddresses []string                   `json:"ip_addresses,omitempty"`
		MacAddress  *string                    `json:"mac_address,omitempty"`
		Name        *string                    `json:"name,omitempty"`
		Status      *InterfaceAttributesStatus `json:"status,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"alias", "description", "index", "ip_addresses", "mac_address", "name", "status"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Alias = all.Alias
	o.Description = all.Description
	o.Index = all.Index
	o.IpAddresses = all.IpAddresses
	o.MacAddress = all.MacAddress
	o.Name = all.Name
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
