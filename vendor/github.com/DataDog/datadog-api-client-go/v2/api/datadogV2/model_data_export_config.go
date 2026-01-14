// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataExportConfig AWS Cost and Usage Report data export configuration.
type DataExportConfig struct {
	// Name of the S3 bucket where the Cost and Usage Report is stored.
	BucketName *string `json:"bucket_name,omitempty"`
	// AWS region of the S3 bucket.
	BucketRegion *string `json:"bucket_region,omitempty"`
	// Name of the Cost and Usage Report.
	ReportName *string `json:"report_name,omitempty"`
	// S3 prefix where the Cost and Usage Report is stored.
	ReportPrefix *string `json:"report_prefix,omitempty"`
	// Type of the Cost and Usage Report.
	ReportType *string `json:"report_type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataExportConfig instantiates a new DataExportConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataExportConfig() *DataExportConfig {
	this := DataExportConfig{}
	return &this
}

// NewDataExportConfigWithDefaults instantiates a new DataExportConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataExportConfigWithDefaults() *DataExportConfig {
	this := DataExportConfig{}
	return &this
}

// GetBucketName returns the BucketName field value if set, zero value otherwise.
func (o *DataExportConfig) GetBucketName() string {
	if o == nil || o.BucketName == nil {
		var ret string
		return ret
	}
	return *o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataExportConfig) GetBucketNameOk() (*string, bool) {
	if o == nil || o.BucketName == nil {
		return nil, false
	}
	return o.BucketName, true
}

// HasBucketName returns a boolean if a field has been set.
func (o *DataExportConfig) HasBucketName() bool {
	return o != nil && o.BucketName != nil
}

// SetBucketName gets a reference to the given string and assigns it to the BucketName field.
func (o *DataExportConfig) SetBucketName(v string) {
	o.BucketName = &v
}

// GetBucketRegion returns the BucketRegion field value if set, zero value otherwise.
func (o *DataExportConfig) GetBucketRegion() string {
	if o == nil || o.BucketRegion == nil {
		var ret string
		return ret
	}
	return *o.BucketRegion
}

// GetBucketRegionOk returns a tuple with the BucketRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataExportConfig) GetBucketRegionOk() (*string, bool) {
	if o == nil || o.BucketRegion == nil {
		return nil, false
	}
	return o.BucketRegion, true
}

// HasBucketRegion returns a boolean if a field has been set.
func (o *DataExportConfig) HasBucketRegion() bool {
	return o != nil && o.BucketRegion != nil
}

// SetBucketRegion gets a reference to the given string and assigns it to the BucketRegion field.
func (o *DataExportConfig) SetBucketRegion(v string) {
	o.BucketRegion = &v
}

// GetReportName returns the ReportName field value if set, zero value otherwise.
func (o *DataExportConfig) GetReportName() string {
	if o == nil || o.ReportName == nil {
		var ret string
		return ret
	}
	return *o.ReportName
}

// GetReportNameOk returns a tuple with the ReportName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataExportConfig) GetReportNameOk() (*string, bool) {
	if o == nil || o.ReportName == nil {
		return nil, false
	}
	return o.ReportName, true
}

// HasReportName returns a boolean if a field has been set.
func (o *DataExportConfig) HasReportName() bool {
	return o != nil && o.ReportName != nil
}

// SetReportName gets a reference to the given string and assigns it to the ReportName field.
func (o *DataExportConfig) SetReportName(v string) {
	o.ReportName = &v
}

// GetReportPrefix returns the ReportPrefix field value if set, zero value otherwise.
func (o *DataExportConfig) GetReportPrefix() string {
	if o == nil || o.ReportPrefix == nil {
		var ret string
		return ret
	}
	return *o.ReportPrefix
}

// GetReportPrefixOk returns a tuple with the ReportPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataExportConfig) GetReportPrefixOk() (*string, bool) {
	if o == nil || o.ReportPrefix == nil {
		return nil, false
	}
	return o.ReportPrefix, true
}

// HasReportPrefix returns a boolean if a field has been set.
func (o *DataExportConfig) HasReportPrefix() bool {
	return o != nil && o.ReportPrefix != nil
}

// SetReportPrefix gets a reference to the given string and assigns it to the ReportPrefix field.
func (o *DataExportConfig) SetReportPrefix(v string) {
	o.ReportPrefix = &v
}

// GetReportType returns the ReportType field value if set, zero value otherwise.
func (o *DataExportConfig) GetReportType() string {
	if o == nil || o.ReportType == nil {
		var ret string
		return ret
	}
	return *o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataExportConfig) GetReportTypeOk() (*string, bool) {
	if o == nil || o.ReportType == nil {
		return nil, false
	}
	return o.ReportType, true
}

// HasReportType returns a boolean if a field has been set.
func (o *DataExportConfig) HasReportType() bool {
	return o != nil && o.ReportType != nil
}

// SetReportType gets a reference to the given string and assigns it to the ReportType field.
func (o *DataExportConfig) SetReportType(v string) {
	o.ReportType = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataExportConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.BucketName != nil {
		toSerialize["bucket_name"] = o.BucketName
	}
	if o.BucketRegion != nil {
		toSerialize["bucket_region"] = o.BucketRegion
	}
	if o.ReportName != nil {
		toSerialize["report_name"] = o.ReportName
	}
	if o.ReportPrefix != nil {
		toSerialize["report_prefix"] = o.ReportPrefix
	}
	if o.ReportType != nil {
		toSerialize["report_type"] = o.ReportType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataExportConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BucketName   *string `json:"bucket_name,omitempty"`
		BucketRegion *string `json:"bucket_region,omitempty"`
		ReportName   *string `json:"report_name,omitempty"`
		ReportPrefix *string `json:"report_prefix,omitempty"`
		ReportType   *string `json:"report_type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bucket_name", "bucket_region", "report_name", "report_prefix", "report_type"})
	} else {
		return err
	}
	o.BucketName = all.BucketName
	o.BucketRegion = all.BucketRegion
	o.ReportName = all.ReportName
	o.ReportPrefix = all.ReportPrefix
	o.ReportType = all.ReportType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
