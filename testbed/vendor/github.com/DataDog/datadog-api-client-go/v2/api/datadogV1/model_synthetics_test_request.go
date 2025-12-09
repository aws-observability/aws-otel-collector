// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestRequest Object describing the Synthetic test request.
type SyntheticsTestRequest struct {
	// Allows loading insecure content for an HTTP request in a multistep test step.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`
	// Object to handle basic authentication when performing the test.
	BasicAuth *SyntheticsBasicAuth `json:"basicAuth,omitempty"`
	// Body to include in the test.
	Body *string `json:"body,omitempty"`
	// Type of the request body.
	BodyType *SyntheticsTestRequestBodyType `json:"bodyType,omitempty"`
	// The type of gRPC call to perform.
	CallType *SyntheticsTestCallType `json:"callType,omitempty"`
	// Client certificate to use when performing the test request.
	Certificate *SyntheticsTestRequestCertificate `json:"certificate,omitempty"`
	// By default, the client certificate is applied on the domain of the starting URL for browser tests. If you want your client certificate to be applied on other domains instead, add them in `certificateDomains`.
	CertificateDomains []string `json:"certificateDomains,omitempty"`
	// Check for certificate revocation.
	CheckCertificateRevocation *bool `json:"checkCertificateRevocation,omitempty"`
	// A protobuf JSON descriptor that needs to be gzipped first then base64 encoded.
	CompressedJsonDescriptor *string `json:"compressedJsonDescriptor,omitempty"`
	// A protobuf file that needs to be gzipped first then base64 encoded.
	CompressedProtoFile *string `json:"compressedProtoFile,omitempty"`
	// Disable fetching intermediate certificates from AIA.
	DisableAiaIntermediateFetching *bool `json:"disableAiaIntermediateFetching,omitempty"`
	// DNS server to use for DNS tests.
	DnsServer *string `json:"dnsServer,omitempty"`
	// DNS server port to use for DNS tests.
	DnsServerPort *SyntheticsTestRequestDNSServerPort `json:"dnsServerPort,omitempty"`
	// Files to be used as part of the request in the test. Only valid if `bodyType` is `multipart/form-data`.
	Files []SyntheticsTestRequestBodyFile `json:"files,omitempty"`
	// Specifies whether or not the request follows redirects.
	FollowRedirects *bool `json:"follow_redirects,omitempty"`
	// Form to be used as part of the request in the test. Only valid if `bodyType` is `multipart/form-data`.
	Form map[string]string `json:"form,omitempty"`
	// Headers to include when performing the test.
	Headers map[string]string `json:"headers,omitempty"`
	// Host name to perform the test with.
	Host *string `json:"host,omitempty"`
	// HTTP version to use for a Synthetic test.
	HttpVersion *SyntheticsTestOptionsHTTPVersion `json:"httpVersion,omitempty"`
	// Whether the message is base64 encoded.
	IsMessageBase64Encoded *bool `json:"isMessageBase64Encoded,omitempty"`
	// Message to send for UDP or WebSocket tests.
	Message *string `json:"message,omitempty"`
	// Metadata to include when performing the gRPC test.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Either the HTTP method/verb to use or a gRPC method available on the service set in the `service` field. Required if `subtype` is `HTTP` or if `subtype` is `grpc` and `callType` is `unary`.
	Method *string `json:"method,omitempty"`
	// Determines whether or not to save the response body.
	NoSavingResponseBody *bool `json:"noSavingResponseBody,omitempty"`
	// Number of pings to use per test.
	NumberOfPackets *int32 `json:"numberOfPackets,omitempty"`
	// Persist cookies across redirects.
	PersistCookies *bool `json:"persistCookies,omitempty"`
	// Port to use when performing the test.
	Port *SyntheticsTestRequestPort `json:"port,omitempty"`
	// The proxy to perform the test.
	Proxy *SyntheticsTestRequestProxy `json:"proxy,omitempty"`
	// Query to use for the test.
	Query interface{} `json:"query,omitempty"`
	// For SSL tests, it specifies on which server you want to initiate the TLS handshake,
	// allowing the server to present one of multiple possible certificates on
	// the same IP address and TCP port number.
	Servername *string `json:"servername,omitempty"`
	// The gRPC service on which you want to perform the gRPC call.
	Service *string `json:"service,omitempty"`
	// Turns on a traceroute probe to discover all gateways along the path to the host destination.
	ShouldTrackHops *bool `json:"shouldTrackHops,omitempty"`
	// Timeout in seconds for the test.
	Timeout *float64 `json:"timeout,omitempty"`
	// URL to perform the test with.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestRequest instantiates a new SyntheticsTestRequest object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestRequest() *SyntheticsTestRequest {
	this := SyntheticsTestRequest{}
	return &this
}

