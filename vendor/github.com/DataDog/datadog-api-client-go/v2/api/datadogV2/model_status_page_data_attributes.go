// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// StatusPageDataAttributes The attributes of a status page.
type StatusPageDataAttributes struct {
	// Base64-encoded image data displayed on the status page.
	CompanyLogo datadog.NullableString `json:"company_logo,omitempty"`
	// Components displayed on the status page.
	Components []StatusPageDataAttributesComponentsItems `json:"components,omitempty"`
	// Timestamp of when the status page was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// If configured, the url that the status page is accessible at.
	CustomDomain datadog.NullableString `json:"custom_domain,omitempty"`
	// Whether the custom domain is configured.
	CustomDomainEnabled *bool `json:"custom_domain_enabled,omitempty"`
	// The subdomain of the status page's url taking the form `https://{domain_prefix}.statuspage.datadoghq.com`. Globally unique across Datadog Status Pages.
	DomainPrefix *string `json:"domain_prefix,omitempty"`
	// Base64-encoded image data included in email notifications sent to status page subscribers.
	EmailHeaderImage datadog.NullableString `json:"email_header_image,omitempty"`
	// Whether the status page is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Base64-encoded image data displayed in the browser tab.
	Favicon datadog.NullableString `json:"favicon,omitempty"`
	// Timestamp of when the status page was last modified.
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	// The name of the status page.
	Name *string `json:"name,omitempty"`
	// The url that the status page is accessible at.
	PageUrl *string `json:"page_url,omitempty"`
	// Whether users can subscribe to the status page.
	SubscriptionsEnabled *bool `json:"subscriptions_enabled,omitempty"`
	// The type of the status page controlling how the status page is accessed.
	Type *CreateStatusPageRequestDataAttributesType `json:"type,omitempty"`
	// The visualization type of the status page.
	VisualizationType *CreateStatusPageRequestDataAttributesVisualizationType `json:"visualization_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewStatusPageDataAttributes instantiates a new StatusPageDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewStatusPageDataAttributes() *StatusPageDataAttributes {
	this := StatusPageDataAttributes{}
	return &this
}

// NewStatusPageDataAttributesWithDefaults instantiates a new StatusPageDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewStatusPageDataAttributesWithDefaults() *StatusPageDataAttributes {
	this := StatusPageDataAttributes{}
	return &this
}

// GetCompanyLogo returns the CompanyLogo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusPageDataAttributes) GetCompanyLogo() string {
	if o == nil || o.CompanyLogo.Get() == nil {
		var ret string
		return ret
	}
	return *o.CompanyLogo.Get()
}

// GetCompanyLogoOk returns a tuple with the CompanyLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StatusPageDataAttributes) GetCompanyLogoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompanyLogo.Get(), o.CompanyLogo.IsSet()
}

// HasCompanyLogo returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasCompanyLogo() bool {
	return o != nil && o.CompanyLogo.IsSet()
}

// SetCompanyLogo gets a reference to the given datadog.NullableString and assigns it to the CompanyLogo field.
func (o *StatusPageDataAttributes) SetCompanyLogo(v string) {
	o.CompanyLogo.Set(&v)
}

// SetCompanyLogoNil sets the value for CompanyLogo to be an explicit nil.
func (o *StatusPageDataAttributes) SetCompanyLogoNil() {
	o.CompanyLogo.Set(nil)
}

// UnsetCompanyLogo ensures that no value is present for CompanyLogo, not even an explicit nil.
func (o *StatusPageDataAttributes) UnsetCompanyLogo() {
	o.CompanyLogo.Unset()
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *StatusPageDataAttributes) GetComponents() []StatusPageDataAttributesComponentsItems {
	if o == nil || o.Components == nil {
		var ret []StatusPageDataAttributesComponentsItems
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageDataAttributes) GetComponentsOk() (*[]StatusPageDataAttributesComponentsItems, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []StatusPageDataAttributesComponentsItems and assigns it to the Components field.
func (o *StatusPageDataAttributes) SetComponents(v []StatusPageDataAttributesComponentsItems) {
	o.Components = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *StatusPageDataAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasCreatedAt() bool {
	return o != nil && o.CreatedAt != nil
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *StatusPageDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCustomDomain returns the CustomDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusPageDataAttributes) GetCustomDomain() string {
	if o == nil || o.CustomDomain.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomDomain.Get()
}

// GetCustomDomainOk returns a tuple with the CustomDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StatusPageDataAttributes) GetCustomDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomDomain.Get(), o.CustomDomain.IsSet()
}

// HasCustomDomain returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasCustomDomain() bool {
	return o != nil && o.CustomDomain.IsSet()
}

// SetCustomDomain gets a reference to the given datadog.NullableString and assigns it to the CustomDomain field.
func (o *StatusPageDataAttributes) SetCustomDomain(v string) {
	o.CustomDomain.Set(&v)
}

// SetCustomDomainNil sets the value for CustomDomain to be an explicit nil.
func (o *StatusPageDataAttributes) SetCustomDomainNil() {
	o.CustomDomain.Set(nil)
}

// UnsetCustomDomain ensures that no value is present for CustomDomain, not even an explicit nil.
func (o *StatusPageDataAttributes) UnsetCustomDomain() {
	o.CustomDomain.Unset()
}

// GetCustomDomainEnabled returns the CustomDomainEnabled field value if set, zero value otherwise.
func (o *StatusPageDataAttributes) GetCustomDomainEnabled() bool {
	if o == nil || o.CustomDomainEnabled == nil {
		var ret bool
		return ret
	}
	return *o.CustomDomainEnabled
}

// GetCustomDomainEnabledOk returns a tuple with the CustomDomainEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageDataAttributes) GetCustomDomainEnabledOk() (*bool, bool) {
	if o == nil || o.CustomDomainEnabled == nil {
		return nil, false
	}
	return o.CustomDomainEnabled, true
}

// HasCustomDomainEnabled returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasCustomDomainEnabled() bool {
	return o != nil && o.CustomDomainEnabled != nil
}

// SetCustomDomainEnabled gets a reference to the given bool and assigns it to the CustomDomainEnabled field.
func (o *StatusPageDataAttributes) SetCustomDomainEnabled(v bool) {
	o.CustomDomainEnabled = &v
}

// GetDomainPrefix returns the DomainPrefix field value if set, zero value otherwise.
func (o *StatusPageDataAttributes) GetDomainPrefix() string {
	if o == nil || o.DomainPrefix == nil {
		var ret string
		return ret
	}
	return *o.DomainPrefix
}

// GetDomainPrefixOk returns a tuple with the DomainPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageDataAttributes) GetDomainPrefixOk() (*string, bool) {
	if o == nil || o.DomainPrefix == nil {
		return nil, false
	}
	return o.DomainPrefix, true
}

// HasDomainPrefix returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasDomainPrefix() bool {
	return o != nil && o.DomainPrefix != nil
}

// SetDomainPrefix gets a reference to the given string and assigns it to the DomainPrefix field.
func (o *StatusPageDataAttributes) SetDomainPrefix(v string) {
	o.DomainPrefix = &v
}

// GetEmailHeaderImage returns the EmailHeaderImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusPageDataAttributes) GetEmailHeaderImage() string {
	if o == nil || o.EmailHeaderImage.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmailHeaderImage.Get()
}

// GetEmailHeaderImageOk returns a tuple with the EmailHeaderImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StatusPageDataAttributes) GetEmailHeaderImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailHeaderImage.Get(), o.EmailHeaderImage.IsSet()
}

// HasEmailHeaderImage returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasEmailHeaderImage() bool {
	return o != nil && o.EmailHeaderImage.IsSet()
}

// SetEmailHeaderImage gets a reference to the given datadog.NullableString and assigns it to the EmailHeaderImage field.
func (o *StatusPageDataAttributes) SetEmailHeaderImage(v string) {
	o.EmailHeaderImage.Set(&v)
}

// SetEmailHeaderImageNil sets the value for EmailHeaderImage to be an explicit nil.
func (o *StatusPageDataAttributes) SetEmailHeaderImageNil() {
	o.EmailHeaderImage.Set(nil)
}

// UnsetEmailHeaderImage ensures that no value is present for EmailHeaderImage, not even an explicit nil.
func (o *StatusPageDataAttributes) UnsetEmailHeaderImage() {
	o.EmailHeaderImage.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StatusPageDataAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *StatusPageDataAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFavicon returns the Favicon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StatusPageDataAttributes) GetFavicon() string {
	if o == nil || o.Favicon.Get() == nil {
		var ret string
		return ret
	}
	return *o.Favicon.Get()
}

// GetFaviconOk returns a tuple with the Favicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *StatusPageDataAttributes) GetFaviconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Favicon.Get(), o.Favicon.IsSet()
}

// HasFavicon returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasFavicon() bool {
	return o != nil && o.Favicon.IsSet()
}

// SetFavicon gets a reference to the given datadog.NullableString and assigns it to the Favicon field.
func (o *StatusPageDataAttributes) SetFavicon(v string) {
	o.Favicon.Set(&v)
}

// SetFaviconNil sets the value for Favicon to be an explicit nil.
func (o *StatusPageDataAttributes) SetFaviconNil() {
	o.Favicon.Set(nil)
}

// UnsetFavicon ensures that no value is present for Favicon, not even an explicit nil.
func (o *StatusPageDataAttributes) UnsetFavicon() {
	o.Favicon.Unset()
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *StatusPageDataAttributes) GetModifiedAt() time.Time {
	if o == nil || o.ModifiedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageDataAttributes) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasModifiedAt() bool {
	return o != nil && o.ModifiedAt != nil
}

// SetModifiedAt gets a reference to the given time.Time and assigns it to the ModifiedAt field.
func (o *StatusPageDataAttributes) SetModifiedAt(v time.Time) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StatusPageDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StatusPageDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetPageUrl returns the PageUrl field value if set, zero value otherwise.
func (o *StatusPageDataAttributes) GetPageUrl() string {
	if o == nil || o.PageUrl == nil {
		var ret string
		return ret
	}
	return *o.PageUrl
}

// GetPageUrlOk returns a tuple with the PageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageDataAttributes) GetPageUrlOk() (*string, bool) {
	if o == nil || o.PageUrl == nil {
		return nil, false
	}
	return o.PageUrl, true
}

// HasPageUrl returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasPageUrl() bool {
	return o != nil && o.PageUrl != nil
}

// SetPageUrl gets a reference to the given string and assigns it to the PageUrl field.
func (o *StatusPageDataAttributes) SetPageUrl(v string) {
	o.PageUrl = &v
}

// GetSubscriptionsEnabled returns the SubscriptionsEnabled field value if set, zero value otherwise.
func (o *StatusPageDataAttributes) GetSubscriptionsEnabled() bool {
	if o == nil || o.SubscriptionsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SubscriptionsEnabled
}

// GetSubscriptionsEnabledOk returns a tuple with the SubscriptionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageDataAttributes) GetSubscriptionsEnabledOk() (*bool, bool) {
	if o == nil || o.SubscriptionsEnabled == nil {
		return nil, false
	}
	return o.SubscriptionsEnabled, true
}

// HasSubscriptionsEnabled returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasSubscriptionsEnabled() bool {
	return o != nil && o.SubscriptionsEnabled != nil
}

// SetSubscriptionsEnabled gets a reference to the given bool and assigns it to the SubscriptionsEnabled field.
func (o *StatusPageDataAttributes) SetSubscriptionsEnabled(v bool) {
	o.SubscriptionsEnabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StatusPageDataAttributes) GetType() CreateStatusPageRequestDataAttributesType {
	if o == nil || o.Type == nil {
		var ret CreateStatusPageRequestDataAttributesType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageDataAttributes) GetTypeOk() (*CreateStatusPageRequestDataAttributesType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given CreateStatusPageRequestDataAttributesType and assigns it to the Type field.
func (o *StatusPageDataAttributes) SetType(v CreateStatusPageRequestDataAttributesType) {
	o.Type = &v
}

// GetVisualizationType returns the VisualizationType field value if set, zero value otherwise.
func (o *StatusPageDataAttributes) GetVisualizationType() CreateStatusPageRequestDataAttributesVisualizationType {
	if o == nil || o.VisualizationType == nil {
		var ret CreateStatusPageRequestDataAttributesVisualizationType
		return ret
	}
	return *o.VisualizationType
}

// GetVisualizationTypeOk returns a tuple with the VisualizationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusPageDataAttributes) GetVisualizationTypeOk() (*CreateStatusPageRequestDataAttributesVisualizationType, bool) {
	if o == nil || o.VisualizationType == nil {
		return nil, false
	}
	return o.VisualizationType, true
}

// HasVisualizationType returns a boolean if a field has been set.
func (o *StatusPageDataAttributes) HasVisualizationType() bool {
	return o != nil && o.VisualizationType != nil
}

// SetVisualizationType gets a reference to the given CreateStatusPageRequestDataAttributesVisualizationType and assigns it to the VisualizationType field.
func (o *StatusPageDataAttributes) SetVisualizationType(v CreateStatusPageRequestDataAttributesVisualizationType) {
	o.VisualizationType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o StatusPageDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompanyLogo.IsSet() {
		toSerialize["company_logo"] = o.CompanyLogo.Get()
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	if o.CreatedAt != nil {
		if o.CreatedAt.Nanosecond() == 0 {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["created_at"] = o.CreatedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.CustomDomain.IsSet() {
		toSerialize["custom_domain"] = o.CustomDomain.Get()
	}
	if o.CustomDomainEnabled != nil {
		toSerialize["custom_domain_enabled"] = o.CustomDomainEnabled
	}
	if o.DomainPrefix != nil {
		toSerialize["domain_prefix"] = o.DomainPrefix
	}
	if o.EmailHeaderImage.IsSet() {
		toSerialize["email_header_image"] = o.EmailHeaderImage.Get()
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Favicon.IsSet() {
		toSerialize["favicon"] = o.Favicon.Get()
	}
	if o.ModifiedAt != nil {
		if o.ModifiedAt.Nanosecond() == 0 {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["modified_at"] = o.ModifiedAt.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PageUrl != nil {
		toSerialize["page_url"] = o.PageUrl
	}
	if o.SubscriptionsEnabled != nil {
		toSerialize["subscriptions_enabled"] = o.SubscriptionsEnabled
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.VisualizationType != nil {
		toSerialize["visualization_type"] = o.VisualizationType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *StatusPageDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompanyLogo          datadog.NullableString                                  `json:"company_logo,omitempty"`
		Components           []StatusPageDataAttributesComponentsItems               `json:"components,omitempty"`
		CreatedAt            *time.Time                                              `json:"created_at,omitempty"`
		CustomDomain         datadog.NullableString                                  `json:"custom_domain,omitempty"`
		CustomDomainEnabled  *bool                                                   `json:"custom_domain_enabled,omitempty"`
		DomainPrefix         *string                                                 `json:"domain_prefix,omitempty"`
		EmailHeaderImage     datadog.NullableString                                  `json:"email_header_image,omitempty"`
		Enabled              *bool                                                   `json:"enabled,omitempty"`
		Favicon              datadog.NullableString                                  `json:"favicon,omitempty"`
		ModifiedAt           *time.Time                                              `json:"modified_at,omitempty"`
		Name                 *string                                                 `json:"name,omitempty"`
		PageUrl              *string                                                 `json:"page_url,omitempty"`
		SubscriptionsEnabled *bool                                                   `json:"subscriptions_enabled,omitempty"`
		Type                 *CreateStatusPageRequestDataAttributesType              `json:"type,omitempty"`
		VisualizationType    *CreateStatusPageRequestDataAttributesVisualizationType `json:"visualization_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"company_logo", "components", "created_at", "custom_domain", "custom_domain_enabled", "domain_prefix", "email_header_image", "enabled", "favicon", "modified_at", "name", "page_url", "subscriptions_enabled", "type", "visualization_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CompanyLogo = all.CompanyLogo
	o.Components = all.Components
	o.CreatedAt = all.CreatedAt
	o.CustomDomain = all.CustomDomain
	o.CustomDomainEnabled = all.CustomDomainEnabled
	o.DomainPrefix = all.DomainPrefix
	o.EmailHeaderImage = all.EmailHeaderImage
	o.Enabled = all.Enabled
	o.Favicon = all.Favicon
	o.ModifiedAt = all.ModifiedAt
	o.Name = all.Name
	o.PageUrl = all.PageUrl
	o.SubscriptionsEnabled = all.SubscriptionsEnabled
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}
	if all.VisualizationType != nil && !all.VisualizationType.IsValid() {
		hasInvalidField = true
	} else {
		o.VisualizationType = all.VisualizationType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
