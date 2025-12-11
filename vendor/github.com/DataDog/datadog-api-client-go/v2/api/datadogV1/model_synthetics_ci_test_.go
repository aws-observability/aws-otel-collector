// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsCITest Configuration for Continuous Testing.
type SyntheticsCITest struct {
	// Disable certificate checks in API tests.
	AllowInsecureCertificates *bool `json:"allowInsecureCertificates,omitempty"`
	// Object to handle basic authentication when performing the test.
	BasicAuth *SyntheticsBasicAuth `json:"basicAuth,omitempty"`
	// Body to include in the test.
	Body *string `json:"body,omitempty"`
	// Type of the data sent in a Synthetic API test.
	BodyType *string `json:"bodyType,omitempty"`
	// Cookies for the request.
	Cookies *string `json:"cookies,omitempty"`
	// For browser test, array with the different device IDs used to run the test.
	DeviceIds []string `json:"deviceIds,omitempty"`
	// For API HTTP test, whether or not the test should follow redirects.
	FollowRedirects *bool `json:"followRedirects,omitempty"`
	// Headers to include when performing the test.
	Headers map[string]string `json:"headers,omitempty"`
	// Array of locations used to run the test.
	Locations []string `json:"locations,omitempty"`
	// Metadata for the Synthetic tests run.
	Metadata *SyntheticsCIBatchMetadata `json:"metadata,omitempty"`
	// The public ID of the Synthetic test to trigger.
	PublicId string `json:"public_id"`
	// Object describing the retry strategy to apply to a Synthetic test.
	Retry *SyntheticsTestOptionsRetry `json:"retry,omitempty"`
	// Starting URL for the browser test.
	StartUrl *string `json:"startUrl,omitempty"`
	// Variables to replace in the test.
	Variables map[string]string `json:"variables,omitempty"`
	// The version number of the Synthetic test version to trigger.
	Version *int64 `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsCITest instantiates a new SyntheticsCITest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsCITest(publicId string) *SyntheticsCITest {
	this := SyntheticsCITest{}
	this.PublicId = publicId
	return &this
}

// NewSyntheticsCITestWithDefaults instantiates a new SyntheticsCITest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsCITestWithDefaults() *SyntheticsCITest {
	this := SyntheticsCITest{}
	return &this
}

// GetAllowInsecureCertificates returns the AllowInsecureCertificates field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetAllowInsecureCertificates() bool {
	if o == nil || o.AllowInsecureCertificates == nil {
		var ret bool
		return ret
	}
	return *o.AllowInsecureCertificates
}

// GetAllowInsecureCertificatesOk returns a tuple with the AllowInsecureCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetAllowInsecureCertificatesOk() (*bool, bool) {
	if o == nil || o.AllowInsecureCertificates == nil {
		return nil, false
	}
	return o.AllowInsecureCertificates, true
}

// HasAllowInsecureCertificates returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasAllowInsecureCertificates() bool {
	return o != nil && o.AllowInsecureCertificates != nil
}

// SetAllowInsecureCertificates gets a reference to the given bool and assigns it to the AllowInsecureCertificates field.
func (o *SyntheticsCITest) SetAllowInsecureCertificates(v bool) {
	o.AllowInsecureCertificates = &v
}

// GetBasicAuth returns the BasicAuth field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetBasicAuth() SyntheticsBasicAuth {
	if o == nil || o.BasicAuth == nil {
		var ret SyntheticsBasicAuth
		return ret
	}
	return *o.BasicAuth
}

// GetBasicAuthOk returns a tuple with the BasicAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetBasicAuthOk() (*SyntheticsBasicAuth, bool) {
	if o == nil || o.BasicAuth == nil {
		return nil, false
	}
	return o.BasicAuth, true
}

// HasBasicAuth returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasBasicAuth() bool {
	return o != nil && o.BasicAuth != nil
}

// SetBasicAuth gets a reference to the given SyntheticsBasicAuth and assigns it to the BasicAuth field.
func (o *SyntheticsCITest) SetBasicAuth(v SyntheticsBasicAuth) {
	o.BasicAuth = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasBody() bool {
	return o != nil && o.Body != nil
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *SyntheticsCITest) SetBody(v string) {
	o.Body = &v
}

// GetBodyType returns the BodyType field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetBodyType() string {
	if o == nil || o.BodyType == nil {
		var ret string
		return ret
	}
	return *o.BodyType
}

// GetBodyTypeOk returns a tuple with the BodyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetBodyTypeOk() (*string, bool) {
	if o == nil || o.BodyType == nil {
		return nil, false
	}
	return o.BodyType, true
}

// HasBodyType returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasBodyType() bool {
	return o != nil && o.BodyType != nil
}

// SetBodyType gets a reference to the given string and assigns it to the BodyType field.
func (o *SyntheticsCITest) SetBodyType(v string) {
	o.BodyType = &v
}

// GetCookies returns the Cookies field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetCookies() string {
	if o == nil || o.Cookies == nil {
		var ret string
		return ret
	}
	return *o.Cookies
}

// GetCookiesOk returns a tuple with the Cookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetCookiesOk() (*string, bool) {
	if o == nil || o.Cookies == nil {
		return nil, false
	}
	return o.Cookies, true
}

// HasCookies returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasCookies() bool {
	return o != nil && o.Cookies != nil
}

// SetCookies gets a reference to the given string and assigns it to the Cookies field.
func (o *SyntheticsCITest) SetCookies(v string) {
	o.Cookies = &v
}

// GetDeviceIds returns the DeviceIds field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetDeviceIds() []string {
	if o == nil || o.DeviceIds == nil {
		var ret []string
		return ret
	}
	return o.DeviceIds
}

// GetDeviceIdsOk returns a tuple with the DeviceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetDeviceIdsOk() (*[]string, bool) {
	if o == nil || o.DeviceIds == nil {
		return nil, false
	}
	return &o.DeviceIds, true
}

// HasDeviceIds returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasDeviceIds() bool {
	return o != nil && o.DeviceIds != nil
}

// SetDeviceIds gets a reference to the given []string and assigns it to the DeviceIds field.
func (o *SyntheticsCITest) SetDeviceIds(v []string) {
	o.DeviceIds = v
}

// GetFollowRedirects returns the FollowRedirects field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetFollowRedirects() bool {
	if o == nil || o.FollowRedirects == nil {
		var ret bool
		return ret
	}
	return *o.FollowRedirects
}

// GetFollowRedirectsOk returns a tuple with the FollowRedirects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetFollowRedirectsOk() (*bool, bool) {
	if o == nil || o.FollowRedirects == nil {
		return nil, false
	}
	return o.FollowRedirects, true
}

// HasFollowRedirects returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasFollowRedirects() bool {
	return o != nil && o.FollowRedirects != nil
}

// SetFollowRedirects gets a reference to the given bool and assigns it to the FollowRedirects field.
func (o *SyntheticsCITest) SetFollowRedirects(v bool) {
	o.FollowRedirects = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetHeaders() map[string]string {
	if o == nil || o.Headers == nil {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasHeaders() bool {
	return o != nil && o.Headers != nil
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *SyntheticsCITest) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetLocations returns the Locations field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetLocations() []string {
	if o == nil || o.Locations == nil {
		var ret []string
		return ret
	}
	return o.Locations
}

// GetLocationsOk returns a tuple with the Locations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetLocationsOk() (*[]string, bool) {
	if o == nil || o.Locations == nil {
		return nil, false
	}
	return &o.Locations, true
}

// HasLocations returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasLocations() bool {
	return o != nil && o.Locations != nil
}

// SetLocations gets a reference to the given []string and assigns it to the Locations field.
func (o *SyntheticsCITest) SetLocations(v []string) {
	o.Locations = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetMetadata() SyntheticsCIBatchMetadata {
	if o == nil || o.Metadata == nil {
		var ret SyntheticsCIBatchMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetMetadataOk() (*SyntheticsCIBatchMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given SyntheticsCIBatchMetadata and assigns it to the Metadata field.
func (o *SyntheticsCITest) SetMetadata(v SyntheticsCIBatchMetadata) {
	o.Metadata = &v
}

// GetPublicId returns the PublicId field value.
func (o *SyntheticsCITest) GetPublicId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetPublicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicId, true
}

// SetPublicId sets field value.
func (o *SyntheticsCITest) SetPublicId(v string) {
	o.PublicId = v
}

// GetRetry returns the Retry field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetRetry() SyntheticsTestOptionsRetry {
	if o == nil || o.Retry == nil {
		var ret SyntheticsTestOptionsRetry
		return ret
	}
	return *o.Retry
}

// GetRetryOk returns a tuple with the Retry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetRetryOk() (*SyntheticsTestOptionsRetry, bool) {
	if o == nil || o.Retry == nil {
		return nil, false
	}
	return o.Retry, true
}

// HasRetry returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasRetry() bool {
	return o != nil && o.Retry != nil
}

// SetRetry gets a reference to the given SyntheticsTestOptionsRetry and assigns it to the Retry field.
func (o *SyntheticsCITest) SetRetry(v SyntheticsTestOptionsRetry) {
	o.Retry = &v
}

// GetStartUrl returns the StartUrl field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetStartUrl() string {
	if o == nil || o.StartUrl == nil {
		var ret string
		return ret
	}
	return *o.StartUrl
}

// GetStartUrlOk returns a tuple with the StartUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetStartUrlOk() (*string, bool) {
	if o == nil || o.StartUrl == nil {
		return nil, false
	}
	return o.StartUrl, true
}

// HasStartUrl returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasStartUrl() bool {
	return o != nil && o.StartUrl != nil
}

// SetStartUrl gets a reference to the given string and assigns it to the StartUrl field.
func (o *SyntheticsCITest) SetStartUrl(v string) {
	o.StartUrl = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetVariables() map[string]string {
	if o == nil || o.Variables == nil {
		var ret map[string]string
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetVariablesOk() (*map[string]string, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return &o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasVariables() bool {
	return o != nil && o.Variables != nil
}

// SetVariables gets a reference to the given map[string]string and assigns it to the Variables field.
func (o *SyntheticsCITest) SetVariables(v map[string]string) {
	o.Variables = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SyntheticsCITest) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsCITest) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SyntheticsCITest) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *SyntheticsCITest) SetVersion(v int64) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsCITest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowInsecureCertificates != nil {
		toSerialize["allowInsecureCertificates"] = o.AllowInsecureCertificates
	}
	if o.BasicAuth != nil {
		toSerialize["basicAuth"] = o.BasicAuth
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.BodyType != nil {
		toSerialize["bodyType"] = o.BodyType
	}
	if o.Cookies != nil {
		toSerialize["cookies"] = o.Cookies
	}
	if o.DeviceIds != nil {
		toSerialize["deviceIds"] = o.DeviceIds
	}
	if o.FollowRedirects != nil {
		toSerialize["followRedirects"] = o.FollowRedirects
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Locations != nil {
		toSerialize["locations"] = o.Locations
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["public_id"] = o.PublicId
	if o.Retry != nil {
		toSerialize["retry"] = o.Retry
	}
	if o.StartUrl != nil {
		toSerialize["startUrl"] = o.StartUrl
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsCITest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowInsecureCertificates *bool                       `json:"allowInsecureCertificates,omitempty"`
		BasicAuth                 *SyntheticsBasicAuth        `json:"basicAuth,omitempty"`
		Body                      *string                     `json:"body,omitempty"`
		BodyType                  *string                     `json:"bodyType,omitempty"`
		Cookies                   *string                     `json:"cookies,omitempty"`
		DeviceIds                 []string                    `json:"deviceIds,omitempty"`
		FollowRedirects           *bool                       `json:"followRedirects,omitempty"`
		Headers                   map[string]string           `json:"headers,omitempty"`
		Locations                 []string                    `json:"locations,omitempty"`
		Metadata                  *SyntheticsCIBatchMetadata  `json:"metadata,omitempty"`
		PublicId                  *string                     `json:"public_id"`
		Retry                     *SyntheticsTestOptionsRetry `json:"retry,omitempty"`
		StartUrl                  *string                     `json:"startUrl,omitempty"`
		Variables                 map[string]string           `json:"variables,omitempty"`
		Version                   *int64                      `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.PublicId == nil {
		return fmt.Errorf("required field public_id missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowInsecureCertificates", "basicAuth", "body", "bodyType", "cookies", "deviceIds", "followRedirects", "headers", "locations", "metadata", "public_id", "retry", "startUrl", "variables", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowInsecureCertificates = all.AllowInsecureCertificates
	o.BasicAuth = all.BasicAuth
	o.Body = all.Body
	o.BodyType = all.BodyType
	o.Cookies = all.Cookies
	o.DeviceIds = all.DeviceIds
	o.FollowRedirects = all.FollowRedirects
	o.Headers = all.Headers
	o.Locations = all.Locations
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	o.PublicId = *all.PublicId
	if all.Retry != nil && all.Retry.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Retry = all.Retry
	o.StartUrl = all.StartUrl
	o.Variables = all.Variables
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