// NewSyntheticsTestRequestWithDefaults instantiates a new SyntheticsTestRequest object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestRequestWithDefaults() *SyntheticsTestRequest {
	this := SyntheticsTestRequest{}
	return &this
}

// GetAllowInsecure returns the AllowInsecure field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetAllowInsecure() bool {
	if o == nil || o.AllowInsecure == nil {
		var ret bool
		return ret
	}
	return *o.AllowInsecure
}

// GetAllowInsecureOk returns a tuple with the AllowInsecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetAllowInsecureOk() (*bool, bool) {
	if o == nil || o.AllowInsecure == nil {
		return nil, false
	}
	return o.AllowInsecure, true
}

// HasAllowInsecure returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasAllowInsecure() bool {
	return o != nil && o.AllowInsecure != nil
}

// SetAllowInsecure gets a reference to the given bool and assigns it to the AllowInsecure field.
func (o *SyntheticsTestRequest) SetAllowInsecure(v bool) {
	o.AllowInsecure = &v
}

// GetBasicAuth returns the BasicAuth field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetBasicAuth() SyntheticsBasicAuth {
	if o == nil || o.BasicAuth == nil {
		var ret SyntheticsBasicAuth
		return ret
	}
	return *o.BasicAuth
}

// GetBasicAuthOk returns a tuple with the BasicAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetBasicAuthOk() (*SyntheticsBasicAuth, bool) {
	if o == nil || o.BasicAuth == nil {
		return nil, false
	}
	return o.BasicAuth, true
}

// HasBasicAuth returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasBasicAuth() bool {
	return o != nil && o.BasicAuth != nil
}

// SetBasicAuth gets a reference to the given SyntheticsBasicAuth and assigns it to the BasicAuth field.
func (o *SyntheticsTestRequest) SetBasicAuth(v SyntheticsBasicAuth) {
	o.BasicAuth = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasBody() bool {
	return o != nil && o.Body != nil
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *SyntheticsTestRequest) SetBody(v string) {
	o.Body = &v
}

// GetBodyType returns the BodyType field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetBodyType() SyntheticsTestRequestBodyType {
	if o == nil || o.BodyType == nil {
		var ret SyntheticsTestRequestBodyType
		return ret
	}
	return *o.BodyType
}

// GetBodyTypeOk returns a tuple with the BodyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetBodyTypeOk() (*SyntheticsTestRequestBodyType, bool) {
	if o == nil || o.BodyType == nil {
		return nil, false
	}
	return o.BodyType, true
}

// HasBodyType returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasBodyType() bool {
	return o != nil && o.BodyType != nil
}

// SetBodyType gets a reference to the given SyntheticsTestRequestBodyType and assigns it to the BodyType field.
func (o *SyntheticsTestRequest) SetBodyType(v SyntheticsTestRequestBodyType) {
	o.BodyType = &v
}

// GetCallType returns the CallType field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetCallType() SyntheticsTestCallType {
	if o == nil || o.CallType == nil {
		var ret SyntheticsTestCallType
		return ret
	}
	return *o.CallType
}

// GetCallTypeOk returns a tuple with the CallType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetCallTypeOk() (*SyntheticsTestCallType, bool) {
	if o == nil || o.CallType == nil {
		return nil, false
	}
	return o.CallType, true
}

// HasCallType returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasCallType() bool {
	return o != nil && o.CallType != nil
}

