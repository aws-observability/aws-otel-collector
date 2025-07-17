// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CsmCloudAccountsCoverageAnalysisAttributes CSM Cloud Accounts Coverage Analysis attributes.
type CsmCloudAccountsCoverageAnalysisAttributes struct {
	// CSM Coverage Analysis.
	AwsCoverage *CsmCoverageAnalysis `json:"aws_coverage,omitempty"`
	// CSM Coverage Analysis.
	AzureCoverage *CsmCoverageAnalysis `json:"azure_coverage,omitempty"`
	// CSM Coverage Analysis.
	GcpCoverage *CsmCoverageAnalysis `json:"gcp_coverage,omitempty"`
	// The ID of your organization.
	OrgId *int64 `json:"org_id,omitempty"`
	// CSM Coverage Analysis.
	TotalCoverage *CsmCoverageAnalysis `json:"total_coverage,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCsmCloudAccountsCoverageAnalysisAttributes instantiates a new CsmCloudAccountsCoverageAnalysisAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCsmCloudAccountsCoverageAnalysisAttributes() *CsmCloudAccountsCoverageAnalysisAttributes {
	this := CsmCloudAccountsCoverageAnalysisAttributes{}
	return &this
}

// NewCsmCloudAccountsCoverageAnalysisAttributesWithDefaults instantiates a new CsmCloudAccountsCoverageAnalysisAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCsmCloudAccountsCoverageAnalysisAttributesWithDefaults() *CsmCloudAccountsCoverageAnalysisAttributes {
	this := CsmCloudAccountsCoverageAnalysisAttributes{}
	return &this
}

// GetAwsCoverage returns the AwsCoverage field value if set, zero value otherwise.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) GetAwsCoverage() CsmCoverageAnalysis {
	if o == nil || o.AwsCoverage == nil {
		var ret CsmCoverageAnalysis
		return ret
	}
	return *o.AwsCoverage
}

// GetAwsCoverageOk returns a tuple with the AwsCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) GetAwsCoverageOk() (*CsmCoverageAnalysis, bool) {
	if o == nil || o.AwsCoverage == nil {
		return nil, false
	}
	return o.AwsCoverage, true
}

// HasAwsCoverage returns a boolean if a field has been set.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) HasAwsCoverage() bool {
	return o != nil && o.AwsCoverage != nil
}

// SetAwsCoverage gets a reference to the given CsmCoverageAnalysis and assigns it to the AwsCoverage field.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) SetAwsCoverage(v CsmCoverageAnalysis) {
	o.AwsCoverage = &v
}

// GetAzureCoverage returns the AzureCoverage field value if set, zero value otherwise.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) GetAzureCoverage() CsmCoverageAnalysis {
	if o == nil || o.AzureCoverage == nil {
		var ret CsmCoverageAnalysis
		return ret
	}
	return *o.AzureCoverage
}

// GetAzureCoverageOk returns a tuple with the AzureCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) GetAzureCoverageOk() (*CsmCoverageAnalysis, bool) {
	if o == nil || o.AzureCoverage == nil {
		return nil, false
	}
	return o.AzureCoverage, true
}

// HasAzureCoverage returns a boolean if a field has been set.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) HasAzureCoverage() bool {
	return o != nil && o.AzureCoverage != nil
}

// SetAzureCoverage gets a reference to the given CsmCoverageAnalysis and assigns it to the AzureCoverage field.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) SetAzureCoverage(v CsmCoverageAnalysis) {
	o.AzureCoverage = &v
}

// GetGcpCoverage returns the GcpCoverage field value if set, zero value otherwise.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) GetGcpCoverage() CsmCoverageAnalysis {
	if o == nil || o.GcpCoverage == nil {
		var ret CsmCoverageAnalysis
		return ret
	}
	return *o.GcpCoverage
}

// GetGcpCoverageOk returns a tuple with the GcpCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) GetGcpCoverageOk() (*CsmCoverageAnalysis, bool) {
	if o == nil || o.GcpCoverage == nil {
		return nil, false
	}
	return o.GcpCoverage, true
}

// HasGcpCoverage returns a boolean if a field has been set.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) HasGcpCoverage() bool {
	return o != nil && o.GcpCoverage != nil
}

// SetGcpCoverage gets a reference to the given CsmCoverageAnalysis and assigns it to the GcpCoverage field.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) SetGcpCoverage(v CsmCoverageAnalysis) {
	o.GcpCoverage = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetTotalCoverage returns the TotalCoverage field value if set, zero value otherwise.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) GetTotalCoverage() CsmCoverageAnalysis {
	if o == nil || o.TotalCoverage == nil {
		var ret CsmCoverageAnalysis
		return ret
	}
	return *o.TotalCoverage
}

// GetTotalCoverageOk returns a tuple with the TotalCoverage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) GetTotalCoverageOk() (*CsmCoverageAnalysis, bool) {
	if o == nil || o.TotalCoverage == nil {
		return nil, false
	}
	return o.TotalCoverage, true
}

// HasTotalCoverage returns a boolean if a field has been set.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) HasTotalCoverage() bool {
	return o != nil && o.TotalCoverage != nil
}

// SetTotalCoverage gets a reference to the given CsmCoverageAnalysis and assigns it to the TotalCoverage field.
func (o *CsmCloudAccountsCoverageAnalysisAttributes) SetTotalCoverage(v CsmCoverageAnalysis) {
	o.TotalCoverage = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o CsmCloudAccountsCoverageAnalysisAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AwsCoverage != nil {
		toSerialize["aws_coverage"] = o.AwsCoverage
	}
	if o.AzureCoverage != nil {
		toSerialize["azure_coverage"] = o.AzureCoverage
	}
	if o.GcpCoverage != nil {
		toSerialize["gcp_coverage"] = o.GcpCoverage
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
func (o *CsmCloudAccountsCoverageAnalysisAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AwsCoverage   *CsmCoverageAnalysis `json:"aws_coverage,omitempty"`
		AzureCoverage *CsmCoverageAnalysis `json:"azure_coverage,omitempty"`
		GcpCoverage   *CsmCoverageAnalysis `json:"gcp_coverage,omitempty"`
		OrgId         *int64               `json:"org_id,omitempty"`
		TotalCoverage *CsmCoverageAnalysis `json:"total_coverage,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aws_coverage", "azure_coverage", "gcp_coverage", "org_id", "total_coverage"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AwsCoverage != nil && all.AwsCoverage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AwsCoverage = all.AwsCoverage
	if all.AzureCoverage != nil && all.AzureCoverage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AzureCoverage = all.AzureCoverage
	if all.GcpCoverage != nil && all.GcpCoverage.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GcpCoverage = all.GcpCoverage
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
