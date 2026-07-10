// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// BudgetAttributesCostsUnit The unit used for all cost values in the response.
type BudgetAttributesCostsUnit struct {
	// The unit family (for example, `currency`).
	Family *string `json:"family,omitempty"`
	// The unique identifier for the unit.
	Id *string `json:"id,omitempty"`
	// The full name of the unit.
	Name *string `json:"name,omitempty"`
	// The plural form of the unit name.
	Plural *string `json:"plural,omitempty"`
	// The scale factor applied to raw cost values.
	ScaleFactor *float64 `json:"scale_factor,omitempty"`
	// The abbreviated unit name.
	ShortName *string `json:"short_name,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewBudgetAttributesCostsUnit instantiates a new BudgetAttributesCostsUnit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewBudgetAttributesCostsUnit() *BudgetAttributesCostsUnit {
	this := BudgetAttributesCostsUnit{}
	return &this
}

// NewBudgetAttributesCostsUnitWithDefaults instantiates a new BudgetAttributesCostsUnit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewBudgetAttributesCostsUnitWithDefaults() *BudgetAttributesCostsUnit {
	this := BudgetAttributesCostsUnit{}
	return &this
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *BudgetAttributesCostsUnit) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributesCostsUnit) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *BudgetAttributesCostsUnit) HasFamily() bool {
	return o != nil && o.Family != nil
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *BudgetAttributesCostsUnit) SetFamily(v string) {
	o.Family = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BudgetAttributesCostsUnit) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributesCostsUnit) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BudgetAttributesCostsUnit) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BudgetAttributesCostsUnit) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BudgetAttributesCostsUnit) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributesCostsUnit) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BudgetAttributesCostsUnit) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BudgetAttributesCostsUnit) SetName(v string) {
	o.Name = &v
}

// GetPlural returns the Plural field value if set, zero value otherwise.
func (o *BudgetAttributesCostsUnit) GetPlural() string {
	if o == nil || o.Plural == nil {
		var ret string
		return ret
	}
	return *o.Plural
}

// GetPluralOk returns a tuple with the Plural field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributesCostsUnit) GetPluralOk() (*string, bool) {
	if o == nil || o.Plural == nil {
		return nil, false
	}
	return o.Plural, true
}

// HasPlural returns a boolean if a field has been set.
func (o *BudgetAttributesCostsUnit) HasPlural() bool {
	return o != nil && o.Plural != nil
}

// SetPlural gets a reference to the given string and assigns it to the Plural field.
func (o *BudgetAttributesCostsUnit) SetPlural(v string) {
	o.Plural = &v
}

// GetScaleFactor returns the ScaleFactor field value if set, zero value otherwise.
func (o *BudgetAttributesCostsUnit) GetScaleFactor() float64 {
	if o == nil || o.ScaleFactor == nil {
		var ret float64
		return ret
	}
	return *o.ScaleFactor
}

// GetScaleFactorOk returns a tuple with the ScaleFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributesCostsUnit) GetScaleFactorOk() (*float64, bool) {
	if o == nil || o.ScaleFactor == nil {
		return nil, false
	}
	return o.ScaleFactor, true
}

// HasScaleFactor returns a boolean if a field has been set.
func (o *BudgetAttributesCostsUnit) HasScaleFactor() bool {
	return o != nil && o.ScaleFactor != nil
}

// SetScaleFactor gets a reference to the given float64 and assigns it to the ScaleFactor field.
func (o *BudgetAttributesCostsUnit) SetScaleFactor(v float64) {
	o.ScaleFactor = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise.
func (o *BudgetAttributesCostsUnit) GetShortName() string {
	if o == nil || o.ShortName == nil {
		var ret string
		return ret
	}
	return *o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BudgetAttributesCostsUnit) GetShortNameOk() (*string, bool) {
	if o == nil || o.ShortName == nil {
		return nil, false
	}
	return o.ShortName, true
}

// HasShortName returns a boolean if a field has been set.
func (o *BudgetAttributesCostsUnit) HasShortName() bool {
	return o != nil && o.ShortName != nil
}

// SetShortName gets a reference to the given string and assigns it to the ShortName field.
func (o *BudgetAttributesCostsUnit) SetShortName(v string) {
	o.ShortName = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o BudgetAttributesCostsUnit) MarshalJSON() ([]byte, error) {
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
func (o *BudgetAttributesCostsUnit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Family      *string  `json:"family,omitempty"`
		Id          *string  `json:"id,omitempty"`
		Name        *string  `json:"name,omitempty"`
		Plural      *string  `json:"plural,omitempty"`
		ScaleFactor *float64 `json:"scale_factor,omitempty"`
		ShortName   *string  `json:"short_name,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
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
