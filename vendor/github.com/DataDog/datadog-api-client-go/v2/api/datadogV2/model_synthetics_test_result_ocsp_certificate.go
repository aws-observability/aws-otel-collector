// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultOCSPCertificate Certificate details returned in an OCSP response.
type SyntheticsTestResultOCSPCertificate struct {
	// Reason code for the revocation, when applicable.
	RevocationReason *string `json:"revocation_reason,omitempty"`
	// Unix timestamp (ms) of the revocation.
	RevocationTime *int64 `json:"revocation_time,omitempty"`
	// Serial number of the certificate.
	SerialNumber *string `json:"serial_number,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultOCSPCertificate instantiates a new SyntheticsTestResultOCSPCertificate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultOCSPCertificate() *SyntheticsTestResultOCSPCertificate {
	this := SyntheticsTestResultOCSPCertificate{}
	return &this
}

// NewSyntheticsTestResultOCSPCertificateWithDefaults instantiates a new SyntheticsTestResultOCSPCertificate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultOCSPCertificateWithDefaults() *SyntheticsTestResultOCSPCertificate {
	this := SyntheticsTestResultOCSPCertificate{}
	return &this
}

// GetRevocationReason returns the RevocationReason field value if set, zero value otherwise.
func (o *SyntheticsTestResultOCSPCertificate) GetRevocationReason() string {
	if o == nil || o.RevocationReason == nil {
		var ret string
		return ret
	}
	return *o.RevocationReason
}

// GetRevocationReasonOk returns a tuple with the RevocationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultOCSPCertificate) GetRevocationReasonOk() (*string, bool) {
	if o == nil || o.RevocationReason == nil {
		return nil, false
	}
	return o.RevocationReason, true
}

// HasRevocationReason returns a boolean if a field has been set.
func (o *SyntheticsTestResultOCSPCertificate) HasRevocationReason() bool {
	return o != nil && o.RevocationReason != nil
}

// SetRevocationReason gets a reference to the given string and assigns it to the RevocationReason field.
func (o *SyntheticsTestResultOCSPCertificate) SetRevocationReason(v string) {
	o.RevocationReason = &v
}

// GetRevocationTime returns the RevocationTime field value if set, zero value otherwise.
func (o *SyntheticsTestResultOCSPCertificate) GetRevocationTime() int64 {
	if o == nil || o.RevocationTime == nil {
		var ret int64
		return ret
	}
	return *o.RevocationTime
}

// GetRevocationTimeOk returns a tuple with the RevocationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultOCSPCertificate) GetRevocationTimeOk() (*int64, bool) {
	if o == nil || o.RevocationTime == nil {
		return nil, false
	}
	return o.RevocationTime, true
}

// HasRevocationTime returns a boolean if a field has been set.
func (o *SyntheticsTestResultOCSPCertificate) HasRevocationTime() bool {
	return o != nil && o.RevocationTime != nil
}

// SetRevocationTime gets a reference to the given int64 and assigns it to the RevocationTime field.
func (o *SyntheticsTestResultOCSPCertificate) SetRevocationTime(v int64) {
	o.RevocationTime = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *SyntheticsTestResultOCSPCertificate) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultOCSPCertificate) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *SyntheticsTestResultOCSPCertificate) HasSerialNumber() bool {
	return o != nil && o.SerialNumber != nil
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *SyntheticsTestResultOCSPCertificate) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultOCSPCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.RevocationReason != nil {
		toSerialize["revocation_reason"] = o.RevocationReason
	}
	if o.RevocationTime != nil {
		toSerialize["revocation_time"] = o.RevocationTime
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultOCSPCertificate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RevocationReason *string `json:"revocation_reason,omitempty"`
		RevocationTime   *int64  `json:"revocation_time,omitempty"`
		SerialNumber     *string `json:"serial_number,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"revocation_reason", "revocation_time", "serial_number"})
	} else {
		return err
	}
	o.RevocationReason = all.RevocationReason
	o.RevocationTime = all.RevocationTime
	o.SerialNumber = all.SerialNumber

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
