// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchStatusPageRequestDataAttributes The supported attributes for updating a status page.
type PatchStatusPageRequestDataAttributes struct {
	// The base64-encoded image data displayed on the status page.
	CompanyLogo *string `json:"company_logo,omitempty"`
	// The subdomain of the status page's url taking the form `https://{domain_prefix}.statuspage.datadoghq.com`. Globally unique across Datadog Status Pages.
	DomainPrefix *string `json:"domain_prefix,omitempty"`
	// The base64-encoded image data displayed in email notifications sent to status page subscribers.
	EmailHeaderImage *string `json:"email_header_image,omitempty"`
	// Whether the status page is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The base64-encoded image data displayed in the browser tab.
	Favicon *string `json:"favicon,omitempty"`
	// The name of the status page.
	Name *string `json:"name,omitempty"`
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

// NewPatchStatusPageRequestDataAttributes instantiates a new PatchStatusPageRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchStatusPageRequestDataAttributes() *PatchStatusPageRequestDataAttributes {
	this := PatchStatusPageRequestDataAttributes{}
	return &this
}

// NewPatchStatusPageRequestDataAttributesWithDefaults instantiates a new PatchStatusPageRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchStatusPageRequestDataAttributesWithDefaults() *PatchStatusPageRequestDataAttributes {
	this := PatchStatusPageRequestDataAttributes{}
	return &this
}

// GetCompanyLogo returns the CompanyLogo field value if set, zero value otherwise.
func (o *PatchStatusPageRequestDataAttributes) GetCompanyLogo() string {
	if o == nil || o.CompanyLogo == nil {
		var ret string
		return ret
	}
	return *o.CompanyLogo
}

// GetCompanyLogoOk returns a tuple with the CompanyLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchStatusPageRequestDataAttributes) GetCompanyLogoOk() (*string, bool) {
	if o == nil || o.CompanyLogo == nil {
		return nil, false
	}
	return o.CompanyLogo, true
}

// HasCompanyLogo returns a boolean if a field has been set.
func (o *PatchStatusPageRequestDataAttributes) HasCompanyLogo() bool {
	return o != nil && o.CompanyLogo != nil
}

// SetCompanyLogo gets a reference to the given string and assigns it to the CompanyLogo field.
func (o *PatchStatusPageRequestDataAttributes) SetCompanyLogo(v string) {
	o.CompanyLogo = &v
}

// GetDomainPrefix returns the DomainPrefix field value if set, zero value otherwise.
func (o *PatchStatusPageRequestDataAttributes) GetDomainPrefix() string {
	if o == nil || o.DomainPrefix == nil {
		var ret string
		return ret
	}
	return *o.DomainPrefix
}

// GetDomainPrefixOk returns a tuple with the DomainPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchStatusPageRequestDataAttributes) GetDomainPrefixOk() (*string, bool) {
	if o == nil || o.DomainPrefix == nil {
		return nil, false
	}
	return o.DomainPrefix, true
}

// HasDomainPrefix returns a boolean if a field has been set.
func (o *PatchStatusPageRequestDataAttributes) HasDomainPrefix() bool {
	return o != nil && o.DomainPrefix != nil
}

// SetDomainPrefix gets a reference to the given string and assigns it to the DomainPrefix field.
func (o *PatchStatusPageRequestDataAttributes) SetDomainPrefix(v string) {
	o.DomainPrefix = &v
}

// GetEmailHeaderImage returns the EmailHeaderImage field value if set, zero value otherwise.
func (o *PatchStatusPageRequestDataAttributes) GetEmailHeaderImage() string {
	if o == nil || o.EmailHeaderImage == nil {
		var ret string
		return ret
	}
	return *o.EmailHeaderImage
}

// GetEmailHeaderImageOk returns a tuple with the EmailHeaderImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchStatusPageRequestDataAttributes) GetEmailHeaderImageOk() (*string, bool) {
	if o == nil || o.EmailHeaderImage == nil {
		return nil, false
	}
	return o.EmailHeaderImage, true
}

// HasEmailHeaderImage returns a boolean if a field has been set.
func (o *PatchStatusPageRequestDataAttributes) HasEmailHeaderImage() bool {
	return o != nil && o.EmailHeaderImage != nil
}

// SetEmailHeaderImage gets a reference to the given string and assigns it to the EmailHeaderImage field.
func (o *PatchStatusPageRequestDataAttributes) SetEmailHeaderImage(v string) {
	o.EmailHeaderImage = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PatchStatusPageRequestDataAttributes) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchStatusPageRequestDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PatchStatusPageRequestDataAttributes) HasEnabled() bool {
	return o != nil && o.Enabled != nil
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PatchStatusPageRequestDataAttributes) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFavicon returns the Favicon field value if set, zero value otherwise.
func (o *PatchStatusPageRequestDataAttributes) GetFavicon() string {
	if o == nil || o.Favicon == nil {
		var ret string
		return ret
	}
	return *o.Favicon
}

// GetFaviconOk returns a tuple with the Favicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchStatusPageRequestDataAttributes) GetFaviconOk() (*string, bool) {
	if o == nil || o.Favicon == nil {
		return nil, false
	}
	return o.Favicon, true
}

