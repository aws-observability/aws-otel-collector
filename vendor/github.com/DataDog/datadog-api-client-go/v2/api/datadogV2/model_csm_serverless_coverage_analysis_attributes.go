// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmServerlessCoverageAnalysisAttributes CSM Serverless Resources Coverage Analysis attributes.
type CsmServerlessCoverageAnalysisAttributes struct {
	// CSM Coverage Analysis.
	CwsCoverage *CsmCoverageAnalysis `json:"cws_coverage,omitempty"`
	// The ID of your organization.
	OrgId *int64 `json:"org_id,omitempty"`
	// CSM Coverage Analysis.
	TotalCoverage *CsmCoverageAnalysis `json:"total_coverage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCsmServerlessCoverageAnalysisAttributes instantiates a new CsmServerlessCoverageAnalysisAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCsmServerlessCoverageAnalysisAttributes() *CsmServerlessCoverageAnalysisAttributes {
	this := CsmServerlessCoverageAnalysisAttributes{}
	return &this
}

// NewCsmServerlessCoverageAnalysisAttributesWithDefaults instantiates a new CsmServerlessCoverageAnalysisAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCsmServerlessCoverageAnalysisAttributesWithDefaults() *CsmServerlessCoverageAnalysisAttributes {
	this := CsmServerlessCoverageAnalysisAttributes{}
	return &this
}

// GetCwsCoverage returns the CwsCoverage field value if set, zero value otherwise.
func (o *CsmServerlessCoverageAnalysisAttributes) GetCwsCoverage() CsmCoverageAnalysis {
	if o == nil || o.CwsCoverage == nil {
		var ret CsmCoverageAnalysis
		return ret
	}
	return *o.CwsCoverage
}

// GetCwsCoverageOk returns a tuple with the CwsCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmServerlessCoverageAnalysisAttributes) GetCwsCoverageOk() (*CsmCoverageAnalysis, bool) {
	if o == nil || o.CwsCoverage == nil {
		return nil, false
	}
	return o.CwsCoverage, true
}

// HasCwsCoverage returns a boolean if a field has been set.
func (o *CsmServerlessCoverageAnalysisAttributes) HasCwsCoverage() bool {
	return o != nil && o.CwsCoverage != nil
}

// SetCwsCoverage gets a reference to the given CsmCoverageAnalysis and assigns it to the CwsCoverage field.
func (o *CsmServerlessCoverageAnalysisAttributes) SetCwsCoverage(v CsmCoverageAnalysis) {
	o.CwsCoverage = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *CsmServerlessCoverageAnalysisAttributes) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmServerlessCoverageAnalysisAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *CsmServerlessCoverageAnalysisAttributes) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *CsmServerlessCoverageAnalysisAttributes) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetTotalCoverage returns the TotalCoverage field value if set, zero value otherwise.
func (o *CsmServerlessCoverageAnalysisAttributes) GetTotalCoverage() CsmCoverageAnalysis {
	if o == nil || o.TotalCoverage == nil {
		var ret CsmCoverageAnalysis
		return ret
	}
	return *o.TotalCoverage
}

// GetTotalCoverageOk returns a tuple with the TotalCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmServerlessCoverageAnalysisAttributes) GetTotalCoverageOk() (*CsmCoverageAnalysis, bool) {
	if o == nil || o.TotalCoverage == nil {
		return nil, false
	}
	return o.TotalCoverage, true
}

// HasTotalCoverage returns a boolean if a field has been set.
func (o *CsmServerlessCoverageAnalysisAttributes) HasTotalCoverage() bool {
	return o != nil && o.TotalCoverage != nil
}

// SetTotalCoverage gets a reference to the given CsmCoverageAnalysis and assigns it to the TotalCoverage field.
func (o *CsmServerlessCoverageAnalysisAttributes) SetTotalCoverage(v CsmCoverageAnalysis) {
	o.TotalCoverage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CsmServerlessCoverageAnalysisAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CsmServerlessCoverageAnalysisAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		CwsCoverage   *CsmCoverageAnalysis `json:"cws_coverage,omitempty"`
		OrgId         *int64               `json:"org_id,omitempty"`
		TotalCoverage *CsmCoverageAnalysis `json:"total_coverage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"cws_coverage", "org_id", "total_coverage"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.CwsCoverage != nil && all.CwsCoverage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.CwsCoverage = all.CwsCoverage
	o.OrgId = all.OrgId
	if all.TotalCoverage != nil && all.TotalCoverage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.TotalCoverage = all.TotalCoverage

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
