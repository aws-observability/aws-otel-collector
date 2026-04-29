// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// GetInvestigationResponseLinks Links related to the investigation.
type GetInvestigationResponseLinks struct {
	// The URL to the investigation in the Datadog app.
	Self string `json:"self"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewGetInvestigationResponseLinks instantiates a new GetInvestigationResponseLinks object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewGetInvestigationResponseLinks(self string) *GetInvestigationResponseLinks {
	this := GetInvestigationResponseLinks{}
	this.Self = self
	return &this
}

// NewGetInvestigationResponseLinksWithDefaults instantiates a new GetInvestigationResponseLinks object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewGetInvestigationResponseLinksWithDefaults() *GetInvestigationResponseLinks {
	this := GetInvestigationResponseLinks{}
	return &this
}

// GetSelf returns the Self field value.
func (o *GetInvestigationResponseLinks) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *GetInvestigationResponseLinks) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value.
func (o *GetInvestigationResponseLinks) SetSelf(v string) {
	o.Self = v
}

// MarshalJSON serializes the struct using spec logic.
func (o GetInvestigationResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["self"] = o.Self

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *GetInvestigationResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Self *string `json:"self"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Self == nil {
		return fmt.Errorf("required field self missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"self"})
	} else {
		return err
	}
	o.Self = *all.Self

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
