// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultOCSPResponse OCSP response received while validating a certificate.
type SyntheticsTestResultOCSPResponse struct {
	// Certificate details returned in an OCSP response.
	Certificate *SyntheticsTestResultOCSPCertificate `json:"certificate,omitempty"`
	// OCSP response status (for example, `good`, `revoked`, `unknown`).
	Status *string `json:"status,omitempty"`
	// OCSP response update timestamps.
	Updates *SyntheticsTestResultOCSPUpdates `json:"updates,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultOCSPResponse instantiates a new SyntheticsTestResultOCSPResponse object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultOCSPResponse() *SyntheticsTestResultOCSPResponse {
	this := SyntheticsTestResultOCSPResponse{}
	return &this
}

// NewSyntheticsTestResultOCSPResponseWithDefaults instantiates a new SyntheticsTestResultOCSPResponse object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultOCSPResponseWithDefaults() *SyntheticsTestResultOCSPResponse {
	this := SyntheticsTestResultOCSPResponse{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *SyntheticsTestResultOCSPResponse) GetCertificate() SyntheticsTestResultOCSPCertificate {
	if o == nil || o.Certificate == nil {
		var ret SyntheticsTestResultOCSPCertificate
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultOCSPResponse) GetCertificateOk() (*SyntheticsTestResultOCSPCertificate, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *SyntheticsTestResultOCSPResponse) HasCertificate() bool {
	return o != nil && o.Certificate != nil
}

// SetCertificate gets a reference to the given SyntheticsTestResultOCSPCertificate and assigns it to the Certificate field.
func (o *SyntheticsTestResultOCSPResponse) SetCertificate(v SyntheticsTestResultOCSPCertificate) {
	o.Certificate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsTestResultOCSPResponse) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultOCSPResponse) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsTestResultOCSPResponse) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SyntheticsTestResultOCSPResponse) SetStatus(v string) {
	o.Status = &v
}

// GetUpdates returns the Updates field value if set, zero value otherwise.
func (o *SyntheticsTestResultOCSPResponse) GetUpdates() SyntheticsTestResultOCSPUpdates {
	if o == nil || o.Updates == nil {
		var ret SyntheticsTestResultOCSPUpdates
		return ret
	}
	return *o.Updates
}

// GetUpdatesOk returns a tuple with the Updates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultOCSPResponse) GetUpdatesOk() (*SyntheticsTestResultOCSPUpdates, bool) {
	if o == nil || o.Updates == nil {
		return nil, false
	}
	return o.Updates, true
}

// HasUpdates returns a boolean if a field has been set.
func (o *SyntheticsTestResultOCSPResponse) HasUpdates() bool {
	return o != nil && o.Updates != nil
}

// SetUpdates gets a reference to the given SyntheticsTestResultOCSPUpdates and assigns it to the Updates field.
func (o *SyntheticsTestResultOCSPResponse) SetUpdates(v SyntheticsTestResultOCSPUpdates) {
	o.Updates = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultOCSPResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Updates != nil {
		toSerialize["updates"] = o.Updates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultOCSPResponse) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Certificate *SyntheticsTestResultOCSPCertificate `json:"certificate,omitempty"`
		Status      *string                              `json:"status,omitempty"`
		Updates     *SyntheticsTestResultOCSPUpdates     `json:"updates,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"certificate", "status", "updates"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Certificate != nil && all.Certificate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Certificate = all.Certificate
	o.Status = all.Status
	if all.Updates != nil && all.Updates.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Updates = all.Updates

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