// SetCallType gets a reference to the given SyntheticsTestCallType and assigns it to the CallType field.
func (o *SyntheticsTestRequest) SetCallType(v SyntheticsTestCallType) {
	o.CallType = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetCertificate() SyntheticsTestRequestCertificate {
	if o == nil || o.Certificate == nil {
		var ret SyntheticsTestRequestCertificate
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetCertificateOk() (*SyntheticsTestRequestCertificate, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasCertificate() bool {
	return o != nil && o.Certificate != nil
}

// SetCertificate gets a reference to the given SyntheticsTestRequestCertificate and assigns it to the Certificate field.
func (o *SyntheticsTestRequest) SetCertificate(v SyntheticsTestRequestCertificate) {
	o.Certificate = &v
}

// GetCertificateDomains returns the CertificateDomains field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetCertificateDomains() []string {
	if o == nil || o.CertificateDomains == nil {
		var ret []string
		return ret
	}
	return o.CertificateDomains
}

// GetCertificateDomainsOk returns a tuple with the CertificateDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetCertificateDomainsOk() (*[]string, bool) {
	if o == nil || o.CertificateDomains == nil {
		return nil, false
	}
	return &o.CertificateDomains, true
}

// HasCertificateDomains returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasCertificateDomains() bool {
	return o != nil && o.CertificateDomains != nil
}

// SetCertificateDomains gets a reference to the given []string and assigns it to the CertificateDomains field.
func (o *SyntheticsTestRequest) SetCertificateDomains(v []string) {
	o.CertificateDomains = v
}

// GetCheckCertificateRevocation returns the CheckCertificateRevocation field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetCheckCertificateRevocation() bool {
	if o == nil || o.CheckCertificateRevocation == nil {
		var ret bool
		return ret
	}
	return *o.CheckCertificateRevocation
}

// GetCheckCertificateRevocationOk returns a tuple with the CheckCertificateRevocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetCheckCertificateRevocationOk() (*bool, bool) {
	if o == nil || o.CheckCertificateRevocation == nil {
		return nil, false
	}
	return o.CheckCertificateRevocation, true
}

// HasCheckCertificateRevocation returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasCheckCertificateRevocation() bool {
	return o != nil && o.CheckCertificateRevocation != nil
}

// SetCheckCertificateRevocation gets a reference to the given bool and assigns it to the CheckCertificateRevocation field.
func (o *SyntheticsTestRequest) SetCheckCertificateRevocation(v bool) {
	o.CheckCertificateRevocation = &v
}

// GetCompressedJsonDescriptor returns the CompressedJsonDescriptor field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetCompressedJsonDescriptor() string {
	if o == nil || o.CompressedJsonDescriptor == nil {
		var ret string
		return ret
	}
	return *o.CompressedJsonDescriptor
}

// GetCompressedJsonDescriptorOk returns a tuple with the CompressedJsonDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetCompressedJsonDescriptorOk() (*string, bool) {
	if o == nil || o.CompressedJsonDescriptor == nil {
		return nil, false
	}
	return o.CompressedJsonDescriptor, true
}

// HasCompressedJsonDescriptor returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasCompressedJsonDescriptor() bool {
	return o != nil && o.CompressedJsonDescriptor != nil
}

// SetCompressedJsonDescriptor gets a reference to the given string and assigns it to the CompressedJsonDescriptor field.
func (o *SyntheticsTestRequest) SetCompressedJsonDescriptor(v string) {
	o.CompressedJsonDescriptor = &v
}

// GetCompressedProtoFile returns the CompressedProtoFile field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetCompressedProtoFile() string {
	if o == nil || o.CompressedProtoFile == nil {
		var ret string
		return ret
	}
	return *o.CompressedProtoFile
}

// GetCompressedProtoFileOk returns a tuple with the CompressedProtoFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetCompressedProtoFileOk() (*string, bool) {
	if o == nil || o.CompressedProtoFile == nil {
		return nil, false
	}
	return o.CompressedProtoFile, true
}

// HasCompressedProtoFile returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasCompressedProtoFile() bool {
	return o != nil && o.CompressedProtoFile != nil
}

