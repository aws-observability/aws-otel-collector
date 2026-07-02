// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SignalsProblemsSampleMetadata Metadata about the sampling quality for a signals and problems query.
type SignalsProblemsSampleMetadata struct {
	// Number of view instances that failed to process.
	Failed int32 `json:"failed"`
	// Number of view instances requested for sampling.
	Requested int32 `json:"requested"`
	// List of RUM view IDs that were sampled.
	SampledViewIds []string `json:"sampled_view_ids"`
	// Number of view instances successfully processed.
	Succeeded int32 `json:"succeeded"`
	// Ratio of successfully processed views to requested views.
	SuccessRate float64 `json:"success_rate"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSignalsProblemsSampleMetadata instantiates a new SignalsProblemsSampleMetadata object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSignalsProblemsSampleMetadata(failed int32, requested int32, sampledViewIds []string, succeeded int32, successRate float64) *SignalsProblemsSampleMetadata {
	this := SignalsProblemsSampleMetadata{}
	this.Failed = failed
	this.Requested = requested
	this.SampledViewIds = sampledViewIds
	this.Succeeded = succeeded
	this.SuccessRate = successRate
	return &this
}

// NewSignalsProblemsSampleMetadataWithDefaults instantiates a new SignalsProblemsSampleMetadata object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSignalsProblemsSampleMetadataWithDefaults() *SignalsProblemsSampleMetadata {
	this := SignalsProblemsSampleMetadata{}
	return &this
}

// GetFailed returns the Failed field value.
func (o *SignalsProblemsSampleMetadata) GetFailed() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Failed
}

// GetFailedOk returns a tuple with the Failed field value
// and a boolean to check if the value has been set.
func (o *SignalsProblemsSampleMetadata) GetFailedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Failed, true
}

// SetFailed sets field value.
func (o *SignalsProblemsSampleMetadata) SetFailed(v int32) {
	o.Failed = v
}

// GetRequested returns the Requested field value.
func (o *SignalsProblemsSampleMetadata) GetRequested() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value
// and a boolean to check if the value has been set.
func (o *SignalsProblemsSampleMetadata) GetRequestedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requested, true
}

// SetRequested sets field value.
func (o *SignalsProblemsSampleMetadata) SetRequested(v int32) {
	o.Requested = v
}

// GetSampledViewIds returns the SampledViewIds field value.
func (o *SignalsProblemsSampleMetadata) GetSampledViewIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SampledViewIds
}

// GetSampledViewIdsOk returns a tuple with the SampledViewIds field value
// and a boolean to check if the value has been set.
func (o *SignalsProblemsSampleMetadata) GetSampledViewIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SampledViewIds, true
}

// SetSampledViewIds sets field value.
func (o *SignalsProblemsSampleMetadata) SetSampledViewIds(v []string) {
	o.SampledViewIds = v
}

// GetSucceeded returns the Succeeded field value.
func (o *SignalsProblemsSampleMetadata) GetSucceeded() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.Succeeded
}

// GetSucceededOk returns a tuple with the Succeeded field value
// and a boolean to check if the value has been set.
func (o *SignalsProblemsSampleMetadata) GetSucceededOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Succeeded, true
}

// SetSucceeded sets field value.
func (o *SignalsProblemsSampleMetadata) SetSucceeded(v int32) {
	o.Succeeded = v
}

// GetSuccessRate returns the SuccessRate field value.
func (o *SignalsProblemsSampleMetadata) GetSuccessRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.SuccessRate
}

// GetSuccessRateOk returns a tuple with the SuccessRate field value
// and a boolean to check if the value has been set.
func (o *SignalsProblemsSampleMetadata) GetSuccessRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccessRate, true
}

// SetSuccessRate sets field value.
func (o *SignalsProblemsSampleMetadata) SetSuccessRate(v float64) {
	o.SuccessRate = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SignalsProblemsSampleMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["failed"] = o.Failed
	toSerialize["requested"] = o.Requested
	toSerialize["sampled_view_ids"] = o.SampledViewIds
	toSerialize["succeeded"] = o.Succeeded
	toSerialize["success_rate"] = o.SuccessRate

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SignalsProblemsSampleMetadata) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Failed         *int32    `json:"failed"`
		Requested      *int32    `json:"requested"`
		SampledViewIds *[]string `json:"sampled_view_ids"`
		Succeeded      *int32    `json:"succeeded"`
		SuccessRate    *float64  `json:"success_rate"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Failed == nil {
		return fmt.Errorf("required field failed missing")
	}
	if all.Requested == nil {
		return fmt.Errorf("required field requested missing")
	}
	if all.SampledViewIds == nil {
		return fmt.Errorf("required field sampled_view_ids missing")
	}
	if all.Succeeded == nil {
		return fmt.Errorf("required field succeeded missing")
	}
	if all.SuccessRate == nil {
		return fmt.Errorf("required field success_rate missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"failed", "requested", "sampled_view_ids", "succeeded", "success_rate"})
	} else {
		return err
	}
	o.Failed = *all.Failed
	o.Requested = *all.Requested
	o.SampledViewIds = *all.SampledViewIds
	o.Succeeded = *all.Succeeded
	o.SuccessRate = *all.SuccessRate

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
