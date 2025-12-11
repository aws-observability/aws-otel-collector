// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ScaRequestDataAttributesCommit
type ScaRequestDataAttributesCommit struct {
	//
	AuthorDate *string `json:"author_date,omitempty"`
	//
	AuthorEmail *string `json:"author_email,omitempty"`
	//
	AuthorName *string `json:"author_name,omitempty"`
	//
	Branch *string `json:"branch,omitempty"`
	//
	CommitterEmail *string `json:"committer_email,omitempty"`
	//
	CommitterName *string `json:"committer_name,omitempty"`
	//
	Sha *string `json:"sha,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewScaRequestDataAttributesCommit instantiates a new ScaRequestDataAttributesCommit object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewScaRequestDataAttributesCommit() *ScaRequestDataAttributesCommit {
	this := ScaRequestDataAttributesCommit{}
	return &this
}

// NewScaRequestDataAttributesCommitWithDefaults instantiates a new ScaRequestDataAttributesCommit object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewScaRequestDataAttributesCommitWithDefaults() *ScaRequestDataAttributesCommit {
	this := ScaRequestDataAttributesCommit{}
	return &this
}

// GetAuthorDate returns the AuthorDate field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesCommit) GetAuthorDate() string {
	if o == nil || o.AuthorDate == nil {
		var ret string
		return ret
	}
	return *o.AuthorDate
}

// GetAuthorDateOk returns a tuple with the AuthorDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesCommit) GetAuthorDateOk() (*string, bool) {
	if o == nil || o.AuthorDate == nil {
		return nil, false
	}
	return o.AuthorDate, true
}

// HasAuthorDate returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesCommit) HasAuthorDate() bool {
	return o != nil && o.AuthorDate != nil
}

// SetAuthorDate gets a reference to the given string and assigns it to the AuthorDate field.
func (o *ScaRequestDataAttributesCommit) SetAuthorDate(v string) {
	o.AuthorDate = &v
}

// GetAuthorEmail returns the AuthorEmail field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesCommit) GetAuthorEmail() string {
	if o == nil || o.AuthorEmail == nil {
		var ret string
		return ret
	}
	return *o.AuthorEmail
}

// GetAuthorEmailOk returns a tuple with the AuthorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesCommit) GetAuthorEmailOk() (*string, bool) {
	if o == nil || o.AuthorEmail == nil {
		return nil, false
	}
	return o.AuthorEmail, true
}

// HasAuthorEmail returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesCommit) HasAuthorEmail() bool {
	return o != nil && o.AuthorEmail != nil
}

// SetAuthorEmail gets a reference to the given string and assigns it to the AuthorEmail field.
func (o *ScaRequestDataAttributesCommit) SetAuthorEmail(v string) {
	o.AuthorEmail = &v
}

// GetAuthorName returns the AuthorName field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesCommit) GetAuthorName() string {
	if o == nil || o.AuthorName == nil {
		var ret string
		return ret
	}
	return *o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesCommit) GetAuthorNameOk() (*string, bool) {
	if o == nil || o.AuthorName == nil {
		return nil, false
	}
	return o.AuthorName, true
}

// HasAuthorName returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesCommit) HasAuthorName() bool {
	return o != nil && o.AuthorName != nil
}

// SetAuthorName gets a reference to the given string and assigns it to the AuthorName field.
func (o *ScaRequestDataAttributesCommit) SetAuthorName(v string) {
	o.AuthorName = &v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesCommit) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesCommit) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesCommit) HasBranch() bool {
	return o != nil && o.Branch != nil
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *ScaRequestDataAttributesCommit) SetBranch(v string) {
	o.Branch = &v
}

// GetCommitterEmail returns the CommitterEmail field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesCommit) GetCommitterEmail() string {
	if o == nil || o.CommitterEmail == nil {
		var ret string
		return ret
	}
	return *o.CommitterEmail
}

// GetCommitterEmailOk returns a tuple with the CommitterEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesCommit) GetCommitterEmailOk() (*string, bool) {
	if o == nil || o.CommitterEmail == nil {
		return nil, false
	}
	return o.CommitterEmail, true
}

// HasCommitterEmail returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesCommit) HasCommitterEmail() bool {
	return o != nil && o.CommitterEmail != nil
}

// SetCommitterEmail gets a reference to the given string and assigns it to the CommitterEmail field.
func (o *ScaRequestDataAttributesCommit) SetCommitterEmail(v string) {
	o.CommitterEmail = &v
}

// GetCommitterName returns the CommitterName field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesCommit) GetCommitterName() string {
	if o == nil || o.CommitterName == nil {
		var ret string
		return ret
	}
	return *o.CommitterName
}

// GetCommitterNameOk returns a tuple with the CommitterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesCommit) GetCommitterNameOk() (*string, bool) {
	if o == nil || o.CommitterName == nil {
		return nil, false
	}
	return o.CommitterName, true
}

// HasCommitterName returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesCommit) HasCommitterName() bool {
	return o != nil && o.CommitterName != nil
}

// SetCommitterName gets a reference to the given string and assigns it to the CommitterName field.
func (o *ScaRequestDataAttributesCommit) SetCommitterName(v string) {
	o.CommitterName = &v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *ScaRequestDataAttributesCommit) GetSha() string {
	if o == nil || o.Sha == nil {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScaRequestDataAttributesCommit) GetShaOk() (*string, bool) {
	if o == nil || o.Sha == nil {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *ScaRequestDataAttributesCommit) HasSha() bool {
	return o != nil && o.Sha != nil
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *ScaRequestDataAttributesCommit) SetSha(v string) {
	o.Sha = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o ScaRequestDataAttributesCommit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.AuthorDate != nil {
		toSerialize["author_date"] = o.AuthorDate
	}
	if o.AuthorEmail != nil {
		toSerialize["author_email"] = o.AuthorEmail
	}
	if o.AuthorName != nil {
		toSerialize["author_name"] = o.AuthorName
	}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	if o.CommitterEmail != nil {
		toSerialize["committer_email"] = o.CommitterEmail
	}
	if o.CommitterName != nil {
		toSerialize["committer_name"] = o.CommitterName
	}
	if o.Sha != nil {
		toSerialize["sha"] = o.Sha
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ScaRequestDataAttributesCommit) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		AuthorDate     *string `json:"author_date,omitempty"`
		AuthorEmail    *string `json:"author_email,omitempty"`
		AuthorName     *string `json:"author_name,omitempty"`
		Branch         *string `json:"branch,omitempty"`
		CommitterEmail *string `json:"committer_email,omitempty"`
		CommitterName  *string `json:"committer_name,omitempty"`
		Sha            *string `json:"sha,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"author_date", "author_email", "author_name", "branch", "committer_email", "committer_name", "sha"})
	} else {
		return err
	}
	o.AuthorDate = all.AuthorDate
	o.AuthorEmail = all.AuthorEmail
	o.AuthorName = all.AuthorName
	o.Branch = all.Branch
	o.CommitterEmail = all.CommitterEmail
	o.CommitterName = all.CommitterName
	o.Sha = all.Sha

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