// SetCompressedProtoFile gets a reference to the given string and assigns it to the CompressedProtoFile field.
func (o *SyntheticsTestRequest) SetCompressedProtoFile(v string) {
	o.CompressedProtoFile = &v
}

// GetDisableAiaIntermediateFetching returns the DisableAiaIntermediateFetching field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetDisableAiaIntermediateFetching() bool {
	if o == nil || o.DisableAiaIntermediateFetching == nil {
		var ret bool
		return ret
	}
	return *o.DisableAiaIntermediateFetching
}

// GetDisableAiaIntermediateFetchingOk returns a tuple with the DisableAiaIntermediateFetching field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetDisableAiaIntermediateFetchingOk() (*bool, bool) {
	if o == nil || o.DisableAiaIntermediateFetching == nil {
		return nil, false
	}
	return o.DisableAiaIntermediateFetching, true
}

// HasDisableAiaIntermediateFetching returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasDisableAiaIntermediateFetching() bool {
	return o != nil && o.DisableAiaIntermediateFetching != nil
}

// SetDisableAiaIntermediateFetching gets a reference to the given bool and assigns it to the DisableAiaIntermediateFetching field.
func (o *SyntheticsTestRequest) SetDisableAiaIntermediateFetching(v bool) {
	o.DisableAiaIntermediateFetching = &v
}

// GetDnsServer returns the DnsServer field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetDnsServer() string {
	if o == nil || o.DnsServer == nil {
		var ret string
		return ret
	}
	return *o.DnsServer
}

// GetDnsServerOk returns a tuple with the DnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetDnsServerOk() (*string, bool) {
	if o == nil || o.DnsServer == nil {
		return nil, false
	}
	return o.DnsServer, true
}

// HasDnsServer returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasDnsServer() bool {
	return o != nil && o.DnsServer != nil
}

// SetDnsServer gets a reference to the given string and assigns it to the DnsServer field.
func (o *SyntheticsTestRequest) SetDnsServer(v string) {
	o.DnsServer = &v
}

// GetDnsServerPort returns the DnsServerPort field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetDnsServerPort() SyntheticsTestRequestDNSServerPort {
	if o == nil || o.DnsServerPort == nil {
		var ret SyntheticsTestRequestDNSServerPort
		return ret
	}
	return *o.DnsServerPort
}

// GetDnsServerPortOk returns a tuple with the DnsServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetDnsServerPortOk() (*SyntheticsTestRequestDNSServerPort, bool) {
	if o == nil || o.DnsServerPort == nil {
		return nil, false
	}
	return o.DnsServerPort, true
}

// HasDnsServerPort returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasDnsServerPort() bool {
	return o != nil && o.DnsServerPort != nil
}

// SetDnsServerPort gets a reference to the given SyntheticsTestRequestDNSServerPort and assigns it to the DnsServerPort field.
func (o *SyntheticsTestRequest) SetDnsServerPort(v SyntheticsTestRequestDNSServerPort) {
	o.DnsServerPort = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetFiles() []SyntheticsTestRequestBodyFile {
	if o == nil || o.Files == nil {
		var ret []SyntheticsTestRequestBodyFile
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetFilesOk() (*[]SyntheticsTestRequestBodyFile, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return &o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasFiles() bool {
	return o != nil && o.Files != nil
}

// SetFiles gets a reference to the given []SyntheticsTestRequestBodyFile and assigns it to the Files field.
func (o *SyntheticsTestRequest) SetFiles(v []SyntheticsTestRequestBodyFile) {
	o.Files = v
}

// GetFollowRedirects returns the FollowRedirects field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetFollowRedirects() bool {
	if o == nil || o.FollowRedirects == nil {
		var ret bool
		return ret
	}
	return *o.FollowRedirects
}

// GetFollowRedirectsOk returns a tuple with the FollowRedirects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetFollowRedirectsOk() (*bool, bool) {
	if o == nil || o.FollowRedirects == nil {
		return nil, false
	}
	return o.FollowRedirects, true
}

// HasFollowRedirects returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasFollowRedirects() bool {
	return o != nil && o.FollowRedirects != nil
}

// SetFollowRedirects gets a reference to the given bool and assigns it to the FollowRedirects field.
func (o *SyntheticsTestRequest) SetFollowRedirects(v bool) {
	o.FollowRedirects = &v
}

// GetForm returns the Form field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetForm() map[string]string {
	if o == nil || o.Form == nil {
		var ret map[string]string
		return ret
	}
	return o.Form
}

