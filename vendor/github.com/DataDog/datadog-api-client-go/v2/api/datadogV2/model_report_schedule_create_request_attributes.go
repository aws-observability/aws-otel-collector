// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/google/uuid"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleCreateRequestAttributes The configuration of the report schedule to create.
type ReportScheduleCreateRequestAttributes struct {
	// How a PDF-export report is delivered. `pdf` attaches a PDF file, `png` embeds
	// an inline PNG image, and `pdf_and_png` delivers both.
	DeliveryFormat *ReportScheduleDeliveryFormat `json:"delivery_format,omitempty"`
	// A description of the report, up to 4096 characters.
	Description string `json:"description"`
	// The recipients of the report. Each entry is an email address, a Slack channel
	// reference in the form `slack:{team_id}.{channel_id}.{channel_name}`, or a Microsoft
	// Teams channel reference in the form `teams:{tenant_id}|{team_id}|{channel_id}`.
	Recipients []string `json:"recipients"`
	// The identifier of the dashboard or integration dashboard to render in the report.
	ResourceId string `json:"resource_id"`
	// The type of dashboard resource the report schedule targets.
	ResourceType ReportScheduleResourceType `json:"resource_type"`
	// The recurrence rule for the schedule, expressed as an iCalendar `RRULE` string.
	Rrule string `json:"rrule"`
	// The identifier of the dashboard tab to render, when the dashboard has tabs.
	TabId *uuid.UUID `json:"tab_id,omitempty"`
	// The dashboard template variables applied when rendering the report.
	TemplateVariables []ReportScheduleTemplateVariable `json:"template_variables"`
	// The relative timeframe of data to include in the report.
	Timeframe string `json:"timeframe"`
	// The IANA time zone identifier the recurrence rule is evaluated in.
	Timezone string `json:"timezone"`
	// The title of the report, between 1 and 78 characters.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportScheduleCreateRequestAttributes instantiates a new ReportScheduleCreateRequestAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportScheduleCreateRequestAttributes(description string, recipients []string, resourceId string, resourceType ReportScheduleResourceType, rrule string, templateVariables []ReportScheduleTemplateVariable, timeframe string, timezone string, title string) *ReportScheduleCreateRequestAttributes {
	this := ReportScheduleCreateRequestAttributes{}
	this.Description = description
	this.Recipients = recipients
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	this.Rrule = rrule
	this.TemplateVariables = templateVariables
	this.Timeframe = timeframe
	this.Timezone = timezone
	this.Title = title
	return &this
}

// NewReportScheduleCreateRequestAttributesWithDefaults instantiates a new ReportScheduleCreateRequestAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportScheduleCreateRequestAttributesWithDefaults() *ReportScheduleCreateRequestAttributes {
	this := ReportScheduleCreateRequestAttributes{}
	return &this
}

// GetDeliveryFormat returns the DeliveryFormat field value if set, zero value otherwise.
func (o *ReportScheduleCreateRequestAttributes) GetDeliveryFormat() ReportScheduleDeliveryFormat {
	if o == nil || o.DeliveryFormat == nil {
		var ret ReportScheduleDeliveryFormat
		return ret
	}
	return *o.DeliveryFormat
}

// GetDeliveryFormatOk returns a tuple with the DeliveryFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportScheduleCreateRequestAttributes) GetDeliveryFormatOk() (*ReportScheduleDeliveryFormat, bool) {
	if o == nil || o.DeliveryFormat == nil {
		return nil, false
	}
	return o.DeliveryFormat, true
}

// HasDeliveryFormat returns a boolean if a field has been set.
func (o *ReportScheduleCreateRequestAttributes) HasDeliveryFormat() bool {
	return o != nil && o.DeliveryFormat != nil
}

// SetDeliveryFormat gets a reference to the given ReportScheduleDeliveryFormat and assigns it to the DeliveryFormat field.
func (o *ReportScheduleCreateRequestAttributes) SetDeliveryFormat(v ReportScheduleDeliveryFormat) {
	o.DeliveryFormat = &v
}

// GetDescription returns the Description field value.
func (o *ReportScheduleCreateRequestAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleCreateRequestAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ReportScheduleCreateRequestAttributes) SetDescription(v string) {
	o.Description = v
}

// GetRecipients returns the Recipients field value.
func (o *ReportScheduleCreateRequestAttributes) GetRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleCreateRequestAttributes) GetRecipientsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value.
func (o *ReportScheduleCreateRequestAttributes) SetRecipients(v []string) {
	o.Recipients = v
}

// GetResourceId returns the ResourceId field value.
func (o *ReportScheduleCreateRequestAttributes) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleCreateRequestAttributes) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value.
func (o *ReportScheduleCreateRequestAttributes) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value.
func (o *ReportScheduleCreateRequestAttributes) GetResourceType() ReportScheduleResourceType {
	if o == nil {
		var ret ReportScheduleResourceType
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleCreateRequestAttributes) GetResourceTypeOk() (*ReportScheduleResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *ReportScheduleCreateRequestAttributes) SetResourceType(v ReportScheduleResourceType) {
	o.ResourceType = v
}

// GetRrule returns the Rrule field value.
func (o *ReportScheduleCreateRequestAttributes) GetRrule() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleCreateRequestAttributes) GetRruleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rrule, true
}

