// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"
	"time"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// LLMObsDatasetVersionDataAttributes Attributes of an LLM Observability dataset version.
type LLMObsDatasetVersionDataAttributes struct {
	// Unique identifier of the dataset this version belongs to.
	DatasetId string `json:"dataset_id"`
	// Timestamp when this dataset version was last referenced. Null if the version has never been used.
	LastUsed datadog.NullableTime `json:"last_used"`
	// Sequential version number for this dataset version.
	VersionNumber int32 `json:"version_number"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewLLMObsDatasetVersionDataAttributes instantiates a new LLMObsDatasetVersionDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewLLMObsDatasetVersionDataAttributes(datasetId string, lastUsed datadog.NullableTime, versionNumber int32) *LLMObsDatasetVersionDataAttributes {
	this := LLMObsDatasetVersionDataAttributes{}
	this.DatasetId = datasetId
	this.LastUsed = lastUsed
	this.VersionNumber = versionNumber
	return &this
}

// NewLLMObsDatasetVersionDataAttributesWithDefaults instantiates a new LLMObsDatasetVersionDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewLLMObsDatasetVersionDataAttributesWithDefaults() *LLMObsDatasetVersionDataAttributes {
	this := LLMObsDatasetVersionDataAttributes{}
	return &this
}

// GetDatasetId returns the DatasetId field value.
func (o *LLMObsDatasetVersionDataAttributes) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetVersionDataAttributes) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value.
func (o *LLMObsDatasetVersionDataAttributes) SetDatasetId(v string) {
	o.DatasetId = v
}

// GetLastUsed returns the LastUsed field value.
// If the value is explicit nil, the zero value for time.Time will be returned.
func (o *LLMObsDatasetVersionDataAttributes) GetLastUsed() time.Time {
	if o == nil || o.LastUsed.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsed.Get()
}

// GetLastUsedOk returns a tuple with the LastUsed field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *LLMObsDatasetVersionDataAttributes) GetLastUsedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsed.Get(), o.LastUsed.IsSet()
}

// SetLastUsed sets field value.
func (o *LLMObsDatasetVersionDataAttributes) SetLastUsed(v time.Time) {
	o.LastUsed.Set(&v)
}

// GetVersionNumber returns the VersionNumber field value.
func (o *LLMObsDatasetVersionDataAttributes) GetVersionNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.VersionNumber
}

// GetVersionNumberOk returns a tuple with the VersionNumber field value
// and a boolean to check if the value has been set.
func (o *LLMObsDatasetVersionDataAttributes) GetVersionNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionNumber, true
}

// SetVersionNumber sets field value.
func (o *LLMObsDatasetVersionDataAttributes) SetVersionNumber(v int32) {
	o.VersionNumber = v
}

// MarshalJSON serializes the struct using spec logic.
func (o LLMObsDatasetVersionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["dataset_id"] = o.DatasetId
	toSerialize["last_used"] = o.LastUsed.Get()
	toSerialize["version_number"] = o.VersionNumber

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *LLMObsDatasetVersionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DatasetId     *string              `json:"dataset_id"`
		LastUsed      datadog.NullableTime `json:"last_used"`
		VersionNumber *int32               `json:"version_number"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.DatasetId == nil {
		return fmt.Errorf("required field dataset_id missing")
	}
	if !all.LastUsed.IsSet() {
		return fmt.Errorf("required field last_used missing")
	}
	if all.VersionNumber == nil {
		return fmt.Errorf("required field version_number missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"dataset_id", "last_used", "version_number"})
	} else {
		return err
	}
	o.DatasetId = *all.DatasetId
	o.LastUsed = all.LastUsed
	o.VersionNumber = *all.VersionNumber

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
