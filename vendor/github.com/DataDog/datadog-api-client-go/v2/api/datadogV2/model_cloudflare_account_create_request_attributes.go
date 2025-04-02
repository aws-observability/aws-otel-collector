// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CloudflareAccountCreateRequestAttributes Attributes object for creating a Cloudflare account.
type CloudflareAccountCreateRequestAttributes struct {
	// The API key (or token) for the Cloudflare account.
	ApiKey string `json:"api_key"`
	// The email associated with the Cloudflare account. If an API key is provided (and not a token), this field is also required.
	Email *string `json:"email,omitempty"`
	// The name of the Cloudflare account.
	Name string `json:"name"`
	// An allowlist of resources to restrict pulling metrics for including `'web', 'dns', 'lb' (load balancer), 'worker'`.
	Resources []string `json:"resources,omitempty"`
	// An allowlist of zones to restrict pulling metrics for.
	Zones []string `json:"zones,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCloudflareAccountCreateRequestAttributes instantiates a new CloudflareAccountCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCloudflareAccountCreateRequestAttributes(apiKey string, name string) *CloudflareAccountCreateRequestAttributes {
	this := CloudflareAccountCreateRequestAttributes{}
	this.ApiKey = apiKey
	this.Name = name
	return &this
}

// NewCloudflareAccountCreateRequestAttributesWithDefaults instantiates a new CloudflareAccountCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCloudflareAccountCreateRequestAttributesWithDefaults() *CloudflareAccountCreateRequestAttributes {
	this := CloudflareAccountCreateRequestAttributes{}
	return &this
}

// GetApiKey returns the ApiKey field value.
func (o *CloudflareAccountCreateRequestAttributes) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *CloudflareAccountCreateRequestAttributes) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value.
func (o *CloudflareAccountCreateRequestAttributes) SetApiKey(v string) {
	o.ApiKey = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CloudflareAccountCreateRequestAttributes) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudflareAccountCreateRequestAttributes) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CloudflareAccountCreateRequestAttributes) HasEmail() bool {
	return o != nil && o.Email != nil
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CloudflareAccountCreateRequestAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value.
func (o *CloudflareAccountCreateRequestAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CloudflareAccountCreateRequestAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CloudflareAccountCreateRequestAttributes) SetName(v string) {
	o.Name = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *CloudflareAccountCreateRequestAttributes) GetResources() []string {
	if o == nil || o.Resources == nil {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudflareAccountCreateRequestAttributes) GetResourcesOk() (*[]string, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return &o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *CloudflareAccountCreateRequestAttributes) HasResources() bool {
	return o != nil && o.Resources != nil
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *CloudflareAccountCreateRequestAttributes) SetResources(v []string) {
	o.Resources = v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *CloudflareAccountCreateRequestAttributes) GetZones() []string {
	if o == nil || o.Zones == nil {
		var ret []string
		return ret
	}
	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudflareAccountCreateRequestAttributes) GetZonesOk() (*[]string, bool) {
	if o == nil || o.Zones == nil {
		return nil, false
	}
	return &o.Zones, true
}

// HasZones returns a boolean if a field has been set.
func (o *CloudflareAccountCreateRequestAttributes) HasZones() bool {
	return o != nil && o.Zones != nil
}

// SetZones gets a reference to the given []string and assigns it to the Zones field.
func (o *CloudflareAccountCreateRequestAttributes) SetZones(v []string) {
	o.Zones = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CloudflareAccountCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["api_key"] = o.ApiKey
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
func (o *CloudflareAccountCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey    *string  `json:"api_key"`
		Email     *string  `json:"email,omitempty"`
		Name      *string  `json:"name"`
		Resources []string `json:"resources,omitempty"`
		Zones     []string `json:"zones,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiKey == nil {
		return fmt.Errorf("required field api_key missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key", "email", "name", "resources", "zones"})
	} else {
		return err
	}
	o.ApiKey = *all.ApiKey
	o.Email = all.Email
	o.Name = *all.Name
	o.Resources = all.Resources
	o.Zones = all.Zones

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