// SetRrule sets field value.
func (o *ReportScheduleCreateRequestAttributes) SetRrule(v string) {
	o.Rrule = v
}

// GetTabId returns the TabId field value if set, zero value otherwise.
func (o *ReportScheduleCreateRequestAttributes) GetTabId() uuid.UUID {
	if o == nil || o.TabId == nil {
		var ret uuid.UUID
		return ret
	}
	return *o.TabId
}

// GetTabIdOk returns a tuple with the TabId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportScheduleCreateRequestAttributes) GetTabIdOk() (*uuid.UUID, bool) {
	if o == nil || o.TabId == nil {
		return nil, false
	}
	return o.TabId, true
}

// HasTabId returns a boolean if a field has been set.
func (o *ReportScheduleCreateRequestAttributes) HasTabId() bool {
	return o != nil && o.TabId != nil
}

// SetTabId gets a reference to the given uuid.UUID and assigns it to the TabId field.
func (o *ReportScheduleCreateRequestAttributes) SetTabId(v uuid.UUID) {
	o.TabId = &v
}

// GetTemplateVariables returns the TemplateVariables field value.
func (o *ReportScheduleCreateRequestAttributes) GetTemplateVariables() []ReportScheduleTemplateVariable {
	if o == nil {
		var ret []ReportScheduleTemplateVariable
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleCreateRequestAttributes) GetTemplateVariablesOk() (*[]ReportScheduleTemplateVariable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateVariables, true
}

// SetTemplateVariables sets field value.
func (o *ReportScheduleCreateRequestAttributes) SetTemplateVariables(v []ReportScheduleTemplateVariable) {
	o.TemplateVariables = v
}

// GetTimeframe returns the Timeframe field value.
func (o *ReportScheduleCreateRequestAttributes) GetTimeframe() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleCreateRequestAttributes) GetTimeframeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeframe, true
}

// SetTimeframe sets field value.
func (o *ReportScheduleCreateRequestAttributes) SetTimeframe(v string) {
	o.Timeframe = v
}

// GetTimezone returns the Timezone field value.
func (o *ReportScheduleCreateRequestAttributes) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleCreateRequestAttributes) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value.
func (o *ReportScheduleCreateRequestAttributes) SetTimezone(v string) {
	o.Timezone = v
}

// GetTitle returns the Title field value.
func (o *ReportScheduleCreateRequestAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleCreateRequestAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *ReportScheduleCreateRequestAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportScheduleCreateRequestAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DeliveryFormat != nil {
		toSerialize["delivery_format"] = o.DeliveryFormat
	}
	toSerialize["description"] = o.Description
	toSerialize["recipients"] = o.Recipients
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["rrule"] = o.Rrule
	if o.TabId != nil {
		toSerialize["tab_id"] = o.TabId
	}
	toSerialize["template_variables"] = o.TemplateVariables
	toSerialize["timeframe"] = o.Timeframe
	toSerialize["timezone"] = o.Timezone
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReportScheduleCreateRequestAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeliveryFormat    *ReportScheduleDeliveryFormat     `json:"delivery_format,omitempty"`
		Description       *string                           `json:"description"`
		Recipients        *[]string                         `json:"recipients"`
		ResourceId        *string                           `json:"resource_id"`
		ResourceType      *ReportScheduleResourceType       `json:"resource_type"`
		Rrule             *string                           `json:"rrule"`
		TabId             *uuid.UUID                        `json:"tab_id,omitempty"`
		TemplateVariables *[]ReportScheduleTemplateVariable `json:"template_variables"`
		Timeframe         *string                           `json:"timeframe"`
		Timezone          *string                           `json:"timezone"`
		Title             *string                           `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if all.Recipients == nil {
		return fmt.Errorf("required field recipients missing")
	}
	if all.ResourceId == nil {
		return fmt.Errorf("required field resource_id missing")
	}
	if all.ResourceType == nil {
		return fmt.Errorf("required field resource_type missing")
	}
	if all.Rrule == nil {
		return fmt.Errorf("required field rrule missing")
	}
	if all.TemplateVariables == nil {
		return fmt.Errorf("required field template_variables missing")
	}
	if all.Timeframe == nil {
		return fmt.Errorf("required field timeframe missing")
	}
	if all.Timezone == nil {
		return fmt.Errorf("required field timezone missing")
	}
	if all.Title == nil {
		return fmt.Errorf("required field title missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"delivery_format", "description", "recipients", "resource_id", "resource_type", "rrule", "tab_id", "template_variables", "timeframe", "timezone", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DeliveryFormat != nil && !all.DeliveryFormat.IsValid() {
		hasInvalidField = true
	} else {
		o.DeliveryFormat = all.DeliveryFormat
	}
	o.Description = *all.Description
	o.Recipients = *all.Recipients
	o.ResourceId = *all.ResourceId
	if !all.ResourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceType = *all.ResourceType
	}
	o.Rrule = *all.Rrule
	o.TabId = all.TabId
	o.TemplateVariables = *all.TemplateVariables
	o.Timeframe = *all.Timeframe
	o.Timezone = *all.Timezone
	o.Title = *all.Title

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
