// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// Remediation Vulnerability remediation.
type Remediation struct {
	// Whether the vulnerability can be resolved when recompiling the package or not.
	AutoSolvable bool `json:"auto_solvable"`
	// Avoided advisories.
	AvoidedAdvisories []Advisory `json:"avoided_advisories"`
	// Remediation fixed advisories.
	FixedAdvisories []Advisory `json:"fixed_advisories"`
	// Library name remediating the vulnerability.
	LibraryName string `json:"library_name"`
	// Library version remediating the vulnerability.
	LibraryVersion string `json:"library_version"`
	// New advisories.
	NewAdvisories []Advisory `json:"new_advisories"`
	// Remaining advisories.
	RemainingAdvisories []Advisory `json:"remaining_advisories"`
	// Remediation type.
	Type string `json:"type"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRemediation instantiates a new Remediation object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRemediation(autoSolvable bool, avoidedAdvisories []Advisory, fixedAdvisories []Advisory, libraryName string, libraryVersion string, newAdvisories []Advisory, remainingAdvisories []Advisory, typeVar string) *Remediation {
	this := Remediation{}
	this.AutoSolvable = autoSolvable
	this.AvoidedAdvisories = avoidedAdvisories
	this.FixedAdvisories = fixedAdvisories
	this.LibraryName = libraryName
	this.LibraryVersion = libraryVersion
	this.NewAdvisories = newAdvisories
	this.RemainingAdvisories = remainingAdvisories
	this.Type = typeVar
	return &this
}

// NewRemediationWithDefaults instantiates a new Remediation object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRemediationWithDefaults() *Remediation {
	this := Remediation{}
	return &this
}

// GetAutoSolvable returns the AutoSolvable field value.
func (o *Remediation) GetAutoSolvable() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.AutoSolvable
}

// GetAutoSolvableOk returns a tuple with the AutoSolvable field value
// and a boolean to check if the value has been set.
func (o *Remediation) GetAutoSolvableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoSolvable, true
}

// SetAutoSolvable sets field value.
func (o *Remediation) SetAutoSolvable(v bool) {
	o.AutoSolvable = v
}

// GetAvoidedAdvisories returns the AvoidedAdvisories field value.
func (o *Remediation) GetAvoidedAdvisories() []Advisory {
	if o == nil {
		var ret []Advisory
		return ret
	}
	return o.AvoidedAdvisories
}

// GetAvoidedAdvisoriesOk returns a tuple with the AvoidedAdvisories field value
// and a boolean to check if the value has been set.
func (o *Remediation) GetAvoidedAdvisoriesOk() (*[]Advisory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvoidedAdvisories, true
}

// SetAvoidedAdvisories sets field value.
func (o *Remediation) SetAvoidedAdvisories(v []Advisory) {
	o.AvoidedAdvisories = v
}

// GetFixedAdvisories returns the FixedAdvisories field value.
func (o *Remediation) GetFixedAdvisories() []Advisory {
	if o == nil {
		var ret []Advisory
		return ret
	}
	return o.FixedAdvisories
}

// GetFixedAdvisoriesOk returns a tuple with the FixedAdvisories field value
// and a boolean to check if the value has been set.
func (o *Remediation) GetFixedAdvisoriesOk() (*[]Advisory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FixedAdvisories, true
}

// SetFixedAdvisories sets field value.
func (o *Remediation) SetFixedAdvisories(v []Advisory) {
	o.FixedAdvisories = v
}

// GetLibraryName returns the LibraryName field value.
func (o *Remediation) GetLibraryName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LibraryName
}

// GetLibraryNameOk returns a tuple with the LibraryName field value
// and a boolean to check if the value has been set.
func (o *Remediation) GetLibraryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LibraryName, true
}

// SetLibraryName sets field value.
func (o *Remediation) SetLibraryName(v string) {
	o.LibraryName = v
}

// GetLibraryVersion returns the LibraryVersion field value.
func (o *Remediation) GetLibraryVersion() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.LibraryVersion
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value
// and a boolean to check if the value has been set.
func (o *Remediation) GetLibraryVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LibraryVersion, true
}

// SetLibraryVersion sets field value.
func (o *Remediation) SetLibraryVersion(v string) {
	o.LibraryVersion = v
}

// GetNewAdvisories returns the NewAdvisories field value.
func (o *Remediation) GetNewAdvisories() []Advisory {
	if o == nil {
		var ret []Advisory
		return ret
	}
	return o.NewAdvisories
}

// GetNewAdvisoriesOk returns a tuple with the NewAdvisories field value
// and a boolean to check if the value has been set.
func (o *Remediation) GetNewAdvisoriesOk() (*[]Advisory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewAdvisories, true
}

// SetNewAdvisories sets field value.
func (o *Remediation) SetNewAdvisories(v []Advisory) {
	o.NewAdvisories = v
}

// GetRemainingAdvisories returns the RemainingAdvisories field value.
func (o *Remediation) GetRemainingAdvisories() []Advisory {
	if o == nil {
		var ret []Advisory
		return ret
	}
	return o.RemainingAdvisories
}

// GetRemainingAdvisoriesOk returns a tuple with the RemainingAdvisories field value
// and a boolean to check if the value has been set.
func (o *Remediation) GetRemainingAdvisoriesOk() (*[]Advisory, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemainingAdvisories, true
}

// SetRemainingAdvisories sets field value.
func (o *Remediation) SetRemainingAdvisories(v []Advisory) {
	o.RemainingAdvisories = v
}

// GetType returns the Type field value.
func (o *Remediation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Remediation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value.
func (o *Remediation) SetType(v string) {
	o.Type = v
}

// MarshalJSON serializes the struct using spec logic.
func (o Remediation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	toSerialize["auto_solvable"] = o.AutoSolvable
	toSerialize["avoided_advisories"] = o.AvoidedAdvisories
	toSerialize["fixed_advisories"] = o.FixedAdvisories
	toSerialize["library_name"] = o.LibraryName
	toSerialize["library_version"] = o.LibraryVersion
	toSerialize["new_advisories"] = o.NewAdvisories
	toSerialize["remaining_advisories"] = o.RemainingAdvisories
	toSerialize["type"] = o.Type

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *Remediation) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AutoSolvable        *bool       `json:"auto_solvable"`
		AvoidedAdvisories   *[]Advisory `json:"avoided_advisories"`
		FixedAdvisories     *[]Advisory `json:"fixed_advisories"`
		LibraryName         *string     `json:"library_name"`
		LibraryVersion      *string     `json:"library_version"`
		NewAdvisories       *[]Advisory `json:"new_advisories"`
		RemainingAdvisories *[]Advisory `json:"remaining_advisories"`
		Type                *string     `json:"type"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.AutoSolvable == nil {
		return fmt.Errorf("required field auto_solvable missing")
	}
	if all.AvoidedAdvisories == nil {
		return fmt.Errorf("required field avoided_advisories missing")
	}
	if all.FixedAdvisories == nil {
		return fmt.Errorf("required field fixed_advisories missing")
	}
	if all.LibraryName == nil {
		return fmt.Errorf("required field library_name missing")
	}
	if all.LibraryVersion == nil {
		return fmt.Errorf("required field library_version missing")
	}
	if all.NewAdvisories == nil {
		return fmt.Errorf("required field new_advisories missing")
	}
	if all.RemainingAdvisories == nil {
		return fmt.Errorf("required field remaining_advisories missing")
	}
	if all.Type == nil {
		return fmt.Errorf("required field type missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"auto_solvable", "avoided_advisories", "fixed_advisories", "library_name", "library_version", "new_advisories", "remaining_advisories", "type"})
	} else {
		return err
	}
	o.AutoSolvable = *all.AutoSolvable
	o.AvoidedAdvisories = *all.AvoidedAdvisories
	o.FixedAdvisories = *all.FixedAdvisories
	o.LibraryName = *all.LibraryName
	o.LibraryVersion = *all.LibraryVersion
	o.NewAdvisories = *all.NewAdvisories
	o.RemainingAdvisories = *all.RemainingAdvisories
	o.Type = *all.Type

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
