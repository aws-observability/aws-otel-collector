// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DatadogAPIKey The definition of the `DatadogAPIKey` object.
type DatadogAPIKey struct {
	// The `DatadogAPIKey` `api_key`.
	ApiKey string `json:"api_key"`
	// The `DatadogAPIKey` `app_key`.
	AppKey string `json:"app_key"`
	// The `DatadogAPIKey` `datacenter`.
	Datacenter string `json:"datacenter"`
	// Custom subdomain used for Datadog URLs generated with this Connection. For example, if this org uses `https://acme.datadoghq.com` to access Datadog, set this field to `acme`. If this field is omitted, generated URLs will use the default site URL for its datacenter (see [https://docs.datadoghq.com/getting_started/site](https://docs.datadoghq.com/getting_started/site)).
	Subdomain *string `json:"subdomain,omitempty"`
	// The definition of the `DatadogAPIKey` object.
	Type DatadogAPIKeyType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDatadogAPIKey instantiates a new DatadogAPIKey object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDatadogAPIKey(apiKey string, appKey string, datacenter string, typeVar DatadogAPIKeyType) *DatadogAPIKey {
	this := DatadogAPIKey{}
	this.ApiKey = apiKey
	this.AppKey = appKey
	this.Datacenter = datacenter
	this.Type = typeVar
	return &this
}

// NewDatadogAPIKeyWithDefaults instantiates a new DatadogAPIKey object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDatadogAPIKeyWithDefaults() *DatadogAPIKey {
	this := DatadogAPIKey{}
	return &this
}

// GetApiKey returns the ApiKey field value.
func (o *DatadogAPIKey) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *DatadogAPIKey) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value.
func (o *DatadogAPIKey) SetApiKey(v string) {
	o.ApiKey = v
}

// GetAppKey returns the AppKey field value.
func (o *DatadogAPIKey) GetAppKey() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AppKey
}

// GetAppKeyOk returns a tuple with the AppKey field value
// and a boolean to check if the value has been set.
func (o *DatadogAPIKey) GetAppKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppKey, true
}

// SetAppKey sets field value.
func (o *DatadogAPIKey) SetAppKey(v string) {
	o.AppKey = v
}

// GetDatacenter returns the Datacenter field value.
func (o *DatadogAPIKey) GetDatacenter() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value
// and a boolean to check if the value has been set.
func (o *DatadogAPIKey) GetDatacenterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datacenter, true
}

// SetDatacenter sets field value.
func (o *DatadogAPIKey) SetDatacenter(v string) {
	o.Datacenter = v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *DatadogAPIKey) GetSubdomain() string {
	if o == nil || o.Subdomain == nil {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatadogAPIKey) GetSubdomainOk() (*string, bool) {
	if o == nil || o.Subdomain == nil {
		return nil, false
	}
	return o.Subdomain, true
}

// HasSubdomain returns a boolean if a field has been set.
func (o *DatadogAPIKey) HasSubdomain() bool {
	return o != nil && o.Subdomain != nil
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *DatadogAPIKey) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetType returns the Type field value.
func (o *DatadogAPIKey) GetType() DatadogAPIKeyType {
	if o == nil {
		var ret DatadogAPIKeyType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatadogAPIKey) GetTypeOk() (*DatadogAPIKeyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *DatadogAPIKey) SetType(v DatadogAPIKeyType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DatadogAPIKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["api_key"] = o.ApiKey
	toSerialize["app_key"] = o.AppKey
	toSerialize["datacenter"] = o.Datacenter
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
func (o *DatadogAPIKey) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ApiKey     *string            `json:"api_key"`
		AppKey     *string            `json:"app_key"`
		Datacenter *string            `json:"datacenter"`
		Subdomain  *string            `json:"subdomain,omitempty"`
		Type       *DatadogAPIKeyType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ApiKey == nil {
		return fmt.Errorf("required field api_key missing")
	}
	if all.AppKey == nil {
		return fmt.Errorf("required field app_key missing")
	}
	if all.Datacenter == nil {
		return fmt.Errorf("required field datacenter missing")
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
	o.ApiKey = *all.ApiKey
	o.AppKey = *all.AppKey
	o.Datacenter = *all.Datacenter
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
