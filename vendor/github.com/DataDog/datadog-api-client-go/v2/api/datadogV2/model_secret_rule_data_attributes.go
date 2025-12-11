// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// SecretRuleDataAttributes
type SecretRuleDataAttributes struct {
	//
	DefaultIncludedKeywords []string `json:"default_included_keywords,omitempty"`
	//
	Description *string `json:"description,omitempty"`
	//
	License *string `json:"license,omitempty"`
	//
	MatchValidation *SecretRuleDataAttributesMatchValidation `json:"match_validation,omitempty"`
	//
	Name *string `json:"name,omitempty"`
	//
	Pattern *string `json:"pattern,omitempty"`
	//
	Priority *string `json:"priority,omitempty"`
	//
	SdsId *string `json:"sds_id,omitempty"`
	//
	Validators []string `json:"validators,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewSecretRuleDataAttributes instantiates a new SecretRuleDataAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewSecretRuleDataAttributes() *SecretRuleDataAttributes {
	this := SecretRuleDataAttributes{}
	return &this
}

// NewSecretRuleDataAttributesWithDefaults instantiates a new SecretRuleDataAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewSecretRuleDataAttributesWithDefaults() *SecretRuleDataAttributes {
	this := SecretRuleDataAttributes{}
	return &this
}

// GetDefaultIncludedKeywords returns the DefaultIncludedKeywords field value if set, zero value otherwise.
func (o *SecretRuleDataAttributes) GetDefaultIncludedKeywords() []string {
	if o == nil || o.DefaultIncludedKeywords == nil {
		var ret []string
		return ret
	}
	return o.DefaultIncludedKeywords
}

// GetDefaultIncludedKeywordsOk returns a tuple with the DefaultIncludedKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributes) GetDefaultIncludedKeywordsOk() (*[]string, bool) {
	if o == nil || o.DefaultIncludedKeywords == nil {
		return nil, false
	}
	return &o.DefaultIncludedKeywords, true
}

// HasDefaultIncludedKeywords returns a boolean if a field has been set.
func (o *SecretRuleDataAttributes) HasDefaultIncludedKeywords() bool {
	return o != nil && o.DefaultIncludedKeywords != nil
}

// SetDefaultIncludedKeywords gets a reference to the given []string and assigns it to the DefaultIncludedKeywords field.
func (o *SecretRuleDataAttributes) SetDefaultIncludedKeywords(v []string) {
	o.DefaultIncludedKeywords = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecretRuleDataAttributes) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributes) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecretRuleDataAttributes) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecretRuleDataAttributes) SetDescription(v string) {
	o.Description = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *SecretRuleDataAttributes) GetLicense() string {
	if o == nil || o.License == nil {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributes) GetLicenseOk() (*string, bool) {
	if o == nil || o.License == nil {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *SecretRuleDataAttributes) HasLicense() bool {
	return o != nil && o.License != nil
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *SecretRuleDataAttributes) SetLicense(v string) {
	o.License = &v
}

// GetMatchValidation returns the MatchValidation field value if set, zero value otherwise.
func (o *SecretRuleDataAttributes) GetMatchValidation() SecretRuleDataAttributesMatchValidation {
	if o == nil || o.MatchValidation == nil {
		var ret SecretRuleDataAttributesMatchValidation
		return ret
	}
	return *o.MatchValidation
}

// GetMatchValidationOk returns a tuple with the MatchValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributes) GetMatchValidationOk() (*SecretRuleDataAttributesMatchValidation, bool) {
	if o == nil || o.MatchValidation == nil {
		return nil, false
	}
	return o.MatchValidation, true
}

// HasMatchValidation returns a boolean if a field has been set.
func (o *SecretRuleDataAttributes) HasMatchValidation() bool {
	return o != nil && o.MatchValidation != nil
}

// SetMatchValidation gets a reference to the given SecretRuleDataAttributesMatchValidation and assigns it to the MatchValidation field.
func (o *SecretRuleDataAttributes) SetMatchValidation(v SecretRuleDataAttributesMatchValidation) {
	o.MatchValidation = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecretRuleDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecretRuleDataAttributes) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecretRuleDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetPattern returns the Pattern field value if set, zero value otherwise.
func (o *SecretRuleDataAttributes) GetPattern() string {
	if o == nil || o.Pattern == nil {
		var ret string
		return ret
	}
	return *o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributes) GetPatternOk() (*string, bool) {
	if o == nil || o.Pattern == nil {
		return nil, false
	}
	return o.Pattern, true
}

// HasPattern returns a boolean if a field has been set.
func (o *SecretRuleDataAttributes) HasPattern() bool {
	return o != nil && o.Pattern != nil
}

// SetPattern gets a reference to the given string and assigns it to the Pattern field.
func (o *SecretRuleDataAttributes) SetPattern(v string) {
	o.Pattern = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *SecretRuleDataAttributes) GetPriority() string {
	if o == nil || o.Priority == nil {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributes) GetPriorityOk() (*string, bool) {
	if o == nil || o.Priority == nil {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *SecretRuleDataAttributes) HasPriority() bool {
	return o != nil && o.Priority != nil
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *SecretRuleDataAttributes) SetPriority(v string) {
	o.Priority = &v
}

// GetSdsId returns the SdsId field value if set, zero value otherwise.
func (o *SecretRuleDataAttributes) GetSdsId() string {
	if o == nil || o.SdsId == nil {
		var ret string
		return ret
	}
	return *o.SdsId
}

// GetSdsIdOk returns a tuple with the SdsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributes) GetSdsIdOk() (*string, bool) {
	if o == nil || o.SdsId == nil {
		return nil, false
	}
	return o.SdsId, true
}

// HasSdsId returns a boolean if a field has been set.
func (o *SecretRuleDataAttributes) HasSdsId() bool {
	return o != nil && o.SdsId != nil
}

// SetSdsId gets a reference to the given string and assigns it to the SdsId field.
func (o *SecretRuleDataAttributes) SetSdsId(v string) {
	o.SdsId = &v
}

// GetValidators returns the Validators field value if set, zero value otherwise.
func (o *SecretRuleDataAttributes) GetValidators() []string {
	if o == nil || o.Validators == nil {
		var ret []string
		return ret
	}
	return o.Validators
}

// GetValidatorsOk returns a tuple with the Validators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecretRuleDataAttributes) GetValidatorsOk() (*[]string, bool) {
	if o == nil || o.Validators == nil {
		return nil, false
	}
	return &o.Validators, true
}

// HasValidators returns a boolean if a field has been set.
func (o *SecretRuleDataAttributes) HasValidators() bool {
	return o != nil && o.Validators != nil
}

// SetValidators gets a reference to the given []string and assigns it to the Validators field.
func (o *SecretRuleDataAttributes) SetValidators(v []string) {
	o.Validators = v
}

// MarshalJSON serializes the struct using spec logic.
func (o SecretRuleDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.DefaultIncludedKeywords != nil {
		toSerialize["default_included_keywords"] = o.DefaultIncludedKeywords
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.License != nil {
		toSerialize["license"] = o.License
	}
	if o.MatchValidation != nil {
		toSerialize["match_validation"] = o.MatchValidation
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Pattern != nil {
		toSerialize["pattern"] = o.Pattern
	}
	if o.Priority != nil {
		toSerialize["priority"] = o.Priority
	}
	if o.SdsId != nil {
		toSerialize["sds_id"] = o.SdsId
	}
	if o.Validators != nil {
		toSerialize["validators"] = o.Validators
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *SecretRuleDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		DefaultIncludedKeywords []string                                 `json:"default_included_keywords,omitempty"`
		Description             *string                                  `json:"description,omitempty"`
		License                 *string                                  `json:"license,omitempty"`
		MatchValidation         *SecretRuleDataAttributesMatchValidation `json:"match_validation,omitempty"`
		Name                    *string                                  `json:"name,omitempty"`
		Pattern                 *string                                  `json:"pattern,omitempty"`
		Priority                *string                                  `json:"priority,omitempty"`
		SdsId                   *string                                  `json:"sds_id,omitempty"`
		Validators              []string                                 `json:"validators,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"default_included_keywords", "description", "license", "match_validation", "name", "pattern", "priority", "sds_id", "validators"})
	} else {
		return err
	}

	hasInvalidField := false
	o.DefaultIncludedKeywords = all.DefaultIncludedKeywords
	o.Description = all.Description
	o.License = all.License
	if all.MatchValidation != nil && all.MatchValidation.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MatchValidation = all.MatchValidation
	o.Name = all.Name
	o.Pattern = all.Pattern
	o.Priority = all.Priority
	o.SdsId = all.SdsId
	o.Validators = all.Validators

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
