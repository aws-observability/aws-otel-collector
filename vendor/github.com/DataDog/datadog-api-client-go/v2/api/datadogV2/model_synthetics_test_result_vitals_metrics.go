// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultVitalsMetrics Web vitals metrics captured during a browser test step.
type SyntheticsTestResultVitalsMetrics struct {
	// Cumulative Layout Shift score.
	Cls *float64 `json:"cls,omitempty"`
	// First Contentful Paint in milliseconds.
	Fcp *float64 `json:"fcp,omitempty"`
	// Interaction to Next Paint in milliseconds.
	Inp *float64 `json:"inp,omitempty"`
	// Largest Contentful Paint in milliseconds.
	Lcp *float64 `json:"lcp,omitempty"`
	// Time To First Byte in milliseconds.
	Ttfb *float64 `json:"ttfb,omitempty"`
	// URL that produced the metrics.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultVitalsMetrics instantiates a new SyntheticsTestResultVitalsMetrics object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultVitalsMetrics() *SyntheticsTestResultVitalsMetrics {
	this := SyntheticsTestResultVitalsMetrics{}
	return &this
}

// NewSyntheticsTestResultVitalsMetricsWithDefaults instantiates a new SyntheticsTestResultVitalsMetrics object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultVitalsMetricsWithDefaults() *SyntheticsTestResultVitalsMetrics {
	this := SyntheticsTestResultVitalsMetrics{}
	return &this
}

// GetCls returns the Cls field value if set, zero value otherwise.
func (o *SyntheticsTestResultVitalsMetrics) GetCls() float64 {
	if o == nil || o.Cls == nil {
		var ret float64
		return ret
	}
	return *o.Cls
}

// GetClsOk returns a tuple with the Cls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVitalsMetrics) GetClsOk() (*float64, bool) {
	if o == nil || o.Cls == nil {
		return nil, false
	}
	return o.Cls, true
}

// HasCls returns a boolean if a field has been set.
func (o *SyntheticsTestResultVitalsMetrics) HasCls() bool {
	return o != nil && o.Cls != nil
}

// SetCls gets a reference to the given float64 and assigns it to the Cls field.
func (o *SyntheticsTestResultVitalsMetrics) SetCls(v float64) {
	o.Cls = &v
}

// GetFcp returns the Fcp field value if set, zero value otherwise.
func (o *SyntheticsTestResultVitalsMetrics) GetFcp() float64 {
	if o == nil || o.Fcp == nil {
		var ret float64
		return ret
	}
	return *o.Fcp
}

// GetFcpOk returns a tuple with the Fcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVitalsMetrics) GetFcpOk() (*float64, bool) {
	if o == nil || o.Fcp == nil {
		return nil, false
	}
	return o.Fcp, true
}

// HasFcp returns a boolean if a field has been set.
func (o *SyntheticsTestResultVitalsMetrics) HasFcp() bool {
	return o != nil && o.Fcp != nil
}

// SetFcp gets a reference to the given float64 and assigns it to the Fcp field.
func (o *SyntheticsTestResultVitalsMetrics) SetFcp(v float64) {
	o.Fcp = &v
}

// GetInp returns the Inp field value if set, zero value otherwise.
func (o *SyntheticsTestResultVitalsMetrics) GetInp() float64 {
	if o == nil || o.Inp == nil {
		var ret float64
		return ret
	}
	return *o.Inp
}

// GetInpOk returns a tuple with the Inp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVitalsMetrics) GetInpOk() (*float64, bool) {
	if o == nil || o.Inp == nil {
		return nil, false
	}
	return o.Inp, true
}

// HasInp returns a boolean if a field has been set.
func (o *SyntheticsTestResultVitalsMetrics) HasInp() bool {
	return o != nil && o.Inp != nil
}

// SetInp gets a reference to the given float64 and assigns it to the Inp field.
func (o *SyntheticsTestResultVitalsMetrics) SetInp(v float64) {
	o.Inp = &v
}

// GetLcp returns the Lcp field value if set, zero value otherwise.
func (o *SyntheticsTestResultVitalsMetrics) GetLcp() float64 {
	if o == nil || o.Lcp == nil {
		var ret float64
		return ret
	}
	return *o.Lcp
}

// GetLcpOk returns a tuple with the Lcp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVitalsMetrics) GetLcpOk() (*float64, bool) {
	if o == nil || o.Lcp == nil {
		return nil, false
	}
	return o.Lcp, true
}

// HasLcp returns a boolean if a field has been set.
func (o *SyntheticsTestResultVitalsMetrics) HasLcp() bool {
	return o != nil && o.Lcp != nil
}

// SetLcp gets a reference to the given float64 and assigns it to the Lcp field.
func (o *SyntheticsTestResultVitalsMetrics) SetLcp(v float64) {
	o.Lcp = &v
}

// GetTtfb returns the Ttfb field value if set, zero value otherwise.
func (o *SyntheticsTestResultVitalsMetrics) GetTtfb() float64 {
	if o == nil || o.Ttfb == nil {
		var ret float64
		return ret
	}
	return *o.Ttfb
}

// GetTtfbOk returns a tuple with the Ttfb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVitalsMetrics) GetTtfbOk() (*float64, bool) {
	if o == nil || o.Ttfb == nil {
		return nil, false
	}
	return o.Ttfb, true
}

// HasTtfb returns a boolean if a field has been set.
func (o *SyntheticsTestResultVitalsMetrics) HasTtfb() bool {
	return o != nil && o.Ttfb != nil
}

// SetTtfb gets a reference to the given float64 and assigns it to the Ttfb field.
func (o *SyntheticsTestResultVitalsMetrics) SetTtfb(v float64) {
	o.Ttfb = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SyntheticsTestResultVitalsMetrics) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultVitalsMetrics) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SyntheticsTestResultVitalsMetrics) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SyntheticsTestResultVitalsMetrics) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultVitalsMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Cls != nil {
		toSerialize["cls"] = o.Cls
	}
	if o.Fcp != nil {
		toSerialize["fcp"] = o.Fcp
	}
	if o.Inp != nil {
		toSerialize["inp"] = o.Inp
	}
	if o.Lcp != nil {
		toSerialize["lcp"] = o.Lcp
	}
	if o.Ttfb != nil {
		toSerialize["ttfb"] = o.Ttfb
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultVitalsMetrics) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cls  *float64 `json:"cls,omitempty"`
		Fcp  *float64 `json:"fcp,omitempty"`
		Inp  *float64 `json:"inp,omitempty"`
		Lcp  *float64 `json:"lcp,omitempty"`
		Ttfb *float64 `json:"ttfb,omitempty"`
		Url  *string  `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cls", "fcp", "inp", "lcp", "ttfb", "url"})
	} else {
		return err
	}
	o.Cls = all.Cls
	o.Fcp = all.Fcp
	o.Inp = all.Inp
	o.Lcp = all.Lcp
	o.Ttfb = all.Ttfb
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
