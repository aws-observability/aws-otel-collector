// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ContainerImageVulnerabilities Vulnerability counts associated with the Container Image.
type ContainerImageVulnerabilities struct {
	// ID of the Container Image.
	AssetId *string `json:"asset_id,omitempty"`
	// Number of vulnerabilities with CVSS Critical severity.
	Critical *int64 `json:"critical,omitempty"`
	// Number of vulnerabilities with CVSS High severity.
	High *int64 `json:"high,omitempty"`
	// Number of vulnerabilities with CVSS Low severity.
	Low *int64 `json:"low,omitempty"`
	// Number of vulnerabilities with CVSS Medium severity.
	Medium *int64 `json:"medium,omitempty"`
	// Number of vulnerabilities with CVSS None severity.
	None *int64 `json:"none,omitempty"`
	// Number of vulnerabilities with an unknown CVSS severity.
	Unknown *int64 `json:"unknown,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewContainerImageVulnerabilities instantiates a new ContainerImageVulnerabilities object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewContainerImageVulnerabilities() *ContainerImageVulnerabilities {
	this := ContainerImageVulnerabilities{}
	return &this
}

// NewContainerImageVulnerabilitiesWithDefaults instantiates a new ContainerImageVulnerabilities object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewContainerImageVulnerabilitiesWithDefaults() *ContainerImageVulnerabilities {
	this := ContainerImageVulnerabilities{}
	return &this
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *ContainerImageVulnerabilities) GetAssetId() string {
	if o == nil || o.AssetId == nil {
		var ret string
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageVulnerabilities) GetAssetIdOk() (*string, bool) {
	if o == nil || o.AssetId == nil {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *ContainerImageVulnerabilities) HasAssetId() bool {
	return o != nil && o.AssetId != nil
}

// SetAssetId gets a reference to the given string and assigns it to the AssetId field.
func (o *ContainerImageVulnerabilities) SetAssetId(v string) {
	o.AssetId = &v
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *ContainerImageVulnerabilities) GetCritical() int64 {
	if o == nil || o.Critical == nil {
		var ret int64
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageVulnerabilities) GetCriticalOk() (*int64, bool) {
	if o == nil || o.Critical == nil {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *ContainerImageVulnerabilities) HasCritical() bool {
	return o != nil && o.Critical != nil
}

// SetCritical gets a reference to the given int64 and assigns it to the Critical field.
func (o *ContainerImageVulnerabilities) SetCritical(v int64) {
	o.Critical = &v
}

// GetHigh returns the High field value if set, zero value otherwise.
func (o *ContainerImageVulnerabilities) GetHigh() int64 {
	if o == nil || o.High == nil {
		var ret int64
		return ret
	}
	return *o.High
}

// GetHighOk returns a tuple with the High field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageVulnerabilities) GetHighOk() (*int64, bool) {
	if o == nil || o.High == nil {
		return nil, false
	}
	return o.High, true
}

// HasHigh returns a boolean if a field has been set.
func (o *ContainerImageVulnerabilities) HasHigh() bool {
	return o != nil && o.High != nil
}

// SetHigh gets a reference to the given int64 and assigns it to the High field.
func (o *ContainerImageVulnerabilities) SetHigh(v int64) {
	o.High = &v
}

// GetLow returns the Low field value if set, zero value otherwise.
func (o *ContainerImageVulnerabilities) GetLow() int64 {
	if o == nil || o.Low == nil {
		var ret int64
		return ret
	}
	return *o.Low
}

// GetLowOk returns a tuple with the Low field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageVulnerabilities) GetLowOk() (*int64, bool) {
	if o == nil || o.Low == nil {
		return nil, false
	}
	return o.Low, true
}

// HasLow returns a boolean if a field has been set.
func (o *ContainerImageVulnerabilities) HasLow() bool {
	return o != nil && o.Low != nil
}

// SetLow gets a reference to the given int64 and assigns it to the Low field.
func (o *ContainerImageVulnerabilities) SetLow(v int64) {
	o.Low = &v
}

// GetMedium returns the Medium field value if set, zero value otherwise.
func (o *ContainerImageVulnerabilities) GetMedium() int64 {
	if o == nil || o.Medium == nil {
		var ret int64
		return ret
	}
	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageVulnerabilities) GetMediumOk() (*int64, bool) {
	if o == nil || o.Medium == nil {
		return nil, false
	}
	return o.Medium, true
}

// HasMedium returns a boolean if a field has been set.
func (o *ContainerImageVulnerabilities) HasMedium() bool {
	return o != nil && o.Medium != nil
}

// SetMedium gets a reference to the given int64 and assigns it to the Medium field.
func (o *ContainerImageVulnerabilities) SetMedium(v int64) {
	o.Medium = &v
}

// GetNone returns the None field value if set, zero value otherwise.
func (o *ContainerImageVulnerabilities) GetNone() int64 {
	if o == nil || o.None == nil {
		var ret int64
		return ret
	}
	return *o.None
}

// GetNoneOk returns a tuple with the None field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageVulnerabilities) GetNoneOk() (*int64, bool) {
	if o == nil || o.None == nil {
		return nil, false
	}
	return o.None, true
}

// HasNone returns a boolean if a field has been set.
func (o *ContainerImageVulnerabilities) HasNone() bool {
	return o != nil && o.None != nil
}

// SetNone gets a reference to the given int64 and assigns it to the None field.
func (o *ContainerImageVulnerabilities) SetNone(v int64) {
	o.None = &v
}

// GetUnknown returns the Unknown field value if set, zero value otherwise.
func (o *ContainerImageVulnerabilities) GetUnknown() int64 {
	if o == nil || o.Unknown == nil {
		var ret int64
		return ret
	}
	return *o.Unknown
}

// GetUnknownOk returns a tuple with the Unknown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerImageVulnerabilities) GetUnknownOk() (*int64, bool) {
	if o == nil || o.Unknown == nil {
		return nil, false
	}
	return o.Unknown, true
}

// HasUnknown returns a boolean if a field has been set.
func (o *ContainerImageVulnerabilities) HasUnknown() bool {
	return o != nil && o.Unknown != nil
}

// SetUnknown gets a reference to the given int64 and assigns it to the Unknown field.
func (o *ContainerImageVulnerabilities) SetUnknown(v int64) {
	o.Unknown = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ContainerImageVulnerabilities) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AssetId != nil {
		toSerialize["asset_id"] = o.AssetId
	}
	if o.Critical != nil {
		toSerialize["critical"] = o.Critical
	}
	if o.High != nil {
		toSerialize["high"] = o.High
	}
	if o.Low != nil {
		toSerialize["low"] = o.Low
	}
	if o.Medium != nil {
		toSerialize["medium"] = o.Medium
	}
	if o.None != nil {
		toSerialize["none"] = o.None
	}
	if o.Unknown != nil {
		toSerialize["unknown"] = o.Unknown
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ContainerImageVulnerabilities) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AssetId  *string `json:"asset_id,omitempty"`
		Critical *int64  `json:"critical,omitempty"`
		High     *int64  `json:"high,omitempty"`
		Low      *int64  `json:"low,omitempty"`
		Medium   *int64  `json:"medium,omitempty"`
		None     *int64  `json:"none,omitempty"`
		Unknown  *int64  `json:"unknown,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"asset_id", "critical", "high", "low", "medium", "none", "unknown"})
	} else {
		return err
	}
	o.AssetId = all.AssetId
	o.Critical = all.Critical
	o.High = all.High
	o.Low = all.Low
	o.Medium = all.Medium
	o.None = all.None
	o.Unknown = all.Unknown

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
