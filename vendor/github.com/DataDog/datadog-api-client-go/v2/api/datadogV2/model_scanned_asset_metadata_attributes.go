// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScannedAssetMetadataAttributes The attributes of a scanned asset metadata.
type ScannedAssetMetadataAttributes struct {
	// The asset of a scanned asset metadata.
	Asset ScannedAssetMetadataAsset `json:"asset"`
	// The timestamp when the scan of the asset was performed for the first time.
	FirstSuccessTimestamp string `json:"first_success_timestamp"`
	// Metadata for the last successful scan of an asset.
	LastSuccess ScannedAssetMetadataLastSuccess `json:"last_success"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScannedAssetMetadataAttributes instantiates a new ScannedAssetMetadataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScannedAssetMetadataAttributes(asset ScannedAssetMetadataAsset, firstSuccessTimestamp string, lastSuccess ScannedAssetMetadataLastSuccess) *ScannedAssetMetadataAttributes {
	this := ScannedAssetMetadataAttributes{}
	this.Asset = asset
	this.FirstSuccessTimestamp = firstSuccessTimestamp
	this.LastSuccess = lastSuccess
	return &this
}

// NewScannedAssetMetadataAttributesWithDefaults instantiates a new ScannedAssetMetadataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScannedAssetMetadataAttributesWithDefaults() *ScannedAssetMetadataAttributes {
	this := ScannedAssetMetadataAttributes{}
	return &this
}

// GetAsset returns the Asset field value.
func (o *ScannedAssetMetadataAttributes) GetAsset() ScannedAssetMetadataAsset {
	if o == nil {
		var ret ScannedAssetMetadataAsset
		return ret
	}
	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *ScannedAssetMetadataAttributes) GetAssetOk() (*ScannedAssetMetadataAsset, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value.
func (o *ScannedAssetMetadataAttributes) SetAsset(v ScannedAssetMetadataAsset) {
	o.Asset = v
}

// GetFirstSuccessTimestamp returns the FirstSuccessTimestamp field value.
func (o *ScannedAssetMetadataAttributes) GetFirstSuccessTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.FirstSuccessTimestamp
}

// GetFirstSuccessTimestampOk returns a tuple with the FirstSuccessTimestamp field value
// and a boolean to check if the value has been set.
func (o *ScannedAssetMetadataAttributes) GetFirstSuccessTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstSuccessTimestamp, true
}

// SetFirstSuccessTimestamp sets field value.
func (o *ScannedAssetMetadataAttributes) SetFirstSuccessTimestamp(v string) {
	o.FirstSuccessTimestamp = v
}

// GetLastSuccess returns the LastSuccess field value.
func (o *ScannedAssetMetadataAttributes) GetLastSuccess() ScannedAssetMetadataLastSuccess {
	if o == nil {
		var ret ScannedAssetMetadataLastSuccess
		return ret
	}
	return o.LastSuccess
}

// GetLastSuccessOk returns a tuple with the LastSuccess field value
// and a boolean to check if the value has been set.
func (o *ScannedAssetMetadataAttributes) GetLastSuccessOk() (*ScannedAssetMetadataLastSuccess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastSuccess, true
}

// SetLastSuccess sets field value.
func (o *ScannedAssetMetadataAttributes) SetLastSuccess(v ScannedAssetMetadataLastSuccess) {
	o.LastSuccess = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScannedAssetMetadataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["asset"] = o.Asset
	toSerialize["first_success_timestamp"] = o.FirstSuccessTimestamp
	toSerialize["last_success"] = o.LastSuccess

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScannedAssetMetadataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Asset                 *ScannedAssetMetadataAsset       `json:"asset"`
		FirstSuccessTimestamp *string                          `json:"first_success_timestamp"`
		LastSuccess           *ScannedAssetMetadataLastSuccess `json:"last_success"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Asset == nil {
		return fmt.Errorf("required field asset missing")
	}
	if all.FirstSuccessTimestamp == nil {
		return fmt.Errorf("required field first_success_timestamp missing")
	}
	if all.LastSuccess == nil {
		return fmt.Errorf("required field last_success missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"asset", "first_success_timestamp", "last_success"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Asset.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Asset = *all.Asset
	o.FirstSuccessTimestamp = *all.FirstSuccessTimestamp
	if all.LastSuccess.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.LastSuccess = *all.LastSuccess

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
