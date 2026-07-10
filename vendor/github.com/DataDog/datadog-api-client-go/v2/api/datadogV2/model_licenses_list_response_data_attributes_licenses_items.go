// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LicensesListResponseDataAttributesLicensesItems An SPDX license entry returned by the licenses list endpoint.
type LicensesListResponseDataAttributesLicensesItems struct {
	// The human-readable name of the license.
	DisplayName string `json:"display_name"`
	// The SPDX identifier of the license.
	Identifier string `json:"identifier"`
	// The short name of the license, typically matching the SPDX identifier.
	ShortName string `json:"short_name"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLicensesListResponseDataAttributesLicensesItems instantiates a new LicensesListResponseDataAttributesLicensesItems object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLicensesListResponseDataAttributesLicensesItems(displayName string, identifier string, shortName string) *LicensesListResponseDataAttributesLicensesItems {
	this := LicensesListResponseDataAttributesLicensesItems{}
	this.DisplayName = displayName
	this.Identifier = identifier
	this.ShortName = shortName
	return &this
}

// NewLicensesListResponseDataAttributesLicensesItemsWithDefaults instantiates a new LicensesListResponseDataAttributesLicensesItems object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLicensesListResponseDataAttributesLicensesItemsWithDefaults() *LicensesListResponseDataAttributesLicensesItems {
	this := LicensesListResponseDataAttributesLicensesItems{}
	return &this
}

// GetDisplayName returns the DisplayName field value.
func (o *LicensesListResponseDataAttributesLicensesItems) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *LicensesListResponseDataAttributesLicensesItems) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value.
func (o *LicensesListResponseDataAttributesLicensesItems) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetIdentifier returns the Identifier field value.
func (o *LicensesListResponseDataAttributesLicensesItems) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *LicensesListResponseDataAttributesLicensesItems) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value.
func (o *LicensesListResponseDataAttributesLicensesItems) SetIdentifier(v string) {
	o.Identifier = v
}

// GetShortName returns the ShortName field value.
func (o *LicensesListResponseDataAttributesLicensesItems) GetShortName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ShortName
}

// GetShortNameOk returns a tuple with the ShortName field value
// and a boolean to check if the value has been set.
func (o *LicensesListResponseDataAttributesLicensesItems) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShortName, true
}

// SetShortName sets field value.
func (o *LicensesListResponseDataAttributesLicensesItems) SetShortName(v string) {
	o.ShortName = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LicensesListResponseDataAttributesLicensesItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["display_name"] = o.DisplayName
	toSerialize["identifier"] = o.Identifier
	toSerialize["short_name"] = o.ShortName

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LicensesListResponseDataAttributesLicensesItems) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DisplayName *string `json:"display_name"`
		Identifier  *string `json:"identifier"`
		ShortName   *string `json:"short_name"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DisplayName == nil {
		return fmt.Errorf("required field display_name missing")
	}
	if all.Identifier == nil {
		return fmt.Errorf("required field identifier missing")
	}
	if all.ShortName == nil {
		return fmt.Errorf("required field short_name missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"display_name", "identifier", "short_name"})
	} else {
		return err
	}
	o.DisplayName = *all.DisplayName
	o.Identifier = *all.Identifier
	o.ShortName = *all.ShortName

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