// HasFavicon returns a boolean if a field has been set.
func (o *PatchStatusPageRequestDataAttributes) HasFavicon() bool {
	return o != nil && o.Favicon != nil
}

// SetFavicon gets a reference to the given string and assigns it to the Favicon field.
func (o *PatchStatusPageRequestDataAttributes) SetFavicon(v string) {
	o.Favicon = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchStatusPageRequestDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchStatusPageRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchStatusPageRequestDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchStatusPageRequestDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetSubscriptionsEnabled returns the SubscriptionsEnabled field value if set, zero value otherwise.
func (o *PatchStatusPageRequestDataAttributes) GetSubscriptionsEnabled() bool {
	if o == nil || o.SubscriptionsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SubscriptionsEnabled
}

// GetSubscriptionsEnabledOk returns a tuple with the SubscriptionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchStatusPageRequestDataAttributes) GetSubscriptionsEnabledOk() (*bool, bool) {
	if o == nil || o.SubscriptionsEnabled == nil {
		return nil, false
	}
	return o.SubscriptionsEnabled, true
}

// HasSubscriptionsEnabled returns a boolean if a field has been set.
func (o *PatchStatusPageRequestDataAttributes) HasSubscriptionsEnabled() bool {
	return o != nil && o.SubscriptionsEnabled != nil
}

// SetSubscriptionsEnabled gets a reference to the given bool and assigns it to the SubscriptionsEnabled field.
func (o *PatchStatusPageRequestDataAttributes) SetSubscriptionsEnabled(v bool) {
	o.SubscriptionsEnabled = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PatchStatusPageRequestDataAttributes) GetType() CreateStatusPageRequestDataAttributesType {
	if o == nil || o.Type == nil {
		var ret CreateStatusPageRequestDataAttributesType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchStatusPageRequestDataAttributes) GetTypeOk() (*CreateStatusPageRequestDataAttributesType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PatchStatusPageRequestDataAttributes) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given CreateStatusPageRequestDataAttributesType and assigns it to the Type field.
func (o *PatchStatusPageRequestDataAttributes) SetType(v CreateStatusPageRequestDataAttributesType) {
	o.Type = &v
}

// GetVisualizationType returns the VisualizationType field value if set, zero value otherwise.
func (o *PatchStatusPageRequestDataAttributes) GetVisualizationType() CreateStatusPageRequestDataAttributesVisualizationType {
	if o == nil || o.VisualizationType == nil {
		var ret CreateStatusPageRequestDataAttributesVisualizationType
		return ret
	}
	return *o.VisualizationType
}

// GetVisualizationTypeOk returns a tuple with the VisualizationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchStatusPageRequestDataAttributes) GetVisualizationTypeOk() (*CreateStatusPageRequestDataAttributesVisualizationType, bool) {
	if o == nil || o.VisualizationType == nil {
		return nil, false
	}
	return o.VisualizationType, true
}

// HasVisualizationType returns a boolean if a field has been set.
func (o *PatchStatusPageRequestDataAttributes) HasVisualizationType() bool {
	return o != nil && o.VisualizationType != nil
}

// SetVisualizationType gets a reference to the given CreateStatusPageRequestDataAttributesVisualizationType and assigns it to the VisualizationType field.
func (o *PatchStatusPageRequestDataAttributes) SetVisualizationType(v CreateStatusPageRequestDataAttributesVisualizationType) {
	o.VisualizationType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchStatusPageRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompanyLogo != nil {
		toSerialize["company_logo"] = o.CompanyLogo
	}
	if o.DomainPrefix != nil {
		toSerialize["domain_prefix"] = o.DomainPrefix
	}
	if o.EmailHeaderImage != nil {
		toSerialize["email_header_image"] = o.EmailHeaderImage
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Favicon != nil {
		toSerialize["favicon"] = o.Favicon
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
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
func (o *PatchStatusPageRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompanyLogo          *string                                                 `json:"company_logo,omitempty"`
		DomainPrefix         *string                                                 `json:"domain_prefix,omitempty"`
		EmailHeaderImage     *string                                                 `json:"email_header_image,omitempty"`
		Enabled              *bool                                                   `json:"enabled,omitempty"`
		Favicon              *string                                                 `json:"favicon,omitempty"`
		Name                 *string                                                 `json:"name,omitempty"`
		SubscriptionsEnabled *bool                                                   `json:"subscriptions_enabled,omitempty"`
		Type                 *CreateStatusPageRequestDataAttributesType              `json:"type,omitempty"`
		VisualizationType    *CreateStatusPageRequestDataAttributesVisualizationType `json:"visualization_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"company_logo", "domain_prefix", "email_header_image", "enabled", "favicon", "name", "subscriptions_enabled", "type", "visualization_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CompanyLogo = all.CompanyLogo
	o.DomainPrefix = all.DomainPrefix
	o.EmailHeaderImage = all.EmailHeaderImage
	o.Enabled = all.Enabled
	o.Favicon = all.Favicon
	o.Name = all.Name
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
