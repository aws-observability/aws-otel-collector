// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// AggregatedHighScriptEval Aggregated high script evaluation detection grouped by source.
type AggregatedHighScriptEval struct {
	// Average script evaluation duration in nanoseconds.
	AvgDuration int64 `json:"avg_duration"`
	// Average forced style/layout duration in nanoseconds.
	AvgForcedStyleLayout int64 `json:"avg_forced_style_layout"`
	// Unique fingerprint identifying this detection group.
	Fingerprint string `json:"fingerprint"`
	// Impact score combining view frequency and duration severity.
	ImpactScore float64 `json:"impact_score"`
	// Total number of detection instances across sampled views.
	InstanceCount int32 `json:"instance_count"`
	// Type of invoker that triggered the script evaluation.
	InvokerType string `json:"invoker_type"`
	// Category of the script source.
	SourceCategory datadog.NullableString `json:"source_category"`
	// Name of the function that triggered the high script evaluation.
	SourceFunctionName string `json:"source_function_name"`
	// URL of the script that triggered the high script evaluation.
	SourceUrl datadog.NullableString `json:"source_url"`
	// Number of sampled views where this detection occurred.
	ViewOccurrences int32 `json:"view_occurrences"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewAggregatedHighScriptEval instantiates a new AggregatedHighScriptEval object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewAggregatedHighScriptEval(avgDuration int64, avgForcedStyleLayout int64, fingerprint string, impactScore float64, instanceCount int32, invokerType string, sourceCategory datadog.NullableString, sourceFunctionName string, sourceUrl datadog.NullableString, viewOccurrences int32) *AggregatedHighScriptEval {
	this := AggregatedHighScriptEval{}
	this.AvgDuration = avgDuration
	this.AvgForcedStyleLayout = avgForcedStyleLayout
	this.Fingerprint = fingerprint
	this.ImpactScore = impactScore
	this.InstanceCount = instanceCount
	this.InvokerType = invokerType
	this.SourceCategory = sourceCategory
	this.SourceFunctionName = sourceFunctionName
	this.SourceUrl = sourceUrl
	this.ViewOccurrences = viewOccurrences
	return &this
}

// NewAggregatedHighScriptEvalWithDefaults instantiates a new AggregatedHighScriptEval object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewAggregatedHighScriptEvalWithDefaults() *AggregatedHighScriptEval {
	this := AggregatedHighScriptEval{}
	return &this
}

// GetAvgDuration returns the AvgDuration field value.
func (o *AggregatedHighScriptEval) GetAvgDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvgDuration
}

// GetAvgDurationOk returns a tuple with the AvgDuration field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighScriptEval) GetAvgDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgDuration, true
}

// SetAvgDuration sets field value.
func (o *AggregatedHighScriptEval) SetAvgDuration(v int64) {
	o.AvgDuration = v
}

// GetAvgForcedStyleLayout returns the AvgForcedStyleLayout field value.
func (o *AggregatedHighScriptEval) GetAvgForcedStyleLayout() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return o.AvgForcedStyleLayout
}

// GetAvgForcedStyleLayoutOk returns a tuple with the AvgForcedStyleLayout field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighScriptEval) GetAvgForcedStyleLayoutOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgForcedStyleLayout, true
}

// SetAvgForcedStyleLayout sets field value.
func (o *AggregatedHighScriptEval) SetAvgForcedStyleLayout(v int64) {
	o.AvgForcedStyleLayout = v
}

// GetFingerprint returns the Fingerprint field value.
func (o *AggregatedHighScriptEval) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighScriptEval) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value.
func (o *AggregatedHighScriptEval) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetImpactScore returns the ImpactScore field value.
func (o *AggregatedHighScriptEval) GetImpactScore() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return o.ImpactScore
}

// GetImpactScoreOk returns a tuple with the ImpactScore field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighScriptEval) GetImpactScoreOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImpactScore, true
}

// SetImpactScore sets field value.
func (o *AggregatedHighScriptEval) SetImpactScore(v float64) {
	o.ImpactScore = v
}

// GetInstanceCount returns the InstanceCount field value.
func (o *AggregatedHighScriptEval) GetInstanceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.InstanceCount
}

// GetInstanceCountOk returns a tuple with the InstanceCount field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighScriptEval) GetInstanceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstanceCount, true
}

// SetInstanceCount sets field value.
func (o *AggregatedHighScriptEval) SetInstanceCount(v int32) {
	o.InstanceCount = v
}

// GetInvokerType returns the InvokerType field value.
func (o *AggregatedHighScriptEval) GetInvokerType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.InvokerType
}

// GetInvokerTypeOk returns a tuple with the InvokerType field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighScriptEval) GetInvokerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvokerType, true
}

// SetInvokerType sets field value.
func (o *AggregatedHighScriptEval) SetInvokerType(v string) {
	o.InvokerType = v
}

// GetSourceCategory returns the SourceCategory field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *AggregatedHighScriptEval) GetSourceCategory() string {
	if o == nil || o.SourceCategory.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceCategory.Get()
}

// GetSourceCategoryOk returns a tuple with the SourceCategory field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AggregatedHighScriptEval) GetSourceCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceCategory.Get(), o.SourceCategory.IsSet()
}

// SetSourceCategory sets field value.
func (o *AggregatedHighScriptEval) SetSourceCategory(v string) {
	o.SourceCategory.Set(&v)
}

// GetSourceFunctionName returns the SourceFunctionName field value.
func (o *AggregatedHighScriptEval) GetSourceFunctionName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.SourceFunctionName
}

// GetSourceFunctionNameOk returns a tuple with the SourceFunctionName field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighScriptEval) GetSourceFunctionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceFunctionName, true
}

// SetSourceFunctionName sets field value.
func (o *AggregatedHighScriptEval) SetSourceFunctionName(v string) {
	o.SourceFunctionName = v
}

// GetSourceUrl returns the SourceUrl field value.
// If the value is explicit nil, the zero value for string will be returned.
func (o *AggregatedHighScriptEval) GetSourceUrl() string {
	if o == nil || o.SourceUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceUrl.Get()
}

// GetSourceUrlOk returns a tuple with the SourceUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *AggregatedHighScriptEval) GetSourceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceUrl.Get(), o.SourceUrl.IsSet()
}

// SetSourceUrl sets field value.
func (o *AggregatedHighScriptEval) SetSourceUrl(v string) {
	o.SourceUrl.Set(&v)
}

// GetViewOccurrences returns the ViewOccurrences field value.
func (o *AggregatedHighScriptEval) GetViewOccurrences() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return o.ViewOccurrences
}

// GetViewOccurrencesOk returns a tuple with the ViewOccurrences field value
// and a boolean to check if the value has been set.
func (o *AggregatedHighScriptEval) GetViewOccurrencesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViewOccurrences, true
}

// SetViewOccurrences sets field value.
func (o *AggregatedHighScriptEval) SetViewOccurrences(v int32) {
	o.ViewOccurrences = v
}

// MarshalJSON serializes the struct using spec logic.
func (o AggregatedHighScriptEval) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["avg_duration"] = o.AvgDuration
	toSerialize["avg_forced_style_layout"] = o.AvgForcedStyleLayout
	toSerialize["fingerprint"] = o.Fingerprint
	toSerialize["impact_score"] = o.ImpactScore
	toSerialize["instance_count"] = o.InstanceCount
	toSerialize["invoker_type"] = o.InvokerType
	toSerialize["source_category"] = o.SourceCategory.Get()
	toSerialize["source_function_name"] = o.SourceFunctionName
	toSerialize["source_url"] = o.SourceUrl.Get()
	toSerialize["view_occurrences"] = o.ViewOccurrences

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *AggregatedHighScriptEval) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AvgDuration          *int64                 `json:"avg_duration"`
		AvgForcedStyleLayout *int64                 `json:"avg_forced_style_layout"`
		Fingerprint          *string                `json:"fingerprint"`
		ImpactScore          *float64               `json:"impact_score"`
		InstanceCount        *int32                 `json:"instance_count"`
		InvokerType          *string                `json:"invoker_type"`
		SourceCategory       datadog.NullableString `json:"source_category"`
		SourceFunctionName   *string                `json:"source_function_name"`
		SourceUrl            datadog.NullableString `json:"source_url"`
		ViewOccurrences      *int32                 `json:"view_occurrences"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AvgDuration == nil {
		return fmt.Errorf("required field avg_duration missing")
	}
	if all.AvgForcedStyleLayout == nil {
		return fmt.Errorf("required field avg_forced_style_layout missing")
	}
	if all.Fingerprint == nil {
		return fmt.Errorf("required field fingerprint missing")
	}
	if all.ImpactScore == nil {
		return fmt.Errorf("required field impact_score missing")
	}
	if all.InstanceCount == nil {
		return fmt.Errorf("required field instance_count missing")
	}
	if all.InvokerType == nil {
		return fmt.Errorf("required field invoker_type missing")
	}
	if !all.SourceCategory.IsSet() {
		return fmt.Errorf("required field source_category missing")
	}
	if all.SourceFunctionName == nil {
		return fmt.Errorf("required field source_function_name missing")
	}
	if !all.SourceUrl.IsSet() {
		return fmt.Errorf("required field source_url missing")
	}
	if all.ViewOccurrences == nil {
		return fmt.Errorf("required field view_occurrences missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.UnmarshalUseNumber(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"avg_duration", "avg_forced_style_layout", "fingerprint", "impact_score", "instance_count", "invoker_type", "source_category", "source_function_name", "source_url", "view_occurrences"})
	} else {
		return err
	}
	o.AvgDuration = *all.AvgDuration
	o.AvgForcedStyleLayout = *all.AvgForcedStyleLayout
	o.Fingerprint = *all.Fingerprint
	o.ImpactScore = *all.ImpactScore
	o.InstanceCount = *all.InstanceCount
	o.InvokerType = *all.InvokerType
	o.SourceCategory = all.SourceCategory
	o.SourceFunctionName = *all.SourceFunctionName
	o.SourceUrl = all.SourceUrl
	o.ViewOccurrences = *all.ViewOccurrences

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
