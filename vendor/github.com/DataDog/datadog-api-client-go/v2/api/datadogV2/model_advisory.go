// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Advisory Advisory.
type Advisory struct {
	// Advisory base severity.
	BaseSeverity string `json:"base_severity"`
	// Advisory id.
	Id string `json:"id"`
	// Advisory Datadog severity.
	Severity *string `json:"severity,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAdvisory instantiates a new Advisory object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAdvisory(baseSeverity string, id string) *Advisory {
	this := Advisory{}
	this.BaseSeverity = baseSeverity
	this.Id = id
	return &this
}

// NewAdvisoryWithDefaults instantiates a new Advisory object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAdvisoryWithDefaults() *Advisory {
	this := Advisory{}
	return &this
}

// GetBaseSeverity returns the BaseSeverity field value.
func (o *Advisory) GetBaseSeverity() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BaseSeverity
}

// GetBaseSeverityOk returns a tuple with the BaseSeverity field value
// and a boolean to check if the value has been set.
func (o *Advisory) GetBaseSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseSeverity, true
}

// SetBaseSeverity sets field value.
func (o *Advisory) SetBaseSeverity(v string) {
	o.BaseSeverity = v
}

// GetId returns the Id field value.
func (o *Advisory) GetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Advisory) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value.
func (o *Advisory) SetId(v string) {
	o.Id = v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *Advisory) GetSeverity() string {
	if o == nil || o.Severity == nil {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Advisory) GetSeverityOk() (*string, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *Advisory) HasSeverity() bool {
	return o != nil && o.Severity != nil
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *Advisory) SetSeverity(v string) {
	o.Severity = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o Advisory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["base_severity"] = o.BaseSeverity
	toSerialize["id"] = o.Id
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Advisory) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BaseSeverity *string `json:"base_severity"`
		Id           *string `json:"id"`
		Severity     *string `json:"severity,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BaseSeverity == nil {
		return fmt.Errorf("required field base_severity missing")
	}
	if all.Id == nil {
		return fmt.Errorf("required field id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"base_severity", "id", "severity"})
	} else {
		return err
	}
	o.BaseSeverity = *all.BaseSeverity
	o.Id = *all.Id
	o.Severity = all.Severity

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
