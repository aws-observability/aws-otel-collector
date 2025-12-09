// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatadogAPIKeyUpdate The definition of the `DatadogAPIKey` object.
type DatadogAPIKeyUpdate struct {
	// The `DatadogAPIKeyUpdate` `api_key`.
	ApiKey *string `json:"api_key,omitempty"`
	// The `DatadogAPIKeyUpdate` `app_key`.
	AppKey *string `json:"app_key,omitempty"`
	// The `DatadogAPIKeyUpdate` `datacenter`.
	Datacenter *string `json:"datacenter,omitempty"`
	// Custom subdomain used for Datadog URLs generated with this Connection. For example, if this org uses `https://acme.datadoghq.com` to access Datadog, set this field to `acme`. If this field is omitted, generated URLs will use the default site URL for its datacenter (see [https://docs.datadoghq.com/getting_started/site](https://docs.datadoghq.com/getting_started/site)).
	Subdomain *string `json:"subdomain,omitempty"`
	// The definition of the `DatadogAPIKey` object.
	Type DatadogAPIKeyType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatadogAPIKeyUpdate instantiates a new DatadogAPIKeyUpdate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatadogAPIKeyUpdate(typeVar DatadogAPIKeyType) *DatadogAPIKeyUpdate {
	this := DatadogAPIKeyUpdate{}
	this.Type = typeVar
	return &this
}

// NewDatadogAPIKeyUpdateWithDefaults instantiates a new DatadogAPIKeyUpdate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatadogAPIKeyUpdateWithDefaults() *DatadogAPIKeyUpdate {
	this := DatadogAPIKeyUpdate{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *DatadogAPIKeyUpdate) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogAPIKeyUpdate) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *DatadogAPIKeyUpdate) HasApiKey() bool {
	return o != nil && o.ApiKey != nil
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *DatadogAPIKeyUpdate) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetAppKey returns the AppKey field value if set, zero value otherwise.
func (o *DatadogAPIKeyUpdate) GetAppKey() string {
	if o == nil || o.AppKey == nil {
		var ret string
		return ret
	}
	return *o.AppKey
}

// GetAppKeyOk returns a tuple with the AppKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogAPIKeyUpdate) GetAppKeyOk() (*string, bool) {
	if o == nil || o.AppKey == nil {
		return nil, false
	}
	return o.AppKey, true
}

// HasAppKey returns a boolean if a field has been set.
func (o *DatadogAPIKeyUpdate) HasAppKey() bool {
	return o != nil && o.AppKey != nil
}

// SetAppKey gets a reference to the given string and assigns it to the AppKey field.
func (o *DatadogAPIKeyUpdate) SetAppKey(v string) {
	o.AppKey = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *DatadogAPIKeyUpdate) GetDatacenter() string {
	if o == nil || o.Datacenter == nil {
		var ret string
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogAPIKeyUpdate) GetDatacenterOk() (*string, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *DatadogAPIKeyUpdate) HasDatacenter() bool {
	return o != nil && o.Datacenter != nil
}

// SetDatacenter gets a reference to the given string and assigns it to the Datacenter field.
func (o *DatadogAPIKeyUpdate) SetDatacenter(v string) {
	o.Datacenter = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *DatadogAPIKeyUpdate) GetSubdomain() string {
	if o == nil || o.Subdomain == nil {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogAPIKeyUpdate) GetSubdomainOk() (*string, bool) {
	if o == nil || o.Subdomain == nil {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *DatadogAPIKeyUpdate) HasSubdomain() bool {
	return o != nil && o.Subdomain != nil
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *DatadogAPIKeyUpdate) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetType returns the Type field value.
func (o *DatadogAPIKeyUpdate) GetType() DatadogAPIKeyType {
	if o == nil {
		var ret DatadogAPIKeyType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatadogAPIKeyUpdate) GetTypeOk() (*DatadogAPIKeyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DatadogAPIKeyUpdate) SetType(v DatadogAPIKeyType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatadogAPIKeyUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.ApiKey != nil {
		toSerialize["api_key"] = o.ApiKey
	}
	if o.AppKey != nil {
		toSerialize["app_key"] = o.AppKey
	}
	if o.Datacenter != nil {
		toSerialize["datacenter"] = o.Datacenter
	}
	if o.Subdomain != nil {
		toSerialize["subdomain"] = o.Subdomain
	}
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DatadogAPIKeyUpdate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey     *string            `json:"api_key,omitempty"`
		AppKey     *string            `json:"app_key,omitempty"`
		Datacenter *string            `json:"datacenter,omitempty"`
		Subdomain  *string            `json:"subdomain,omitempty"`
		Type       *DatadogAPIKeyType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"api_key", "app_key", "datacenter", "subdomain", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.ApiKey = all.ApiKey
	o.AppKey = all.AppKey
	o.Datacenter = all.Datacenter
	o.Subdomain = all.Subdomain
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
