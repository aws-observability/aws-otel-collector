// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultRequestInfo Details of the outgoing request made during the test execution.
type SyntheticsTestResultRequestInfo struct {
	// Whether insecure certificates are allowed for this request.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`
	// Body sent with the request.
	Body *string `json:"body,omitempty"`
	// gRPC call type (for example, `unary`, `healthCheck`, or `reflection`).
	CallType *string `json:"call_type,omitempty"`
	// Destination service for a Network Path test.
	DestinationService *string `json:"destination_service,omitempty"`
	// DNS server used to resolve the target host.
	DnsServer *string `json:"dns_server,omitempty"`
	// Port of the DNS server used for resolution.
	DnsServerPort *int64 `json:"dns_server_port,omitempty"`
	// Number of end-to-end probe queries issued.
	E2eQueries *int64 `json:"e2e_queries,omitempty"`
	// Files attached to the request.
	Files []SyntheticsTestResultFileRef `json:"files,omitempty"`
	// Headers sent with the request.
	Headers map[string]interface{} `json:"headers,omitempty"`
	// Host targeted by the request.
	Host *string `json:"host,omitempty"`
	// Maximum TTL for network probe packets.
	MaxTtl *int64 `json:"max_ttl,omitempty"`
	// Message sent with the request (for WebSocket/TCP/UDP tests).
	Message *string `json:"message,omitempty"`
	// HTTP method used for the request.
	Method *string `json:"method,omitempty"`
	// Whether the response body was not saved.
	NoSavingResponseBody *bool `json:"no_saving_response_body,omitempty"`
	// Port targeted by the request. Can be a number or a string variable reference.
	Port interface{} `json:"port,omitempty"`
	// Service name targeted by the request (for gRPC tests).
	Service *string `json:"service,omitempty"`
	// Source service for a Network Path test.
	SourceService *string `json:"source_service,omitempty"`
	// Request timeout in milliseconds.
	Timeout *int64 `json:"timeout,omitempty"`
	// Name of the MCP tool called (MCP tests only).
	ToolName *string `json:"tool_name,omitempty"`
	// Number of traceroute probe queries issued.
	TracerouteQueries *int64 `json:"traceroute_queries,omitempty"`
	// URL targeted by the request.
	Url *string `json:"url,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultRequestInfo instantiates a new SyntheticsTestResultRequestInfo object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultRequestInfo() *SyntheticsTestResultRequestInfo {
	this := SyntheticsTestResultRequestInfo{}
	return &this
}

// NewSyntheticsTestResultRequestInfoWithDefaults instantiates a new SyntheticsTestResultRequestInfo object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultRequestInfoWithDefaults() *SyntheticsTestResultRequestInfo {
	this := SyntheticsTestResultRequestInfo{}
	return &this
}

// GetAllowInsecure returns the AllowInsecure field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetAllowInsecure() bool {
	if o == nil || o.AllowInsecure == nil {
		var ret bool
		return ret
	}
	return *o.AllowInsecure
}

// GetAllowInsecureOk returns a tuple with the AllowInsecure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetAllowInsecureOk() (*bool, bool) {
	if o == nil || o.AllowInsecure == nil {
		return nil, false
	}
	return o.AllowInsecure, true
}

// HasAllowInsecure returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasAllowInsecure() bool {
	return o != nil && o.AllowInsecure != nil
}

// SetAllowInsecure gets a reference to the given bool and assigns it to the AllowInsecure field.
func (o *SyntheticsTestResultRequestInfo) SetAllowInsecure(v bool) {
	o.AllowInsecure = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasBody() bool {
	return o != nil && o.Body != nil
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *SyntheticsTestResultRequestInfo) SetBody(v string) {
	o.Body = &v
}

// GetCallType returns the CallType field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetCallType() string {
	if o == nil || o.CallType == nil {
		var ret string
		return ret
	}
	return *o.CallType
}

// GetCallTypeOk returns a tuple with the CallType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetCallTypeOk() (*string, bool) {
	if o == nil || o.CallType == nil {
		return nil, false
	}
	return o.CallType, true
}

// HasCallType returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasCallType() bool {
	return o != nil && o.CallType != nil
}

// SetCallType gets a reference to the given string and assigns it to the CallType field.
func (o *SyntheticsTestResultRequestInfo) SetCallType(v string) {
	o.CallType = &v
}

// GetDestinationService returns the DestinationService field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetDestinationService() string {
	if o == nil || o.DestinationService == nil {
		var ret string
		return ret
	}
	return *o.DestinationService
}

// GetDestinationServiceOk returns a tuple with the DestinationService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetDestinationServiceOk() (*string, bool) {
	if o == nil || o.DestinationService == nil {
		return nil, false
	}
	return o.DestinationService, true
}

// HasDestinationService returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasDestinationService() bool {
	return o != nil && o.DestinationService != nil
}

// SetDestinationService gets a reference to the given string and assigns it to the DestinationService field.
func (o *SyntheticsTestResultRequestInfo) SetDestinationService(v string) {
	o.DestinationService = &v
}

// GetDnsServer returns the DnsServer field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetDnsServer() string {
	if o == nil || o.DnsServer == nil {
		var ret string
		return ret
	}
	return *o.DnsServer
}

// GetDnsServerOk returns a tuple with the DnsServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetDnsServerOk() (*string, bool) {
	if o == nil || o.DnsServer == nil {
		return nil, false
	}
	return o.DnsServer, true
}

// HasDnsServer returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasDnsServer() bool {
	return o != nil && o.DnsServer != nil
}

// SetDnsServer gets a reference to the given string and assigns it to the DnsServer field.
func (o *SyntheticsTestResultRequestInfo) SetDnsServer(v string) {
	o.DnsServer = &v
}

// GetDnsServerPort returns the DnsServerPort field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetDnsServerPort() int64 {
	if o == nil || o.DnsServerPort == nil {
		var ret int64
		return ret
	}
	return *o.DnsServerPort
}

// GetDnsServerPortOk returns a tuple with the DnsServerPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetDnsServerPortOk() (*int64, bool) {
	if o == nil || o.DnsServerPort == nil {
		return nil, false
	}
	return o.DnsServerPort, true
}

// HasDnsServerPort returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasDnsServerPort() bool {
	return o != nil && o.DnsServerPort != nil
}

// SetDnsServerPort gets a reference to the given int64 and assigns it to the DnsServerPort field.
func (o *SyntheticsTestResultRequestInfo) SetDnsServerPort(v int64) {
	o.DnsServerPort = &v
}

// GetE2eQueries returns the E2eQueries field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetE2eQueries() int64 {
	if o == nil || o.E2eQueries == nil {
		var ret int64
		return ret
	}
	return *o.E2eQueries
}

// GetE2eQueriesOk returns a tuple with the E2eQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetE2eQueriesOk() (*int64, bool) {
	if o == nil || o.E2eQueries == nil {
		return nil, false
	}
	return o.E2eQueries, true
}

// HasE2eQueries returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasE2eQueries() bool {
	return o != nil && o.E2eQueries != nil
}

// SetE2eQueries gets a reference to the given int64 and assigns it to the E2eQueries field.
func (o *SyntheticsTestResultRequestInfo) SetE2eQueries(v int64) {
	o.E2eQueries = &v
}

// GetFiles returns the Files field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetFiles() []SyntheticsTestResultFileRef {
	if o == nil || o.Files == nil {
		var ret []SyntheticsTestResultFileRef
		return ret
	}
	return o.Files
}

// GetFilesOk returns a tuple with the Files field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetFilesOk() (*[]SyntheticsTestResultFileRef, bool) {
	if o == nil || o.Files == nil {
		return nil, false
	}
	return &o.Files, true
}

// HasFiles returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasFiles() bool {
	return o != nil && o.Files != nil
}

// SetFiles gets a reference to the given []SyntheticsTestResultFileRef and assigns it to the Files field.
func (o *SyntheticsTestResultRequestInfo) SetFiles(v []SyntheticsTestResultFileRef) {
	o.Files = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetHeaders() map[string]interface{} {
	if o == nil || o.Headers == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetHeadersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return &o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasHeaders() bool {
	return o != nil && o.Headers != nil
}

// SetHeaders gets a reference to the given map[string]interface{} and assigns it to the Headers field.
func (o *SyntheticsTestResultRequestInfo) SetHeaders(v map[string]interface{}) {
	o.Headers = v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasHost() bool {
	return o != nil && o.Host != nil
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *SyntheticsTestResultRequestInfo) SetHost(v string) {
	o.Host = &v
}

// GetMaxTtl returns the MaxTtl field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetMaxTtl() int64 {
	if o == nil || o.MaxTtl == nil {
		var ret int64
		return ret
	}
	return *o.MaxTtl
}

// GetMaxTtlOk returns a tuple with the MaxTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetMaxTtlOk() (*int64, bool) {
	if o == nil || o.MaxTtl == nil {
		return nil, false
	}
	return o.MaxTtl, true
}

// HasMaxTtl returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasMaxTtl() bool {
	return o != nil && o.MaxTtl != nil
}

// SetMaxTtl gets a reference to the given int64 and assigns it to the MaxTtl field.
func (o *SyntheticsTestResultRequestInfo) SetMaxTtl(v int64) {
	o.MaxTtl = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasMessage() bool {
	return o != nil && o.Message != nil
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *SyntheticsTestResultRequestInfo) SetMessage(v string) {
	o.Message = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasMethod() bool {
	return o != nil && o.Method != nil
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *SyntheticsTestResultRequestInfo) SetMethod(v string) {
	o.Method = &v
}

// GetNoSavingResponseBody returns the NoSavingResponseBody field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetNoSavingResponseBody() bool {
	if o == nil || o.NoSavingResponseBody == nil {
		var ret bool
		return ret
	}
	return *o.NoSavingResponseBody
}

// GetNoSavingResponseBodyOk returns a tuple with the NoSavingResponseBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetNoSavingResponseBodyOk() (*bool, bool) {
	if o == nil || o.NoSavingResponseBody == nil {
		return nil, false
	}
	return o.NoSavingResponseBody, true
}

// HasNoSavingResponseBody returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasNoSavingResponseBody() bool {
	return o != nil && o.NoSavingResponseBody != nil
}

// SetNoSavingResponseBody gets a reference to the given bool and assigns it to the NoSavingResponseBody field.
func (o *SyntheticsTestResultRequestInfo) SetNoSavingResponseBody(v bool) {
	o.NoSavingResponseBody = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetPort() interface{} {
	if o == nil || o.Port == nil {
		var ret interface{}
		return ret
	}
	return o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetPortOk() (*interface{}, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return &o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasPort() bool {
	return o != nil && o.Port != nil
}

// SetPort gets a reference to the given interface{} and assigns it to the Port field.
func (o *SyntheticsTestResultRequestInfo) SetPort(v interface{}) {
	o.Port = v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *SyntheticsTestResultRequestInfo) SetService(v string) {
	o.Service = &v
}

// GetSourceService returns the SourceService field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetSourceService() string {
	if o == nil || o.SourceService == nil {
		var ret string
		return ret
	}
	return *o.SourceService
}

// GetSourceServiceOk returns a tuple with the SourceService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetSourceServiceOk() (*string, bool) {
	if o == nil || o.SourceService == nil {
		return nil, false
	}
	return o.SourceService, true
}

// HasSourceService returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasSourceService() bool {
	return o != nil && o.SourceService != nil
}

// SetSourceService gets a reference to the given string and assigns it to the SourceService field.
func (o *SyntheticsTestResultRequestInfo) SetSourceService(v string) {
	o.SourceService = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasTimeout() bool {
	return o != nil && o.Timeout != nil
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *SyntheticsTestResultRequestInfo) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetToolName returns the ToolName field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetToolName() string {
	if o == nil || o.ToolName == nil {
		var ret string
		return ret
	}
	return *o.ToolName
}

// GetToolNameOk returns a tuple with the ToolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetToolNameOk() (*string, bool) {
	if o == nil || o.ToolName == nil {
		return nil, false
	}
	return o.ToolName, true
}

// HasToolName returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasToolName() bool {
	return o != nil && o.ToolName != nil
}

// SetToolName gets a reference to the given string and assigns it to the ToolName field.
func (o *SyntheticsTestResultRequestInfo) SetToolName(v string) {
	o.ToolName = &v
}

// GetTracerouteQueries returns the TracerouteQueries field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetTracerouteQueries() int64 {
	if o == nil || o.TracerouteQueries == nil {
		var ret int64
		return ret
	}
	return *o.TracerouteQueries
}

// GetTracerouteQueriesOk returns a tuple with the TracerouteQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetTracerouteQueriesOk() (*int64, bool) {
	if o == nil || o.TracerouteQueries == nil {
		return nil, false
	}
	return o.TracerouteQueries, true
}

// HasTracerouteQueries returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasTracerouteQueries() bool {
	return o != nil && o.TracerouteQueries != nil
}

// SetTracerouteQueries gets a reference to the given int64 and assigns it to the TracerouteQueries field.
func (o *SyntheticsTestResultRequestInfo) SetTracerouteQueries(v int64) {
	o.TracerouteQueries = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *SyntheticsTestResultRequestInfo) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultRequestInfo) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *SyntheticsTestResultRequestInfo) HasUrl() bool {
	return o != nil && o.Url != nil
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *SyntheticsTestResultRequestInfo) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowInsecure != nil {
		toSerialize["allow_insecure"] = o.AllowInsecure
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.CallType != nil {
		toSerialize["call_type"] = o.CallType
	}
	if o.DestinationService != nil {
		toSerialize["destination_service"] = o.DestinationService
	}
	if o.DnsServer != nil {
		toSerialize["dns_server"] = o.DnsServer
	}
	if o.DnsServerPort != nil {
		toSerialize["dns_server_port"] = o.DnsServerPort
	}
	if o.E2eQueries != nil {
		toSerialize["e2e_queries"] = o.E2eQueries
	}
	if o.Files != nil {
		toSerialize["files"] = o.Files
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.MaxTtl != nil {
		toSerialize["max_ttl"] = o.MaxTtl
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.NoSavingResponseBody != nil {
		toSerialize["no_saving_response_body"] = o.NoSavingResponseBody
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.SourceService != nil {
		toSerialize["source_service"] = o.SourceService
	}
	if o.Timeout != nil {
		toSerialize["timeout"] = o.Timeout
	}
	if o.ToolName != nil {
		toSerialize["tool_name"] = o.ToolName
	}
	if o.TracerouteQueries != nil {
		toSerialize["traceroute_queries"] = o.TracerouteQueries
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
func (o *SyntheticsTestResultRequestInfo) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowInsecure        *bool                         `json:"allow_insecure,omitempty"`
		Body                 *string                       `json:"body,omitempty"`
		CallType             *string                       `json:"call_type,omitempty"`
		DestinationService   *string                       `json:"destination_service,omitempty"`
		DnsServer            *string                       `json:"dns_server,omitempty"`
		DnsServerPort        *int64                        `json:"dns_server_port,omitempty"`
		E2eQueries           *int64                        `json:"e2e_queries,omitempty"`
		Files                []SyntheticsTestResultFileRef `json:"files,omitempty"`
		Headers              map[string]interface{}        `json:"headers,omitempty"`
		Host                 *string                       `json:"host,omitempty"`
		MaxTtl               *int64                        `json:"max_ttl,omitempty"`
		Message              *string                       `json:"message,omitempty"`
		Method               *string                       `json:"method,omitempty"`
		NoSavingResponseBody *bool                         `json:"no_saving_response_body,omitempty"`
		Port                 interface{}                   `json:"port,omitempty"`
		Service              *string                       `json:"service,omitempty"`
		SourceService        *string                       `json:"source_service,omitempty"`
		Timeout              *int64                        `json:"timeout,omitempty"`
		ToolName             *string                       `json:"tool_name,omitempty"`
		TracerouteQueries    *int64                        `json:"traceroute_queries,omitempty"`
		Url                  *string                       `json:"url,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allow_insecure", "body", "call_type", "destination_service", "dns_server", "dns_server_port", "e2e_queries", "files", "headers", "host", "max_ttl", "message", "method", "no_saving_response_body", "port", "service", "source_service", "timeout", "tool_name", "traceroute_queries", "url"})
	} else {
		return err
	}
	o.AllowInsecure = all.AllowInsecure
	o.Body = all.Body
	o.CallType = all.CallType
	o.DestinationService = all.DestinationService
	o.DnsServer = all.DnsServer
	o.DnsServerPort = all.DnsServerPort
	o.E2eQueries = all.E2eQueries
	o.Files = all.Files
	o.Headers = all.Headers
	o.Host = all.Host
	o.MaxTtl = all.MaxTtl
	o.Message = all.Message
	o.Method = all.Method
	o.NoSavingResponseBody = all.NoSavingResponseBody
	o.Port = all.Port
	o.Service = all.Service
	o.SourceService = all.SourceService
	o.Timeout = all.Timeout
	o.ToolName = all.ToolName
	o.TracerouteQueries = all.TracerouteQueries
	o.Url = all.Url

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
