// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// IoCIndicator An indicator of compromise with threat intelligence data.
type IoCIndicator struct {
	// Geographic location information for an IP indicator.
	AsGeo *IoCGeoLocation `json:"as_geo,omitempty"`
	// Autonomous system type.
	AsType *string `json:"as_type,omitempty"`
	// Threat intelligence sources that flagged this indicator as benign.
	BenignSources []IoCSource `json:"benign_sources,omitempty"`
	// Threat categories associated with the indicator.
	Categories []string `json:"categories,omitempty"`
	// Timestamp when the indicator was first seen.
	FirstSeen *time.Time `json:"first_seen,omitempty"`
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
	// Number of security signals that matched this indicator.
	SignalMatches *int64 `json:"signal_matches,omitempty"`
	// Signal tier level.
	SignalTier *int64 `json:"signal_tier,omitempty"`
	// Threat intelligence sources that flagged this indicator as suspicious.
	SuspiciousSources []IoCSource `json:"suspicious_sources,omitempty"`
	// Tags associated with the indicator.
	Tags []string `json:"tags,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewIoCIndicator instantiates a new IoCIndicator object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewIoCIndicator() *IoCIndicator {
	this := IoCIndicator{}
	return &this
}

// NewIoCIndicatorWithDefaults instantiates a new IoCIndicator object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewIoCIndicatorWithDefaults() *IoCIndicator {
	this := IoCIndicator{}
	return &this
}

// GetAsGeo returns the AsGeo field value if set, zero value otherwise.
func (o *IoCIndicator) GetAsGeo() IoCGeoLocation {
	if o == nil || o.AsGeo == nil {
		var ret IoCGeoLocation
		return ret
	}
	return *o.AsGeo
}

// GetAsGeoOk returns a tuple with the AsGeo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetAsGeoOk() (*IoCGeoLocation, bool) {
	if o == nil || o.AsGeo == nil {
		return nil, false
	}
	return o.AsGeo, true
}

// HasAsGeo returns a boolean if a field has been set.
func (o *IoCIndicator) HasAsGeo() bool {
	return o != nil && o.AsGeo != nil
}

// SetAsGeo gets a reference to the given IoCGeoLocation and assigns it to the AsGeo field.
func (o *IoCIndicator) SetAsGeo(v IoCGeoLocation) {
	o.AsGeo = &v
}

// GetAsType returns the AsType field value if set, zero value otherwise.
func (o *IoCIndicator) GetAsType() string {
	if o == nil || o.AsType == nil {
		var ret string
		return ret
	}
	return *o.AsType
}

// GetAsTypeOk returns a tuple with the AsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetAsTypeOk() (*string, bool) {
	if o == nil || o.AsType == nil {
		return nil, false
	}
	return o.AsType, true
}

// HasAsType returns a boolean if a field has been set.
func (o *IoCIndicator) HasAsType() bool {
	return o != nil && o.AsType != nil
}

// SetAsType gets a reference to the given string and assigns it to the AsType field.
func (o *IoCIndicator) SetAsType(v string) {
	o.AsType = &v
}

// GetBenignSources returns the BenignSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IoCIndicator) GetBenignSources() []IoCSource {
	if o == nil {
		var ret []IoCSource
		return ret
	}
	return o.BenignSources
}

// GetBenignSourcesOk returns a tuple with the BenignSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IoCIndicator) GetBenignSourcesOk() (*[]IoCSource, bool) {
	if o == nil || o.BenignSources == nil {
		return nil, false
	}
	return &o.BenignSources, true
}

// HasBenignSources returns a boolean if a field has been set.
func (o *IoCIndicator) HasBenignSources() bool {
	return o != nil && o.BenignSources != nil
}

// SetBenignSources gets a reference to the given []IoCSource and assigns it to the BenignSources field.
func (o *IoCIndicator) SetBenignSources(v []IoCSource) {
	o.BenignSources = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *IoCIndicator) GetCategories() []string {
	if o == nil || o.Categories == nil {
		var ret []string
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetCategoriesOk() (*[]string, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return &o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *IoCIndicator) HasCategories() bool {
	return o != nil && o.Categories != nil
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *IoCIndicator) SetCategories(v []string) {
	o.Categories = v
}

// GetFirstSeen returns the FirstSeen field value if set, zero value otherwise.
func (o *IoCIndicator) GetFirstSeen() time.Time {
	if o == nil || o.FirstSeen == nil {
		var ret time.Time
		return ret
	}
	return *o.FirstSeen
}

// GetFirstSeenOk returns a tuple with the FirstSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetFirstSeenOk() (*time.Time, bool) {
	if o == nil || o.FirstSeen == nil {
		return nil, false
	}
	return o.FirstSeen, true
}

// HasFirstSeen returns a boolean if a field has been set.
func (o *IoCIndicator) HasFirstSeen() bool {
	return o != nil && o.FirstSeen != nil
}

// SetFirstSeen gets a reference to the given time.Time and assigns it to the FirstSeen field.
func (o *IoCIndicator) SetFirstSeen(v time.Time) {
	o.FirstSeen = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IoCIndicator) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IoCIndicator) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IoCIndicator) SetId(v string) {
	o.Id = &v
}

// GetIndicator returns the Indicator field value if set, zero value otherwise.
func (o *IoCIndicator) GetIndicator() string {
	if o == nil || o.Indicator == nil {
		var ret string
		return ret
	}
	return *o.Indicator
}

// GetIndicatorOk returns a tuple with the Indicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetIndicatorOk() (*string, bool) {
	if o == nil || o.Indicator == nil {
		return nil, false
	}
	return o.Indicator, true
}

// HasIndicator returns a boolean if a field has been set.
func (o *IoCIndicator) HasIndicator() bool {
	return o != nil && o.Indicator != nil
}

// SetIndicator gets a reference to the given string and assigns it to the Indicator field.
func (o *IoCIndicator) SetIndicator(v string) {
	o.Indicator = &v
}

// GetIndicatorType returns the IndicatorType field value if set, zero value otherwise.
func (o *IoCIndicator) GetIndicatorType() string {
	if o == nil || o.IndicatorType == nil {
		var ret string
		return ret
	}
	return *o.IndicatorType
}

// GetIndicatorTypeOk returns a tuple with the IndicatorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetIndicatorTypeOk() (*string, bool) {
	if o == nil || o.IndicatorType == nil {
		return nil, false
	}
	return o.IndicatorType, true
}

// HasIndicatorType returns a boolean if a field has been set.
func (o *IoCIndicator) HasIndicatorType() bool {
	return o != nil && o.IndicatorType != nil
}

// SetIndicatorType gets a reference to the given string and assigns it to the IndicatorType field.
func (o *IoCIndicator) SetIndicatorType(v string) {
	o.IndicatorType = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *IoCIndicator) GetLastSeen() time.Time {
	if o == nil || o.LastSeen == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetLastSeenOk() (*time.Time, bool) {
	if o == nil || o.LastSeen == nil {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *IoCIndicator) HasLastSeen() bool {
	return o != nil && o.LastSeen != nil
}

// SetLastSeen gets a reference to the given time.Time and assigns it to the LastSeen field.
func (o *IoCIndicator) SetLastSeen(v time.Time) {
	o.LastSeen = &v
}

// GetLogMatches returns the LogMatches field value if set, zero value otherwise.
func (o *IoCIndicator) GetLogMatches() int64 {
	if o == nil || o.LogMatches == nil {
		var ret int64
		return ret
	}
	return *o.LogMatches
}

// GetLogMatchesOk returns a tuple with the LogMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetLogMatchesOk() (*int64, bool) {
	if o == nil || o.LogMatches == nil {
		return nil, false
	}
	return o.LogMatches, true
}

// HasLogMatches returns a boolean if a field has been set.
func (o *IoCIndicator) HasLogMatches() bool {
	return o != nil && o.LogMatches != nil
}

// SetLogMatches gets a reference to the given int64 and assigns it to the LogMatches field.
func (o *IoCIndicator) SetLogMatches(v int64) {
	o.LogMatches = &v
}

// GetMAsType returns the MAsType field value if set, zero value otherwise.
func (o *IoCIndicator) GetMAsType() IoCScoreEffect {
	if o == nil || o.MAsType == nil {
		var ret IoCScoreEffect
		return ret
	}
	return *o.MAsType
}

// GetMAsTypeOk returns a tuple with the MAsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetMAsTypeOk() (*IoCScoreEffect, bool) {
	if o == nil || o.MAsType == nil {
		return nil, false
	}
	return o.MAsType, true
}

// HasMAsType returns a boolean if a field has been set.
func (o *IoCIndicator) HasMAsType() bool {
	return o != nil && o.MAsType != nil
}

// SetMAsType gets a reference to the given IoCScoreEffect and assigns it to the MAsType field.
func (o *IoCIndicator) SetMAsType(v IoCScoreEffect) {
	o.MAsType = &v
}

// GetMPersistence returns the MPersistence field value if set, zero value otherwise.
func (o *IoCIndicator) GetMPersistence() IoCScoreEffect {
	if o == nil || o.MPersistence == nil {
		var ret IoCScoreEffect
		return ret
	}
	return *o.MPersistence
}

// GetMPersistenceOk returns a tuple with the MPersistence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetMPersistenceOk() (*IoCScoreEffect, bool) {
	if o == nil || o.MPersistence == nil {
		return nil, false
	}
	return o.MPersistence, true
}

// HasMPersistence returns a boolean if a field has been set.
func (o *IoCIndicator) HasMPersistence() bool {
	return o != nil && o.MPersistence != nil
}

// SetMPersistence gets a reference to the given IoCScoreEffect and assigns it to the MPersistence field.
func (o *IoCIndicator) SetMPersistence(v IoCScoreEffect) {
	o.MPersistence = &v
}

// GetMSignal returns the MSignal field value if set, zero value otherwise.
func (o *IoCIndicator) GetMSignal() IoCScoreEffect {
	if o == nil || o.MSignal == nil {
		var ret IoCScoreEffect
		return ret
	}
	return *o.MSignal
}

// GetMSignalOk returns a tuple with the MSignal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetMSignalOk() (*IoCScoreEffect, bool) {
	if o == nil || o.MSignal == nil {
		return nil, false
	}
	return o.MSignal, true
}

// HasMSignal returns a boolean if a field has been set.
func (o *IoCIndicator) HasMSignal() bool {
	return o != nil && o.MSignal != nil
}

// SetMSignal gets a reference to the given IoCScoreEffect and assigns it to the MSignal field.
func (o *IoCIndicator) SetMSignal(v IoCScoreEffect) {
	o.MSignal = &v
}

// GetMSources returns the MSources field value if set, zero value otherwise.
func (o *IoCIndicator) GetMSources() IoCScoreEffect {
	if o == nil || o.MSources == nil {
		var ret IoCScoreEffect
		return ret
	}
	return *o.MSources
}

// GetMSourcesOk returns a tuple with the MSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetMSourcesOk() (*IoCScoreEffect, bool) {
	if o == nil || o.MSources == nil {
		return nil, false
	}
	return o.MSources, true
}

// HasMSources returns a boolean if a field has been set.
func (o *IoCIndicator) HasMSources() bool {
	return o != nil && o.MSources != nil
}

// SetMSources gets a reference to the given IoCScoreEffect and assigns it to the MSources field.
func (o *IoCIndicator) SetMSources(v IoCScoreEffect) {
	o.MSources = &v
}

// GetMaliciousSources returns the MaliciousSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IoCIndicator) GetMaliciousSources() []IoCSource {
	if o == nil {
		var ret []IoCSource
		return ret
	}
	return o.MaliciousSources
}

// GetMaliciousSourcesOk returns a tuple with the MaliciousSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IoCIndicator) GetMaliciousSourcesOk() (*[]IoCSource, bool) {
	if o == nil || o.MaliciousSources == nil {
		return nil, false
	}
	return &o.MaliciousSources, true
}

// HasMaliciousSources returns a boolean if a field has been set.
func (o *IoCIndicator) HasMaliciousSources() bool {
	return o != nil && o.MaliciousSources != nil
}

// SetMaliciousSources gets a reference to the given []IoCSource and assigns it to the MaliciousSources field.
func (o *IoCIndicator) SetMaliciousSources(v []IoCSource) {
	o.MaliciousSources = v
}

// GetMaxTrustScore returns the MaxTrustScore field value if set, zero value otherwise.
func (o *IoCIndicator) GetMaxTrustScore() IoCScoreEffect {
	if o == nil || o.MaxTrustScore == nil {
		var ret IoCScoreEffect
		return ret
	}
	return *o.MaxTrustScore
}

// GetMaxTrustScoreOk returns a tuple with the MaxTrustScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetMaxTrustScoreOk() (*IoCScoreEffect, bool) {
	if o == nil || o.MaxTrustScore == nil {
		return nil, false
	}
	return o.MaxTrustScore, true
}

// HasMaxTrustScore returns a boolean if a field has been set.
func (o *IoCIndicator) HasMaxTrustScore() bool {
	return o != nil && o.MaxTrustScore != nil
}

// SetMaxTrustScore gets a reference to the given IoCScoreEffect and assigns it to the MaxTrustScore field.
func (o *IoCIndicator) SetMaxTrustScore(v IoCScoreEffect) {
	o.MaxTrustScore = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *IoCIndicator) GetScore() float64 {
	if o == nil || o.Score == nil {
		var ret float64
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetScoreOk() (*float64, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *IoCIndicator) HasScore() bool {
	return o != nil && o.Score != nil
}

// SetScore gets a reference to the given float64 and assigns it to the Score field.
func (o *IoCIndicator) SetScore(v float64) {
	o.Score = &v
}

// GetSignalMatches returns the SignalMatches field value if set, zero value otherwise.
func (o *IoCIndicator) GetSignalMatches() int64 {
	if o == nil || o.SignalMatches == nil {
		var ret int64
		return ret
	}
	return *o.SignalMatches
}

// GetSignalMatchesOk returns a tuple with the SignalMatches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetSignalMatchesOk() (*int64, bool) {
	if o == nil || o.SignalMatches == nil {
		return nil, false
	}
	return o.SignalMatches, true
}

// HasSignalMatches returns a boolean if a field has been set.
func (o *IoCIndicator) HasSignalMatches() bool {
	return o != nil && o.SignalMatches != nil
}

// SetSignalMatches gets a reference to the given int64 and assigns it to the SignalMatches field.
func (o *IoCIndicator) SetSignalMatches(v int64) {
	o.SignalMatches = &v
}

// GetSignalTier returns the SignalTier field value if set, zero value otherwise.
func (o *IoCIndicator) GetSignalTier() int64 {
	if o == nil || o.SignalTier == nil {
		var ret int64
		return ret
	}
	return *o.SignalTier
}

// GetSignalTierOk returns a tuple with the SignalTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetSignalTierOk() (*int64, bool) {
	if o == nil || o.SignalTier == nil {
		return nil, false
	}
	return o.SignalTier, true
}

// HasSignalTier returns a boolean if a field has been set.
func (o *IoCIndicator) HasSignalTier() bool {
	return o != nil && o.SignalTier != nil
}

// SetSignalTier gets a reference to the given int64 and assigns it to the SignalTier field.
func (o *IoCIndicator) SetSignalTier(v int64) {
	o.SignalTier = &v
}

// GetSuspiciousSources returns the SuspiciousSources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IoCIndicator) GetSuspiciousSources() []IoCSource {
	if o == nil {
		var ret []IoCSource
		return ret
	}
	return o.SuspiciousSources
}

// GetSuspiciousSourcesOk returns a tuple with the SuspiciousSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *IoCIndicator) GetSuspiciousSourcesOk() (*[]IoCSource, bool) {
	if o == nil || o.SuspiciousSources == nil {
		return nil, false
	}
	return &o.SuspiciousSources, true
}

// HasSuspiciousSources returns a boolean if a field has been set.
func (o *IoCIndicator) HasSuspiciousSources() bool {
	return o != nil && o.SuspiciousSources != nil
}

// SetSuspiciousSources gets a reference to the given []IoCSource and assigns it to the SuspiciousSources field.
func (o *IoCIndicator) SetSuspiciousSources(v []IoCSource) {
	o.SuspiciousSources = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IoCIndicator) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IoCIndicator) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IoCIndicator) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *IoCIndicator) SetTags(v []string) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o IoCIndicator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AsGeo != nil {
		toSerialize["as_geo"] = o.AsGeo
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
	if o.FirstSeen != nil {
		if o.FirstSeen.Nanosecond() == 0 {
			toSerialize["first_seen"] = o.FirstSeen.Format("2006-01-02T15:04:05Z07:00")
		} else {
			toSerialize["first_seen"] = o.FirstSeen.Format("2006-01-02T15:04:05.000Z07:00")
		}
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
	if o.SignalMatches != nil {
		toSerialize["signal_matches"] = o.SignalMatches
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *IoCIndicator) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AsGeo             *IoCGeoLocation `json:"as_geo,omitempty"`
		AsType            *string         `json:"as_type,omitempty"`
		BenignSources     []IoCSource     `json:"benign_sources,omitempty"`
		Categories        []string        `json:"categories,omitempty"`
		FirstSeen         *time.Time      `json:"first_seen,omitempty"`
		Id                *string         `json:"id,omitempty"`
		Indicator         *string         `json:"indicator,omitempty"`
		IndicatorType     *string         `json:"indicator_type,omitempty"`
		LastSeen          *time.Time      `json:"last_seen,omitempty"`
		LogMatches        *int64          `json:"log_matches,omitempty"`
		MAsType           *IoCScoreEffect `json:"m_as_type,omitempty"`
		MPersistence      *IoCScoreEffect `json:"m_persistence,omitempty"`
		MSignal           *IoCScoreEffect `json:"m_signal,omitempty"`
		MSources          *IoCScoreEffect `json:"m_sources,omitempty"`
		MaliciousSources  []IoCSource     `json:"malicious_sources,omitempty"`
		MaxTrustScore     *IoCScoreEffect `json:"max_trust_score,omitempty"`
		Score             *float64        `json:"score,omitempty"`
		SignalMatches     *int64          `json:"signal_matches,omitempty"`
		SignalTier        *int64          `json:"signal_tier,omitempty"`
		SuspiciousSources []IoCSource     `json:"suspicious_sources,omitempty"`
		Tags              []string        `json:"tags,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"as_geo", "as_type", "benign_sources", "categories", "first_seen", "id", "indicator", "indicator_type", "last_seen", "log_matches", "m_as_type", "m_persistence", "m_signal", "m_sources", "malicious_sources", "max_trust_score", "score", "signal_matches", "signal_tier", "suspicious_sources", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AsGeo != nil && all.AsGeo.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AsGeo = all.AsGeo
	o.AsType = all.AsType
	o.BenignSources = all.BenignSources
	o.Categories = all.Categories
	o.FirstSeen = all.FirstSeen
	o.Id = all.Id
	o.Indicator = all.Indicator
	o.IndicatorType = all.IndicatorType
	o.LastSeen = all.LastSeen
	o.LogMatches = all.LogMatches
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
	o.SignalMatches = all.SignalMatches
	o.SignalTier = all.SignalTier
	o.SuspiciousSources = all.SuspiciousSources
	o.Tags = all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
