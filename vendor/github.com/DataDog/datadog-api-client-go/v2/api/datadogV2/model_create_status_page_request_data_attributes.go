// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CreateStatusPageRequestDataAttributes The supported attributes for creating a status page.
type CreateStatusPageRequestDataAttributes struct {
	// The base64-encoded image data displayed on the status page.
	CompanyLogo *string `json:"company_logo,omitempty"`
	// The components displayed on the status page.
	Components []CreateStatusPageRequestDataAttributesComponentsItems `json:"components,omitempty"`
	// The subdomain of the status page's url taking the form `https://{domain_prefix}.statuspage.datadoghq.com`. Globally unique across Datadog Status Pages.
	DomainPrefix string `json:"domain_prefix"`
	// Base64-encoded image data included in email notifications sent to status page subscribers.
	EmailHeaderImage *string `json:"email_header_image,omitempty"`
	// Whether the status page is enabled.
	Enabled bool `json:"enabled"`
	// Base64-encoded image data displayed in the browser tab.
	Favicon *string `json:"favicon,omitempty"`
	// The name of the status page.
	Name string `json:"name"`
	// Whether users can subscribe to the status page.
	SubscriptionsEnabled *bool `json:"subscriptions_enabled,omitempty"`
	// The type of the status page controlling how the status page is accessed.
	Type CreateStatusPageRequestDataAttributesType `json:"type"`
	// The visualization type of the status page.
	VisualizationType CreateStatusPageRequestDataAttributesVisualizationType `json:"visualization_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCreateStatusPageRequestDataAttributes instantiates a new CreateStatusPageRequestDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCreateStatusPageRequestDataAttributes(domainPrefix string, enabled bool, name string, typeVar CreateStatusPageRequestDataAttributesType, visualizationType CreateStatusPageRequestDataAttributesVisualizationType) *CreateStatusPageRequestDataAttributes {
	this := CreateStatusPageRequestDataAttributes{}
	this.DomainPrefix = domainPrefix
	this.Enabled = enabled
	this.Name = name
	this.Type = typeVar
	this.VisualizationType = visualizationType
	return &this
}

// NewCreateStatusPageRequestDataAttributesWithDefaults instantiates a new CreateStatusPageRequestDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCreateStatusPageRequestDataAttributesWithDefaults() *CreateStatusPageRequestDataAttributes {
	this := CreateStatusPageRequestDataAttributes{}
	return &this
}

// GetCompanyLogo returns the CompanyLogo field value if set, zero value otherwise.
func (o *CreateStatusPageRequestDataAttributes) GetCompanyLogo() string {
	if o == nil || o.CompanyLogo == nil {
		var ret string
		return ret
	}
	return *o.CompanyLogo
}

// GetCompanyLogoOk returns a tuple with the CompanyLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStatusPageRequestDataAttributes) GetCompanyLogoOk() (*string, bool) {
	if o == nil || o.CompanyLogo == nil {
		return nil, false
	}
	return o.CompanyLogo, true
}

// HasCompanyLogo returns a boolean if a field has been set.
func (o *CreateStatusPageRequestDataAttributes) HasCompanyLogo() bool {
	return o != nil && o.CompanyLogo != nil
}

// SetCompanyLogo gets a reference to the given string and assigns it to the CompanyLogo field.
func (o *CreateStatusPageRequestDataAttributes) SetCompanyLogo(v string) {
	o.CompanyLogo = &v
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *CreateStatusPageRequestDataAttributes) GetComponents() []CreateStatusPageRequestDataAttributesComponentsItems {
	if o == nil || o.Components == nil {
		var ret []CreateStatusPageRequestDataAttributesComponentsItems
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStatusPageRequestDataAttributes) GetComponentsOk() (*[]CreateStatusPageRequestDataAttributesComponentsItems, bool) {
	if o == nil || o.Components == nil {
		return nil, false
	}
	return &o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *CreateStatusPageRequestDataAttributes) HasComponents() bool {
	return o != nil && o.Components != nil
}

// SetComponents gets a reference to the given []CreateStatusPageRequestDataAttributesComponentsItems and assigns it to the Components field.
func (o *CreateStatusPageRequestDataAttributes) SetComponents(v []CreateStatusPageRequestDataAttributesComponentsItems) {
	o.Components = v
}

// GetDomainPrefix returns the DomainPrefix field value.
func (o *CreateStatusPageRequestDataAttributes) GetDomainPrefix() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DomainPrefix
}

// GetDomainPrefixOk returns a tuple with the DomainPrefix field value
// and a boolean to check if the value has been set.
func (o *CreateStatusPageRequestDataAttributes) GetDomainPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DomainPrefix, true
}

// SetDomainPrefix sets field value.
func (o *CreateStatusPageRequestDataAttributes) SetDomainPrefix(v string) {
	o.DomainPrefix = v
}

// GetEmailHeaderImage returns the EmailHeaderImage field value if set, zero value otherwise.
func (o *CreateStatusPageRequestDataAttributes) GetEmailHeaderImage() string {
	if o == nil || o.EmailHeaderImage == nil {
		var ret string
		return ret
	}
	return *o.EmailHeaderImage
}

// GetEmailHeaderImageOk returns a tuple with the EmailHeaderImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStatusPageRequestDataAttributes) GetEmailHeaderImageOk() (*string, bool) {
	if o == nil || o.EmailHeaderImage == nil {
		return nil, false
	}
	return o.EmailHeaderImage, true
}

// HasEmailHeaderImage returns a boolean if a field has been set.
func (o *CreateStatusPageRequestDataAttributes) HasEmailHeaderImage() bool {
	return o != nil && o.EmailHeaderImage != nil
}

// SetEmailHeaderImage gets a reference to the given string and assigns it to the EmailHeaderImage field.
func (o *CreateStatusPageRequestDataAttributes) SetEmailHeaderImage(v string) {
	o.EmailHeaderImage = &v
}

// GetEnabled returns the Enabled field value.
func (o *CreateStatusPageRequestDataAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *CreateStatusPageRequestDataAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *CreateStatusPageRequestDataAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFavicon returns the Favicon field value if set, zero value otherwise.
func (o *CreateStatusPageRequestDataAttributes) GetFavicon() string {
	if o == nil || o.Favicon == nil {
		var ret string
		return ret
	}
	return *o.Favicon
}

// GetFaviconOk returns a tuple with the Favicon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStatusPageRequestDataAttributes) GetFaviconOk() (*string, bool) {
	if o == nil || o.Favicon == nil {
		return nil, false
	}
	return o.Favicon, true
}

// HasFavicon returns a boolean if a field has been set.
func (o *CreateStatusPageRequestDataAttributes) HasFavicon() bool {
	return o != nil && o.Favicon != nil
}

// SetFavicon gets a reference to the given string and assigns it to the Favicon field.
func (o *CreateStatusPageRequestDataAttributes) SetFavicon(v string) {
	o.Favicon = &v
}

// GetName returns the Name field value.
func (o *CreateStatusPageRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateStatusPageRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CreateStatusPageRequestDataAttributes) SetName(v string) {
	o.Name = v
}

// GetSubscriptionsEnabled returns the SubscriptionsEnabled field value if set, zero value otherwise.
func (o *CreateStatusPageRequestDataAttributes) GetSubscriptionsEnabled() bool {
	if o == nil || o.SubscriptionsEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SubscriptionsEnabled
}

// GetSubscriptionsEnabledOk returns a tuple with the SubscriptionsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateStatusPageRequestDataAttributes) GetSubscriptionsEnabledOk() (*bool, bool) {
	if o == nil || o.SubscriptionsEnabled == nil {
		return nil, false
	}
	return o.SubscriptionsEnabled, true
}

// HasSubscriptionsEnabled returns a boolean if a field has been set.
func (o *CreateStatusPageRequestDataAttributes) HasSubscriptionsEnabled() bool {
	return o != nil && o.SubscriptionsEnabled != nil
}

// SetSubscriptionsEnabled gets a reference to the given bool and assigns it to the SubscriptionsEnabled field.
func (o *CreateStatusPageRequestDataAttributes) SetSubscriptionsEnabled(v bool) {
	o.SubscriptionsEnabled = &v
}

// GetType returns the Type field value.
func (o *CreateStatusPageRequestDataAttributes) GetType() CreateStatusPageRequestDataAttributesType {
	if o == nil {
		var ret CreateStatusPageRequestDataAttributesType
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateStatusPageRequestDataAttributes) GetTypeOk() (*CreateStatusPageRequestDataAttributesType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *CreateStatusPageRequestDataAttributes) SetType(v CreateStatusPageRequestDataAttributesType) {
	o.Type = v
}

// GetVisualizationType returns the VisualizationType field value.
func (o *CreateStatusPageRequestDataAttributes) GetVisualizationType() CreateStatusPageRequestDataAttributesVisualizationType {
	if o == nil {
		var ret CreateStatusPageRequestDataAttributesVisualizationType
		return ret
	}
	return o.VisualizationType
}

// GetVisualizationTypeOk returns a tuple with the VisualizationType field value
// and a boolean to check if the value has been set.
func (o *CreateStatusPageRequestDataAttributes) GetVisualizationTypeOk() (*CreateStatusPageRequestDataAttributesVisualizationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisualizationType, true
}

// SetVisualizationType sets field value.
func (o *CreateStatusPageRequestDataAttributes) SetVisualizationType(v CreateStatusPageRequestDataAttributesVisualizationType) {
	o.VisualizationType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CreateStatusPageRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CompanyLogo != nil {
		toSerialize["company_logo"] = o.CompanyLogo
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	toSerialize["domain_prefix"] = o.DomainPrefix
	if o.EmailHeaderImage != nil {
		toSerialize["email_header_image"] = o.EmailHeaderImage
	}
	toSerialize["enabled"] = o.Enabled
	if o.Favicon != nil {
		toSerialize["favicon"] = o.Favicon
	}
	toSerialize["name"] = o.Name
	if o.SubscriptionsEnabled != nil {
		toSerialize["subscriptions_enabled"] = o.SubscriptionsEnabled
	}
	toSerialize["type"] = o.Type
	toSerialize["visualization_type"] = o.VisualizationType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CreateStatusPageRequestDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CompanyLogo          *string                                                 `json:"company_logo,omitempty"`
		Components           []CreateStatusPageRequestDataAttributesComponentsItems  `json:"components,omitempty"`
		DomainPrefix         *string                                                 `json:"domain_prefix"`
		EmailHeaderImage     *string                                                 `json:"email_header_image,omitempty"`
		Enabled              *bool                                                   `json:"enabled"`
		Favicon              *string                                                 `json:"favicon,omitempty"`
		Name                 *string                                                 `json:"name"`
		SubscriptionsEnabled *bool                                                   `json:"subscriptions_enabled,omitempty"`
		Type                 *CreateStatusPageRequestDataAttributesType              `json:"type"`
		VisualizationType    *CreateStatusPageRequestDataAttributesVisualizationType `json:"visualization_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DomainPrefix == nil {
		return fmt.Errorf("required field domain_prefix missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	if all.VisualizationType == nil {
		return fmt.Errorf("required field visualization_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"company_logo", "components", "domain_prefix", "email_header_image", "enabled", "favicon", "name", "subscriptions_enabled", "type", "visualization_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.CompanyLogo = all.CompanyLogo
	o.Components = all.Components
	o.DomainPrefix = *all.DomainPrefix
	o.EmailHeaderImage = all.EmailHeaderImage
	o.Enabled = *all.Enabled
	o.Favicon = all.Favicon
	o.Name = *all.Name
	o.SubscriptionsEnabled = all.SubscriptionsEnabled
	if !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = *all.Type
	}
	if !all.VisualizationType.IsValid() {
		hasInvalidField = true
	} else {
		o.VisualizationType = *all.VisualizationType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