// GetFormOk returns a tuple with the Form field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetFormOk() (*map[string]string, bool) {
	if o == nil || o.Form == nil {
		return nil, false
	}
	return &o.Form, true
}

// HasForm returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasForm() bool {
	return o != nil && o.Form != nil
}

// SetForm gets a reference to the given map[string]string and assigns it to the Form field.
func (o *SyntheticsTestRequest) SetForm(v map[string]string) {
	o.Form = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetHeaders() map[string]string {
	if o == nil || o.Headers == nil {
		var ret map[string]string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasHeaders() bool {
	return o != nil && o.Headers != nil
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *SyntheticsTestRequest) SetHeaders(v map[string]string) {
	o.Headers = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *SyntheticsTestRequest) SetHost(v string) {
	o.Host = &v
}

// GetHttpVersion returns the HttpVersion field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetHttpVersion() SyntheticsTestOptionsHTTPVersion {
	if o == nil || o.HttpVersion == nil {
		var ret SyntheticsTestOptionsHTTPVersion
		return ret
	}
	return *o.HttpVersion
}

// GetHttpVersionOk returns a tuple with the HttpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetHttpVersionOk() (*SyntheticsTestOptionsHTTPVersion, bool) {
	if o == nil || o.HttpVersion == nil {
		return nil, false
	}
	return o.HttpVersion, true
}

// HasHttpVersion returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasHttpVersion() bool {
	return o != nil && o.HttpVersion != nil
}

// SetHttpVersion gets a reference to the given SyntheticsTestOptionsHTTPVersion and assigns it to the HttpVersion field.
func (o *SyntheticsTestRequest) SetHttpVersion(v SyntheticsTestOptionsHTTPVersion) {
	o.HttpVersion = &v
}

// GetIsMessageBase64Encoded returns the IsMessageBase64Encoded field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetIsMessageBase64Encoded() bool {
	if o == nil || o.IsMessageBase64Encoded == nil {
		var ret bool
		return ret
	}
	return *o.IsMessageBase64Encoded
}

// GetIsMessageBase64EncodedOk returns a tuple with the IsMessageBase64Encoded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetIsMessageBase64EncodedOk() (*bool, bool) {
	if o == nil || o.IsMessageBase64Encoded == nil {
		return nil, false
	}
	return o.IsMessageBase64Encoded, true
}

// HasIsMessageBase64Encoded returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasIsMessageBase64Encoded() bool {
	return o != nil && o.IsMessageBase64Encoded != nil
}

// SetIsMessageBase64Encoded gets a reference to the given bool and assigns it to the IsMessageBase64Encoded field.
func (o *SyntheticsTestRequest) SetIsMessageBase64Encoded(v bool) {
	o.IsMessageBase64Encoded = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SyntheticsTestRequest) SetMessage(v string) {
	o.Message = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *SyntheticsTestRequest) SetMetadata(v map[string]string) {
	o.Metadata = v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasMethod() bool {
	return o != nil && o.Method != nil
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *SyntheticsTestRequest) SetMethod(v string) {
	o.Method = &v
}

// GetNoSavingResponseBody returns the NoSavingResponseBody field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetNoSavingResponseBody() bool {
	if o == nil || o.NoSavingResponseBody == nil {
		var ret bool
		return ret
	}
	return *o.NoSavingResponseBody
}

// GetNoSavingResponseBodyOk returns a tuple with the NoSavingResponseBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetNoSavingResponseBodyOk() (*bool, bool) {
	if o == nil || o.NoSavingResponseBody == nil {
		return nil, false
	}
	return o.NoSavingResponseBody, true
}

// HasNoSavingResponseBody returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasNoSavingResponseBody() bool {
	return o != nil && o.NoSavingResponseBody != nil
}

