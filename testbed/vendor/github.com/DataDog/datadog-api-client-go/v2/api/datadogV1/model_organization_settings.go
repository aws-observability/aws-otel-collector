// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OrganizationSettings A JSON array of settings.
type OrganizationSettings struct {
	// Whether or not the organization users can share widgets outside of Datadog.
	PrivateWidgetShare *bool `json:"private_widget_share,omitempty"`
	// Set the boolean property enabled to enable or disable single sign on with SAML.
	// See the SAML documentation for more information about all SAML settings.
	Saml *OrganizationSettingsSaml `json:"saml,omitempty"`
	// The access role of the user. Options are **st** (standard user), **adm** (admin user), or **ro** (read-only user).
	SamlAutocreateAccessRole NullableAccessRole `json:"saml_autocreate_access_role,omitempty"`
	// Has two properties, `enabled` (boolean) and `domains`, which is a list of domains without the @ symbol.
	SamlAutocreateUsersDomains *OrganizationSettingsSamlAutocreateUsersDomains `json:"saml_autocreate_users_domains,omitempty"`
	// Whether or not SAML can be enabled for this organization.
	SamlCanBeEnabled *bool `json:"saml_can_be_enabled,omitempty"`
	// Identity provider endpoint for SAML authentication.
	SamlIdpEndpoint *string `json:"saml_idp_endpoint,omitempty"`
	// Has one property enabled (boolean).
	SamlIdpInitiatedLogin *OrganizationSettingsSamlIdpInitiatedLogin `json:"saml_idp_initiated_login,omitempty"`
	// Whether or not a SAML identity provider metadata file was provided to the Datadog organization.
	SamlIdpMetadataUploaded *bool `json:"saml_idp_metadata_uploaded,omitempty"`
	// URL for SAML logging.
	SamlLoginUrl *string `json:"saml_login_url,omitempty"`
	// Has one property enabled (boolean).
	SamlStrictMode *OrganizationSettingsSamlStrictMode `json:"saml_strict_mode,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOrganizationSettings instantiates a new OrganizationSettings object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOrganizationSettings() *OrganizationSettings {
	this := OrganizationSettings{}
	return &this
}

// NewOrganizationSettingsWithDefaults instantiates a new OrganizationSettings object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOrganizationSettingsWithDefaults() *OrganizationSettings {
	this := OrganizationSettings{}
	return &this
}

// GetPrivateWidgetShare returns the PrivateWidgetShare field value if set, zero value otherwise.
func (o *OrganizationSettings) GetPrivateWidgetShare() bool {
	if o == nil || o.PrivateWidgetShare == nil {
		var ret bool
		return ret
	}
	return *o.PrivateWidgetShare
}

// GetPrivateWidgetShareOk returns a tuple with the PrivateWidgetShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettings) GetPrivateWidgetShareOk() (*bool, bool) {
	if o == nil || o.PrivateWidgetShare == nil {
		return nil, false
	}
	return o.PrivateWidgetShare, true
}

// HasPrivateWidgetShare returns a boolean if a field has been set.
func (o *OrganizationSettings) HasPrivateWidgetShare() bool {
	return o != nil && o.PrivateWidgetShare != nil
}

// SetPrivateWidgetShare gets a reference to the given bool and assigns it to the PrivateWidgetShare field.
func (o *OrganizationSettings) SetPrivateWidgetShare(v bool) {
	o.PrivateWidgetShare = &v
}

// GetSaml returns the Saml field value if set, zero value otherwise.
func (o *OrganizationSettings) GetSaml() OrganizationSettingsSaml {
	if o == nil || o.Saml == nil {
		var ret OrganizationSettingsSaml
		return ret
	}
	return *o.Saml
}

// GetSamlOk returns a tuple with the Saml field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettings) GetSamlOk() (*OrganizationSettingsSaml, bool) {
	if o == nil || o.Saml == nil {
		return nil, false
	}
	return o.Saml, true
}

// HasSaml returns a boolean if a field has been set.
func (o *OrganizationSettings) HasSaml() bool {
	return o != nil && o.Saml != nil
}

// SetSaml gets a reference to the given OrganizationSettingsSaml and assigns it to the Saml field.
func (o *OrganizationSettings) SetSaml(v OrganizationSettingsSaml) {
	o.Saml = &v
}

// GetSamlAutocreateAccessRole returns the SamlAutocreateAccessRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OrganizationSettings) GetSamlAutocreateAccessRole() AccessRole {
	if o == nil || o.SamlAutocreateAccessRole.Get() == nil {
		var ret AccessRole
		return ret
	}
	return *o.SamlAutocreateAccessRole.Get()
}

// GetSamlAutocreateAccessRoleOk returns a tuple with the SamlAutocreateAccessRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *OrganizationSettings) GetSamlAutocreateAccessRoleOk() (*AccessRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.SamlAutocreateAccessRole.Get(), o.SamlAutocreateAccessRole.IsSet()
}

// HasSamlAutocreateAccessRole returns a boolean if a field has been set.
func (o *OrganizationSettings) HasSamlAutocreateAccessRole() bool {
	return o != nil && o.SamlAutocreateAccessRole.IsSet()
}

// SetSamlAutocreateAccessRole gets a reference to the given NullableAccessRole and assigns it to the SamlAutocreateAccessRole field.
func (o *OrganizationSettings) SetSamlAutocreateAccessRole(v AccessRole) {
	o.SamlAutocreateAccessRole.Set(&v)
}

// SetSamlAutocreateAccessRoleNil sets the value for SamlAutocreateAccessRole to be an explicit nil.
func (o *OrganizationSettings) SetSamlAutocreateAccessRoleNil() {
	o.SamlAutocreateAccessRole.Set(nil)
}

// UnsetSamlAutocreateAccessRole ensures that no value is present for SamlAutocreateAccessRole, not even an explicit nil.
func (o *OrganizationSettings) UnsetSamlAutocreateAccessRole() {
	o.SamlAutocreateAccessRole.Unset()
}

// GetSamlAutocreateUsersDomains returns the SamlAutocreateUsersDomains field value if set, zero value otherwise.
func (o *OrganizationSettings) GetSamlAutocreateUsersDomains() OrganizationSettingsSamlAutocreateUsersDomains {
	if o == nil || o.SamlAutocreateUsersDomains == nil {
		var ret OrganizationSettingsSamlAutocreateUsersDomains
		return ret
	}
	return *o.SamlAutocreateUsersDomains
}

// GetSamlAutocreateUsersDomainsOk returns a tuple with the SamlAutocreateUsersDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettings) GetSamlAutocreateUsersDomainsOk() (*OrganizationSettingsSamlAutocreateUsersDomains, bool) {
	if o == nil || o.SamlAutocreateUsersDomains == nil {
		return nil, false
	}
	return o.SamlAutocreateUsersDomains, true
}

// HasSamlAutocreateUsersDomains returns a boolean if a field has been set.
func (o *OrganizationSettings) HasSamlAutocreateUsersDomains() bool {
	return o != nil && o.SamlAutocreateUsersDomains != nil
}

// SetSamlAutocreateUsersDomains gets a reference to the given OrganizationSettingsSamlAutocreateUsersDomains and assigns it to the SamlAutocreateUsersDomains field.
func (o *OrganizationSettings) SetSamlAutocreateUsersDomains(v OrganizationSettingsSamlAutocreateUsersDomains) {
	o.SamlAutocreateUsersDomains = &v
}

// GetSamlCanBeEnabled returns the SamlCanBeEnabled field value if set, zero value otherwise.
func (o *OrganizationSettings) GetSamlCanBeEnabled() bool {
	if o == nil || o.SamlCanBeEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SamlCanBeEnabled
}

// GetSamlCanBeEnabledOk returns a tuple with the SamlCanBeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettings) GetSamlCanBeEnabledOk() (*bool, bool) {
	if o == nil || o.SamlCanBeEnabled == nil {
		return nil, false
	}
	return o.SamlCanBeEnabled, true
}

// HasSamlCanBeEnabled returns a boolean if a field has been set.
func (o *OrganizationSettings) HasSamlCanBeEnabled() bool {
	return o != nil && o.SamlCanBeEnabled != nil
}

// SetSamlCanBeEnabled gets a reference to the given bool and assigns it to the SamlCanBeEnabled field.
func (o *OrganizationSettings) SetSamlCanBeEnabled(v bool) {
	o.SamlCanBeEnabled = &v
}

// GetSamlIdpEndpoint returns the SamlIdpEndpoint field value if set, zero value otherwise.
func (o *OrganizationSettings) GetSamlIdpEndpoint() string {
	if o == nil || o.SamlIdpEndpoint == nil {
		var ret string
		return ret
	}
	return *o.SamlIdpEndpoint
}

// GetSamlIdpEndpointOk returns a tuple with the SamlIdpEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettings) GetSamlIdpEndpointOk() (*string, bool) {
	if o == nil || o.SamlIdpEndpoint == nil {
		return nil, false
	}
	return o.SamlIdpEndpoint, true
}

// HasSamlIdpEndpoint returns a boolean if a field has been set.
func (o *OrganizationSettings) HasSamlIdpEndpoint() bool {
	return o != nil && o.SamlIdpEndpoint != nil
}

// SetSamlIdpEndpoint gets a reference to the given string and assigns it to the SamlIdpEndpoint field.
func (o *OrganizationSettings) SetSamlIdpEndpoint(v string) {
	o.SamlIdpEndpoint = &v
}

// GetSamlIdpInitiatedLogin returns the SamlIdpInitiatedLogin field value if set, zero value otherwise.
func (o *OrganizationSettings) GetSamlIdpInitiatedLogin() OrganizationSettingsSamlIdpInitiatedLogin {
	if o == nil || o.SamlIdpInitiatedLogin == nil {
		var ret OrganizationSettingsSamlIdpInitiatedLogin
		return ret
	}
	return *o.SamlIdpInitiatedLogin
}

// GetSamlIdpInitiatedLoginOk returns a tuple with the SamlIdpInitiatedLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettings) GetSamlIdpInitiatedLoginOk() (*OrganizationSettingsSamlIdpInitiatedLogin, bool) {
	if o == nil || o.SamlIdpInitiatedLogin == nil {
		return nil, false
	}
	return o.SamlIdpInitiatedLogin, true
}

// HasSamlIdpInitiatedLogin returns a boolean if a field has been set.
func (o *OrganizationSettings) HasSamlIdpInitiatedLogin() bool {
	return o != nil && o.SamlIdpInitiatedLogin != nil
}

// SetSamlIdpInitiatedLogin gets a reference to the given OrganizationSettingsSamlIdpInitiatedLogin and assigns it to the SamlIdpInitiatedLogin field.
func (o *OrganizationSettings) SetSamlIdpInitiatedLogin(v OrganizationSettingsSamlIdpInitiatedLogin) {
	o.SamlIdpInitiatedLogin = &v
}

// GetSamlIdpMetadataUploaded returns the SamlIdpMetadataUploaded field value if set, zero value otherwise.
func (o *OrganizationSettings) GetSamlIdpMetadataUploaded() bool {
	if o == nil || o.SamlIdpMetadataUploaded == nil {
		var ret bool
		return ret
	}
	return *o.SamlIdpMetadataUploaded
}

// GetSamlIdpMetadataUploadedOk returns a tuple with the SamlIdpMetadataUploaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettings) GetSamlIdpMetadataUploadedOk() (*bool, bool) {
	if o == nil || o.SamlIdpMetadataUploaded == nil {
		return nil, false
	}
	return o.SamlIdpMetadataUploaded, true
}

// HasSamlIdpMetadataUploaded returns a boolean if a field has been set.
func (o *OrganizationSettings) HasSamlIdpMetadataUploaded() bool {
	return o != nil && o.SamlIdpMetadataUploaded != nil
}

// SetSamlIdpMetadataUploaded gets a reference to the given bool and assigns it to the SamlIdpMetadataUploaded field.
func (o *OrganizationSettings) SetSamlIdpMetadataUploaded(v bool) {
	o.SamlIdpMetadataUploaded = &v
}

// GetSamlLoginUrl returns the SamlLoginUrl field value if set, zero value otherwise.
func (o *OrganizationSettings) GetSamlLoginUrl() string {
	if o == nil || o.SamlLoginUrl == nil {
		var ret string
		return ret
	}
	return *o.SamlLoginUrl
}

// GetSamlLoginUrlOk returns a tuple with the SamlLoginUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettings) GetSamlLoginUrlOk() (*string, bool) {
	if o == nil || o.SamlLoginUrl == nil {
		return nil, false
	}
	return o.SamlLoginUrl, true
}

// HasSamlLoginUrl returns a boolean if a field has been set.
func (o *OrganizationSettings) HasSamlLoginUrl() bool {
	return o != nil && o.SamlLoginUrl != nil
}

// SetSamlLoginUrl gets a reference to the given string and assigns it to the SamlLoginUrl field.
func (o *OrganizationSettings) SetSamlLoginUrl(v string) {
	o.SamlLoginUrl = &v
}

// GetSamlStrictMode returns the SamlStrictMode field value if set, zero value otherwise.
func (o *OrganizationSettings) GetSamlStrictMode() OrganizationSettingsSamlStrictMode {
	if o == nil || o.SamlStrictMode == nil {
		var ret OrganizationSettingsSamlStrictMode
		return ret
	}
	return *o.SamlStrictMode
}

// GetSamlStrictModeOk returns a tuple with the SamlStrictMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationSettings) GetSamlStrictModeOk() (*OrganizationSettingsSamlStrictMode, bool) {
	if o == nil || o.SamlStrictMode == nil {
		return nil, false
	}
	return o.SamlStrictMode, true
}

// HasSamlStrictMode returns a boolean if a field has been set.
func (o *OrganizationSettings) HasSamlStrictMode() bool {
	return o != nil && o.SamlStrictMode != nil
}

// SetSamlStrictMode gets a reference to the given OrganizationSettingsSamlStrictMode and assigns it to the SamlStrictMode field.
func (o *OrganizationSettings) SetSamlStrictMode(v OrganizationSettingsSamlStrictMode) {
	o.SamlStrictMode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OrganizationSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.PrivateWidgetShare != nil {
		toSerialize["private_widget_share"] = o.PrivateWidgetShare
	}
	if o.Saml != nil {
		toSerialize["saml"] = o.Saml
	}
	if o.SamlAutocreateAccessRole.IsSet() {
		toSerialize["saml_autocreate_access_role"] = o.SamlAutocreateAccessRole.Get()
	}
	if o.SamlAutocreateUsersDomains != nil {
		toSerialize["saml_autocreate_users_domains"] = o.SamlAutocreateUsersDomains
	}
	if o.SamlCanBeEnabled != nil {
		toSerialize["saml_can_be_enabled"] = o.SamlCanBeEnabled
	}
	if o.SamlIdpEndpoint != nil {
		toSerialize["saml_idp_endpoint"] = o.SamlIdpEndpoint
	}
	if o.SamlIdpInitiatedLogin != nil {
		toSerialize["saml_idp_initiated_login"] = o.SamlIdpInitiatedLogin
	}
	if o.SamlIdpMetadataUploaded != nil {
		toSerialize["saml_idp_metadata_uploaded"] = o.SamlIdpMetadataUploaded
	}
	if o.SamlLoginUrl != nil {
		toSerialize["saml_login_url"] = o.SamlLoginUrl
	}
	if o.SamlStrictMode != nil {
		toSerialize["saml_strict_mode"] = o.SamlStrictMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OrganizationSettings) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		PrivateWidgetShare         *bool                                           `json:"private_widget_share,omitempty"`
		Saml                       *OrganizationSettingsSaml                       `json:"saml,omitempty"`
		SamlAutocreateAccessRole   NullableAccessRole                              `json:"saml_autocreate_access_role,omitempty"`
		SamlAutocreateUsersDomains *OrganizationSettingsSamlAutocreateUsersDomains `json:"saml_autocreate_users_domains,omitempty"`
		SamlCanBeEnabled           *bool                                           `json:"saml_can_be_enabled,omitempty"`
		SamlIdpEndpoint            *string                                         `json:"saml_idp_endpoint,omitempty"`
		SamlIdpInitiatedLogin      *OrganizationSettingsSamlIdpInitiatedLogin      `json:"saml_idp_initiated_login,omitempty"`
		SamlIdpMetadataUploaded    *bool                                           `json:"saml_idp_metadata_uploaded,omitempty"`
		SamlLoginUrl               *string                                         `json:"saml_login_url,omitempty"`
		SamlStrictMode             *OrganizationSettingsSamlStrictMode             `json:"saml_strict_mode,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"private_widget_share", "saml", "saml_autocreate_access_role", "saml_autocreate_users_domains", "saml_can_be_enabled", "saml_idp_endpoint", "saml_idp_initiated_login", "saml_idp_metadata_uploaded", "saml_login_url", "saml_strict_mode"})
	} else {
		return err
	}

	hasInvalidField := false
	o.PrivateWidgetShare = all.PrivateWidgetShare
	if all.Saml != nil && all.Saml.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Saml = all.Saml
	if all.SamlAutocreateAccessRole.Get() != nil && !all.SamlAutocreateAccessRole.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.SamlAutocreateAccessRole = all.SamlAutocreateAccessRole
	}
	if all.SamlAutocreateUsersDomains != nil && all.SamlAutocreateUsersDomains.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SamlAutocreateUsersDomains = all.SamlAutocreateUsersDomains
	o.SamlCanBeEnabled = all.SamlCanBeEnabled
	o.SamlIdpEndpoint = all.SamlIdpEndpoint
	if all.SamlIdpInitiatedLogin != nil && all.SamlIdpInitiatedLogin.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SamlIdpInitiatedLogin = all.SamlIdpInitiatedLogin
	o.SamlIdpMetadataUploaded = all.SamlIdpMetadataUploaded
	o.SamlLoginUrl = all.SamlLoginUrl
	if all.SamlStrictMode != nil && all.SamlStrictMode.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.SamlStrictMode = all.SamlStrictMode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
