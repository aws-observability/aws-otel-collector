// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// TableResultV2DataAttributesFileMetadataOneOfAccessDetails Cloud storage access configuration for the reference table data file.
type TableResultV2DataAttributesFileMetadataOneOfAccessDetails struct {
	// Amazon Web Services S3 storage access configuration.
	AwsDetail *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail `json:"aws_detail,omitempty"`
	// Azure Blob Storage access configuration.
	AzureDetail *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail `json:"azure_detail,omitempty"`
	// Google Cloud Platform storage access configuration.
	GcpDetail *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsGcpDetail `json:"gcp_detail,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewTableResultV2DataAttributesFileMetadataOneOfAccessDetails instantiates a new TableResultV2DataAttributesFileMetadataOneOfAccessDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewTableResultV2DataAttributesFileMetadataOneOfAccessDetails() *TableResultV2DataAttributesFileMetadataOneOfAccessDetails {
	this := TableResultV2DataAttributesFileMetadataOneOfAccessDetails{}
	return &this
}

// NewTableResultV2DataAttributesFileMetadataOneOfAccessDetailsWithDefaults instantiates a new TableResultV2DataAttributesFileMetadataOneOfAccessDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewTableResultV2DataAttributesFileMetadataOneOfAccessDetailsWithDefaults() *TableResultV2DataAttributesFileMetadataOneOfAccessDetails {
	this := TableResultV2DataAttributesFileMetadataOneOfAccessDetails{}
	return &this
}

// GetAwsDetail returns the AwsDetail field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetails) GetAwsDetail() TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail {
	if o == nil || o.AwsDetail == nil {
		var ret TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail
		return ret
	}
	return *o.AwsDetail
}

// GetAwsDetailOk returns a tuple with the AwsDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetails) GetAwsDetailOk() (*TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail, bool) {
	if o == nil || o.AwsDetail == nil {
		return nil, false
	}
	return o.AwsDetail, true
}

// HasAwsDetail returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetails) HasAwsDetail() bool {
	return o != nil && o.AwsDetail != nil
}

// SetAwsDetail gets a reference to the given TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail and assigns it to the AwsDetail field.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetails) SetAwsDetail(v TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail) {
	o.AwsDetail = &v
}

// GetAzureDetail returns the AzureDetail field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetails) GetAzureDetail() TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail {
	if o == nil || o.AzureDetail == nil {
		var ret TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail
		return ret
	}
	return *o.AzureDetail
}

// GetAzureDetailOk returns a tuple with the AzureDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetails) GetAzureDetailOk() (*TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail, bool) {
	if o == nil || o.AzureDetail == nil {
		return nil, false
	}
	return o.AzureDetail, true
}

// HasAzureDetail returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetails) HasAzureDetail() bool {
	return o != nil && o.AzureDetail != nil
}

// SetAzureDetail gets a reference to the given TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail and assigns it to the AzureDetail field.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetails) SetAzureDetail(v TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail) {
	o.AzureDetail = &v
}

// GetGcpDetail returns the GcpDetail field value if set, zero value otherwise.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetails) GetGcpDetail() TableResultV2DataAttributesFileMetadataOneOfAccessDetailsGcpDetail {
	if o == nil || o.GcpDetail == nil {
		var ret TableResultV2DataAttributesFileMetadataOneOfAccessDetailsGcpDetail
		return ret
	}
	return *o.GcpDetail
}

// GetGcpDetailOk returns a tuple with the GcpDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetails) GetGcpDetailOk() (*TableResultV2DataAttributesFileMetadataOneOfAccessDetailsGcpDetail, bool) {
	if o == nil || o.GcpDetail == nil {
		return nil, false
	}
	return o.GcpDetail, true
}

// HasGcpDetail returns a boolean if a field has been set.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetails) HasGcpDetail() bool {
	return o != nil && o.GcpDetail != nil
}

// SetGcpDetail gets a reference to the given TableResultV2DataAttributesFileMetadataOneOfAccessDetailsGcpDetail and assigns it to the GcpDetail field.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetails) SetGcpDetail(v TableResultV2DataAttributesFileMetadataOneOfAccessDetailsGcpDetail) {
	o.GcpDetail = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o TableResultV2DataAttributesFileMetadataOneOfAccessDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AwsDetail != nil {
		toSerialize["aws_detail"] = o.AwsDetail
	}
	if o.AzureDetail != nil {
		toSerialize["azure_detail"] = o.AzureDetail
	}
	if o.GcpDetail != nil {
		toSerialize["gcp_detail"] = o.GcpDetail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *TableResultV2DataAttributesFileMetadataOneOfAccessDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AwsDetail   *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAwsDetail   `json:"aws_detail,omitempty"`
		AzureDetail *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsAzureDetail `json:"azure_detail,omitempty"`
		GcpDetail   *TableResultV2DataAttributesFileMetadataOneOfAccessDetailsGcpDetail   `json:"gcp_detail,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"aws_detail", "azure_detail", "gcp_detail"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.AwsDetail != nil && all.AwsDetail.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AwsDetail = all.AwsDetail
	if all.AzureDetail != nil && all.AzureDetail.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.AzureDetail = all.AzureDetail
	if all.GcpDetail != nil && all.GcpDetail.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.GcpDetail = all.GcpDetail

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
