// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecurityMonitoringSignalSuggestedActionAttributes Attributes of a suggested action for a security signal. The available fields depend on the action type.
type SecurityMonitoringSignalSuggestedActionAttributes struct {
	// The name of the investigation log query.
	Name *string `json:"name,omitempty"`
	// The log query filter for the investigation.
	QueryFilter *string `json:"query_filter,omitempty"`
	// Template variables applied to the investigation log query, mapping attribute paths to values extracted from the signal.
	TemplateVariables map[string][]string `json:"template_variables,omitempty"`
	// The title of the recommended blog post.
	Title *string `json:"title,omitempty"`
	// The URL of the suggested action.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecurityMonitoringSignalSuggestedActionAttributes instantiates a new SecurityMonitoringSignalSuggestedActionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecurityMonitoringSignalSuggestedActionAttributes() *SecurityMonitoringSignalSuggestedActionAttributes {
	this := SecurityMonitoringSignalSuggestedActionAttributes{}
	return &this
}

// NewSecurityMonitoringSignalSuggestedActionAttributesWithDefaults instantiates a new SecurityMonitoringSignalSuggestedActionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecurityMonitoringSignalSuggestedActionAttributesWithDefaults() *SecurityMonitoringSignalSuggestedActionAttributes {
	this := SecurityMonitoringSignalSuggestedActionAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) SetName(v string) {
	o.Name = &v
}

// GetQueryFilter returns the QueryFilter field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) GetQueryFilter() string {
	if o == nil || o.QueryFilter == nil {
		var ret string
		return ret
	}
	return *o.QueryFilter
}

// GetQueryFilterOk returns a tuple with the QueryFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) GetQueryFilterOk() (*string, bool) {
	if o == nil || o.QueryFilter == nil {
		return nil, false
	}
	return o.QueryFilter, true
}

// HasQueryFilter returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) HasQueryFilter() bool {
	return o != nil && o.QueryFilter != nil
}

// SetQueryFilter gets a reference to the given string and assigns it to the QueryFilter field.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) SetQueryFilter(v string) {
	o.QueryFilter = &v
}

// GetTemplateVariables returns the TemplateVariables field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) GetTemplateVariables() map[string][]string {
	if o == nil || o.TemplateVariables == nil {
		var ret map[string][]string
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) GetTemplateVariablesOk() (*map[string][]string, bool) {
	if o == nil || o.TemplateVariables == nil {
		return nil, false
	}
	return &o.TemplateVariables, true
}

// HasTemplateVariables returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) HasTemplateVariables() bool {
	return o != nil && o.TemplateVariables != nil
}

// SetTemplateVariables gets a reference to the given map[string][]string and assigns it to the TemplateVariables field.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) SetTemplateVariables(v map[string][]string) {
	o.TemplateVariables = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) HasTitle() bool {
	return o != nil && o.Title != nil
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecurityMonitoringSignalSuggestedActionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.QueryFilter != nil {
		toSerialize["query_filter"] = o.QueryFilter
	}
	if o.TemplateVariables != nil {
		toSerialize["template_variables"] = o.TemplateVariables
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecurityMonitoringSignalSuggestedActionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name              *string             `json:"name,omitempty"`
		QueryFilter       *string             `json:"query_filter,omitempty"`
		TemplateVariables map[string][]string `json:"template_variables,omitempty"`
		Title             *string             `json:"title,omitempty"`
		Url               *string             `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"name", "query_filter", "template_variables", "title", "url"})
	} else {
		return err
	}
	o.Name = all.Name
	o.QueryFilter = all.QueryFilter
	o.TemplateVariables = all.TemplateVariables
	o.Title = all.Title
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
