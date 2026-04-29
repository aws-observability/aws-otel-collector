// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsTestResultDetail Full result details for a Synthetic test execution.
type SyntheticsTestResultDetail struct {
	// Assertion results produced by the test.
	Assertions []SyntheticsTestResultAssertionResult `json:"assertions,omitempty"`
	// Storage bucket keys for artifacts produced during a step or test.
	BucketKeys *SyntheticsTestResultBucketKeys `json:"bucket_keys,omitempty"`
	// gRPC call type (for example, `unary`, `healthCheck`, or `reflection`).
	CallType *string `json:"call_type,omitempty"`
	// SSL/TLS certificate information returned from an SSL test.
	Cert *SyntheticsTestResultCertificate `json:"cert,omitempty"`
	// Compressed JSON descriptor for the test (internal format).
	CompressedJsonDescriptor *string `json:"compressed_json_descriptor,omitempty"`
	// Compressed representation of the test steps (internal format).
	CompressedSteps *string `json:"compressed_steps,omitempty"`
	// Outcome of the connection attempt (for example, `established`, `refused`).
	ConnectionOutcome *string `json:"connection_outcome,omitempty"`
	// DNS resolution details recorded during the test execution.
	DnsResolution *SyntheticsTestResultDnsResolution `json:"dns_resolution,omitempty"`
	// Duration of the test execution (in milliseconds).
	Duration *float64 `json:"duration,omitempty"`
	// Whether the test exited early because a step marked with `exitIfSucceed` passed.
	ExitedOnStepSuccess *bool `json:"exited_on_step_success,omitempty"`
	// Details about the failure of a Synthetic test.
	Failure *SyntheticsTestResultFailure `json:"failure,omitempty"`
	// Timestamp of when the test finished (in milliseconds).
	FinishedAt *int64 `json:"finished_at,omitempty"`
	// Handshake request and response for protocol-level tests.
	Handshake *SyntheticsTestResultHandshake `json:"handshake,omitempty"`
	// The unique identifier for this result.
	Id *string `json:"id,omitempty"`
	// The initial result ID before any retries.
	InitialId *string `json:"initial_id,omitempty"`
	// Whether this result is from a fast retry.
	IsFastRetry *bool `json:"is_fast_retry,omitempty"`
	// Whether this result is from the last retry.
	IsLastRetry *bool `json:"is_last_retry,omitempty"`
	// Network Path test result capturing the path between source and destination.
	Netpath *SyntheticsTestResultNetpath `json:"netpath,omitempty"`
	// Aggregated network statistics from the test execution.
	Netstats *SyntheticsTestResultNetstats `json:"netstats,omitempty"`
	// OCSP response received while validating a certificate.
	Ocsp *SyntheticsTestResultOCSPResponse `json:"ocsp,omitempty"`
	// A network probe result, used for traceroute hops and ping summaries.
	Ping *SyntheticsTestResultTracerouteHop `json:"ping,omitempty"`
	// Number of emails received during the test (email tests).
	ReceivedEmailCount *int64 `json:"received_email_count,omitempty"`
	// Message received from the target (for WebSocket/TCP/UDP tests).
	ReceivedMessage *string `json:"received_message,omitempty"`
	// Details of the outgoing request made during the test execution.
	Request *SyntheticsTestResultRequestInfo `json:"request,omitempty"`
	// IP address resolved for the target host.
	ResolvedIp *string `json:"resolved_ip,omitempty"`
	// Details of the response received during the test execution.
	Response *SyntheticsTestResultResponseInfo `json:"response,omitempty"`
	// The type of run for a Synthetic test result.
	RunType *SyntheticsTestResultRunType `json:"run_type,omitempty"`
	// Message sent to the target (for WebSocket/TCP/UDP tests).
	SentMessage *string `json:"sent_message,omitempty"`
	// Start URL for the test (browser tests).
	StartUrl *string `json:"start_url,omitempty"`
	// Timestamp of when the test started (in milliseconds).
	StartedAt *int64 `json:"started_at,omitempty"`
	// Status of a Synthetic test result.
	Status *SyntheticsTestResultStatus `json:"status,omitempty"`
	// Step results (for browser, mobile, and multistep API tests).
	Steps []SyntheticsTestResultStep `json:"steps,omitempty"`
	// Time to interactive in milliseconds (browser tests).
	TimeToInteractive *int64 `json:"time_to_interactive,omitempty"`
	// Timing breakdown of the test request phases (for example, DNS, TCP, TLS, first byte).
	Timings map[string]interface{} `json:"timings,omitempty"`
	// Trace identifiers associated with a Synthetic test result.
	Trace *SyntheticsTestResultTrace `json:"trace,omitempty"`
	// Traceroute hop results (for network tests).
	Traceroute []SyntheticsTestResultTracerouteHop `json:"traceroute,omitempty"`
	// Timestamp of when the test was triggered (in milliseconds).
	TriggeredAt *int64 `json:"triggered_at,omitempty"`
	// Whether the test was executed through a tunnel.
	Tunnel *bool `json:"tunnel,omitempty"`
	// Turns executed by a goal-based browser test.
	Turns []SyntheticsTestResultTurn `json:"turns,omitempty"`
	// Whether the test runner was unhealthy at the time of execution.
	Unhealthy *bool `json:"unhealthy,omitempty"`
	// Variables captured during a test step.
	Variables *SyntheticsTestResultVariables `json:"variables,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsTestResultDetail instantiates a new SyntheticsTestResultDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsTestResultDetail() *SyntheticsTestResultDetail {
	this := SyntheticsTestResultDetail{}
	return &this
}

// NewSyntheticsTestResultDetailWithDefaults instantiates a new SyntheticsTestResultDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsTestResultDetailWithDefaults() *SyntheticsTestResultDetail {
	this := SyntheticsTestResultDetail{}
	return &this
}

// GetAssertions returns the Assertions field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetAssertions() []SyntheticsTestResultAssertionResult {
	if o == nil || o.Assertions == nil {
		var ret []SyntheticsTestResultAssertionResult
		return ret
	}
	return o.Assertions
}

// GetAssertionsOk returns a tuple with the Assertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetAssertionsOk() (*[]SyntheticsTestResultAssertionResult, bool) {
	if o == nil || o.Assertions == nil {
		return nil, false
	}
	return &o.Assertions, true
}

// HasAssertions returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasAssertions() bool {
	return o != nil && o.Assertions != nil
}

// SetAssertions gets a reference to the given []SyntheticsTestResultAssertionResult and assigns it to the Assertions field.
func (o *SyntheticsTestResultDetail) SetAssertions(v []SyntheticsTestResultAssertionResult) {
	o.Assertions = v
}

// GetBucketKeys returns the BucketKeys field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetBucketKeys() SyntheticsTestResultBucketKeys {
	if o == nil || o.BucketKeys == nil {
		var ret SyntheticsTestResultBucketKeys
		return ret
	}
	return *o.BucketKeys
}

// GetBucketKeysOk returns a tuple with the BucketKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetBucketKeysOk() (*SyntheticsTestResultBucketKeys, bool) {
	if o == nil || o.BucketKeys == nil {
		return nil, false
	}
	return o.BucketKeys, true
}

// HasBucketKeys returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasBucketKeys() bool {
	return o != nil && o.BucketKeys != nil
}

// SetBucketKeys gets a reference to the given SyntheticsTestResultBucketKeys and assigns it to the BucketKeys field.
func (o *SyntheticsTestResultDetail) SetBucketKeys(v SyntheticsTestResultBucketKeys) {
	o.BucketKeys = &v
}

// GetCallType returns the CallType field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetCallType() string {
	if o == nil || o.CallType == nil {
		var ret string
		return ret
	}
	return *o.CallType
}

// GetCallTypeOk returns a tuple with the CallType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetCallTypeOk() (*string, bool) {
	if o == nil || o.CallType == nil {
		return nil, false
	}
	return o.CallType, true
}

// HasCallType returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasCallType() bool {
	return o != nil && o.CallType != nil
}

// SetCallType gets a reference to the given string and assigns it to the CallType field.
func (o *SyntheticsTestResultDetail) SetCallType(v string) {
	o.CallType = &v
}

// GetCert returns the Cert field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetCert() SyntheticsTestResultCertificate {
	if o == nil || o.Cert == nil {
		var ret SyntheticsTestResultCertificate
		return ret
	}
	return *o.Cert
}

// GetCertOk returns a tuple with the Cert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetCertOk() (*SyntheticsTestResultCertificate, bool) {
	if o == nil || o.Cert == nil {
		return nil, false
	}
	return o.Cert, true
}

// HasCert returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasCert() bool {
	return o != nil && o.Cert != nil
}

// SetCert gets a reference to the given SyntheticsTestResultCertificate and assigns it to the Cert field.
func (o *SyntheticsTestResultDetail) SetCert(v SyntheticsTestResultCertificate) {
	o.Cert = &v
}

// GetCompressedJsonDescriptor returns the CompressedJsonDescriptor field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetCompressedJsonDescriptor() string {
	if o == nil || o.CompressedJsonDescriptor == nil {
		var ret string
		return ret
	}
	return *o.CompressedJsonDescriptor
}

// GetCompressedJsonDescriptorOk returns a tuple with the CompressedJsonDescriptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetCompressedJsonDescriptorOk() (*string, bool) {
	if o == nil || o.CompressedJsonDescriptor == nil {
		return nil, false
	}
	return o.CompressedJsonDescriptor, true
}

// HasCompressedJsonDescriptor returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasCompressedJsonDescriptor() bool {
	return o != nil && o.CompressedJsonDescriptor != nil
}

// SetCompressedJsonDescriptor gets a reference to the given string and assigns it to the CompressedJsonDescriptor field.
func (o *SyntheticsTestResultDetail) SetCompressedJsonDescriptor(v string) {
	o.CompressedJsonDescriptor = &v
}

// GetCompressedSteps returns the CompressedSteps field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetCompressedSteps() string {
	if o == nil || o.CompressedSteps == nil {
		var ret string
		return ret
	}
	return *o.CompressedSteps
}

// GetCompressedStepsOk returns a tuple with the CompressedSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetCompressedStepsOk() (*string, bool) {
	if o == nil || o.CompressedSteps == nil {
		return nil, false
	}
	return o.CompressedSteps, true
}

// HasCompressedSteps returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasCompressedSteps() bool {
	return o != nil && o.CompressedSteps != nil
}

// SetCompressedSteps gets a reference to the given string and assigns it to the CompressedSteps field.
func (o *SyntheticsTestResultDetail) SetCompressedSteps(v string) {
	o.CompressedSteps = &v
}

// GetConnectionOutcome returns the ConnectionOutcome field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetConnectionOutcome() string {
	if o == nil || o.ConnectionOutcome == nil {
		var ret string
		return ret
	}
	return *o.ConnectionOutcome
}

// GetConnectionOutcomeOk returns a tuple with the ConnectionOutcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetConnectionOutcomeOk() (*string, bool) {
	if o == nil || o.ConnectionOutcome == nil {
		return nil, false
	}
	return o.ConnectionOutcome, true
}

// HasConnectionOutcome returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasConnectionOutcome() bool {
	return o != nil && o.ConnectionOutcome != nil
}

// SetConnectionOutcome gets a reference to the given string and assigns it to the ConnectionOutcome field.
func (o *SyntheticsTestResultDetail) SetConnectionOutcome(v string) {
	o.ConnectionOutcome = &v
}

// GetDnsResolution returns the DnsResolution field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetDnsResolution() SyntheticsTestResultDnsResolution {
	if o == nil || o.DnsResolution == nil {
		var ret SyntheticsTestResultDnsResolution
		return ret
	}
	return *o.DnsResolution
}

// GetDnsResolutionOk returns a tuple with the DnsResolution field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetDnsResolutionOk() (*SyntheticsTestResultDnsResolution, bool) {
	if o == nil || o.DnsResolution == nil {
		return nil, false
	}
	return o.DnsResolution, true
}

// HasDnsResolution returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasDnsResolution() bool {
	return o != nil && o.DnsResolution != nil
}

// SetDnsResolution gets a reference to the given SyntheticsTestResultDnsResolution and assigns it to the DnsResolution field.
func (o *SyntheticsTestResultDetail) SetDnsResolution(v SyntheticsTestResultDnsResolution) {
	o.DnsResolution = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetDuration() float64 {
	if o == nil || o.Duration == nil {
		var ret float64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetDurationOk() (*float64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given float64 and assigns it to the Duration field.
func (o *SyntheticsTestResultDetail) SetDuration(v float64) {
	o.Duration = &v
}

// GetExitedOnStepSuccess returns the ExitedOnStepSuccess field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetExitedOnStepSuccess() bool {
	if o == nil || o.ExitedOnStepSuccess == nil {
		var ret bool
		return ret
	}
	return *o.ExitedOnStepSuccess
}

// GetExitedOnStepSuccessOk returns a tuple with the ExitedOnStepSuccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetExitedOnStepSuccessOk() (*bool, bool) {
	if o == nil || o.ExitedOnStepSuccess == nil {
		return nil, false
	}
	return o.ExitedOnStepSuccess, true
}

// HasExitedOnStepSuccess returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasExitedOnStepSuccess() bool {
	return o != nil && o.ExitedOnStepSuccess != nil
}

// SetExitedOnStepSuccess gets a reference to the given bool and assigns it to the ExitedOnStepSuccess field.
func (o *SyntheticsTestResultDetail) SetExitedOnStepSuccess(v bool) {
	o.ExitedOnStepSuccess = &v
}

// GetFailure returns the Failure field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetFailure() SyntheticsTestResultFailure {
	if o == nil || o.Failure == nil {
		var ret SyntheticsTestResultFailure
		return ret
	}
	return *o.Failure
}

// GetFailureOk returns a tuple with the Failure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetFailureOk() (*SyntheticsTestResultFailure, bool) {
	if o == nil || o.Failure == nil {
		return nil, false
	}
	return o.Failure, true
}

// HasFailure returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasFailure() bool {
	return o != nil && o.Failure != nil
}

// SetFailure gets a reference to the given SyntheticsTestResultFailure and assigns it to the Failure field.
func (o *SyntheticsTestResultDetail) SetFailure(v SyntheticsTestResultFailure) {
	o.Failure = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetFinishedAt() int64 {
	if o == nil || o.FinishedAt == nil {
		var ret int64
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetFinishedAtOk() (*int64, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasFinishedAt() bool {
	return o != nil && o.FinishedAt != nil
}

// SetFinishedAt gets a reference to the given int64 and assigns it to the FinishedAt field.
func (o *SyntheticsTestResultDetail) SetFinishedAt(v int64) {
	o.FinishedAt = &v
}

// GetHandshake returns the Handshake field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetHandshake() SyntheticsTestResultHandshake {
	if o == nil || o.Handshake == nil {
		var ret SyntheticsTestResultHandshake
		return ret
	}
	return *o.Handshake
}

// GetHandshakeOk returns a tuple with the Handshake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetHandshakeOk() (*SyntheticsTestResultHandshake, bool) {
	if o == nil || o.Handshake == nil {
		return nil, false
	}
	return o.Handshake, true
}

// HasHandshake returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasHandshake() bool {
	return o != nil && o.Handshake != nil
}

// SetHandshake gets a reference to the given SyntheticsTestResultHandshake and assigns it to the Handshake field.
func (o *SyntheticsTestResultDetail) SetHandshake(v SyntheticsTestResultHandshake) {
	o.Handshake = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsTestResultDetail) SetId(v string) {
	o.Id = &v
}

// GetInitialId returns the InitialId field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetInitialId() string {
	if o == nil || o.InitialId == nil {
		var ret string
		return ret
	}
	return *o.InitialId
}

// GetInitialIdOk returns a tuple with the InitialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetInitialIdOk() (*string, bool) {
	if o == nil || o.InitialId == nil {
		return nil, false
	}
	return o.InitialId, true
}

// HasInitialId returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasInitialId() bool {
	return o != nil && o.InitialId != nil
}

// SetInitialId gets a reference to the given string and assigns it to the InitialId field.
func (o *SyntheticsTestResultDetail) SetInitialId(v string) {
	o.InitialId = &v
}

// GetIsFastRetry returns the IsFastRetry field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetIsFastRetry() bool {
	if o == nil || o.IsFastRetry == nil {
		var ret bool
		return ret
	}
	return *o.IsFastRetry
}

// GetIsFastRetryOk returns a tuple with the IsFastRetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetIsFastRetryOk() (*bool, bool) {
	if o == nil || o.IsFastRetry == nil {
		return nil, false
	}
	return o.IsFastRetry, true
}

// HasIsFastRetry returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasIsFastRetry() bool {
	return o != nil && o.IsFastRetry != nil
}

// SetIsFastRetry gets a reference to the given bool and assigns it to the IsFastRetry field.
func (o *SyntheticsTestResultDetail) SetIsFastRetry(v bool) {
	o.IsFastRetry = &v
}

// GetIsLastRetry returns the IsLastRetry field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetIsLastRetry() bool {
	if o == nil || o.IsLastRetry == nil {
		var ret bool
		return ret
	}
	return *o.IsLastRetry
}

// GetIsLastRetryOk returns a tuple with the IsLastRetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetIsLastRetryOk() (*bool, bool) {
	if o == nil || o.IsLastRetry == nil {
		return nil, false
	}
	return o.IsLastRetry, true
}

// HasIsLastRetry returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasIsLastRetry() bool {
	return o != nil && o.IsLastRetry != nil
}

// SetIsLastRetry gets a reference to the given bool and assigns it to the IsLastRetry field.
func (o *SyntheticsTestResultDetail) SetIsLastRetry(v bool) {
	o.IsLastRetry = &v
}

// GetNetpath returns the Netpath field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetNetpath() SyntheticsTestResultNetpath {
	if o == nil || o.Netpath == nil {
		var ret SyntheticsTestResultNetpath
		return ret
	}
	return *o.Netpath
}

// GetNetpathOk returns a tuple with the Netpath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetNetpathOk() (*SyntheticsTestResultNetpath, bool) {
	if o == nil || o.Netpath == nil {
		return nil, false
	}
	return o.Netpath, true
}

// HasNetpath returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasNetpath() bool {
	return o != nil && o.Netpath != nil
}

// SetNetpath gets a reference to the given SyntheticsTestResultNetpath and assigns it to the Netpath field.
func (o *SyntheticsTestResultDetail) SetNetpath(v SyntheticsTestResultNetpath) {
	o.Netpath = &v
}

// GetNetstats returns the Netstats field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetNetstats() SyntheticsTestResultNetstats {
	if o == nil || o.Netstats == nil {
		var ret SyntheticsTestResultNetstats
		return ret
	}
	return *o.Netstats
}

// GetNetstatsOk returns a tuple with the Netstats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetNetstatsOk() (*SyntheticsTestResultNetstats, bool) {
	if o == nil || o.Netstats == nil {
		return nil, false
	}
	return o.Netstats, true
}

// HasNetstats returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasNetstats() bool {
	return o != nil && o.Netstats != nil
}

// SetNetstats gets a reference to the given SyntheticsTestResultNetstats and assigns it to the Netstats field.
func (o *SyntheticsTestResultDetail) SetNetstats(v SyntheticsTestResultNetstats) {
	o.Netstats = &v
}

// GetOcsp returns the Ocsp field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetOcsp() SyntheticsTestResultOCSPResponse {
	if o == nil || o.Ocsp == nil {
		var ret SyntheticsTestResultOCSPResponse
		return ret
	}
	return *o.Ocsp
}

// GetOcspOk returns a tuple with the Ocsp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetOcspOk() (*SyntheticsTestResultOCSPResponse, bool) {
	if o == nil || o.Ocsp == nil {
		return nil, false
	}
	return o.Ocsp, true
}

// HasOcsp returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasOcsp() bool {
	return o != nil && o.Ocsp != nil
}

// SetOcsp gets a reference to the given SyntheticsTestResultOCSPResponse and assigns it to the Ocsp field.
func (o *SyntheticsTestResultDetail) SetOcsp(v SyntheticsTestResultOCSPResponse) {
	o.Ocsp = &v
}

// GetPing returns the Ping field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetPing() SyntheticsTestResultTracerouteHop {
	if o == nil || o.Ping == nil {
		var ret SyntheticsTestResultTracerouteHop
		return ret
	}
	return *o.Ping
}

// GetPingOk returns a tuple with the Ping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetPingOk() (*SyntheticsTestResultTracerouteHop, bool) {
	if o == nil || o.Ping == nil {
		return nil, false
	}
	return o.Ping, true
}

// HasPing returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasPing() bool {
	return o != nil && o.Ping != nil
}

// SetPing gets a reference to the given SyntheticsTestResultTracerouteHop and assigns it to the Ping field.
func (o *SyntheticsTestResultDetail) SetPing(v SyntheticsTestResultTracerouteHop) {
	o.Ping = &v
}

// GetReceivedEmailCount returns the ReceivedEmailCount field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetReceivedEmailCount() int64 {
	if o == nil || o.ReceivedEmailCount == nil {
		var ret int64
		return ret
	}
	return *o.ReceivedEmailCount
}

// GetReceivedEmailCountOk returns a tuple with the ReceivedEmailCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetReceivedEmailCountOk() (*int64, bool) {
	if o == nil || o.ReceivedEmailCount == nil {
		return nil, false
	}
	return o.ReceivedEmailCount, true
}

// HasReceivedEmailCount returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasReceivedEmailCount() bool {
	return o != nil && o.ReceivedEmailCount != nil
}

// SetReceivedEmailCount gets a reference to the given int64 and assigns it to the ReceivedEmailCount field.
func (o *SyntheticsTestResultDetail) SetReceivedEmailCount(v int64) {
	o.ReceivedEmailCount = &v
}

// GetReceivedMessage returns the ReceivedMessage field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetReceivedMessage() string {
	if o == nil || o.ReceivedMessage == nil {
		var ret string
		return ret
	}
	return *o.ReceivedMessage
}

// GetReceivedMessageOk returns a tuple with the ReceivedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetReceivedMessageOk() (*string, bool) {
	if o == nil || o.ReceivedMessage == nil {
		return nil, false
	}
	return o.ReceivedMessage, true
}

// HasReceivedMessage returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasReceivedMessage() bool {
	return o != nil && o.ReceivedMessage != nil
}

// SetReceivedMessage gets a reference to the given string and assigns it to the ReceivedMessage field.
func (o *SyntheticsTestResultDetail) SetReceivedMessage(v string) {
	o.ReceivedMessage = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetRequest() SyntheticsTestResultRequestInfo {
	if o == nil || o.Request == nil {
		var ret SyntheticsTestResultRequestInfo
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetRequestOk() (*SyntheticsTestResultRequestInfo, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasRequest() bool {
	return o != nil && o.Request != nil
}

// SetRequest gets a reference to the given SyntheticsTestResultRequestInfo and assigns it to the Request field.
func (o *SyntheticsTestResultDetail) SetRequest(v SyntheticsTestResultRequestInfo) {
	o.Request = &v
}

// GetResolvedIp returns the ResolvedIp field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetResolvedIp() string {
	if o == nil || o.ResolvedIp == nil {
		var ret string
		return ret
	}
	return *o.ResolvedIp
}

// GetResolvedIpOk returns a tuple with the ResolvedIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetResolvedIpOk() (*string, bool) {
	if o == nil || o.ResolvedIp == nil {
		return nil, false
	}
	return o.ResolvedIp, true
}

// HasResolvedIp returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasResolvedIp() bool {
	return o != nil && o.ResolvedIp != nil
}

// SetResolvedIp gets a reference to the given string and assigns it to the ResolvedIp field.
func (o *SyntheticsTestResultDetail) SetResolvedIp(v string) {
	o.ResolvedIp = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetResponse() SyntheticsTestResultResponseInfo {
	if o == nil || o.Response == nil {
		var ret SyntheticsTestResultResponseInfo
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetResponseOk() (*SyntheticsTestResultResponseInfo, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasResponse() bool {
	return o != nil && o.Response != nil
}

// SetResponse gets a reference to the given SyntheticsTestResultResponseInfo and assigns it to the Response field.
func (o *SyntheticsTestResultDetail) SetResponse(v SyntheticsTestResultResponseInfo) {
	o.Response = &v
}

// GetRunType returns the RunType field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetRunType() SyntheticsTestResultRunType {
	if o == nil || o.RunType == nil {
		var ret SyntheticsTestResultRunType
		return ret
	}
	return *o.RunType
}

// GetRunTypeOk returns a tuple with the RunType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetRunTypeOk() (*SyntheticsTestResultRunType, bool) {
	if o == nil || o.RunType == nil {
		return nil, false
	}
	return o.RunType, true
}

// HasRunType returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasRunType() bool {
	return o != nil && o.RunType != nil
}

// SetRunType gets a reference to the given SyntheticsTestResultRunType and assigns it to the RunType field.
func (o *SyntheticsTestResultDetail) SetRunType(v SyntheticsTestResultRunType) {
	o.RunType = &v
}

// GetSentMessage returns the SentMessage field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetSentMessage() string {
	if o == nil || o.SentMessage == nil {
		var ret string
		return ret
	}
	return *o.SentMessage
}

// GetSentMessageOk returns a tuple with the SentMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetSentMessageOk() (*string, bool) {
	if o == nil || o.SentMessage == nil {
		return nil, false
	}
	return o.SentMessage, true
}

// HasSentMessage returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasSentMessage() bool {
	return o != nil && o.SentMessage != nil
}

// SetSentMessage gets a reference to the given string and assigns it to the SentMessage field.
func (o *SyntheticsTestResultDetail) SetSentMessage(v string) {
	o.SentMessage = &v
}

// GetStartUrl returns the StartUrl field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetStartUrl() string {
	if o == nil || o.StartUrl == nil {
		var ret string
		return ret
	}
	return *o.StartUrl
}

// GetStartUrlOk returns a tuple with the StartUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetStartUrlOk() (*string, bool) {
	if o == nil || o.StartUrl == nil {
		return nil, false
	}
	return o.StartUrl, true
}

// HasStartUrl returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasStartUrl() bool {
	return o != nil && o.StartUrl != nil
}

// SetStartUrl gets a reference to the given string and assigns it to the StartUrl field.
func (o *SyntheticsTestResultDetail) SetStartUrl(v string) {
	o.StartUrl = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetStartedAt() int64 {
	if o == nil || o.StartedAt == nil {
		var ret int64
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetStartedAtOk() (*int64, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given int64 and assigns it to the StartedAt field.
func (o *SyntheticsTestResultDetail) SetStartedAt(v int64) {
	o.StartedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetStatus() SyntheticsTestResultStatus {
	if o == nil || o.Status == nil {
		var ret SyntheticsTestResultStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetStatusOk() (*SyntheticsTestResultStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given SyntheticsTestResultStatus and assigns it to the Status field.
func (o *SyntheticsTestResultDetail) SetStatus(v SyntheticsTestResultStatus) {
	o.Status = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetSteps() []SyntheticsTestResultStep {
	if o == nil || o.Steps == nil {
		var ret []SyntheticsTestResultStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetStepsOk() (*[]SyntheticsTestResultStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []SyntheticsTestResultStep and assigns it to the Steps field.
func (o *SyntheticsTestResultDetail) SetSteps(v []SyntheticsTestResultStep) {
	o.Steps = v
}

// GetTimeToInteractive returns the TimeToInteractive field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetTimeToInteractive() int64 {
	if o == nil || o.TimeToInteractive == nil {
		var ret int64
		return ret
	}
	return *o.TimeToInteractive
}

// GetTimeToInteractiveOk returns a tuple with the TimeToInteractive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetTimeToInteractiveOk() (*int64, bool) {
	if o == nil || o.TimeToInteractive == nil {
		return nil, false
	}
	return o.TimeToInteractive, true
}

// HasTimeToInteractive returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasTimeToInteractive() bool {
	return o != nil && o.TimeToInteractive != nil
}

// SetTimeToInteractive gets a reference to the given int64 and assigns it to the TimeToInteractive field.
func (o *SyntheticsTestResultDetail) SetTimeToInteractive(v int64) {
	o.TimeToInteractive = &v
}

// GetTimings returns the Timings field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetTimings() map[string]interface{} {
	if o == nil || o.Timings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Timings
}

// GetTimingsOk returns a tuple with the Timings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetTimingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Timings == nil {
		return nil, false
	}
	return &o.Timings, true
}

// HasTimings returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasTimings() bool {
	return o != nil && o.Timings != nil
}

// SetTimings gets a reference to the given map[string]interface{} and assigns it to the Timings field.
func (o *SyntheticsTestResultDetail) SetTimings(v map[string]interface{}) {
	o.Timings = v
}

// GetTrace returns the Trace field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetTrace() SyntheticsTestResultTrace {
	if o == nil || o.Trace == nil {
		var ret SyntheticsTestResultTrace
		return ret
	}
	return *o.Trace
}

// GetTraceOk returns a tuple with the Trace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetTraceOk() (*SyntheticsTestResultTrace, bool) {
	if o == nil || o.Trace == nil {
		return nil, false
	}
	return o.Trace, true
}

// HasTrace returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasTrace() bool {
	return o != nil && o.Trace != nil
}

// SetTrace gets a reference to the given SyntheticsTestResultTrace and assigns it to the Trace field.
func (o *SyntheticsTestResultDetail) SetTrace(v SyntheticsTestResultTrace) {
	o.Trace = &v
}

// GetTraceroute returns the Traceroute field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetTraceroute() []SyntheticsTestResultTracerouteHop {
	if o == nil || o.Traceroute == nil {
		var ret []SyntheticsTestResultTracerouteHop
		return ret
	}
	return o.Traceroute
}

// GetTracerouteOk returns a tuple with the Traceroute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetTracerouteOk() (*[]SyntheticsTestResultTracerouteHop, bool) {
	if o == nil || o.Traceroute == nil {
		return nil, false
	}
	return &o.Traceroute, true
}

// HasTraceroute returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasTraceroute() bool {
	return o != nil && o.Traceroute != nil
}

// SetTraceroute gets a reference to the given []SyntheticsTestResultTracerouteHop and assigns it to the Traceroute field.
func (o *SyntheticsTestResultDetail) SetTraceroute(v []SyntheticsTestResultTracerouteHop) {
	o.Traceroute = v
}

// GetTriggeredAt returns the TriggeredAt field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetTriggeredAt() int64 {
	if o == nil || o.TriggeredAt == nil {
		var ret int64
		return ret
	}
	return *o.TriggeredAt
}

// GetTriggeredAtOk returns a tuple with the TriggeredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetTriggeredAtOk() (*int64, bool) {
	if o == nil || o.TriggeredAt == nil {
		return nil, false
	}
	return o.TriggeredAt, true
}

// HasTriggeredAt returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasTriggeredAt() bool {
	return o != nil && o.TriggeredAt != nil
}

// SetTriggeredAt gets a reference to the given int64 and assigns it to the TriggeredAt field.
func (o *SyntheticsTestResultDetail) SetTriggeredAt(v int64) {
	o.TriggeredAt = &v
}

// GetTunnel returns the Tunnel field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetTunnel() bool {
	if o == nil || o.Tunnel == nil {
		var ret bool
		return ret
	}
	return *o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetTunnelOk() (*bool, bool) {
	if o == nil || o.Tunnel == nil {
		return nil, false
	}
	return o.Tunnel, true
}

// HasTunnel returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasTunnel() bool {
	return o != nil && o.Tunnel != nil
}

// SetTunnel gets a reference to the given bool and assigns it to the Tunnel field.
func (o *SyntheticsTestResultDetail) SetTunnel(v bool) {
	o.Tunnel = &v
}

// GetTurns returns the Turns field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetTurns() []SyntheticsTestResultTurn {
	if o == nil || o.Turns == nil {
		var ret []SyntheticsTestResultTurn
		return ret
	}
	return o.Turns
}

// GetTurnsOk returns a tuple with the Turns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetTurnsOk() (*[]SyntheticsTestResultTurn, bool) {
	if o == nil || o.Turns == nil {
		return nil, false
	}
	return &o.Turns, true
}

// HasTurns returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasTurns() bool {
	return o != nil && o.Turns != nil
}

// SetTurns gets a reference to the given []SyntheticsTestResultTurn and assigns it to the Turns field.
func (o *SyntheticsTestResultDetail) SetTurns(v []SyntheticsTestResultTurn) {
	o.Turns = v
}

// GetUnhealthy returns the Unhealthy field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetUnhealthy() bool {
	if o == nil || o.Unhealthy == nil {
		var ret bool
		return ret
	}
	return *o.Unhealthy
}

// GetUnhealthyOk returns a tuple with the Unhealthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetUnhealthyOk() (*bool, bool) {
	if o == nil || o.Unhealthy == nil {
		return nil, false
	}
	return o.Unhealthy, true
}

// HasUnhealthy returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasUnhealthy() bool {
	return o != nil && o.Unhealthy != nil
}

// SetUnhealthy gets a reference to the given bool and assigns it to the Unhealthy field.
func (o *SyntheticsTestResultDetail) SetUnhealthy(v bool) {
	o.Unhealthy = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *SyntheticsTestResultDetail) GetVariables() SyntheticsTestResultVariables {
	if o == nil || o.Variables == nil {
		var ret SyntheticsTestResultVariables
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsTestResultDetail) GetVariablesOk() (*SyntheticsTestResultVariables, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *SyntheticsTestResultDetail) HasVariables() bool {
	return o != nil && o.Variables != nil
}

// SetVariables gets a reference to the given SyntheticsTestResultVariables and assigns it to the Variables field.
func (o *SyntheticsTestResultDetail) SetVariables(v SyntheticsTestResultVariables) {
	o.Variables = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsTestResultDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assertions != nil {
		toSerialize["assertions"] = o.Assertions
	}
	if o.BucketKeys != nil {
		toSerialize["bucket_keys"] = o.BucketKeys
	}
	if o.CallType != nil {
		toSerialize["call_type"] = o.CallType
	}
	if o.Cert != nil {
		toSerialize["cert"] = o.Cert
	}
	if o.CompressedJsonDescriptor != nil {
		toSerialize["compressed_json_descriptor"] = o.CompressedJsonDescriptor
	}
	if o.CompressedSteps != nil {
		toSerialize["compressed_steps"] = o.CompressedSteps
	}
	if o.ConnectionOutcome != nil {
		toSerialize["connection_outcome"] = o.ConnectionOutcome
	}
	if o.DnsResolution != nil {
		toSerialize["dns_resolution"] = o.DnsResolution
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.ExitedOnStepSuccess != nil {
		toSerialize["exited_on_step_success"] = o.ExitedOnStepSuccess
	}
	if o.Failure != nil {
		toSerialize["failure"] = o.Failure
	}
	if o.FinishedAt != nil {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if o.Handshake != nil {
		toSerialize["handshake"] = o.Handshake
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.InitialId != nil {
		toSerialize["initial_id"] = o.InitialId
	}
	if o.IsFastRetry != nil {
		toSerialize["is_fast_retry"] = o.IsFastRetry
	}
	if o.IsLastRetry != nil {
		toSerialize["is_last_retry"] = o.IsLastRetry
	}
	if o.Netpath != nil {
		toSerialize["netpath"] = o.Netpath
	}
	if o.Netstats != nil {
		toSerialize["netstats"] = o.Netstats
	}
	if o.Ocsp != nil {
		toSerialize["ocsp"] = o.Ocsp
	}
	if o.Ping != nil {
		toSerialize["ping"] = o.Ping
	}
	if o.ReceivedEmailCount != nil {
		toSerialize["received_email_count"] = o.ReceivedEmailCount
	}
	if o.ReceivedMessage != nil {
		toSerialize["received_message"] = o.ReceivedMessage
	}
	if o.Request != nil {
		toSerialize["request"] = o.Request
	}
	if o.ResolvedIp != nil {
		toSerialize["resolved_ip"] = o.ResolvedIp
	}
	if o.Response != nil {
		toSerialize["response"] = o.Response
	}
	if o.RunType != nil {
		toSerialize["run_type"] = o.RunType
	}
	if o.SentMessage != nil {
		toSerialize["sent_message"] = o.SentMessage
	}
	if o.StartUrl != nil {
		toSerialize["start_url"] = o.StartUrl
	}
	if o.StartedAt != nil {
		toSerialize["started_at"] = o.StartedAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.TimeToInteractive != nil {
		toSerialize["time_to_interactive"] = o.TimeToInteractive
	}
	if o.Timings != nil {
		toSerialize["timings"] = o.Timings
	}
	if o.Trace != nil {
		toSerialize["trace"] = o.Trace
	}
	if o.Traceroute != nil {
		toSerialize["traceroute"] = o.Traceroute
	}
	if o.TriggeredAt != nil {
		toSerialize["triggered_at"] = o.TriggeredAt
	}
	if o.Tunnel != nil {
		toSerialize["tunnel"] = o.Tunnel
	}
	if o.Turns != nil {
		toSerialize["turns"] = o.Turns
	}
	if o.Unhealthy != nil {
		toSerialize["unhealthy"] = o.Unhealthy
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsTestResultDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assertions               []SyntheticsTestResultAssertionResult `json:"assertions,omitempty"`
		BucketKeys               *SyntheticsTestResultBucketKeys       `json:"bucket_keys,omitempty"`
		CallType                 *string                               `json:"call_type,omitempty"`
		Cert                     *SyntheticsTestResultCertificate      `json:"cert,omitempty"`
		CompressedJsonDescriptor *string                               `json:"compressed_json_descriptor,omitempty"`
		CompressedSteps          *string                               `json:"compressed_steps,omitempty"`
		ConnectionOutcome        *string                               `json:"connection_outcome,omitempty"`
		DnsResolution            *SyntheticsTestResultDnsResolution    `json:"dns_resolution,omitempty"`
		Duration                 *float64                              `json:"duration,omitempty"`
		ExitedOnStepSuccess      *bool                                 `json:"exited_on_step_success,omitempty"`
		Failure                  *SyntheticsTestResultFailure          `json:"failure,omitempty"`
		FinishedAt               *int64                                `json:"finished_at,omitempty"`
		Handshake                *SyntheticsTestResultHandshake        `json:"handshake,omitempty"`
		Id                       *string                               `json:"id,omitempty"`
		InitialId                *string                               `json:"initial_id,omitempty"`
		IsFastRetry              *bool                                 `json:"is_fast_retry,omitempty"`
		IsLastRetry              *bool                                 `json:"is_last_retry,omitempty"`
		Netpath                  *SyntheticsTestResultNetpath          `json:"netpath,omitempty"`
		Netstats                 *SyntheticsTestResultNetstats         `json:"netstats,omitempty"`
		Ocsp                     *SyntheticsTestResultOCSPResponse     `json:"ocsp,omitempty"`
		Ping                     *SyntheticsTestResultTracerouteHop    `json:"ping,omitempty"`
		ReceivedEmailCount       *int64                                `json:"received_email_count,omitempty"`
		ReceivedMessage          *string                               `json:"received_message,omitempty"`
		Request                  *SyntheticsTestResultRequestInfo      `json:"request,omitempty"`
		ResolvedIp               *string                               `json:"resolved_ip,omitempty"`
		Response                 *SyntheticsTestResultResponseInfo     `json:"response,omitempty"`
		RunType                  *SyntheticsTestResultRunType          `json:"run_type,omitempty"`
		SentMessage              *string                               `json:"sent_message,omitempty"`
		StartUrl                 *string                               `json:"start_url,omitempty"`
		StartedAt                *int64                                `json:"started_at,omitempty"`
		Status                   *SyntheticsTestResultStatus           `json:"status,omitempty"`
		Steps                    []SyntheticsTestResultStep            `json:"steps,omitempty"`
		TimeToInteractive        *int64                                `json:"time_to_interactive,omitempty"`
		Timings                  map[string]interface{}                `json:"timings,omitempty"`
		Trace                    *SyntheticsTestResultTrace            `json:"trace,omitempty"`
		Traceroute               []SyntheticsTestResultTracerouteHop   `json:"traceroute,omitempty"`
		TriggeredAt              *int64                                `json:"triggered_at,omitempty"`
		Tunnel                   *bool                                 `json:"tunnel,omitempty"`
		Turns                    []SyntheticsTestResultTurn            `json:"turns,omitempty"`
		Unhealthy                *bool                                 `json:"unhealthy,omitempty"`
		Variables                *SyntheticsTestResultVariables        `json:"variables,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assertions", "bucket_keys", "call_type", "cert", "compressed_json_descriptor", "compressed_steps", "connection_outcome", "dns_resolution", "duration", "exited_on_step_success", "failure", "finished_at", "handshake", "id", "initial_id", "is_fast_retry", "is_last_retry", "netpath", "netstats", "ocsp", "ping", "received_email_count", "received_message", "request", "resolved_ip", "response", "run_type", "sent_message", "start_url", "started_at", "status", "steps", "time_to_interactive", "timings", "trace", "traceroute", "triggered_at", "tunnel", "turns", "unhealthy", "variables"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Assertions = all.Assertions
	if all.BucketKeys != nil && all.BucketKeys.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.BucketKeys = all.BucketKeys
	o.CallType = all.CallType
	if all.Cert != nil && all.Cert.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cert = all.Cert
	o.CompressedJsonDescriptor = all.CompressedJsonDescriptor
	o.CompressedSteps = all.CompressedSteps
	o.ConnectionOutcome = all.ConnectionOutcome
	if all.DnsResolution != nil && all.DnsResolution.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.DnsResolution = all.DnsResolution
	o.Duration = all.Duration
	o.ExitedOnStepSuccess = all.ExitedOnStepSuccess
	if all.Failure != nil && all.Failure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Failure = all.Failure
	o.FinishedAt = all.FinishedAt
	if all.Handshake != nil && all.Handshake.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Handshake = all.Handshake
	o.Id = all.Id
	o.InitialId = all.InitialId
	o.IsFastRetry = all.IsFastRetry
	o.IsLastRetry = all.IsLastRetry
	if all.Netpath != nil && all.Netpath.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Netpath = all.Netpath
	if all.Netstats != nil && all.Netstats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Netstats = all.Netstats
	if all.Ocsp != nil && all.Ocsp.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Ocsp = all.Ocsp
	if all.Ping != nil && all.Ping.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Ping = all.Ping
	o.ReceivedEmailCount = all.ReceivedEmailCount
	o.ReceivedMessage = all.ReceivedMessage
	if all.Request != nil && all.Request.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Request = all.Request
	o.ResolvedIp = all.ResolvedIp
	if all.Response != nil && all.Response.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Response = all.Response
	if all.RunType != nil && !all.RunType.IsValid() {
		hasInvalidField = true
	} else {
		o.RunType = all.RunType
	}
	o.SentMessage = all.SentMessage
	o.StartUrl = all.StartUrl
	o.StartedAt = all.StartedAt
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Steps = all.Steps
	o.TimeToInteractive = all.TimeToInteractive
	o.Timings = all.Timings
	if all.Trace != nil && all.Trace.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Trace = all.Trace
	o.Traceroute = all.Traceroute
	o.TriggeredAt = all.TriggeredAt
	o.Tunnel = all.Tunnel
	o.Turns = all.Turns
	o.Unhealthy = all.Unhealthy
	if all.Variables != nil && all.Variables.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Variables = all.Variables

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
