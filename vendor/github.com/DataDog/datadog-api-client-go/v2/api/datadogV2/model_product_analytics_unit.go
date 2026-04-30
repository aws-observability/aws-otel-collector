// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ProductAnalyticsUnit A unit definition for metric values.
type ProductAnalyticsUnit struct {
	// The unit family (e.g., time, bytes).
	Family *string `json:"family,omitempty"`
	// Numeric identifier for the unit.
	Id *int64 `json:"id,omitempty"`
	// The full name of the unit (e.g., nanosecond).
	Name *string `json:"name,omitempty"`
	// Plural form of the unit name (e.g., nanoseconds).
	Plural *string `json:"plural,omitempty"`
	// Conversion factor relative to the base unit of the family.
	ScaleFactor *float64 `json:"scale_factor,omitempty"`
	// Abbreviated unit name (e.g., ns).
	ShortName *string `json:"short_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewProductAnalyticsUnit instantiates a new ProductAnalyticsUnit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewProductAnalyticsUnit() *ProductAnalyticsUnit {
	this := ProductAnalyticsUnit{}
	return &this
}

// NewProductAnalyticsUnitWithDefaults instantiates a new ProductAnalyticsUnit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewProductAnalyticsUnitWithDefaults() *ProductAnalyticsUnit {
	this := ProductAnalyticsUnit{}
	return &this
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *ProductAnalyticsUnit) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsUnit) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *ProductAnalyticsUnit) HasFamily() bool {
	return o != nil && o.Family != nil
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *ProductAnalyticsUnit) SetFamily(v string) {
	o.Family = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProductAnalyticsUnit) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsUnit) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProductAnalyticsUnit) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ProductAnalyticsUnit) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProductAnalyticsUnit) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsUnit) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProductAnalyticsUnit) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProductAnalyticsUnit) SetName(v string) {
	o.Name = &v
}

// GetPlural returns the Plural field value if set, zero value otherwise.
func (o *ProductAnalyticsUnit) GetPlural() string {
	if o == nil || o.Plural == nil {
		var ret string
		return ret
	}
	return *o.Plural
}

// GetPluralOk returns a tuple with the Plural field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsUnit) GetPluralOk() (*string, bool) {
	if o == nil || o.Plural == nil {
		return nil, false
	}
	return o.Plural, true
}

// HasPlural returns a boolean if a field has been set.
func (o *ProductAnalyticsUnit) HasPlural() bool {
	return o != nil && o.Plural != nil
}

// SetPlural gets a reference to the given string and assigns it to the Plural field.
func (o *ProductAnalyticsUnit) SetPlural(v string) {
	o.Plural = &v
}

// GetScaleFactor returns the ScaleFactor field value if set, zero value otherwise.
func (o *ProductAnalyticsUnit) GetScaleFactor() float64 {
	if o == nil || o.ScaleFactor == nil {
		var ret float64
		return ret
	}
	return *o.ScaleFactor
}

// GetScaleFactorOk returns a tuple with the ScaleFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsUnit) GetScaleFactorOk() (*float64, bool) {
	if o == nil || o.ScaleFactor == nil {
		return nil, false
	}
	return o.ScaleFactor, true
}

// HasScaleFactor returns a boolean if a field has been set.
func (o *ProductAnalyticsUnit) HasScaleFactor() bool {
	return o != nil && o.ScaleFactor != nil
}

// SetScaleFactor gets a reference to the given float64 and assigns it to the ScaleFactor field.
func (o *ProductAnalyticsUnit) SetScaleFactor(v float64) {
	o.ScaleFactor = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *ProductAnalyticsUnit) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAnalyticsUnit) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *ProductAnalyticsUnit) HasShortName() bool {
	return o != nil && o.ShortName != nil
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *ProductAnalyticsUnit) SetShortName(v string) {
	o.ShortName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ProductAnalyticsUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Plural != nil {
		toSerialize["plural"] = o.Plural
	}
	if o.ScaleFactor != nil {
		toSerialize["scale_factor"] = o.ScaleFactor
	}
	if o.ShortName != nil {
		toSerialize["short_name"] = o.ShortName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ProductAnalyticsUnit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Family      *string  `json:"family,omitempty"`
		Id          *int64   `json:"id,omitempty"`
		Name        *string  `json:"name,omitempty"`
		Plural      *string  `json:"plural,omitempty"`
		ScaleFactor *float64 `json:"scale_factor,omitempty"`
		ShortName   *string  `json:"short_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"family", "id", "name", "plural", "scale_factor", "short_name"})
	} else {
		return err
	}
	o.Family = all.Family
	o.Id = all.Id
	o.Name = all.Name
	o.Plural = all.Plural
	o.ScaleFactor = all.ScaleFactor
	o.ShortName = all.ShortName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
