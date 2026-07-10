// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityServiceAttributes Application Security details describing a service in a given environment.
type ApplicationSecurityServiceAttributes struct {
	// The Datadog Agent versions reporting for the service.
	AgentVersions []string `json:"agent_versions"`
	// The application type of the service, such as `web` or `serverless`.
	AppType string `json:"app_type"`
	// Whether the service is compatible with Application Security Management (Threats).
	AsmThreatCompatible bool `json:"asm_threat_compatible"`
	// The number of backend WAF events detected for the service.
	BackendWafEventCount int64 `json:"backend_waf_event_count"`
	// The enabled business logic detection rules for the service.
	BusinessLogic []string `json:"business_logic"`
	// Deprecated: a display color associated with the service in the UI.
	// Deprecated
	Color string `json:"color"`
	// The environment the service runs in.
	Env string `json:"env"`
	// The number of Application Security events detected for the service.
	EventCount int64 `json:"event_count"`
	// Deprecated: the trend of Application Security events over time.
	// Deprecated
	EventTrend []int64 `json:"event_trend"`
	// Whether Application Security Management (Threats) is enabled for the service.
	HasAppsecEnabled bool `json:"has_appsec_enabled"`
	// Deprecated: the number of hits for the service.
	// Deprecated
	Hits int64 `json:"hits"`
	// Whether Interactive Application Security Testing (IAST) is enabled for the service.
	IastProductActivation bool `json:"iast_product_activation"`
	// The Interactive Application Security Testing (IAST) compatibility status of the service.
	IastProductCompatibility string `json:"iast_product_compatibility"`
	// The reasons explaining the Interactive Application Security Testing (IAST) compatibility status.
	IastProductCompatibilityReasons []string `json:"iast_product_compatibility_reasons"`
	// The programming languages detected for the service.
	Languages []string `json:"languages"`
	// The Unix timestamp, in seconds, of the last ingested span for the service.
	LastIngestedSpans int64 `json:"last_ingested_spans"`
	// The Remote Configuration capabilities reported by the service.
	RcCapabilities []string `json:"rc_capabilities"`
	// The recommended business logic detection rules for the service.
	RecommendedBusinessLogic []string `json:"recommended_business_logic"`
	// Whether Software Composition Analysis (SCA) is enabled for the service.
	RiskProductActivation bool `json:"risk_product_activation"`
	// The Software Composition Analysis (SCA) compatibility status of the service.
	RiskProductCompatibility string `json:"risk_product_compatibility"`
	// The reasons explaining the Software Composition Analysis (SCA) compatibility status.
	RiskProductCompatibilityReasons []string `json:"risk_product_compatibility_reasons"`
	// The WAF rules versions applied to the service.
	RulesVersion []string `json:"rules_version"`
	// The name of the service.
	Service string `json:"service"`
	// Deprecated: the number of security signals for the service.
	// Deprecated
	SignalCount int64 `json:"signal_count"`
	// Deprecated: the trend of security signals over time.
	// Deprecated
	SignalTrend []int64 `json:"signal_trend"`
	// The data sources that contributed information about the service.
	Source []string `json:"source"`
	// The teams that own the service.
	Teams []string `json:"teams"`
	// The Datadog tracing library versions reporting for the service.
	TracerVersions []string `json:"tracer_versions"`
	// The Vulnerability Management activation status of the service.
	VmActivation string `json:"vm-activation"`
	// Deprecated: the number of critical-severity vulnerabilities for the service.
	// Deprecated
	VulnCriticalCount int64 `json:"vuln_critical_count"`
	// Deprecated: the number of high-severity vulnerabilities for the service.
	// Deprecated
	VulnHighCount int64 `json:"vuln_high_count"`
	// The total number of services available without applying the service filter.
	WithoutFilterServices int64 `json:"without_filter_services"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityServiceAttributes instantiates a new ApplicationSecurityServiceAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityServiceAttributes(agentVersions []string, appType string, asmThreatCompatible bool, backendWafEventCount int64, businessLogic []string, color string, env string, eventCount int64, eventTrend []int64, hasAppsecEnabled bool, hits int64, iastProductActivation bool, iastProductCompatibility string, iastProductCompatibilityReasons []string, languages []string, lastIngestedSpans int64, rcCapabilities []string, recommendedBusinessLogic []string, riskProductActivation bool, riskProductCompatibility string, riskProductCompatibilityReasons []string, rulesVersion []string, service string, signalCount int64, signalTrend []int64, source []string, teams []string, tracerVersions []string, vmActivation string, vulnCriticalCount int64, vulnHighCount int64, withoutFilterServices int64) *ApplicationSecurityServiceAttributes {
	this := ApplicationSecurityServiceAttributes{}
	this.AgentVersions = agentVersions
	this.AppType = appType
	this.AsmThreatCompatible = asmThreatCompatible
	this.BackendWafEventCount = backendWafEventCount
	this.BusinessLogic = businessLogic
	this.Color = color
	this.Env = env
	this.EventCount = eventCount
	this.EventTrend = eventTrend
	this.HasAppsecEnabled = hasAppsecEnabled
	this.Hits = hits
	this.IastProductActivation = iastProductActivation
	this.IastProductCompatibility = iastProductCompatibility
	this.IastProductCompatibilityReasons = iastProductCompatibilityReasons
	this.Languages = languages
	this.LastIngestedSpans = lastIngestedSpans
	this.RcCapabilities = rcCapabilities
	this.RecommendedBusinessLogic = recommendedBusinessLogic
	this.RiskProductActivation = riskProductActivation
	this.RiskProductCompatibility = riskProductCompatibility
	this.RiskProductCompatibilityReasons = riskProductCompatibilityReasons
	this.RulesVersion = rulesVersion
	this.Service = service
	this.SignalCount = signalCount
	this.SignalTrend = signalTrend
	this.Source = source
	this.Teams = teams
	this.TracerVersions = tracerVersions
	this.VmActivation = vmActivation
	this.VulnCriticalCount = vulnCriticalCount
	this.VulnHighCount = vulnHighCount
	this.WithoutFilterServices = withoutFilterServices
	return &this
}

// NewApplicationSecurityServiceAttributesWithDefaults instantiates a new ApplicationSecurityServiceAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityServiceAttributesWithDefaults() *ApplicationSecurityServiceAttributes {
	this := ApplicationSecurityServiceAttributes{}
	return &this
}

// GetAgentVersions returns the AgentVersions field value.
func (o *ApplicationSecurityServiceAttributes) GetAgentVersions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AgentVersions
}

// GetAgentVersionsOk returns a tuple with the AgentVersions field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetAgentVersionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentVersions, true
}

// SetAgentVersions sets field value.
func (o *ApplicationSecurityServiceAttributes) SetAgentVersions(v []string) {
	o.AgentVersions = v
}

// GetAppType returns the AppType field value.
func (o *ApplicationSecurityServiceAttributes) GetAppType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.AppType
}

// GetAppTypeOk returns a tuple with the AppType field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetAppTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppType, true
}

// SetAppType sets field value.
func (o *ApplicationSecurityServiceAttributes) SetAppType(v string) {
	o.AppType = v
}

// GetAsmThreatCompatible returns the AsmThreatCompatible field value.
func (o *ApplicationSecurityServiceAttributes) GetAsmThreatCompatible() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AsmThreatCompatible
}

// GetAsmThreatCompatibleOk returns a tuple with the AsmThreatCompatible field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetAsmThreatCompatibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AsmThreatCompatible, true
}

// SetAsmThreatCompatible sets field value.
func (o *ApplicationSecurityServiceAttributes) SetAsmThreatCompatible(v bool) {
	o.AsmThreatCompatible = v
}

// GetBackendWafEventCount returns the BackendWafEventCount field value.
func (o *ApplicationSecurityServiceAttributes) GetBackendWafEventCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.BackendWafEventCount
}

// GetBackendWafEventCountOk returns a tuple with the BackendWafEventCount field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetBackendWafEventCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackendWafEventCount, true
}

// SetBackendWafEventCount sets field value.
func (o *ApplicationSecurityServiceAttributes) SetBackendWafEventCount(v int64) {
	o.BackendWafEventCount = v
}

// GetBusinessLogic returns the BusinessLogic field value.
func (o *ApplicationSecurityServiceAttributes) GetBusinessLogic() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.BusinessLogic
}

// GetBusinessLogicOk returns a tuple with the BusinessLogic field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetBusinessLogicOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessLogic, true
}

// SetBusinessLogic sets field value.
func (o *ApplicationSecurityServiceAttributes) SetBusinessLogic(v []string) {
	o.BusinessLogic = v
}

// GetColor returns the Color field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) SetColor(v string) {
	o.Color = v
}

// GetEnv returns the Env field value.
func (o *ApplicationSecurityServiceAttributes) GetEnv() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetEnvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Env, true
}

// SetEnv sets field value.
func (o *ApplicationSecurityServiceAttributes) SetEnv(v string) {
	o.Env = v
}

// GetEventCount returns the EventCount field value.
func (o *ApplicationSecurityServiceAttributes) GetEventCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.EventCount
}

// GetEventCountOk returns a tuple with the EventCount field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetEventCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCount, true
}

// SetEventCount sets field value.
func (o *ApplicationSecurityServiceAttributes) SetEventCount(v int64) {
	o.EventCount = v
}

// GetEventTrend returns the EventTrend field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetEventTrend() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.EventTrend
}

// GetEventTrendOk returns a tuple with the EventTrend field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetEventTrendOk() (*[]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventTrend, true
}

// SetEventTrend sets field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) SetEventTrend(v []int64) {
	o.EventTrend = v
}

// GetHasAppsecEnabled returns the HasAppsecEnabled field value.
func (o *ApplicationSecurityServiceAttributes) GetHasAppsecEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.HasAppsecEnabled
}

// GetHasAppsecEnabledOk returns a tuple with the HasAppsecEnabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetHasAppsecEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasAppsecEnabled, true
}

// SetHasAppsecEnabled sets field value.
func (o *ApplicationSecurityServiceAttributes) SetHasAppsecEnabled(v bool) {
	o.HasAppsecEnabled = v
}

// GetHits returns the Hits field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetHits() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.Hits
}

// GetHitsOk returns a tuple with the Hits field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetHitsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hits, true
}

// SetHits sets field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) SetHits(v int64) {
	o.Hits = v
}

// GetIastProductActivation returns the IastProductActivation field value.
func (o *ApplicationSecurityServiceAttributes) GetIastProductActivation() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.IastProductActivation
}

// GetIastProductActivationOk returns a tuple with the IastProductActivation field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetIastProductActivationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IastProductActivation, true
}

// SetIastProductActivation sets field value.
func (o *ApplicationSecurityServiceAttributes) SetIastProductActivation(v bool) {
	o.IastProductActivation = v
}

// GetIastProductCompatibility returns the IastProductCompatibility field value.
func (o *ApplicationSecurityServiceAttributes) GetIastProductCompatibility() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.IastProductCompatibility
}

// GetIastProductCompatibilityOk returns a tuple with the IastProductCompatibility field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetIastProductCompatibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IastProductCompatibility, true
}

// SetIastProductCompatibility sets field value.
func (o *ApplicationSecurityServiceAttributes) SetIastProductCompatibility(v string) {
	o.IastProductCompatibility = v
}

// GetIastProductCompatibilityReasons returns the IastProductCompatibilityReasons field value.
func (o *ApplicationSecurityServiceAttributes) GetIastProductCompatibilityReasons() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.IastProductCompatibilityReasons
}

// GetIastProductCompatibilityReasonsOk returns a tuple with the IastProductCompatibilityReasons field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetIastProductCompatibilityReasonsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IastProductCompatibilityReasons, true
}

// SetIastProductCompatibilityReasons sets field value.
func (o *ApplicationSecurityServiceAttributes) SetIastProductCompatibilityReasons(v []string) {
	o.IastProductCompatibilityReasons = v
}

// GetLanguages returns the Languages field value.
func (o *ApplicationSecurityServiceAttributes) GetLanguages() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetLanguagesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Languages, true
}

// SetLanguages sets field value.
func (o *ApplicationSecurityServiceAttributes) SetLanguages(v []string) {
	o.Languages = v
}

// GetLastIngestedSpans returns the LastIngestedSpans field value.
func (o *ApplicationSecurityServiceAttributes) GetLastIngestedSpans() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.LastIngestedSpans
}

// GetLastIngestedSpansOk returns a tuple with the LastIngestedSpans field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetLastIngestedSpansOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastIngestedSpans, true
}

// SetLastIngestedSpans sets field value.
func (o *ApplicationSecurityServiceAttributes) SetLastIngestedSpans(v int64) {
	o.LastIngestedSpans = v
}

// GetRcCapabilities returns the RcCapabilities field value.
func (o *ApplicationSecurityServiceAttributes) GetRcCapabilities() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RcCapabilities
}

// GetRcCapabilitiesOk returns a tuple with the RcCapabilities field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetRcCapabilitiesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RcCapabilities, true
}

// SetRcCapabilities sets field value.
func (o *ApplicationSecurityServiceAttributes) SetRcCapabilities(v []string) {
	o.RcCapabilities = v
}

// GetRecommendedBusinessLogic returns the RecommendedBusinessLogic field value.
func (o *ApplicationSecurityServiceAttributes) GetRecommendedBusinessLogic() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RecommendedBusinessLogic
}

// GetRecommendedBusinessLogicOk returns a tuple with the RecommendedBusinessLogic field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetRecommendedBusinessLogicOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendedBusinessLogic, true
}

// SetRecommendedBusinessLogic sets field value.
func (o *ApplicationSecurityServiceAttributes) SetRecommendedBusinessLogic(v []string) {
	o.RecommendedBusinessLogic = v
}

// GetRiskProductActivation returns the RiskProductActivation field value.
func (o *ApplicationSecurityServiceAttributes) GetRiskProductActivation() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.RiskProductActivation
}

// GetRiskProductActivationOk returns a tuple with the RiskProductActivation field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetRiskProductActivationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskProductActivation, true
}

// SetRiskProductActivation sets field value.
func (o *ApplicationSecurityServiceAttributes) SetRiskProductActivation(v bool) {
	o.RiskProductActivation = v
}

// GetRiskProductCompatibility returns the RiskProductCompatibility field value.
func (o *ApplicationSecurityServiceAttributes) GetRiskProductCompatibility() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.RiskProductCompatibility
}

// GetRiskProductCompatibilityOk returns a tuple with the RiskProductCompatibility field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetRiskProductCompatibilityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskProductCompatibility, true
}

// SetRiskProductCompatibility sets field value.
func (o *ApplicationSecurityServiceAttributes) SetRiskProductCompatibility(v string) {
	o.RiskProductCompatibility = v
}

// GetRiskProductCompatibilityReasons returns the RiskProductCompatibilityReasons field value.
func (o *ApplicationSecurityServiceAttributes) GetRiskProductCompatibilityReasons() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RiskProductCompatibilityReasons
}

// GetRiskProductCompatibilityReasonsOk returns a tuple with the RiskProductCompatibilityReasons field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetRiskProductCompatibilityReasonsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RiskProductCompatibilityReasons, true
}

// SetRiskProductCompatibilityReasons sets field value.
func (o *ApplicationSecurityServiceAttributes) SetRiskProductCompatibilityReasons(v []string) {
	o.RiskProductCompatibilityReasons = v
}

// GetRulesVersion returns the RulesVersion field value.
func (o *ApplicationSecurityServiceAttributes) GetRulesVersion() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RulesVersion
}

// GetRulesVersionOk returns a tuple with the RulesVersion field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetRulesVersionOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RulesVersion, true
}

// SetRulesVersion sets field value.
func (o *ApplicationSecurityServiceAttributes) SetRulesVersion(v []string) {
	o.RulesVersion = v
}

// GetService returns the Service field value.
func (o *ApplicationSecurityServiceAttributes) GetService() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value.
func (o *ApplicationSecurityServiceAttributes) SetService(v string) {
	o.Service = v
}

// GetSignalCount returns the SignalCount field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetSignalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.SignalCount
}

// GetSignalCountOk returns a tuple with the SignalCount field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetSignalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalCount, true
}

// SetSignalCount sets field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) SetSignalCount(v int64) {
	o.SignalCount = v
}

// GetSignalTrend returns the SignalTrend field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetSignalTrend() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.SignalTrend
}

// GetSignalTrendOk returns a tuple with the SignalTrend field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetSignalTrendOk() (*[]int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalTrend, true
}

// SetSignalTrend sets field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) SetSignalTrend(v []int64) {
	o.SignalTrend = v
}

// GetSource returns the Source field value.
func (o *ApplicationSecurityServiceAttributes) GetSource() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetSourceOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value.
func (o *ApplicationSecurityServiceAttributes) SetSource(v []string) {
	o.Source = v
}

// GetTeams returns the Teams field value.
func (o *ApplicationSecurityServiceAttributes) GetTeams() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetTeamsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Teams, true
}

// SetTeams sets field value.
func (o *ApplicationSecurityServiceAttributes) SetTeams(v []string) {
	o.Teams = v
}

// GetTracerVersions returns the TracerVersions field value.
func (o *ApplicationSecurityServiceAttributes) GetTracerVersions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.TracerVersions
}

// GetTracerVersionsOk returns a tuple with the TracerVersions field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetTracerVersionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TracerVersions, true
}

// SetTracerVersions sets field value.
func (o *ApplicationSecurityServiceAttributes) SetTracerVersions(v []string) {
	o.TracerVersions = v
}

// GetVmActivation returns the VmActivation field value.
func (o *ApplicationSecurityServiceAttributes) GetVmActivation() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.VmActivation
}

// GetVmActivationOk returns a tuple with the VmActivation field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetVmActivationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmActivation, true
}

// SetVmActivation sets field value.
func (o *ApplicationSecurityServiceAttributes) SetVmActivation(v string) {
	o.VmActivation = v
}

// GetVulnCriticalCount returns the VulnCriticalCount field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetVulnCriticalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.VulnCriticalCount
}

// GetVulnCriticalCountOk returns a tuple with the VulnCriticalCount field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetVulnCriticalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnCriticalCount, true
}

// SetVulnCriticalCount sets field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) SetVulnCriticalCount(v int64) {
	o.VulnCriticalCount = v
}

// GetVulnHighCount returns the VulnHighCount field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetVulnHighCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.VulnHighCount
}

// GetVulnHighCountOk returns a tuple with the VulnHighCount field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) GetVulnHighCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnHighCount, true
}

// SetVulnHighCount sets field value.
// Deprecated
func (o *ApplicationSecurityServiceAttributes) SetVulnHighCount(v int64) {
	o.VulnHighCount = v
}

// GetWithoutFilterServices returns the WithoutFilterServices field value.
func (o *ApplicationSecurityServiceAttributes) GetWithoutFilterServices() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.WithoutFilterServices
}

// GetWithoutFilterServicesOk returns a tuple with the WithoutFilterServices field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityServiceAttributes) GetWithoutFilterServicesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithoutFilterServices, true
}

// SetWithoutFilterServices sets field value.
func (o *ApplicationSecurityServiceAttributes) SetWithoutFilterServices(v int64) {
	o.WithoutFilterServices = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityServiceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["agent_versions"] = o.AgentVersions
	toSerialize["app_type"] = o.AppType
	toSerialize["asm_threat_compatible"] = o.AsmThreatCompatible
	toSerialize["backend_waf_event_count"] = o.BackendWafEventCount
	toSerialize["business_logic"] = o.BusinessLogic
	toSerialize["color"] = o.Color
	toSerialize["env"] = o.Env
	toSerialize["event_count"] = o.EventCount
	toSerialize["event_trend"] = o.EventTrend
	toSerialize["has_appsec_enabled"] = o.HasAppsecEnabled
	toSerialize["hits"] = o.Hits
	toSerialize["iast_product_activation"] = o.IastProductActivation
	toSerialize["iast_product_compatibility"] = o.IastProductCompatibility
	toSerialize["iast_product_compatibility_reasons"] = o.IastProductCompatibilityReasons
	toSerialize["languages"] = o.Languages
	toSerialize["last_ingested_spans"] = o.LastIngestedSpans
	toSerialize["rc_capabilities"] = o.RcCapabilities
	toSerialize["recommended_business_logic"] = o.RecommendedBusinessLogic
	toSerialize["risk_product_activation"] = o.RiskProductActivation
	toSerialize["risk_product_compatibility"] = o.RiskProductCompatibility
	toSerialize["risk_product_compatibility_reasons"] = o.RiskProductCompatibilityReasons
	toSerialize["rules_version"] = o.RulesVersion
	toSerialize["service"] = o.Service
	toSerialize["signal_count"] = o.SignalCount
	toSerialize["signal_trend"] = o.SignalTrend
	toSerialize["source"] = o.Source
	toSerialize["teams"] = o.Teams
	toSerialize["tracer_versions"] = o.TracerVersions
	toSerialize["vm-activation"] = o.VmActivation
	toSerialize["vuln_critical_count"] = o.VulnCriticalCount
	toSerialize["vuln_high_count"] = o.VulnHighCount
	toSerialize["without_filter_services"] = o.WithoutFilterServices

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityServiceAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AgentVersions                   *[]string `json:"agent_versions"`
		AppType                         *string   `json:"app_type"`
		AsmThreatCompatible             *bool     `json:"asm_threat_compatible"`
		BackendWafEventCount            *int64    `json:"backend_waf_event_count"`
		BusinessLogic                   *[]string `json:"business_logic"`
		Color                           *string   `json:"color"`
		Env                             *string   `json:"env"`
		EventCount                      *int64    `json:"event_count"`
		EventTrend                      *[]int64  `json:"event_trend"`
		HasAppsecEnabled                *bool     `json:"has_appsec_enabled"`
		Hits                            *int64    `json:"hits"`
		IastProductActivation           *bool     `json:"iast_product_activation"`
		IastProductCompatibility        *string   `json:"iast_product_compatibility"`
		IastProductCompatibilityReasons *[]string `json:"iast_product_compatibility_reasons"`
		Languages                       *[]string `json:"languages"`
		LastIngestedSpans               *int64    `json:"last_ingested_spans"`
		RcCapabilities                  *[]string `json:"rc_capabilities"`
		RecommendedBusinessLogic        *[]string `json:"recommended_business_logic"`
		RiskProductActivation           *bool     `json:"risk_product_activation"`
		RiskProductCompatibility        *string   `json:"risk_product_compatibility"`
		RiskProductCompatibilityReasons *[]string `json:"risk_product_compatibility_reasons"`
		RulesVersion                    *[]string `json:"rules_version"`
		Service                         *string   `json:"service"`
		SignalCount                     *int64    `json:"signal_count"`
		SignalTrend                     *[]int64  `json:"signal_trend"`
		Source                          *[]string `json:"source"`
		Teams                           *[]string `json:"teams"`
		TracerVersions                  *[]string `json:"tracer_versions"`
		VmActivation                    *string   `json:"vm-activation"`
		VulnCriticalCount               *int64    `json:"vuln_critical_count"`
		VulnHighCount                   *int64    `json:"vuln_high_count"`
		WithoutFilterServices           *int64    `json:"without_filter_services"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AgentVersions == nil {
		return fmt.Errorf("required field agent_versions missing")
	}
	if all.AppType == nil {
		return fmt.Errorf("required field app_type missing")
	}
	if all.AsmThreatCompatible == nil {
		return fmt.Errorf("required field asm_threat_compatible missing")
	}
	if all.BackendWafEventCount == nil {
		return fmt.Errorf("required field backend_waf_event_count missing")
	}
	if all.BusinessLogic == nil {
		return fmt.Errorf("required field business_logic missing")
	}
	if all.Color == nil {
		return fmt.Errorf("required field color missing")
	}
	if all.Env == nil {
		return fmt.Errorf("required field env missing")
	}
	if all.EventCount == nil {
		return fmt.Errorf("required field event_count missing")
	}
	if all.EventTrend == nil {
		return fmt.Errorf("required field event_trend missing")
	}
	if all.HasAppsecEnabled == nil {
		return fmt.Errorf("required field has_appsec_enabled missing")
	}
	if all.Hits == nil {
		return fmt.Errorf("required field hits missing")
	}
	if all.IastProductActivation == nil {
		return fmt.Errorf("required field iast_product_activation missing")
	}
	if all.IastProductCompatibility == nil {
		return fmt.Errorf("required field iast_product_compatibility missing")
	}
	if all.IastProductCompatibilityReasons == nil {
		return fmt.Errorf("required field iast_product_compatibility_reasons missing")
	}
	if all.Languages == nil {
		return fmt.Errorf("required field languages missing")
	}
	if all.LastIngestedSpans == nil {
		return fmt.Errorf("required field last_ingested_spans missing")
	}
	if all.RcCapabilities == nil {
		return fmt.Errorf("required field rc_capabilities missing")
	}
	if all.RecommendedBusinessLogic == nil {
		return fmt.Errorf("required field recommended_business_logic missing")
	}
	if all.RiskProductActivation == nil {
		return fmt.Errorf("required field risk_product_activation missing")
	}
	if all.RiskProductCompatibility == nil {
		return fmt.Errorf("required field risk_product_compatibility missing")
	}
	if all.RiskProductCompatibilityReasons == nil {
		return fmt.Errorf("required field risk_product_compatibility_reasons missing")
	}
	if all.RulesVersion == nil {
		return fmt.Errorf("required field rules_version missing")
	}
	if all.Service == nil {
		return fmt.Errorf("required field service missing")
	}
	if all.SignalCount == nil {
		return fmt.Errorf("required field signal_count missing")
	}
	if all.SignalTrend == nil {
		return fmt.Errorf("required field signal_trend missing")
	}
	if all.Source == nil {
		return fmt.Errorf("required field source missing")
	}
	if all.Teams == nil {
		return fmt.Errorf("required field teams missing")
	}
	if all.TracerVersions == nil {
		return fmt.Errorf("required field tracer_versions missing")
	}
	if all.VmActivation == nil {
		return fmt.Errorf("required field vm-activation missing")
	}
	if all.VulnCriticalCount == nil {
		return fmt.Errorf("required field vuln_critical_count missing")
	}
	if all.VulnHighCount == nil {
		return fmt.Errorf("required field vuln_high_count missing")
	}
	if all.WithoutFilterServices == nil {
		return fmt.Errorf("required field without_filter_services missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"agent_versions", "app_type", "asm_threat_compatible", "backend_waf_event_count", "business_logic", "color", "env", "event_count", "event_trend", "has_appsec_enabled", "hits", "iast_product_activation", "iast_product_compatibility", "iast_product_compatibility_reasons", "languages", "last_ingested_spans", "rc_capabilities", "recommended_business_logic", "risk_product_activation", "risk_product_compatibility", "risk_product_compatibility_reasons", "rules_version", "service", "signal_count", "signal_trend", "source", "teams", "tracer_versions", "vm-activation", "vuln_critical_count", "vuln_high_count", "without_filter_services"})
	} else {
		return err
	}
	o.AgentVersions = *all.AgentVersions
	o.AppType = *all.AppType
	o.AsmThreatCompatible = *all.AsmThreatCompatible
	o.BackendWafEventCount = *all.BackendWafEventCount
	o.BusinessLogic = *all.BusinessLogic
	o.Color = *all.Color
	o.Env = *all.Env
	o.EventCount = *all.EventCount
	o.EventTrend = *all.EventTrend
	o.HasAppsecEnabled = *all.HasAppsecEnabled
	o.Hits = *all.Hits
	o.IastProductActivation = *all.IastProductActivation
	o.IastProductCompatibility = *all.IastProductCompatibility
	o.IastProductCompatibilityReasons = *all.IastProductCompatibilityReasons
	o.Languages = *all.Languages
	o.LastIngestedSpans = *all.LastIngestedSpans
	o.RcCapabilities = *all.RcCapabilities
	o.RecommendedBusinessLogic = *all.RecommendedBusinessLogic
	o.RiskProductActivation = *all.RiskProductActivation
	o.RiskProductCompatibility = *all.RiskProductCompatibility
	o.RiskProductCompatibilityReasons = *all.RiskProductCompatibilityReasons
	o.RulesVersion = *all.RulesVersion
	o.Service = *all.Service
	o.SignalCount = *all.SignalCount
	o.SignalTrend = *all.SignalTrend
	o.Source = *all.Source
	o.Teams = *all.Teams
	o.TracerVersions = *all.TracerVersions
	o.VmActivation = *all.VmActivation
	o.VulnCriticalCount = *all.VulnCriticalCount
	o.VulnHighCount = *all.VulnHighCount
	o.WithoutFilterServices = *all.WithoutFilterServices

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
