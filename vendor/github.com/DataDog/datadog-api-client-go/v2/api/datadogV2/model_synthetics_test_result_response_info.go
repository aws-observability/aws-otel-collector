// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultResponseInfo Details of the response received during the test execution.
type SyntheticsTestResultResponseInfo struct {
	// Body of the response.
	Body *string `json:"body,omitempty"`
	// Compressed representation of the response body.
	BodyCompressed *string `json:"body_compressed,omitempty"`
	// Hashes computed over the response body.
	BodyHashes *string `json:"body_hashes,omitempty"`
	// Size of the response body in bytes.
	BodySize *int64 `json:"body_size,omitempty"`
	// Cache-related response headers.
	CacheHeaders map[string]string `json:"cache_headers,omitempty"`
	// CDN provider details inferred from response headers.
	Cdn *SyntheticsTestResultCdnProviderInfo `json:"cdn,omitempty"`
	// WebSocket close frame information for WebSocket test responses.
	Close *SyntheticsTestResultWebSocketClose `json:"close,omitempty"`
	// Compressed representation of the response message.
	CompressedMessage *string `json:"compressed_message,omitempty"`
	// Response headers.
	Headers map[string]interface{} `json:"headers,omitempty"`
	// Health check information returned from a gRPC health check call.
	Healthcheck *SyntheticsTestResultHealthCheck `json:"healthcheck,omitempty"`
	// HTTP version of the response.
	HttpVersion *string `json:"http_version,omitempty"`
	// Whether the response body was truncated.
	IsBodyTruncated *bool `json:"is_body_truncated,omitempty"`
	// Whether the response message was truncated.
	IsMessageTruncated *bool `json:"is_message_truncated,omitempty"`
	// Message received in the response (for WebSocket/TCP/UDP tests).
	Message *string `json:"message,omitempty"`
	// Additional metadata returned with the response.
	Metadata map[string]string `json:"metadata,omitempty"`
	// DNS records returned in the response (DNS tests only).
	Records []SyntheticsTestResultDnsRecord `json:"records,omitempty"`
	// Redirect hops encountered while performing the request.
	Redirects []SyntheticsTestResultRedirect `json:"redirects,omitempty"`
	// HTTP status code of the response.
	StatusCode *int64 `json:"status_code,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultResponseInfo instantiates a new SyntheticsTestResultResponseInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultResponseInfo() *SyntheticsTestResultResponseInfo {
	this := SyntheticsTestResultResponseInfo{}
	return &this
}

// NewSyntheticsTestResultResponseInfoWithDefaults instantiates a new SyntheticsTestResultResponseInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultResponseInfoWithDefaults() *SyntheticsTestResultResponseInfo {
	this := SyntheticsTestResultResponseInfo{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasBody() bool {
	return o != nil && o.Body != nil
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *SyntheticsTestResultResponseInfo) SetBody(v string) {
	o.Body = &v
}

// GetBodyCompressed returns the BodyCompressed field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetBodyCompressed() string {
	if o == nil || o.BodyCompressed == nil {
		var ret string
		return ret
	}
	return *o.BodyCompressed
}

// GetBodyCompressedOk returns a tuple with the BodyCompressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetBodyCompressedOk() (*string, bool) {
	if o == nil || o.BodyCompressed == nil {
		return nil, false
	}
	return o.BodyCompressed, true
}

// HasBodyCompressed returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasBodyCompressed() bool {
	return o != nil && o.BodyCompressed != nil
}

// SetBodyCompressed gets a reference to the given string and assigns it to the BodyCompressed field.
func (o *SyntheticsTestResultResponseInfo) SetBodyCompressed(v string) {
	o.BodyCompressed = &v
}

// GetBodyHashes returns the BodyHashes field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetBodyHashes() string {
	if o == nil || o.BodyHashes == nil {
		var ret string
		return ret
	}
	return *o.BodyHashes
}

// GetBodyHashesOk returns a tuple with the BodyHashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetBodyHashesOk() (*string, bool) {
	if o == nil || o.BodyHashes == nil {
		return nil, false
	}
	return o.BodyHashes, true
}

// HasBodyHashes returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasBodyHashes() bool {
	return o != nil && o.BodyHashes != nil
}

// SetBodyHashes gets a reference to the given string and assigns it to the BodyHashes field.
func (o *SyntheticsTestResultResponseInfo) SetBodyHashes(v string) {
	o.BodyHashes = &v
}

// GetBodySize returns the BodySize field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetBodySize() int64 {
	if o == nil || o.BodySize == nil {
		var ret int64
		return ret
	}
	return *o.BodySize
}

// GetBodySizeOk returns a tuple with the BodySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetBodySizeOk() (*int64, bool) {
	if o == nil || o.BodySize == nil {
		return nil, false
	}
	return o.BodySize, true
}

// HasBodySize returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasBodySize() bool {
	return o != nil && o.BodySize != nil
}

// SetBodySize gets a reference to the given int64 and assigns it to the BodySize field.
func (o *SyntheticsTestResultResponseInfo) SetBodySize(v int64) {
	o.BodySize = &v
}

// GetCacheHeaders returns the CacheHeaders field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetCacheHeaders() map[string]string {
	if o == nil || o.CacheHeaders == nil {
		var ret map[string]string
		return ret
	}
	return o.CacheHeaders
}

// GetCacheHeadersOk returns a tuple with the CacheHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetCacheHeadersOk() (*map[string]string, bool) {
	if o == nil || o.CacheHeaders == nil {
		return nil, false
	}
	return &o.CacheHeaders, true
}

// HasCacheHeaders returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasCacheHeaders() bool {
	return o != nil && o.CacheHeaders != nil
}

// SetCacheHeaders gets a reference to the given map[string]string and assigns it to the CacheHeaders field.
func (o *SyntheticsTestResultResponseInfo) SetCacheHeaders(v map[string]string) {
	o.CacheHeaders = v
}

// GetCdn returns the Cdn field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetCdn() SyntheticsTestResultCdnProviderInfo {
	if o == nil || o.Cdn == nil {
		var ret SyntheticsTestResultCdnProviderInfo
		return ret
	}
	return *o.Cdn
}

// GetCdnOk returns a tuple with the Cdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetCdnOk() (*SyntheticsTestResultCdnProviderInfo, bool) {
	if o == nil || o.Cdn == nil {
		return nil, false
	}
	return o.Cdn, true
}

// HasCdn returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasCdn() bool {
	return o != nil && o.Cdn != nil
}

// SetCdn gets a reference to the given SyntheticsTestResultCdnProviderInfo and assigns it to the Cdn field.
func (o *SyntheticsTestResultResponseInfo) SetCdn(v SyntheticsTestResultCdnProviderInfo) {
	o.Cdn = &v
}

// GetClose returns the Close field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetClose() SyntheticsTestResultWebSocketClose {
	if o == nil || o.Close == nil {
		var ret SyntheticsTestResultWebSocketClose
		return ret
	}
	return *o.Close
}

// GetCloseOk returns a tuple with the Close field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetCloseOk() (*SyntheticsTestResultWebSocketClose, bool) {
	if o == nil || o.Close == nil {
		return nil, false
	}
	return o.Close, true
}

// HasClose returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasClose() bool {
	return o != nil && o.Close != nil
}

// SetClose gets a reference to the given SyntheticsTestResultWebSocketClose and assigns it to the Close field.
func (o *SyntheticsTestResultResponseInfo) SetClose(v SyntheticsTestResultWebSocketClose) {
	o.Close = &v
}

// GetCompressedMessage returns the CompressedMessage field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetCompressedMessage() string {
	if o == nil || o.CompressedMessage == nil {
		var ret string
		return ret
	}
	return *o.CompressedMessage
}

// GetCompressedMessageOk returns a tuple with the CompressedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetCompressedMessageOk() (*string, bool) {
	if o == nil || o.CompressedMessage == nil {
		return nil, false
	}
	return o.CompressedMessage, true
}

// HasCompressedMessage returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasCompressedMessage() bool {
	return o != nil && o.CompressedMessage != nil
}

// SetCompressedMessage gets a reference to the given string and assigns it to the CompressedMessage field.
func (o *SyntheticsTestResultResponseInfo) SetCompressedMessage(v string) {
	o.CompressedMessage = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetHeaders() map[string]interface{} {
	if o == nil || o.Headers == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetHeadersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasHeaders() bool {
	return o != nil && o.Headers != nil
}

// SetHeaders gets a reference to the given map[string]interface{} and assigns it to the Headers field.
func (o *SyntheticsTestResultResponseInfo) SetHeaders(v map[string]interface{}) {
	o.Headers = v
}

// GetHealthcheck returns the Healthcheck field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetHealthcheck() SyntheticsTestResultHealthCheck {
	if o == nil || o.Healthcheck == nil {
		var ret SyntheticsTestResultHealthCheck
		return ret
	}
	return *o.Healthcheck
}

// GetHealthcheckOk returns a tuple with the Healthcheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetHealthcheckOk() (*SyntheticsTestResultHealthCheck, bool) {
	if o == nil || o.Healthcheck == nil {
		return nil, false
	}
	return o.Healthcheck, true
}

// HasHealthcheck returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasHealthcheck() bool {
	return o != nil && o.Healthcheck != nil
}

// SetHealthcheck gets a reference to the given SyntheticsTestResultHealthCheck and assigns it to the Healthcheck field.
func (o *SyntheticsTestResultResponseInfo) SetHealthcheck(v SyntheticsTestResultHealthCheck) {
	o.Healthcheck = &v
}

// GetHttpVersion returns the HttpVersion field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetHttpVersion() string {
	if o == nil || o.HttpVersion == nil {
		var ret string
		return ret
	}
	return *o.HttpVersion
}

// GetHttpVersionOk returns a tuple with the HttpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetHttpVersionOk() (*string, bool) {
	if o == nil || o.HttpVersion == nil {
		return nil, false
	}
	return o.HttpVersion, true
}

// HasHttpVersion returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasHttpVersion() bool {
	return o != nil && o.HttpVersion != nil
}

// SetHttpVersion gets a reference to the given string and assigns it to the HttpVersion field.
func (o *SyntheticsTestResultResponseInfo) SetHttpVersion(v string) {
	o.HttpVersion = &v
}

// GetIsBodyTruncated returns the IsBodyTruncated field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetIsBodyTruncated() bool {
	if o == nil || o.IsBodyTruncated == nil {
		var ret bool
		return ret
	}
	return *o.IsBodyTruncated
}

// GetIsBodyTruncatedOk returns a tuple with the IsBodyTruncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetIsBodyTruncatedOk() (*bool, bool) {
	if o == nil || o.IsBodyTruncated == nil {
		return nil, false
	}
	return o.IsBodyTruncated, true
}

// HasIsBodyTruncated returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasIsBodyTruncated() bool {
	return o != nil && o.IsBodyTruncated != nil
}

// SetIsBodyTruncated gets a reference to the given bool and assigns it to the IsBodyTruncated field.
func (o *SyntheticsTestResultResponseInfo) SetIsBodyTruncated(v bool) {
	o.IsBodyTruncated = &v
}

// GetIsMessageTruncated returns the IsMessageTruncated field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetIsMessageTruncated() bool {
	if o == nil || o.IsMessageTruncated == nil {
		var ret bool
		return ret
	}
	return *o.IsMessageTruncated
}

// GetIsMessageTruncatedOk returns a tuple with the IsMessageTruncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetIsMessageTruncatedOk() (*bool, bool) {
	if o == nil || o.IsMessageTruncated == nil {
		return nil, false
	}
	return o.IsMessageTruncated, true
}

// HasIsMessageTruncated returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasIsMessageTruncated() bool {
	return o != nil && o.IsMessageTruncated != nil
}

// SetIsMessageTruncated gets a reference to the given bool and assigns it to the IsMessageTruncated field.
func (o *SyntheticsTestResultResponseInfo) SetIsMessageTruncated(v bool) {
	o.IsMessageTruncated = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SyntheticsTestResultResponseInfo) SetMessage(v string) {
	o.Message = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *SyntheticsTestResultResponseInfo) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetRecords() []SyntheticsTestResultDnsRecord {
	if o == nil || o.Records == nil {
		var ret []SyntheticsTestResultDnsRecord
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetRecordsOk() (*[]SyntheticsTestResultDnsRecord, bool) {
	if o == nil || o.Records == nil {
		return nil, false
	}
	return &o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasRecords() bool {
	return o != nil && o.Records != nil
}

// SetRecords gets a reference to the given []SyntheticsTestResultDnsRecord and assigns it to the Records field.
func (o *SyntheticsTestResultResponseInfo) SetRecords(v []SyntheticsTestResultDnsRecord) {
	o.Records = v
}

// GetRedirects returns the Redirects field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetRedirects() []SyntheticsTestResultRedirect {
	if o == nil || o.Redirects == nil {
		var ret []SyntheticsTestResultRedirect
		return ret
	}
	return o.Redirects
}

// GetRedirectsOk returns a tuple with the Redirects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetRedirectsOk() (*[]SyntheticsTestResultRedirect, bool) {
	if o == nil || o.Redirects == nil {
		return nil, false
	}
	return &o.Redirects, true
}

// HasRedirects returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasRedirects() bool {
	return o != nil && o.Redirects != nil
}

// SetRedirects gets a reference to the given []SyntheticsTestResultRedirect and assigns it to the Redirects field.
func (o *SyntheticsTestResultResponseInfo) SetRedirects(v []SyntheticsTestResultRedirect) {
	o.Redirects = v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *SyntheticsTestResultResponseInfo) GetStatusCode() int64 {
	if o == nil || o.StatusCode == nil {
		var ret int64
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultResponseInfo) GetStatusCodeOk() (*int64, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *SyntheticsTestResultResponseInfo) HasStatusCode() bool {
	return o != nil && o.StatusCode != nil
}

// SetStatusCode gets a reference to the given int64 and assigns it to the StatusCode field.
func (o *SyntheticsTestResultResponseInfo) SetStatusCode(v int64) {
	o.StatusCode = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultResponseInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.BodyCompressed != nil {
		toSerialize["body_compressed"] = o.BodyCompressed
	}
	if o.BodyHashes != nil {
		toSerialize["body_hashes"] = o.BodyHashes
	}
	if o.BodySize != nil {
		toSerialize["body_size"] = o.BodySize
	}
	if o.CacheHeaders != nil {
		toSerialize["cache_headers"] = o.CacheHeaders
	}
	if o.Cdn != nil {
		toSerialize["cdn"] = o.Cdn
	}
	if o.Close != nil {
		toSerialize["close"] = o.Close
	}
	if o.CompressedMessage != nil {
		toSerialize["compressed_message"] = o.CompressedMessage
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Healthcheck != nil {
		toSerialize["healthcheck"] = o.Healthcheck
	}
	if o.HttpVersion != nil {
		toSerialize["http_version"] = o.HttpVersion
	}
	if o.IsBodyTruncated != nil {
		toSerialize["is_body_truncated"] = o.IsBodyTruncated
	}
	if o.IsMessageTruncated != nil {
		toSerialize["is_message_truncated"] = o.IsMessageTruncated
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Records != nil {
		toSerialize["records"] = o.Records
	}
	if o.Redirects != nil {
		toSerialize["redirects"] = o.Redirects
	}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultResponseInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Body               *string                              `json:"body,omitempty"`
		BodyCompressed     *string                              `json:"body_compressed,omitempty"`
		BodyHashes         *string                              `json:"body_hashes,omitempty"`
		BodySize           *int64                               `json:"body_size,omitempty"`
		CacheHeaders       map[string]string                    `json:"cache_headers,omitempty"`
		Cdn                *SyntheticsTestResultCdnProviderInfo `json:"cdn,omitempty"`
		Close              *SyntheticsTestResultWebSocketClose  `json:"close,omitempty"`
		CompressedMessage  *string                              `json:"compressed_message,omitempty"`
		Headers            map[string]interface{}               `json:"headers,omitempty"`
		Healthcheck        *SyntheticsTestResultHealthCheck     `json:"healthcheck,omitempty"`
		HttpVersion        *string                              `json:"http_version,omitempty"`
		IsBodyTruncated    *bool                                `json:"is_body_truncated,omitempty"`
		IsMessageTruncated *bool                                `json:"is_message_truncated,omitempty"`
		Message            *string                              `json:"message,omitempty"`
		Metadata           map[string]string                    `json:"metadata,omitempty"`
		Records            []SyntheticsTestResultDnsRecord      `json:"records,omitempty"`
		Redirects          []SyntheticsTestResultRedirect       `json:"redirects,omitempty"`
		StatusCode         *int64                               `json:"status_code,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"body", "body_compressed", "body_hashes", "body_size", "cache_headers", "cdn", "close", "compressed_message", "headers", "healthcheck", "http_version", "is_body_truncated", "is_message_truncated", "message", "metadata", "records", "redirects", "status_code"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Body = all.Body
	o.BodyCompressed = all.BodyCompressed
	o.BodyHashes = all.BodyHashes
	o.BodySize = all.BodySize
	o.CacheHeaders = all.CacheHeaders
	if all.Cdn != nil && all.Cdn.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cdn = all.Cdn
	if all.Close != nil && all.Close.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Close = all.Close
	o.CompressedMessage = all.CompressedMessage
	o.Headers = all.Headers
	if all.Healthcheck != nil && all.Healthcheck.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Healthcheck = all.Healthcheck
	o.HttpVersion = all.HttpVersion
	o.IsBodyTruncated = all.IsBodyTruncated
	o.IsMessageTruncated = all.IsMessageTruncated
	o.Message = all.Message
	o.Metadata = all.Metadata
	o.Records = all.Records
	o.Redirects = all.Redirects
	o.StatusCode = all.StatusCode

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
