// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedResourceTimingBreakdown Average timing breakdown per network phase for a resource.
type AggregatedResourceTimingBreakdown struct {
	// Average TCP connect duration in milliseconds.
	AvgConnectMs float64 `json:"avg_connect_ms"`
	// Average DNS resolution duration in milliseconds.
	AvgDnsMs float64 `json:"avg_dns_ms"`
	// Average download phase duration in milliseconds.
	AvgDownloadMs float64 `json:"avg_download_ms"`
	// Average time to first byte in milliseconds.
	AvgFirstByteMs float64 `json:"avg_first_byte_ms"`
	// Average redirect phase duration in milliseconds.
	AvgRedirectMs float64 `json:"avg_redirect_ms"`
	// Average SSL handshake duration in milliseconds.
	AvgSslMs float64 `json:"avg_ssl_ms"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedResourceTimingBreakdown instantiates a new AggregatedResourceTimingBreakdown object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedResourceTimingBreakdown(avgConnectMs float64, avgDnsMs float64, avgDownloadMs float64, avgFirstByteMs float64, avgRedirectMs float64, avgSslMs float64) *AggregatedResourceTimingBreakdown {
	this := AggregatedResourceTimingBreakdown{}
	this.AvgConnectMs = avgConnectMs
	this.AvgDnsMs = avgDnsMs
	this.AvgDownloadMs = avgDownloadMs
	this.AvgFirstByteMs = avgFirstByteMs
	this.AvgRedirectMs = avgRedirectMs
	this.AvgSslMs = avgSslMs
	return &this
}

// NewAggregatedResourceTimingBreakdownWithDefaults instantiates a new AggregatedResourceTimingBreakdown object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedResourceTimingBreakdownWithDefaults() *AggregatedResourceTimingBreakdown {
	this := AggregatedResourceTimingBreakdown{}
	return &this
}

// GetAvgConnectMs returns the AvgConnectMs field value.
func (o *AggregatedResourceTimingBreakdown) GetAvgConnectMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AvgConnectMs
}

// GetAvgConnectMsOk returns a tuple with the AvgConnectMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedResourceTimingBreakdown) GetAvgConnectMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgConnectMs, true
}

// SetAvgConnectMs sets field value.
func (o *AggregatedResourceTimingBreakdown) SetAvgConnectMs(v float64) {
	o.AvgConnectMs = v
}

// GetAvgDnsMs returns the AvgDnsMs field value.
func (o *AggregatedResourceTimingBreakdown) GetAvgDnsMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AvgDnsMs
}

// GetAvgDnsMsOk returns a tuple with the AvgDnsMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedResourceTimingBreakdown) GetAvgDnsMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgDnsMs, true
}

// SetAvgDnsMs sets field value.
func (o *AggregatedResourceTimingBreakdown) SetAvgDnsMs(v float64) {
	o.AvgDnsMs = v
}

// GetAvgDownloadMs returns the AvgDownloadMs field value.
func (o *AggregatedResourceTimingBreakdown) GetAvgDownloadMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AvgDownloadMs
}

// GetAvgDownloadMsOk returns a tuple with the AvgDownloadMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedResourceTimingBreakdown) GetAvgDownloadMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgDownloadMs, true
}

// SetAvgDownloadMs sets field value.
func (o *AggregatedResourceTimingBreakdown) SetAvgDownloadMs(v float64) {
	o.AvgDownloadMs = v
}

// GetAvgFirstByteMs returns the AvgFirstByteMs field value.
func (o *AggregatedResourceTimingBreakdown) GetAvgFirstByteMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AvgFirstByteMs
}

// GetAvgFirstByteMsOk returns a tuple with the AvgFirstByteMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedResourceTimingBreakdown) GetAvgFirstByteMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgFirstByteMs, true
}

// SetAvgFirstByteMs sets field value.
func (o *AggregatedResourceTimingBreakdown) SetAvgFirstByteMs(v float64) {
	o.AvgFirstByteMs = v
}

// GetAvgRedirectMs returns the AvgRedirectMs field value.
func (o *AggregatedResourceTimingBreakdown) GetAvgRedirectMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AvgRedirectMs
}

// GetAvgRedirectMsOk returns a tuple with the AvgRedirectMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedResourceTimingBreakdown) GetAvgRedirectMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgRedirectMs, true
}

// SetAvgRedirectMs sets field value.
func (o *AggregatedResourceTimingBreakdown) SetAvgRedirectMs(v float64) {
	o.AvgRedirectMs = v
}

// GetAvgSslMs returns the AvgSslMs field value.
func (o *AggregatedResourceTimingBreakdown) GetAvgSslMs() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.AvgSslMs
}

// GetAvgSslMsOk returns a tuple with the AvgSslMs field value
// and a boolean to check if the value has been set.
func (o *AggregatedResourceTimingBreakdown) GetAvgSslMsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgSslMs, true
}

// SetAvgSslMs sets field value.
func (o *AggregatedResourceTimingBreakdown) SetAvgSslMs(v float64) {
	o.AvgSslMs = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedResourceTimingBreakdown) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["avg_connect_ms"] = o.AvgConnectMs
	toSerialize["avg_dns_ms"] = o.AvgDnsMs
	toSerialize["avg_download_ms"] = o.AvgDownloadMs
	toSerialize["avg_first_byte_ms"] = o.AvgFirstByteMs
	toSerialize["avg_redirect_ms"] = o.AvgRedirectMs
	toSerialize["avg_ssl_ms"] = o.AvgSslMs

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedResourceTimingBreakdown) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvgConnectMs   *float64 `json:"avg_connect_ms"`
		AvgDnsMs       *float64 `json:"avg_dns_ms"`
		AvgDownloadMs  *float64 `json:"avg_download_ms"`
		AvgFirstByteMs *float64 `json:"avg_first_byte_ms"`
		AvgRedirectMs  *float64 `json:"avg_redirect_ms"`
		AvgSslMs       *float64 `json:"avg_ssl_ms"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AvgConnectMs == nil {
		return fmt.Errorf("required field avg_connect_ms missing")
	}
	if all.AvgDnsMs == nil {
		return fmt.Errorf("required field avg_dns_ms missing")
	}
	if all.AvgDownloadMs == nil {
		return fmt.Errorf("required field avg_download_ms missing")
	}
	if all.AvgFirstByteMs == nil {
		return fmt.Errorf("required field avg_first_byte_ms missing")
	}
	if all.AvgRedirectMs == nil {
		return fmt.Errorf("required field avg_redirect_ms missing")
	}
	if all.AvgSslMs == nil {
		return fmt.Errorf("required field avg_ssl_ms missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"avg_connect_ms", "avg_dns_ms", "avg_download_ms", "avg_first_byte_ms", "avg_redirect_ms", "avg_ssl_ms"})
	} else {
		return err
	}
	o.AvgConnectMs = *all.AvgConnectMs
	o.AvgDnsMs = *all.AvgDnsMs
	o.AvgDownloadMs = *all.AvgDownloadMs
	o.AvgFirstByteMs = *all.AvgFirstByteMs
	o.AvgRedirectMs = *all.AvgRedirectMs
	o.AvgSslMs = *all.AvgSslMs

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
