// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleResourceAttributes Attributes of an included report target resource.
type ReportScheduleResourceAttributes struct {
	// The type of dashboard resource the report schedule targets.
	ResourceType ReportScheduleResourceType `json:"resource_type"`
	// Template variable metadata from the dashboard resource, when available.
	TemplateVariables []ReportScheduleIndexTemplateVariable `json:"template_variables,omitempty"`
	// The title of the dashboard or integration dashboard resource, when available.
	Title datadog.NullableString `json:"title,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportScheduleResourceAttributes instantiates a new ReportScheduleResourceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportScheduleResourceAttributes(resourceType ReportScheduleResourceType) *ReportScheduleResourceAttributes {
	this := ReportScheduleResourceAttributes{}
	this.ResourceType = resourceType
	return &this
}

// NewReportScheduleResourceAttributesWithDefaults instantiates a new ReportScheduleResourceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportScheduleResourceAttributesWithDefaults() *ReportScheduleResourceAttributes {
	this := ReportScheduleResourceAttributes{}
	return &this
}

// GetResourceType returns the ResourceType field value.
func (o *ReportScheduleResourceAttributes) GetResourceType() ReportScheduleResourceType {
	if o == nil {
		var ret ReportScheduleResourceType
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResourceAttributes) GetResourceTypeOk() (*ReportScheduleResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *ReportScheduleResourceAttributes) SetResourceType(v ReportScheduleResourceType) {
	o.ResourceType = v
}

// GetTemplateVariables returns the TemplateVariables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduleResourceAttributes) GetTemplateVariables() []ReportScheduleIndexTemplateVariable {
	if o == nil {
		var ret []ReportScheduleIndexTemplateVariable
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleResourceAttributes) GetTemplateVariablesOk() (*[]ReportScheduleIndexTemplateVariable, bool) {
	if o == nil || o.TemplateVariables == nil {
		return nil, false
	}
	return &o.TemplateVariables, true
}

// HasTemplateVariables returns a boolean if a field has been set.
func (o *ReportScheduleResourceAttributes) HasTemplateVariables() bool {
	return o != nil && o.TemplateVariables != nil
}

// SetTemplateVariables gets a reference to the given []ReportScheduleIndexTemplateVariable and assigns it to the TemplateVariables field.
func (o *ReportScheduleResourceAttributes) SetTemplateVariables(v []ReportScheduleIndexTemplateVariable) {
	o.TemplateVariables = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduleResourceAttributes) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleResourceAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ReportScheduleResourceAttributes) HasTitle() bool {
	return o != nil && o.Title.IsSet()
}

// SetTitle gets a reference to the given datadog.NullableString and assigns it to the Title field.
func (o *ReportScheduleResourceAttributes) SetTitle(v string) {
	o.Title.Set(&v)
}

// SetTitleNil sets the value for Title to be an explicit nil.
func (o *ReportScheduleResourceAttributes) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil.
func (o *ReportScheduleResourceAttributes) UnsetTitle() {
	o.Title.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportScheduleResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["resource_type"] = o.ResourceType
	if o.TemplateVariables != nil {
		toSerialize["template_variables"] = o.TemplateVariables
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReportScheduleResourceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ResourceType      *ReportScheduleResourceType           `json:"resource_type"`
		TemplateVariables []ReportScheduleIndexTemplateVariable `json:"template_variables,omitempty"`
		Title             datadog.NullableString                `json:"title,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.ResourceType == nil {
		return fmt.Errorf("required field resource_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"resource_type", "template_variables", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	if !all.ResourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceType = *all.ResourceType
	}
	o.TemplateVariables = all.TemplateVariables
	o.Title = all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
