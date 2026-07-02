// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// OpsgenieAccountResponseAttributes The attributes from an Opsgenie account response.
type OpsgenieAccountResponseAttributes struct {
	// The region for the Opsgenie service.
	Region *OpsgenieServiceRegionType `json:"region,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewOpsgenieAccountResponseAttributes instantiates a new OpsgenieAccountResponseAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewOpsgenieAccountResponseAttributes() *OpsgenieAccountResponseAttributes {
	this := OpsgenieAccountResponseAttributes{}
	return &this
}

// NewOpsgenieAccountResponseAttributesWithDefaults instantiates a new OpsgenieAccountResponseAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewOpsgenieAccountResponseAttributesWithDefaults() *OpsgenieAccountResponseAttributes {
	this := OpsgenieAccountResponseAttributes{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *OpsgenieAccountResponseAttributes) GetRegion() OpsgenieServiceRegionType {
	if o == nil || o.Region == nil {
		var ret OpsgenieServiceRegionType
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsgenieAccountResponseAttributes) GetRegionOk() (*OpsgenieServiceRegionType, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *OpsgenieAccountResponseAttributes) HasRegion() bool {
	return o != nil && o.Region != nil
}

// SetRegion gets a reference to the given OpsgenieServiceRegionType and assigns it to the Region field.
func (o *OpsgenieAccountResponseAttributes) SetRegion(v OpsgenieServiceRegionType) {
	o.Region = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o OpsgenieAccountResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *OpsgenieAccountResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Region *OpsgenieServiceRegionType `json:"region,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"region"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Region != nil && !all.Region.IsValid() {
		hasInvalidField = true
	} else {
		o.Region = all.Region
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
