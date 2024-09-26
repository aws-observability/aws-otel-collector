// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonthlyCostAttributionAttributes Cost Attribution by Tag for a given organization.
type MonthlyCostAttributionAttributes struct {
	// Datetime in ISO-8601 format, UTC, precise to hour: `[YYYY-MM-DDThh]`.
	Month *time.Time `json:"month,omitempty"`
	// The name of the organization.
	OrgName *string `json:"org_name,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// The source of the cost attribution tag configuration and the selected tags in the format `<source_org_name>:::<selected tag 1>///<selected tag 2>///<selected tag 3>`.
	TagConfigSource *string `json:"tag_config_source,omitempty"`
	// Tag keys and values.
	// A `null` value here means that the requested tag breakdown cannot be applied because it does not match the [tags
	// configured for usage attribution](https://docs.datadoghq.com/account_management/billing/usage_attribution/#getting-started).
	// In this scenario the API returns the total cost, not broken down by tags.
	Tags map[string][]string `json:"tags,omitempty"`
	// Shows the most recent hour in the current months for all organizations for which all costs were calculated.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// Fields in Cost Attribution by tag(s). Example: `infra_host_on_demand_cost`, `infra_host_committed_cost`, `infra_host_total_cost`, `infra_host_percentage_in_org`, `infra_host_percentage_in_account`.
	Values interface{} `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonthlyCostAttributionAttributes instantiates a new MonthlyCostAttributionAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonthlyCostAttributionAttributes() *MonthlyCostAttributionAttributes {
	this := MonthlyCostAttributionAttributes{}
	return &this
}

// NewMonthlyCostAttributionAttributesWithDefaults instantiates a new MonthlyCostAttributionAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonthlyCostAttributionAttributesWithDefaults() *MonthlyCostAttributionAttributes {
	this := MonthlyCostAttributionAttributes{}
	return &this
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *MonthlyCostAttributionAttributes) GetMonth() time.Time {
	if o == nil || o.Month == nil {
		var ret time.Time
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyCostAttributionAttributes) GetMonthOk() (*time.Time, bool) {
	if o == nil || o.Month == nil {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *MonthlyCostAttributionAttributes) HasMonth() bool {
	return o != nil && o.Month != nil
}

// SetMonth gets a reference to the given time.Time and assigns it to the Month field.
func (o *MonthlyCostAttributionAttributes) SetMonth(v time.Time) {
	o.Month = &v
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *MonthlyCostAttributionAttributes) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyCostAttributionAttributes) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *MonthlyCostAttributionAttributes) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *MonthlyCostAttributionAttributes) SetOrgName(v string) {
	o.OrgName = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *MonthlyCostAttributionAttributes) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyCostAttributionAttributes) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *MonthlyCostAttributionAttributes) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *MonthlyCostAttributionAttributes) SetPublicId(v string) {
	o.PublicId = &v
}

// GetTagConfigSource returns the TagConfigSource field value if set, zero value otherwise.
func (o *MonthlyCostAttributionAttributes) GetTagConfigSource() string {
	if o == nil || o.TagConfigSource == nil {
		var ret string
		return ret
	}
	return *o.TagConfigSource
}

// GetTagConfigSourceOk returns a tuple with the TagConfigSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyCostAttributionAttributes) GetTagConfigSourceOk() (*string, bool) {
	if o == nil || o.TagConfigSource == nil {
		return nil, false
	}
	return o.TagConfigSource, true
}

// HasTagConfigSource returns a boolean if a field has been set.
func (o *MonthlyCostAttributionAttributes) HasTagConfigSource() bool {
	return o != nil && o.TagConfigSource != nil
}

// SetTagConfigSource gets a reference to the given string and assigns it to the TagConfigSource field.
func (o *MonthlyCostAttributionAttributes) SetTagConfigSource(v string) {
	o.TagConfigSource = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonthlyCostAttributionAttributes) GetTags() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonthlyCostAttributionAttributes) GetTagsOk() (*map[string][]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MonthlyCostAttributionAttributes) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given map[string][]string and assigns it to the Tags field.
func (o *MonthlyCostAttributionAttributes) SetTags(v map[string][]string) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MonthlyCostAttributionAttributes) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyCostAttributionAttributes) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MonthlyCostAttributionAttributes) HasUpdatedAt() bool {
	return o != nil && o.UpdatedAt != nil
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *MonthlyCostAttributionAttributes) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *MonthlyCostAttributionAttributes) GetValues() interface{} {
	if o == nil || o.Values == nil {
		var ret interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonthlyCostAttributionAttributes) GetValuesOk() (*interface{}, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *MonthlyCostAttributionAttributes) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given interface{} and assigns it to the Values field.
func (o *MonthlyCostAttributionAttributes) SetValues(v interface{}) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonthlyCostAttributionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Month != nil {
		if o.Month.Nanosecond() == 0 {
			toSerialize["month"] = o.Month.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["month"] = o.Month.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.TagConfigSource != nil {
		toSerialize["tag_config_source"] = o.TagConfigSource
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonthlyCostAttributionAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Month           *time.Time          `json:"month,omitempty"`
		OrgName         *string             `json:"org_name,omitempty"`
		PublicId        *string             `json:"public_id,omitempty"`
		TagConfigSource *string             `json:"tag_config_source,omitempty"`
		Tags            map[string][]string `json:"tags,omitempty"`
		UpdatedAt       *string             `json:"updated_at,omitempty"`
		Values          interface{}         `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"month", "org_name", "public_id", "tag_config_source", "tags", "updated_at", "values"})
	} else {
		return err
	}
	o.Month = all.Month
	o.OrgName = all.OrgName
	o.PublicId = all.PublicId
	o.TagConfigSource = all.TagConfigSource
	o.Tags = all.Tags
	o.UpdatedAt = all.UpdatedAt
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
