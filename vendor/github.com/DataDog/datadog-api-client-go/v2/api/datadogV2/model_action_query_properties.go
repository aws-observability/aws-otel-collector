// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ActionQueryProperties The properties of the action query.
type ActionQueryProperties struct {
	// Whether to run this query. If specified, the query will only run if this condition evaluates to `true` in JavaScript and all other conditions are also met.
	Condition *ActionQueryCondition `json:"condition,omitempty"`
	// The minimum time in milliseconds that must pass before the query can be triggered again. This is useful for preventing accidental double-clicks from triggering the query multiple times.
	DebounceInMs *ActionQueryDebounceInMs `json:"debounceInMs,omitempty"`
	// The mocked outputs of the action query. This is useful for testing the app without actually running the action.
	MockedOutputs *ActionQueryMockedOutputs `json:"mockedOutputs,omitempty"`
	// Determines when this query is executed. If set to `false`, the query will run when the app loads and whenever any query arguments change. If set to `true`, the query will only run when manually triggered from elsewhere in the app.
	OnlyTriggerManually *ActionQueryOnlyTriggerManually `json:"onlyTriggerManually,omitempty"`
	// The post-query transformation function, which is a JavaScript function that changes the query's `.outputs` property after the query's execution.
	Outputs *string `json:"outputs,omitempty"`
	// If specified, the app will poll the query at the specified interval in milliseconds. The minimum polling interval is 15 seconds. The query will only poll when the app's browser tab is active.
	PollingIntervalInMs *ActionQueryPollingIntervalInMs `json:"pollingIntervalInMs,omitempty"`
	// Whether to prompt the user to confirm this query before it runs.
	RequiresConfirmation *ActionQueryRequiresConfirmation `json:"requiresConfirmation,omitempty"`
	// Whether to display a toast to the user when the query returns an error.
	ShowToastOnError *ActionQueryShowToastOnError `json:"showToastOnError,omitempty"`
	// The definition of the action query.
	Spec ActionQuerySpec `json:"spec"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewActionQueryProperties instantiates a new ActionQueryProperties object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewActionQueryProperties(spec ActionQuerySpec) *ActionQueryProperties {
	this := ActionQueryProperties{}
	this.Spec = spec
	return &this
}

// NewActionQueryPropertiesWithDefaults instantiates a new ActionQueryProperties object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewActionQueryPropertiesWithDefaults() *ActionQueryProperties {
	this := ActionQueryProperties{}
	return &this
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *ActionQueryProperties) GetCondition() ActionQueryCondition {
	if o == nil || o.Condition == nil {
		var ret ActionQueryCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionQueryProperties) GetConditionOk() (*ActionQueryCondition, bool) {
	if o == nil || o.Condition == nil {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *ActionQueryProperties) HasCondition() bool {
	return o != nil && o.Condition != nil
}

// SetCondition gets a reference to the given ActionQueryCondition and assigns it to the Condition field.
func (o *ActionQueryProperties) SetCondition(v ActionQueryCondition) {
	o.Condition = &v
}

// GetDebounceInMs returns the DebounceInMs field value if set, zero value otherwise.
func (o *ActionQueryProperties) GetDebounceInMs() ActionQueryDebounceInMs {
	if o == nil || o.DebounceInMs == nil {
		var ret ActionQueryDebounceInMs
		return ret
	}
	return *o.DebounceInMs
}

// GetDebounceInMsOk returns a tuple with the DebounceInMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionQueryProperties) GetDebounceInMsOk() (*ActionQueryDebounceInMs, bool) {
	if o == nil || o.DebounceInMs == nil {
		return nil, false
	}
	return o.DebounceInMs, true
}

// HasDebounceInMs returns a boolean if a field has been set.
func (o *ActionQueryProperties) HasDebounceInMs() bool {
	return o != nil && o.DebounceInMs != nil
}

// SetDebounceInMs gets a reference to the given ActionQueryDebounceInMs and assigns it to the DebounceInMs field.
func (o *ActionQueryProperties) SetDebounceInMs(v ActionQueryDebounceInMs) {
	o.DebounceInMs = &v
}

// GetMockedOutputs returns the MockedOutputs field value if set, zero value otherwise.
func (o *ActionQueryProperties) GetMockedOutputs() ActionQueryMockedOutputs {
	if o == nil || o.MockedOutputs == nil {
		var ret ActionQueryMockedOutputs
		return ret
	}
	return *o.MockedOutputs
}

// GetMockedOutputsOk returns a tuple with the MockedOutputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionQueryProperties) GetMockedOutputsOk() (*ActionQueryMockedOutputs, bool) {
	if o == nil || o.MockedOutputs == nil {
		return nil, false
	}
	return o.MockedOutputs, true
}

// HasMockedOutputs returns a boolean if a field has been set.
func (o *ActionQueryProperties) HasMockedOutputs() bool {
	return o != nil && o.MockedOutputs != nil
}

// SetMockedOutputs gets a reference to the given ActionQueryMockedOutputs and assigns it to the MockedOutputs field.
func (o *ActionQueryProperties) SetMockedOutputs(v ActionQueryMockedOutputs) {
	o.MockedOutputs = &v
}

// GetOnlyTriggerManually returns the OnlyTriggerManually field value if set, zero value otherwise.
func (o *ActionQueryProperties) GetOnlyTriggerManually() ActionQueryOnlyTriggerManually {
	if o == nil || o.OnlyTriggerManually == nil {
		var ret ActionQueryOnlyTriggerManually
		return ret
	}
	return *o.OnlyTriggerManually
}

// GetOnlyTriggerManuallyOk returns a tuple with the OnlyTriggerManually field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionQueryProperties) GetOnlyTriggerManuallyOk() (*ActionQueryOnlyTriggerManually, bool) {
	if o == nil || o.OnlyTriggerManually == nil {
		return nil, false
	}
	return o.OnlyTriggerManually, true
}

// HasOnlyTriggerManually returns a boolean if a field has been set.
func (o *ActionQueryProperties) HasOnlyTriggerManually() bool {
	return o != nil && o.OnlyTriggerManually != nil
}

// SetOnlyTriggerManually gets a reference to the given ActionQueryOnlyTriggerManually and assigns it to the OnlyTriggerManually field.
func (o *ActionQueryProperties) SetOnlyTriggerManually(v ActionQueryOnlyTriggerManually) {
	o.OnlyTriggerManually = &v
}

// GetOutputs returns the Outputs field value if set, zero value otherwise.
func (o *ActionQueryProperties) GetOutputs() string {
	if o == nil || o.Outputs == nil {
		var ret string
		return ret
	}
	return *o.Outputs
}

// GetOutputsOk returns a tuple with the Outputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionQueryProperties) GetOutputsOk() (*string, bool) {
	if o == nil || o.Outputs == nil {
		return nil, false
	}
	return o.Outputs, true
}

// HasOutputs returns a boolean if a field has been set.
func (o *ActionQueryProperties) HasOutputs() bool {
	return o != nil && o.Outputs != nil
}

// SetOutputs gets a reference to the given string and assigns it to the Outputs field.
func (o *ActionQueryProperties) SetOutputs(v string) {
	o.Outputs = &v
}

// GetPollingIntervalInMs returns the PollingIntervalInMs field value if set, zero value otherwise.
func (o *ActionQueryProperties) GetPollingIntervalInMs() ActionQueryPollingIntervalInMs {
	if o == nil || o.PollingIntervalInMs == nil {
		var ret ActionQueryPollingIntervalInMs
		return ret
	}
	return *o.PollingIntervalInMs
}

// GetPollingIntervalInMsOk returns a tuple with the PollingIntervalInMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionQueryProperties) GetPollingIntervalInMsOk() (*ActionQueryPollingIntervalInMs, bool) {
	if o == nil || o.PollingIntervalInMs == nil {
		return nil, false
	}
	return o.PollingIntervalInMs, true
}

// HasPollingIntervalInMs returns a boolean if a field has been set.
func (o *ActionQueryProperties) HasPollingIntervalInMs() bool {
	return o != nil && o.PollingIntervalInMs != nil
}

// SetPollingIntervalInMs gets a reference to the given ActionQueryPollingIntervalInMs and assigns it to the PollingIntervalInMs field.
func (o *ActionQueryProperties) SetPollingIntervalInMs(v ActionQueryPollingIntervalInMs) {
	o.PollingIntervalInMs = &v
}

// GetRequiresConfirmation returns the RequiresConfirmation field value if set, zero value otherwise.
func (o *ActionQueryProperties) GetRequiresConfirmation() ActionQueryRequiresConfirmation {
	if o == nil || o.RequiresConfirmation == nil {
		var ret ActionQueryRequiresConfirmation
		return ret
	}
	return *o.RequiresConfirmation
}

// GetRequiresConfirmationOk returns a tuple with the RequiresConfirmation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionQueryProperties) GetRequiresConfirmationOk() (*ActionQueryRequiresConfirmation, bool) {
	if o == nil || o.RequiresConfirmation == nil {
		return nil, false
	}
	return o.RequiresConfirmation, true
}

// HasRequiresConfirmation returns a boolean if a field has been set.
func (o *ActionQueryProperties) HasRequiresConfirmation() bool {
	return o != nil && o.RequiresConfirmation != nil
}

// SetRequiresConfirmation gets a reference to the given ActionQueryRequiresConfirmation and assigns it to the RequiresConfirmation field.
func (o *ActionQueryProperties) SetRequiresConfirmation(v ActionQueryRequiresConfirmation) {
	o.RequiresConfirmation = &v
}

// GetShowToastOnError returns the ShowToastOnError field value if set, zero value otherwise.
func (o *ActionQueryProperties) GetShowToastOnError() ActionQueryShowToastOnError {
	if o == nil || o.ShowToastOnError == nil {
		var ret ActionQueryShowToastOnError
		return ret
	}
	return *o.ShowToastOnError
}

// GetShowToastOnErrorOk returns a tuple with the ShowToastOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActionQueryProperties) GetShowToastOnErrorOk() (*ActionQueryShowToastOnError, bool) {
	if o == nil || o.ShowToastOnError == nil {
		return nil, false
	}
	return o.ShowToastOnError, true
}

// HasShowToastOnError returns a boolean if a field has been set.
func (o *ActionQueryProperties) HasShowToastOnError() bool {
	return o != nil && o.ShowToastOnError != nil
}

// SetShowToastOnError gets a reference to the given ActionQueryShowToastOnError and assigns it to the ShowToastOnError field.
func (o *ActionQueryProperties) SetShowToastOnError(v ActionQueryShowToastOnError) {
	o.ShowToastOnError = &v
}

// GetSpec returns the Spec field value.
func (o *ActionQueryProperties) GetSpec() ActionQuerySpec {
	if o == nil {
		var ret ActionQuerySpec
		return ret
	}
	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *ActionQueryProperties) GetSpecOk() (*ActionQuerySpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value.
func (o *ActionQueryProperties) SetSpec(v ActionQuerySpec) {
	o.Spec = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ActionQueryProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Condition != nil {
		toSerialize["condition"] = o.Condition
	}
	if o.DebounceInMs != nil {
		toSerialize["debounceInMs"] = o.DebounceInMs
	}
	if o.MockedOutputs != nil {
		toSerialize["mockedOutputs"] = o.MockedOutputs
	}
	if o.OnlyTriggerManually != nil {
		toSerialize["onlyTriggerManually"] = o.OnlyTriggerManually
	}
	if o.Outputs != nil {
		toSerialize["outputs"] = o.Outputs
	}
	if o.PollingIntervalInMs != nil {
		toSerialize["pollingIntervalInMs"] = o.PollingIntervalInMs
	}
	if o.RequiresConfirmation != nil {
		toSerialize["requiresConfirmation"] = o.RequiresConfirmation
	}
	if o.ShowToastOnError != nil {
		toSerialize["showToastOnError"] = o.ShowToastOnError
	}
	toSerialize["spec"] = o.Spec

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ActionQueryProperties) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Condition            *ActionQueryCondition            `json:"condition,omitempty"`
		DebounceInMs         *ActionQueryDebounceInMs         `json:"debounceInMs,omitempty"`
		MockedOutputs        *ActionQueryMockedOutputs        `json:"mockedOutputs,omitempty"`
		OnlyTriggerManually  *ActionQueryOnlyTriggerManually  `json:"onlyTriggerManually,omitempty"`
		Outputs              *string                          `json:"outputs,omitempty"`
		PollingIntervalInMs  *ActionQueryPollingIntervalInMs  `json:"pollingIntervalInMs,omitempty"`
		RequiresConfirmation *ActionQueryRequiresConfirmation `json:"requiresConfirmation,omitempty"`
		ShowToastOnError     *ActionQueryShowToastOnError     `json:"showToastOnError,omitempty"`
		Spec                 *ActionQuerySpec                 `json:"spec"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Spec == nil {
		return fmt.Errorf("required field spec missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"condition", "debounceInMs", "mockedOutputs", "onlyTriggerManually", "outputs", "pollingIntervalInMs", "requiresConfirmation", "showToastOnError", "spec"})
	} else {
		return err
	}
	o.Condition = all.Condition
	o.DebounceInMs = all.DebounceInMs
	o.MockedOutputs = all.MockedOutputs
	o.OnlyTriggerManually = all.OnlyTriggerManually
	o.Outputs = all.Outputs
	o.PollingIntervalInMs = all.PollingIntervalInMs
	o.RequiresConfirmation = all.RequiresConfirmation
	o.ShowToastOnError = all.ShowToastOnError
	o.Spec = *all.Spec

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
