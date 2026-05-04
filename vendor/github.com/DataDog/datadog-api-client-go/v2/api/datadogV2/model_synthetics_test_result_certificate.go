// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultCertificate SSL/TLS certificate information returned from an SSL test.
type SyntheticsTestResultCertificate struct {
	// Cipher used for the TLS connection.
	Cipher *string `json:"cipher,omitempty"`
	// RSA exponent of the certificate.
	Exponent *int64 `json:"exponent,omitempty"`
	// Extended key usage extensions for the certificate.
	ExtKeyUsage []string `json:"ext_key_usage,omitempty"`
	// SHA-1 fingerprint of the certificate.
	Fingerprint *string `json:"fingerprint,omitempty"`
	// SHA-256 fingerprint of the certificate.
	Fingerprint256 *string `json:"fingerprint256,omitempty"`
	// Certificate issuer details.
	Issuer map[string]string `json:"issuer,omitempty"`
	// RSA modulus of the certificate.
	Modulus *string `json:"modulus,omitempty"`
	// TLS protocol used (for example, `TLSv1.2`).
	Protocol *string `json:"protocol,omitempty"`
	// Serial number of the certificate.
	SerialNumber *string `json:"serial_number,omitempty"`
	// Certificate subject details.
	Subject map[string]string `json:"subject,omitempty"`
	// TLS protocol version.
	TlsVersion *float64 `json:"tls_version,omitempty"`
	// Validity window of a certificate.
	Valid *SyntheticsTestResultCertificateValidity `json:"valid,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultCertificate instantiates a new SyntheticsTestResultCertificate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultCertificate() *SyntheticsTestResultCertificate {
	this := SyntheticsTestResultCertificate{}
	return &this
}

// NewSyntheticsTestResultCertificateWithDefaults instantiates a new SyntheticsTestResultCertificate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultCertificateWithDefaults() *SyntheticsTestResultCertificate {
	this := SyntheticsTestResultCertificate{}
	return &this
}

// GetCipher returns the Cipher field value if set, zero value otherwise.
func (o *SyntheticsTestResultCertificate) GetCipher() string {
	if o == nil || o.Cipher == nil {
		var ret string
		return ret
	}
	return *o.Cipher
}

// GetCipherOk returns a tuple with the Cipher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCertificate) GetCipherOk() (*string, bool) {
	if o == nil || o.Cipher == nil {
		return nil, false
	}
	return o.Cipher, true
}

// HasCipher returns a boolean if a field has been set.
func (o *SyntheticsTestResultCertificate) HasCipher() bool {
	return o != nil && o.Cipher != nil
}

// SetCipher gets a reference to the given string and assigns it to the Cipher field.
func (o *SyntheticsTestResultCertificate) SetCipher(v string) {
	o.Cipher = &v
}

// GetExponent returns the Exponent field value if set, zero value otherwise.
func (o *SyntheticsTestResultCertificate) GetExponent() int64 {
	if o == nil || o.Exponent == nil {
		var ret int64
		return ret
	}
	return *o.Exponent
}

// GetExponentOk returns a tuple with the Exponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCertificate) GetExponentOk() (*int64, bool) {
	if o == nil || o.Exponent == nil {
		return nil, false
	}
	return o.Exponent, true
}

// HasExponent returns a boolean if a field has been set.
func (o *SyntheticsTestResultCertificate) HasExponent() bool {
	return o != nil && o.Exponent != nil
}

// SetExponent gets a reference to the given int64 and assigns it to the Exponent field.
func (o *SyntheticsTestResultCertificate) SetExponent(v int64) {
	o.Exponent = &v
}

// GetExtKeyUsage returns the ExtKeyUsage field value if set, zero value otherwise.
func (o *SyntheticsTestResultCertificate) GetExtKeyUsage() []string {
	if o == nil || o.ExtKeyUsage == nil {
		var ret []string
		return ret
	}
	return o.ExtKeyUsage
}

// GetExtKeyUsageOk returns a tuple with the ExtKeyUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCertificate) GetExtKeyUsageOk() (*[]string, bool) {
	if o == nil || o.ExtKeyUsage == nil {
		return nil, false
	}
	return &o.ExtKeyUsage, true
}

// HasExtKeyUsage returns a boolean if a field has been set.
func (o *SyntheticsTestResultCertificate) HasExtKeyUsage() bool {
	return o != nil && o.ExtKeyUsage != nil
}

// SetExtKeyUsage gets a reference to the given []string and assigns it to the ExtKeyUsage field.
func (o *SyntheticsTestResultCertificate) SetExtKeyUsage(v []string) {
	o.ExtKeyUsage = v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *SyntheticsTestResultCertificate) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCertificate) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *SyntheticsTestResultCertificate) HasFingerprint() bool {
	return o != nil && o.Fingerprint != nil
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *SyntheticsTestResultCertificate) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetFingerprint256 returns the Fingerprint256 field value if set, zero value otherwise.
func (o *SyntheticsTestResultCertificate) GetFingerprint256() string {
	if o == nil || o.Fingerprint256 == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint256
}

// GetFingerprint256Ok returns a tuple with the Fingerprint256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCertificate) GetFingerprint256Ok() (*string, bool) {
	if o == nil || o.Fingerprint256 == nil {
		return nil, false
	}
	return o.Fingerprint256, true
}

// HasFingerprint256 returns a boolean if a field has been set.
func (o *SyntheticsTestResultCertificate) HasFingerprint256() bool {
	return o != nil && o.Fingerprint256 != nil
}

// SetFingerprint256 gets a reference to the given string and assigns it to the Fingerprint256 field.
func (o *SyntheticsTestResultCertificate) SetFingerprint256(v string) {
	o.Fingerprint256 = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *SyntheticsTestResultCertificate) GetIssuer() map[string]string {
	if o == nil || o.Issuer == nil {
		var ret map[string]string
		return ret
	}
	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCertificate) GetIssuerOk() (*map[string]string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *SyntheticsTestResultCertificate) HasIssuer() bool {
	return o != nil && o.Issuer != nil
}

// SetIssuer gets a reference to the given map[string]string and assigns it to the Issuer field.
func (o *SyntheticsTestResultCertificate) SetIssuer(v map[string]string) {
	o.Issuer = v
}

// GetModulus returns the Modulus field value if set, zero value otherwise.
func (o *SyntheticsTestResultCertificate) GetModulus() string {
	if o == nil || o.Modulus == nil {
		var ret string
		return ret
	}
	return *o.Modulus
}

// GetModulusOk returns a tuple with the Modulus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCertificate) GetModulusOk() (*string, bool) {
	if o == nil || o.Modulus == nil {
		return nil, false
	}
	return o.Modulus, true
}

// HasModulus returns a boolean if a field has been set.
func (o *SyntheticsTestResultCertificate) HasModulus() bool {
	return o != nil && o.Modulus != nil
}

// SetModulus gets a reference to the given string and assigns it to the Modulus field.
func (o *SyntheticsTestResultCertificate) SetModulus(v string) {
	o.Modulus = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *SyntheticsTestResultCertificate) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCertificate) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *SyntheticsTestResultCertificate) HasProtocol() bool {
	return o != nil && o.Protocol != nil
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *SyntheticsTestResultCertificate) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *SyntheticsTestResultCertificate) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCertificate) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *SyntheticsTestResultCertificate) HasSerialNumber() bool {
	return o != nil && o.SerialNumber != nil
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *SyntheticsTestResultCertificate) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SyntheticsTestResultCertificate) GetSubject() map[string]string {
	if o == nil || o.Subject == nil {
		var ret map[string]string
		return ret
	}
	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCertificate) GetSubjectOk() (*map[string]string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return &o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SyntheticsTestResultCertificate) HasSubject() bool {
	return o != nil && o.Subject != nil
}

// SetSubject gets a reference to the given map[string]string and assigns it to the Subject field.
func (o *SyntheticsTestResultCertificate) SetSubject(v map[string]string) {
	o.Subject = v
}

// GetTlsVersion returns the TlsVersion field value if set, zero value otherwise.
func (o *SyntheticsTestResultCertificate) GetTlsVersion() float64 {
	if o == nil || o.TlsVersion == nil {
		var ret float64
		return ret
	}
	return *o.TlsVersion
}

// GetTlsVersionOk returns a tuple with the TlsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCertificate) GetTlsVersionOk() (*float64, bool) {
	if o == nil || o.TlsVersion == nil {
		return nil, false
	}
	return o.TlsVersion, true
}

// HasTlsVersion returns a boolean if a field has been set.
func (o *SyntheticsTestResultCertificate) HasTlsVersion() bool {
	return o != nil && o.TlsVersion != nil
}

// SetTlsVersion gets a reference to the given float64 and assigns it to the TlsVersion field.
func (o *SyntheticsTestResultCertificate) SetTlsVersion(v float64) {
	o.TlsVersion = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *SyntheticsTestResultCertificate) GetValid() SyntheticsTestResultCertificateValidity {
	if o == nil || o.Valid == nil {
		var ret SyntheticsTestResultCertificateValidity
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultCertificate) GetValidOk() (*SyntheticsTestResultCertificateValidity, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *SyntheticsTestResultCertificate) HasValid() bool {
	return o != nil && o.Valid != nil
}

// SetValid gets a reference to the given SyntheticsTestResultCertificateValidity and assigns it to the Valid field.
func (o *SyntheticsTestResultCertificate) SetValid(v SyntheticsTestResultCertificateValidity) {
	o.Valid = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cipher != nil {
		toSerialize["cipher"] = o.Cipher
	}
	if o.Exponent != nil {
		toSerialize["exponent"] = o.Exponent
	}
	if o.ExtKeyUsage != nil {
		toSerialize["ext_key_usage"] = o.ExtKeyUsage
	}
	if o.Fingerprint != nil {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if o.Fingerprint256 != nil {
		toSerialize["fingerprint256"] = o.Fingerprint256
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.Modulus != nil {
		toSerialize["modulus"] = o.Modulus
	}
	if o.Protocol != nil {
		toSerialize["protocol"] = o.Protocol
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.TlsVersion != nil {
		toSerialize["tls_version"] = o.TlsVersion
	}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultCertificate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cipher         *string                                  `json:"cipher,omitempty"`
		Exponent       *int64                                   `json:"exponent,omitempty"`
		ExtKeyUsage    []string                                 `json:"ext_key_usage,omitempty"`
		Fingerprint    *string                                  `json:"fingerprint,omitempty"`
		Fingerprint256 *string                                  `json:"fingerprint256,omitempty"`
		Issuer         map[string]string                        `json:"issuer,omitempty"`
		Modulus        *string                                  `json:"modulus,omitempty"`
		Protocol       *string                                  `json:"protocol,omitempty"`
		SerialNumber   *string                                  `json:"serial_number,omitempty"`
		Subject        map[string]string                        `json:"subject,omitempty"`
		TlsVersion     *float64                                 `json:"tls_version,omitempty"`
		Valid          *SyntheticsTestResultCertificateValidity `json:"valid,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cipher", "exponent", "ext_key_usage", "fingerprint", "fingerprint256", "issuer", "modulus", "protocol", "serial_number", "subject", "tls_version", "valid"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Cipher = all.Cipher
	o.Exponent = all.Exponent
	o.ExtKeyUsage = all.ExtKeyUsage
	o.Fingerprint = all.Fingerprint
	o.Fingerprint256 = all.Fingerprint256
	o.Issuer = all.Issuer
	o.Modulus = all.Modulus
	o.Protocol = all.Protocol
	o.SerialNumber = all.SerialNumber
	o.Subject = all.Subject
	o.TlsVersion = all.TlsVersion
	if all.Valid != nil && all.Valid.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Valid = all.Valid

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
