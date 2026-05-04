// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SyntheticsFastTestResultDetail Detailed result data for the fast test run. The exact shape of nested fields
// (`request`, `response`, `assertions`, etc.) depends on the test subtype.
type SyntheticsFastTestResultDetail struct {
	// Results of each assertion evaluated during the test.
	Assertions []SyntheticsTestResultAssertionResult `json:"assertions,omitempty"`
	// gRPC call type (for example, `unary`, `healthCheck`, or `reflection`).
	CallType *string `json:"call_type,omitempty"`
	// SSL/TLS certificate information returned from an SSL test.
	Cert *SyntheticsTestResultCertificate `json:"cert,omitempty"`
	// Total duration of the test in milliseconds.
	Duration *float64 `json:"duration,omitempty"`
	// Details about the failure of a Synthetic test.
	Failure *SyntheticsTestResultFailure `json:"failure,omitempty"`
	// Unix timestamp (ms) of when the test finished.
	FinishedAt *int64 `json:"finished_at,omitempty"`
	// The result ID. Set to the fast test UUID because no persistent result ID exists for fast tests.
	Id *string `json:"id,omitempty"`
	// Whether this result is from an automatic fast retry.
	IsFastRetry *bool `json:"is_fast_retry,omitempty"`
	// Details of the outgoing request made during the test execution.
	Request *SyntheticsTestResultRequestInfo `json:"request,omitempty"`
	// IP address resolved for the target host.
	ResolvedIp *string `json:"resolved_ip,omitempty"`
	// Details of the response received during the test execution.
	Response *SyntheticsTestResultResponseInfo `json:"response,omitempty"`
	// The type of run for a Synthetic test result.
	RunType *SyntheticsTestResultRunType `json:"run_type,omitempty"`
	// Unix timestamp (ms) of when the test started.
	StartedAt *int64 `json:"started_at,omitempty"`
	// Status of the test result (`passed` or `failed`).
	Status *string `json:"status,omitempty"`
	// Step results for multistep API tests.
	Steps []SyntheticsTestResultStep `json:"steps,omitempty"`
	// Timing breakdown of the test request phases (for example, DNS, TCP, TLS, first byte).
	Timings map[string]interface{} `json:"timings,omitempty"`
	// Traceroute hop results, present for ICMP and TCP tests.
	Traceroute []SyntheticsTestResultTracerouteHop `json:"traceroute,omitempty"`
	// Unix timestamp (ms) of when the test was triggered.
	TriggeredAt *int64 `json:"triggered_at,omitempty"`
	// Whether the test was run through a Synthetics tunnel.
	Tunnel *bool `json:"tunnel,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSyntheticsFastTestResultDetail instantiates a new SyntheticsFastTestResultDetail object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSyntheticsFastTestResultDetail() *SyntheticsFastTestResultDetail {
	this := SyntheticsFastTestResultDetail{}
	return &this
}

// NewSyntheticsFastTestResultDetailWithDefaults instantiates a new SyntheticsFastTestResultDetail object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSyntheticsFastTestResultDetailWithDefaults() *SyntheticsFastTestResultDetail {
	this := SyntheticsFastTestResultDetail{}
	return &this
}

// GetAssertions returns the Assertions field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetAssertions() []SyntheticsTestResultAssertionResult {
	if o == nil || o.Assertions == nil {
		var ret []SyntheticsTestResultAssertionResult
		return ret
	}
	return o.Assertions
}

// GetAssertionsOk returns a tuple with the Assertions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetAssertionsOk() (*[]SyntheticsTestResultAssertionResult, bool) {
	if o == nil || o.Assertions == nil {
		return nil, false
	}
	return &o.Assertions, true
}

// HasAssertions returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasAssertions() bool {
	return o != nil && o.Assertions != nil
}

// SetAssertions gets a reference to the given []SyntheticsTestResultAssertionResult and assigns it to the Assertions field.
func (o *SyntheticsFastTestResultDetail) SetAssertions(v []SyntheticsTestResultAssertionResult) {
	o.Assertions = v
}

// GetCallType returns the CallType field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetCallType() string {
	if o == nil || o.CallType == nil {
		var ret string
		return ret
	}
	return *o.CallType
}

// GetCallTypeOk returns a tuple with the CallType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetCallTypeOk() (*string, bool) {
	if o == nil || o.CallType == nil {
		return nil, false
	}
	return o.CallType, true
}

// HasCallType returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasCallType() bool {
	return o != nil && o.CallType != nil
}

// SetCallType gets a reference to the given string and assigns it to the CallType field.
func (o *SyntheticsFastTestResultDetail) SetCallType(v string) {
	o.CallType = &v
}

// GetCert returns the Cert field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetCert() SyntheticsTestResultCertificate {
	if o == nil || o.Cert == nil {
		var ret SyntheticsTestResultCertificate
		return ret
	}
	return *o.Cert
}

// GetCertOk returns a tuple with the Cert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetCertOk() (*SyntheticsTestResultCertificate, bool) {
	if o == nil || o.Cert == nil {
		return nil, false
	}
	return o.Cert, true
}

// HasCert returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasCert() bool {
	return o != nil && o.Cert != nil
}

// SetCert gets a reference to the given SyntheticsTestResultCertificate and assigns it to the Cert field.
func (o *SyntheticsFastTestResultDetail) SetCert(v SyntheticsTestResultCertificate) {
	o.Cert = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetDuration() float64 {
	if o == nil || o.Duration == nil {
		var ret float64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetDurationOk() (*float64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasDuration() bool {
	return o != nil && o.Duration != nil
}

// SetDuration gets a reference to the given float64 and assigns it to the Duration field.
func (o *SyntheticsFastTestResultDetail) SetDuration(v float64) {
	o.Duration = &v
}

// GetFailure returns the Failure field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetFailure() SyntheticsTestResultFailure {
	if o == nil || o.Failure == nil {
		var ret SyntheticsTestResultFailure
		return ret
	}
	return *o.Failure
}

// GetFailureOk returns a tuple with the Failure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetFailureOk() (*SyntheticsTestResultFailure, bool) {
	if o == nil || o.Failure == nil {
		return nil, false
	}
	return o.Failure, true
}

// HasFailure returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasFailure() bool {
	return o != nil && o.Failure != nil
}

// SetFailure gets a reference to the given SyntheticsTestResultFailure and assigns it to the Failure field.
func (o *SyntheticsFastTestResultDetail) SetFailure(v SyntheticsTestResultFailure) {
	o.Failure = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetFinishedAt() int64 {
	if o == nil || o.FinishedAt == nil {
		var ret int64
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetFinishedAtOk() (*int64, bool) {
	if o == nil || o.FinishedAt == nil {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasFinishedAt() bool {
	return o != nil && o.FinishedAt != nil
}

// SetFinishedAt gets a reference to the given int64 and assigns it to the FinishedAt field.
func (o *SyntheticsFastTestResultDetail) SetFinishedAt(v int64) {
	o.FinishedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SyntheticsFastTestResultDetail) SetId(v string) {
	o.Id = &v
}

// GetIsFastRetry returns the IsFastRetry field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetIsFastRetry() bool {
	if o == nil || o.IsFastRetry == nil {
		var ret bool
		return ret
	}
	return *o.IsFastRetry
}

// GetIsFastRetryOk returns a tuple with the IsFastRetry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetIsFastRetryOk() (*bool, bool) {
	if o == nil || o.IsFastRetry == nil {
		return nil, false
	}
	return o.IsFastRetry, true
}

// HasIsFastRetry returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasIsFastRetry() bool {
	return o != nil && o.IsFastRetry != nil
}

// SetIsFastRetry gets a reference to the given bool and assigns it to the IsFastRetry field.
func (o *SyntheticsFastTestResultDetail) SetIsFastRetry(v bool) {
	o.IsFastRetry = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetRequest() SyntheticsTestResultRequestInfo {
	if o == nil || o.Request == nil {
		var ret SyntheticsTestResultRequestInfo
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetRequestOk() (*SyntheticsTestResultRequestInfo, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasRequest() bool {
	return o != nil && o.Request != nil
}

// SetRequest gets a reference to the given SyntheticsTestResultRequestInfo and assigns it to the Request field.
func (o *SyntheticsFastTestResultDetail) SetRequest(v SyntheticsTestResultRequestInfo) {
	o.Request = &v
}

// GetResolvedIp returns the ResolvedIp field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetResolvedIp() string {
	if o == nil || o.ResolvedIp == nil {
		var ret string
		return ret
	}
	return *o.ResolvedIp
}

// GetResolvedIpOk returns a tuple with the ResolvedIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetResolvedIpOk() (*string, bool) {
	if o == nil || o.ResolvedIp == nil {
		return nil, false
	}
	return o.ResolvedIp, true
}

// HasResolvedIp returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasResolvedIp() bool {
	return o != nil && o.ResolvedIp != nil
}

// SetResolvedIp gets a reference to the given string and assigns it to the ResolvedIp field.
func (o *SyntheticsFastTestResultDetail) SetResolvedIp(v string) {
	o.ResolvedIp = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetResponse() SyntheticsTestResultResponseInfo {
	if o == nil || o.Response == nil {
		var ret SyntheticsTestResultResponseInfo
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetResponseOk() (*SyntheticsTestResultResponseInfo, bool) {
	if o == nil || o.Response == nil {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasResponse() bool {
	return o != nil && o.Response != nil
}

// SetResponse gets a reference to the given SyntheticsTestResultResponseInfo and assigns it to the Response field.
func (o *SyntheticsFastTestResultDetail) SetResponse(v SyntheticsTestResultResponseInfo) {
	o.Response = &v
}

// GetRunType returns the RunType field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetRunType() SyntheticsTestResultRunType {
	if o == nil || o.RunType == nil {
		var ret SyntheticsTestResultRunType
		return ret
	}
	return *o.RunType
}

// GetRunTypeOk returns a tuple with the RunType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetRunTypeOk() (*SyntheticsTestResultRunType, bool) {
	if o == nil || o.RunType == nil {
		return nil, false
	}
	return o.RunType, true
}

// HasRunType returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasRunType() bool {
	return o != nil && o.RunType != nil
}

// SetRunType gets a reference to the given SyntheticsTestResultRunType and assigns it to the RunType field.
func (o *SyntheticsFastTestResultDetail) SetRunType(v SyntheticsTestResultRunType) {
	o.RunType = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetStartedAt() int64 {
	if o == nil || o.StartedAt == nil {
		var ret int64
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetStartedAtOk() (*int64, bool) {
	if o == nil || o.StartedAt == nil {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasStartedAt() bool {
	return o != nil && o.StartedAt != nil
}

// SetStartedAt gets a reference to the given int64 and assigns it to the StartedAt field.
func (o *SyntheticsFastTestResultDetail) SetStartedAt(v int64) {
	o.StartedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SyntheticsFastTestResultDetail) SetStatus(v string) {
	o.Status = &v
}

// GetSteps returns the Steps field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetSteps() []SyntheticsTestResultStep {
	if o == nil || o.Steps == nil {
		var ret []SyntheticsTestResultStep
		return ret
	}
	return o.Steps
}

// GetStepsOk returns a tuple with the Steps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetStepsOk() (*[]SyntheticsTestResultStep, bool) {
	if o == nil || o.Steps == nil {
		return nil, false
	}
	return &o.Steps, true
}

// HasSteps returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasSteps() bool {
	return o != nil && o.Steps != nil
}

// SetSteps gets a reference to the given []SyntheticsTestResultStep and assigns it to the Steps field.
func (o *SyntheticsFastTestResultDetail) SetSteps(v []SyntheticsTestResultStep) {
	o.Steps = v
}

// GetTimings returns the Timings field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetTimings() map[string]interface{} {
	if o == nil || o.Timings == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Timings
}

// GetTimingsOk returns a tuple with the Timings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetTimingsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Timings == nil {
		return nil, false
	}
	return &o.Timings, true
}

// HasTimings returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasTimings() bool {
	return o != nil && o.Timings != nil
}

// SetTimings gets a reference to the given map[string]interface{} and assigns it to the Timings field.
func (o *SyntheticsFastTestResultDetail) SetTimings(v map[string]interface{}) {
	o.Timings = v
}

// GetTraceroute returns the Traceroute field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetTraceroute() []SyntheticsTestResultTracerouteHop {
	if o == nil || o.Traceroute == nil {
		var ret []SyntheticsTestResultTracerouteHop
		return ret
	}
	return o.Traceroute
}

// GetTracerouteOk returns a tuple with the Traceroute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetTracerouteOk() (*[]SyntheticsTestResultTracerouteHop, bool) {
	if o == nil || o.Traceroute == nil {
		return nil, false
	}
	return &o.Traceroute, true
}

// HasTraceroute returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasTraceroute() bool {
	return o != nil && o.Traceroute != nil
}

// SetTraceroute gets a reference to the given []SyntheticsTestResultTracerouteHop and assigns it to the Traceroute field.
func (o *SyntheticsFastTestResultDetail) SetTraceroute(v []SyntheticsTestResultTracerouteHop) {
	o.Traceroute = v
}

// GetTriggeredAt returns the TriggeredAt field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetTriggeredAt() int64 {
	if o == nil || o.TriggeredAt == nil {
		var ret int64
		return ret
	}
	return *o.TriggeredAt
}

// GetTriggeredAtOk returns a tuple with the TriggeredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetTriggeredAtOk() (*int64, bool) {
	if o == nil || o.TriggeredAt == nil {
		return nil, false
	}
	return o.TriggeredAt, true
}

// HasTriggeredAt returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasTriggeredAt() bool {
	return o != nil && o.TriggeredAt != nil
}

// SetTriggeredAt gets a reference to the given int64 and assigns it to the TriggeredAt field.
func (o *SyntheticsFastTestResultDetail) SetTriggeredAt(v int64) {
	o.TriggeredAt = &v
}

// GetTunnel returns the Tunnel field value if set, zero value otherwise.
func (o *SyntheticsFastTestResultDetail) GetTunnel() bool {
	if o == nil || o.Tunnel == nil {
		var ret bool
		return ret
	}
	return *o.Tunnel
}

// GetTunnelOk returns a tuple with the Tunnel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsFastTestResultDetail) GetTunnelOk() (*bool, bool) {
	if o == nil || o.Tunnel == nil {
		return nil, false
	}
	return o.Tunnel, true
}

// HasTunnel returns a boolean if a field has been set.
func (o *SyntheticsFastTestResultDetail) HasTunnel() bool {
	return o != nil && o.Tunnel != nil
}

// SetTunnel gets a reference to the given bool and assigns it to the Tunnel field.
func (o *SyntheticsFastTestResultDetail) SetTunnel(v bool) {
	o.Tunnel = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o SyntheticsFastTestResultDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Assertions != nil {
		toSerialize["assertions"] = o.Assertions
	}
	if o.CallType != nil {
		toSerialize["call_type"] = o.CallType
	}
	if o.Cert != nil {
		toSerialize["cert"] = o.Cert
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Failure != nil {
		toSerialize["failure"] = o.Failure
	}
	if o.FinishedAt != nil {
		toSerialize["finished_at"] = o.FinishedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsFastRetry != nil {
		toSerialize["is_fast_retry"] = o.IsFastRetry
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
	if o.StartedAt != nil {
		toSerialize["started_at"] = o.StartedAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Steps != nil {
		toSerialize["steps"] = o.Steps
	}
	if o.Timings != nil {
		toSerialize["timings"] = o.Timings
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SyntheticsFastTestResultDetail) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Assertions  []SyntheticsTestResultAssertionResult `json:"assertions,omitempty"`
		CallType    *string                               `json:"call_type,omitempty"`
		Cert        *SyntheticsTestResultCertificate      `json:"cert,omitempty"`
		Duration    *float64                              `json:"duration,omitempty"`
		Failure     *SyntheticsTestResultFailure          `json:"failure,omitempty"`
		FinishedAt  *int64                                `json:"finished_at,omitempty"`
		Id          *string                               `json:"id,omitempty"`
		IsFastRetry *bool                                 `json:"is_fast_retry,omitempty"`
		Request     *SyntheticsTestResultRequestInfo      `json:"request,omitempty"`
		ResolvedIp  *string                               `json:"resolved_ip,omitempty"`
		Response    *SyntheticsTestResultResponseInfo     `json:"response,omitempty"`
		RunType     *SyntheticsTestResultRunType          `json:"run_type,omitempty"`
		StartedAt   *int64                                `json:"started_at,omitempty"`
		Status      *string                               `json:"status,omitempty"`
		Steps       []SyntheticsTestResultStep            `json:"steps,omitempty"`
		Timings     map[string]interface{}                `json:"timings,omitempty"`
		Traceroute  []SyntheticsTestResultTracerouteHop   `json:"traceroute,omitempty"`
		TriggeredAt *int64                                `json:"triggered_at,omitempty"`
		Tunnel      *bool                                 `json:"tunnel,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"assertions", "call_type", "cert", "duration", "failure", "finished_at", "id", "is_fast_retry", "request", "resolved_ip", "response", "run_type", "started_at", "status", "steps", "timings", "traceroute", "triggered_at", "tunnel"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Assertions = all.Assertions
	o.CallType = all.CallType
	if all.Cert != nil && all.Cert.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Cert = all.Cert
	o.Duration = all.Duration
	if all.Failure != nil && all.Failure.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Failure = all.Failure
	o.FinishedAt = all.FinishedAt
	o.Id = all.Id
	o.IsFastRetry = all.IsFastRetry
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
	o.StartedAt = all.StartedAt
	o.Status = all.Status
	o.Steps = all.Steps
	o.Timings = all.Timings
	o.Traceroute = all.Traceroute
	o.TriggeredAt = all.TriggeredAt
	o.Tunnel = all.Tunnel

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
