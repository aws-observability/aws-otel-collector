// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// DataExportConfig AWS Cost and Usage Report data export configuration.
type DataExportConfig struct {
	// Name of the S3 bucket where the Cost and Usage Report is stored.
	BucketName string `json:"bucket_name"`
	// AWS region of the S3 bucket.
	BucketRegion string `json:"bucket_region"`
	// Name of the Cost and Usage Report.
	ReportName string `json:"report_name"`
	// S3 prefix where the Cost and Usage Report is stored.
	ReportPrefix string `json:"report_prefix"`
	// Type of the Cost and Usage Report. Currently only `CUR2.0` is supported.
	ReportType string `json:"report_type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDataExportConfig instantiates a new DataExportConfig object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDataExportConfig(bucketName string, bucketRegion string, reportName string, reportPrefix string, reportType string) *DataExportConfig {
	this := DataExportConfig{}
	this.BucketName = bucketName
	this.BucketRegion = bucketRegion
	this.ReportName = reportName
	this.ReportPrefix = reportPrefix
	this.ReportType = reportType
	return &this
}

// NewDataExportConfigWithDefaults instantiates a new DataExportConfig object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDataExportConfigWithDefaults() *DataExportConfig {
	this := DataExportConfig{}
	return &this
}

// GetBucketName returns the BucketName field value.
func (o *DataExportConfig) GetBucketName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BucketName
}

// GetBucketNameOk returns a tuple with the BucketName field value
// and a boolean to check if the value has been set.
func (o *DataExportConfig) GetBucketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketName, true
}

// SetBucketName sets field value.
func (o *DataExportConfig) SetBucketName(v string) {
	o.BucketName = v
}

// GetBucketRegion returns the BucketRegion field value.
func (o *DataExportConfig) GetBucketRegion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.BucketRegion
}

// GetBucketRegionOk returns a tuple with the BucketRegion field value
// and a boolean to check if the value has been set.
func (o *DataExportConfig) GetBucketRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketRegion, true
}

// SetBucketRegion sets field value.
func (o *DataExportConfig) SetBucketRegion(v string) {
	o.BucketRegion = v
}

// GetReportName returns the ReportName field value.
func (o *DataExportConfig) GetReportName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReportName
}

// GetReportNameOk returns a tuple with the ReportName field value
// and a boolean to check if the value has been set.
func (o *DataExportConfig) GetReportNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportName, true
}

// SetReportName sets field value.
func (o *DataExportConfig) SetReportName(v string) {
	o.ReportName = v
}

// GetReportPrefix returns the ReportPrefix field value.
func (o *DataExportConfig) GetReportPrefix() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReportPrefix
}

// GetReportPrefixOk returns a tuple with the ReportPrefix field value
// and a boolean to check if the value has been set.
func (o *DataExportConfig) GetReportPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportPrefix, true
}

// SetReportPrefix sets field value.
func (o *DataExportConfig) SetReportPrefix(v string) {
	o.ReportPrefix = v
}

// GetReportType returns the ReportType field value.
func (o *DataExportConfig) GetReportType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.ReportType
}

// GetReportTypeOk returns a tuple with the ReportType field value
// and a boolean to check if the value has been set.
func (o *DataExportConfig) GetReportTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportType, true
}

// SetReportType sets field value.
func (o *DataExportConfig) SetReportType(v string) {
	o.ReportType = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DataExportConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["bucket_name"] = o.BucketName
	toSerialize["bucket_region"] = o.BucketRegion
	toSerialize["report_name"] = o.ReportName
	toSerialize["report_prefix"] = o.ReportPrefix
	toSerialize["report_type"] = o.ReportType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DataExportConfig) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		BucketName   *string `json:"bucket_name"`
		BucketRegion *string `json:"bucket_region"`
		ReportName   *string `json:"report_name"`
		ReportPrefix *string `json:"report_prefix"`
		ReportType   *string `json:"report_type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.BucketName == nil {
		return fmt.Errorf("required field bucket_name missing")
	}
	if all.BucketRegion == nil {
		return fmt.Errorf("required field bucket_region missing")
	}
	if all.ReportName == nil {
		return fmt.Errorf("required field report_name missing")
	}
	if all.ReportPrefix == nil {
		return fmt.Errorf("required field report_prefix missing")
	}
	if all.ReportType == nil {
		return fmt.Errorf("required field report_type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"bucket_name", "bucket_region", "report_name", "report_prefix", "report_type"})
	} else {
		return err
	}
	o.BucketName = *all.BucketName
	o.BucketRegion = *all.BucketRegion
	o.ReportName = *all.ReportName
	o.ReportPrefix = *all.ReportPrefix
	o.ReportType = *all.ReportType

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
