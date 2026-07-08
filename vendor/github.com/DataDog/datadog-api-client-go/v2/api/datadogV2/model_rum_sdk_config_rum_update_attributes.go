// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// RumSdkConfigRumUpdateAttributes The RUM SDK settings to apply when updating a configuration.
type RumSdkConfigRumUpdateAttributes struct {
	// A list of URL configurations for distributed tracing.
	AllowedTracingUrls []RumSdkConfigTracingUrlConfig `json:"allowed_tracing_urls,omitempty"`
	// A list of origin patterns allowed for cross-origin session tracking.
	AllowedTrackingOrigins []RumSdkConfigMatchOption `json:"allowed_tracking_origins,omitempty"`
	// A list of dynamic option key-value pairs.
	Context []RumSdkConfigDynamicOptionPair `json:"context,omitempty"`
	// The default privacy masking level applied to all RUM data.
	DefaultPrivacyLevel string `json:"default_privacy_level"`
	// Whether to mask user-interaction action names for privacy.
	EnablePrivacyForActionName bool `json:"enable_privacy_for_action_name"`
	// The environment tag for the RUM application.
	Env *string `json:"env,omitempty"`
	// The service name tag for the RUM application.
	Service *string `json:"service,omitempty"`
	// The percentage of collected sessions for which a replay is captured (0–100).
	SessionReplaySampleRate int64 `json:"session_replay_sample_rate"`
	// The percentage of user sessions to collect (0–100).
	SessionSampleRate int64 `json:"session_sample_rate"`
	// The percentage of requests to forward as APM traces (0–100).
	TraceSampleRate *int64 `json:"trace_sample_rate,omitempty"`
	// Whether to share a session across subdomains of the same site.
	TrackSessionAcrossSubdomains *bool `json:"track_session_across_subdomains,omitempty"`
	// A list of dynamic option key-value pairs.
	User []RumSdkConfigDynamicOptionPair `json:"user,omitempty"`
	// A dynamic configuration option that extracts a value at runtime using a specified strategy.
	Version *RumSdkConfigDynamicOption `json:"version,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRumSdkConfigRumUpdateAttributes instantiates a new RumSdkConfigRumUpdateAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRumSdkConfigRumUpdateAttributes(defaultPrivacyLevel string, enablePrivacyForActionName bool, sessionReplaySampleRate int64, sessionSampleRate int64) *RumSdkConfigRumUpdateAttributes {
	this := RumSdkConfigRumUpdateAttributes{}
	this.DefaultPrivacyLevel = defaultPrivacyLevel
	this.EnablePrivacyForActionName = enablePrivacyForActionName
	this.SessionReplaySampleRate = sessionReplaySampleRate
	this.SessionSampleRate = sessionSampleRate
	return &this
}

// NewRumSdkConfigRumUpdateAttributesWithDefaults instantiates a new RumSdkConfigRumUpdateAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRumSdkConfigRumUpdateAttributesWithDefaults() *RumSdkConfigRumUpdateAttributes {
	this := RumSdkConfigRumUpdateAttributes{}
	return &this
}

// GetAllowedTracingUrls returns the AllowedTracingUrls field value if set, zero value otherwise.
func (o *RumSdkConfigRumUpdateAttributes) GetAllowedTracingUrls() []RumSdkConfigTracingUrlConfig {
	if o == nil || o.AllowedTracingUrls == nil {
		var ret []RumSdkConfigTracingUrlConfig
		return ret
	}
	return o.AllowedTracingUrls
}

// GetAllowedTracingUrlsOk returns a tuple with the AllowedTracingUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSdkConfigRumUpdateAttributes) GetAllowedTracingUrlsOk() (*[]RumSdkConfigTracingUrlConfig, bool) {
	if o == nil || o.AllowedTracingUrls == nil {
		return nil, false
	}
	return &o.AllowedTracingUrls, true
}

// HasAllowedTracingUrls returns a boolean if a field has been set.
func (o *RumSdkConfigRumUpdateAttributes) HasAllowedTracingUrls() bool {
	return o != nil && o.AllowedTracingUrls != nil
}

// SetAllowedTracingUrls gets a reference to the given []RumSdkConfigTracingUrlConfig and assigns it to the AllowedTracingUrls field.
func (o *RumSdkConfigRumUpdateAttributes) SetAllowedTracingUrls(v []RumSdkConfigTracingUrlConfig) {
	o.AllowedTracingUrls = v
}

// GetAllowedTrackingOrigins returns the AllowedTrackingOrigins field value if set, zero value otherwise.
func (o *RumSdkConfigRumUpdateAttributes) GetAllowedTrackingOrigins() []RumSdkConfigMatchOption {
	if o == nil || o.AllowedTrackingOrigins == nil {
		var ret []RumSdkConfigMatchOption
		return ret
	}
	return o.AllowedTrackingOrigins
}

// GetAllowedTrackingOriginsOk returns a tuple with the AllowedTrackingOrigins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSdkConfigRumUpdateAttributes) GetAllowedTrackingOriginsOk() (*[]RumSdkConfigMatchOption, bool) {
	if o == nil || o.AllowedTrackingOrigins == nil {
		return nil, false
	}
	return &o.AllowedTrackingOrigins, true
}

// HasAllowedTrackingOrigins returns a boolean if a field has been set.
func (o *RumSdkConfigRumUpdateAttributes) HasAllowedTrackingOrigins() bool {
	return o != nil && o.AllowedTrackingOrigins != nil
}

// SetAllowedTrackingOrigins gets a reference to the given []RumSdkConfigMatchOption and assigns it to the AllowedTrackingOrigins field.
func (o *RumSdkConfigRumUpdateAttributes) SetAllowedTrackingOrigins(v []RumSdkConfigMatchOption) {
	o.AllowedTrackingOrigins = v
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *RumSdkConfigRumUpdateAttributes) GetContext() []RumSdkConfigDynamicOptionPair {
	if o == nil || o.Context == nil {
		var ret []RumSdkConfigDynamicOptionPair
		return ret
	}
	return o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSdkConfigRumUpdateAttributes) GetContextOk() (*[]RumSdkConfigDynamicOptionPair, bool) {
	if o == nil || o.Context == nil {
		return nil, false
	}
	return &o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *RumSdkConfigRumUpdateAttributes) HasContext() bool {
	return o != nil && o.Context != nil
}

// SetContext gets a reference to the given []RumSdkConfigDynamicOptionPair and assigns it to the Context field.
func (o *RumSdkConfigRumUpdateAttributes) SetContext(v []RumSdkConfigDynamicOptionPair) {
	o.Context = v
}

// GetDefaultPrivacyLevel returns the DefaultPrivacyLevel field value.
func (o *RumSdkConfigRumUpdateAttributes) GetDefaultPrivacyLevel() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DefaultPrivacyLevel
}

// GetDefaultPrivacyLevelOk returns a tuple with the DefaultPrivacyLevel field value
// and a boolean to check if the value has been set.
func (o *RumSdkConfigRumUpdateAttributes) GetDefaultPrivacyLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultPrivacyLevel, true
}

// SetDefaultPrivacyLevel sets field value.
func (o *RumSdkConfigRumUpdateAttributes) SetDefaultPrivacyLevel(v string) {
	o.DefaultPrivacyLevel = v
}

// GetEnablePrivacyForActionName returns the EnablePrivacyForActionName field value.
func (o *RumSdkConfigRumUpdateAttributes) GetEnablePrivacyForActionName() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.EnablePrivacyForActionName
}

// GetEnablePrivacyForActionNameOk returns a tuple with the EnablePrivacyForActionName field value
// and a boolean to check if the value has been set.
func (o *RumSdkConfigRumUpdateAttributes) GetEnablePrivacyForActionNameOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnablePrivacyForActionName, true
}

// SetEnablePrivacyForActionName sets field value.
func (o *RumSdkConfigRumUpdateAttributes) SetEnablePrivacyForActionName(v bool) {
	o.EnablePrivacyForActionName = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *RumSdkConfigRumUpdateAttributes) GetEnv() string {
	if o == nil || o.Env == nil {
		var ret string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSdkConfigRumUpdateAttributes) GetEnvOk() (*string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *RumSdkConfigRumUpdateAttributes) HasEnv() bool {
	return o != nil && o.Env != nil
}

// SetEnv gets a reference to the given string and assigns it to the Env field.
func (o *RumSdkConfigRumUpdateAttributes) SetEnv(v string) {
	o.Env = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *RumSdkConfigRumUpdateAttributes) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSdkConfigRumUpdateAttributes) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *RumSdkConfigRumUpdateAttributes) HasService() bool {
	return o != nil && o.Service != nil
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *RumSdkConfigRumUpdateAttributes) SetService(v string) {
	o.Service = &v
}

// GetSessionReplaySampleRate returns the SessionReplaySampleRate field value.
func (o *RumSdkConfigRumUpdateAttributes) GetSessionReplaySampleRate() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionReplaySampleRate
}

// GetSessionReplaySampleRateOk returns a tuple with the SessionReplaySampleRate field value
// and a boolean to check if the value has been set.
func (o *RumSdkConfigRumUpdateAttributes) GetSessionReplaySampleRateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionReplaySampleRate, true
}

// SetSessionReplaySampleRate sets field value.
func (o *RumSdkConfigRumUpdateAttributes) SetSessionReplaySampleRate(v int64) {
	o.SessionReplaySampleRate = v
}

// GetSessionSampleRate returns the SessionSampleRate field value.
func (o *RumSdkConfigRumUpdateAttributes) GetSessionSampleRate() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SessionSampleRate
}

// GetSessionSampleRateOk returns a tuple with the SessionSampleRate field value
// and a boolean to check if the value has been set.
func (o *RumSdkConfigRumUpdateAttributes) GetSessionSampleRateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionSampleRate, true
}

// SetSessionSampleRate sets field value.
func (o *RumSdkConfigRumUpdateAttributes) SetSessionSampleRate(v int64) {
	o.SessionSampleRate = v
}

// GetTraceSampleRate returns the TraceSampleRate field value if set, zero value otherwise.
func (o *RumSdkConfigRumUpdateAttributes) GetTraceSampleRate() int64 {
	if o == nil || o.TraceSampleRate == nil {
		var ret int64
		return ret
	}
	return *o.TraceSampleRate
}

// GetTraceSampleRateOk returns a tuple with the TraceSampleRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSdkConfigRumUpdateAttributes) GetTraceSampleRateOk() (*int64, bool) {
	if o == nil || o.TraceSampleRate == nil {
		return nil, false
	}
	return o.TraceSampleRate, true
}

// HasTraceSampleRate returns a boolean if a field has been set.
func (o *RumSdkConfigRumUpdateAttributes) HasTraceSampleRate() bool {
	return o != nil && o.TraceSampleRate != nil
}

// SetTraceSampleRate gets a reference to the given int64 and assigns it to the TraceSampleRate field.
func (o *RumSdkConfigRumUpdateAttributes) SetTraceSampleRate(v int64) {
	o.TraceSampleRate = &v
}

// GetTrackSessionAcrossSubdomains returns the TrackSessionAcrossSubdomains field value if set, zero value otherwise.
func (o *RumSdkConfigRumUpdateAttributes) GetTrackSessionAcrossSubdomains() bool {
	if o == nil || o.TrackSessionAcrossSubdomains == nil {
		var ret bool
		return ret
	}
	return *o.TrackSessionAcrossSubdomains
}

// GetTrackSessionAcrossSubdomainsOk returns a tuple with the TrackSessionAcrossSubdomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSdkConfigRumUpdateAttributes) GetTrackSessionAcrossSubdomainsOk() (*bool, bool) {
	if o == nil || o.TrackSessionAcrossSubdomains == nil {
		return nil, false
	}
	return o.TrackSessionAcrossSubdomains, true
}

// HasTrackSessionAcrossSubdomains returns a boolean if a field has been set.
func (o *RumSdkConfigRumUpdateAttributes) HasTrackSessionAcrossSubdomains() bool {
	return o != nil && o.TrackSessionAcrossSubdomains != nil
}

// SetTrackSessionAcrossSubdomains gets a reference to the given bool and assigns it to the TrackSessionAcrossSubdomains field.
func (o *RumSdkConfigRumUpdateAttributes) SetTrackSessionAcrossSubdomains(v bool) {
	o.TrackSessionAcrossSubdomains = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RumSdkConfigRumUpdateAttributes) GetUser() []RumSdkConfigDynamicOptionPair {
	if o == nil || o.User == nil {
		var ret []RumSdkConfigDynamicOptionPair
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSdkConfigRumUpdateAttributes) GetUserOk() (*[]RumSdkConfigDynamicOptionPair, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return &o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RumSdkConfigRumUpdateAttributes) HasUser() bool {
	return o != nil && o.User != nil
}

// SetUser gets a reference to the given []RumSdkConfigDynamicOptionPair and assigns it to the User field.
func (o *RumSdkConfigRumUpdateAttributes) SetUser(v []RumSdkConfigDynamicOptionPair) {
	o.User = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RumSdkConfigRumUpdateAttributes) GetVersion() RumSdkConfigDynamicOption {
	if o == nil || o.Version == nil {
		var ret RumSdkConfigDynamicOption
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RumSdkConfigRumUpdateAttributes) GetVersionOk() (*RumSdkConfigDynamicOption, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RumSdkConfigRumUpdateAttributes) HasVersion() bool {
	return o != nil && o.Version != nil
}

// SetVersion gets a reference to the given RumSdkConfigDynamicOption and assigns it to the Version field.
func (o *RumSdkConfigRumUpdateAttributes) SetVersion(v RumSdkConfigDynamicOption) {
	o.Version = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RumSdkConfigRumUpdateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AllowedTracingUrls != nil {
		toSerialize["allowed_tracing_urls"] = o.AllowedTracingUrls
	}
	if o.AllowedTrackingOrigins != nil {
		toSerialize["allowed_tracking_origins"] = o.AllowedTrackingOrigins
	}
	if o.Context != nil {
		toSerialize["context"] = o.Context
	}
	toSerialize["default_privacy_level"] = o.DefaultPrivacyLevel
	toSerialize["enable_privacy_for_action_name"] = o.EnablePrivacyForActionName
	if o.Env != nil {
		toSerialize["env"] = o.Env
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	toSerialize["session_replay_sample_rate"] = o.SessionReplaySampleRate
	toSerialize["session_sample_rate"] = o.SessionSampleRate
	if o.TraceSampleRate != nil {
		toSerialize["trace_sample_rate"] = o.TraceSampleRate
	}
	if o.TrackSessionAcrossSubdomains != nil {
		toSerialize["track_session_across_subdomains"] = o.TrackSessionAcrossSubdomains
	}
	if o.User != nil {
		toSerialize["user"] = o.User
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
func (o *RumSdkConfigRumUpdateAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AllowedTracingUrls           []RumSdkConfigTracingUrlConfig  `json:"allowed_tracing_urls,omitempty"`
		AllowedTrackingOrigins       []RumSdkConfigMatchOption       `json:"allowed_tracking_origins,omitempty"`
		Context                      []RumSdkConfigDynamicOptionPair `json:"context,omitempty"`
		DefaultPrivacyLevel          *string                         `json:"default_privacy_level"`
		EnablePrivacyForActionName   *bool                           `json:"enable_privacy_for_action_name"`
		Env                          *string                         `json:"env,omitempty"`
		Service                      *string                         `json:"service,omitempty"`
		SessionReplaySampleRate      *int64                          `json:"session_replay_sample_rate"`
		SessionSampleRate            *int64                          `json:"session_sample_rate"`
		TraceSampleRate              *int64                          `json:"trace_sample_rate,omitempty"`
		TrackSessionAcrossSubdomains *bool                           `json:"track_session_across_subdomains,omitempty"`
		User                         []RumSdkConfigDynamicOptionPair `json:"user,omitempty"`
		Version                      *RumSdkConfigDynamicOption      `json:"version,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DefaultPrivacyLevel == nil {
		return fmt.Errorf("required field default_privacy_level missing")
	}
	if all.EnablePrivacyForActionName == nil {
		return fmt.Errorf("required field enable_privacy_for_action_name missing")
	}
	if all.SessionReplaySampleRate == nil {
		return fmt.Errorf("required field session_replay_sample_rate missing")
	}
	if all.SessionSampleRate == nil {
		return fmt.Errorf("required field session_sample_rate missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"allowed_tracing_urls", "allowed_tracking_origins", "context", "default_privacy_level", "enable_privacy_for_action_name", "env", "service", "session_replay_sample_rate", "session_sample_rate", "trace_sample_rate", "track_session_across_subdomains", "user", "version"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AllowedTracingUrls = all.AllowedTracingUrls
	o.AllowedTrackingOrigins = all.AllowedTrackingOrigins
	o.Context = all.Context
	o.DefaultPrivacyLevel = *all.DefaultPrivacyLevel
	o.EnablePrivacyForActionName = *all.EnablePrivacyForActionName
	o.Env = all.Env
	o.Service = all.Service
	o.SessionReplaySampleRate = *all.SessionReplaySampleRate
	o.SessionSampleRate = *all.SessionSampleRate
	o.TraceSampleRate = all.TraceSampleRate
	o.TrackSessionAcrossSubdomains = all.TrackSessionAcrossSubdomains
	o.User = all.User
	if all.Version != nil && all.Version.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Version = all.Version

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
