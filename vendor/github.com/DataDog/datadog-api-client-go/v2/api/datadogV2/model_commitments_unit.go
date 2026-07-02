// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CommitmentsUnit Unit metadata for a numeric metric.
type CommitmentsUnit struct {
	// The unit family (for example, percentage or money).
	Family string `json:"family"`
	// The unit identifier.
	Id int64 `json:"id"`
	// The unit name (for example, percent or dollar).
	Name string `json:"name"`
	// The plural form of the unit name.
	Plural string `json:"plural"`
	// The scale factor for the unit.
	ScaleFactor float64 `json:"scale_factor"`
	// The abbreviated unit name (for example, % or $).
	ShortName string `json:"short_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCommitmentsUnit instantiates a new CommitmentsUnit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCommitmentsUnit(family string, id int64, name string, plural string, scaleFactor float64, shortName string) *CommitmentsUnit {
	this := CommitmentsUnit{}
	this.Family = family
	this.Id = id
	this.Name = name
	this.Plural = plural
	this.ScaleFactor = scaleFactor
	this.ShortName = shortName
	return &this
}

// NewCommitmentsUnitWithDefaults instantiates a new CommitmentsUnit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCommitmentsUnitWithDefaults() *CommitmentsUnit {
	this := CommitmentsUnit{}
	return &this
}

// GetFamily returns the Family field value.
func (o *CommitmentsUnit) GetFamily() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *CommitmentsUnit) GetFamilyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value.
func (o *CommitmentsUnit) SetFamily(v string) {
	o.Family = v
}

// GetId returns the Id field value.
func (o *CommitmentsUnit) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CommitmentsUnit) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *CommitmentsUnit) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value.
func (o *CommitmentsUnit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CommitmentsUnit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *CommitmentsUnit) SetName(v string) {
	o.Name = v
}

// GetPlural returns the Plural field value.
func (o *CommitmentsUnit) GetPlural() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Plural
}

// GetPluralOk returns a tuple with the Plural field value
// and a boolean to check if the value has been set.
func (o *CommitmentsUnit) GetPluralOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plural, true
}

// SetPlural sets field value.
func (o *CommitmentsUnit) SetPlural(v string) {
	o.Plural = v
}

// GetScaleFactor returns the ScaleFactor field value.
func (o *CommitmentsUnit) GetScaleFactor() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ScaleFactor
}

// GetScaleFactorOk returns a tuple with the ScaleFactor field value
// and a boolean to check if the value has been set.
func (o *CommitmentsUnit) GetScaleFactorOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScaleFactor, true
}

// SetScaleFactor sets field value.
func (o *CommitmentsUnit) SetScaleFactor(v float64) {
	o.ScaleFactor = v
}

// GetShortName returns the ShortName field value.
func (o *CommitmentsUnit) GetShortName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value
// and a boolean to check if the value has been set.
func (o *CommitmentsUnit) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortName, true
}

// SetShortName sets field value.
func (o *CommitmentsUnit) SetShortName(v string) {
	o.ShortName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CommitmentsUnit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["family"] = o.Family
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["plural"] = o.Plural
	toSerialize["scale_factor"] = o.ScaleFactor
	toSerialize["short_name"] = o.ShortName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CommitmentsUnit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Family      *string  `json:"family"`
		Id          *int64   `json:"id"`
		Name        *string  `json:"name"`
		Plural      *string  `json:"plural"`
		ScaleFactor *float64 `json:"scale_factor"`
		ShortName   *string  `json:"short_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Family == nil {
		return fmt.Errorf("required field family missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Plural == nil {
		return fmt.Errorf("required field plural missing")
	}
	if all.ScaleFactor == nil {
		return fmt.Errorf("required field scale_factor missing")
	}
	if all.ShortName == nil {
		return fmt.Errorf("required field short_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"family", "id", "name", "plural", "scale_factor", "short_name"})
	} else {
		return err
	}
	o.Family = *all.Family
	o.Id = *all.Id
	o.Name = *all.Name
	o.Plural = *all.Plural
	o.ScaleFactor = *all.ScaleFactor
	o.ShortName = *all.ShortName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
