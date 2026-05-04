// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IoCIndicatorDetailed An indicator of compromise with extended context from your environment.
type IoCIndicatorDetailed struct {
	// Additional domain-specific context from threat intelligence sources.
	AdditionalData map[string]interface{} `json:"additional_data,omitempty"`
	// Autonomous system CIDR block.
	AsCidrBlock *string `json:"as_cidr_block,omitempty"`
	// Geographic location information for an IP indicator.
	AsGeo *IoCGeoLocation `json:"as_geo,omitempty"`
	// Autonomous system number.
	AsNumber *string `json:"as_number,omitempty"`
	// Autonomous system organization name.
	AsOrganization *string `json:"as_organization,omitempty"`
	// Autonomous system type.
	AsType *string `json:"as_type,omitempty"`
	// Threat intelligence sources that flagged this indicator as benign.
	BenignSources []IoCSource `json:"benign_sources,omitempty"`
	// Threat categories associated with the indicator.
	Categories []string `json:"categories,omitempty"`
	// Critical assets associated with this indicator.
	CriticalAssets []string `json:"critical_assets,omitempty"`
	// Timestamp when the indicator was first seen.
	FirstSeen *time.Time `json:"first_seen,omitempty"`
	// Hosts associated with this indicator.
	Hosts []string `json:"hosts,omitempty"`
	// Unique identifier for the indicator.
	Id *string `json:"id,omitempty"`
	// The indicator value (for example, an IP address or domain).
	Indicator *string `json:"indicator,omitempty"`
	// Type of indicator (for example, IP address or domain).
	IndicatorType *string `json:"indicator_type,omitempty"`
	// Timestamp when the indicator was last seen.
	LastSeen *time.Time `json:"last_seen,omitempty"`
	// Number of logs that matched this indicator.
	LogMatches *int64 `json:"log_matches,omitempty"`
	// Log sources where this indicator was observed.
	LogSources []string `json:"log_sources,omitempty"`
	// Effect of a scoring factor on the indicator's threat score.
	MAsType *IoCScoreEffect `json:"m_as_type,omitempty"`
	// Effect of a scoring factor on the indicator's threat score.
	MPersistence *IoCScoreEffect `json:"m_persistence,omitempty"`
	// Effect of a scoring factor on the indicator's threat score.
	MSignal *IoCScoreEffect `json:"m_signal,omitempty"`
	// Effect of a scoring factor on the indicator's threat score.
	MSources *IoCScoreEffect `json:"m_sources,omitempty"`
	// Threat intelligence sources that flagged this indicator as malicious.
	MaliciousSources []IoCSource `json:"malicious_sources,omitempty"`
	// Effect of a scoring factor on the indicator's threat score.
	MaxTrustScore *IoCScoreEffect `json:"max_trust_score,omitempty"`
	// Threat score for the indicator (0-100).
	Score *float64 `json:"score,omitempty"`
	// Services where this indicator was observed.
	Services []string `json:"services,omitempty"`
	// Number of security signals that matched this indicator.
	SignalMatches *int64 `json:"signal_matches,omitempty"`
	// Breakdown of security signals by severity.
	SignalSeverity []IoCSignalSeverityCount `json:"signal_severity,omitempty"`
	// Signal tier level.
	SignalTier *int64 `json:"signal_tier,omitempty"`
	// Threat intelligence sources that flagged this indicator as suspicious.
	SuspiciousSources []IoCSource `json:"suspicious_sources,omitempty"`
	// Tags associated with the indicator.
	Tags []string `json:"tags,omitempty"`
	// Users associated with this indicator, grouped by category.
	Users map[string][]string `json:"users,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIoCIndicatorDetailed instantiates a new IoCIndicatorDetailed object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIoCIndicatorDetailed() *IoCIndicatorDetailed {
	this := IoCIndicatorDetailed{}
	return &this
}

// NewIoCIndicatorDetailedWithDefaults instantiates a new IoCIndicatorDetailed object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIoCIndicatorDetailedWithDefaults() *IoCIndicatorDetailed {
	this := IoCIndicatorDetailed{}
	return &this
}

// GetAdditionalData returns the AdditionalData field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetAdditionalData() map[string]interface{} {
	if o == nil || o.AdditionalData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalData
}

// GetAdditionalDataOk returns a tuple with the AdditionalData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetAdditionalDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.AdditionalData == nil {
		return nil, false
	}
	return &o.AdditionalData, true
}

// HasAdditionalData returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasAdditionalData() bool {
	return o != nil && o.AdditionalData != nil
}

// SetAdditionalData gets a reference to the given map[string]interface{} and assigns it to the AdditionalData field.
func (o *IoCIndicatorDetailed) SetAdditionalData(v map[string]interface{}) {
	o.AdditionalData = v
}

// GetAsCidrBlock returns the AsCidrBlock field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetAsCidrBlock() string {
	if o == nil || o.AsCidrBlock == nil {
		var ret string
		return ret
	}
	return *o.AsCidrBlock
}

// GetAsCidrBlockOk returns a tuple with the AsCidrBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetAsCidrBlockOk() (*string, bool) {
	if o == nil || o.AsCidrBlock == nil {
		return nil, false
	}
	return o.AsCidrBlock, true
}

// HasAsCidrBlock returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasAsCidrBlock() bool {
	return o != nil && o.AsCidrBlock != nil
}

// SetAsCidrBlock gets a reference to the given string and assigns it to the AsCidrBlock field.
func (o *IoCIndicatorDetailed) SetAsCidrBlock(v string) {
	o.AsCidrBlock = &v
}

// GetAsGeo returns the AsGeo field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetAsGeo() IoCGeoLocation {
	if o == nil || o.AsGeo == nil {
		var ret IoCGeoLocation
		return ret
	}
	return *o.AsGeo
}

// GetAsGeoOk returns a tuple with the AsGeo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetAsGeoOk() (*IoCGeoLocation, bool) {
	if o == nil || o.AsGeo == nil {
		return nil, false
	}
	return o.AsGeo, true
}

// HasAsGeo returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasAsGeo() bool {
	return o != nil && o.AsGeo != nil
}

// SetAsGeo gets a reference to the given IoCGeoLocation and assigns it to the AsGeo field.
func (o *IoCIndicatorDetailed) SetAsGeo(v IoCGeoLocation) {
	o.AsGeo = &v
}

// GetAsNumber returns the AsNumber field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetAsNumber() string {
	if o == nil || o.AsNumber == nil {
		var ret string
		return ret
	}
	return *o.AsNumber
}

// GetAsNumberOk returns a tuple with the AsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetAsNumberOk() (*string, bool) {
	if o == nil || o.AsNumber == nil {
		return nil, false
	}
	return o.AsNumber, true
}

// HasAsNumber returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasAsNumber() bool {
	return o != nil && o.AsNumber != nil
}

// SetAsNumber gets a reference to the given string and assigns it to the AsNumber field.
func (o *IoCIndicatorDetailed) SetAsNumber(v string) {
	o.AsNumber = &v
}

// GetAsOrganization returns the AsOrganization field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetAsOrganization() string {
	if o == nil || o.AsOrganization == nil {
		var ret string
		return ret
	}
	return *o.AsOrganization
}

// GetAsOrganizationOk returns a tuple with the AsOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetAsOrganizationOk() (*string, bool) {
	if o == nil || o.AsOrganization == nil {
		return nil, false
	}
	return o.AsOrganization, true
}

// HasAsOrganization returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasAsOrganization() bool {
	return o != nil && o.AsOrganization != nil
}

// SetAsOrganization gets a reference to the given string and assigns it to the AsOrganization field.
func (o *IoCIndicatorDetailed) SetAsOrganization(v string) {
	o.AsOrganization = &v
}

// GetAsType returns the AsType field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetAsType() string {
	if o == nil || o.AsType == nil {
		var ret string
		return ret
	}
	return *o.AsType
}

// GetAsTypeOk returns a tuple with the AsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetAsTypeOk() (*string, bool) {
	if o == nil || o.AsType == nil {
		return nil, false
	}
	return o.AsType, true
}

// HasAsType returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasAsType() bool {
	return o != nil && o.AsType != nil
}

// SetAsType gets a reference to the given string and assigns it to the AsType field.
func (o *IoCIndicatorDetailed) SetAsType(v string) {
	o.AsType = &v
}

// GetBenignSources returns the BenignSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IoCIndicatorDetailed) GetBenignSources() []IoCSource {
	if o == nil {
		var ret []IoCSource
		return ret
	}
	return o.BenignSources
}

// GetBenignSourcesOk returns a tuple with the BenignSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IoCIndicatorDetailed) GetBenignSourcesOk() (*[]IoCSource, bool) {
	if o == nil || o.BenignSources == nil {
		return nil, false
	}
	return &o.BenignSources, true
}

// HasBenignSources returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasBenignSources() bool {
	return o != nil && o.BenignSources != nil
}

// SetBenignSources gets a reference to the given []IoCSource and assigns it to the BenignSources field.
func (o *IoCIndicatorDetailed) SetBenignSources(v []IoCSource) {
	o.BenignSources = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetCategories() []string {
	if o == nil || o.Categories == nil {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetCategoriesOk() (*[]string, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return &o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasCategories() bool {
	return o != nil && o.Categories != nil
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *IoCIndicatorDetailed) SetCategories(v []string) {
	o.Categories = v
}

// GetCriticalAssets returns the CriticalAssets field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetCriticalAssets() []string {
	if o == nil || o.CriticalAssets == nil {
		var ret []string
		return ret
	}
	return o.CriticalAssets
}

// GetCriticalAssetsOk returns a tuple with the CriticalAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetCriticalAssetsOk() (*[]string, bool) {
	if o == nil || o.CriticalAssets == nil {
		return nil, false
	}
	return &o.CriticalAssets, true
}

// HasCriticalAssets returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasCriticalAssets() bool {
	return o != nil && o.CriticalAssets != nil
}

// SetCriticalAssets gets a reference to the given []string and assigns it to the CriticalAssets field.
func (o *IoCIndicatorDetailed) SetCriticalAssets(v []string) {
	o.CriticalAssets = v
}

// GetFirstSeen returns the FirstSeen field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetFirstSeen() time.Time {
	if o == nil || o.FirstSeen == nil {
		var ret time.Time
		return ret
	}
	return *o.FirstSeen
}

// GetFirstSeenOk returns a tuple with the FirstSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetFirstSeenOk() (*time.Time, bool) {
	if o == nil || o.FirstSeen == nil {
		return nil, false
	}
	return o.FirstSeen, true
}

// HasFirstSeen returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasFirstSeen() bool {
	return o != nil && o.FirstSeen != nil
}

// SetFirstSeen gets a reference to the given time.Time and assigns it to the FirstSeen field.
func (o *IoCIndicatorDetailed) SetFirstSeen(v time.Time) {
	o.FirstSeen = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetHosts() []string {
	if o == nil || o.Hosts == nil {
		var ret []string
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetHostsOk() (*[]string, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return &o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasHosts() bool {
	return o != nil && o.Hosts != nil
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *IoCIndicatorDetailed) SetHosts(v []string) {
	o.Hosts = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IoCIndicatorDetailed) SetId(v string) {
	o.Id = &v
}

// GetIndicator returns the Indicator field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetIndicator() string {
	if o == nil || o.Indicator == nil {
		var ret string
		return ret
	}
	return *o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetIndicatorOk() (*string, bool) {
	if o == nil || o.Indicator == nil {
		return nil, false
	}
	return o.Indicator, true
}

// HasIndicator returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasIndicator() bool {
	return o != nil && o.Indicator != nil
}

// SetIndicator gets a reference to the given string and assigns it to the Indicator field.
func (o *IoCIndicatorDetailed) SetIndicator(v string) {
	o.Indicator = &v
}

// GetIndicatorType returns the IndicatorType field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetIndicatorType() string {
	if o == nil || o.IndicatorType == nil {
		var ret string
		return ret
	}
	return *o.IndicatorType
}

// GetIndicatorTypeOk returns a tuple with the IndicatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetIndicatorTypeOk() (*string, bool) {
	if o == nil || o.IndicatorType == nil {
		return nil, false
	}
	return o.IndicatorType, true
}

// HasIndicatorType returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasIndicatorType() bool {
	return o != nil && o.IndicatorType != nil
}

// SetIndicatorType gets a reference to the given string and assigns it to the IndicatorType field.
func (o *IoCIndicatorDetailed) SetIndicatorType(v string) {
	o.IndicatorType = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetLastSeen() time.Time {
	if o == nil || o.LastSeen == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetLastSeenOk() (*time.Time, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasLastSeen() bool {
	return o != nil && o.LastSeen != nil
}

// SetLastSeen gets a reference to the given time.Time and assigns it to the LastSeen field.
func (o *IoCIndicatorDetailed) SetLastSeen(v time.Time) {
	o.LastSeen = &v
}

// GetLogMatches returns the LogMatches field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetLogMatches() int64 {
	if o == nil || o.LogMatches == nil {
		var ret int64
		return ret
	}
	return *o.LogMatches
}

// GetLogMatchesOk returns a tuple with the LogMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetLogMatchesOk() (*int64, bool) {
	if o == nil || o.LogMatches == nil {
		return nil, false
	}
	return o.LogMatches, true
}

// HasLogMatches returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasLogMatches() bool {
	return o != nil && o.LogMatches != nil
}

// SetLogMatches gets a reference to the given int64 and assigns it to the LogMatches field.
func (o *IoCIndicatorDetailed) SetLogMatches(v int64) {
	o.LogMatches = &v
}

// GetLogSources returns the LogSources field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetLogSources() []string {
	if o == nil || o.LogSources == nil {
		var ret []string
		return ret
	}
	return o.LogSources
}

// GetLogSourcesOk returns a tuple with the LogSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetLogSourcesOk() (*[]string, bool) {
	if o == nil || o.LogSources == nil {
		return nil, false
	}
	return &o.LogSources, true
}

// HasLogSources returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasLogSources() bool {
	return o != nil && o.LogSources != nil
}

// SetLogSources gets a reference to the given []string and assigns it to the LogSources field.
func (o *IoCIndicatorDetailed) SetLogSources(v []string) {
	o.LogSources = v
}

// GetMAsType returns the MAsType field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetMAsType() IoCScoreEffect {
	if o == nil || o.MAsType == nil {
		var ret IoCScoreEffect
		return ret
	}
	return *o.MAsType
}

// GetMAsTypeOk returns a tuple with the MAsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetMAsTypeOk() (*IoCScoreEffect, bool) {
	if o == nil || o.MAsType == nil {
		return nil, false
	}
	return o.MAsType, true
}

// HasMAsType returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasMAsType() bool {
	return o != nil && o.MAsType != nil
}

// SetMAsType gets a reference to the given IoCScoreEffect and assigns it to the MAsType field.
func (o *IoCIndicatorDetailed) SetMAsType(v IoCScoreEffect) {
	o.MAsType = &v
}

// GetMPersistence returns the MPersistence field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetMPersistence() IoCScoreEffect {
	if o == nil || o.MPersistence == nil {
		var ret IoCScoreEffect
		return ret
	}
	return *o.MPersistence
}

// GetMPersistenceOk returns a tuple with the MPersistence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetMPersistenceOk() (*IoCScoreEffect, bool) {
	if o == nil || o.MPersistence == nil {
		return nil, false
	}
	return o.MPersistence, true
}

// HasMPersistence returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasMPersistence() bool {
	return o != nil && o.MPersistence != nil
}

// SetMPersistence gets a reference to the given IoCScoreEffect and assigns it to the MPersistence field.
func (o *IoCIndicatorDetailed) SetMPersistence(v IoCScoreEffect) {
	o.MPersistence = &v
}

// GetMSignal returns the MSignal field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetMSignal() IoCScoreEffect {
	if o == nil || o.MSignal == nil {
		var ret IoCScoreEffect
		return ret
	}
	return *o.MSignal
}

// GetMSignalOk returns a tuple with the MSignal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetMSignalOk() (*IoCScoreEffect, bool) {
	if o == nil || o.MSignal == nil {
		return nil, false
	}
	return o.MSignal, true
}

// HasMSignal returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasMSignal() bool {
	return o != nil && o.MSignal != nil
}

// SetMSignal gets a reference to the given IoCScoreEffect and assigns it to the MSignal field.
func (o *IoCIndicatorDetailed) SetMSignal(v IoCScoreEffect) {
	o.MSignal = &v
}

// GetMSources returns the MSources field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetMSources() IoCScoreEffect {
	if o == nil || o.MSources == nil {
		var ret IoCScoreEffect
		return ret
	}
	return *o.MSources
}

// GetMSourcesOk returns a tuple with the MSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetMSourcesOk() (*IoCScoreEffect, bool) {
	if o == nil || o.MSources == nil {
		return nil, false
	}
	return o.MSources, true
}

// HasMSources returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasMSources() bool {
	return o != nil && o.MSources != nil
}

// SetMSources gets a reference to the given IoCScoreEffect and assigns it to the MSources field.
func (o *IoCIndicatorDetailed) SetMSources(v IoCScoreEffect) {
	o.MSources = &v
}

// GetMaliciousSources returns the MaliciousSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IoCIndicatorDetailed) GetMaliciousSources() []IoCSource {
	if o == nil {
		var ret []IoCSource
		return ret
	}
	return o.MaliciousSources
}

// GetMaliciousSourcesOk returns a tuple with the MaliciousSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IoCIndicatorDetailed) GetMaliciousSourcesOk() (*[]IoCSource, bool) {
	if o == nil || o.MaliciousSources == nil {
		return nil, false
	}
	return &o.MaliciousSources, true
}

// HasMaliciousSources returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasMaliciousSources() bool {
	return o != nil && o.MaliciousSources != nil
}

// SetMaliciousSources gets a reference to the given []IoCSource and assigns it to the MaliciousSources field.
func (o *IoCIndicatorDetailed) SetMaliciousSources(v []IoCSource) {
	o.MaliciousSources = v
}

// GetMaxTrustScore returns the MaxTrustScore field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetMaxTrustScore() IoCScoreEffect {
	if o == nil || o.MaxTrustScore == nil {
		var ret IoCScoreEffect
		return ret
	}
	return *o.MaxTrustScore
}

// GetMaxTrustScoreOk returns a tuple with the MaxTrustScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetMaxTrustScoreOk() (*IoCScoreEffect, bool) {
	if o == nil || o.MaxTrustScore == nil {
		return nil, false
	}
	return o.MaxTrustScore, true
}

// HasMaxTrustScore returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasMaxTrustScore() bool {
	return o != nil && o.MaxTrustScore != nil
}

// SetMaxTrustScore gets a reference to the given IoCScoreEffect and assigns it to the MaxTrustScore field.
func (o *IoCIndicatorDetailed) SetMaxTrustScore(v IoCScoreEffect) {
	o.MaxTrustScore = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetScore() float64 {
	if o == nil || o.Score == nil {
		var ret float64
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetScoreOk() (*float64, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasScore() bool {
	return o != nil && o.Score != nil
}

// SetScore gets a reference to the given float64 and assigns it to the Score field.
func (o *IoCIndicatorDetailed) SetScore(v float64) {
	o.Score = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *IoCIndicatorDetailed) SetServices(v []string) {
	o.Services = v
}

// GetSignalMatches returns the SignalMatches field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetSignalMatches() int64 {
	if o == nil || o.SignalMatches == nil {
		var ret int64
		return ret
	}
	return *o.SignalMatches
}

// GetSignalMatchesOk returns a tuple with the SignalMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetSignalMatchesOk() (*int64, bool) {
	if o == nil || o.SignalMatches == nil {
		return nil, false
	}
	return o.SignalMatches, true
}

// HasSignalMatches returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasSignalMatches() bool {
	return o != nil && o.SignalMatches != nil
}

// SetSignalMatches gets a reference to the given int64 and assigns it to the SignalMatches field.
func (o *IoCIndicatorDetailed) SetSignalMatches(v int64) {
	o.SignalMatches = &v
}

// GetSignalSeverity returns the SignalSeverity field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetSignalSeverity() []IoCSignalSeverityCount {
	if o == nil || o.SignalSeverity == nil {
		var ret []IoCSignalSeverityCount
		return ret
	}
	return o.SignalSeverity
}

// GetSignalSeverityOk returns a tuple with the SignalSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetSignalSeverityOk() (*[]IoCSignalSeverityCount, bool) {
	if o == nil || o.SignalSeverity == nil {
		return nil, false
	}
	return &o.SignalSeverity, true
}

// HasSignalSeverity returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasSignalSeverity() bool {
	return o != nil && o.SignalSeverity != nil
}

// SetSignalSeverity gets a reference to the given []IoCSignalSeverityCount and assigns it to the SignalSeverity field.
func (o *IoCIndicatorDetailed) SetSignalSeverity(v []IoCSignalSeverityCount) {
	o.SignalSeverity = v
}

// GetSignalTier returns the SignalTier field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetSignalTier() int64 {
	if o == nil || o.SignalTier == nil {
		var ret int64
		return ret
	}
	return *o.SignalTier
}

// GetSignalTierOk returns a tuple with the SignalTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetSignalTierOk() (*int64, bool) {
	if o == nil || o.SignalTier == nil {
		return nil, false
	}
	return o.SignalTier, true
}

// HasSignalTier returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasSignalTier() bool {
	return o != nil && o.SignalTier != nil
}

// SetSignalTier gets a reference to the given int64 and assigns it to the SignalTier field.
func (o *IoCIndicatorDetailed) SetSignalTier(v int64) {
	o.SignalTier = &v
}

// GetSuspiciousSources returns the SuspiciousSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IoCIndicatorDetailed) GetSuspiciousSources() []IoCSource {
	if o == nil {
		var ret []IoCSource
		return ret
	}
	return o.SuspiciousSources
}

// GetSuspiciousSourcesOk returns a tuple with the SuspiciousSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IoCIndicatorDetailed) GetSuspiciousSourcesOk() (*[]IoCSource, bool) {
	if o == nil || o.SuspiciousSources == nil {
		return nil, false
	}
	return &o.SuspiciousSources, true
}

// HasSuspiciousSources returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasSuspiciousSources() bool {
	return o != nil && o.SuspiciousSources != nil
}

// SetSuspiciousSources gets a reference to the given []IoCSource and assigns it to the SuspiciousSources field.
func (o *IoCIndicatorDetailed) SetSuspiciousSources(v []IoCSource) {
	o.SuspiciousSources = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *IoCIndicatorDetailed) SetTags(v []string) {
	o.Tags = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *IoCIndicatorDetailed) GetUsers() map[string][]string {
	if o == nil || o.Users == nil {
		var ret map[string][]string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicatorDetailed) GetUsersOk() (*map[string][]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return &o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *IoCIndicatorDetailed) HasUsers() bool {
	return o != nil && o.Users != nil
}

// SetUsers gets a reference to the given map[string][]string and assigns it to the Users field.
func (o *IoCIndicatorDetailed) SetUsers(v map[string][]string) {
	o.Users = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IoCIndicatorDetailed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AdditionalData != nil {
		toSerialize["additional_data"] = o.AdditionalData
	}
	if o.AsCidrBlock != nil {
		toSerialize["as_cidr_block"] = o.AsCidrBlock
	}
	if o.AsGeo != nil {
		toSerialize["as_geo"] = o.AsGeo
	}
	if o.AsNumber != nil {
		toSerialize["as_number"] = o.AsNumber
	}
	if o.AsOrganization != nil {
		toSerialize["as_organization"] = o.AsOrganization
	}
	if o.AsType != nil {
		toSerialize["as_type"] = o.AsType
	}
	if o.BenignSources != nil {
		toSerialize["benign_sources"] = o.BenignSources
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.CriticalAssets != nil {
		toSerialize["critical_assets"] = o.CriticalAssets
	}
	if o.FirstSeen != nil {
		if o.FirstSeen.Nanosecond() == 0 {
			toSerialize["first_seen"] = o.FirstSeen.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["first_seen"] = o.FirstSeen.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Indicator != nil {
		toSerialize["indicator"] = o.Indicator
	}
	if o.IndicatorType != nil {
		toSerialize["indicator_type"] = o.IndicatorType
	}
	if o.LastSeen != nil {
		if o.LastSeen.Nanosecond() == 0 {
			toSerialize["last_seen"] = o.LastSeen.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["last_seen"] = o.LastSeen.Format("2006-01-02T15:04:05.000Z07:00")
		}
	}
	if o.LogMatches != nil {
		toSerialize["log_matches"] = o.LogMatches
	}
	if o.LogSources != nil {
		toSerialize["log_sources"] = o.LogSources
	}
	if o.MAsType != nil {
		toSerialize["m_as_type"] = o.MAsType
	}
	if o.MPersistence != nil {
		toSerialize["m_persistence"] = o.MPersistence
	}
	if o.MSignal != nil {
		toSerialize["m_signal"] = o.MSignal
	}
	if o.MSources != nil {
		toSerialize["m_sources"] = o.MSources
	}
	if o.MaliciousSources != nil {
		toSerialize["malicious_sources"] = o.MaliciousSources
	}
	if o.MaxTrustScore != nil {
		toSerialize["max_trust_score"] = o.MaxTrustScore
	}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.SignalMatches != nil {
		toSerialize["signal_matches"] = o.SignalMatches
	}
	if o.SignalSeverity != nil {
		toSerialize["signal_severity"] = o.SignalSeverity
	}
	if o.SignalTier != nil {
		toSerialize["signal_tier"] = o.SignalTier
	}
	if o.SuspiciousSources != nil {
		toSerialize["suspicious_sources"] = o.SuspiciousSources
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IoCIndicatorDetailed) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AdditionalData    map[string]interface{}   `json:"additional_data,omitempty"`
		AsCidrBlock       *string                  `json:"as_cidr_block,omitempty"`
		AsGeo             *IoCGeoLocation          `json:"as_geo,omitempty"`
		AsNumber          *string                  `json:"as_number,omitempty"`
		AsOrganization    *string                  `json:"as_organization,omitempty"`
		AsType            *string                  `json:"as_type,omitempty"`
		BenignSources     []IoCSource              `json:"benign_sources,omitempty"`
		Categories        []string                 `json:"categories,omitempty"`
		CriticalAssets    []string                 `json:"critical_assets,omitempty"`
		FirstSeen         *time.Time               `json:"first_seen,omitempty"`
		Hosts             []string                 `json:"hosts,omitempty"`
		Id                *string                  `json:"id,omitempty"`
		Indicator         *string                  `json:"indicator,omitempty"`
		IndicatorType     *string                  `json:"indicator_type,omitempty"`
		LastSeen          *time.Time               `json:"last_seen,omitempty"`
		LogMatches        *int64                   `json:"log_matches,omitempty"`
		LogSources        []string                 `json:"log_sources,omitempty"`
		MAsType           *IoCScoreEffect          `json:"m_as_type,omitempty"`
		MPersistence      *IoCScoreEffect          `json:"m_persistence,omitempty"`
		MSignal           *IoCScoreEffect          `json:"m_signal,omitempty"`
		MSources          *IoCScoreEffect          `json:"m_sources,omitempty"`
		MaliciousSources  []IoCSource              `json:"malicious_sources,omitempty"`
		MaxTrustScore     *IoCScoreEffect          `json:"max_trust_score,omitempty"`
		Score             *float64                 `json:"score,omitempty"`
		Services          []string                 `json:"services,omitempty"`
		SignalMatches     *int64                   `json:"signal_matches,omitempty"`
		SignalSeverity    []IoCSignalSeverityCount `json:"signal_severity,omitempty"`
		SignalTier        *int64                   `json:"signal_tier,omitempty"`
		SuspiciousSources []IoCSource              `json:"suspicious_sources,omitempty"`
		Tags              []string                 `json:"tags,omitempty"`
		Users             map[string][]string      `json:"users,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"additional_data", "as_cidr_block", "as_geo", "as_number", "as_organization", "as_type", "benign_sources", "categories", "critical_assets", "first_seen", "hosts", "id", "indicator", "indicator_type", "last_seen", "log_matches", "log_sources", "m_as_type", "m_persistence", "m_signal", "m_sources", "malicious_sources", "max_trust_score", "score", "services", "signal_matches", "signal_severity", "signal_tier", "suspicious_sources", "tags", "users"})
	} else {
		return err
	}

	hasInvalidField := false
	o.AdditionalData = all.AdditionalData
	o.AsCidrBlock = all.AsCidrBlock
	if all.AsGeo != nil && all.AsGeo.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AsGeo = all.AsGeo
	o.AsNumber = all.AsNumber
	o.AsOrganization = all.AsOrganization
	o.AsType = all.AsType
	o.BenignSources = all.BenignSources
	o.Categories = all.Categories
	o.CriticalAssets = all.CriticalAssets
	o.FirstSeen = all.FirstSeen
	o.Hosts = all.Hosts
	o.Id = all.Id
	o.Indicator = all.Indicator
	o.IndicatorType = all.IndicatorType
	o.LastSeen = all.LastSeen
	o.LogMatches = all.LogMatches
	o.LogSources = all.LogSources
	if all.MAsType != nil && !all.MAsType.IsValid() {
		hasInvalidField = true
	} else {
		o.MAsType = all.MAsType
	}
	if all.MPersistence != nil && !all.MPersistence.IsValid() {
		hasInvalidField = true
	} else {
		o.MPersistence = all.MPersistence
	}
	if all.MSignal != nil && !all.MSignal.IsValid() {
		hasInvalidField = true
	} else {
		o.MSignal = all.MSignal
	}
	if all.MSources != nil && !all.MSources.IsValid() {
		hasInvalidField = true
	} else {
		o.MSources = all.MSources
	}
	o.MaliciousSources = all.MaliciousSources
	if all.MaxTrustScore != nil && !all.MaxTrustScore.IsValid() {
		hasInvalidField = true
	} else {
		o.MaxTrustScore = all.MaxTrustScore
	}
	o.Score = all.Score
	o.Services = all.Services
	o.SignalMatches = all.SignalMatches
	o.SignalSeverity = all.SignalSeverity
	o.SignalTier = all.SignalTier
	o.SuspiciousSources = all.SuspiciousSources
	o.Tags = all.Tags
	o.Users = all.Users

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
