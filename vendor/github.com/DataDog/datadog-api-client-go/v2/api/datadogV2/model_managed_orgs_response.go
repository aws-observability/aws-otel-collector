// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ManagedOrgsResponse Response containing the current organization and its managed organizations.
type ManagedOrgsResponse struct {
	// The managed organizations resource.
	Data ManagedOrgsData `json:"data"`
	// Included organization resources.
	Included []OrgData `json:"included"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewManagedOrgsResponse instantiates a new ManagedOrgsResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewManagedOrgsResponse(data ManagedOrgsData, included []OrgData) *ManagedOrgsResponse {
	this := ManagedOrgsResponse{}
	this.Data = data
	this.Included = included
	return &this
}

// NewManagedOrgsResponseWithDefaults instantiates a new ManagedOrgsResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewManagedOrgsResponseWithDefaults() *ManagedOrgsResponse {
	this := ManagedOrgsResponse{}
	return &this
}

// GetData returns the Data field value.
func (o *ManagedOrgsResponse) GetData() ManagedOrgsData {
	if o == nil {
		var ret ManagedOrgsData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ManagedOrgsResponse) GetDataOk() (*ManagedOrgsData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value.
func (o *ManagedOrgsResponse) SetData(v ManagedOrgsData) {
	o.Data = v
}

// GetIncluded returns the Included field value.
func (o *ManagedOrgsResponse) GetIncluded() []OrgData {
	if o == nil {
		var ret []OrgData
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value
// and a boolean to check if the value has been set.
func (o *ManagedOrgsResponse) GetIncludedOk() (*[]OrgData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Included, true
}

// SetIncluded sets field value.
func (o *ManagedOrgsResponse) SetIncluded(v []OrgData) {
	o.Included = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ManagedOrgsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["data"] = o.Data
	toSerialize["included"] = o.Included

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ManagedOrgsResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Data     *ManagedOrgsData `json:"data"`
		Included *[]OrgData       `json:"included"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Data == nil {
		return fmt.Errorf("required field data missing")
	}
	if all.Included == nil {
		return fmt.Errorf("required field included missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"data", "included"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Data.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Data = *all.Data
	o.Included = *all.Included

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
