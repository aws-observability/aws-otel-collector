// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AzureTenant The definition of the `AzureTenant` object.
type AzureTenant struct {
	// The Client ID, also known as the Application ID in Azure, is a unique identifier for an application. It's used to identify the application during the authentication process. Your Application (client) ID is listed in the application's overview page. You can navigate to your application via the Azure Directory.
	AppClientId string `json:"app_client_id"`
	// The Client Secret is a confidential piece of information known only to the application and Azure AD. It's used to prove the application's identity. Your Client Secret is available from the application’s secrets page. You can navigate to your application via the Azure Directory.
	ClientSecret string `json:"client_secret"`
	// If provided, the custom scope to be requested from Microsoft when acquiring an OAuth 2 access token. This custom scope is used only in conjunction with the HTTP action. A resource's scope is constructed by using the identifier URI for the resource and .default, separated by a forward slash (/) as follows:{identifierURI}/.default.
	CustomScopes *string `json:"custom_scopes,omitempty"`
	// The Tenant ID, also known as the Directory ID in Azure, is a unique identifier that represents an Azure AD instance. Your Tenant ID (Directory ID) is listed in your Active Directory overview page under the 'Tenant information' section.
	TenantId string `json:"tenant_id"`
	// The definition of the `AzureTenant` object.
	Type AzureTenantType `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAzureTenant instantiates a new AzureTenant object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAzureTenant(appClientId string, clientSecret string, tenantId string, typeVar AzureTenantType) *AzureTenant {
	this := AzureTenant{}
	this.AppClientId = appClientId
	this.ClientSecret = clientSecret
	this.TenantId = tenantId
	this.Type = typeVar
	return &this
}

// NewAzureTenantWithDefaults instantiates a new AzureTenant object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAzureTenantWithDefaults() *AzureTenant {
	this := AzureTenant{}
	return &this
}

// GetAppClientId returns the AppClientId field value.
func (o *AzureTenant) GetAppClientId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AppClientId
}

// GetAppClientIdOk returns a tuple with the AppClientId field value
// and a boolean to check if the value has been set.
func (o *AzureTenant) GetAppClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppClientId, true
}

// SetAppClientId sets field value.
func (o *AzureTenant) SetAppClientId(v string) {
	o.AppClientId = v
}

// GetClientSecret returns the ClientSecret field value.
func (o *AzureTenant) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *AzureTenant) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value.
func (o *AzureTenant) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetCustomScopes returns the CustomScopes field value if set, zero value otherwise.
func (o *AzureTenant) GetCustomScopes() string {
	if o == nil || o.CustomScopes == nil {
		var ret string
		return ret
	}
	return *o.CustomScopes
}

// GetCustomScopesOk returns a tuple with the CustomScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureTenant) GetCustomScopesOk() (*string, bool) {
	if o == nil || o.CustomScopes == nil {
		return nil, false
	}
	return o.CustomScopes, true
}

// HasCustomScopes returns a boolean if a field has been set.
func (o *AzureTenant) HasCustomScopes() bool {
	return o != nil && o.CustomScopes != nil
}

// SetCustomScopes gets a reference to the given string and assigns it to the CustomScopes field.
func (o *AzureTenant) SetCustomScopes(v string) {
	o.CustomScopes = &v
}

// GetTenantId returns the TenantId field value.
func (o *AzureTenant) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *AzureTenant) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value.
func (o *AzureTenant) SetTenantId(v string) {
	o.TenantId = v
}

// GetType returns the Type field value.
func (o *AzureTenant) GetType() AzureTenantType {
	if o == nil {
		var ret AzureTenantType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AzureTenant) GetTypeOk() (*AzureTenantType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *AzureTenant) SetType(v AzureTenantType) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AzureTenant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["app_client_id"] = o.AppClientId
	toSerialize["client_secret"] = o.ClientSecret
	if o.CustomScopes != nil {
		toSerialize["custom_scopes"] = o.CustomScopes
	}
	toSerialize["tenant_id"] = o.TenantId
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AzureTenant) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AppClientId  *string          `json:"app_client_id"`
		ClientSecret *string          `json:"client_secret"`
		CustomScopes *string          `json:"custom_scopes,omitempty"`
		TenantId     *string          `json:"tenant_id"`
		Type         *AzureTenantType `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AppClientId == nil {
		return fmt.Errorf("required field app_client_id missing")
	}
	if all.ClientSecret == nil {
		return fmt.Errorf("required field client_secret missing")
	}
	if all.TenantId == nil {
		return fmt.Errorf("required field tenant_id missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"app_client_id", "client_secret", "custom_scopes", "tenant_id", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AppClientId = *all.AppClientId
	o.ClientSecret = *all.ClientSecret
	o.CustomScopes = all.CustomScopes
	o.TenantId = *all.TenantId
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
