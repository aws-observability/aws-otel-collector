// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultDnsRecord A DNS record returned in a DNS test response.
type SyntheticsTestResultDnsRecord struct {
	// DNS record type (for example, `A`, `AAAA`, `CNAME`).
	Type *string `json:"type,omitempty"`
	// Values associated with the DNS record.
	Values []string `json:"values,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultDnsRecord instantiates a new SyntheticsTestResultDnsRecord object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultDnsRecord() *SyntheticsTestResultDnsRecord {
	this := SyntheticsTestResultDnsRecord{}
	return &this
}

// NewSyntheticsTestResultDnsRecordWithDefaults instantiates a new SyntheticsTestResultDnsRecord object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultDnsRecordWithDefaults() *SyntheticsTestResultDnsRecord {
	this := SyntheticsTestResultDnsRecord{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SyntheticsTestResultDnsRecord) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDnsRecord) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SyntheticsTestResultDnsRecord) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SyntheticsTestResultDnsRecord) SetType(v string) {
	o.Type = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *SyntheticsTestResultDnsRecord) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDnsRecord) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *SyntheticsTestResultDnsRecord) HasValues() bool {
	return o != nil && o.Values != nil
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *SyntheticsTestResultDnsRecord) SetValues(v []string) {
	o.Values = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultDnsRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
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
func (o *SyntheticsTestResultDnsRecord) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Type   *string  `json:"type,omitempty"`
		Values []string `json:"values,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"type", "values"})
	} else {
		return err
	}
	o.Type = all.Type
	o.Values = all.Values

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
