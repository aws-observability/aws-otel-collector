// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ReportScheduleResponseAttributes The configuration and derived state of a report schedule.
type ReportScheduleResponseAttributes struct {
	// The delivery format for dashboard report schedules, or `null` if not set.
	DeliveryFormat NullableReportScheduleResponseAttributesDeliveryFormat `json:"delivery_format,omitempty"`
	// The description of the report.
	Description string `json:"description"`
	// The Unix timestamp, in milliseconds, of the next scheduled delivery, or `null` if none is scheduled.
	NextRecurrence datadog.NullableInt64 `json:"next_recurrence"`
	// The recipients of the report (email addresses, Slack channel references, or Microsoft Teams channel references).
	Recipients []string `json:"recipients"`
	// The identifier of the resource rendered in the report.
	ResourceId string `json:"resource_id"`
	// The type of dashboard resource the report schedule targets.
	ResourceType ReportScheduleResourceType `json:"resource_type"`
	// The recurrence rule for the schedule, expressed as an iCalendar `RRULE` string.
	Rrule string `json:"rrule"`
	// Whether the schedule is currently delivering reports (`active`) or paused (`inactive`).
	Status ReportScheduleStatus `json:"status"`
	// The identifier of the dashboard tab rendered in the report, or `null` if not set.
	TabId datadog.NullableString `json:"tab_id"`
	// The dashboard template variables applied when rendering the report.
	TemplateVariables []ReportScheduleTemplateVariable `json:"template_variables"`
	// The relative timeframe of data included in the report, or `null` if not set.
	Timeframe datadog.NullableString `json:"timeframe"`
	// The IANA time zone identifier the recurrence rule is evaluated in.
	Timezone string `json:"timezone"`
	// The title of the report.
	Title string `json:"title"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewReportScheduleResponseAttributes instantiates a new ReportScheduleResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewReportScheduleResponseAttributes(description string, nextRecurrence datadog.NullableInt64, recipients []string, resourceId string, resourceType ReportScheduleResourceType, rrule string, status ReportScheduleStatus, tabId datadog.NullableString, templateVariables []ReportScheduleTemplateVariable, timeframe datadog.NullableString, timezone string, title string) *ReportScheduleResponseAttributes {
	this := ReportScheduleResponseAttributes{}
	this.Description = description
	this.NextRecurrence = nextRecurrence
	this.Recipients = recipients
	this.ResourceId = resourceId
	this.ResourceType = resourceType
	this.Rrule = rrule
	this.Status = status
	this.TabId = tabId
	this.TemplateVariables = templateVariables
	this.Timeframe = timeframe
	this.Timezone = timezone
	this.Title = title
	return &this
}

// NewReportScheduleResponseAttributesWithDefaults instantiates a new ReportScheduleResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewReportScheduleResponseAttributesWithDefaults() *ReportScheduleResponseAttributes {
	this := ReportScheduleResponseAttributes{}
	return &this
}

// GetDeliveryFormat returns the DeliveryFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportScheduleResponseAttributes) GetDeliveryFormat() ReportScheduleResponseAttributesDeliveryFormat {
	if o == nil || o.DeliveryFormat.Get() == nil {
		var ret ReportScheduleResponseAttributesDeliveryFormat
		return ret
	}
	return *o.DeliveryFormat.Get()
}

// GetDeliveryFormatOk returns a tuple with the DeliveryFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleResponseAttributes) GetDeliveryFormatOk() (*ReportScheduleResponseAttributesDeliveryFormat, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeliveryFormat.Get(), o.DeliveryFormat.IsSet()
}

// HasDeliveryFormat returns a boolean if a field has been set.
func (o *ReportScheduleResponseAttributes) HasDeliveryFormat() bool {
	return o != nil && o.DeliveryFormat.IsSet()
}

// SetDeliveryFormat gets a reference to the given NullableReportScheduleResponseAttributesDeliveryFormat and assigns it to the DeliveryFormat field.
func (o *ReportScheduleResponseAttributes) SetDeliveryFormat(v ReportScheduleResponseAttributesDeliveryFormat) {
	o.DeliveryFormat.Set(&v)
}

// SetDeliveryFormatNil sets the value for DeliveryFormat to be an explicit nil.
func (o *ReportScheduleResponseAttributes) SetDeliveryFormatNil() {
	o.DeliveryFormat.Set(nil)
}

// UnsetDeliveryFormat ensures that no value is present for DeliveryFormat, not even an explicit nil.
func (o *ReportScheduleResponseAttributes) UnsetDeliveryFormat() {
	o.DeliveryFormat.Unset()
}

// GetDescription returns the Description field value.
func (o *ReportScheduleResponseAttributes) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value.
func (o *ReportScheduleResponseAttributes) SetDescription(v string) {
	o.Description = v
}

// GetNextRecurrence returns the NextRecurrence field value.
// If the value is explicit nil, the zero value for int64 will be returned.
func (o *ReportScheduleResponseAttributes) GetNextRecurrence() int64 {
	if o == nil || o.NextRecurrence.Get() == nil {
		var ret int64
		return ret
	}
	return *o.NextRecurrence.Get()
}

// GetNextRecurrenceOk returns a tuple with the NextRecurrence field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleResponseAttributes) GetNextRecurrenceOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextRecurrence.Get(), o.NextRecurrence.IsSet()
}

// SetNextRecurrence sets field value.
func (o *ReportScheduleResponseAttributes) SetNextRecurrence(v int64) {
	o.NextRecurrence.Set(&v)
}

// GetRecipients returns the Recipients field value.
func (o *ReportScheduleResponseAttributes) GetRecipients() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseAttributes) GetRecipientsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipients, true
}

// SetRecipients sets field value.
func (o *ReportScheduleResponseAttributes) SetRecipients(v []string) {
	o.Recipients = v
}

// GetResourceId returns the ResourceId field value.
func (o *ReportScheduleResponseAttributes) GetResourceId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseAttributes) GetResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceId, true
}

// SetResourceId sets field value.
func (o *ReportScheduleResponseAttributes) SetResourceId(v string) {
	o.ResourceId = v
}

// GetResourceType returns the ResourceType field value.
func (o *ReportScheduleResponseAttributes) GetResourceType() ReportScheduleResourceType {
	if o == nil {
		var ret ReportScheduleResourceType
		return ret
	}
	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseAttributes) GetResourceTypeOk() (*ReportScheduleResourceType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value.
func (o *ReportScheduleResponseAttributes) SetResourceType(v ReportScheduleResourceType) {
	o.ResourceType = v
}

// GetRrule returns the Rrule field value.
func (o *ReportScheduleResponseAttributes) GetRrule() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Rrule
}

// GetRruleOk returns a tuple with the Rrule field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseAttributes) GetRruleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rrule, true
}

// SetRrule sets field value.
func (o *ReportScheduleResponseAttributes) SetRrule(v string) {
	o.Rrule = v
}

// GetStatus returns the Status field value.
func (o *ReportScheduleResponseAttributes) GetStatus() ReportScheduleStatus {
	if o == nil {
		var ret ReportScheduleStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseAttributes) GetStatusOk() (*ReportScheduleStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value.
func (o *ReportScheduleResponseAttributes) SetStatus(v ReportScheduleStatus) {
	o.Status = v
}

// GetTabId returns the TabId field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ReportScheduleResponseAttributes) GetTabId() string {
	if o == nil || o.TabId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TabId.Get()
}

// GetTabIdOk returns a tuple with the TabId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleResponseAttributes) GetTabIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TabId.Get(), o.TabId.IsSet()
}

// SetTabId sets field value.
func (o *ReportScheduleResponseAttributes) SetTabId(v string) {
	o.TabId.Set(&v)
}

// GetTemplateVariables returns the TemplateVariables field value.
func (o *ReportScheduleResponseAttributes) GetTemplateVariables() []ReportScheduleTemplateVariable {
	if o == nil {
		var ret []ReportScheduleTemplateVariable
		return ret
	}
	return o.TemplateVariables
}

// GetTemplateVariablesOk returns a tuple with the TemplateVariables field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseAttributes) GetTemplateVariablesOk() (*[]ReportScheduleTemplateVariable, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateVariables, true
}

// SetTemplateVariables sets field value.
func (o *ReportScheduleResponseAttributes) SetTemplateVariables(v []ReportScheduleTemplateVariable) {
	o.TemplateVariables = v
}

// GetTimeframe returns the Timeframe field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *ReportScheduleResponseAttributes) GetTimeframe() string {
	if o == nil || o.Timeframe.Get() == nil {
		var ret string
		return ret
	}
	return *o.Timeframe.Get()
}

// GetTimeframeOk returns a tuple with the Timeframe field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *ReportScheduleResponseAttributes) GetTimeframeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timeframe.Get(), o.Timeframe.IsSet()
}

// SetTimeframe sets field value.
func (o *ReportScheduleResponseAttributes) SetTimeframe(v string) {
	o.Timeframe.Set(&v)
}

// GetTimezone returns the Timezone field value.
func (o *ReportScheduleResponseAttributes) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseAttributes) GetTimezoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value.
func (o *ReportScheduleResponseAttributes) SetTimezone(v string) {
	o.Timezone = v
}

// GetTitle returns the Title field value.
func (o *ReportScheduleResponseAttributes) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ReportScheduleResponseAttributes) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value.
func (o *ReportScheduleResponseAttributes) SetTitle(v string) {
	o.Title = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ReportScheduleResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DeliveryFormat.IsSet() {
		toSerialize["delivery_format"] = o.DeliveryFormat.Get()
	}
	toSerialize["description"] = o.Description
	toSerialize["next_recurrence"] = o.NextRecurrence.Get()
	toSerialize["recipients"] = o.Recipients
	toSerialize["resource_id"] = o.ResourceId
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["rrule"] = o.Rrule
	toSerialize["status"] = o.Status
	toSerialize["tab_id"] = o.TabId.Get()
	toSerialize["template_variables"] = o.TemplateVariables
	toSerialize["timeframe"] = o.Timeframe.Get()
	toSerialize["timezone"] = o.Timezone
	toSerialize["title"] = o.Title

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ReportScheduleResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DeliveryFormat    NullableReportScheduleResponseAttributesDeliveryFormat `json:"delivery_format,omitempty"`
		Description       *string                                                `json:"description"`
		NextRecurrence    datadog.NullableInt64                                  `json:"next_recurrence"`
		Recipients        *[]string                                              `json:"recipients"`
		ResourceId        *string                                                `json:"resource_id"`
		ResourceType      *ReportScheduleResourceType                            `json:"resource_type"`
		Rrule             *string                                                `json:"rrule"`
		Status            *ReportScheduleStatus                                  `json:"status"`
		TabId             datadog.NullableString                                 `json:"tab_id"`
		TemplateVariables *[]ReportScheduleTemplateVariable                      `json:"template_variables"`
		Timeframe         datadog.NullableString                                 `json:"timeframe"`
		Timezone          *string                                                `json:"timezone"`
		Title             *string                                                `json:"title"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Description == nil {
		return fmt.Errorf("required field description missing")
	}
	if !all.NextRecurrence.IsSet() {
		return fmt.Errorf("required field next_recurrence missing")
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
	if all.Status == nil {
		return fmt.Errorf("required field status missing")
	}
	if !all.TabId.IsSet() {
		return fmt.Errorf("required field tab_id missing")
	}
	if all.TemplateVariables == nil {
		return fmt.Errorf("required field template_variables missing")
	}
	if !all.Timeframe.IsSet() {
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
		datadog.DeleteKeys(additionalProperties, &[]string{"delivery_format", "description", "next_recurrence", "recipients", "resource_id", "resource_type", "rrule", "status", "tab_id", "template_variables", "timeframe", "timezone", "title"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.DeliveryFormat.Get() != nil && !all.DeliveryFormat.Get().IsValid() {
		hasInvalidField = true
	} else {
		o.DeliveryFormat = all.DeliveryFormat
	}
	o.Description = *all.Description
	o.NextRecurrence = all.NextRecurrence
	o.Recipients = *all.Recipients
	o.ResourceId = *all.ResourceId
	if !all.ResourceType.IsValid() {
		hasInvalidField = true
	} else {
		o.ResourceType = *all.ResourceType
	}
	o.Rrule = *all.Rrule
	if !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = *all.Status
	}
	o.TabId = all.TabId
	o.TemplateVariables = *all.TemplateVariables
	o.Timeframe = all.Timeframe
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
