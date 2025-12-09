// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails Cloud storage access configuration for the reference table data file.
type PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails struct {
	// Amazon Web Services S3 storage access configuration.
	AwsDetail *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail `json:"aws_detail,omitempty"`
	// Azure Blob Storage access configuration.
	AzureDetail *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail `json:"azure_detail,omitempty"`
	// Google Cloud Platform storage access configuration.
	GcpDetail *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail `json:"gcp_detail,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewPatchTableRequestDataAttributesFileMetadataOneOfAccessDetails instantiates a new PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewPatchTableRequestDataAttributesFileMetadataOneOfAccessDetails() *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails {
	this := PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails{}
	return &this
}

// NewPatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsWithDefaults instantiates a new PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewPatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsWithDefaults() *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails {
	this := PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails{}
	return &this
}

// GetAwsDetail returns the AwsDetail field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) GetAwsDetail() PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail {
	if o == nil || o.AwsDetail == nil {
		var ret PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail
		return ret
	}
	return *o.AwsDetail
}

// GetAwsDetailOk returns a tuple with the AwsDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) GetAwsDetailOk() (*PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail, bool) {
	if o == nil || o.AwsDetail == nil {
		return nil, false
	}
	return o.AwsDetail, true
}

// HasAwsDetail returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) HasAwsDetail() bool {
	return o != nil && o.AwsDetail != nil
}

// SetAwsDetail gets a reference to the given PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail and assigns it to the AwsDetail field.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) SetAwsDetail(v PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail) {
	o.AwsDetail = &v
}

// GetAzureDetail returns the AzureDetail field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) GetAzureDetail() PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail {
	if o == nil || o.AzureDetail == nil {
		var ret PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail
		return ret
	}
	return *o.AzureDetail
}

// GetAzureDetailOk returns a tuple with the AzureDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) GetAzureDetailOk() (*PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail, bool) {
	if o == nil || o.AzureDetail == nil {
		return nil, false
	}
	return o.AzureDetail, true
}

// HasAzureDetail returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) HasAzureDetail() bool {
	return o != nil && o.AzureDetail != nil
}

// SetAzureDetail gets a reference to the given PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail and assigns it to the AzureDetail field.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) SetAzureDetail(v PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail) {
	o.AzureDetail = &v
}

// GetGcpDetail returns the GcpDetail field value if set, zero value otherwise.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) GetGcpDetail() PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail {
	if o == nil || o.GcpDetail == nil {
		var ret PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail
		return ret
	}
	return *o.GcpDetail
}

// GetGcpDetailOk returns a tuple with the GcpDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) GetGcpDetailOk() (*PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail, bool) {
	if o == nil || o.GcpDetail == nil {
		return nil, false
	}
	return o.GcpDetail, true
}

// HasGcpDetail returns a boolean if a field has been set.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) HasGcpDetail() bool {
	return o != nil && o.GcpDetail != nil
}

// SetGcpDetail gets a reference to the given PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail and assigns it to the GcpDetail field.
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) SetGcpDetail(v PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail) {
	o.GcpDetail = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) MarshalJSON() ([]byte, error) {
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
func (o *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetails) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AwsDetail   *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAwsDetail   `json:"aws_detail,omitempty"`
		AzureDetail *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsAzureDetail `json:"azure_detail,omitempty"`
		GcpDetail   *PatchTableRequestDataAttributesFileMetadataOneOfAccessDetailsGcpDetail   `json:"gcp_detail,omitempty"`
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