// SetNoSavingResponseBody gets a reference to the given bool and assigns it to the NoSavingResponseBody field.
func (o *SyntheticsTestRequest) SetNoSavingResponseBody(v bool) {
	o.NoSavingResponseBody = &v
}

// GetNumberOfPackets returns the NumberOfPackets field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetNumberOfPackets() int32 {
	if o == nil || o.NumberOfPackets == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfPackets
}

// GetNumberOfPacketsOk returns a tuple with the NumberOfPackets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetNumberOfPacketsOk() (*int32, bool) {
	if o == nil || o.NumberOfPackets == nil {
		return nil, false
	}
	return o.NumberOfPackets, true
}

// HasNumberOfPackets returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasNumberOfPackets() bool {
	return o != nil && o.NumberOfPackets != nil
}

// SetNumberOfPackets gets a reference to the given int32 and assigns it to the NumberOfPackets field.
func (o *SyntheticsTestRequest) SetNumberOfPackets(v int32) {
	o.NumberOfPackets = &v
}

// GetPersistCookies returns the PersistCookies field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetPersistCookies() bool {
	if o == nil || o.PersistCookies == nil {
		var ret bool
		return ret
	}
	return *o.PersistCookies
}

// GetPersistCookiesOk returns a tuple with the PersistCookies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetPersistCookiesOk() (*bool, bool) {
	if o == nil || o.PersistCookies == nil {
		return nil, false
	}
	return o.PersistCookies, true
}

// HasPersistCookies returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasPersistCookies() bool {
	return o != nil && o.PersistCookies != nil
}

