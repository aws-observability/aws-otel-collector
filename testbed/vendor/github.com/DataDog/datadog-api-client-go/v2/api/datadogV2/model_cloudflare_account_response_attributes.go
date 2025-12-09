// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudflareAccountResponseAttributes Attributes object of a Cloudflare account.
type CloudflareAccountResponseAttributes struct {
	// The email associated with the Cloudflare account.
	Email *string `json:"email,omitempty"`
	// The name of the Cloudflare account.
	Name string `json:"name"`
	// An allowlist of resources, such as `web`, `dns`, `lb` (load balancer), `worker`, that restricts pulling metrics from those resources.
	Resources []string `json:"resources,omitempty"`
	// An allowlist of zones to restrict pulling metrics for.
	Zones []string `json:"zones,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudflareAccountResponseAttributes instantiates a new CloudflareAccountResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudflareAccountResponseAttributes(name string) *CloudflareAccountResponseAttributes {
	this := CloudflareAccountResponseAttributes{}
	this.Name = name
	return &this
}

// NewCloudflareAccountResponseAttributesWithDefaults instantiates a new CloudflareAccountResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudflareAccountResponseAttributesWithDefaults() *CloudflareAccountResponseAttributes {
	this := CloudflareAccountResponseAttributes{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CloudflareAccountResponseAttributes) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudflareAccountResponseAttributes) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CloudflareAccountResponseAttributes) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CloudflareAccountResponseAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value.
func (o *CloudflareAccountResponseAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudflareAccountResponseAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CloudflareAccountResponseAttributes) SetName(v string) {
	o.Name = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *CloudflareAccountResponseAttributes) GetResources() []string {
	if o == nil || o.Resources == nil {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudflareAccountResponseAttributes) GetResourcesOk() (*[]string, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return &o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *CloudflareAccountResponseAttributes) HasResources() bool {
	return o != nil && o.Resources != nil
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *CloudflareAccountResponseAttributes) SetResources(v []string) {
	o.Resources = v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *CloudflareAccountResponseAttributes) GetZones() []string {
	if o == nil || o.Zones == nil {
		var ret []string
		return ret
	}
	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudflareAccountResponseAttributes) GetZonesOk() (*[]string, bool) {
	if o == nil || o.Zones == nil {
		return nil, false
	}
	return &o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *CloudflareAccountResponseAttributes) HasZones() bool {
	return o != nil && o.Zones != nil
}

// SetZones gets a reference to the given []string and assigns it to the Zones field.
func (o *CloudflareAccountResponseAttributes) SetZones(v []string) {
	o.Zones = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudflareAccountResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	toSerialize["name"] = o.Name
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.Zones != nil {
		toSerialize["zones"] = o.Zones
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CloudflareAccountResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Email     *string  `json:"email,omitempty"`
		Name      *string  `json:"name"`
		Resources []string `json:"resources,omitempty"`
		Zones     []string `json:"zones,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"email", "name", "resources", "zones"})
	} else {
		return err
	}
	o.Email = all.Email
	o.Name = *all.Name
	o.Resources = all.Resources
	o.Zones = all.Zones

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
