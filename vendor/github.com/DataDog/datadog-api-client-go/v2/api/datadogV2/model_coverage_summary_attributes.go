// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CoverageSummaryAttributes Attributes object for code coverage summary response.
type CoverageSummaryAttributes struct {
	// Coverage statistics broken down by code owner.
	Codeowners map[string]CoverageSummaryCodeownerStats `json:"codeowners,omitempty"`
	// Total number of coverage flags evaluated.
	EvaluatedFlagsCount *int64 `json:"evaluated_flags_count,omitempty"`
	// Total number of coverage reports evaluated.
	EvaluatedReportsCount *int64 `json:"evaluated_reports_count,omitempty"`
	// Overall patch coverage percentage.
	PatchCoverage datadog.NullableFloat64 `json:"patch_coverage,omitempty"`
	// Coverage statistics broken down by service.
	Services map[string]CoverageSummaryServiceStats `json:"services,omitempty"`
	// Overall total coverage percentage.
	TotalCoverage datadog.NullableFloat64 `json:"total_coverage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCoverageSummaryAttributes instantiates a new CoverageSummaryAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCoverageSummaryAttributes() *CoverageSummaryAttributes {
	this := CoverageSummaryAttributes{}
	return &this
}

// NewCoverageSummaryAttributesWithDefaults instantiates a new CoverageSummaryAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCoverageSummaryAttributesWithDefaults() *CoverageSummaryAttributes {
	this := CoverageSummaryAttributes{}
	return &this
}

// GetCodeowners returns the Codeowners field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoverageSummaryAttributes) GetCodeowners() map[string]CoverageSummaryCodeownerStats {
	if o == nil {
		var ret map[string]CoverageSummaryCodeownerStats
		return ret
	}
	return o.Codeowners
}

// GetCodeownersOk returns a tuple with the Codeowners field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CoverageSummaryAttributes) GetCodeownersOk() (*map[string]CoverageSummaryCodeownerStats, bool) {
	if o == nil || o.Codeowners == nil {
		return nil, false
	}
	return &o.Codeowners, true
}

// HasCodeowners returns a boolean if a field has been set.
func (o *CoverageSummaryAttributes) HasCodeowners() bool {
	return o != nil && o.Codeowners != nil
}

// SetCodeowners gets a reference to the given map[string]CoverageSummaryCodeownerStats and assigns it to the Codeowners field.
func (o *CoverageSummaryAttributes) SetCodeowners(v map[string]CoverageSummaryCodeownerStats) {
	o.Codeowners = v
}

// GetEvaluatedFlagsCount returns the EvaluatedFlagsCount field value if set, zero value otherwise.
func (o *CoverageSummaryAttributes) GetEvaluatedFlagsCount() int64 {
	if o == nil || o.EvaluatedFlagsCount == nil {
		var ret int64
		return ret
	}
	return *o.EvaluatedFlagsCount
}

// GetEvaluatedFlagsCountOk returns a tuple with the EvaluatedFlagsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageSummaryAttributes) GetEvaluatedFlagsCountOk() (*int64, bool) {
	if o == nil || o.EvaluatedFlagsCount == nil {
		return nil, false
	}
	return o.EvaluatedFlagsCount, true
}

// HasEvaluatedFlagsCount returns a boolean if a field has been set.
func (o *CoverageSummaryAttributes) HasEvaluatedFlagsCount() bool {
	return o != nil && o.EvaluatedFlagsCount != nil
}

// SetEvaluatedFlagsCount gets a reference to the given int64 and assigns it to the EvaluatedFlagsCount field.
func (o *CoverageSummaryAttributes) SetEvaluatedFlagsCount(v int64) {
	o.EvaluatedFlagsCount = &v
}

// GetEvaluatedReportsCount returns the EvaluatedReportsCount field value if set, zero value otherwise.
func (o *CoverageSummaryAttributes) GetEvaluatedReportsCount() int64 {
	if o == nil || o.EvaluatedReportsCount == nil {
		var ret int64
		return ret
	}
	return *o.EvaluatedReportsCount
}

// GetEvaluatedReportsCountOk returns a tuple with the EvaluatedReportsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoverageSummaryAttributes) GetEvaluatedReportsCountOk() (*int64, bool) {
	if o == nil || o.EvaluatedReportsCount == nil {
		return nil, false
	}
	return o.EvaluatedReportsCount, true
}

// HasEvaluatedReportsCount returns a boolean if a field has been set.
func (o *CoverageSummaryAttributes) HasEvaluatedReportsCount() bool {
	return o != nil && o.EvaluatedReportsCount != nil
}

// SetEvaluatedReportsCount gets a reference to the given int64 and assigns it to the EvaluatedReportsCount field.
func (o *CoverageSummaryAttributes) SetEvaluatedReportsCount(v int64) {
	o.EvaluatedReportsCount = &v
}

// GetPatchCoverage returns the PatchCoverage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoverageSummaryAttributes) GetPatchCoverage() float64 {
	if o == nil || o.PatchCoverage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.PatchCoverage.Get()
}

// GetPatchCoverageOk returns a tuple with the PatchCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CoverageSummaryAttributes) GetPatchCoverageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PatchCoverage.Get(), o.PatchCoverage.IsSet()
}

// HasPatchCoverage returns a boolean if a field has been set.
func (o *CoverageSummaryAttributes) HasPatchCoverage() bool {
	return o != nil && o.PatchCoverage.IsSet()
}

// SetPatchCoverage gets a reference to the given datadog.NullableFloat64 and assigns it to the PatchCoverage field.
func (o *CoverageSummaryAttributes) SetPatchCoverage(v float64) {
	o.PatchCoverage.Set(&v)
}

// SetPatchCoverageNil sets the value for PatchCoverage to be an explicit nil.
func (o *CoverageSummaryAttributes) SetPatchCoverageNil() {
	o.PatchCoverage.Set(nil)
}

// UnsetPatchCoverage ensures that no value is present for PatchCoverage, not even an explicit nil.
func (o *CoverageSummaryAttributes) UnsetPatchCoverage() {
	o.PatchCoverage.Unset()
}

// GetServices returns the Services field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoverageSummaryAttributes) GetServices() map[string]CoverageSummaryServiceStats {
	if o == nil {
		var ret map[string]CoverageSummaryServiceStats
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CoverageSummaryAttributes) GetServicesOk() (*map[string]CoverageSummaryServiceStats, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return &o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *CoverageSummaryAttributes) HasServices() bool {
	return o != nil && o.Services != nil
}

// SetServices gets a reference to the given map[string]CoverageSummaryServiceStats and assigns it to the Services field.
func (o *CoverageSummaryAttributes) SetServices(v map[string]CoverageSummaryServiceStats) {
	o.Services = v
}

// GetTotalCoverage returns the TotalCoverage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CoverageSummaryAttributes) GetTotalCoverage() float64 {
	if o == nil || o.TotalCoverage.Get() == nil {
		var ret float64
		return ret
	}
	return *o.TotalCoverage.Get()
}

// GetTotalCoverageOk returns a tuple with the TotalCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *CoverageSummaryAttributes) GetTotalCoverageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalCoverage.Get(), o.TotalCoverage.IsSet()
}

// HasTotalCoverage returns a boolean if a field has been set.
func (o *CoverageSummaryAttributes) HasTotalCoverage() bool {
	return o != nil && o.TotalCoverage.IsSet()
}

// SetTotalCoverage gets a reference to the given datadog.NullableFloat64 and assigns it to the TotalCoverage field.
func (o *CoverageSummaryAttributes) SetTotalCoverage(v float64) {
	o.TotalCoverage.Set(&v)
}

// SetTotalCoverageNil sets the value for TotalCoverage to be an explicit nil.
func (o *CoverageSummaryAttributes) SetTotalCoverageNil() {
	o.TotalCoverage.Set(nil)
}

// UnsetTotalCoverage ensures that no value is present for TotalCoverage, not even an explicit nil.
func (o *CoverageSummaryAttributes) UnsetTotalCoverage() {
	o.TotalCoverage.Unset()
}

// MarshalJSON serializes the struct using spec logic.
func (o CoverageSummaryAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Codeowners != nil {
		toSerialize["codeowners"] = o.Codeowners
	}
	if o.EvaluatedFlagsCount != nil {
		toSerialize["evaluated_flags_count"] = o.EvaluatedFlagsCount
	}
	if o.EvaluatedReportsCount != nil {
		toSerialize["evaluated_reports_count"] = o.EvaluatedReportsCount
	}
	if o.PatchCoverage.IsSet() {
		toSerialize["patch_coverage"] = o.PatchCoverage.Get()
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.TotalCoverage.IsSet() {
		toSerialize["total_coverage"] = o.TotalCoverage.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CoverageSummaryAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Codeowners            map[string]CoverageSummaryCodeownerStats `json:"codeowners,omitempty"`
		EvaluatedFlagsCount   *int64                                   `json:"evaluated_flags_count,omitempty"`
		EvaluatedReportsCount *int64                                   `json:"evaluated_reports_count,omitempty"`
		PatchCoverage         datadog.NullableFloat64                  `json:"patch_coverage,omitempty"`
		Services              map[string]CoverageSummaryServiceStats   `json:"services,omitempty"`
		TotalCoverage         datadog.NullableFloat64                  `json:"total_coverage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"codeowners", "evaluated_flags_count", "evaluated_reports_count", "patch_coverage", "services", "total_coverage"})
	} else {
		return err
	}
	o.Codeowners = all.Codeowners
	o.EvaluatedFlagsCount = all.EvaluatedFlagsCount
	o.EvaluatedReportsCount = all.EvaluatedReportsCount
	o.PatchCoverage = all.PatchCoverage
	o.Services = all.Services
	o.TotalCoverage = all.TotalCoverage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
