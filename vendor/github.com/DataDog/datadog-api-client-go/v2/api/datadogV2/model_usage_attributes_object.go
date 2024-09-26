// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// UsageAttributesObject Usage attributes data.
type UsageAttributesObject struct {
	// The organization name.
	OrgName *string `json:"org_name,omitempty"`
	// The product for which usage is being reported.
	ProductFamily *string `json:"product_family,omitempty"`
	// The organization public ID.
	PublicId *string `json:"public_id,omitempty"`
	// The region of the Datadog instance that the organization belongs to.
	Region *string `json:"region,omitempty"`
	// List of usage data reported for each requested hour.
	Timeseries []UsageTimeSeriesObject `json:"timeseries,omitempty"`
	// Usage type that is being measured.
	UsageType *HourlyUsageType `json:"usage_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewUsageAttributesObject instantiates a new UsageAttributesObject object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewUsageAttributesObject() *UsageAttributesObject {
	this := UsageAttributesObject{}
	return &this
}

// NewUsageAttributesObjectWithDefaults instantiates a new UsageAttributesObject object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewUsageAttributesObjectWithDefaults() *UsageAttributesObject {
	this := UsageAttributesObject{}
	return &this
}

// GetOrgName returns the OrgName field value if set, zero value otherwise.
func (o *UsageAttributesObject) GetOrgName() string {
	if o == nil || o.OrgName == nil {
		var ret string
		return ret
	}
	return *o.OrgName
}

// GetOrgNameOk returns a tuple with the OrgName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAttributesObject) GetOrgNameOk() (*string, bool) {
	if o == nil || o.OrgName == nil {
		return nil, false
	}
	return o.OrgName, true
}

// HasOrgName returns a boolean if a field has been set.
func (o *UsageAttributesObject) HasOrgName() bool {
	return o != nil && o.OrgName != nil
}

// SetOrgName gets a reference to the given string and assigns it to the OrgName field.
func (o *UsageAttributesObject) SetOrgName(v string) {
	o.OrgName = &v
}

// GetProductFamily returns the ProductFamily field value if set, zero value otherwise.
func (o *UsageAttributesObject) GetProductFamily() string {
	if o == nil || o.ProductFamily == nil {
		var ret string
		return ret
	}
	return *o.ProductFamily
}

// GetProductFamilyOk returns a tuple with the ProductFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAttributesObject) GetProductFamilyOk() (*string, bool) {
	if o == nil || o.ProductFamily == nil {
		return nil, false
	}
	return o.ProductFamily, true
}

// HasProductFamily returns a boolean if a field has been set.
func (o *UsageAttributesObject) HasProductFamily() bool {
	return o != nil && o.ProductFamily != nil
}

// SetProductFamily gets a reference to the given string and assigns it to the ProductFamily field.
func (o *UsageAttributesObject) SetProductFamily(v string) {
	o.ProductFamily = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *UsageAttributesObject) GetPublicId() string {
	if o == nil || o.PublicId == nil {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAttributesObject) GetPublicIdOk() (*string, bool) {
	if o == nil || o.PublicId == nil {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *UsageAttributesObject) HasPublicId() bool {
	return o != nil && o.PublicId != nil
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *UsageAttributesObject) SetPublicId(v string) {
	o.PublicId = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *UsageAttributesObject) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAttributesObject) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *UsageAttributesObject) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *UsageAttributesObject) SetRegion(v string) {
	o.Region = &v
}

// GetTimeseries returns the Timeseries field value if set, zero value otherwise.
func (o *UsageAttributesObject) GetTimeseries() []UsageTimeSeriesObject {
	if o == nil || o.Timeseries == nil {
		var ret []UsageTimeSeriesObject
		return ret
	}
	return o.Timeseries
}

// GetTimeseriesOk returns a tuple with the Timeseries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAttributesObject) GetTimeseriesOk() (*[]UsageTimeSeriesObject, bool) {
	if o == nil || o.Timeseries == nil {
		return nil, false
	}
	return &o.Timeseries, true
}

// HasTimeseries returns a boolean if a field has been set.
func (o *UsageAttributesObject) HasTimeseries() bool {
	return o != nil && o.Timeseries != nil
}

// SetTimeseries gets a reference to the given []UsageTimeSeriesObject and assigns it to the Timeseries field.
func (o *UsageAttributesObject) SetTimeseries(v []UsageTimeSeriesObject) {
	o.Timeseries = v
}

// GetUsageType returns the UsageType field value if set, zero value otherwise.
func (o *UsageAttributesObject) GetUsageType() HourlyUsageType {
	if o == nil || o.UsageType == nil {
		var ret HourlyUsageType
		return ret
	}
	return *o.UsageType
}

// GetUsageTypeOk returns a tuple with the UsageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageAttributesObject) GetUsageTypeOk() (*HourlyUsageType, bool) {
	if o == nil || o.UsageType == nil {
		return nil, false
	}
	return o.UsageType, true
}

// HasUsageType returns a boolean if a field has been set.
func (o *UsageAttributesObject) HasUsageType() bool {
	return o != nil && o.UsageType != nil
}

// SetUsageType gets a reference to the given HourlyUsageType and assigns it to the UsageType field.
func (o *UsageAttributesObject) SetUsageType(v HourlyUsageType) {
	o.UsageType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o UsageAttributesObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.OrgName != nil {
		toSerialize["org_name"] = o.OrgName
	}
	if o.ProductFamily != nil {
		toSerialize["product_family"] = o.ProductFamily
	}
	if o.PublicId != nil {
		toSerialize["public_id"] = o.PublicId
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Timeseries != nil {
		toSerialize["timeseries"] = o.Timeseries
	}
	if o.UsageType != nil {
		toSerialize["usage_type"] = o.UsageType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *UsageAttributesObject) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		OrgName       *string                 `json:"org_name,omitempty"`
		ProductFamily *string                 `json:"product_family,omitempty"`
		PublicId      *string                 `json:"public_id,omitempty"`
		Region        *string                 `json:"region,omitempty"`
		Timeseries    []UsageTimeSeriesObject `json:"timeseries,omitempty"`
		UsageType     *HourlyUsageType        `json:"usage_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"org_name", "product_family", "public_id", "region", "timeseries", "usage_type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.OrgName = all.OrgName
	o.ProductFamily = all.ProductFamily
	o.PublicId = all.PublicId
	o.Region = all.Region
	o.Timeseries = all.Timeseries
	if all.UsageType != nil && !all.UsageType.IsValid() {
		hasInvalidField = true
	} else {
		o.UsageType = all.UsageType
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