// SetPersistCookies gets a reference to the given bool and assigns it to the PersistCookies field.
func (o *SyntheticsTestRequest) SetPersistCookies(v bool) {
	o.PersistCookies = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetPort() SyntheticsTestRequestPort {
	if o == nil || o.Port == nil {
		var ret SyntheticsTestRequestPort
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetPortOk() (*SyntheticsTestRequestPort, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given SyntheticsTestRequestPort and assigns it to the Port field.
func (o *SyntheticsTestRequest) SetPort(v SyntheticsTestRequestPort) {
	o.Port = &v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetProxy() SyntheticsTestRequestProxy {
	if o == nil || o.Proxy == nil {
		var ret SyntheticsTestRequestProxy
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetProxyOk() (*SyntheticsTestRequestProxy, bool) {
	if o == nil || o.Proxy == nil {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasProxy() bool {
	return o != nil && o.Proxy != nil
}

// SetProxy gets a reference to the given SyntheticsTestRequestProxy and assigns it to the Proxy field.
func (o *SyntheticsTestRequest) SetProxy(v SyntheticsTestRequestProxy) {
	o.Proxy = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetQuery() interface{} {
	if o == nil || o.Query == nil {
		var ret interface{}
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetQueryOk() (*interface{}, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return &o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given interface{} and assigns it to the Query field.
func (o *SyntheticsTestRequest) SetQuery(v interface{}) {
	o.Query = v
}

// GetServername returns the Servername field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetServername() string {
	if o == nil || o.Servername == nil {
		var ret string
		return ret
	}
	return *o.Servername
}

// GetServernameOk returns a tuple with the Servername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetServernameOk() (*string, bool) {
	if o == nil || o.Servername == nil {
		return nil, false
	}
	return o.Servername, true
}

// HasServername returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasServername() bool {
	return o != nil && o.Servername != nil
}

// SetServername gets a reference to the given string and assigns it to the Servername field.
func (o *SyntheticsTestRequest) SetServername(v string) {
	o.Servername = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *SyntheticsTestRequest) SetService(v string) {
	o.Service = &v
}

// GetShouldTrackHops returns the ShouldTrackHops field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetShouldTrackHops() bool {
	if o == nil || o.ShouldTrackHops == nil {
		var ret bool
		return ret
	}
	return *o.ShouldTrackHops
}

// GetShouldTrackHopsOk returns a tuple with the ShouldTrackHops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetShouldTrackHopsOk() (*bool, bool) {
	if o == nil || o.ShouldTrackHops == nil {
		return nil, false
	}
	return o.ShouldTrackHops, true
}

// HasShouldTrackHops returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasShouldTrackHops() bool {
	return o != nil && o.ShouldTrackHops != nil
}

// SetShouldTrackHops gets a reference to the given bool and assigns it to the ShouldTrackHops field.
func (o *SyntheticsTestRequest) SetShouldTrackHops(v bool) {
	o.ShouldTrackHops = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetTimeout() float64 {
	if o == nil || o.Timeout == nil {
		var ret float64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetTimeoutOk() (*float64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasTimeout() bool {
	return o != nil && o.Timeout != nil
}

// SetTimeout gets a reference to the given float64 and assigns it to the Timeout field.
func (o *SyntheticsTestRequest) SetTimeout(v float64) {
	o.Timeout = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SyntheticsTestRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SyntheticsTestRequest) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SyntheticsTestRequest) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowInsecure != nil {
		toSerialize["allow_insecure"] = o.AllowInsecure
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
	if o.CallType != nil {
		toSerialize["callType"] = o.CallType
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.CertificateDomains != nil {
		toSerialize["certificateDomains"] = o.CertificateDomains
	}
	if o.CheckCertificateRevocation != nil {
		toSerialize["checkCertificateRevocation"] = o.CheckCertificateRevocation
	}
	if o.CompressedJsonDescriptor != nil {
		toSerialize["compressedJsonDescriptor"] = o.CompressedJsonDescriptor
	}
	if o.CompressedProtoFile != nil {
		toSerialize["compressedProtoFile"] = o.CompressedProtoFile
	}
	if o.DisableAiaIntermediateFetching != nil {
		toSerialize["disableAiaIntermediateFetching"] = o.DisableAiaIntermediateFetching
	}
	if o.DnsServer != nil {
		toSerialize["dnsServer"] = o.DnsServer
	}
	if o.DnsServerPort != nil {
		toSerialize["dnsServerPort"] = o.DnsServerPort
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.FollowRedirects != nil {
		toSerialize["follow_redirects"] = o.FollowRedirects
	}
	if o.Form != nil {
		toSerialize["form"] = o.Form
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.HttpVersion != nil {
		toSerialize["httpVersion"] = o.HttpVersion
	}
	if o.IsMessageBase64Encoded != nil {
		toSerialize["isMessageBase64Encoded"] = o.IsMessageBase64Encoded
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.NoSavingResponseBody != nil {
		toSerialize["noSavingResponseBody"] = o.NoSavingResponseBody
	}
	if o.NumberOfPackets != nil {
		toSerialize["numberOfPackets"] = o.NumberOfPackets
	}
	if o.PersistCookies != nil {
		toSerialize["persistCookies"] = o.PersistCookies
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Proxy != nil {
		toSerialize["proxy"] = o.Proxy
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Servername != nil {
		toSerialize["servername"] = o.Servername
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.ShouldTrackHops != nil {
		toSerialize["shouldTrackHops"] = o.ShouldTrackHops
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
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
func (o *SyntheticsTestRequest) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowInsecure                  *bool                               `json:"allow_insecure,omitempty"`
		BasicAuth                      *SyntheticsBasicAuth                `json:"basicAuth,omitempty"`
		Body                           *string                             `json:"body,omitempty"`
		BodyType                       *SyntheticsTestRequestBodyType      `json:"bodyType,omitempty"`
		CallType                       *SyntheticsTestCallType             `json:"callType,omitempty"`
		Certificate                    *SyntheticsTestRequestCertificate   `json:"certificate,omitempty"`
		CertificateDomains             []string                            `json:"certificateDomains,omitempty"`
		CheckCertificateRevocation     *bool                               `json:"checkCertificateRevocation,omitempty"`
		CompressedJsonDescriptor       *string                             `json:"compressedJsonDescriptor,omitempty"`
		CompressedProtoFile            *string                             `json:"compressedProtoFile,omitempty"`
		DisableAiaIntermediateFetching *bool                               `json:"disableAiaIntermediateFetching,omitempty"`
		DnsServer                      *string                             `json:"dnsServer,omitempty"`
		DnsServerPort                  *SyntheticsTestRequestDNSServerPort `json:"dnsServerPort,omitempty"`
		Files                          []SyntheticsTestRequestBodyFile     `json:"files,omitempty"`
		FollowRedirects                *bool                               `json:"follow_redirects,omitempty"`
		Form                           map[string]string                   `json:"form,omitempty"`
		Headers                        map[string]string                   `json:"headers,omitempty"`
		Host                           *string                             `json:"host,omitempty"`
		HttpVersion                    *SyntheticsTestOptionsHTTPVersion   `json:"httpVersion,omitempty"`
		IsMessageBase64Encoded         *bool                               `json:"isMessageBase64Encoded,omitempty"`
		Message                        *string                             `json:"message,omitempty"`
		Metadata                       map[string]string                   `json:"metadata,omitempty"`
		Method                         *string                             `json:"method,omitempty"`
		NoSavingResponseBody           *bool                               `json:"noSavingResponseBody,omitempty"`
		NumberOfPackets                *int32                              `json:"numberOfPackets,omitempty"`
		PersistCookies                 *bool                               `json:"persistCookies,omitempty"`
		Port                           *SyntheticsTestRequestPort          `json:"port,omitempty"`
		Proxy                          *SyntheticsTestRequestProxy         `json:"proxy,omitempty"`
		Query                          interface{}                         `json:"query,omitempty"`
		Servername                     *string                             `json:"servername,omitempty"`
		Service                        *string                             `json:"service,omitempty"`
		ShouldTrackHops                *bool                               `json:"shouldTrackHops,omitempty"`
		Timeout                        *float64                            `json:"timeout,omitempty"`
		Url                            *string                             `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allow_insecure", "basicAuth", "body", "bodyType", "callType", "certificate", "certificateDomains", "checkCertificateRevocation", "compressedJsonDescriptor", "compressedProtoFile", "disableAiaIntermediateFetching", "dnsServer", "dnsServerPort", "files", "follow_redirects", "form", "headers", "host", "httpVersion", "isMessageBase64Encoded", "message", "metadata", "method", "noSavingResponseBody", "numberOfPackets", "persistCookies", "port", "proxy", "query", "servername", "service", "shouldTrackHops", "timeout", "url"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowInsecure = all.AllowInsecure
	o.BasicAuth = all.BasicAuth
	o.Body = all.Body
	if all.BodyType != nil && !all.BodyType.IsValid() {
		hasInvalidField = true
	} else {
		o.BodyType = all.BodyType
	}
	if all.CallType != nil && !all.CallType.IsValid() {
		hasInvalidField = true
	} else {
		o.CallType = all.CallType
	}
	if all.Certificate != nil && all.Certificate.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Certificate = all.Certificate
	o.CertificateDomains = all.CertificateDomains
	o.CheckCertificateRevocation = all.CheckCertificateRevocation
	o.CompressedJsonDescriptor = all.CompressedJsonDescriptor
	o.CompressedProtoFile = all.CompressedProtoFile
	o.DisableAiaIntermediateFetching = all.DisableAiaIntermediateFetching
	o.DnsServer = all.DnsServer
	o.DnsServerPort = all.DnsServerPort
	o.Files = all.Files
	o.FollowRedirects = all.FollowRedirects
	o.Form = all.Form
	o.Headers = all.Headers
	o.Host = all.Host
	if all.HttpVersion != nil && !all.HttpVersion.IsValid() {
		hasInvalidField = true
	} else {
		o.HttpVersion = all.HttpVersion
	}
	o.IsMessageBase64Encoded = all.IsMessageBase64Encoded
	o.Message = all.Message
	o.Metadata = all.Metadata
	o.Method = all.Method
	o.NoSavingResponseBody = all.NoSavingResponseBody
	o.NumberOfPackets = all.NumberOfPackets
	o.PersistCookies = all.PersistCookies
	o.Port = all.Port
	if all.Proxy != nil && all.Proxy.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Proxy = all.Proxy
	o.Query = all.Query
	o.Servername = all.Servername
	o.Service = all.Service
	o.ShouldTrackHops = all.ShouldTrackHops
	o.Timeout = all.Timeout
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
