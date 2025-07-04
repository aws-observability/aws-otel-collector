// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmHostsAndContainersCoverageAnalysisAttributes CSM Hosts and Containers Coverage Analysis attributes.
type CsmHostsAndContainersCoverageAnalysisAttributes struct {
	// CSM Coverage Analysis.
	CspmCoverage *CsmCoverageAnalysis `json:"cspm_coverage,omitempty"`
	// CSM Coverage Analysis.
	CwsCoverage *CsmCoverageAnalysis `json:"cws_coverage,omitempty"`
	// The ID of your organization.
	OrgId *int64 `json:"org_id,omitempty"`
	// CSM Coverage Analysis.
	TotalCoverage *CsmCoverageAnalysis `json:"total_coverage,omitempty"`
	// CSM Coverage Analysis.
	VmCoverage *CsmCoverageAnalysis `json:"vm_coverage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCsmHostsAndContainersCoverageAnalysisAttributes instantiates a new CsmHostsAndContainersCoverageAnalysisAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCsmHostsAndContainersCoverageAnalysisAttributes() *CsmHostsAndContainersCoverageAnalysisAttributes {
	this := CsmHostsAndContainersCoverageAnalysisAttributes{}
	return &this
}

// NewCsmHostsAndContainersCoverageAnalysisAttributesWithDefaults instantiates a new CsmHostsAndContainersCoverageAnalysisAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCsmHostsAndContainersCoverageAnalysisAttributesWithDefaults() *CsmHostsAndContainersCoverageAnalysisAttributes {
	this := CsmHostsAndContainersCoverageAnalysisAttributes{}
	return &this
}

// GetCspmCoverage returns the CspmCoverage field value if set, zero value otherwise.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) GetCspmCoverage() CsmCoverageAnalysis {
	if o == nil || o.CspmCoverage == nil {
		var ret CsmCoverageAnalysis
		return ret
	}
	return *o.CspmCoverage
}

// GetCspmCoverageOk returns a tuple with the CspmCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) GetCspmCoverageOk() (*CsmCoverageAnalysis, bool) {
	if o == nil || o.CspmCoverage == nil {
		return nil, false
	}
	return o.CspmCoverage, true
}

// HasCspmCoverage returns a boolean if a field has been set.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) HasCspmCoverage() bool {
	return o != nil && o.CspmCoverage != nil
}

// SetCspmCoverage gets a reference to the given CsmCoverageAnalysis and assigns it to the CspmCoverage field.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) SetCspmCoverage(v CsmCoverageAnalysis) {
	o.CspmCoverage = &v
}

// GetCwsCoverage returns the CwsCoverage field value if set, zero value otherwise.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) GetCwsCoverage() CsmCoverageAnalysis {
	if o == nil || o.CwsCoverage == nil {
		var ret CsmCoverageAnalysis
		return ret
	}
	return *o.CwsCoverage
}

// GetCwsCoverageOk returns a tuple with the CwsCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) GetCwsCoverageOk() (*CsmCoverageAnalysis, bool) {
	if o == nil || o.CwsCoverage == nil {
		return nil, false
	}
	return o.CwsCoverage, true
}

// HasCwsCoverage returns a boolean if a field has been set.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) HasCwsCoverage() bool {
	return o != nil && o.CwsCoverage != nil
}

// SetCwsCoverage gets a reference to the given CsmCoverageAnalysis and assigns it to the CwsCoverage field.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) SetCwsCoverage(v CsmCoverageAnalysis) {
	o.CwsCoverage = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetTotalCoverage returns the TotalCoverage field value if set, zero value otherwise.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) GetTotalCoverage() CsmCoverageAnalysis {
	if o == nil || o.TotalCoverage == nil {
		var ret CsmCoverageAnalysis
		return ret
	}
	return *o.TotalCoverage
}

// GetTotalCoverageOk returns a tuple with the TotalCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) GetTotalCoverageOk() (*CsmCoverageAnalysis, bool) {
	if o == nil || o.TotalCoverage == nil {
		return nil, false
	}
	return o.TotalCoverage, true
}

// HasTotalCoverage returns a boolean if a field has been set.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) HasTotalCoverage() bool {
	return o != nil && o.TotalCoverage != nil
}

// SetTotalCoverage gets a reference to the given CsmCoverageAnalysis and assigns it to the TotalCoverage field.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) SetTotalCoverage(v CsmCoverageAnalysis) {
	o.TotalCoverage = &v
}

// GetVmCoverage returns the VmCoverage field value if set, zero value otherwise.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) GetVmCoverage() CsmCoverageAnalysis {
	if o == nil || o.VmCoverage == nil {
		var ret CsmCoverageAnalysis
		return ret
	}
	return *o.VmCoverage
}

// GetVmCoverageOk returns a tuple with the VmCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) GetVmCoverageOk() (*CsmCoverageAnalysis, bool) {
	if o == nil || o.VmCoverage == nil {
		return nil, false
	}
	return o.VmCoverage, true
}

// HasVmCoverage returns a boolean if a field has been set.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) HasVmCoverage() bool {
	return o != nil && o.VmCoverage != nil
}

// SetVmCoverage gets a reference to the given CsmCoverageAnalysis and assigns it to the VmCoverage field.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) SetVmCoverage(v CsmCoverageAnalysis) {
	o.VmCoverage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CsmHostsAndContainersCoverageAnalysisAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.CspmCoverage != nil {
		toSerialize["cspm_coverage"] = o.CspmCoverage
	}
	if o.CwsCoverage != nil {
		toSerialize["cws_coverage"] = o.CwsCoverage
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.TotalCoverage != nil {
		toSerialize["total_coverage"] = o.TotalCoverage
	}
	if o.VmCoverage != nil {
		toSerialize["vm_coverage"] = o.VmCoverage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CsmHostsAndContainersCoverageAnalysisAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CspmCoverage  *CsmCoverageAnalysis `json:"cspm_coverage,omitempty"`
		CwsCoverage   *CsmCoverageAnalysis `json:"cws_coverage,omitempty"`
		OrgId         *int64               `json:"org_id,omitempty"`
		TotalCoverage *CsmCoverageAnalysis `json:"total_coverage,omitempty"`
		VmCoverage    *CsmCoverageAnalysis `json:"vm_coverage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cspm_coverage", "cws_coverage", "org_id", "total_coverage", "vm_coverage"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CspmCoverage != nil && all.CspmCoverage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CspmCoverage = all.CspmCoverage
	if all.CwsCoverage != nil && all.CwsCoverage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CwsCoverage = all.CwsCoverage
	o.OrgId = all.OrgId
	if all.TotalCoverage != nil && all.TotalCoverage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TotalCoverage = all.TotalCoverage
	if all.VmCoverage != nil && all.VmCoverage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.VmCoverage = all.VmCoverage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
